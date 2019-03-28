# @title Computable Datatrust
# @notice Handle the details pertaining to Computable Datatrust functionality
# @author Computable

# constants
REGISTRATION: constant(uint256) = 4 # candidate.kind

# external contracts
contract EtherToken:
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

contract Parameterizer:
  def getBackendPayment() -> uint256: constant
  def getCostPerByte() -> wei_value: constant
  def getStake() -> uint256(wei): constant
  def getQuorum() -> uint256: constant
  def getVoteBy() -> uint256(sec): constant

# events
Registered: event({hash: indexed(bytes32), registrant: indexed(address)})
RegistrationSucceeded: event({hash: indexed(bytes32), registrant: indexed(address)})
RegistrationFailed: event({hash: indexed(bytes32), registrant: indexed(address)})

# state vars
data_hashes: map(bytes32, bytes32) # listing_hash -> data_hash
byte_credits: map(address, wei_value) # maps user-address to byte-credits
backend_url: string[128]
backend_address: address
ether_token: EtherToken
voting: Voting
parameterizer: Parameterizer
factory_address: address
listing_address: address

@public
def __init__(ether_token_addr: address, voting_addr: address, p11r_addr: address):
  self.factory_address = msg.sender
  self.ether_token = EtherToken(ether_token_addr)
  self.voting = Voting(voting_addr)
  self.parameterizer = Parameterizer(p11r_addr)


@public
@constant
def getPrivileged() -> address:
  """
  @notice Fetch a list of each privileged address recognized by this contract
  @return Listing contract address
  """
  return self.listing_address


@public
def setPrivileged(listing: address):
  """
  @notice Allow the Market Factory to set privileged contract addresses
  """
  assert msg.sender == self.factory_address
  self.listing_address = listing


@public
@constant
def getHash(url: string[128]) -> bytes32:
  """
  @notice Return the same hashed value, given a url string (max length 128 chars), that we generate internally when registering
  """
  return keccak256(url)


@public
@constant
def getBackendAddress() -> address:
  """
  @notice Return the address of the currently registered backend
  """
  return self.backend_address


@public
@constant
def getBackendUrl() -> string[128]:
  """
  @notice Return the URL of the currently registered backend
  """
  return self.backend_url


@public
def setBackendUrl(url: string[128]):
  """
  @notice Allow a registered backend to change its URL
  @param str The URL string
  """
  assert msg.sender == self.backend_address
  self.backend_url = url


@public
@constant
def getDataHash(hash: bytes32) -> bytes32:
  """
  @notice Return the data_hash for a given listing_hash if a backend has reported it
  @param hash The hashed name of a Listing whose data_hash we may possess
  """
  return self.data_hashes[hash]


@public
def setDataHash(listing: bytes32, data: bytes32):
  """
  @notice Allow a registered backend to set the data_hash for a given listing
  @param listing The hashed identifier of a current market listing
  @param data The hashed data held by the backend for said listing
  """
  assert msg.sender == self.backend_address
  self.data_hashes[listing] = data

@public
def removeDataHash(hash: bytes32):
  """
  @notice Allow the listing contract to call for the removal of a data_hash
  whose listing has been removed
  @dev Restricted to the Listing contract
  @param hash The Listing whose data_hash to clear
  """
  assert msg.sender == self.listing_address
  clear(self.data_hashes[hash])


@public
def register(url: string[128]):
  """
  @notice Allow a backend to register as a candidate
  @param url The location of this backend
  """
  assert msg.sender != self.backend_address # don't register 2x
  self.backend_url = url # we'll clear this if the registration fails
  hash: bytes32 = keccak256(url)
  assert not self.voting.isCandidate(hash)
  self.voting.addCandidate(hash, REGISTRATION, msg.sender, self.parameterizer.getStake(), self.parameterizer.getVoteBy())
  log.Registered(hash, msg.sender)


@public
def resolveRegistration(hash: bytes32):
  """
  @notice Set internal backend in use if approved (remove if not)
  @param hash The hashed string url of the backend candidate
  """
  assert self.voting.candidateIs(hash, REGISTRATION)
  assert self.voting.pollClosed(hash)
  owner: address = self.voting.getCandidateOwner(hash)
  # case: listing accepted
  if self.voting.didPass(hash, self.parameterizer.getQuorum()):
    self.backend_address = owner
    log.RegistrationSucceeded(hash, owner)
  else: # application did not pass vote
    log.RegistrationFailed(hash, owner)
    clear(self.backend_url)
  # regardless, the candidate is pruned
  self.voting.removeCandidate(hash)


@public
def purchaseByteCredits(amount: wei_value):
  """
  @notice Allows a user to store byte_credits which can be used later to purchase data from the market
  @dev Msg.sender will have needed to approve the datatrust contract to spend on their behalf
  """
  self.ether_token.transferFrom(msg.sender, self, amount) # note that reserve_payment is implicitly in here
  self.byte_credits[msg.sender] += amount


@public
@constant
def getByteCredits(addr: address) -> wei_value:
  """
  @notice return the amount, in wei, of byte credits a user has purchased
  """
  return self.byte_credits[addr]


@public
def purchaseBytes(addr: address, amount: uint256):
  """
  @notice Pay for listing access in bytes via a user's byte credits
  @dev The requested amount of bytes, at a cost of current cost_per_byte PLUS
  the cost of the backend_payment must be present or this method will revert.
  @param amount The number of bytes requested to purchase
  """
  assert msg.sender == self.backend_address
  cost: wei_value = self.parameterizer.getCostPerByte()
  fee_pct: uint256 = self.parameterizer.getBackendPayment()
  sub_total: wei_value = amount * cost
  fee: wei_value = sub_total / (100 / fee_pct)
  # TODO we _could_ assert credits >= (sub+fee) here, but technically the below operation guards that
  self.byte_credits[addr] -= (sub_total + fee) # clear the users credits first
  # backend is given its fee
  self.ether_token.transfer(self.backend_address, fee)

