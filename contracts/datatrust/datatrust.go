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
const DatatrustABI = "[{\"name\":\"Registered\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"registrant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RegistrationSucceeded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"registrant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RegistrationFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"registrant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"DeliveryRequested\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"requester\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Delivered\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"bytes32\",\"name\":\"url\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"ether_token_addr\"},{\"type\":\"address\",\"name\":\"voting_addr\"},{\"type\":\"address\",\"name\":\"p11r_addr\"},{\"type\":\"address\",\"name\":\"res_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getPrivileged\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":551},{\"name\":\"getReserve\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":581},{\"name\":\"setPrivileged\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"listing\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35998},{\"name\":\"getHash\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"string\",\"name\":\"url\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":665},{\"name\":\"getBackendAddress\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":671},{\"name\":\"getBackendUrl\",\"outputs\":[{\"type\":\"string\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":22831},{\"name\":\"setBackendUrl\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"url\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":176987},{\"name\":\"getDataHash\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":876},{\"name\":\"setDataHash\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"listing\"},{\"type\":\"bytes32\",\"name\":\"data\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35963},{\"name\":\"removeDataHash\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":20990},{\"name\":\"register\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"url\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":186333},{\"name\":\"resolveRegistration\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":48388},{\"name\":\"requestDelivery\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":116397},{\"name\":\"getBytesPurchased\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1095},{\"name\":\"getDelivery\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2037},{\"name\":\"listingAccessed\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"listing\"},{\"type\":\"bytes32\",\"name\":\"delivery\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":109369},{\"name\":\"getBytesAccessed\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1146},{\"name\":\"bytesAccessedClaimed\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"unit\":\"wei\",\"name\":\"fee\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":23804},{\"name\":\"delivered\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"delivery\"},{\"type\":\"bytes32\",\"name\":\"url\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":77140}]"

// DatatrustBin is the compiled bytecode used for deploying new contracts.
const DatatrustBin = `0x740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052608061128b6101403934156100a157600080fd5b602061128b60c03960c05160205181106100ba57600080fd5b506020602061128b0160c03960c05160205181106100d757600080fd5b506020604061128b0160c03960c05160205181106100f457600080fd5b506020606061128b0160c03960c051602051811061011157600080fd5b50336009556101405160065561016051600755610180516008556101a051600a5561127356600436101561000d57611136565b600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263d4c1753960005114156100c85734156100ba57600080fd5b600b5460005260206000f350005b6359bf5d3960005114156100ef5734156100e157600080fd5b600a5460005260206000f350005b632ecace9c600051141561013d57341561010857600080fd5b600435602051811061011957600080fd5b50600954331461012857600080fd5b600b541561013557600080fd5b600435600b55005b635b6beeb9600051141561018d57341561015657600080fd5b60a060043560040161014037608060043560040135111561017657600080fd5b61014080516020820120905060005260206000f350005b63edb39a4060005114156101b45734156101a657600080fd5b60055460005260206000f350005b6376e1263560005114156102995734156101cd57600080fd5b60048060c052602060c020610180602082540161012060006005818352015b826101205160200211156101ff57610221565b61012051850154610120516020028501525b81516001018083528114156101ec575b5050505050506101805160206001820306601f8201039050610240610180516080818352015b826102405110151561025857610274565b6000610240516101a001535b8151600101808352811415610247575b5050506020610160526040610180510160206001820306601f8201039050610160f350005b6341d28f90600051141561033c5734156102b257600080fd5b60a06004356004016101403760806004356004013511156102d257600080fd5b60055433146102e057600080fd5b61014080600460c052602060c020602082510161012060006005818352015b8261012051602002111561031257610334565b61012051602002850151610120518501555b81516001018083528114156102ff575b505050505050005b637a639f6e600051141561037157341561035557600080fd5b600060043560e05260c052604060c0205460005260206000f350005b63b818bf0260005114156103ae57341561038a57600080fd5b600554331461039857600080fd5b602435600060043560e05260c052604060c02055005b63bd82badc60005114156103ea5734156103c757600080fd5b600b5433146103d557600080fd5b6000600060043560e05260c052604060c02055005b63f2c298be60005114156105fb57341561040357600080fd5b60a060043560040161014037608060043560040135111561042357600080fd5b600554331861043157600080fd5b61014080600460c052602060c020602082510161012060006005818352015b8261012051602002111561046357610485565b61012051602002850151610120518501555b8151600101808352811415610450575b505050505050610140805160208201209050610200526007543b6104a857600080fd5b60075430186104b657600080fd5b60206102a0602463b89694c661022052610200516102405261023c6007545afa6104df57600080fd5b6000506102a051156104f057600080fd5b6007543b6104fd57600080fd5b600754301861050b57600080fd5b6000600060a463bb12c49e6103c052610200516103e05260046104005233610420526008543b61053a57600080fd5b600854301861054857600080fd5b6020610320600463fc0e3d906102c0526102dc6008545afa61056957600080fd5b60005061032051610440526008543b61058157600080fd5b600854301861058f57600080fd5b60206103a06004632d0d2bc66103405261035c6008545afa6105b057600080fd5b6000506103a051610460526103dc60006007545af16105ce57600080fd5b33610200517f7d917fcbc9a29a9705ff9936ffa599500e4fd902e4486bae317414fe967b307c60006000a3005b6384e1fe15600051141561089f57341561061457600080fd5b6007543b61062157600080fd5b600754301861062f57600080fd5b60206101e0604463af61f760610140526004356101605260046101805261015c6007545afa61065d57600080fd5b6000506101e05161066d57600080fd5b6007543b61067a57600080fd5b600754301861068857600080fd5b6020610280602463327322c8610200526004356102205261021c6007545afa6106b057600080fd5b600050610280516106c057600080fd5b6007543b6106cd57600080fd5b60075430186106db57600080fd5b6020610340602463eb3014fd6102c0526004356102e0526102dc6007545afa61070357600080fd5b600050610340516102a0526007543b61071b57600080fd5b600754301861072957600080fd5b60206104806044638f354b796103e052600435610400526008543b61074d57600080fd5b600854301861075b57600080fd5b60206103c06004633d1e37d56103605261037c6008545afa61077c57600080fd5b6000506103c051610420526103fc6007545afa61079857600080fd5b60005061048051156107dd576102a0516005556102a0516004357ff9749f013eb1a881b147fd6da901e63089fadfb6fb375d6e56babcbcb5e0be4e60006000a3610859565b6102a0516004357ff83db24154eb020b1b0c94c9566e464732758c2c6bc070062458777d038baa3c60006000a3600080600460c052602060c020600161012060006001818352015b8261012051602002111561083857610852565b6000610120518501555b8151600101808352811415610825575b5050505050505b6007543b61086657600080fd5b600754301861087457600080fd5b6000600060246389bb617c6104a0526004356104c0526104bc60006007545af161089d57600080fd5b005b630ee206f56000511415610afa5734156108b857600080fd5b600360043560e05260c052604060c02060c052602060c02054156108db57600080fd5b6008543b6108e857600080fd5b60085430186108f657600080fd5b60206101c0600463cd2d54386101605261017c6008545afa61091757600080fd5b6000506101c0516101e0526101e0511515610933576000610956565b6024356101e0516024356101e05102041461094d57600080fd5b6024356101e051025b61014052606461096557600080fd5b60646008543b61097457600080fd5b600854301861098257600080fd5b6020610280600463ffe44da86102205261023c6008545afa6109a357600080fd5b600050610280516102a0526101405115156109bf5760006109e5565b6102a051610140516102a051610140510204146109db57600080fd5b6102a05161014051025b04610200526006543b6109f757600080fd5b6006543018610a0557600080fd5b602061038060646323b872dd6102c052336102e052306103005261014051610320526102dc60006006545af1610a3a57600080fd5b600050610380506006543b610a4e57600080fd5b6006543018610a5c57600080fd5b6020610440604463a9059cbb6103a052600a546103c052610200516103e0526103bc60006006545af1610a8e57600080fd5b6000506104405060013360e05260c052604060c02080546024358254011015610ab657600080fd5b60243581540181555033600360043560e05260c052604060c02060c052602060c020556024356001600360043560e05260c052604060c02060c052602060c0200155005b63f2cbc3fe6000511415610b41573415610b1357600080fd5b6004356020518110610b2457600080fd5b50600160043560e05260c052604060c0205460005260206000f350005b63edab743e6000511415610bcc573415610b5a57600080fd5b606061014052610160600360043560e05260c052604060c02060c052602060c0205481526001600360043560e05260c052604060c02060c052602060c020015481602001526002600360043560e05260c052604060c02060c052602060c020015481604001525061014051610160f350005b63043b21666000511415610cb4573415610be557600080fd5b6005543314610bf357600080fd5b6000600060043560e05260c052604060c0205418610c1057600080fd5b600260043560e05260c052604060c02080546044358254011015610c3357600080fd5b6044358154018155506001600360243560e05260c052604060c02060c052602060c0205460e05260c052604060c02060443581541015610c7257600080fd5b6044358154038155506002600360243560e05260c052604060c02060c052602060c0200180546044358254011015610ca957600080fd5b604435815401815550005b63e80239d56000511415610ce9573415610ccd57600080fd5b600260043560e05260c052604060c0205460005260206000f350005b639d38d6da6000511415610d78573415610d0257600080fd5b600b543314610d1057600080fd5b6000600260043560e05260c052604060c020556006543b610d3057600080fd5b6006543018610d3e57600080fd5b60206101e0604463a9059cbb61014052600a54610160526024356101805261015c60006006545af1610d6f57600080fd5b6000506101e050005b63d96804126000511415611135573415610d9157600080fd5b6005543314610d9f57600080fd5b600360043560e05260c052604060c02060c052602060c02054610140526001600360043560e05260c052604060c02060c052602060c020015461016052610160516002600360043560e05260c052604060c02060c052602060c02001541015610e0757600080fd5b600360043560e05260c052604060c02060c052602060c020600081556000600182015560006002820155506064610e3d57600080fd5b60646008543b610e4c57600080fd5b6008543018610e5a57600080fd5b60206102a06004634b8b7dce6102405261025c6008545afa610e7b57600080fd5b6000506102a0516102c0526008543b610e9357600080fd5b6008543018610ea157600080fd5b6020610200600463cd2d54386101a0526101bc6008545afa610ec257600080fd5b6000506102005161022052610220511515610ede576000610f04565b61016051610220516101605161022051020414610efa57600080fd5b6101605161022051025b1515610f115760006110a5565b6102c0516008543b610f2257600080fd5b6008543018610f3057600080fd5b6020610200600463cd2d54386101a0526101bc6008545afa610f5157600080fd5b6000506102005161022052610220511515610f6d576000610f93565b61016051610220516101605161022051020414610f8957600080fd5b6101605161022051025b6102c0516008543b610fa457600080fd5b6008543018610fb257600080fd5b6020610200600463cd2d54386101a0526101bc6008545afa610fd357600080fd5b6000506102005161022052610220511515610fef576000611015565b6101605161022051610160516102205102041461100b57600080fd5b6101605161022051025b02041461102157600080fd5b6102c0516008543b61103257600080fd5b600854301861104057600080fd5b6020610200600463cd2d54386101a0526101bc6008545afa61106157600080fd5b600050610200516102205261022051151561107d5760006110a3565b6101605161022051610160516102205102041461109957600080fd5b6101605161022051025b025b04610180526006543b6110b757600080fd5b60065430186110c557600080fd5b6020610380604463a9059cbb6102e0526005546103005261018051610320526102fc60006006545af16110f757600080fd5b600050610380506024356103a052610140516004357fa6bd58fcd4da90af9fea785d8cf919f762887f0f05b0656fb2ac5f7227183d5a60206103a0a3005b5b60006000fd5b61013761127303610137600039610137611273036000f3`

// DeployDatatrust deploys a new Ethereum contract, binding an instance of Datatrust to it.
func DeployDatatrust(auth *bind.TransactOpts, backend bind.ContractBackend, ether_token_addr common.Address, voting_addr common.Address, p11r_addr common.Address, res_addr common.Address) (common.Address, *types.Transaction, *Datatrust, error) {
	parsed, err := abi.JSON(strings.NewReader(DatatrustABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DatatrustBin), backend, ether_token_addr, voting_addr, p11r_addr, res_addr)
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

// GetBytesAccessed is a free data retrieval call binding the contract method 0xe80239d5.
//
// Solidity: function getBytesAccessed(hash bytes32) constant returns(out uint256)
func (_Datatrust *DatatrustCaller) GetBytesAccessed(opts *bind.CallOpts, hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getBytesAccessed", hash)
	return *ret0, err
}

// GetBytesAccessed is a free data retrieval call binding the contract method 0xe80239d5.
//
// Solidity: function getBytesAccessed(hash bytes32) constant returns(out uint256)
func (_Datatrust *DatatrustSession) GetBytesAccessed(hash [32]byte) (*big.Int, error) {
	return _Datatrust.Contract.GetBytesAccessed(&_Datatrust.CallOpts, hash)
}

// GetBytesAccessed is a free data retrieval call binding the contract method 0xe80239d5.
//
// Solidity: function getBytesAccessed(hash bytes32) constant returns(out uint256)
func (_Datatrust *DatatrustCallerSession) GetBytesAccessed(hash [32]byte) (*big.Int, error) {
	return _Datatrust.Contract.GetBytesAccessed(&_Datatrust.CallOpts, hash)
}

// GetBytesPurchased is a free data retrieval call binding the contract method 0xf2cbc3fe.
//
// Solidity: function getBytesPurchased(addr address) constant returns(out uint256)
func (_Datatrust *DatatrustCaller) GetBytesPurchased(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getBytesPurchased", addr)
	return *ret0, err
}

// GetBytesPurchased is a free data retrieval call binding the contract method 0xf2cbc3fe.
//
// Solidity: function getBytesPurchased(addr address) constant returns(out uint256)
func (_Datatrust *DatatrustSession) GetBytesPurchased(addr common.Address) (*big.Int, error) {
	return _Datatrust.Contract.GetBytesPurchased(&_Datatrust.CallOpts, addr)
}

// GetBytesPurchased is a free data retrieval call binding the contract method 0xf2cbc3fe.
//
// Solidity: function getBytesPurchased(addr address) constant returns(out uint256)
func (_Datatrust *DatatrustCallerSession) GetBytesPurchased(addr common.Address) (*big.Int, error) {
	return _Datatrust.Contract.GetBytesPurchased(&_Datatrust.CallOpts, addr)
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

// GetDelivery is a free data retrieval call binding the contract method 0xedab743e.
//
// Solidity: function getDelivery(hash bytes32) constant returns(out address, out uint256, out uint256)
func (_Datatrust *DatatrustCaller) GetDelivery(opts *bind.CallOpts, hash [32]byte) (common.Address, *big.Int, *big.Int, error) {
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
	err := _Datatrust.contract.Call(opts, out, "getDelivery", hash)
	return *ret0, *ret1, *ret2, err
}

// GetDelivery is a free data retrieval call binding the contract method 0xedab743e.
//
// Solidity: function getDelivery(hash bytes32) constant returns(out address, out uint256, out uint256)
func (_Datatrust *DatatrustSession) GetDelivery(hash [32]byte) (common.Address, *big.Int, *big.Int, error) {
	return _Datatrust.Contract.GetDelivery(&_Datatrust.CallOpts, hash)
}

// GetDelivery is a free data retrieval call binding the contract method 0xedab743e.
//
// Solidity: function getDelivery(hash bytes32) constant returns(out address, out uint256, out uint256)
func (_Datatrust *DatatrustCallerSession) GetDelivery(hash [32]byte) (common.Address, *big.Int, *big.Int, error) {
	return _Datatrust.Contract.GetDelivery(&_Datatrust.CallOpts, hash)
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

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() constant returns(out address)
func (_Datatrust *DatatrustCaller) GetReserve(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Datatrust.contract.Call(opts, out, "getReserve")
	return *ret0, err
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() constant returns(out address)
func (_Datatrust *DatatrustSession) GetReserve() (common.Address, error) {
	return _Datatrust.Contract.GetReserve(&_Datatrust.CallOpts)
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() constant returns(out address)
func (_Datatrust *DatatrustCallerSession) GetReserve() (common.Address, error) {
	return _Datatrust.Contract.GetReserve(&_Datatrust.CallOpts)
}

// BytesAccessedClaimed is a paid mutator transaction binding the contract method 0x9d38d6da.
//
// Solidity: function bytesAccessedClaimed(hash bytes32, fee uint256) returns()
func (_Datatrust *DatatrustTransactor) BytesAccessedClaimed(opts *bind.TransactOpts, hash [32]byte, fee *big.Int) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "bytesAccessedClaimed", hash, fee)
}

// BytesAccessedClaimed is a paid mutator transaction binding the contract method 0x9d38d6da.
//
// Solidity: function bytesAccessedClaimed(hash bytes32, fee uint256) returns()
func (_Datatrust *DatatrustSession) BytesAccessedClaimed(hash [32]byte, fee *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.BytesAccessedClaimed(&_Datatrust.TransactOpts, hash, fee)
}

// BytesAccessedClaimed is a paid mutator transaction binding the contract method 0x9d38d6da.
//
// Solidity: function bytesAccessedClaimed(hash bytes32, fee uint256) returns()
func (_Datatrust *DatatrustTransactorSession) BytesAccessedClaimed(hash [32]byte, fee *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.BytesAccessedClaimed(&_Datatrust.TransactOpts, hash, fee)
}

// Delivered is a paid mutator transaction binding the contract method 0xd9680412.
//
// Solidity: function delivered(delivery bytes32, url bytes32) returns()
func (_Datatrust *DatatrustTransactor) Delivered(opts *bind.TransactOpts, delivery [32]byte, url [32]byte) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "delivered", delivery, url)
}

// Delivered is a paid mutator transaction binding the contract method 0xd9680412.
//
// Solidity: function delivered(delivery bytes32, url bytes32) returns()
func (_Datatrust *DatatrustSession) Delivered(delivery [32]byte, url [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.Delivered(&_Datatrust.TransactOpts, delivery, url)
}

// Delivered is a paid mutator transaction binding the contract method 0xd9680412.
//
// Solidity: function delivered(delivery bytes32, url bytes32) returns()
func (_Datatrust *DatatrustTransactorSession) Delivered(delivery [32]byte, url [32]byte) (*types.Transaction, error) {
	return _Datatrust.Contract.Delivered(&_Datatrust.TransactOpts, delivery, url)
}

// ListingAccessed is a paid mutator transaction binding the contract method 0x043b2166.
//
// Solidity: function listingAccessed(listing bytes32, delivery bytes32, amount uint256) returns()
func (_Datatrust *DatatrustTransactor) ListingAccessed(opts *bind.TransactOpts, listing [32]byte, delivery [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "listingAccessed", listing, delivery, amount)
}

// ListingAccessed is a paid mutator transaction binding the contract method 0x043b2166.
//
// Solidity: function listingAccessed(listing bytes32, delivery bytes32, amount uint256) returns()
func (_Datatrust *DatatrustSession) ListingAccessed(listing [32]byte, delivery [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.ListingAccessed(&_Datatrust.TransactOpts, listing, delivery, amount)
}

// ListingAccessed is a paid mutator transaction binding the contract method 0x043b2166.
//
// Solidity: function listingAccessed(listing bytes32, delivery bytes32, amount uint256) returns()
func (_Datatrust *DatatrustTransactorSession) ListingAccessed(listing [32]byte, delivery [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.ListingAccessed(&_Datatrust.TransactOpts, listing, delivery, amount)
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

// RequestDelivery is a paid mutator transaction binding the contract method 0x0ee206f5.
//
// Solidity: function requestDelivery(hash bytes32, amount uint256) returns()
func (_Datatrust *DatatrustTransactor) RequestDelivery(opts *bind.TransactOpts, hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.contract.Transact(opts, "requestDelivery", hash, amount)
}

// RequestDelivery is a paid mutator transaction binding the contract method 0x0ee206f5.
//
// Solidity: function requestDelivery(hash bytes32, amount uint256) returns()
func (_Datatrust *DatatrustSession) RequestDelivery(hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.RequestDelivery(&_Datatrust.TransactOpts, hash, amount)
}

// RequestDelivery is a paid mutator transaction binding the contract method 0x0ee206f5.
//
// Solidity: function requestDelivery(hash bytes32, amount uint256) returns()
func (_Datatrust *DatatrustTransactorSession) RequestDelivery(hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Datatrust.Contract.RequestDelivery(&_Datatrust.TransactOpts, hash, amount)
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

// DatatrustDeliveredIterator is returned from FilterDelivered and is used to iterate over the raw logs and unpacked data for Delivered events raised by the Datatrust contract.
type DatatrustDeliveredIterator struct {
	Event *DatatrustDelivered // Event containing the contract specifics and raw log

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
func (it *DatatrustDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatatrustDelivered)
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
		it.Event = new(DatatrustDelivered)
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
func (it *DatatrustDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatatrustDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatatrustDelivered represents a Delivered event raised by the Datatrust contract.
type DatatrustDelivered struct {
	Hash  [32]byte
	Owner common.Address
	Url   [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDelivered is a free log retrieval operation binding the contract event 0xa6bd58fcd4da90af9fea785d8cf919f762887f0f05b0656fb2ac5f7227183d5a.
//
// Solidity: e Delivered(hash indexed bytes32, owner indexed address, url bytes32)
func (_Datatrust *DatatrustFilterer) FilterDelivered(opts *bind.FilterOpts, hash [][32]byte, owner []common.Address) (*DatatrustDeliveredIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Datatrust.contract.FilterLogs(opts, "Delivered", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &DatatrustDeliveredIterator{contract: _Datatrust.contract, event: "Delivered", logs: logs, sub: sub}, nil
}

// WatchDelivered is a free log subscription operation binding the contract event 0xa6bd58fcd4da90af9fea785d8cf919f762887f0f05b0656fb2ac5f7227183d5a.
//
// Solidity: e Delivered(hash indexed bytes32, owner indexed address, url bytes32)
func (_Datatrust *DatatrustFilterer) WatchDelivered(opts *bind.WatchOpts, sink chan<- *DatatrustDelivered, hash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Datatrust.contract.WatchLogs(opts, "Delivered", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatatrustDelivered)
				if err := _Datatrust.contract.UnpackLog(event, "Delivered", log); err != nil {
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

// DatatrustDeliveryRequestedIterator is returned from FilterDeliveryRequested and is used to iterate over the raw logs and unpacked data for DeliveryRequested events raised by the Datatrust contract.
type DatatrustDeliveryRequestedIterator struct {
	Event *DatatrustDeliveryRequested // Event containing the contract specifics and raw log

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
func (it *DatatrustDeliveryRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatatrustDeliveryRequested)
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
		it.Event = new(DatatrustDeliveryRequested)
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
func (it *DatatrustDeliveryRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatatrustDeliveryRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatatrustDeliveryRequested represents a DeliveryRequested event raised by the Datatrust contract.
type DatatrustDeliveryRequested struct {
	Hash      [32]byte
	Requester common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeliveryRequested is a free log retrieval operation binding the contract event 0x898564ce29843afc6dc09bcabe85faaec3c18c45244d2c42e811a25bc0fd82c1.
//
// Solidity: e DeliveryRequested(hash indexed bytes32, requester indexed address, amount uint256)
func (_Datatrust *DatatrustFilterer) FilterDeliveryRequested(opts *bind.FilterOpts, hash [][32]byte, requester []common.Address) (*DatatrustDeliveryRequestedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Datatrust.contract.FilterLogs(opts, "DeliveryRequested", hashRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &DatatrustDeliveryRequestedIterator{contract: _Datatrust.contract, event: "DeliveryRequested", logs: logs, sub: sub}, nil
}

// WatchDeliveryRequested is a free log subscription operation binding the contract event 0x898564ce29843afc6dc09bcabe85faaec3c18c45244d2c42e811a25bc0fd82c1.
//
// Solidity: e DeliveryRequested(hash indexed bytes32, requester indexed address, amount uint256)
func (_Datatrust *DatatrustFilterer) WatchDeliveryRequested(opts *bind.WatchOpts, sink chan<- *DatatrustDeliveryRequested, hash [][32]byte, requester []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Datatrust.contract.WatchLogs(opts, "DeliveryRequested", hashRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatatrustDeliveryRequested)
				if err := _Datatrust.contract.UnpackLog(event, "DeliveryRequested", log); err != nil {
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
