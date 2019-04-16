# @title Computable Datatrust
# @notice Handle the details pertaining to Computable Datatrust functionality
# @author Computable

# constants
REGISTRATION: constant(uint256) = 4 # candidate.kind

# structs
struct Delivery:
  owner:address
  bytes_requested: uint256
  bytes_delivered: uint256

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
  def getReservePayment() -> uint256: constant
  def getCostPerByte() -> wei_value: constant
  def getStake() -> uint256(wei): constant
  def getQuorum() -> uint256: constant
  def getVoteBy() -> uint256(sec): constant

# events
Registered: event({hash: indexed(bytes32), registrant: indexed(address)})
RegistrationSucceeded: event({hash: indexed(bytes32), registrant: indexed(address)})
RegistrationFailed: event({hash: indexed(bytes32), registrant: indexed(address)})
DeliveryRequested: event({hash: indexed(bytes32), requester: indexed(address), amount: uint256})
Delivered: event({hash: indexed(bytes32), owner: indexed(address), url: bytes32})

# state vars
data_hashes: map(bytes32, bytes32) # listing_hash -> data_hash
byte_credits: map(address, wei_value) # maps user-address to byte-credits
access_credits: map(bytes32, wei_value) # maps a listing to any accumulated access rewards it may claim
deliveries: map(bytes32, Delivery) # maps delivery_hash -> delivery
backend_url: string[128]
backend_address: address
ether_token: EtherToken
voting: Voting
parameterizer: Parameterizer
factory_address: address
investing_address: address
listing_address: address

@public
def __init__(ether_token_addr: address, voting_addr: address, p11r_addr: address, inv_addr: address):
  self.factory_address = msg.sender
  self.ether_token = EtherToken(ether_token_addr)
  self.voting = Voting(voting_addr)
  self.parameterizer = Parameterizer(p11r_addr)
  self.investing_address = inv_addr # does not consume the contract, only knows the address


@public
@constant
def getPrivileged() -> address:
  """
  @notice Fetch a list of each privileged address recognized by this contract
  @return Listing contract address
  """
  return self.listing_address


@public
@constant
def getInvesting() -> address:
  """
  @notice Return the address at which the datatrust contract transfers reserve payments to
  @dev This should always be the address of the Investing contract deployed with this market
  """
  return self.investing_address


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
def requestDelivery(hash: bytes32, amount: uint256):
  """
  @notice Allow a user to request that a delivery of 'amount' (bytes) be made available to them.
  @param hash This is a keccack hash recieved by a client that uniquely identifies a request. NOTE care should be taken
  by the client to insure this is uinque.
  @param amount The number of bytes the user is paying for.
  """
  assert self.deliveries[hash].owner == ZERO_ADDRESS # not already a request
  total: wei_value = self.parameterizer.getCostPerByte() * amount
  res_fee: wei_value = total / (100 / self.parameterizer.getReservePayment())
  self.ether_token.transferFrom(msg.sender, self, total) # take the total payment
  self.ether_token.transfer(self.investing_address, res_fee) # transfer res_pct to reserve
  self.byte_credits[msg.sender] += total # should reflect the total payment
  self.deliveries[hash].owner = msg.sender
  self.deliveries[hash].bytes_requested = amount


@public
@constant
def getByteCredits(addr: address) -> wei_value:
  """
  @notice return the amount, in wei, of byte credits a user has purchased
  """
  return self.byte_credits[addr]


@public
@constant
def getDelivery(hash: bytes32) -> (address, uint256, uint256):
  """
  @notice Return the data present in a given delivery request
  @param hash The hashed delivery identifier
  """
  return (self.deliveries[hash].owner, self.deliveries[hash].bytes_requested,
    self.deliveries[hash].bytes_delivered)


@public
def listingAccessed(listing:bytes32, delivery: bytes32, amount: uint256):
  """
  @dev Only a registered backend may call. Enforce that the claimed listing exists.
  @param listing The listing that was accessed
  @param delivery Which delivery object this access was for
  @param amount How many bytes were accessed
  """
  assert msg.sender == self.backend_address
  assert self.data_hashes[listing] != EMPTY_BYTES32
  total: wei_value = self.parameterizer.getCostPerByte() * amount
  # this can be claimed later by the listing owner, and are subtractive to byte_credits
  self.access_credits[listing] += total
  self.byte_credits[self.deliveries[delivery].owner] -= total
  # bytes_delivered must eq (or exceed) bytes_requested in order for a datatrust to claim delivery
  self.deliveries[delivery].bytes_delivered += amount


@public
@constant
def getAccessCredits(hash: bytes32) -> wei_value:
  """
  @notice return the amount, in wei, of access credits a listing has accumulated
  """
  return self.access_credits[hash]


@public
def delivered(delivery: bytes32, url: bytes32):
  """
  @notice Allow a backend to collect its payment.
  @dev We check that a backend has delivered, at least, the amount of bytes requested
  @param delivery Identifier of the delivery in question
  @param url A hash of the URL that the backend delivered to
  """
  assert msg.sender == self.backend_address
  owner: address = self.deliveries[delivery].owner
  delivered: uint256 = self.deliveries[delivery].bytes_delivered
  assert delivered >= self.deliveries[delivery].bytes_requested
  # clear the delivery record first
  clear(self.deliveries[delivery])
  # now pay the datatrust from the banked delivery request
  back_fee: wei_value = (self.parameterizer.getCostPerByte() * delivered) / (100 / self.parameterizer.getBackendPayment())
  self.ether_token.transfer(self.backend_address, back_fee)
  log.Delivered(delivery, owner, url)

