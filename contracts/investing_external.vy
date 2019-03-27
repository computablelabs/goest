
# External Contracts
contract Investing:
    def getInvestmentPrice() -> uint256(wei): constant
    def invest(offer: uint256(wei)): modifying
    def getDivestmentProceeds(addr: address) -> uint256(wei): constant
    def divest(): modifying


