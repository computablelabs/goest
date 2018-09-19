// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plcrvoting

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
const PLCRVotingABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getTotalNumberOfTokensForWinningOption\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"numTokens\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getInsertPointForNumTokens\",\"outputs\":[{\"name\":\"prevNode\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"name\":\"commitDuration\",\"type\":\"uint256\"},{\"name\":\"revealDuration\",\"type\":\"uint256\"}],\"name\":\"startPoll\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getLastNode\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"revealPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isPassed\",\"outputs\":[{\"name\":\"passed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getLockedTokens\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"secretHash\",\"type\":\"bytes32\"},{\"name\":\"numTokens\",\"type\":\"uint256\"},{\"name\":\"prevId\",\"type\":\"uint256\"}],\"name\":\"commitVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"didCommit\",\"outputs\":[{\"name\":\"committed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"prevId\",\"type\":\"uint256\"},{\"name\":\"nextId\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"validPosition\",\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"pollExists\",\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"rescueTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"requestVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"commitPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"didReveal\",\"outputs\":[{\"name\":\"revealed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"voteOption\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"revealVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getNumPassingTokens\",\"outputs\":[{\"name\":\"correctVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getNumTokens\",\"outputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getCommitHash\",\"outputs\":[{\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"terminationDate\",\"type\":\"uint256\"}],\"name\":\"isExpired\",\"outputs\":[{\"name\":\"expired\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"withdrawVotingRights\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"pollEnded\",\"outputs\":[{\"name\":\"ended\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VoteCommittedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesFor\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votesAgainst\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"choice\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VoteRevealedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commitExpiry\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealExpiry\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"PollCreatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VotingRightsGrantedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"numTokens\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VotingRightsWithdrawnEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"TokensRescuedEvent\",\"type\":\"event\"}]"

// PLCRVotingBin is the compiled bytecode used for deploying new contracts.
const PLCRVotingBin = `0x608060405234801561001057600080fd5b50604051602080612415833981016040525160058054600160a060020a03909216600160a060020a031990921691909117905560006002556123be806100576000396000f3006080604052600436106101275763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663053e71a6811461012c5780632c0520311461015657806332ed3d601461017d578063427fa1d21461019b578063441c77c0146101bc57806349403183146101e85780636b2d95d4146102005780636cbf9c5e146102215780637f97e83614610244578063819b02931461026857806388d21ff31461029257806397603560146102aa578063a25236fe146102c2578063a4439dc5146102da578063aa7ca464146102f2578063b11d8bb814610316578063b43bd06914610334578063d13820921461035b578063d901402b1461037f578063d9548e53146103a3578063e7b1d43c146103bb578063ee684830146103d3575b600080fd5b34801561013857600080fd5b506101446004356103eb565b60408051918252519081900360200190f35b34801561016257600080fd5b50610144600160a060020a03600435166024356044356104d8565b34801561018957600080fd5b50610144600435602435604435610549565b3480156101a757600080fd5b50610144600160a060020a0360043516610639565b3480156101c857600080fd5b506101d460043561064c565b604080519115158252519081900360200190f35b3480156101f457600080fd5b506101d4600435610706565b34801561020c57600080fd5b50610144600160a060020a0360043516610829565b34801561022d57600080fd5b5061024260043560243560443560643561083d565b005b34801561025057600080fd5b506101d4600160a060020a0360043516602435610c02565b34801561027457600080fd5b506101d4600435602435600160a060020a0360443516606435610cb8565b34801561029e57600080fd5b506101d4600435610cf9565b3480156102b657600080fd5b50610242600435610d0e565b3480156102ce57600080fd5b50610242600435610e65565b3480156102e657600080fd5b506101d4600435611120565b3480156102fe57600080fd5b506101d4600160a060020a03600435166024356111bf565b34801561032257600080fd5b50610242600435602435604435611275565b34801561034057600080fd5b50610144600160a060020a0360043516602435604435611674565b34801561036757600080fd5b50610144600160a060020a03600435166024356118fb565b34801561038b57600080fd5b50610144600160a060020a036004351660243561194a565b3480156103af57600080fd5b506101d4600435611992565b3480156103c757600080fd5b50610242600435611997565b3480156103df57600080fd5b506101d4600435611bea565b60006103f682611bea565b1515610498576040805160e560020a62461bcd02815260206004820152604a60248201527f4572726f723a566f74696e672e676574546f74616c4e756d6265724f66546f6b60448201527f656e73466f7257696e6e696e674f7074696f6e202d20506f6c6c206d7573742060648201527f6861766520656e64656400000000000000000000000000000000000000000000608482015290519081900360a40190fd5b6104a182610706565b156104bf5750600081815260036020819052604090912001546104d3565b506000818152600360205260409020600401545b919050565b60008060006104e686610639565b91506104f286836118fb565b90505b811561053c5761050586836118fb565b905084811161052b5783821415610523576105208683611c8c565b91505b819250610540565b6105358683611c8c565b91506104f5565b8192505b50509392505050565b60008060006105646001600254611cb790919063ffffffff16565b600255610577428663ffffffff611cb716565b9150610589828563ffffffff611cb716565b6040805160a08101825284815260208082018481528284018b815260006060808601828152608087018381526002805485526003808952948a90209851895595516001890155935187860155519186019190915590516004909401939093555483518b8152918201879052818401859052925193945033937ff7e7200f3afa7bbc8aca2cfa0ab33e741504a20fed2abc59869413cba2a33400929181900390910190a35050600254949350505050565b6000610646826000611c8c565b92915050565b600061065782610cf9565b15156106d3576040805160e560020a62461bcd02815260206004820152603160248201527f4572726f723a566f74696e672e72657665616c506572696f644163746976652060448201527f2d20506f6c6c206d757374206578697374000000000000000000000000000000606482015290519081900360840190fd5b6000828152600360205260409020600101546106ee90611992565b15801561064657506106ff82611120565b1592915050565b6000610710612322565b61071983611bea565b1515610795576040805160e560020a62461bcd02815260206004820152602c60248201527f4572726f723a566f74696e672e6973506173736564202d20506f6c6c206d757360448201527f74206861766520656e6465640000000000000000000000000000000000000000606482015290519081900360840190fd5b50600082815260036020818152604092839020835160a08101855281548152600182015492810192909252600281015493820193909352908201546060820181905260049092015460808201819052909161080b916107fa919063ffffffff611cb716565b60408301519063ffffffff611cc416565b606082015161082190606463ffffffff611cc416565b119392505050565b60006106468261083884610639565b6118fb565b60008061084986611120565b15156108c5576040805160e560020a62461bcd02815260206004820152603460248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d2054617267657460448201527f20706f6c6c206d75737420626520616374697665000000000000000000000000606482015290519081900360840190fd5b33600090815260046020526040902054841115610952576040805160e560020a62461bcd02815260206004820152603160248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d2053656e64657260448201527f2063616e6e6f74206f7665727370656e64000000000000000000000000000000606482015290519081900360840190fd5b8515156109cf576040805160e560020a62461bcd02815260206004820152602d60248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d20506f6c6c204960448201527f442063616e6e6f74206265203000000000000000000000000000000000000000606482015290519081900360840190fd5b8215806109e157506109e13384611ced565b1515610a5d576040805160e560020a62461bcd02815260206004820152602960248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d20506f6c6c206d60448201527f7573742065786973740000000000000000000000000000000000000000000000606482015290519081900360840190fd5b610a673384611d6b565b915085821415610a7e57610a7b3387611d6b565b91505b610a8a83833387610cb8565b1515610b06576040805160e560020a62461bcd02815260206004820152603460248201527f4572726f723a566f74696e672e636f6d6d6974566f7465202d20506f6c6c207060448201527f6f736974696f6e206973206e6f742076616c6964000000000000000000000000606482015290519081900360840190fd5b610b1233848885611d93565b610b1c3387611ff7565b9050610b5e816040805190810160405280600981526020017f6e756d546f6b656e730000000000000000000000000000000000000000000000815250866120a4565b60408051808201909152600a81527f636f6d6d697448617368000000000000000000000000000000000000000000006020820152610b9e908290876120a4565b6000868152600360209081526040808320338085526005909101835292819020805460ff191660011790558051878152905189927f673f4a020a46f562a6e9fee9c8d4a179a11924190b777bec160133dd00c9f1de928290030190a3505050505050565b6000610c0d82610cf9565b1515610c89576040805160e560020a62461bcd02815260206004820152602f60248201527f4572726f723a506172616d65746572697a65722e646964436f6d6d6974202d2060448201527f506f6c6c206d7573742065786973740000000000000000000000000000000000606482015290519081900360840190fd5b506000908152600360209081526040808320600160a060020a0394909416835260059093019052205460ff1690565b6000806000610cc785886118fb565b8410159150610cd685876118fb565b84111580610ce2575085155b9050818015610cee5750805b979650505050505050565b60008115801590610646575050600254101590565b600081815260036020526040902060010154610d2990611992565b1515610da5576040805160e560020a62461bcd02815260206004820152603060248201527f4572726f723a566f74696e672e726573637565546f6b656e73202d20506f6c6c60448201527f206d757374206861766520656e64656400000000000000000000000000000000606482015290519081900360840190fd5b610daf3382611ced565b1515610e2b576040805160e560020a62461bcd02815260206004820152602a60248201527f4572726f723a566f74696e672e726573637565546f6b656e73202d20506f6c6c60448201527f206e6f7420666f756e6400000000000000000000000000000000000000000000606482015290519081900360840190fd5b610e353382612188565b604051339082907f1e897e4fe81e0db0459f0219545995fde305b17af5ca5cba7d6b6e8738df5dd090600090a350565b600554604080517f70a0823100000000000000000000000000000000000000000000000000000000815233600482015290518392600160a060020a0316916370a082319160248083019260209291908290030181600087803b158015610eca57600080fd5b505af1158015610ede573d6000803e3d6000fd5b505050506040513d6020811015610ef457600080fd5b50511015610f98576040805160e560020a62461bcd02815260206004820152604860248201527f4572726f723a566f74696e672e72657175657374566f74696e6752696768747360448201527f202d2053656e64657220646f6573206e6f74206861766520737566666963696560648201527f6e742066756e6473000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b33600090815260046020526040902054610fb8908263ffffffff611cb716565b3360008181526004602081815260408084209590955560055485517f23b872dd00000000000000000000000000000000000000000000000000000000815292830194909452306024830152604482018690529351600160a060020a03909316936323b872dd9360648084019492939192918390030190829087803b15801561103f57600080fd5b505af1158015611053573d6000803e3d6000fd5b505050506040513d602081101561106957600080fd5b505115156110e7576040805160e560020a62461bcd02815260206004820152603d60248201527f4572726f723a566f74696e672e72657175657374566f74696e6752696768747360448201527f202d20546f6b656e2e7472616e7366657246726f6d206661696c757265000000606482015290519081900360840190fd5b60408051828152905133917f0fbea84792de88539ad3d7f184543a65916440f4c6f9e2ef136c97f13109c45d919081900360200190a250565b600061112b82610cf9565b15156111a7576040805160e560020a62461bcd02815260206004820152603160248201527f4572726f723a566f74696e672e636f6d6d6974506572696f644163746976652060448201527f2d20506f6c6c206d757374206578697374000000000000000000000000000000606482015290519081900360840190fd5b6000828152600360205260409020546106ff90611992565b60006111ca82610cf9565b1515611246576040805160e560020a62461bcd02815260206004820152602860248201527f4572726f723a566f74696e672e64696452657665616c202d20506f6c6c206d7560448201527f7374206578697374000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b506000908152600360209081526040808320600160a060020a0394909416835260069093019052205460ff1690565b60006112808461064c565b15156112fc576040805160e560020a62461bcd02815260206004820152603660248201527f4572726f723a566f74696e672e72657665616c566f7465202d2052657665616c60448201527f20706572696f64206d7573742062652061637469766500000000000000000000606482015290519081900360840190fd5b600084815260036020908152604080832033845260050190915290205460ff161515611398576040805160e560020a62461bcd02815260206004820152603f60248201527f4572726f723a566f74696e672e72657665616c566f7465202d2053656e64657260448201527f206d757374206861766520636f6d6d697474656420746865697220766f746500606482015290519081900360840190fd5b600084815260036020908152604080832033845260060190915290205460ff1615611459576040805160e560020a62461bcd02815260206004820152604860248201527f4572726f723a566f74696e672e72657665616c566f7465202d2053656e64657260448201527f2063616e6e6f742072657665616c20746865697220766f7465206d756c74697060648201527f6c652074696d6573000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b611463338561194a565b604080516020808201879052818301869052825180830384018152606090920192839052815191929182918401908083835b602083106114b45780518252601f199092019160209182019101611495565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916141515611562576040805160e560020a62461bcd02815260206004820152602760248201527f4572726f723a566f74696e672e72657665616c566f7465202d2048617368206d60448201527f69736d6174636800000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b61156c33856118fb565b905082600114156115b2576000848152600360208190526040909120015461159a908263ffffffff611cb716565b600085815260036020819052604090912001556115e7565b6000848152600360205260409020600401546115d4908263ffffffff611cb716565b6000858152600360205260409020600401555b6115f13385612188565b600084815260036020818152604080842033808652600682018452828620805460ff191660011790559489905283835292830154600490930154815186815292830193909352818101929092529051859187917f448a2036d0f6dc799c1d3fdf26c4fe80832817bdbe81133808ffa6777c5826e09181900360600190a450505050565b60008060008061168386611bea565b15156116ff576040805160e560020a62461bcd02815260206004820152603760248201527f4572726f723a566f74696e672e6765744e756d50617373696e67546f6b656e7360448201527f202d20506f6c6c206d757374206861766520656e646564000000000000000000606482015290519081900360840190fd5b6000868152600360209081526040808320600160a060020a038b16845260060190915290205460ff1615156117ca576040805160e560020a62461bcd02815260206004820152604660248201527f4572726f723a566f74696e672e6765744e756d50617373696e67546f6b656e7360448201527f202d20566f746572206d75737420686176652072657665616c6564207468656960648201527f7220766f74650000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b6117d386610706565b6117de5760006117e1565b60015b60ff169250828560405160200180838152602001828152602001925050506040516020818303038152906040526040518082805190602001908083835b6020831061183d5780518252601f19909201916020918201910161181e565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209150611876878761194a565b90508181146118f5576040805160e560020a62461bcd02815260206004820152603060248201527f4572726f723a566f74696e672e6765744e756d50617373696e67546f6b656e7360448201527f202d2048617368206d69736d6174636800000000000000000000000000000000606482015290519081900360840190fd5b610cee87875b600061194361190a8484611ff7565b60408051808201909152600981527f6e756d546f6b656e73000000000000000000000000000000000000000000000060208201526121fb565b9392505050565b60006119436119598484611ff7565b60408051808201909152600a81527f636f6d6d6974486173680000000000000000000000000000000000000000000060208201526121fb565b421190565b60006119c16119a533610829565b336000908152600460205260409020549063ffffffff6122e516565b905081811015611a67576040805160e560020a62461bcd02815260206004820152604160248201527f4572726f723a566f74696e672e7769746864726177566f74696e67526967687460448201527f73202d20496e73756666696369656e7420617661696c61626c6520746f6b656e60648201527f7300000000000000000000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b33600090815260046020526040902054611a87908363ffffffff6122e516565b3360008181526004602081815260408084209590955560055485517fa9059cbb00000000000000000000000000000000000000000000000000000000815292830194909452602482018790529351600160a060020a039093169363a9059cbb9360448084019492939192918390030190829087803b158015611b0857600080fd5b505af1158015611b1c573d6000803e3d6000fd5b505050506040513d6020811015611b3257600080fd5b50511515611bb0576040805160e560020a62461bcd02815260206004820152603a60248201527f4572726f723a566f74696e672e7769746864726177566f74696e67526967687460448201527f73202d20546f6b656e2e7472616e73666572206661696c757265000000000000606482015290519081900360840190fd5b60408051838152905133917f029a58e9abaa250a4ec75d14b364d764e890adec2d902a297fa24a6c5dfd097b919081900360200190a25050565b6000611bf582610cf9565b1515611c71576040805160e560020a62461bcd02815260206004820152602860248201527f4572726f723a566f74696e672e706f6c6c456e646564202d20506f6c6c206d7560448201527f7374206578697374000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008281526003602052604090206001015461064690611992565b600160a060020a03919091166000908152600160208181526040808420948452939052919020015490565b8181018281101561064657fe5b6000821515611cd557506000610646565b50818102818382811515611ce557fe5b041461064657fe5b6000806000611cfb856122f7565b80611d04575083155b15611d125760009250611d63565b83611d1c8661230a565b148015611d30575083611d2e86612316565b145b91506000611d3e8686611d6b565b148015611d5457506000611d528686611c8c565b145b90508180611d60575080155b92505b505092915050565b600160a060020a03919091166000908152600160209081526040808320938352929052205490565b811515611e10576040805160e560020a62461bcd02815260206004820152603a60248201527f4572726f723a504c4352566f74696e672e646c6c496e73657274202d2049442060448201527f6f662061206e6577206e6f6465206d6179206e6f742062652030000000000000606482015290519081900360840190fd5b611e1a8483612188565b821580611e2c5750611e2c8484611ced565b1515611e84576040805160e560020a62461bcd02815260206004820152603360248201526000805160206123738339815191526044820152600080516020612353833981519152606482015290519081900360840190fd5b801580611e965750611e968482611ced565b1515611eee576040805160e560020a62461bcd02815260206004820152603360248201526000805160206123738339815191526044820152600080516020612353833981519152606482015290519081900360840190fd5b80611ef98585611d6b565b14611f50576040805160e560020a62461bcd02815260206004820152603360248201526000805160206123738339815191526044820152600080516020612353833981519152606482015290519081900360840190fd5b82611f5b8583611c8c565b14611fb2576040805160e560020a62461bcd02815260206004820152603360248201526000805160206123738339815191526044820152600080516020612353833981519152606482015290519081900360840190fd5b600160a060020a039390931660009081526001602081815260408084208585529091528083208083018690558690559382528382208390559381529190912090910155565b600082826040516020018083600160a060020a0316600160a060020a03166c01000000000000000000000000028152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106120715780518252601f199092019160209182019101612052565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091209695505050505050565b604080516020808201868152855160009488948894910191908401908083835b602083106120e35780518252601f1990920191602091820191016120c4565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b602083106121475780518252601f199092019160209182019101612128565b51815160209384036101000a600019018019909216911617905260408051929094018290039091206000908152908190529190912094909455505050505050565b6000806121958484611ced565b15156121a0576121f5565b6121aa8484611d6b565b91506121b68484611c8c565b600160a060020a038516600090815260016020818152604080842087855290915280832082018490558383528083208690558683528220828155015590505b50505050565b600080838360405160200180836000191660001916815260200182805190602001908083835b602083106122405780518252601f199092019160209182019101612221565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b602083106122a45780518252601f199092019160209182019101612285565b51815160209384036101000a600019018019909216911617905260408051929094018290039091206000908152908190529190912054979650505050505050565b6000828211156122f157fe5b50900390565b6000806123038361230a565b1492915050565b60006106468282611d6b565b60006106468282611c8c565b60a0604051908101604052806000815260200160008152602001600081526020016000815260200160008152509056006e6f7420696e7365727420696e746f20444c4c000000000000000000000000004572726f723a504c4352566f74696e672e644c4c496e73657274202d20436f6ea165627a7a7230582091fef9ae27d9b6df88ffb9725f59e00541b6350b82eeb001b0ba8043cf2b101a0029`

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
