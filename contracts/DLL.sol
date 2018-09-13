pragma solidity 0.4.24;


library DLL {

  uint constant NULL_NODE_ID = 0;

  struct Node {
    uint next;
    uint prev;
  }

  struct Data {
    mapping(uint => Node) dll;
  }

  function isEmpty(Data storage self) public view returns (bool) {
    return getStart(self) == NULL_NODE_ID;
  }

  function contains(Data storage self, uint curr) public view returns (bool) {
    if (isEmpty(self) || curr == NULL_NODE_ID) {
      return false;
    }

    bool isSingleNode = (getStart(self) == curr) && (getEnd(self) == curr);
    bool isNullNode = (getNext(self, curr) == NULL_NODE_ID) && (getPrev(self, curr) == NULL_NODE_ID);
    return isSingleNode || !isNullNode;
  }

  function getNext(Data storage self, uint curr) public view returns (uint) {
    return self.dll[curr].next;
  }

  function getPrev(Data storage self, uint curr) public view returns (uint) {
    return self.dll[curr].prev;
  }

  function getStart(Data storage self) public view returns (uint) {
    return getNext(self, NULL_NODE_ID);
  }

  function getEnd(Data storage self) public view returns (uint) {
    return getPrev(self, NULL_NODE_ID);
  }

  /**
  @dev Inserts a new node between _prev and _next. When inserting a node already existing in
  the list it will be automatically removed from the old position.
  @param prev the node which _new will be inserted after
  @param curr the id of the new node being inserted
  @param next the node which _new will be inserted before
  */
  function insert(
    Data storage self,
    uint prev,
    uint curr,
    uint next
  ) public
  {
    require(curr != NULL_NODE_ID, "Error:DLL.insert - ID of a new node may not be 0");

    remove(self, curr);

    require(prev == NULL_NODE_ID || contains(self, prev), "Error:DLL.insert - DLL must contain a previous node");
    require(next == NULL_NODE_ID || contains(self, next), "Error:DLL.insert - DLL must contain a next node");

    require(getNext(self, prev) == next, "Error:DLL.insert - DLL must have the previous node");
    require(getPrev(self, next) == prev, "Error:DLL.insert - DLL must have the next node");

    self.dll[curr].prev = prev;
    self.dll[curr].next = next;

    self.dll[prev].next = curr;
    self.dll[next].prev = curr;
  }

  function remove(Data storage self, uint curr) public {
    if (!contains(self, curr)) {
      return;
    }

    uint next = getNext(self, curr);
    uint prev = getPrev(self, curr);

    self.dll[next].prev = prev;
    self.dll[prev].next = next;

    delete self.dll[curr];
  }
}
