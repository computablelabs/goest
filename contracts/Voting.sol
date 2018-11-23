pragma solidity 0.4.25;

import "./SafeMath.sol";

contract Voting {

  using SafeMath for uint;

  struct Candidate {
    uint index; // where is this key in CandidateKeys?
    bytes kind; // listing | reparameterization | challenge
    uint voteBy; // timestamp via `block.timestamp` (seconds since epoch)
    uint votes; // tally of votes 'for' the candidate
    mapping(address => bool) voted;  // indicates whether an address committed a vote for this poll
  }

  // a mapping to hold the actual candidate data
  mapping(bytes32 => Candidate) private selfCandidates;
  // we use a simple dynamic array to house all candidate keys (hashes)
  bytes32[] private selfCandidateKeys;
  // a simple mapping of current countil members to an index in the council array
  mapping(address => uint) private selfCouncil;
  // dynamic array holding current council member's adrresses
  address[] private selfCouncilKeys;
  // the market and parameterizer have priveledge to add candidates, council members etc
  address private selfMarketAddress;
  address private selfParameterizerAddress;
  // this will be the market factory
  address private selfOwner;

  constructor() public {
    selfOwner = msg.sender;
  }

  /**
  @dev Place a new candidate into the candidates array with the necessary info
  @param kind string 'application' | 'challenge' | 'reparam'. NOTE: this is passed by the contract itself, not input from a user
  @param hash bytes32 hash which key's this candidate's kind
  @param voteBy some amount of time, in seconds, from "now" that the poll for this candidate will close
  @return true if a success
  */
  function addCandidate(string kind, bytes32 hash, uint voteBy) external hasPrivilege returns(bool) {
    // can't be dupe candidates
    require(!isCandidate(hash), "Error:Voting.addCandidate - Candidate already exists");

    uint end = block.timestamp.add(voteBy);

    Candidate storage c = selfCandidates[hash];

    // both push the hash into the array, and get the index. NOTE: push returns array length (hence - 1)
    c.index = selfCandidateKeys.push(hash) - 1; // no need to use .sub here
    c.kind = bytes(kind); // again, not checking the string as we control its passing
    c.voteBy = end;
    c.votes = 0;

    emit CandidateAddedEvent(kind, hash, voteBy);

    return true;
  }

  function addToCouncil(address member) external hasPrivilege returns(bool) {
    require(!inCouncil(member), "Error:Voting.addToCouncil - Already a council member");
    selfCouncil[member] = selfCouncilKeys.push(member) - 1;
  }

  /**
    @dev Determines if a given candidate is of the given kind
    @param hash The key of this candidate
    @param kind The type of candidate we expect it to be
    @return boolean
  */
  function candidateIs(bytes32 hash, string kind) external view returns (bool) {
    require(isCandidate(hash), "Error:Voting.pollClosed - Candidate does not exist");
    // fast check for an obvious case
    if (selfCandidates[hash].kind.length != bytes(kind).length) return false;
    // more thorough if the lengths match. we can't compare strings but we can compare hashes
    else return keccak256(string(selfCandidates[hash].kind)) == keccak256(kind);
  }

  /**
  @notice Determines if a candidate has enough yes votes to pass
  @dev Check if votes for exceeds quorum threshold (requires polling to be closed)
  @param hash  identifier associated with target listing or reparam
  @param quorum the current % of 100 majority that this candidate needs to pass
  */
  function didPass(bytes32 hash, uint quorum) public view returns (bool) {
    require(isCandidate(hash), "Error:Voting.pass - Candidate does not exist");
    require(selfCandidates[hash].voteBy < block.timestamp, "Error:Voting.pass - Polling must be closed for this candidate");
    require(selfCouncilKeys.length > 0, "Error:Voting.didPass - No council members");

    return ((selfCandidates[hash].votes.div(selfCouncilKeys.length)).mul(100) > quorum);
  }

  function didVote(bytes32 hash, address member) public view returns(bool) {
    // likely do not need to check candidate validity here TODO
    return selfCandidates[hash].voted[member] == true;
  }

  function getCandidate(bytes32 hash) external view returns (string, uint, uint) {
    require(isCandidate(hash), "Error:Voting.getCandidate - Candidate does not exist");

    // NOTE: we don't need to return the quorum, and we cannot return the voted mapping...
    return (
      string(selfCandidates[hash].kind),
      selfCandidates[hash].voteBy,
      selfCandidates[hash].votes
    );
  }

  function getCandidates() external view returns(bytes32[]) {
    return selfCandidateKeys;
  }

  function getCouncil() external view returns(address[]) {
    return selfCouncilKeys;
  }

  function getPrivilegedAddresses() external view returns(address, address) {
    return (selfMarketAddress, selfParameterizerAddress);
  }

  modifier hasPrivilege() {
    require(msg.sender == selfMarketAddress || msg.sender == selfParameterizerAddress, "Error:Voting.hasPrivilege - Sender must be a privileged contract");
    _;
  }

  modifier isOwner() {
    require(msg.sender == selfOwner, "Error:Voting.isOwner - Sender must be owner");
    _;
  }

  function inCouncil(address member) public view returns (bool) {
    // if the keys array is empty, obv not...
    if(selfCouncilKeys.length == 0) return false;
    // now just confirm that they match
    return(selfCouncilKeys[selfCouncil[member]] == member);
  }

  function isCandidate(bytes32 hash) public view returns (bool) {
    // if the keys array is empty, obv not...
    if(selfCandidateKeys.length == 0) return false;
    // now just confirm that they match
    return(selfCandidateKeys[selfCandidates[hash].index] == hash);
  }

  function pollClosed(bytes32 hash) external view returns (bool) {
    require(isCandidate(hash), "Error:Voting.pollClosed - Candidate does not exist");
    return selfCandidates[hash].voteBy < block.timestamp;
  }

  // not sold that we need to forbid the removing of a candidate with an 'active' voteBy,
  // as this is triggered internally, so not implementing for now...
  function removeCandidate(bytes32 hash) external hasPrivilege returns(bool) {
    require(isCandidate(hash), "Error:Voting.removeCandidate - Candidate does not exist");

    // first let's efficiently prune the array of keys
    uint deleted = selfCandidates[hash].index; // getting rid of this one
    bytes32 moved = selfCandidateKeys[selfCandidateKeys.length - 1]; // moving this one to where 'deleted' was
    selfCandidateKeys[deleted] = moved; // delete target now overwritten
    selfCandidateKeys.length--; // drops the last item, which we already moved

    // now update the index of the moved candidate
    selfCandidates[moved].index = deleted;

    // finally zero out the old candidate
    delete selfCandidates[hash];

    emit CandidateRemovedEvent(hash);

    return true;
  }

  function removeFromCouncil(address member) external hasPrivilege returns(bool) {
    require(inCouncil(member), "Error:Voting.removeFromCouncil - Not a member");

    // first let's efficiently prune the array of keys
    uint deleted = selfCouncil[member]; // getting rid of this one
    address moved = selfCouncilKeys[selfCouncilKeys.length - 1]; // moving this one to where 'deleted' was
    selfCouncilKeys[deleted] = moved; // delete target now overwritten
    selfCouncilKeys.length--; // drops the last item, which we already moved

    // now update the index of the moved member
    selfCouncil[moved] = deleted;

    // finally zero out the removed member
    delete selfCouncil[member];

    emit CouncilMemberRemovedEvent(member);

    return true;
  }

  function setPrivilegedContracts(address market, address parameterizer) external isOwner returns (bool) {
    // can only ever be set once
    require(selfMarketAddress == address(0), "Error:Voting.setPrivilegedContracts - Market address already set");
    require(selfParameterizerAddress == address(0), "Error:Voting.setPrivilegedContracts - Parameterizer address already set");
    selfMarketAddress = market;
    selfParameterizerAddress = parameterizer;
    return true;
  }

  /**
  @notice Commits vote using hash of choice and secret salt to conceal vote until reveal
  @param hash bytes32 identifier associated with target listing or reparam
  @return true if a success
  */
  function vote(bytes32 hash) external returns (bool) {
    require(inCouncil(msg.sender), "Error:Voting.vote - Sender must be council member");
    require(isCandidate(hash), "Error:Voting.vote - Candidate does not exist");
    require(selfCandidates[hash].voteBy > block.timestamp, "Error:Voting.vote - Polling is closed for this candidate");
    require(!didVote(hash, msg.sender), "Error:Voting.vote - Sender has already voted");

    selfCandidates[hash].voted[msg.sender] = true; // we will keep track of who voted
    selfCandidates[hash].votes = selfCandidates[hash].votes.add(1);

    // NOTE: this creates the public record of the vote being cast (read: not concealed)
    emit VotedEvent(msg.sender, hash);
    return true;
  }

  event VotedEvent(address indexed voter, bytes32 indexed hash);
  event CandidateAddedEvent(string indexed kind, bytes32 indexed hash, uint indexed voteBy);
  event CandidateRemovedEvent(bytes32 indexed hash);
  event CouncilMemberAddedEvent(address indexed member);
  event CouncilMemberRemovedEvent(address indexed member);
}
