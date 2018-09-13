pragma solidity 0.4.24;

import "./Standard.sol";


/**
 * A subclass of the StandardToken with a constructor.
 */
contract ConstructableStandard is Standard {

  constructor(address initialAccount, uint256 initialBalance) public {
    balances[initialAccount] = initialBalance;
    supply = initialBalance;
  }
}
