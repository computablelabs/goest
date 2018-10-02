pragma solidity 0.4.25;

import "./SafeMath.sol";

/*
 * The Market token is an ERC20 Standard + Mintable + Burnable
 * NOTE: We do not use any sort of 'ownable' library, using instead a simplified
 * private owner variable, a public getter, and a modifier. If, in the future, we
 * need to provide more ownership functionality we will just add it here.
*/
contract MarketToken {
  using SafeMath for uint256;

  uint256 private selfSupply;
  bool private selfMintingStopped = false;
  address private selfOwner; // will be the market factory
  address private selfMarket; // immutable once set by the owner (market factory)
  mapping(address => uint256) private selfBalances;
  mapping (address => mapping (address => uint256)) private selfAllowed;

  constructor(address initialAccount, uint256 initialBalance) public {
    selfOwner = msg.sender;
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
   * Burns a specific amount of tokens.
   * @param value The amount of token to be burned.
   */
  function burn(uint256 value) external {
    doBurn(msg.sender, value);
  }

  /**
   * Burns a specific amount of tokens from the target address and decrements allowance
   * @param from address The address which you want to send tokens from
   * @param value uint256 The amount of token to be burned
   */
  function burnFrom(address from, uint256 value) external {
    require(value <= selfAllowed[from][msg.sender], "Error:MarketToken.burnFrom - Value exceeds allowed amount");
    // Should https://github.com/OpenZeppelin/zeppelin-solidity/issues/707 be accepted,
    // this function needs to emit an event with the updated approval.
    selfAllowed[from][msg.sender] = selfAllowed[from][msg.sender].sub(value);
    doBurn(from, value);
  }

  /**
   * If minting functionality has been explicitly stopped, we don't mint any longer
   */
  modifier canMint() {
    require(!selfMintingStopped, "Error:MarketToken.canMint - Minting has been stopped");
    _;
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
   * Abstracted logic for burning used by both `burn` and `burnFrom`
   */
  function doBurn(address who, uint256 value) private {
    require(value <= selfBalances[who], "Error:NetworkToken.doBurn - Value exceeds available balance");
    // no need to require value <= supply, since that would imply the
    // sender's balance is greater than the supply, which *should* be an assertion failure

    selfBalances[who] = selfBalances[who].sub(value);
    selfSupply = selfSupply.sub(value);
    emit BurnEvent(who, value);
    emit TransferEvent(who, address(0), value); // 'to' address will be empty in this case
  }

  /**
   * @dev Return the address of the contract owner, should be the MarketFactory TODO
   */
  function getOwner() external view returns(address) {
    return selfOwner;
  }

  /**
   * Only the market contract has mint permission
   */
  modifier hasMintPermission() {
    require(msg.sender == selfMarket, "Error:MarketToken.hasMintPermission - Caller must be market contract");
    _;
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

  modifier isOwner() {
    require(msg.sender == selfOwner, "Error:MarketToken.isOwner - Caller must be owner");
    _;
  }

  /**
   * @dev Create new tokens for the given amount, and give them to the given address
   * @param to The recipient of the new tokens
   * @param amount How many tokens to mint
   */
  function mint(address to, uint256 amount) external hasMintPermission canMint returns (bool) {
    require(to != 0, "Error:MarketToken.mint - 'to' account may not be 0");

    selfSupply = selfSupply.add(amount);
    selfBalances[to] = selfBalances[to].add(amount);
    emit TransferEvent(address(0), to, amount);
    return true;
  }

  /**
   * @dev Return the boolean flag telling if minting has ceased
   */
  function mintingStopped() external view returns(bool) {
    return selfMintingStopped;
  }

  function setMarket(address market) external isOwner returns (bool) {
    // we only allow this once, so the current val of market must be 0 initialized still
    require(selfMarket == address(0), "Error:MarketToken.setMarket - Market address is immutable once set");
    selfMarket = market;
    return true;
  }

  /**
   * @dev explicity state that no more tokens may be minted
   */
  function stopMinting() external hasMintPermission canMint returns (bool) {
    selfMintingStopped = true;
    emit MintStoppedEvent();
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
    require(to != address(0), "Error:Basic.transfer - An address must be specified");
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
    require(to != address(0), "Error:Standard.transferFrom - 'to' address must be specified");
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
  event BurnEvent(address indexed burner, uint256 value);
  event MintStoppedEvent();
}
