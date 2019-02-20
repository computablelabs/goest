
# External Contracts
contract Market:
    def isListed(hash: bytes32) -> bool: constant
    def isListing(hash: bytes32) -> bool: constant
    def wasListing(hash: bytes32) -> bool: constant
    def isListingOwner(addr: address) -> bool: constant
    def isInvestor(addr: address) -> bool: constant
    def depositToListing(hash: bytes32, amount: uint256(wei)): modifying
    def withdrawFromListing(hash: bytes32, amount: uint256(wei)): modifying
    def list(listing: string[64], data_hash: bytes32, amount: uint256(wei)): modifying
    def getListingCount() -> int128: constant
    def getListingKey(index: int128) -> bytes32: constant
    def getListing(hash: bytes32) -> (bool, address, bytes32, uint256(wei), uint256(wei)): constant
    def resolveApplication(hash: bytes32): modifying
    def challenge(chall: string[64], hash: bytes32): modifying
    def getChallenge(hash: bytes32) -> (address, uint256(wei), uint256(wei)): constant
    def resolveChallenge(hash: bytes32): modifying
    def getInvested() -> uint256(wei): constant
    def getInvestorCount() -> int128: constant
    def getInvestment(addr: address) -> uint256(wei): constant
    def getInvestmentPrice() -> uint256(wei): constant
    def invest(offer: uint256(wei)): modifying
    def getDivestmentProceeds(addr: address) -> uint256(wei): constant
    def divest(): modifying
    def exit(hash: bytes32): modifying


