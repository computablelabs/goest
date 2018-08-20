pragma solidity 0.4.24;


import "./ERC20Basic.sol";
import "./SafeMath.sol";


/**
 * Basic token
 * Basic version of StandardToken, with no allowances.
 */
contract Basic is ERC20Basic {
  using SafeMath for uint256;

  mapping(address => uint256) balances;

  uint256 supply;

  /**
  * @dev Total number of tokens in existence
  */
  function totalSupply() public view returns (uint256) {
    return supply;
  }

  /**
  * Transfer token for a specified address
  * @param to The address to transfer to.
  * @param value The amount to be transferred.
  */
  function transfer(address to, uint256 value) public returns (bool) {
    require(to != address(0), "Error:Basic.transfer - 'to' address cannot be owner");
    require(value <= balances[msg.sender], "Error:Basic.transfer - Value exceeds the balance of msg.sender");

    balances[msg.sender] = balances[msg.sender].sub(value);
    balances[to] = balances[to].add(value);
    emit Transfer(msg.sender, to, value);
    return true;
  }

  /**
  * Gets the balance of the specified address.
  * @param owner The address to query the the balance of.
  * @return An uint256 representing the amount owned by the passed address.
  */
  function balanceOf(address owner) public view returns (uint256) {
    return balances[owner];
  }
}
