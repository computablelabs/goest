pragma solidity 0.4.24;

import "./Basic.sol";


/**
 * A subclass of the BasicToken with a constructor.
 */
contract ConstructableBasic is Basic {

  constructor(address initialAccount, uint256 initialBalance) public {
    balances[initialAccount] = initialBalance;
    supply = initialBalance;
  }
}
