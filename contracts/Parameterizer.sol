pragma solidity 0.5.1;

import "./Voting.sol";
import "./SafeMath.sol";

contract Parameterizer {
  using SafeMath for uint;

  struct Reparam {
    address proposer; // who asked for the change
    uint8 param;
    uint value; // what to change 'name' to
  }

  // so that we don't get into casting strings <-> bytes for parameter names
  enum Params {
    undefined, // 0
    challengeStake, // 1
    conversionRate, // 2
    conversionSlopeDenominator, // 3
    conversionSlopeNumerator, // 4
    listReward, // 5
    quorum, // 6
    voteBy // 7
  }

  mapping(bytes32 => Reparam) private selfReparams;

  Voting private selfVoting;

  // the values counterpart for the Param keys.
  uint private selfChallengeStake;
  uint private selfConversionRate;
  uint private selfConversionSlopeDenominator;
  uint private selfConversionSlopeNumerator;
  uint private selfListReward;
  uint private selfQuorum;
  uint private selfVoteBy;

  // we don't need a full kinds enum here, just the one type pertinent to the p11r
  uint8 constant REPARAM  = 3;
  // uint constant VOTEBY = 604800; // 7 days

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

  function getChallengeStake() external view returns(uint) {
    return selfChallengeStake;
  }

  function getConversionRate() external view returns(uint) {
    return selfConversionRate;
  }

  function getConversionSlopeDenominator() external view returns(uint) {
    return selfConversionSlopeDenominator;
  }

  function getConversionSlopeNumerator() external view returns(uint) {
    return selfConversionSlopeNumerator;
  }

  function getListReward() external view returns(uint) {
    return selfListReward;
  }

  /**
    @dev a purely external version of how the Parameterizer generates the candidate paramHash.
    @param param The string name to get packed and hashed
    @param value The proposed value change to get packed and hashed
    @return The single hashed value of both combined args
  */
  function getParamHash(uint8 param, uint value) external pure returns(bytes32) {
    return keccak256(abi.encodePacked(param, value));
  }

  function getQuorum() external view returns(uint) {
    return selfQuorum;
  }

  function getReparam(bytes32 paramHash) external view returns(address, uint8, uint) {
    return (
      selfReparams[paramHash].proposer,
      selfReparams[paramHash].param,
      selfReparams[paramHash].value
    );
  }

  function getVoteBy() external view returns(uint) {
    return selfVoteBy;
  }

  /**
    @dev Determine if a reparam proposal collected the necessary votes, setting it if so.
    @notice This function must be called by a council member.
    @param paramHash the proposal to make a determination and possible state transition for
    TODO revisit for possibly using an enum
  */
  function resolveReparam(bytes32 paramHash) public {
    require(selfVoting.inCouncil(msg.sender) == true, "Error:Parameterizer.resolveReparam - Sender must be council member");
    require(selfVoting.candidateIs(paramHash, REPARAM) == true, "Error:Parameterizer.resolveReparam - Must be a Reparam");
    require(selfVoting.pollClosed(paramHash) == true, "Error:Parameterizer.resolveReparam - Polls for this candidate must be closed");

    // Case: reparam accepted
    if(selfVoting.didPass(paramHash, selfQuorum)) {
      uint8 param = selfReparams[paramHash].param;
      // a switch would have been nice...
      if (param == uint8(Params.challengeStake)) {
        selfChallengeStake = selfReparams[paramHash].value;
      } else if (param == uint8(Params.conversionRate)) {
        selfConversionRate = selfReparams[paramHash].value;
      } else if (param == uint8(Params.conversionSlopeDenominator)) {
        selfConversionSlopeDenominator = selfReparams[paramHash].value;
      } else if (param == uint8(Params.conversionSlopeNumerator)) {
        selfConversionSlopeNumerator = selfReparams[paramHash].value;
      } else if (param == uint8(Params.listReward)) {
        selfListReward = selfReparams[paramHash].value;
      } else if (param == uint8(Params.quorum)) {
        selfQuorum = selfReparams[paramHash].value;
      } else if (param == uint8(Params.voteBy)) {
        selfVoteBy = selfReparams[paramHash].value;
      }
    }

    // Pass or not, clean up the reparam and candidate
    selfVoting.removeCandidate(paramHash);
    delete selfReparams[paramHash];
  }

  /**
    @dev propose a reparamaterization of name's value to value.
    @param param the name of the proposed param to be set
    @param value the proposed value to set the param to be set
  */
  function reparameterize(uint8 param, uint value) external {
    require(selfVoting.inCouncil(msg.sender) == true, "Error:Parameterizer.reparameterize - Sender must be council member");

    bytes32 paramHash = keccak256(abi.encodePacked(param, value));

    // TODO any type checks or anything else for auto-rejecting reparams on certain params?

    require(selfVoting.isCandidate(paramHash) != true, "Error:Parameterizer.proposeReparam - Proposed reparam is already a voting candidate"); // Forbid duplicate proposals

    // first, stash the proposition details here
    selfReparams[paramHash] = Reparam({
      proposer: msg.sender,
      param: param,
      value: value
    });

    // now create the voting candidate for it
    selfVoting.addCandidate(
      paramHash,
      REPARAM,
      selfVoteBy
    );

    emit ReparamProposedEvent(
      msg.sender,
      paramHash,
      param,
      value
    );
  }

  event ReparamProposedEvent(address indexed proposer, bytes32 indexed paramHash, uint8 indexed param, uint value);
  event ReparamFailedEvent(bytes32 indexed propHash, string indexed name, uint value);
  event ReparamSucceededEvent(bytes32 indexed propHash, string indexed name, uint value);
}
