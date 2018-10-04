pragma solidity 0.4.25;

import "./SafeMath.sol";

/*
 * The Computable Network token is a stand-alone (no inheritance) implementation of an ERC20 Standard Token
 * // TODO determine if the network token can be a basic token vs standard
*/
contract NetworkToken {
  using SafeMath for uint256;

  uint256 private selfSupply;
  mapping(address => uint256) private selfBalances;
  mapping (address => mapping (address => uint256)) private selfAllowed;

  constructor(address initialAccount, uint256 initialBalance) public {
    selfBalances[initialAccount] = initialBalance;
    selfSupply = initialBalance;
  }

  /**
   * Function to check the amount of tokens that an owner allowed to a spender.
   * @param holder address The address which owns the funds.
   * @param spender address The address which will spend the funds.
   * @return A uint256 specifying the amount of tokens still available for the spender.
   */
  function allowance(address holder, address spender) external view returns (uint256) {
    return selfAllowed[holder][spender];
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
  function approve(address spender, uint256 value) external returns (bool) {
    selfAllowed[msg.sender][spender] = value;
    emit ApprovalEvent(msg.sender, spender, value);
    return true;
  }

  /**
  * Gets the balance of the specified address.
  * @param holder The address to query the the balance of.
  * @return An uint256 representing the amount owned by the passed address.
  */
  function balanceOf(address holder) external view returns (uint256) {
    return selfBalances[holder];
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
  function decreaseApproval(address spender, uint256 subtractedValue) external returns (bool) {
    uint256 oldValue = selfAllowed[msg.sender][spender];
    if (subtractedValue > oldValue) {
      selfAllowed[msg.sender][spender] = 0;
    } else {
      selfAllowed[msg.sender][spender] = oldValue.sub(subtractedValue);
    }
    emit ApprovalEvent(msg.sender, spender, selfAllowed[msg.sender][spender]);
    return true;
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
  function increaseApproval(address spender, uint256 addedValue) external returns (bool) {
    selfAllowed[msg.sender][spender] = (
      selfAllowed[msg.sender][spender].add(addedValue));
    emit ApprovalEvent(msg.sender, spender, selfAllowed[msg.sender][spender]);
    return true;
  }

  /**
  * @dev Total number of tokens in existence
  */
  function totalSupply() external view returns (uint256) {
    return selfSupply;
  }

  /**
  * Transfer token for a specified address
  * @param to The address to transfer to.
  * @param value The amount to be transferred.
  */
  function transfer(address to, uint256 value) external returns (bool) {
    require(to != address(0), "Error:Basic.transfer - 'to' cannot be the zero-address");
    require(value <= selfBalances[msg.sender], "Error:Basic.transfer - Value exceeds the balance of msg.sender");

    selfBalances[msg.sender] = selfBalances[msg.sender].sub(value);
    selfBalances[to] = selfBalances[to].add(value);
    emit TransferEvent(msg.sender, to, value);
    return true;
  }

  /**
   * Transfer tokens from one address to another
   * @param from address The address which you want to send tokens from
   * @param to address The address which you want to transfer to
   * @param value uint256 the amount of tokens to be transferred
   */
  function transferFrom(address from, address to, uint256 value) external returns (bool) {
    require(to != address(0), "Error:Standard.transferFrom - 'to' may not be the zero-address");
    require(value <= selfBalances[from], "Error:Standard.transferFrom - Value exceeds available balance");
    require(value <= selfAllowed[from][msg.sender], "Error.Standard.transferFrom - Value exceeds allowed amount");

    selfBalances[from] = selfBalances[from].sub(value);
    selfBalances[to] = selfBalances[to].add(value);
    selfAllowed[from][msg.sender] = selfAllowed[from][msg.sender].sub(value);
    emit TransferEvent(from, to, value);
    return true;
  }

  event ApprovalEvent(address indexed holder, address indexed spender, uint256 value);
  event TransferEvent(address indexed from, address indexed to, uint256 value);
}
