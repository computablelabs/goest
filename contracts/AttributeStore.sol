pragma solidity 0.4.24;


library AttributeStore {
  struct Data {
    mapping(bytes32 => uint) store;
  }

  function getAttribute(Data storage self, bytes32 UUID, string attrName)
  public view returns (uint)
  {
    bytes32 key = keccak256(abi.encodePacked(UUID, attrName));
    return self.store[key];
  }

  function setAttribute(
    Data storage self,
    bytes32 UUID,
    string attrName,
    uint attrVal) public
  {
    bytes32 key = keccak256(abi.encodePacked(UUID, attrName));
    self.store[key] = attrVal;
  }
}
