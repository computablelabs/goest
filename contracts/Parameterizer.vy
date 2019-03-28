# @title Computable Parameterizer
# @notice A simplified parameterizer, hard-coded to our protocol.
# @author Computable

# We refer to any 'params' as integers to avoid conversions to and from strings
CONVERSION_RATE: constant(uint256) = 2
INVEST_DENOMINATOR: constant(uint256) = 3
INVEST_NUMERATOR: constant(uint256) = 4
LIST_REWARD: constant(uint256) = 5
ACCESS_REWARD: constant(uint256) = 10
STAKE: constant(uint256) = 1
VOTE_BY: constant(uint256) = 7
QUORUM: constant(uint256) = 6
# _PAYMENTs are % based so the sum of both should be < 100 as there is an implicit reserve_payment
BACKEND_PAYMENT: constant(uint256) = 8
MAKER_PAYMENT: constant(uint256) = 9
COST_PER_BYTE: constant(uint256) = 11

# The sole Candidate 'kind' known to the Parameterizer
REPARAM: constant(uint256) = 3

struct Reparam:
  param: uint256
  value: uint256

# Parameterizer has access to the Voting contract, being recognized by it as privileged
contract Voting:
  def candidateIs(hash: bytes32, kind: uint256) -> bool: constant
  def isCandidate(hash: bytes32) -> bool: constant
  def addCandidate(hash: bytes32, kind: uint256, owner: address, stake: uint256(wei), vote_by: uint256(sec)): modifying
  def removeCandidate(hash: bytes32): modifying
  def didPass(hash: bytes32, quorum: uint256) -> bool: constant
  def pollClosed(hash: bytes32) -> bool: constant

ReparamProposed: event({owner: indexed(address), hash: indexed(bytes32), param: indexed(uint256), value: uint256})
ReparamFailed: event({hash: indexed(bytes32), param: indexed(uint256), value: uint256})
ReparamSucceeded: event({hash: indexed(bytes32), param: indexed(uint256), value: uint256})

reparams: map(bytes32, Reparam)
conversion_rate: wei_value
invest_denominator: uint256
invest_numerator: uint256
list_reward: wei_value
access_reward: wei_value
stake: wei_value
vote_by: timedelta
quorum: uint256
backend_payment: uint256
maker_payment: uint256
cost_per_byte: wei_value
voting: Voting

@public
def __init__(voting_addr: address, rate: wei_value, denominator: uint256, numerator: uint256, list_re: wei_value,
  comp_re: wei_value, stk: wei_value, vote_by_delta: timedelta, quorum_pct: uint256, back_pay: uint256, maker_pay: uint256, cost: wei_value):
    self.voting = Voting(voting_addr)
    self.conversion_rate = rate
    self.invest_denominator = denominator
    self.invest_numerator = numerator
    self.list_reward = list_re
    self.access_reward = comp_re
    self.stake = stk
    self.vote_by = vote_by_delta
    self.quorum = quorum_pct
    self.backend_payment = back_pay
    self.maker_payment = maker_pay
    self.cost_per_byte = cost


@public
@constant
def getBackendPayment() -> uint256:
  """
  @notice The current amount of market token, in wei, awarded to the backend per list usage
  @return % of 100 given to the backend
  """
  return self.backend_payment


@public
@constant
def getMakerPayment() -> uint256:
  """
  @notice The current amount of market token, in wei, awarded to the maker (list owner) per list usage
  @return % of 100 given to the listing owner
  """
  return self.maker_payment


@public
@constant
def getCostPerByte() -> wei_value:
  """
  @notice Return the amount, in wei, to be paid per output byte of listing data
  """
  return self.cost_per_byte


@public
@constant
def getStake() -> wei_value:
  """
  @notice Return the current voting/challenge stake value in wei
  """
  return self.stake


@public
@constant
def getConversionRate() -> wei_value:
  """
  @notice Return the current Ethertoken to MarketToken conversion rate in wei
  """
  return self.conversion_rate


@public
@constant
def getHash(param: uint256, value: uint256) -> bytes32:
  """
  @notice Given the param: value keypair - generate a hash
  @dev Returns the same value we use internally for constructing mapping keys
  @param param the integer representing a parameterizer member
  @param value what to set it to
  @return The generated hash
  """
  return keccak256(convert((param + value), bytes32))


@public
@constant
def getInvestDenominator() -> uint256:
  """
  @notice Return the current conversion_slope_denominator scaling factor
  """
  return self.invest_denominator


@public
@constant
def getInvestNumerator() -> uint256:
  """
  @notice Return the current conversion_slope_numerator scaling factor
  """
  return self.invest_numerator


@public
@constant
def getListReward() -> wei_value:
  """
  @notice Return the current amount of market token, in wei, awarded to a new listing
  """
  return self.list_reward

@public
@constant
def getAccessReward() -> wei_value:
  """
  @notice Return the current amount of market token, in wei, awarded to a listing when used
  """
  return self.access_reward


@public
@constant
def getQuorum() -> uint256:
  """
  @notice Return the percent of 100 needed by a candidate to 'pass' in a poll
  """
  return self.quorum


@public
@constant
def getReparam(hash: bytes32) -> (uint256, uint256):
  """
  @notice Return the data about the given Reparam
  @param hash The Reparam identifier
  """
  return (self.reparams[hash].param, self.reparams[hash].value)


@public
@constant
def getVoteBy() -> timedelta:
  """
  @notice Return the number of seconds, from when begun, that a candidate poll should close
  """
  return self.vote_by


@public
def reparameterize(param: uint256, value: uint256):
  """
  @notice Suggest a change to a Parameterizer attribute, creating a candidate for it
  @dev Sender must not be a matching candidate already open
  @param param The attribute to change
  @param value What to change it to
  """
  # hashed identifier made up of the prop and its proposed value
  hash: bytes32 = keccak256(convert((param + value), bytes32)) # TODO may not need to SHA this
  assert not self.voting.isCandidate(hash)
  self.reparams[hash] = Reparam({param: param, value:value})
  self.voting.addCandidate(hash, REPARAM, msg.sender, self.stake, self.vote_by)
  log.ReparamProposed(msg.sender, hash, param, value)


@public
def resolveReparam(hash: bytes32):
  """
  @notice Determine if a Reparam Candidate collected enough votes to pass, setting it if so
  @dev This method enforces that the candidate is of the correct type and its poll is closed
  @param hash The Reparam identifier
  """
  assert self.voting.candidateIs(hash, REPARAM)
  assert self.voting.pollClosed(hash)
  # ascertain which param,value we are looking at
  param: uint256 = self.reparams[hash].param
  value: uint256 = self.reparams[hash].value
  if self.voting.didPass(hash, self.quorum):
    #TODO in time we can likely tell an optimal order for these...
    if param == CONVERSION_RATE:
      self.conversion_rate = value
    elif param == INVEST_DENOMINATOR:
      self.invest_denominator = value
    elif param == INVEST_NUMERATOR:
      self.invest_numerator = value
    elif param == LIST_REWARD:
      self.list_reward = value
    elif param == ACCESS_REWARD:
      self.access_reward = value
    elif param == STAKE:
      self.stake = value
    elif param == VOTE_BY:
      self.vote_by = value
    elif param == QUORUM:
      self.quorum = value
    elif param == MAKER_PAYMENT:
      self.maker_payment = value
    elif param == BACKEND_PAYMENT:
      self.backend_payment = value
    elif param == COST_PER_BYTE:
      self.cost_per_byte = value
    log.ReparamSucceeded(hash, param, value)
  # did not get enough votes...
  log.ReparamFailed(hash, param, value)
  # regardless, cleanup the reparam and candidate
  self.voting.removeCandidate(hash)
  # TODO make sure this works as expected
  clear(self.reparams[hash])
