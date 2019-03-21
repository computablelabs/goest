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
const InvestingABI = "[{\"name\":\"Divested\",\"inputs\":[{\"type\":\"address\",\"name\":\"investor\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"transferred\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Invested\",\"inputs\":[{\"type\":\"address\",\"name\":\"investor\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"offered\",\"indexed\":false,\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"minted\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"ether_token_addr\"},{\"type\":\"address\",\"name\":\"market_token_addr\"},{\"type\":\"address\",\"name\":\"voting_addr\"},{\"type\":\"address\",\"name\":\"p11r_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getInvested\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":483},{\"name\":\"getInvestment\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":685},{\"name\":\"getInvestmentPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":12340},{\"name\":\"invest\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"offer\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":99872},{\"name\":\"getDivestmentProceeds\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6110},{\"name\":\"divest\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":36891}]"

// InvestingBin is the compiled bytecode used for deploying new contracts.
const InvestingBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a0526080610ddc6101403934156100a757600080fd5b6020610ddc60c03960c05160205181106100c057600080fd5b5060206020610ddc0160c03960c05160205181106100dd57600080fd5b5060206040610ddc0160c03960c05160205181106100fa57600080fd5b5060206060610ddc0160c03960c051602051811061011757600080fd5b506101405160025561016051600355610180516004556101a051600555610dc456600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263befc3e2b60005114156100b95734156100ac57600080fd5b60015460005260206000f3005b63146b58df600051141561010857602060046101403734156100da57600080fd5b60043560205181106100eb57600080fd5b5060006101405160e05260c052604060c0205460005260206000f3005b63a3d6602b600051141561069157341561012157600080fd5b6005543b61012e57600080fd5b60055430141561013d57600080fd5b60206101c0600463f36089ec6101605261017c6005545afa61015e57600080fd5b6000506101c051610140526005543b61017657600080fd5b60055430141561018557600080fd5b60206102606004634ba135d86102005261021c6005545afa6101a657600080fd5b600050610260516101e0526005543b6101be57600080fd5b6005543014156101cd57600080fd5b602061030060046328dee5706102a0526102bc6005545afa6101ee57600080fd5b60005061030051610280526002543b61020657600080fd5b60025430141561021557600080fd5b60206103c060246370a0823161034052306103605261035c6002545afa61023b57600080fd5b6000506103c051610320526003543b61025357600080fd5b60035430141561026257600080fd5b602061046060046318160ddd6104005261041c6003545afa61028357600080fd5b600050610460516103e052670de0b6b3a76400006103e051101561035957610140516101e0516102b257600080fd5b6101e0516102805115156102c75760006102ed565b610320516102805161032051610280510204146102e357600080fd5b6103205161028051025b04610140510110156102fe57600080fd5b6101e05161030b57600080fd5b6101e051610280511515610320576000610346565b6103205161028051610320516102805102041461033c57600080fd5b6103205161028051025b04610140510160005260206000f361068f565b610140516101e051151561036e576000610394565b6103e0516101e0516103e0516101e05102041461038a57600080fd5b6103e0516101e051025b61039d57600080fd5b6101e05115156103ae5760006103d4565b6103e0516101e0516103e0516101e0510204146103ca57600080fd5b6103e0516101e051025b6102805115156103e557600061040b565b6103205161028051610320516102805102041461040157600080fd5b6103205161028051025b15156104185760006104e6565b670de0b6b3a7640000610280511515610432576000610458565b6103205161028051610320516102805102041461044e57600080fd5b6103205161028051025b670de0b6b3a7640000610280511515610472576000610498565b6103205161028051610320516102805102041461048e57600080fd5b6103205161028051025b0204146104a457600080fd5b670de0b6b3a76400006102805115156104be5760006104e4565b610320516102805161032051610280510204146104da57600080fd5b6103205161028051025b025b04610140510110156104f757600080fd5b6101e051151561050857600061052e565b6103e0516101e0516103e0516101e05102041461052457600080fd5b6103e0516101e051025b61053757600080fd5b6101e051151561054857600061056e565b6103e0516101e0516103e0516101e05102041461056457600080fd5b6103e0516101e051025b61028051151561057f5760006105a5565b6103205161028051610320516102805102041461059b57600080fd5b6103205161028051025b15156105b2576000610680565b670de0b6b3a76400006102805115156105cc5760006105f2565b610320516102805161032051610280510204146105e857600080fd5b6103205161028051025b670de0b6b3a764000061028051151561060c576000610632565b6103205161028051610320516102805102041461062857600080fd5b6103205161028051025b02041461063e57600080fd5b670de0b6b3a764000061028051151561065857600061067e565b6103205161028051610320516102805102041461067457600080fd5b6103205161028051025b025b04610140510160005260206000f35b005b632afcf480600051141561097e57602060046101403734156106b257600080fd5b60206101e0600463a3d6602b6101805261019c6000305af16106d357600080fd5b6101e05161016052610160516101405110156106ee57600080fd5b6002543b6106fb57600080fd5b60025430141561070a57600080fd5b60206102c060646323b872dd6102005233610220523061024052610140516102605261021c60006002545af161073f57600080fd5b6000506102c0506101605161075357600080fd5b61016051610140510415156107695760006107c8565b633b9aca006101605161077b57600080fd5b610160516101405104633b9aca006101605161079657600080fd5b6101605161014051040204146107ab57600080fd5b633b9aca00610160516107bd57600080fd5b610160516101405104025b6102e0526003543b6107d957600080fd5b6003543014156107e857600080fd5b60006000602463a0712d68610300526102e0516103205261031c60006003545af161081257600080fd5b6003543b61081f57600080fd5b60035430141561082e57600080fd5b6020610420604463a9059cbb61038052336103a0526102e0516103c05261039c60006003545af161085e57600080fd5b6000506104205060003360e05260c052604060c020805461014051825401101561088757600080fd5b61014051815401815550600180546101405182540110156108a757600080fd5b610140518154018155506004543b6108be57600080fd5b6004543014156108cd57600080fd5b60206104c060246362f9a55d61044052336104605261045c6004545afa6108f357600080fd5b6000506104c0511515610944576004543b61090d57600080fd5b60045430141561091c57600080fd5b60006000602463471bb3096104e05233610500526104fc60006004545af161094357600080fd5b5b61014051610560526102e05161058052337f9e9d071824fd57d062ca63fd8b786d8da48a6807eebbcb2d83f9e8d21398e28c6040610560a2005b63998339036000511415610afa576020600461014037341561099f57600080fd5b60043560205181106109b057600080fd5b5060006101405114156109c257600080fd5b6003543b6109cf57600080fd5b6003543014156109de57600080fd5b602061020060246370a0823161018052610140516101a05261019c6003545afa610a0757600080fd5b60005061020051610160526002543b610a1f57600080fd5b600254301415610a2e57600080fd5b60206102c060246370a0823161024052306102605261025c6002545afa610a5457600080fd5b6000506102c051610220526003543b610a6c57600080fd5b600354301415610a7b57600080fd5b602061036060046318160ddd6103005261031c6003545afa610a9c57600080fd5b600050610360516102e0526102e051610ab457600080fd5b6102e051610160511515610ac9576000610aef565b61022051610160516102205161016051020414610ae557600080fd5b6102205161016051025b0460005260206000f3005b63058aace16000511415610c85573415610b1357600080fd5b60206101e06024639983390361016052336101805261017c6000305af1610b3957600080fd5b6101e0516101405260006101405111610b5157600080fd5b6003543b610b5e57600080fd5b600354301415610b6d57600080fd5b600060006024637e9d2ac161020052336102205261021c60006003545af1610b9457600080fd5b600060003360e05260c052604060c020541115610c0057600060003360e05260c052604060c020556004543b610bc957600080fd5b600454301415610bd857600080fd5b60006000602463335d808061028052336102a05261029c60006004545af1610bff57600080fd5b5b6002543b610c0d57600080fd5b600254301415610c1c57600080fd5b60206103a0604463a9059cbb610300523361032052610140516103405261031c60006002545af1610c4c57600080fd5b6000506103a050610140516103c052337f2253aebe2fe8682635bbe60d9b78df72efaf785a596910a8ad66e8c6e37584fd60206103c0a2005b60006000fd5b610139610dc403610139600039610139610dc4036000f3`

// DeployInvesting deploys a new Ethereum contract, binding an instance of Investing to it.
func DeployInvesting(auth *bind.TransactOpts, backend bind.ContractBackend, ether_token_addr common.Address, market_token_addr common.Address, voting_addr common.Address, p11r_addr common.Address) (common.Address, *types.Transaction, *Investing, error) {
	parsed, err := abi.JSON(strings.NewReader(InvestingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InvestingBin), backend, ether_token_addr, market_token_addr, voting_addr, p11r_addr)
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
