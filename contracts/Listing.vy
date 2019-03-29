# @title Computable Listing
# @notice Handle the details pertaining to a Listing in a Computable Market
# @author Computable

# constants
APPLICATION: constant(uint256) = 1 # candidate.kind
CHALLENGE: constant(uint256) = 2 # candidate.kind

# structs
struct Listing:
  owner: address
  supply: wei_value

contract MarketToken:
  def burn(amount: uint256(wei)): modifying
  def mint(amount: uint256(wei)): modifying
  def transfer(to: address, amount: uint256(wei)) -> bool: modifying
  def transferFrom(source: address, to: address, amount: uint256(wei)) -> bool: modifying

contract Voting:
  def candidateIs(hash: bytes32, kind: uint256) -> bool: constant
  def isCandidate(hash: bytes32) -> bool: constant
  def getCandidateOwner(hash: bytes32) -> address: constant
  def addCandidate(hash: bytes32, kind: uint256, owner: address, stake: uint256(wei), vote_by: uint256(sec)): modifying
  def removeCandidate(hash: bytes32): modifying
  def didPass(hash: bytes32, quorum: uint256) -> bool: constant
  def pollClosed(hash: bytes32) -> bool: constant
  def transferStake(hash: bytes32, addr: address): modifying

contract Parameterizer:
  def getStake() -> uint256(wei): constant
  def getListReward() -> uint256(wei): constant
  def getQuorum() -> uint256: constant
  def getVoteBy() -> uint256(sec): constant

contract Datatrust:
  def getDataHash(hash: bytes32) -> bytes32: constant
  def removeDataHash(hash: bytes32): modifying

# events
ApplicationFailed: event({hash: indexed(bytes32), applicant: indexed(address)})
Applied: event({hash: indexed(bytes32), applicant: indexed(address)})
Challenged: event({hash: indexed(bytes32), challenger: indexed(address)})
ChallengeFailed: event({hash: indexed(bytes32), challenger: indexed(address)})
ChallengeSucceeded: event({hash: indexed(bytes32), challenger: indexed(address)})
Listed: event({hash: indexed(bytes32), owner: indexed(address), reward: wei_value})
ListingRemoved: event({hash: indexed(bytes32)})
ListingWithdraw: event({hash: indexed(bytes32), owner: indexed(address), withdrawn: wei_value})

# state vars
listings: map(bytes32, Listing)
market_token: MarketToken
voting: Voting
parameterizer: Parameterizer
datatrust: Datatrust

@public
def __init__(market_token_addr: address, voting_addr: address, p11r_addr: address, data_addr: address):
    self.market_token = MarketToken(market_token_addr)
    self.voting = Voting(voting_addr)
    self.parameterizer = Parameterizer(p11r_addr)
    self.datatrust = Datatrust(data_addr)


@public
@constant
def isListed(hash: bytes32) -> bool:
  """
  @notice Return a boolean representing whether a Listing has been listed
  """
  return self.listings[hash].owner != ZERO_ADDRESS


@public
def withdrawFromListing(hash: bytes32, amount: wei_value):
  """
  @notice Allow a maker to witdhraw from the supply of an owned listing
  @param hash The listing identifier
  @param amount The funds to withdraw
  """
  assert self.listings[hash].owner == msg.sender
  self.listings[hash].supply -= amount
  self.market_token.transfer(msg.sender, amount)
  log.ListingWithdraw(hash, msg.sender, amount)


@public
def list(listing: string[64]):
  """
  @notice Allows a maker to propose a new listing to a Market, creating a candidate for voting
  @dev Listing cannot already exist, in any active form. Owner not set here as it's an application
  @param listing A string (max length of 64 chars) serving as a unique identifier for this Listing when hashed
  """
  hash: bytes32 = keccak256(listing)
  assert not self.voting.isCandidate(hash) # not an applicant
  assert self.listings[hash].owner == ZERO_ADDRESS # not already listed
  self.voting.addCandidate(hash, APPLICATION, msg.sender, self.parameterizer.getStake(), self.parameterizer.getVoteBy())
  log.Applied(hash, msg.sender)


@public
@constant
def getHash(str: string[64]) -> bytes32:
  """
  @notice Return the same hashed value, given a string (max length 64 chars), that we generate internally when listing
  """
  return keccak256(str)


@public
@constant
def getListing(hash: bytes32) -> (address, wei_value):
  """
  @notice Return pertinent information about a listing
  @param hash The listing identifier
  """
  return (self.listings[hash].owner, self.listings[hash].supply)


@private
def removeListing(hash: bytes32):
  """
  @notice Used internally to remove listings
  @dev We clear the members of a struct pointed to by the listing hash (enabling re-use)
  @param hash The listing to remove
  """
  assert self.listings[hash].owner != ZERO_ADDRESS
  if self.listings[hash].supply > 0:
    self.market_token.burn(self.listings[hash].supply)
    clear(self.listings[hash].supply)
  clear(self.listings[hash]) # TODO assure we don't need to do this by hand
  # datatrust now needs to clear the data hash
  self.datatrust.removeDataHash(hash)
  log.ListingRemoved(hash)


@public
def resolveApplication(hash: bytes32):
  """
  @notice Decide if an application becomes a listing.
  @dev Added as a listing if passing vote. Candidate is cleared regardless. Datatrust data_hash for this listing must be set.
  @param hash The identifier for said applicant
  """
  assert self.voting.candidateIs(hash, APPLICATION)
  assert self.voting.pollClosed(hash)
  data_hash: bytes32 = self.datatrust.getDataHash(hash)
  owner: address = self.voting.getCandidateOwner(hash)
  # case: listing accepted
  if data_hash != EMPTY_BYTES32:
    if self.voting.didPass(hash, self.parameterizer.getQuorum()):
      self.listings[hash].owner = owner # is now 'listed'
      amount: wei_value = self.parameterizer.getListReward()
      self.market_token.mint(amount)
      self.listings[hash].supply = amount
      log.Listed(hash, owner, amount)
    else: # we have a data_hash but vote didn't pass - remove it
      self.datatrust.removeDataHash(hash)
      log.ApplicationFailed(hash, owner)
  else: # application did not pass vote. there is no listing to remove
    log.ApplicationFailed(hash, owner)
  # regardless, the candidate is pruned
  self.voting.removeCandidate(hash)


@public
def challenge(hash: bytes32):
  """
  @notice Challenge a current listing, creating a candidate for voting if it should remain
  @dev Must actually be listed, and not already challenged.
  @param hash The identifier for the listing being challenged
  """
  assert self.listings[hash].owner != ZERO_ADDRESS
  assert not self.voting.isCandidate(hash) # assure not already challenged
  self.voting.addCandidate(hash, CHALLENGE, msg.sender, self.parameterizer.getStake(), self.parameterizer.getVoteBy())
  log.Challenged(hash, msg.sender)


@public
def resolveChallenge(hash: bytes32):
  """
  @notice Determines the winner for a given Challenge. Rewards winner, possibly removing a listing (if appropriate)
  @dev Challenge stakes are unlocked, replinished and rewarded accordingly. Note that the ability to fund the stake
  takes precedence over voting
  @param hash The identifier for a Given challenge
  """
  assert self.voting.candidateIs(hash, CHALLENGE)
  assert self.voting.pollClosed(hash)
  owner: address = self.voting.getCandidateOwner(hash) # TODO this could likely be removed now
  # Case: challenge won
  if self.voting.didPass(hash, self.parameterizer.getQuorum()):
    self.removeListing(hash)
    log.ChallengeSucceeded(hash, owner)
  else: # Case: listing won
    # the voting contract will make the stake avail (via unstake) to the list owner
    self.voting.transferStake(hash, self.listings[hash].owner)
    log.ChallengeFailed(hash, owner)
  # regardless, clean up the candidate
  self.voting.removeCandidate(hash)


@public
def exit(hash: bytes32):
  """
  @notice Allows the owner of a Listing to remove it from the market
  @dev Returns all Listing Supply to the owner. Burns all minted/rewarded funds associated with the Listing.
  Must actually be listed, and must not be involved in a Challenge
  @param hash Listing identifier
  """
  assert self.listings[hash].owner == msg.sender
  assert not self.voting.isCandidate(hash)
  self.removeListing(hash)
