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
const MarketABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isListingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"exit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"resolveApplication\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offered\",\"type\":\"uint256\"}],\"name\":\"invest\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToListing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"apply\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInvestmentPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromListing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isInvestor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"challenge\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"resolveChallenge\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"isListed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"networkTokenAddr\",\"type\":\"address\"},{\"name\":\"marketTokenAddr\",\"type\":\"address\"},{\"name\":\"parameterizerAddr\",\"type\":\"address\"},{\"name\":\"votingAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"ApplicationFailedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"applicant\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"voteBy\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AppliedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"ChallengedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"ChallengeFailedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"ChallengeSucceededEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"offered\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"taken\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minted\",\"type\":\"uint256\"}],\"name\":\"InvestedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"ListedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"added\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"ListingDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"ListingDeletedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"added\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"ListingWithdrawEvent\",\"type\":\"event\"}]"

// MarketBin is the compiled bytecode used for deploying new contracts.
const MarketBin = `0x60806040523480156200001157600080fd5b50604051620037103803806200371083398101604090815281516020808401519284015160608501516080860151600c8054600160a060020a031916331790559390950180519095919391926200006e91600b91880190620000c4565b5060078054600160a060020a03958616600160a060020a0319918216179091556008805494861694821694909417909355600a805492851692841692909217909155600980549190931691161790555062000169565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200010757805160ff191683800117855562000137565b8280016001018555821562000137579182015b82811115620001375782518255916020019190600101906200011a565b506200014592915062000149565b5090565b6200016691905b8082111562000145576000815560010162000150565b90565b61359780620001796000396000f3006080604052600436106100ab5763ffffffff60e060020a60003504166302a47cd781146100b05780630ca36263146100e557806317d7de7c146100fd57806318e28e1c146101875780632afcf4801461019f578063341ff089146101c957806389210b8f146101e4578063a3d6602b14610202578063a7e7867514610217578063cee2a9cf14610234578063cffd46dc14610255578063d32c943a1461026d578063ecefbdc614610285575b600080fd5b3480156100bc57600080fd5b506100d1600160a060020a036004351661029d565b604080519115158252519081900360200190f35b3480156100f157600080fd5b506100d16004356102c0565b34801561010957600080fd5b5061011261057f565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561014c578181015183820152602001610134565b50505050905090810190601f1680156101795780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561019357600080fd5b506100d1600435610616565b3480156101ab57600080fd5b506101b7600435610e50565b60408051918252519081900360200190f35b3480156101d557600080fd5b506100d16004356024356114f0565b3480156101f057600080fd5b506100d1600435602435604435611703565b34801561020e57600080fd5b506101b7611cb0565b34801561022357600080fd5b50610232600435602435611e83565b005b34801561024057600080fd5b506100d1600160a060020a0360043516612173565b34801561026157600080fd5b506100d16004356121cb565b34801561027957600080fd5b506100d160043561288a565b34801561029157600080fd5b506100d1600435613102565b600160a060020a038116600090815260036020526040902054600111155b919050565b6000818152600160208190526040822001546101009004600160a060020a0316331461035c576040805160e560020a62461bcd02815260206004820152602960248201527f4572726f723a4d61726b65742e65786974202d204d757374206265206c69737460448201527f696e67206f776e65720000000000000000000000000000000000000000000000606482015290519081900360840190fd5b61036582613102565b15156103e1576040805160e560020a62461bcd02815260206004820152602260248201527f4572726f723a4d61726b65742e65786974202d204d757374206265206c69737460448201527f6564000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600954604080517fb89694c6000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a039092169163b89694c6916024808201926020929091908290030181600087803b15801561044857600080fd5b505af115801561045c573d6000803e3d6000fd5b505050506040513d602081101561047257600080fd5b5051156104ef576040805160e560020a62461bcd02815260206004820152603260248201527f4572726f723a4d61726b65742e65786974202d2043616e6e6f7420657869742060448201527f647572696e672061206368616c6c656e67650000000000000000000000000000606482015290519081900360840190fd5b6104f88261311b565b1515600114610577576040805160e560020a62461bcd02815260206004820152602c60248201527f4572726f723a4d61726b65742e65786974202d20436f756c64206e6f7420726560448201527f6d6f7665206c697374696e670000000000000000000000000000000000000000606482015290519081900360840190fd5b506001919050565b600b8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561060b5780601f106105e05761010080835404028352916020019161060b565b820191906000526020600020905b8154815290600101906020018083116105ee57829003601f168201915b505050505090505b90565b600954604080517f62f9a55d00000000000000000000000000000000000000000000000000000000815233600482015290516000928392600160a060020a03909116916362f9a55d9160248082019260209290919082900301818787803b15801561068057600080fd5b505af1158015610694573d6000803e3d6000fd5b505050506040513d60208110156106aa57600080fd5b50511515610716576040805160e560020a62461bcd02815260206004820152603f602482015260008051602061352c83398151915260448201527f2d2053656e646572206d75737420626520636f756e63696c206d656d62657200606482015290519081900360840190fd5b6000838152600160208190526040909120015460ff1615610795576040805160e560020a62461bcd028152602060048201526030602482015260008051602061352c83398151915260448201527f2d20416c7265616479206c697374656400000000000000000000000000000000606482015290519081900360840190fd5b600954604080517f11b883820000000000000000000000000000000000000000000000000000000081526004810186905260248101829052600b60448201527f6170706c69636174696f6e00000000000000000000000000000000000000000060648201529051600160a060020a03909216916311b88382916084808201926020929091908290030181600087803b15801561083057600080fd5b505af1158015610844573d6000803e3d6000fd5b505050506040513d602081101561085a57600080fd5b505115156108c6576040805160e560020a62461bcd028152602060048201526038602482015260008051602061352c83398151915260448201527f2d204d75737420626520616e206170706c69636174696f6e0000000000000000606482015290519081900360840190fd5b600954604080517f327322c8000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a039092169163327322c8916024808201926020929091908290030181600087803b15801561092d57600080fd5b505af1158015610941573d6000803e3d6000fd5b505050506040513d602081101561095757600080fd5b505115156109e9576040805160e560020a62461bcd028152602060048201526049602482015260008051602061352c83398151915260448201527f2d20506f6c6c7320666f7220746869732063616e646964617465206d7573742060648201527f626520636c6f7365640000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b600954600a54604080517fc26c12eb0000000000000000000000000000000000000000000000000000000081529051600160a060020a0393841693638f354b7993889391169163c26c12eb916004808201926020929091908290030181600087803b158015610a5757600080fd5b505af1158015610a6b573d6000803e3d6000fd5b505050506040513d6020811015610a8157600080fd5b50516040805160e060020a63ffffffff8616028152600481019390935260248301919091525160448083019260209291908290030181600087803b158015610ac857600080fd5b505af1158015610adc573d6000803e3d6000fd5b505050506040513d6020811015610af257600080fd5b505115610d0e5760008381526001602081815260408084208301805460ff1916909317909255600a5482517fdc2b28530000000000000000000000000000000000000000000000000000000081529251600160a060020a039091169363dc2b2853936004808201949392918390030190829087803b158015610b7357600080fd5b505af1158015610b87573d6000803e3d6000fd5b505050506040513d6020811015610b9d57600080fd5b5051600854604080517fa0712d68000000000000000000000000000000000000000000000000000000008152600481018490529051929350600160a060020a039091169163a0712d68916024808201926020929091908290030181600087803b158015610c0957600080fd5b505af1158015610c1d573d6000803e3d6000fd5b505050506040513d6020811015610c3357600080fd5b50511515610c9f576040805160e560020a62461bcd028152602060048201526038602482015260008051602061352c83398151915260448201527f2d20436f756c64206e6f74206d696e7420726577617264730000000000000000606482015290519081900360840190fd5b60008381526001602081815260409283902060048101859055918201546002909201548351858152918201528251610100909204600160a060020a03169286927fb08f341fa4f58287c76e79e69c47e07939983551f847f728f4dc9e06a49e308d9281900390910190a3610d84565b610d178361311b565b1515600114610d84576040805160e560020a62461bcd02815260206004820152603c602482015260008051602061352c83398151915260448201527f2d20436f756c64206e6f742072656d6f7665206170706c6963616e7400000000606482015290519081900360840190fd5b600954604080517f89bb617c000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a03909216916389bb617c916024808201926020929091908290030181600087803b158015610deb57600080fd5b505af1158015610dff573d6000803e3d6000fd5b505050506040513d6020811015610e1557600080fd5b50511515610e47576040805160e560020a62461bcd028152602060048201526000602482015290519081900360640190fd5b50600192915050565b600080600080600080610e623361029d565b15610f03576040805160e560020a62461bcd02815260206004820152604260248201527f4572726f723a4d61726b65742e696e76657374202d2043616e6e6f7420696e7660448201527f657374207768696c65206f776e696e6720616e20616374697665206c6973746960648201527f6e67000000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b610f0b611cb0565b945084871015610fb1576040805160e560020a62461bcd02815260206004820152605e60248201527f4572726f723a4d61726b65742e696e76657374202d20416d6f756e74206f666660448201527f65726564206d7573742062652067726561746572207468616e206f722065717560648201527f616c20746f2063757272656e7420696e766573746d656e742070726963650000608482015290519081900360a40190fd5b8487811515610fbc57fe5b069350831515610fcc5783610fdc565b610fdc878563ffffffff61348216565b6007546040805160e060020a6323b872dd028152336004820152306024820152604481018490529051929550600160a060020a03909116916323b872dd916064808201926020929091908290030181600087803b15801561103c57600080fd5b505af1158015611050573d6000803e3d6000fd5b505050506040513d602081101561106657600080fd5b506110799050838663ffffffff61349916565b600854604080517fa0712d68000000000000000000000000000000000000000000000000000000008152600481018490529051929450600160a060020a039091169163a0712d68916024808201926020929091908290030181600087803b1580156110e357600080fd5b505af11580156110f7573d6000803e3d6000fd5b505050506040513d602081101561110d57600080fd5b5051151561118b576040805160e560020a62461bcd02815260206004820152602b60248201527f4572726f723a4d61726b65742e696e76657374202d20436f756c64206e6f742060448201527f6d696e7420746f6b656e73000000000000000000000000000000000000000000606482015290519081900360840190fd5b600854604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018590529051600160a060020a039092169163a9059cbb916044808201926020929091908290030181600087803b1580156111f857600080fd5b505af115801561120c573d6000803e3d6000fd5b505050506040513d602081101561122257600080fd5b505115156112a0576040805160e560020a62461bcd02815260206004820152602f60248201527f4572726f723a4d61726b65742e696e76657374202d20436f756c64206e6f742060448201527f7472616e7366657220746f6b656e730000000000000000000000000000000000606482015290519081900360840190fd5b6112a933612173565b1561131457336000908152600460205260409020600101546112d1908463ffffffff6134ae16565b3360009081526004602052604090206001810191909155600201546112fc908363ffffffff6134ae16565b3360009081526004602052604090206002015561148d565b5033600081815260046020818152604080842060018082018990556002820188905560058054918201815586527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db08101805473ffffffffffffffffffffffffffffffffffffffff191688179055815560095482517f471bb3090000000000000000000000000000000000000000000000000000000081529485019690965290519094600160a060020a03169363471bb309936024808201949392918390030190829087803b1580156113e557600080fd5b505af11580156113f9573d6000803e3d6000fd5b505050506040513d602081101561140f57600080fd5b5051151561148d576040805160e560020a62461bcd02815260206004820152603760248201527f4572726f723a4d61726b65742e696e76657374202d20636f756c64206e6f742060448201527f61646420696e766573746f7220746f20636f756e63696c000000000000000000606482015290519081900360840190fd5b6006546114a0908463ffffffff6134ae16565b6006556040805188815260208101859052808201849052905133917fe2eb2b75e74d90e641625df497b148251bf493fadfa040f83f78c20a405fca0b919081900360600190a25095945050505050565b60008281526001602081905260408220908101546101009004600160a060020a0316331461158e576040805160e560020a62461bcd02815260206004820152602c60248201527f4572726f723a4d61726b65742e6465706f736974202d204d757374206265206c60448201527f697374696e67206f776e65720000000000000000000000000000000000000000606482015290519081900360840190fd5b6008546040805160e060020a6323b872dd028152336004820152306024820152604481018690529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b1580156115eb57600080fd5b505af11580156115ff573d6000803e3d6000fd5b505050506040513d602081101561161557600080fd5b50511515611693576040805160e560020a62461bcd02815260206004820152603760248201527f4572726f723a4d61726b65742e6465706f736974202d20436f756c64206e6f7460448201527f207472616e73666572204d61726b657420546f6b656e73000000000000000000606482015290519081900360840190fd5b60028101546116a8908463ffffffff6134ae16565b6002820181905560048201546040805186815260208101939093528281019190915251339186917f7f1c6a478f1fbf07f8b135e162e7322085eb5371d2b77163e27dc39d4154d1369181900360600190a35060019392505050565b600954604080517fb89694c600000000000000000000000000000000000000000000000000000000815260048101869052905160009283928392600160a060020a039092169163b89694c69160248082019260209290919082900301818787803b15801561177057600080fd5b505af1158015611784573d6000803e3d6000fd5b505050506040513d602081101561179a57600080fd5b505115611817576040805160e560020a62461bcd02815260206004820152602860248201527f4572726f723a4d61726b65742e6170706c79202d20416c72656164792061206360448201527f616e646964617465000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b61182086613102565b1561189b576040805160e560020a62461bcd02815260206004820152602360248201527f4572726f723a4d61726b65742e6170706c79202d20416c7265616479206c697360448201527f7465640000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008681526001602081905260408220908101805474ffffffffffffffffffffffffffffffffffffffff00191633610100021790556003810187905592508411156119f8576002820184905560085460018301546040805160e060020a6323b872dd028152610100909204600160a060020a0390811660048401523060248401526044830188905290519216916323b872dd916064808201926020929091908290030181600087803b15801561195057600080fd5b505af1158015611964573d6000803e3d6000fd5b505050506040513d602081101561197a57600080fd5b505115156119f8576040805160e560020a62461bcd02815260206004820152603560248201527f4572726f723a4d61726b65742e6170706c79202d20436f756c64206e6f74207460448201527f72616e73666572204d61726b657420546f6b656e730000000000000000000000606482015290519081900360840190fd5b600954600a54604080517f2d0d2bc60000000000000000000000000000000000000000000000000000000081529051600160a060020a039384169363364f46a4938b93911691632d0d2bc6916004808201926020929091908290030181600087803b158015611a6657600080fd5b505af1158015611a7a573d6000803e3d6000fd5b505050506040513d6020811015611a9057600080fd5b50516040805160e060020a63ffffffff86160281526024810193909352604483019190915260606004830152600b60648301527f6170706c69636174696f6e00000000000000000000000000000000000000000060848301525160a48083019260209291908290030181600087803b158015611b0b57600080fd5b505af1158015611b1f573d6000803e3d6000fd5b505050506040513d6020811015611b3557600080fd5b50519050801515611bdc576040805160e560020a62461bcd02815260206004820152604660248201527f4572726f723a4d61726b65742e6170706c79202d20436f756c64206e6f74206160448201527f6464206170706c6963616e7420746f20766f74696e672063616e64696461746560648201527f73206c6973740000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b600a54604080517f2d0d2bc60000000000000000000000000000000000000000000000000000000081529051879233928a927f06198c59f70529156e65d8211777fdd8fdd009977545a9e8c44050f7ff2cb13392600160a060020a031691632d0d2bc69160048083019260209291908290030181600087803b158015611c6157600080fd5b505af1158015611c75573d6000803e3d6000fd5b505050506040513d6020811015611c8b57600080fd5b505160408051918252602082018a90528051918290030190a450600195945050505050565b600080600080600a60009054906101000a9004600160a060020a0316600160a060020a031663f36089ec6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015611d0957600080fd5b505af1158015611d1d573d6000803e3d6000fd5b505050506040513d6020811015611d3357600080fd5b5051600a54604080517ffefb9dde0000000000000000000000000000000000000000000000000000000081529051929550600160a060020a039091169163fefb9dde916004808201926020929091908290030181600087803b158015611d9857600080fd5b505af1158015611dac573d6000803e3d6000fd5b505050506040513d6020811015611dc257600080fd5b5051600754604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051929450600160a060020a03909116916370a08231916024808201926020929091908290030181600087803b158015611e2d57600080fd5b505af1158015611e41573d6000803e3d6000fd5b505050506040513d6020811015611e5757600080fd5b50519050611e7b611e6e838363ffffffff6134bb16565b849063ffffffff6134ae16565b935050505090565b6000828152600160208190526040909120908101546101009004600160a060020a03163314611f22576040805160e560020a62461bcd02815260206004820152602d60248201527f4572726f723a4d61726b65742e7769746864726177202d204d7573742062652060448201527f6c697374696e67206f776e657200000000000000000000000000000000000000606482015290519081900360840190fd5b6002810154821115611fca576040805160e560020a62461bcd02815260206004820152604f60248201527f4572726f723a4d61726b65742e7769746864726177202d20616d6f756e74206d60448201527f757374206265206c657373207468616e206f7220657175616c20746f2074686560648201527f206c697374696e6720737570706c790000000000000000000000000000000000608482015290519081900360a40190fd5b6002810154611fdf908363ffffffff61348216565b6002820155600854604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018590529051600160a060020a039092169163a9059cbb916044808201926020929091908290030181600087803b15801561205157600080fd5b505af1158015612065573d6000803e3d6000fd5b505050506040513d602081101561207b57600080fd5b5051151561211f576040805160e560020a62461bcd02815260206004820152604260248201527f4572726f723a4d61726b65742e7769746864726177202d20436f756c64206e6f60448201527f74207472616e7366657220746f6b656e7320746f206c697374696e67206f776e60648201527f6572000000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b600281015460048201546040805185815260208101939093528281019190915251339185917f261e65ee6edcc0072cc284d3d0665cbc0c581ab1fc708222b6bc3d20c452750d9181900360600190a3505050565b6005546000901515612187575060006102bb565b600160a060020a0382166000818152600460205260409020546005805490919081106121af57fe5b600091825260209091200154600160a060020a03161492915050565b6000806000806000806000600a60009054906101000a9004600160a060020a0316600160a060020a0316635aebe7686040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561222957600080fd5b505af115801561223d573d6000803e3d6000fd5b505050506040513d602081101561225357600080fd5b5051600089815260016020819052604090912080820154929850965060ff9091161515146122f1576040805160e560020a62461bcd02815260206004820152603b60248201527f4572726f723a4d61726b65742e6368616c6c656e6765202d204d75737420626560448201527f2061206c697374696e6720746f206265206368616c6c656e6765640000000000606482015290519081900360840190fd5b600954604080517fb89694c6000000000000000000000000000000000000000000000000000000008152600481018b90529051600160a060020a039092169163b89694c6916024808201926020929091908290030181600087803b15801561235857600080fd5b505af115801561236c573d6000803e3d6000fd5b505050506040513d602081101561238257600080fd5b5051156123ff576040805160e560020a62461bcd02815260206004820152602b60248201527f4572726f723a4d61726b65742e6368616c6c656e6765202d20416c726561647960448201527f206368616c6c656e676564000000000000000000000000000000000000000000606482015290519081900360840190fd5b8561241b866004015487600201546134ae90919063ffffffff16565b10156124ae5761242a8861311b565b15156001146124a9576040805160e560020a62461bcd02815260206004820152603660248201527f4572726f723a4d61726b65742e6368616c6c656e6765202d20436f756c64206e60448201527f6f74206175746f2072656d6f7665206c697374696e6700000000000000000000606482015290519081900360840190fd5b61287c565b6008546040805160e060020a6323b872dd028152336004820152306024820152604481018990529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b15801561250b57600080fd5b505af115801561251f573d6000803e3d6000fd5b505050506040513d602081101561253557600080fd5b505115156125b3576040805160e560020a62461bcd02815260206004820152603960248201527f4572726f723a4d61726b65742e6368616c6c656e6765202d20436f756c64206e60448201527f6f74207472616e73666572204d61726b657420746f6b656e7300000000000000606482015290519081900360840190fd5b600285015486116125e05760028501546125d3908763ffffffff61348216565b6002860155859350612613565b600285018054600090915560048601549094508487039250612608908363ffffffff61348216565b600486015590915081905b60408051606081018252338152602080820187815282840187815260008d81528084528581209451855473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03918216178655925160018601559051600290940193909355600954600a5485517f2d0d2bc600000000000000000000000000000000000000000000000000000000815295519183169563364f46a4958f959290941693632d0d2bc693600480840194939192918390030190829087803b1580156126d857600080fd5b505af11580156126ec573d6000803e3d6000fd5b505050506040513d602081101561270257600080fd5b50516040805160e060020a63ffffffff86160281526024810193909352604483019190915260606004830152600960648301527f6368616c6c656e6765000000000000000000000000000000000000000000000060848301525160a48083019260209291908290030181600087803b15801561277d57600080fd5b505af1158015612791573d6000803e3d6000fd5b505050506040513d60208110156127a757600080fd5b5051905080151561284e576040805160e560020a62461bcd02815260206004820152604a60248201527f4572726f723a4d61726b65742e6368616c6c656e6765202d20436f756c64206e60448201527f6f7420616464206368616c6c656e676520746f20766f74696e672063616e646960648201527f6461746573206c69737400000000000000000000000000000000000000000000608482015290519081900360a40190fd5b604051339089907f4dd89117b2f54a955a15b7ca70efe9d9cda23696ebd2a8290ce6bb42b1594dba90600090a35b506001979650505050505050565b600954604080517f62f9a55d00000000000000000000000000000000000000000000000000000000815233600482015290516000928392600160a060020a03909116916362f9a55d9160248082019260209290919082900301818787803b1580156128f457600080fd5b505af1158015612908573d6000803e3d6000fd5b505050506040513d602081101561291e57600080fd5b5051151561298a576040805160e560020a62461bcd02815260206004820152603d602482015260008051602061354c83398151915260448201527f53656e646572206d75737420626520636f756e63696c206d656d626572000000606482015290519081900360840190fd5b60098054604080517f11b88382000000000000000000000000000000000000000000000000000000008152600481018790526024810182905260448101939093527f6368616c6c656e67650000000000000000000000000000000000000000000000606484015251600160a060020a03909116916311b883829160848083019260209291908290030181600087803b158015612a2557600080fd5b505af1158015612a39573d6000803e3d6000fd5b505050506040513d6020811015612a4f57600080fd5b50511515612abb576040805160e560020a62461bcd028152602060048201526033602482015260008051602061354c83398151915260448201527f4d7573742062652061206368616c6c656e676500000000000000000000000000606482015290519081900360840190fd5b600954604080517f327322c8000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a039092169163327322c8916024808201926020929091908290030181600087803b158015612b2257600080fd5b505af1158015612b36573d6000803e3d6000fd5b505050506040513d6020811015612b4c57600080fd5b50511515612bde576040805160e560020a62461bcd028152602060048201526047602482015260008051602061354c83398151915260448201527f506f6c6c7320666f7220746869732063616e646964617465206d75737420626560648201527f20636c6f73656400000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b600a60009054906101000a9004600160a060020a0316600160a060020a0316635aebe7686040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612c3157600080fd5b505af1158015612c45573d6000803e3d6000fd5b505050506040513d6020811015612c5b57600080fd5b5051600954600a54604080517fc26c12eb0000000000000000000000000000000000000000000000000000000081529051939450600160a060020a0392831693638f354b79938893169163c26c12eb9160048083019260209291908290030181600087803b158015612ccc57600080fd5b505af1158015612ce0573d6000803e3d6000fd5b505050506040513d6020811015612cf657600080fd5b50516040805160e060020a63ffffffff8616028152600481019390935260248301919091525160448083019260209291908290030181600087803b158015612d3d57600080fd5b505af1158015612d51573d6000803e3d6000fd5b505050506040513d6020811015612d6757600080fd5b505115612efc57612d778361311b565b50600854600084815260208190526040902054600160a060020a039182169163a9059cbb9116612dae84600263ffffffff6134bb16565b6040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015612e0057600080fd5b505af1158015612e14573d6000803e3d6000fd5b505050506040513d6020811015612e2a57600080fd5b50511515612ea8576040805160e560020a62461bcd02815260206004820152603e60248201527f4572726f723a6d61726b65742e7265736f6c76654368616c6c656e6765202d2060448201527f436f756c64206e6f74207472616e73666572204d61726b6574546f6b656e0000606482015290519081900360840190fd5b600083815260208181526040918290205482518481529251600160a060020a039091169286927f26adc8b945cdd9b054615b9f327f3945c817f77568b5fc0b7adaecd3fa5c149b92918290030190a3612fd2565b600083815260208181526040808320600190810154925290912060020154612f299163ffffffff6134ae16565b6000848152600160208181526040808420600280820196909655848352932090930154925260040154612f73918391612f679163ffffffff6134ae16565b9063ffffffff6134ae16565b600084815260016020908152604080832060040193909355818152908290205482518481529251600160a060020a039091169286927f303714bcbac27ef8d90dd5bce88c73b089d8c012ea32c762355112f66320c5fe92918290030190a35b600083815260208181526040808320805473ffffffffffffffffffffffffffffffffffffffff191681556001810184905560020183905560095481517f89bb617c000000000000000000000000000000000000000000000000000000008152600481018890529151600160a060020a03909116936389bb617c93602480850194919392918390030190829087803b15801561306c57600080fd5b505af1158015613080573d6000803e3d6000fd5b505050506040513d602081101561309657600080fd5b50511515610e47576040805160e560020a62461bcd02815260206004820152603a602482015260008051602061354c83398151915260448201527f436f756c64206e6f742072656d6f76652063616e646964617465000000000000606482015290519081900360840190fd5b6000908152600160208190526040909120015460ff1690565b600081815260016020819052604082208101548291829160ff161515146131b2576040805160e560020a62461bcd02815260206004820152602e60248201527f4572726f723a4d61726b65742e72656d6f76654c697374696e67202d204d757360448201527f742062652061206c697374696e67000000000000000000000000000000000000606482015290519081900360840190fd5b60008481526001602052604081206002015411156132fa5760085460008581526001602081815260408084209283015460029093015481517fa9059cbb000000000000000000000000000000000000000000000000000000008152610100909404600160a060020a0390811660048601526024850191909152905194169363a9059cbb936044808501948390030190829087803b15801561325257600080fd5b505af1158015613266573d6000803e3d6000fd5b505050506040513d602081101561327c57600080fd5b505115156132fa576040805160e560020a62461bcd02815260206004820152603e60248201527f4572726f723a4d61726b65742e72656d6f76654c697374696e67202d20436f7560448201527f6c64206e6f7420726566756e642072656d61696e696e6720737570706c790000606482015290519081900360840190fd5b600084815260016020526040812060040154111561339e5760085460008581526001602052604080822060049081015482517f42966c68000000000000000000000000000000000000000000000000000000008152918201529051600160a060020a03909316926342966c689260248084019391929182900301818387803b15801561338557600080fd5b505af1158015613399573d6000803e3d6000fd5b505050505b600084815260016020526040902054600280549193509060001981019081106133c357fe5b90600052602060002001549050806002838154811015156133e057fe5b60009182526020909120015560028054906133ff9060001983016134e4565b506000818152600160208190526040808320859055868352808320838155918201805474ffffffffffffffffffffffffffffffffffffffffff19169055600282018390556003820183905560049091018290555185917fbf8f571a7fa2903765660aa54ed716098a6b456db5be6fba8e17bffbce48c85e91a25060019392505050565b60008282111561348e57fe5b508082035b92915050565b600081838115156134a657fe5b049392505050565b8181018281101561349357fe5b60008215156134cc57506000613493565b508181028183828115156134dc57fe5b041461349357fe5b8154818355818111156135085760008381526020902061350891810190830161350d565b505050565b61061391905b808211156135275760008155600101613513565b509056004572726f723a4d61726b65742e7265736f6c76654170706c69636174696f6e204572726f723a4d61726b65742e7265736f6c76654368616c6c656e6765202d20a165627a7a72305820e2ddbd1d0d6db89e59596220eca48a2daee73e94fb9045615f63efcc9abac21e0029`

// DeployMarket deploys a new Ethereum contract, binding an instance of Market to it.
func DeployMarket(auth *bind.TransactOpts, backend bind.ContractBackend, name string, networkTokenAddr common.Address, marketTokenAddr common.Address, parameterizerAddr common.Address, votingAddr common.Address) (common.Address, *types.Transaction, *Market, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarketBin), backend, name, networkTokenAddr, marketTokenAddr, parameterizerAddr, votingAddr)
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

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_Market *MarketCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_Market *MarketSession) GetName() (string, error) {
	return _Market.Contract.GetName(&_Market.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_Market *MarketCallerSession) GetName() (string, error) {
	return _Market.Contract.GetName(&_Market.CallOpts)
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

// Apply is a paid mutator transaction binding the contract method 0x89210b8f.
//
// Solidity: function apply(listingHash bytes32, dataHash bytes32, amount uint256) returns(bool)
func (_Market *MarketTransactor) Apply(opts *bind.TransactOpts, listingHash [32]byte, dataHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "apply", listingHash, dataHash, amount)
}

// Apply is a paid mutator transaction binding the contract method 0x89210b8f.
//
// Solidity: function apply(listingHash bytes32, dataHash bytes32, amount uint256) returns(bool)
func (_Market *MarketSession) Apply(listingHash [32]byte, dataHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Apply(&_Market.TransactOpts, listingHash, dataHash, amount)
}

// Apply is a paid mutator transaction binding the contract method 0x89210b8f.
//
// Solidity: function apply(listingHash bytes32, dataHash bytes32, amount uint256) returns(bool)
func (_Market *MarketTransactorSession) Apply(listingHash [32]byte, dataHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Apply(&_Market.TransactOpts, listingHash, dataHash, amount)
}

// Challenge is a paid mutator transaction binding the contract method 0xcffd46dc.
//
// Solidity: function challenge(listingHash bytes32) returns(bool)
func (_Market *MarketTransactor) Challenge(opts *bind.TransactOpts, listingHash [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "challenge", listingHash)
}

// Challenge is a paid mutator transaction binding the contract method 0xcffd46dc.
//
// Solidity: function challenge(listingHash bytes32) returns(bool)
func (_Market *MarketSession) Challenge(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Challenge(&_Market.TransactOpts, listingHash)
}

// Challenge is a paid mutator transaction binding the contract method 0xcffd46dc.
//
// Solidity: function challenge(listingHash bytes32) returns(bool)
func (_Market *MarketTransactorSession) Challenge(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Challenge(&_Market.TransactOpts, listingHash)
}

// DepositToListing is a paid mutator transaction binding the contract method 0x341ff089.
//
// Solidity: function depositToListing(listingHash bytes32, amount uint256) returns(bool)
func (_Market *MarketTransactor) DepositToListing(opts *bind.TransactOpts, listingHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "depositToListing", listingHash, amount)
}

// DepositToListing is a paid mutator transaction binding the contract method 0x341ff089.
//
// Solidity: function depositToListing(listingHash bytes32, amount uint256) returns(bool)
func (_Market *MarketSession) DepositToListing(listingHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.DepositToListing(&_Market.TransactOpts, listingHash, amount)
}

// DepositToListing is a paid mutator transaction binding the contract method 0x341ff089.
//
// Solidity: function depositToListing(listingHash bytes32, amount uint256) returns(bool)
func (_Market *MarketTransactorSession) DepositToListing(listingHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.DepositToListing(&_Market.TransactOpts, listingHash, amount)
}

// Exit is a paid mutator transaction binding the contract method 0x0ca36263.
//
// Solidity: function exit(listingHash bytes32) returns(bool)
func (_Market *MarketTransactor) Exit(opts *bind.TransactOpts, listingHash [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "exit", listingHash)
}

// Exit is a paid mutator transaction binding the contract method 0x0ca36263.
//
// Solidity: function exit(listingHash bytes32) returns(bool)
func (_Market *MarketSession) Exit(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Exit(&_Market.TransactOpts, listingHash)
}

// Exit is a paid mutator transaction binding the contract method 0x0ca36263.
//
// Solidity: function exit(listingHash bytes32) returns(bool)
func (_Market *MarketTransactorSession) Exit(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Exit(&_Market.TransactOpts, listingHash)
}

// Invest is a paid mutator transaction binding the contract method 0x2afcf480.
//
// Solidity: function invest(offered uint256) returns(uint256)
func (_Market *MarketTransactor) Invest(opts *bind.TransactOpts, offered *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "invest", offered)
}

// Invest is a paid mutator transaction binding the contract method 0x2afcf480.
//
// Solidity: function invest(offered uint256) returns(uint256)
func (_Market *MarketSession) Invest(offered *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Invest(&_Market.TransactOpts, offered)
}

// Invest is a paid mutator transaction binding the contract method 0x2afcf480.
//
// Solidity: function invest(offered uint256) returns(uint256)
func (_Market *MarketTransactorSession) Invest(offered *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Invest(&_Market.TransactOpts, offered)
}

// ResolveApplication is a paid mutator transaction binding the contract method 0x18e28e1c.
//
// Solidity: function resolveApplication(listingHash bytes32) returns(bool)
func (_Market *MarketTransactor) ResolveApplication(opts *bind.TransactOpts, listingHash [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "resolveApplication", listingHash)
}

// ResolveApplication is a paid mutator transaction binding the contract method 0x18e28e1c.
//
// Solidity: function resolveApplication(listingHash bytes32) returns(bool)
func (_Market *MarketSession) ResolveApplication(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.ResolveApplication(&_Market.TransactOpts, listingHash)
}

// ResolveApplication is a paid mutator transaction binding the contract method 0x18e28e1c.
//
// Solidity: function resolveApplication(listingHash bytes32) returns(bool)
func (_Market *MarketTransactorSession) ResolveApplication(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.ResolveApplication(&_Market.TransactOpts, listingHash)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0xd32c943a.
//
// Solidity: function resolveChallenge(listingHash bytes32) returns(bool)
func (_Market *MarketTransactor) ResolveChallenge(opts *bind.TransactOpts, listingHash [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "resolveChallenge", listingHash)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0xd32c943a.
//
// Solidity: function resolveChallenge(listingHash bytes32) returns(bool)
func (_Market *MarketSession) ResolveChallenge(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.ResolveChallenge(&_Market.TransactOpts, listingHash)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0xd32c943a.
//
// Solidity: function resolveChallenge(listingHash bytes32) returns(bool)
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
	Taken    *big.Int
	Minted   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestedEvent is a free log retrieval operation binding the contract event 0xe2eb2b75e74d90e641625df497b148251bf493fadfa040f83f78c20a405fca0b.
//
// Solidity: e InvestedEvent(investor indexed address, offered uint256, taken uint256, minted uint256)
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

// WatchInvestedEvent is a free log subscription operation binding the contract event 0xe2eb2b75e74d90e641625df497b148251bf493fadfa040f83f78c20a405fca0b.
//
// Solidity: e InvestedEvent(investor indexed address, offered uint256, taken uint256, minted uint256)
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
const MarketTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stopMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedAmount\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"market\",\"type\":\"address\"}],\"name\":\"setPrivilegedContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedAmount\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintingStopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialAccount\",\"type\":\"address\"},{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ApprovalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurnEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintStoppedEvent\",\"type\":\"event\"}]"

// MarketTokenBin is the compiled bytecode used for deploying new contracts.
const MarketTokenBin = `0x60806040526001805460ff1916905534801561001a57600080fd5b50604051604080611084833981016040908152815160209283015160018054610100330261010060a860020a0319909116179055600160a060020a0390911660009081526003909352908220819055905561100a8061007a6000396000f3006080604052600436106100c45763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b381146100c957806318160ddd1461010157806323b872dd146101285780633e3e0b121461015257806342966c6814610167578063661884631461018157806370a08231146101a5578063730e6e84146101c6578063a0712d68146101e7578063a9059cbb146101ff578063d73dd62314610223578063dd62ed3e14610247578063f339292f1461026e575b600080fd5b3480156100d557600080fd5b506100ed600160a060020a0360043516602435610283565b604080519115158252519081900360200190f35b34801561010d57600080fd5b506101166102e9565b60408051918252519081900360200190f35b34801561013457600080fd5b506100ed600160a060020a03600435811690602435166044356102ef565b34801561015e57600080fd5b506100ed6105a9565b34801561017357600080fd5b5061017f600435610716565b005b34801561018d57600080fd5b506100ed600160a060020a03600435166024356108f7565b3480156101b157600080fd5b50610116600160a060020a03600435166109e7565b3480156101d257600080fd5b506100ed600160a060020a0360043516610a02565b3480156101f357600080fd5b506100ed600435610b71565b34801561020b57600080fd5b506100ed600160a060020a0360043516602435610d19565b34801561022f57600080fd5b506100ed600160a060020a0360043516602435610ecc565b34801561025357600080fd5b50610116600160a060020a0360043581169060243516610f65565b34801561027a57600080fd5b506100ed610f90565b336000818152600460209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e928290030190a350600192915050565b60005490565b6000600160a060020a0383161515610377576040805160e560020a62461bcd02815260206004820152603c60248201527f4572726f723a5374616e646172642e7472616e7366657246726f6d202d20277460448201527f6f272061646472657373206d7573742062652073706563696669656400000000606482015290519081900360840190fd5b600160a060020a03841660009081526003602052604090205482111561040d576040805160e560020a62461bcd02815260206004820152603e60248201527f4572726f723a5374616e646172642e7472616e7366657246726f6d202d20416d60448201527f6f756e74206578636565647320617661696c61626c652062616c616e63650000606482015290519081900360840190fd5b600160a060020a03841660009081526004602090815260408083203384529091529020548211156104ae576040805160e560020a62461bcd02815260206004820152603b60248201527f4572726f722e5374616e646172642e7472616e7366657246726f6d202d20416d60448201527f6f756e74206578636565647320616c6c6f77656420616d6f756e740000000000606482015290519081900360840190fd5b600160a060020a0384166000908152600360205260409020546104d7908363ffffffff610f9916565b600160a060020a03808616600090815260036020526040808220939093559085168152205461050c908363ffffffff610fab16565b600160a060020a038085166000908152600360209081526040808320949094559187168152600482528281203382529091522054610550908363ffffffff610f9916565b600160a060020a0380861660008181526004602090815260408083203384528252918290209490945580518681529051928716939192600080516020610fbf833981519152929181900390910190a35060019392505050565b600254600090600160a060020a0316331461065a576040805160e560020a62461bcd02815260206004820152604660248201527f4572726f723a4d61726b6574546f6b656e2e68617350726976696c656765202d60448201527f2043616c6c6572206d75737420626520612070726976696c656765642020636f60648201527f6e74726163740000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b60015460ff16156106db576040805160e560020a62461bcd02815260206004820152603460248201527f4572726f723a4d61726b6574546f6b656e2e63616e4d696e74202d204d696e7460448201527f696e6720686173206265656e2073746f70706564000000000000000000000000606482015290519081900360840190fd5b6001805460ff1916811790556040517f6ae65cb1ae147362b8906c28687a17e1712ca926bea50710ba279849cec3347190600090a150600190565b600254600160a060020a031633146107c4576040805160e560020a62461bcd02815260206004820152604660248201527f4572726f723a4d61726b6574546f6b656e2e68617350726976696c656765202d60448201527f2043616c6c6572206d75737420626520612070726976696c656765642020636f60648201527f6e74726163740000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b33600090815260036020526040902054811115610851576040805160e560020a62461bcd02815260206004820152603a60248201527f4572726f723a4e6574776f726b546f6b656e2e6275726e202d20416d6f756e7460448201527f206578636565647320617661696c61626c652062616c616e6365000000000000606482015290519081900360840190fd5b33600090815260036020526040902054610871908263ffffffff610f9916565b3360009081526003602052604081209190915554610895908263ffffffff610f9916565b60005560408051828152905133917f512586160ebd4dc6945ba9ec5d21a1f723f26f3c7aa36cdffb6818d4e7b88030919081900360200190a26040805182815290516000913391600080516020610fbf8339815191529181900360200190a350565b336000908152600460209081526040808320600160a060020a03861684529091528120548083111561094c57336000908152600460209081526040808320600160a060020a0388168452909152812055610981565b61095c818463ffffffff610f9916565b336000908152600460209081526040808320600160a060020a03891684529091529020555b336000818152600460209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e929181900390910190a35060019392505050565b600160a060020a031660009081526003602052604090205490565b6001546000906101009004600160a060020a03163314610a92576040805160e560020a62461bcd02815260206004820152603060248201527f4572726f723a4d61726b6574546f6b656e2e69734f776e6572202d2043616c6c60448201527f6572206d757374206265206f776e657200000000000000000000000000000000606482015290519081900360840190fd5b600254600160a060020a031615610b3f576040805160e560020a62461bcd02815260206004820152604560248201527f4572726f723a4d61726b6574546f6b656e2e73657450726976696c656765644360448201527f6f6e747261637473202d204d61726b6574206164647265737320616c7265616460648201527f7920736574000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b5060028054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600254600090600160a060020a03163314610c22576040805160e560020a62461bcd02815260206004820152604660248201527f4572726f723a4d61726b6574546f6b656e2e68617350726976696c656765202d60448201527f2043616c6c6572206d75737420626520612070726976696c656765642020636f60648201527f6e74726163740000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b60015460ff1615610ca3576040805160e560020a62461bcd02815260206004820152603460248201527f4572726f723a4d61726b6574546f6b656e2e63616e4d696e74202d204d696e7460448201527f696e6720686173206265656e2073746f70706564000000000000000000000000606482015290519081900360840190fd5b600054610cb6908363ffffffff610fab16565b600090815533815260036020526040902054610cd8908363ffffffff610fab16565b336000818152600360209081526040808320949094558351868152935192939192600080516020610fbf8339815191529281900390910190a3506001919050565b6000600160a060020a0383161515610da1576040805160e560020a62461bcd02815260206004820152603360248201527f4572726f723a42617369632e7472616e73666572202d20416e2061646472657360448201527f73206d7573742062652073706563696669656400000000000000000000000000606482015290519081900360840190fd5b33600090815260036020526040902054821115610e2e576040805160e560020a62461bcd02815260206004820152603f60248201527f4572726f723a42617369632e7472616e73666572202d20416d6f756e7420657860448201527f6365656473207468652062616c616e6365206f66206d73672e73656e64657200606482015290519081900360840190fd5b33600090815260036020526040902054610e4e908363ffffffff610f9916565b3360009081526003602052604080822092909255600160a060020a03851681522054610e80908363ffffffff610fab16565b600160a060020a038416600081815260036020908152604091829020939093558051858152905191923392600080516020610fbf8339815191529281900390910190a350600192915050565b336000908152600460209081526040808320600160a060020a0386168452909152812054610f00908363ffffffff610fab16565b336000818152600460209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e929081900390910190a350600192915050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205490565b60015460ff1690565b600082821115610fa557fe5b50900390565b81810182811015610fb857fe5b929150505600eaf1c4b3ce0f4f62a2bae7eb3e68225c75f7e6ff4422073b7437b9a78d25f170a165627a7a72305820f130493a4193cd84842dc04fb412d72ffc3dd7739bb1aff4642c4df7d0dcb9020029`

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

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, subtractedAmount uint256) returns(bool)
func (_MarketToken *MarketTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedAmount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "decreaseApproval", spender, subtractedAmount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, subtractedAmount uint256) returns(bool)
func (_MarketToken *MarketTokenSession) DecreaseApproval(spender common.Address, subtractedAmount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.DecreaseApproval(&_MarketToken.TransactOpts, spender, subtractedAmount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, subtractedAmount uint256) returns(bool)
func (_MarketToken *MarketTokenTransactorSession) DecreaseApproval(spender common.Address, subtractedAmount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.DecreaseApproval(&_MarketToken.TransactOpts, spender, subtractedAmount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, addedAmount uint256) returns(bool)
func (_MarketToken *MarketTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedAmount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "increaseApproval", spender, addedAmount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, addedAmount uint256) returns(bool)
func (_MarketToken *MarketTokenSession) IncreaseApproval(spender common.Address, addedAmount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.IncreaseApproval(&_MarketToken.TransactOpts, spender, addedAmount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, addedAmount uint256) returns(bool)
func (_MarketToken *MarketTokenTransactorSession) IncreaseApproval(spender common.Address, addedAmount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.IncreaseApproval(&_MarketToken.TransactOpts, spender, addedAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(amount uint256) returns(bool)
func (_MarketToken *MarketTokenTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(amount uint256) returns(bool)
func (_MarketToken *MarketTokenSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Mint(&_MarketToken.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(amount uint256) returns(bool)
func (_MarketToken *MarketTokenTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Mint(&_MarketToken.TransactOpts, amount)
}

// SetPrivilegedContracts is a paid mutator transaction binding the contract method 0x730e6e84.
//
// Solidity: function setPrivilegedContracts(market address) returns(bool)
func (_MarketToken *MarketTokenTransactor) SetPrivilegedContracts(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "setPrivilegedContracts", market)
}

// SetPrivilegedContracts is a paid mutator transaction binding the contract method 0x730e6e84.
//
// Solidity: function setPrivilegedContracts(market address) returns(bool)
func (_MarketToken *MarketTokenSession) SetPrivilegedContracts(market common.Address) (*types.Transaction, error) {
	return _MarketToken.Contract.SetPrivilegedContracts(&_MarketToken.TransactOpts, market)
}

// SetPrivilegedContracts is a paid mutator transaction binding the contract method 0x730e6e84.
//
// Solidity: function setPrivilegedContracts(market address) returns(bool)
func (_MarketToken *MarketTokenTransactorSession) SetPrivilegedContracts(market common.Address) (*types.Transaction, error) {
	return _MarketToken.Contract.SetPrivilegedContracts(&_MarketToken.TransactOpts, market)
}

// StopMinting is a paid mutator transaction binding the contract method 0x3e3e0b12.
//
// Solidity: function stopMinting() returns(bool)
func (_MarketToken *MarketTokenTransactor) StopMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "stopMinting")
}

// StopMinting is a paid mutator transaction binding the contract method 0x3e3e0b12.
//
// Solidity: function stopMinting() returns(bool)
func (_MarketToken *MarketTokenSession) StopMinting() (*types.Transaction, error) {
	return _MarketToken.Contract.StopMinting(&_MarketToken.TransactOpts)
}

// StopMinting is a paid mutator transaction binding the contract method 0x3e3e0b12.
//
// Solidity: function stopMinting() returns(bool)
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
const ParameterizerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getVoteBy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"paramHash\",\"type\":\"bytes32\"}],\"name\":\"resolveReparam\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getChallengeStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"reparameterize\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQuorum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getListReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getConversionRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getConversionSlope\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"},{\"name\":\"votingAddr\",\"type\":\"address\"},{\"name\":\"challengeStake\",\"type\":\"uint256\"},{\"name\":\"conversionRate\",\"type\":\"uint256\"},{\"name\":\"conversionSlope\",\"type\":\"uint256\"},{\"name\":\"listReward\",\"type\":\"uint256\"},{\"name\":\"quorum\",\"type\":\"uint256\"},{\"name\":\"voteBy\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"paramHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReparamProposedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReparamFailedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReparamSucceededEvent\",\"type\":\"event\"}]"

// ParameterizerBin is the compiled bytecode used for deploying new contracts.
const ParameterizerBin = `0x608060405234801561001057600080fd5b5060405161010080610f848339810160409081526020820151908201516060830151608084015160a085015160c086015160e09096015160018054600160a060020a031916600160a060020a0390971696909617909555600293909355600391909155600455600555600691909155600755610ef3806100916000396000f3006080604052600436106100745763ffffffff60e060020a6000350416632d0d2bc68114610079578063435c709a146100a05780635aebe768146100cc57806379bcf86f146100e1578063c26c12eb14610105578063dc2b28531461011a578063f36089ec1461012f578063fefb9dde14610144575b600080fd5b34801561008557600080fd5b5061008e610159565b60408051918252519081900360200190f35b3480156100ac57600080fd5b506100b8600435610160565b604080519115158252519081900360200190f35b3480156100d857600080fd5b5061008e610864565b3480156100ed57600080fd5b5061008e60246004803582810192910135903561086a565b34801561011157600080fd5b5061008e610dd0565b34801561012657600080fd5b5061008e610dd6565b34801561013b57600080fd5b5061008e610ddc565b34801561015057600080fd5b5061008e610de2565b6007545b90565b600154604080517f62f9a55d00000000000000000000000000000000000000000000000000000000815233600482015290516000928392600160a060020a03909116916362f9a55d9160248082019260209290919082900301818787803b1580156101ca57600080fd5b505af11580156101de573d6000803e3d6000fd5b505050506040513d60208110156101f457600080fd5b50511515610298576040805160e560020a62461bcd02815260206004820152604260248201527f4572726f723a506172616d65746572697a65722e7265736f6c7665526570617260448201527f616d202d2053656e646572206d75737420626520636f756e63696c206d656d6260648201527f6572000000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b600154604080517f11b883820000000000000000000000000000000000000000000000000000000081526004810186905260248101829052600760448201527f7265706172616d0000000000000000000000000000000000000000000000000060648201529051600160a060020a03909216916311b88382916084808201926020929091908290030181600087803b15801561033357600080fd5b505af1158015610347573d6000803e3d6000fd5b505050506040513d602081101561035d57600080fd5b505115156103db576040805160e560020a62461bcd02815260206004820152603660248201527f4572726f723a506172616d65746572697a65722e7265736f6c7665526570617260448201527f616d202d204d7573742062652061205265706172616d00000000000000000000606482015290519081900360840190fd5b600154604080517f327322c8000000000000000000000000000000000000000000000000000000008152600481018690529051600160a060020a039092169163327322c8916024808201926020929091908290030181600087803b15801561044257600080fd5b505af1158015610456573d6000803e3d6000fd5b505050506040513d602081101561046c57600080fd5b50511515610510576040805160e560020a62461bcd02815260206004820152604c60248201527f4572726f723a506172616d65746572697a65722e7265736f6c7665526570617260448201527f616d202d20506f6c6c7320666f7220746869732063616e646964617465206d7560648201527f737420626520636c6f7365640000000000000000000000000000000000000000608482015290519081900360a40190fd5b600154600160a060020a0316638f354b798461052a610dd0565b6040518363ffffffff1660e060020a02815260040180836000191660001916815260200182815260200192505050602060405180830381600087803b15801561057257600080fd5b505af1158015610586573d6000803e3d6000fd5b505050506040513d602081101561059c57600080fd5b50511561081a576000808460001916600019168152602001908152602001600020600101604051808280546001816001161561010002031660029004801561061b5780601f106105f957610100808354040283529182019161061b565b820191906000526020600020905b815481529060010190602001808311610607575b5050604080519182900382207f6368616c6c656e67655374616b650000000000000000000000000000000000008352905191829003600e01909120909350831415915061067f9050576000838152602081905260409020600290810154905561081a565b604080517f636f6e76657273696f6e526174650000000000000000000000000000000000008152905190819003600e0190208114156106d25760008381526020819052604090206002015460035561081a565b604080517f636f6e76657273696f6e536c6f706500000000000000000000000000000000008152905190819003600f0190208114156107255760008381526020819052604090206002015460045561081a565b604080517f6c697374526577617264000000000000000000000000000000000000000000008152905190819003600a0190208114156107785760008381526020819052604090206002015460055561081a565b604080517f71756f72756d0000000000000000000000000000000000000000000000000000815290519081900360060190208114156107cb5760008381526020819052604090206002015460065561081a565b604080517f766f7465427900000000000000000000000000000000000000000000000000008152905190819003600601902081141561081a576000838152602081905260409020600201546007555b6000838152602081905260408120805473ffffffffffffffffffffffffffffffffffffffff19168155906108516001830182610de8565b5060006002919091015550600192915050565b60025490565b600154604080517f62f9a55d000000000000000000000000000000000000000000000000000000008152336004820152905160009283928392600160a060020a03909216916362f9a55d9160248082019260209290919082900301818787803b1580156108d657600080fd5b505af11580156108ea573d6000803e3d6000fd5b505050506040513d602081101561090057600080fd5b505115156109a4576040805160e560020a62461bcd02815260206004820152604260248201527f4572726f723a506172616d65746572697a65722e7265706172616d657465726960448201527f7a65202d2053656e646572206d75737420626520636f756e63696c206d656d6260648201527f6572000000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b85858560405160200180848480828437820191505082815260200193505050506040516020818303038152906040526040518082805190602001908083835b60208310610a025780518252601f1990920191602091820191016109e3565b51815160209384036101000a6000190180199092169116179052604080519290940182900382206001547fb89694c6000000000000000000000000000000000000000000000000000000008452600484018290529451909850600160a060020a03909416955063b89694c6945060248083019491935090918290030181600087803b158015610a9057600080fd5b505af1158015610aa4573d6000803e3d6000fd5b505050506040513d6020811015610aba57600080fd5b505115610b5d576040805160e560020a62461bcd02815260206004820152605360248201527f4572726f723a506172616d65746572697a65722e70726f706f7365526570617260448201527f616d202d2050726f706f736564207265706172616d20697320616c726561647960648201527f206120766f74696e672063616e64696461746500000000000000000000000000608482015290519081900360a40190fd5b60606040519081016040528033600160a060020a0316815260200187878080601f01602080910402602001604051908101604052809392919081815260200183838082843750505092845250505060209081018690526000848152808252604090208251815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039091161781558282015180519192610c0292600185019290910190610e2f565b5060409190910151600290910155600154600160a060020a031663364f46a483610c2a610159565b6040805160e060020a63ffffffff86160281526024810193909352604483019190915260606004830152600760648301527f7265706172616d0000000000000000000000000000000000000000000000000060848301525160a48083019260209291908290030181600087803b158015610ca357600080fd5b505af1158015610cb7573d6000803e3d6000fd5b505050506040513d6020811015610ccd57600080fd5b50519050801515610d74576040805160e560020a62461bcd02815260206004820152605460248201527f4572726f723a506172616d65746572697a65722e7265706172616d657465726960448201527f7a65202d20436f756c64206e6f7420616464207265706172616d20746f20766f60648201527f74696e672063616e64696461746573206c697374000000000000000000000000608482015290519081900360a40190fd5b858560405180838380828437604080519190930181900381208a825292519295508794503393507f6e6b1d8a6fe9513dd2b9ed502e413958e63440558f0e1f17b15a1e688cc8612192908190036020019150a450949350505050565b60065490565b60055490565b60035490565b60045490565b50805460018160011615610100020316600290046000825580601f10610e0e5750610e2c565b601f016020900490600052602060002090810190610e2c9190610ead565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e7057805160ff1916838001178555610e9d565b82800160010185558215610e9d579182015b82811115610e9d578251825591602001919060010190610e82565b50610ea9929150610ead565b5090565b61015d91905b80821115610ea95760008155600101610eb35600a165627a7a723058203010afdee5be1d88496ca7d412133723efbf177e4075b99dbb3e9971765a2f170029`

// DeployParameterizer deploys a new Ethereum contract, binding an instance of Parameterizer to it.
func DeployParameterizer(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddr common.Address, votingAddr common.Address, challengeStake *big.Int, conversionRate *big.Int, conversionSlope *big.Int, listReward *big.Int, quorum *big.Int, voteBy *big.Int) (common.Address, *types.Transaction, *Parameterizer, error) {
	parsed, err := abi.JSON(strings.NewReader(ParameterizerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParameterizerBin), backend, tokenAddr, votingAddr, challengeStake, conversionRate, conversionSlope, listReward, quorum, voteBy)
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

// GetConversionSlope is a free data retrieval call binding the contract method 0xfefb9dde.
//
// Solidity: function getConversionSlope() constant returns(uint256)
func (_Parameterizer *ParameterizerCaller) GetConversionSlope(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getConversionSlope")
	return *ret0, err
}

// GetConversionSlope is a free data retrieval call binding the contract method 0xfefb9dde.
//
// Solidity: function getConversionSlope() constant returns(uint256)
func (_Parameterizer *ParameterizerSession) GetConversionSlope() (*big.Int, error) {
	return _Parameterizer.Contract.GetConversionSlope(&_Parameterizer.CallOpts)
}

// GetConversionSlope is a free data retrieval call binding the contract method 0xfefb9dde.
//
// Solidity: function getConversionSlope() constant returns(uint256)
func (_Parameterizer *ParameterizerCallerSession) GetConversionSlope() (*big.Int, error) {
	return _Parameterizer.Contract.GetConversionSlope(&_Parameterizer.CallOpts)
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

// Reparameterize is a paid mutator transaction binding the contract method 0x79bcf86f.
//
// Solidity: function reparameterize(name string, value uint256) returns(bytes32)
func (_Parameterizer *ParameterizerTransactor) Reparameterize(opts *bind.TransactOpts, name string, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "reparameterize", name, value)
}

// Reparameterize is a paid mutator transaction binding the contract method 0x79bcf86f.
//
// Solidity: function reparameterize(name string, value uint256) returns(bytes32)
func (_Parameterizer *ParameterizerSession) Reparameterize(name string, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.Reparameterize(&_Parameterizer.TransactOpts, name, value)
}

// Reparameterize is a paid mutator transaction binding the contract method 0x79bcf86f.
//
// Solidity: function reparameterize(name string, value uint256) returns(bytes32)
func (_Parameterizer *ParameterizerTransactorSession) Reparameterize(name string, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.Reparameterize(&_Parameterizer.TransactOpts, name, value)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(paramHash bytes32) returns(bool)
func (_Parameterizer *ParameterizerTransactor) ResolveReparam(opts *bind.TransactOpts, paramHash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "resolveReparam", paramHash)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(paramHash bytes32) returns(bool)
func (_Parameterizer *ParameterizerSession) ResolveReparam(paramHash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.Contract.ResolveReparam(&_Parameterizer.TransactOpts, paramHash)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(paramHash bytes32) returns(bool)
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
	Name     common.Hash
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReparamFailedEvent is a free log retrieval operation binding the contract event 0xb738f684b674f369ce7b32bc2f124bcbcfe6532c6629cb816c30bda0558ef621.
//
// Solidity: e ReparamFailedEvent(propHash indexed bytes32, name indexed string, value uint256)
func (_Parameterizer *ParameterizerFilterer) FilterReparamFailedEvent(opts *bind.FilterOpts, propHash [][32]byte, name []string) (*ParameterizerReparamFailedEventIterator, error) {

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

// WatchReparamFailedEvent is a free log subscription operation binding the contract event 0xb738f684b674f369ce7b32bc2f124bcbcfe6532c6629cb816c30bda0558ef621.
//
// Solidity: e ReparamFailedEvent(propHash indexed bytes32, name indexed string, value uint256)
func (_Parameterizer *ParameterizerFilterer) WatchReparamFailedEvent(opts *bind.WatchOpts, sink chan<- *ParameterizerReparamFailedEvent, propHash [][32]byte, name []string) (event.Subscription, error) {

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
	Name      common.Hash
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReparamProposedEvent is a free log retrieval operation binding the contract event 0x6e6b1d8a6fe9513dd2b9ed502e413958e63440558f0e1f17b15a1e688cc86121.
//
// Solidity: e ReparamProposedEvent(proposer indexed address, paramHash indexed bytes32, name indexed string, value uint256)
func (_Parameterizer *ParameterizerFilterer) FilterReparamProposedEvent(opts *bind.FilterOpts, proposer []common.Address, paramHash [][32]byte, name []string) (*ParameterizerReparamProposedEventIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var paramHashRule []interface{}
	for _, paramHashItem := range paramHash {
		paramHashRule = append(paramHashRule, paramHashItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ReparamProposedEvent", proposerRule, paramHashRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerReparamProposedEventIterator{contract: _Parameterizer.contract, event: "ReparamProposedEvent", logs: logs, sub: sub}, nil
}

// WatchReparamProposedEvent is a free log subscription operation binding the contract event 0x6e6b1d8a6fe9513dd2b9ed502e413958e63440558f0e1f17b15a1e688cc86121.
//
// Solidity: e ReparamProposedEvent(proposer indexed address, paramHash indexed bytes32, name indexed string, value uint256)
func (_Parameterizer *ParameterizerFilterer) WatchReparamProposedEvent(opts *bind.WatchOpts, sink chan<- *ParameterizerReparamProposedEvent, proposer []common.Address, paramHash [][32]byte, name []string) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var paramHashRule []interface{}
	for _, paramHashItem := range paramHash {
		paramHashRule = append(paramHashRule, paramHashItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ReparamProposedEvent", proposerRule, paramHashRule, nameRule)
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
	Name     common.Hash
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReparamSucceededEvent is a free log retrieval operation binding the contract event 0x8a225cd09da1d9bccd14e377451a47138eb1314a007912b6d83d820f6d989fdf.
//
// Solidity: e ReparamSucceededEvent(propHash indexed bytes32, name indexed string, value uint256)
func (_Parameterizer *ParameterizerFilterer) FilterReparamSucceededEvent(opts *bind.FilterOpts, propHash [][32]byte, name []string) (*ParameterizerReparamSucceededEventIterator, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ReparamSucceededEvent", propHashRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerReparamSucceededEventIterator{contract: _Parameterizer.contract, event: "ReparamSucceededEvent", logs: logs, sub: sub}, nil
}

// WatchReparamSucceededEvent is a free log subscription operation binding the contract event 0x8a225cd09da1d9bccd14e377451a47138eb1314a007912b6d83d820f6d989fdf.
//
// Solidity: e ReparamSucceededEvent(propHash indexed bytes32, name indexed string, value uint256)
func (_Parameterizer *ParameterizerFilterer) WatchReparamSucceededEvent(opts *bind.WatchOpts, sink chan<- *ParameterizerReparamSucceededEvent, propHash [][32]byte, name []string) (event.Subscription, error) {

	var propHashRule []interface{}
	for _, propHashItem := range propHash {
		propHashRule = append(propHashRule, propHashItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ReparamSucceededEvent", propHashRule, nameRule)
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

// VotingABI is the input ABI used to generate the binding from.
const VotingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"kind\",\"type\":\"string\"}],\"name\":\"candidateIs\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"pollClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"member\",\"type\":\"address\"}],\"name\":\"removeFromCouncil\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"kind\",\"type\":\"string\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"voteBy\",\"type\":\"uint256\"}],\"name\":\"addCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"member\",\"type\":\"address\"}],\"name\":\"addToCouncil\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"member\",\"type\":\"address\"}],\"name\":\"inCouncil\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"removeCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"quorum\",\"type\":\"uint256\"}],\"name\":\"didPass\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"market\",\"type\":\"address\"},{\"name\":\"parameterizer\",\"type\":\"address\"}],\"name\":\"setPrivilegedContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCouncil\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"VotedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"kind\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"voteBy\",\"type\":\"uint256\"}],\"name\":\"CandidateAddedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"CandidateRemovedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"member\",\"type\":\"address\"}],\"name\":\"CouncilMemberAddedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"member\",\"type\":\"address\"}],\"name\":\"CouncilMemberRemovedEvent\",\"type\":\"event\"}]"

// VotingBin is the compiled bytecode used for deploying new contracts.
const VotingBin = `0x608060405234801561001057600080fd5b5060068054600160a060020a031916331790556116b1806100326000396000f3006080604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306a49fce81146100d457806311b8838214610139578063327322c814610171578063335d808014610189578063364f46a4146101aa578063471bb309146101d157806362f9a55d146101f257806389bb617c146102135780638f354b791461022b578063a69beaba14610246578063a7e4479014610260578063b181369514610287578063b89694c61461029c578063dfb6419f146102b4575b600080fd5b3480156100e057600080fd5b506100e9610352565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561012557818101518382015260200161010d565b505050509050019250505060405180910390f35b34801561014557600080fd5b5061015d6004803590602480359081019101356103ac565b604080519115158252519081900360200190f35b34801561017d57600080fd5b5061015d600435610464565b34801561019557600080fd5b5061015d600160a060020a0360043516610506565b3480156101b657600080fd5b5061015d602460048035828101929101359035604435610708565b3480156101dd57600080fd5b5061015d600160a060020a03600435166108d8565b3480156101fe57600080fd5b5061015d600160a060020a0360043516610a48565b34801561021f57600080fd5b5061015d600435610aa0565b34801561023757600080fd5b5061015d600435602435610c75565b34801561025257600080fd5b5061025e600435610e4f565b005b34801561026c57600080fd5b5061015d600160a060020a036004358116906024351661110b565b34801561029357600080fd5b506100e9611309565b3480156102a857600080fd5b5061015d60043561136a565b3480156102c057600080fd5b506102cc6004356113af565b6040518080602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b838110156103155781810151838201526020016102fd565b50505050905090810190601f1680156103425780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b606060018054806020026020016040519081016040528092919081815260200182805480156103a157602002820191906000526020600020905b8154815260019091019060200180831161038c575b505050505090505b90565b60006103b78461136a565b1515610433576040805160e560020a62461bcd02815260206004820152603260248201527f4572726f723a566f74696e672e706f6c6c436c6f736564202d2043616e64696460448201527f61746520646f6573206e6f742065786973740000000000000000000000000000606482015290519081900360840190fd5b5060008381526020819052604090206001908101546002610100928216159290920260001901160481149392505050565b600061046f8261136a565b15156104eb576040805160e560020a62461bcd02815260206004820152603260248201527f4572726f723a566f74696e672e706f6c6c436c6f736564202d2043616e64696460448201527f61746520646f6573206e6f742065786973740000000000000000000000000000606482015290519081900360840190fd5b5060008181526020819052604090206002015442115b919050565b60045460009081908190600160a060020a03163314806105305750600554600160a060020a031633145b1515610588576040805160e560020a62461bcd02815260206004820152602481018290526000805160206116468339815191526044820152600080516020611666833981519152606482015290519081900360840190fd5b61059184610a48565b151561060d576040805160e560020a62461bcd02815260206004820152602d60248201527f4572726f723a566f74696e672e72656d6f766546726f6d436f756e63696c202d60448201527f204e6f742061206d656d62657200000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0384166000908152600260205260409020546003805491935090600019810190811061063c57fe5b60009182526020909120015460038054600160a060020a03909216925082918490811061066557fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560038054906106ae90600019830161153d565b50600160a060020a0380821660009081526002602052604080822085905591861680825282822082905591517f564e70d59b50a6b32776c3d4a1e4e378a77dc9e24c3e2fbfa04a79166a67489b9190a25060019392505050565b60045460009081908190600160a060020a03163314806107325750600554600160a060020a031633145b151561078a576040805160e560020a62461bcd02815260206004820152602481018290526000805160206116468339815191526044820152600080516020611666833981519152606482015290519081900360840190fd5b6107938561136a565b1561080e576040805160e560020a62461bcd02815260206004820152603460248201527f4572726f723a566f74696e672e61646443616e646964617465202d2043616e6460448201527f696461746520616c726561647920657869737473000000000000000000000000606482015290519081900360840190fd5b61081e428563ffffffff6114f216565b6000868152602081905260408120600180548082018255928190527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf683018990559181559193509091506108759082018888611566565b5060028101829055600060038201556040518490869089908990808383808284376040519201829003822094507fb8c6b8ac5e262d053e2f12f7a1fb6c2a024843810aaa1e8ef854ff04fb27dd9393506000925050a45060019695505050505050565b600454600090600160a060020a03163314806108fe5750600554600160a060020a031633145b1515610956576040805160e560020a62461bcd02815260206004820152602481018290526000805160206116468339815191526044820152600080516020611666833981519152606482015290519081900360840190fd5b61095f82610a48565b156109da576040805160e560020a62461bcd02815260206004820152603460248201527f4572726f723a566f74696e672e616464546f436f756e63696c202d20416c726560448201527f616479206120636f756e63696c206d656d626572000000000000000000000000606482015290519081900360840190fd5b60038054600181019091557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b81018054600160a060020a0390941673ffffffffffffffffffffffffffffffffffffffff19909416841790556000928352600260205260409092209190915590565b6003546000901515610a5c57506000610501565b600160a060020a038216600081815260026020526040902054600380549091908110610a8457fe5b600091825260209091200154600160a060020a03161492915050565b60045460009081908190600160a060020a0316331480610aca5750600554600160a060020a031633145b1515610b22576040805160e560020a62461bcd02815260206004820152602481018290526000805160206116468339815191526044820152600080516020611666833981519152606482015290519081900360840190fd5b610b2b8461136a565b1515610ba7576040805160e560020a62461bcd02815260206004820152603760248201527f4572726f723a566f74696e672e72656d6f766543616e646964617465202d204360448201527f616e64696461746520646f6573206e6f74206578697374000000000000000000606482015290519081900360840190fd5b60008481526020819052604090205460018054919350906000198101908110610bcc57fe5b9060005260206000200154905080600183815481101515610be957fe5b6000918252602090912001556001805490610c0890600019830161153d565b50600081815260208190526040808220849055858252812081815590610c3160018301826115e4565b50600060028201819055600390910181905560405185917f53ce6706735b5cd3034bf68c2a532801c1efb4ba72410b5af2167a8462f15b7d91a25060019392505050565b6000610c808361136a565b1515610cfc576040805160e560020a62461bcd02815260206004820152602c60248201527f4572726f723a566f74696e672e70617373202d2043616e64696461746520646f60448201527f6573206e6f742065786973740000000000000000000000000000000000000000606482015290519081900360840190fd5b6000838152602081905260409020600201544211610d8a576040805160e560020a62461bcd02815260206004820152603d60248201527f4572726f723a566f74696e672e70617373202d20506f6c6c696e67206d75737460448201527f20626520636c6f73656420666f7220746869732063616e646964617465000000606482015290519081900360840190fd5b600354600010610e0a576040805160e560020a62461bcd02815260206004820152602960248201527f4572726f723a566f74696e672e64696450617373202d204e6f20636f756e636960448201527f6c206d656d626572730000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600380546000858152602081905260409020909101548391610e4591606491610e39919063ffffffff6114ff16565b9063ffffffff61151416565b1190505b92915050565b610e5833610a48565b1515610ed4576040805160e560020a62461bcd02815260206004820152603160248201527f4572726f723a566f74696e672e766f7465202d2053656e646572206d7573742060448201527f626520636f756e63696c206d656d626572000000000000000000000000000000606482015290519081900360840190fd5b610edd8161136a565b1515610f59576040805160e560020a62461bcd02815260206004820152602c60248201527f4572726f723a566f74696e672e766f7465202d2043616e64696461746520646f60448201527f6573206e6f742065786973740000000000000000000000000000000000000000606482015290519081900360840190fd5b6000818152602081905260409020600201544210610fe7576040805160e560020a62461bcd02815260206004820152603860248201527f4572726f723a566f74696e672e766f7465202d20506f6c6c696e67206973206360448201527f6c6f73656420666f7220746869732063616e6469646174650000000000000000606482015290519081900360840190fd5b60008181526020818152604080832033845260040190915290205460ff16151560011415611085576040805160e560020a62461bcd02815260206004820152602c60248201527f4572726f723a566f74696e672e766f7465202d2053656e64657220686173206160448201527f6c726561647920766f7465640000000000000000000000000000000000000000606482015290519081900360840190fd5b600081815260208181526040808320338452600481018352908320805460ff1916600190811790915584845292909152600301546110c89163ffffffff6114f216565b600082815260208190526040808220600301929092559051829133917f81681507a95428bb04c5f5ac127e404caf6422d0dcf2ff77612b9201c3aa33bf9190a350565b600654600090600160a060020a03163314611196576040805160e560020a62461bcd02815260206004820152602b60248201527f4572726f723a566f74696e672e69734f776e6572202d2053656e646572206d7560448201527f7374206265206f776e6572000000000000000000000000000000000000000000606482015290519081900360840190fd5b600454600160a060020a03161561121d576040805160e560020a62461bcd02815260206004820152602481018290527f4572726f723a566f74696e672e73657450726976696c65676564436f6e74726160448201527f637473202d204d61726b6574206164647265737320616c726561647920736574606482015290519081900360840190fd5b600554600160a060020a0316156112ca576040805160e560020a62461bcd02815260206004820152604760248201527f4572726f723a566f74696e672e73657450726976696c65676564436f6e74726160448201527f637473202d20506172616d65746572697a6572206164647265737320616c726560648201527f6164792073657400000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b5060048054600160a060020a0393841673ffffffffffffffffffffffffffffffffffffffff199182161790915560058054929093169116179055600190565b606060038054806020026020016040519081016040528092919081815260200182805480156103a157602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311611343575050505050905090565b600154600090151561137e57506000610501565b60008281526020819052604090205460018054849290811061139c57fe5b6000918252602090912001541492915050565b60606000806113bd8461136a565b1515611439576040805160e560020a62461bcd02815260206004820152603460248201527f4572726f723a566f74696e672e67657443616e646964617465202d2043616e6460448201527f696461746520646f6573206e6f74206578697374000000000000000000000000606482015290519081900360840190fd5b600084815260208181526040918290206002808201546003830154600193840180548751601f610100978316159790970260001901909116949094049485018690048602840186019096528383529093909285918301828280156114de5780601f106114b3576101008083540402835291602001916114de565b820191906000526020600020905b8154815290600101906020018083116114c157829003601f168201915b505050505092509250925092509193909250565b81810182811015610e4957fe5b6000818381151561150c57fe5b049392505050565b600082151561152557506000610e49565b5081810281838281151561153557fe5b0414610e4957fe5b8154818355818111156115615760008381526020902061156191810190830161162b565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106115a75782800160ff198235161785556115d4565b828001600101855582156115d4579182015b828111156115d45782358255916020019190600101906115b9565b506115e092915061162b565b5090565b50805460018160011615610100020316600290046000825580601f1061160a5750611628565b601f016020900490600052602060002090810190611628919061162b565b50565b6103a991905b808211156115e0576000815560010161163156004572726f723a566f74696e672e68617350726976696c656765202d2053656e646572206d75737420626520612070726976696c6567656420636f6e7472616374a165627a7a723058201d48c8471d621c712f68aa15142e9aac2945a03d1600cb6a662f3bf7dc2d54310029`

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

// CandidateIs is a free data retrieval call binding the contract method 0x11b88382.
//
// Solidity: function candidateIs(hash bytes32, kind string) constant returns(bool)
func (_Voting *VotingCaller) CandidateIs(opts *bind.CallOpts, hash [32]byte, kind string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "candidateIs", hash, kind)
	return *ret0, err
}

// CandidateIs is a free data retrieval call binding the contract method 0x11b88382.
//
// Solidity: function candidateIs(hash bytes32, kind string) constant returns(bool)
func (_Voting *VotingSession) CandidateIs(hash [32]byte, kind string) (bool, error) {
	return _Voting.Contract.CandidateIs(&_Voting.CallOpts, hash, kind)
}

// CandidateIs is a free data retrieval call binding the contract method 0x11b88382.
//
// Solidity: function candidateIs(hash bytes32, kind string) constant returns(bool)
func (_Voting *VotingCallerSession) CandidateIs(hash [32]byte, kind string) (bool, error) {
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

// GetCandidate is a free data retrieval call binding the contract method 0xdfb6419f.
//
// Solidity: function getCandidate(hash bytes32) constant returns(string, uint256, uint256)
func (_Voting *VotingCaller) GetCandidate(opts *bind.CallOpts, hash [32]byte) (string, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
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
// Solidity: function getCandidate(hash bytes32) constant returns(string, uint256, uint256)
func (_Voting *VotingSession) GetCandidate(hash [32]byte) (string, *big.Int, *big.Int, error) {
	return _Voting.Contract.GetCandidate(&_Voting.CallOpts, hash)
}

// GetCandidate is a free data retrieval call binding the contract method 0xdfb6419f.
//
// Solidity: function getCandidate(hash bytes32) constant returns(string, uint256, uint256)
func (_Voting *VotingCallerSession) GetCandidate(hash [32]byte) (string, *big.Int, *big.Int, error) {
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

// AddCandidate is a paid mutator transaction binding the contract method 0x364f46a4.
//
// Solidity: function addCandidate(kind string, hash bytes32, voteBy uint256) returns(bool)
func (_Voting *VotingTransactor) AddCandidate(opts *bind.TransactOpts, kind string, hash [32]byte, voteBy *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "addCandidate", kind, hash, voteBy)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x364f46a4.
//
// Solidity: function addCandidate(kind string, hash bytes32, voteBy uint256) returns(bool)
func (_Voting *VotingSession) AddCandidate(kind string, hash [32]byte, voteBy *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, kind, hash, voteBy)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x364f46a4.
//
// Solidity: function addCandidate(kind string, hash bytes32, voteBy uint256) returns(bool)
func (_Voting *VotingTransactorSession) AddCandidate(kind string, hash [32]byte, voteBy *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, kind, hash, voteBy)
}

// AddToCouncil is a paid mutator transaction binding the contract method 0x471bb309.
//
// Solidity: function addToCouncil(member address) returns(bool)
func (_Voting *VotingTransactor) AddToCouncil(opts *bind.TransactOpts, member common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "addToCouncil", member)
}

// AddToCouncil is a paid mutator transaction binding the contract method 0x471bb309.
//
// Solidity: function addToCouncil(member address) returns(bool)
func (_Voting *VotingSession) AddToCouncil(member common.Address) (*types.Transaction, error) {
	return _Voting.Contract.AddToCouncil(&_Voting.TransactOpts, member)
}

// AddToCouncil is a paid mutator transaction binding the contract method 0x471bb309.
//
// Solidity: function addToCouncil(member address) returns(bool)
func (_Voting *VotingTransactorSession) AddToCouncil(member common.Address) (*types.Transaction, error) {
	return _Voting.Contract.AddToCouncil(&_Voting.TransactOpts, member)
}

// RemoveCandidate is a paid mutator transaction binding the contract method 0x89bb617c.
//
// Solidity: function removeCandidate(hash bytes32) returns(bool)
func (_Voting *VotingTransactor) RemoveCandidate(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "removeCandidate", hash)
}

// RemoveCandidate is a paid mutator transaction binding the contract method 0x89bb617c.
//
// Solidity: function removeCandidate(hash bytes32) returns(bool)
func (_Voting *VotingSession) RemoveCandidate(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.RemoveCandidate(&_Voting.TransactOpts, hash)
}

// RemoveCandidate is a paid mutator transaction binding the contract method 0x89bb617c.
//
// Solidity: function removeCandidate(hash bytes32) returns(bool)
func (_Voting *VotingTransactorSession) RemoveCandidate(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.RemoveCandidate(&_Voting.TransactOpts, hash)
}

// RemoveFromCouncil is a paid mutator transaction binding the contract method 0x335d8080.
//
// Solidity: function removeFromCouncil(member address) returns(bool)
func (_Voting *VotingTransactor) RemoveFromCouncil(opts *bind.TransactOpts, member common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "removeFromCouncil", member)
}

// RemoveFromCouncil is a paid mutator transaction binding the contract method 0x335d8080.
//
// Solidity: function removeFromCouncil(member address) returns(bool)
func (_Voting *VotingSession) RemoveFromCouncil(member common.Address) (*types.Transaction, error) {
	return _Voting.Contract.RemoveFromCouncil(&_Voting.TransactOpts, member)
}

// RemoveFromCouncil is a paid mutator transaction binding the contract method 0x335d8080.
//
// Solidity: function removeFromCouncil(member address) returns(bool)
func (_Voting *VotingTransactorSession) RemoveFromCouncil(member common.Address) (*types.Transaction, error) {
	return _Voting.Contract.RemoveFromCouncil(&_Voting.TransactOpts, member)
}

// SetPrivilegedContracts is a paid mutator transaction binding the contract method 0xa7e44790.
//
// Solidity: function setPrivilegedContracts(market address, parameterizer address) returns(bool)
func (_Voting *VotingTransactor) SetPrivilegedContracts(opts *bind.TransactOpts, market common.Address, parameterizer common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "setPrivilegedContracts", market, parameterizer)
}

// SetPrivilegedContracts is a paid mutator transaction binding the contract method 0xa7e44790.
//
// Solidity: function setPrivilegedContracts(market address, parameterizer address) returns(bool)
func (_Voting *VotingSession) SetPrivilegedContracts(market common.Address, parameterizer common.Address) (*types.Transaction, error) {
	return _Voting.Contract.SetPrivilegedContracts(&_Voting.TransactOpts, market, parameterizer)
}

// SetPrivilegedContracts is a paid mutator transaction binding the contract method 0xa7e44790.
//
// Solidity: function setPrivilegedContracts(market address, parameterizer address) returns(bool)
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
	Kind   common.Hash
	Hash   [32]byte
	VoteBy *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCandidateAddedEvent is a free log retrieval operation binding the contract event 0xb8c6b8ac5e262d053e2f12f7a1fb6c2a024843810aaa1e8ef854ff04fb27dd93.
//
// Solidity: e CandidateAddedEvent(kind indexed string, hash indexed bytes32, voteBy indexed uint256)
func (_Voting *VotingFilterer) FilterCandidateAddedEvent(opts *bind.FilterOpts, kind []string, hash [][32]byte, voteBy []*big.Int) (*VotingCandidateAddedEventIterator, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var voteByRule []interface{}
	for _, voteByItem := range voteBy {
		voteByRule = append(voteByRule, voteByItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CandidateAddedEvent", kindRule, hashRule, voteByRule)
	if err != nil {
		return nil, err
	}
	return &VotingCandidateAddedEventIterator{contract: _Voting.contract, event: "CandidateAddedEvent", logs: logs, sub: sub}, nil
}

// WatchCandidateAddedEvent is a free log subscription operation binding the contract event 0xb8c6b8ac5e262d053e2f12f7a1fb6c2a024843810aaa1e8ef854ff04fb27dd93.
//
// Solidity: e CandidateAddedEvent(kind indexed string, hash indexed bytes32, voteBy indexed uint256)
func (_Voting *VotingFilterer) WatchCandidateAddedEvent(opts *bind.WatchOpts, sink chan<- *VotingCandidateAddedEvent, kind []string, hash [][32]byte, voteBy []*big.Int) (event.Subscription, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var voteByRule []interface{}
	for _, voteByItem := range voteBy {
		voteByRule = append(voteByRule, voteByItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CandidateAddedEvent", kindRule, hashRule, voteByRule)
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
