# @title Computable Voting
# @notice A simplified, weightless, public Voting contract
# @author Computable

# No poll will collect more than 1k votes
MAX_LENGTH: constant(uint256) = 1000

struct Candidate:
  kind: uint256 # one of [1,2,3,4] representing an application, challenge, reparam or registration respectively
  owner: address
  vote_by: timestamp
  votes: address[MAX_LENGTH]
  votes_length: int128

CandidateAdded: event({hash: indexed(bytes32), kind: indexed(uint256), owner: indexed(address), voteBy: timestamp})
CandidateRemoved: event({hash: indexed(bytes32)})
CouncilMemberAdded: event({member: indexed(address)})
CouncilMemberRemoved: event({member: indexed(address)})
Voted: event({hash: indexed(bytes32), voter: indexed(address)})

candidates: map(bytes32, Candidate)
council: map(address, bool)
council_count: uint256
parameterizer_address: address
datatrust_address: address
listing_address: address
investing_address: address
factory_address: address

@public
def __init__():
  self.factory_address = msg.sender


@public
@constant
def getPrivileged() -> (address, address, address, address, address):
  """
  @notice Fetch a list of each privileged address recognized by this contract
  @return factory, market, and parameterizer addresses
  """
  return (self.factory_address, self.parameterizer_address, self.datatrust_address, self.listing_address, self.investing_address)


@public
def setPrivileged(parameterizer: address, datatrust: address, listing: address, investing: address):
  """
  @notice Allow the Market Factory to set privileged contract addresses
  """
  assert msg.sender == self.factory_address
  self.parameterizer_address = parameterizer
  self.datatrust_address = datatrust
  self.listing_address = listing
  self.investing_address = investing


@private
@constant
def has_privilege(sender: address) -> bool:
  """
  @notice Return a bool indicating whether the given address is a member of this contracts privileged group
  @return bool
  """
  return (sender == self.factory_address or sender == self.parameterizer_address or sender == self.datatrust_address or sender == self.listing_address or sender == self.investing_address)


@public
@constant
def getCouncilCount() -> uint256:
  """
  @notice returns the current number of council members
  """
  return self.council_count


@public
@constant
def inCouncil(member: address) -> bool:
  """
  @notice Check if a given address is a member of the council
  @return bool
  """
  return self.council[member]


@public
def addToCouncil(member: address):
  """
  @notice Add a member to the Market Council
  @dev Only privileged contracts may call this method
  @param member The address of the possible new council member
  """
  assert self.has_privilege(msg.sender)
  assert not self.council[member]
  self.council[member] = True
  self.council_count += 1
  log.CouncilMemberAdded(member)


@public
def removeFromCouncil(member: address):
  """
  @notice Remove a member from the council
  """
  assert self.has_privilege(msg.sender)
  assert self.council[member]
  clear(self.council[member])
  self.council_count -= 1
  log.CouncilMemberRemoved(member)


@public
@constant
def candidateIs(hash: bytes32, kind: uint256) -> bool:
  """
  @notice Is a given candidate the given kind?
  @param hash The candidate to check
  @param kind The kind to compare
  @return bool
  """
  return self.candidates[hash].kind == kind


@public
@constant
def isCandidate(hash: bytes32) -> bool:
  """
  @notice Check if a given hash identifier is a current candidate
  @return bool
  """
  return self.candidates[hash].owner != ZERO_ADDRESS


@public
@constant
def getCandidate(hash: bytes32) -> (uint256, address, timestamp, int128):
  """
  @notice Return information about the given candidate identified by the given hash
  @dev Hash argument keys a candidate struct in the candidates mapping
  @return The type, vote_by timestamp and number of votes recieved
  """
  return (self.candidates[hash].kind, self.candidates[hash].owner, self.candidates[hash].vote_by, self.candidates[hash].votes_length)


@public
@constant
def getCandidateOwner(hash: bytes32) -> address:
  """
  @notice return the owner address of a given Candidate.
  """
  return self.candidates[hash].owner


@public
def addCandidate(hash: bytes32, kind: uint256, owner: address, vote_by: timedelta):
  """
  @notice Given a listing or parameter hash, create a new voting candidate
  @dev Only priveliged contracts may call this method
  @param hash The identifier for the listing or reparameterization candidate
  @param kind The type of candidate we are adding
  @param owner The adress which owns this created candidate
  @param vote_by How long into the future until polls for this candidate close
  """
  assert self.has_privilege(msg.sender)
  assert self.candidates[hash].owner == ZERO_ADDRESS
  end: timestamp = block.timestamp + vote_by
  self.candidates[hash].kind = kind
  self.candidates[hash].owner = owner
  self.candidates[hash].vote_by = end
  log.CandidateAdded(hash, kind, owner, end)


@public
def removeCandidate(hash: bytes32):
  """
  @notice Remove a candidate from the current list
  @dev Clears all members from a Candidate pointed to by a hash (enabling re-use)
  """
  assert self.has_privilege(msg.sender)
  assert self.candidates[hash].owner != ZERO_ADDRESS
  clear(self.candidates[hash].kind)
  clear(self.candidates[hash].owner)
  clear(self.candidates[hash].vote_by)
  # we must clear individual votes that exist
  for i in range(1000):
    if i == self.candidates[hash].votes_length: # don't interate past the actual number of votes
      break
    else:
      clear(self.candidates[hash].votes[i])
  clear(self.candidates[hash].votes_length)
  log.CandidateRemoved(hash)


@public
@constant
def didPass(hash: bytes32, quorum: uint256) -> bool:
  """
  @notice Return a bool indicating whether a given candidate recieved enough votes to exceed the quorum
  @dev The poll must be closed. Also we cover the corner case that no one voted.
  @return bool
  """
  assert self.candidates[hash].owner != ZERO_ADDRESS
  assert self.candidates[hash].vote_by < block.timestamp
  # edge case that no one voted
  if self.candidates[hash].votes_length == 0:
    # theoretically a market could have a 0 quorum
    return quorum == 0
  else:
    return (convert((self.candidates[hash].votes_length * 100), uint256) / self.council_count) >= quorum


@public
@constant
def didVote(hash: bytes32, member: address) -> bool:
  """
  @notice Check to see if a given member has voted for a given candidate
  @return bool
  """
  voted: bool = False
  # NOTE we must use a literal here as vyper won't allow the const as an upper bound
  for i in range(1000):
    if i == self.candidates[hash].votes_length: # don't interate past the actual number of votes
      break
    elif self.candidates[hash].votes[i] == member:
      voted = True
      break
  return voted


@public
@constant
def pollClosed(hash: bytes32) -> bool:
  """
  @notice Check to see if a given candidate's polling has closed
  @return bool
  """
  assert self.candidates[hash].owner != ZERO_ADDRESS
  return self.candidates[hash].vote_by < block.timestamp


@public
def vote(hash: bytes32):
  """
  @notice Cast a vote for a given candidate
  @dev Voter must be a council member, and not have already voted
  """
  assert self.council[msg.sender]
  assert self.candidates[hash].owner != ZERO_ADDRESS
  assert self.candidates[hash].vote_by > block.timestamp
  assert not self.didVote(hash, msg.sender)
  # here we use the number of votes as an index
  self.candidates[hash].votes[self.candidates[hash].votes_length] = msg.sender
  self.candidates[hash].votes_length += 1
