# @title Computable Market
# @notice The Computable Market Contract
# @author Computable

# constants
MAX_LENGTH: constant(uint256) = 1000000000 # no market may have more than 1B active listings
APPLICATION: constant(uint256) = 1 # market understands this candidate.kind
CHALLENGE: constant(uint256) = 2 # market also knows this candidate.kind

# structs
struct Listing:
  index: int128
  listed: bool
  owner: address
  data_hash: bytes32
  supply: wei_value
  rewards: wei_value

struct Challenge:
  challenger: address
  from_challengee_supply: wei_value
  from_challengee_rewards: wei_value

struct Investor:
  index: int128
  invested: wei_value

# external contracts
contract EtherToken:
  def balanceOf(owner: address) -> uint256(wei): constant
  def transfer(to: address, amount: uint256(wei)) -> bool: modifying
  def transferFrom(source: address, to: address, amount: uint256(wei)) -> bool: modifying

contract MarketToken:
  def balanceOf(owner: address) -> uint256(wei): constant
  def burn(amount: uint256(wei)): modifying
  def burnAll(owner: address): modifying
  def mint(amount: uint256(wei)): modifying
  def totalSupply() -> uint256(wei): constant
  def transfer(to: address, amount: uint256(wei)) -> bool: modifying
  def transferFrom(source: address, to: address, amount: uint256(wei)) -> bool: modifying

contract Voting:
  def inCouncil(member: address) -> bool: constant
  def addToCouncil(member: address): modifying
  def removeFromCouncil(member: address): modifying
  def candidateIs(hash: bytes32, kind: uint256) -> bool: constant
  def isCandidate(hash: bytes32) -> bool: constant
  def getCandidateCount() -> int128: constant
  def getCandidateKey(index: int128) -> bytes32: constant
  def addCandidate(hash: bytes32, kind: uint256, vote_by: uint256(sec)): modifying
  def removeCandidate(hash: bytes32): modifying
  def didPass(hash: bytes32, quorum: uint256) -> bool: constant
  def pollClosed(hash: bytes32) -> bool: constant
  def willAddCandidate(hash: bytes32) -> bool: constant
  def willAddToCouncil(addr: address) -> bool: constant

contract Parameterizer:
  def getChallengeStake() -> uint256(wei): constant
  def getConversionRate() -> uint256(wei): constant
  def getInvestDenominator() -> uint256: constant
  def getInvestNumerator() -> uint256: constant
  def getListReward() -> uint256(wei): constant
  def getQuorum() -> uint256: constant
  def getVoteBy() -> uint256(sec): constant

# events
ApplicationFailed: event({hash: indexed(bytes32)})
Applied: event({hash: indexed(bytes32), applicant: indexed(address), dataHash: bytes32, amount: wei_value})
Challenged: event({hash: indexed(bytes32), challenger: indexed(address)})
ChallengeFailed: event({hash: indexed(bytes32), challenger: indexed(address), reward: wei_value})
ChallengeSucceeded: event({hash: indexed(bytes32), challenger: indexed(address), reward: wei_value})
Divested: event({investor: indexed(address), transferred: wei_value})
Invested: event({investor: indexed(address), offered: wei_value, minted: wei_value})
Listed: event({hash: indexed(bytes32), owner: indexed(address), reward: wei_value})
ListingDeposit: event({hash: indexed(bytes32), owner: indexed(address), deposited: wei_value})
ListingRemoved: event({hash: indexed(bytes32)})
ListingWithdraw: event({hash: indexed(bytes32), owner: indexed(address), withdrawn: wei_value})

# state vars
listings: map(bytes32, Listing)
listings_length: int128 # the actual number of active listings
listing_keys: bytes32[MAX_LENGTH]
listing_owners: map(address, uint256) # maps makers to number of listings owned
investors: map(address, Investor)
investors_length: int128 # the actual number of investors
investor_keys: address[MAX_LENGTH] # perhaps a bit optimistic, but :shrug...
invested: wei_value # running total of all EtherToken invested in this Market
challenges: map(bytes32, Challenge)
factory_address: address
ether_token: EtherToken
market_token: MarketToken
voting: Voting
parameterizer: Parameterizer

@public
def __init__(ether_token_address: address, market_token_address: address,
  voting_address: address, parameterizer_address: address):
    self.factory_address = msg.sender
    self.ether_token = ether_token_address
    self.market_token = market_token_address
    self.voting = voting_address
    self.parameterizer = parameterizer_address


@public
@constant
def isListed(hash: bytes32) -> bool:
  """
  @notice Return a boolean representing whether a Listing has been listed
  """
  return self.listings[hash].listed == True


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
@constant
def isInvestor(addr: address) -> bool:
  """
  @notice Return a bool indicating if the given address has invested in this market
  """
  return self.investor_keys[self.investors[addr].index] == addr


@public
def depositToListing(hash: bytes32, amount: wei_value):
  """
  @notice Allow the supply of a listing to be increased
  @param hash Identifier of the listing to deposit supply to
  @param amount How much to deposit
  """
  assert self.isListing(hash)
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
  assert self.isListing(hash)
  assert self.listings[hash].owner == msg.sender
  self.listings[hash].supply -= amount
  self.market_token.transfer(msg.sender, amount)
  log.ListingWithdraw(hash, msg.sender, amount)


@public
def list(listing: string[64], data_hash: bytes32, amount: wei_value):
  """
  @notice Allows a maker to propose a new listing to a Market, creating a candidate for voting
  @dev Sender cannot be an investor, nor can the listing already exist, nor be or have ever been a candidate
  @param listing A string (max length of 64 chars) serving as a unique identifier for this Listing when hashed
  @param data_hash Magical construct which flies across the sky pooping out unicorns which, in turn, poop rainbows
  @param amount Optional funding to be deposited to the listing's supply
  """
  assert not self.isInvestor(msg.sender)
  hash: bytes32 = keccak256(listing) # not calling self method so as not to pass string around...
  # NOTE we do not need to check if hash is a current or past listing as the candidacy check disallows dupes
  # TODO implement one for sheer redundancy if desired
  assert self.voting.willAddCandidate(hash)
  self.voting.addCandidate(hash, APPLICATION, self.parameterizer.getVoteBy())
  self.listings[hash].owner = msg.sender
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
def getListing(hash: bytes32) -> (bool, address, bytes32, wei_value, wei_value):
  """
  @notice Return pertinent information about a listing
  @param hash The listing identifier
  """
  return (self.listings[hash].listed, self.listings[hash].owner, self.listings[hash].data_hash,
    self.listings[hash].supply, self.listings[hash].rewards)


@private
def removeListing(hash: bytes32):
  """
  @notice Used internally to remove both listings and failed applications
  @dev listings are never fully removed from the mapping, only indicated as removed by `index: -1` and `listed: False` (if it were ever True)
  @param hash The listing to remove
  """
  assert self.isListing(hash)
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
  if self.listings[hash].listed == True: # a failed application will not be True here
    self.listings[hash].listed = False # technically not necessary, but, seems the correct thing TODO
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
  assert self.listings[hash].listed != True
  assert self.voting.pollClosed(hash)
  # case: listing accepted
  if self.voting.didPass(hash, self.parameterizer.getQuorum()):
    self.listings[hash].listed = True
    # how much minted token is awarded (banked by the Market on the listing's behalf)
    amount: wei_value = self.parameterizer.getListReward()
    self.market_token.mint(amount)
    self.listings[hash].rewards = amount
    # currently any new listing owner becomes a council member TODO revisit when we change threshold rules
    if self.voting.willAddToCouncil(self.listings[hash].owner):
      self.voting.addToCouncil(self.listings[hash].owner)
    log.Listed(hash, self.listings[hash].owner, amount)
  else: # application did not pass vote
    self.removeListing(hash)
  # regardless, the candidate is pruned
  self.voting.removeCandidate(hash)


@public
def challenge(chall: string[64], hash: bytes32):
  """
  @notice Challenge a current listing, creating a candidate for voting if it should remain
  @dev Must actually be listed, and not already challenged. Also note that the challengee must
  have enough funds to pay the challenge stake (in supply or rewards or both)
  @param chall A string, max-length 64 chars, used to generate a hash which must be unique
  @param hash The listing identifier
  """
  assert self.listings[hash].listed
  stake: wei_value = self.parameterizer.getChallengeStake()
  # if a listing does not have the combined funds to survive a challenge, remove it.
  if (self.listings[hash].supply + self.listings[hash].rewards) < stake:
    self.removeListing(hash)
  else:
    hashed: bytes32 = keccak256(chall)
    assert self.voting.willAddCandidate(hashed)
    # TODO make it possible to stake from a listing?
    self.market_token.transferFrom(msg.sender, self, stake)
    from_supply: wei_value
    from_rewards: wei_value
    # if possible we'll lock only supply
    if self.listings[hash].supply >= stake:
      self.listings[hash].supply -= stake
      from_supply = stake
    else: # lock supply and rewards if we must
      from_supply = self.listings[hash].supply
      remainder: wei_value = stake - self.listings[hash].supply
      self.listings[hash].supply = 0
      self.listings[hash].rewards -= remainder
      from_rewards = remainder
    self.challenges[hash] = Challenge({challenger: msg.sender,
      from_challengee_supply: from_supply, from_challengee_rewards: from_rewards})
    self.voting.addCandidate(hash, CHALLENGE, self.parameterizer.getVoteBy())
    log.Challenged(hash, msg.sender)


@public
@constant
def getChallenge(hash: bytes32) -> (address, wei_value, wei_value):
  """
  @notice Return pertinent information about a given challenge
  @param hash The Challenge in question
  """
  return (self.challenges[hash].challenger, self.challenges[hash].from_challengee_supply,
    self.challenges[hash].from_challengee_rewards)


@public
def resolveChallenge(hash: bytes32):
  """
  @notice Determines the winner for a given Challenge. Rewards winner, possibly removing a listing (if appropriate)
  @dev Challenge stakes are unlocked, replinished and rewarded accordingly.
  @param hash The identifier for a Given challenge
  """
  assert self.voting.inCouncil(msg.sender)
  assert self.voting.candidateIs(hash, CHALLENGE)
  assert self.voting.pollClosed(hash)
  stake: wei_value = self.parameterizer.getChallengeStake()
  # case: challenge won
  if self.voting.didPass(hash, self.parameterizer.getQuorum()):
    self.removeListing(hash)
    # x2 as challenger gets back their own stake + the listing's
    self.market_token.transfer(self.challenges[hash].challenger, stake*2)
    log.ChallengeSucceeded(hash, self.challenges[hash].challenger, stake)
  else: # listing won
    self.listings[hash].supply += self.challenges[hash].from_challengee_supply
    self.listings[hash].rewards += (self.challenges[hash].from_challengee_rewards + stake)
    log.ChallengeFailed(hash, self.challenges[hash].challenger, stake)
  # regardless, clean up the challenge and candidate
  clear(self.challenges[hash])
  self.voting.removeCandidate(hash)


@public
@constant
def getInvested() -> wei_value:
  """
  @notice Return the total amount of Ether Token invested in this Market
  """
  return self.invested


@public
@constant
def getInvestorCount() -> int128:
  """
  @notice Return the current number of investors
  @dev returning the actual number of addresses in the investor_keys unordered list
  @return The count
  """
  return self.investors_length


@public
@constant
def getInvestment(addr: address) -> wei_value:
  """
  @notice Return the amount of EtherToken a given investor has invested in this market
  """
  return self.investors[addr].invested


@public
@constant
def getInvestmentPrice() -> wei_value:
  """
  @notice Return the amount of Ether token (in wei) needed to purchase one billionth of a Market token
  @dev WIP
  """
  rate: wei_value = self.parameterizer.getConversionRate()
  invest_d: uint256 = self.parameterizer.getInvestDenominator()
  invest_n: uint256 = self.parameterizer.getInvestNumerator()
  reserve: wei_value = self.ether_token.balanceOf(self)
  total: wei_value = self.market_token.totalSupply()
  if total < 1000000000: # that is, is total supply less than one-billionth token
    return rate + invest_n * reserve / invest_d
  else:
    return rate + (invest_n * reserve) / (invest_d * total)


@public
def invest(offer: wei_value):
  """
  @notice Allow an investor to purchase MarketToken with EtherToken priced according to the "buy-curve"
  @dev WIP
  @param offer An amount of Ether Token in Wei
  """
  assert not self.isListingOwner(msg.sender)
  price: wei_value = self.getInvestmentPrice()
  assert offer >= price # you cannot buy less than one billionth of a market token
  self.ether_token.transferFrom(msg.sender, self, offer)
  minted: uint256 = (offer / price) * 1000000000 # NOTE using wei_value here throws TypeMismatch
  self.market_token.mint(minted)
  self.market_token.transfer(msg.sender, minted)
  if not self.isInvestor(msg.sender):
    self.investors[msg.sender].index = self.investors_length
    self.investor_keys[self.investors_length] = msg.sender # push new investor into the unordered list
    self.investors_length += 1 # move the pointer
    # currently any investor is made a council member TODO revisit when we implement threshold rules
    if self.voting.willAddToCouncil(msg.sender):
      self.voting.addToCouncil(msg.sender)
  self.investors[msg.sender].invested += offer
  self.invested += offer
  log.Invested(msg.sender, offer, minted)


@public
@constant
def getDivestmentProceeds(addr: address) -> wei_value:
  """
  @notice Return the amount of Ether Token funds an investor would recieve for divestment
  @dev TBD
  @return Amount in Ether Token Wei
  """
  assert addr != ZERO_ADDRESS
  bal: wei_value = self.market_token.balanceOf(addr)
  reserve: wei_value = self.ether_token.balanceOf(self)
  total: wei_value = self.market_token.totalSupply()
  return bal * reserve / total


@public
def divest():
  """
  @notice Allows a stakeholder to exit the market. Burning any market token owned and
  withdrawing their share of the reserve.
  @dev Stakeholder may not be engaged in any current challenge.
  """
  is_challenger: bool = False
  candidates: int128 = self.voting.getCandidateCount()
  for i in range(1000): # NOTE vyper will not let us use candidates here, use Voting MAX_LENGTH instead
    k: bytes32 = self.voting.getCandidateKey(i)
    if i == candidates: # don't exceed the actual number of candidates
      break
    elif self.challenges[k].challenger == msg.sender:
      is_challenger = True
      break
  assert is_challenger == False
  divested: wei_value = self.getDivestmentProceeds(msg.sender)
  # before any transfer, burn their market tokens and remove if an investor
  self.market_token.burnAll(msg.sender)
  if self.isInvestor(msg.sender):
    self.investors_length -=1 # move the pointer back to latest entry
    if self.investors_length > 0:
      deleted: int128 = self.investors[msg.sender].index
      moved: address = self.investor_keys[self.investors_length]
      self.investor_keys[deleted] = moved
      self.investors[moved].index = deleted
      clear(self.investors[msg.sender])
      # ATM all investors are council members TODO edit when threshold rules implemented
      self.voting.removeFromCouncil(msg.sender)
    # in either case, we zero out the latest entry
    clear(self.investor_keys[self.investors_length])
  self.ether_token.transfer(msg.sender, divested)
  log.Divested(msg.sender, divested)


@public
def exit(hash: bytes32):
  """
  @notice Allows the owner of a Listing to remove it from the market
  @dev Returns all Listing Supply to the owner. Burns all minted/rewarded funds associated with the Listing.
  Must actually be listed, and must not be involved in a Challenge
  @param hash Listing identifier
  """
  assert self.listings[hash].owner == msg.sender
  assert self.listings[hash].listed == True
  assert not self.voting.isCandidate(hash)
  self.removeListing(hash)
