# @title Computable Voting
# @notice A simplified, weightless, public Voting contract
# @author Computable

# No council may exceed 1000 members. There may not be more than 1000 active polls at once.
MAX_LENGTH: constant(uint256) = 1000

struct Candidate:
  index: int128
  kind: uint256 # one of [1,2,3] representing an application, challenge or reparam respectively
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
candidates_length: int128 # represents the actual number of candidates currently
candidate_keys: bytes32[MAX_LENGTH]
council: map(address, int128)
council_length: int128 # represents the actual number of council members
council_keys: address[MAX_LENGTH]
listing_address: address
investing_address: address
parameterizer_address: address
factory_address: address

@public
def __init__():
  self.factory_address = msg.sender


@public
@constant
def getPrivileged() -> (address, address, address, address):
  """
  @notice Fetch a list of each privileged address recognized by this contract
  @return factory, market, and parameterizer addresses
  """
  return (self.factory_address, self.parameterizer_address, self.listing_address, self.investing_address)


@public
def setPrivileged(parameterizer: address, listing: address, investing: address):
  """
  @notice Allow the Market Factory to set privileged contract addresses
  """
  assert msg.sender == self.factory_address
  self.parameterizer_address = parameterizer
  self.listing_address = listing
  self.investing_address = investing


@private
@constant
def has_privilege(sender: address) -> bool:
  """
  @notice Return a bool indicating whether the given address is a member of this contracts privileged group
  @return bool
  """
  return (sender == self.factory_address or sender == self.parameterizer_address or sender == self.listing_address or sender == self.investing_address)


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
  @dev As only privileged contracts may call this method, they will first assert that
  `willAddToCouncil` is true (thus we do not repeat that here)
  @param member The address of the possible new council member
  """
  assert self.has_privilege(msg.sender)
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
  if self.council_length > 0:
    # this target member is being overwritten
    deleted: int128 = self.council[member]
    # the latest member is used to overwrite said target
    moved: address = self.council_keys[self.council_length]
    self.council_keys[deleted] = moved
    # update the index of the moved member
    self.council[moved] = deleted
  # zero out the latest in the unordered list
  self.council_keys[self.council_length] = ZERO_ADDRESS
  # we are fine with just zeroed council membership (as it may recur)
  clear(self.council[member])
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
def getCandidateCount() -> int128:
  """
  @notice Return the number of current candidates
  @dev This count includes those whose vote_by has passed, but have not been resolved yet
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
  @dev The priveliged contracts which call this method perform `willAddCandidate()` check first
  @param hash The identifier for the listing or reparameterization candidate
  @param kind The type of candidate we are adding
  @param owner The adress which owns this created candidate
  @param vote_by How long into the future until polls for this candidate close
  """
  assert self.has_privilege(msg.sender)
  end: timestamp = block.timestamp + vote_by
  # place candidate into the mapping with unordered list index pointer along with kind and vote-by
  self.candidates[hash].index = self.candidates_length
  self.candidates[hash].kind = kind
  self.candidates[hash].owner = owner
  self.candidates[hash].vote_by = end
  # "push" the new candidate into the unordered list
  self.candidate_keys[self.candidates_length] = hash
  # move the pointer to the next empty list location
  self.candidates_length += 1
  log.CandidateAdded(hash, kind, owner, end)


@public
def removeCandidate(hash: bytes32):
  """
  @notice Remove a candidate from the current list
  @dev Simply overwrite the target with the latest candidate, then update bookkeeping.
  NOTE: this does mean that all candidate hashes must be unique
  Note that candidates are never fully removed from the mapping, but are set to `index: -1` to indicate status
  """
  assert self.has_privilege(msg.sender)
  assert self.isCandidate(hash)
  # go ahead and back this up to point at the last entry
  self.candidates_length -= 1
  # in the case of a single current candidate, the operation is simplified
  if self.candidates_length > 0:
    # this target candidate is being overwritten
    deleted: int128 = self.candidates[hash].index
    # the latest candidate is used to overwrite said target
    moved: bytes32 = self.candidate_keys[self.candidates_length]
    self.candidate_keys[deleted] = moved
    # update the index of the moved candidate
    self.candidates[moved].index = deleted
  # zero out the latest entry (now a dupe)
  clear(self.candidate_keys[self.candidates_length])
  clear(self.candidates[hash].index)
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
  assert self.isCandidate(hash)
  assert self.candidates[hash].vote_by < block.timestamp
  # edge case that no one voted
  if self.candidates[hash].votes_length == 0:
    # theoretically a market could have a 0 quorum
    return quorum == 0
  else:
    return ((self.candidates[hash].votes_length * 100) / self.council_length) >= convert(quorum, int128)


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
  assert self.isCandidate(hash)
  return self.candidates[hash].vote_by < block.timestamp


@public
def vote(hash: bytes32):
  """
  @notice Cast a vote for a given candidate
  @dev Voter must be a council member, and not have already voted
  """
  assert self.inCouncil(msg.sender)
  assert self.isCandidate(hash)
  assert self.candidates[hash].vote_by > block.timestamp
  assert not self.didVote(hash, msg.sender)
  # here we use the number of votes as an index
  self.candidates[hash].votes[self.candidates[hash].votes_length] = msg.sender
  self.candidates[hash].votes_length += 1


@public
@constant
def willAddCandidate(hash: bytes32) -> bool:
  """
  @notice Return a bool indicating if a possible candidate would be added
  @param hash Identifier for a (suppsosedly) currently non-existant candidate
  """
  # TODO possible don't call the internal methods and just do the comparisons here
  return self.candidates_length < MAX_LENGTH and not self.isCandidate(hash)


@public
@constant
def willAddToCouncil(addr: address) -> bool:
  """
  @notice Return a bool indicating if a possible council member would be added
  @pram addr Address of the proposed council member
  """
  # TODO possible don't call the internal methods and just do the comparisons here
  return self.council_length < MAX_LENGTH and not self.inCouncil(addr)
