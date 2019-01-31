# @dev an experimental contract to toy with some math

user_market_token_balance: wei_value
conversion_rate: wei_value # we likely could create a custom gwei_value type
reserve: wei_value
slope_numerator: uint256
slope_denominator: uint256
slope: decimal # could work as a decimal? unsure of how it would be dynamic once deployed
total_supply: wei_value

@public
def __init__():
    # We'll set some defaults here, and just use setters to change around in spec
    self.user_market_token_balance = 0
    # are we sticking to a conversion rate that is 1 gwei?
    self.conversion_rate = as_wei_value(1, 'gwei')
    self.reserve = 0
    self.slope_numerator = 100
    self.slope_denominator = 101
    # self.slope = 0.5 # hypothetical slope as a decimal?
    self.total_supply = as_wei_value(2, 'ether')


# interesting that public methods will still need to be camelCased to fit with eth style,
# private methods could use snake_casing... something to consider
@public
def setUserMarketTokenBalance(new_balance: wei_value):
    self.user_market_token_balance = new_balance


@public
def setConversionRate(new_rate: wei_value):
    self.conversion_rate = new_rate


@public
def setReserve(new_reserve: wei_value):
    self.reserve = new_reserve

@public
@constant
def getReserve() -> wei_value:
    return self.reserve

@public
def setSlopeNumerator(new_n: uint256):
    self.slope_numerator = new_n


@public
def setSlopeDenominator(new_d: uint256):
    self.slope_denominator = new_d

# if the slope were a decimal...
#  @public
#  def setSlope(new_slope: decimal):
    #  self.slope = new_slope


@public
def setTotalSupply(new_supply: wei_value):
    self.total_supply = new_supply


# we can just return the values these generate to peek at what is happening cuz i still don't grok this
@public
@constant
def getInvestmentPrice() -> wei_value:
    # this is what is in-place in market.sol
    # return rate.add(slopeN.mul(reserve).div(slopeD.mul(1e9)));

    # TODO may not need the parens...
    return self.conversion_rate + ((self.slope_numerator * self.reserve) / (self.slope_denominator * 1000000000))


@public
@constant
def getDivestmentProceeds() -> wei_value:
    return self.user_market_token_balance * self.reserve / self.total_supply

