pragma solidity 0.4.24;

import "./Mintable.sol";


/**
 * A subclass of the MintableToken with a constructor.
 */
contract ConstructableMintable is Mintable {

  constructor(address initialAccount, uint256 initialBalance) public {
    owner = msg.sender;
    balances[initialAccount] = initialBalance;
    supply = initialBalance;
  }
}
