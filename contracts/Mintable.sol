pragma solidity 0.4.24;

import "./Standard.sol";


/**
 * Mintable Token
 */
contract Mintable is Standard {

  event Mint(address indexed to, uint256 amount);
  event MintStopped();

  bool public mintingStopped = false;
  address public owner;

  modifier canMint() {
    require(!mintingStopped, "Error:Mintable.canMint - Minting has been stopped");
    _;
  }

  modifier hasMintPermission() {
    require(msg.sender == owner, "Error:Mintable.hasMintPermission - Caller must be owner");
    _;
  }

  function mint(address to, uint256 amount) public hasMintPermission canMint returns (bool) {
    require(to != 0, "Error:Mintable.mint - 'to' account may not be 0");

    supply = supply.add(amount);
    balances[to] = balances[to].add(amount);
    emit Transfer(address(0), to, amount);
    emit Mint(to, amount);
    return true;
  }

  function stopMinting() public hasMintPermission canMint returns (bool) {
    mintingStopped = true;
    emit MintStopped();
    return true;
  }
}
