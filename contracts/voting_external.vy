
# External Contracts
contract Voting:
    def getPrivileged() -> (address, address, address, address, address): constant
    def setPrivileged(parameterizer: address, datatrust: address, listing: address, investing: address): modifying
    def getCouncilCount() -> uint256: constant
    def inCouncil(member: address) -> bool: constant
    def addToCouncil(member: address): modifying
    def removeFromCouncil(member: address): modifying
    def candidateIs(hash: bytes32, kind: uint256) -> bool: constant
    def isCandidate(hash: bytes32) -> bool: constant
    def getCandidate(hash: bytes32) -> (uint256, address, uint256(sec, positional), int128): constant
    def getCandidateOwner(hash: bytes32) -> address: constant
    def addCandidate(hash: bytes32, kind: uint256, owner: address, vote_by: uint256(sec)): modifying
    def removeCandidate(hash: bytes32): modifying
    def didPass(hash: bytes32, quorum: uint256) -> bool: constant
    def didVote(hash: bytes32, member: address) -> bool: constant
    def pollClosed(hash: bytes32) -> bool: constant
    def vote(hash: bytes32): modifying


