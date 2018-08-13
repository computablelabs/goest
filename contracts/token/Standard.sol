pragma solidity 0.4.24;

import "./Basic.sol";
import "./ERC20.sol";


/**
 * Standard ERC20 token
 *
 * Implementation of the basic standard token.
 * https://github.com/ethereum/EIPs/issues/20
 * Based on code by Open Zeppelin: https://github.com/openzeppelin-solidity
 */
contract Standard is ERC20, Basic {

  mapping (address => mapping (address => uint256)) internal allowed;


  /**
   * Transfer tokens from one address to another
   * @param from address The address which you want to send tokens from
   * @param to address The address which you want to transfer to
   * @param value uint256 the amount of tokens to be transferred
   */
  function transferFrom(
    address from,
    address to,
    uint256 value
  )
    public
    returns (bool)
  {
    require(to != address(0));
    require(value <= balances[from]);
    require(value <= allowed[from][msg.sender]);

    balances[from] = balances[from].sub(value);
    balances[to] = balances[to].add(value);
    allowed[from][msg.sender] = allowed[from][msg.sender].sub(value);
    emit Transfer(from, to, value);
    return true;
  }

  /**
   * Approve the passed address to spend the specified amount of tokens on behalf of msg.sender.
   * Beware that changing an allowance with this method brings the risk that someone may use both the old
   * and the new allowance by unfortunate transaction ordering. One possible solution to mitigate this
   * race condition is to first reduce the spender's allowance to 0 and set the desired value afterwards:
   * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
   * @param spender The address which will spend the funds.
   * @param value The amount of tokens to be spent.
   */
  function approve(address spender, uint256 value) public returns (bool) {
    allowed[msg.sender][spender] = value;
    emit Approval(msg.sender, spender, value);
    return true;
  }

  /**
   * Function to check the amount of tokens that an owner allowed to a spender.
   * @param owner address The address which owns the funds.
   * @param spender address The address which will spend the funds.
   * @return A uint256 specifying the amount of tokens still available for the spender.
   */
  function allowance(
    address owner,
    address spender
   )
    public
    view
    returns (uint256)
  {
    return allowed[owner][spender];
  }

  /**
   * Increase the amount of tokens that an owner allowed to a spender.
   * approve should be called when allowed[spender] == 0. To increment
   * allowed value is better to use this function to avoid 2 calls (and wait until
   * the first transaction is mined)
   * From MonolithDAO Token.sol
   * @param spender The address which will spend the funds.
   * @param addedValue The amount of tokens to increase the allowance by.
   */
  function increaseApproval(
    address spender,
    uint256 addedValue
  )
    public
    returns (bool)
  {
    allowed[msg.sender][spender] = (
      allowed[msg.sender][spender].add(addedValue));
    emit Approval(msg.sender, spender, allowed[msg.sender][spender]);
    return true;
  }

  /**
   * Decrease the amount of tokens that an owner allowed to a spender.
   * approve should be called when allowed[spender] == 0. To decrement
   * allowed value is better to use this function to avoid 2 calls (and wait until
   * the first transaction is mined)
   * From MonolithDAO Token.sol
   * @param spender The address which will spend the funds.
   * @param subtractedValue The amount of tokens to decrease the allowance by.
   */
  function decreaseApproval(
    address spender,
    uint256 subtractedValue
  )
    public
    returns (bool)
  {
    uint256 oldValue = allowed[msg.sender][spender];
    if (subtractedValue > oldValue) {
      allowed[msg.sender][spender] = 0;
    } else {
      allowed[msg.sender][spender] = oldValue.sub(subtractedValue);
    }
    emit Approval(msg.sender, spender, allowed[msg.sender][spender]);
    return true;
  }
}
