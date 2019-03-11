# @title Computable Investing
# @notice Handle the details of Investing in the Computable Market
# @author Computable

# constants
MAX_LENGTH: constant(uint256) = 100000 # no market may have more than 100k active investors

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
  def burnAll(owner: address): modifying
  def mint(amount: uint256(wei)): modifying
  def totalSupply() -> uint256(wei): constant
  def transfer(to: address, amount: uint256(wei)) -> bool: modifying

contract Voting:
  def addToCouncil(member: address): modifying
  def removeFromCouncil(member: address): modifying
  def getCandidateCount() -> int128: constant
  def getCandidateKey(index: int128) -> bytes32: constant
  def getCandidateOwner(hash: bytes32) -> address: constant
  def willAddToCouncil(addr: address) -> bool: constant

contract Parameterizer:
  def getConversionRate() -> uint256(wei): constant
  def getInvestDenominator() -> uint256: constant
  def getInvestNumerator() -> uint256: constant

# events
Divested: event({investor: indexed(address), transferred: wei_value})
Invested: event({investor: indexed(address), offered: wei_value, minted: wei_value})

# state vars
investors: map(address, Investor)
investors_length: int128 # the actual number of investors
investor_keys: address[MAX_LENGTH] # perhaps a bit optimistic, but :shrug...
invested: wei_value # running total of all EtherToken invested in this Market
factory_address: address
ether_token: EtherToken
market_token: MarketToken
voting: Voting
parameterizer: Parameterizer

@public
def __init__(ether_token_addr: address, market_token_addr: address,
  voting_addr: address, p11r_addr: address):
    self.factory_address = msg.sender
    self.ether_token = EtherToken(ether_token_addr)
    self.market_token = MarketToken(market_token_addr)
    self.voting = Voting(voting_addr)
    self.parameterizer = Parameterizer(p11r_addr)


@public
@constant
def isInvestor(addr: address) -> bool:
  """
  @notice Return a bool indicating if the given address has invested in this market
  """
  return self.investor_keys[self.investors[addr].index] == addr


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
  if total < 1000000000000000000: # that is, is total supply less than one token in wei
    return rate + invest_n * reserve / invest_d
  else:
    return rate + (invest_n * reserve * 1000000000000000000) / (invest_d * total)


@public
def invest(offer: wei_value):
  """
  @notice Allow an investor to purchase MarketToken with EtherToken priced according to the "buy-curve"
  @dev WIP
  @param offer An amount of Ether Token in Wei
  """
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
  @dev Stakeholder may not be engaged in any currently active poll (cannot own a candidate)
  """
  in_poll: bool = False
  candidates: int128 = self.voting.getCandidateCount()
  for i in range(1000): # NOTE vyper will not let us use candidates here, use Voting MAX_LENGTH instead
    hash: bytes32 = self.voting.getCandidateKey(i)
    owner: address = self.voting.getCandidateOwner(hash)
    if i == candidates: # don't exceed the actual number of candidates
      break
    elif owner == msg.sender:
      in_poll = True
      break
  assert in_poll == False
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
