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

// ParameterizerABI is the input ABI used to generate the binding from.
const ParameterizerABI = "[{\"name\":\"ReparamProposed\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"param\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ReparamFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"param\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ReparamSucceeded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"param\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"voting_addr\"},{\"type\":\"uint256\",\"name\":\"rate\",\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"denominator\"},{\"type\":\"uint256\",\"name\":\"numerator\"},{\"type\":\"uint256\",\"name\":\"list_re\",\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"comp_re\",\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"stk\",\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"vote_by_delta\",\"unit\":\"sec\"},{\"type\":\"uint256\",\"name\":\"quorum_pct\"},{\"type\":\"uint256\",\"name\":\"back_pay\"},{\"type\":\"uint256\",\"name\":\"maker_pay\"},{\"type\":\"uint256\",\"name\":\"cost\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getBackendPayment\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":483},{\"name\":\"getMakerPayment\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":513},{\"name\":\"getCostPerByte\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":543},{\"name\":\"getStake\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":573},{\"name\":\"getConversionRate\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":603},{\"name\":\"getHash\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"param\"},{\"type\":\"uint256\",\"name\":\"value\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":687},{\"name\":\"getInvestDenominator\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":663},{\"name\":\"getInvestNumerator\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":693},{\"name\":\"getListReward\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":723},{\"name\":\"getAccessReward\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":753},{\"name\":\"getQuorum\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":783},{\"name\":\"getReparam\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1489},{\"name\":\"getVoteBy\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"sec\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":843},{\"name\":\"reparameterize\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"param\"},{\"type\":\"uint256\",\"name\":\"value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":78290},{\"name\":\"resolveReparam\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":89520}]"

// ParameterizerBin is the compiled bytecode used for deploying new contracts.
const ParameterizerBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a0526101806108ed6101403934156100a857600080fd5b60206108ed60c03960c05160205181106100c157600080fd5b5061014051600c5561016051600155610180516002556101a0516003556101c0516004556101e0516005556102005160065561022051600755610240516008556102605160095561028051600a556102a051600b556108d556600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052634b8b7dce60005114156100b95734156100ac57600080fd5b60095460005260206000f3005b63ccc3412f60005114156100df5734156100d257600080fd5b600a5460005260206000f3005b63cd2d543860005114156101055734156100f857600080fd5b600b5460005260206000f3005b63fc0e3d90600051141561012b57341561011e57600080fd5b60065460005260206000f3005b63f36089ec600051141561015157341561014457600080fd5b60015460005260206000f3005b631b85504460005114156101a5576040600461014037341561017257600080fd5b61014051610160516101405101101561018a57600080fd5b61016051610140510160c052602060c02060005260206000f3005b634ba135d860005114156101cb5734156101be57600080fd5b60025460005260206000f3005b6328dee57060005114156101f15734156101e457600080fd5b60035460005260206000f3005b63dc2b2853600051141561021757341561020a57600080fd5b60045460005260206000f3005b631730eb2a600051141561023d57341561023057600080fd5b60055460005260206000f3005b63c26c12eb600051141561026357341561025657600080fd5b60085460005260206000f3005b63d90f8f0d60005114156102d6576020600461014037341561028457600080fd5b60406101605261018060006101405160e05260c052604060c02060c052602060c020548152600160006101405160e05260c052604060c02060c052602060c020015481602001525061016051610180f3005b632d0d2bc660005114156102fc5734156102ef57600080fd5b60075460005260206000f3005b63c18362186000511415610462576040600461014037341561031d57600080fd5b61014051610160516101405101101561033557600080fd5b61016051610140510160c052602060c02061018052600c543b61035757600080fd5b600c5430141561036657600080fd5b6020610220602463b89694c66101a052610180516101c0526101bc600c545afa61038f57600080fd5b60005061022051156103a057600080fd5b60006101805160e05260c052604060c02060c052602060c02061014051815561016051600182015550600c543b6103d657600080fd5b600c543014156103e557600080fd5b6000600060a463bb12c49e610240526101805161026052600361028052336102a0526006546102c0526007546102e05261025c6000600c545af161042857600080fd5b61016051610340526101405161018051337fb38eb1d14189ba2de4ca92c3ead8d4c2a61d566d266d46b487c4fa82ae7328426020610340a4005b63435c709a60005114156107b4576020600461014037341561048357600080fd5b600c543b61049057600080fd5b600c5430141561049f57600080fd5b6020610200604463af61f76061016052610140516101805260036101a05261017c600c545afa6104ce57600080fd5b600050610200516104de57600080fd5b600c543b6104eb57600080fd5b600c543014156104fa57600080fd5b60206102a0602463327322c861022052610140516102405261023c600c545afa61052357600080fd5b6000506102a05161053357600080fd5b60006101405160e05260c052604060c02060c052602060c020546102c052600160006101405160e05260c052604060c02060c052602060c02001546102e052600c543b61057f57600080fd5b600c5430141561058e57600080fd5b60206103a06044638f354b796103005261014051610320526008546103405261031c600c545afa6105be57600080fd5b6000506103a051156107105760026102c05114156105e2576102e0516001556106d8565b60036102c05114156105fa576102e0516002556106d7565b60046102c0511415610612576102e0516003556106d6565b60056102c051141561062a576102e0516004556106d5565b600a6102c0511415610642576102e0516005556106d4565b60016102c051141561065a576102e0516006556106d3565b60076102c0511415610672576102e0516007556106d2565b60066102c051141561068a576102e0516008556106d1565b60096102c05114156106a2576102e051600a556106d0565b60086102c05114156106ba576102e0516009556106cf565b600b6102c05114156106ce576102e051600b555b5b5b5b5b5b5b5b5b5b5b6102e0516103c0526102c051610140517fe6f29ce1d705c668b223a35e77096d95bbd994c121503fba4b29e65be92bc7e860206103c0a35b6102e0516103e0526102c051610140517f1f050cdbc74210070c4a499a73459732a158933d789331de2757eb4d3563025360206103e0a3600c543b61075457600080fd5b600c5430141561076357600080fd5b6000600060246389bb617c61040052610140516104205261041c6000600c545af161078d57600080fd5b60006101405160e05260c052604060c02060c052602060c020600081556000600182015550005b60006000fd5b61011b6108d50361011b60003961011b6108d5036000f3`

// DeployParameterizer deploys a new Ethereum contract, binding an instance of Parameterizer to it.
func DeployParameterizer(auth *bind.TransactOpts, backend bind.ContractBackend, voting_addr common.Address, rate *big.Int, denominator *big.Int, numerator *big.Int, list_re *big.Int, comp_re *big.Int, stk *big.Int, vote_by_delta *big.Int, quorum_pct *big.Int, back_pay *big.Int, maker_pay *big.Int, cost *big.Int) (common.Address, *types.Transaction, *Parameterizer, error) {
	parsed, err := abi.JSON(strings.NewReader(ParameterizerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParameterizerBin), backend, voting_addr, rate, denominator, numerator, list_re, comp_re, stk, vote_by_delta, quorum_pct, back_pay, maker_pay, cost)
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

// GetAccessReward is a free data retrieval call binding the contract method 0x1730eb2a.
//
// Solidity: function getAccessReward() constant returns(out uint256)
func (_Parameterizer *ParameterizerCaller) GetAccessReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getAccessReward")
	return *ret0, err
}

// GetAccessReward is a free data retrieval call binding the contract method 0x1730eb2a.
//
// Solidity: function getAccessReward() constant returns(out uint256)
func (_Parameterizer *ParameterizerSession) GetAccessReward() (*big.Int, error) {
	return _Parameterizer.Contract.GetAccessReward(&_Parameterizer.CallOpts)
}

// GetAccessReward is a free data retrieval call binding the contract method 0x1730eb2a.
//
// Solidity: function getAccessReward() constant returns(out uint256)
func (_Parameterizer *ParameterizerCallerSession) GetAccessReward() (*big.Int, error) {
	return _Parameterizer.Contract.GetAccessReward(&_Parameterizer.CallOpts)
}

// GetBackendPayment is a free data retrieval call binding the contract method 0x4b8b7dce.
//
// Solidity: function getBackendPayment() constant returns(out uint256)
func (_Parameterizer *ParameterizerCaller) GetBackendPayment(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getBackendPayment")
	return *ret0, err
}

// GetBackendPayment is a free data retrieval call binding the contract method 0x4b8b7dce.
//
// Solidity: function getBackendPayment() constant returns(out uint256)
func (_Parameterizer *ParameterizerSession) GetBackendPayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetBackendPayment(&_Parameterizer.CallOpts)
}

// GetBackendPayment is a free data retrieval call binding the contract method 0x4b8b7dce.
//
// Solidity: function getBackendPayment() constant returns(out uint256)
func (_Parameterizer *ParameterizerCallerSession) GetBackendPayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetBackendPayment(&_Parameterizer.CallOpts)
}

// GetConversionRate is a free data retrieval call binding the contract method 0xf36089ec.
//
// Solidity: function getConversionRate() constant returns(out uint256)
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
// Solidity: function getConversionRate() constant returns(out uint256)
func (_Parameterizer *ParameterizerSession) GetConversionRate() (*big.Int, error) {
	return _Parameterizer.Contract.GetConversionRate(&_Parameterizer.CallOpts)
}

// GetConversionRate is a free data retrieval call binding the contract method 0xf36089ec.
//
// Solidity: function getConversionRate() constant returns(out uint256)
func (_Parameterizer *ParameterizerCallerSession) GetConversionRate() (*big.Int, error) {
	return _Parameterizer.Contract.GetConversionRate(&_Parameterizer.CallOpts)
}

// GetCostPerByte is a free data retrieval call binding the contract method 0xcd2d5438.
//
// Solidity: function getCostPerByte() constant returns(out uint256)
func (_Parameterizer *ParameterizerCaller) GetCostPerByte(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getCostPerByte")
	return *ret0, err
}

// GetCostPerByte is a free data retrieval call binding the contract method 0xcd2d5438.
//
// Solidity: function getCostPerByte() constant returns(out uint256)
func (_Parameterizer *ParameterizerSession) GetCostPerByte() (*big.Int, error) {
	return _Parameterizer.Contract.GetCostPerByte(&_Parameterizer.CallOpts)
}

// GetCostPerByte is a free data retrieval call binding the contract method 0xcd2d5438.
//
// Solidity: function getCostPerByte() constant returns(out uint256)
func (_Parameterizer *ParameterizerCallerSession) GetCostPerByte() (*big.Int, error) {
	return _Parameterizer.Contract.GetCostPerByte(&_Parameterizer.CallOpts)
}

// GetHash is a free data retrieval call binding the contract method 0x1b855044.
//
// Solidity: function getHash(param uint256, value uint256) constant returns(out bytes32)
func (_Parameterizer *ParameterizerCaller) GetHash(opts *bind.CallOpts, param *big.Int, value *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getHash", param, value)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0x1b855044.
//
// Solidity: function getHash(param uint256, value uint256) constant returns(out bytes32)
func (_Parameterizer *ParameterizerSession) GetHash(param *big.Int, value *big.Int) ([32]byte, error) {
	return _Parameterizer.Contract.GetHash(&_Parameterizer.CallOpts, param, value)
}

// GetHash is a free data retrieval call binding the contract method 0x1b855044.
//
// Solidity: function getHash(param uint256, value uint256) constant returns(out bytes32)
func (_Parameterizer *ParameterizerCallerSession) GetHash(param *big.Int, value *big.Int) ([32]byte, error) {
	return _Parameterizer.Contract.GetHash(&_Parameterizer.CallOpts, param, value)
}

// GetInvestDenominator is a free data retrieval call binding the contract method 0x4ba135d8.
//
// Solidity: function getInvestDenominator() constant returns(out uint256)
func (_Parameterizer *ParameterizerCaller) GetInvestDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getInvestDenominator")
	return *ret0, err
}

// GetInvestDenominator is a free data retrieval call binding the contract method 0x4ba135d8.
//
// Solidity: function getInvestDenominator() constant returns(out uint256)
func (_Parameterizer *ParameterizerSession) GetInvestDenominator() (*big.Int, error) {
	return _Parameterizer.Contract.GetInvestDenominator(&_Parameterizer.CallOpts)
}

// GetInvestDenominator is a free data retrieval call binding the contract method 0x4ba135d8.
//
// Solidity: function getInvestDenominator() constant returns(out uint256)
func (_Parameterizer *ParameterizerCallerSession) GetInvestDenominator() (*big.Int, error) {
	return _Parameterizer.Contract.GetInvestDenominator(&_Parameterizer.CallOpts)
}

// GetInvestNumerator is a free data retrieval call binding the contract method 0x28dee570.
//
// Solidity: function getInvestNumerator() constant returns(out uint256)
func (_Parameterizer *ParameterizerCaller) GetInvestNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getInvestNumerator")
	return *ret0, err
}

// GetInvestNumerator is a free data retrieval call binding the contract method 0x28dee570.
//
// Solidity: function getInvestNumerator() constant returns(out uint256)
func (_Parameterizer *ParameterizerSession) GetInvestNumerator() (*big.Int, error) {
	return _Parameterizer.Contract.GetInvestNumerator(&_Parameterizer.CallOpts)
}

// GetInvestNumerator is a free data retrieval call binding the contract method 0x28dee570.
//
// Solidity: function getInvestNumerator() constant returns(out uint256)
func (_Parameterizer *ParameterizerCallerSession) GetInvestNumerator() (*big.Int, error) {
	return _Parameterizer.Contract.GetInvestNumerator(&_Parameterizer.CallOpts)
}

// GetListReward is a free data retrieval call binding the contract method 0xdc2b2853.
//
// Solidity: function getListReward() constant returns(out uint256)
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
// Solidity: function getListReward() constant returns(out uint256)
func (_Parameterizer *ParameterizerSession) GetListReward() (*big.Int, error) {
	return _Parameterizer.Contract.GetListReward(&_Parameterizer.CallOpts)
}

// GetListReward is a free data retrieval call binding the contract method 0xdc2b2853.
//
// Solidity: function getListReward() constant returns(out uint256)
func (_Parameterizer *ParameterizerCallerSession) GetListReward() (*big.Int, error) {
	return _Parameterizer.Contract.GetListReward(&_Parameterizer.CallOpts)
}

// GetMakerPayment is a free data retrieval call binding the contract method 0xccc3412f.
//
// Solidity: function getMakerPayment() constant returns(out uint256)
func (_Parameterizer *ParameterizerCaller) GetMakerPayment(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getMakerPayment")
	return *ret0, err
}

// GetMakerPayment is a free data retrieval call binding the contract method 0xccc3412f.
//
// Solidity: function getMakerPayment() constant returns(out uint256)
func (_Parameterizer *ParameterizerSession) GetMakerPayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetMakerPayment(&_Parameterizer.CallOpts)
}

// GetMakerPayment is a free data retrieval call binding the contract method 0xccc3412f.
//
// Solidity: function getMakerPayment() constant returns(out uint256)
func (_Parameterizer *ParameterizerCallerSession) GetMakerPayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetMakerPayment(&_Parameterizer.CallOpts)
}

// GetQuorum is a free data retrieval call binding the contract method 0xc26c12eb.
//
// Solidity: function getQuorum() constant returns(out uint256)
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
// Solidity: function getQuorum() constant returns(out uint256)
func (_Parameterizer *ParameterizerSession) GetQuorum() (*big.Int, error) {
	return _Parameterizer.Contract.GetQuorum(&_Parameterizer.CallOpts)
}

// GetQuorum is a free data retrieval call binding the contract method 0xc26c12eb.
//
// Solidity: function getQuorum() constant returns(out uint256)
func (_Parameterizer *ParameterizerCallerSession) GetQuorum() (*big.Int, error) {
	return _Parameterizer.Contract.GetQuorum(&_Parameterizer.CallOpts)
}

// GetReparam is a free data retrieval call binding the contract method 0xd90f8f0d.
//
// Solidity: function getReparam(hash bytes32) constant returns(out uint256, out uint256)
func (_Parameterizer *ParameterizerCaller) GetReparam(opts *bind.CallOpts, hash [32]byte) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Parameterizer.contract.Call(opts, out, "getReparam", hash)
	return *ret0, *ret1, err
}

// GetReparam is a free data retrieval call binding the contract method 0xd90f8f0d.
//
// Solidity: function getReparam(hash bytes32) constant returns(out uint256, out uint256)
func (_Parameterizer *ParameterizerSession) GetReparam(hash [32]byte) (*big.Int, *big.Int, error) {
	return _Parameterizer.Contract.GetReparam(&_Parameterizer.CallOpts, hash)
}

// GetReparam is a free data retrieval call binding the contract method 0xd90f8f0d.
//
// Solidity: function getReparam(hash bytes32) constant returns(out uint256, out uint256)
func (_Parameterizer *ParameterizerCallerSession) GetReparam(hash [32]byte) (*big.Int, *big.Int, error) {
	return _Parameterizer.Contract.GetReparam(&_Parameterizer.CallOpts, hash)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(out uint256)
func (_Parameterizer *ParameterizerCaller) GetStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getStake")
	return *ret0, err
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(out uint256)
func (_Parameterizer *ParameterizerSession) GetStake() (*big.Int, error) {
	return _Parameterizer.Contract.GetStake(&_Parameterizer.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(out uint256)
func (_Parameterizer *ParameterizerCallerSession) GetStake() (*big.Int, error) {
	return _Parameterizer.Contract.GetStake(&_Parameterizer.CallOpts)
}

// GetVoteBy is a free data retrieval call binding the contract method 0x2d0d2bc6.
//
// Solidity: function getVoteBy() constant returns(out uint256)
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
// Solidity: function getVoteBy() constant returns(out uint256)
func (_Parameterizer *ParameterizerSession) GetVoteBy() (*big.Int, error) {
	return _Parameterizer.Contract.GetVoteBy(&_Parameterizer.CallOpts)
}

// GetVoteBy is a free data retrieval call binding the contract method 0x2d0d2bc6.
//
// Solidity: function getVoteBy() constant returns(out uint256)
func (_Parameterizer *ParameterizerCallerSession) GetVoteBy() (*big.Int, error) {
	return _Parameterizer.Contract.GetVoteBy(&_Parameterizer.CallOpts)
}

// Reparameterize is a paid mutator transaction binding the contract method 0xc1836218.
//
// Solidity: function reparameterize(param uint256, value uint256) returns()
func (_Parameterizer *ParameterizerTransactor) Reparameterize(opts *bind.TransactOpts, param *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "reparameterize", param, value)
}

// Reparameterize is a paid mutator transaction binding the contract method 0xc1836218.
//
// Solidity: function reparameterize(param uint256, value uint256) returns()
func (_Parameterizer *ParameterizerSession) Reparameterize(param *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.Reparameterize(&_Parameterizer.TransactOpts, param, value)
}

// Reparameterize is a paid mutator transaction binding the contract method 0xc1836218.
//
// Solidity: function reparameterize(param uint256, value uint256) returns()
func (_Parameterizer *ParameterizerTransactorSession) Reparameterize(param *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.Reparameterize(&_Parameterizer.TransactOpts, param, value)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(hash bytes32) returns()
func (_Parameterizer *ParameterizerTransactor) ResolveReparam(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "resolveReparam", hash)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(hash bytes32) returns()
func (_Parameterizer *ParameterizerSession) ResolveReparam(hash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.Contract.ResolveReparam(&_Parameterizer.TransactOpts, hash)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(hash bytes32) returns()
func (_Parameterizer *ParameterizerTransactorSession) ResolveReparam(hash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.Contract.ResolveReparam(&_Parameterizer.TransactOpts, hash)
}

// ParameterizerReparamFailedIterator is returned from FilterReparamFailed and is used to iterate over the raw logs and unpacked data for ReparamFailed events raised by the Parameterizer contract.
type ParameterizerReparamFailedIterator struct {
	Event *ParameterizerReparamFailed // Event containing the contract specifics and raw log

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
func (it *ParameterizerReparamFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerReparamFailed)
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
		it.Event = new(ParameterizerReparamFailed)
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
func (it *ParameterizerReparamFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerReparamFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerReparamFailed represents a ReparamFailed event raised by the Parameterizer contract.
type ParameterizerReparamFailed struct {
	Hash  [32]byte
	Param *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReparamFailed is a free log retrieval operation binding the contract event 0x1f050cdbc74210070c4a499a73459732a158933d789331de2757eb4d35630253.
//
// Solidity: e ReparamFailed(hash indexed bytes32, param indexed uint256, value uint256)
func (_Parameterizer *ParameterizerFilterer) FilterReparamFailed(opts *bind.FilterOpts, hash [][32]byte, param []*big.Int) (*ParameterizerReparamFailedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ReparamFailed", hashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerReparamFailedIterator{contract: _Parameterizer.contract, event: "ReparamFailed", logs: logs, sub: sub}, nil
}

// WatchReparamFailed is a free log subscription operation binding the contract event 0x1f050cdbc74210070c4a499a73459732a158933d789331de2757eb4d35630253.
//
// Solidity: e ReparamFailed(hash indexed bytes32, param indexed uint256, value uint256)
func (_Parameterizer *ParameterizerFilterer) WatchReparamFailed(opts *bind.WatchOpts, sink chan<- *ParameterizerReparamFailed, hash [][32]byte, param []*big.Int) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ReparamFailed", hashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerReparamFailed)
				if err := _Parameterizer.contract.UnpackLog(event, "ReparamFailed", log); err != nil {
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

// ParameterizerReparamProposedIterator is returned from FilterReparamProposed and is used to iterate over the raw logs and unpacked data for ReparamProposed events raised by the Parameterizer contract.
type ParameterizerReparamProposedIterator struct {
	Event *ParameterizerReparamProposed // Event containing the contract specifics and raw log

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
func (it *ParameterizerReparamProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerReparamProposed)
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
		it.Event = new(ParameterizerReparamProposed)
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
func (it *ParameterizerReparamProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerReparamProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerReparamProposed represents a ReparamProposed event raised by the Parameterizer contract.
type ParameterizerReparamProposed struct {
	Owner common.Address
	Hash  [32]byte
	Param *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReparamProposed is a free log retrieval operation binding the contract event 0xb38eb1d14189ba2de4ca92c3ead8d4c2a61d566d266d46b487c4fa82ae732842.
//
// Solidity: e ReparamProposed(owner indexed address, hash indexed bytes32, param indexed uint256, value uint256)
func (_Parameterizer *ParameterizerFilterer) FilterReparamProposed(opts *bind.FilterOpts, owner []common.Address, hash [][32]byte, param []*big.Int) (*ParameterizerReparamProposedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ReparamProposed", ownerRule, hashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerReparamProposedIterator{contract: _Parameterizer.contract, event: "ReparamProposed", logs: logs, sub: sub}, nil
}

// WatchReparamProposed is a free log subscription operation binding the contract event 0xb38eb1d14189ba2de4ca92c3ead8d4c2a61d566d266d46b487c4fa82ae732842.
//
// Solidity: e ReparamProposed(owner indexed address, hash indexed bytes32, param indexed uint256, value uint256)
func (_Parameterizer *ParameterizerFilterer) WatchReparamProposed(opts *bind.WatchOpts, sink chan<- *ParameterizerReparamProposed, owner []common.Address, hash [][32]byte, param []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ReparamProposed", ownerRule, hashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerReparamProposed)
				if err := _Parameterizer.contract.UnpackLog(event, "ReparamProposed", log); err != nil {
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

// ParameterizerReparamSucceededIterator is returned from FilterReparamSucceeded and is used to iterate over the raw logs and unpacked data for ReparamSucceeded events raised by the Parameterizer contract.
type ParameterizerReparamSucceededIterator struct {
	Event *ParameterizerReparamSucceeded // Event containing the contract specifics and raw log

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
func (it *ParameterizerReparamSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParameterizerReparamSucceeded)
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
		it.Event = new(ParameterizerReparamSucceeded)
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
func (it *ParameterizerReparamSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParameterizerReparamSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParameterizerReparamSucceeded represents a ReparamSucceeded event raised by the Parameterizer contract.
type ParameterizerReparamSucceeded struct {
	Hash  [32]byte
	Param *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReparamSucceeded is a free log retrieval operation binding the contract event 0xe6f29ce1d705c668b223a35e77096d95bbd994c121503fba4b29e65be92bc7e8.
//
// Solidity: e ReparamSucceeded(hash indexed bytes32, param indexed uint256, value uint256)
func (_Parameterizer *ParameterizerFilterer) FilterReparamSucceeded(opts *bind.FilterOpts, hash [][32]byte, param []*big.Int) (*ParameterizerReparamSucceededIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.FilterLogs(opts, "ReparamSucceeded", hashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return &ParameterizerReparamSucceededIterator{contract: _Parameterizer.contract, event: "ReparamSucceeded", logs: logs, sub: sub}, nil
}

// WatchReparamSucceeded is a free log subscription operation binding the contract event 0xe6f29ce1d705c668b223a35e77096d95bbd994c121503fba4b29e65be92bc7e8.
//
// Solidity: e ReparamSucceeded(hash indexed bytes32, param indexed uint256, value uint256)
func (_Parameterizer *ParameterizerFilterer) WatchReparamSucceeded(opts *bind.WatchOpts, sink chan<- *ParameterizerReparamSucceeded, hash [][32]byte, param []*big.Int) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var paramRule []interface{}
	for _, paramItem := range param {
		paramRule = append(paramRule, paramItem)
	}

	logs, sub, err := _Parameterizer.contract.WatchLogs(opts, "ReparamSucceeded", hashRule, paramRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParameterizerReparamSucceeded)
				if err := _Parameterizer.contract.UnpackLog(event, "ReparamSucceeded", log); err != nil {
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
