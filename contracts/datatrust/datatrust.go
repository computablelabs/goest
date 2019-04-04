// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package datatrust

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

// DatatrustABI is the input ABI used to generate the binding from.
const DatatrustABI = "[{\"name\":\"Registered\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"registrant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RegistrationSucceeded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"registrant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RegistrationFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"registrant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"ether_token_addr\"},{\"type\":\"address\",\"name\":\"voting_addr\"},{\"type\":\"address\",\"name\":\"p11r_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getPrivileged\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":483},{\"name\":\"setPrivileged\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"listing\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35657},{\"name\":\"getHash\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"string\",\"name\":\"url\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":585},{\"name\":\"getBackendAddress\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":573},{\"name\":\"getBackendUrl\",\"outputs\":[{\"type\":\"string\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":21837},{\"name\":\"setBackendUrl\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"url\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":176937},{\"name\":\"getDataHash\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":796},{\"name\":\"setDataHash\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"listing\"},{\"type\":\"bytes32\",\"name\":\"data\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35919},{\"name\":\"removeDataHash\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":20940},{\"name\":\"register\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"url\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":186298},{\"name\":\"resolveRegistration\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":48356},{\"name\":\"purchaseByteCredits\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":38884},{\"name\":\"getByteCredits\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1015}]"

// DatatrustBin is the compiled bytecode used for deploying new contracts.
const DatatrustBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a0526060610adb6101403934156100a757600080fd5b6020610adb60c03960c05160205181106100c057600080fd5b5060206020610adb0160c03960c05160205181106100dd57600080fd5b5060206040610adb0160c03960c05160205181106100fa57600080fd5b5033600755610140516004556101605160055561018051600655610ac356600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263d4c1753960005114156100b95734156100ac57600080fd5b60085460005260206000f3005b632ecace9c600051141561010357602060046101403734156100da57600080fd5b60043560205181106100eb57600080fd5b5060075433146100fa57600080fd5b61014051600855005b635b6beeb9600051141561015a576020600461014037341561012457600080fd5b60a060043560040161016037608060043560040135111561014457600080fd5b61016080516020820120905060005260206000f3005b63edb39a40600051141561018057341561017357600080fd5b60035460005260206000f3005b6376e12635600051141561026357341561019957600080fd5b60028060c052602060c020610180602082540161012060006005818352015b826101205160200211156101cb576101ed565b61012051850154610120516020028501525b81516001018083528114156101b8575b5050505050506101805160206001820306601f8201039050610240610180516080818352015b826102405111156102235761023f565b6000610240516101a001535b8151600101808352811415610213575b5050506020610160526040610180510160206001820306601f8201039050610160f3005b6341d28f90600051141561030e576020600461014037341561028457600080fd5b60a06004356004016101603760806004356004013511156102a457600080fd5b60035433146102b257600080fd5b61016080600260c052602060c020602082510161012060006005818352015b826101205160200211156102e457610306565b61012051602002850151610120518501555b81516001018083528114156102d1575b505050505050005b637a639f6e600051141561034b576020600461014037341561032f57600080fd5b60006101405160e05260c052604060c0205460005260206000f3005b63b818bf026000511415610392576040600461014037341561036c57600080fd5b600354331461037a57600080fd5b6101605160006101405160e05260c052604060c02055005b63bd82badc60005114156103d757602060046101403734156103b357600080fd5b60085433146103c157600080fd5b600060006101405160e05260c052604060c02055005b63f2c298be60005114156105f557602060046101403734156103f857600080fd5b60a060043560040161016037608060043560040135111561041857600080fd5b60035433141561042757600080fd5b61016080600260c052602060c020602082510161012060006005818352015b826101205160200211156104595761047b565b61012051602002850151610120518501555b8151600101808352811415610446575b505050505050610160805160208201209050610220526005543b61049e57600080fd5b6005543014156104ad57600080fd5b60206102c0602463b89694c661024052610220516102605261025c6005545afa6104d657600080fd5b6000506102c051156104e757600080fd5b6005543b6104f457600080fd5b60055430141561050357600080fd5b6000600060a463bb12c49e6103e052610220516104005260046104205233610440526006543b61053257600080fd5b60065430141561054157600080fd5b6020610340600463fc0e3d906102e0526102fc6006545afa61056257600080fd5b60005061034051610460526006543b61057a57600080fd5b60065430141561058957600080fd5b60206103c06004632d0d2bc66103605261037c6006545afa6105aa57600080fd5b6000506103c051610480526103fc60006005545af16105c857600080fd5b33610220517f7d917fcbc9a29a9705ff9936ffa599500e4fd902e4486bae317414fe967b307c60006000a3005b6384e1fe1560005114156108ae576020600461014037341561061657600080fd5b6005543b61062357600080fd5b60055430141561063257600080fd5b6020610200604463af61f76061016052610140516101805260046101a05261017c6005545afa61066157600080fd5b6000506102005161067157600080fd5b6005543b61067e57600080fd5b60055430141561068d57600080fd5b60206102a0602463327322c861022052610140516102405261023c6005545afa6106b657600080fd5b6000506102a0516106c657600080fd5b6005543b6106d357600080fd5b6005543014156106e257600080fd5b6020610360602463eb3014fd6102e05261014051610300526102fc6005545afa61070b57600080fd5b600050610360516102c0526005543b61072357600080fd5b60055430141561073257600080fd5b60206104a06044638f354b796104005261014051610420526006543b61075757600080fd5b60065430141561076657600080fd5b60206103e0600463c26c12eb6103805261039c6006545afa61078757600080fd5b6000506103e0516104405261041c6005545afa6107a357600080fd5b6000506104a051156107e9576102c0516003556102c051610140517ff9749f013eb1a881b147fd6da901e63089fadfb6fb375d6e56babcbcb5e0be4e60006000a3610866565b6102c051610140517ff83db24154eb020b1b0c94c9566e464732758c2c6bc070062458777d038baa3c60006000a3600080600260c052602060c020600161012060006001818352015b826101205160200211156108455761085f565b6000610120518501555b8151600101808352811415610832575b5050505050505b6005543b61087357600080fd5b60055430141561088257600080fd5b6000600060246389bb617c6104c052610140516104e0526104dc60006005545af16108ac57600080fd5b005b6373fa31a9600051141561095557602060046101403734156108cf57600080fd5b6004543b6108dc57600080fd5b6004543014156108eb57600080fd5b602061022060646323b872dd610160523361018052306101a052610140516101c05261017c60006004545af161092057600080fd5b6000506102205060013360e05260c052604060c020805461014051825401101561094957600080fd5b61014051815401815550005b637ebdb30860005114156109a4576020600461014037341561097657600080fd5b600435602051811061098757600080fd5b5060016101405160e05260c052604060c0205460005260206000f3005b60006000fd5b610119610ac303610119600039610119610ac3036000f3`

// DeployDatatrust deploys a new Ethereum contract, binding an instance of Datatrust to it.
func DeployDatatrust(auth *bind.TransactOpts, backend bind.ContractBackend, ether_token_addr common.Address, voting_addr common.Address, p11r_addr common.Address) (common.Address, *types.Transaction, *Datatrust, error) {
	parsed, err := abi.JSON(strings.NewReader(DatatrustABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DatatrustBin), backend, ether_token_addr, voting_addr, p11r_addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Datatrust{DatatrustCaller: DatatrustCaller{contract: contract}, DatatrustTransactor: DatatrustTransactor{contract: contract}, DatatrustFilterer: DatatrustFilterer{contract: contract}}, nil
}

// Datatrust is an auto generated Go binding around an Ethereum contract.
type Datatrust struct {
	DatatrustCaller     // Read-only binding to the contract
	DatatrustTransactor // Write-only binding to the contract
	DatatrustFilterer   // Log filterer for contract events
}

// DatatrustCaller is an auto generated read-only Go binding around an Ethereum contract.
type DatatrustCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatatrustTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DatatrustTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatatrustFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DatatrustFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatatrustSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DatatrustSession struct {
	Contract     *Datatrust        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DatatrustCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DatatrustCallerSession struct {
	Contract *DatatrustCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DatatrustTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DatatrustTransactorSession struct {
	Contract     *DatatrustTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DatatrustRaw is an auto generated low-level Go binding around an Ethereum contract.
type DatatrustRaw struct {
	Contract *Datatrust // Generic contract binding to access the raw methods on
}

// DatatrustCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DatatrustCallerRaw struct {
	Contract *DatatrustCaller // Generic read-only contract binding to access the raw methods on
}

// DatatrustTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DatatrustTransactorRaw struct {
	Contract *DatatrustTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDatatrust creates a new instance of Datatrust, bound to a specific deployed contract.
func NewDatatrust(address common.Address, backend bind.ContractBackend) (*Datatrust, error) {
	contract, err := bindDatatrust(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Datatrust{DatatrustCaller: DatatrustCaller{contract: contract}, DatatrustTransactor: DatatrustTransactor{contract: contract}, DatatrustFilterer: DatatrustFilterer{contract: contract}}, nil
}

// NewDatatrustCaller creates a new read-only instance of Datatrust, bound to a specific deployed contract.
func NewDatatrustCaller(address common.Address, caller bind.ContractCaller) (*DatatrustCaller, error) {
	contract, err := bindDatatrust(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DatatrustCaller{contract: contract}, nil
}

// NewDatatrustTransactor creates a new write-only instance of Datatrust, bound to a specific deployed contract.
func NewDatatrustTransactor(address common.Address, transactor bind.ContractTransactor) (*DatatrustTransactor, error) {
	contract, err := bindDatatrust(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DatatrustTransactor{contract: contract}, nil
}

// NewDatatrustFilterer creates a new log filterer instance of Datatrust, bound to a specific deployed contract.
func NewDatatrustFilterer(address common.Address, filterer bind.ContractFilterer) (*DatatrustFilterer, error) {
	contract, err := bindDatatrust(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DatatrustFilterer{contract: contract}, nil
}

// bindDatatrust binds a generic wrapper to an already deployed contract.
func bindDatatrust(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DatatrustABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datatrust *DatatrustRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Datatrust.Contract.DatatrustCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datatrust *DatatrustRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datatrust.Contract.DatatrustTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datatrust *DatatrustRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datatrust.Contract.DatatrustTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datatrust *DatatrustCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Datatrust.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datatrust *DatatrustTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datatrust.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datatrust *DatatrustTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datatrust.Contract.contract.Transact(opts, method, params...)
}

// GetBackendAddress is a free data retrieval call binding the contract method 0xedb39a40.
//
// Solidity: function getBackendAddress() constant returns(out address)
func (_Datatrust *DatatrustCaller) GetBackendAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getBackendAddress")
	return *ret0, err
}

// GetBackendAddress is a free data retrieval call binding the contract method 0xedb39a40.
//
// Solidity: function getBackendAddress() constant returns(out address)
func (_Datatrust *DatatrustSession) GetBackendAddress() (common.Address, error) {
	return _Datatrust.Contract.GetBackendAddress(&_Datatrust.CallOpts)
}

// GetBackendAddress is a free data retrieval call binding the contract method 0xedb39a40.
//
// Solidity: function getBackendAddress() constant returns(out address)
func (_Datatrust *DatatrustCallerSession) GetBackendAddress() (common.Address, error) {
	return _Datatrust.Contract.GetBackendAddress(&_Datatrust.CallOpts)
}

// GetBackendUrl is a free data retrieval call binding the contract method 0x76e12635.
//
// Solidity: function getBackendUrl() constant returns(out string)
func (_Datatrust *DatatrustCaller) GetBackendUrl(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getBackendUrl")
	return *ret0, err
}

// GetBackendUrl is a free data retrieval call binding the contract method 0x76e12635.
//
// Solidity: function getBackendUrl() constant returns(out string)
func (_Datatrust *DatatrustSession) GetBackendUrl() (string, error) {
	return _Datatrust.Contract.GetBackendUrl(&_Datatrust.CallOpts)
}

// GetBackendUrl is a free data retrieval call binding the contract method 0x76e12635.
//
// Solidity: function getBackendUrl() constant returns(out string)
func (_Datatrust *DatatrustCallerSession) GetBackendUrl() (string, error) {
	return _Datatrust.Contract.GetBackendUrl(&_Datatrust.CallOpts)
}

// GetByteCredits is a free data retrieval call binding the contract method 0x7ebdb308.
//
// Solidity: function getByteCredits(addr address) constant returns(out uint256)
func (_Datatrust *DatatrustCaller) GetByteCredits(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getByteCredits", addr)
	return *ret0, err
}

// GetByteCredits is a free data retrieval call binding the contract method 0x7ebdb308.
//
// Solidity: function getByteCredits(addr address) constant returns(out uint256)
func (_Datatrust *DatatrustSession) GetByteCredits(addr common.Address) (*big.Int, error) {
	return _Datatrust.Contract.GetByteCredits(&_Datatrust.CallOpts, addr)
}

// GetByteCredits is a free data retrieval call binding the contract method 0x7ebdb308.
//
// Solidity: function getByteCredits(addr address) constant returns(out uint256)
func (_Datatrust *DatatrustCallerSession) GetByteCredits(addr common.Address) (*big.Int, error) {
	return _Datatrust.Contract.GetByteCredits(&_Datatrust.CallOpts, addr)
}

// GetDataHash is a free data retrieval call binding the contract method 0x7a639f6e.
//
// Solidity: function getDataHash(hash bytes32) constant returns(out bytes32)
func (_Datatrust *DatatrustCaller) GetDataHash(opts *bind.CallOpts, hash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getDataHash", hash)
	return *ret0, err
}

// GetDataHash is a free data retrieval call binding the contract method 0x7a639f6e.
//
// Solidity: function getDataHash(hash bytes32) constant returns(out bytes32)
func (_Datatrust *DatatrustSession) GetDataHash(hash [32]byte) ([32]byte, error) {
	return _Datatrust.Contract.GetDataHash(&_Datatrust.CallOpts, hash)
}

// GetDataHash is a free data retrieval call binding the contract method 0x7a639f6e.
//
// Solidity: function getDataHash(hash bytes32) constant returns(out bytes32)
func (_Datatrust *DatatrustCallerSession) GetDataHash(hash [32]byte) ([32]byte, error) {
	return _Datatrust.Contract.GetDataHash(&_Datatrust.CallOpts, hash)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(url string) constant returns(out bytes32)
func (_Datatrust *DatatrustCaller) GetHash(opts *bind.CallOpts, url string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getHash", url)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(url string) constant returns(out bytes32)
func (_Datatrust *DatatrustSession) GetHash(url string) ([32]byte, error) {
	return _Datatrust.Contract.GetHash(&_Datatrust.CallOpts, url)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(url string) constant returns(out bytes32)
func (_Datatrust *DatatrustCallerSession) GetHash(url string) ([32]byte, error) {
	return _Datatrust.Contract.GetHash(&_Datatrust.CallOpts, url)
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address)
func (_Datatrust *DatatrustCaller) GetPrivileged(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getPrivileged")
	return *ret0, err
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address)
func (_Datatrust *DatatrustSession) GetPrivileged() (common.Address, error) {
	return _Datatrust.Contract.GetPrivileged(&_Datatrust.CallOpts)
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address)
func (_Datatrust *DatatrustCallerSession) GetPrivileged() (common.Address, error) {
	return _Datatrust.Contract.GetPrivileged(&_Datatrust.CallOpts)
}

// PurchaseByteCredits is a paid mutator transaction binding the contract method 0x73fa31a9.
//
// Solidity: function purchaseByteCredits(amount uint256) returns()
func (_Datatrust *DatatrustTransactor) PurchaseByteCredits(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "purchaseByteCredits", amount)
}

// PurchaseByteCredits is a paid mutator transaction binding the contract method 0x73fa31a9.
//
// Solidity: function purchaseByteCredits(amount uint256) returns()
func (_Datatrust *DatatrustSession) PurchaseByteCredits(amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.PurchaseByteCredits(&_Datatrust.TransactOpts, amount)
}

// PurchaseByteCredits is a paid mutator transaction binding the contract method 0x73fa31a9.
//
// Solidity: function purchaseByteCredits(amount uint256) returns()
func (_Datatrust *DatatrustTransactorSession) PurchaseByteCredits(amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.PurchaseByteCredits(&_Datatrust.TransactOpts, amount)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(url string) returns()
func (_Datatrust *DatatrustTransactor) Register(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "register", url)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(url string) returns()
func (_Datatrust *DatatrustSession) Register(url string) (*types.Transaction, error) {
	return _Datatrust.Contract.Register(&_Datatrust.TransactOpts, url)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(url string) returns()
func (_Datatrust *DatatrustTransactorSession) Register(url string) (*types.Transaction, error) {
	return _Datatrust.Contract.Register(&_Datatrust.TransactOpts, url)
}

// RemoveDataHash is a paid mutator transaction binding the contract method 0xbd82badc.
//
// Solidity: function removeDataHash(hash bytes32) returns()
func (_Datatrust *DatatrustTransactor) RemoveDataHash(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "removeDataHash", hash)
}

// RemoveDataHash is a paid mutator transaction binding the contract method 0xbd82badc.
//
// Solidity: function removeDataHash(hash bytes32) returns()
func (_Datatrust *DatatrustSession) RemoveDataHash(hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.RemoveDataHash(&_Datatrust.TransactOpts, hash)
}

// RemoveDataHash is a paid mutator transaction binding the contract method 0xbd82badc.
//
// Solidity: function removeDataHash(hash bytes32) returns()
func (_Datatrust *DatatrustTransactorSession) RemoveDataHash(hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.RemoveDataHash(&_Datatrust.TransactOpts, hash)
}

// ResolveRegistration is a paid mutator transaction binding the contract method 0x84e1fe15.
//
// Solidity: function resolveRegistration(hash bytes32) returns()
func (_Datatrust *DatatrustTransactor) ResolveRegistration(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "resolveRegistration", hash)
}

// ResolveRegistration is a paid mutator transaction binding the contract method 0x84e1fe15.
//
// Solidity: function resolveRegistration(hash bytes32) returns()
func (_Datatrust *DatatrustSession) ResolveRegistration(hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.ResolveRegistration(&_Datatrust.TransactOpts, hash)
}

// ResolveRegistration is a paid mutator transaction binding the contract method 0x84e1fe15.
//
// Solidity: function resolveRegistration(hash bytes32) returns()
func (_Datatrust *DatatrustTransactorSession) ResolveRegistration(hash [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.ResolveRegistration(&_Datatrust.TransactOpts, hash)
}

// SetBackendUrl is a paid mutator transaction binding the contract method 0x41d28f90.
//
// Solidity: function setBackendUrl(url string) returns()
func (_Datatrust *DatatrustTransactor) SetBackendUrl(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "setBackendUrl", url)
}

// SetBackendUrl is a paid mutator transaction binding the contract method 0x41d28f90.
//
// Solidity: function setBackendUrl(url string) returns()
func (_Datatrust *DatatrustSession) SetBackendUrl(url string) (*types.Transaction, error) {
	return _Datatrust.Contract.SetBackendUrl(&_Datatrust.TransactOpts, url)
}

// SetBackendUrl is a paid mutator transaction binding the contract method 0x41d28f90.
//
// Solidity: function setBackendUrl(url string) returns()
func (_Datatrust *DatatrustTransactorSession) SetBackendUrl(url string) (*types.Transaction, error) {
	return _Datatrust.Contract.SetBackendUrl(&_Datatrust.TransactOpts, url)
}

// SetDataHash is a paid mutator transaction binding the contract method 0xb818bf02.
//
// Solidity: function setDataHash(listing bytes32, data bytes32) returns()
func (_Datatrust *DatatrustTransactor) SetDataHash(opts *bind.TransactOpts, listing [32]byte, data [32]byte) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "setDataHash", listing, data)
}

// SetDataHash is a paid mutator transaction binding the contract method 0xb818bf02.
//
// Solidity: function setDataHash(listing bytes32, data bytes32) returns()
func (_Datatrust *DatatrustSession) SetDataHash(listing [32]byte, data [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.SetDataHash(&_Datatrust.TransactOpts, listing, data)
}

// SetDataHash is a paid mutator transaction binding the contract method 0xb818bf02.
//
// Solidity: function setDataHash(listing bytes32, data bytes32) returns()
func (_Datatrust *DatatrustTransactorSession) SetDataHash(listing [32]byte, data [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.SetDataHash(&_Datatrust.TransactOpts, listing, data)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0x2ecace9c.
//
// Solidity: function setPrivileged(listing address) returns()
func (_Datatrust *DatatrustTransactor) SetPrivileged(opts *bind.TransactOpts, listing common.Address) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "setPrivileged", listing)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0x2ecace9c.
//
// Solidity: function setPrivileged(listing address) returns()
func (_Datatrust *DatatrustSession) SetPrivileged(listing common.Address) (*types.Transaction, error) {
	return _Datatrust.Contract.SetPrivileged(&_Datatrust.TransactOpts, listing)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0x2ecace9c.
//
// Solidity: function setPrivileged(listing address) returns()
func (_Datatrust *DatatrustTransactorSession) SetPrivileged(listing common.Address) (*types.Transaction, error) {
	return _Datatrust.Contract.SetPrivileged(&_Datatrust.TransactOpts, listing)
}

// DatatrustRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Datatrust contract.
type DatatrustRegisteredIterator struct {
	Event *DatatrustRegistered // Event containing the contract specifics and raw log

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
func (it *DatatrustRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatatrustRegistered)
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
		it.Event = new(DatatrustRegistered)
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
func (it *DatatrustRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatatrustRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatatrustRegistered represents a Registered event raised by the Datatrust contract.
type DatatrustRegistered struct {
	Hash       [32]byte
	Registrant common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x7d917fcbc9a29a9705ff9936ffa599500e4fd902e4486bae317414fe967b307c.
//
// Solidity: e Registered(hash indexed bytes32, registrant indexed address)
func (_Datatrust *DatatrustFilterer) FilterRegistered(opts *bind.FilterOpts, hash [][32]byte, registrant []common.Address) (*DatatrustRegisteredIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _Datatrust.contract.FilterLogs(opts, "Registered", hashRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return &DatatrustRegisteredIterator{contract: _Datatrust.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x7d917fcbc9a29a9705ff9936ffa599500e4fd902e4486bae317414fe967b307c.
//
// Solidity: e Registered(hash indexed bytes32, registrant indexed address)
func (_Datatrust *DatatrustFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *DatatrustRegistered, hash [][32]byte, registrant []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _Datatrust.contract.WatchLogs(opts, "Registered", hashRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatatrustRegistered)
				if err := _Datatrust.contract.UnpackLog(event, "Registered", log); err != nil {
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

// DatatrustRegistrationFailedIterator is returned from FilterRegistrationFailed and is used to iterate over the raw logs and unpacked data for RegistrationFailed events raised by the Datatrust contract.
type DatatrustRegistrationFailedIterator struct {
	Event *DatatrustRegistrationFailed // Event containing the contract specifics and raw log

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
func (it *DatatrustRegistrationFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatatrustRegistrationFailed)
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
		it.Event = new(DatatrustRegistrationFailed)
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
func (it *DatatrustRegistrationFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatatrustRegistrationFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatatrustRegistrationFailed represents a RegistrationFailed event raised by the Datatrust contract.
type DatatrustRegistrationFailed struct {
	Hash       [32]byte
	Registrant common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistrationFailed is a free log retrieval operation binding the contract event 0xf83db24154eb020b1b0c94c9566e464732758c2c6bc070062458777d038baa3c.
//
// Solidity: e RegistrationFailed(hash indexed bytes32, registrant indexed address)
func (_Datatrust *DatatrustFilterer) FilterRegistrationFailed(opts *bind.FilterOpts, hash [][32]byte, registrant []common.Address) (*DatatrustRegistrationFailedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _Datatrust.contract.FilterLogs(opts, "RegistrationFailed", hashRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return &DatatrustRegistrationFailedIterator{contract: _Datatrust.contract, event: "RegistrationFailed", logs: logs, sub: sub}, nil
}

// WatchRegistrationFailed is a free log subscription operation binding the contract event 0xf83db24154eb020b1b0c94c9566e464732758c2c6bc070062458777d038baa3c.
//
// Solidity: e RegistrationFailed(hash indexed bytes32, registrant indexed address)
func (_Datatrust *DatatrustFilterer) WatchRegistrationFailed(opts *bind.WatchOpts, sink chan<- *DatatrustRegistrationFailed, hash [][32]byte, registrant []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _Datatrust.contract.WatchLogs(opts, "RegistrationFailed", hashRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatatrustRegistrationFailed)
				if err := _Datatrust.contract.UnpackLog(event, "RegistrationFailed", log); err != nil {
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

// DatatrustRegistrationSucceededIterator is returned from FilterRegistrationSucceeded and is used to iterate over the raw logs and unpacked data for RegistrationSucceeded events raised by the Datatrust contract.
type DatatrustRegistrationSucceededIterator struct {
	Event *DatatrustRegistrationSucceeded // Event containing the contract specifics and raw log

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
func (it *DatatrustRegistrationSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatatrustRegistrationSucceeded)
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
		it.Event = new(DatatrustRegistrationSucceeded)
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
func (it *DatatrustRegistrationSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatatrustRegistrationSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatatrustRegistrationSucceeded represents a RegistrationSucceeded event raised by the Datatrust contract.
type DatatrustRegistrationSucceeded struct {
	Hash       [32]byte
	Registrant common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistrationSucceeded is a free log retrieval operation binding the contract event 0xf9749f013eb1a881b147fd6da901e63089fadfb6fb375d6e56babcbcb5e0be4e.
//
// Solidity: e RegistrationSucceeded(hash indexed bytes32, registrant indexed address)
func (_Datatrust *DatatrustFilterer) FilterRegistrationSucceeded(opts *bind.FilterOpts, hash [][32]byte, registrant []common.Address) (*DatatrustRegistrationSucceededIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _Datatrust.contract.FilterLogs(opts, "RegistrationSucceeded", hashRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return &DatatrustRegistrationSucceededIterator{contract: _Datatrust.contract, event: "RegistrationSucceeded", logs: logs, sub: sub}, nil
}

// WatchRegistrationSucceeded is a free log subscription operation binding the contract event 0xf9749f013eb1a881b147fd6da901e63089fadfb6fb375d6e56babcbcb5e0be4e.
//
// Solidity: e RegistrationSucceeded(hash indexed bytes32, registrant indexed address)
func (_Datatrust *DatatrustFilterer) WatchRegistrationSucceeded(opts *bind.WatchOpts, sink chan<- *DatatrustRegistrationSucceeded, hash [][32]byte, registrant []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var registrantRule []interface{}
	for _, registrantItem := range registrant {
		registrantRule = append(registrantRule, registrantItem)
	}

	logs, sub, err := _Datatrust.contract.WatchLogs(opts, "RegistrationSucceeded", hashRule, registrantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatatrustRegistrationSucceeded)
				if err := _Datatrust.contract.UnpackLog(event, "RegistrationSucceeded", log); err != nil {
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
