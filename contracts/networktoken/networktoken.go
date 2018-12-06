// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package networktoken

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// NetworkTokenABI is the input ABI used to generate the binding from.
const NetworkTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialAccount\",\"type\":\"address\"},{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ApprovalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferEvent\",\"type\":\"event\"}]"

// NetworkTokenBin is the compiled bytecode used for deploying new contracts.
const NetworkTokenBin = `0x608060405234801561001057600080fd5b506040516040806109798339810180604052604081101561003057600080fd5b508051602091820151600160a060020a0390911660009081526001909252604082208190559055610913806100666000396000f3fe60806040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009257806318160ddd146100cd57806323b872dd146100f4578063661884631461013757806370a0823114610170578063a9059cbb146101a3578063d73dd623146101dc578063dd62ed3e14610215575b600080fd5b34801561009e57600080fd5b506100cb600480360360408110156100b557600080fd5b50600160a060020a038135169060200135610250565b005b3480156100d957600080fd5b506100e26102b2565b60408051918252519081900360200190f35b34801561010057600080fd5b506100cb6004803603606081101561011757600080fd5b50600160a060020a038135811691602081013590911690604001356102b8565b34801561014357600080fd5b506100cb6004803603604081101561015a57600080fd5b50600160a060020a03813516906020013561057d565b34801561017c57600080fd5b506100e26004803603602081101561019357600080fd5b5035600160a060020a031661068a565b3480156101af57600080fd5b506100cb600480360360408110156101c657600080fd5b50600160a060020a0381351690602001356106a5565b3480156101e857600080fd5b506100cb600480360360408110156101ff57600080fd5b50600160a060020a038135169060200135610863565b34801561022157600080fd5b506100e26004803603604081101561023857600080fd5b50600160a060020a0381358116916020013516610897565b336000818152600260209081526040808320600160a060020a03871680855290835292819020859055805185815290519293927f08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e929181900390910190a35050565b60005490565b600160a060020a038216151561033e576040805160e560020a62461bcd02815260206004820152603c60248201527f4572726f723a5374616e646172642e7472616e7366657246726f6d202d20277460448201527f6f272061646472657373206d7573742062652073706563696669656400000000606482015290519081900360840190fd5b600160a060020a0383166000908152600160205260409020548111156103d4576040805160e560020a62461bcd02815260206004820152603e60248201527f4572726f723a5374616e646172642e7472616e7366657246726f6d202d20416d60448201527f6f756e74206578636565647320617661696c61626c652062616c616e63650000606482015290519081900360840190fd5b600160a060020a0383166000908152600260209081526040808320338452909152902054811115610475576040805160e560020a62461bcd02815260206004820152603b60248201527f4572726f722e5374616e646172642e7472616e7366657246726f6d202d20416d60448201527f6f756e74206578636565647320616c6c6f77656420616d6f756e740000000000606482015290519081900360840190fd5b600160a060020a03831660009081526001602052604090205461049e908263ffffffff6108c216565b600160a060020a0380851660009081526001602052604080822093909355908416815220546104d3908263ffffffff6108d416565b600160a060020a038084166000908152600160209081526040808320949094559186168152600282528281203382529091522054610517908263ffffffff6108c216565b600160a060020a03808516600081815260026020908152604080832033845282529182902094909455805185815290519286169391927feaf1c4b3ce0f4f62a2bae7eb3e68225c75f7e6ff4422073b7437b9a78d25f170929181900390910190a3505050565b336000908152600260209081526040808320600160a060020a03861684529091529020548111156105d157336000908152600260209081526040808320600160a060020a038616845290915281205561062a565b336000908152600260209081526040808320600160a060020a0386168452909152902054610605908263ffffffff6108c216565b336000908152600260209081526040808320600160a060020a03871684529091529020555b336000818152600260209081526040808320600160a060020a0387168085529083529281902054815190815290519293927f08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e929181900390910190a35050565b600160a060020a031660009081526001602052604090205490565b600160a060020a038216151561072b576040805160e560020a62461bcd02815260206004820152603560248201527f4572726f723a42617369632e7472616e73666572202d2027746f27206164647260448201527f657373206d757374206265207370656369666965640000000000000000000000606482015290519081900360840190fd5b336000908152600160205260409020548111156107b8576040805160e560020a62461bcd02815260206004820152603f60248201527f4572726f723a42617369632e7472616e73666572202d20416d6f756e7420657860448201527f6365656473207468652062616c616e6365206f66206d73672e73656e64657200606482015290519081900360840190fd5b336000908152600160205260409020546107d8908263ffffffff6108c216565b3360009081526001602052604080822092909255600160a060020a0384168152205461080a908263ffffffff6108d416565b600160a060020a0383166000818152600160209081526040918290209390935580518481529051919233927feaf1c4b3ce0f4f62a2bae7eb3e68225c75f7e6ff4422073b7437b9a78d25f1709281900390910190a35050565b336000908152600260209081526040808320600160a060020a0386168452909152902054610605908263ffffffff6108d416565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b6000828211156108ce57fe5b50900390565b818101828110156108e157fe5b9291505056fea165627a7a723058201f8df9679cb94fb57c17f215c7d60a1adc057ff8534491ac8ca9b903d325d6220029`

// DeployNetworkToken deploys a new Ethereum contract, binding an instance of NetworkToken to it.
func DeployNetworkToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialAccount common.Address, initialBalance *big.Int) (common.Address, *types.Transaction, *NetworkToken, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NetworkTokenBin), backend, initialAccount, initialBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NetworkToken{NetworkTokenCaller: NetworkTokenCaller{contract: contract}, NetworkTokenTransactor: NetworkTokenTransactor{contract: contract}, NetworkTokenFilterer: NetworkTokenFilterer{contract: contract}}, nil
}

// NetworkToken is an auto generated Go binding around an Ethereum contract.
type NetworkToken struct {
	NetworkTokenCaller     // Read-only binding to the contract
	NetworkTokenTransactor // Write-only binding to the contract
	NetworkTokenFilterer   // Log filterer for contract events
}

// NetworkTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NetworkTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkTokenSession struct {
	Contract     *NetworkToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NetworkTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkTokenCallerSession struct {
	Contract *NetworkTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NetworkTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkTokenTransactorSession struct {
	Contract     *NetworkTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NetworkTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkTokenRaw struct {
	Contract *NetworkToken // Generic contract binding to access the raw methods on
}

// NetworkTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkTokenCallerRaw struct {
	Contract *NetworkTokenCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkTokenTransactorRaw struct {
	Contract *NetworkTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetworkToken creates a new instance of NetworkToken, bound to a specific deployed contract.
func NewNetworkToken(address common.Address, backend bind.ContractBackend) (*NetworkToken, error) {
	contract, err := bindNetworkToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NetworkToken{NetworkTokenCaller: NetworkTokenCaller{contract: contract}, NetworkTokenTransactor: NetworkTokenTransactor{contract: contract}, NetworkTokenFilterer: NetworkTokenFilterer{contract: contract}}, nil
}

// NewNetworkTokenCaller creates a new read-only instance of NetworkToken, bound to a specific deployed contract.
func NewNetworkTokenCaller(address common.Address, caller bind.ContractCaller) (*NetworkTokenCaller, error) {
	contract, err := bindNetworkToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkTokenCaller{contract: contract}, nil
}

// NewNetworkTokenTransactor creates a new write-only instance of NetworkToken, bound to a specific deployed contract.
func NewNetworkTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkTokenTransactor, error) {
	contract, err := bindNetworkToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkTokenTransactor{contract: contract}, nil
}

// NewNetworkTokenFilterer creates a new log filterer instance of NetworkToken, bound to a specific deployed contract.
func NewNetworkTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*NetworkTokenFilterer, error) {
	contract, err := bindNetworkToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NetworkTokenFilterer{contract: contract}, nil
}

// bindNetworkToken binds a generic wrapper to an already deployed contract.
func bindNetworkToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkToken *NetworkTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NetworkToken.Contract.NetworkTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkToken *NetworkTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkToken.Contract.NetworkTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkToken *NetworkTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkToken.Contract.NetworkTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkToken *NetworkTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NetworkToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkToken *NetworkTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkToken *NetworkTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(holder address, spender address) constant returns(uint256)
func (_NetworkToken *NetworkTokenCaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkToken.contract.Call(opts, out, "allowance", holder, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(holder address, spender address) constant returns(uint256)
func (_NetworkToken *NetworkTokenSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _NetworkToken.Contract.Allowance(&_NetworkToken.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(holder address, spender address) constant returns(uint256)
func (_NetworkToken *NetworkTokenCallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _NetworkToken.Contract.Allowance(&_NetworkToken.CallOpts, holder, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(holder address) constant returns(uint256)
func (_NetworkToken *NetworkTokenCaller) BalanceOf(opts *bind.CallOpts, holder common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkToken.contract.Call(opts, out, "balanceOf", holder)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(holder address) constant returns(uint256)
func (_NetworkToken *NetworkTokenSession) BalanceOf(holder common.Address) (*big.Int, error) {
	return _NetworkToken.Contract.BalanceOf(&_NetworkToken.CallOpts, holder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(holder address) constant returns(uint256)
func (_NetworkToken *NetworkTokenCallerSession) BalanceOf(holder common.Address) (*big.Int, error) {
	return _NetworkToken.Contract.BalanceOf(&_NetworkToken.CallOpts, holder)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NetworkToken *NetworkTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NetworkToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NetworkToken *NetworkTokenSession) TotalSupply() (*big.Int, error) {
	return _NetworkToken.Contract.TotalSupply(&_NetworkToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NetworkToken *NetworkTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _NetworkToken.Contract.TotalSupply(&_NetworkToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns()
func (_NetworkToken *NetworkTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns()
func (_NetworkToken *NetworkTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.Contract.Approve(&_NetworkToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns()
func (_NetworkToken *NetworkTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.Contract.Approve(&_NetworkToken.TransactOpts, spender, amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, amount uint256) returns()
func (_NetworkToken *NetworkTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.contract.Transact(opts, "decreaseApproval", spender, amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, amount uint256) returns()
func (_NetworkToken *NetworkTokenSession) DecreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.Contract.DecreaseApproval(&_NetworkToken.TransactOpts, spender, amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, amount uint256) returns()
func (_NetworkToken *NetworkTokenTransactorSession) DecreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.Contract.DecreaseApproval(&_NetworkToken.TransactOpts, spender, amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, amount uint256) returns()
func (_NetworkToken *NetworkTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.contract.Transact(opts, "increaseApproval", spender, amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, amount uint256) returns()
func (_NetworkToken *NetworkTokenSession) IncreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.Contract.IncreaseApproval(&_NetworkToken.TransactOpts, spender, amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, amount uint256) returns()
func (_NetworkToken *NetworkTokenTransactorSession) IncreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.Contract.IncreaseApproval(&_NetworkToken.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, amount uint256) returns()
func (_NetworkToken *NetworkTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, amount uint256) returns()
func (_NetworkToken *NetworkTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.Contract.Transfer(&_NetworkToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, amount uint256) returns()
func (_NetworkToken *NetworkTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.Contract.Transfer(&_NetworkToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, amount uint256) returns()
func (_NetworkToken *NetworkTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, amount uint256) returns()
func (_NetworkToken *NetworkTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.Contract.TransferFrom(&_NetworkToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, amount uint256) returns()
func (_NetworkToken *NetworkTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkToken.Contract.TransferFrom(&_NetworkToken.TransactOpts, from, to, amount)
}

// NetworkTokenApprovalEventIterator is returned from FilterApprovalEvent and is used to iterate over the raw logs and unpacked data for ApprovalEvent events raised by the NetworkToken contract.
type NetworkTokenApprovalEventIterator struct {
	Event *NetworkTokenApprovalEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NetworkTokenApprovalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkTokenApprovalEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NetworkTokenApprovalEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NetworkTokenApprovalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkTokenApprovalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkTokenApprovalEvent represents a ApprovalEvent event raised by the NetworkToken contract.
type NetworkTokenApprovalEvent struct {
	Holder  common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApprovalEvent is a free log retrieval operation binding the contract event 0x08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e.
//
// Solidity: e ApprovalEvent(holder indexed address, spender indexed address, amount uint256)
func (_NetworkToken *NetworkTokenFilterer) FilterApprovalEvent(opts *bind.FilterOpts, holder []common.Address, spender []common.Address) (*NetworkTokenApprovalEventIterator, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NetworkToken.contract.FilterLogs(opts, "ApprovalEvent", holderRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NetworkTokenApprovalEventIterator{contract: _NetworkToken.contract, event: "ApprovalEvent", logs: logs, sub: sub}, nil
}

// WatchApprovalEvent is a free log subscription operation binding the contract event 0x08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e.
//
// Solidity: e ApprovalEvent(holder indexed address, spender indexed address, amount uint256)
func (_NetworkToken *NetworkTokenFilterer) WatchApprovalEvent(opts *bind.WatchOpts, sink chan<- *NetworkTokenApprovalEvent, holder []common.Address, spender []common.Address) (event.Subscription, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NetworkToken.contract.WatchLogs(opts, "ApprovalEvent", holderRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkTokenApprovalEvent)
				if err := _NetworkToken.contract.UnpackLog(event, "ApprovalEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// NetworkTokenTransferEventIterator is returned from FilterTransferEvent and is used to iterate over the raw logs and unpacked data for TransferEvent events raised by the NetworkToken contract.
type NetworkTokenTransferEventIterator struct {
	Event *NetworkTokenTransferEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NetworkTokenTransferEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkTokenTransferEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NetworkTokenTransferEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NetworkTokenTransferEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkTokenTransferEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkTokenTransferEvent represents a TransferEvent event raised by the NetworkToken contract.
type NetworkTokenTransferEvent struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferEvent is a free log retrieval operation binding the contract event 0xeaf1c4b3ce0f4f62a2bae7eb3e68225c75f7e6ff4422073b7437b9a78d25f170.
//
// Solidity: e TransferEvent(from indexed address, to indexed address, amount uint256)
func (_NetworkToken *NetworkTokenFilterer) FilterTransferEvent(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NetworkTokenTransferEventIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NetworkToken.contract.FilterLogs(opts, "TransferEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NetworkTokenTransferEventIterator{contract: _NetworkToken.contract, event: "TransferEvent", logs: logs, sub: sub}, nil
}

// WatchTransferEvent is a free log subscription operation binding the contract event 0xeaf1c4b3ce0f4f62a2bae7eb3e68225c75f7e6ff4422073b7437b9a78d25f170.
//
// Solidity: e TransferEvent(from indexed address, to indexed address, amount uint256)
func (_NetworkToken *NetworkTokenFilterer) WatchTransferEvent(opts *bind.WatchOpts, sink chan<- *NetworkTokenTransferEvent, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NetworkToken.contract.WatchLogs(opts, "TransferEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkTokenTransferEvent)
				if err := _NetworkToken.contract.UnpackLog(event, "TransferEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582089a940306f69292337d84ef032ffbab1c6a08ad82e8342c5675a54f479fb70550029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
