
# External Contracts
contract Parameterizer:
    def getChallengeStake() -> uint256(wei): constant
    def getConversionRate() -> uint256(wei): constant
    def getConversionSlopeDenominator() -> uint256: constant
    def getConversionSlopeNumerator() -> uint256: constant
    def getListReward() -> uint256(wei): constant
    def getQuorum() -> uint256: constant
    def getVoteBy() -> uint256(sec): constant
    def getParamHash(reparam: string[64]) -> bytes32: constant
    def getReparam(hash: bytes32) -> (address, uint256, uint256): constant
    def resolveReparam(hash: bytes32): modifying
    def reparameterize(reparam: string[64], param: uint256, value: uint256): modifying


