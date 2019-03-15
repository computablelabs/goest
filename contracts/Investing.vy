# @title Computable Investing
# @notice Handle the details of Investing in the Computable Market
# @author Computable

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
  def inCouncil(member: address) -> bool: constant
  def addToCouncil(member: address): modifying
  def removeFromCouncil(member: address): modifying
  def getCandidateCount() -> int128: constant
  def getCandidateKey(index: int128) -> bytes32: constant

contract Parameterizer:
  def getConversionRate() -> uint256(wei): constant
  def getInvestDenominator() -> uint256: constant
  def getInvestNumerator() -> uint256: constant

# events
Divested: event({investor: indexed(address), transferred: wei_value})
Invested: event({investor: indexed(address), offered: wei_value, minted: wei_value})

# state vars
investors: map(address, wei_value)
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
def getInvested() -> wei_value:
  """
  @notice Return the total amount of Ether Token invested in this Market
  """
  return self.invested


@public
@constant
def getInvestment(addr: address) -> wei_value:
  """
  @notice Return the amount of EtherToken a given investor has invested in this market
  """
  return self.investors[addr]


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
  self.investors[msg.sender] += offer
  self.invested += offer
  # currently any investor is made a council member TODO revisit when we implement threshold rules
  if not self.voting.inCouncil(msg.sender):
    self.voting.addToCouncil(msg.sender)
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
  @dev Stakeholder, if owning a challenge, may want to wait until that is over (in case they win)
  """
  divested: wei_value = self.getDivestmentProceeds(msg.sender)
  assert divested > 0
  # before any transfer, burn their market tokens and remove if an investor
  self.market_token.burnAll(msg.sender)
  if self.investors[msg.sender] > 0:
    clear(self.investors[msg.sender])
    # ATM all investors are council members TODO edit when threshold rules implemented
    self.voting.removeFromCouncil(msg.sender)
  self.ether_token.transfer(msg.sender, divested)
  log.Divested(msg.sender, divested)
