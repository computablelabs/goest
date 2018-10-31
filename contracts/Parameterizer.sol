pragma solidity 0.4.25;

import "./PLCRVoting.sol";
import "./IERC20.sol";
import "./SafeMath.sol";


contract Parameterizer {
  using SafeMath for uint;

  struct ParamProposal {
    uint appExpiry;
    uint challenge;
    uint deposit;
    bytes name; // storing the byte array and not a string
    address owner;
    uint processBy;
    uint value;
  }

  struct Challenge {
    uint rewardPool;        // (remaining) pool of tokens distributed amongst winning voters
    address challenger;     // owner of Challenge
    bool resolved;          // indication of if challenge is resolved
    uint stake;             // number of tokens at risk for either party during challenge
    uint winningTokens;     // (remaining) amount of tokens used for voting by the winning side
    mapping(address => bool) rewardsClaimed;
  }

  mapping(bytes32 => uint) private selfParams;

  // maps challenge IDs to associated challenge data
  mapping(uint => Challenge) private selfChallenges;

  // maps pollIDs to intended data change if poll passes
  mapping(bytes32 => ParamProposal) private selfProposals;

  // Global Variables
  IERC20 private selfToken;
  PLCRVoting private selfVoting;
  uint private PROCESSBY = 604800; // 7 days TODO should we make this mutable ?

  /**
  @dev constructor
  @param tokenAddr Address of the token which parameterizes this system
  @param votingAddr Address of a voting contract for the provided token
  @param minDeposit Minimum deposit for listing to be whitelisted
  @param applyStageLen Period over which applicants wait to be whitelisted
  @param dispensationPct Percentage of losing party's deposit distributed to winning party
  @param commitStageLen Length of commit period for voting
  @param revealStageLen Length of reveal period for voting
  @param voteQuorum Type of majority out of 100 necessary for vote success
  @param listReward The amount of market tokens minted for a listing if accepted
  @param conversionRate TODO
  @param conversionSlope TODO
  */
  constructor(
    address tokenAddr,
    address votingAddr,
    uint minDeposit,
    uint applyStageLen,
    uint commitStageLen,
    uint revealStageLen,
    uint dispensationPct,
    uint voteQuorum,
    uint listReward,
    uint conversionRate,
    uint conversionSlope
    ) public
  {
    selfToken = IERC20(tokenAddr);
    selfVoting = PLCRVoting(votingAddr);

    set("minDeposit", minDeposit);
    set("applyStageLen", applyStageLen);
    set("commitStageLen", commitStageLen);
    set("revealStageLen", revealStageLen);
    set("dispensationPct", dispensationPct);
    set("voteQuorum", voteQuorum);
    set("listReward", listReward);
    set("conversionRate", conversionRate);
    set("conversionSlope", conversionSlope);
  }

  /**
  @notice Determines whether a proposal passed its application stage without a challenge
  @param propHash The proposal ID for which to determine whether its application stage passed without a challenge
  */
  function canBeSet(bytes32 propHash) public view returns (bool) {
    ParamProposal memory prop = selfProposals[propHash];

    return (now > prop.appExpiry && now < prop.processBy && prop.challenge == 0);
  }

  /**
  @notice Determines whether the provided proposal ID has a challenge which can be resolved
  @param propHash The proposal ID whose challenge to inspect
  */
  function challengeCanBeResolved(bytes32 propHash) public view returns (bool) {
    ParamProposal memory prop = selfProposals[propHash];
    Challenge memory challenge = selfChallenges[prop.challenge];

    return (prop.challenge > 0 && challenge.resolved == false && selfVoting.pollEnded(prop.challenge));
  }

   /**
  @notice challenge the provided proposal ID, and put tokens at stake to do so.
  @param propHash the proposal ID to challenge
  */
  function challengeReparameterization(bytes32 propHash) external returns (uint) {
    ParamProposal memory prop = selfProposals[propHash];
    uint deposit = prop.deposit;

    require(propExists(propHash) && prop.challenge == 0, "Error:Parameterizer.challengeReparameterization - Proposal does not exist");

    //start poll
    uint id = selfVoting.startPoll(
      get("voteQuorum"),
      get("commitStageLen"),
      get("revealStageLen")
    );

    selfChallenges[id] = Challenge({
      challenger: msg.sender,
      rewardPool: SafeMath.sub(100, get("dispensationPct")).mul(deposit).div(100),
      stake: deposit,
      resolved: false,
      winningTokens: 0
    });

    selfProposals[propHash].challenge = id; // update listing to store most recent challenge

    //take tokens from challenger
    require(selfToken.transferFrom(msg.sender, this, deposit), "Error:Parameterizer.challengeReparameterization - Could not transfer funds");

    uint commitExpiry = selfVoting.getPollCommitExpiry(id);
    uint revealExpiry = selfVoting.getPollRevealExpiry(id);

    emit NewChallengeEvent(
      propHash,
      id,
      commitExpiry,
      revealExpiry,
      msg.sender
    );

    return id;
  }

  /**
  @notice Determines the number of tokens to awarded to the winning party in a challenge
  @param id The challenge ID to determine a reward for
  */
  function challengeWinnerReward(uint id) public view returns (uint) {
    if (selfVoting.getTotalNumberOfTokensForWinningOption(id) == 0) {
      // Edge case, nobody voted, give all tokens to the challenger.
      return selfChallenges[id].stake.mul(2);
    }

    return (selfChallenges[id].stake.mul(2)).sub(selfChallenges[id].rewardPool);
  }

  /**
  @notice claim the tokens owed for the msg.sender in the provided challenge
  @param id the challenge ID to claim tokens for
  @param salt the salt used to vote in the challenge being withdrawn for
  */
  function claimReward(uint id, uint salt) public {
    // ensure voter has not already claimed tokens and challenge results have been processed
    require(selfChallenges[id].rewardsClaimed[msg.sender] == false, "Error:Parameterizer.claimReward - Tokens have already been claimed");
    require(selfChallenges[id].resolved == true, "Error:Parameterizer.claimReward - Challenge results have not been processed yet");

    uint voterTokens = selfVoting.getNumPassingTokens(msg.sender, id, salt);
    uint reward = voterReward(msg.sender, id, salt);

    // subtract voter's information to preserve the participation ratios of other voters
    // compared to the remaining pool of rewards
    selfChallenges[id].winningTokens = selfChallenges[id].winningTokens.sub(voterTokens);
    selfChallenges[id].rewardPool = selfChallenges[id].rewardPool.sub(reward);

    // ensures a voter cannot claim tokens again
    selfChallenges[id].rewardsClaimed[msg.sender] = true;

    emit RewardsClaimedEvent(id, reward, msg.sender);
    require(selfToken.transfer(msg.sender, reward), "Error:Parameterizer.claimReward - Could not transfer funds");
  }

  /**
  @notice gets the parameter keyed by the provided name value from the params mapping
  @param name the key whose value is to be determined
  */
  function get(string name) public view returns (uint value) {
    return selfParams[keccak256(abi.encodePacked(name))];
  }

  /**
  @notice for the provided proposal ID, set it, resolve its challenge, or delete it depending on whether it can be set, has a challenge which can be resolved, or if its "process by" date has passed
  @param propHash the proposal ID to make a determination and state transition for
  */
  function processProposal(bytes32 propHash) public {
    ParamProposal storage prop = selfProposals[propHash];
    address propOwner = prop.owner;
    uint propDeposit = prop.deposit;

    // Before any token transfers, deleting the proposal will ensure that if reentrancy occurs the
    // prop.owner and prop.deposit will be 0, thereby preventing theft
    if (canBeSet(propHash)) {
      // There is no challenge against the proposal. The processBy date for the proposal has not
      // passed, but the proposal's appExpirty date has passed.
      set(string(prop.name), prop.value);
      emit ProposalAcceptedEvent(propHash, string(prop.name), prop.value);
      delete selfProposals[propHash];
      require(selfToken.transfer(propOwner, propDeposit), "Error:Parameterizer.processProposal - Could not transfer funds");
    } else if (challengeCanBeResolved(propHash)) {
      // There is a challenge against the proposal.
      resolveChallenge(propHash);
    } else if (now > prop.processBy) {
      // There is no challenge against the proposal, but the processBy date has passed.
      emit ProposalExpiredEvent(propHash);
      delete selfProposals[propHash];
      require(selfToken.transfer(propOwner, propDeposit), "Error:Parameterizer.processProposal - Could not transfer funds");
    } else {
      revert("There is no challenge against the proposal, and neither the appExpiry date nor the processBy date has passed");
    }

    assert(get("dispensationPct") <= 100);

    // verify that future proposal appExpiry and processBy times will not overflow
    now.add(get("applyStageLen"))
      .add(get("commitStageLen"))
      .add(get("revealStageLen"))
      .add(PROCESSBY);

    delete selfProposals[propHash];
  }

  /**
  @notice Determines whether a proposal exists for the provided proposal ID
  @param propHash The proposal ID whose existance is to be determined
  */
  function propExists(bytes32 propHash) public view returns (bool) {
    return selfProposals[propHash].processBy > 0;
  }

  /**
  @notice propose a reparamaterization of the key name's value to value.
  @param name the name of the proposed param to be set
  @param value the proposed value to set the param to be set
  */
  function proposeReparameterization(string name, uint value) external returns (bytes32) {
    uint deposit = get("minDeposit");
    bytes32 propHash = keccak256(abi.encodePacked(name, value));

    if (keccak256(abi.encodePacked(name)) == keccak256(abi.encodePacked("dispensationPct")) ||
       keccak256(abi.encodePacked(name)) == keccak256(abi.encodePacked("pDispensationPct"))) {
      require(value <= 100, "Error:Parameterizer.proposeReparameterization - `value` must be less than or equal to 100");
    }

    require(!propExists(propHash), "Error:Parameterizer.proposeReparameterization - Duplicate proposals not allowed"); // Forbid duplicate proposals
    require(get(name) != value, "Error:Parameterizer.proposeReparameterization - Arguments cannot be identical"); // Forbid NOOP reparameterizations

    // attach name and value to poll ID
    selfProposals[propHash] = ParamProposal({
      appExpiry: now.add(get("applyStageLen")),
      challenge: 0,
      deposit: deposit,
      name: bytes(name), // let's not store a bunch of strings
      owner: msg.sender,
      processBy: now.add(get("applyStageLen"))
        .add(get("commitStageLen"))
        .add(get("revealStageLen"))
        .add(PROCESSBY),
      value: value
    });

    require(selfToken.transferFrom(msg.sender, this, deposit), "Error:Parameterizer.proposeReparameterization - Could not transfer funds"); // escrow tokens (deposit amt)

    emit ReparameterizationProposalEvent(
      name,
      value,
      propHash,
      deposit,
      selfProposals[propHash].appExpiry,
      msg.sender
    );

    return propHash;
  }

  /**
  @dev resolves a challenge for the provided prop ID. It must be checked in advance whether the prop ID has a challenge on it
  @param propHash the proposal ID whose challenge is to be resolved.
  */
  function resolveChallenge(bytes32 propHash) private {
    ParamProposal memory prop = selfProposals[propHash];
    Challenge storage challenge = selfChallenges[prop.challenge];

    // winner gets back their full staked deposit, and dispensationPct*loser's stake
    uint reward = challengeWinnerReward(prop.challenge);

    challenge.winningTokens = selfVoting.getTotalNumberOfTokensForWinningOption(prop.challenge);
    challenge.resolved = true;

    if (selfVoting.isPassed(prop.challenge)) { // The challenge failed
      if (prop.processBy > now) {
        set(string(prop.name), prop.value);
      }
      emit ChallengeFailedEvent(
        propHash,
        prop.challenge,
        challenge.rewardPool,
        challenge.winningTokens
      );

      require(selfToken.transfer(prop.owner, reward), "Error:Parameterizer.resolveChallenge - Could not transfer funds");
    } else { // The challenge succeeded or nobody voted
      emit ChallengeSucceededEvent(
        propHash,
        prop.challenge,
        challenge.rewardPool,
        challenge.winningTokens
      );

      require(selfToken.transfer(selfChallenges[prop.challenge].challenger, reward), "Error:Parameterizer.resolveChallenge - Could not transfer funds");
    }
  }

  /**
  @dev Getter for Challenge rewardsClaimed mappings
  @param id The challenge ID to query
  @param voter The voter whose claim status to query for the provided challenge ID
  */
  function rewardsClaimed(uint id, address voter) external view returns (bool) {
    return selfChallenges[id].rewardsClaimed[voter];
  }

  /**
  @dev sets the param keted by the provided name to the provided value
  @param name the name of the param to be set
  @param value the value to set the param to be set
  */
  function set(string name, uint value) private {
    selfParams[keccak256(abi.encodePacked(name))] = value;
  }

  /**
  @dev Calculates the provided voter's token reward for the given poll.
  @param voter The address of the voter whose reward balance is to be returned
  @param id The ID of the challenge the voter's reward is being calculated for
  @param salt The salt of the voter's commit hash in the given poll
  @return The uint indicating the voter's reward
  */
  function voterReward(address voter, uint id, uint salt)
  public view returns (uint)
  {
    uint winningTokens = selfChallenges[id].winningTokens;
    uint rewardPool = selfChallenges[id].rewardPool;
    uint voterTokens = selfVoting.getNumPassingTokens(voter, id, salt);
    return voterTokens.mul(rewardPool).div(winningTokens);
  }

  event ReparameterizationProposalEvent(string name, uint value, bytes32 propHash, uint deposit, uint appExpiry, address indexed proposer);
  event NewChallengeEvent(bytes32 indexed propHash, uint id, uint commitExpiry, uint revealExpiry, address indexed challenger);
  event ProposalAcceptedEvent(bytes32 indexed propHash, string name, uint value);
  event ProposalExpiredEvent(bytes32 indexed propHash);
  event ChallengeSucceededEvent(bytes32 indexed propHash, uint indexed id, uint rewardPool, uint totalTokens);
  event ChallengeFailedEvent(bytes32 indexed propHash, uint indexed id, uint rewardPool, uint totalTokens);
  event RewardsClaimedEvent(uint indexed id, uint reward, address indexed voter);
}
