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
  address private selfMarketAddress; // address of the market contract which has some priviledge here
  mapping(address => uint256) private selfBalances;
  mapping (address => mapping (address => uint256)) private selfAllowed;

  constructor(address initialAccount, uint256 initialBalance) public {
    selfOwner = msg.sender; // TODO this may not be the case
    selfBalances[initialAccount] = initialBalance; // likely to change as well TODO
    selfSupply = initialBalance; // all-the-changes ...
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
  function approve(address spender, uint256 amount) external returns (bool) {
    selfAllowed[msg.sender][spender] = amount;
    emit ApprovalEvent(msg.sender, spender, amount);
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
   * @param amount The amount of token to be burned.
   */
  function burn(uint256 amount) external hasPrivilege {
    require(amount <= selfBalances[msg.sender], "Error:NetworkToken.burn - Amount exceeds available balance");
    // no need to require amout <= supply, since that would imply the
    // sender's balance is greater than the supply, which *should* be an assertion failure

    selfBalances[msg.sender] = selfBalances[msg.sender].sub(amount);
    selfSupply = selfSupply.sub(amount);
    emit BurnEvent(msg.sender, amount);
    emit TransferEvent(msg.sender, address(0), amount); // from: market, to: nobody
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
   * @param subtractedAmount The amount of tokens to decrease the allowance by.
   */
  function decreaseApproval(address spender, uint256 subtractedAmount) external returns (bool) {
    uint256 oldAmount= selfAllowed[msg.sender][spender];
    if (subtractedAmount > oldAmount) {
      selfAllowed[msg.sender][spender] = 0;
    } else {
      selfAllowed[msg.sender][spender] = oldAmount.sub(subtractedAmount);
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
   * @param addedAmount The amount of tokens to increase the allowance by.
   */
  function increaseApproval(address spender, uint256 addedAmount) external returns (bool) {
    selfAllowed[msg.sender][spender] = (
      selfAllowed[msg.sender][spender].add(addedAmount));
    emit ApprovalEvent(msg.sender, spender, selfAllowed[msg.sender][spender]);
    return true;
  }

  /**
   * Only the market contract has mint and burn permissions
   */
  modifier hasPrivilege() {
    require(msg.sender == selfMarketAddress, "Error:MarketToken.hasPrivilege - Caller must be a privileged  contract");
    _;
  }

  modifier isOwner() {
    require(msg.sender == selfOwner, "Error:MarketToken.isOwner - Caller must be owner");
    _;
  }

  /**
   * @dev Create new tokens for the given amount, and bank them with the market until they are either burned or transferred
   * @param amount How many tokens to mint
   */
  function mint(uint256 amount) external hasPrivilege canMint returns (bool) {
    selfSupply = selfSupply.add(amount);
    selfBalances[msg.sender] = selfBalances[msg.sender].add(amount);
    emit TransferEvent(address(0), msg.sender, amount); // from: nobody, to: market
    return true;
  }

  /**
   * @dev Return the boolean flag telling if minting has ceased
   */
  function mintingStopped() external view returns(bool) {
    return selfMintingStopped;
  }

  function setPrivilegedContracts(address market) external isOwner returns (bool) {
    // we only allow this once, so the current val of market must be 0 initialized still
    require(selfMarketAddress == address(0), "Error:MarketToken.setPrivilegedContracts - Market address already set");
    selfMarketAddress = market;
    return true;
  }

  /**
   * @dev explicity state that no more tokens may be minted
   */
  function stopMinting() external hasPrivilege canMint returns (bool) {
    selfMintingStopped = true;
    emit MintStoppedEvent();
    return true;
  }

  /**
  * @dev Total number of tokens in existence
  * I'd prefer this to be `getSupply` to go along with our style for getters,
  * but we'll stick to the ERC20 interface for this one.
  */
  function totalSupply() external view returns (uint256) {
    return selfSupply;
  }

  /**
  * Transfer token for a specified address
  * @param to The address to transfer to.
  * @param amount The amount to be transferred.
  */
  function transfer(address to, uint256 amount) external returns (bool) {
    require(to != address(0), "Error:Basic.transfer - An address must be specified");
    require(amount <= selfBalances[msg.sender], "Error:Basic.transfer - Amount exceeds the balance of msg.sender");

    selfBalances[msg.sender] = selfBalances[msg.sender].sub(amount);
    selfBalances[to] = selfBalances[to].add(amount);
    emit TransferEvent(msg.sender, to, amount);
    return true;
  }

  /**
   * Transfer tokens from one address to another
   * @param from address The address which you want to send tokens from
   * @param to address The address which you want to transfer to
   * @param amount uint256 the amount of tokens to be transferred
   */
  function transferFrom(address from, address to, uint256 amount) external returns (bool) {
    require(to != address(0), "Error:Standard.transferFrom - 'to' address must be specified");
    require(amount <= selfBalances[from], "Error:Standard.transferFrom - Amount exceeds available balance");
    require(amount <= selfAllowed[from][msg.sender], "Error.Standard.transferFrom - Amount exceeds allowed amount");

    selfBalances[from] = selfBalances[from].sub(amount);
    selfBalances[to] = selfBalances[to].add(amount);
    selfAllowed[from][msg.sender] = selfAllowed[from][msg.sender].sub(amount);
    emit TransferEvent(from, to, amount);
    return true;
  }

  event ApprovalEvent(address indexed holder, address indexed spender, uint256 amount);
  event TransferEvent(address indexed from, address indexed to, uint256 amount);
  event BurnEvent(address indexed burner, uint256 amount);
  event MintStoppedEvent();
}
