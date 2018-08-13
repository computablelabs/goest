pragma solidity 0.4.24;

import "./Burnable.sol";


/**
 * A subclass of the BurnableToken with a constructor.
 */
contract ConstructableBurnable is Burnable {

  constructor(address initialAccount, uint256 initialBalance) public {
    balances[initialAccount] = initialBalance;
    supply = initialBalance;
  }
}
