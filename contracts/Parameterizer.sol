pragma solidity 0.4.25;

import "./Voting.sol";
import "./SafeMath.sol";

contract Parameterizer {
  using SafeMath for uint;

  struct Reparam {
    address proposer; // who asked for the change
    bytes name; // string name of the param converted to a byte array
    uint value; // what to change 'name' to
  }

  mapping(bytes32 => Reparam) private selfReparams;

  // Global Variables
  Voting private selfVoting;
  uint private selfChallengeStake;
  uint private selfConversionRate;
  uint private selfConversionSlopeDenominator;
  uint private selfConversionSlopeNumerator;
  uint private selfListReward;
  uint private selfQuorum;
  uint private selfVoteBy;

  // uint private VOTEBY = 604800; // 7 days

  /**
  @dev constructor
  @param votingAddr Address of a voting contract for the provided token
  @param challengeStake the amount, in tokenWei, that will be wagered in a challenge, ex: 10**18 (1 token)
  @param conversionRate expressed in tokenWei, ex: 10**17 (.1)
  @param conversionSlopeDenominator a scaling factor, ex: 101
  @param conversionSlopeNumerator a scaling factor, ex: 100
  @param listReward The amount of market tokens minted for a listing if accepted, ex: 10**18 (1 token)
  @param quorum Type of majority out of 100 necessary for vote success, ex: 50
  @param voteBy Amount of time, in seconds, that any voting poll will remain open for votes to be cast
  */
  constructor(
    address votingAddr,
    uint challengeStake,
    uint conversionRate,
    uint conversionSlopeDenominator,
    uint conversionSlopeNumerator,
    uint listReward,
    uint quorum,
    uint voteBy
    ) public
  {
    selfVoting = Voting(votingAddr);
    selfChallengeStake = challengeStake;
    selfConversionRate = conversionRate;
    selfConversionSlopeDenominator = conversionSlopeDenominator;
    selfConversionSlopeNumerator = conversionSlopeNumerator;
    selfListReward = listReward;
    selfQuorum = quorum;
    selfVoteBy = voteBy;
  }

  function getChallengeStake() public view returns(uint) {
    return selfChallengeStake;
  }

  function getConversionRate() public view returns(uint) {
    return selfConversionRate;
  }

  function getConversionSlopeDenominator() public view returns(uint) {
    return selfConversionSlopeDenominator;
  }

  function getConversionSlopeNumerator() public view returns(uint) {
    return selfConversionSlopeNumerator;
  }

  function getListReward() public view returns(uint) {
    return selfListReward;
  }

  function getQuorum() public view returns(uint) {
    return selfQuorum;
  }

  function getVoteBy() public view returns(uint) {
    return selfVoteBy;
  }

  /**
  @dev Determine if a reparam proposal collected the necessary votes, setting it if so
  @param paramHash the proposal to make a determination and possible state transition for
  */
  function resolveReparam(bytes32 paramHash) public returns (bool) {
    require(selfVoting.inCouncil(msg.sender), "Error:Parameterizer.resolveReparam - Sender must be council member");
    require(selfVoting.candidateIs(paramHash, "reparam"), "Error:Parameterizer.resolveReparam - Must be a Reparam");
    require(selfVoting.pollClosed(paramHash), "Error:Parameterizer.resolveReparam - Polls for this candidate must be closed");

    // Case: reparam accepted
    if(selfVoting.didPass(paramHash, getQuorum())) {
      bytes32 cmp = keccak256(string(selfReparams[paramHash].name));
      // a switch would have been nice... TODO better way to compare?
      if (cmp == keccak256("challengeStake")) {
        selfChallengeStake = selfReparams[paramHash].value;
      } else if (cmp == keccak256("conversionRate")) {
        selfConversionRate = selfReparams[paramHash].value;
      } else if (cmp == keccak256("conversionSlopeDenominator")) {
        selfConversionSlopeDenominator = selfReparams[paramHash].value;
      } else if (cmp == keccak256("conversionSlopeNumerator")) {
        selfConversionSlopeNumerator = selfReparams[paramHash].value;
      } else if (cmp == keccak256("listReward")) {
        selfListReward = selfReparams[paramHash].value;
      } else if (cmp == keccak256("quorum")) {
        selfQuorum = selfReparams[paramHash].value;
      } else if (cmp == keccak256("voteBy")) {
        selfVoteBy = selfReparams[paramHash].value;
      }
    }

    // clean up the reparam and candidate
    delete selfReparams[paramHash];
    return true;
  }

  /**
  @notice propose a reparamaterization of the key name's value to value.
  @param name the name of the proposed param to be set
  @param value the proposed value to set the param to be set
  */
  function reparameterize(string name, uint value) external returns (bytes32) {
    require(selfVoting.inCouncil(msg.sender), "Error:Parameterizer.reparameterize - Sender must be council member");

    bytes32 paramHash = keccak256(abi.encodePacked(name, value));

    // TODO any type checks or anything else for auto-rejecting reparams on certain params?

    require(!selfVoting.isCandidate(paramHash), "Error:Parameterizer.proposeReparam - Proposed reparam is already a voting candidate"); // Forbid duplicate proposals

    // first, stash the proposition details here
    selfReparams[paramHash] = Reparam({
      proposer: msg.sender,
      name: bytes(name),
      value: value
    });

    // now create the voting candidate for it
    bool added = selfVoting.addCandidate(
      "reparam",
      paramHash,
      getVoteBy()
    );

    require(added, "Error:Parameterizer.reparameterize - Could not add reparam to voting candidates list");

    emit ReparamProposedEvent(
      msg.sender,
      paramHash,
      name,
      value
    );

    return paramHash;
  }

  event ReparamProposedEvent(address indexed proposer, bytes32 indexed paramHash, string indexed name, uint value);
  event ReparamFailedEvent(bytes32 indexed propHash, string indexed name, uint value);
  event ReparamSucceededEvent(bytes32 indexed propHash, string indexed name, uint value);
}
