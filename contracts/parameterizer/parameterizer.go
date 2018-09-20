// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package parameterizer

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

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ApprovalEvent\",\"type\":\"event\"}]"

// IERC20Bin is the compiled bytecode used for deploying new contracts.
const IERC20Bin = `0x`

// DeployIERC20 deploys a new Ethereum contract, binding an instance of IERC20 to it.
func DeployIERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(holder address, spender address) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", holder, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(holder address, spender address) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(holder address, spender address) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, holder, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(holder address) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, holder common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", holder)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(holder address) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(holder common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, holder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(holder address) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(holder common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, holder)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalEventIterator is returned from FilterApprovalEvent and is used to iterate over the raw logs and unpacked data for ApprovalEvent events raised by the IERC20 contract.
type IERC20ApprovalEventIterator struct {
	Event *IERC20ApprovalEvent // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20ApprovalEvent)
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
		it.Event = new(IERC20ApprovalEvent)
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
func (it *IERC20ApprovalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20ApprovalEvent represents a ApprovalEvent event raised by the IERC20 contract.
type IERC20ApprovalEvent struct {
	Holder  common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApprovalEvent is a free log retrieval operation binding the contract event 0x08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e.
//
// Solidity: e ApprovalEvent(holder indexed address, spender indexed address, value uint256)
func (_IERC20 *IERC20Filterer) FilterApprovalEvent(opts *bind.FilterOpts, holder []common.Address, spender []common.Address) (*IERC20ApprovalEventIterator, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "ApprovalEvent", holderRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalEventIterator{contract: _IERC20.contract, event: "ApprovalEvent", logs: logs, sub: sub}, nil
}

// WatchApprovalEvent is a free log subscription operation binding the contract event 0x08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e.
//
// Solidity: e ApprovalEvent(holder indexed address, spender indexed address, value uint256)
func (_IERC20 *IERC20Filterer) WatchApprovalEvent(opts *bind.WatchOpts, sink chan<- *IERC20ApprovalEvent, holder []common.Address, spender []common.Address) (event.Subscription, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "ApprovalEvent", holderRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20ApprovalEvent)
				if err := _IERC20.contract.UnpackLog(event, "ApprovalEvent", log); err != nil {
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

// IERC20TransferEventIterator is returned from FilterTransferEvent and is used to iterate over the raw logs and unpacked data for TransferEvent events raised by the IERC20 contract.
type IERC20TransferEventIterator struct {
	Event *IERC20TransferEvent // Event containing the contract specifics and raw log

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
func (it *IERC20TransferEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20TransferEvent)
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
		it.Event = new(IERC20TransferEvent)
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
func (it *IERC20TransferEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20TransferEvent represents a TransferEvent event raised by the IERC20 contract.
type IERC20TransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferEvent is a free log retrieval operation binding the contract event 0xeaf1c4b3ce0f4f62a2bae7eb3e68225c75f7e6ff4422073b7437b9a78d25f170.
//
// Solidity: e TransferEvent(from indexed address, to indexed address, value uint256)
func (_IERC20 *IERC20Filterer) FilterTransferEvent(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferEventIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "TransferEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferEventIterator{contract: _IERC20.contract, event: "TransferEvent", logs: logs, sub: sub}, nil
}

// WatchTransferEvent is a free log subscription operation binding the contract event 0xeaf1c4b3ce0f4f62a2bae7eb3e68225c75f7e6ff4422073b7437b9a78d25f170.
//
// Solidity: e TransferEvent(from indexed address, to indexed address, value uint256)
func (_IERC20 *IERC20Filterer) WatchTransferEvent(opts *bind.WatchOpts, sink chan<- *IERC20TransferEvent, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "TransferEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20TransferEvent)
				if err := _IERC20.contract.UnpackLog(event, "TransferEvent", log); err != nil {
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

// PLCRVotingABI is the input ABI used to generate the binding from.
const PLCRVotingABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getTotalNumberOfTokensForWinningOption\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getPollRevealExpiry\",\"outputs\":[{\"name\":\"expiry\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"numTokens\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getInsertPointForNumTokens\",\"outputs\":[{\"name\":\"prevNode\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"name\":\"commitDuration\",\"type\":\"uint256\"},{\"name\":\"revealDuration\",\"type\":\"uint256\"}],\"name\":\"startPoll\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getLastNode\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"revealPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isPassed\",\"outputs\":[{\"name\":\"passed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getLockedTokens\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"secretHash\",\"type\":\"bytes32\"},{\"name\":\"numTokens\",\"type\":\"uint256\"},{\"name\":\"prevId\",\"type\":\"uint256\"}],\"name\":\"commitVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"didCommit\",\"outputs\":[{\"name\":\"committed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"prevId\",\"type\":\"uint256\"},{\"name\":\"nextId\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"validPosition\",\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"pollExists\",\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"rescueTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"requestVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"commitPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"didReveal\",\"outputs\":[{\"name\":\"revealed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"voteOption\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"revealVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getNumPassingTokens\",\"outputs\":[{\"name\":\"correctVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getNumTokens\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getCommitHash\",\"outputs\":[{\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"terminationDate\",\"type\":\"uint256\"}],\"name\":\"isExpired\",\"outputs\":[{\"name\":\"expired\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getPollCommitExpiry\",\"outputs\":[{\"name\":\"expiry\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"withdrawVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"pollEnded\",\"outputs\":[{\"name\":\"ended\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VoteCommittedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesFor\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesAgainst\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"choice\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VoteRevealedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commitExpiry\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealExpiry\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"PollCreatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VotingRightsGrantedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VotingRightsWithdrawnEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"TokensRescuedEvent\",\"type\":\"event\"}]"

// PLCRVotingBin is the compiled bytecode used for deploying new contracts.
const PLCRVotingBin = `0x608060405234801561001057600080fd5b50604051602080612592833981016040525160058054600160a060020a03909216600160a060020a0319909216919091179055600060025561253b806100576000396000f30060806040526004361061013d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663053e71a681146101425780631be8ea501461016c5780632c0520311461018457806332ed3d60146101ab578063427fa1d2146101c9578063441c77c0146101ea57806349403183146102165780636b2d95d41461022e5780636cbf9c5e1461024f5780637f97e83614610272578063819b02931461029657806388d21ff3146102c057806397603560146102d8578063a25236fe146102f0578063a4439dc514610308578063aa7ca46414610320578063b11d8bb814610344578063b43bd06914610362578063d138209214610389578063d901402b146103ad578063d9548e53146103d1578063dc1b946f146103e9578063e7b1d43c14610401578063ee68483014610419575b600080fd5b34801561014e57600080fd5b5061015a600435610431565b60408051918252519081900360200190f35b34801561017857600080fd5b5061015a60043561051e565b34801561019057600080fd5b5061015a600160a060020a03600435166024356044356105bb565b3480156101b757600080fd5b5061015a60043560243560443561062c565b3480156101d557600080fd5b5061015a600160a060020a036004351661071c565b3480156101f657600080fd5b5061020260043561072f565b604080519115158252519081900360200190f35b34801561022257600080fd5b506102026004356107e9565b34801561023a57600080fd5b5061015a600160a060020a036004351661090c565b34801561025b57600080fd5b50610270600435602435604435606435610920565b005b34801561027e57600080fd5b50610202600160a060020a0360043516602435610ce5565b3480156102a257600080fd5b50610202600435602435600160a060020a0360443516606435610d9b565b3480156102cc57600080fd5b50610202600435610ddc565b3480156102e457600080fd5b50610270600435610df1565b3480156102fc57600080fd5b50610270600435610f48565b34801561031457600080fd5b50610202600435611203565b34801561032c57600080fd5b50610202600160a060020a03600435166024356112a2565b34801561035057600080fd5b50610270600435602435604435611358565b34801561036e57600080fd5b5061015a600160a060020a0360043516602435604435611757565b34801561039557600080fd5b5061015a600160a060020a03600435166024356119de565b3480156103b957600080fd5b5061015a600160a060020a0360043516602435611a2d565b3480156103dd57600080fd5b50610202600435611a75565b3480156103f557600080fd5b5061015a600435611a7a565b34801561040d57600080fd5b50610270600435611b14565b34801561042557600080fd5b50610202600435611d67565b600061043c82611d67565b15156104de576040805160e560020a62461bcd02815260206004820152604a60248201527f4572726f723a566f74696e672e676574546f74616c4e756d6265724f66546f6b60448201527f656e73466f7257696e6e696e674f7074696f6e202d20506f6c6c206d7573742060648201527f6861766520656e64656400000000000000000000000000000000000000000000608482015290519081900360a40190fd5b6104e7826107e9565b15610505575060008181526003602081905260409091200154610519565b506000818152600360205260409020600401545b919050565b600061052982610ddc565b15156105a5576040805160e560020a62461bcd02815260206004820152603260248201527f4572726f723a566f74696e672e676574506f6c6c72657665616c45787069727960448201527f202d20506f6c6c206d7573742065786973740000000000000000000000000000606482015290519081900360840190fd5b5060009081526003602052604090206001015490565b60008060006105c98661071c565b91506105d586836119de565b90505b811561061f576105e886836119de565b905084811161060e5783821415610606576106038683611e09565b91505b819250610623565b6106188683611e09565b91506105d8565b8192505b50509392505050565b60008060006106476001600254611e3490919063ffffffff16565b60025561065a428663ffffffff611e3416565b915061066c828563ffffffff611e3416565b6040805160a08101825284815260208082018481528284018b815260006060808601828152608087018381526002805485526003808952948a90209851895595516001890155935187860155519186019190915590516004909401939093555483518b8152918201879052818401859052925193945033937ff7e7200f3afa7bbc8aca2cfa0ab33e741504a20fed2abc59869413cba2a33400929181900390910190a35050600254949350505050565b6000610729826000611e09565b92915050565b600061073a82610ddc565b15156107b6576040805160e560020a62461bcd02815260206004820152603160248201527f4572726f723a566f74696e672e72657665616c506572696f644163746976652060448201527f2d20506f6c6c206d757374206578697374000000000000000000000000000000606482015290519081900360840190fd5b6000828152600360205260409020600101546107d190611a75565b15801561072957506107e282611203565b1592915050565b60006107f361249f565b6107fc83611d67565b1515610878576040805160e560020a62461bcd02815260206004820152602c60248201527f4572726f723a566f74696e672e6973506173736564202d20506f6c6c206d757360448201527f74206861766520656e6465640000000000000000000000000000000000000000606482015290519081900360840190fd5b50600082815260036020818152604092839020835160a0810185528154815260018201549281019290925260028101549382019390935290820154606082018190526004909201546080820181905290916108ee916108dd919063ffffffff611e3416565b60408301519063ffffffff611e4116565b606082015161090490606463ffffffff611e4116565b119392505050565b60006107298261091b8461071c565b6119de565b60008061092c86611203565b15156109a8576040805160e560020a62461bcd02815260206004820152603460248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d2054617267657460448201527f20706f6c6c206d75737420626520616374697665000000000000000000000000606482015290519081900360840190fd5b33600090815260046020526040902054841115610a35576040805160e560020a62461bcd02815260206004820152603160248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d2053656e64657260448201527f2063616e6e6f74206f7665727370656e64000000000000000000000000000000606482015290519081900360840190fd5b851515610ab2576040805160e560020a62461bcd02815260206004820152602d60248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d20506f6c6c204960448201527f442063616e6e6f74206265203000000000000000000000000000000000000000606482015290519081900360840190fd5b821580610ac45750610ac43384611e6a565b1515610b40576040805160e560020a62461bcd02815260206004820152602960248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d20506f6c6c206d60448201527f7573742065786973740000000000000000000000000000000000000000000000606482015290519081900360840190fd5b610b4a3384611ee8565b915085821415610b6157610b5e3387611ee8565b91505b610b6d83833387610d9b565b1515610be9576040805160e560020a62461bcd02815260206004820152603460248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d20506f6c6c207060448201527f6f736974696f6e206973206e6f742076616c6964000000000000000000000000606482015290519081900360840190fd5b610bf533848885611f10565b610bff3387612174565b9050610c41816040805190810160405280600981526020017f6e756d546f6b656e73000000000000000000000000000000000000000000000081525086612221565b60408051808201909152600a81527f636f6d6d697448617368000000000000000000000000000000000000000000006020820152610c8190829087612221565b6000868152600360209081526040808320338085526005909101835292819020805460ff191660011790558051878152905189927f673f4a020a46f562a6e9fee9c8d4a179a11924190b777bec160133dd00c9f1de928290030190a3505050505050565b6000610cf082610ddc565b1515610d6c576040805160e560020a62461bcd02815260206004820152602f60248201527f4572726f723a506172616d65746572697a65722e646964436f6d6d6974202d2060448201527f506f6c6c206d7573742065786973740000000000000000000000000000000000606482015290519081900360840190fd5b506000908152600360209081526040808320600160a060020a0394909416835260059093019052205460ff1690565b6000806000610daa85886119de565b8410159150610db985876119de565b84111580610dc5575085155b9050818015610dd15750805b979650505050505050565b60008115801590610729575050600254101590565b600081815260036020526040902060010154610e0c90611a75565b1515610e88576040805160e560020a62461bcd02815260206004820152603060248201527f4572726f723a566f74696e672e726573637565546f6b656e73202d20506f6c6c60448201527f206d757374206861766520656e64656400000000000000000000000000000000606482015290519081900360840190fd5b610e923382611e6a565b1515610f0e576040805160e560020a62461bcd02815260206004820152602a60248201527f4572726f723a566f74696e672e726573637565546f6b656e73202d20506f6c6c60448201527f206e6f7420666f756e6400000000000000000000000000000000000000000000606482015290519081900360840190fd5b610f183382612305565b604051339082907f1e897e4fe81e0db0459f0219545995fde305b17af5ca5cba7d6b6e8738df5dd090600090a350565b600554604080517f70a0823100000000000000000000000000000000000000000000000000000000815233600482015290518392600160a060020a0316916370a082319160248083019260209291908290030181600087803b158015610fad57600080fd5b505af1158015610fc1573d6000803e3d6000fd5b505050506040513d6020811015610fd757600080fd5b5051101561107b576040805160e560020a62461bcd02815260206004820152604860248201527f4572726f723a566f74696e672e72657175657374566f74696e6752696768747360448201527f202d2053656e64657220646f6573206e6f74206861766520737566666963696560648201527f6e742066756e6473000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b3360009081526004602052604090205461109b908263ffffffff611e3416565b3360008181526004602081815260408084209590955560055485517f23b872dd00000000000000000000000000000000000000000000000000000000815292830194909452306024830152604482018690529351600160a060020a03909316936323b872dd9360648084019492939192918390030190829087803b15801561112257600080fd5b505af1158015611136573d6000803e3d6000fd5b505050506040513d602081101561114c57600080fd5b505115156111ca576040805160e560020a62461bcd02815260206004820152603d60248201527f4572726f723a566f74696e672e72657175657374566f74696e6752696768747360448201527f202d20546f6b656e2e7472616e7366657246726f6d206661696c757265000000606482015290519081900360840190fd5b60408051828152905133917f0fbea84792de88539ad3d7f184543a65916440f4c6f9e2ef136c97f13109c45d919081900360200190a250565b600061120e82610ddc565b151561128a576040805160e560020a62461bcd02815260206004820152603160248201527f4572726f723a566f74696e672e636f6d6d6974506572696f644163746976652060448201527f2d20506f6c6c206d757374206578697374000000000000000000000000000000606482015290519081900360840190fd5b6000828152600360205260409020546107e290611a75565b60006112ad82610ddc565b1515611329576040805160e560020a62461bcd02815260206004820152602860248201527f4572726f723a566f74696e672e64696452657665616c202d20506f6c6c206d7560448201527f7374206578697374000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b506000908152600360209081526040808320600160a060020a0394909416835260069093019052205460ff1690565b60006113638461072f565b15156113df576040805160e560020a62461bcd02815260206004820152603660248201527f4572726f723a566f74696e672e72657665616c566f7465202d2052657665616c60448201527f20706572696f64206d7573742062652061637469766500000000000000000000606482015290519081900360840190fd5b600084815260036020908152604080832033845260050190915290205460ff16151561147b576040805160e560020a62461bcd02815260206004820152603f60248201527f4572726f723a566f74696e672e72657665616c566f7465202d2053656e64657260448201527f206d757374206861766520636f6d6d697474656420746865697220766f746500606482015290519081900360840190fd5b600084815260036020908152604080832033845260060190915290205460ff161561153c576040805160e560020a62461bcd02815260206004820152604860248201527f4572726f723a566f74696e672e72657665616c566f7465202d2053656e64657260448201527f2063616e6e6f742072657665616c20746865697220766f7465206d756c74697060648201527f6c652074696d6573000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b6115463385611a2d565b604080516020808201879052818301869052825180830384018152606090920192839052815191929182918401908083835b602083106115975780518252601f199092019160209182019101611578565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916141515611645576040805160e560020a62461bcd02815260206004820152602760248201527f4572726f723a566f74696e672e72657665616c566f7465202d2048617368206d60448201527f69736d6174636800000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b61164f33856119de565b90508260011415611695576000848152600360208190526040909120015461167d908263ffffffff611e3416565b600085815260036020819052604090912001556116ca565b6000848152600360205260409020600401546116b7908263ffffffff611e3416565b6000858152600360205260409020600401555b6116d43385612305565b600084815260036020818152604080842033808652600682018452828620805460ff191660011790559489905283835292830154600490930154815186815292830193909352818101929092529051859187917f448a2036d0f6dc799c1d3fdf26c4fe80832817bdbe81133808ffa6777c5826e09181900360600190a450505050565b60008060008061176686611d67565b15156117e2576040805160e560020a62461bcd02815260206004820152603760248201527f4572726f723a566f74696e672e6765744e756d50617373696e67546f6b656e7360448201527f202d20506f6c6c206d757374206861766520656e646564000000000000000000606482015290519081900360840190fd5b6000868152600360209081526040808320600160a060020a038b16845260060190915290205460ff1615156118ad576040805160e560020a62461bcd02815260206004820152604660248201527f4572726f723a566f74696e672e6765744e756d50617373696e67546f6b656e7360448201527f202d20566f746572206d75737420686176652072657665616c6564207468656960648201527f7220766f74650000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b6118b6866107e9565b6118c15760006118c4565b60015b60ff169250828560405160200180838152602001828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106119205780518252601f199092019160209182019101611901565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902091506119598787611a2d565b90508181146119d8576040805160e560020a62461bcd02815260206004820152603060248201527f4572726f723a566f74696e672e6765744e756d50617373696e67546f6b656e7360448201527f202d2048617368206d69736d6174636800000000000000000000000000000000606482015290519081900360840190fd5b610dd187875b6000611a266119ed8484612174565b60408051808201909152600981527f6e756d546f6b656e7300000000000000000000000000000000000000000000006020820152612378565b9392505050565b6000611a26611a3c8484612174565b60408051808201909152600a81527f636f6d6d697448617368000000000000000000000000000000000000000000006020820152612378565b421190565b6000611a8582610ddc565b1515611b01576040805160e560020a62461bcd02815260206004820152603260248201527f4572726f723a566f74696e672e676574506f6c6c436f6d6d697445787069727960448201527f202d20506f6c6c206d7573742065786973740000000000000000000000000000606482015290519081900360840190fd5b5060009081526003602052604090205490565b6000611b3e611b223361090c565b336000908152600460205260409020549063ffffffff61246216565b905081811015611be4576040805160e560020a62461bcd02815260206004820152604160248201527f4572726f723a566f74696e672e7769746864726177566f74696e67526967687460448201527f73202d20496e73756666696369656e7420617661696c61626c6520746f6b656e60648201527f7300000000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b33600090815260046020526040902054611c04908363ffffffff61246216565b3360008181526004602081815260408084209590955560055485517fa9059cbb00000000000000000000000000000000000000000000000000000000815292830194909452602482018790529351600160a060020a039093169363a9059cbb9360448084019492939192918390030190829087803b158015611c8557600080fd5b505af1158015611c99573d6000803e3d6000fd5b505050506040513d6020811015611caf57600080fd5b50511515611d2d576040805160e560020a62461bcd02815260206004820152603a60248201527f4572726f723a566f74696e672e7769746864726177566f74696e67526967687460448201527f73202d20546f6b656e2e7472616e73666572206661696c757265000000000000606482015290519081900360840190fd5b60408051838152905133917f029a58e9abaa250a4ec75d14b364d764e890adec2d902a297fa24a6c5dfd097b919081900360200190a25050565b6000611d7282610ddc565b1515611dee576040805160e560020a62461bcd02815260206004820152602860248201527f4572726f723a566f74696e672e706f6c6c456e646564202d20506f6c6c206d7560448201527f7374206578697374000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008281526003602052604090206001015461072990611a75565b600160a060020a03919091166000908152600160208181526040808420948452939052919020015490565b8181018281101561072957fe5b6000821515611e5257506000610729565b50818102818382811515611e6257fe5b041461072957fe5b6000806000611e7885612474565b80611e81575083155b15611e8f5760009250611ee0565b83611e9986612487565b148015611ead575083611eab86612493565b145b91506000611ebb8686611ee8565b148015611ed157506000611ecf8686611e09565b145b90508180611edd575080155b92505b505092915050565b600160a060020a03919091166000908152600160209081526040808320938352929052205490565b811515611f8d576040805160e560020a62461bcd02815260206004820152603a60248201527f4572726f723a504c4352566f74696e672e646c6c496e73657274202d2049442060448201527f6f662061206e6577206e6f6465206d6179206e6f742062652030000000000000606482015290519081900360840190fd5b611f978483612305565b821580611fa95750611fa98484611e6a565b1515612001576040805160e560020a62461bcd02815260206004820152603360248201526000805160206124f083398151915260448201526000805160206124d0833981519152606482015290519081900360840190fd5b80158061201357506120138482611e6a565b151561206b576040805160e560020a62461bcd02815260206004820152603360248201526000805160206124f083398151915260448201526000805160206124d0833981519152606482015290519081900360840190fd5b806120768585611ee8565b146120cd576040805160e560020a62461bcd02815260206004820152603360248201526000805160206124f083398151915260448201526000805160206124d0833981519152606482015290519081900360840190fd5b826120d88583611e09565b1461212f576040805160e560020a62461bcd02815260206004820152603360248201526000805160206124f083398151915260448201526000805160206124d0833981519152606482015290519081900360840190fd5b600160a060020a039390931660009081526001602081815260408084208585529091528083208083018690558690559382528382208390559381529190912090910155565b600082826040516020018083600160a060020a0316600160a060020a03166c01000000000000000000000000028152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106121ee5780518252601f1990920191602091820191016121cf565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091209695505050505050565b604080516020808201868152855160009488948894910191908401908083835b602083106122605780518252601f199092019160209182019101612241565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b602083106122c45780518252601f1990920191602091820191016122a5565b51815160209384036101000a600019018019909216911617905260408051929094018290039091206000908152908190529190912094909455505050505050565b6000806123128484611e6a565b151561231d57612372565b6123278484611ee8565b91506123338484611e09565b600160a060020a038516600090815260016020818152604080842087855290915280832082018490558383528083208690558683528220828155015590505b50505050565b600080838360405160200180836000191660001916815260200182805190602001908083835b602083106123bd5780518252601f19909201916020918201910161239e565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b602083106124215780518252601f199092019160209182019101612402565b51815160209384036101000a600019018019909216911617905260408051929094018290039091206000908152908190529190912054979650505050505050565b60008282111561246e57fe5b50900390565b60008061248083612487565b1492915050565b60006107298282611ee8565b60006107298282611e09565b60a0604051908101604052806000815260200160008152602001600081526020016000815260200160008152509056006e6f7420696e7365727420696e746f20444c4c000000000000000000000000004572726f723a504c4352566f74696e672e644c4c496e73657274202d20436f6ea165627a7a72305820b31d57707843bf73d71bed033124c471c24ca1597ff3d1132c32ff6034a091c00029`

// DeployPLCRVoting deploys a new Ethereum contract, binding an instance of PLCRVoting to it.
func DeployPLCRVoting(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddr common.Address) (common.Address, *types.Transaction, *PLCRVoting, error) {
	parsed, err := abi.JSON(strings.NewReader(PLCRVotingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PLCRVotingBin), backend, tokenAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PLCRVoting{PLCRVotingCaller: PLCRVotingCaller{contract: contract}, PLCRVotingTransactor: PLCRVotingTransactor{contract: contract}, PLCRVotingFilterer: PLCRVotingFilterer{contract: contract}}, nil
}

// PLCRVoting is an auto generated Go binding around an Ethereum contract.
type PLCRVoting struct {
	PLCRVotingCaller     // Read-only binding to the contract
	PLCRVotingTransactor // Write-only binding to the contract
	PLCRVotingFilterer   // Log filterer for contract events
}

// PLCRVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PLCRVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PLCRVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PLCRVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PLCRVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PLCRVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PLCRVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PLCRVotingSession struct {
	Contract     *PLCRVoting       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PLCRVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PLCRVotingCallerSession struct {
	Contract *PLCRVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PLCRVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PLCRVotingTransactorSession struct {
	Contract     *PLCRVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PLCRVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PLCRVotingRaw struct {
	Contract *PLCRVoting // Generic contract binding to access the raw methods on
}

// PLCRVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PLCRVotingCallerRaw struct {
	Contract *PLCRVotingCaller // Generic read-only contract binding to access the raw methods on
}

// PLCRVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PLCRVotingTransactorRaw struct {
	Contract *PLCRVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPLCRVoting creates a new instance of PLCRVoting, bound to a specific deployed contract.
func NewPLCRVoting(address common.Address, backend bind.ContractBackend) (*PLCRVoting, error) {
	contract, err := bindPLCRVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PLCRVoting{PLCRVotingCaller: PLCRVotingCaller{contract: contract}, PLCRVotingTransactor: PLCRVotingTransactor{contract: contract}, PLCRVotingFilterer: PLCRVotingFilterer{contract: contract}}, nil
}

// NewPLCRVotingCaller creates a new read-only instance of PLCRVoting, bound to a specific deployed contract.
func NewPLCRVotingCaller(address common.Address, caller bind.ContractCaller) (*PLCRVotingCaller, error) {
	contract, err := bindPLCRVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingCaller{contract: contract}, nil
}

// NewPLCRVotingTransactor creates a new write-only instance of PLCRVoting, bound to a specific deployed contract.
func NewPLCRVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*PLCRVotingTransactor, error) {
	contract, err := bindPLCRVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingTransactor{contract: contract}, nil
}

// NewPLCRVotingFilterer creates a new log filterer instance of PLCRVoting, bound to a specific deployed contract.
func NewPLCRVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*PLCRVotingFilterer, error) {
	contract, err := bindPLCRVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingFilterer{contract: contract}, nil
}

// bindPLCRVoting binds a generic wrapper to an already deployed contract.
func bindPLCRVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PLCRVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PLCRVoting *PLCRVotingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PLCRVoting.Contract.PLCRVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PLCRVoting *PLCRVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PLCRVoting.Contract.PLCRVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PLCRVoting *PLCRVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PLCRVoting.Contract.PLCRVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PLCRVoting *PLCRVotingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PLCRVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PLCRVoting *PLCRVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PLCRVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PLCRVoting *PLCRVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PLCRVoting.Contract.contract.Transact(opts, method, params...)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(id uint256) constant returns(active bool)
func (_PLCRVoting *PLCRVotingCaller) CommitPeriodActive(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "commitPeriodActive", id)
	return *ret0, err
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(id uint256) constant returns(active bool)
func (_PLCRVoting *PLCRVotingSession) CommitPeriodActive(id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.CommitPeriodActive(&_PLCRVoting.CallOpts, id)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(id uint256) constant returns(active bool)
func (_PLCRVoting *PLCRVotingCallerSession) CommitPeriodActive(id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.CommitPeriodActive(&_PLCRVoting.CallOpts, id)
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(voter address, id uint256) constant returns(committed bool)
func (_PLCRVoting *PLCRVotingCaller) DidCommit(opts *bind.CallOpts, voter common.Address, id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "didCommit", voter, id)
	return *ret0, err
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(voter address, id uint256) constant returns(committed bool)
func (_PLCRVoting *PLCRVotingSession) DidCommit(voter common.Address, id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.DidCommit(&_PLCRVoting.CallOpts, voter, id)
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(voter address, id uint256) constant returns(committed bool)
func (_PLCRVoting *PLCRVotingCallerSession) DidCommit(voter common.Address, id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.DidCommit(&_PLCRVoting.CallOpts, voter, id)
}

// DidReveal is a free data retrieval call binding the contract method 0xaa7ca464.
//
// Solidity: function didReveal(voter address, id uint256) constant returns(revealed bool)
func (_PLCRVoting *PLCRVotingCaller) DidReveal(opts *bind.CallOpts, voter common.Address, id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "didReveal", voter, id)
	return *ret0, err
}

// DidReveal is a free data retrieval call binding the contract method 0xaa7ca464.
//
// Solidity: function didReveal(voter address, id uint256) constant returns(revealed bool)
func (_PLCRVoting *PLCRVotingSession) DidReveal(voter common.Address, id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.DidReveal(&_PLCRVoting.CallOpts, voter, id)
}

// DidReveal is a free data retrieval call binding the contract method 0xaa7ca464.
//
// Solidity: function didReveal(voter address, id uint256) constant returns(revealed bool)
func (_PLCRVoting *PLCRVotingCallerSession) DidReveal(voter common.Address, id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.DidReveal(&_PLCRVoting.CallOpts, voter, id)
}

// GetCommitHash is a free data retrieval call binding the contract method 0xd901402b.
//
// Solidity: function getCommitHash(voter address, id uint256) constant returns(commitHash bytes32)
func (_PLCRVoting *PLCRVotingCaller) GetCommitHash(opts *bind.CallOpts, voter common.Address, id *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "getCommitHash", voter, id)
	return *ret0, err
}

// GetCommitHash is a free data retrieval call binding the contract method 0xd901402b.
//
// Solidity: function getCommitHash(voter address, id uint256) constant returns(commitHash bytes32)
func (_PLCRVoting *PLCRVotingSession) GetCommitHash(voter common.Address, id *big.Int) ([32]byte, error) {
	return _PLCRVoting.Contract.GetCommitHash(&_PLCRVoting.CallOpts, voter, id)
}

// GetCommitHash is a free data retrieval call binding the contract method 0xd901402b.
//
// Solidity: function getCommitHash(voter address, id uint256) constant returns(commitHash bytes32)
func (_PLCRVoting *PLCRVotingCallerSession) GetCommitHash(voter common.Address, id *big.Int) ([32]byte, error) {
	return _PLCRVoting.Contract.GetCommitHash(&_PLCRVoting.CallOpts, voter, id)
}

// GetInsertPointForNumTokens is a free data retrieval call binding the contract method 0x2c052031.
//
// Solidity: function getInsertPointForNumTokens(voter address, numTokens uint256, id uint256) constant returns(prevNode uint256)
func (_PLCRVoting *PLCRVotingCaller) GetInsertPointForNumTokens(opts *bind.CallOpts, voter common.Address, numTokens *big.Int, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "getInsertPointForNumTokens", voter, numTokens, id)
	return *ret0, err
}

// GetInsertPointForNumTokens is a free data retrieval call binding the contract method 0x2c052031.
//
// Solidity: function getInsertPointForNumTokens(voter address, numTokens uint256, id uint256) constant returns(prevNode uint256)
func (_PLCRVoting *PLCRVotingSession) GetInsertPointForNumTokens(voter common.Address, numTokens *big.Int, id *big.Int) (*big.Int, error) {
	return _PLCRVoting.Contract.GetInsertPointForNumTokens(&_PLCRVoting.CallOpts, voter, numTokens, id)
}

// GetInsertPointForNumTokens is a free data retrieval call binding the contract method 0x2c052031.
//
// Solidity: function getInsertPointForNumTokens(voter address, numTokens uint256, id uint256) constant returns(prevNode uint256)
func (_PLCRVoting *PLCRVotingCallerSession) GetInsertPointForNumTokens(voter common.Address, numTokens *big.Int, id *big.Int) (*big.Int, error) {
	return _PLCRVoting.Contract.GetInsertPointForNumTokens(&_PLCRVoting.CallOpts, voter, numTokens, id)
}

// GetLastNode is a free data retrieval call binding the contract method 0x427fa1d2.
//
// Solidity: function getLastNode(voter address) constant returns(id uint256)
func (_PLCRVoting *PLCRVotingCaller) GetLastNode(opts *bind.CallOpts, voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "getLastNode", voter)
	return *ret0, err
}

// GetLastNode is a free data retrieval call binding the contract method 0x427fa1d2.
//
// Solidity: function getLastNode(voter address) constant returns(id uint256)
func (_PLCRVoting *PLCRVotingSession) GetLastNode(voter common.Address) (*big.Int, error) {
	return _PLCRVoting.Contract.GetLastNode(&_PLCRVoting.CallOpts, voter)
}

// GetLastNode is a free data retrieval call binding the contract method 0x427fa1d2.
//
// Solidity: function getLastNode(voter address) constant returns(id uint256)
func (_PLCRVoting *PLCRVotingCallerSession) GetLastNode(voter common.Address) (*big.Int, error) {
	return _PLCRVoting.Contract.GetLastNode(&_PLCRVoting.CallOpts, voter)
}

// GetLockedTokens is a free data retrieval call binding the contract method 0x6b2d95d4.
//
// Solidity: function getLockedTokens(voter address) constant returns(numTokens uint256)
func (_PLCRVoting *PLCRVotingCaller) GetLockedTokens(opts *bind.CallOpts, voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "getLockedTokens", voter)
	return *ret0, err
}

// GetLockedTokens is a free data retrieval call binding the contract method 0x6b2d95d4.
//
// Solidity: function getLockedTokens(voter address) constant returns(numTokens uint256)
func (_PLCRVoting *PLCRVotingSession) GetLockedTokens(voter common.Address) (*big.Int, error) {
	return _PLCRVoting.Contract.GetLockedTokens(&_PLCRVoting.CallOpts, voter)
}

// GetLockedTokens is a free data retrieval call binding the contract method 0x6b2d95d4.
//
// Solidity: function getLockedTokens(voter address) constant returns(numTokens uint256)
func (_PLCRVoting *PLCRVotingCallerSession) GetLockedTokens(voter common.Address) (*big.Int, error) {
	return _PLCRVoting.Contract.GetLockedTokens(&_PLCRVoting.CallOpts, voter)
}

// GetNumPassingTokens is a free data retrieval call binding the contract method 0xb43bd069.
//
// Solidity: function getNumPassingTokens(voter address, id uint256, salt uint256) constant returns(correctVotes uint256)
func (_PLCRVoting *PLCRVotingCaller) GetNumPassingTokens(opts *bind.CallOpts, voter common.Address, id *big.Int, salt *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "getNumPassingTokens", voter, id, salt)
	return *ret0, err
}

// GetNumPassingTokens is a free data retrieval call binding the contract method 0xb43bd069.
//
// Solidity: function getNumPassingTokens(voter address, id uint256, salt uint256) constant returns(correctVotes uint256)
func (_PLCRVoting *PLCRVotingSession) GetNumPassingTokens(voter common.Address, id *big.Int, salt *big.Int) (*big.Int, error) {
	return _PLCRVoting.Contract.GetNumPassingTokens(&_PLCRVoting.CallOpts, voter, id, salt)
}

// GetNumPassingTokens is a free data retrieval call binding the contract method 0xb43bd069.
//
// Solidity: function getNumPassingTokens(voter address, id uint256, salt uint256) constant returns(correctVotes uint256)
func (_PLCRVoting *PLCRVotingCallerSession) GetNumPassingTokens(voter common.Address, id *big.Int, salt *big.Int) (*big.Int, error) {
	return _PLCRVoting.Contract.GetNumPassingTokens(&_PLCRVoting.CallOpts, voter, id, salt)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xd1382092.
//
// Solidity: function getNumTokens(voter address, id uint256) constant returns(numTokens uint256)
func (_PLCRVoting *PLCRVotingCaller) GetNumTokens(opts *bind.CallOpts, voter common.Address, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "getNumTokens", voter, id)
	return *ret0, err
}

// GetNumTokens is a free data retrieval call binding the contract method 0xd1382092.
//
// Solidity: function getNumTokens(voter address, id uint256) constant returns(numTokens uint256)
func (_PLCRVoting *PLCRVotingSession) GetNumTokens(voter common.Address, id *big.Int) (*big.Int, error) {
	return _PLCRVoting.Contract.GetNumTokens(&_PLCRVoting.CallOpts, voter, id)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xd1382092.
//
// Solidity: function getNumTokens(voter address, id uint256) constant returns(numTokens uint256)
func (_PLCRVoting *PLCRVotingCallerSession) GetNumTokens(voter common.Address, id *big.Int) (*big.Int, error) {
	return _PLCRVoting.Contract.GetNumTokens(&_PLCRVoting.CallOpts, voter, id)
}

// GetPollCommitExpiry is a free data retrieval call binding the contract method 0xdc1b946f.
//
// Solidity: function getPollCommitExpiry(id uint256) constant returns(expiry uint256)
func (_PLCRVoting *PLCRVotingCaller) GetPollCommitExpiry(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "getPollCommitExpiry", id)
	return *ret0, err
}

// GetPollCommitExpiry is a free data retrieval call binding the contract method 0xdc1b946f.
//
// Solidity: function getPollCommitExpiry(id uint256) constant returns(expiry uint256)
func (_PLCRVoting *PLCRVotingSession) GetPollCommitExpiry(id *big.Int) (*big.Int, error) {
	return _PLCRVoting.Contract.GetPollCommitExpiry(&_PLCRVoting.CallOpts, id)
}

// GetPollCommitExpiry is a free data retrieval call binding the contract method 0xdc1b946f.
//
// Solidity: function getPollCommitExpiry(id uint256) constant returns(expiry uint256)
func (_PLCRVoting *PLCRVotingCallerSession) GetPollCommitExpiry(id *big.Int) (*big.Int, error) {
	return _PLCRVoting.Contract.GetPollCommitExpiry(&_PLCRVoting.CallOpts, id)
}

// GetPollRevealExpiry is a free data retrieval call binding the contract method 0x1be8ea50.
//
// Solidity: function getPollRevealExpiry(id uint256) constant returns(expiry uint256)
func (_PLCRVoting *PLCRVotingCaller) GetPollRevealExpiry(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "getPollRevealExpiry", id)
	return *ret0, err
}

// GetPollRevealExpiry is a free data retrieval call binding the contract method 0x1be8ea50.
//
// Solidity: function getPollRevealExpiry(id uint256) constant returns(expiry uint256)
func (_PLCRVoting *PLCRVotingSession) GetPollRevealExpiry(id *big.Int) (*big.Int, error) {
	return _PLCRVoting.Contract.GetPollRevealExpiry(&_PLCRVoting.CallOpts, id)
}

// GetPollRevealExpiry is a free data retrieval call binding the contract method 0x1be8ea50.
//
// Solidity: function getPollRevealExpiry(id uint256) constant returns(expiry uint256)
func (_PLCRVoting *PLCRVotingCallerSession) GetPollRevealExpiry(id *big.Int) (*big.Int, error) {
	return _PLCRVoting.Contract.GetPollRevealExpiry(&_PLCRVoting.CallOpts, id)
}

// GetTotalNumberOfTokensForWinningOption is a free data retrieval call binding the contract method 0x053e71a6.
//
// Solidity: function getTotalNumberOfTokensForWinningOption(id uint256) constant returns(numTokens uint256)
func (_PLCRVoting *PLCRVotingCaller) GetTotalNumberOfTokensForWinningOption(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "getTotalNumberOfTokensForWinningOption", id)
	return *ret0, err
}

// GetTotalNumberOfTokensForWinningOption is a free data retrieval call binding the contract method 0x053e71a6.
//
// Solidity: function getTotalNumberOfTokensForWinningOption(id uint256) constant returns(numTokens uint256)
func (_PLCRVoting *PLCRVotingSession) GetTotalNumberOfTokensForWinningOption(id *big.Int) (*big.Int, error) {
	return _PLCRVoting.Contract.GetTotalNumberOfTokensForWinningOption(&_PLCRVoting.CallOpts, id)
}

// GetTotalNumberOfTokensForWinningOption is a free data retrieval call binding the contract method 0x053e71a6.
//
// Solidity: function getTotalNumberOfTokensForWinningOption(id uint256) constant returns(numTokens uint256)
func (_PLCRVoting *PLCRVotingCallerSession) GetTotalNumberOfTokensForWinningOption(id *big.Int) (*big.Int, error) {
	return _PLCRVoting.Contract.GetTotalNumberOfTokensForWinningOption(&_PLCRVoting.CallOpts, id)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(terminationDate uint256) constant returns(expired bool)
func (_PLCRVoting *PLCRVotingCaller) IsExpired(opts *bind.CallOpts, terminationDate *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "isExpired", terminationDate)
	return *ret0, err
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(terminationDate uint256) constant returns(expired bool)
func (_PLCRVoting *PLCRVotingSession) IsExpired(terminationDate *big.Int) (bool, error) {
	return _PLCRVoting.Contract.IsExpired(&_PLCRVoting.CallOpts, terminationDate)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(terminationDate uint256) constant returns(expired bool)
func (_PLCRVoting *PLCRVotingCallerSession) IsExpired(terminationDate *big.Int) (bool, error) {
	return _PLCRVoting.Contract.IsExpired(&_PLCRVoting.CallOpts, terminationDate)
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(id uint256) constant returns(passed bool)
func (_PLCRVoting *PLCRVotingCaller) IsPassed(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "isPassed", id)
	return *ret0, err
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(id uint256) constant returns(passed bool)
func (_PLCRVoting *PLCRVotingSession) IsPassed(id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.IsPassed(&_PLCRVoting.CallOpts, id)
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(id uint256) constant returns(passed bool)
func (_PLCRVoting *PLCRVotingCallerSession) IsPassed(id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.IsPassed(&_PLCRVoting.CallOpts, id)
}

// PollEnded is a free data retrieval call binding the contract method 0xee684830.
//
// Solidity: function pollEnded(id uint256) constant returns(ended bool)
func (_PLCRVoting *PLCRVotingCaller) PollEnded(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "pollEnded", id)
	return *ret0, err
}

// PollEnded is a free data retrieval call binding the contract method 0xee684830.
//
// Solidity: function pollEnded(id uint256) constant returns(ended bool)
func (_PLCRVoting *PLCRVotingSession) PollEnded(id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.PollEnded(&_PLCRVoting.CallOpts, id)
}

// PollEnded is a free data retrieval call binding the contract method 0xee684830.
//
// Solidity: function pollEnded(id uint256) constant returns(ended bool)
func (_PLCRVoting *PLCRVotingCallerSession) PollEnded(id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.PollEnded(&_PLCRVoting.CallOpts, id)
}

// PollExists is a free data retrieval call binding the contract method 0x88d21ff3.
//
// Solidity: function pollExists(id uint256) constant returns(exists bool)
func (_PLCRVoting *PLCRVotingCaller) PollExists(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "pollExists", id)
	return *ret0, err
}

// PollExists is a free data retrieval call binding the contract method 0x88d21ff3.
//
// Solidity: function pollExists(id uint256) constant returns(exists bool)
func (_PLCRVoting *PLCRVotingSession) PollExists(id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.PollExists(&_PLCRVoting.CallOpts, id)
}

// PollExists is a free data retrieval call binding the contract method 0x88d21ff3.
//
// Solidity: function pollExists(id uint256) constant returns(exists bool)
func (_PLCRVoting *PLCRVotingCallerSession) PollExists(id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.PollExists(&_PLCRVoting.CallOpts, id)
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x441c77c0.
//
// Solidity: function revealPeriodActive(id uint256) constant returns(active bool)
func (_PLCRVoting *PLCRVotingCaller) RevealPeriodActive(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "revealPeriodActive", id)
	return *ret0, err
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x441c77c0.
//
// Solidity: function revealPeriodActive(id uint256) constant returns(active bool)
func (_PLCRVoting *PLCRVotingSession) RevealPeriodActive(id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.RevealPeriodActive(&_PLCRVoting.CallOpts, id)
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x441c77c0.
//
// Solidity: function revealPeriodActive(id uint256) constant returns(active bool)
func (_PLCRVoting *PLCRVotingCallerSession) RevealPeriodActive(id *big.Int) (bool, error) {
	return _PLCRVoting.Contract.RevealPeriodActive(&_PLCRVoting.CallOpts, id)
}

// ValidPosition is a free data retrieval call binding the contract method 0x819b0293.
//
// Solidity: function validPosition(prevId uint256, nextId uint256, voter address, numTokens uint256) constant returns(valid bool)
func (_PLCRVoting *PLCRVotingCaller) ValidPosition(opts *bind.CallOpts, prevId *big.Int, nextId *big.Int, voter common.Address, numTokens *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PLCRVoting.contract.Call(opts, out, "validPosition", prevId, nextId, voter, numTokens)
	return *ret0, err
}

// ValidPosition is a free data retrieval call binding the contract method 0x819b0293.
//
// Solidity: function validPosition(prevId uint256, nextId uint256, voter address, numTokens uint256) constant returns(valid bool)
func (_PLCRVoting *PLCRVotingSession) ValidPosition(prevId *big.Int, nextId *big.Int, voter common.Address, numTokens *big.Int) (bool, error) {
	return _PLCRVoting.Contract.ValidPosition(&_PLCRVoting.CallOpts, prevId, nextId, voter, numTokens)
}

// ValidPosition is a free data retrieval call binding the contract method 0x819b0293.
//
// Solidity: function validPosition(prevId uint256, nextId uint256, voter address, numTokens uint256) constant returns(valid bool)
func (_PLCRVoting *PLCRVotingCallerSession) ValidPosition(prevId *big.Int, nextId *big.Int, voter common.Address, numTokens *big.Int) (bool, error) {
	return _PLCRVoting.Contract.ValidPosition(&_PLCRVoting.CallOpts, prevId, nextId, voter, numTokens)
}

// CommitVote is a paid mutator transaction binding the contract method 0x6cbf9c5e.
//
// Solidity: function commitVote(id uint256, secretHash bytes32, numTokens uint256, prevId uint256) returns()
func (_PLCRVoting *PLCRVotingTransactor) CommitVote(opts *bind.TransactOpts, id *big.Int, secretHash [32]byte, numTokens *big.Int, prevId *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.contract.Transact(opts, "commitVote", id, secretHash, numTokens, prevId)
}

// CommitVote is a paid mutator transaction binding the contract method 0x6cbf9c5e.
//
// Solidity: function commitVote(id uint256, secretHash bytes32, numTokens uint256, prevId uint256) returns()
func (_PLCRVoting *PLCRVotingSession) CommitVote(id *big.Int, secretHash [32]byte, numTokens *big.Int, prevId *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.Contract.CommitVote(&_PLCRVoting.TransactOpts, id, secretHash, numTokens, prevId)
}

// CommitVote is a paid mutator transaction binding the contract method 0x6cbf9c5e.
//
// Solidity: function commitVote(id uint256, secretHash bytes32, numTokens uint256, prevId uint256) returns()
func (_PLCRVoting *PLCRVotingTransactorSession) CommitVote(id *big.Int, secretHash [32]byte, numTokens *big.Int, prevId *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.Contract.CommitVote(&_PLCRVoting.TransactOpts, id, secretHash, numTokens, prevId)
}

// RequestVotingRights is a paid mutator transaction binding the contract method 0xa25236fe.
//
// Solidity: function requestVotingRights(numTokens uint256) returns()
func (_PLCRVoting *PLCRVotingTransactor) RequestVotingRights(opts *bind.TransactOpts, numTokens *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.contract.Transact(opts, "requestVotingRights", numTokens)
}

// RequestVotingRights is a paid mutator transaction binding the contract method 0xa25236fe.
//
// Solidity: function requestVotingRights(numTokens uint256) returns()
func (_PLCRVoting *PLCRVotingSession) RequestVotingRights(numTokens *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.Contract.RequestVotingRights(&_PLCRVoting.TransactOpts, numTokens)
}

// RequestVotingRights is a paid mutator transaction binding the contract method 0xa25236fe.
//
// Solidity: function requestVotingRights(numTokens uint256) returns()
func (_PLCRVoting *PLCRVotingTransactorSession) RequestVotingRights(numTokens *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.Contract.RequestVotingRights(&_PLCRVoting.TransactOpts, numTokens)
}

// RescueTokens is a paid mutator transaction binding the contract method 0x97603560.
//
// Solidity: function rescueTokens(id uint256) returns()
func (_PLCRVoting *PLCRVotingTransactor) RescueTokens(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.contract.Transact(opts, "rescueTokens", id)
}

// RescueTokens is a paid mutator transaction binding the contract method 0x97603560.
//
// Solidity: function rescueTokens(id uint256) returns()
func (_PLCRVoting *PLCRVotingSession) RescueTokens(id *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.Contract.RescueTokens(&_PLCRVoting.TransactOpts, id)
}

// RescueTokens is a paid mutator transaction binding the contract method 0x97603560.
//
// Solidity: function rescueTokens(id uint256) returns()
func (_PLCRVoting *PLCRVotingTransactorSession) RescueTokens(id *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.Contract.RescueTokens(&_PLCRVoting.TransactOpts, id)
}

// RevealVote is a paid mutator transaction binding the contract method 0xb11d8bb8.
//
// Solidity: function revealVote(id uint256, voteOption uint256, salt uint256) returns()
func (_PLCRVoting *PLCRVotingTransactor) RevealVote(opts *bind.TransactOpts, id *big.Int, voteOption *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.contract.Transact(opts, "revealVote", id, voteOption, salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0xb11d8bb8.
//
// Solidity: function revealVote(id uint256, voteOption uint256, salt uint256) returns()
func (_PLCRVoting *PLCRVotingSession) RevealVote(id *big.Int, voteOption *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.Contract.RevealVote(&_PLCRVoting.TransactOpts, id, voteOption, salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0xb11d8bb8.
//
// Solidity: function revealVote(id uint256, voteOption uint256, salt uint256) returns()
func (_PLCRVoting *PLCRVotingTransactorSession) RevealVote(id *big.Int, voteOption *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.Contract.RevealVote(&_PLCRVoting.TransactOpts, id, voteOption, salt)
}

// StartPoll is a paid mutator transaction binding the contract method 0x32ed3d60.
//
// Solidity: function startPoll(voteQuorum uint256, commitDuration uint256, revealDuration uint256) returns(id uint256)
func (_PLCRVoting *PLCRVotingTransactor) StartPoll(opts *bind.TransactOpts, voteQuorum *big.Int, commitDuration *big.Int, revealDuration *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.contract.Transact(opts, "startPoll", voteQuorum, commitDuration, revealDuration)
}

// StartPoll is a paid mutator transaction binding the contract method 0x32ed3d60.
//
// Solidity: function startPoll(voteQuorum uint256, commitDuration uint256, revealDuration uint256) returns(id uint256)
func (_PLCRVoting *PLCRVotingSession) StartPoll(voteQuorum *big.Int, commitDuration *big.Int, revealDuration *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.Contract.StartPoll(&_PLCRVoting.TransactOpts, voteQuorum, commitDuration, revealDuration)
}

// StartPoll is a paid mutator transaction binding the contract method 0x32ed3d60.
//
// Solidity: function startPoll(voteQuorum uint256, commitDuration uint256, revealDuration uint256) returns(id uint256)
func (_PLCRVoting *PLCRVotingTransactorSession) StartPoll(voteQuorum *big.Int, commitDuration *big.Int, revealDuration *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.Contract.StartPoll(&_PLCRVoting.TransactOpts, voteQuorum, commitDuration, revealDuration)
}

// WithdrawVotingRights is a paid mutator transaction binding the contract method 0xe7b1d43c.
//
// Solidity: function withdrawVotingRights(numTokens uint256) returns()
func (_PLCRVoting *PLCRVotingTransactor) WithdrawVotingRights(opts *bind.TransactOpts, numTokens *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.contract.Transact(opts, "withdrawVotingRights", numTokens)
}

// WithdrawVotingRights is a paid mutator transaction binding the contract method 0xe7b1d43c.
//
// Solidity: function withdrawVotingRights(numTokens uint256) returns()
func (_PLCRVoting *PLCRVotingSession) WithdrawVotingRights(numTokens *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.Contract.WithdrawVotingRights(&_PLCRVoting.TransactOpts, numTokens)
}

// WithdrawVotingRights is a paid mutator transaction binding the contract method 0xe7b1d43c.
//
// Solidity: function withdrawVotingRights(numTokens uint256) returns()
func (_PLCRVoting *PLCRVotingTransactorSession) WithdrawVotingRights(numTokens *big.Int) (*types.Transaction, error) {
	return _PLCRVoting.Contract.WithdrawVotingRights(&_PLCRVoting.TransactOpts, numTokens)
}

// PLCRVotingPollCreatedEventIterator is returned from FilterPollCreatedEvent and is used to iterate over the raw logs and unpacked data for PollCreatedEvent events raised by the PLCRVoting contract.
type PLCRVotingPollCreatedEventIterator struct {
	Event *PLCRVotingPollCreatedEvent // Event containing the contract specifics and raw log

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
func (it *PLCRVotingPollCreatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PLCRVotingPollCreatedEvent)
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
		it.Event = new(PLCRVotingPollCreatedEvent)
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
func (it *PLCRVotingPollCreatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PLCRVotingPollCreatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PLCRVotingPollCreatedEvent represents a PollCreatedEvent event raised by the PLCRVoting contract.
type PLCRVotingPollCreatedEvent struct {
	VoteQuorum   *big.Int
	CommitExpiry *big.Int
	RevealExpiry *big.Int
	Id           *big.Int
	Creator      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPollCreatedEvent is a free log retrieval operation binding the contract event 0xf7e7200f3afa7bbc8aca2cfa0ab33e741504a20fed2abc59869413cba2a33400.
//
// Solidity: e PollCreatedEvent(voteQuorum uint256, commitExpiry uint256, revealExpiry uint256, id indexed uint256, creator indexed address)
func (_PLCRVoting *PLCRVotingFilterer) FilterPollCreatedEvent(opts *bind.FilterOpts, id []*big.Int, creator []common.Address) (*PLCRVotingPollCreatedEventIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _PLCRVoting.contract.FilterLogs(opts, "PollCreatedEvent", idRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingPollCreatedEventIterator{contract: _PLCRVoting.contract, event: "PollCreatedEvent", logs: logs, sub: sub}, nil
}

// WatchPollCreatedEvent is a free log subscription operation binding the contract event 0xf7e7200f3afa7bbc8aca2cfa0ab33e741504a20fed2abc59869413cba2a33400.
//
// Solidity: e PollCreatedEvent(voteQuorum uint256, commitExpiry uint256, revealExpiry uint256, id indexed uint256, creator indexed address)
func (_PLCRVoting *PLCRVotingFilterer) WatchPollCreatedEvent(opts *bind.WatchOpts, sink chan<- *PLCRVotingPollCreatedEvent, id []*big.Int, creator []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _PLCRVoting.contract.WatchLogs(opts, "PollCreatedEvent", idRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PLCRVotingPollCreatedEvent)
				if err := _PLCRVoting.contract.UnpackLog(event, "PollCreatedEvent", log); err != nil {
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

// PLCRVotingTokensRescuedEventIterator is returned from FilterTokensRescuedEvent and is used to iterate over the raw logs and unpacked data for TokensRescuedEvent events raised by the PLCRVoting contract.
type PLCRVotingTokensRescuedEventIterator struct {
	Event *PLCRVotingTokensRescuedEvent // Event containing the contract specifics and raw log

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
func (it *PLCRVotingTokensRescuedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PLCRVotingTokensRescuedEvent)
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
		it.Event = new(PLCRVotingTokensRescuedEvent)
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
func (it *PLCRVotingTokensRescuedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PLCRVotingTokensRescuedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PLCRVotingTokensRescuedEvent represents a TokensRescuedEvent event raised by the PLCRVoting contract.
type PLCRVotingTokensRescuedEvent struct {
	Id    *big.Int
	Voter common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokensRescuedEvent is a free log retrieval operation binding the contract event 0x1e897e4fe81e0db0459f0219545995fde305b17af5ca5cba7d6b6e8738df5dd0.
//
// Solidity: e TokensRescuedEvent(id indexed uint256, voter indexed address)
func (_PLCRVoting *PLCRVotingFilterer) FilterTokensRescuedEvent(opts *bind.FilterOpts, id []*big.Int, voter []common.Address) (*PLCRVotingTokensRescuedEventIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVoting.contract.FilterLogs(opts, "TokensRescuedEvent", idRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingTokensRescuedEventIterator{contract: _PLCRVoting.contract, event: "TokensRescuedEvent", logs: logs, sub: sub}, nil
}

// WatchTokensRescuedEvent is a free log subscription operation binding the contract event 0x1e897e4fe81e0db0459f0219545995fde305b17af5ca5cba7d6b6e8738df5dd0.
//
// Solidity: e TokensRescuedEvent(id indexed uint256, voter indexed address)
func (_PLCRVoting *PLCRVotingFilterer) WatchTokensRescuedEvent(opts *bind.WatchOpts, sink chan<- *PLCRVotingTokensRescuedEvent, id []*big.Int, voter []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVoting.contract.WatchLogs(opts, "TokensRescuedEvent", idRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PLCRVotingTokensRescuedEvent)
				if err := _PLCRVoting.contract.UnpackLog(event, "TokensRescuedEvent", log); err != nil {
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

// PLCRVotingVoteCommittedEventIterator is returned from FilterVoteCommittedEvent and is used to iterate over the raw logs and unpacked data for VoteCommittedEvent events raised by the PLCRVoting contract.
type PLCRVotingVoteCommittedEventIterator struct {
	Event *PLCRVotingVoteCommittedEvent // Event containing the contract specifics and raw log

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
func (it *PLCRVotingVoteCommittedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PLCRVotingVoteCommittedEvent)
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
		it.Event = new(PLCRVotingVoteCommittedEvent)
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
func (it *PLCRVotingVoteCommittedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PLCRVotingVoteCommittedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PLCRVotingVoteCommittedEvent represents a VoteCommittedEvent event raised by the PLCRVoting contract.
type PLCRVotingVoteCommittedEvent struct {
	Id        *big.Int
	NumTokens *big.Int
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoteCommittedEvent is a free log retrieval operation binding the contract event 0x673f4a020a46f562a6e9fee9c8d4a179a11924190b777bec160133dd00c9f1de.
//
// Solidity: e VoteCommittedEvent(id indexed uint256, numTokens uint256, voter indexed address)
func (_PLCRVoting *PLCRVotingFilterer) FilterVoteCommittedEvent(opts *bind.FilterOpts, id []*big.Int, voter []common.Address) (*PLCRVotingVoteCommittedEventIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVoting.contract.FilterLogs(opts, "VoteCommittedEvent", idRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingVoteCommittedEventIterator{contract: _PLCRVoting.contract, event: "VoteCommittedEvent", logs: logs, sub: sub}, nil
}

// WatchVoteCommittedEvent is a free log subscription operation binding the contract event 0x673f4a020a46f562a6e9fee9c8d4a179a11924190b777bec160133dd00c9f1de.
//
// Solidity: e VoteCommittedEvent(id indexed uint256, numTokens uint256, voter indexed address)
func (_PLCRVoting *PLCRVotingFilterer) WatchVoteCommittedEvent(opts *bind.WatchOpts, sink chan<- *PLCRVotingVoteCommittedEvent, id []*big.Int, voter []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVoting.contract.WatchLogs(opts, "VoteCommittedEvent", idRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PLCRVotingVoteCommittedEvent)
				if err := _PLCRVoting.contract.UnpackLog(event, "VoteCommittedEvent", log); err != nil {
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

// PLCRVotingVoteRevealedEventIterator is returned from FilterVoteRevealedEvent and is used to iterate over the raw logs and unpacked data for VoteRevealedEvent events raised by the PLCRVoting contract.
type PLCRVotingVoteRevealedEventIterator struct {
	Event *PLCRVotingVoteRevealedEvent // Event containing the contract specifics and raw log

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
func (it *PLCRVotingVoteRevealedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PLCRVotingVoteRevealedEvent)
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
		it.Event = new(PLCRVotingVoteRevealedEvent)
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
func (it *PLCRVotingVoteRevealedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PLCRVotingVoteRevealedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PLCRVotingVoteRevealedEvent represents a VoteRevealedEvent event raised by the PLCRVoting contract.
type PLCRVotingVoteRevealedEvent struct {
	Id           *big.Int
	NumTokens    *big.Int
	VotesFor     *big.Int
	VotesAgainst *big.Int
	Choice       *big.Int
	Voter        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVoteRevealedEvent is a free log retrieval operation binding the contract event 0x448a2036d0f6dc799c1d3fdf26c4fe80832817bdbe81133808ffa6777c5826e0.
//
// Solidity: e VoteRevealedEvent(id indexed uint256, numTokens uint256, votesFor uint256, votesAgainst uint256, choice indexed uint256, voter indexed address)
func (_PLCRVoting *PLCRVotingFilterer) FilterVoteRevealedEvent(opts *bind.FilterOpts, id []*big.Int, choice []*big.Int, voter []common.Address) (*PLCRVotingVoteRevealedEventIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var choiceRule []interface{}
	for _, choiceItem := range choice {
		choiceRule = append(choiceRule, choiceItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVoting.contract.FilterLogs(opts, "VoteRevealedEvent", idRule, choiceRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingVoteRevealedEventIterator{contract: _PLCRVoting.contract, event: "VoteRevealedEvent", logs: logs, sub: sub}, nil
}

// WatchVoteRevealedEvent is a free log subscription operation binding the contract event 0x448a2036d0f6dc799c1d3fdf26c4fe80832817bdbe81133808ffa6777c5826e0.
//
// Solidity: e VoteRevealedEvent(id indexed uint256, numTokens uint256, votesFor uint256, votesAgainst uint256, choice indexed uint256, voter indexed address)
func (_PLCRVoting *PLCRVotingFilterer) WatchVoteRevealedEvent(opts *bind.WatchOpts, sink chan<- *PLCRVotingVoteRevealedEvent, id []*big.Int, choice []*big.Int, voter []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var choiceRule []interface{}
	for _, choiceItem := range choice {
		choiceRule = append(choiceRule, choiceItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVoting.contract.WatchLogs(opts, "VoteRevealedEvent", idRule, choiceRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PLCRVotingVoteRevealedEvent)
				if err := _PLCRVoting.contract.UnpackLog(event, "VoteRevealedEvent", log); err != nil {
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

// PLCRVotingVotingRightsGrantedEventIterator is returned from FilterVotingRightsGrantedEvent and is used to iterate over the raw logs and unpacked data for VotingRightsGrantedEvent events raised by the PLCRVoting contract.
type PLCRVotingVotingRightsGrantedEventIterator struct {
	Event *PLCRVotingVotingRightsGrantedEvent // Event containing the contract specifics and raw log

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
func (it *PLCRVotingVotingRightsGrantedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PLCRVotingVotingRightsGrantedEvent)
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
		it.Event = new(PLCRVotingVotingRightsGrantedEvent)
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
func (it *PLCRVotingVotingRightsGrantedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PLCRVotingVotingRightsGrantedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PLCRVotingVotingRightsGrantedEvent represents a VotingRightsGrantedEvent event raised by the PLCRVoting contract.
type PLCRVotingVotingRightsGrantedEvent struct {
	NumTokens *big.Int
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVotingRightsGrantedEvent is a free log retrieval operation binding the contract event 0x0fbea84792de88539ad3d7f184543a65916440f4c6f9e2ef136c97f13109c45d.
//
// Solidity: e VotingRightsGrantedEvent(numTokens uint256, voter indexed address)
func (_PLCRVoting *PLCRVotingFilterer) FilterVotingRightsGrantedEvent(opts *bind.FilterOpts, voter []common.Address) (*PLCRVotingVotingRightsGrantedEventIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVoting.contract.FilterLogs(opts, "VotingRightsGrantedEvent", voterRule)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingVotingRightsGrantedEventIterator{contract: _PLCRVoting.contract, event: "VotingRightsGrantedEvent", logs: logs, sub: sub}, nil
}

// WatchVotingRightsGrantedEvent is a free log subscription operation binding the contract event 0x0fbea84792de88539ad3d7f184543a65916440f4c6f9e2ef136c97f13109c45d.
//
// Solidity: e VotingRightsGrantedEvent(numTokens uint256, voter indexed address)
func (_PLCRVoting *PLCRVotingFilterer) WatchVotingRightsGrantedEvent(opts *bind.WatchOpts, sink chan<- *PLCRVotingVotingRightsGrantedEvent, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVoting.contract.WatchLogs(opts, "VotingRightsGrantedEvent", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PLCRVotingVotingRightsGrantedEvent)
				if err := _PLCRVoting.contract.UnpackLog(event, "VotingRightsGrantedEvent", log); err != nil {
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

// PLCRVotingVotingRightsWithdrawnEventIterator is returned from FilterVotingRightsWithdrawnEvent and is used to iterate over the raw logs and unpacked data for VotingRightsWithdrawnEvent events raised by the PLCRVoting contract.
type PLCRVotingVotingRightsWithdrawnEventIterator struct {
	Event *PLCRVotingVotingRightsWithdrawnEvent // Event containing the contract specifics and raw log

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
func (it *PLCRVotingVotingRightsWithdrawnEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PLCRVotingVotingRightsWithdrawnEvent)
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
		it.Event = new(PLCRVotingVotingRightsWithdrawnEvent)
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
func (it *PLCRVotingVotingRightsWithdrawnEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PLCRVotingVotingRightsWithdrawnEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PLCRVotingVotingRightsWithdrawnEvent represents a VotingRightsWithdrawnEvent event raised by the PLCRVoting contract.
type PLCRVotingVotingRightsWithdrawnEvent struct {
	NumTokens *big.Int
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVotingRightsWithdrawnEvent is a free log retrieval operation binding the contract event 0x029a58e9abaa250a4ec75d14b364d764e890adec2d902a297fa24a6c5dfd097b.
//
// Solidity: e VotingRightsWithdrawnEvent(numTokens uint256, voter indexed address)
func (_PLCRVoting *PLCRVotingFilterer) FilterVotingRightsWithdrawnEvent(opts *bind.FilterOpts, voter []common.Address) (*PLCRVotingVotingRightsWithdrawnEventIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVoting.contract.FilterLogs(opts, "VotingRightsWithdrawnEvent", voterRule)
	if err != nil {
		return nil, err
	}
	return &PLCRVotingVotingRightsWithdrawnEventIterator{contract: _PLCRVoting.contract, event: "VotingRightsWithdrawnEvent", logs: logs, sub: sub}, nil
}

// WatchVotingRightsWithdrawnEvent is a free log subscription operation binding the contract event 0x029a58e9abaa250a4ec75d14b364d764e890adec2d902a297fa24a6c5dfd097b.
//
// Solidity: e VotingRightsWithdrawnEvent(numTokens uint256, voter indexed address)
func (_PLCRVoting *PLCRVotingFilterer) WatchVotingRightsWithdrawnEvent(opts *bind.WatchOpts, sink chan<- *PLCRVotingVotingRightsWithdrawnEvent, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _PLCRVoting.contract.WatchLogs(opts, "VotingRightsWithdrawnEvent", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PLCRVotingVotingRightsWithdrawnEvent)
				if err := _PLCRVoting.contract.UnpackLog(event, "VotingRightsWithdrawnEvent", log); err != nil {
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

// ParameterizerABI is the input ABI used to generate the binding from.
const ParameterizerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"propHash\",\"type\":\"bytes32\"}],\"name\":\"processProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"propHash\",\"type\":\"bytes32\"}],\"name\":\"propExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"challengeWinnerReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"propHash\",\"type\":\"bytes32\"}],\"name\":\"challengeCanBeResolved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"propHash\",\"type\":\"bytes32\"}],\"name\":\"challengeReparameterization\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"rewardsClaimed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"voterReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"proposeReparameterization\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"propHash\",\"type\":\"bytes32\"}],\"name\":\"canBeSet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"},{\"name\":\"votingAddr\",\"type\":\"address\"},{\"name\":\"minDeposit\",\"type\":\"uint256\"},{\"name\":\"applyStageLen\",\"type\":\"uint256\"},{\"name\":\"commitStageLen\",\"type\":\"uint256\"},{\"name\":\"revealStageLen\",\"type\":\"uint256\"},{\"name\":\"dispensationPct\",\"type\":\"uint256\"},{\"name\":\"voteQuorum\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"appExpiry\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"ReparameterizationProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commitExpiry\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealExpiry\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"NewChallengeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ProposalAcceptedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalExpiredEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"name\":\"ChallengeSucceededEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"name\":\"ChallengeFailedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"RewardsClaimedEvent\",\"type\":\"event\"}]"

// ParameterizerBin is the compiled bytecode used for deploying new contracts.
const ParameterizerBin = `0x608060405262093a806005553480156200001857600080fd5b506040516101008062002d318339810160408181528251602080850151838601516060870151608088015160a089015160c08a015160e0909a015160038054600160a060020a03808b16600160a060020a0319928316179092556004805492891692909116919091179055888a01909852600a89527f6d696e4465706f73697400000000000000000000000000000000000000000000958901959095529497929691959094939291620000d5908764010000000062000250810204565b60408051808201909152600d81527f6170706c7953746167654c656e0000000000000000000000000000000000000060208201526200011e908664010000000062000250810204565b60408051808201909152600e81527f636f6d6d697453746167654c656e000000000000000000000000000000000000602082015262000167908564010000000062000250810204565b60408051808201909152600e81527f72657665616c53746167654c656e0000000000000000000000000000000000006020820152620001b0908464010000000062000250810204565b60408051808201909152600f81527f64697370656e736174696f6e50637400000000000000000000000000000000006020820152620001f9908364010000000062000250810204565b60408051808201909152600a81527f766f746551756f72756d00000000000000000000000000000000000000000000602082015262000242908264010000000062000250810204565b505050505050505062000330565b80600080846040516020018082805190602001908083835b60208310620002895780518252601f19909201916020918201910162000268565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b60208310620002ee5780518252601f199092019160209182019101620002cd565b51815160209384036101000a60001901801990921691161790526040805192909401829003909120865285019590955292909201600020939093555050505050565b6129f180620003406000396000f3006080604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166330490e9181146100b357806335300990146100cd57806350411552146100f9578063693ec85e1461012357806377609a411461017c5780638240ae4b1461019457806386bb8f37146101ac5780639a835e91146101c7578063a7aad3db146101eb578063bade1c5414610212578063c51131fb14610236575b600080fd5b3480156100bf57600080fd5b506100cb60043561024e565b005b3480156100d957600080fd5b506100e5600435610814565b604080519115158252519081900360200190f35b34801561010557600080fd5b5061011160043561082d565b60408051918252519081900360200190f35b34801561012f57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261011194369492936024939284019190819084018382808284375094975061092c9650505050505050565b34801561018857600080fd5b506100e5600435610a08565b3480156101a057600080fd5b50610111600435610c32565b3480156101b857600080fd5b506100cb600435602435611336565b3480156101d357600080fd5b506100e5600435600160a060020a036024351661172b565b3480156101f757600080fd5b50610111600160a060020a036004351660243560443561175a565b34801561021e57600080fd5b50610111602460048035828101929101359035611836565b34801561024257600080fd5b506100e5600435612073565b60008181526002602081905260409091206004810154918101549091600160a060020a03169061027d84612073565b1561054857600383018054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815261032093909290918301828280156103115780601f106102e657610100808354040283529160200191610311565b820191906000526020600020905b8154815290600101906020018083116102f457829003601f168201915b50505050508460060154612196565b600683015460408051602081018390528181526003860180546002610100600183161502600019019091160492820183905287937f0fca8490116d823913d2468dc21632e9226d20e54379af9cf34712a3ff2bc30c939192909181906060820190859080156103d05780601f106103a5576101008083540402835291602001916103d0565b820191906000526020600020905b8154815290600101906020018083116103b357829003601f168201915b5050935050505060405180910390a260008481526002602081905260408220828155600181018390559081018290559061040d6003830182612841565b506004818101805473ffffffffffffffffffffffffffffffffffffffff1916905560006005830181905560069092018290556003546040805160e060020a63a9059cbb028152600160a060020a0387811694820194909452602481018690529051929091169263a9059cbb926044808401936020939083900390910190829087803b15801561049b57600080fd5b505af11580156104af573d6000803e3d6000fd5b505050506040513d60208110156104c557600080fd5b50511515610543576040805160e560020a62461bcd02815260206004820152603e60248201527f4572726f723a506172616d65746572697a65722e70726f6365737350726f706f60448201527f73616c202d20436f756c64206e6f74207472616e736665722066756e64730000606482015290519081900360840190fd5b610686565b61055184610a08565b1561055f5761054384612272565b82600501544211156105c45760405184907f456cbba32d3f5372c3ef95ba54f18060130b9da12721ebf89f1146a8f406132490600090a260008481526002602081905260408220828155600181018390559081018290559061040d6003830182612841565b6040805160e560020a62461bcd02815260206004820152606c60248201527f5468657265206973206e6f206368616c6c656e676520616761696e737420746860448201527f652070726f706f73616c2c20616e64206e65697468657220746865206170704560648201527f78706972792064617465206e6f72207468652070726f6365737342792064617460848201527f652068617320706173736564000000000000000000000000000000000000000060a482015290519081900360c40190fd5b60646106c66040805190810160405280600f81526020017f64697370656e736174696f6e506374000000000000000000000000000000000081525061092c565b11156106ce57fe5b6107b06005546107a46107156040805190810160405280600e81526020017f72657665616c53746167654c656e00000000000000000000000000000000000081525061092c565b6107a46107566040805190810160405280600e81526020017f636f6d6d697453746167654c656e00000000000000000000000000000000000081525061092c565b6107a46107976040805190810160405280600d81526020017f6170706c7953746167654c656e0000000000000000000000000000000000000081525061092c565b429063ffffffff6127e416565b9063ffffffff6127e416565b506000848152600260208190526040822082815560018101839055908101829055906107df6003830182612841565b5060048101805473ffffffffffffffffffffffffffffffffffffffff1916905560006005820181905560069091015550505050565b600081815260026020526040812060050154115b919050565b60048054604080517f053e71a600000000000000000000000000000000000000000000000000000000815292830184905251600092600160a060020a039092169163053e71a691602480830192602092919082900301818787803b15801561089457600080fd5b505af11580156108a8573d6000803e3d6000fd5b505050506040513d60208110156108be57600080fd5b505115156108f05760008281526001602052604090206002908101546108e99163ffffffff6127f116565b9050610828565b600082815260016020526040902080546002918201546109269261091a919063ffffffff6127f116565b9063ffffffff61281a16565b92915050565b6000806000836040516020018082805190602001908083835b602083106109645780518252601f199092019160209182019101610945565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106109c75780518252601f1990920191602091820191016109a8565b51815160209384036101000a600019018019909216911617905260408051929094018290039091208652850195909552929092016000205495945050505050565b6000610a12612888565b610a1a6128cf565b600084815260026020818152604092839020835160e0810185528154815260018083015482850152828501548287015260038301805487516000199382161561010002939093011695909504601f810185900485028201850190965285815290949193606086019391929091830182828015610ad75780601f10610aac57610100808354040283529160200191610ad7565b820191906000526020600020905b815481529060010190602001808311610aba57829003601f168201915b50505091835250506004820154600160a060020a039081166020808401919091526005840154604080850191909152600690940154606093840152848101805160009081526001808452868220875160a0810189528154815291810154958616948201949094527401000000000000000000000000000000000000000090940460ff1615159584019590955260028201549383019390935260030154608082015290519294509250108015610b8e57506040810151155b8015610c2a575060048054602080850151604080517fee6848300000000000000000000000000000000000000000000000000000000081529485019190915251600160a060020a039092169263ee68483092602480830193928290030181600087803b158015610bfd57600080fd5b505af1158015610c11573d6000803e3d6000fd5b505050506040513d6020811015610c2757600080fd5b50515b949350505050565b6000610c3c612888565b6000838152600260208181526040808420815160e0810183528154815260018083015482860152828601548285015260038301805485516000199382161561010002939093011696909604601f810186900486028201860190945283815286958695869593949360608601939291830182828015610cfb5780601f10610cd057610100808354040283529160200191610cfb565b820191906000526020600020905b815481529060010190602001808311610cde57829003601f168201915b50505091835250506004820154600160a060020a0316602082015260058201546040808301919091526006909201546060909101528101519095509350610d4187610814565b8015610d4f57506020850151155b1515610df1576040805160e560020a62461bcd02815260206004820152604960248201527f4572726f723a506172616d65746572697a65722e6368616c6c656e676552657060448201527f6172616d65746572697a6174696f6e202d2050726f706f73616c20646f65732060648201527f6e6f742065786973740000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b60045460408051808201909152600a81527f766f746551756f72756d000000000000000000000000000000000000000000006020820152600160a060020a03909116906332ed3d6090610e439061092c565b610e816040805190810160405280600e81526020017f636f6d6d697453746167654c656e00000000000000000000000000000000000081525061092c565b610ebf6040805190810160405280600e81526020017f72657665616c53746167654c656e00000000000000000000000000000000000081525061092c565b6040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808481526020018381526020018281526020019350505050602060405180830381600087803b158015610f1f57600080fd5b505af1158015610f33573d6000803e3d6000fd5b505050506040513d6020811015610f4957600080fd5b50516040805160e08101909152600f60a082019081527f64697370656e736174696f6e506374000000000000000000000000000000000060c0830152919450908190610fc490606490610fb8908990610fac908490610fa79061092c565b61281a565b9063ffffffff6127f116565b9063ffffffff61282c16565b81523360208083018290526000604080850182905260608086018b905260809586018390528983526001808552828420885181558886015181830180548b87015173ffffffffffffffffffffffffffffffffffffffff19909116600160a060020a039384161774ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000911515919091021790559289015160028083019190915598909701516003978801558e8452968452818320909601899055935484517f23b872dd0000000000000000000000000000000000000000000000000000000081526004810194909452306024850152604484018a9052935193909416936323b872dd936064808501949192918390030190829087803b1580156110f157600080fd5b505af1158015611105573d6000803e3d6000fd5b505050506040513d602081101561111b57600080fd5b505115156111bf576040805160e560020a62461bcd02815260206004820152604a60248201527f4572726f723a506172616d65746572697a65722e6368616c6c656e676552657060448201527f6172616d65746572697a6174696f6e202d20436f756c64206e6f74207472616e60648201527f736665722066756e647300000000000000000000000000000000000000000000608482015290519081900360a40190fd5b60048054604080517fdc1b946f00000000000000000000000000000000000000000000000000000000815292830186905251600160a060020a039091169163dc1b946f9160248083019260209291908290030181600087803b15801561122457600080fd5b505af1158015611238573d6000803e3d6000fd5b505050506040513d602081101561124e57600080fd5b505160048054604080517f1be8ea5000000000000000000000000000000000000000000000000000000000815292830187905251929450600160a060020a031691631be8ea50916024808201926020929091908290030181600087803b1580156112b757600080fd5b505af11580156112cb573d6000803e3d6000fd5b505050506040513d60208110156112e157600080fd5b505160408051858152602081018590528082018390529051919250339189917f4e6815e2c453c1de363fd2b483078339512560f8fd777420d06dd69291e060c2919081900360600190a3509095945050505050565b6000828152600160209081526040808320338452600401909152812054819060ff16156113f9576040805160e560020a62461bcd02815260206004820152604260248201527f4572726f723a506172616d65746572697a65722e636c61696d5265776172642060448201527f2d20546f6b656e73206861766520616c7265616479206265656e20636c61696d60648201527f6564000000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b600084815260016020819052604090912081015474010000000000000000000000000000000000000000900460ff161515146114cb576040805160e560020a62461bcd02815260206004820152604f60248201527f4572726f723a506172616d65746572697a65722e636c61696d5265776172642060448201527f2d204368616c6c656e676520726573756c74732068617665206e6f742062656560648201527f6e2070726f636573736564207965740000000000000000000000000000000000608482015290519081900360a40190fd5b60048054604080517fb43bd0690000000000000000000000000000000000000000000000000000000081523393810193909352602483018790526044830186905251600160a060020a039091169163b43bd0699160648083019260209291908290030181600087803b15801561154057600080fd5b505af1158015611554573d6000803e3d6000fd5b505050506040513d602081101561156a57600080fd5b5051915061157933858561175a565b60008581526001602052604090206003015490915061159e908363ffffffff61281a16565b60008581526001602052604090206003810191909155546115c5908263ffffffff61281a16565b6000858152600160208181526040808420948555338085526004909501825292839020805460ff19169092179091558151848152915187927f17ccb5aed827c5078139dc28b7fb924b006f56b3dcf163c4acf21fa3beb1314492908290030190a36003546040805160e060020a63a9059cbb028152336004820152602481018490529051600160a060020a039092169163a9059cbb916044808201926020929091908290030181600087803b15801561167d57600080fd5b505af1158015611691573d6000803e3d6000fd5b505050506040513d60208110156116a757600080fd5b50511515611725576040805160e560020a62461bcd02815260206004820152603a60248201527f4572726f723a506172616d65746572697a65722e636c61696d5265776172642060448201527f2d20436f756c64206e6f74207472616e736665722066756e6473000000000000606482015290519081900360840190fd5b50505050565b6000828152600160209081526040808320600160a060020a038516845260040190915290205460ff1692915050565b6000828152600160209081526040808320600381015490546004805484517fb43bd069000000000000000000000000000000000000000000000000000000008152600160a060020a038b811693820193909352602481018a9052604481018990529451939592948794929091169263b43bd0699260648084019382900301818787803b1580156117e957600080fd5b505af11580156117fd573d6000803e3d6000fd5b505050506040513d602081101561181357600080fd5b5051905061182b83610fb8838563ffffffff6127f116565b979650505050505050565b60008060006118796040805190810160405280600a81526020017f6d696e4465706f7369740000000000000000000000000000000000000000000081525061092c565b915085858560405160200180848480828437820191505082815260200193505050506040516020818303038152906040526040518082805190602001908083835b602083106118d95780518252601f1990920191602091820191016118ba565b51815160209384036101000a6000190180199092169116179052604080519290940182900382207f64697370656e736174696f6e5063740000000000000000000000000000000000838301528451600f818503018152602f9093019485905282519097509195509293508392850191508083835b6020831061196c5780518252601f19909201916020918201910161194d565b51815160209384036101000a600019018019909216911617905260405191909301819003812094508a93508992019050808383808284378201915050925050506040516020818303038152906040526040518082805190602001908083835b602083106119ea5780518252601f1990920191602091820191016119cb565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020600019161480611b42575060405160200180807f7044697370656e736174696f6e5063740000000000000000000000000000000081525060100190506040516020818303038152906040526040518082805190602001908083835b60208310611a915780518252601f199092019160209182019101611a72565b51815160209384036101000a600019018019909216911617905260405191909301819003812094508a93508992019050808383808284378201915050925050506040516020818303038152906040526040518082805190602001908083835b60208310611b0f5780518252601f199092019160209182019101611af0565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916145b15611bda576064841115611bda576040805160e560020a62461bcd02815260206004820152605960248201526000805160206129a683398151915260448201527f616d65746572697a6174696f6e202d206076616c756560206d7573742062652060648201527f6c657373207468616e206f7220657175616c20746f2031303000000000000000608482015290519081900360a40190fd5b611be381610814565b15611c72576040805160e560020a62461bcd02815260206004820152604f60248201526000805160206129a683398151915260448201527f616d65746572697a6174696f6e202d204475706c69636174652070726f706f7360648201527f616c73206e6f7420616c6c6f7765640000000000000000000000000000000000608482015290519081900360a40190fd5b83611cac87878080601f0160208091040260200160405190810160405280939291908181526020018383808284375061092c945050505050565b1415611d3c576040805160e560020a62461bcd02815260206004820152604d60248201526000805160206129a683398151915260448201527f616d65746572697a6174696f6e202d20417267756d656e74732063616e6e6f7460648201527f206265206964656e746963616c00000000000000000000000000000000000000608482015290519081900360a40190fd5b604080516101208101909152600d60e082019081527f6170706c7953746167654c656e000000000000000000000000000000000000006101008301528190611d87906107979061092c565b81526020016000815260200183815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200133600160a060020a03168152602001611e276005546107a46107156040805190810160405280600e81526020017f72657665616c53746167654c656e00000000000000000000000000000000000081525061092c565b8152602090810186905260008381526002808352604091829020845181558484015160018201559184015190820155606083015180519192611e719260038501929091019061290a565b5060808201516004828101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0393841617905560a0840151600584015560c090930151600690920191909155600354604080517f23b872dd000000000000000000000000000000000000000000000000000000008152339481019490945230602485015260448401869052519116916323b872dd9160648083019260209291908290030181600087803b158015611f2657600080fd5b505af1158015611f3a573d6000803e3d6000fd5b505050506040513d6020811015611f5057600080fd5b50511515611fe2576040805160e560020a62461bcd02815260206004820152604860248201526000805160206129a683398151915260448201527f616d65746572697a6174696f6e202d20436f756c64206e6f74207472616e736660648201527f65722066756e6473000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b600081815260026020908152604091829020548251918201879052918101839052606081018490526080810182905260a0808252810187905233917ff6aafb6c8410e85ab57ef6d7df18fcb782806ee3b89634f3b52046846236f97a9189918991899187918991908060c081018888808284376040519201829003995090975050505050505050a295945050505050565b600061207d612888565b600083815260026020818152604092839020835160e0810185528154815260018083015482850152828501548287015260038301805487516000199382161561010002939093011695909504601f81018590048502820185019096528581529094919360608601939192909183018282801561213a5780601f1061210f5761010080835404028352916020019161213a565b820191906000526020600020905b81548152906001019060200180831161211d57829003601f168201915b50505091835250506004820154600160a060020a031660208201526005820154604082015260069091015460609091015280519091504211801561218157508060a0015142105b801561218f57506020810151155b9392505050565b80600080846040516020018082805190602001908083835b602083106121cd5780518252601f1990920191602091820191016121ae565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106122305780518252601f199092019160209182019101612211565b51815160209384036101000a60001901801990921691161790526040805192909401829003909120865285019590955292909201600020939093555050505050565b61227a612888565b6000828152600260208181526040808420815160e0810183528154815260018083015482860152828601548285015260038301805485516000199382161561010002939093011696909604601f810186900486028201860190945283815286959194929360608601939192918301828280156123375780601f1061230c57610100808354040283529160200191612337565b820191906000526020600020905b81548152906001019060200180831161231a57829003601f168201915b505050505081526020016004820160009054906101000a9004600160a060020a0316600160a060020a0316600160a060020a03168152602001600582015481526020016006820154815250509250600160008460200151815260200190815260200160002091506123ab836020015161082d565b60048054602080870151604080517f053e71a60000000000000000000000000000000000000000000000000000000081529485019190915251939450600160a060020a039091169263053e71a69260248082019392918290030181600087803b15801561241757600080fd5b505af115801561242b573d6000803e3d6000fd5b505050506040513d602081101561244157600080fd5b5051600383015560018201805474ff000000000000000000000000000000000000000019167401000000000000000000000000000000000000000017905560048054602085810151604080517f494031830000000000000000000000000000000000000000000000000000000081529485019190915251600160a060020a0390921692634940318392602480830193928290030181600087803b1580156124e757600080fd5b505af11580156124fb573d6000803e3d6000fd5b505050506040513d602081101561251157600080fd5b50511561268957428360a0015111156125365761253683606001518460c00151612196565b6020808401518354600385015460408051928352938201528251919287927f824960b36dcff66c0b43953e552b1b0d622dcb3f3cf71cf1cbe88039a781a8539281900390910190a360035460808401516040805160e060020a63a9059cbb028152600160a060020a039283166004820152602481018590529051919092169163a9059cbb9160448083019260209291908290030181600087803b1580156125dc57600080fd5b505af11580156125f0573d6000803e3d6000fd5b505050506040513d602081101561260657600080fd5b50511515612684576040805160e560020a62461bcd02815260206004820152603f60248201527f4572726f723a506172616d65746572697a65722e7265736f6c76654368616c6c60448201527f656e6765202d20436f756c64206e6f74207472616e736665722066756e647300606482015290519081900360840190fd5b611725565b6020808401518354600385015460408051928352938201528251919287927f3dddcf5cd7cb20cfa68f0643b9422cdbc130a6602933781bf9855e10fd205e429281900390910190a360035460208085015160009081526001808352604080832090910154815160e060020a63a9059cbb028152600160a060020a03918216600482015260248101879052915194169363a9059cbb9360448084019491938390030190829087803b15801561273c57600080fd5b505af1158015612750573d6000803e3d6000fd5b505050506040513d602081101561276657600080fd5b50511515611725576040805160e560020a62461bcd02815260206004820152603f60248201527f4572726f723a506172616d65746572697a65722e7265736f6c76654368616c6c60448201527f656e6765202d20436f756c64206e6f74207472616e736665722066756e647300606482015290519081900360840190fd5b8181018281101561092657fe5b600082151561280257506000610926565b5081810281838281151561281257fe5b041461092657fe5b60008282111561282657fe5b50900390565b6000818381151561283957fe5b049392505050565b50805460018160011615610100020316600290046000825580601f106128675750612885565b601f0160209004906000526020600020908101906128859190612988565b50565b60e060405190810160405280600081526020016000815260200160008152602001606081526020016000600160a060020a0316815260200160008152602001600081525090565b60a060405190810160405280600081526020016000600160a060020a0316815260200160001515815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061294b57805160ff1916838001178555612978565b82800160010185558215612978579182015b8281111561297857825182559160200191906001019061295d565b50612984929150612988565b5090565b6129a291905b80821115612984576000815560010161298e565b9056004572726f723a506172616d65746572697a65722e70726f706f73655265706172a165627a7a72305820e9689ffbe11c3f4626d1aad6354f091e887c83bc06cc493f724511d8485a04d20029`

// DeployParameterizer deploys a new Ethereum contract, binding an instance of Parameterizer to it.
func DeployParameterizer(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddr common.Address, votingAddr common.Address, minDeposit *big.Int, applyStageLen *big.Int, commitStageLen *big.Int, revealStageLen *big.Int, dispensationPct *big.Int, voteQuorum *big.Int) (common.Address, *types.Transaction, *Parameterizer, error) {
	parsed, err := abi.JSON(strings.NewReader(ParameterizerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParameterizerBin), backend, tokenAddr, votingAddr, minDeposit, applyStageLen, commitStageLen, revealStageLen, dispensationPct, voteQuorum)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Parameterizer{ParameterizerCaller: ParameterizerCaller{contract: contract}, ParameterizerTransactor: ParameterizerTransactor{contract: contract}, ParameterizerFilterer: ParameterizerFilterer{contract: contract}}, nil
}

// Parameterizer is an auto generated Go binding around an Ethereum contract.
type Parameterizer struct {
	ParameterizerCaller     // Read-only binding to the contract
	ParameterizerTransactor // Write-only binding to the contract
	ParameterizerFilterer   // Log filterer for contract events
}

// ParameterizerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParameterizerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParameterizerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParameterizerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParameterizerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParameterizerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParameterizerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParameterizerSession struct {
	Contract     *Parameterizer    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParameterizerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParameterizerCallerSession struct {
	Contract *ParameterizerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ParameterizerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParameterizerTransactorSession struct {
	Contract     *ParameterizerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ParameterizerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParameterizerRaw struct {
	Contract *Parameterizer // Generic contract binding to access the raw methods on
}

// ParameterizerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParameterizerCallerRaw struct {
	Contract *ParameterizerCaller // Generic read-only contract binding to access the raw methods on
}

// ParameterizerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParameterizerTransactorRaw struct {
	Contract *ParameterizerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParameterizer creates a new instance of Parameterizer, bound to a specific deployed contract.
func NewParameterizer(address common.Address, backend bind.ContractBackend) (*Parameterizer, error) {
	contract, err := bindParameterizer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Parameterizer{ParameterizerCaller: ParameterizerCaller{contract: contract}, ParameterizerTransactor: ParameterizerTransactor{contract: contract}, ParameterizerFilterer: ParameterizerFilterer{contract: contract}}, nil
}

// NewParameterizerCaller creates a new read-only instance of Parameterizer, bound to a specific deployed contract.
func NewParameterizerCaller(address common.Address, caller bind.ContractCaller) (*ParameterizerCaller, error) {
	contract, err := bindParameterizer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParameterizerCaller{contract: contract}, nil
}

// NewParameterizerTransactor creates a new write-only instance of Parameterizer, bound to a specific deployed contract.
func NewParameterizerTransactor(address common.Address, transactor bind.ContractTransactor) (*ParameterizerTransactor, error) {
	contract, err := bindParameterizer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParameterizerTransactor{contract: contract}, nil
}

// NewParameterizerFilterer creates a new log filterer instance of Parameterizer, bound to a specific deployed contract.
func NewParameterizerFilterer(address common.Address, filterer bind.ContractFilterer) (*ParameterizerFilterer, error) {
	contract, err := bindParameterizer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParameterizerFilterer{contract: contract}, nil
}

// bindParameterizer binds a generic wrapper to an already deployed contract.
func bindParameterizer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParameterizerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parameterizer *ParameterizerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Parameterizer.Contract.ParameterizerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parameterizer *ParameterizerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parameterizer.Contract.ParameterizerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parameterizer *ParameterizerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parameterizer.Contract.ParameterizerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parameterizer *ParameterizerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Parameterizer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parameterizer *ParameterizerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parameterizer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parameterizer *ParameterizerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parameterizer.Contract.contract.Transact(opts, method, params...)
}

// CanBeSet is a free data retrieval call binding the contract method 0xc51131fb.
//
// Solidity: function canBeSet(propHash bytes32) constant returns(bool)
func (_Parameterizer *ParameterizerCaller) CanBeSet(opts *bind.CallOpts, propHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "canBeSet", propHash)
	return *ret0, err
}

// CanBeSet is a free data retrieval call binding the contract method 0xc51131fb.
//
// Solidity: function canBeSet(propHash bytes32) constant returns(bool)
func (_Parameterizer *ParameterizerSession) CanBeSet(propHash [32]byte) (bool, error) {
	return _Parameterizer.Contract.CanBeSet(&_Parameterizer.CallOpts, propHash)
}

// CanBeSet is a free data retrieval call binding the contract method 0xc51131fb.
//
// Solidity: function canBeSet(propHash bytes32) constant returns(bool)
func (_Parameterizer *ParameterizerCallerSession) CanBeSet(propHash [32]byte) (bool, error) {
	return _Parameterizer.Contract.CanBeSet(&_Parameterizer.CallOpts, propHash)
}

// ChallengeCanBeResolved is a free data retrieval call binding the contract method 0x77609a41.
//
// Solidity: function challengeCanBeResolved(propHash bytes32) constant returns(bool)
func (_Parameterizer *ParameterizerCaller) ChallengeCanBeResolved(opts *bind.CallOpts, propHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "challengeCanBeResolved", propHash)
	return *ret0, err
}

// ChallengeCanBeResolved is a free data retrieval call binding the contract method 0x77609a41.
//
// Solidity: function challengeCanBeResolved(propHash bytes32) constant returns(bool)
func (_Parameterizer *ParameterizerSession) ChallengeCanBeResolved(propHash [32]byte) (bool, error) {
	return _Parameterizer.Contract.ChallengeCanBeResolved(&_Parameterizer.CallOpts, propHash)
}

// ChallengeCanBeResolved is a free data retrieval call binding the contract method 0x77609a41.
//
// Solidity: function challengeCanBeResolved(propHash bytes32) constant returns(bool)
func (_Parameterizer *ParameterizerCallerSession) ChallengeCanBeResolved(propHash [32]byte) (bool, error) {
	return _Parameterizer.Contract.ChallengeCanBeResolved(&_Parameterizer.CallOpts, propHash)
}

// ChallengeWinnerReward is a free data retrieval call binding the contract method 0x50411552.
//
// Solidity: function challengeWinnerReward(id uint256) constant returns(uint256)
func (_Parameterizer *ParameterizerCaller) ChallengeWinnerReward(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "challengeWinnerReward", id)
	return *ret0, err
}

// ChallengeWinnerReward is a free data retrieval call binding the contract method 0x50411552.
//
// Solidity: function challengeWinnerReward(id uint256) constant returns(uint256)
func (_Parameterizer *ParameterizerSession) ChallengeWinnerReward(id *big.Int) (*big.Int, error) {
	return _Parameterizer.Contract.ChallengeWinnerReward(&_Parameterizer.CallOpts, id)
}

// ChallengeWinnerReward is a free data retrieval call binding the contract method 0x50411552.
//
// Solidity: function challengeWinnerReward(id uint256) constant returns(uint256)
func (_Parameterizer *ParameterizerCallerSession) ChallengeWinnerReward(id *big.Int) (*big.Int, error) {
	return _Parameterizer.Contract.ChallengeWinnerReward(&_Parameterizer.CallOpts, id)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(name string) constant returns(value uint256)
func (_Parameterizer *ParameterizerCaller) Get(opts *bind.CallOpts, name string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "get", name)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(name string) constant returns(value uint256)
func (_Parameterizer *ParameterizerSession) Get(name string) (*big.Int, error) {
	return _Parameterizer.Contract.Get(&_Parameterizer.CallOpts, name)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(name string) constant returns(value uint256)
func (_Parameterizer *ParameterizerCallerSession) Get(name string) (*big.Int, error) {
	return _Parameterizer.Contract.Get(&_Parameterizer.CallOpts, name)
}

// PropExists is a free data retrieval call binding the contract method 0x35300990.
//
// Solidity: function propExists(propHash bytes32) constant returns(bool)
func (_Parameterizer *ParameterizerCaller) PropExists(opts *bind.CallOpts, propHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "propExists", propHash)
	return *ret0, err
}

// PropExists is a free data retrieval call binding the contract method 0x35300990.
//
// Solidity: function propExists(propHash bytes32) constant returns(bool)
func (_Parameterizer *ParameterizerSession) PropExists(propHash [32]byte) (bool, error) {
	return _Parameterizer.Contract.PropExists(&_Parameterizer.CallOpts, propHash)
}

// PropExists is a free data retrieval call binding the contract method 0x35300990.
//
// Solidity: function propExists(propHash bytes32) constant returns(bool)
func (_Parameterizer *ParameterizerCallerSession) PropExists(propHash [32]byte) (bool, error) {
	return _Parameterizer.Contract.PropExists(&_Parameterizer.CallOpts, propHash)
}

// RewardsClaimed is a free data retrieval call binding the contract method 0x9a835e91.
//
// Solidity: function rewardsClaimed(id uint256, voter address) constant returns(bool)
func (_Parameterizer *ParameterizerCaller) RewardsClaimed(opts *bind.CallOpts, id *big.Int, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "rewardsClaimed", id, voter)
	return *ret0, err
}

// RewardsClaimed is a free data retrieval call binding the contract method 0x9a835e91.
//
// Solidity: function rewardsClaimed(id uint256, voter address) constant returns(bool)
func (_Parameterizer *ParameterizerSession) RewardsClaimed(id *big.Int, voter common.Address) (bool, error) {
	return _Parameterizer.Contract.RewardsClaimed(&_Parameterizer.CallOpts, id, voter)
}

// RewardsClaimed is a free data retrieval call binding the contract method 0x9a835e91.
//
// Solidity: function rewardsClaimed(id uint256, voter address) constant returns(bool)
func (_Parameterizer *ParameterizerCallerSession) RewardsClaimed(id *big.Int, voter common.Address) (bool, error) {
	return _Parameterizer.Contract.RewardsClaimed(&_Parameterizer.CallOpts, id, voter)
}

// VoterReward is a free data retrieval call binding the contract method 0xa7aad3db.
//
// Solidity: function voterReward(voter address, id uint256, salt uint256) constant returns(uint256)
func (_Parameterizer *ParameterizerCaller) VoterReward(opts *bind.CallOpts, voter common.Address, id *big.Int, salt *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "voterReward", voter, id, salt)
	return *ret0, err
}

// VoterReward is a free data retrieval call binding the contract method 0xa7aad3db.
//
// Solidity: function voterReward(voter address, id uint256, salt uint256) constant returns(uint256)
func (_Parameterizer *ParameterizerSession) VoterReward(voter common.Address, id *big.Int, salt *big.Int) (*big.Int, error) {
	return _Parameterizer.Contract.VoterReward(&_Parameterizer.CallOpts, voter, id, salt)
}

// VoterReward is a free data retrieval call binding the contract method 0xa7aad3db.
//
// Solidity: function voterReward(voter address, id uint256, salt uint256) constant returns(uint256)
func (_Parameterizer *ParameterizerCallerSession) VoterReward(voter common.Address, id *big.Int, salt *big.Int) (*big.Int, error) {
	return _Parameterizer.Contract.VoterReward(&_Parameterizer.CallOpts, voter, id, salt)
}

// ChallengeReparameterization is a paid mutator transaction binding the contract method 0x8240ae4b.
//
// Solidity: function challengeReparameterization(propHash bytes32) returns(uint256)
func (_Parameterizer *ParameterizerTransactor) ChallengeReparameterization(opts *bind.TransactOpts, propHash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "challengeReparameterization", propHash)
}

// ChallengeReparameterization is a paid mutator transaction binding the contract method 0x8240ae4b.
//
// Solidity: function challengeReparameterization(propHash bytes32) returns(uint256)
func (_Parameterizer *ParameterizerSession) ChallengeReparameterization(propHash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.Contract.ChallengeReparameterization(&_Parameterizer.TransactOpts, propHash)
}

// ChallengeReparameterization is a paid mutator transaction binding the contract method 0x8240ae4b.
//
// Solidity: function challengeReparameterization(propHash bytes32) returns(uint256)
func (_Parameterizer *ParameterizerTransactorSession) ChallengeReparameterization(propHash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.Contract.ChallengeReparameterization(&_Parameterizer.TransactOpts, propHash)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x86bb8f37.
//
// Solidity: function claimReward(id uint256, salt uint256) returns()
func (_Parameterizer *ParameterizerTransactor) ClaimReward(opts *bind.TransactOpts, id *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "claimReward", id, salt)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x86bb8f37.
//
// Solidity: function claimReward(id uint256, salt uint256) returns()
func (_Parameterizer *ParameterizerSession) ClaimReward(id *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.ClaimReward(&_Parameterizer.TransactOpts, id, salt)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x86bb8f37.
//
// Solidity: function claimReward(id uint256, salt uint256) returns()
func (_Parameterizer *ParameterizerTransactorSession) ClaimReward(id *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.ClaimReward(&_Parameterizer.TransactOpts, id, salt)
}

// ProcessProposal is a paid mutator transaction binding the contract method 0x30490e91.
//
// Solidity: function processProposal(propHash bytes32) returns()
func (_Parameterizer *ParameterizerTransactor) ProcessProposal(opts *bind.TransactOpts, propHash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "processProposal", propHash)
}

// ProcessProposal is a paid mutator transaction binding the contract method 0x30490e91.
//
// Solidity: function processProposal(propHash bytes32) returns()
func (_Parameterizer *ParameterizerSession) ProcessProposal(propHash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.Contract.ProcessProposal(&_Parameterizer.TransactOpts, propHash)
}

// ProcessProposal is a paid mutator transaction binding the contract method 0x30490e91.
//
// Solidity: function processProposal(propHash bytes32) returns()
func (_Parameterizer *ParameterizerTransactorSession) ProcessProposal(propHash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.Contract.ProcessProposal(&_Parameterizer.TransactOpts, propHash)
}

// ProposeReparameterization is a paid mutator transaction binding the contract method 0xbade1c54.
//
// Solidity: function proposeReparameterization(name string, value uint256) returns(bytes32)
func (_Parameterizer *ParameterizerTransactor) ProposeReparameterization(opts *bind.TransactOpts, name string, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "proposeReparameterization", name, value)
}

// ProposeReparameterization is a paid mutator transaction binding the contract method 0xbade1c54.
//
// Solidity: function proposeReparameterization(name string, value uint256) returns(bytes32)
func (_Parameterizer *ParameterizerSession) ProposeReparameterization(name string, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.ProposeReparameterization(&_Parameterizer.TransactOpts, name, value)
}

// ProposeReparameterization is a paid mutator transaction binding the contract method 0xbade1c54.
//
// Solidity: function proposeReparameterization(name string, value uint256) returns(bytes32)
func (_Parameterizer *ParameterizerTransactorSession) ProposeReparameterization(name string, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.ProposeReparameterization(&_Parameterizer.TransactOpts, name, value)
}

// ParameterizerChallengeFailedEventIterator is returned from FilterChallengeFailedEvent and is used to iterate over the raw logs and unpacked data for ChallengeFailedEvent events raised by the Parameterizer contract.
type ParameterizerChallengeFailedEventIterator struct {
	Event *ParameterizerChallengeFailedEvent // Event containing the contract specifics and raw log

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
func (it *ParameterizerChallengeFailedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerChallengeFailedEvent)
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
		it.Event = new(ParameterizerChallengeFailedEvent)
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
func (it *ParameterizerChallengeFailedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerChallengeFailedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerChallengeFailedEvent represents a ChallengeFailedEvent event raised by the Parameterizer contract.
type ParameterizerChallengeFailedEvent struct {
	PropHash    [32]byte
	Id          *big.Int
	RewardPool  *big.Int
	TotalTokens *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChallengeFailedEvent is a free log retrieval operation binding the contract event 0x824960b36dcff66c0b43953e552b1b0d622dcb3f3cf71cf1cbe88039a781a853.
//
// Solidity: e ChallengeFailedEvent(propHash indexed bytes32, id indexed uint256, rewardPool uint256, totalTokens uint256)
func (_Parameterizer *ParameterizerFilterer) FilterChallengeFailedEvent(opts *bind.FilterOpts, propHash [][32]byte, id []*big.Int) (*ParameterizerChallengeFailedEventIterator, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ChallengeFailedEvent", propHashRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerChallengeFailedEventIterator{contract: _Parameterizer.contract, event: "ChallengeFailedEvent", logs: logs, sub: sub}, nil
}

// WatchChallengeFailedEvent is a free log subscription operation binding the contract event 0x824960b36dcff66c0b43953e552b1b0d622dcb3f3cf71cf1cbe88039a781a853.
//
// Solidity: e ChallengeFailedEvent(propHash indexed bytes32, id indexed uint256, rewardPool uint256, totalTokens uint256)
func (_Parameterizer *ParameterizerFilterer) WatchChallengeFailedEvent(opts *bind.WatchOpts, sink chan<- *ParameterizerChallengeFailedEvent, propHash [][32]byte, id []*big.Int) (event.Subscription, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ChallengeFailedEvent", propHashRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerChallengeFailedEvent)
				if err := _Parameterizer.contract.UnpackLog(event, "ChallengeFailedEvent", log); err != nil {
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

// ParameterizerChallengeSucceededEventIterator is returned from FilterChallengeSucceededEvent and is used to iterate over the raw logs and unpacked data for ChallengeSucceededEvent events raised by the Parameterizer contract.
type ParameterizerChallengeSucceededEventIterator struct {
	Event *ParameterizerChallengeSucceededEvent // Event containing the contract specifics and raw log

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
func (it *ParameterizerChallengeSucceededEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerChallengeSucceededEvent)
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
		it.Event = new(ParameterizerChallengeSucceededEvent)
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
func (it *ParameterizerChallengeSucceededEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerChallengeSucceededEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerChallengeSucceededEvent represents a ChallengeSucceededEvent event raised by the Parameterizer contract.
type ParameterizerChallengeSucceededEvent struct {
	PropHash    [32]byte
	Id          *big.Int
	RewardPool  *big.Int
	TotalTokens *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChallengeSucceededEvent is a free log retrieval operation binding the contract event 0x3dddcf5cd7cb20cfa68f0643b9422cdbc130a6602933781bf9855e10fd205e42.
//
// Solidity: e ChallengeSucceededEvent(propHash indexed bytes32, id indexed uint256, rewardPool uint256, totalTokens uint256)
func (_Parameterizer *ParameterizerFilterer) FilterChallengeSucceededEvent(opts *bind.FilterOpts, propHash [][32]byte, id []*big.Int) (*ParameterizerChallengeSucceededEventIterator, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ChallengeSucceededEvent", propHashRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerChallengeSucceededEventIterator{contract: _Parameterizer.contract, event: "ChallengeSucceededEvent", logs: logs, sub: sub}, nil
}

// WatchChallengeSucceededEvent is a free log subscription operation binding the contract event 0x3dddcf5cd7cb20cfa68f0643b9422cdbc130a6602933781bf9855e10fd205e42.
//
// Solidity: e ChallengeSucceededEvent(propHash indexed bytes32, id indexed uint256, rewardPool uint256, totalTokens uint256)
func (_Parameterizer *ParameterizerFilterer) WatchChallengeSucceededEvent(opts *bind.WatchOpts, sink chan<- *ParameterizerChallengeSucceededEvent, propHash [][32]byte, id []*big.Int) (event.Subscription, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ChallengeSucceededEvent", propHashRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerChallengeSucceededEvent)
				if err := _Parameterizer.contract.UnpackLog(event, "ChallengeSucceededEvent", log); err != nil {
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

// ParameterizerNewChallengeEventIterator is returned from FilterNewChallengeEvent and is used to iterate over the raw logs and unpacked data for NewChallengeEvent events raised by the Parameterizer contract.
type ParameterizerNewChallengeEventIterator struct {
	Event *ParameterizerNewChallengeEvent // Event containing the contract specifics and raw log

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
func (it *ParameterizerNewChallengeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerNewChallengeEvent)
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
		it.Event = new(ParameterizerNewChallengeEvent)
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
func (it *ParameterizerNewChallengeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerNewChallengeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerNewChallengeEvent represents a NewChallengeEvent event raised by the Parameterizer contract.
type ParameterizerNewChallengeEvent struct {
	PropHash     [32]byte
	Id           *big.Int
	CommitExpiry *big.Int
	RevealExpiry *big.Int
	Challenger   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewChallengeEvent is a free log retrieval operation binding the contract event 0x4e6815e2c453c1de363fd2b483078339512560f8fd777420d06dd69291e060c2.
//
// Solidity: e NewChallengeEvent(propHash indexed bytes32, id uint256, commitExpiry uint256, revealExpiry uint256, challenger indexed address)
func (_Parameterizer *ParameterizerFilterer) FilterNewChallengeEvent(opts *bind.FilterOpts, propHash [][32]byte, challenger []common.Address) (*ParameterizerNewChallengeEventIterator, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "NewChallengeEvent", propHashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerNewChallengeEventIterator{contract: _Parameterizer.contract, event: "NewChallengeEvent", logs: logs, sub: sub}, nil
}

// WatchNewChallengeEvent is a free log subscription operation binding the contract event 0x4e6815e2c453c1de363fd2b483078339512560f8fd777420d06dd69291e060c2.
//
// Solidity: e NewChallengeEvent(propHash indexed bytes32, id uint256, commitExpiry uint256, revealExpiry uint256, challenger indexed address)
func (_Parameterizer *ParameterizerFilterer) WatchNewChallengeEvent(opts *bind.WatchOpts, sink chan<- *ParameterizerNewChallengeEvent, propHash [][32]byte, challenger []common.Address) (event.Subscription, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "NewChallengeEvent", propHashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerNewChallengeEvent)
				if err := _Parameterizer.contract.UnpackLog(event, "NewChallengeEvent", log); err != nil {
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

// ParameterizerProposalAcceptedEventIterator is returned from FilterProposalAcceptedEvent and is used to iterate over the raw logs and unpacked data for ProposalAcceptedEvent events raised by the Parameterizer contract.
type ParameterizerProposalAcceptedEventIterator struct {
	Event *ParameterizerProposalAcceptedEvent // Event containing the contract specifics and raw log

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
func (it *ParameterizerProposalAcceptedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerProposalAcceptedEvent)
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
		it.Event = new(ParameterizerProposalAcceptedEvent)
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
func (it *ParameterizerProposalAcceptedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerProposalAcceptedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerProposalAcceptedEvent represents a ProposalAcceptedEvent event raised by the Parameterizer contract.
type ParameterizerProposalAcceptedEvent struct {
	PropHash [32]byte
	Name     string
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProposalAcceptedEvent is a free log retrieval operation binding the contract event 0x0fca8490116d823913d2468dc21632e9226d20e54379af9cf34712a3ff2bc30c.
//
// Solidity: e ProposalAcceptedEvent(propHash indexed bytes32, name string, value uint256)
func (_Parameterizer *ParameterizerFilterer) FilterProposalAcceptedEvent(opts *bind.FilterOpts, propHash [][32]byte) (*ParameterizerProposalAcceptedEventIterator, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ProposalAcceptedEvent", propHashRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerProposalAcceptedEventIterator{contract: _Parameterizer.contract, event: "ProposalAcceptedEvent", logs: logs, sub: sub}, nil
}

// WatchProposalAcceptedEvent is a free log subscription operation binding the contract event 0x0fca8490116d823913d2468dc21632e9226d20e54379af9cf34712a3ff2bc30c.
//
// Solidity: e ProposalAcceptedEvent(propHash indexed bytes32, name string, value uint256)
func (_Parameterizer *ParameterizerFilterer) WatchProposalAcceptedEvent(opts *bind.WatchOpts, sink chan<- *ParameterizerProposalAcceptedEvent, propHash [][32]byte) (event.Subscription, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ProposalAcceptedEvent", propHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerProposalAcceptedEvent)
				if err := _Parameterizer.contract.UnpackLog(event, "ProposalAcceptedEvent", log); err != nil {
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

// ParameterizerProposalExpiredEventIterator is returned from FilterProposalExpiredEvent and is used to iterate over the raw logs and unpacked data for ProposalExpiredEvent events raised by the Parameterizer contract.
type ParameterizerProposalExpiredEventIterator struct {
	Event *ParameterizerProposalExpiredEvent // Event containing the contract specifics and raw log

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
func (it *ParameterizerProposalExpiredEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerProposalExpiredEvent)
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
		it.Event = new(ParameterizerProposalExpiredEvent)
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
func (it *ParameterizerProposalExpiredEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerProposalExpiredEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerProposalExpiredEvent represents a ProposalExpiredEvent event raised by the Parameterizer contract.
type ParameterizerProposalExpiredEvent struct {
	PropHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProposalExpiredEvent is a free log retrieval operation binding the contract event 0x456cbba32d3f5372c3ef95ba54f18060130b9da12721ebf89f1146a8f4061324.
//
// Solidity: e ProposalExpiredEvent(propHash indexed bytes32)
func (_Parameterizer *ParameterizerFilterer) FilterProposalExpiredEvent(opts *bind.FilterOpts, propHash [][32]byte) (*ParameterizerProposalExpiredEventIterator, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ProposalExpiredEvent", propHashRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerProposalExpiredEventIterator{contract: _Parameterizer.contract, event: "ProposalExpiredEvent", logs: logs, sub: sub}, nil
}

// WatchProposalExpiredEvent is a free log subscription operation binding the contract event 0x456cbba32d3f5372c3ef95ba54f18060130b9da12721ebf89f1146a8f4061324.
//
// Solidity: e ProposalExpiredEvent(propHash indexed bytes32)
func (_Parameterizer *ParameterizerFilterer) WatchProposalExpiredEvent(opts *bind.WatchOpts, sink chan<- *ParameterizerProposalExpiredEvent, propHash [][32]byte) (event.Subscription, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ProposalExpiredEvent", propHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerProposalExpiredEvent)
				if err := _Parameterizer.contract.UnpackLog(event, "ProposalExpiredEvent", log); err != nil {
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

// ParameterizerReparameterizationProposalEventIterator is returned from FilterReparameterizationProposalEvent and is used to iterate over the raw logs and unpacked data for ReparameterizationProposalEvent events raised by the Parameterizer contract.
type ParameterizerReparameterizationProposalEventIterator struct {
	Event *ParameterizerReparameterizationProposalEvent // Event containing the contract specifics and raw log

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
func (it *ParameterizerReparameterizationProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerReparameterizationProposalEvent)
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
		it.Event = new(ParameterizerReparameterizationProposalEvent)
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
func (it *ParameterizerReparameterizationProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerReparameterizationProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerReparameterizationProposalEvent represents a ReparameterizationProposalEvent event raised by the Parameterizer contract.
type ParameterizerReparameterizationProposalEvent struct {
	Name      string
	Value     *big.Int
	PropHash  [32]byte
	Deposit   *big.Int
	AppExpiry *big.Int
	Proposer  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReparameterizationProposalEvent is a free log retrieval operation binding the contract event 0xf6aafb6c8410e85ab57ef6d7df18fcb782806ee3b89634f3b52046846236f97a.
//
// Solidity: e ReparameterizationProposalEvent(name string, value uint256, propHash bytes32, deposit uint256, appExpiry uint256, proposer indexed address)
func (_Parameterizer *ParameterizerFilterer) FilterReparameterizationProposalEvent(opts *bind.FilterOpts, proposer []common.Address) (*ParameterizerReparameterizationProposalEventIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ReparameterizationProposalEvent", proposerRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerReparameterizationProposalEventIterator{contract: _Parameterizer.contract, event: "ReparameterizationProposalEvent", logs: logs, sub: sub}, nil
}

// WatchReparameterizationProposalEvent is a free log subscription operation binding the contract event 0xf6aafb6c8410e85ab57ef6d7df18fcb782806ee3b89634f3b52046846236f97a.
//
// Solidity: e ReparameterizationProposalEvent(name string, value uint256, propHash bytes32, deposit uint256, appExpiry uint256, proposer indexed address)
func (_Parameterizer *ParameterizerFilterer) WatchReparameterizationProposalEvent(opts *bind.WatchOpts, sink chan<- *ParameterizerReparameterizationProposalEvent, proposer []common.Address) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ReparameterizationProposalEvent", proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerReparameterizationProposalEvent)
				if err := _Parameterizer.contract.UnpackLog(event, "ReparameterizationProposalEvent", log); err != nil {
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

// ParameterizerRewardsClaimedEventIterator is returned from FilterRewardsClaimedEvent and is used to iterate over the raw logs and unpacked data for RewardsClaimedEvent events raised by the Parameterizer contract.
type ParameterizerRewardsClaimedEventIterator struct {
	Event *ParameterizerRewardsClaimedEvent // Event containing the contract specifics and raw log

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
func (it *ParameterizerRewardsClaimedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerRewardsClaimedEvent)
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
		it.Event = new(ParameterizerRewardsClaimedEvent)
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
func (it *ParameterizerRewardsClaimedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerRewardsClaimedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerRewardsClaimedEvent represents a RewardsClaimedEvent event raised by the Parameterizer contract.
type ParameterizerRewardsClaimedEvent struct {
	Id     *big.Int
	Reward *big.Int
	Voter  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimedEvent is a free log retrieval operation binding the contract event 0x17ccb5aed827c5078139dc28b7fb924b006f56b3dcf163c4acf21fa3beb13144.
//
// Solidity: e RewardsClaimedEvent(id indexed uint256, reward uint256, voter indexed address)
func (_Parameterizer *ParameterizerFilterer) FilterRewardsClaimedEvent(opts *bind.FilterOpts, id []*big.Int, voter []common.Address) (*ParameterizerRewardsClaimedEventIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "RewardsClaimedEvent", idRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerRewardsClaimedEventIterator{contract: _Parameterizer.contract, event: "RewardsClaimedEvent", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimedEvent is a free log subscription operation binding the contract event 0x17ccb5aed827c5078139dc28b7fb924b006f56b3dcf163c4acf21fa3beb13144.
//
// Solidity: e RewardsClaimedEvent(id indexed uint256, reward uint256, voter indexed address)
func (_Parameterizer *ParameterizerFilterer) WatchRewardsClaimedEvent(opts *bind.WatchOpts, sink chan<- *ParameterizerRewardsClaimedEvent, id []*big.Int, voter []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "RewardsClaimedEvent", idRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerRewardsClaimedEvent)
				if err := _Parameterizer.contract.UnpackLog(event, "RewardsClaimedEvent", log); err != nil {
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
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820a242470ccd5de57331f6473b35ebb9044e0504422e500ebad03b24591f26c27c0029`

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
