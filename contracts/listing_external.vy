
# External Contracts
contract Listing:
    def isListed(hash: bytes32) -> bool: constant
    def depositToListing(hash: bytes32, amount: uint256(wei)): modifying
    def withdrawFromListing(hash: bytes32, amount: uint256(wei)): modifying
    def list(listing: string[64]): modifying
    def getHash(str: string[64]) -> bytes32: constant
    def getListing(hash: bytes32) -> (address, uint256(wei), uint256(wei)): constant
    def convertListing(hash: bytes32): modifying
    def resolveApplication(hash: bytes32): modifying
    def challenge(hash: bytes32): modifying
    def resolveChallenge(hash: bytes32): modifying
    def exit(hash: bytes32): modifying


