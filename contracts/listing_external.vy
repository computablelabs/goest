
# External Contracts
contract Listing:
    def isListed(hash: bytes32) -> bool: constant
    def isListing(hash: bytes32) -> bool: constant
    def wasListing(hash: bytes32) -> bool: constant
    def isListingOwner(addr: address) -> bool: constant
    def depositToListing(hash: bytes32, amount: uint256(wei)): modifying
    def withdrawFromListing(hash: bytes32, amount: uint256(wei)): modifying
    def list(listing: string[64], data_hash: bytes32, amount: uint256(wei)): modifying
    def getListingCount() -> int128: constant
    def getListingKey(index: int128) -> bytes32: constant
    def getListing(hash: bytes32) -> (bool, address, bytes32, uint256(wei), uint256(wei)): constant
    def convertListing(hash: bytes32): modifying
    def resolveApplication(hash: bytes32): modifying
    def challenge(chall: string[64], hash: bytes32): modifying
    def getChallenge(hash: bytes32) -> address: constant
    def resolveChallenge(hash: bytes32): modifying
    def exit(hash: bytes32): modifying


