pragma solidity 0.5.1;

import "./SafeMath.sol";

contract Voting {

  using SafeMath for uint;

  /**
    Note that we use uint8s for Candidate.kind
    {
      application: 1,
      challenge: 2,
      reparam: 3
    }

    we could store these in an enum, but as they are passed from elsewhere, its rather useless
  */

  struct Candidate {
    uint index; // where is this key in CandidateKeys?
    uint8 kind; // member of the enum...
    uint voteBy; // timestamp via `block.timestamp` (seconds since epoch)
    uint votes; // tally of votes 'for' the candidate
    address[] voted; // an address present in this array inidicates that user voted for this poll
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
  @param hash hash which key's this candidate's kind
  @param kind uint8 that is one of our Kinds enum NOTE: this is passed by the contract itself, not input from a user
  @param voteBy some amount of time, in seconds, from "now" that the poll for this candidate will close
  */
  function addCandidate(bytes32 hash, uint8 kind, uint voteBy) external hasPrivilege {
    // can't be dupe candidates
    require(isCandidate(hash) != true, "Error:Voting.addCandidate - Candidate already exists");

    uint end = block.timestamp.add(voteBy);
    Candidate storage c = selfCandidates[hash];
    // both push the hash into the array, and get the index. NOTE: push returns array length (hence - 1)
    c.index = selfCandidateKeys.push(hash) - 1; // no need to use .sub here
    c.kind = kind; // not checking this input as we control its passing...
    c.voteBy = end;
    // c.votes = 0; this is already 0...

    emit CandidateAddedEvent(hash, kind, voteBy);
  }

  function addToCouncil(address member) external hasPrivilege {
    require(inCouncil(member) != true, "Error:Voting.addToCouncil - Already a council member");

    selfCouncil[member] = selfCouncilKeys.push(member) - 1;
  }

  /**
    @dev Determines if a given candidate is of the given kind
    @param hash The key of this candidate
    @param kind The type of candidate we expect it to be
    @return boolean
  */
  function candidateIs(bytes32 hash, uint8 kind) external view returns (bool) {
    return selfCandidates[hash].kind == kind;
  }

  /**
  @notice Determines if a candidate has enough yes votes to pass
  @dev Check if votes for exceeds quorum threshold (requires polling to be closed)
  @param hash  identifier associated with target listing or reparam
  @param quorum the current % of 100 majority that this candidate needs to pass
  */
  function didPass(bytes32 hash, uint quorum) public view returns (bool) {
    require(isCandidate(hash) == true, "Error:Voting.didPass - Candidate does not exist");
    require(selfCandidates[hash].voteBy < block.timestamp, "Error:Voting.pass - Polling must be closed for this candidate");
    require(selfCouncilKeys.length > 0, "Error:Voting.didPass - No council members");

    // Case: no one voted
    if (selfCandidates[hash].votes == 0) {
      // there _could_ be a market with a quorum set to zero
      return quorum == 0;
    } else {
      // (councilMembers*100 - (delta between councilMembers*100 - votes*100)) divided by actual number of council members
      return ((selfCouncilKeys.length * 100) - ((selfCouncilKeys.length * 100) - (selfCandidates[hash].votes * 100))) / selfCouncilKeys.length >= quorum;
    }
  }

  function didVote(bytes32 hash, address member) public view returns(bool) {
    require(isCandidate(hash) == true, "Error:Voting.didVote - Candidate does not exist");

    // return selfCandidates[hash].voted[member] == true;
    for (uint i=0; i < selfCandidates[hash].voted.length; i++) {
      // sol says we can directly compare addresses
      if (selfCandidates[hash].voted[i] == member) {
        return true;
      }
    }

    return false;
  }

  function getCandidate(bytes32 hash) external view returns (uint8, uint, uint) {
    // Not requiring existance with getters like this. TODO

    // NOTE: we don't need to return the quorum, and we cannot return the voted mapping...
    return (
      selfCandidates[hash].kind,
      selfCandidates[hash].voteBy,
      selfCandidates[hash].votes
    );
  }

  function getCandidates() external view returns(bytes32[] memory) {
    return selfCandidateKeys;
  }

  function getCouncil() external view returns(address[] memory) {
    return selfCouncilKeys;
  }

  function getPrivilegedAddresses() external view returns(address, address) {
    return (selfMarketAddress, selfParameterizerAddress);
  }

  modifier hasPrivilege() {
    require(
      msg.sender == selfOwner ||
      msg.sender == selfMarketAddress ||
      msg.sender == selfParameterizerAddress, "Error:Voting.hasPrivilege - Sender must be a privileged contract");
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
    // NOTE: we _are_ checking existance here as otherwise result might be surprising
    require(isCandidate(hash) == true, "Error:Voting.pollClosed - Candidate does not exist");

    return selfCandidates[hash].voteBy < block.timestamp;
  }

  // not sold that we need to forbid the removing of a candidate with an 'active' voteBy,
  // as this is triggered internally, so not implementing for now...
  function removeCandidate(bytes32 hash) external hasPrivilege {
    require(isCandidate(hash) == true, "Error:Voting.removeCandidate - Candidate does not exist");

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
  }

  function removeFromCouncil(address member) external hasPrivilege {
    if (inCouncil(member) == true) {
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
    }
  }

  function setPrivilegedContracts(address market, address parameterizer) external isOwner {
    // can only ever be set once
    require(selfMarketAddress == address(0), "Error:Voting.setPrivilegedContracts - Market address already set");
    require(selfParameterizerAddress == address(0), "Error:Voting.setPrivilegedContracts - Parameterizer address already set");
    selfMarketAddress = market;
    selfParameterizerAddress = parameterizer;
  }

  /**
  @notice Commits vote using hash of choice and secret salt to conceal vote until reveal
  @param hash bytes32 identifier associated with target listing or reparam
  @return true if a success
  */
  function vote(bytes32 hash) external {
    require(inCouncil(msg.sender) == true, "Error:Voting.vote - Sender must be council member");
    require(isCandidate(hash) == true, "Error:Voting.vote - Candidate does not exist");
    require(selfCandidates[hash].voteBy > block.timestamp, "Error:Voting.vote - Polling is closed for this candidate");
    require(didVote(hash, msg.sender) != true, "Error:Voting.vote - Sender has already voted");

    selfCandidates[hash].voted.push(msg.sender); // we will keep track of who voted
    selfCandidates[hash].votes = selfCandidates[hash].votes.add(1);

    // NOTE: this creates the public record of the vote being cast (read: not concealed)
    emit VotedEvent(msg.sender, hash);
  }

  event VotedEvent(address indexed voter, bytes32 indexed hash);
  event CandidateAddedEvent(bytes32 indexed hash, uint8 indexed kind, uint indexed voteBy);
  event CandidateRemovedEvent(bytes32 indexed hash);
  event CouncilMemberAddedEvent(address indexed member);
  event CouncilMemberRemovedEvent(address indexed member);
}
