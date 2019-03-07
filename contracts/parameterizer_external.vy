
# External Contracts
contract Parameterizer:
    def getChallengeStake() -> uint256(wei): constant
    def getConversionRate() -> uint256(wei): constant
    def getHash(str: string[64]) -> bytes32: constant
    def getInvestDenominator() -> uint256: constant
    def getInvestNumerator() -> uint256: constant
    def getListReward() -> uint256(wei): constant
    def getQuorum() -> uint256: constant
    def getReparam(hash: bytes32) -> (uint256, uint256): constant
    def getVoteBy() -> uint256(sec): constant
    def reparameterize(reparam: string[64], param: uint256, value: uint256): modifying
    def resolveReparam(hash: bytes32): modifying


