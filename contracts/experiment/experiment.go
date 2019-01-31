// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package experiment

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ExperimentABI is the input ABI used to generate the binding from.
const ExperimentABI = "[{\"name\":\"__init__\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"setUserMarketTokenBalance\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"new_balance\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35295},{\"name\":\"setConversionRate\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"new_rate\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35325},{\"name\":\"setReserve\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"new_reserve\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35355},{\"name\":\"getReserve\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":573},{\"name\":\"setSlopeNumerator\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"new_n\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35415},{\"name\":\"setSlopeDenominator\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"new_d\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35445},{\"name\":\"setTotalSupply\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"new_supply\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35475},{\"name\":\"getInvestmentPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":8622},{\"name\":\"getDivestmentProceeds\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2638}]"

// ExperimentBin is the compiled bytecode used for deploying new contracts.
const ExperimentBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052341561009e57600080fd5b6000600055633b9aca00600155600060025560646003556065600455671bc16d674ec8000060065561046e56600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052639ea70ff460005114156100bd57602060046101403734156100b457600080fd5b61014051600055005b63d2e8049460005114156100e757602060046101403734156100de57600080fd5b61014051600155005b634256dbe36000511415610111576020600461014037341561010857600080fd5b61014051600255005b6359bf5d39600051141561013757341561012a57600080fd5b60025460005260206000f3005b631de7223b6000511415610161576020600461014037341561015857600080fd5b61014051600355005b63d58ee471600051141561018b576020600461014037341561018257600080fd5b61014051600455005b63f7ea7a3d60005114156101b557602060046101403734156101ac57600080fd5b61014051600655005b63a3d6602b600051141561033a5734156101ce57600080fd5b60015460045415156101e1576000610207565b633b9aca00600454633b9aca006004540204146101fd57600080fd5b633b9aca00600454025b61021057600080fd5b6004541515610220576000610246565b633b9aca00600454633b9aca0060045402041461023c57600080fd5b633b9aca00600454025b6003541515610256576000610276565b60025460035460025460035402041461026e57600080fd5b600254600354025b0460015401101561028657600080fd5b60045415156102965760006102bc565b633b9aca00600454633b9aca006004540204146102b257600080fd5b633b9aca00600454025b6102c557600080fd5b60045415156102d55760006102fb565b633b9aca00600454633b9aca006004540204146102f157600080fd5b633b9aca00600454025b600354151561030b57600061032b565b60025460035460025460035402041461032357600080fd5b600254600354025b046001540160005260206000f3005b634fef9077600051141561039d57341561035357600080fd5b60065461035f57600080fd5b6006546000541515610372576000610392565b60025460005460025460005402041461038a57600080fd5b600254600054025b0460005260206000f3005b60006000fd5b6100cb61046e036100cb6000396100cb61046e036000f3`

// DeployExperiment deploys a new Ethereum contract, binding an instance of Experiment to it.
func DeployExperiment(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Experiment, error) {
	parsed, err := abi.JSON(strings.NewReader(ExperimentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExperimentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Experiment{ExperimentCaller: ExperimentCaller{contract: contract}, ExperimentTransactor: ExperimentTransactor{contract: contract}, ExperimentFilterer: ExperimentFilterer{contract: contract}}, nil
}

// Experiment is an auto generated Go binding around an Ethereum contract.
type Experiment struct {
	ExperimentCaller     // Read-only binding to the contract
	ExperimentTransactor // Write-only binding to the contract
	ExperimentFilterer   // Log filterer for contract events
}

// ExperimentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExperimentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExperimentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExperimentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExperimentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExperimentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExperimentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExperimentSession struct {
	Contract     *Experiment       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExperimentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExperimentCallerSession struct {
	Contract *ExperimentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ExperimentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExperimentTransactorSession struct {
	Contract     *ExperimentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ExperimentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExperimentRaw struct {
	Contract *Experiment // Generic contract binding to access the raw methods on
}

// ExperimentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExperimentCallerRaw struct {
	Contract *ExperimentCaller // Generic read-only contract binding to access the raw methods on
}

// ExperimentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExperimentTransactorRaw struct {
	Contract *ExperimentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExperiment creates a new instance of Experiment, bound to a specific deployed contract.
func NewExperiment(address common.Address, backend bind.ContractBackend) (*Experiment, error) {
	contract, err := bindExperiment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Experiment{ExperimentCaller: ExperimentCaller{contract: contract}, ExperimentTransactor: ExperimentTransactor{contract: contract}, ExperimentFilterer: ExperimentFilterer{contract: contract}}, nil
}

// NewExperimentCaller creates a new read-only instance of Experiment, bound to a specific deployed contract.
func NewExperimentCaller(address common.Address, caller bind.ContractCaller) (*ExperimentCaller, error) {
	contract, err := bindExperiment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExperimentCaller{contract: contract}, nil
}

// NewExperimentTransactor creates a new write-only instance of Experiment, bound to a specific deployed contract.
func NewExperimentTransactor(address common.Address, transactor bind.ContractTransactor) (*ExperimentTransactor, error) {
	contract, err := bindExperiment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExperimentTransactor{contract: contract}, nil
}

// NewExperimentFilterer creates a new log filterer instance of Experiment, bound to a specific deployed contract.
func NewExperimentFilterer(address common.Address, filterer bind.ContractFilterer) (*ExperimentFilterer, error) {
	contract, err := bindExperiment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExperimentFilterer{contract: contract}, nil
}

// bindExperiment binds a generic wrapper to an already deployed contract.
func bindExperiment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExperimentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Experiment *ExperimentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Experiment.Contract.ExperimentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Experiment *ExperimentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Experiment.Contract.ExperimentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Experiment *ExperimentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Experiment.Contract.ExperimentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Experiment *ExperimentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Experiment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Experiment *ExperimentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Experiment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Experiment *ExperimentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Experiment.Contract.contract.Transact(opts, method, params...)
}

// GetDivestmentProceeds is a free data retrieval call binding the contract method 0x4fef9077.
//
// Solidity: function getDivestmentProceeds() constant returns(out uint256)
func (_Experiment *ExperimentCaller) GetDivestmentProceeds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Experiment.contract.Call(opts, out, "getDivestmentProceeds")
	return *ret0, err
}

// GetDivestmentProceeds is a free data retrieval call binding the contract method 0x4fef9077.
//
// Solidity: function getDivestmentProceeds() constant returns(out uint256)
func (_Experiment *ExperimentSession) GetDivestmentProceeds() (*big.Int, error) {
	return _Experiment.Contract.GetDivestmentProceeds(&_Experiment.CallOpts)
}

// GetDivestmentProceeds is a free data retrieval call binding the contract method 0x4fef9077.
//
// Solidity: function getDivestmentProceeds() constant returns(out uint256)
func (_Experiment *ExperimentCallerSession) GetDivestmentProceeds() (*big.Int, error) {
	return _Experiment.Contract.GetDivestmentProceeds(&_Experiment.CallOpts)
}

// GetInvestmentPrice is a free data retrieval call binding the contract method 0xa3d6602b.
//
// Solidity: function getInvestmentPrice() constant returns(out uint256)
func (_Experiment *ExperimentCaller) GetInvestmentPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Experiment.contract.Call(opts, out, "getInvestmentPrice")
	return *ret0, err
}

// GetInvestmentPrice is a free data retrieval call binding the contract method 0xa3d6602b.
//
// Solidity: function getInvestmentPrice() constant returns(out uint256)
func (_Experiment *ExperimentSession) GetInvestmentPrice() (*big.Int, error) {
	return _Experiment.Contract.GetInvestmentPrice(&_Experiment.CallOpts)
}

// GetInvestmentPrice is a free data retrieval call binding the contract method 0xa3d6602b.
//
// Solidity: function getInvestmentPrice() constant returns(out uint256)
func (_Experiment *ExperimentCallerSession) GetInvestmentPrice() (*big.Int, error) {
	return _Experiment.Contract.GetInvestmentPrice(&_Experiment.CallOpts)
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() constant returns(out uint256)
func (_Experiment *ExperimentCaller) GetReserve(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Experiment.contract.Call(opts, out, "getReserve")
	return *ret0, err
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() constant returns(out uint256)
func (_Experiment *ExperimentSession) GetReserve() (*big.Int, error) {
	return _Experiment.Contract.GetReserve(&_Experiment.CallOpts)
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() constant returns(out uint256)
func (_Experiment *ExperimentCallerSession) GetReserve() (*big.Int, error) {
	return _Experiment.Contract.GetReserve(&_Experiment.CallOpts)
}

// SetConversionRate is a paid mutator transaction binding the contract method 0xd2e80494.
//
// Solidity: function setConversionRate(new_rate uint256) returns()
func (_Experiment *ExperimentTransactor) SetConversionRate(opts *bind.TransactOpts, new_rate *big.Int) (*types.Transaction, error) {
	return _Experiment.contract.Transact(opts, "setConversionRate", new_rate)
}

// SetConversionRate is a paid mutator transaction binding the contract method 0xd2e80494.
//
// Solidity: function setConversionRate(new_rate uint256) returns()
func (_Experiment *ExperimentSession) SetConversionRate(new_rate *big.Int) (*types.Transaction, error) {
	return _Experiment.Contract.SetConversionRate(&_Experiment.TransactOpts, new_rate)
}

// SetConversionRate is a paid mutator transaction binding the contract method 0xd2e80494.
//
// Solidity: function setConversionRate(new_rate uint256) returns()
func (_Experiment *ExperimentTransactorSession) SetConversionRate(new_rate *big.Int) (*types.Transaction, error) {
	return _Experiment.Contract.SetConversionRate(&_Experiment.TransactOpts, new_rate)
}

// SetReserve is a paid mutator transaction binding the contract method 0x4256dbe3.
//
// Solidity: function setReserve(new_reserve uint256) returns()
func (_Experiment *ExperimentTransactor) SetReserve(opts *bind.TransactOpts, new_reserve *big.Int) (*types.Transaction, error) {
	return _Experiment.contract.Transact(opts, "setReserve", new_reserve)
}

// SetReserve is a paid mutator transaction binding the contract method 0x4256dbe3.
//
// Solidity: function setReserve(new_reserve uint256) returns()
func (_Experiment *ExperimentSession) SetReserve(new_reserve *big.Int) (*types.Transaction, error) {
	return _Experiment.Contract.SetReserve(&_Experiment.TransactOpts, new_reserve)
}

// SetReserve is a paid mutator transaction binding the contract method 0x4256dbe3.
//
// Solidity: function setReserve(new_reserve uint256) returns()
func (_Experiment *ExperimentTransactorSession) SetReserve(new_reserve *big.Int) (*types.Transaction, error) {
	return _Experiment.Contract.SetReserve(&_Experiment.TransactOpts, new_reserve)
}

// SetSlopeDenominator is a paid mutator transaction binding the contract method 0xd58ee471.
//
// Solidity: function setSlopeDenominator(new_d uint256) returns()
func (_Experiment *ExperimentTransactor) SetSlopeDenominator(opts *bind.TransactOpts, new_d *big.Int) (*types.Transaction, error) {
	return _Experiment.contract.Transact(opts, "setSlopeDenominator", new_d)
}

// SetSlopeDenominator is a paid mutator transaction binding the contract method 0xd58ee471.
//
// Solidity: function setSlopeDenominator(new_d uint256) returns()
func (_Experiment *ExperimentSession) SetSlopeDenominator(new_d *big.Int) (*types.Transaction, error) {
	return _Experiment.Contract.SetSlopeDenominator(&_Experiment.TransactOpts, new_d)
}

// SetSlopeDenominator is a paid mutator transaction binding the contract method 0xd58ee471.
//
// Solidity: function setSlopeDenominator(new_d uint256) returns()
func (_Experiment *ExperimentTransactorSession) SetSlopeDenominator(new_d *big.Int) (*types.Transaction, error) {
	return _Experiment.Contract.SetSlopeDenominator(&_Experiment.TransactOpts, new_d)
}

// SetSlopeNumerator is a paid mutator transaction binding the contract method 0x1de7223b.
//
// Solidity: function setSlopeNumerator(new_n uint256) returns()
func (_Experiment *ExperimentTransactor) SetSlopeNumerator(opts *bind.TransactOpts, new_n *big.Int) (*types.Transaction, error) {
	return _Experiment.contract.Transact(opts, "setSlopeNumerator", new_n)
}

// SetSlopeNumerator is a paid mutator transaction binding the contract method 0x1de7223b.
//
// Solidity: function setSlopeNumerator(new_n uint256) returns()
func (_Experiment *ExperimentSession) SetSlopeNumerator(new_n *big.Int) (*types.Transaction, error) {
	return _Experiment.Contract.SetSlopeNumerator(&_Experiment.TransactOpts, new_n)
}

// SetSlopeNumerator is a paid mutator transaction binding the contract method 0x1de7223b.
//
// Solidity: function setSlopeNumerator(new_n uint256) returns()
func (_Experiment *ExperimentTransactorSession) SetSlopeNumerator(new_n *big.Int) (*types.Transaction, error) {
	return _Experiment.Contract.SetSlopeNumerator(&_Experiment.TransactOpts, new_n)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(new_supply uint256) returns()
func (_Experiment *ExperimentTransactor) SetTotalSupply(opts *bind.TransactOpts, new_supply *big.Int) (*types.Transaction, error) {
	return _Experiment.contract.Transact(opts, "setTotalSupply", new_supply)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(new_supply uint256) returns()
func (_Experiment *ExperimentSession) SetTotalSupply(new_supply *big.Int) (*types.Transaction, error) {
	return _Experiment.Contract.SetTotalSupply(&_Experiment.TransactOpts, new_supply)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(new_supply uint256) returns()
func (_Experiment *ExperimentTransactorSession) SetTotalSupply(new_supply *big.Int) (*types.Transaction, error) {
	return _Experiment.Contract.SetTotalSupply(&_Experiment.TransactOpts, new_supply)
}

// SetUserMarketTokenBalance is a paid mutator transaction binding the contract method 0x9ea70ff4.
//
// Solidity: function setUserMarketTokenBalance(new_balance uint256) returns()
func (_Experiment *ExperimentTransactor) SetUserMarketTokenBalance(opts *bind.TransactOpts, new_balance *big.Int) (*types.Transaction, error) {
	return _Experiment.contract.Transact(opts, "setUserMarketTokenBalance", new_balance)
}

// SetUserMarketTokenBalance is a paid mutator transaction binding the contract method 0x9ea70ff4.
//
// Solidity: function setUserMarketTokenBalance(new_balance uint256) returns()
func (_Experiment *ExperimentSession) SetUserMarketTokenBalance(new_balance *big.Int) (*types.Transaction, error) {
	return _Experiment.Contract.SetUserMarketTokenBalance(&_Experiment.TransactOpts, new_balance)
}

// SetUserMarketTokenBalance is a paid mutator transaction binding the contract method 0x9ea70ff4.
//
// Solidity: function setUserMarketTokenBalance(new_balance uint256) returns()
func (_Experiment *ExperimentTransactorSession) SetUserMarketTokenBalance(new_balance *big.Int) (*types.Transaction, error) {
	return _Experiment.Contract.SetUserMarketTokenBalance(&_Experiment.TransactOpts, new_balance)
}
