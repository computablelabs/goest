// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reserve

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

// ReserveABI is the input ABI used to generate the binding from.
const ReserveABI = "[{\"name\":\"Withdrawn\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"transferred\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Supported\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"offered\",\"indexed\":false,\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"minted\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"ether_token_addr\"},{\"type\":\"address\",\"name\":\"market_token_addr\"},{\"type\":\"address\",\"name\":\"p11r_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getSupportPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":10492},{\"name\":\"support\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"offer\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":22186},{\"name\":\"getWithdrawalProceeds\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6038},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":13931}]"

// ReserveBin is the compiled bytecode used for deploying new contracts.
const ReserveBin = `0x740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a0526060610b5d6101403934156100a157600080fd5b6020610b5d60c03960c05160205181106100ba57600080fd5b5060206020610b5d0160c03960c05160205181106100d757600080fd5b5060206040610b5d0160c03960c05160205181106100f457600080fd5b50610140516000556101605160015561018051600255610b4556600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263a056a5b960005114156105905734156100ac57600080fd5b6002543b6100b957600080fd5b60025430186100c757600080fd5b60206101c060046303bca7d36101605261017c6002545afa6100e857600080fd5b6000506101c051610140526002543b61010057600080fd5b600254301861010e57600080fd5b60206102606004637b3cee606102005261021c6002545afa61012f57600080fd5b600050610260516101e0526000543b61014757600080fd5b600054301861015557600080fd5b602061032060246370a082316102a052306102c0526102bc6000545afa61017b57600080fd5b60005061032051610280526001543b61019357600080fd5b60015430186101a157600080fd5b60206103c060046318160ddd6103605261037c6001545afa6101c257600080fd5b6000506103c05161034052670de0b6b3a7640000610340511015610290576101405160646101ef57600080fd5b60646101e0511515610202576000610228565b610280516101e051610280516101e05102041461021e57600080fd5b610280516101e051025b046101405101101561023957600080fd5b606461024457600080fd5b60646101e051151561025757600061027d565b610280516101e051610280516101e05102041461027357600080fd5b610280516101e051025b04610140510160005260206000f361058e565b61014051606415156102a35760006102c3565b6103405160646103405160640204146102bb57600080fd5b610340516064025b6102cc57600080fd5b606415156102db5760006102fb565b6103405160646103405160640204146102f357600080fd5b610340516064025b6101e051151561030c576000610332565b610280516101e051610280516101e05102041461032857600080fd5b610280516101e051025b151561033f576000610401565b633b9aca006101e051151561035557600061037b565b610280516101e051610280516101e05102041461037157600080fd5b610280516101e051025b633b9aca006101e05115156103915760006103b7565b610280516101e051610280516101e0510204146103ad57600080fd5b610280516101e051025b0204146103c357600080fd5b633b9aca006101e05115156103d95760006103ff565b610280516101e051610280516101e0510204146103f557600080fd5b610280516101e051025b025b046101405101101561041257600080fd5b60641515610421576000610441565b61034051606461034051606402041461043957600080fd5b610340516064025b61044a57600080fd5b60641515610459576000610479565b61034051606461034051606402041461047157600080fd5b610340516064025b6101e051151561048a5760006104b0565b610280516101e051610280516101e0510204146104a657600080fd5b610280516101e051025b15156104bd57600061057f565b633b9aca006101e05115156104d35760006104f9565b610280516101e051610280516101e0510204146104ef57600080fd5b610280516101e051025b633b9aca006101e051151561050f576000610535565b610280516101e051610280516101e05102041461052b57600080fd5b610280516101e051025b02041461054157600080fd5b633b9aca006101e051151561055757600061057d565b610280516101e051610280516101e05102041461057357600080fd5b610280516101e051025b025b04610140510160005260206000f35b005b6356c9493f600051141561079b57602060046101403734156105b157600080fd5b60206101e0600463a056a5b96101805261019c6000305af16105d257600080fd5b6101e05161016052610160516101405110156105ed57600080fd5b6000543b6105fa57600080fd5b600054301861060857600080fd5b60206102c060646323b872dd6102005233610220523061024052610140516102605261021c60006000545af161063d57600080fd5b6000506102c0506101605161065157600080fd5b61016051610140510415156106675760006106c6565b633b9aca006101605161067957600080fd5b610160516101405104633b9aca006101605161069457600080fd5b6101605161014051040204146106a957600080fd5b633b9aca00610160516106bb57600080fd5b610160516101405104025b6102e0526001543b6106d757600080fd5b60015430186106e557600080fd5b60006000602463a0712d68610300526102e0516103205261031c60006001545af161070f57600080fd5b6001543b61071c57600080fd5b600154301861072a57600080fd5b6020610420604463a9059cbb61038052336103a0526102e0516103c05261039c60006001545af161075a57600080fd5b6000506104205061014051610440526102e05161046052337fb19981ad2efd084bc7ddcdc94541b38fd493a1c176ffe2cd1e512e6ee0c34fe96040610440a2005b634633020d600051141561091357602060046101403734156107bc57600080fd5b60043560205181106107cd57600080fd5b50600061014051186107de57600080fd5b6001543b6107eb57600080fd5b60015430186107f957600080fd5b602061020060246370a0823161018052610140516101a05261019c6001545afa61082257600080fd5b60005061020051610160526000543b61083a57600080fd5b600054301861084857600080fd5b60206102c060246370a0823161024052306102605261025c6000545afa61086e57600080fd5b6000506102c051610220526001543b61088657600080fd5b600154301861089457600080fd5b602061036060046318160ddd6103005261031c6001545afa6108b557600080fd5b600050610360516102e0526102e0516108cd57600080fd5b6102e0516101605115156108e2576000610908565b610220516101605161022051610160510204146108fe57600080fd5b6102205161016051025b0460005260206000f3005b633ccfd60b6000511415610a3057341561092c57600080fd5b60206101e06024634633020d61016052336101805261017c6000305af161095257600080fd5b6101e051610140526000610140511161096a57600080fd5b6001543b61097757600080fd5b600154301861098557600080fd5b600060006024637e9d2ac161020052336102205261021c60006001545af16109ac57600080fd5b6000543b6109b957600080fd5b60005430186109c757600080fd5b6020610320604463a9059cbb61028052336102a052610140516102c05261029c60006000545af16109f757600080fd5b600050610320506101405161034052337f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d56020610340a2005b60006000fd5b61010f610b450361010f60003961010f610b45036000f3`

// DeployReserve deploys a new Ethereum contract, binding an instance of Reserve to it.
func DeployReserve(auth *bind.TransactOpts, backend bind.ContractBackend, ether_token_addr common.Address, market_token_addr common.Address, p11r_addr common.Address) (common.Address, *types.Transaction, *Reserve, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReserveBin), backend, ether_token_addr, market_token_addr, p11r_addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reserve{ReserveCaller: ReserveCaller{contract: contract}, ReserveTransactor: ReserveTransactor{contract: contract}, ReserveFilterer: ReserveFilterer{contract: contract}}, nil
}

// Reserve is an auto generated Go binding around an Ethereum contract.
type Reserve struct {
	ReserveCaller     // Read-only binding to the contract
	ReserveTransactor // Write-only binding to the contract
	ReserveFilterer   // Log filterer for contract events
}

// ReserveCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReserveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReserveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReserveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReserveSession struct {
	Contract     *Reserve          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReserveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReserveCallerSession struct {
	Contract *ReserveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ReserveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReserveTransactorSession struct {
	Contract     *ReserveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReserveRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReserveRaw struct {
	Contract *Reserve // Generic contract binding to access the raw methods on
}

// ReserveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReserveCallerRaw struct {
	Contract *ReserveCaller // Generic read-only contract binding to access the raw methods on
}

// ReserveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReserveTransactorRaw struct {
	Contract *ReserveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReserve creates a new instance of Reserve, bound to a specific deployed contract.
func NewReserve(address common.Address, backend bind.ContractBackend) (*Reserve, error) {
	contract, err := bindReserve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reserve{ReserveCaller: ReserveCaller{contract: contract}, ReserveTransactor: ReserveTransactor{contract: contract}, ReserveFilterer: ReserveFilterer{contract: contract}}, nil
}

// NewReserveCaller creates a new read-only instance of Reserve, bound to a specific deployed contract.
func NewReserveCaller(address common.Address, caller bind.ContractCaller) (*ReserveCaller, error) {
	contract, err := bindReserve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveCaller{contract: contract}, nil
}

// NewReserveTransactor creates a new write-only instance of Reserve, bound to a specific deployed contract.
func NewReserveTransactor(address common.Address, transactor bind.ContractTransactor) (*ReserveTransactor, error) {
	contract, err := bindReserve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveTransactor{contract: contract}, nil
}

// NewReserveFilterer creates a new log filterer instance of Reserve, bound to a specific deployed contract.
func NewReserveFilterer(address common.Address, filterer bind.ContractFilterer) (*ReserveFilterer, error) {
	contract, err := bindReserve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReserveFilterer{contract: contract}, nil
}

// bindReserve binds a generic wrapper to an already deployed contract.
func bindReserve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reserve *ReserveRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reserve.Contract.ReserveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reserve *ReserveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.Contract.ReserveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reserve *ReserveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reserve.Contract.ReserveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reserve *ReserveCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reserve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reserve *ReserveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reserve *ReserveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reserve.Contract.contract.Transact(opts, method, params...)
}

// GetSupportPrice is a free data retrieval call binding the contract method 0xa056a5b9.
//
// Solidity: function getSupportPrice() constant returns(out uint256)
func (_Reserve *ReserveCaller) GetSupportPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "getSupportPrice")
	return *ret0, err
}

// GetSupportPrice is a free data retrieval call binding the contract method 0xa056a5b9.
//
// Solidity: function getSupportPrice() constant returns(out uint256)
func (_Reserve *ReserveSession) GetSupportPrice() (*big.Int, error) {
	return _Reserve.Contract.GetSupportPrice(&_Reserve.CallOpts)
}

// GetSupportPrice is a free data retrieval call binding the contract method 0xa056a5b9.
//
// Solidity: function getSupportPrice() constant returns(out uint256)
func (_Reserve *ReserveCallerSession) GetSupportPrice() (*big.Int, error) {
	return _Reserve.Contract.GetSupportPrice(&_Reserve.CallOpts)
}

// GetWithdrawalProceeds is a free data retrieval call binding the contract method 0x4633020d.
//
// Solidity: function getWithdrawalProceeds(addr address) constant returns(out uint256)
func (_Reserve *ReserveCaller) GetWithdrawalProceeds(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "getWithdrawalProceeds", addr)
	return *ret0, err
}

// GetWithdrawalProceeds is a free data retrieval call binding the contract method 0x4633020d.
//
// Solidity: function getWithdrawalProceeds(addr address) constant returns(out uint256)
func (_Reserve *ReserveSession) GetWithdrawalProceeds(addr common.Address) (*big.Int, error) {
	return _Reserve.Contract.GetWithdrawalProceeds(&_Reserve.CallOpts, addr)
}

// GetWithdrawalProceeds is a free data retrieval call binding the contract method 0x4633020d.
//
// Solidity: function getWithdrawalProceeds(addr address) constant returns(out uint256)
func (_Reserve *ReserveCallerSession) GetWithdrawalProceeds(addr common.Address) (*big.Int, error) {
	return _Reserve.Contract.GetWithdrawalProceeds(&_Reserve.CallOpts, addr)
}

// Support is a paid mutator transaction binding the contract method 0x56c9493f.
//
// Solidity: function support(offer uint256) returns()
func (_Reserve *ReserveTransactor) Support(opts *bind.TransactOpts, offer *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "support", offer)
}

// Support is a paid mutator transaction binding the contract method 0x56c9493f.
//
// Solidity: function support(offer uint256) returns()
func (_Reserve *ReserveSession) Support(offer *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.Support(&_Reserve.TransactOpts, offer)
}

// Support is a paid mutator transaction binding the contract method 0x56c9493f.
//
// Solidity: function support(offer uint256) returns()
func (_Reserve *ReserveTransactorSession) Support(offer *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.Support(&_Reserve.TransactOpts, offer)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Reserve *ReserveTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Reserve *ReserveSession) Withdraw() (*types.Transaction, error) {
	return _Reserve.Contract.Withdraw(&_Reserve.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Reserve *ReserveTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Reserve.Contract.Withdraw(&_Reserve.TransactOpts)
}

// ReserveSupportedIterator is returned from FilterSupported and is used to iterate over the raw logs and unpacked data for Supported events raised by the Reserve contract.
type ReserveSupportedIterator struct {
	Event *ReserveSupported // Event containing the contract specifics and raw log

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
func (it *ReserveSupportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveSupported)
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
		it.Event = new(ReserveSupported)
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
func (it *ReserveSupportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveSupportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveSupported represents a Supported event raised by the Reserve contract.
type ReserveSupported struct {
	Owner   common.Address
	Offered *big.Int
	Minted  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSupported is a free log retrieval operation binding the contract event 0xb19981ad2efd084bc7ddcdc94541b38fd493a1c176ffe2cd1e512e6ee0c34fe9.
//
// Solidity: e Supported(owner indexed address, offered uint256, minted uint256)
func (_Reserve *ReserveFilterer) FilterSupported(opts *bind.FilterOpts, owner []common.Address) (*ReserveSupportedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "Supported", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ReserveSupportedIterator{contract: _Reserve.contract, event: "Supported", logs: logs, sub: sub}, nil
}

// WatchSupported is a free log subscription operation binding the contract event 0xb19981ad2efd084bc7ddcdc94541b38fd493a1c176ffe2cd1e512e6ee0c34fe9.
//
// Solidity: e Supported(owner indexed address, offered uint256, minted uint256)
func (_Reserve *ReserveFilterer) WatchSupported(opts *bind.WatchOpts, sink chan<- *ReserveSupported, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "Supported", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveSupported)
				if err := _Reserve.contract.UnpackLog(event, "Supported", log); err != nil {
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

// ReserveWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Reserve contract.
type ReserveWithdrawnIterator struct {
	Event *ReserveWithdrawn // Event containing the contract specifics and raw log

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
func (it *ReserveWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveWithdrawn)
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
		it.Event = new(ReserveWithdrawn)
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
func (it *ReserveWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveWithdrawn represents a Withdrawn event raised by the Reserve contract.
type ReserveWithdrawn struct {
	Owner       common.Address
	Transferred *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(owner indexed address, transferred uint256)
func (_Reserve *ReserveFilterer) FilterWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*ReserveWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "Withdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ReserveWithdrawnIterator{contract: _Reserve.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(owner indexed address, transferred uint256)
func (_Reserve *ReserveFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ReserveWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "Withdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveWithdrawn)
				if err := _Reserve.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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