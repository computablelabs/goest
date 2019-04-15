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

contract Parameterizer:
  def getConversionRate() -> uint256(wei): constant
  def getSpread() -> uint256: constant

# events
Divested: event({investor: indexed(address), transferred: wei_value})
Invested: event({investor: indexed(address), offered: wei_value, minted: wei_value})

# state vars
ether_token: EtherToken
market_token: MarketToken
parameterizer: Parameterizer

@public
def __init__(ether_token_addr: address, market_token_addr: address, p11r_addr: address):
    self.ether_token = EtherToken(ether_token_addr)
    self.market_token = MarketToken(market_token_addr)
    self.parameterizer = Parameterizer(p11r_addr)


@public
@constant
def getInvestmentPrice() -> wei_value:
  """
  @notice Return the amount of Ether token (in wei) needed to purchase one billionth of a Market token
  """
  rate: wei_value = self.parameterizer.getConversionRate()
  spread: uint256 = self.parameterizer.getSpread()
  reserve: wei_value = self.ether_token.balanceOf(self)
  total: wei_value = self.market_token.totalSupply()
  if total < 1000000000000000000: # that is, is total supply less than one token in wei
    return rate + ((spread * reserve) / 100)
  else:
    return rate + ((spread * reserve * 1000000000000000000) / (100 * total))


@public
def invest(offer: wei_value):
  """
  @notice Allow an investor to purchase MarketToken with EtherToken priced according to the "buy-curve"
  @param offer An amount of Ether Token in Wei
  """
  price: wei_value = self.getInvestmentPrice()
  assert offer >= price # you cannot buy less than one billionth of a market token
  self.ether_token.transferFrom(msg.sender, self, offer)
  minted: uint256 = (offer / price) * 1000000000 # NOTE using wei_value here throws TypeMismatch
  self.market_token.mint(minted)
  self.market_token.transfer(msg.sender, minted)
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
  return (bal * reserve) / total


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
  self.ether_token.transfer(msg.sender, divested)
  log.Divested(msg.sender, divested)
