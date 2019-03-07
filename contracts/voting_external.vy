
# External Contracts
contract Voting:
    def getPrivileged() -> (address, address, address, address): constant
    def setPrivileged(parameterizer: address, listing: address, investing: address): modifying
    def inCouncil(member: address) -> bool: constant
    def getCouncilCount() -> int128: constant
    def addToCouncil(member: address): modifying
    def removeFromCouncil(member: address): modifying
    def candidateIs(hash: bytes32, kind: uint256) -> bool: constant
    def isCandidate(hash: bytes32) -> bool: constant
    def wasCandidate(hash: bytes32) -> bool: constant
    def getCandidateCount() -> int128: constant
    def getCandidateKey(index: int128) -> bytes32: constant
    def getCandidate(hash: bytes32) -> (uint256, uint256(sec, positional), uint256): constant
    def addCandidate(hash: bytes32, kind: uint256, vote_by: uint256(sec)): modifying
    def removeCandidate(hash: bytes32): modifying
    def didPass(hash: bytes32, quorum: uint256) -> bool: constant
    def didVote(hash: bytes32, member: address) -> bool: constant
    def pollClosed(hash: bytes32) -> bool: constant
    def vote(hash: bytes32): modifying
    def willAddCandidate(hash: bytes32) -> bool: constant
    def willAddToCouncil(addr: address) -> bool: constant


