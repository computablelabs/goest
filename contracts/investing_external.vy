
# External Contracts
contract Investing:
    def isInvestor(addr: address) -> bool: constant
    def getInvested() -> uint256(wei): constant
    def getInvestorCount() -> int128: constant
    def getInvestment(addr: address) -> uint256(wei): constant
    def getInvestmentPrice() -> uint256(wei): constant
    def invest(offer: uint256(wei)): modifying
    def getDivestmentProceeds(addr: address) -> uint256(wei): constant
    def divest(): modifying


