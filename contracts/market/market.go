// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package market

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

// MarketABI is the input ABI used to generate the binding from.
const MarketABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isListingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"divest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getInvestment\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"getListing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"resolveApplication\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offered\",\"type\":\"uint256\"}],\"name\":\"invest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToListing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"getChallenge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listing\",\"type\":\"string\"},{\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"list\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"isListing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listing\",\"type\":\"string\"}],\"name\":\"getListingHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInvestmentPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromListing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInvested\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isInvestor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"challenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"resolveChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"isListed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getListings\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"marketTokenAddr\",\"type\":\"address\"},{\"name\":\"etherTokenAddr\",\"type\":\"address\"},{\"name\":\"parameterizerAddr\",\"type\":\"address\"},{\"name\":\"votingAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"ApplicationFailedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"applicant\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"voteBy\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AppliedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"ChallengedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"ChallengeFailedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"ChallengeSucceededEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"transferred\",\"type\":\"uint256\"}],\"name\":\"DivestedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"offered\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minted\",\"type\":\"uint256\"}],\"name\":\"InvestedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"ListedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"added\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"ListingDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"ListingDeletedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"added\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"ListingWithdrawEvent\",\"type\":\"event\"}]"

// MarketBin is the compiled bytecode used for deploying new contracts.
const MarketBin = `0x608060405234801561001057600080fd5b506040516080806138ce8339810180604052608081101561003057600080fd5b50805160208201516040830151606090930151600b805433600160a060020a031991821617909155600880548216600160a060020a039586161790556007805482169385169390931790925560098054831694841694909417909355600a805490911691909216179055613825806100a96000396000f3fe608060405234801561001057600080fd5b50600436106101305760003560e060020a900480636aa0966b116100b1578063cee2a9cf11610075578063cee2a9cf146103f8578063cffd46dc1461041e578063d32c943a1461043b578063ecefbdc614610458578063f1b2d6a31461047557610130565b80636aa0966b14610338578063a2c45b3914610355578063a3d6602b146103c5578063a7e78675146103cd578063befc3e2b146103f057610130565b806318e28e1c116100f857806318e28e1c146102205780632afcf4801461023d578063341ff0891461025a578063458d2bf11461027d5780635148cd0f146102c257610130565b806302a47cd714610135578063058aace11461016f5780630ca3626314610179578063146b58df14610196578063175c0d16146101ce575b600080fd5b61015b6004803603602081101561014b57600080fd5b5035600160a060020a03166104cd565b604080519115158252519081900360200190f35b6101776104f0565b005b6101776004803603602081101561018f57600080fd5b5035610b1c565b6101bc600480360360208110156101ac57600080fd5b5035600160a060020a0316610ca7565b60408051918252519081900360200190f35b6101eb600480360360208110156101e457600080fd5b5035610cc5565b604080519515158652600160a060020a0390941660208601528484019290925260608401526080830152519081900360a00190f35b6101776004803603602081101561023657600080fd5b5035610d03565b6101776004803603602081101561025357600080fd5b5035611447565b6101776004803603604081101561027057600080fd5b5080359060200135611815565b61029a6004803603602081101561029357600080fd5b503561196c565b60408051600160a060020a039094168452602084019290925282820152519081900360600190f35b610177600480360360608110156102d857600080fd5b8101906020810181356401000000008111156102f357600080fd5b82018360208201111561030557600080fd5b8035906020019184600183028401116401000000008311171561032757600080fd5b919350915080359060200135611995565b61015b6004803603602081101561034e57600080fd5b5035611ddb565b6101bc6004803603602081101561036b57600080fd5b81019060208101813564010000000081111561038657600080fd5b82018360208201111561039857600080fd5b803590602001918460018302840111640100000000831117156103ba57600080fd5b509092509050611e20565b6101bc611e46565b610177600480360360408110156103e357600080fd5b50803590602001356120bf565b6101bc6122d4565b61015b6004803603602081101561040e57600080fd5b5035600160a060020a03166122da565b6101776004803603602081101561043457600080fd5b5035612332565b6101776004803603602081101561045157600080fd5b503561279d565b61015b6004803603602081101561046e57600080fd5b5035612e46565b61047d612e63565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156104b95781810151838201526020016104a1565b505050509050019250505060405180910390f35b600160a060020a038116600090815260036020526040902054600111155b919050565b6104f9336122da565b15156105395760405160e560020a62461bcd0281526004018080602001828103825260298152602001806133766029913960400191505060405180910390fd5b60006060600a60009054906101000a9004600160a060020a0316600160a060020a03166306a49fce6040518163ffffffff1660e060020a02815260040160006040518083038186803b15801561058e57600080fd5b505afa1580156105a2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156105cb57600080fd5b8101908080516401000000008111156105e357600080fd5b820160208101848111156105f657600080fd5b815185602082028301116401000000008211171561061357600080fd5b50909450600093505050505b815181101561067d5733600160a060020a0316600080848481518110151561064357fe5b6020908102909101810151825281019190915260400160002054600160a060020a03161415610675576001925061067d565b60010161061f565b50600182151514156106c35760405160e560020a62461bcd0281526004018080602001828103825260478152602001806136736047913960600191505060405180910390fd5b600754604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600092600160a060020a0316916370a08231916024808301926020929190829003018186803b15801561072757600080fd5b505afa15801561073b573d6000803e3d6000fd5b505050506040513d602081101561075157600080fd5b5051600854604080517f18160ddd0000000000000000000000000000000000000000000000000000000081529051929350600092600160a060020a03909216916318160ddd91600480820192602092909190829003018186803b1580156107b757600080fd5b505afa1580156107cb573d6000803e3d6000fd5b505050506040513d60208110156107e157600080fd5b5051600854604080517f70a082310000000000000000000000000000000000000000000000000000000081523360048201529051929350600092600160a060020a03909216916370a0823191602480820192602092909190829003018186803b15801561084d57600080fd5b505afa158015610861573d6000803e3d6000fd5b505050506040513d602081101561087757600080fd5b5051600854604080517f7e9d2ac10000000000000000000000000000000000000000000000000000000081523360048201529051929350600160a060020a0390911691637e9d2ac19160248082019260009290919082900301818387803b1580156108e157600080fd5b505af11580156108f5573d6000803e3d6000fd5b50503360009081526004602052604081205460058054919450919250600019810190811061091f57fe5b60009182526020909120015460058054600160a060020a03909216925082918490811061094857fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560058054906109919060001983016132bf565b50600160a060020a03808216600090815260046020819052604080832086905533808452818420848155600101849055600a5482517f335d808000000000000000000000000000000000000000000000000000000000815293840191909152905193169263335d80809260248084019391929182900301818387803b158015610a1957600080fd5b505af1158015610a2d573d6000803e3d6000fd5b505050506000610a5685610a4a8887612ebb90919063ffffffff16565b9063ffffffff612eea16565b6007546040805160e060020a63a9059cbb028152336004820152602481018490529051929350600160a060020a039091169163a9059cbb916044808201926020929091908290030181600087803b158015610ab057600080fd5b505af1158015610ac4573d6000803e3d6000fd5b505050506040513d6020811015610ada57600080fd5b505060408051828152905133917f31b5a476fc1c356476ae05b02dcec8b3c6aaab43b21d4d270089270be10da644919081900360200190a25050505050505050565b600081815260016020819052604090912001546101009004600160a060020a03163314610b7d5760405160e560020a62461bcd0281526004018080602001828103825260298152602001806134006029913960400191505060405180910390fd5b610b8681612e46565b1515610bc65760405160e560020a62461bcd0281526004018080602001828103825260228152602001806135066022913960400191505060405180910390fd5b600a54604080517fb89694c6000000000000000000000000000000000000000000000000000000008152600481018490529051600160a060020a039092169163b89694c691602480820192602092909190829003018186803b158015610c2b57600080fd5b505afa158015610c3f573d6000803e3d6000fd5b505050506040513d6020811015610c5557600080fd5b5051151560011415610c9b5760405160e560020a62461bcd0281526004018080602001828103825260328152602001806133446032913960400191505060405180910390fd5b610ca481612f0b565b50565b600160a060020a031660009081526004602052604090206001015490565b6000908152600160208190526040909120908101546002820154600383015460049093015460ff831694610100909304600160a060020a0316939192565b600a54604080517f62f9a55d0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a03909216916362f9a55d91602480820192602092909190829003018186803b158015610d6757600080fd5b505afa158015610d7b573d6000803e3d6000fd5b505050506040513d6020811015610d9157600080fd5b50511515610dd35760405160e560020a62461bcd02815260040180806020018281038252603f8152602001806135c6603f913960400191505060405180910390fd5b600a54604080517f06db129800000000000000000000000000000000000000000000000000000000815260048101849052600160248201529051600160a060020a03909216916306db129891604480820192602092909190829003018186803b158015610e3f57600080fd5b505afa158015610e53573d6000803e3d6000fd5b505050506040513d6020811015610e6957600080fd5b50511515600114610eae5760405160e560020a62461bcd0281526004018080602001828103825260388152602001806136056038913960400191505060405180910390fd5b600081815260016020819052604090912081015460ff1615151415610f075760405160e560020a62461bcd0281526004018080602001828103825260308152602001806134296030913960400191505060405180910390fd5b600a54604080517f327322c8000000000000000000000000000000000000000000000000000000008152600481018490529051600160a060020a039092169163327322c891602480820192602092909190829003018186803b158015610f6c57600080fd5b505afa158015610f80573d6000803e3d6000fd5b505050506040513d6020811015610f9657600080fd5b50511515610fd85760405160e560020a62461bcd02815260040180806020018281038252604981526020018061357d6049913960600191505060405180910390fd5b600a54600954604080517fc26c12eb0000000000000000000000000000000000000000000000000000000081529051600160a060020a0393841693638f354b7993869391169163c26c12eb91600480820192602092909190829003018186803b15801561104457600080fd5b505afa158015611058573d6000803e3d6000fd5b505050506040513d602081101561106e57600080fd5b50516040805160e060020a63ffffffff861602815260048101939093526024830191909152516044808301926020929190829003018186803b1580156110b357600080fd5b505afa1580156110c7573d6000803e3d6000fd5b505050506040513d60208110156110dd57600080fd5b5051156113bd5760008181526001602081815260408084208301805460ff191690931790925560095482517fdc2b28530000000000000000000000000000000000000000000000000000000081529251600160a060020a039091169263dc2b2853926004808301939192829003018186803b15801561115b57600080fd5b505afa15801561116f573d6000803e3d6000fd5b505050506040513d602081101561118557600080fd5b5051600854604080517fa0712d68000000000000000000000000000000000000000000000000000000008152600481018490529051929350600160a060020a039091169163a0712d689160248082019260009290919082900301818387803b1580156111f057600080fd5b505af1158015611204573d6000803e3d6000fd5b5050506000838152600160208181526040928390206004808201879055600a54919093015484517f62f9a55d000000000000000000000000000000000000000000000000000000008152600160a060020a036101009092048216948101949094529351931693506362f9a55d92602480840193829003018186803b15801561128b57600080fd5b505afa15801561129f573d6000803e3d6000fd5b505050506040513d60208110156112b557600080fd5b5051151560011461135457600a5460008381526001602081905260408083209091015481517f471bb309000000000000000000000000000000000000000000000000000000008152600160a060020a0361010090920482166004820152915193169263471bb3099260248084019391929182900301818387803b15801561133b57600080fd5b505af115801561134f573d6000803e3d6000fd5b505050505b600082815260016020818152604092839020918201546002909201548351858152918201528251610100909204600160a060020a03169285927fb08f341fa4f58287c76e79e69c47e07939983551f847f728f4dc9e06a49e308d9281900390910190a3506113c6565b6113c681612f0b565b600a54604080517f89bb617c000000000000000000000000000000000000000000000000000000008152600481018490529051600160a060020a03909216916389bb617c9160248082019260009290919082900301818387803b15801561142c57600080fd5b505af1158015611440573d6000803e3d6000fd5b5050505050565b611450336104cd565b1515600114156114945760405160e560020a62461bcd0281526004018080602001828103825260428152602001806134c46042913960600191505060405180910390fd5b600061149e611e46565b9050808210156114e25760405160e560020a62461bcd02815260040180806020018281038252605e8152602001806136ba605e913960600191505060405180910390fd5b6007546040805160e060020a6323b872dd028152336004820152306024820152604481018590529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b15801561153f57600080fd5b505af1158015611553573d6000803e3d6000fd5b505050506040513d602081101561156957600080fd5b5060009050611592633b9aca00611586858563ffffffff612eea16565b9063ffffffff612ebb16565b600854604080517fa0712d68000000000000000000000000000000000000000000000000000000008152600481018490529051929350600160a060020a039091169163a0712d689160248082019260009290919082900301818387803b1580156115fb57600080fd5b505af115801561160f573d6000803e3d6000fd5b50506008546040805160e060020a63a9059cbb028152336004820152602481018690529051600160a060020a03909216935063a9059cbb92506044808201926020929091908290030181600087803b15801561166a57600080fd5b505af115801561167e573d6000803e3d6000fd5b505050506040513d602081101561169457600080fd5b506116a09050336122da565b156116e057336000908152600460205260409020600101546116c8908463ffffffff61329e16565b336000908152600460205260409020600101556117c0565b336000818152600460208190526040808320600180820189905560058054918201815585527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db08101805473ffffffffffffffffffffffffffffffffffffffff1916871790558155600a5482517f471bb3090000000000000000000000000000000000000000000000000000000081529384019590955290519093600160a060020a03169263471bb309926024808201939182900301818387803b1580156117a657600080fd5b505af11580156117ba573d6000803e3d6000fd5b50505050505b6006546117d3908463ffffffff61329e16565b6006556040805184815260208101839052815133927f1be1f5b0e87aa0e79eed7f75d2113999609d67094a95ff7f8e6c5625a35d9a39928290030190a2505050565b61181e82611ddb565b151561185e5760405160e560020a62461bcd02815260040180806020018281038252603681526020018061363d6036913960400191505060405180910390fd5b6008546040805160e060020a6323b872dd028152336004820152306024820152604481018490529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b1580156118bb57600080fd5b505af11580156118cf573d6000803e3d6000fd5b505050506040513d60208110156118e557600080fd5b5050600082815260016020526040902060020154611909908263ffffffff61329e16565b60008381526001602090815260409182902060028101849055600401548251858152918201939093528082019290925251339184917f7f1c6a478f1fbf07f8b135e162e7322085eb5371d2b77163e27dc39d4154d1369181900360600190a35050565b600090815260208190526040902080546001820154600290920154600160a060020a0390911692565b600084846040518083838082843760408051919093018190038120600a547fb89694c6000000000000000000000000000000000000000000000000000000008352600483018290529351909750600160a060020a03909316955063b89694c6945060248082019450602093925090829003018186803b158015611a1757600080fd5b505afa158015611a2b573d6000803e3d6000fd5b505050506040513d6020811015611a4157600080fd5b5051151560011415611a875760405160e560020a62461bcd02815260040180806020018281038252602881526020018061339f6028913960400191505060405180910390fd5b611a9081612e46565b151560011415611ad45760405160e560020a62461bcd0281526004018080602001828103825260228152602001806135286022913960400191505060405180910390fd5b6000818152600160208181526040808420808401805474ffffffffffffffffffffffffffffffffffffffff0019163361010081029190911790915560038083018a9055600280548088019091557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace8101899055835590865290925290922054611b5c9161329e565b33600090815260036020526040812091909155831115611c135760085460018201546040805160e060020a6323b872dd028152610100909204600160a060020a0390811660048401523060248401526044830187905290519216916323b872dd916064808201926020929091908290030181600087803b158015611bdf57600080fd5b505af1158015611bf3573d6000803e3d6000fd5b505050506040513d6020811015611c0957600080fd5b5050600281018390555b600a54600954604080517f2d0d2bc60000000000000000000000000000000000000000000000000000000081529051600160a060020a03938416936364b34e7d9387936001939190921691632d0d2bc6916004808301926020929190829003018186803b158015611c8357600080fd5b505afa158015611c97573d6000803e3d6000fd5b505050506040513d6020811015611cad57600080fd5b50516040805160e060020a63ffffffff8716028152600481019490945260ff9092166024840152604483015251606480830192600092919082900301818387803b158015611cfa57600080fd5b505af1158015611d0e573d6000803e3d6000fd5b505050508333600160a060020a0316837f06198c59f70529156e65d8211777fdd8fdd009977545a9e8c44050f7ff2cb133600960009054906101000a9004600160a060020a0316600160a060020a0316632d0d2bc66040518163ffffffff1660e060020a02815260040160206040518083038186803b158015611d9057600080fd5b505afa158015611da4573d6000803e3d6000fd5b505050506040513d6020811015611dba57600080fd5b505160408051918252602082018990528051918290030190a4505050505050565b6002546000901515611def575060006104eb565b600082815260016020526040902054600280548492908110611e0d57fe5b9060005260206000200154149050919050565b600082826040518083838082843760405192018290039091209450505050505b92915050565b600080600960009054906101000a9004600160a060020a0316600160a060020a031663f36089ec6040518163ffffffff1660e060020a02815260040160206040518083038186803b158015611e9a57600080fd5b505afa158015611eae573d6000803e3d6000fd5b505050506040513d6020811015611ec457600080fd5b5051600954604080517f7f58af4a0000000000000000000000000000000000000000000000000000000081529051929350600092600160a060020a0390921691637f58af4a91600480820192602092909190829003018186803b158015611f2a57600080fd5b505afa158015611f3e573d6000803e3d6000fd5b505050506040513d6020811015611f5457600080fd5b5051600954604080517f2de2b6ac0000000000000000000000000000000000000000000000000000000081529051929350600092600160a060020a0390921691632de2b6ac91600480820192602092909190829003018186803b158015611fba57600080fd5b505afa158015611fce573d6000803e3d6000fd5b505050506040513d6020811015611fe457600080fd5b5051600754604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051929350600092600160a060020a03909216916370a0823191602480820192602092909190829003018186803b15801561205057600080fd5b505afa158015612064573d6000803e3d6000fd5b505050506040513d602081101561207a57600080fd5b505190506120b56120a861209885633b9aca0063ffffffff612ebb16565b610a4a858563ffffffff612ebb16565b859063ffffffff61329e16565b9450505050505b90565b6120c882611ddb565b15156121085760405160e560020a62461bcd0281526004018080602001828103825260398152602001806133c76039913960400191505060405180910390fd5b600082815260016020819052604090912001546101009004600160a060020a031633146121695760405160e560020a62461bcd02815260040180806020018281038252602d8152602001806137cd602d913960400191505060405180910390fd5b6000828152600160205260409020600201548111156121bc5760405160e560020a62461bcd02815260040180806020018281038252604f815260200180613718604f913960600191505060405180910390fd5b6000828152600160205260409020600201546121de908263ffffffff6132ad16565b600083815260016020908152604080832060020193909355600854835160e060020a63a9059cbb028152336004820152602481018690529351600160a060020a039091169363a9059cbb9360448083019493928390030190829087803b15801561224757600080fd5b505af115801561225b573d6000803e3d6000fd5b505050506040513d602081101561227157600080fd5b50506000828152600160209081526040918290206002810154600490910154835185815292830191909152818301529051339184917f261e65ee6edcc0072cc284d3d0665cbc0c581ab1fc708222b6bc3d20c452750d9181900360600190a35050565b60065490565b60055460009015156122ee575060006104eb565b600160a060020a03821660008181526004602052604090205460058054909190811061231657fe5b600091825260209091200154600160a060020a03161492915050565b600954604080517f5aebe7680000000000000000000000000000000000000000000000000000000081529051600092600160a060020a031691635aebe768916004808301926020929190829003018186803b15801561239057600080fd5b505afa1580156123a4573d6000803e3d6000fd5b505050506040513d60208110156123ba57600080fd5b50516000838152600160208190526040909120808201549293509160ff161515146124195760405160e560020a62461bcd02815260040180806020018281038252603b815260200180613792603b913960400191505060405180910390fd5b600a54604080517fb89694c6000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a039092169163b89694c691602480820192602092909190829003018186803b15801561247e57600080fd5b505afa158015612492573d6000803e3d6000fd5b505050506040513d60208110156124a857600080fd5b50511515600114156124ee5760405160e560020a62461bcd02815260040180806020018281038252602b815260200180613767602b913960400191505060405180910390fd5b8161250a8260040154836002015461329e90919063ffffffff16565b101561251e5761251983612f0b565b612798565b6008546040805160e060020a6323b872dd028152336004820152306024820152604481018590529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b15801561257b57600080fd5b505af115801561258f573d6000803e3d6000fd5b505050506040513d60208110156125a557600080fd5b50506002810154600090819084116125d95760028301546125cc908563ffffffff6132ad16565b6002840155839150612619565b6002830154915060006125f2858463ffffffff6132ad16565b600060028601556004850154909150612611908263ffffffff6132ad16565b600485015590505b60408051606081018252338152602080820185815282840185815260008a81528084528590209351845473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039182161785559151600185015551600293840155600a5460095485517f2d0d2bc60000000000000000000000000000000000000000000000000000000081529551918316956364b34e7d958c9590949290921692632d0d2bc692600480840193919291829003018186803b1580156126d957600080fd5b505afa1580156126ed573d6000803e3d6000fd5b505050506040513d602081101561270357600080fd5b50516040805160e060020a63ffffffff8716028152600481019490945260ff9092166024840152604483015251606480830192600092919082900301818387803b15801561275057600080fd5b505af1158015612764573d6000803e3d6000fd5b50506040513392508791507f4dd89117b2f54a955a15b7ca70efe9d9cda23696ebd2a8290ce6bb42b1594dba90600090a350505b505050565b600a54604080517f62f9a55d0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a03909216916362f9a55d91602480820192602092909190829003018186803b15801561280157600080fd5b505afa158015612815573d6000803e3d6000fd5b505050506040513d602081101561282b57600080fd5b5051151561286d5760405160e560020a62461bcd02815260040180806020018281038252603d815260200180613459603d913960400191505060405180910390fd5b600a54604080517f06db129800000000000000000000000000000000000000000000000000000000815260048101849052600260248201529051600160a060020a03909216916306db129891604480820192602092909190829003018186803b1580156128d957600080fd5b505afa1580156128ed573d6000803e3d6000fd5b505050506040513d602081101561290357600080fd5b505115156129455760405160e560020a62461bcd02815260040180806020018281038252603381526020018061354a6033913960400191505060405180910390fd5b600a54604080517f327322c8000000000000000000000000000000000000000000000000000000008152600481018490529051600160a060020a039092169163327322c891602480820192602092909190829003018186803b1580156129aa57600080fd5b505afa1580156129be573d6000803e3d6000fd5b505050506040513d60208110156129d457600080fd5b50511515612a165760405160e560020a62461bcd0281526004018080602001828103825260478152602001806132fd6047913960600191505060405180910390fd5b600954604080517f5aebe7680000000000000000000000000000000000000000000000000000000081529051600092600160a060020a031691635aebe768916004808301926020929190829003018186803b158015612a7457600080fd5b505afa158015612a88573d6000803e3d6000fd5b505050506040513d6020811015612a9e57600080fd5b5051600a54600954604080517fc26c12eb0000000000000000000000000000000000000000000000000000000081529051939450600160a060020a0392831693638f354b79938793169163c26c12eb916004808301926020929190829003018186803b158015612b0d57600080fd5b505afa158015612b21573d6000803e3d6000fd5b505050506040513d6020811015612b3757600080fd5b50516040805160e060020a63ffffffff861602815260048101939093526024830191909152516044808301926020929190829003018186803b158015612b7c57600080fd5b505afa158015612b90573d6000803e3d6000fd5b505050506040513d6020811015612ba657600080fd5b505115612cbe57612bb682612f0b565b600854600083815260208190526040902054600160a060020a039182169163a9059cbb9116612bec84600263ffffffff612ebb16565b6040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015612c3e57600080fd5b505af1158015612c52573d6000803e3d6000fd5b505050506040513d6020811015612c6857600080fd5b5050600082815260208181526040918290205482518481529251600160a060020a039091169285927f26adc8b945cdd9b054615b9f327f3945c817f77568b5fc0b7adaecd3fa5c149b92918290030190a3612d94565b600082815260208181526040808320600190810154925290912060020154612ceb9163ffffffff61329e16565b6000838152600160208181526040808420600280820196909655848352932090930154925260040154612d35918391612d299163ffffffff61329e16565b9063ffffffff61329e16565b600083815260016020908152604080832060040193909355818152908290205482518481529251600160a060020a039091169285927f303714bcbac27ef8d90dd5bce88c73b089d8c012ea32c762355112f66320c5fe92918290030190a35b600082815260208190526040808220805473ffffffffffffffffffffffffffffffffffffffff1916815560018101839055600201829055600a5481517f89bb617c000000000000000000000000000000000000000000000000000000008152600481018690529151600160a060020a03909116926389bb617c926024808201939182900301818387803b158015612e2a57600080fd5b505af1158015612e3e573d6000803e3d6000fd5b505050505050565b600090815260016020819052604090912081015460ff1615151490565b60606002805480602002602001604051908101604052809291908181526020018280548015612eb157602002820191906000526020600020905b815481526020019060010190808311612e9d575b5050505050905090565b6000821515612ecc57506000611e40565b828202828482811515612edb57fe5b0414612ee357fe5b9392505050565b6000808211612ef557fe5b60008284811515612f0257fe5b04949350505050565b612f1481611ddb565b1515612f545760405160e560020a62461bcd02815260040180806020018281038252602e815260200180613496602e913960400191505060405180910390fd5b600081815260016020526040812060020154111561300b57600854600082815260016020818152604080842092830154600290930154815160e060020a63a9059cbb028152610100909404600160a060020a0390811660048601526024850191909152905194169363a9059cbb936044808501948390030190829087803b158015612fde57600080fd5b505af1158015612ff2573d6000803e3d6000fd5b505050506040513d602081101561300857600080fd5b50505b60008181526001602052604081206004015411156130af5760085460008281526001602052604080822060049081015482517f42966c68000000000000000000000000000000000000000000000000000000008152918201529051600160a060020a03909316926342966c689260248084019391929182900301818387803b15801561309657600080fd5b505af11580156130aa573d6000803e3d6000fd5b505050505b6000818152600160205260408120546002805491929160001981019081106130d357fe5b90600052602060002001549050806002838154811015156130f057fe5b600091825260209091200155600280549061310f9060001983016132bf565b5060008181526001602081815260408084208690558684528084208301546101009004600160a060020a0316845260039091529091205461314f916132ad565b600084815260016020818152604080842083018054600160a060020a03610100918290048116875260038086528488209890985591540416845293905291902054101561322a57600a5460008481526001602081905260408083209091015481517f335d8080000000000000000000000000000000000000000000000000000000008152600160a060020a0361010090920482166004820152915193169263335d80809260248084019391929182900301818387803b15801561321157600080fd5b505af1158015613225573d6000803e3d6000fd5b505050505b6000838152600160208190526040808320838155918201805474ffffffffffffffffffffffffffffffffffffffffff19169055600282018390556003820183905560049091018290555184917fbf8f571a7fa2903765660aa54ed716098a6b456db5be6fba8e17bffbce48c85e91a2505050565b600082820183811015612ee357fe5b6000828211156132b957fe5b50900390565b815481835581811115612798576000838152602090206127989181019083016120bc91905b808211156132f857600081556001016132e4565b509056fe4572726f723a4d61726b65742e7265736f6c76654368616c6c656e6765202d20506f6c6c7320666f7220746869732063616e646964617465206d75737420626520636c6f7365644572726f723a4d61726b65742e65786974202d2043616e6e6f74206578697420647572696e672061206368616c6c656e67654572726f723a4d61726b65742e646976657374202d204d75737420626520616e20696e766573746f724572726f723a4d61726b65742e6170706c79202d20416c726561647920612063616e6469646174654572726f723a4d61726b65742e776974686472617746726f6d4c697374696e67202d204c697374696e6720646f6573206e6f742065786973744572726f723a4d61726b65742e65786974202d204d757374206265206c697374696e67206f776e65724572726f723a4d61726b65742e7265736f6c76654170706c69636174696f6e202d20416c7265616479206c69737465644572726f723a4d61726b65742e7265736f6c76654368616c6c656e6765202d2053656e646572206d75737420626520636f756e63696c206d656d6265724572726f723a4d61726b65742e72656d6f76654c697374696e67202d204d7573742062652061206c697374696e674572726f723a4d61726b65742e696e76657374202d2043616e6e6f7420696e76657374207768696c65206f776e696e6720616e20616374697665206c697374696e674572726f723a4d61726b65742e65786974202d204d757374206265206c69737465644572726f723a4d61726b65742e6c697374202d20416c7265616479206c69737465644572726f723a4d61726b65742e7265736f6c76654368616c6c656e6765202d204d7573742062652061206368616c6c656e67654572726f723a4d61726b65742e7265736f6c76654170706c69636174696f6e202d20506f6c6c7320666f7220746869732063616e646964617465206d75737420626520636c6f7365644572726f723a4d61726b65742e7265736f6c76654170706c69636174696f6e202d2053656e646572206d75737420626520636f756e63696c206d656d6265724572726f723a4d61726b65742e7265736f6c76654170706c69636174696f6e202d204d75737420626520616e206170706c69636174696f6e4572726f723a4d61726b65742e6465706f736974546f4c697374696e67202d204c697374696e6720646f6573206e6f742065786973744572726f723a4d61726b65742e646976657374202d204d757374206e6f7420626520746865206368616c6c656e67657220696e20616e20616374697665206368616c6c656e67654572726f723a4d61726b65742e696e76657374202d20416d6f756e74206f666665726564206d7573742062652067726561746572207468616e206f7220657175616c20746f2063757272656e7420696e766573746d656e742070726963654572726f723a4d61726b65742e7769746864726177202d20616d6f756e74206d757374206265206c657373207468616e206f7220657175616c20746f20746865206c697374696e6720737570706c794572726f723a4d61726b65742e6368616c6c656e6765202d20416c7265616479206368616c6c656e6765644572726f723a4d61726b65742e6368616c6c656e6765202d204d7573742062652061206c697374696e6720746f206265206368616c6c656e6765644572726f723a4d61726b65742e7769746864726177202d204d757374206265206c697374696e67206f776e6572a165627a7a72305820382bd1c72455f614694a767d47a0bdcda2d764e2f890e334859e7a46b1faff5a0029`

// DeployMarket deploys a new Ethereum contract, binding an instance of Market to it.
func DeployMarket(auth *bind.TransactOpts, backend bind.ContractBackend, marketTokenAddr common.Address, etherTokenAddr common.Address, parameterizerAddr common.Address, votingAddr common.Address) (common.Address, *types.Transaction, *Market, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarketBin), backend, marketTokenAddr, etherTokenAddr, parameterizerAddr, votingAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// Market is an auto generated Go binding around an Ethereum contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(listingHash bytes32) constant returns(address, uint256, uint256)
func (_Market *MarketCaller) GetChallenge(opts *bind.CallOpts, listingHash [32]byte) (common.Address, *big.Int, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Market.contract.Call(opts, out, "getChallenge", listingHash)
	return *ret0, *ret1, *ret2, err
}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(listingHash bytes32) constant returns(address, uint256, uint256)
func (_Market *MarketSession) GetChallenge(listingHash [32]byte) (common.Address, *big.Int, *big.Int, error) {
	return _Market.Contract.GetChallenge(&_Market.CallOpts, listingHash)
}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(listingHash bytes32) constant returns(address, uint256, uint256)
func (_Market *MarketCallerSession) GetChallenge(listingHash [32]byte) (common.Address, *big.Int, *big.Int, error) {
	return _Market.Contract.GetChallenge(&_Market.CallOpts, listingHash)
}

// GetInvested is a free data retrieval call binding the contract method 0xbefc3e2b.
//
// Solidity: function getInvested() constant returns(uint256)
func (_Market *MarketCaller) GetInvested(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getInvested")
	return *ret0, err
}

// GetInvested is a free data retrieval call binding the contract method 0xbefc3e2b.
//
// Solidity: function getInvested() constant returns(uint256)
func (_Market *MarketSession) GetInvested() (*big.Int, error) {
	return _Market.Contract.GetInvested(&_Market.CallOpts)
}

// GetInvested is a free data retrieval call binding the contract method 0xbefc3e2b.
//
// Solidity: function getInvested() constant returns(uint256)
func (_Market *MarketCallerSession) GetInvested() (*big.Int, error) {
	return _Market.Contract.GetInvested(&_Market.CallOpts)
}

// GetInvestment is a free data retrieval call binding the contract method 0x146b58df.
//
// Solidity: function getInvestment(addr address) constant returns(uint256)
func (_Market *MarketCaller) GetInvestment(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getInvestment", addr)
	return *ret0, err
}

// GetInvestment is a free data retrieval call binding the contract method 0x146b58df.
//
// Solidity: function getInvestment(addr address) constant returns(uint256)
func (_Market *MarketSession) GetInvestment(addr common.Address) (*big.Int, error) {
	return _Market.Contract.GetInvestment(&_Market.CallOpts, addr)
}

// GetInvestment is a free data retrieval call binding the contract method 0x146b58df.
//
// Solidity: function getInvestment(addr address) constant returns(uint256)
func (_Market *MarketCallerSession) GetInvestment(addr common.Address) (*big.Int, error) {
	return _Market.Contract.GetInvestment(&_Market.CallOpts, addr)
}

// GetInvestmentPrice is a free data retrieval call binding the contract method 0xa3d6602b.
//
// Solidity: function getInvestmentPrice() constant returns(uint256)
func (_Market *MarketCaller) GetInvestmentPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getInvestmentPrice")
	return *ret0, err
}

// GetInvestmentPrice is a free data retrieval call binding the contract method 0xa3d6602b.
//
// Solidity: function getInvestmentPrice() constant returns(uint256)
func (_Market *MarketSession) GetInvestmentPrice() (*big.Int, error) {
	return _Market.Contract.GetInvestmentPrice(&_Market.CallOpts)
}

// GetInvestmentPrice is a free data retrieval call binding the contract method 0xa3d6602b.
//
// Solidity: function getInvestmentPrice() constant returns(uint256)
func (_Market *MarketCallerSession) GetInvestmentPrice() (*big.Int, error) {
	return _Market.Contract.GetInvestmentPrice(&_Market.CallOpts)
}

// GetListing is a free data retrieval call binding the contract method 0x175c0d16.
//
// Solidity: function getListing(listingHash bytes32) constant returns(bool, address, uint256, bytes32, uint256)
func (_Market *MarketCaller) GetListing(opts *bind.CallOpts, listingHash [32]byte) (bool, common.Address, *big.Int, [32]byte, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new([32]byte)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Market.contract.Call(opts, out, "getListing", listingHash)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetListing is a free data retrieval call binding the contract method 0x175c0d16.
//
// Solidity: function getListing(listingHash bytes32) constant returns(bool, address, uint256, bytes32, uint256)
func (_Market *MarketSession) GetListing(listingHash [32]byte) (bool, common.Address, *big.Int, [32]byte, *big.Int, error) {
	return _Market.Contract.GetListing(&_Market.CallOpts, listingHash)
}

// GetListing is a free data retrieval call binding the contract method 0x175c0d16.
//
// Solidity: function getListing(listingHash bytes32) constant returns(bool, address, uint256, bytes32, uint256)
func (_Market *MarketCallerSession) GetListing(listingHash [32]byte) (bool, common.Address, *big.Int, [32]byte, *big.Int, error) {
	return _Market.Contract.GetListing(&_Market.CallOpts, listingHash)
}

// GetListingHash is a free data retrieval call binding the contract method 0xa2c45b39.
//
// Solidity: function getListingHash(listing string) constant returns(bytes32)
func (_Market *MarketCaller) GetListingHash(opts *bind.CallOpts, listing string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getListingHash", listing)
	return *ret0, err
}

// GetListingHash is a free data retrieval call binding the contract method 0xa2c45b39.
//
// Solidity: function getListingHash(listing string) constant returns(bytes32)
func (_Market *MarketSession) GetListingHash(listing string) ([32]byte, error) {
	return _Market.Contract.GetListingHash(&_Market.CallOpts, listing)
}

// GetListingHash is a free data retrieval call binding the contract method 0xa2c45b39.
//
// Solidity: function getListingHash(listing string) constant returns(bytes32)
func (_Market *MarketCallerSession) GetListingHash(listing string) ([32]byte, error) {
	return _Market.Contract.GetListingHash(&_Market.CallOpts, listing)
}

// GetListings is a free data retrieval call binding the contract method 0xf1b2d6a3.
//
// Solidity: function getListings() constant returns(bytes32[])
func (_Market *MarketCaller) GetListings(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getListings")
	return *ret0, err
}

// GetListings is a free data retrieval call binding the contract method 0xf1b2d6a3.
//
// Solidity: function getListings() constant returns(bytes32[])
func (_Market *MarketSession) GetListings() ([][32]byte, error) {
	return _Market.Contract.GetListings(&_Market.CallOpts)
}

// GetListings is a free data retrieval call binding the contract method 0xf1b2d6a3.
//
// Solidity: function getListings() constant returns(bytes32[])
func (_Market *MarketCallerSession) GetListings() ([][32]byte, error) {
	return _Market.Contract.GetListings(&_Market.CallOpts)
}

// IsInvestor is a free data retrieval call binding the contract method 0xcee2a9cf.
//
// Solidity: function isInvestor(addr address) constant returns(bool)
func (_Market *MarketCaller) IsInvestor(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "isInvestor", addr)
	return *ret0, err
}

// IsInvestor is a free data retrieval call binding the contract method 0xcee2a9cf.
//
// Solidity: function isInvestor(addr address) constant returns(bool)
func (_Market *MarketSession) IsInvestor(addr common.Address) (bool, error) {
	return _Market.Contract.IsInvestor(&_Market.CallOpts, addr)
}

// IsInvestor is a free data retrieval call binding the contract method 0xcee2a9cf.
//
// Solidity: function isInvestor(addr address) constant returns(bool)
func (_Market *MarketCallerSession) IsInvestor(addr common.Address) (bool, error) {
	return _Market.Contract.IsInvestor(&_Market.CallOpts, addr)
}

// IsListed is a free data retrieval call binding the contract method 0xecefbdc6.
//
// Solidity: function isListed(listingHash bytes32) constant returns(bool)
func (_Market *MarketCaller) IsListed(opts *bind.CallOpts, listingHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "isListed", listingHash)
	return *ret0, err
}

// IsListed is a free data retrieval call binding the contract method 0xecefbdc6.
//
// Solidity: function isListed(listingHash bytes32) constant returns(bool)
func (_Market *MarketSession) IsListed(listingHash [32]byte) (bool, error) {
	return _Market.Contract.IsListed(&_Market.CallOpts, listingHash)
}

// IsListed is a free data retrieval call binding the contract method 0xecefbdc6.
//
// Solidity: function isListed(listingHash bytes32) constant returns(bool)
func (_Market *MarketCallerSession) IsListed(listingHash [32]byte) (bool, error) {
	return _Market.Contract.IsListed(&_Market.CallOpts, listingHash)
}

// IsListing is a free data retrieval call binding the contract method 0x6aa0966b.
//
// Solidity: function isListing(listingHash bytes32) constant returns(bool)
func (_Market *MarketCaller) IsListing(opts *bind.CallOpts, listingHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "isListing", listingHash)
	return *ret0, err
}

// IsListing is a free data retrieval call binding the contract method 0x6aa0966b.
//
// Solidity: function isListing(listingHash bytes32) constant returns(bool)
func (_Market *MarketSession) IsListing(listingHash [32]byte) (bool, error) {
	return _Market.Contract.IsListing(&_Market.CallOpts, listingHash)
}

// IsListing is a free data retrieval call binding the contract method 0x6aa0966b.
//
// Solidity: function isListing(listingHash bytes32) constant returns(bool)
func (_Market *MarketCallerSession) IsListing(listingHash [32]byte) (bool, error) {
	return _Market.Contract.IsListing(&_Market.CallOpts, listingHash)
}

// IsListingOwner is a free data retrieval call binding the contract method 0x02a47cd7.
//
// Solidity: function isListingOwner(addr address) constant returns(bool)
func (_Market *MarketCaller) IsListingOwner(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "isListingOwner", addr)
	return *ret0, err
}

// IsListingOwner is a free data retrieval call binding the contract method 0x02a47cd7.
//
// Solidity: function isListingOwner(addr address) constant returns(bool)
func (_Market *MarketSession) IsListingOwner(addr common.Address) (bool, error) {
	return _Market.Contract.IsListingOwner(&_Market.CallOpts, addr)
}

// IsListingOwner is a free data retrieval call binding the contract method 0x02a47cd7.
//
// Solidity: function isListingOwner(addr address) constant returns(bool)
func (_Market *MarketCallerSession) IsListingOwner(addr common.Address) (bool, error) {
	return _Market.Contract.IsListingOwner(&_Market.CallOpts, addr)
}

// Challenge is a paid mutator transaction binding the contract method 0xcffd46dc.
//
// Solidity: function challenge(listingHash bytes32) returns()
func (_Market *MarketTransactor) Challenge(opts *bind.TransactOpts, listingHash [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "challenge", listingHash)
}

// Challenge is a paid mutator transaction binding the contract method 0xcffd46dc.
//
// Solidity: function challenge(listingHash bytes32) returns()
func (_Market *MarketSession) Challenge(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Challenge(&_Market.TransactOpts, listingHash)
}

// Challenge is a paid mutator transaction binding the contract method 0xcffd46dc.
//
// Solidity: function challenge(listingHash bytes32) returns()
func (_Market *MarketTransactorSession) Challenge(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Challenge(&_Market.TransactOpts, listingHash)
}

// DepositToListing is a paid mutator transaction binding the contract method 0x341ff089.
//
// Solidity: function depositToListing(listingHash bytes32, amount uint256) returns()
func (_Market *MarketTransactor) DepositToListing(opts *bind.TransactOpts, listingHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "depositToListing", listingHash, amount)
}

// DepositToListing is a paid mutator transaction binding the contract method 0x341ff089.
//
// Solidity: function depositToListing(listingHash bytes32, amount uint256) returns()
func (_Market *MarketSession) DepositToListing(listingHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.DepositToListing(&_Market.TransactOpts, listingHash, amount)
}

// DepositToListing is a paid mutator transaction binding the contract method 0x341ff089.
//
// Solidity: function depositToListing(listingHash bytes32, amount uint256) returns()
func (_Market *MarketTransactorSession) DepositToListing(listingHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.DepositToListing(&_Market.TransactOpts, listingHash, amount)
}

// Divest is a paid mutator transaction binding the contract method 0x058aace1.
//
// Solidity: function divest() returns()
func (_Market *MarketTransactor) Divest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "divest")
}

// Divest is a paid mutator transaction binding the contract method 0x058aace1.
//
// Solidity: function divest() returns()
func (_Market *MarketSession) Divest() (*types.Transaction, error) {
	return _Market.Contract.Divest(&_Market.TransactOpts)
}

// Divest is a paid mutator transaction binding the contract method 0x058aace1.
//
// Solidity: function divest() returns()
func (_Market *MarketTransactorSession) Divest() (*types.Transaction, error) {
	return _Market.Contract.Divest(&_Market.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0x0ca36263.
//
// Solidity: function exit(listingHash bytes32) returns()
func (_Market *MarketTransactor) Exit(opts *bind.TransactOpts, listingHash [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "exit", listingHash)
}

// Exit is a paid mutator transaction binding the contract method 0x0ca36263.
//
// Solidity: function exit(listingHash bytes32) returns()
func (_Market *MarketSession) Exit(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Exit(&_Market.TransactOpts, listingHash)
}

// Exit is a paid mutator transaction binding the contract method 0x0ca36263.
//
// Solidity: function exit(listingHash bytes32) returns()
func (_Market *MarketTransactorSession) Exit(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Exit(&_Market.TransactOpts, listingHash)
}

// Invest is a paid mutator transaction binding the contract method 0x2afcf480.
//
// Solidity: function invest(offered uint256) returns()
func (_Market *MarketTransactor) Invest(opts *bind.TransactOpts, offered *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "invest", offered)
}

// Invest is a paid mutator transaction binding the contract method 0x2afcf480.
//
// Solidity: function invest(offered uint256) returns()
func (_Market *MarketSession) Invest(offered *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Invest(&_Market.TransactOpts, offered)
}

// Invest is a paid mutator transaction binding the contract method 0x2afcf480.
//
// Solidity: function invest(offered uint256) returns()
func (_Market *MarketTransactorSession) Invest(offered *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Invest(&_Market.TransactOpts, offered)
}

// List is a paid mutator transaction binding the contract method 0x5148cd0f.
//
// Solidity: function list(listing string, dataHash bytes32, amount uint256) returns()
func (_Market *MarketTransactor) List(opts *bind.TransactOpts, listing string, dataHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "list", listing, dataHash, amount)
}

// List is a paid mutator transaction binding the contract method 0x5148cd0f.
//
// Solidity: function list(listing string, dataHash bytes32, amount uint256) returns()
func (_Market *MarketSession) List(listing string, dataHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.List(&_Market.TransactOpts, listing, dataHash, amount)
}

// List is a paid mutator transaction binding the contract method 0x5148cd0f.
//
// Solidity: function list(listing string, dataHash bytes32, amount uint256) returns()
func (_Market *MarketTransactorSession) List(listing string, dataHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.List(&_Market.TransactOpts, listing, dataHash, amount)
}

// ResolveApplication is a paid mutator transaction binding the contract method 0x18e28e1c.
//
// Solidity: function resolveApplication(listingHash bytes32) returns()
func (_Market *MarketTransactor) ResolveApplication(opts *bind.TransactOpts, listingHash [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "resolveApplication", listingHash)
}

// ResolveApplication is a paid mutator transaction binding the contract method 0x18e28e1c.
//
// Solidity: function resolveApplication(listingHash bytes32) returns()
func (_Market *MarketSession) ResolveApplication(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.ResolveApplication(&_Market.TransactOpts, listingHash)
}

// ResolveApplication is a paid mutator transaction binding the contract method 0x18e28e1c.
//
// Solidity: function resolveApplication(listingHash bytes32) returns()
func (_Market *MarketTransactorSession) ResolveApplication(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.ResolveApplication(&_Market.TransactOpts, listingHash)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0xd32c943a.
//
// Solidity: function resolveChallenge(listingHash bytes32) returns()
func (_Market *MarketTransactor) ResolveChallenge(opts *bind.TransactOpts, listingHash [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "resolveChallenge", listingHash)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0xd32c943a.
//
// Solidity: function resolveChallenge(listingHash bytes32) returns()
func (_Market *MarketSession) ResolveChallenge(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.ResolveChallenge(&_Market.TransactOpts, listingHash)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0xd32c943a.
//
// Solidity: function resolveChallenge(listingHash bytes32) returns()
func (_Market *MarketTransactorSession) ResolveChallenge(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.ResolveChallenge(&_Market.TransactOpts, listingHash)
}

// WithdrawFromListing is a paid mutator transaction binding the contract method 0xa7e78675.
//
// Solidity: function withdrawFromListing(listingHash bytes32, amount uint256) returns()
func (_Market *MarketTransactor) WithdrawFromListing(opts *bind.TransactOpts, listingHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "withdrawFromListing", listingHash, amount)
}

// WithdrawFromListing is a paid mutator transaction binding the contract method 0xa7e78675.
//
// Solidity: function withdrawFromListing(listingHash bytes32, amount uint256) returns()
func (_Market *MarketSession) WithdrawFromListing(listingHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.WithdrawFromListing(&_Market.TransactOpts, listingHash, amount)
}

// WithdrawFromListing is a paid mutator transaction binding the contract method 0xa7e78675.
//
// Solidity: function withdrawFromListing(listingHash bytes32, amount uint256) returns()
func (_Market *MarketTransactorSession) WithdrawFromListing(listingHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.WithdrawFromListing(&_Market.TransactOpts, listingHash, amount)
}

// MarketApplicationFailedEventIterator is returned from FilterApplicationFailedEvent and is used to iterate over the raw logs and unpacked data for ApplicationFailedEvent events raised by the Market contract.
type MarketApplicationFailedEventIterator struct {
	Event *MarketApplicationFailedEvent // Event containing the contract specifics and raw log

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
func (it *MarketApplicationFailedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketApplicationFailedEvent)
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
		it.Event = new(MarketApplicationFailedEvent)
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
func (it *MarketApplicationFailedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketApplicationFailedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketApplicationFailedEvent represents a ApplicationFailedEvent event raised by the Market contract.
type MarketApplicationFailedEvent struct {
	ListingHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterApplicationFailedEvent is a free log retrieval operation binding the contract event 0x914b32eae964b80c5a5d1e61c1b5617ffb64c0c0b3bf61ab0e0208ebe5c0db35.
//
// Solidity: e ApplicationFailedEvent(listingHash indexed bytes32)
func (_Market *MarketFilterer) FilterApplicationFailedEvent(opts *bind.FilterOpts, listingHash [][32]byte) (*MarketApplicationFailedEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ApplicationFailedEvent", listingHashRule)
	if err != nil {
		return nil, err
	}
	return &MarketApplicationFailedEventIterator{contract: _Market.contract, event: "ApplicationFailedEvent", logs: logs, sub: sub}, nil
}

// WatchApplicationFailedEvent is a free log subscription operation binding the contract event 0x914b32eae964b80c5a5d1e61c1b5617ffb64c0c0b3bf61ab0e0208ebe5c0db35.
//
// Solidity: e ApplicationFailedEvent(listingHash indexed bytes32)
func (_Market *MarketFilterer) WatchApplicationFailedEvent(opts *bind.WatchOpts, sink chan<- *MarketApplicationFailedEvent, listingHash [][32]byte) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ApplicationFailedEvent", listingHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketApplicationFailedEvent)
				if err := _Market.contract.UnpackLog(event, "ApplicationFailedEvent", log); err != nil {
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

// MarketAppliedEventIterator is returned from FilterAppliedEvent and is used to iterate over the raw logs and unpacked data for AppliedEvent events raised by the Market contract.
type MarketAppliedEventIterator struct {
	Event *MarketAppliedEvent // Event containing the contract specifics and raw log

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
func (it *MarketAppliedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketAppliedEvent)
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
		it.Event = new(MarketAppliedEvent)
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
func (it *MarketAppliedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketAppliedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketAppliedEvent represents a AppliedEvent event raised by the Market contract.
type MarketAppliedEvent struct {
	ListingHash [32]byte
	Applicant   common.Address
	DataHash    [32]byte
	VoteBy      *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAppliedEvent is a free log retrieval operation binding the contract event 0x06198c59f70529156e65d8211777fdd8fdd009977545a9e8c44050f7ff2cb133.
//
// Solidity: e AppliedEvent(listingHash indexed bytes32, applicant indexed address, dataHash indexed bytes32, voteBy uint256, amount uint256)
func (_Market *MarketFilterer) FilterAppliedEvent(opts *bind.FilterOpts, listingHash [][32]byte, applicant []common.Address, dataHash [][32]byte) (*MarketAppliedEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var applicantRule []interface{}
	for _, applicantItem := range applicant {
		applicantRule = append(applicantRule, applicantItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "AppliedEvent", listingHashRule, applicantRule, dataHashRule)
	if err != nil {
		return nil, err
	}
	return &MarketAppliedEventIterator{contract: _Market.contract, event: "AppliedEvent", logs: logs, sub: sub}, nil
}

// WatchAppliedEvent is a free log subscription operation binding the contract event 0x06198c59f70529156e65d8211777fdd8fdd009977545a9e8c44050f7ff2cb133.
//
// Solidity: e AppliedEvent(listingHash indexed bytes32, applicant indexed address, dataHash indexed bytes32, voteBy uint256, amount uint256)
func (_Market *MarketFilterer) WatchAppliedEvent(opts *bind.WatchOpts, sink chan<- *MarketAppliedEvent, listingHash [][32]byte, applicant []common.Address, dataHash [][32]byte) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var applicantRule []interface{}
	for _, applicantItem := range applicant {
		applicantRule = append(applicantRule, applicantItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "AppliedEvent", listingHashRule, applicantRule, dataHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketAppliedEvent)
				if err := _Market.contract.UnpackLog(event, "AppliedEvent", log); err != nil {
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

// MarketChallengeFailedEventIterator is returned from FilterChallengeFailedEvent and is used to iterate over the raw logs and unpacked data for ChallengeFailedEvent events raised by the Market contract.
type MarketChallengeFailedEventIterator struct {
	Event *MarketChallengeFailedEvent // Event containing the contract specifics and raw log

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
func (it *MarketChallengeFailedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketChallengeFailedEvent)
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
		it.Event = new(MarketChallengeFailedEvent)
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
func (it *MarketChallengeFailedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketChallengeFailedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketChallengeFailedEvent represents a ChallengeFailedEvent event raised by the Market contract.
type MarketChallengeFailedEvent struct {
	ListingHash [32]byte
	Challenger  common.Address
	Reward      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChallengeFailedEvent is a free log retrieval operation binding the contract event 0x303714bcbac27ef8d90dd5bce88c73b089d8c012ea32c762355112f66320c5fe.
//
// Solidity: e ChallengeFailedEvent(listingHash indexed bytes32, challenger indexed address, reward uint256)
func (_Market *MarketFilterer) FilterChallengeFailedEvent(opts *bind.FilterOpts, listingHash [][32]byte, challenger []common.Address) (*MarketChallengeFailedEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ChallengeFailedEvent", listingHashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &MarketChallengeFailedEventIterator{contract: _Market.contract, event: "ChallengeFailedEvent", logs: logs, sub: sub}, nil
}

// WatchChallengeFailedEvent is a free log subscription operation binding the contract event 0x303714bcbac27ef8d90dd5bce88c73b089d8c012ea32c762355112f66320c5fe.
//
// Solidity: e ChallengeFailedEvent(listingHash indexed bytes32, challenger indexed address, reward uint256)
func (_Market *MarketFilterer) WatchChallengeFailedEvent(opts *bind.WatchOpts, sink chan<- *MarketChallengeFailedEvent, listingHash [][32]byte, challenger []common.Address) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ChallengeFailedEvent", listingHashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketChallengeFailedEvent)
				if err := _Market.contract.UnpackLog(event, "ChallengeFailedEvent", log); err != nil {
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

// MarketChallengeSucceededEventIterator is returned from FilterChallengeSucceededEvent and is used to iterate over the raw logs and unpacked data for ChallengeSucceededEvent events raised by the Market contract.
type MarketChallengeSucceededEventIterator struct {
	Event *MarketChallengeSucceededEvent // Event containing the contract specifics and raw log

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
func (it *MarketChallengeSucceededEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketChallengeSucceededEvent)
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
		it.Event = new(MarketChallengeSucceededEvent)
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
func (it *MarketChallengeSucceededEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketChallengeSucceededEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketChallengeSucceededEvent represents a ChallengeSucceededEvent event raised by the Market contract.
type MarketChallengeSucceededEvent struct {
	ListingHash [32]byte
	Challenger  common.Address
	Reward      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChallengeSucceededEvent is a free log retrieval operation binding the contract event 0x26adc8b945cdd9b054615b9f327f3945c817f77568b5fc0b7adaecd3fa5c149b.
//
// Solidity: e ChallengeSucceededEvent(listingHash indexed bytes32, challenger indexed address, reward uint256)
func (_Market *MarketFilterer) FilterChallengeSucceededEvent(opts *bind.FilterOpts, listingHash [][32]byte, challenger []common.Address) (*MarketChallengeSucceededEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ChallengeSucceededEvent", listingHashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &MarketChallengeSucceededEventIterator{contract: _Market.contract, event: "ChallengeSucceededEvent", logs: logs, sub: sub}, nil
}

// WatchChallengeSucceededEvent is a free log subscription operation binding the contract event 0x26adc8b945cdd9b054615b9f327f3945c817f77568b5fc0b7adaecd3fa5c149b.
//
// Solidity: e ChallengeSucceededEvent(listingHash indexed bytes32, challenger indexed address, reward uint256)
func (_Market *MarketFilterer) WatchChallengeSucceededEvent(opts *bind.WatchOpts, sink chan<- *MarketChallengeSucceededEvent, listingHash [][32]byte, challenger []common.Address) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ChallengeSucceededEvent", listingHashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketChallengeSucceededEvent)
				if err := _Market.contract.UnpackLog(event, "ChallengeSucceededEvent", log); err != nil {
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

// MarketChallengedEventIterator is returned from FilterChallengedEvent and is used to iterate over the raw logs and unpacked data for ChallengedEvent events raised by the Market contract.
type MarketChallengedEventIterator struct {
	Event *MarketChallengedEvent // Event containing the contract specifics and raw log

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
func (it *MarketChallengedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketChallengedEvent)
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
		it.Event = new(MarketChallengedEvent)
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
func (it *MarketChallengedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketChallengedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketChallengedEvent represents a ChallengedEvent event raised by the Market contract.
type MarketChallengedEvent struct {
	ListingHash [32]byte
	Challenger  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChallengedEvent is a free log retrieval operation binding the contract event 0x4dd89117b2f54a955a15b7ca70efe9d9cda23696ebd2a8290ce6bb42b1594dba.
//
// Solidity: e ChallengedEvent(listingHash indexed bytes32, challenger indexed address)
func (_Market *MarketFilterer) FilterChallengedEvent(opts *bind.FilterOpts, listingHash [][32]byte, challenger []common.Address) (*MarketChallengedEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ChallengedEvent", listingHashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &MarketChallengedEventIterator{contract: _Market.contract, event: "ChallengedEvent", logs: logs, sub: sub}, nil
}

// WatchChallengedEvent is a free log subscription operation binding the contract event 0x4dd89117b2f54a955a15b7ca70efe9d9cda23696ebd2a8290ce6bb42b1594dba.
//
// Solidity: e ChallengedEvent(listingHash indexed bytes32, challenger indexed address)
func (_Market *MarketFilterer) WatchChallengedEvent(opts *bind.WatchOpts, sink chan<- *MarketChallengedEvent, listingHash [][32]byte, challenger []common.Address) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ChallengedEvent", listingHashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketChallengedEvent)
				if err := _Market.contract.UnpackLog(event, "ChallengedEvent", log); err != nil {
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

// MarketDivestedEventIterator is returned from FilterDivestedEvent and is used to iterate over the raw logs and unpacked data for DivestedEvent events raised by the Market contract.
type MarketDivestedEventIterator struct {
	Event *MarketDivestedEvent // Event containing the contract specifics and raw log

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
func (it *MarketDivestedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketDivestedEvent)
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
		it.Event = new(MarketDivestedEvent)
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
func (it *MarketDivestedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketDivestedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketDivestedEvent represents a DivestedEvent event raised by the Market contract.
type MarketDivestedEvent struct {
	Investor    common.Address
	Transferred *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDivestedEvent is a free log retrieval operation binding the contract event 0x31b5a476fc1c356476ae05b02dcec8b3c6aaab43b21d4d270089270be10da644.
//
// Solidity: e DivestedEvent(investor indexed address, transferred uint256)
func (_Market *MarketFilterer) FilterDivestedEvent(opts *bind.FilterOpts, investor []common.Address) (*MarketDivestedEventIterator, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "DivestedEvent", investorRule)
	if err != nil {
		return nil, err
	}
	return &MarketDivestedEventIterator{contract: _Market.contract, event: "DivestedEvent", logs: logs, sub: sub}, nil
}

// WatchDivestedEvent is a free log subscription operation binding the contract event 0x31b5a476fc1c356476ae05b02dcec8b3c6aaab43b21d4d270089270be10da644.
//
// Solidity: e DivestedEvent(investor indexed address, transferred uint256)
func (_Market *MarketFilterer) WatchDivestedEvent(opts *bind.WatchOpts, sink chan<- *MarketDivestedEvent, investor []common.Address) (event.Subscription, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "DivestedEvent", investorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketDivestedEvent)
				if err := _Market.contract.UnpackLog(event, "DivestedEvent", log); err != nil {
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

// MarketInvestedEventIterator is returned from FilterInvestedEvent and is used to iterate over the raw logs and unpacked data for InvestedEvent events raised by the Market contract.
type MarketInvestedEventIterator struct {
	Event *MarketInvestedEvent // Event containing the contract specifics and raw log

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
func (it *MarketInvestedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketInvestedEvent)
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
		it.Event = new(MarketInvestedEvent)
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
func (it *MarketInvestedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketInvestedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketInvestedEvent represents a InvestedEvent event raised by the Market contract.
type MarketInvestedEvent struct {
	Investor common.Address
	Offered  *big.Int
	Minted   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestedEvent is a free log retrieval operation binding the contract event 0x1be1f5b0e87aa0e79eed7f75d2113999609d67094a95ff7f8e6c5625a35d9a39.
//
// Solidity: e InvestedEvent(investor indexed address, offered uint256, minted uint256)
func (_Market *MarketFilterer) FilterInvestedEvent(opts *bind.FilterOpts, investor []common.Address) (*MarketInvestedEventIterator, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "InvestedEvent", investorRule)
	if err != nil {
		return nil, err
	}
	return &MarketInvestedEventIterator{contract: _Market.contract, event: "InvestedEvent", logs: logs, sub: sub}, nil
}

// WatchInvestedEvent is a free log subscription operation binding the contract event 0x1be1f5b0e87aa0e79eed7f75d2113999609d67094a95ff7f8e6c5625a35d9a39.
//
// Solidity: e InvestedEvent(investor indexed address, offered uint256, minted uint256)
func (_Market *MarketFilterer) WatchInvestedEvent(opts *bind.WatchOpts, sink chan<- *MarketInvestedEvent, investor []common.Address) (event.Subscription, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "InvestedEvent", investorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketInvestedEvent)
				if err := _Market.contract.UnpackLog(event, "InvestedEvent", log); err != nil {
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

// MarketListedEventIterator is returned from FilterListedEvent and is used to iterate over the raw logs and unpacked data for ListedEvent events raised by the Market contract.
type MarketListedEventIterator struct {
	Event *MarketListedEvent // Event containing the contract specifics and raw log

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
func (it *MarketListedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketListedEvent)
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
		it.Event = new(MarketListedEvent)
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
func (it *MarketListedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketListedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketListedEvent represents a ListedEvent event raised by the Market contract.
type MarketListedEvent struct {
	ListingHash [32]byte
	Owner       common.Address
	Reward      *big.Int
	Supply      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterListedEvent is a free log retrieval operation binding the contract event 0xb08f341fa4f58287c76e79e69c47e07939983551f847f728f4dc9e06a49e308d.
//
// Solidity: e ListedEvent(listingHash indexed bytes32, owner indexed address, reward uint256, supply uint256)
func (_Market *MarketFilterer) FilterListedEvent(opts *bind.FilterOpts, listingHash [][32]byte, owner []common.Address) (*MarketListedEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ListedEvent", listingHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MarketListedEventIterator{contract: _Market.contract, event: "ListedEvent", logs: logs, sub: sub}, nil
}

// WatchListedEvent is a free log subscription operation binding the contract event 0xb08f341fa4f58287c76e79e69c47e07939983551f847f728f4dc9e06a49e308d.
//
// Solidity: e ListedEvent(listingHash indexed bytes32, owner indexed address, reward uint256, supply uint256)
func (_Market *MarketFilterer) WatchListedEvent(opts *bind.WatchOpts, sink chan<- *MarketListedEvent, listingHash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ListedEvent", listingHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketListedEvent)
				if err := _Market.contract.UnpackLog(event, "ListedEvent", log); err != nil {
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

// MarketListingDeletedEventIterator is returned from FilterListingDeletedEvent and is used to iterate over the raw logs and unpacked data for ListingDeletedEvent events raised by the Market contract.
type MarketListingDeletedEventIterator struct {
	Event *MarketListingDeletedEvent // Event containing the contract specifics and raw log

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
func (it *MarketListingDeletedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketListingDeletedEvent)
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
		it.Event = new(MarketListingDeletedEvent)
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
func (it *MarketListingDeletedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketListingDeletedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketListingDeletedEvent represents a ListingDeletedEvent event raised by the Market contract.
type MarketListingDeletedEvent struct {
	ListingHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterListingDeletedEvent is a free log retrieval operation binding the contract event 0xbf8f571a7fa2903765660aa54ed716098a6b456db5be6fba8e17bffbce48c85e.
//
// Solidity: e ListingDeletedEvent(listingHash indexed bytes32)
func (_Market *MarketFilterer) FilterListingDeletedEvent(opts *bind.FilterOpts, listingHash [][32]byte) (*MarketListingDeletedEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ListingDeletedEvent", listingHashRule)
	if err != nil {
		return nil, err
	}
	return &MarketListingDeletedEventIterator{contract: _Market.contract, event: "ListingDeletedEvent", logs: logs, sub: sub}, nil
}

// WatchListingDeletedEvent is a free log subscription operation binding the contract event 0xbf8f571a7fa2903765660aa54ed716098a6b456db5be6fba8e17bffbce48c85e.
//
// Solidity: e ListingDeletedEvent(listingHash indexed bytes32)
func (_Market *MarketFilterer) WatchListingDeletedEvent(opts *bind.WatchOpts, sink chan<- *MarketListingDeletedEvent, listingHash [][32]byte) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ListingDeletedEvent", listingHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketListingDeletedEvent)
				if err := _Market.contract.UnpackLog(event, "ListingDeletedEvent", log); err != nil {
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

// MarketListingDepositEventIterator is returned from FilterListingDepositEvent and is used to iterate over the raw logs and unpacked data for ListingDepositEvent events raised by the Market contract.
type MarketListingDepositEventIterator struct {
	Event *MarketListingDepositEvent // Event containing the contract specifics and raw log

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
func (it *MarketListingDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketListingDepositEvent)
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
		it.Event = new(MarketListingDepositEvent)
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
func (it *MarketListingDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketListingDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketListingDepositEvent represents a ListingDepositEvent event raised by the Market contract.
type MarketListingDepositEvent struct {
	ListingHash [32]byte
	Owner       common.Address
	Added       *big.Int
	Supply      *big.Int
	Rewards     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterListingDepositEvent is a free log retrieval operation binding the contract event 0x7f1c6a478f1fbf07f8b135e162e7322085eb5371d2b77163e27dc39d4154d136.
//
// Solidity: e ListingDepositEvent(listingHash indexed bytes32, owner indexed address, added uint256, supply uint256, rewards uint256)
func (_Market *MarketFilterer) FilterListingDepositEvent(opts *bind.FilterOpts, listingHash [][32]byte, owner []common.Address) (*MarketListingDepositEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ListingDepositEvent", listingHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MarketListingDepositEventIterator{contract: _Market.contract, event: "ListingDepositEvent", logs: logs, sub: sub}, nil
}

// WatchListingDepositEvent is a free log subscription operation binding the contract event 0x7f1c6a478f1fbf07f8b135e162e7322085eb5371d2b77163e27dc39d4154d136.
//
// Solidity: e ListingDepositEvent(listingHash indexed bytes32, owner indexed address, added uint256, supply uint256, rewards uint256)
func (_Market *MarketFilterer) WatchListingDepositEvent(opts *bind.WatchOpts, sink chan<- *MarketListingDepositEvent, listingHash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ListingDepositEvent", listingHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketListingDepositEvent)
				if err := _Market.contract.UnpackLog(event, "ListingDepositEvent", log); err != nil {
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

// MarketListingWithdrawEventIterator is returned from FilterListingWithdrawEvent and is used to iterate over the raw logs and unpacked data for ListingWithdrawEvent events raised by the Market contract.
type MarketListingWithdrawEventIterator struct {
	Event *MarketListingWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *MarketListingWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketListingWithdrawEvent)
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
		it.Event = new(MarketListingWithdrawEvent)
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
func (it *MarketListingWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketListingWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketListingWithdrawEvent represents a ListingWithdrawEvent event raised by the Market contract.
type MarketListingWithdrawEvent struct {
	ListingHash [32]byte
	Owner       common.Address
	Added       *big.Int
	Supply      *big.Int
	Rewards     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterListingWithdrawEvent is a free log retrieval operation binding the contract event 0x261e65ee6edcc0072cc284d3d0665cbc0c581ab1fc708222b6bc3d20c452750d.
//
// Solidity: e ListingWithdrawEvent(listingHash indexed bytes32, owner indexed address, added uint256, supply uint256, rewards uint256)
func (_Market *MarketFilterer) FilterListingWithdrawEvent(opts *bind.FilterOpts, listingHash [][32]byte, owner []common.Address) (*MarketListingWithdrawEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ListingWithdrawEvent", listingHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MarketListingWithdrawEventIterator{contract: _Market.contract, event: "ListingWithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchListingWithdrawEvent is a free log subscription operation binding the contract event 0x261e65ee6edcc0072cc284d3d0665cbc0c581ab1fc708222b6bc3d20c452750d.
//
// Solidity: e ListingWithdrawEvent(listingHash indexed bytes32, owner indexed address, added uint256, supply uint256, rewards uint256)
func (_Market *MarketFilterer) WatchListingWithdrawEvent(opts *bind.WatchOpts, sink chan<- *MarketListingWithdrawEvent, listingHash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ListingWithdrawEvent", listingHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketListingWithdrawEvent)
				if err := _Market.contract.UnpackLog(event, "ListingWithdrawEvent", log); err != nil {
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

// MarketTokenABI is the input ABI used to generate the binding from.
const MarketTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stopMinting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"market\",\"type\":\"address\"}],\"name\":\"setPrivilegedContracts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"burnAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPrivilegedAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintingStopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialAccount\",\"type\":\"address\"},{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ApprovalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurnEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MintEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintStoppedEvent\",\"type\":\"event\"}]"

// MarketTokenBin is the compiled bytecode used for deploying new contracts.
const MarketTokenBin = `0x60806040526001805460ff1916905534801561001a57600080fd5b506040516040806111198339810180604052604081101561003a57600080fd5b50805160209182015160018054610100330261010060a860020a0319909116179055600160a060020a0390911660009081526003909252604082208190559055611090806100896000396000f3fe608060405234801561001057600080fd5b5060043610610112576000357c010000000000000000000000000000000000000000000000000000000090048063730e6e84116100b4578063b6896e5f11610083578063b6896e5f146102b5578063d73dd623146102d9578063dd62ed3e14610305578063f339292f1461033357610112565b8063730e6e84146102205780637e9d2ac114610246578063a0712d681461026c578063a9059cbb1461028957610112565b80633e3e0b12116100f05780633e3e0b12146101a757806342966c68146101b157806366188463146101ce57806370a08231146101fa57610112565b8063095ea7b31461011757806318160ddd1461015757806323b872dd14610171575b600080fd5b6101436004803603604081101561012d57600080fd5b50600160a060020a03813516906020013561033b565b604080519115158252519081900360200190f35b61015f6103a1565b60408051918252519081900360200190f35b6101436004803603606081101561018757600080fd5b50600160a060020a038135811691602081013590911690604001356103a7565b6101af6105bf565b005b6101af600480360360208110156101c757600080fd5b503561068b565b6101af600480360360408110156101e457600080fd5b50600160a060020a0381351690602001356107a8565b61015f6004803603602081101561021057600080fd5b5035600160a060020a03166108b5565b6101af6004803603602081101561023657600080fd5b5035600160a060020a03166108d0565b6101af6004803603602081101561025c57600080fd5b5035600160a060020a031661099b565b6101af6004803603602081101561028257600080fd5b5035610ac8565b6101436004803603604081101561029f57600080fd5b50600160a060020a038135169060200135610bde565b6102bd610d2b565b60408051600160a060020a039092168252519081900360200190f35b6101af600480360360408110156102ef57600080fd5b50600160a060020a038135169060200135610d3a565b61015f6004803603604081101561031b57600080fd5b50600160a060020a0381358116916020013516610d6e565b610143610d99565b336000818152600460209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e928290030190a350600192915050565b60005490565b6000600160a060020a03831615156103f35760405160e560020a62461bcd02815260040180806020018281038252603f815260200180611026603f913960400191505060405180910390fd5b600160a060020a03841660009081526003602052604090205482111561044d5760405160e560020a62461bcd028152600401808060200182810382526041815260200180610ecd6041913960600191505060405180910390fd5b600160a060020a03841660009081526004602090815260408083203384529091529020548211156104b25760405160e560020a62461bcd02815260040180806020018281038252603e815260200180610dcb603e913960400191505060405180910390fd5b600160a060020a0384166000908152600360205260409020546104db908363ffffffff610da216565b600160a060020a038086166000908152600360205260408082209390935590851681522054610510908363ffffffff610db416565b600160a060020a038085166000908152600360209081526040808320949094559187168152600482528281203382529091522054610554908363ffffffff610da216565b600160a060020a03808616600081815260046020908152604080832033845282529182902094909455805186815290519287169391927feaf1c4b3ce0f4f62a2bae7eb3e68225c75f7e6ff4422073b7437b9a78d25f170929181900390910190a35060019392505050565b600254600160a060020a0316331461060b5760405160e560020a62461bcd028152600401808060200182810382526046815260200180610e426046913960600191505060405180910390fd5b6001805460ff16151514156106545760405160e560020a62461bcd028152600401808060200182810382526034815260200180610f0e6034913960400191505060405180910390fd5b6001805460ff1916811790556040517f6ae65cb1ae147362b8906c28687a17e1712ca926bea50710ba279849cec3347190600090a1565b600254600160a060020a031633146106d75760405160e560020a62461bcd028152600401808060200182810382526046815260200180610e426046913960600191505060405180910390fd5b336000908152600360205260409020548111156107285760405160e560020a62461bcd028152600401808060200182810382526039815260200180610e096039913960400191505060405180910390fd5b33600090815260036020526040902054610748908263ffffffff610da216565b336000908152600360205260408120919091555461076c908263ffffffff610da216565b60005560408051828152905133917f512586160ebd4dc6945ba9ec5d21a1f723f26f3c7aa36cdffb6818d4e7b88030919081900360200190a250565b336000908152600460209081526040808320600160a060020a03861684529091529020548111156107fc57336000908152600460209081526040808320600160a060020a0386168452909152812055610855565b336000908152600460209081526040808320600160a060020a0386168452909152902054610830908263ffffffff610da216565b336000908152600460209081526040808320600160a060020a03871684529091529020555b336000818152600460209081526040808320600160a060020a0387168085529083529281902054815190815290519293927f08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e929181900390910190a35050565b600160a060020a031660009081526003602052604090205490565b6001546101009004600160a060020a031633146109215760405160e560020a62461bcd028152600401808060200182810382526030815260200180610f7b6030913960400191505060405180910390fd5b600254600160a060020a03161561096c5760405160e560020a62461bcd028152600401808060200182810382526045815260200180610fab6045913960600191505060405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a031633146109e75760405160e560020a62461bcd028152600401808060200182810382526046815260200180610e426046913960600191505060405180910390fd5b600160a060020a0381161515610a315760405160e560020a62461bcd028152600401808060200182810382526036815260200180610ff06036913960400191505060405180910390fd5b600160a060020a0381166000908152600360205260408120549054610a5c908263ffffffff610da216565b6000908155600160a060020a038316808252600360209081526040808420849055600482528084203385528252808420939093558251848152925191927f512586160ebd4dc6945ba9ec5d21a1f723f26f3c7aa36cdffb6818d4e7b88030929081900390910190a25050565b600254600160a060020a03163314610b145760405160e560020a62461bcd028152600401808060200182810382526046815260200180610e426046913960600191505060405180910390fd5b6001805460ff1615151415610b5d5760405160e560020a62461bcd028152600401808060200182810382526034815260200180610f0e6034913960400191505060405180910390fd5b600054610b70908263ffffffff610db416565b600090815533815260036020526040902054610b92908263ffffffff610db416565b33600081815260036020908152604091829020939093558051848152905191927f3fffaa5804a26fcec0d70b1d0fb0a2d0031df3a5f9c8af2127c2f4360e97b46392918290030190a250565b6000600160a060020a0383161515610c2a5760405160e560020a62461bcd028152600401808060200182810382526039815260200180610f426039913960400191505060405180910390fd5b33600090815260036020526040902054821115610c7b5760405160e560020a62461bcd028152600401808060200182810382526045815260200180610e886045913960600191505060405180910390fd5b33600090815260036020526040902054610c9b908363ffffffff610da216565b3360009081526003602052604080822092909255600160a060020a03851681522054610ccd908363ffffffff610db416565b600160a060020a0384166000818152600360209081526040918290209390935580518581529051919233927feaf1c4b3ce0f4f62a2bae7eb3e68225c75f7e6ff4422073b7437b9a78d25f1709281900390910190a350600192915050565b600254600160a060020a031690565b336000908152600460209081526040808320600160a060020a0386168452909152902054610830908263ffffffff610db416565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205490565b60015460ff1690565b600082821115610dae57fe5b50900390565b600082820183811015610dc357fe5b939250505056fe4572726f722e4d61726b6574546f6b656e2e7472616e7366657246726f6d202d20416d6f756e74206578636565647320616c6c6f77656420616d6f756e744572726f723a4d61726b6574546f6b656e2e6275726e202d20416d6f756e74206578636565647320617661696c61626c652062616c616e63654572726f723a4d61726b6574546f6b656e2e68617350726976696c656765202d2043616c6c6572206d75737420626520612070726976696c656765642020636f6e74726163744572726f723a4d61726b6574546f6b656e2e7472616e73666572202d20416d6f756e742065786365656473207468652062616c616e6365206f66206d73672e73656e6465724572726f723a4d61726b6574546f6b656e2e7472616e7366657246726f6d202d20416d6f756e74206578636565647320617661696c61626c652062616c616e63654572726f723a4d61726b6574546f6b656e2e63616e4d696e74202d204d696e74696e6720686173206265656e2073746f707065644572726f723a4d61726b6574546f6b656e2e7472616e73666572202d20416e2061646472657373206d757374206265207370656369666965644572726f723a4d61726b6574546f6b656e2e69734f776e6572202d2043616c6c6572206d757374206265206f776e65724572726f723a4d61726b6574546f6b656e2e73657450726976696c65676564436f6e747261637473202d204d61726b6574206164647265737320616c7265616479207365744572726f723a4d61726b6574546f6b656e2e6275726e46726f6d202d2061646472657373206d757374206265207370656369666965644572726f723a4d61726b6574546f6b656e2e7472616e7366657246726f6d202d2027746f272061646472657373206d75737420626520737065636966696564a165627a7a723058201598338df2915f348023e95e7f546868bcd3cf5686c2da025bbddac9dc75dc940029`

// DeployMarketToken deploys a new Ethereum contract, binding an instance of MarketToken to it.
func DeployMarketToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialAccount common.Address, initialBalance *big.Int) (common.Address, *types.Transaction, *MarketToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarketTokenBin), backend, initialAccount, initialBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketToken{MarketTokenCaller: MarketTokenCaller{contract: contract}, MarketTokenTransactor: MarketTokenTransactor{contract: contract}, MarketTokenFilterer: MarketTokenFilterer{contract: contract}}, nil
}

// MarketToken is an auto generated Go binding around an Ethereum contract.
type MarketToken struct {
	MarketTokenCaller     // Read-only binding to the contract
	MarketTokenTransactor // Write-only binding to the contract
	MarketTokenFilterer   // Log filterer for contract events
}

// MarketTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketTokenSession struct {
	Contract     *MarketToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketTokenCallerSession struct {
	Contract *MarketTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTokenTransactorSession struct {
	Contract     *MarketTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketTokenRaw struct {
	Contract *MarketToken // Generic contract binding to access the raw methods on
}

// MarketTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketTokenCallerRaw struct {
	Contract *MarketTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTokenTransactorRaw struct {
	Contract *MarketTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketToken creates a new instance of MarketToken, bound to a specific deployed contract.
func NewMarketToken(address common.Address, backend bind.ContractBackend) (*MarketToken, error) {
	contract, err := bindMarketToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketToken{MarketTokenCaller: MarketTokenCaller{contract: contract}, MarketTokenTransactor: MarketTokenTransactor{contract: contract}, MarketTokenFilterer: MarketTokenFilterer{contract: contract}}, nil
}

// NewMarketTokenCaller creates a new read-only instance of MarketToken, bound to a specific deployed contract.
func NewMarketTokenCaller(address common.Address, caller bind.ContractCaller) (*MarketTokenCaller, error) {
	contract, err := bindMarketToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTokenCaller{contract: contract}, nil
}

// NewMarketTokenTransactor creates a new write-only instance of MarketToken, bound to a specific deployed contract.
func NewMarketTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTokenTransactor, error) {
	contract, err := bindMarketToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTokenTransactor{contract: contract}, nil
}

// NewMarketTokenFilterer creates a new log filterer instance of MarketToken, bound to a specific deployed contract.
func NewMarketTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketTokenFilterer, error) {
	contract, err := bindMarketToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketTokenFilterer{contract: contract}, nil
}

// bindMarketToken binds a generic wrapper to an already deployed contract.
func bindMarketToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketToken *MarketTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MarketToken.Contract.MarketTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketToken *MarketTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketToken.Contract.MarketTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketToken *MarketTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketToken.Contract.MarketTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketToken *MarketTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MarketToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketToken *MarketTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketToken *MarketTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(holder address, spender address) constant returns(uint256)
func (_MarketToken *MarketTokenCaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "allowance", holder, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(holder address, spender address) constant returns(uint256)
func (_MarketToken *MarketTokenSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _MarketToken.Contract.Allowance(&_MarketToken.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(holder address, spender address) constant returns(uint256)
func (_MarketToken *MarketTokenCallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _MarketToken.Contract.Allowance(&_MarketToken.CallOpts, holder, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(holder address) constant returns(uint256)
func (_MarketToken *MarketTokenCaller) BalanceOf(opts *bind.CallOpts, holder common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "balanceOf", holder)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(holder address) constant returns(uint256)
func (_MarketToken *MarketTokenSession) BalanceOf(holder common.Address) (*big.Int, error) {
	return _MarketToken.Contract.BalanceOf(&_MarketToken.CallOpts, holder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(holder address) constant returns(uint256)
func (_MarketToken *MarketTokenCallerSession) BalanceOf(holder common.Address) (*big.Int, error) {
	return _MarketToken.Contract.BalanceOf(&_MarketToken.CallOpts, holder)
}

// GetPrivilegedAddresses is a free data retrieval call binding the contract method 0xb6896e5f.
//
// Solidity: function getPrivilegedAddresses() constant returns(address)
func (_MarketToken *MarketTokenCaller) GetPrivilegedAddresses(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "getPrivilegedAddresses")
	return *ret0, err
}

// GetPrivilegedAddresses is a free data retrieval call binding the contract method 0xb6896e5f.
//
// Solidity: function getPrivilegedAddresses() constant returns(address)
func (_MarketToken *MarketTokenSession) GetPrivilegedAddresses() (common.Address, error) {
	return _MarketToken.Contract.GetPrivilegedAddresses(&_MarketToken.CallOpts)
}

// GetPrivilegedAddresses is a free data retrieval call binding the contract method 0xb6896e5f.
//
// Solidity: function getPrivilegedAddresses() constant returns(address)
func (_MarketToken *MarketTokenCallerSession) GetPrivilegedAddresses() (common.Address, error) {
	return _MarketToken.Contract.GetPrivilegedAddresses(&_MarketToken.CallOpts)
}

// MintingStopped is a free data retrieval call binding the contract method 0xf339292f.
//
// Solidity: function mintingStopped() constant returns(bool)
func (_MarketToken *MarketTokenCaller) MintingStopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "mintingStopped")
	return *ret0, err
}

// MintingStopped is a free data retrieval call binding the contract method 0xf339292f.
//
// Solidity: function mintingStopped() constant returns(bool)
func (_MarketToken *MarketTokenSession) MintingStopped() (bool, error) {
	return _MarketToken.Contract.MintingStopped(&_MarketToken.CallOpts)
}

// MintingStopped is a free data retrieval call binding the contract method 0xf339292f.
//
// Solidity: function mintingStopped() constant returns(bool)
func (_MarketToken *MarketTokenCallerSession) MintingStopped() (bool, error) {
	return _MarketToken.Contract.MintingStopped(&_MarketToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MarketToken *MarketTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MarketToken *MarketTokenSession) TotalSupply() (*big.Int, error) {
	return _MarketToken.Contract.TotalSupply(&_MarketToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MarketToken *MarketTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MarketToken.Contract.TotalSupply(&_MarketToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns(bool)
func (_MarketToken *MarketTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns(bool)
func (_MarketToken *MarketTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Approve(&_MarketToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns(bool)
func (_MarketToken *MarketTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Approve(&_MarketToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns()
func (_MarketToken *MarketTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns()
func (_MarketToken *MarketTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Burn(&_MarketToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns()
func (_MarketToken *MarketTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Burn(&_MarketToken.TransactOpts, amount)
}

// BurnAll is a paid mutator transaction binding the contract method 0x7e9d2ac1.
//
// Solidity: function burnAll(holder address) returns()
func (_MarketToken *MarketTokenTransactor) BurnAll(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "burnAll", holder)
}

// BurnAll is a paid mutator transaction binding the contract method 0x7e9d2ac1.
//
// Solidity: function burnAll(holder address) returns()
func (_MarketToken *MarketTokenSession) BurnAll(holder common.Address) (*types.Transaction, error) {
	return _MarketToken.Contract.BurnAll(&_MarketToken.TransactOpts, holder)
}

// BurnAll is a paid mutator transaction binding the contract method 0x7e9d2ac1.
//
// Solidity: function burnAll(holder address) returns()
func (_MarketToken *MarketTokenTransactorSession) BurnAll(holder common.Address) (*types.Transaction, error) {
	return _MarketToken.Contract.BurnAll(&_MarketToken.TransactOpts, holder)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, amount uint256) returns()
func (_MarketToken *MarketTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "decreaseApproval", spender, amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, amount uint256) returns()
func (_MarketToken *MarketTokenSession) DecreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.DecreaseApproval(&_MarketToken.TransactOpts, spender, amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, amount uint256) returns()
func (_MarketToken *MarketTokenTransactorSession) DecreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.DecreaseApproval(&_MarketToken.TransactOpts, spender, amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, amount uint256) returns()
func (_MarketToken *MarketTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "increaseApproval", spender, amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, amount uint256) returns()
func (_MarketToken *MarketTokenSession) IncreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.IncreaseApproval(&_MarketToken.TransactOpts, spender, amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, amount uint256) returns()
func (_MarketToken *MarketTokenTransactorSession) IncreaseApproval(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.IncreaseApproval(&_MarketToken.TransactOpts, spender, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(amount uint256) returns()
func (_MarketToken *MarketTokenTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(amount uint256) returns()
func (_MarketToken *MarketTokenSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Mint(&_MarketToken.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(amount uint256) returns()
func (_MarketToken *MarketTokenTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Mint(&_MarketToken.TransactOpts, amount)
}

// SetPrivilegedContracts is a paid mutator transaction binding the contract method 0x730e6e84.
//
// Solidity: function setPrivilegedContracts(market address) returns()
func (_MarketToken *MarketTokenTransactor) SetPrivilegedContracts(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "setPrivilegedContracts", market)
}

// SetPrivilegedContracts is a paid mutator transaction binding the contract method 0x730e6e84.
//
// Solidity: function setPrivilegedContracts(market address) returns()
func (_MarketToken *MarketTokenSession) SetPrivilegedContracts(market common.Address) (*types.Transaction, error) {
	return _MarketToken.Contract.SetPrivilegedContracts(&_MarketToken.TransactOpts, market)
}

// SetPrivilegedContracts is a paid mutator transaction binding the contract method 0x730e6e84.
//
// Solidity: function setPrivilegedContracts(market address) returns()
func (_MarketToken *MarketTokenTransactorSession) SetPrivilegedContracts(market common.Address) (*types.Transaction, error) {
	return _MarketToken.Contract.SetPrivilegedContracts(&_MarketToken.TransactOpts, market)
}

// StopMinting is a paid mutator transaction binding the contract method 0x3e3e0b12.
//
// Solidity: function stopMinting() returns()
func (_MarketToken *MarketTokenTransactor) StopMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "stopMinting")
}

// StopMinting is a paid mutator transaction binding the contract method 0x3e3e0b12.
//
// Solidity: function stopMinting() returns()
func (_MarketToken *MarketTokenSession) StopMinting() (*types.Transaction, error) {
	return _MarketToken.Contract.StopMinting(&_MarketToken.TransactOpts)
}

// StopMinting is a paid mutator transaction binding the contract method 0x3e3e0b12.
//
// Solidity: function stopMinting() returns()
func (_MarketToken *MarketTokenTransactorSession) StopMinting() (*types.Transaction, error) {
	return _MarketToken.Contract.StopMinting(&_MarketToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, amount uint256) returns(bool)
func (_MarketToken *MarketTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, amount uint256) returns(bool)
func (_MarketToken *MarketTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Transfer(&_MarketToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, amount uint256) returns(bool)
func (_MarketToken *MarketTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Transfer(&_MarketToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, amount uint256) returns(bool)
func (_MarketToken *MarketTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, amount uint256) returns(bool)
func (_MarketToken *MarketTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.TransferFrom(&_MarketToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, amount uint256) returns(bool)
func (_MarketToken *MarketTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.TransferFrom(&_MarketToken.TransactOpts, from, to, amount)
}

// MarketTokenApprovalEventIterator is returned from FilterApprovalEvent and is used to iterate over the raw logs and unpacked data for ApprovalEvent events raised by the MarketToken contract.
type MarketTokenApprovalEventIterator struct {
	Event *MarketTokenApprovalEvent // Event containing the contract specifics and raw log

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
func (it *MarketTokenApprovalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketTokenApprovalEvent)
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
		it.Event = new(MarketTokenApprovalEvent)
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
func (it *MarketTokenApprovalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketTokenApprovalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketTokenApprovalEvent represents a ApprovalEvent event raised by the MarketToken contract.
type MarketTokenApprovalEvent struct {
	Holder  common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApprovalEvent is a free log retrieval operation binding the contract event 0x08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e.
//
// Solidity: e ApprovalEvent(holder indexed address, spender indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) FilterApprovalEvent(opts *bind.FilterOpts, holder []common.Address, spender []common.Address) (*MarketTokenApprovalEventIterator, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MarketToken.contract.FilterLogs(opts, "ApprovalEvent", holderRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MarketTokenApprovalEventIterator{contract: _MarketToken.contract, event: "ApprovalEvent", logs: logs, sub: sub}, nil
}

// WatchApprovalEvent is a free log subscription operation binding the contract event 0x08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e.
//
// Solidity: e ApprovalEvent(holder indexed address, spender indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) WatchApprovalEvent(opts *bind.WatchOpts, sink chan<- *MarketTokenApprovalEvent, holder []common.Address, spender []common.Address) (event.Subscription, error) {

	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MarketToken.contract.WatchLogs(opts, "ApprovalEvent", holderRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketTokenApprovalEvent)
				if err := _MarketToken.contract.UnpackLog(event, "ApprovalEvent", log); err != nil {
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

// MarketTokenBurnEventIterator is returned from FilterBurnEvent and is used to iterate over the raw logs and unpacked data for BurnEvent events raised by the MarketToken contract.
type MarketTokenBurnEventIterator struct {
	Event *MarketTokenBurnEvent // Event containing the contract specifics and raw log

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
func (it *MarketTokenBurnEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketTokenBurnEvent)
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
		it.Event = new(MarketTokenBurnEvent)
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
func (it *MarketTokenBurnEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketTokenBurnEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketTokenBurnEvent represents a BurnEvent event raised by the MarketToken contract.
type MarketTokenBurnEvent struct {
	Burner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnEvent is a free log retrieval operation binding the contract event 0x512586160ebd4dc6945ba9ec5d21a1f723f26f3c7aa36cdffb6818d4e7b88030.
//
// Solidity: e BurnEvent(burner indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) FilterBurnEvent(opts *bind.FilterOpts, burner []common.Address) (*MarketTokenBurnEventIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _MarketToken.contract.FilterLogs(opts, "BurnEvent", burnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketTokenBurnEventIterator{contract: _MarketToken.contract, event: "BurnEvent", logs: logs, sub: sub}, nil
}

// WatchBurnEvent is a free log subscription operation binding the contract event 0x512586160ebd4dc6945ba9ec5d21a1f723f26f3c7aa36cdffb6818d4e7b88030.
//
// Solidity: e BurnEvent(burner indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) WatchBurnEvent(opts *bind.WatchOpts, sink chan<- *MarketTokenBurnEvent, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _MarketToken.contract.WatchLogs(opts, "BurnEvent", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketTokenBurnEvent)
				if err := _MarketToken.contract.UnpackLog(event, "BurnEvent", log); err != nil {
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

// MarketTokenMintEventIterator is returned from FilterMintEvent and is used to iterate over the raw logs and unpacked data for MintEvent events raised by the MarketToken contract.
type MarketTokenMintEventIterator struct {
	Event *MarketTokenMintEvent // Event containing the contract specifics and raw log

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
func (it *MarketTokenMintEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketTokenMintEvent)
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
		it.Event = new(MarketTokenMintEvent)
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
func (it *MarketTokenMintEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketTokenMintEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketTokenMintEvent represents a MintEvent event raised by the MarketToken contract.
type MarketTokenMintEvent struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintEvent is a free log retrieval operation binding the contract event 0x3fffaa5804a26fcec0d70b1d0fb0a2d0031df3a5f9c8af2127c2f4360e97b463.
//
// Solidity: e MintEvent(to indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) FilterMintEvent(opts *bind.FilterOpts, to []common.Address) (*MarketTokenMintEventIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MarketToken.contract.FilterLogs(opts, "MintEvent", toRule)
	if err != nil {
		return nil, err
	}
	return &MarketTokenMintEventIterator{contract: _MarketToken.contract, event: "MintEvent", logs: logs, sub: sub}, nil
}

// WatchMintEvent is a free log subscription operation binding the contract event 0x3fffaa5804a26fcec0d70b1d0fb0a2d0031df3a5f9c8af2127c2f4360e97b463.
//
// Solidity: e MintEvent(to indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) WatchMintEvent(opts *bind.WatchOpts, sink chan<- *MarketTokenMintEvent, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MarketToken.contract.WatchLogs(opts, "MintEvent", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketTokenMintEvent)
				if err := _MarketToken.contract.UnpackLog(event, "MintEvent", log); err != nil {
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

// MarketTokenMintStoppedEventIterator is returned from FilterMintStoppedEvent and is used to iterate over the raw logs and unpacked data for MintStoppedEvent events raised by the MarketToken contract.
type MarketTokenMintStoppedEventIterator struct {
	Event *MarketTokenMintStoppedEvent // Event containing the contract specifics and raw log

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
func (it *MarketTokenMintStoppedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketTokenMintStoppedEvent)
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
		it.Event = new(MarketTokenMintStoppedEvent)
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
func (it *MarketTokenMintStoppedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketTokenMintStoppedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketTokenMintStoppedEvent represents a MintStoppedEvent event raised by the MarketToken contract.
type MarketTokenMintStoppedEvent struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintStoppedEvent is a free log retrieval operation binding the contract event 0x6ae65cb1ae147362b8906c28687a17e1712ca926bea50710ba279849cec33471.
//
// Solidity: e MintStoppedEvent()
func (_MarketToken *MarketTokenFilterer) FilterMintStoppedEvent(opts *bind.FilterOpts) (*MarketTokenMintStoppedEventIterator, error) {

	logs, sub, err := _MarketToken.contract.FilterLogs(opts, "MintStoppedEvent")
	if err != nil {
		return nil, err
	}
	return &MarketTokenMintStoppedEventIterator{contract: _MarketToken.contract, event: "MintStoppedEvent", logs: logs, sub: sub}, nil
}

// WatchMintStoppedEvent is a free log subscription operation binding the contract event 0x6ae65cb1ae147362b8906c28687a17e1712ca926bea50710ba279849cec33471.
//
// Solidity: e MintStoppedEvent()
func (_MarketToken *MarketTokenFilterer) WatchMintStoppedEvent(opts *bind.WatchOpts, sink chan<- *MarketTokenMintStoppedEvent) (event.Subscription, error) {

	logs, sub, err := _MarketToken.contract.WatchLogs(opts, "MintStoppedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketTokenMintStoppedEvent)
				if err := _MarketToken.contract.UnpackLog(event, "MintStoppedEvent", log); err != nil {
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

// MarketTokenTransferEventIterator is returned from FilterTransferEvent and is used to iterate over the raw logs and unpacked data for TransferEvent events raised by the MarketToken contract.
type MarketTokenTransferEventIterator struct {
	Event *MarketTokenTransferEvent // Event containing the contract specifics and raw log

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
func (it *MarketTokenTransferEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketTokenTransferEvent)
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
		it.Event = new(MarketTokenTransferEvent)
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
func (it *MarketTokenTransferEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketTokenTransferEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketTokenTransferEvent represents a TransferEvent event raised by the MarketToken contract.
type MarketTokenTransferEvent struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferEvent is a free log retrieval operation binding the contract event 0xeaf1c4b3ce0f4f62a2bae7eb3e68225c75f7e6ff4422073b7437b9a78d25f170.
//
// Solidity: e TransferEvent(from indexed address, to indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) FilterTransferEvent(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MarketTokenTransferEventIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MarketToken.contract.FilterLogs(opts, "TransferEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MarketTokenTransferEventIterator{contract: _MarketToken.contract, event: "TransferEvent", logs: logs, sub: sub}, nil
}

// WatchTransferEvent is a free log subscription operation binding the contract event 0xeaf1c4b3ce0f4f62a2bae7eb3e68225c75f7e6ff4422073b7437b9a78d25f170.
//
// Solidity: e TransferEvent(from indexed address, to indexed address, amount uint256)
func (_MarketToken *MarketTokenFilterer) WatchTransferEvent(opts *bind.WatchOpts, sink chan<- *MarketTokenTransferEvent, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MarketToken.contract.WatchLogs(opts, "TransferEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketTokenTransferEvent)
				if err := _MarketToken.contract.UnpackLog(event, "TransferEvent", log); err != nil {
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
const ParameterizerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getVoteBy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getConversionSlopeNumerator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"param\",\"type\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"getParamHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"param\",\"type\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"reparameterize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"paramHash\",\"type\":\"bytes32\"}],\"name\":\"resolveReparam\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getChallengeStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getConversionSlopeDenominator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQuorum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"paramHash\",\"type\":\"bytes32\"}],\"name\":\"getReparam\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getListReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getConversionRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"votingAddr\",\"type\":\"address\"},{\"name\":\"challengeStake\",\"type\":\"uint256\"},{\"name\":\"conversionRate\",\"type\":\"uint256\"},{\"name\":\"conversionSlopeDenominator\",\"type\":\"uint256\"},{\"name\":\"conversionSlopeNumerator\",\"type\":\"uint256\"},{\"name\":\"listReward\",\"type\":\"uint256\"},{\"name\":\"quorum\",\"type\":\"uint256\"},{\"name\":\"voteBy\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"paramHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"param\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReparamProposedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"name\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReparamFailedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"param\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReparamSucceededEvent\",\"type\":\"event\"}]"

// ParameterizerBin is the compiled bytecode used for deploying new contracts.
const ParameterizerBin = `0x608060405234801561001057600080fd5b5060405161010080610d76833981018060405261010081101561003257600080fd5b508051602082015160408301516060840151608085015160a086015160c087015160e09097015160018054600160a060020a031916600160a060020a0390981697909717909655600294909455600392909255600455600555600655600791909155600855610cd0806100a66000396000f3fe608060405234801561001057600080fd5b50600436106100c6576000357c0100000000000000000000000000000000000000000000000000000000900480635aebe7681161008e5780635aebe768146101585780637f58af4a14610160578063c26c12eb14610168578063d90f8f0d14610170578063dc2b2853146101b7578063f36089ec146101bf576100c6565b80632d0d2bc6146100cb5780632de2b6ac146100e55780632f6b6533146100ed5780633598cdfb14610113578063435c709a1461013b575b600080fd5b6100d36101c7565b60408051918252519081900360200190f35b6100d36101cd565b6100d36004803603604081101561010357600080fd5b5060ff81351690602001356101d3565b6101396004803603604081101561012957600080fd5b5060ff813516906020013561022a565b005b6101396004803603602081101561015157600080fd5b5035610568565b6100d3610afe565b6100d3610b04565b6100d3610b0a565b61018d6004803603602081101561018657600080fd5b5035610b10565b60408051600160a060020a03909416845260ff909216602084015282820152519081900360600190f35b6100d3610b3f565b6100d3610b45565b60085490565b60055490565b6040805160ff939093167f0100000000000000000000000000000000000000000000000000000000000000026020808501919091526021808501939093528151808503909301835260419093019052805191012090565b600154604080517f62f9a55d0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a03909216916362f9a55d91602480820192602092909190829003018186803b15801561028e57600080fd5b505afa1580156102a2573d6000803e3d6000fd5b505050506040513d60208110156102b857600080fd5b505115156102fa5760405160e560020a62461bcd028152600401808060200182810382526042815260200180610bd56042913960600191505060405180910390fd5b604080517f010000000000000000000000000000000000000000000000000000000000000060ff851602602080830191909152602180830185905283518084039091018152604183018085528151918301919091206001547fb89694c600000000000000000000000000000000000000000000000000000000909252604584018190529351600160a060020a039091169263b89694c6926065808301939192829003018186803b1580156103ad57600080fd5b505afa1580156103c1573d6000803e3d6000fd5b505050506040513d60208110156103d757600080fd5b505115156001141561041d5760405160e560020a62461bcd028152600401808060200182810382526053815260200180610b4c6053913960600191505060405180910390fd5b6040805160608101825233815260ff858116602080840191825283850187815260008781529182905285822094518554935173ffffffffffffffffffffffffffffffffffffffff19909416600160a060020a039182161774ff0000000000000000000000000000000000000000191660a060020a94909516939093029390931784559151600193840155915460085484517f64b34e7d0000000000000000000000000000000000000000000000000000000081526004810187905260036024820152604481019190915293519216926364b34e7d926064808301939282900301818387803b15801561050e57600080fd5b505af1158015610522573d6000803e3d6000fd5b505060408051858152905160ff8716935084925033917f774b91ae28180f7f3c68df1dac76e3990f1fc4605d38dbd3d3d3cb5331fcc7d2919081900360200190a4505050565b600154604080517f62f9a55d0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a03909216916362f9a55d91602480820192602092909190829003018186803b1580156105cc57600080fd5b505afa1580156105e0573d6000803e3d6000fd5b505050506040513d60208110156105f657600080fd5b505115156106385760405160e560020a62461bcd028152600401808060200182810382526042815260200180610c176042913960600191505060405180910390fd5b600154604080517f06db129800000000000000000000000000000000000000000000000000000000815260048101849052600360248201529051600160a060020a03909216916306db129891604480820192602092909190829003018186803b1580156106a457600080fd5b505afa1580156106b8573d6000803e3d6000fd5b505050506040513d60208110156106ce57600080fd5b505115156107105760405160e560020a62461bcd028152600401808060200182810382526036815260200180610b9f6036913960400191505060405180910390fd5b600154604080517f327322c8000000000000000000000000000000000000000000000000000000008152600481018490529051600160a060020a039092169163327322c891602480820192602092909190829003018186803b15801561077557600080fd5b505afa158015610789573d6000803e3d6000fd5b505050506040513d602081101561079f57600080fd5b505115156107e15760405160e560020a62461bcd02815260040180806020018281038252604c815260200180610c59604c913960600191505060405180910390fd5b600154600754604080517f8f354b7900000000000000000000000000000000000000000000000000000000815260048101859052602481019290925251600160a060020a0390921691638f354b7991604480820192602092909190829003018186803b15801561085057600080fd5b505afa158015610864573d6000803e3d6000fd5b505050506040513d602081101561087a57600080fd5b5051156109f95760008181526020819052604090205460a060020a900460ff1660018114156108bd5760008281526020819052604090206001015460025561099d565b60ff8116600214156108e35760008281526020819052604090206001015460035561099d565b60ff8116600314156109095760008281526020819052604090206001015460045561099d565b60ff81166004141561092f5760008281526020819052604090206001015460055561099d565b60ff8116600514156109555760008281526020819052604090206001015460065561099d565b60ff81166006141561097b5760008281526020819052604090206001015460075561099d565b60ff81166007141561099d576000828152602081905260409020600101546008555b6000828152602081815260409182902080546001909101548351908152925160a060020a90910460ff169285927f97eec321a0bc05b245b111890db5afb982d71e8f521f0c61ca80a25a874eee9492918290030190a350610a50565b6000818152602081815260409182902080546001909101548351908152925160a060020a90910460ff169284927f0e4936379fe6406e53422c2e745236484eed6057ef5d46401f7de744f04e39f892918290030190a35b600154604080517f89bb617c000000000000000000000000000000000000000000000000000000008152600481018490529051600160a060020a03909216916389bb617c9160248082019260009290919082900301818387803b158015610ab657600080fd5b505af1158015610aca573d6000803e3d6000fd5b505050600091825250602081905260408120805474ffffffffffffffffffffffffffffffffffffffffff1916815560010155565b60025490565b60045490565b60075490565b60009081526020819052604090208054600190910154600160a060020a0382169260a060020a90920460ff1691565b60065490565b6003549056fe4572726f723a506172616d65746572697a65722e70726f706f73655265706172616d202d2050726f706f736564207265706172616d20697320616c7265616479206120766f74696e672063616e6469646174654572726f723a506172616d65746572697a65722e7265736f6c76655265706172616d202d204d7573742062652061205265706172616d4572726f723a506172616d65746572697a65722e7265706172616d65746572697a65202d2053656e646572206d75737420626520636f756e63696c206d656d6265724572726f723a506172616d65746572697a65722e7265736f6c76655265706172616d202d2053656e646572206d75737420626520636f756e63696c206d656d6265724572726f723a506172616d65746572697a65722e7265736f6c76655265706172616d202d20506f6c6c7320666f7220746869732063616e646964617465206d75737420626520636c6f736564a165627a7a723058203470db7a0d71a76044cdabb19b6d662f68d89980927e196a9a7e08275336ea3b0029`

// DeployParameterizer deploys a new Ethereum contract, binding an instance of Parameterizer to it.
func DeployParameterizer(auth *bind.TransactOpts, backend bind.ContractBackend, votingAddr common.Address, challengeStake *big.Int, conversionRate *big.Int, conversionSlopeDenominator *big.Int, conversionSlopeNumerator *big.Int, listReward *big.Int, quorum *big.Int, voteBy *big.Int) (common.Address, *types.Transaction, *Parameterizer, error) {
	parsed, err := abi.JSON(strings.NewReader(ParameterizerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParameterizerBin), backend, votingAddr, challengeStake, conversionRate, conversionSlopeDenominator, conversionSlopeNumerator, listReward, quorum, voteBy)
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

// GetChallengeStake is a free data retrieval call binding the contract method 0x5aebe768.
//
// Solidity: function getChallengeStake() constant returns(uint256)
func (_Parameterizer *ParameterizerCaller) GetChallengeStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getChallengeStake")
	return *ret0, err
}

// GetChallengeStake is a free data retrieval call binding the contract method 0x5aebe768.
//
// Solidity: function getChallengeStake() constant returns(uint256)
func (_Parameterizer *ParameterizerSession) GetChallengeStake() (*big.Int, error) {
	return _Parameterizer.Contract.GetChallengeStake(&_Parameterizer.CallOpts)
}

// GetChallengeStake is a free data retrieval call binding the contract method 0x5aebe768.
//
// Solidity: function getChallengeStake() constant returns(uint256)
func (_Parameterizer *ParameterizerCallerSession) GetChallengeStake() (*big.Int, error) {
	return _Parameterizer.Contract.GetChallengeStake(&_Parameterizer.CallOpts)
}

// GetConversionRate is a free data retrieval call binding the contract method 0xf36089ec.
//
// Solidity: function getConversionRate() constant returns(uint256)
func (_Parameterizer *ParameterizerCaller) GetConversionRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getConversionRate")
	return *ret0, err
}

// GetConversionRate is a free data retrieval call binding the contract method 0xf36089ec.
//
// Solidity: function getConversionRate() constant returns(uint256)
func (_Parameterizer *ParameterizerSession) GetConversionRate() (*big.Int, error) {
	return _Parameterizer.Contract.GetConversionRate(&_Parameterizer.CallOpts)
}

// GetConversionRate is a free data retrieval call binding the contract method 0xf36089ec.
//
// Solidity: function getConversionRate() constant returns(uint256)
func (_Parameterizer *ParameterizerCallerSession) GetConversionRate() (*big.Int, error) {
	return _Parameterizer.Contract.GetConversionRate(&_Parameterizer.CallOpts)
}

// GetConversionSlopeDenominator is a free data retrieval call binding the contract method 0x7f58af4a.
//
// Solidity: function getConversionSlopeDenominator() constant returns(uint256)
func (_Parameterizer *ParameterizerCaller) GetConversionSlopeDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getConversionSlopeDenominator")
	return *ret0, err
}

// GetConversionSlopeDenominator is a free data retrieval call binding the contract method 0x7f58af4a.
//
// Solidity: function getConversionSlopeDenominator() constant returns(uint256)
func (_Parameterizer *ParameterizerSession) GetConversionSlopeDenominator() (*big.Int, error) {
	return _Parameterizer.Contract.GetConversionSlopeDenominator(&_Parameterizer.CallOpts)
}

// GetConversionSlopeDenominator is a free data retrieval call binding the contract method 0x7f58af4a.
//
// Solidity: function getConversionSlopeDenominator() constant returns(uint256)
func (_Parameterizer *ParameterizerCallerSession) GetConversionSlopeDenominator() (*big.Int, error) {
	return _Parameterizer.Contract.GetConversionSlopeDenominator(&_Parameterizer.CallOpts)
}

// GetConversionSlopeNumerator is a free data retrieval call binding the contract method 0x2de2b6ac.
//
// Solidity: function getConversionSlopeNumerator() constant returns(uint256)
func (_Parameterizer *ParameterizerCaller) GetConversionSlopeNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getConversionSlopeNumerator")
	return *ret0, err
}

// GetConversionSlopeNumerator is a free data retrieval call binding the contract method 0x2de2b6ac.
//
// Solidity: function getConversionSlopeNumerator() constant returns(uint256)
func (_Parameterizer *ParameterizerSession) GetConversionSlopeNumerator() (*big.Int, error) {
	return _Parameterizer.Contract.GetConversionSlopeNumerator(&_Parameterizer.CallOpts)
}

// GetConversionSlopeNumerator is a free data retrieval call binding the contract method 0x2de2b6ac.
//
// Solidity: function getConversionSlopeNumerator() constant returns(uint256)
func (_Parameterizer *ParameterizerCallerSession) GetConversionSlopeNumerator() (*big.Int, error) {
	return _Parameterizer.Contract.GetConversionSlopeNumerator(&_Parameterizer.CallOpts)
}

// GetListReward is a free data retrieval call binding the contract method 0xdc2b2853.
//
// Solidity: function getListReward() constant returns(uint256)
func (_Parameterizer *ParameterizerCaller) GetListReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getListReward")
	return *ret0, err
}

// GetListReward is a free data retrieval call binding the contract method 0xdc2b2853.
//
// Solidity: function getListReward() constant returns(uint256)
func (_Parameterizer *ParameterizerSession) GetListReward() (*big.Int, error) {
	return _Parameterizer.Contract.GetListReward(&_Parameterizer.CallOpts)
}

// GetListReward is a free data retrieval call binding the contract method 0xdc2b2853.
//
// Solidity: function getListReward() constant returns(uint256)
func (_Parameterizer *ParameterizerCallerSession) GetListReward() (*big.Int, error) {
	return _Parameterizer.Contract.GetListReward(&_Parameterizer.CallOpts)
}

// GetParamHash is a free data retrieval call binding the contract method 0x2f6b6533.
//
// Solidity: function getParamHash(param uint8, value uint256) constant returns(bytes32)
func (_Parameterizer *ParameterizerCaller) GetParamHash(opts *bind.CallOpts, param uint8, value *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getParamHash", param, value)
	return *ret0, err
}

// GetParamHash is a free data retrieval call binding the contract method 0x2f6b6533.
//
// Solidity: function getParamHash(param uint8, value uint256) constant returns(bytes32)
func (_Parameterizer *ParameterizerSession) GetParamHash(param uint8, value *big.Int) ([32]byte, error) {
	return _Parameterizer.Contract.GetParamHash(&_Parameterizer.CallOpts, param, value)
}

// GetParamHash is a free data retrieval call binding the contract method 0x2f6b6533.
//
// Solidity: function getParamHash(param uint8, value uint256) constant returns(bytes32)
func (_Parameterizer *ParameterizerCallerSession) GetParamHash(param uint8, value *big.Int) ([32]byte, error) {
	return _Parameterizer.Contract.GetParamHash(&_Parameterizer.CallOpts, param, value)
}

// GetQuorum is a free data retrieval call binding the contract method 0xc26c12eb.
//
// Solidity: function getQuorum() constant returns(uint256)
func (_Parameterizer *ParameterizerCaller) GetQuorum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getQuorum")
	return *ret0, err
}

// GetQuorum is a free data retrieval call binding the contract method 0xc26c12eb.
//
// Solidity: function getQuorum() constant returns(uint256)
func (_Parameterizer *ParameterizerSession) GetQuorum() (*big.Int, error) {
	return _Parameterizer.Contract.GetQuorum(&_Parameterizer.CallOpts)
}

// GetQuorum is a free data retrieval call binding the contract method 0xc26c12eb.
//
// Solidity: function getQuorum() constant returns(uint256)
func (_Parameterizer *ParameterizerCallerSession) GetQuorum() (*big.Int, error) {
	return _Parameterizer.Contract.GetQuorum(&_Parameterizer.CallOpts)
}

// GetReparam is a free data retrieval call binding the contract method 0xd90f8f0d.
//
// Solidity: function getReparam(paramHash bytes32) constant returns(address, uint8, uint256)
func (_Parameterizer *ParameterizerCaller) GetReparam(opts *bind.CallOpts, paramHash [32]byte) (common.Address, uint8, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(uint8)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Parameterizer.contract.Call(opts, out, "getReparam", paramHash)
	return *ret0, *ret1, *ret2, err
}

// GetReparam is a free data retrieval call binding the contract method 0xd90f8f0d.
//
// Solidity: function getReparam(paramHash bytes32) constant returns(address, uint8, uint256)
func (_Parameterizer *ParameterizerSession) GetReparam(paramHash [32]byte) (common.Address, uint8, *big.Int, error) {
	return _Parameterizer.Contract.GetReparam(&_Parameterizer.CallOpts, paramHash)
}

// GetReparam is a free data retrieval call binding the contract method 0xd90f8f0d.
//
// Solidity: function getReparam(paramHash bytes32) constant returns(address, uint8, uint256)
func (_Parameterizer *ParameterizerCallerSession) GetReparam(paramHash [32]byte) (common.Address, uint8, *big.Int, error) {
	return _Parameterizer.Contract.GetReparam(&_Parameterizer.CallOpts, paramHash)
}

// GetVoteBy is a free data retrieval call binding the contract method 0x2d0d2bc6.
//
// Solidity: function getVoteBy() constant returns(uint256)
func (_Parameterizer *ParameterizerCaller) GetVoteBy(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getVoteBy")
	return *ret0, err
}

// GetVoteBy is a free data retrieval call binding the contract method 0x2d0d2bc6.
//
// Solidity: function getVoteBy() constant returns(uint256)
func (_Parameterizer *ParameterizerSession) GetVoteBy() (*big.Int, error) {
	return _Parameterizer.Contract.GetVoteBy(&_Parameterizer.CallOpts)
}

// GetVoteBy is a free data retrieval call binding the contract method 0x2d0d2bc6.
//
// Solidity: function getVoteBy() constant returns(uint256)
func (_Parameterizer *ParameterizerCallerSession) GetVoteBy() (*big.Int, error) {
	return _Parameterizer.Contract.GetVoteBy(&_Parameterizer.CallOpts)
}

// Reparameterize is a paid mutator transaction binding the contract method 0x3598cdfb.
//
// Solidity: function reparameterize(param uint8, value uint256) returns()
func (_Parameterizer *ParameterizerTransactor) Reparameterize(opts *bind.TransactOpts, param uint8, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "reparameterize", param, value)
}

// Reparameterize is a paid mutator transaction binding the contract method 0x3598cdfb.
//
// Solidity: function reparameterize(param uint8, value uint256) returns()
func (_Parameterizer *ParameterizerSession) Reparameterize(param uint8, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.Reparameterize(&_Parameterizer.TransactOpts, param, value)
}

// Reparameterize is a paid mutator transaction binding the contract method 0x3598cdfb.
//
// Solidity: function reparameterize(param uint8, value uint256) returns()
func (_Parameterizer *ParameterizerTransactorSession) Reparameterize(param uint8, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.Reparameterize(&_Parameterizer.TransactOpts, param, value)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(paramHash bytes32) returns()
func (_Parameterizer *ParameterizerTransactor) ResolveReparam(opts *bind.TransactOpts, paramHash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "resolveReparam", paramHash)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(paramHash bytes32) returns()
func (_Parameterizer *ParameterizerSession) ResolveReparam(paramHash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.Contract.ResolveReparam(&_Parameterizer.TransactOpts, paramHash)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(paramHash bytes32) returns()
func (_Parameterizer *ParameterizerTransactorSession) ResolveReparam(paramHash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.Contract.ResolveReparam(&_Parameterizer.TransactOpts, paramHash)
}

// ParameterizerReparamFailedEventIterator is returned from FilterReparamFailedEvent and is used to iterate over the raw logs and unpacked data for ReparamFailedEvent events raised by the Parameterizer contract.
type ParameterizerReparamFailedEventIterator struct {
	Event *ParameterizerReparamFailedEvent // Event containing the contract specifics and raw log

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
func (it *ParameterizerReparamFailedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerReparamFailedEvent)
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
		it.Event = new(ParameterizerReparamFailedEvent)
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
func (it *ParameterizerReparamFailedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerReparamFailedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerReparamFailedEvent represents a ReparamFailedEvent event raised by the Parameterizer contract.
type ParameterizerReparamFailedEvent struct {
	PropHash [32]byte
	Name     uint8
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReparamFailedEvent is a free log retrieval operation binding the contract event 0x0e4936379fe6406e53422c2e745236484eed6057ef5d46401f7de744f04e39f8.
//
// Solidity: e ReparamFailedEvent(propHash indexed bytes32, name indexed uint8, value uint256)
func (_Parameterizer *ParameterizerFilterer) FilterReparamFailedEvent(opts *bind.FilterOpts, propHash [][32]byte, name []uint8) (*ParameterizerReparamFailedEventIterator, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ReparamFailedEvent", propHashRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerReparamFailedEventIterator{contract: _Parameterizer.contract, event: "ReparamFailedEvent", logs: logs, sub: sub}, nil
}

// WatchReparamFailedEvent is a free log subscription operation binding the contract event 0x0e4936379fe6406e53422c2e745236484eed6057ef5d46401f7de744f04e39f8.
//
// Solidity: e ReparamFailedEvent(propHash indexed bytes32, name indexed uint8, value uint256)
func (_Parameterizer *ParameterizerFilterer) WatchReparamFailedEvent(opts *bind.WatchOpts, sink chan<- *ParameterizerReparamFailedEvent, propHash [][32]byte, name []uint8) (event.Subscription, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ReparamFailedEvent", propHashRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerReparamFailedEvent)
				if err := _Parameterizer.contract.UnpackLog(event, "ReparamFailedEvent", log); err != nil {
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

// ParameterizerReparamProposedEventIterator is returned from FilterReparamProposedEvent and is used to iterate over the raw logs and unpacked data for ReparamProposedEvent events raised by the Parameterizer contract.
type ParameterizerReparamProposedEventIterator struct {
	Event *ParameterizerReparamProposedEvent // Event containing the contract specifics and raw log

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
func (it *ParameterizerReparamProposedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerReparamProposedEvent)
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
		it.Event = new(ParameterizerReparamProposedEvent)
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
func (it *ParameterizerReparamProposedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerReparamProposedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerReparamProposedEvent represents a ReparamProposedEvent event raised by the Parameterizer contract.
type ParameterizerReparamProposedEvent struct {
	Proposer  common.Address
	ParamHash [32]byte
	Param     uint8
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReparamProposedEvent is a free log retrieval operation binding the contract event 0x774b91ae28180f7f3c68df1dac76e3990f1fc4605d38dbd3d3d3cb5331fcc7d2.
//
// Solidity: e ReparamProposedEvent(proposer indexed address, paramHash indexed bytes32, param indexed uint8, value uint256)
func (_Parameterizer *ParameterizerFilterer) FilterReparamProposedEvent(opts *bind.FilterOpts, proposer []common.Address, paramHash [][32]byte, param []uint8) (*ParameterizerReparamProposedEventIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var paramHashRule []interface{}
	for _, paramHashItem := range paramHash {
		paramHashRule = append(paramHashRule, paramHashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ReparamProposedEvent", proposerRule, paramHashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerReparamProposedEventIterator{contract: _Parameterizer.contract, event: "ReparamProposedEvent", logs: logs, sub: sub}, nil
}

// WatchReparamProposedEvent is a free log subscription operation binding the contract event 0x774b91ae28180f7f3c68df1dac76e3990f1fc4605d38dbd3d3d3cb5331fcc7d2.
//
// Solidity: e ReparamProposedEvent(proposer indexed address, paramHash indexed bytes32, param indexed uint8, value uint256)
func (_Parameterizer *ParameterizerFilterer) WatchReparamProposedEvent(opts *bind.WatchOpts, sink chan<- *ParameterizerReparamProposedEvent, proposer []common.Address, paramHash [][32]byte, param []uint8) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var paramHashRule []interface{}
	for _, paramHashItem := range paramHash {
		paramHashRule = append(paramHashRule, paramHashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ReparamProposedEvent", proposerRule, paramHashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerReparamProposedEvent)
				if err := _Parameterizer.contract.UnpackLog(event, "ReparamProposedEvent", log); err != nil {
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

// ParameterizerReparamSucceededEventIterator is returned from FilterReparamSucceededEvent and is used to iterate over the raw logs and unpacked data for ReparamSucceededEvent events raised by the Parameterizer contract.
type ParameterizerReparamSucceededEventIterator struct {
	Event *ParameterizerReparamSucceededEvent // Event containing the contract specifics and raw log

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
func (it *ParameterizerReparamSucceededEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerReparamSucceededEvent)
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
		it.Event = new(ParameterizerReparamSucceededEvent)
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
func (it *ParameterizerReparamSucceededEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerReparamSucceededEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerReparamSucceededEvent represents a ReparamSucceededEvent event raised by the Parameterizer contract.
type ParameterizerReparamSucceededEvent struct {
	PropHash [32]byte
	Param    uint8
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReparamSucceededEvent is a free log retrieval operation binding the contract event 0x97eec321a0bc05b245b111890db5afb982d71e8f521f0c61ca80a25a874eee94.
//
// Solidity: e ReparamSucceededEvent(propHash indexed bytes32, param indexed uint8, value uint256)
func (_Parameterizer *ParameterizerFilterer) FilterReparamSucceededEvent(opts *bind.FilterOpts, propHash [][32]byte, param []uint8) (*ParameterizerReparamSucceededEventIterator, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ReparamSucceededEvent", propHashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerReparamSucceededEventIterator{contract: _Parameterizer.contract, event: "ReparamSucceededEvent", logs: logs, sub: sub}, nil
}

// WatchReparamSucceededEvent is a free log subscription operation binding the contract event 0x97eec321a0bc05b245b111890db5afb982d71e8f521f0c61ca80a25a874eee94.
//
// Solidity: e ReparamSucceededEvent(propHash indexed bytes32, param indexed uint8, value uint256)
func (_Parameterizer *ParameterizerFilterer) WatchReparamSucceededEvent(opts *bind.WatchOpts, sink chan<- *ParameterizerReparamSucceededEvent, propHash [][32]byte, param []uint8) (event.Subscription, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ReparamSucceededEvent", propHashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerReparamSucceededEvent)
				if err := _Parameterizer.contract.UnpackLog(event, "ReparamSucceededEvent", log); err != nil {
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
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820d864a81554d5dcee11d396cd7287901c09ccc4ee4795cc188a40fb3b6c0977140029`

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

// VotingABI is the input ABI used to generate the binding from.
const VotingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"candidateIs\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"member\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"pollClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"member\",\"type\":\"address\"}],\"name\":\"removeFromCouncil\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"member\",\"type\":\"address\"}],\"name\":\"addToCouncil\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"member\",\"type\":\"address\"}],\"name\":\"inCouncil\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"voteBy\",\"type\":\"uint256\"}],\"name\":\"addCandidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"removeCandidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"quorum\",\"type\":\"uint256\"}],\"name\":\"didPass\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"market\",\"type\":\"address\"},{\"name\":\"parameterizer\",\"type\":\"address\"}],\"name\":\"setPrivilegedContracts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCouncil\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPrivilegedAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"VotedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":true,\"name\":\"voteBy\",\"type\":\"uint256\"}],\"name\":\"CandidateAddedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"CandidateRemovedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"member\",\"type\":\"address\"}],\"name\":\"CouncilMemberAddedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"member\",\"type\":\"address\"}],\"name\":\"CouncilMemberRemovedEvent\",\"type\":\"event\"}]"

// VotingBin is the compiled bytecode used for deploying new contracts.
const VotingBin = `0x608060405234801561001057600080fd5b5060068054600160a060020a0319163317905561144c806100326000396000f3fe608060405234801561001057600080fd5b506004361061011d576000357c01000000000000000000000000000000000000000000000000000000009004806389bb617c116100b4578063b181369511610083578063b181369514610328578063b6896e5f14610330578063b89694c61461035e578063dfb6419f1461037b5761011d565b806389bb617c1461029d5780638f354b79146102ba578063a69beaba146102dd578063a7e44790146102fa5761011d565b8063335d8080116100f0578063335d8080146101fd578063471bb3091461022557806362f9a55d1461024b57806364b34e7d146102715761011d565b806306a49fce1461012257806306db12981461017a578063187a9c07146101b4578063327322c8146101e0575b600080fd5b61012a6103ba565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561016657818101518382015260200161014e565b505050509050019250505060405180910390f35b6101a06004803603604081101561019057600080fd5b508035906020013560ff16610413565b604080519115158252519081900360200190f35b6101a0600480360360408110156101ca57600080fd5b5080359060200135600160a060020a0316610434565b6101a0600480360360208110156101f657600080fd5b50356104fb565b6102236004803603602081101561021357600080fd5b5035600160a060020a0316610561565b005b6102236004803603602081101561023b57600080fd5b5035600160a060020a03166106dd565b6101a06004803603602081101561026157600080fd5b5035600160a060020a031661080c565b6102236004803603606081101561028757600080fd5b5080359060ff6020820135169060400135610864565b610223600480360360208110156102b357600080fd5b50356109cc565b6101a0600480360360408110156102d057600080fd5b5080359060200135610b62565b610223600480360360208110156102f357600080fd5b5035610c92565b6102236004803603604081101561031057600080fd5b50600160a060020a0381358116916020013516610e61565b61012a610f7e565b610338610fdf565b60408051600160a060020a03938416815291909216602082015281519081900390910190f35b6101a06004803603602081101561037457600080fd5b5035610ff6565b6103986004803603602081101561039157600080fd5b503561103b565b6040805160ff9094168452602084019290925282820152519081900360600190f35b6060600180548060200260200160405190810160405280929190818152602001828054801561040857602002820191906000526020600020905b8154815260200190600101908083116103f4575b505050505090505b90565b60008281526020819052604090206001015460ff8281169116145b92915050565b600061043f83610ff6565b151561047f5760405160e560020a62461bcd02815260040180806020018281038252602f815260200180611326602f913960400191505060405180910390fd5b60005b6000848152602081905260409020600401548110156104f15760008481526020819052604090206004018054600160a060020a0385169190839081106104c457fe5b600091825260209091200154600160a060020a031614156104e957600191505061042e565b600101610482565b5060009392505050565b600061050682610ff6565b15156105465760405160e560020a62461bcd0281526004018080602001828103825260328152602001806112c96032913960400191505060405180910390fd5b5060008181526020819052604090206002015442115b919050565b600654600160a060020a03163314806105845750600454600160a060020a031633145b806105995750600554600160a060020a031633145b15156105d95760405160e560020a62461bcd0281526004018080602001828103825260408152602001806112006040913960400191505060405180910390fd5b6105e28161080c565b156106da57600160a060020a03811660009081526002602052604081205460038054919291600019810190811061061557fe5b60009182526020909120015460038054600160a060020a03909216925082918490811061063e57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790556003805490610687906000198301611077565b50600160a060020a0380821660009081526002602052604080822085905591851680825282822082905591517f564e70d59b50a6b32776c3d4a1e4e378a77dc9e24c3e2fbfa04a79166a67489b9190a250505b50565b600654600160a060020a03163314806107005750600454600160a060020a031633145b806107155750600554600160a060020a031633145b15156107555760405160e560020a62461bcd0281526004018080602001828103825260408152602001806112006040913960400191505060405180910390fd5b61075e8161080c565b1515600114156107a25760405160e560020a62461bcd0281526004018080602001828103825260348152602001806113926034913960400191505060405180910390fd5b60038054600181019091557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b81018054600160a060020a0390931673ffffffffffffffffffffffffffffffffffffffff199093168317905560009182526002602052604090912055565b60035460009015156108205750600061055c565b600160a060020a03821660008181526002602052604090205460038054909190811061084857fe5b600091825260209091200154600160a060020a03161492915050565b600654600160a060020a03163314806108875750600454600160a060020a031633145b8061089c5750600554600160a060020a031633145b15156108dc5760405160e560020a62461bcd0281526004018080602001828103825260408152602001806112006040913960400191505060405180910390fd5b6108e583610ff6565b1515600114156109295760405160e560020a62461bcd0281526004018080602001828103825260348152602001806112956034913960400191505060405180910390fd5b600061093b428363ffffffff61106116565b6000858152602081905260408082206001805480820182558185527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf681018a905582558101805460ff191660ff891690811790915560028201859055915193945092859288917fb7fdb794d49b8471a97fc8b1f90d1af97e75e0afc5f56d3b12b696e6d80f01869190a45050505050565b600654600160a060020a03163314806109ef5750600454600160a060020a031633145b80610a045750600554600160a060020a031633145b1515610a445760405160e560020a62461bcd0281526004018080602001828103825260408152602001806112006040913960400191505060405180910390fd5b610a4d81610ff6565b1515610a8d5760405160e560020a62461bcd0281526004018080602001828103825260378152602001806111c96037913960400191505060405180910390fd5b600081815260208190526040812054600180549192916000198101908110610ab157fe5b9060005260206000200154905080600183815481101515610ace57fe5b6000918252602090912001556001805490610aed906000198301611077565b50600081815260208190526040808220849055848252812081815560018101805460ff19169055600281018290556003810182905590610b3060048301826110a0565b505060405183907f53ce6706735b5cd3034bf68c2a532801c1efb4ba72410b5af2167a8462f15b7d90600090a2505050565b6000610b6d83610ff6565b1515610bad5760405160e560020a62461bcd02815260040180806020018281038252602f8152602001806113c6602f913960400191505060405180910390fd5b6000838152602081905260409020600201544211610bff5760405160e560020a62461bcd02815260040180806020018281038252603d815260200180611355603d913960400191505060405180910390fd5b600354600010610c435760405160e560020a62461bcd02815260040180806020018281038252602981526020018061126c6029913960400191505060405180910390fd5b6000838152602081905260409020600301541515610c635750801561042e565b60038054600085815260208190526040902090910154839190606402811515610c8857fe5b041015905061042e565b610c9b3361080c565b1515610cdb5760405160e560020a62461bcd0281526004018080602001828103825260318152602001806111586031913960400191505060405180910390fd5b610ce481610ff6565b1515610d245760405160e560020a62461bcd02815260040180806020018281038252602c815260200180611240602c913960400191505060405180910390fd5b6000818152602081905260409020600201544210610d765760405160e560020a62461bcd0281526004018080602001828103825260388152602001806110d96038913960400191505060405180910390fd5b610d808133610434565b151560011415610dc45760405160e560020a62461bcd02815260040180806020018281038252602c8152602001806113f5602c913960400191505060405180910390fd5b600081815260208181526040822060048101805460018181018355918552928420909201805473ffffffffffffffffffffffffffffffffffffffff19163317905591839052600390910154610e1e9163ffffffff61106116565b600082815260208190526040808220600301929092559051829133917f81681507a95428bb04c5f5ac127e404caf6422d0dcf2ff77612b9201c3aa33bf9190a350565b600654600160a060020a03163314610ead5760405160e560020a62461bcd02815260040180806020018281038252602b8152602001806112fb602b913960400191505060405180910390fd5b600454600160a060020a031615610ef85760405160e560020a62461bcd0281526004018080602001828103825260408152602001806111896040913960400191505060405180910390fd5b600554600160a060020a031615610f435760405160e560020a62461bcd0281526004018080602001828103825260478152602001806111116047913960600191505060405180910390fd5b60048054600160a060020a0393841673ffffffffffffffffffffffffffffffffffffffff199182161790915560058054929093169116179055565b6060600380548060200260200160405190810160405280929190818152602001828054801561040857602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610fb8575050505050905090565b600454600554600160a060020a0391821691169091565b600154600090151561100a5750600061055c565b60008281526020819052604090205460018054849290811061102857fe5b9060005260206000200154149050919050565b60009081526020819052604090206001810154600282015460039092015460ff90911692565b60008282018381101561107057fe5b9392505050565b81548183558181111561109b5760008381526020902061109b9181019083016110ba565b505050565b50805460008255906000526020600020908101906106da91905b61041091905b808211156110d457600081556001016110c0565b509056fe4572726f723a566f74696e672e766f7465202d20506f6c6c696e6720697320636c6f73656420666f7220746869732063616e6469646174654572726f723a566f74696e672e73657450726976696c65676564436f6e747261637473202d20506172616d65746572697a6572206164647265737320616c7265616479207365744572726f723a566f74696e672e766f7465202d2053656e646572206d75737420626520636f756e63696c206d656d6265724572726f723a566f74696e672e73657450726976696c65676564436f6e747261637473202d204d61726b6574206164647265737320616c7265616479207365744572726f723a566f74696e672e72656d6f766543616e646964617465202d2043616e64696461746520646f6573206e6f742065786973744572726f723a566f74696e672e68617350726976696c656765202d2053656e646572206d75737420626520612070726976696c6567656420636f6e74726163744572726f723a566f74696e672e766f7465202d2043616e64696461746520646f6573206e6f742065786973744572726f723a566f74696e672e64696450617373202d204e6f20636f756e63696c206d656d626572734572726f723a566f74696e672e61646443616e646964617465202d2043616e64696461746520616c7265616479206578697374734572726f723a566f74696e672e706f6c6c436c6f736564202d2043616e64696461746520646f6573206e6f742065786973744572726f723a566f74696e672e69734f776e6572202d2053656e646572206d757374206265206f776e65724572726f723a566f74696e672e646964566f7465202d2043616e64696461746520646f6573206e6f742065786973744572726f723a566f74696e672e70617373202d20506f6c6c696e67206d75737420626520636c6f73656420666f7220746869732063616e6469646174654572726f723a566f74696e672e616464546f436f756e63696c202d20416c7265616479206120636f756e63696c206d656d6265724572726f723a566f74696e672e64696450617373202d2043616e64696461746520646f6573206e6f742065786973744572726f723a566f74696e672e766f7465202d2053656e6465722068617320616c726561647920766f746564a165627a7a7230582090e4e3a6c60dcb80e6f1c779d468ad71d3eb1825a92302e2b49e37348f58be290029`

// DeployVoting deploys a new Ethereum contract, binding an instance of Voting to it.
func DeployVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Voting, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// CandidateIs is a free data retrieval call binding the contract method 0x06db1298.
//
// Solidity: function candidateIs(hash bytes32, kind uint8) constant returns(bool)
func (_Voting *VotingCaller) CandidateIs(opts *bind.CallOpts, hash [32]byte, kind uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "candidateIs", hash, kind)
	return *ret0, err
}

// CandidateIs is a free data retrieval call binding the contract method 0x06db1298.
//
// Solidity: function candidateIs(hash bytes32, kind uint8) constant returns(bool)
func (_Voting *VotingSession) CandidateIs(hash [32]byte, kind uint8) (bool, error) {
	return _Voting.Contract.CandidateIs(&_Voting.CallOpts, hash, kind)
}

// CandidateIs is a free data retrieval call binding the contract method 0x06db1298.
//
// Solidity: function candidateIs(hash bytes32, kind uint8) constant returns(bool)
func (_Voting *VotingCallerSession) CandidateIs(hash [32]byte, kind uint8) (bool, error) {
	return _Voting.Contract.CandidateIs(&_Voting.CallOpts, hash, kind)
}

// DidPass is a free data retrieval call binding the contract method 0x8f354b79.
//
// Solidity: function didPass(hash bytes32, quorum uint256) constant returns(bool)
func (_Voting *VotingCaller) DidPass(opts *bind.CallOpts, hash [32]byte, quorum *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "didPass", hash, quorum)
	return *ret0, err
}

// DidPass is a free data retrieval call binding the contract method 0x8f354b79.
//
// Solidity: function didPass(hash bytes32, quorum uint256) constant returns(bool)
func (_Voting *VotingSession) DidPass(hash [32]byte, quorum *big.Int) (bool, error) {
	return _Voting.Contract.DidPass(&_Voting.CallOpts, hash, quorum)
}

// DidPass is a free data retrieval call binding the contract method 0x8f354b79.
//
// Solidity: function didPass(hash bytes32, quorum uint256) constant returns(bool)
func (_Voting *VotingCallerSession) DidPass(hash [32]byte, quorum *big.Int) (bool, error) {
	return _Voting.Contract.DidPass(&_Voting.CallOpts, hash, quorum)
}

// DidVote is a free data retrieval call binding the contract method 0x187a9c07.
//
// Solidity: function didVote(hash bytes32, member address) constant returns(bool)
func (_Voting *VotingCaller) DidVote(opts *bind.CallOpts, hash [32]byte, member common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "didVote", hash, member)
	return *ret0, err
}

// DidVote is a free data retrieval call binding the contract method 0x187a9c07.
//
// Solidity: function didVote(hash bytes32, member address) constant returns(bool)
func (_Voting *VotingSession) DidVote(hash [32]byte, member common.Address) (bool, error) {
	return _Voting.Contract.DidVote(&_Voting.CallOpts, hash, member)
}

// DidVote is a free data retrieval call binding the contract method 0x187a9c07.
//
// Solidity: function didVote(hash bytes32, member address) constant returns(bool)
func (_Voting *VotingCallerSession) DidVote(hash [32]byte, member common.Address) (bool, error) {
	return _Voting.Contract.DidVote(&_Voting.CallOpts, hash, member)
}

// GetCandidate is a free data retrieval call binding the contract method 0xdfb6419f.
//
// Solidity: function getCandidate(hash bytes32) constant returns(uint8, uint256, uint256)
func (_Voting *VotingCaller) GetCandidate(opts *bind.CallOpts, hash [32]byte) (uint8, *big.Int, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Voting.contract.Call(opts, out, "getCandidate", hash)
	return *ret0, *ret1, *ret2, err
}

// GetCandidate is a free data retrieval call binding the contract method 0xdfb6419f.
//
// Solidity: function getCandidate(hash bytes32) constant returns(uint8, uint256, uint256)
func (_Voting *VotingSession) GetCandidate(hash [32]byte) (uint8, *big.Int, *big.Int, error) {
	return _Voting.Contract.GetCandidate(&_Voting.CallOpts, hash)
}

// GetCandidate is a free data retrieval call binding the contract method 0xdfb6419f.
//
// Solidity: function getCandidate(hash bytes32) constant returns(uint8, uint256, uint256)
func (_Voting *VotingCallerSession) GetCandidate(hash [32]byte) (uint8, *big.Int, *big.Int, error) {
	return _Voting.Contract.GetCandidate(&_Voting.CallOpts, hash)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(bytes32[])
func (_Voting *VotingCaller) GetCandidates(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "getCandidates")
	return *ret0, err
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(bytes32[])
func (_Voting *VotingSession) GetCandidates() ([][32]byte, error) {
	return _Voting.Contract.GetCandidates(&_Voting.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(bytes32[])
func (_Voting *VotingCallerSession) GetCandidates() ([][32]byte, error) {
	return _Voting.Contract.GetCandidates(&_Voting.CallOpts)
}

// GetCouncil is a free data retrieval call binding the contract method 0xb1813695.
//
// Solidity: function getCouncil() constant returns(address[])
func (_Voting *VotingCaller) GetCouncil(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "getCouncil")
	return *ret0, err
}

// GetCouncil is a free data retrieval call binding the contract method 0xb1813695.
//
// Solidity: function getCouncil() constant returns(address[])
func (_Voting *VotingSession) GetCouncil() ([]common.Address, error) {
	return _Voting.Contract.GetCouncil(&_Voting.CallOpts)
}

// GetCouncil is a free data retrieval call binding the contract method 0xb1813695.
//
// Solidity: function getCouncil() constant returns(address[])
func (_Voting *VotingCallerSession) GetCouncil() ([]common.Address, error) {
	return _Voting.Contract.GetCouncil(&_Voting.CallOpts)
}

// GetPrivilegedAddresses is a free data retrieval call binding the contract method 0xb6896e5f.
//
// Solidity: function getPrivilegedAddresses() constant returns(address, address)
func (_Voting *VotingCaller) GetPrivilegedAddresses(opts *bind.CallOpts) (common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Voting.contract.Call(opts, out, "getPrivilegedAddresses")
	return *ret0, *ret1, err
}

// GetPrivilegedAddresses is a free data retrieval call binding the contract method 0xb6896e5f.
//
// Solidity: function getPrivilegedAddresses() constant returns(address, address)
func (_Voting *VotingSession) GetPrivilegedAddresses() (common.Address, common.Address, error) {
	return _Voting.Contract.GetPrivilegedAddresses(&_Voting.CallOpts)
}

// GetPrivilegedAddresses is a free data retrieval call binding the contract method 0xb6896e5f.
//
// Solidity: function getPrivilegedAddresses() constant returns(address, address)
func (_Voting *VotingCallerSession) GetPrivilegedAddresses() (common.Address, common.Address, error) {
	return _Voting.Contract.GetPrivilegedAddresses(&_Voting.CallOpts)
}

// InCouncil is a free data retrieval call binding the contract method 0x62f9a55d.
//
// Solidity: function inCouncil(member address) constant returns(bool)
func (_Voting *VotingCaller) InCouncil(opts *bind.CallOpts, member common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "inCouncil", member)
	return *ret0, err
}

// InCouncil is a free data retrieval call binding the contract method 0x62f9a55d.
//
// Solidity: function inCouncil(member address) constant returns(bool)
func (_Voting *VotingSession) InCouncil(member common.Address) (bool, error) {
	return _Voting.Contract.InCouncil(&_Voting.CallOpts, member)
}

// InCouncil is a free data retrieval call binding the contract method 0x62f9a55d.
//
// Solidity: function inCouncil(member address) constant returns(bool)
func (_Voting *VotingCallerSession) InCouncil(member common.Address) (bool, error) {
	return _Voting.Contract.InCouncil(&_Voting.CallOpts, member)
}

// IsCandidate is a free data retrieval call binding the contract method 0xb89694c6.
//
// Solidity: function isCandidate(hash bytes32) constant returns(bool)
func (_Voting *VotingCaller) IsCandidate(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "isCandidate", hash)
	return *ret0, err
}

// IsCandidate is a free data retrieval call binding the contract method 0xb89694c6.
//
// Solidity: function isCandidate(hash bytes32) constant returns(bool)
func (_Voting *VotingSession) IsCandidate(hash [32]byte) (bool, error) {
	return _Voting.Contract.IsCandidate(&_Voting.CallOpts, hash)
}

// IsCandidate is a free data retrieval call binding the contract method 0xb89694c6.
//
// Solidity: function isCandidate(hash bytes32) constant returns(bool)
func (_Voting *VotingCallerSession) IsCandidate(hash [32]byte) (bool, error) {
	return _Voting.Contract.IsCandidate(&_Voting.CallOpts, hash)
}

// PollClosed is a free data retrieval call binding the contract method 0x327322c8.
//
// Solidity: function pollClosed(hash bytes32) constant returns(bool)
func (_Voting *VotingCaller) PollClosed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "pollClosed", hash)
	return *ret0, err
}

// PollClosed is a free data retrieval call binding the contract method 0x327322c8.
//
// Solidity: function pollClosed(hash bytes32) constant returns(bool)
func (_Voting *VotingSession) PollClosed(hash [32]byte) (bool, error) {
	return _Voting.Contract.PollClosed(&_Voting.CallOpts, hash)
}

// PollClosed is a free data retrieval call binding the contract method 0x327322c8.
//
// Solidity: function pollClosed(hash bytes32) constant returns(bool)
func (_Voting *VotingCallerSession) PollClosed(hash [32]byte) (bool, error) {
	return _Voting.Contract.PollClosed(&_Voting.CallOpts, hash)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x64b34e7d.
//
// Solidity: function addCandidate(hash bytes32, kind uint8, voteBy uint256) returns()
func (_Voting *VotingTransactor) AddCandidate(opts *bind.TransactOpts, hash [32]byte, kind uint8, voteBy *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "addCandidate", hash, kind, voteBy)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x64b34e7d.
//
// Solidity: function addCandidate(hash bytes32, kind uint8, voteBy uint256) returns()
func (_Voting *VotingSession) AddCandidate(hash [32]byte, kind uint8, voteBy *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, hash, kind, voteBy)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x64b34e7d.
//
// Solidity: function addCandidate(hash bytes32, kind uint8, voteBy uint256) returns()
func (_Voting *VotingTransactorSession) AddCandidate(hash [32]byte, kind uint8, voteBy *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, hash, kind, voteBy)
}

// AddToCouncil is a paid mutator transaction binding the contract method 0x471bb309.
//
// Solidity: function addToCouncil(member address) returns()
func (_Voting *VotingTransactor) AddToCouncil(opts *bind.TransactOpts, member common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "addToCouncil", member)
}

// AddToCouncil is a paid mutator transaction binding the contract method 0x471bb309.
//
// Solidity: function addToCouncil(member address) returns()
func (_Voting *VotingSession) AddToCouncil(member common.Address) (*types.Transaction, error) {
	return _Voting.Contract.AddToCouncil(&_Voting.TransactOpts, member)
}

// AddToCouncil is a paid mutator transaction binding the contract method 0x471bb309.
//
// Solidity: function addToCouncil(member address) returns()
func (_Voting *VotingTransactorSession) AddToCouncil(member common.Address) (*types.Transaction, error) {
	return _Voting.Contract.AddToCouncil(&_Voting.TransactOpts, member)
}

// RemoveCandidate is a paid mutator transaction binding the contract method 0x89bb617c.
//
// Solidity: function removeCandidate(hash bytes32) returns()
func (_Voting *VotingTransactor) RemoveCandidate(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "removeCandidate", hash)
}

// RemoveCandidate is a paid mutator transaction binding the contract method 0x89bb617c.
//
// Solidity: function removeCandidate(hash bytes32) returns()
func (_Voting *VotingSession) RemoveCandidate(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.RemoveCandidate(&_Voting.TransactOpts, hash)
}

// RemoveCandidate is a paid mutator transaction binding the contract method 0x89bb617c.
//
// Solidity: function removeCandidate(hash bytes32) returns()
func (_Voting *VotingTransactorSession) RemoveCandidate(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.RemoveCandidate(&_Voting.TransactOpts, hash)
}

// RemoveFromCouncil is a paid mutator transaction binding the contract method 0x335d8080.
//
// Solidity: function removeFromCouncil(member address) returns()
func (_Voting *VotingTransactor) RemoveFromCouncil(opts *bind.TransactOpts, member common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "removeFromCouncil", member)
}

// RemoveFromCouncil is a paid mutator transaction binding the contract method 0x335d8080.
//
// Solidity: function removeFromCouncil(member address) returns()
func (_Voting *VotingSession) RemoveFromCouncil(member common.Address) (*types.Transaction, error) {
	return _Voting.Contract.RemoveFromCouncil(&_Voting.TransactOpts, member)
}

// RemoveFromCouncil is a paid mutator transaction binding the contract method 0x335d8080.
//
// Solidity: function removeFromCouncil(member address) returns()
func (_Voting *VotingTransactorSession) RemoveFromCouncil(member common.Address) (*types.Transaction, error) {
	return _Voting.Contract.RemoveFromCouncil(&_Voting.TransactOpts, member)
}

// SetPrivilegedContracts is a paid mutator transaction binding the contract method 0xa7e44790.
//
// Solidity: function setPrivilegedContracts(market address, parameterizer address) returns()
func (_Voting *VotingTransactor) SetPrivilegedContracts(opts *bind.TransactOpts, market common.Address, parameterizer common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "setPrivilegedContracts", market, parameterizer)
}

// SetPrivilegedContracts is a paid mutator transaction binding the contract method 0xa7e44790.
//
// Solidity: function setPrivilegedContracts(market address, parameterizer address) returns()
func (_Voting *VotingSession) SetPrivilegedContracts(market common.Address, parameterizer common.Address) (*types.Transaction, error) {
	return _Voting.Contract.SetPrivilegedContracts(&_Voting.TransactOpts, market, parameterizer)
}

// SetPrivilegedContracts is a paid mutator transaction binding the contract method 0xa7e44790.
//
// Solidity: function setPrivilegedContracts(market address, parameterizer address) returns()
func (_Voting *VotingTransactorSession) SetPrivilegedContracts(market common.Address, parameterizer common.Address) (*types.Transaction, error) {
	return _Voting.Contract.SetPrivilegedContracts(&_Voting.TransactOpts, market, parameterizer)
}

// Vote is a paid mutator transaction binding the contract method 0xa69beaba.
//
// Solidity: function vote(hash bytes32) returns()
func (_Voting *VotingTransactor) Vote(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "vote", hash)
}

// Vote is a paid mutator transaction binding the contract method 0xa69beaba.
//
// Solidity: function vote(hash bytes32) returns()
func (_Voting *VotingSession) Vote(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, hash)
}

// Vote is a paid mutator transaction binding the contract method 0xa69beaba.
//
// Solidity: function vote(hash bytes32) returns()
func (_Voting *VotingTransactorSession) Vote(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, hash)
}

// VotingCandidateAddedEventIterator is returned from FilterCandidateAddedEvent and is used to iterate over the raw logs and unpacked data for CandidateAddedEvent events raised by the Voting contract.
type VotingCandidateAddedEventIterator struct {
	Event *VotingCandidateAddedEvent // Event containing the contract specifics and raw log

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
func (it *VotingCandidateAddedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingCandidateAddedEvent)
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
		it.Event = new(VotingCandidateAddedEvent)
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
func (it *VotingCandidateAddedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingCandidateAddedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingCandidateAddedEvent represents a CandidateAddedEvent event raised by the Voting contract.
type VotingCandidateAddedEvent struct {
	Hash   [32]byte
	Kind   uint8
	VoteBy *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCandidateAddedEvent is a free log retrieval operation binding the contract event 0xb7fdb794d49b8471a97fc8b1f90d1af97e75e0afc5f56d3b12b696e6d80f0186.
//
// Solidity: e CandidateAddedEvent(hash indexed bytes32, kind indexed uint8, voteBy indexed uint256)
func (_Voting *VotingFilterer) FilterCandidateAddedEvent(opts *bind.FilterOpts, hash [][32]byte, kind []uint8, voteBy []*big.Int) (*VotingCandidateAddedEventIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var voteByRule []interface{}
	for _, voteByItem := range voteBy {
		voteByRule = append(voteByRule, voteByItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CandidateAddedEvent", hashRule, kindRule, voteByRule)
	if err != nil {
		return nil, err
	}
	return &VotingCandidateAddedEventIterator{contract: _Voting.contract, event: "CandidateAddedEvent", logs: logs, sub: sub}, nil
}

// WatchCandidateAddedEvent is a free log subscription operation binding the contract event 0xb7fdb794d49b8471a97fc8b1f90d1af97e75e0afc5f56d3b12b696e6d80f0186.
//
// Solidity: e CandidateAddedEvent(hash indexed bytes32, kind indexed uint8, voteBy indexed uint256)
func (_Voting *VotingFilterer) WatchCandidateAddedEvent(opts *bind.WatchOpts, sink chan<- *VotingCandidateAddedEvent, hash [][32]byte, kind []uint8, voteBy []*big.Int) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var voteByRule []interface{}
	for _, voteByItem := range voteBy {
		voteByRule = append(voteByRule, voteByItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CandidateAddedEvent", hashRule, kindRule, voteByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingCandidateAddedEvent)
				if err := _Voting.contract.UnpackLog(event, "CandidateAddedEvent", log); err != nil {
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

// VotingCandidateRemovedEventIterator is returned from FilterCandidateRemovedEvent and is used to iterate over the raw logs and unpacked data for CandidateRemovedEvent events raised by the Voting contract.
type VotingCandidateRemovedEventIterator struct {
	Event *VotingCandidateRemovedEvent // Event containing the contract specifics and raw log

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
func (it *VotingCandidateRemovedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingCandidateRemovedEvent)
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
		it.Event = new(VotingCandidateRemovedEvent)
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
func (it *VotingCandidateRemovedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingCandidateRemovedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingCandidateRemovedEvent represents a CandidateRemovedEvent event raised by the Voting contract.
type VotingCandidateRemovedEvent struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCandidateRemovedEvent is a free log retrieval operation binding the contract event 0x53ce6706735b5cd3034bf68c2a532801c1efb4ba72410b5af2167a8462f15b7d.
//
// Solidity: e CandidateRemovedEvent(hash indexed bytes32)
func (_Voting *VotingFilterer) FilterCandidateRemovedEvent(opts *bind.FilterOpts, hash [][32]byte) (*VotingCandidateRemovedEventIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CandidateRemovedEvent", hashRule)
	if err != nil {
		return nil, err
	}
	return &VotingCandidateRemovedEventIterator{contract: _Voting.contract, event: "CandidateRemovedEvent", logs: logs, sub: sub}, nil
}

// WatchCandidateRemovedEvent is a free log subscription operation binding the contract event 0x53ce6706735b5cd3034bf68c2a532801c1efb4ba72410b5af2167a8462f15b7d.
//
// Solidity: e CandidateRemovedEvent(hash indexed bytes32)
func (_Voting *VotingFilterer) WatchCandidateRemovedEvent(opts *bind.WatchOpts, sink chan<- *VotingCandidateRemovedEvent, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CandidateRemovedEvent", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingCandidateRemovedEvent)
				if err := _Voting.contract.UnpackLog(event, "CandidateRemovedEvent", log); err != nil {
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

// VotingCouncilMemberAddedEventIterator is returned from FilterCouncilMemberAddedEvent and is used to iterate over the raw logs and unpacked data for CouncilMemberAddedEvent events raised by the Voting contract.
type VotingCouncilMemberAddedEventIterator struct {
	Event *VotingCouncilMemberAddedEvent // Event containing the contract specifics and raw log

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
func (it *VotingCouncilMemberAddedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingCouncilMemberAddedEvent)
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
		it.Event = new(VotingCouncilMemberAddedEvent)
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
func (it *VotingCouncilMemberAddedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingCouncilMemberAddedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingCouncilMemberAddedEvent represents a CouncilMemberAddedEvent event raised by the Voting contract.
type VotingCouncilMemberAddedEvent struct {
	Member common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCouncilMemberAddedEvent is a free log retrieval operation binding the contract event 0xf8897212eda7751d264aa9bac4715cbefe140b3124d688b928373f4c910cb829.
//
// Solidity: e CouncilMemberAddedEvent(member indexed address)
func (_Voting *VotingFilterer) FilterCouncilMemberAddedEvent(opts *bind.FilterOpts, member []common.Address) (*VotingCouncilMemberAddedEventIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CouncilMemberAddedEvent", memberRule)
	if err != nil {
		return nil, err
	}
	return &VotingCouncilMemberAddedEventIterator{contract: _Voting.contract, event: "CouncilMemberAddedEvent", logs: logs, sub: sub}, nil
}

// WatchCouncilMemberAddedEvent is a free log subscription operation binding the contract event 0xf8897212eda7751d264aa9bac4715cbefe140b3124d688b928373f4c910cb829.
//
// Solidity: e CouncilMemberAddedEvent(member indexed address)
func (_Voting *VotingFilterer) WatchCouncilMemberAddedEvent(opts *bind.WatchOpts, sink chan<- *VotingCouncilMemberAddedEvent, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CouncilMemberAddedEvent", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingCouncilMemberAddedEvent)
				if err := _Voting.contract.UnpackLog(event, "CouncilMemberAddedEvent", log); err != nil {
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

// VotingCouncilMemberRemovedEventIterator is returned from FilterCouncilMemberRemovedEvent and is used to iterate over the raw logs and unpacked data for CouncilMemberRemovedEvent events raised by the Voting contract.
type VotingCouncilMemberRemovedEventIterator struct {
	Event *VotingCouncilMemberRemovedEvent // Event containing the contract specifics and raw log

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
func (it *VotingCouncilMemberRemovedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingCouncilMemberRemovedEvent)
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
		it.Event = new(VotingCouncilMemberRemovedEvent)
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
func (it *VotingCouncilMemberRemovedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingCouncilMemberRemovedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingCouncilMemberRemovedEvent represents a CouncilMemberRemovedEvent event raised by the Voting contract.
type VotingCouncilMemberRemovedEvent struct {
	Member common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCouncilMemberRemovedEvent is a free log retrieval operation binding the contract event 0x564e70d59b50a6b32776c3d4a1e4e378a77dc9e24c3e2fbfa04a79166a67489b.
//
// Solidity: e CouncilMemberRemovedEvent(member indexed address)
func (_Voting *VotingFilterer) FilterCouncilMemberRemovedEvent(opts *bind.FilterOpts, member []common.Address) (*VotingCouncilMemberRemovedEventIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CouncilMemberRemovedEvent", memberRule)
	if err != nil {
		return nil, err
	}
	return &VotingCouncilMemberRemovedEventIterator{contract: _Voting.contract, event: "CouncilMemberRemovedEvent", logs: logs, sub: sub}, nil
}

// WatchCouncilMemberRemovedEvent is a free log subscription operation binding the contract event 0x564e70d59b50a6b32776c3d4a1e4e378a77dc9e24c3e2fbfa04a79166a67489b.
//
// Solidity: e CouncilMemberRemovedEvent(member indexed address)
func (_Voting *VotingFilterer) WatchCouncilMemberRemovedEvent(opts *bind.WatchOpts, sink chan<- *VotingCouncilMemberRemovedEvent, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CouncilMemberRemovedEvent", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingCouncilMemberRemovedEvent)
				if err := _Voting.contract.UnpackLog(event, "CouncilMemberRemovedEvent", log); err != nil {
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

// VotingVotedEventIterator is returned from FilterVotedEvent and is used to iterate over the raw logs and unpacked data for VotedEvent events raised by the Voting contract.
type VotingVotedEventIterator struct {
	Event *VotingVotedEvent // Event containing the contract specifics and raw log

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
func (it *VotingVotedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingVotedEvent)
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
		it.Event = new(VotingVotedEvent)
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
func (it *VotingVotedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingVotedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingVotedEvent represents a VotedEvent event raised by the Voting contract.
type VotingVotedEvent struct {
	Voter common.Address
	Hash  [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVotedEvent is a free log retrieval operation binding the contract event 0x81681507a95428bb04c5f5ac127e404caf6422d0dcf2ff77612b9201c3aa33bf.
//
// Solidity: e VotedEvent(voter indexed address, hash indexed bytes32)
func (_Voting *VotingFilterer) FilterVotedEvent(opts *bind.FilterOpts, voter []common.Address, hash [][32]byte) (*VotingVotedEventIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "VotedEvent", voterRule, hashRule)
	if err != nil {
		return nil, err
	}
	return &VotingVotedEventIterator{contract: _Voting.contract, event: "VotedEvent", logs: logs, sub: sub}, nil
}

// WatchVotedEvent is a free log subscription operation binding the contract event 0x81681507a95428bb04c5f5ac127e404caf6422d0dcf2ff77612b9201c3aa33bf.
//
// Solidity: e VotedEvent(voter indexed address, hash indexed bytes32)
func (_Voting *VotingFilterer) WatchVotedEvent(opts *bind.WatchOpts, sink chan<- *VotingVotedEvent, voter []common.Address, hash [][32]byte) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "VotedEvent", voterRule, hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingVotedEvent)
				if err := _Voting.contract.UnpackLog(event, "VotedEvent", log); err != nil {
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
