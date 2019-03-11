// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package investing

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

// InvestingABI is the input ABI used to generate the binding from.
const InvestingABI = "[{\"name\":\"Divested\",\"inputs\":[{\"type\":\"address\",\"name\":\"investor\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"transferred\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Invested\",\"inputs\":[{\"type\":\"address\",\"name\":\"investor\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"offered\",\"indexed\":false,\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"minted\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"ether_token_addr\"},{\"type\":\"address\",\"name\":\"market_token_addr\"},{\"type\":\"address\",\"name\":\"voting_addr\"},{\"type\":\"address\",\"name\":\"p11r_addr\"},{\"type\":\"address\",\"name\":\"listing_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"isInvestor\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1042},{\"name\":\"getInvested\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":513},{\"name\":\"getInvestorCount\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":543},{\"name\":\"getInvestment\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":823},{\"name\":\"getInvestmentPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":12400},{\"name\":\"invest\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"offer\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":208145},{\"name\":\"getDivestmentProceeds\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6170},{\"name\":\"divest\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":3854752}]"

// InvestingBin is the compiled bytecode used for deploying new contracts.
const InvestingBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05260a06111836101403934156100a757600080fd5b602061118360c03960c05160205181106100c057600080fd5b50602060206111830160c03960c05160205181106100dd57600080fd5b50602060406111830160c03960c05160205181106100fa57600080fd5b50602060606111830160c03960c051602051811061011757600080fd5b50602060806111830160c03960c051602051811061013457600080fd5b50336004556101405160055561016051600655610180516007556101a0516008556101c05160095561116b56600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263cee2a9cf600051141561010a57602060046101403734156100b457600080fd5b60043560205181106100c557600080fd5b506101405160006101405160e05260c052604060c02060c052602060c02054620186a081106100f357600080fd5b600260c052602060c02001541460005260206000f3005b63befc3e2b600051141561013057341561012357600080fd5b60035460005260206000f3005b63960524e3600051141561015657341561014957600080fd5b60015460005260206000f3005b63146b58df60005114156101b0576020600461014037341561017757600080fd5b600435602051811061018857600080fd5b50600160006101405160e05260c052604060c02060c052602060c020015460005260206000f3005b63a3d6602b60005114156107395734156101c957600080fd5b6008543b6101d657600080fd5b6008543014156101e557600080fd5b60206101c0600463f36089ec6101605261017c6008545afa61020657600080fd5b6000506101c051610140526008543b61021e57600080fd5b60085430141561022d57600080fd5b60206102606004634ba135d86102005261021c6008545afa61024e57600080fd5b600050610260516101e0526008543b61026657600080fd5b60085430141561027557600080fd5b602061030060046328dee5706102a0526102bc6008545afa61029657600080fd5b60005061030051610280526005543b6102ae57600080fd5b6005543014156102bd57600080fd5b60206103c060246370a0823161034052306103605261035c6005545afa6102e357600080fd5b6000506103c051610320526006543b6102fb57600080fd5b60065430141561030a57600080fd5b602061046060046318160ddd6104005261041c6006545afa61032b57600080fd5b600050610460516103e052670de0b6b3a76400006103e051101561040157610140516101e05161035a57600080fd5b6101e05161028051151561036f576000610395565b6103205161028051610320516102805102041461038b57600080fd5b6103205161028051025b04610140510110156103a657600080fd5b6101e0516103b357600080fd5b6101e0516102805115156103c85760006103ee565b610320516102805161032051610280510204146103e457600080fd5b6103205161028051025b04610140510160005260206000f3610737565b610140516101e051151561041657600061043c565b6103e0516101e0516103e0516101e05102041461043257600080fd5b6103e0516101e051025b61044557600080fd5b6101e051151561045657600061047c565b6103e0516101e0516103e0516101e05102041461047257600080fd5b6103e0516101e051025b61028051151561048d5760006104b3565b610320516102805161032051610280510204146104a957600080fd5b6103205161028051025b15156104c057600061058e565b670de0b6b3a76400006102805115156104da576000610500565b610320516102805161032051610280510204146104f657600080fd5b6103205161028051025b670de0b6b3a764000061028051151561051a576000610540565b6103205161028051610320516102805102041461053657600080fd5b6103205161028051025b02041461054c57600080fd5b670de0b6b3a764000061028051151561056657600061058c565b6103205161028051610320516102805102041461058257600080fd5b6103205161028051025b025b046101405101101561059f57600080fd5b6101e05115156105b05760006105d6565b6103e0516101e0516103e0516101e0510204146105cc57600080fd5b6103e0516101e051025b6105df57600080fd5b6101e05115156105f0576000610616565b6103e0516101e0516103e0516101e05102041461060c57600080fd5b6103e0516101e051025b61028051151561062757600061064d565b6103205161028051610320516102805102041461064357600080fd5b6103205161028051025b151561065a576000610728565b670de0b6b3a764000061028051151561067457600061069a565b6103205161028051610320516102805102041461069057600080fd5b6103205161028051025b670de0b6b3a76400006102805115156106b45760006106da565b610320516102805161032051610280510204146106d057600080fd5b6103205161028051025b0204146106e657600080fd5b670de0b6b3a7640000610280511515610700576000610726565b6103205161028051610320516102805102041461071c57600080fd5b6103205161028051025b025b04610140510160005260206000f35b005b632afcf4806000511415610ac5576020600461014037341561075a57600080fd5b60206101e0600463a3d6602b6101805261019c6000305af161077b57600080fd5b6101e051610160526101605161014051101561079657600080fd5b6005543b6107a357600080fd5b6005543014156107b257600080fd5b60206102c060646323b872dd6102005233610220523061024052610140516102605261021c60006005545af16107e757600080fd5b6000506102c050610160516107fb57600080fd5b6101605161014051041515610811576000610870565b633b9aca006101605161082357600080fd5b610160516101405104633b9aca006101605161083e57600080fd5b61016051610140510402041461085357600080fd5b633b9aca006101605161086557600080fd5b610160516101405104025b6102e0526006543b61088157600080fd5b60065430141561089057600080fd5b60006000602463a0712d68610300526102e0516103205261031c60006006545af16108ba57600080fd5b6006543b6108c757600080fd5b6006543014156108d657600080fd5b6020610420604463a9059cbb61038052336103a0526102e0516103c05261039c60006006545af161090657600080fd5b6000506104205060206104c0602463cee2a9cf61044052336104605261045c6000305af161093357600080fd5b6104c0511515610a345760015460003360e05260c052604060c02060c052602060c0205533600154620186a0811061096a57600080fd5b600260c052602060c0200155600160605160018254018060405190131561099057600080fd5b809190121561099e57600080fd5b8155506007543b6109ae57600080fd5b6007543014156109bd57600080fd5b6020610560602463c666b5806104e05233610500526104fc6007545afa6109e357600080fd5b6000506105605115610a33576007543b6109fc57600080fd5b600754301415610a0b57600080fd5b60006000602463471bb30961058052336105a05261059c60006007545af1610a3257600080fd5b5b5b600160003360e05260c052604060c02060c052602060c020018054610140518254011015610a6157600080fd5b6101405181540181555060038054610140518254011015610a8157600080fd5b6101405181540181555061014051610600526102e05161062052337f9e9d071824fd57d062ca63fd8b786d8da48a6807eebbcb2d83f9e8d21398e28c6040610600a2005b63998339036000511415610c415760206004610140373415610ae657600080fd5b6004356020518110610af757600080fd5b506000610140511415610b0957600080fd5b6006543b610b1657600080fd5b600654301415610b2557600080fd5b602061020060246370a0823161018052610140516101a05261019c6006545afa610b4e57600080fd5b60005061020051610160526005543b610b6657600080fd5b600554301415610b7557600080fd5b60206102c060246370a0823161024052306102605261025c6005545afa610b9b57600080fd5b6000506102c051610220526006543b610bb357600080fd5b600654301415610bc257600080fd5b602061036060046318160ddd6103005261031c6006545afa610be357600080fd5b600050610360516102e0526102e051610bfb57600080fd5b6102e051610160511515610c10576000610c36565b61022051610160516102205161016051020414610c2c57600080fd5b6102205161016051025b0460005260206000f3005b63058aace16000511415611004573415610c5a57600080fd5b6000610140526007543b610c6d57600080fd5b600754301415610c7c57600080fd5b60206101e060046330a563476101805261019c6007545afa610c9d57600080fd5b6000506101e0516101605261020060006103e8818352015b6007543b610cc257600080fd5b600754301415610cd157600080fd5b60206102c060246379b7392461024052610200516102605261025c6007545afa610cfa57600080fd5b6000506102c051610220526009543b610d1257600080fd5b600954301415610d2157600080fd5b6020610380602463458d2bf161030052610220516103205261031c6009545afa610d4a57600080fd5b600050610380516102e05261016051610200511415610d6c57610d9456610d83565b336102e0511415610d8257600161014052610d94565b5b5b8151600101808352811415610cb5575b50506101405115610da457600080fd5b6020610440602463998339036103c052336103e0526103dc6000305af1610dca57600080fd5b610440516103a0526006543b610ddf57600080fd5b600654301415610dee57600080fd5b600060006024637e9d2ac161046052336104805261047c60006006545af1610e1557600080fd5b6020610560602463cee2a9cf6104e05233610500526104fc6000305af1610e3b57600080fd5b6105605115610f7f576001606051600182540380604051901315610e5e57600080fd5b8091901215610e6c57600080fd5b81555060006001541315610f5e5760003360e05260c052604060c02060c052602060c0205461058052600154620186a08110610ea757600080fd5b600260c052602060c02001546105a0526105a05161058051620186a08110610ece57600080fd5b600260c052602060c02001556105805160006105a05160e05260c052604060c02060c052602060c0205560003360e05260c052604060c02060c052602060c0206000815560006001820155506007543b610f2757600080fd5b600754301415610f3657600080fd5b60006000602463335d80806105c052336105e0526105dc60006007545af1610f5d57600080fd5b5b6000600154620186a08110610f7257600080fd5b600260c052602060c02001555b6005543b610f8c57600080fd5b600554301415610f9b57600080fd5b60206106e0604463a9059cbb6106405233610660526103a0516106805261065c60006005545af1610fcb57600080fd5b6000506106e0506103a05161070052337f2253aebe2fe8682635bbe60d9b78df72efaf785a596910a8ad66e8c6e37584fd6020610700a2005b60006000fd5b61016161116b0361016160003961016161116b036000f3`

// DeployInvesting deploys a new Ethereum contract, binding an instance of Investing to it.
func DeployInvesting(auth *bind.TransactOpts, backend bind.ContractBackend, ether_token_addr common.Address, market_token_addr common.Address, voting_addr common.Address, p11r_addr common.Address, listing_addr common.Address) (common.Address, *types.Transaction, *Investing, error) {
	parsed, err := abi.JSON(strings.NewReader(InvestingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InvestingBin), backend, ether_token_addr, market_token_addr, voting_addr, p11r_addr, listing_addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Investing{InvestingCaller: InvestingCaller{contract: contract}, InvestingTransactor: InvestingTransactor{contract: contract}, InvestingFilterer: InvestingFilterer{contract: contract}}, nil
}

// Investing is an auto generated Go binding around an Ethereum contract.
type Investing struct {
	InvestingCaller     // Read-only binding to the contract
	InvestingTransactor // Write-only binding to the contract
	InvestingFilterer   // Log filterer for contract events
}

// InvestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type InvestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InvestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InvestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InvestingSession struct {
	Contract     *Investing        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InvestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InvestingCallerSession struct {
	Contract *InvestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// InvestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InvestingTransactorSession struct {
	Contract     *InvestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// InvestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type InvestingRaw struct {
	Contract *Investing // Generic contract binding to access the raw methods on
}

// InvestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InvestingCallerRaw struct {
	Contract *InvestingCaller // Generic read-only contract binding to access the raw methods on
}

// InvestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InvestingTransactorRaw struct {
	Contract *InvestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInvesting creates a new instance of Investing, bound to a specific deployed contract.
func NewInvesting(address common.Address, backend bind.ContractBackend) (*Investing, error) {
	contract, err := bindInvesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Investing{InvestingCaller: InvestingCaller{contract: contract}, InvestingTransactor: InvestingTransactor{contract: contract}, InvestingFilterer: InvestingFilterer{contract: contract}}, nil
}

// NewInvestingCaller creates a new read-only instance of Investing, bound to a specific deployed contract.
func NewInvestingCaller(address common.Address, caller bind.ContractCaller) (*InvestingCaller, error) {
	contract, err := bindInvesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InvestingCaller{contract: contract}, nil
}

// NewInvestingTransactor creates a new write-only instance of Investing, bound to a specific deployed contract.
func NewInvestingTransactor(address common.Address, transactor bind.ContractTransactor) (*InvestingTransactor, error) {
	contract, err := bindInvesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InvestingTransactor{contract: contract}, nil
}

// NewInvestingFilterer creates a new log filterer instance of Investing, bound to a specific deployed contract.
func NewInvestingFilterer(address common.Address, filterer bind.ContractFilterer) (*InvestingFilterer, error) {
	contract, err := bindInvesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InvestingFilterer{contract: contract}, nil
}

// bindInvesting binds a generic wrapper to an already deployed contract.
func bindInvesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InvestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Investing *InvestingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Investing.Contract.InvestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Investing *InvestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Investing.Contract.InvestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Investing *InvestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Investing.Contract.InvestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Investing *InvestingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Investing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Investing *InvestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Investing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Investing *InvestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Investing.Contract.contract.Transact(opts, method, params...)
}

// GetDivestmentProceeds is a free data retrieval call binding the contract method 0x99833903.
//
// Solidity: function getDivestmentProceeds(addr address) constant returns(out uint256)
func (_Investing *InvestingCaller) GetDivestmentProceeds(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Investing.contract.Call(opts, out, "getDivestmentProceeds", addr)
	return *ret0, err
}

// GetDivestmentProceeds is a free data retrieval call binding the contract method 0x99833903.
//
// Solidity: function getDivestmentProceeds(addr address) constant returns(out uint256)
func (_Investing *InvestingSession) GetDivestmentProceeds(addr common.Address) (*big.Int, error) {
	return _Investing.Contract.GetDivestmentProceeds(&_Investing.CallOpts, addr)
}

// GetDivestmentProceeds is a free data retrieval call binding the contract method 0x99833903.
//
// Solidity: function getDivestmentProceeds(addr address) constant returns(out uint256)
func (_Investing *InvestingCallerSession) GetDivestmentProceeds(addr common.Address) (*big.Int, error) {
	return _Investing.Contract.GetDivestmentProceeds(&_Investing.CallOpts, addr)
}

// GetInvested is a free data retrieval call binding the contract method 0xbefc3e2b.
//
// Solidity: function getInvested() constant returns(out uint256)
func (_Investing *InvestingCaller) GetInvested(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Investing.contract.Call(opts, out, "getInvested")
	return *ret0, err
}

// GetInvested is a free data retrieval call binding the contract method 0xbefc3e2b.
//
// Solidity: function getInvested() constant returns(out uint256)
func (_Investing *InvestingSession) GetInvested() (*big.Int, error) {
	return _Investing.Contract.GetInvested(&_Investing.CallOpts)
}

// GetInvested is a free data retrieval call binding the contract method 0xbefc3e2b.
//
// Solidity: function getInvested() constant returns(out uint256)
func (_Investing *InvestingCallerSession) GetInvested() (*big.Int, error) {
	return _Investing.Contract.GetInvested(&_Investing.CallOpts)
}

// GetInvestment is a free data retrieval call binding the contract method 0x146b58df.
//
// Solidity: function getInvestment(addr address) constant returns(out uint256)
func (_Investing *InvestingCaller) GetInvestment(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Investing.contract.Call(opts, out, "getInvestment", addr)
	return *ret0, err
}

// GetInvestment is a free data retrieval call binding the contract method 0x146b58df.
//
// Solidity: function getInvestment(addr address) constant returns(out uint256)
func (_Investing *InvestingSession) GetInvestment(addr common.Address) (*big.Int, error) {
	return _Investing.Contract.GetInvestment(&_Investing.CallOpts, addr)
}

// GetInvestment is a free data retrieval call binding the contract method 0x146b58df.
//
// Solidity: function getInvestment(addr address) constant returns(out uint256)
func (_Investing *InvestingCallerSession) GetInvestment(addr common.Address) (*big.Int, error) {
	return _Investing.Contract.GetInvestment(&_Investing.CallOpts, addr)
}

// GetInvestmentPrice is a free data retrieval call binding the contract method 0xa3d6602b.
//
// Solidity: function getInvestmentPrice() constant returns(out uint256)
func (_Investing *InvestingCaller) GetInvestmentPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Investing.contract.Call(opts, out, "getInvestmentPrice")
	return *ret0, err
}

// GetInvestmentPrice is a free data retrieval call binding the contract method 0xa3d6602b.
//
// Solidity: function getInvestmentPrice() constant returns(out uint256)
func (_Investing *InvestingSession) GetInvestmentPrice() (*big.Int, error) {
	return _Investing.Contract.GetInvestmentPrice(&_Investing.CallOpts)
}

// GetInvestmentPrice is a free data retrieval call binding the contract method 0xa3d6602b.
//
// Solidity: function getInvestmentPrice() constant returns(out uint256)
func (_Investing *InvestingCallerSession) GetInvestmentPrice() (*big.Int, error) {
	return _Investing.Contract.GetInvestmentPrice(&_Investing.CallOpts)
}

// GetInvestorCount is a free data retrieval call binding the contract method 0x960524e3.
//
// Solidity: function getInvestorCount() constant returns(out int128)
func (_Investing *InvestingCaller) GetInvestorCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Investing.contract.Call(opts, out, "getInvestorCount")
	return *ret0, err
}

// GetInvestorCount is a free data retrieval call binding the contract method 0x960524e3.
//
// Solidity: function getInvestorCount() constant returns(out int128)
func (_Investing *InvestingSession) GetInvestorCount() (*big.Int, error) {
	return _Investing.Contract.GetInvestorCount(&_Investing.CallOpts)
}

// GetInvestorCount is a free data retrieval call binding the contract method 0x960524e3.
//
// Solidity: function getInvestorCount() constant returns(out int128)
func (_Investing *InvestingCallerSession) GetInvestorCount() (*big.Int, error) {
	return _Investing.Contract.GetInvestorCount(&_Investing.CallOpts)
}

// IsInvestor is a free data retrieval call binding the contract method 0xcee2a9cf.
//
// Solidity: function isInvestor(addr address) constant returns(out bool)
func (_Investing *InvestingCaller) IsInvestor(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Investing.contract.Call(opts, out, "isInvestor", addr)
	return *ret0, err
}

// IsInvestor is a free data retrieval call binding the contract method 0xcee2a9cf.
//
// Solidity: function isInvestor(addr address) constant returns(out bool)
func (_Investing *InvestingSession) IsInvestor(addr common.Address) (bool, error) {
	return _Investing.Contract.IsInvestor(&_Investing.CallOpts, addr)
}

// IsInvestor is a free data retrieval call binding the contract method 0xcee2a9cf.
//
// Solidity: function isInvestor(addr address) constant returns(out bool)
func (_Investing *InvestingCallerSession) IsInvestor(addr common.Address) (bool, error) {
	return _Investing.Contract.IsInvestor(&_Investing.CallOpts, addr)
}

// Divest is a paid mutator transaction binding the contract method 0x058aace1.
//
// Solidity: function divest() returns()
func (_Investing *InvestingTransactor) Divest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Investing.contract.Transact(opts, "divest")
}

// Divest is a paid mutator transaction binding the contract method 0x058aace1.
//
// Solidity: function divest() returns()
func (_Investing *InvestingSession) Divest() (*types.Transaction, error) {
	return _Investing.Contract.Divest(&_Investing.TransactOpts)
}

// Divest is a paid mutator transaction binding the contract method 0x058aace1.
//
// Solidity: function divest() returns()
func (_Investing *InvestingTransactorSession) Divest() (*types.Transaction, error) {
	return _Investing.Contract.Divest(&_Investing.TransactOpts)
}

// Invest is a paid mutator transaction binding the contract method 0x2afcf480.
//
// Solidity: function invest(offer uint256) returns()
func (_Investing *InvestingTransactor) Invest(opts *bind.TransactOpts, offer *big.Int) (*types.Transaction, error) {
	return _Investing.contract.Transact(opts, "invest", offer)
}

// Invest is a paid mutator transaction binding the contract method 0x2afcf480.
//
// Solidity: function invest(offer uint256) returns()
func (_Investing *InvestingSession) Invest(offer *big.Int) (*types.Transaction, error) {
	return _Investing.Contract.Invest(&_Investing.TransactOpts, offer)
}

// Invest is a paid mutator transaction binding the contract method 0x2afcf480.
//
// Solidity: function invest(offer uint256) returns()
func (_Investing *InvestingTransactorSession) Invest(offer *big.Int) (*types.Transaction, error) {
	return _Investing.Contract.Invest(&_Investing.TransactOpts, offer)
}

// InvestingDivestedIterator is returned from FilterDivested and is used to iterate over the raw logs and unpacked data for Divested events raised by the Investing contract.
type InvestingDivestedIterator struct {
	Event *InvestingDivested // Event containing the contract specifics and raw log

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
func (it *InvestingDivestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InvestingDivested)
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
		it.Event = new(InvestingDivested)
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
func (it *InvestingDivestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InvestingDivestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InvestingDivested represents a Divested event raised by the Investing contract.
type InvestingDivested struct {
	Investor    common.Address
	Transferred *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDivested is a free log retrieval operation binding the contract event 0x2253aebe2fe8682635bbe60d9b78df72efaf785a596910a8ad66e8c6e37584fd.
//
// Solidity: e Divested(investor indexed address, transferred uint256)
func (_Investing *InvestingFilterer) FilterDivested(opts *bind.FilterOpts, investor []common.Address) (*InvestingDivestedIterator, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}

	logs, sub, err := _Investing.contract.FilterLogs(opts, "Divested", investorRule)
	if err != nil {
		return nil, err
	}
	return &InvestingDivestedIterator{contract: _Investing.contract, event: "Divested", logs: logs, sub: sub}, nil
}

// WatchDivested is a free log subscription operation binding the contract event 0x2253aebe2fe8682635bbe60d9b78df72efaf785a596910a8ad66e8c6e37584fd.
//
// Solidity: e Divested(investor indexed address, transferred uint256)
func (_Investing *InvestingFilterer) WatchDivested(opts *bind.WatchOpts, sink chan<- *InvestingDivested, investor []common.Address) (event.Subscription, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}

	logs, sub, err := _Investing.contract.WatchLogs(opts, "Divested", investorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InvestingDivested)
				if err := _Investing.contract.UnpackLog(event, "Divested", log); err != nil {
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

// InvestingInvestedIterator is returned from FilterInvested and is used to iterate over the raw logs and unpacked data for Invested events raised by the Investing contract.
type InvestingInvestedIterator struct {
	Event *InvestingInvested // Event containing the contract specifics and raw log

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
func (it *InvestingInvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InvestingInvested)
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
		it.Event = new(InvestingInvested)
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
func (it *InvestingInvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InvestingInvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InvestingInvested represents a Invested event raised by the Investing contract.
type InvestingInvested struct {
	Investor common.Address
	Offered  *big.Int
	Minted   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvested is a free log retrieval operation binding the contract event 0x9e9d071824fd57d062ca63fd8b786d8da48a6807eebbcb2d83f9e8d21398e28c.
//
// Solidity: e Invested(investor indexed address, offered uint256, minted uint256)
func (_Investing *InvestingFilterer) FilterInvested(opts *bind.FilterOpts, investor []common.Address) (*InvestingInvestedIterator, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}

	logs, sub, err := _Investing.contract.FilterLogs(opts, "Invested", investorRule)
	if err != nil {
		return nil, err
	}
	return &InvestingInvestedIterator{contract: _Investing.contract, event: "Invested", logs: logs, sub: sub}, nil
}

// WatchInvested is a free log subscription operation binding the contract event 0x9e9d071824fd57d062ca63fd8b786d8da48a6807eebbcb2d83f9e8d21398e28c.
//
// Solidity: e Invested(investor indexed address, offered uint256, minted uint256)
func (_Investing *InvestingFilterer) WatchInvested(opts *bind.WatchOpts, sink chan<- *InvestingInvested, investor []common.Address) (event.Subscription, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}

	logs, sub, err := _Investing.contract.WatchLogs(opts, "Invested", investorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InvestingInvested)
				if err := _Investing.contract.UnpackLog(event, "Invested", log); err != nil {
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
