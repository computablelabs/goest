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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ParameterizerABI is the input ABI used to generate the binding from.
const ParameterizerABI = "[{\"name\":\"ReparamProposed\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"param\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ReparamFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"param\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ReparamSucceeded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"param\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"mkt_addr\"},{\"type\":\"address\",\"name\":\"v_addr\"},{\"type\":\"uint256\",\"name\":\"pr_fl\",\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"spd\"},{\"type\":\"uint256\",\"name\":\"list_re\",\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"stk\",\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"vote_by_d\",\"unit\":\"sec\"},{\"type\":\"uint256\",\"name\":\"pl\"},{\"type\":\"uint256\",\"name\":\"back_p\"},{\"type\":\"uint256\",\"name\":\"maker_p\"},{\"type\":\"uint256\",\"name\":\"cost\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getBackendPayment\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":483},{\"name\":\"getMakerPayment\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":513},{\"name\":\"getReservePayment\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2766},{\"name\":\"getCostPerByte\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":573},{\"name\":\"getStake\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":603},{\"name\":\"getPriceFloor\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":633},{\"name\":\"getHash\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"param\"},{\"type\":\"uint256\",\"name\":\"value\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":722},{\"name\":\"getSpread\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":693},{\"name\":\"getListReward\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":723},{\"name\":\"getPlurality\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":753},{\"name\":\"getReparam\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1459},{\"name\":\"getVoteBy\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"sec\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":813},{\"name\":\"reparameterize\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"param\"},{\"type\":\"uint256\",\"name\":\"value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":78259},{\"name\":\"resolveReparam\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":88970}]"

// ParameterizerBin is the compiled bytecode used for deploying new contracts.
const ParameterizerBin = `0x740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052610160610a156101403934156100a257600080fd5b6020610a1560c03960c05160205181106100bb57600080fd5b5060206020610a150160c03960c05160205181106100d857600080fd5b5061014051600b5561016051600a55610180516001556101a0516002556101c0516003556101e05160045561020051600555610220516006556102405160075561026051600855610280516009556109fd56600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052634b8b7dce60005114156100b95734156100ac57600080fd5b60075460005260206000f3005b63ccc3412f60005114156100df5734156100d257600080fd5b60085460005260206000f3005b63ffe44da8600051141561014a5734156100f857600080fd5b60075460085460075401101561010d57600080fd5b600854600754016064101561012157600080fd5b60075460085460075401101561013657600080fd5b6008546007540160640360005260206000f3005b63cd2d5438600051141561017057341561016357600080fd5b60095460005260206000f3005b63fc0e3d90600051141561019657341561018957600080fd5b60045460005260206000f3005b6303bca7d360005114156101bc5734156101af57600080fd5b60015460005260206000f3005b631b855044600051141561022257604060046101403734156101dd57600080fd5b600061014051602082610180010152602081019050610160516020826101800101526020810190508061018052610180905080516020820120905060005260206000f3005b637b3cee60600051141561024857341561023b57600080fd5b60025460005260206000f3005b63dc2b2853600051141561026e57341561026157600080fd5b60035460005260206000f3005b633d1e37d5600051141561029457341561028757600080fd5b60065460005260206000f3005b63d90f8f0d600051141561030757602060046101403734156102b557600080fd5b60406101605261018060006101405160e05260c052604060c02060c052602060c020548152600160006101405160e05260c052604060c02060c052602060c020015481602001525061016051610180f3005b632d0d2bc6600051141561032d57341561032057600080fd5b60055460005260206000f3005b63c183621860005114156104a3576040600461014037341561034e57600080fd5b6000610140516020826101a0010152602081019050610160516020826101a0010152602081019050806101a0526101a0905080516020820120905061018052600a543b61039a57600080fd5b600a5430186103a857600080fd5b60206102a0602463b89694c661022052610180516102405261023c600a545afa6103d157600080fd5b6000506102a051156103e257600080fd5b60006101805160e05260c052604060c02060c052602060c02061014051815561016051600182015550600a543b61041857600080fd5b600a54301861042657600080fd5b6000600060a463bb12c49e6102c052610180516102e052600361030052336103205260045461034052600554610360526102dc6000600a545af161046957600080fd5b610160516103c0526101405161018051337fb38eb1d14189ba2de4ca92c3ead8d4c2a61d566d266d46b487c4fa82ae73284260206103c0a4005b63435c709a60005114156108cc57602060046101403734156104c457600080fd5b600a543b6104d157600080fd5b600a5430186104df57600080fd5b6020610200604463af61f76061016052610140516101805260036101a05261017c600a545afa61050e57600080fd5b6000506102005161051e57600080fd5b600a543b61052b57600080fd5b600a54301861053957600080fd5b60206102a0602463327322c861022052610140516102405261023c600a545afa61056257600080fd5b6000506102a05161057257600080fd5b60006101405160e05260c052604060c02060c052602060c020546102c052600160006101405160e05260c052604060c02060c052602060c02001546102e052600a543b6105be57600080fd5b600a5430186105cc57600080fd5b60206103c06044638f354b796103205261014051610340526006546103605261033c600a545afa6105fc57600080fd5b6000506103c051156108285760026102c0511415610620576102e0516001556107ec565b60046102c05114156106495760646102e051101561063d57600080fd5b6102e0516002556107eb565b60056102c0511415610661576102e0516003556107ea565b60016102c05114156106e157600b543b61067a57600080fd5b600b54301861068857600080fd5b602061044060046318160ddd6103e0526103fc600b545afa6106a957600080fd5b600050610440516104605260036106bf57600080fd5b600361046051046102e05111156106d557600080fd5b6102e0516004556107e9565b60076102c051141561071f57620151806102e051101561070057600080fd5b621275006102e051111561071357600080fd5b6102e0516005556107e8565b60066102c05114156107485760646102e051111561073c57600080fd5b6102e0516006556107e7565b60096102c051141561078c5760646102e0516007546102e05101101561076d57600080fd5b6007546102e05101111561078057600080fd5b6102e0516008556107e6565b60086102c05114156107d05760646102e0516008546102e0510110156107b157600080fd5b6008546102e0510111156107c457600080fd5b6102e0516007556107e5565b600b6102c05114156107e4576102e0516009555b5b5b5b5b5b5b5b5b6102e051610480526102c051610140517fe6f29ce1d705c668b223a35e77096d95bbd994c121503fba4b29e65be92bc7e86020610480a3610860565b6102e051610300526102c051610140517f1f050cdbc74210070c4a499a73459732a158933d789331de2757eb4d356302536020610300a35b600a543b61086d57600080fd5b600a54301861087b57600080fd5b6000600060246389bb617c6104a052610140516104c0526104bc6000600a545af16108a557600080fd5b60006101405160e05260c052604060c02060c052602060c020600081556000600182015550005b60006000fd5b61012b6109fd0361012b60003961012b6109fd036000f3`

// DeployParameterizer deploys a new Ethereum contract, binding an instance of Parameterizer to it.
func DeployParameterizer(auth *bind.TransactOpts, backend bind.ContractBackend, mkt_addr common.Address, v_addr common.Address, pr_fl *big.Int, spd *big.Int, list_re *big.Int, stk *big.Int, vote_by_d *big.Int, pl *big.Int, back_p *big.Int, maker_p *big.Int, cost *big.Int) (common.Address, *types.Transaction, *Parameterizer, error) {
	parsed, err := abi.JSON(strings.NewReader(ParameterizerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParameterizerBin), backend, mkt_addr, v_addr, pr_fl, spd, list_re, stk, vote_by_d, pl, back_p, maker_p, cost)
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

// GetBackendPayment is a free data retrieval call binding the contract method 0x4b8b7dce.
//
// Solidity: function getBackendPayment() constant returns(uint256 out)
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
// Solidity: function getBackendPayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetBackendPayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetBackendPayment(&_Parameterizer.CallOpts)
}

// GetBackendPayment is a free data retrieval call binding the contract method 0x4b8b7dce.
//
// Solidity: function getBackendPayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetBackendPayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetBackendPayment(&_Parameterizer.CallOpts)
}

// GetCostPerByte is a free data retrieval call binding the contract method 0xcd2d5438.
//
// Solidity: function getCostPerByte() constant returns(uint256 out)
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
// Solidity: function getCostPerByte() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetCostPerByte() (*big.Int, error) {
	return _Parameterizer.Contract.GetCostPerByte(&_Parameterizer.CallOpts)
}

// GetCostPerByte is a free data retrieval call binding the contract method 0xcd2d5438.
//
// Solidity: function getCostPerByte() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetCostPerByte() (*big.Int, error) {
	return _Parameterizer.Contract.GetCostPerByte(&_Parameterizer.CallOpts)
}

// GetHash is a free data retrieval call binding the contract method 0x1b855044.
//
// Solidity: function getHash(uint256 param, uint256 value) constant returns(bytes32 out)
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
// Solidity: function getHash(uint256 param, uint256 value) constant returns(bytes32 out)
func (_Parameterizer *ParameterizerSession) GetHash(param *big.Int, value *big.Int) ([32]byte, error) {
	return _Parameterizer.Contract.GetHash(&_Parameterizer.CallOpts, param, value)
}

// GetHash is a free data retrieval call binding the contract method 0x1b855044.
//
// Solidity: function getHash(uint256 param, uint256 value) constant returns(bytes32 out)
func (_Parameterizer *ParameterizerCallerSession) GetHash(param *big.Int, value *big.Int) ([32]byte, error) {
	return _Parameterizer.Contract.GetHash(&_Parameterizer.CallOpts, param, value)
}

// GetListReward is a free data retrieval call binding the contract method 0xdc2b2853.
//
// Solidity: function getListReward() constant returns(uint256 out)
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
// Solidity: function getListReward() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetListReward() (*big.Int, error) {
	return _Parameterizer.Contract.GetListReward(&_Parameterizer.CallOpts)
}

// GetListReward is a free data retrieval call binding the contract method 0xdc2b2853.
//
// Solidity: function getListReward() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetListReward() (*big.Int, error) {
	return _Parameterizer.Contract.GetListReward(&_Parameterizer.CallOpts)
}

// GetMakerPayment is a free data retrieval call binding the contract method 0xccc3412f.
//
// Solidity: function getMakerPayment() constant returns(uint256 out)
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
// Solidity: function getMakerPayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetMakerPayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetMakerPayment(&_Parameterizer.CallOpts)
}

// GetMakerPayment is a free data retrieval call binding the contract method 0xccc3412f.
//
// Solidity: function getMakerPayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetMakerPayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetMakerPayment(&_Parameterizer.CallOpts)
}

// GetPlurality is a free data retrieval call binding the contract method 0x3d1e37d5.
//
// Solidity: function getPlurality() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetPlurality(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getPlurality")
	return *ret0, err
}

// GetPlurality is a free data retrieval call binding the contract method 0x3d1e37d5.
//
// Solidity: function getPlurality() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetPlurality() (*big.Int, error) {
	return _Parameterizer.Contract.GetPlurality(&_Parameterizer.CallOpts)
}

// GetPlurality is a free data retrieval call binding the contract method 0x3d1e37d5.
//
// Solidity: function getPlurality() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetPlurality() (*big.Int, error) {
	return _Parameterizer.Contract.GetPlurality(&_Parameterizer.CallOpts)
}

// GetPriceFloor is a free data retrieval call binding the contract method 0x03bca7d3.
//
// Solidity: function getPriceFloor() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetPriceFloor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getPriceFloor")
	return *ret0, err
}

// GetPriceFloor is a free data retrieval call binding the contract method 0x03bca7d3.
//
// Solidity: function getPriceFloor() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetPriceFloor() (*big.Int, error) {
	return _Parameterizer.Contract.GetPriceFloor(&_Parameterizer.CallOpts)
}

// GetPriceFloor is a free data retrieval call binding the contract method 0x03bca7d3.
//
// Solidity: function getPriceFloor() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetPriceFloor() (*big.Int, error) {
	return _Parameterizer.Contract.GetPriceFloor(&_Parameterizer.CallOpts)
}

// GetReparam is a free data retrieval call binding the contract method 0xd90f8f0d.
//
// Solidity: function getReparam(bytes32 hash) constant returns(uint256 out, uint256 out)
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
// Solidity: function getReparam(bytes32 hash) constant returns(uint256 out, uint256 out)
func (_Parameterizer *ParameterizerSession) GetReparam(hash [32]byte) (*big.Int, *big.Int, error) {
	return _Parameterizer.Contract.GetReparam(&_Parameterizer.CallOpts, hash)
}

// GetReparam is a free data retrieval call binding the contract method 0xd90f8f0d.
//
// Solidity: function getReparam(bytes32 hash) constant returns(uint256 out, uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetReparam(hash [32]byte) (*big.Int, *big.Int, error) {
	return _Parameterizer.Contract.GetReparam(&_Parameterizer.CallOpts, hash)
}

// GetReservePayment is a free data retrieval call binding the contract method 0xffe44da8.
//
// Solidity: function getReservePayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetReservePayment(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getReservePayment")
	return *ret0, err
}

// GetReservePayment is a free data retrieval call binding the contract method 0xffe44da8.
//
// Solidity: function getReservePayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetReservePayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetReservePayment(&_Parameterizer.CallOpts)
}

// GetReservePayment is a free data retrieval call binding the contract method 0xffe44da8.
//
// Solidity: function getReservePayment() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetReservePayment() (*big.Int, error) {
	return _Parameterizer.Contract.GetReservePayment(&_Parameterizer.CallOpts)
}

// GetSpread is a free data retrieval call binding the contract method 0x7b3cee60.
//
// Solidity: function getSpread() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCaller) GetSpread(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameterizer.contract.Call(opts, out, "getSpread")
	return *ret0, err
}

// GetSpread is a free data retrieval call binding the contract method 0x7b3cee60.
//
// Solidity: function getSpread() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetSpread() (*big.Int, error) {
	return _Parameterizer.Contract.GetSpread(&_Parameterizer.CallOpts)
}

// GetSpread is a free data retrieval call binding the contract method 0x7b3cee60.
//
// Solidity: function getSpread() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetSpread() (*big.Int, error) {
	return _Parameterizer.Contract.GetSpread(&_Parameterizer.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256 out)
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
// Solidity: function getStake() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetStake() (*big.Int, error) {
	return _Parameterizer.Contract.GetStake(&_Parameterizer.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xfc0e3d90.
//
// Solidity: function getStake() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetStake() (*big.Int, error) {
	return _Parameterizer.Contract.GetStake(&_Parameterizer.CallOpts)
}

// GetVoteBy is a free data retrieval call binding the contract method 0x2d0d2bc6.
//
// Solidity: function getVoteBy() constant returns(uint256 out)
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
// Solidity: function getVoteBy() constant returns(uint256 out)
func (_Parameterizer *ParameterizerSession) GetVoteBy() (*big.Int, error) {
	return _Parameterizer.Contract.GetVoteBy(&_Parameterizer.CallOpts)
}

// GetVoteBy is a free data retrieval call binding the contract method 0x2d0d2bc6.
//
// Solidity: function getVoteBy() constant returns(uint256 out)
func (_Parameterizer *ParameterizerCallerSession) GetVoteBy() (*big.Int, error) {
	return _Parameterizer.Contract.GetVoteBy(&_Parameterizer.CallOpts)
}

// Reparameterize is a paid mutator transaction binding the contract method 0xc1836218.
//
// Solidity: function reparameterize(uint256 param, uint256 value) returns()
func (_Parameterizer *ParameterizerTransactor) Reparameterize(opts *bind.TransactOpts, param *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "reparameterize", param, value)
}

// Reparameterize is a paid mutator transaction binding the contract method 0xc1836218.
//
// Solidity: function reparameterize(uint256 param, uint256 value) returns()
func (_Parameterizer *ParameterizerSession) Reparameterize(param *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.Reparameterize(&_Parameterizer.TransactOpts, param, value)
}

// Reparameterize is a paid mutator transaction binding the contract method 0xc1836218.
//
// Solidity: function reparameterize(uint256 param, uint256 value) returns()
func (_Parameterizer *ParameterizerTransactorSession) Reparameterize(param *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Parameterizer.Contract.Reparameterize(&_Parameterizer.TransactOpts, param, value)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(bytes32 hash) returns()
func (_Parameterizer *ParameterizerTransactor) ResolveReparam(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.contract.Transact(opts, "resolveReparam", hash)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(bytes32 hash) returns()
func (_Parameterizer *ParameterizerSession) ResolveReparam(hash [32]byte) (*types.Transaction, error) {
	return _Parameterizer.Contract.ResolveReparam(&_Parameterizer.TransactOpts, hash)
}

// ResolveReparam is a paid mutator transaction binding the contract method 0x435c709a.
//
// Solidity: function resolveReparam(bytes32 hash) returns()
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
// Solidity: event ReparamFailed(bytes32 indexed hash, uint256 indexed param, uint256 value)
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
// Solidity: event ReparamFailed(bytes32 indexed hash, uint256 indexed param, uint256 value)
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
// Solidity: event ReparamProposed(address indexed owner, bytes32 indexed hash, uint256 indexed param, uint256 value)
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
// Solidity: event ReparamProposed(address indexed owner, bytes32 indexed hash, uint256 indexed param, uint256 value)
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
// Solidity: event ReparamSucceeded(bytes32 indexed hash, uint256 indexed param, uint256 value)
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
// Solidity: event ReparamSucceeded(bytes32 indexed hash, uint256 indexed param, uint256 value)
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
