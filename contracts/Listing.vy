# @title Computable Listing
# @notice Handle the details pertaining to a Listing in a Computable Market
# @author Computable

# constants
MAX_LENGTH: constant(uint256) = 1000000 # no market may have more than 1M active listings
APPLICATION: constant(uint256) = 1 # market understands this candidate.kind
CHALLENGE: constant(uint256) = 2 # market also knows this candidate.kind

# structs
struct Listing:
  index: int128
  owner: address
  data_hash: bytes32
  supply: wei_value
  rewards: wei_value

contract MarketToken:
  def burn(amount: uint256(wei)): modifying
  def mint(amount: uint256(wei)): modifying
  def transfer(to: address, amount: uint256(wei)) -> bool: modifying
  def transferFrom(source: address, to: address, amount: uint256(wei)) -> bool: modifying

contract Voting:
  def inCouncil(member: address) -> bool: constant
  def addToCouncil(member: address): modifying
  def removeFromCouncil(member: address): modifying
  def candidateIs(hash: bytes32, kind: uint256) -> bool: constant
  def isCandidate(hash: bytes32) -> bool: constant
  def getCandidateOwner(hash: bytes32) -> address: constant
  def addCandidate(hash: bytes32, kind: uint256, owner: address, vote_by: uint256(sec)): modifying
  def removeCandidate(hash: bytes32): modifying
  def didPass(hash: bytes32, quorum: uint256) -> bool: constant
  def pollClosed(hash: bytes32) -> bool: constant
  def willAddCandidate(hash: bytes32) -> bool: constant
  def willAddToCouncil(addr: address) -> bool: constant

contract Parameterizer:
  def getChallengeStake() -> uint256(wei): constant
  def getListReward() -> uint256(wei): constant
  def getQuorum() -> uint256: constant
  def getVoteBy() -> uint256(sec): constant

# events
ApplicationFailed: event({hash: indexed(bytes32)})
Applied: event({hash: indexed(bytes32), applicant: indexed(address), dataHash: bytes32, amount: wei_value})
Challenged: event({hash: indexed(bytes32), challenger: indexed(address)})
ChallengeFailed: event({hash: indexed(bytes32), challenger: indexed(address), reward: wei_value})
ChallengeSucceeded: event({hash: indexed(bytes32), challenger: indexed(address), reward: wei_value})
Listed: event({hash: indexed(bytes32), owner: indexed(address), reward: wei_value})
ListingConverted: event({ hash: indexed(bytes32)})
ListingDeposit: event({hash: indexed(bytes32), owner: indexed(address), deposited: wei_value})
ListingRemoved: event({hash: indexed(bytes32)})
ListingWithdraw: event({hash: indexed(bytes32), owner: indexed(address), withdrawn: wei_value})

# state vars
listings: map(bytes32, Listing)
listings_length: int128 # the actual number of active listings
listing_keys: bytes32[MAX_LENGTH]
listing_owners: map(address, uint256) # maps makers to number of listings owned
factory_address: address
market_token: MarketToken
voting: Voting
parameterizer: Parameterizer

@public
def __init__(market_token_addr: address, voting_addr: address, p11r_addr: address):
    self.factory_address = msg.sender
    self.market_token = MarketToken(market_token_addr)
    self.voting = Voting(voting_addr)
    self.parameterizer = Parameterizer(p11r_addr)


@public
@constant
def isListed(hash: bytes32) -> bool:
  """
  @notice Return a boolean representing whether a Listing has been listed
  @dev We don't store the owner until an applicatication passes voting (if it does)
  """
  return self.listings[hash].owner != ZERO_ADDRESS


@public
@constant
def isListing(hash: bytes32) -> bool:
  """
  @notice Returns a bool telling if a given hash is an active listing, but not necessarily `listed`
  """
  return self.listing_keys[self.listings[hash].index] == hash


@public
@constant
def wasListing(hash: bytes32) -> bool:
  """
  @notice Return true if a given hash points to a removed listing
  @dev An non-active "removed" listing is one whose index is -1
  @return bool
  """
  return self.listings[hash].index == -1


@public
@constant
def isListingOwner(addr: address) -> bool:
  """
  @notice Return a bool indicating if the given address owns any listings
  """
  return self.listing_owners[addr] >= 1


@public
def depositToListing(hash: bytes32, amount: wei_value):
  """
  @notice Allow the supply of a listing to be increased
  @param hash Identifier of the listing to deposit supply to
  @param amount How much to deposit
  """
  assert self.listing_keys[self.listings[hash].index] == hash # isListing
  self.market_token.transferFrom(msg.sender, self, amount)
  self.listings[hash].supply += amount
  log.ListingDeposit(hash, msg.sender, amount)


@public
def withdrawFromListing(hash: bytes32, amount: wei_value):
  """
  @notice Allow a maker to witdhraw from the supply of an owned listing
  @dev Funds may come from supply only
  @param hash The listing identifier
  @param amount The funds to withdraw
  """
  assert self.listing_keys[self.listings[hash].index] == hash # isListing
  assert self.listings[hash].owner == msg.sender
  self.listings[hash].supply -= amount
  self.market_token.transfer(msg.sender, amount)
  log.ListingWithdraw(hash, msg.sender, amount)


@public
def list(listing: string[64], data_hash: bytes32, amount: wei_value):
  """
  @notice Allows a maker to propose a new listing to a Market, creating a candidate for voting
  @dev We cannot already be at listings MAX_LENGTH. Listing cannot already exist, nor be or have ever been a candidate
  @param listing A string (max length of 64 chars) serving as a unique identifier for this Listing when hashed
  @param data_hash Magical construct which flies across the sky pooping out unicorns which, in turn, poop rainbows
  @param amount Optional funding to be deposited to the listing's supply
  """
  assert self.listings_length < MAX_LENGTH
  hash: bytes32 = keccak256(listing) # not calling self method so as not to pass string around...
  # NOTE we do not need to check if hash is a current or past listing as the candidacy check disallows dupes
  # TODO implement one for sheer redundancy if desired
  assert self.voting.willAddCandidate(hash)
  self.voting.addCandidate(hash, APPLICATION, msg.sender, self.parameterizer.getVoteBy())
  self.listings[hash].data_hash = data_hash
  self.listings[hash].index = self.listings_length
  # "push" the new listing into the unordered list
  self.listing_keys[self.listings_length] = hash
  # move the pointer to an empty slot
  self.listings_length += 1
  self.listing_owners[msg.sender] += 1
  if amount > 0:
    self.market_token.transferFrom(msg.sender, self, amount)
    self.listings[hash].supply = amount
  log.Applied(hash, msg.sender, data_hash, amount)


@public
@constant
def getListingCount() -> int128:
  """
  @notice Return the current number of Listings
  @dev this is both Listings and Applicants (but not removed listings)
  @return The count
  """
  return self.listings_length


@public
@constant
def getListingKey(index: int128) -> bytes32:
  """
  @notice Fetch the hashed listing identifier for a given inhex in our listing unordered list
  """
  return self.listing_keys[index]


@public
@constant
def getListing(hash: bytes32) -> (address, bytes32, wei_value, wei_value):
  """
  @notice Return pertinent information about a listing
  @param hash The listing identifier
  """
  return (self.listings[hash].owner, self.listings[hash].data_hash,
    self.listings[hash].supply, self.listings[hash].rewards)


@public
def convertListing(hash:bytes32):
  """
  @notice Allow a Listing owner to claim their rewards (and any supply present) by turning over ownership to the Market
  @param hash The listing identifier
  """
  assert self.listings[hash].owner == msg.sender
  funds: wei_value = self.listings[hash].supply + self.listings[hash].rewards
  if funds > 0:
    self.listings[hash].supply = 0
    self.listings[hash].rewards = 0
    self.market_token.transfer(self.listings[hash].owner, funds)
  self.listings[hash].owner = self
  log.ListingConverted(hash)


@private
def removeListing(hash: bytes32):
  """
  @notice Used internally to remove both listings and failed applications
  @dev listings are never fully removed from the mapping, only indicated as removed by `index: -1` (other fields are cleared)
  @param hash The listing to remove
  """
  assert self.listing_keys[self.listings[hash].index] == hash # isListing
  if self.listings[hash].supply > 0:
    self.market_token.transfer(self.listings[hash].owner, self.listings[hash].supply)
    clear(self.listings[hash].supply)
  if self.listings[hash].rewards > 0:
    self.market_token.burn(self.listings[hash].rewards)
    clear(self.listings[hash].rewards)
  # move the pointer back to the latest entry
  self.listings_length -= 1
  if self.listings_length > 0:
    deleted: int128 = self.listings[hash].index # target being overwritten
    moved: bytes32 = self.listing_keys[self.listings_length] # target being used to overwrite with
    self.listing_keys[deleted] = moved
    self.listings[moved].index = deleted # update index of the moved target
  clear(self.listing_keys[self.listings_length]) # zero out what is now a dupe
  # regardless, this owner has one less listing now
  self.listing_owners[self.listings[hash].owner] -= 1
  # possibly remove from council TODO revisit when we change council threshold rules
  if self.listing_owners[self.listings[hash].owner] < 1:
    self.voting.removeFromCouncil(self.listings[hash].owner)
  # finally indicate the removed listing's status
  self.listings[hash].index = -1
  clear(self.listings[hash].owner) # inactive listings are not listed
  clear(self.listings[hash].data_hash) # meaningless on an inactive listing
  log.ListingRemoved(hash)


@public
def resolveApplication(hash: bytes32):
  """
  @notice Decide if an application becomes a listing.
  @dev Either lists or removes the application. Regardless the voting candidate is removed
  @param hash The identifier for said applicant
  """
  assert self.voting.inCouncil(msg.sender)
  assert self.voting.candidateIs(hash, APPLICATION)
  assert self.listings[hash].owner == ZERO_ADDRESS # not yet listed
  assert self.voting.pollClosed(hash)
  # case: listing accepted
  if self.voting.didPass(hash, self.parameterizer.getQuorum()):
    owner: address = self.voting.getCandidateOwner(hash)
    self.listings[hash].owner = owner
    # how much minted token is awarded (banked by the Market on the listing's behalf)
    amount: wei_value = self.parameterizer.getListReward()
    self.market_token.mint(amount)
    self.listings[hash].rewards = amount
    # currently any new listing owner becomes a council member TODO revisit when we change threshold rules
    if self.voting.willAddToCouncil(owner):
      self.voting.addToCouncil(owner)
    log.Listed(hash, owner, amount)
  else: # application did not pass vote
    self.removeListing(hash)
  # regardless, the candidate is pruned
  self.voting.removeCandidate(hash)


@public
def challenge(chall: string[64], hash: bytes32):
  """
  @notice Challenge a current listing, creating a candidate for voting if it should remain
  @dev Must actually be listed, and not already challenged.
  @param chall A string, max-length 64 chars, used to generate a hash which must be unique
  @param hash The listing identifier
  """
  assert self.listings[hash].owner != ZERO_ADDRESS
  stake: wei_value = self.parameterizer.getChallengeStake()
  hashed: bytes32 = keccak256(chall)
  assert self.voting.willAddCandidate(hashed)
  # TODO make it possible to stake from a listing?
  self.market_token.transferFrom(msg.sender, self, stake)
  self.voting.addCandidate(hash, CHALLENGE, msg.sender, self.parameterizer.getVoteBy())
  log.Challenged(hash, msg.sender)


@public
def resolveChallenge(hash: bytes32):
  """
  @notice Determines the winner for a given Challenge. Rewards winner, possibly removing a listing (if appropriate)
  @dev Challenge stakes are unlocked, replinished and rewarded accordingly. Note that the ability to fund the stake
  takes precedence over voting
  @param hash The identifier for a Given challenge
  """
  assert self.voting.inCouncil(msg.sender)
  assert self.voting.candidateIs(hash, CHALLENGE)
  assert self.voting.pollClosed(hash)
  owner: address = self.voting.getCandidateOwner(hash)
  stake: wei_value = self.parameterizer.getChallengeStake()
  # before we tally votes, we check that the challengee has the funds to stake with
  supply: wei_value = self.listings[hash].supply
  rewards: wei_value = self.listings[hash].rewards
  if (supply + rewards) < stake:
    # we allow a challenger to recieve whatever balances a listing may have had, if present
    if (supply + rewards) > 0:
      self.listings[hash].supply = 0
      self.listings[hash].rewards = 0
      self.market_token.transfer(owner, (supply + rewards))
    self.removeListing(hash)
  else:
    # Now we can check voting. Case: challenge won
    if self.voting.didPass(hash, self.parameterizer.getQuorum()):
      # adjust the supply and rewards depending on where we were able to assemble the stake from
      if supply >= stake:
        self.listings[hash].supply -= stake
      else:
        self.listings[hash].supply = 0
        self.listings[hash].rewards -= (stake - supply)
      # x2 as challenger gets back their own stake + the listing's
      self.market_token.transfer(owner, stake*2)
      self.removeListing(hash)
      log.ChallengeSucceeded(hash, owner, stake)
    else: # Case: listing won
      self.listings[hash].rewards += stake
      log.ChallengeFailed(hash, owner, stake)
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
  assert not self.voting.isCandidate(hash) # not challenged
  self.removeListing(hash)
