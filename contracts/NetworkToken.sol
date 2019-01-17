pragma solidity 0.5.2;

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
   * @param amount The amount of tokens to be spent.
   */
  function approve(address spender, uint256 amount) external {
    selfAllowed[msg.sender][spender] = amount;
    emit ApprovalEvent(msg.sender, spender, amount);
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
   * @param spender The address which will spend the funds.
   * @param amount The amount of tokens to decrease the allowance by.
   */
  function decreaseApproval(address spender, uint256 amount) external {
    if (amount > selfAllowed[msg.sender][spender]) {
      selfAllowed[msg.sender][spender] = 0;
    } else {
      selfAllowed[msg.sender][spender] = selfAllowed[msg.sender][spender].sub(amount);
    }
    emit ApprovalEvent(msg.sender, spender, selfAllowed[msg.sender][spender]);
  }

  /**
   * Increase the amount of tokens that an owner allowed to a spender.
   * approve should be called when allowed[spender] == 0. To increment
   * allowed value is better to use this function to avoid 2 calls (and wait until
   * the first transaction is mined)
   * @param spender The address which will spend the funds.
   * @param amount The amount of tokens to increase the allowance by.
   */
  function increaseApproval(address spender, uint256 amount) external {
    selfAllowed[msg.sender][spender] = (
      selfAllowed[msg.sender][spender].add(amount));
    emit ApprovalEvent(msg.sender, spender, selfAllowed[msg.sender][spender]);
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
  * @param amount The amount to be transferred.
  */
  function transfer(address to, uint256 amount) external {
    require(to != address(0), "Error:Basic.transfer - 'to' address must be specified");
    require(amount <= selfBalances[msg.sender], "Error:Basic.transfer - Amount exceeds the balance of msg.sender");

    selfBalances[msg.sender] = selfBalances[msg.sender].sub(amount);
    selfBalances[to] = selfBalances[to].add(amount);
    emit TransferEvent(msg.sender, to, amount);
  }

  /**
   * Transfer tokens from one address to another
   * @param from address The address which you want to send tokens from
   * @param to address The address which you want to transfer to
   * @param amount uint256 the amount of tokens to be transferred
   */
  function transferFrom(address from, address to, uint256 amount) external {
    require(to != address(0), "Error:Standard.transferFrom - 'to' address must be specified");
    require(amount <= selfBalances[from], "Error:Standard.transferFrom - Amount exceeds available balance");
    require(amount <= selfAllowed[from][msg.sender], "Error.Standard.transferFrom - Amount exceeds allowed amount");

    selfBalances[from] = selfBalances[from].sub(amount);
    selfBalances[to] = selfBalances[to].add(amount);
    selfAllowed[from][msg.sender] = selfAllowed[from][msg.sender].sub(amount);
    emit TransferEvent(from, to, amount);
  }

  event ApprovalEvent(address indexed holder, address indexed spender, uint256 amount);
  event TransferEvent(address indexed from, address indexed to, uint256 amount);
}
