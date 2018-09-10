pragma solidity 0.4.24;

import "./PLCRVoting.sol";
import "./ERC20.sol";
import "./SafeMath.sol";


contract Parameterizer {

  // ------
  // EVENTS
  // ------

  event ReparameterizationProposal(string name, uint value, bytes32 propID, uint deposit, uint appEndDate, address indexed proposer);
  event NewChallenge(bytes32 indexed propID, uint challengeID, uint commitEndDate, uint revealEndDate, address indexed challenger);
  event ProposalAccepted(bytes32 indexed propID, string name, uint value);
  event ProposalExpired(bytes32 indexed propID);
  event ChallengeSucceeded(bytes32 indexed propID, uint indexed challengeID, uint rewardPool, uint totalTokens);
  event ChallengeFailed(bytes32 indexed propID, uint indexed challengeID, uint rewardPool, uint totalTokens);
  event RewardClaimed(uint indexed challengeID, uint reward, address indexed voter);


  // ------
  // DATA STRUCTURES
  // ------

  using SafeMath for uint;

  struct ParamProposal {
    uint appExpiry;
    uint challengeID;
    uint deposit;
    string name;
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
    mapping(address => bool) tokenClaims;
  }

  // ------
  // STATE
  // ------

  mapping(bytes32 => uint) public params;

  // maps challengeIDs to associated challenge data
  mapping(uint => Challenge) public challenges;

  // maps pollIDs to intended data change if poll passes
  mapping(bytes32 => ParamProposal) public proposals;

  // Global Variables
  ERC20 public token;
  PLCRVoting public voting;
  uint public PROCESSBY = 604800; // 7 days

  // ------------
  // CONSTRUCTOR
  // ------------

  /**
  @dev constructor
  @param tokenAddr        address of the token which parameterizes this system
  @param votingAddr       address of a voting contract for the provided token
  @param minDeposit       minimum deposit for listing to be whitelisted
  @param applyStageLen    period over which applicants wait to be whitelisted
  @param dispensationPct  percentage of losing party's deposit distributed to winning party
  @param commitStageLen  length of commit period for voting
  @param pCommitStageLen length of commit period for voting in parameterizer
  @param revealStageLen  length of reveal period for voting
  @param voteQuorum       type of majority out of 100 necessary for vote success
  */
  constructor(
    address tokenAddr,
    address votingAddr,
    uint minDeposit,
    uint applyStageLen,
    uint commitStageLen,
    uint revealStageLen,
    uint dispensationPct,
    uint voteQuorum
    ) public
  {
    token = ERC20(tokenAddr);
    voting = PLCRVoting(votingAddr);

    set("minDeposit", minDeposit);
    set("applyStageLen", applyStageLen);
    set("commitStageLen", commitStageLen);
    set("revealStageLen", revealStageLen);
    set("dispensationPct", dispensationPct);
    set("voteQuorum", voteQuorum);
  }

  // -----------------------
  // TOKEN HOLDER INTERFACE
  // -----------------------

  /**
  @notice propose a reparamaterization of the key name's value to value.
  @param name the name of the proposed param to be set
  @param value the proposed value to set the param to be set
  */
  function proposeReparameterization(string name, uint value) public returns (bytes32) {
    uint deposit = get("minDeposit");
    bytes32 propID = keccak256(abi.encodePacked(name, value));

    if (keccak256(abi.encodePacked(name)) == keccak256(abi.encodePacked("dispensationPct")) ||
       keccak256(abi.encodePacked(name)) == keccak256(abi.encodePacked("pDispensationPct"))) {
      require(value <= 100, "Error:Parameterizer.proposeReparameterization - `value` must be less than or equal to 100");
    }

    require(!propExists(propID), "Error:Parameterizer.proposeReparameterization - Duplicate proposals not allowed"); // Forbid duplicate proposals
    require(get(name) != value, "Error:Parameterizer.proposeReparameterization - Arguments cannot be identical"); // Forbid NOOP reparameterizations

    // attach name and value to pollID
    proposals[propID] = ParamProposal({
      appExpiry: now.add(get("applyStageLen")),
      challengeID: 0,
      deposit: deposit,
      name: name,
      owner: msg.sender,
      processBy: now.add(get("applyStageLen"))
        .add(get("commitStageLen"))
        .add(get("revealStageLen"))
        .add(PROCESSBY),
      value: value
    });

    require(token.transferFrom(msg.sender, this, deposit), "Error:Parameterizer.proposeReparameterization - Could not transfer funds"); // escrow tokens (deposit amt)

    emit ReparameterizationProposal(
      name,
      value,
      propID,
      deposit,
      proposals[propID].appExpiry,
      msg.sender
    );

    return propID;
  }

  /**
  @notice challenge the provided proposal ID, and put tokens at stake to do so.
  @param propID the proposal ID to challenge
  */
  function challengeReparameterization(bytes32 propID) public returns (uint challengeID) {
    ParamProposal memory prop = proposals[propID];
    uint deposit = prop.deposit;

    require(propExists(propID) && prop.challengeID == 0, "Error:Parameterizer.challengeReparameterization - Proposal does not exist");

    //start poll
    uint pollID = voting.startPoll(
      get("voteQuorum"),
      get("commitStageLen"),
      get("revealStageLen")
    );

    challenges[pollID] = Challenge({
      challenger: msg.sender,
      rewardPool: SafeMath.sub(100, get("dispensationPct")).mul(deposit).div(100),
      stake: deposit,
      resolved: false,
      winningTokens: 0
    });

    proposals[propID].challengeID = pollID; // update listing to store most recent challenge

    //take tokens from challenger
    require(token.transferFrom(msg.sender, this, deposit), "Error:Parameterizer.challengeReparameterization - Could not transfer funds");

    uint commitEndDate;
    uint revealEndDate;
    (commitEndDate, revealEndDate,) = voting.pollMap(pollID);

    emit NewChallenge(
      propID,
      pollID,
      commitEndDate,
      revealEndDate,
      msg.sender
    );

    return pollID;
  }

  /**
  @notice for the provided proposal ID, set it, resolve its challenge, or delete it depending on whether it can be set, has a challenge which can be resolved, or if its "process by" date has passed
  @param propID the proposal ID to make a determination and state transition for
  */
  function processProposal(bytes32 propID) public {
    ParamProposal storage prop = proposals[propID];
    address propOwner = prop.owner;
    uint propDeposit = prop.deposit;

    // Before any token transfers, deleting the proposal will ensure that if reentrancy occurs the
    // prop.owner and prop.deposit will be 0, thereby preventing theft
    if (canBeSet(propID)) {
      // There is no challenge against the proposal. The processBy date for the proposal has not
      // passed, but the proposal's appExpirty date has passed.
      set(prop.name, prop.value);
      emit ProposalAccepted(propID, prop.name, prop.value);
      delete proposals[propID];
      require(token.transfer(propOwner, propDeposit), "Error:Parameterizer.processProposal - Could not transfer funds");
    } else if (challengeCanBeResolved(propID)) {
      // There is a challenge against the proposal.
      resolveChallenge(propID);
    } else if (now > prop.processBy) {
      // There is no challenge against the proposal, but the processBy date has passed.
      emit ProposalExpired(propID);
      delete proposals[propID];
      require(token.transfer(propOwner, propDeposit), "Error:Parameterizer.processProposal - Could not transfer funds");
    } else {
      revert("There is no challenge against the proposal, and neither the appExpiry date nor the processBy date has passed");
    }

    assert(get("dispensationPct") <= 100);

    // verify that future proposal appExpiry and processBy times will not overflow
    now.add(get("applyStageLen"))
      .add(get("commitStageLen"))
      .add(get("revealStageLen"))
      .add(PROCESSBY);

    delete proposals[propID];
  }

  /**
  @notice claim the tokens owed for the msg.sender in the provided challenge
  @param challengeID the challenge ID to claim tokens for
  @param salt the salt used to vote in the challenge being withdrawn for
  */
  function claimReward(uint challengeID, uint salt) public {
    // ensure voter has not already claimed tokens and challenge results have been processed
    require(challenges[challengeID].tokenClaims[msg.sender] == false, "Error:Parameterizer.claimReward - Tokens have already been claimed");
    require(challenges[challengeID].resolved == true, "Error:Parameterizer.claimReward - Challenge results have not been processed yet");

    uint voterTokens = voting.getNumPassingTokens(msg.sender, challengeID, salt);
    uint reward = voterReward(msg.sender, challengeID, salt);

    // subtract voter's information to preserve the participation ratios of other voters
    // compared to the remaining pool of rewards
    challenges[challengeID].winningTokens = challenges[challengeID].winningTokens.sub(voterTokens);
    challenges[challengeID].rewardPool = challenges[challengeID].rewardPool.sub(reward);

    // ensures a voter cannot claim tokens again
    challenges[challengeID].tokenClaims[msg.sender] = true;

    emit RewardClaimed(challengeID, reward, msg.sender);
    require(token.transfer(msg.sender, reward), "Error:Parameterizer.claimReward - Could not transfer funds");
  }

  // --------
  // GETTERS
  // --------

  /**
  @dev                Calculates the provided voter's token reward for the given poll.
  @param voter       The address of the voter whose reward balance is to be returned
  @param challengeID The ID of the challenge the voter's reward is being calculated for
  @param salt        The salt of the voter's commit hash in the given poll
  @return             The uint indicating the voter's reward
  */
  function voterReward(address voter, uint challengeID, uint salt)
  public view returns (uint)
  {
    uint winningTokens = challenges[challengeID].winningTokens;
    uint rewardPool = challenges[challengeID].rewardPool;
    uint voterTokens = voting.getNumPassingTokens(voter, challengeID, salt);
    return voterTokens.mul(rewardPool).div(winningTokens);
  }

  /**
  @notice Determines whether a proposal passed its application stage without a challenge
  @param propID The proposal ID for which to determine whether its application stage passed without a challenge
  */
  function canBeSet(bytes32 propID) public view returns (bool) {
    ParamProposal memory prop = proposals[propID];

    return (now > prop.appExpiry && now < prop.processBy && prop.challengeID == 0);
  }

  /**
  @notice Determines whether a proposal exists for the provided proposal ID
  @param propID The proposal ID whose existance is to be determined
  */
  function propExists(bytes32 propID) public view returns (bool) {
    return proposals[propID].processBy > 0;
  }

  /**
  @notice Determines whether the provided proposal ID has a challenge which can be resolved
  @param propID The proposal ID whose challenge to inspect
  */
  function challengeCanBeResolved(bytes32 propID) public view returns (bool) {
    ParamProposal memory prop = proposals[propID];
    Challenge memory challenge = challenges[prop.challengeID];

    return (prop.challengeID > 0 && challenge.resolved == false && voting.pollEnded(prop.challengeID));
  }

  /**
  @notice Determines the number of tokens to awarded to the winning party in a challenge
  @param challengeID The challengeID to determine a reward for
  */
  function challengeWinnerReward(uint challengeID) public view returns (uint) {
    if (voting.getTotalNumberOfTokensForWinningOption(challengeID) == 0) {
      // Edge case, nobody voted, give all tokens to the challenger.
      return challenges[challengeID].stake.mul(2);
    }

    return (challenges[challengeID].stake.mul(2)).sub(challenges[challengeID].rewardPool);
  }

  /**
  @notice gets the parameter keyed by the provided name value from the params mapping
  @param name the key whose value is to be determined
  */
  function get(string name) public view returns (uint value) {
    return params[keccak256(abi.encodePacked(name))];
  }

  /**
  @dev                Getter for Challenge tokenClaims mappings
  @param challengeID The challengeID to query
  @param voter       The voter whose claim status to query for the provided challengeID
  */
  function tokenClaims(uint challengeID, address voter) public view returns (bool) {
    return challenges[challengeID].tokenClaims[voter];
  }

  // ----------------
  // PRIVATE FUNCTIONS
  // ----------------

  /**
  @dev resolves a challenge for the provided propID. It must be checked in advance whether the propID has a challenge on it
  @param propID the proposal ID whose challenge is to be resolved.
  */
  function resolveChallenge(bytes32 propID) private {
    ParamProposal memory prop = proposals[propID];
    Challenge storage challenge = challenges[prop.challengeID];

    // winner gets back their full staked deposit, and dispensationPct*loser's stake
    uint reward = challengeWinnerReward(prop.challengeID);

    challenge.winningTokens = voting.getTotalNumberOfTokensForWinningOption(prop.challengeID);
    challenge.resolved = true;

    if (voting.isPassed(prop.challengeID)) { // The challenge failed
      if (prop.processBy > now) {
        set(prop.name, prop.value);
      }
      emit ChallengeFailed(
        propID,
        prop.challengeID,
        challenge.rewardPool,
        challenge.winningTokens
      );

      require(token.transfer(prop.owner, reward), "Error:Parameterizer.resolveChallenge - Could not transfer funds");
    } else { // The challenge succeeded or nobody voted
      emit ChallengeSucceeded(
        propID,
        prop.challengeID,
        challenge.rewardPool,
        challenge.winningTokens
      );

      require(token.transfer(challenges[prop.challengeID].challenger, reward), "Error:Parameterizer.resolveChallenge - Could not transfer funds");
    }
  }

  /**
  @dev sets the param keted by the provided name to the provided value
  @param name the name of the param to be set
  @param value the value to set the param to be set
  */
  function set(string name, uint value) private {
    params[keccak256(abi.encodePacked(name))] = value;
  }
}
