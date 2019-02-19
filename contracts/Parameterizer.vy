# @title Computable Parameterizer
# @notice A simplified parameterizer, hard-coded to our protocol.
# @author Computable

# We refer to any 'params' as integers to avoid conversions to and from strings
CHALLENGE_STAKE: constant(uint256) = 1
CONVERSION_RATE: constant(uint256) = 2
CONVERSION_SLOPE_DENOMINATOR: constant(uint256) = 3
CONVERSION_SLOPE_NUMERATOR: constant(uint256) = 4
LIST_REWARD: constant(uint256) = 5
QUORUM: constant(uint256) = 6
VOTE_BY: constant(uint256) = 7

# The sole Candidate 'kind' known to the Parameterizer
REPARAM: constant(uint256) = 3

struct Reparam:
  proposer: address
  param: uint256
  value: uint256

# Parameterizer has access to the Voting contract, being recognized by it as privileged
contract Voting:
  def inCouncil(member: address) -> bool: constant
  def candidateIs(hash: bytes32, kind: uint256) -> bool: constant
  def isCandidate(hash: bytes32) -> bool: constant
  def addCandidate(hash: bytes32, kind: uint256, vote_by: uint256(sec)): modifying
  def removeCandidate(hash: bytes32): modifying
  def didPass(hash: bytes32, quorum: uint256) -> bool: constant
  def pollClosed(hash: bytes32) -> bool: constant
  def willAddCandidate(hash: bytes32) -> bool: constant

ReparamProposed: event({proposer: indexed(address), hash: indexed(bytes32), param: indexed(uint256), value: uint256})
ReparamFailed: event({hash: indexed(bytes32), param: indexed(uint256), value: uint256})
ReparamSucceeded: event({hash: indexed(bytes32), param: indexed(uint256), value: uint256})

reparams: map(bytes32, Reparam)
challenge_stake: wei_value
conversion_rate: wei_value
conversion_slope_denominator: uint256
conversion_slope_numerator: uint256
list_reward: wei_value
quorum: uint256
vote_by: timedelta
voting: Voting

@public
def __init__(voting_addr: address, stake: wei_value, rate: wei_value, denominator: uint256, numerator: uint256,
  reward: wei_value, quorum_pct: uint256, vote_by_delta: timedelta):
    self.voting = voting_addr
    self.challenge_stake = stake
    self.conversion_rate = rate
    self.conversion_slope_denominator = denominator
    self.conversion_slope_numerator = numerator
    self.list_reward = reward
    self.quorum = quorum_pct
    self.vote_by = vote_by_delta


@public
@constant
def getChallengeStake() -> wei_value:
  """
  @notice Return the current challenge stake value in wei
  """
  return self.challenge_stake


@public
@constant
def getConversionRate() -> wei_value:
  """
  @notice Return the current Ethertoken to MarketToken conversion rate in wei
  """
  return self.conversion_rate


@public
@constant
def getConversionSlopeDenominator() -> uint256:
  """
  @notice Return the current conversion_slope_denominator scaling factor
  """
  return self.conversion_slope_denominator


@public
@constant
def getConversionSlopeNumerator() -> uint256:
  """
  @notice Return the current conversion_slope_numerator scaling factor
  """
  return self.conversion_slope_numerator


@public
@constant
def getListReward() -> wei_value:
  """
  @notice Return the current amount of market token, in wei, awarded to a new listing
  """
  return self.list_reward


@public
@constant
def getQuorum() -> uint256:
  """
  @notice Return the percent of 100 needed by a candidate to 'pass' in a poll
  """
  return self.quorum


@public
@constant
def getVoteBy() -> timedelta:
  """
  @notice Return the number of seconds, from when begun, that a candidate poll should close
  """
  return self.vote_by


@public
@constant
def getParamHash(reparam: string[64]) -> bytes32:
  """
  @notice Given a string of max-length 64 chars, generate a hash
  @dev Used as a key in the voting candidates mapping. Must be unique
  @param reparam A string name supplied by the user / client
  @return The generated hash
  """
  return keccak256(reparam)


@public
@constant
def getReparam(hash: bytes32) -> (address, uint256, uint256):
  """
  @notice Return the data about the given Reparam
  @param hash The Reparam identifier
  """
  return (self.reparams[hash].proposer, self.reparams[hash].param, self.reparams[hash].value)


@public
def resolveReparam(hash: bytes32):
  """
  @notice Determine if a Reparam Candidate collected enough votes to pass, setting it if so
  @dev This method enforces that the caller is a council member, the candidate is of the correct type and its poll is closed
  @param hash The Reparam identifier
  """
  assert self.voting.inCouncil(msg.sender)
  assert self.voting.candidateIs(hash, REPARAM)
  assert self.voting.pollClosed(hash)
  # ascertain which param,value we are looking at
  param: uint256 = self.reparams[hash].param
  value: uint256 = self.reparams[hash].value
  if self.voting.didPass(hash, self.quorum):
    if param == CHALLENGE_STAKE:
      self.challenge_stake = value
    elif param == CONVERSION_RATE:
      self.conversion_rate = value
    elif param == CONVERSION_SLOPE_DENOMINATOR:
      self.conversion_slope_denominator = value
    elif param == CONVERSION_SLOPE_NUMERATOR:
      self.conversion_slope_numerator = value
    elif param == LIST_REWARD:
      self.list_reward = value
    elif param == QUORUM:
      self.quorum = value
    elif param == VOTE_BY:
      self.vote_by = value
    log.ReparamSucceeded(hash, param, value)
  # did not get enough votes...
  log.ReparamFailed(hash, param, value)
  # regardless, cleanup the reparam and candidate
  self.voting.removeCandidate(hash)
  clear(self.reparams[hash])


@public
def reparameterize(reparam: string[64], param: uint256, value: uint256):
  """
  @notice Suggest a change to a Parameterizer attribute, creating a candidate for it
  @dev Sender must be in council, and there must not be a matching candidate already open
  @param reparam A string used to generate the reparam hash, which must be unique
  @param param The attribute to change
  @param value What to change it to
  """
  assert self.voting.inCouncil(msg.sender)
  hash: bytes32 = keccak256(reparam)
  assert self.voting.willAddCandidate(hash)
  self.reparams[hash] = Reparam({proposer: msg.sender, param: param, value:value})
  self.voting.addCandidate(hash, REPARAM, self.vote_by)
  log.ReparamProposed(msg.sender, hash, param, value)
