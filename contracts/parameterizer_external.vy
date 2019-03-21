
# External Contracts
contract Parameterizer:
    def getBackendPayment() -> uint256: constant
    def getMakerPayment() -> uint256: constant
    def getCostPerByte() -> uint256(wei): constant
    def getChallengeStake() -> uint256(wei): constant
    def getConversionRate() -> uint256(wei): constant
    def getHash(param: uint256, value: uint256) -> bytes32: constant
    def getInvestDenominator() -> uint256: constant
    def getInvestNumerator() -> uint256: constant
    def getListReward() -> uint256(wei): constant
    def getAccessReward() -> uint256(wei): constant
    def getQuorum() -> uint256: constant
    def getReparam(hash: bytes32) -> (uint256, uint256): constant
    def getVoteBy() -> uint256(sec): constant
    def reparameterize(param: uint256, value: uint256): modifying
    def resolveReparam(hash: bytes32): modifying


