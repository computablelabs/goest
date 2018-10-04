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
const MarketABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"canBeListed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"challengeExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToListing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"applied\",\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"challengeCanBeResolved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"string\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"updateStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"voterReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromListing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"determineReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"challenge\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"isListed\",\"outputs\":[{\"name\":\"listed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"marketTokenAddr\",\"type\":\"address\"},{\"name\":\"votingAddr\",\"type\":\"address\"},{\"name\":\"parameterizerAddr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"ApplicationDeletedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"applicationExpiry\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"applicant\",\"type\":\"address\"}],\"name\":\"AppliedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commitExpiry\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealExpiry\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"ChallengeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"name\":\"ChallengeFailedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"name\":\"ChallengeSucceededEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minted\",\"type\":\"uint256\"}],\"name\":\"ListedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"added\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minted\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ListingDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"}],\"name\":\"ListingDeletedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"subtracted\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minted\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ListingWithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"RewardClaimedEvent\",\"type\":\"event\"}]"

// MarketBin is the compiled bytecode used for deploying new contracts.
const MarketBin = `0x60806040523480156200001157600080fd5b5060405162002b9438038062002b94833981016040908152815160208084015192840151606085015160028054600160a060020a03808716600160a060020a0319928316179092556003805483891690831617905560048054928516929091169190911790559094018051929492909162000092916005918401906200009d565b505050505062000142565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000e057805160ff191683800117855562000110565b8280016001018555821562000110579182015b8281111562000110578251825591602001919060010190620000f3565b506200011e92915062000122565b5090565b6200013f91905b808211156200011e576000815560010162000129565b90565b612a4280620001526000396000f3006080604052600436106100da5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ca3626381146100df5780630dd606d2146100f957806317d7de7c146101255780631b7bbecb146101af578063341ff089146101c757806343e55908146101e257806377609a41146101fa57806386bb8f371461021257806389bb55c71461022d5780638a59eb5614610255578063a7aad3db1461026d578063a7e78675146102a6578063c8187cf1146102c1578063cffd46dc146102d9578063ecefbdc6146102f1575b600080fd5b3480156100eb57600080fd5b506100f7600435610309565b005b34801561010557600080fd5b506101116004356104e4565b604080519115158252519081900360200190f35b34801561013157600080fd5b5061013a610574565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561017457818101518382015260200161015c565b50505050905090810190601f1680156101a15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101bb57600080fd5b5061011160043561060b565b3480156101d357600080fd5b506100f760043560243561064c565b3480156101ee57600080fd5b50610111600435610876565b34801561020657600080fd5b5061011160043561088c565b34801561021e57600080fd5b506100f76004356024356109bd565b34801561023957600080fd5b506100f760048035906024803591604435918201910135610bfc565b34801561026157600080fd5b506100f76004356110f3565b34801561027957600080fd5b50610294600160a060020a0360043516602435604435611129565b60408051918252519081900360200190f35b3480156102b257600080fd5b506100f760043560243561120a565b3480156102cd57600080fd5b50610294600435611659565b3480156102e557600080fd5b50610294600435611880565b3480156102fd57600080fd5b50610111600435612170565b6000818152600160208190526040909120908101546101009004600160a060020a031633146103a8576040805160e560020a62461bcd02815260206004820152602960248201527f4572726f723a4d61726b65742e65786974202d204d757374206265206c69737460448201527f696e67206f776e65720000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6103b182612170565b151561042d576040805160e560020a62461bcd02815260206004820152602260248201527f4572726f723a4d61726b65742e65786974202d204d757374206265206c69737460448201527f6564000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6003810154158061045b5750600381015460009081526020819052604090206001015460a060020a900460ff165b15156104d7576040805160e560020a62461bcd02815260206004820152603260248201527f4572726f723a4d61726b65742e65786974202d2043616e6e6f7420657869742060448201527f647572696e672061206368616c6c656e67650000000000000000000000000000606482015290519081900360840190fd5b6104e082612189565b5050565b6000818152600160205260408120600301546104ff83610876565b8015610518575060008381526001602052604090205442115b801561052a575061052883612170565b155b801561055b575080158061055b5750600081815260208190526040902060019081015460a060020a900460ff161515145b15610569576001915061056e565b600091505b50919050565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106005780601f106105d557610100808354040283529160200191610600565b820191906000526020600020905b8154815290600101906020018083116105e357829003601f168201915b505050505090505b90565b6000818152600160205260408120600301548181118015610645575060008181526020819052604090206001015460a060020a900460ff16155b9392505050565b6000828152600160208190526040909120908101546101009004600160a060020a031633146106eb576040805160e560020a62461bcd02815260206004820152602c60248201527f4572726f723a4d61726b65742e6465706f736974202d204d757374206265206c60448201527f697374696e67206f776e65720000000000000000000000000000000000000000606482015290519081900360840190fd5b6002810154610700908363ffffffff6123fd16565b60028083019190915554604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b15801561077a57600080fd5b505af115801561078e573d6000803e3d6000fd5b505050506040513d60208110156107a457600080fd5b50511515610822576040805160e560020a62461bcd02815260206004820152603760248201527f4572726f723a4d61726b65742e6465706f736974202d20436f756c64206e6f7460448201527f207472616e73666572204d61726b657420546f6b656e73000000000000000000606482015290519081900360840190fd5b600281015460058201546040805185815260208101939093528281019190915251339185917fc37633a15533e70debba779cecd6f0ac9d5ff599e6b5a6c2cfae4d3c62dd03389181900360600190a3505050565b600081815260016020526040812054115b919050565b6000818152600160205260408120600301546108a78361060b565b1515610923576040805160e560020a62461bcd02815260206004820152603e60248201527f4572726f723a4d61726b65742e6368616c6c656e676543616e42655265736f6c60448201527f766564202d204368616c6c656e676520646f6573206e6f742065786973740000606482015290519081900360840190fd5b600354604080517fee684830000000000000000000000000000000000000000000000000000000008152600481018490529051600160a060020a039092169163ee684830916024808201926020929091908290030181600087803b15801561098a57600080fd5b505af115801561099e573d6000803e3d6000fd5b505050506040513d60208110156109b457600080fd5b50519392505050565b600082815260208181526040808320338452600401909152812054819060ff16156109e757600080fd5b600084815260208190526040902060019081015460a060020a900460ff16151514610a1157600080fd5b600354604080517fb43bd06900000000000000000000000000000000000000000000000000000000815233600482015260248101879052604481018690529051600160a060020a039092169163b43bd069916064808201926020929091908290030181600087803b158015610a8557600080fd5b505af1158015610a99573d6000803e3d6000fd5b505050506040513d6020811015610aaf57600080fd5b50519150610abe338585611129565b600085815260208190526040902060030154909150610ae3908363ffffffff61240a16565b6000858152602081905260409020600381019190915554610b0a908263ffffffff61240a16565b6000858152602081815260408083209384553380845260049485018352818420805460ff19166001179055600254825160e060020a63a9059cbb02815295860191909152602485018690529051600160a060020a03919091169363a9059cbb9360448083019493928390030190829087803b158015610b8857600080fd5b505af1158015610b9c573d6000803e3d6000fd5b505050506040513d6020811015610bb257600080fd5b50511515610bbf57600080fd5b604080518281529051339186917f1444f3f25ac00f3a484d68ef4ce9dd0630c3f5943d4f05253dfabc4286643a639181900360200190a350505050565b6000610c0785612170565b15610c82576040805160e560020a62461bcd02815260206004820152602360248201527f4572726f723a4d61726b65742e6170706c79202d20416c7265616479206c697360448201527f7465640000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b610c8b85610876565b15610d05576040805160e560020a62461bcd028152602060048201526024808201527f4572726f723a4d61726b65742e6170706c79202d20416c72656164792061707060448201527f6c69656400000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600480546040805160e160020a63349f642f0281526020938101849052600a60248201527f6d696e4465706f7369740000000000000000000000000000000000000000000060448201529051600160a060020a039092169263693ec85e926064808401938290030181600087803b158015610d7f57600080fd5b505af1158015610d93573d6000803e3d6000fd5b505050506040513d6020811015610da957600080fd5b5051841015610e4e576040805160e560020a62461bcd02815260206004820152605260248201527f4572726f723a4d61726b65742e6170706c79202d20616d6f756e74206d75737460448201527f2062652067726561746572207468616e2c206f7220657175616c20746f2c207460648201527f6865206d696e696d756d206465706f7369740000000000000000000000000000608482015290519081900360a40190fd5b506000848152600160208181526040808420928301805474ffffffffffffffffffffffffffffffffffffffff001916336101000217905560048054825160e160020a63349f642f028152918201849052600d60248301527f6170706c7953746167654c656e00000000000000000000000000000000000000604483015291519394610f3e94600160a060020a039093169363693ec85e93606480850194929391928390030190829087803b158015610f0557600080fd5b505af1158015610f19573d6000803e3d6000fd5b505050506040513d6020811015610f2f57600080fd5b5051429063ffffffff6123fd16565b81556002810184905560006005820155610f5c60048201848461293e565b506002546001820154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152610100909204600160a060020a0390811660048401523060248401526044830188905290519216916323b872dd916064808201926020929091908290030181600087803b158015610fdc57600080fd5b505af1158015610ff0573d6000803e3d6000fd5b505050506040513d602081101561100657600080fd5b50511515611084576040805160e560020a62461bcd02815260206004820152603560248201527f4572726f723a4d61726b65742e6170706c79202d20436f756c64206e6f74207460448201527f72616e73666572204d61726b657420546f6b656e730000000000000000000000606482015290519081900360840190fd5b805460408051868152602081018390526060918101828152918101859052339288927fa323908d99431e979128331420090c1eaaaec95e1a53484463417a7ee79c956c928992918991899160808201848480828437604051920182900397509095505050505050a35050505050565b6110fc816104e4565b1561110f5761110a8161241c565b611126565b6111188161088c565b156100da5761110a816125e5565b50565b6000828152602081815260408083206003808201549154905483517fb43bd069000000000000000000000000000000000000000000000000000000008152600160a060020a038a81166004830152602482018a905260448201899052945193959294879492169263b43bd0699260648084019382900301818787803b1580156111b157600080fd5b505af11580156111c5573d6000803e3d6000fd5b505050506040513d60208110156111db57600080fd5b505190506111ff836111f3838563ffffffff61290016565b9063ffffffff61292916565b979650505050505050565b6000828152600160208190526040909120908101546101009004600160a060020a031633146112a9576040805160e560020a62461bcd02815260206004820152602d60248201527f4572726f723a4d61726b65742e7769746864726177202d204d7573742062652060448201527f6c697374696e67206f776e657200000000000000000000000000000000000000606482015290519081900360840190fd5b6002810154821115611351576040805160e560020a62461bcd02815260206004820152604f60248201527f4572726f723a4d61726b65742e7769746864726177202d20616d6f756e74206d60448201527f757374206265206c657373207468616e206f7220657175616c20746f2074686560648201527f206c697374696e6720737570706c790000000000000000000000000000000000608482015290519081900360a40190fd5b600480546040805160e160020a63349f642f0281526020938101849052600a60248201527f6d696e4465706f7369740000000000000000000000000000000000000000000060448201529051600160a060020a039092169263693ec85e926064808401938290030181600087803b1580156113cb57600080fd5b505af11580156113df573d6000803e3d6000fd5b505050506040513d60208110156113f557600080fd5b5051600582015460028301546114229185916114169163ffffffff6123fd16565b9063ffffffff61240a16565b10156114c4576040805160e560020a62461bcd02815260206004820152604c60248201527f4572726f723a4d61726b65742e7769746864726177202d2043616e6e6f74207760448201527f69746864726177206c697374696e6720737570706c792062656c6f77206d696e60648201527f696d756d206465706f7369740000000000000000000000000000000000000000608482015290519081900360a40190fd5b60028101546114d9908363ffffffff61240a16565b600280830191909155546040805160e060020a63a9059cbb028152336004820152602481018590529051600160a060020a039092169163a9059cbb916044808201926020929091908290030181600087803b15801561153757600080fd5b505af115801561154b573d6000803e3d6000fd5b505050506040513d602081101561156157600080fd5b50511515611605576040805160e560020a62461bcd02815260206004820152604260248201527f4572726f723a4d61726b65742e7769746864726177202d20436f756c64206e6f60448201527f74207472616e7366657220746f6b656e7320746f206c697374696e67206f776e60648201527f6572000000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b600281015460058201546040805185815260208101939093528281019190915251339185917fda21bce65544aeea410c866bc9fe650b80110de14e8bb84bc79c50e44091f91e9181900360600190a3505050565b60008181526020819052604081206001015460a060020a900460ff161580156117115750600354604080517fee684830000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a039092169163ee684830916024808201926020929091908290030181600087803b1580156116e457600080fd5b505af11580156116f8573d6000803e3d6000fd5b505050506040513d602081101561170e57600080fd5b50515b151561178d576040805160e560020a62461bcd02815260206004820152603360248201527f4572726f723a4d61726b65742e64657465726d696e65526577617264202d205060448201527f6f6c6c206d757374206861766520656e64656400000000000000000000000000606482015290519081900360840190fd5b600354604080517f053e71a6000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a039092169163053e71a6916024808201926020929091908290030181600087803b1580156117f457600080fd5b505af1158015611808573d6000803e3d6000fd5b505050506040513d602081101561181e57600080fd5b5051151561185157600082815260208190526040902060029081015461184a919063ffffffff61290016565b9050610887565b6000828152602081905260409020805460029182015461187a926114169163ffffffff61290016565b92915050565b600081815260016020908152604080832060048054835160e160020a63349f642f028152918201859052600a60248301527f6d696e4465706f73697400000000000000000000000000000000000000000000604483015292519193859384938493600160a060020a039092169263693ec85e926064808201939182900301818787803b15801561190f57600080fd5b505af1158015611923573d6000803e3d6000fd5b505050506040513d602081101561193957600080fd5b5051925061194686610876565b806119555750600184015460ff165b15156119d1576040805160e560020a62461bcd02815260206004820152603c60248201527f4572726f723a4d61726b65742e6368616c6c656e6765202d204d75737420626560448201527f2061206c697374696e67206f7220616e206170706c69636174696f6e00000000606482015290519081900360840190fd5b600384015415806119ff5750600384015460009081526020819052604090206001015460a060020a900460ff165b1515611a7b576040805160e560020a62461bcd02815260206004820152602b60248201527f4572726f723a4d61726b65742e6368616c6c656e6765202d20416c726561647960448201527f206368616c6c656e676564000000000000000000000000000000000000000000606482015290519081900360840190fd5b8284600501548560020154011015611a9f57611a9686612189565b60009450612167565b600354600480546040805160e160020a63349f642f0281526020938101849052600a60248201527f766f746551756f72756d0000000000000000000000000000000000000000000060448201529051600160a060020a03948516946332ed3d609493169263693ec85e92606480820193918290030181600087803b158015611b2657600080fd5b505af1158015611b3a573d6000803e3d6000fd5b505050506040513d6020811015611b5057600080fd5b5051600480546040805160e160020a63349f642f0281526020938101849052600e60248201527f636f6d6d697453746167654c656e00000000000000000000000000000000000060448201529051600160a060020a039092169263693ec85e926064808401938290030181600087803b158015611bcc57600080fd5b505af1158015611be0573d6000803e3d6000fd5b505050506040513d6020811015611bf657600080fd5b5051600480546040805160e160020a63349f642f0281526020938101849052600e60248201527f72657665616c53746167654c656e00000000000000000000000000000000000060448201529051600160a060020a039092169263693ec85e926064808401938290030181600087803b158015611c7257600080fd5b505af1158015611c86573d6000803e3d6000fd5b505050506040513d6020811015611c9c57600080fd5b5051604080517c010000000000000000000000000000000000000000000000000000000063ffffffff87160281526004810194909452602484019290925260448301525160648083019260209291908290030181600087803b158015611d0157600080fd5b505af1158015611d15573d6000803e3d6000fd5b505050506040513d6020811015611d2b57600080fd5b50516040805160a0810180835260045460e160020a63349f642f02909152602060a48301819052600f60c48401527f64697370656e736174696f6e506374000000000000000000000000000000000060e4840152925193985090928392611e10926064926111f3928a92611e0492600160a060020a039091169163693ec85e91610104808b01929190818c030181600087803b158015611dca57600080fd5b505af1158015611dde573d6000803e3d6000fd5b505050506040513d6020811015611df457600080fd5b505160649063ffffffff61240a16565b9063ffffffff61290016565b81523360208083019190915260006040808401829052606080850189905260809485018390528a835282845291819020855181559285015160018401805492870151151560a060020a0274ff000000000000000000000000000000000000000019600160a060020a039390931673ffffffffffffffffffffffffffffffffffffffff19909416939093179190911691909117905583015160028083019190915592909101516003918201558501869055840154611ed3908463ffffffff61240a16565b60028086019190915554604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018690529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b158015611f4d57600080fd5b505af1158015611f61573d6000803e3d6000fd5b505050506040513d6020811015611f7757600080fd5b50511515611ff5576040805160e560020a62461bcd02815260206004820152603960248201527f4572726f723a4d61726b65742e6368616c6c656e6765202d20436f756c64206e60448201527f6f74207472616e73666572204d61726b657420746f6b656e7300000000000000606482015290519081900360840190fd5b600354604080517fdc1b946f000000000000000000000000000000000000000000000000000000008152600481018890529051600160a060020a039092169163dc1b946f916024808201926020929091908290030181600087803b15801561205c57600080fd5b505af1158015612070573d6000803e3d6000fd5b505050506040513d602081101561208657600080fd5b5051600354604080517f1be8ea50000000000000000000000000000000000000000000000000000000008152600481018990529051929450600160a060020a0390911691631be8ea50916024808201926020929091908290030181600087803b1580156120f257600080fd5b505af1158015612106573d6000803e3d6000fd5b505050506040513d602081101561211c57600080fd5b505160408051878152602081018590528082018390529051919250339188917f06f6018cb708e7f4406b7192a2316e87cb4e103f0083b9ca1545d38623bafcb4919081900360600190a35b50505050919050565b6000908152600160208190526040909120015460ff1690565b60008181526001602081905260409091209081015460ff16156121d65760405182907fbf8f571a7fa2903765660aa54ed716098a6b456db5be6fba8e17bffbce48c85e90600090a2612202565b60405182907fc2a99ac7c5ef757a2eb1b5e61e4b96d45431017bc7680e2d8c463ec4e830ef9790600090a25b60008281526001602081905260408220828155908101805474ffffffffffffffffffffffffffffffffffffffffff1916905560028101829055600381018290559061225060048301826129bc565b600582016000905550506000816002015411156104e057600280546001830154918301546040805160e060020a63a9059cbb028152610100909404600160a060020a03908116600486015260248501929092525191169163a9059cbb9160448083019260209291908290030181600087803b1580156122ce57600080fd5b505af11580156122e2573d6000803e3d6000fd5b505050506040513d60208110156122f857600080fd5b50511515612376576040805160e560020a62461bcd02815260206004820152603e60248201527f4572726f723a4d61726b65742e64656c6574654c697374696e67202d20436f7560448201527f6c64206e6f7420726566756e642072656d61696e696e6720737570706c790000606482015290519081900360840190fd5b6002546005820154604080517f42966c68000000000000000000000000000000000000000000000000000000008152600481019290925251600160a060020a03909216916342966c689160248082019260009290919082900301818387803b1580156123e157600080fd5b505af11580156123f5573d6000803e3d6000fd5b505050505050565b8181018281101561187a57fe5b60008282111561241657fe5b50900390565b600081815260016020819052604082200154819060ff1615156125e0576000838152600160208181526040808420808401805460ff191690941790935560048054825160e160020a63349f642f028152918201849052600a60248301527f6c6973745265776172640000000000000000000000000000000000000000000060448301529151939650600160a060020a039091169363693ec85e93606480840194939192918390030190829087803b1580156124d657600080fd5b505af11580156124ea573d6000803e3d6000fd5b505050506040513d602081101561250057600080fd5b5051600254604080517fa0712d68000000000000000000000000000000000000000000000000000000008152600481018490529051929350600160a060020a039091169163a0712d68916024808201926020929091908290030181600087803b15801561256c57600080fd5b505af1158015612580573d6000803e3d6000fd5b505050506040513d602081101561259657600080fd5b50506005820181905560028201546040805191825260208201839052805185927f50b5f2e84e4525dee7bfd028bc77620921b1ca3bedd9c83a99698c603c8fe6fb92908290030190a25b505050565b6000818152600160205260408120600301549061260182611659565b600083815260208181526040808320600101805474ff0000000000000000000000000000000000000000191660a060020a17905560035481517f053e71a6000000000000000000000000000000000000000000000000000000008152600481018890529151949550600160a060020a03169363053e71a693602480840194938390030190829087803b15801561269657600080fd5b505af11580156126aa573d6000803e3d6000fd5b505050506040513d60208110156126c057600080fd5b5051600083815260208181526040808320600390810194909455925483517f49403183000000000000000000000000000000000000000000000000000000008152600481018790529351600160a060020a039091169363494031839360248083019493928390030190829087803b15801561273a57600080fd5b505af115801561274e573d6000803e3d6000fd5b505050506040513d602081101561276457600080fd5b505115612802576127748361241c565b600083815260016020526040902060020154612796908263ffffffff6123fd16565b6000848152600160209081526040808320600201939093558482528181529082902080546003909101548351918252918101919091528151849286927f824960b36dcff66c0b43953e552b1b0d622dcb3f3cf71cf1cbe88039a781a853929081900390910190a36125e0565b61280b83612189565b60025460008381526020818152604080832060010154815160e060020a63a9059cbb028152600160a060020a03918216600482015260248101879052915194169363a9059cbb93604480840194938390030190829087803b15801561286f57600080fd5b505af1158015612883573d6000803e3d6000fd5b505050506040513d602081101561289957600080fd5b505115156128a657600080fd5b6000828152602081815260409182902080546003909101548351918252918101919091528151849286927f3dddcf5cd7cb20cfa68f0643b9422cdbc130a6602933781bf9855e10fd205e42929081900390910190a3505050565b60008215156129115750600061187a565b5081810281838281151561292157fe5b041461187a57fe5b6000818381151561293657fe5b049392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061297f5782800160ff198235161785556129ac565b828001600101855582156129ac579182015b828111156129ac578235825591602001919060010190612991565b506129b89291506129fc565b5090565b50805460018160011615610100020316600290046000825580601f106129e25750611126565b601f01602090049060005260206000209081019061112691905b61060891905b808211156129b85760008155600101612a025600a165627a7a72305820407433c1fa70f699c6ae97782481399af82e802810e77202bf29dcb911839e710029`

// DeployMarket deploys a new Ethereum contract, binding an instance of Market to it.
func DeployMarket(auth *bind.TransactOpts, backend bind.ContractBackend, marketTokenAddr common.Address, votingAddr common.Address, parameterizerAddr common.Address, name string) (common.Address, *types.Transaction, *Market, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarketBin), backend, marketTokenAddr, votingAddr, parameterizerAddr, name)
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

// Applied is a free data retrieval call binding the contract method 0x43e55908.
//
// Solidity: function applied(listingHash bytes32) constant returns(exists bool)
func (_Market *MarketCaller) Applied(opts *bind.CallOpts, listingHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "applied", listingHash)
	return *ret0, err
}

// Applied is a free data retrieval call binding the contract method 0x43e55908.
//
// Solidity: function applied(listingHash bytes32) constant returns(exists bool)
func (_Market *MarketSession) Applied(listingHash [32]byte) (bool, error) {
	return _Market.Contract.Applied(&_Market.CallOpts, listingHash)
}

// Applied is a free data retrieval call binding the contract method 0x43e55908.
//
// Solidity: function applied(listingHash bytes32) constant returns(exists bool)
func (_Market *MarketCallerSession) Applied(listingHash [32]byte) (bool, error) {
	return _Market.Contract.Applied(&_Market.CallOpts, listingHash)
}

// CanBeListed is a free data retrieval call binding the contract method 0x0dd606d2.
//
// Solidity: function canBeListed(listingHash bytes32) constant returns(bool)
func (_Market *MarketCaller) CanBeListed(opts *bind.CallOpts, listingHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "canBeListed", listingHash)
	return *ret0, err
}

// CanBeListed is a free data retrieval call binding the contract method 0x0dd606d2.
//
// Solidity: function canBeListed(listingHash bytes32) constant returns(bool)
func (_Market *MarketSession) CanBeListed(listingHash [32]byte) (bool, error) {
	return _Market.Contract.CanBeListed(&_Market.CallOpts, listingHash)
}

// CanBeListed is a free data retrieval call binding the contract method 0x0dd606d2.
//
// Solidity: function canBeListed(listingHash bytes32) constant returns(bool)
func (_Market *MarketCallerSession) CanBeListed(listingHash [32]byte) (bool, error) {
	return _Market.Contract.CanBeListed(&_Market.CallOpts, listingHash)
}

// ChallengeCanBeResolved is a free data retrieval call binding the contract method 0x77609a41.
//
// Solidity: function challengeCanBeResolved(listingHash bytes32) constant returns(bool)
func (_Market *MarketCaller) ChallengeCanBeResolved(opts *bind.CallOpts, listingHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "challengeCanBeResolved", listingHash)
	return *ret0, err
}

// ChallengeCanBeResolved is a free data retrieval call binding the contract method 0x77609a41.
//
// Solidity: function challengeCanBeResolved(listingHash bytes32) constant returns(bool)
func (_Market *MarketSession) ChallengeCanBeResolved(listingHash [32]byte) (bool, error) {
	return _Market.Contract.ChallengeCanBeResolved(&_Market.CallOpts, listingHash)
}

// ChallengeCanBeResolved is a free data retrieval call binding the contract method 0x77609a41.
//
// Solidity: function challengeCanBeResolved(listingHash bytes32) constant returns(bool)
func (_Market *MarketCallerSession) ChallengeCanBeResolved(listingHash [32]byte) (bool, error) {
	return _Market.Contract.ChallengeCanBeResolved(&_Market.CallOpts, listingHash)
}

// ChallengeExists is a free data retrieval call binding the contract method 0x1b7bbecb.
//
// Solidity: function challengeExists(listingHash bytes32) constant returns(bool)
func (_Market *MarketCaller) ChallengeExists(opts *bind.CallOpts, listingHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "challengeExists", listingHash)
	return *ret0, err
}

// ChallengeExists is a free data retrieval call binding the contract method 0x1b7bbecb.
//
// Solidity: function challengeExists(listingHash bytes32) constant returns(bool)
func (_Market *MarketSession) ChallengeExists(listingHash [32]byte) (bool, error) {
	return _Market.Contract.ChallengeExists(&_Market.CallOpts, listingHash)
}

// ChallengeExists is a free data retrieval call binding the contract method 0x1b7bbecb.
//
// Solidity: function challengeExists(listingHash bytes32) constant returns(bool)
func (_Market *MarketCallerSession) ChallengeExists(listingHash [32]byte) (bool, error) {
	return _Market.Contract.ChallengeExists(&_Market.CallOpts, listingHash)
}

// DetermineReward is a free data retrieval call binding the contract method 0xc8187cf1.
//
// Solidity: function determineReward(id uint256) constant returns(uint256)
func (_Market *MarketCaller) DetermineReward(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "determineReward", id)
	return *ret0, err
}

// DetermineReward is a free data retrieval call binding the contract method 0xc8187cf1.
//
// Solidity: function determineReward(id uint256) constant returns(uint256)
func (_Market *MarketSession) DetermineReward(id *big.Int) (*big.Int, error) {
	return _Market.Contract.DetermineReward(&_Market.CallOpts, id)
}

// DetermineReward is a free data retrieval call binding the contract method 0xc8187cf1.
//
// Solidity: function determineReward(id uint256) constant returns(uint256)
func (_Market *MarketCallerSession) DetermineReward(id *big.Int) (*big.Int, error) {
	return _Market.Contract.DetermineReward(&_Market.CallOpts, id)
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

// IsListed is a free data retrieval call binding the contract method 0xecefbdc6.
//
// Solidity: function isListed(listingHash bytes32) constant returns(listed bool)
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
// Solidity: function isListed(listingHash bytes32) constant returns(listed bool)
func (_Market *MarketSession) IsListed(listingHash [32]byte) (bool, error) {
	return _Market.Contract.IsListed(&_Market.CallOpts, listingHash)
}

// IsListed is a free data retrieval call binding the contract method 0xecefbdc6.
//
// Solidity: function isListed(listingHash bytes32) constant returns(listed bool)
func (_Market *MarketCallerSession) IsListed(listingHash [32]byte) (bool, error) {
	return _Market.Contract.IsListed(&_Market.CallOpts, listingHash)
}

// VoterReward is a free data retrieval call binding the contract method 0xa7aad3db.
//
// Solidity: function voterReward(voter address, id uint256, salt uint256) constant returns(uint256)
func (_Market *MarketCaller) VoterReward(opts *bind.CallOpts, voter common.Address, id *big.Int, salt *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "voterReward", voter, id, salt)
	return *ret0, err
}

// VoterReward is a free data retrieval call binding the contract method 0xa7aad3db.
//
// Solidity: function voterReward(voter address, id uint256, salt uint256) constant returns(uint256)
func (_Market *MarketSession) VoterReward(voter common.Address, id *big.Int, salt *big.Int) (*big.Int, error) {
	return _Market.Contract.VoterReward(&_Market.CallOpts, voter, id, salt)
}

// VoterReward is a free data retrieval call binding the contract method 0xa7aad3db.
//
// Solidity: function voterReward(voter address, id uint256, salt uint256) constant returns(uint256)
func (_Market *MarketCallerSession) VoterReward(voter common.Address, id *big.Int, salt *big.Int) (*big.Int, error) {
	return _Market.Contract.VoterReward(&_Market.CallOpts, voter, id, salt)
}

// Apply is a paid mutator transaction binding the contract method 0x89bb55c7.
//
// Solidity: function apply(listingHash bytes32, amount uint256, data string) returns()
func (_Market *MarketTransactor) Apply(opts *bind.TransactOpts, listingHash [32]byte, amount *big.Int, data string) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "apply", listingHash, amount, data)
}

// Apply is a paid mutator transaction binding the contract method 0x89bb55c7.
//
// Solidity: function apply(listingHash bytes32, amount uint256, data string) returns()
func (_Market *MarketSession) Apply(listingHash [32]byte, amount *big.Int, data string) (*types.Transaction, error) {
	return _Market.Contract.Apply(&_Market.TransactOpts, listingHash, amount, data)
}

// Apply is a paid mutator transaction binding the contract method 0x89bb55c7.
//
// Solidity: function apply(listingHash bytes32, amount uint256, data string) returns()
func (_Market *MarketTransactorSession) Apply(listingHash [32]byte, amount *big.Int, data string) (*types.Transaction, error) {
	return _Market.Contract.Apply(&_Market.TransactOpts, listingHash, amount, data)
}

// Challenge is a paid mutator transaction binding the contract method 0xcffd46dc.
//
// Solidity: function challenge(listingHash bytes32) returns(id uint256)
func (_Market *MarketTransactor) Challenge(opts *bind.TransactOpts, listingHash [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "challenge", listingHash)
}

// Challenge is a paid mutator transaction binding the contract method 0xcffd46dc.
//
// Solidity: function challenge(listingHash bytes32) returns(id uint256)
func (_Market *MarketSession) Challenge(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Challenge(&_Market.TransactOpts, listingHash)
}

// Challenge is a paid mutator transaction binding the contract method 0xcffd46dc.
//
// Solidity: function challenge(listingHash bytes32) returns(id uint256)
func (_Market *MarketTransactorSession) Challenge(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.Challenge(&_Market.TransactOpts, listingHash)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x86bb8f37.
//
// Solidity: function claimReward(id uint256, salt uint256) returns()
func (_Market *MarketTransactor) ClaimReward(opts *bind.TransactOpts, id *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "claimReward", id, salt)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x86bb8f37.
//
// Solidity: function claimReward(id uint256, salt uint256) returns()
func (_Market *MarketSession) ClaimReward(id *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _Market.Contract.ClaimReward(&_Market.TransactOpts, id, salt)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x86bb8f37.
//
// Solidity: function claimReward(id uint256, salt uint256) returns()
func (_Market *MarketTransactorSession) ClaimReward(id *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _Market.Contract.ClaimReward(&_Market.TransactOpts, id, salt)
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

// UpdateStatus is a paid mutator transaction binding the contract method 0x8a59eb56.
//
// Solidity: function updateStatus(listingHash bytes32) returns()
func (_Market *MarketTransactor) UpdateStatus(opts *bind.TransactOpts, listingHash [32]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "updateStatus", listingHash)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x8a59eb56.
//
// Solidity: function updateStatus(listingHash bytes32) returns()
func (_Market *MarketSession) UpdateStatus(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.UpdateStatus(&_Market.TransactOpts, listingHash)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x8a59eb56.
//
// Solidity: function updateStatus(listingHash bytes32) returns()
func (_Market *MarketTransactorSession) UpdateStatus(listingHash [32]byte) (*types.Transaction, error) {
	return _Market.Contract.UpdateStatus(&_Market.TransactOpts, listingHash)
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

// MarketApplicationDeletedEventIterator is returned from FilterApplicationDeletedEvent and is used to iterate over the raw logs and unpacked data for ApplicationDeletedEvent events raised by the Market contract.
type MarketApplicationDeletedEventIterator struct {
	Event *MarketApplicationDeletedEvent // Event containing the contract specifics and raw log

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
func (it *MarketApplicationDeletedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketApplicationDeletedEvent)
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
		it.Event = new(MarketApplicationDeletedEvent)
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
func (it *MarketApplicationDeletedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketApplicationDeletedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketApplicationDeletedEvent represents a ApplicationDeletedEvent event raised by the Market contract.
type MarketApplicationDeletedEvent struct {
	ListingHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterApplicationDeletedEvent is a free log retrieval operation binding the contract event 0xc2a99ac7c5ef757a2eb1b5e61e4b96d45431017bc7680e2d8c463ec4e830ef97.
//
// Solidity: e ApplicationDeletedEvent(listingHash indexed bytes32)
func (_Market *MarketFilterer) FilterApplicationDeletedEvent(opts *bind.FilterOpts, listingHash [][32]byte) (*MarketApplicationDeletedEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ApplicationDeletedEvent", listingHashRule)
	if err != nil {
		return nil, err
	}
	return &MarketApplicationDeletedEventIterator{contract: _Market.contract, event: "ApplicationDeletedEvent", logs: logs, sub: sub}, nil
}

// WatchApplicationDeletedEvent is a free log subscription operation binding the contract event 0xc2a99ac7c5ef757a2eb1b5e61e4b96d45431017bc7680e2d8c463ec4e830ef97.
//
// Solidity: e ApplicationDeletedEvent(listingHash indexed bytes32)
func (_Market *MarketFilterer) WatchApplicationDeletedEvent(opts *bind.WatchOpts, sink chan<- *MarketApplicationDeletedEvent, listingHash [][32]byte) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ApplicationDeletedEvent", listingHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketApplicationDeletedEvent)
				if err := _Market.contract.UnpackLog(event, "ApplicationDeletedEvent", log); err != nil {
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
	ListingHash       [32]byte
	Deposit           *big.Int
	ApplicationExpiry *big.Int
	Data              string
	Applicant         common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAppliedEvent is a free log retrieval operation binding the contract event 0xa323908d99431e979128331420090c1eaaaec95e1a53484463417a7ee79c956c.
//
// Solidity: e AppliedEvent(listingHash indexed bytes32, deposit uint256, applicationExpiry uint256, data string, applicant indexed address)
func (_Market *MarketFilterer) FilterAppliedEvent(opts *bind.FilterOpts, listingHash [][32]byte, applicant []common.Address) (*MarketAppliedEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	var applicantRule []interface{}
	for _, applicantItem := range applicant {
		applicantRule = append(applicantRule, applicantItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "AppliedEvent", listingHashRule, applicantRule)
	if err != nil {
		return nil, err
	}
	return &MarketAppliedEventIterator{contract: _Market.contract, event: "AppliedEvent", logs: logs, sub: sub}, nil
}

// WatchAppliedEvent is a free log subscription operation binding the contract event 0xa323908d99431e979128331420090c1eaaaec95e1a53484463417a7ee79c956c.
//
// Solidity: e AppliedEvent(listingHash indexed bytes32, deposit uint256, applicationExpiry uint256, data string, applicant indexed address)
func (_Market *MarketFilterer) WatchAppliedEvent(opts *bind.WatchOpts, sink chan<- *MarketAppliedEvent, listingHash [][32]byte, applicant []common.Address) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	var applicantRule []interface{}
	for _, applicantItem := range applicant {
		applicantRule = append(applicantRule, applicantItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "AppliedEvent", listingHashRule, applicantRule)
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

// MarketChallengeEventIterator is returned from FilterChallengeEvent and is used to iterate over the raw logs and unpacked data for ChallengeEvent events raised by the Market contract.
type MarketChallengeEventIterator struct {
	Event *MarketChallengeEvent // Event containing the contract specifics and raw log

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
func (it *MarketChallengeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketChallengeEvent)
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
		it.Event = new(MarketChallengeEvent)
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
func (it *MarketChallengeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketChallengeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketChallengeEvent represents a ChallengeEvent event raised by the Market contract.
type MarketChallengeEvent struct {
	ListingHash  [32]byte
	Id           *big.Int
	CommitExpiry *big.Int
	RevealExpiry *big.Int
	Challenger   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChallengeEvent is a free log retrieval operation binding the contract event 0x06f6018cb708e7f4406b7192a2316e87cb4e103f0083b9ca1545d38623bafcb4.
//
// Solidity: e ChallengeEvent(listingHash indexed bytes32, id uint256, commitExpiry uint256, revealExpiry uint256, challenger indexed address)
func (_Market *MarketFilterer) FilterChallengeEvent(opts *bind.FilterOpts, listingHash [][32]byte, challenger []common.Address) (*MarketChallengeEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ChallengeEvent", listingHashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &MarketChallengeEventIterator{contract: _Market.contract, event: "ChallengeEvent", logs: logs, sub: sub}, nil
}

// WatchChallengeEvent is a free log subscription operation binding the contract event 0x06f6018cb708e7f4406b7192a2316e87cb4e103f0083b9ca1545d38623bafcb4.
//
// Solidity: e ChallengeEvent(listingHash indexed bytes32, id uint256, commitExpiry uint256, revealExpiry uint256, challenger indexed address)
func (_Market *MarketFilterer) WatchChallengeEvent(opts *bind.WatchOpts, sink chan<- *MarketChallengeEvent, listingHash [][32]byte, challenger []common.Address) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ChallengeEvent", listingHashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketChallengeEvent)
				if err := _Market.contract.UnpackLog(event, "ChallengeEvent", log); err != nil {
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
	Id          *big.Int
	RewardPool  *big.Int
	TotalTokens *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChallengeFailedEvent is a free log retrieval operation binding the contract event 0x824960b36dcff66c0b43953e552b1b0d622dcb3f3cf71cf1cbe88039a781a853.
//
// Solidity: e ChallengeFailedEvent(listingHash indexed bytes32, id indexed uint256, rewardPool uint256, totalTokens uint256)
func (_Market *MarketFilterer) FilterChallengeFailedEvent(opts *bind.FilterOpts, listingHash [][32]byte, id []*big.Int) (*MarketChallengeFailedEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ChallengeFailedEvent", listingHashRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MarketChallengeFailedEventIterator{contract: _Market.contract, event: "ChallengeFailedEvent", logs: logs, sub: sub}, nil
}

// WatchChallengeFailedEvent is a free log subscription operation binding the contract event 0x824960b36dcff66c0b43953e552b1b0d622dcb3f3cf71cf1cbe88039a781a853.
//
// Solidity: e ChallengeFailedEvent(listingHash indexed bytes32, id indexed uint256, rewardPool uint256, totalTokens uint256)
func (_Market *MarketFilterer) WatchChallengeFailedEvent(opts *bind.WatchOpts, sink chan<- *MarketChallengeFailedEvent, listingHash [][32]byte, id []*big.Int) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ChallengeFailedEvent", listingHashRule, idRule)
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
	Id          *big.Int
	RewardPool  *big.Int
	TotalTokens *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChallengeSucceededEvent is a free log retrieval operation binding the contract event 0x3dddcf5cd7cb20cfa68f0643b9422cdbc130a6602933781bf9855e10fd205e42.
//
// Solidity: e ChallengeSucceededEvent(listingHash indexed bytes32, id indexed uint256, rewardPool uint256, totalTokens uint256)
func (_Market *MarketFilterer) FilterChallengeSucceededEvent(opts *bind.FilterOpts, listingHash [][32]byte, id []*big.Int) (*MarketChallengeSucceededEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ChallengeSucceededEvent", listingHashRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MarketChallengeSucceededEventIterator{contract: _Market.contract, event: "ChallengeSucceededEvent", logs: logs, sub: sub}, nil
}

// WatchChallengeSucceededEvent is a free log subscription operation binding the contract event 0x3dddcf5cd7cb20cfa68f0643b9422cdbc130a6602933781bf9855e10fd205e42.
//
// Solidity: e ChallengeSucceededEvent(listingHash indexed bytes32, id indexed uint256, rewardPool uint256, totalTokens uint256)
func (_Market *MarketFilterer) WatchChallengeSucceededEvent(opts *bind.WatchOpts, sink chan<- *MarketChallengeSucceededEvent, listingHash [][32]byte, id []*big.Int) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ChallengeSucceededEvent", listingHashRule, idRule)
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
	Supply      *big.Int
	Minted      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterListedEvent is a free log retrieval operation binding the contract event 0x50b5f2e84e4525dee7bfd028bc77620921b1ca3bedd9c83a99698c603c8fe6fb.
//
// Solidity: e ListedEvent(listingHash indexed bytes32, supply uint256, minted uint256)
func (_Market *MarketFilterer) FilterListedEvent(opts *bind.FilterOpts, listingHash [][32]byte) (*MarketListedEventIterator, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ListedEvent", listingHashRule)
	if err != nil {
		return nil, err
	}
	return &MarketListedEventIterator{contract: _Market.contract, event: "ListedEvent", logs: logs, sub: sub}, nil
}

// WatchListedEvent is a free log subscription operation binding the contract event 0x50b5f2e84e4525dee7bfd028bc77620921b1ca3bedd9c83a99698c603c8fe6fb.
//
// Solidity: e ListedEvent(listingHash indexed bytes32, supply uint256, minted uint256)
func (_Market *MarketFilterer) WatchListedEvent(opts *bind.WatchOpts, sink chan<- *MarketListedEvent, listingHash [][32]byte) (event.Subscription, error) {

	var listingHashRule []interface{}
	for _, listingHashItem := range listingHash {
		listingHashRule = append(listingHashRule, listingHashItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ListedEvent", listingHashRule)
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
	Added       *big.Int
	Supply      *big.Int
	Minted      *big.Int
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterListingDepositEvent is a free log retrieval operation binding the contract event 0xc37633a15533e70debba779cecd6f0ac9d5ff599e6b5a6c2cfae4d3c62dd0338.
//
// Solidity: e ListingDepositEvent(listingHash indexed bytes32, added uint256, supply uint256, minted uint256, owner indexed address)
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

// WatchListingDepositEvent is a free log subscription operation binding the contract event 0xc37633a15533e70debba779cecd6f0ac9d5ff599e6b5a6c2cfae4d3c62dd0338.
//
// Solidity: e ListingDepositEvent(listingHash indexed bytes32, added uint256, supply uint256, minted uint256, owner indexed address)
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
	Subtracted  *big.Int
	Supply      *big.Int
	Minted      *big.Int
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterListingWithdrawEvent is a free log retrieval operation binding the contract event 0xda21bce65544aeea410c866bc9fe650b80110de14e8bb84bc79c50e44091f91e.
//
// Solidity: e ListingWithdrawEvent(listingHash indexed bytes32, subtracted uint256, supply uint256, minted uint256, owner indexed address)
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

// WatchListingWithdrawEvent is a free log subscription operation binding the contract event 0xda21bce65544aeea410c866bc9fe650b80110de14e8bb84bc79c50e44091f91e.
//
// Solidity: e ListingWithdrawEvent(listingHash indexed bytes32, subtracted uint256, supply uint256, minted uint256, owner indexed address)
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

// MarketRewardClaimedEventIterator is returned from FilterRewardClaimedEvent and is used to iterate over the raw logs and unpacked data for RewardClaimedEvent events raised by the Market contract.
type MarketRewardClaimedEventIterator struct {
	Event *MarketRewardClaimedEvent // Event containing the contract specifics and raw log

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
func (it *MarketRewardClaimedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketRewardClaimedEvent)
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
		it.Event = new(MarketRewardClaimedEvent)
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
func (it *MarketRewardClaimedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketRewardClaimedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketRewardClaimedEvent represents a RewardClaimedEvent event raised by the Market contract.
type MarketRewardClaimedEvent struct {
	Id     *big.Int
	Reward *big.Int
	Voter  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimedEvent is a free log retrieval operation binding the contract event 0x1444f3f25ac00f3a484d68ef4ce9dd0630c3f5943d4f05253dfabc4286643a63.
//
// Solidity: e RewardClaimedEvent(id indexed uint256, reward uint256, voter indexed address)
func (_Market *MarketFilterer) FilterRewardClaimedEvent(opts *bind.FilterOpts, id []*big.Int, voter []common.Address) (*MarketRewardClaimedEventIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "RewardClaimedEvent", idRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &MarketRewardClaimedEventIterator{contract: _Market.contract, event: "RewardClaimedEvent", logs: logs, sub: sub}, nil
}

// WatchRewardClaimedEvent is a free log subscription operation binding the contract event 0x1444f3f25ac00f3a484d68ef4ce9dd0630c3f5943d4f05253dfabc4286643a63.
//
// Solidity: e RewardClaimedEvent(id indexed uint256, reward uint256, voter indexed address)
func (_Market *MarketFilterer) WatchRewardClaimedEvent(opts *bind.WatchOpts, sink chan<- *MarketRewardClaimedEvent, id []*big.Int, voter []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "RewardClaimedEvent", idRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketRewardClaimedEvent)
				if err := _Market.contract.UnpackLog(event, "RewardClaimedEvent", log); err != nil {
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
const MarketTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stopMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"market\",\"type\":\"address\"}],\"name\":\"setMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintingStopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialAccount\",\"type\":\"address\"},{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ApprovalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BurnEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintStoppedEvent\",\"type\":\"event\"}]"

// MarketTokenBin is the compiled bytecode used for deploying new contracts.
const MarketTokenBin = `0x60806040526001805460ff1916905534801561001a57600080fd5b5060405160408061106b833981016040908152815160209283015160018054610100330261010060a860020a0319909116179055600160a060020a03909116600090815260039093529082208190559055610ff18061007a6000396000f3006080604052600436106100da5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b381146100df57806323b872dd146101175780633e3e0b121461014157806342966c681461015657806366188463146101705780636c9c2faf146101945780636dcea85f146101bb57806370a08231146101dc578063893d20e8146101fd578063a0712d681461022e578063a9059cbb14610246578063d73dd6231461026a578063dd62ed3e1461028e578063f1be1679146102b5578063f339292f146102ca575b600080fd5b3480156100eb57600080fd5b50610103600160a060020a03600435166024356102df565b604080519115158252519081900360200190f35b34801561012357600080fd5b50610103600160a060020a0360043581169060243516604435610345565b34801561014d57600080fd5b506101036105ff565b34801561016257600080fd5b5061016e600435610746565b005b34801561017c57600080fd5b50610103600160a060020a0360043516602435610901565b3480156101a057600080fd5b506101a96109f1565b60408051918252519081900360200190f35b3480156101c757600080fd5b50610103600160a060020a03600435166109f7565b3480156101e857600080fd5b506101a9600160a060020a0360043516610b40565b34801561020957600080fd5b50610212610b5b565b60408051600160a060020a039092168252519081900360200190f35b34801561023a57600080fd5b50610103600435610b6f565b34801561025257600080fd5b50610103600160a060020a0360043516602435610cf1565b34801561027657600080fd5b50610103600160a060020a0360043516602435610ea4565b34801561029a57600080fd5b506101a9600160a060020a0360043581169060243516610f3d565b3480156102c157600080fd5b50610212610f68565b3480156102d657600080fd5b50610103610f77565b336000818152600460209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e928290030190a350600192915050565b6000600160a060020a03831615156103cd576040805160e560020a62461bcd02815260206004820152603c60248201527f4572726f723a5374616e646172642e7472616e7366657246726f6d202d20277460448201527f6f272061646472657373206d7573742062652073706563696669656400000000606482015290519081900360840190fd5b600160a060020a038416600090815260036020526040902054821115610463576040805160e560020a62461bcd02815260206004820152603d60248201527f4572726f723a5374616e646172642e7472616e7366657246726f6d202d20566160448201527f6c7565206578636565647320617661696c61626c652062616c616e6365000000606482015290519081900360840190fd5b600160a060020a0384166000908152600460209081526040808320338452909152902054821115610504576040805160e560020a62461bcd02815260206004820152603a60248201527f4572726f722e5374616e646172642e7472616e7366657246726f6d202d20566160448201527f6c7565206578636565647320616c6c6f77656420616d6f756e74000000000000606482015290519081900360840190fd5b600160a060020a03841660009081526003602052604090205461052d908363ffffffff610f8016565b600160a060020a038086166000908152600360205260408082209390935590851681522054610562908363ffffffff610f9216565b600160a060020a0380851660009081526003602090815260408083209490945591871681526004825282812033825290915220546105a6908363ffffffff610f8016565b600160a060020a0380861660008181526004602090815260408083203384528252918290209490945580518681529051928716939192600080516020610fa6833981519152929181900390910190a35060019392505050565b600254600090600160a060020a0316331461068a576040805160e560020a62461bcd02815260206004820152603b60248201527f4572726f723a4d61726b6574546f6b656e2e69734d61726b6574202d2043616c60448201527f6c6572206d757374206265206d61726b657420636f6e74726163740000000000606482015290519081900360840190fd5b60015460ff161561070b576040805160e560020a62461bcd02815260206004820152603460248201527f4572726f723a4d61726b6574546f6b656e2e63616e4d696e74202d204d696e7460448201527f696e6720686173206265656e2073746f70706564000000000000000000000000606482015290519081900360840190fd5b6001805460ff1916811790556040517f6ae65cb1ae147362b8906c28687a17e1712ca926bea50710ba279849cec3347190600090a150600190565b600254600160a060020a031633146107ce576040805160e560020a62461bcd02815260206004820152603b60248201527f4572726f723a4d61726b6574546f6b656e2e69734d61726b6574202d2043616c60448201527f6c6572206d757374206265206d61726b657420636f6e74726163740000000000606482015290519081900360840190fd5b3360009081526003602052604090205481111561085b576040805160e560020a62461bcd02815260206004820152603960248201527f4572726f723a4e6574776f726b546f6b656e2e6275726e202d2056616c75652060448201527f6578636565647320617661696c61626c652062616c616e636500000000000000606482015290519081900360840190fd5b3360009081526003602052604090205461087b908263ffffffff610f8016565b336000908152600360205260408120919091555461089f908263ffffffff610f8016565b60005560408051828152905133917f512586160ebd4dc6945ba9ec5d21a1f723f26f3c7aa36cdffb6818d4e7b88030919081900360200190a26040805182815290516000913391600080516020610fa68339815191529181900360200190a350565b336000908152600460209081526040808320600160a060020a03861684529091528120548083111561095657336000908152600460209081526040808320600160a060020a038816845290915281205561098b565b610966818463ffffffff610f8016565b336000908152600460209081526040808320600160a060020a03891684529091529020555b336000818152600460209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e929181900390910190a35060019392505050565b60005490565b6001546000906101009004600160a060020a03163314610a87576040805160e560020a62461bcd02815260206004820152603060248201527f4572726f723a4d61726b6574546f6b656e2e69734f776e6572202d2043616c6c60448201527f6572206d757374206265206f776e657200000000000000000000000000000000606482015290519081900360840190fd5b600254600160a060020a031615610b0e576040805160e560020a62461bcd02815260206004820152603860248201527f4572726f723a4d61726b6574546f6b656e2e7365744d61726b6574202d204d6160448201527f726b6574206164647265737320616c7265616479207365740000000000000000606482015290519081900360840190fd5b5060028054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600160a060020a031660009081526003602052604090205490565b6001546101009004600160a060020a031690565b600254600090600160a060020a03163314610bfa576040805160e560020a62461bcd02815260206004820152603b60248201527f4572726f723a4d61726b6574546f6b656e2e69734d61726b6574202d2043616c60448201527f6c6572206d757374206265206d61726b657420636f6e74726163740000000000606482015290519081900360840190fd5b60015460ff1615610c7b576040805160e560020a62461bcd02815260206004820152603460248201527f4572726f723a4d61726b6574546f6b656e2e63616e4d696e74202d204d696e7460448201527f696e6720686173206265656e2073746f70706564000000000000000000000000606482015290519081900360840190fd5b600054610c8e908363ffffffff610f9216565b600090815533815260036020526040902054610cb0908363ffffffff610f9216565b336000818152600360209081526040808320949094558351868152935192939192600080516020610fa68339815191529281900390910190a3506001919050565b6000600160a060020a0383161515610d79576040805160e560020a62461bcd02815260206004820152603360248201527f4572726f723a42617369632e7472616e73666572202d20416e2061646472657360448201527f73206d7573742062652073706563696669656400000000000000000000000000606482015290519081900360840190fd5b33600090815260036020526040902054821115610e06576040805160e560020a62461bcd02815260206004820152603e60248201527f4572726f723a42617369632e7472616e73666572202d2056616c75652065786360448201527f65656473207468652062616c616e6365206f66206d73672e73656e6465720000606482015290519081900360840190fd5b33600090815260036020526040902054610e26908363ffffffff610f8016565b3360009081526003602052604080822092909255600160a060020a03851681522054610e58908363ffffffff610f9216565b600160a060020a038416600081815260036020908152604091829020939093558051858152905191923392600080516020610fa68339815191529281900390910190a350600192915050565b336000908152600460209081526040808320600160a060020a0386168452909152812054610ed8908363ffffffff610f9216565b336000818152600460209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e929081900390910190a350600192915050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205490565b600254600160a060020a031690565b60015460ff1690565b600082821115610f8c57fe5b50900390565b81810182811015610f9f57fe5b929150505600eaf1c4b3ce0f4f62a2bae7eb3e68225c75f7e6ff4422073b7437b9a78d25f170a165627a7a72305820a4a99ee9bbf5e18833f777f6c8f99614aa0e6777a4768a2f8222a99166b147b20029`

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

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_MarketToken *MarketTokenCaller) GetMarket(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "getMarket")
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_MarketToken *MarketTokenSession) GetMarket() (common.Address, error) {
	return _MarketToken.Contract.GetMarket(&_MarketToken.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xf1be1679.
//
// Solidity: function getMarket() constant returns(address)
func (_MarketToken *MarketTokenCallerSession) GetMarket() (common.Address, error) {
	return _MarketToken.Contract.GetMarket(&_MarketToken.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_MarketToken *MarketTokenCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_MarketToken *MarketTokenSession) GetOwner() (common.Address, error) {
	return _MarketToken.Contract.GetOwner(&_MarketToken.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_MarketToken *MarketTokenCallerSession) GetOwner() (common.Address, error) {
	return _MarketToken.Contract.GetOwner(&_MarketToken.CallOpts)
}

// GetSupply is a free data retrieval call binding the contract method 0x6c9c2faf.
//
// Solidity: function getSupply() constant returns(uint256)
func (_MarketToken *MarketTokenCaller) GetSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MarketToken.contract.Call(opts, out, "getSupply")
	return *ret0, err
}

// GetSupply is a free data retrieval call binding the contract method 0x6c9c2faf.
//
// Solidity: function getSupply() constant returns(uint256)
func (_MarketToken *MarketTokenSession) GetSupply() (*big.Int, error) {
	return _MarketToken.Contract.GetSupply(&_MarketToken.CallOpts)
}

// GetSupply is a free data retrieval call binding the contract method 0x6c9c2faf.
//
// Solidity: function getSupply() constant returns(uint256)
func (_MarketToken *MarketTokenCallerSession) GetSupply() (*big.Int, error) {
	return _MarketToken.Contract.GetSupply(&_MarketToken.CallOpts)
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

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MarketToken *MarketTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MarketToken *MarketTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Approve(&_MarketToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MarketToken *MarketTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Approve(&_MarketToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(value uint256) returns()
func (_MarketToken *MarketTokenTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(value uint256) returns()
func (_MarketToken *MarketTokenSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Burn(&_MarketToken.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(value uint256) returns()
func (_MarketToken *MarketTokenTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Burn(&_MarketToken.TransactOpts, value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, subtractedValue uint256) returns(bool)
func (_MarketToken *MarketTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "decreaseApproval", spender, subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, subtractedValue uint256) returns(bool)
func (_MarketToken *MarketTokenSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.DecreaseApproval(&_MarketToken.TransactOpts, spender, subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, subtractedValue uint256) returns(bool)
func (_MarketToken *MarketTokenTransactorSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.DecreaseApproval(&_MarketToken.TransactOpts, spender, subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, addedValue uint256) returns(bool)
func (_MarketToken *MarketTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "increaseApproval", spender, addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, addedValue uint256) returns(bool)
func (_MarketToken *MarketTokenSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.IncreaseApproval(&_MarketToken.TransactOpts, spender, addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, addedValue uint256) returns(bool)
func (_MarketToken *MarketTokenTransactorSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.IncreaseApproval(&_MarketToken.TransactOpts, spender, addedValue)
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

// SetMarket is a paid mutator transaction binding the contract method 0x6dcea85f.
//
// Solidity: function setMarket(market address) returns(bool)
func (_MarketToken *MarketTokenTransactor) SetMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "setMarket", market)
}

// SetMarket is a paid mutator transaction binding the contract method 0x6dcea85f.
//
// Solidity: function setMarket(market address) returns(bool)
func (_MarketToken *MarketTokenSession) SetMarket(market common.Address) (*types.Transaction, error) {
	return _MarketToken.Contract.SetMarket(&_MarketToken.TransactOpts, market)
}

// SetMarket is a paid mutator transaction binding the contract method 0x6dcea85f.
//
// Solidity: function setMarket(market address) returns(bool)
func (_MarketToken *MarketTokenTransactorSession) SetMarket(market common.Address) (*types.Transaction, error) {
	return _MarketToken.Contract.SetMarket(&_MarketToken.TransactOpts, market)
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
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MarketToken *MarketTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MarketToken *MarketTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Transfer(&_MarketToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MarketToken *MarketTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.Transfer(&_MarketToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MarketToken *MarketTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MarketToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MarketToken *MarketTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.TransferFrom(&_MarketToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MarketToken *MarketTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MarketToken.Contract.TransferFrom(&_MarketToken.TransactOpts, from, to, value)
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
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApprovalEvent is a free log retrieval operation binding the contract event 0x08245b82180b1f5e514e503c113ab0197093b2cb542145037c0a31b54b1d998e.
//
// Solidity: e ApprovalEvent(holder indexed address, spender indexed address, value uint256)
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
// Solidity: e ApprovalEvent(holder indexed address, spender indexed address, value uint256)
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
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnEvent is a free log retrieval operation binding the contract event 0x512586160ebd4dc6945ba9ec5d21a1f723f26f3c7aa36cdffb6818d4e7b88030.
//
// Solidity: e BurnEvent(burner indexed address, value uint256)
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
// Solidity: e BurnEvent(burner indexed address, value uint256)
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
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferEvent is a free log retrieval operation binding the contract event 0xeaf1c4b3ce0f4f62a2bae7eb3e68225c75f7e6ff4422073b7437b9a78d25f170.
//
// Solidity: e TransferEvent(from indexed address, to indexed address, value uint256)
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
// Solidity: e TransferEvent(from indexed address, to indexed address, value uint256)
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

// PLCRVotingABI is the input ABI used to generate the binding from.
const PLCRVotingABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getTotalNumberOfTokensForWinningOption\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getPollRevealExpiry\",\"outputs\":[{\"name\":\"expiry\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"numTokens\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getInsertPointForNumTokens\",\"outputs\":[{\"name\":\"prevNode\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"name\":\"commitDuration\",\"type\":\"uint256\"},{\"name\":\"revealDuration\",\"type\":\"uint256\"}],\"name\":\"startPoll\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"revealPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isPassed\",\"outputs\":[{\"name\":\"passed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getLockedTokens\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"secretHash\",\"type\":\"bytes32\"},{\"name\":\"numTokens\",\"type\":\"uint256\"},{\"name\":\"prevId\",\"type\":\"uint256\"}],\"name\":\"commitVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"didCommit\",\"outputs\":[{\"name\":\"committed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"prevId\",\"type\":\"uint256\"},{\"name\":\"nextId\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"validPosition\",\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"pollExists\",\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"rescueTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"requestVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"commitPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"didReveal\",\"outputs\":[{\"name\":\"revealed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"voteOption\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"revealVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getNumPassingTokens\",\"outputs\":[{\"name\":\"correctVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getNumTokens\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getCommitHash\",\"outputs\":[{\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"terminationDate\",\"type\":\"uint256\"}],\"name\":\"isExpired\",\"outputs\":[{\"name\":\"expired\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getPollCommitExpiry\",\"outputs\":[{\"name\":\"expiry\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"withdrawVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"pollEnded\",\"outputs\":[{\"name\":\"ended\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VoteCommittedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesFor\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesAgainst\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"choice\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VoteRevealedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commitExpiry\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealExpiry\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"PollCreatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VotingRightsGrantedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VotingRightsWithdrawnEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"TokensRescuedEvent\",\"type\":\"event\"}]"

// PLCRVotingBin is the compiled bytecode used for deploying new contracts.
const PLCRVotingBin = `0x608060405234801561001057600080fd5b5060405160208061255d833981016040525160058054600160a060020a03909216600160a060020a03199092169190911790556000600255612506806100576000396000f3006080604052600436106101325763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663053e71a681146101375780631be8ea50146101615780632c0520311461017957806332ed3d60146101a0578063441c77c0146101be57806349403183146101ea5780636b2d95d4146102025780636cbf9c5e146102235780637f97e83614610246578063819b02931461026a57806388d21ff31461029457806397603560146102ac578063a25236fe146102c4578063a4439dc5146102dc578063aa7ca464146102f4578063b11d8bb814610318578063b43bd06914610336578063d13820921461035d578063d901402b14610381578063d9548e53146103a5578063dc1b946f146103bd578063e7b1d43c146103d5578063ee684830146103ed575b600080fd5b34801561014357600080fd5b5061014f600435610405565b60408051918252519081900360200190f35b34801561016d57600080fd5b5061014f6004356104f2565b34801561018557600080fd5b5061014f600160a060020a036004351660243560443561058f565b3480156101ac57600080fd5b5061014f600435602435604435610600565b3480156101ca57600080fd5b506101d66004356106f0565b604080519115158252519081900360200190f35b3480156101f657600080fd5b506101d66004356107ab565b34801561020e57600080fd5b5061014f600160a060020a03600435166108ce565b34801561022f57600080fd5b506102446004356024356044356064356108e2565b005b34801561025257600080fd5b506101d6600160a060020a0360043516602435610ca7565b34801561027657600080fd5b506101d6600435602435600160a060020a0360443516606435610d5d565b3480156102a057600080fd5b506101d6600435610d9e565b3480156102b857600080fd5b50610244600435610db3565b3480156102d057600080fd5b50610244600435610f0a565b3480156102e857600080fd5b506101d66004356111c5565b34801561030057600080fd5b506101d6600160a060020a0360043516602435611264565b34801561032457600080fd5b5061024460043560243560443561131a565b34801561034257600080fd5b5061014f600160a060020a0360043516602435604435611719565b34801561036957600080fd5b5061014f600160a060020a03600435166024356119a0565b34801561038d57600080fd5b5061014f600160a060020a03600435166024356119ef565b3480156103b157600080fd5b506101d6600435611a37565b3480156103c957600080fd5b5061014f600435611a3c565b3480156103e157600080fd5b50610244600435611ad6565b3480156103f957600080fd5b506101d6600435611d29565b600061041082611d29565b15156104b2576040805160e560020a62461bcd02815260206004820152604a60248201527f4572726f723a566f74696e672e676574546f74616c4e756d6265724f66546f6b60448201527f656e73466f7257696e6e696e674f7074696f6e202d20506f6c6c206d7573742060648201527f6861766520656e64656400000000000000000000000000000000000000000000608482015290519081900360a40190fd5b6104bb826107ab565b156104d95750600081815260036020819052604090912001546104ed565b506000818152600360205260409020600401545b919050565b60006104fd82610d9e565b1515610579576040805160e560020a62461bcd02815260206004820152603260248201527f4572726f723a566f74696e672e676574506f6c6c72657665616c45787069727960448201527f202d20506f6c6c206d7573742065786973740000000000000000000000000000606482015290519081900360840190fd5b5060009081526003602052604090206001015490565b600080600061059d86611dcb565b91506105a986836119a0565b90505b81156105f3576105bc86836119a0565b90508481116105e257838214156105da576105d78683611dd4565b91505b8192506105f7565b6105ec8683611dd4565b91506105ac565b8192505b50509392505050565b600080600061061b6001600254611dff90919063ffffffff16565b60025561062e428663ffffffff611dff16565b9150610640828563ffffffff611dff16565b6040805160a08101825284815260208082018481528284018b815260006060808601828152608087018381526002805485526003808952948a90209851895595516001890155935187860155519186019190915590516004909401939093555483518b8152918201879052818401859052925193945033937ff7e7200f3afa7bbc8aca2cfa0ab33e741504a20fed2abc59869413cba2a33400929181900390910190a35050600254949350505050565b60006106fb82610d9e565b1515610777576040805160e560020a62461bcd02815260206004820152603160248201527f4572726f723a566f74696e672e72657665616c506572696f644163746976652060448201527f2d20506f6c6c206d757374206578697374000000000000000000000000000000606482015290519081900360840190fd5b60008281526003602052604090206001015461079290611a37565b1580156107a557506107a3826111c5565b155b92915050565b60006107b561246a565b6107be83611d29565b151561083a576040805160e560020a62461bcd02815260206004820152602c60248201527f4572726f723a566f74696e672e6973506173736564202d20506f6c6c206d757360448201527f74206861766520656e6465640000000000000000000000000000000000000000606482015290519081900360840190fd5b50600082815260036020818152604092839020835160a0810185528154815260018201549281019290925260028101549382019390935290820154606082018190526004909201546080820181905290916108b09161089f919063ffffffff611dff16565b60408301519063ffffffff611e0c16565b60608201516108c690606463ffffffff611e0c16565b119392505050565b60006107a5826108dd84611dcb565b6119a0565b6000806108ee866111c5565b151561096a576040805160e560020a62461bcd02815260206004820152603460248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d2054617267657460448201527f20706f6c6c206d75737420626520616374697665000000000000000000000000606482015290519081900360840190fd5b336000908152600460205260409020548411156109f7576040805160e560020a62461bcd02815260206004820152603160248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d2053656e64657260448201527f2063616e6e6f74206f7665727370656e64000000000000000000000000000000606482015290519081900360840190fd5b851515610a74576040805160e560020a62461bcd02815260206004820152602d60248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d20506f6c6c204960448201527f442063616e6e6f74206265203000000000000000000000000000000000000000606482015290519081900360840190fd5b821580610a865750610a863384611e35565b1515610b02576040805160e560020a62461bcd02815260206004820152602960248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d20506f6c6c206d60448201527f7573742065786973740000000000000000000000000000000000000000000000606482015290519081900360840190fd5b610b0c3384611eb3565b915085821415610b2357610b203387611eb3565b91505b610b2f83833387610d5d565b1515610bab576040805160e560020a62461bcd02815260206004820152603460248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d20506f6c6c207060448201527f6f736974696f6e206973206e6f742076616c6964000000000000000000000000606482015290519081900360840190fd5b610bb733848885611edb565b610bc1338761213f565b9050610c03816040805190810160405280600981526020017f6e756d546f6b656e730000000000000000000000000000000000000000000000815250866121ec565b60408051808201909152600a81527f636f6d6d697448617368000000000000000000000000000000000000000000006020820152610c43908290876121ec565b6000868152600360209081526040808320338085526005909101835292819020805460ff191660011790558051878152905189927f673f4a020a46f562a6e9fee9c8d4a179a11924190b777bec160133dd00c9f1de928290030190a3505050505050565b6000610cb282610d9e565b1515610d2e576040805160e560020a62461bcd02815260206004820152602f60248201527f4572726f723a506172616d65746572697a65722e646964436f6d6d6974202d2060448201527f506f6c6c206d7573742065786973740000000000000000000000000000000000606482015290519081900360840190fd5b506000908152600360209081526040808320600160a060020a0394909416835260059093019052205460ff1690565b6000806000610d6c85886119a0565b8410159150610d7b85876119a0565b84111580610d87575085155b9050818015610d935750805b979650505050505050565b600081158015906107a5575050600254101590565b600081815260036020526040902060010154610dce90611a37565b1515610e4a576040805160e560020a62461bcd02815260206004820152603060248201527f4572726f723a566f74696e672e726573637565546f6b656e73202d20506f6c6c60448201527f206d757374206861766520656e64656400000000000000000000000000000000606482015290519081900360840190fd5b610e543382611e35565b1515610ed0576040805160e560020a62461bcd02815260206004820152602a60248201527f4572726f723a566f74696e672e726573637565546f6b656e73202d20506f6c6c60448201527f206e6f7420666f756e6400000000000000000000000000000000000000000000606482015290519081900360840190fd5b610eda33826122d0565b604051339082907f1e897e4fe81e0db0459f0219545995fde305b17af5ca5cba7d6b6e8738df5dd090600090a350565b600554604080517f70a0823100000000000000000000000000000000000000000000000000000000815233600482015290518392600160a060020a0316916370a082319160248083019260209291908290030181600087803b158015610f6f57600080fd5b505af1158015610f83573d6000803e3d6000fd5b505050506040513d6020811015610f9957600080fd5b5051101561103d576040805160e560020a62461bcd02815260206004820152604860248201527f4572726f723a566f74696e672e72657175657374566f74696e6752696768747360448201527f202d2053656e64657220646f6573206e6f74206861766520737566666963696560648201527f6e742066756e6473000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b3360009081526004602052604090205461105d908263ffffffff611dff16565b3360008181526004602081815260408084209590955560055485517f23b872dd00000000000000000000000000000000000000000000000000000000815292830194909452306024830152604482018690529351600160a060020a03909316936323b872dd9360648084019492939192918390030190829087803b1580156110e457600080fd5b505af11580156110f8573d6000803e3d6000fd5b505050506040513d602081101561110e57600080fd5b5051151561118c576040805160e560020a62461bcd02815260206004820152603d60248201527f4572726f723a566f74696e672e72657175657374566f74696e6752696768747360448201527f202d20546f6b656e2e7472616e7366657246726f6d206661696c757265000000606482015290519081900360840190fd5b60408051828152905133917f0fbea84792de88539ad3d7f184543a65916440f4c6f9e2ef136c97f13109c45d919081900360200190a250565b60006111d082610d9e565b151561124c576040805160e560020a62461bcd02815260206004820152603160248201527f4572726f723a566f74696e672e636f6d6d6974506572696f644163746976652060448201527f2d20506f6c6c206d757374206578697374000000000000000000000000000000606482015290519081900360840190fd5b6000828152600360205260409020546107a390611a37565b600061126f82610d9e565b15156112eb576040805160e560020a62461bcd02815260206004820152602860248201527f4572726f723a566f74696e672e64696452657665616c202d20506f6c6c206d7560448201527f7374206578697374000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b506000908152600360209081526040808320600160a060020a0394909416835260069093019052205460ff1690565b6000611325846106f0565b15156113a1576040805160e560020a62461bcd02815260206004820152603660248201527f4572726f723a566f74696e672e72657665616c566f7465202d2052657665616c60448201527f20706572696f64206d7573742062652061637469766500000000000000000000606482015290519081900360840190fd5b600084815260036020908152604080832033845260050190915290205460ff16151561143d576040805160e560020a62461bcd02815260206004820152603f60248201527f4572726f723a566f74696e672e72657665616c566f7465202d2053656e64657260448201527f206d757374206861766520636f6d6d697474656420746865697220766f746500606482015290519081900360840190fd5b600084815260036020908152604080832033845260060190915290205460ff16156114fe576040805160e560020a62461bcd02815260206004820152604860248201527f4572726f723a566f74696e672e72657665616c566f7465202d2053656e64657260448201527f2063616e6e6f742072657665616c20746865697220766f7465206d756c74697060648201527f6c652074696d6573000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b61150833856119ef565b604080516020808201879052818301869052825180830384018152606090920192839052815191929182918401908083835b602083106115595780518252601f19909201916020918201910161153a565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916141515611607576040805160e560020a62461bcd02815260206004820152602760248201527f4572726f723a566f74696e672e72657665616c566f7465202d2048617368206d60448201527f69736d6174636800000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b61161133856119a0565b90508260011415611657576000848152600360208190526040909120015461163f908263ffffffff611dff16565b6000858152600360208190526040909120015561168c565b600084815260036020526040902060040154611679908263ffffffff611dff16565b6000858152600360205260409020600401555b61169633856122d0565b600084815260036020818152604080842033808652600682018452828620805460ff191660011790559489905283835292830154600490930154815186815292830193909352818101929092529051859187917f448a2036d0f6dc799c1d3fdf26c4fe80832817bdbe81133808ffa6777c5826e09181900360600190a450505050565b60008060008061172886611d29565b15156117a4576040805160e560020a62461bcd02815260206004820152603760248201527f4572726f723a566f74696e672e6765744e756d50617373696e67546f6b656e7360448201527f202d20506f6c6c206d757374206861766520656e646564000000000000000000606482015290519081900360840190fd5b6000868152600360209081526040808320600160a060020a038b16845260060190915290205460ff16151561186f576040805160e560020a62461bcd02815260206004820152604660248201527f4572726f723a566f74696e672e6765744e756d50617373696e67546f6b656e7360448201527f202d20566f746572206d75737420686176652072657665616c6564207468656960648201527f7220766f74650000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b611878866107ab565b611883576000611886565b60015b60ff169250828560405160200180838152602001828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106118e25780518252601f1990920191602091820191016118c3565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020915061191b87876119ef565b905081811461199a576040805160e560020a62461bcd02815260206004820152603060248201527f4572726f723a566f74696e672e6765744e756d50617373696e67546f6b656e7360448201527f202d2048617368206d69736d6174636800000000000000000000000000000000606482015290519081900360840190fd5b610d9387875b60006119e86119af848461213f565b60408051808201909152600981527f6e756d546f6b656e7300000000000000000000000000000000000000000000006020820152612343565b9392505050565b60006119e86119fe848461213f565b60408051808201909152600a81527f636f6d6d697448617368000000000000000000000000000000000000000000006020820152612343565b421190565b6000611a4782610d9e565b1515611ac3576040805160e560020a62461bcd02815260206004820152603260248201527f4572726f723a566f74696e672e676574506f6c6c436f6d6d697445787069727960448201527f202d20506f6c6c206d7573742065786973740000000000000000000000000000606482015290519081900360840190fd5b5060009081526003602052604090205490565b6000611b00611ae4336108ce565b336000908152600460205260409020549063ffffffff61242d16565b905081811015611ba6576040805160e560020a62461bcd02815260206004820152604160248201527f4572726f723a566f74696e672e7769746864726177566f74696e67526967687460448201527f73202d20496e73756666696369656e7420617661696c61626c6520746f6b656e60648201527f7300000000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b33600090815260046020526040902054611bc6908363ffffffff61242d16565b3360008181526004602081815260408084209590955560055485517fa9059cbb00000000000000000000000000000000000000000000000000000000815292830194909452602482018790529351600160a060020a039093169363a9059cbb9360448084019492939192918390030190829087803b158015611c4757600080fd5b505af1158015611c5b573d6000803e3d6000fd5b505050506040513d6020811015611c7157600080fd5b50511515611cef576040805160e560020a62461bcd02815260206004820152603a60248201527f4572726f723a566f74696e672e7769746864726177566f74696e67526967687460448201527f73202d20546f6b656e2e7472616e73666572206661696c757265000000000000606482015290519081900360840190fd5b60408051838152905133917f029a58e9abaa250a4ec75d14b364d764e890adec2d902a297fa24a6c5dfd097b919081900360200190a25050565b6000611d3482610d9e565b1515611db0576040805160e560020a62461bcd02815260206004820152602860248201527f4572726f723a566f74696e672e706f6c6c456e646564202d20506f6c6c206d7560448201527f7374206578697374000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000828152600360205260409020600101546107a590611a37565b60006107a58260005b600160a060020a03919091166000908152600160208181526040808420948452939052919020015490565b818101828110156107a557fe5b6000821515611e1d575060006107a5565b50818102818382811515611e2d57fe5b04146107a557fe5b6000806000611e438561243f565b80611e4c575083155b15611e5a5760009250611eab565b83611e6486612452565b148015611e78575083611e768661245e565b145b91506000611e868686611eb3565b148015611e9c57506000611e9a8686611dd4565b145b90508180611ea8575080155b92505b505092915050565b600160a060020a03919091166000908152600160209081526040808320938352929052205490565b811515611f58576040805160e560020a62461bcd02815260206004820152603a60248201527f4572726f723a504c4352566f74696e672e646c6c496e73657274202d2049442060448201527f6f662061206e6577206e6f6465206d6179206e6f742062652030000000000000606482015290519081900360840190fd5b611f6284836122d0565b821580611f745750611f748484611e35565b1515611fcc576040805160e560020a62461bcd02815260206004820152603360248201526000805160206124bb833981519152604482015260008051602061249b833981519152606482015290519081900360840190fd5b801580611fde5750611fde8482611e35565b1515612036576040805160e560020a62461bcd02815260206004820152603360248201526000805160206124bb833981519152604482015260008051602061249b833981519152606482015290519081900360840190fd5b806120418585611eb3565b14612098576040805160e560020a62461bcd02815260206004820152603360248201526000805160206124bb833981519152604482015260008051602061249b833981519152606482015290519081900360840190fd5b826120a38583611dd4565b146120fa576040805160e560020a62461bcd02815260206004820152603360248201526000805160206124bb833981519152604482015260008051602061249b833981519152606482015290519081900360840190fd5b600160a060020a039390931660009081526001602081815260408084208585529091528083208083018690558690559382528382208390559381529190912090910155565b600082826040516020018083600160a060020a0316600160a060020a03166c01000000000000000000000000028152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106121b95780518252601f19909201916020918201910161219a565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091209695505050505050565b604080516020808201868152855160009488948894910191908401908083835b6020831061222b5780518252601f19909201916020918201910161220c565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b6020831061228f5780518252601f199092019160209182019101612270565b51815160209384036101000a600019018019909216911617905260408051929094018290039091206000908152908190529190912094909455505050505050565b6000806122dd8484611e35565b15156122e85761233d565b6122f28484611eb3565b91506122fe8484611dd4565b600160a060020a038516600090815260016020818152604080842087855290915280832082018490558383528083208690558683528220828155015590505b50505050565b600080838360405160200180836000191660001916815260200182805190602001908083835b602083106123885780518252601f199092019160209182019101612369565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b602083106123ec5780518252601f1990920191602091820191016123cd565b51815160209384036101000a600019018019909216911617905260408051929094018290039091206000908152908190529190912054979650505050505050565b60008282111561243957fe5b50900390565b60008061244b83612452565b1492915050565b60006107a58282611eb3565b60006107a58282611dd4565b60a0604051908101604052806000815260200160008152602001600081526020016000815260200160008152509056006e6f7420696e7365727420696e746f20444c4c000000000000000000000000004572726f723a504c4352566f74696e672e644c4c496e73657274202d20436f6ea165627a7a72305820be196c13f7bf098e733064d4253e587de19ae43c0b5fe490863e8052663adb9d0029`

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
const ParameterizerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"propHash\",\"type\":\"bytes32\"}],\"name\":\"processProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"propHash\",\"type\":\"bytes32\"}],\"name\":\"propExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"challengeWinnerReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"propHash\",\"type\":\"bytes32\"}],\"name\":\"challengeCanBeResolved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"propHash\",\"type\":\"bytes32\"}],\"name\":\"challengeReparameterization\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"rewardsClaimed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"voterReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"proposeReparameterization\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"propHash\",\"type\":\"bytes32\"}],\"name\":\"canBeSet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"},{\"name\":\"votingAddr\",\"type\":\"address\"},{\"name\":\"minDeposit\",\"type\":\"uint256\"},{\"name\":\"applyStageLen\",\"type\":\"uint256\"},{\"name\":\"commitStageLen\",\"type\":\"uint256\"},{\"name\":\"revealStageLen\",\"type\":\"uint256\"},{\"name\":\"dispensationPct\",\"type\":\"uint256\"},{\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"name\":\"listReward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"appExpiry\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"ReparameterizationProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commitExpiry\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealExpiry\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"NewChallengeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ProposalAcceptedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalExpiredEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"name\":\"ChallengeSucceededEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"propHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"name\":\"ChallengeFailedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"RewardsClaimedEvent\",\"type\":\"event\"}]"

// ParameterizerBin is the compiled bytecode used for deploying new contracts.
const ParameterizerBin = `0x608060405262093a806005553480156200001857600080fd5b506040516101208062002d878339810160408181528251602080850151838601516060870151608088015160a089015160c08a015160e08b0151610100909b015160038054600160a060020a03808c16600160a060020a03199283161790925560048054928a1692909116919091179055898b01909952600a8a527f6d696e4465706f73697400000000000000000000000000000000000000000000968a01969096529598939792969195909490939092909190620000e19088640100000000620002a6810204565b60408051808201909152600d81527f6170706c7953746167654c656e0000000000000000000000000000000000000060208201526200012a9087640100000000620002a6810204565b60408051808201909152600e81527f636f6d6d697453746167654c656e0000000000000000000000000000000000006020820152620001739086640100000000620002a6810204565b60408051808201909152600e81527f72657665616c53746167654c656e0000000000000000000000000000000000006020820152620001bc9085640100000000620002a6810204565b60408051808201909152600f81527f64697370656e736174696f6e50637400000000000000000000000000000000006020820152620002059084640100000000620002a6810204565b60408051808201909152600a81527f766f746551756f72756d0000000000000000000000000000000000000000000060208201526200024e9083640100000000620002a6810204565b60408051808201909152600a81527f6c697374526577617264000000000000000000000000000000000000000000006020820152620002979082640100000000620002a6810204565b50505050505050505062000386565b80600080846040516020018082805190602001908083835b60208310620002df5780518252601f199092019160209182019101620002be565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b60208310620003445780518252601f19909201916020918201910162000323565b51815160209384036101000a60001901801990921691161790526040805192909401829003909120865285019590955292909201600020939093555050505050565b6129f180620003966000396000f3006080604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166330490e9181146100b357806335300990146100cd57806350411552146100f9578063693ec85e1461012357806377609a411461017c5780638240ae4b1461019457806386bb8f37146101ac5780639a835e91146101c7578063a7aad3db146101eb578063bade1c5414610212578063c51131fb14610236575b600080fd5b3480156100bf57600080fd5b506100cb60043561024e565b005b3480156100d957600080fd5b506100e5600435610814565b604080519115158252519081900360200190f35b34801561010557600080fd5b5061011160043561082d565b60408051918252519081900360200190f35b34801561012f57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261011194369492936024939284019190819084018382808284375094975061092c9650505050505050565b34801561018857600080fd5b506100e5600435610a08565b3480156101a057600080fd5b50610111600435610c32565b3480156101b857600080fd5b506100cb600435602435611336565b3480156101d357600080fd5b506100e5600435600160a060020a036024351661172b565b3480156101f757600080fd5b50610111600160a060020a036004351660243560443561175a565b34801561021e57600080fd5b50610111602460048035828101929101359035611836565b34801561024257600080fd5b506100e5600435612073565b60008181526002602081905260409091206004810154918101549091600160a060020a03169061027d84612073565b1561054857600383018054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815261032093909290918301828280156103115780601f106102e657610100808354040283529160200191610311565b820191906000526020600020905b8154815290600101906020018083116102f457829003601f168201915b50505050508460060154612196565b600683015460408051602081018390528181526003860180546002610100600183161502600019019091160492820183905287937f0fca8490116d823913d2468dc21632e9226d20e54379af9cf34712a3ff2bc30c939192909181906060820190859080156103d05780601f106103a5576101008083540402835291602001916103d0565b820191906000526020600020905b8154815290600101906020018083116103b357829003601f168201915b5050935050505060405180910390a260008481526002602081905260408220828155600181018390559081018290559061040d6003830182612841565b506004818101805473ffffffffffffffffffffffffffffffffffffffff1916905560006005830181905560069092018290556003546040805160e060020a63a9059cbb028152600160a060020a0387811694820194909452602481018690529051929091169263a9059cbb926044808401936020939083900390910190829087803b15801561049b57600080fd5b505af11580156104af573d6000803e3d6000fd5b505050506040513d60208110156104c557600080fd5b50511515610543576040805160e560020a62461bcd02815260206004820152603e60248201527f4572726f723a506172616d65746572697a65722e70726f6365737350726f706f60448201527f73616c202d20436f756c64206e6f74207472616e736665722066756e64730000606482015290519081900360840190fd5b610686565b61055184610a08565b1561055f5761054384612272565b82600501544211156105c45760405184907f456cbba32d3f5372c3ef95ba54f18060130b9da12721ebf89f1146a8f406132490600090a260008481526002602081905260408220828155600181018390559081018290559061040d6003830182612841565b6040805160e560020a62461bcd02815260206004820152606c60248201527f5468657265206973206e6f206368616c6c656e676520616761696e737420746860448201527f652070726f706f73616c2c20616e64206e65697468657220746865206170704560648201527f78706972792064617465206e6f72207468652070726f6365737342792064617460848201527f652068617320706173736564000000000000000000000000000000000000000060a482015290519081900360c40190fd5b60646106c66040805190810160405280600f81526020017f64697370656e736174696f6e506374000000000000000000000000000000000081525061092c565b11156106ce57fe5b6107b06005546107a46107156040805190810160405280600e81526020017f72657665616c53746167654c656e00000000000000000000000000000000000081525061092c565b6107a46107566040805190810160405280600e81526020017f636f6d6d697453746167654c656e00000000000000000000000000000000000081525061092c565b6107a46107976040805190810160405280600d81526020017f6170706c7953746167654c656e0000000000000000000000000000000000000081525061092c565b429063ffffffff6127e416565b9063ffffffff6127e416565b506000848152600260208190526040822082815560018101839055908101829055906107df6003830182612841565b5060048101805473ffffffffffffffffffffffffffffffffffffffff1916905560006005820181905560069091015550505050565b600081815260026020526040812060050154115b919050565b60048054604080517f053e71a600000000000000000000000000000000000000000000000000000000815292830184905251600092600160a060020a039092169163053e71a691602480830192602092919082900301818787803b15801561089457600080fd5b505af11580156108a8573d6000803e3d6000fd5b505050506040513d60208110156108be57600080fd5b505115156108f05760008281526001602052604090206002908101546108e99163ffffffff6127f116565b9050610828565b600082815260016020526040902080546002918201546109269261091a919063ffffffff6127f116565b9063ffffffff61281a16565b92915050565b6000806000836040516020018082805190602001908083835b602083106109645780518252601f199092019160209182019101610945565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106109c75780518252601f1990920191602091820191016109a8565b51815160209384036101000a600019018019909216911617905260408051929094018290039091208652850195909552929092016000205495945050505050565b6000610a12612888565b610a1a6128cf565b600084815260026020818152604092839020835160e0810185528154815260018083015482850152828501548287015260038301805487516000199382161561010002939093011695909504601f810185900485028201850190965285815290949193606086019391929091830182828015610ad75780601f10610aac57610100808354040283529160200191610ad7565b820191906000526020600020905b815481529060010190602001808311610aba57829003601f168201915b50505091835250506004820154600160a060020a039081166020808401919091526005840154604080850191909152600690940154606093840152848101805160009081526001808452868220875160a0810189528154815291810154958616948201949094527401000000000000000000000000000000000000000090940460ff1615159584019590955260028201549383019390935260030154608082015290519294509250108015610b8e57506040810151155b8015610c2a575060048054602080850151604080517fee6848300000000000000000000000000000000000000000000000000000000081529485019190915251600160a060020a039092169263ee68483092602480830193928290030181600087803b158015610bfd57600080fd5b505af1158015610c11573d6000803e3d6000fd5b505050506040513d6020811015610c2757600080fd5b50515b949350505050565b6000610c3c612888565b6000838152600260208181526040808420815160e0810183528154815260018083015482860152828601548285015260038301805485516000199382161561010002939093011696909604601f810186900486028201860190945283815286958695869593949360608601939291830182828015610cfb5780601f10610cd057610100808354040283529160200191610cfb565b820191906000526020600020905b815481529060010190602001808311610cde57829003601f168201915b50505091835250506004820154600160a060020a0316602082015260058201546040808301919091526006909201546060909101528101519095509350610d4187610814565b8015610d4f57506020850151155b1515610df1576040805160e560020a62461bcd02815260206004820152604960248201527f4572726f723a506172616d65746572697a65722e6368616c6c656e676552657060448201527f6172616d65746572697a6174696f6e202d2050726f706f73616c20646f65732060648201527f6e6f742065786973740000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b60045460408051808201909152600a81527f766f746551756f72756d000000000000000000000000000000000000000000006020820152600160a060020a03909116906332ed3d6090610e439061092c565b610e816040805190810160405280600e81526020017f636f6d6d697453746167654c656e00000000000000000000000000000000000081525061092c565b610ebf6040805190810160405280600e81526020017f72657665616c53746167654c656e00000000000000000000000000000000000081525061092c565b6040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808481526020018381526020018281526020019350505050602060405180830381600087803b158015610f1f57600080fd5b505af1158015610f33573d6000803e3d6000fd5b505050506040513d6020811015610f4957600080fd5b50516040805160e08101909152600f60a082019081527f64697370656e736174696f6e506374000000000000000000000000000000000060c0830152919450908190610fc490606490610fb8908990610fac908490610fa79061092c565b61281a565b9063ffffffff6127f116565b9063ffffffff61282c16565b81523360208083018290526000604080850182905260608086018b905260809586018390528983526001808552828420885181558886015181830180548b87015173ffffffffffffffffffffffffffffffffffffffff19909116600160a060020a039384161774ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000911515919091021790559289015160028083019190915598909701516003978801558e8452968452818320909601899055935484517f23b872dd0000000000000000000000000000000000000000000000000000000081526004810194909452306024850152604484018a9052935193909416936323b872dd936064808501949192918390030190829087803b1580156110f157600080fd5b505af1158015611105573d6000803e3d6000fd5b505050506040513d602081101561111b57600080fd5b505115156111bf576040805160e560020a62461bcd02815260206004820152604a60248201527f4572726f723a506172616d65746572697a65722e6368616c6c656e676552657060448201527f6172616d65746572697a6174696f6e202d20436f756c64206e6f74207472616e60648201527f736665722066756e647300000000000000000000000000000000000000000000608482015290519081900360a40190fd5b60048054604080517fdc1b946f00000000000000000000000000000000000000000000000000000000815292830186905251600160a060020a039091169163dc1b946f9160248083019260209291908290030181600087803b15801561122457600080fd5b505af1158015611238573d6000803e3d6000fd5b505050506040513d602081101561124e57600080fd5b505160048054604080517f1be8ea5000000000000000000000000000000000000000000000000000000000815292830187905251929450600160a060020a031691631be8ea50916024808201926020929091908290030181600087803b1580156112b757600080fd5b505af11580156112cb573d6000803e3d6000fd5b505050506040513d60208110156112e157600080fd5b505160408051858152602081018590528082018390529051919250339189917f4e6815e2c453c1de363fd2b483078339512560f8fd777420d06dd69291e060c2919081900360600190a3509095945050505050565b6000828152600160209081526040808320338452600401909152812054819060ff16156113f9576040805160e560020a62461bcd02815260206004820152604260248201527f4572726f723a506172616d65746572697a65722e636c61696d5265776172642060448201527f2d20546f6b656e73206861766520616c7265616479206265656e20636c61696d60648201527f6564000000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b600084815260016020819052604090912081015474010000000000000000000000000000000000000000900460ff161515146114cb576040805160e560020a62461bcd02815260206004820152604f60248201527f4572726f723a506172616d65746572697a65722e636c61696d5265776172642060448201527f2d204368616c6c656e676520726573756c74732068617665206e6f742062656560648201527f6e2070726f636573736564207965740000000000000000000000000000000000608482015290519081900360a40190fd5b60048054604080517fb43bd0690000000000000000000000000000000000000000000000000000000081523393810193909352602483018790526044830186905251600160a060020a039091169163b43bd0699160648083019260209291908290030181600087803b15801561154057600080fd5b505af1158015611554573d6000803e3d6000fd5b505050506040513d602081101561156a57600080fd5b5051915061157933858561175a565b60008581526001602052604090206003015490915061159e908363ffffffff61281a16565b60008581526001602052604090206003810191909155546115c5908263ffffffff61281a16565b6000858152600160208181526040808420948555338085526004909501825292839020805460ff19169092179091558151848152915187927f17ccb5aed827c5078139dc28b7fb924b006f56b3dcf163c4acf21fa3beb1314492908290030190a36003546040805160e060020a63a9059cbb028152336004820152602481018490529051600160a060020a039092169163a9059cbb916044808201926020929091908290030181600087803b15801561167d57600080fd5b505af1158015611691573d6000803e3d6000fd5b505050506040513d60208110156116a757600080fd5b50511515611725576040805160e560020a62461bcd02815260206004820152603a60248201527f4572726f723a506172616d65746572697a65722e636c61696d5265776172642060448201527f2d20436f756c64206e6f74207472616e736665722066756e6473000000000000606482015290519081900360840190fd5b50505050565b6000828152600160209081526040808320600160a060020a038516845260040190915290205460ff1692915050565b6000828152600160209081526040808320600381015490546004805484517fb43bd069000000000000000000000000000000000000000000000000000000008152600160a060020a038b811693820193909352602481018a9052604481018990529451939592948794929091169263b43bd0699260648084019382900301818787803b1580156117e957600080fd5b505af11580156117fd573d6000803e3d6000fd5b505050506040513d602081101561181357600080fd5b5051905061182b83610fb8838563ffffffff6127f116565b979650505050505050565b60008060006118796040805190810160405280600a81526020017f6d696e4465706f7369740000000000000000000000000000000000000000000081525061092c565b915085858560405160200180848480828437820191505082815260200193505050506040516020818303038152906040526040518082805190602001908083835b602083106118d95780518252601f1990920191602091820191016118ba565b51815160209384036101000a6000190180199092169116179052604080519290940182900382207f64697370656e736174696f6e5063740000000000000000000000000000000000838301528451600f818503018152602f9093019485905282519097509195509293508392850191508083835b6020831061196c5780518252601f19909201916020918201910161194d565b51815160209384036101000a600019018019909216911617905260405191909301819003812094508a93508992019050808383808284378201915050925050506040516020818303038152906040526040518082805190602001908083835b602083106119ea5780518252601f1990920191602091820191016119cb565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020600019161480611b42575060405160200180807f7044697370656e736174696f6e5063740000000000000000000000000000000081525060100190506040516020818303038152906040526040518082805190602001908083835b60208310611a915780518252601f199092019160209182019101611a72565b51815160209384036101000a600019018019909216911617905260405191909301819003812094508a93508992019050808383808284378201915050925050506040516020818303038152906040526040518082805190602001908083835b60208310611b0f5780518252601f199092019160209182019101611af0565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916145b15611bda576064841115611bda576040805160e560020a62461bcd02815260206004820152605960248201526000805160206129a683398151915260448201527f616d65746572697a6174696f6e202d206076616c756560206d7573742062652060648201527f6c657373207468616e206f7220657175616c20746f2031303000000000000000608482015290519081900360a40190fd5b611be381610814565b15611c72576040805160e560020a62461bcd02815260206004820152604f60248201526000805160206129a683398151915260448201527f616d65746572697a6174696f6e202d204475706c69636174652070726f706f7360648201527f616c73206e6f7420616c6c6f7765640000000000000000000000000000000000608482015290519081900360a40190fd5b83611cac87878080601f0160208091040260200160405190810160405280939291908181526020018383808284375061092c945050505050565b1415611d3c576040805160e560020a62461bcd02815260206004820152604d60248201526000805160206129a683398151915260448201527f616d65746572697a6174696f6e202d20417267756d656e74732063616e6e6f7460648201527f206265206964656e746963616c00000000000000000000000000000000000000608482015290519081900360a40190fd5b604080516101208101909152600d60e082019081527f6170706c7953746167654c656e000000000000000000000000000000000000006101008301528190611d87906107979061092c565b81526020016000815260200183815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200133600160a060020a03168152602001611e276005546107a46107156040805190810160405280600e81526020017f72657665616c53746167654c656e00000000000000000000000000000000000081525061092c565b8152602090810186905260008381526002808352604091829020845181558484015160018201559184015190820155606083015180519192611e719260038501929091019061290a565b5060808201516004828101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0393841617905560a0840151600584015560c090930151600690920191909155600354604080517f23b872dd000000000000000000000000000000000000000000000000000000008152339481019490945230602485015260448401869052519116916323b872dd9160648083019260209291908290030181600087803b158015611f2657600080fd5b505af1158015611f3a573d6000803e3d6000fd5b505050506040513d6020811015611f5057600080fd5b50511515611fe2576040805160e560020a62461bcd02815260206004820152604860248201526000805160206129a683398151915260448201527f616d65746572697a6174696f6e202d20436f756c64206e6f74207472616e736660648201527f65722066756e6473000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b600081815260026020908152604091829020548251918201879052918101839052606081018490526080810182905260a0808252810187905233917ff6aafb6c8410e85ab57ef6d7df18fcb782806ee3b89634f3b52046846236f97a9189918991899187918991908060c081018888808284376040519201829003995090975050505050505050a295945050505050565b600061207d612888565b600083815260026020818152604092839020835160e0810185528154815260018083015482850152828501548287015260038301805487516000199382161561010002939093011695909504601f81018590048502820185019096528581529094919360608601939192909183018282801561213a5780601f1061210f5761010080835404028352916020019161213a565b820191906000526020600020905b81548152906001019060200180831161211d57829003601f168201915b50505091835250506004820154600160a060020a031660208201526005820154604082015260069091015460609091015280519091504211801561218157508060a0015142105b801561218f57506020810151155b9392505050565b80600080846040516020018082805190602001908083835b602083106121cd5780518252601f1990920191602091820191016121ae565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106122305780518252601f199092019160209182019101612211565b51815160209384036101000a60001901801990921691161790526040805192909401829003909120865285019590955292909201600020939093555050505050565b61227a612888565b6000828152600260208181526040808420815160e0810183528154815260018083015482860152828601548285015260038301805485516000199382161561010002939093011696909604601f810186900486028201860190945283815286959194929360608601939192918301828280156123375780601f1061230c57610100808354040283529160200191612337565b820191906000526020600020905b81548152906001019060200180831161231a57829003601f168201915b505050505081526020016004820160009054906101000a9004600160a060020a0316600160a060020a0316600160a060020a03168152602001600582015481526020016006820154815250509250600160008460200151815260200190815260200160002091506123ab836020015161082d565b60048054602080870151604080517f053e71a60000000000000000000000000000000000000000000000000000000081529485019190915251939450600160a060020a039091169263053e71a69260248082019392918290030181600087803b15801561241757600080fd5b505af115801561242b573d6000803e3d6000fd5b505050506040513d602081101561244157600080fd5b5051600383015560018201805474ff000000000000000000000000000000000000000019167401000000000000000000000000000000000000000017905560048054602085810151604080517f494031830000000000000000000000000000000000000000000000000000000081529485019190915251600160a060020a0390921692634940318392602480830193928290030181600087803b1580156124e757600080fd5b505af11580156124fb573d6000803e3d6000fd5b505050506040513d602081101561251157600080fd5b50511561268957428360a0015111156125365761253683606001518460c00151612196565b6020808401518354600385015460408051928352938201528251919287927f824960b36dcff66c0b43953e552b1b0d622dcb3f3cf71cf1cbe88039a781a8539281900390910190a360035460808401516040805160e060020a63a9059cbb028152600160a060020a039283166004820152602481018590529051919092169163a9059cbb9160448083019260209291908290030181600087803b1580156125dc57600080fd5b505af11580156125f0573d6000803e3d6000fd5b505050506040513d602081101561260657600080fd5b50511515612684576040805160e560020a62461bcd02815260206004820152603f60248201527f4572726f723a506172616d65746572697a65722e7265736f6c76654368616c6c60448201527f656e6765202d20436f756c64206e6f74207472616e736665722066756e647300606482015290519081900360840190fd5b611725565b6020808401518354600385015460408051928352938201528251919287927f3dddcf5cd7cb20cfa68f0643b9422cdbc130a6602933781bf9855e10fd205e429281900390910190a360035460208085015160009081526001808352604080832090910154815160e060020a63a9059cbb028152600160a060020a03918216600482015260248101879052915194169363a9059cbb9360448084019491938390030190829087803b15801561273c57600080fd5b505af1158015612750573d6000803e3d6000fd5b505050506040513d602081101561276657600080fd5b50511515611725576040805160e560020a62461bcd02815260206004820152603f60248201527f4572726f723a506172616d65746572697a65722e7265736f6c76654368616c6c60448201527f656e6765202d20436f756c64206e6f74207472616e736665722066756e647300606482015290519081900360840190fd5b8181018281101561092657fe5b600082151561280257506000610926565b5081810281838281151561281257fe5b041461092657fe5b60008282111561282657fe5b50900390565b6000818381151561283957fe5b049392505050565b50805460018160011615610100020316600290046000825580601f106128675750612885565b601f0160209004906000526020600020908101906128859190612988565b50565b60e060405190810160405280600081526020016000815260200160008152602001606081526020016000600160a060020a0316815260200160008152602001600081525090565b60a060405190810160405280600081526020016000600160a060020a0316815260200160001515815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061294b57805160ff1916838001178555612978565b82800160010185558215612978579182015b8281111561297857825182559160200191906001019061295d565b50612984929150612988565b5090565b6129a291905b80821115612984576000815560010161298e565b9056004572726f723a506172616d65746572697a65722e70726f706f73655265706172a165627a7a72305820b7fc3c7818977ac884a72928a5505c563d9f389357c07ceff7c11a025ba294180029`

// DeployParameterizer deploys a new Ethereum contract, binding an instance of Parameterizer to it.
func DeployParameterizer(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddr common.Address, votingAddr common.Address, minDeposit *big.Int, applyStageLen *big.Int, commitStageLen *big.Int, revealStageLen *big.Int, dispensationPct *big.Int, voteQuorum *big.Int, listReward *big.Int) (common.Address, *types.Transaction, *Parameterizer, error) {
	parsed, err := abi.JSON(strings.NewReader(ParameterizerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParameterizerBin), backend, tokenAddr, votingAddr, minDeposit, applyStageLen, commitStageLen, revealStageLen, dispensationPct, voteQuorum, listReward)
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
