pragma solidity 0.4.24;

import "./ERC20Basic.sol";


/**
 * ERC20 interface
 * See https://github.com/ethereum/EIPs/issues/20
 */
contract ERC20 is ERC20Basic {
  function allowance(address holder, address spender)
    public view returns (uint256);

  function transferFrom(address from, address to, uint256 value)
    public returns (bool);

  function approve(address spender, uint256 value) public returns (bool);
  event Approval(
    address indexed holder,
    address indexed spender,
    uint256 value
  );
}
