# @title Computable Voting
# @notice A simplified, weightless, public Voting contract
# @author Computable

# No council may exceed 1000 members. There may not be more than 1000 active polls at once.
MAX_LENGTH: constant(uint256) = 1000

struct Candidate:
  index: int128
  kind: uint256 # one of [1,2,3] representing an application, challenge or reparam respectively
  voteBy: timestamp
  voted: map(address, bool)
  votes: uint256

CandidateAdded: event({hash: indexed(bytes32), kind: indexed(uint256), voteBy: indexed(timestamp)})
CandidateRemoved: event({hash: indexed(bytes32)})
CouncilMemberAdded: event({member: indexed(address)})
CouncilMemberRemoved: event({member: indexed(address)})
Voted: event({hash: indexed(bytes32), voter: indexed(address)})

candidates: map(bytes32, Candidate)
candidates_length: int128 # represents the actual number of candidates currently
candidate_keys: bytes32[MAX_LENGTH]
council: map(address, int128)
council_length: int128 # represents the actual number of council members
council_keys: address[MAX_LENGTH]
market_address: address
parameterizer_address: address
factory_address: address

@public
def __init__():
  self.factory_address = msg.sender


@public
@constant
def getPrivileged() -> (address, address, address):
  """
  @notice Fetch a list of each privileged address recognized by this contract
  @return factory, market, and parameterizer addresses
  """
  return (self.factory_address, self.market_address, self.parameterizer_address)


@public
def setPrivileged(market: address, parameterizer: address):
  """
  @notice Allow the Market Factory to set privileged contract addresses
  """
  assert msg.sender == self.factory_address
  self.market_address = market
  self.parameterizer_address = parameterizer


@private
@constant
def has_privilege(sender: address) -> bool:
  """
  @notice Return a bool indicating whether the given address is a member of this contracts privileged group
  @return bool
  """
  return (sender == self.factory_address or sender == self.market_address or sender == self.parameterizer_address)


@public
@constant
def inCouncil(member: address) -> bool:
  """
  @notice Check if a given address is a member of the council
  @return bool
  """
  if self.council_length == 0:
    return False
  return self.council_keys[self.council[member]] == member


@public
@constant
def getCouncilCount() -> int128:
  """
  @notice Return the current number of council members
  """
  return self.council_length


@public
def addToCouncil(member: address):
  """
  @notice Add a member to the Market Council
  @dev The MAX_LENGTH is enforced here, as well as the privilege of the caller
  @param member The address of the possible new council member
  """
  assert self.has_privilege(msg.sender)
  assert self.council_length < MAX_LENGTH
  assert not self.inCouncil(member)
  # the council mapping simply points to an index in the unordered list
  self.council[member] = self.council_length
  # actually place it in said list...
  self.council_keys[self.council_length] = member
  # move the pointer along
  self.council_length += 1
  log.CouncilMemberAdded(member)


@public
def removeFromCouncil(member: address):
  """
  @notice Remove a member from the council
  @dev Like removeCandidate, we simply overwrite the removed member and update
  """
  assert self.has_privilege(msg.sender)
  assert self.inCouncil(member)
  # move the pointer to the latest entry
  self.council_length -= 1
  if self.council_length == 0:
    self.council_keys[0] = ZERO_ADDRESS
  else:
    # this target member is being overwritten
    deleted: int128 = self.council[member]
    # the latest member is used to overwrite said target
    moved: address = self.council_keys[self.council_length]
    self.council_keys[deleted] = moved
    # zero out this index in the unordered list
    self.council_keys[self.council_length] = ZERO_ADDRESS
    # update the index of the moved member
    self.council[moved] = deleted
  # either way, a -1 assures this index is invalid
  self.council[member] = -1
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
  if self.candidates_length == 0:
    return False
  return self.candidate_keys[self.candidates[hash].index] == hash


@public
@constant
def wasCandidate(hash: bytes32) -> bool:
  """
  @notice Return true if a given hash points to a removed candidate
  @dev An non-active "removed" candidate mapping is one whose index is -1
  @return bool
  """
  return self.candidates[hash].index == -1


@public
@constant
def getCandidateCount() -> int128:
  """
  @notice Return the number of current candidates
  @dev This count includes those whose voteBy has passed, but have not been resolved yet
  @return The number of candidates
  """
  return self.candidates_length


@public
@constant
def getCandidateKey(index: int128) -> bytes32:
  """
  @notice Return the hashed identifier that points to a given candidate
  @dev The given int is the index at which said hashed identifier is located in our candidate_keys list.
  @return The hashed identifier
  """
  return self.candidate_keys[index]


@public
@constant
def getCandidate(hash: bytes32) -> (uint256, timestamp, uint256):
  """
  @notice Return information about the given candidate identified by the given hash
  @dev Hash argument keys a candidate struct in the candidates mapping
  @return The type, voteBy timestamp and number of votes recieved
  """
  return (self.candidates[hash].kind, self.candidates[hash].voteBy, self.candidates[hash].votes)


@public
def addCandidate(hash: bytes32, kind: uint256, voteBy: timedelta):
  """
  @notice Given a listing or parameter hash, create a new voting candidate
  @dev we enforce that this hash is not a current candidate, nor a past one (hashes must always be unique)
  @param hash The identifier for the listing or reparameterization candidate
  @param kind The type of candidate we are adding
  """
  assert self.has_privilege(msg.sender)
  assert self.candidates_length < MAX_LENGTH
  assert not self.isCandidate(hash)
  assert not self.wasCandidate(hash)
  end: timestamp = block.timestamp + voteBy
  # place candidate into the mapping with unordered list index pointer along with kind and vote-by
  self.candidates[hash].index = self.candidates_length
  self.candidates[hash].kind = kind
  self.candidates[hash].voteBy = end
  # "push" the new candidate into the unordered list
  self.candidate_keys[self.candidates_length] = hash
  # move the pointer to the next empty list location
  self.candidates_length += 1
  log.CandidateAdded(hash, kind, end)


@public
def removeCandidate(hash: bytes32):
  """
  @notice Remove a candidate from the current list
  @dev Simply overwrite the target with the latest candidate, then update bookkeeping
  """
  assert self.has_privilege(msg.sender)
  assert self.isCandidate(hash)
  # go ahead and back this up to point at the last entry
  self.candidates_length -= 1
  # in the case of a single current candidate, the operation is simplified
  if self.candidates_length == 0:
    self.candidate_keys[0] = EMPTY_BYTES32
  else:
    # this target candidate is being overwritten
    deleted: int128 = self.candidates[hash].index
    # the latest candidate is used to overwrite said target
    moved: bytes32 = self.candidate_keys[self.candidates_length]
    self.candidate_keys[deleted] = moved
    # zero out this index of the unordered list
    self.candidate_keys[self.candidates_length] = EMPTY_BYTES32
    # update the index of the moved candidate
    self.candidates[moved].index = deleted
  # regardless, we denote this mapping entry as removed with a negative index
  self.candidates[hash].index = -1
  log.CandidateRemoved(hash)


@public
@constant
def didPass(hash: bytes32, quorum: uint256) -> bool:
  """
  @notice Return a bool indicating whether a given candidate recieved enough votes to exceed the quorum
  @dev The poll must be closed. Also we cover the corner case that no one voted.
  @return bool
  """
  assert self.isCandidate(hash)
  assert self.candidates[hash].voteBy < block.timestamp
  # edge case that no one voted
  if self.candidates[hash].votes == 0:
    # theoretically a market could have a 0 quorum
    return quorum == 0
  else:
    return (self.candidates[hash].votes * 100) / (convert(self.council_length, uint256)) >= quorum


@public
@constant
def didVote(hash: bytes32, member: address) -> bool:
  """
  @notice Check to see if a given member has voted for a given candidate
  @return bool
  """
  assert self.isCandidate(hash)
  return self.candidates[hash].voted[member] == True

@public
@constant
def pollClosed(hash: bytes32) -> bool:
  """
  @notice Check to see if a given candidate's polling has closed
  @return bool
  """
  assert self.isCandidate(hash)
  return self.candidates[hash].voteBy < block.timestamp


@public
def vote(hash: bytes32):
  """
  @notice Cast a vote for a given candidate
  @dev Voter must be a council member, and not have already voted
  """
  assert self.inCouncil(msg.sender)
  assert self.isCandidate(hash)
  assert self.candidates[hash].voteBy > block.timestamp
  assert not self.didVote(hash, msg.sender)
  # here we use the number of votes as an index
  self.candidates[hash].voted[msg.sender] = True
  self.candidates[hash].votes += 1
