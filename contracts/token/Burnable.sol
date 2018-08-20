pragma solidity 0.4.24;

import "./Standard.sol";


/**
 * Burnable Token
 * Token that can be irreversibly burned (destroyed).
 */
contract Burnable is Standard {

  event Burn(address indexed burner, uint256 value);

  /**
   * Burns a specific amount of tokens.
   * @param value The amount of token to be burned.
   */
  function burn(uint256 value) public {
    doBurn(msg.sender, value);
  }

  /**
   * Burns a specific amount of tokens from the target address and decrements allowance
   * @param from address The address which you want to send tokens from
   * @param value uint256 The amount of token to be burned
   */
  function burnFrom(address from, uint256 value) public {
    require(value <= allowed[from][msg.sender], "Error:Burnable.burnFrom - Value exceeds allowed amount");
    // Should https://github.com/OpenZeppelin/zeppelin-solidity/issues/707 be accepted,
    // this function needs to emit an event with the updated approval.
    allowed[from][msg.sender] = allowed[from][msg.sender].sub(value);
    doBurn(from, value);
  }

  /**
   * Abstracted logic for burning used by both `burn` and `burnFrom`
   */
  function doBurn(address who, uint256 value) internal {
    require(value <= balances[who], "Error:Burnable.doBurn - Value exceeds available balance");
    // no need to require value <= supply, since that would imply the
    // sender's balance is greater than the supply, which *should* be an assertion failure

    balances[who] = balances[who].sub(value);
    supply = supply.sub(value);
    emit Burn(who, value);
    emit Transfer(who, address(0), value);
  }
}
