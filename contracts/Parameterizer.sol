pragma solidity 0.4.25;

import "./Voting.sol";
import "./IERC20.sol";
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
  IERC20 private selfToken;
  Voting private selfVoting;
  uint private selfChallengeStake;
  uint private selfConversionRate;
  uint private selfConversionSlope;
  uint private selfListReward;
  uint private selfQuorum;
  uint private selfVoteBy;

  // uint private VOTEBY = 604800; // 7 days

  /**
  @dev constructor
  @param tokenAddr Address of the token which parameterizes this system
  @param votingAddr Address of a voting contract for the provided token
  @param challengeStake the amount, in tokenWei, that will be wagered in a challenge
  @param conversionRate TODO
  @param conversionSlope TODO
  @param listReward The amount of market tokens minted for a listing if accepted
  @param quorum Type of majority out of 100 necessary for vote success
  @param voteBy Amount of time, in seconds, that any voting poll will "run"
  */
  constructor(
    address tokenAddr,
    address votingAddr,
    uint challengeStake,
    uint conversionRate,
    uint conversionSlope,
    uint listReward,
    uint quorum,
    uint voteBy
    ) public
  {
    selfToken = IERC20(tokenAddr);
    selfVoting = Voting(votingAddr);
    selfChallengeStake = challengeStake;
    selfConversionRate = conversionRate;
    selfConversionSlope = conversionSlope;
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

  function getConversionSlope() public view returns(uint) {
    return selfConversionSlope;
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
  function resolveReparam(bytes32 paramHash) public {
    // TODO
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
