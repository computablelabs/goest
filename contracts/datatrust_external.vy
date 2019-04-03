
# External Contracts
contract Datatrust:
    def getPrivileged() -> address: constant
    def setPrivileged(listing: address): modifying
    def getHash(url: string[128]) -> bytes32: constant
    def getBackendAddress() -> address: constant
    def getBackendUrl() -> string[128]: constant
    def setBackendUrl(url: string[128]): modifying
    def getDataHash(hash: bytes32) -> bytes32: constant
    def setDataHash(listing: bytes32, data: bytes32): modifying
    def removeDataHash(hash: bytes32): modifying
    def register(url: string[128]): modifying
    def resolveRegistration(hash: bytes32): modifying
    def purchaseByteCredits(amount: uint256(wei)): modifying
    def getByteCredits(addr: address) -> uint256(wei): constant


