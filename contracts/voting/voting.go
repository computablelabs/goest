// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voting

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

// VotingABI is the input ABI used to generate the binding from.
const VotingABI = "[{\"name\":\"CandidateAdded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"kind\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"voteBy\",\"indexed\":true,\"unit\":\"sec\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CandidateRemoved\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CouncilMemberAdded\",\"inputs\":[{\"type\":\"address\",\"name\":\"member\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CouncilMemberRemoved\",\"inputs\":[{\"type\":\"address\",\"name\":\"member\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Voted\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"voter\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getPrivileged\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":976},{\"name\":\"setPrivileged\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"market\"},{\"type\":\"address\",\"name\":\"parameterizer\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":70711},{\"name\":\"inCouncil\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"member\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1330},{\"name\":\"getCouncilCount\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":603},{\"name\":\"addToCouncil\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"member\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":108859},{\"name\":\"removeFromCouncil\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"member\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":152452},{\"name\":\"candidateIs\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"kind\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":919},{\"name\":\"isCandidate\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1513},{\"name\":\"wasCandidate\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":964},{\"name\":\"getCandidateCount\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":783},{\"name\":\"getCandidateKey\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"index\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1030},{\"name\":\"getCandidate\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"sec\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1933},{\"name\":\"addCandidate\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"kind\"},{\"type\":\"uint256\",\"name\":\"vote_by\",\"unit\":\"sec\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":180492},{\"name\":\"removeCandidate\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":168052},{\"name\":\"didPass\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"quorum\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6201},{\"name\":\"didVote\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"address\",\"name\":\"member\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":3943},{\"name\":\"pollClosed\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":3812},{\"name\":\"vote\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":84195},{\"name\":\"willAddCandidate\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":5574},{\"name\":\"willAddToCouncil\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":3502}]"

// VotingBin is the compiled bytecode used for deploying new contracts.
const VotingBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052341561009e57600080fd5b33600855610fd156600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263d4c1753960005114156100d55734156100ac57600080fd5b6060610140526101606008548152600654816020015260075481604001525061014051610160f3005b63b01e369f600051141561013857604060046101403734156100f657600080fd5b600435602051811061010757600080fd5b50602435602051811061011957600080fd5b50600854331461012857600080fd5b6101405160065561016051600755005b600015610170575b61016052610140526006546101405114600854610140511417600754610140511417600052600051610160515650005b6362f9a55d60005114156101f2576020600461014037341561019157600080fd5b60043560205181106101a257600080fd5b5060045415156101b757600060005260206000f35b6101405160036101405160e05260c052604060c020546103e881106101db57600080fd5b600560c052602060c02001541460005260206000f3005b634627da50600051141561021857341561020b57600080fd5b60045460005260206000f3005b63471bb309600051141561030c576020600461014037341561023957600080fd5b600435602051811061024a57600080fd5b5061014051636a395d236101605233610180526101805160065801610140565b6101e052610140526101e05161027f57600080fd5b60045460036101405160e05260c052604060c02055610140516004546103e881106102a957600080fd5b600560c052602060c020015560046060516001825401806040519013156102cf57600080fd5b80919012156102dd57600080fd5b815550610140517fdcb46b4634af6065fbbeab63938d70a7cc803a67bc6a5bb9f28fb8d9a5c4ce1460006000a2005b63335d808060005114156104ae576020600461014037341561032d57600080fd5b600435602051811061033e57600080fd5b5061014051636a395d236101605233610180526101805160065801610140565b6101e052610140526101e05161037357600080fd5b602061028060246362f9a55d61020052610140516102205261021c6000305af161039c57600080fd5b610280516103a957600080fd5b60046060516001825403806040519013156103c357600080fd5b80919012156103d157600080fd5b8155506000600454131561044f5760036101405160e05260c052604060c020546102a0526004546103e8811061040657600080fd5b600560c052602060c02001546102c0526102c0516102a0516103e8811061042c57600080fd5b600560c052602060c02001556102a05160036102c05160e05260c052604060c020555b60006004546103e8811061046257600080fd5b600560c052602060c0200155600060036101405160e05260c052604060c02055610140517fce6ed85c24e62fbd23486204879240944c169dfcde2952b0c15f12bf1202966960006000a2005b63af61f76060005114156104fb57604060046101403734156104cf57600080fd5b61016051600160006101405160e05260c052604060c02060c052602060c02001541460005260206000f3005b63b89694c66000511415610573576020600461014037341561051c57600080fd5b600154151561053057600060005260206000f35b6101405160006101405160e05260c052604060c02060c052602060c020546103e8811061055c57600080fd5b600260c052602060c02001541460005260206000f3005b638033814260005114156105da576020600461014037341561059457600080fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60006101405160e05260c052604060c02060c052602060c020541460005260206000f3005b6330a5634760005114156106005734156105f357600080fd5b60015460005260206000f3005b6379b73924600051141561066e576020600461014037341561062157600080fd5b6060516004358060405190131561063757600080fd5b809190121561064557600080fd5b50610140516103e8811061065857600080fd5b600260c052602060c020015460005260206000f3005b63dfb6419f6000511415610706576020600461014037341561068f57600080fd5b606061016052610180600160006101405160e05260c052604060c02060c052602060c02001548152600260006101405160e05260c052604060c02060c052602060c02001548160200152600460006101405160e05260c052604060c02060c052602060c020015481604001525061016051610180f3005b637ff45d8b6000511415610866576060600461014037341561072757600080fd5b610140516101605161018051636a395d236101a052336101c0526101c05160065801610140565b610220526101805261016052610140526102205161076b57600080fd5b42610180514201101561077d57600080fd5b6101805142016102405260015460006101405160e05260c052604060c02060c052602060c0205561016051600160006101405160e05260c052604060c02060c052602060c020015561024051600260006101405160e05260c052604060c02060c052602060c0200155610140516001546103e881106107fb57600080fd5b600260c052602060c0200155600160605160018254018060405190131561082157600080fd5b809190121561082f57600080fd5b8155506102405161016051610140517f4566e4146679483436367293d719cef8eb5d21b6b774ec4cec37b82f224db16160006000a4005b6389bb617c6000511415610a2d576020600461014037341561088757600080fd5b61014051636a395d236101605233610180526101805160065801610140565b6101e052610140526101e0516108bb57600080fd5b6020610280602463b89694c661020052610140516102205261021c6000305af16108e457600080fd5b610280516108f157600080fd5b600160605160018254038060405190131561090b57600080fd5b809190121561091957600080fd5b815550600060015413156109a75760006101405160e05260c052604060c02060c052602060c020546102a0526001546103e8811061095657600080fd5b600260c052602060c02001546102c0526102c0516102a0516103e8811061097c57600080fd5b600260c052602060c02001556102a05160006102c05160e05260c052604060c02060c052602060c020555b60006001546103e881106109ba57600080fd5b600260c052602060c02001557fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60006101405160e05260c052604060c02060c052602060c02055610140517fb314642ce4aa3156c1090458cfccabff50f32bc85d110d7799c369e1a660bcf460006000a2005b638f354b796000511415610bae5760406004610140373415610a4e57600080fd5b6020610200602463b89694c661018052610140516101a05261019c6000305af1610a7757600080fd5b61020051610a8457600080fd5b42600260006101405160e05260c052604060c02060c052602060c020015410610aac57600080fd5b600460006101405160e05260c052604060c02060c052602060c02001541515610ae157610160511560005260206000f3610bac565b610160516004546000811215610af657600080fd5b6102205261022051610b0757600080fd5b61022051600460006101405160e05260c052604060c02060c052602060c02001541515610b35576000610ba0565b6064600460006101405160e05260c052604060c02060c052602060c02001546064600460006101405160e05260c052604060c02060c052602060c0200154020414610b7f57600080fd5b6064600460006101405160e05260c052604060c02060c052602060c0200154025b04101560005260206000f35b005b63187a9c076000511415610c505760406004610140373415610bcf57600080fd5b6024356020518110610be057600080fd5b506020610200602463b89694c661018052610140516101a05261019c6000305af1610c0a57600080fd5b61020051610c1757600080fd5b6001600360006101405160e05260c052604060c02060c052602060c020016101605160e05260c052604060c020541460005260206000f3005b63327322c86000511415610cd05760206004610140373415610c7157600080fd5b60206101e0602463b89694c661016052610140516101805261017c6000305af1610c9a57600080fd5b6101e051610ca757600080fd5b42600260006101405160e05260c052604060c02060c052602060c02001541060005260206000f3005b63a69beaba6000511415610e215760206004610140373415610cf157600080fd5b60206101e060246362f9a55d61016052336101805261017c6000305af1610d1757600080fd5b6101e051610d2457600080fd5b6020610280602463b89694c661020052610140516102205261021c6000305af1610d4d57600080fd5b61028051610d5a57600080fd5b42600260006101405160e05260c052604060c02060c052602060c020015411610d8257600080fd5b6020610340604463187a9c076102a052610140516102c052336102e0526102bc6000305af1610db057600080fd5b6103405115610dbe57600080fd5b6001600360006101405160e05260c052604060c02060c052602060c020013360e05260c052604060c02055600460006101405160e05260c052604060c02060c052602060c02001805460018254011015610e1757600080fd5b6001815401815550005b631f1b02996000511415610eb15760206004610140373415610e4257600080fd5b6020610320602463b89694c66102a052610140516102c0526102bc6000305af1610e6b57600080fd5b61032051156103e8600154101660206103c06024638033814261034052610140516103605261035c6000305af1610ea157600080fd5b6103c051151660005260206000f3005b63c666b5806000511415610f245760206004610140373415610ed257600080fd5b6004356020518110610ee357600080fd5b50602061028060246362f9a55d61020052610140516102205261021c6000305af1610f0d57600080fd5b61028051156103e8600454101660005260206000f3005b60006000fd5b6100a7610fd1036100a76000396100a7610fd1036000f3`

// DeployVoting deploys a new Ethereum contract, binding an instance of Voting to it.
func DeployVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Voting, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// CandidateIs is a free data retrieval call binding the contract method 0xaf61f760.
//
// Solidity: function candidateIs(hash bytes32, kind uint256) constant returns(out bool)
func (_Voting *VotingCaller) CandidateIs(opts *bind.CallOpts, hash [32]byte, kind *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "candidateIs", hash, kind)
	return *ret0, err
}

// CandidateIs is a free data retrieval call binding the contract method 0xaf61f760.
//
// Solidity: function candidateIs(hash bytes32, kind uint256) constant returns(out bool)
func (_Voting *VotingSession) CandidateIs(hash [32]byte, kind *big.Int) (bool, error) {
	return _Voting.Contract.CandidateIs(&_Voting.CallOpts, hash, kind)
}

// CandidateIs is a free data retrieval call binding the contract method 0xaf61f760.
//
// Solidity: function candidateIs(hash bytes32, kind uint256) constant returns(out bool)
func (_Voting *VotingCallerSession) CandidateIs(hash [32]byte, kind *big.Int) (bool, error) {
	return _Voting.Contract.CandidateIs(&_Voting.CallOpts, hash, kind)
}

// DidPass is a free data retrieval call binding the contract method 0x8f354b79.
//
// Solidity: function didPass(hash bytes32, quorum uint256) constant returns(out bool)
func (_Voting *VotingCaller) DidPass(opts *bind.CallOpts, hash [32]byte, quorum *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "didPass", hash, quorum)
	return *ret0, err
}

// DidPass is a free data retrieval call binding the contract method 0x8f354b79.
//
// Solidity: function didPass(hash bytes32, quorum uint256) constant returns(out bool)
func (_Voting *VotingSession) DidPass(hash [32]byte, quorum *big.Int) (bool, error) {
	return _Voting.Contract.DidPass(&_Voting.CallOpts, hash, quorum)
}

// DidPass is a free data retrieval call binding the contract method 0x8f354b79.
//
// Solidity: function didPass(hash bytes32, quorum uint256) constant returns(out bool)
func (_Voting *VotingCallerSession) DidPass(hash [32]byte, quorum *big.Int) (bool, error) {
	return _Voting.Contract.DidPass(&_Voting.CallOpts, hash, quorum)
}

// DidVote is a free data retrieval call binding the contract method 0x187a9c07.
//
// Solidity: function didVote(hash bytes32, member address) constant returns(out bool)
func (_Voting *VotingCaller) DidVote(opts *bind.CallOpts, hash [32]byte, member common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "didVote", hash, member)
	return *ret0, err
}

// DidVote is a free data retrieval call binding the contract method 0x187a9c07.
//
// Solidity: function didVote(hash bytes32, member address) constant returns(out bool)
func (_Voting *VotingSession) DidVote(hash [32]byte, member common.Address) (bool, error) {
	return _Voting.Contract.DidVote(&_Voting.CallOpts, hash, member)
}

// DidVote is a free data retrieval call binding the contract method 0x187a9c07.
//
// Solidity: function didVote(hash bytes32, member address) constant returns(out bool)
func (_Voting *VotingCallerSession) DidVote(hash [32]byte, member common.Address) (bool, error) {
	return _Voting.Contract.DidVote(&_Voting.CallOpts, hash, member)
}

// GetCandidate is a free data retrieval call binding the contract method 0xdfb6419f.
//
// Solidity: function getCandidate(hash bytes32) constant returns(out uint256, out uint256, out uint256)
func (_Voting *VotingCaller) GetCandidate(opts *bind.CallOpts, hash [32]byte) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Voting.contract.Call(opts, out, "getCandidate", hash)
	return *ret0, *ret1, *ret2, err
}

// GetCandidate is a free data retrieval call binding the contract method 0xdfb6419f.
//
// Solidity: function getCandidate(hash bytes32) constant returns(out uint256, out uint256, out uint256)
func (_Voting *VotingSession) GetCandidate(hash [32]byte) (*big.Int, *big.Int, *big.Int, error) {
	return _Voting.Contract.GetCandidate(&_Voting.CallOpts, hash)
}

// GetCandidate is a free data retrieval call binding the contract method 0xdfb6419f.
//
// Solidity: function getCandidate(hash bytes32) constant returns(out uint256, out uint256, out uint256)
func (_Voting *VotingCallerSession) GetCandidate(hash [32]byte) (*big.Int, *big.Int, *big.Int, error) {
	return _Voting.Contract.GetCandidate(&_Voting.CallOpts, hash)
}

// GetCandidateCount is a free data retrieval call binding the contract method 0x30a56347.
//
// Solidity: function getCandidateCount() constant returns(out int128)
func (_Voting *VotingCaller) GetCandidateCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "getCandidateCount")
	return *ret0, err
}

// GetCandidateCount is a free data retrieval call binding the contract method 0x30a56347.
//
// Solidity: function getCandidateCount() constant returns(out int128)
func (_Voting *VotingSession) GetCandidateCount() (*big.Int, error) {
	return _Voting.Contract.GetCandidateCount(&_Voting.CallOpts)
}

// GetCandidateCount is a free data retrieval call binding the contract method 0x30a56347.
//
// Solidity: function getCandidateCount() constant returns(out int128)
func (_Voting *VotingCallerSession) GetCandidateCount() (*big.Int, error) {
	return _Voting.Contract.GetCandidateCount(&_Voting.CallOpts)
}

// GetCandidateKey is a free data retrieval call binding the contract method 0x79b73924.
//
// Solidity: function getCandidateKey(index int128) constant returns(out bytes32)
func (_Voting *VotingCaller) GetCandidateKey(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "getCandidateKey", index)
	return *ret0, err
}

// GetCandidateKey is a free data retrieval call binding the contract method 0x79b73924.
//
// Solidity: function getCandidateKey(index int128) constant returns(out bytes32)
func (_Voting *VotingSession) GetCandidateKey(index *big.Int) ([32]byte, error) {
	return _Voting.Contract.GetCandidateKey(&_Voting.CallOpts, index)
}

// GetCandidateKey is a free data retrieval call binding the contract method 0x79b73924.
//
// Solidity: function getCandidateKey(index int128) constant returns(out bytes32)
func (_Voting *VotingCallerSession) GetCandidateKey(index *big.Int) ([32]byte, error) {
	return _Voting.Contract.GetCandidateKey(&_Voting.CallOpts, index)
}

// GetCouncilCount is a free data retrieval call binding the contract method 0x4627da50.
//
// Solidity: function getCouncilCount() constant returns(out int128)
func (_Voting *VotingCaller) GetCouncilCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "getCouncilCount")
	return *ret0, err
}

// GetCouncilCount is a free data retrieval call binding the contract method 0x4627da50.
//
// Solidity: function getCouncilCount() constant returns(out int128)
func (_Voting *VotingSession) GetCouncilCount() (*big.Int, error) {
	return _Voting.Contract.GetCouncilCount(&_Voting.CallOpts)
}

// GetCouncilCount is a free data retrieval call binding the contract method 0x4627da50.
//
// Solidity: function getCouncilCount() constant returns(out int128)
func (_Voting *VotingCallerSession) GetCouncilCount() (*big.Int, error) {
	return _Voting.Contract.GetCouncilCount(&_Voting.CallOpts)
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address, out address, out address)
func (_Voting *VotingCaller) GetPrivileged(opts *bind.CallOpts) (common.Address, common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Voting.contract.Call(opts, out, "getPrivileged")
	return *ret0, *ret1, *ret2, err
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address, out address, out address)
func (_Voting *VotingSession) GetPrivileged() (common.Address, common.Address, common.Address, error) {
	return _Voting.Contract.GetPrivileged(&_Voting.CallOpts)
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address, out address, out address)
func (_Voting *VotingCallerSession) GetPrivileged() (common.Address, common.Address, common.Address, error) {
	return _Voting.Contract.GetPrivileged(&_Voting.CallOpts)
}

// InCouncil is a free data retrieval call binding the contract method 0x62f9a55d.
//
// Solidity: function inCouncil(member address) constant returns(out bool)
func (_Voting *VotingCaller) InCouncil(opts *bind.CallOpts, member common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "inCouncil", member)
	return *ret0, err
}

// InCouncil is a free data retrieval call binding the contract method 0x62f9a55d.
//
// Solidity: function inCouncil(member address) constant returns(out bool)
func (_Voting *VotingSession) InCouncil(member common.Address) (bool, error) {
	return _Voting.Contract.InCouncil(&_Voting.CallOpts, member)
}

// InCouncil is a free data retrieval call binding the contract method 0x62f9a55d.
//
// Solidity: function inCouncil(member address) constant returns(out bool)
func (_Voting *VotingCallerSession) InCouncil(member common.Address) (bool, error) {
	return _Voting.Contract.InCouncil(&_Voting.CallOpts, member)
}

// IsCandidate is a free data retrieval call binding the contract method 0xb89694c6.
//
// Solidity: function isCandidate(hash bytes32) constant returns(out bool)
func (_Voting *VotingCaller) IsCandidate(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "isCandidate", hash)
	return *ret0, err
}

// IsCandidate is a free data retrieval call binding the contract method 0xb89694c6.
//
// Solidity: function isCandidate(hash bytes32) constant returns(out bool)
func (_Voting *VotingSession) IsCandidate(hash [32]byte) (bool, error) {
	return _Voting.Contract.IsCandidate(&_Voting.CallOpts, hash)
}

// IsCandidate is a free data retrieval call binding the contract method 0xb89694c6.
//
// Solidity: function isCandidate(hash bytes32) constant returns(out bool)
func (_Voting *VotingCallerSession) IsCandidate(hash [32]byte) (bool, error) {
	return _Voting.Contract.IsCandidate(&_Voting.CallOpts, hash)
}

// PollClosed is a free data retrieval call binding the contract method 0x327322c8.
//
// Solidity: function pollClosed(hash bytes32) constant returns(out bool)
func (_Voting *VotingCaller) PollClosed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "pollClosed", hash)
	return *ret0, err
}

// PollClosed is a free data retrieval call binding the contract method 0x327322c8.
//
// Solidity: function pollClosed(hash bytes32) constant returns(out bool)
func (_Voting *VotingSession) PollClosed(hash [32]byte) (bool, error) {
	return _Voting.Contract.PollClosed(&_Voting.CallOpts, hash)
}

// PollClosed is a free data retrieval call binding the contract method 0x327322c8.
//
// Solidity: function pollClosed(hash bytes32) constant returns(out bool)
func (_Voting *VotingCallerSession) PollClosed(hash [32]byte) (bool, error) {
	return _Voting.Contract.PollClosed(&_Voting.CallOpts, hash)
}

// WasCandidate is a free data retrieval call binding the contract method 0x80338142.
//
// Solidity: function wasCandidate(hash bytes32) constant returns(out bool)
func (_Voting *VotingCaller) WasCandidate(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "wasCandidate", hash)
	return *ret0, err
}

// WasCandidate is a free data retrieval call binding the contract method 0x80338142.
//
// Solidity: function wasCandidate(hash bytes32) constant returns(out bool)
func (_Voting *VotingSession) WasCandidate(hash [32]byte) (bool, error) {
	return _Voting.Contract.WasCandidate(&_Voting.CallOpts, hash)
}

// WasCandidate is a free data retrieval call binding the contract method 0x80338142.
//
// Solidity: function wasCandidate(hash bytes32) constant returns(out bool)
func (_Voting *VotingCallerSession) WasCandidate(hash [32]byte) (bool, error) {
	return _Voting.Contract.WasCandidate(&_Voting.CallOpts, hash)
}

// WillAddCandidate is a free data retrieval call binding the contract method 0x1f1b0299.
//
// Solidity: function willAddCandidate(hash bytes32) constant returns(out bool)
func (_Voting *VotingCaller) WillAddCandidate(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "willAddCandidate", hash)
	return *ret0, err
}

// WillAddCandidate is a free data retrieval call binding the contract method 0x1f1b0299.
//
// Solidity: function willAddCandidate(hash bytes32) constant returns(out bool)
func (_Voting *VotingSession) WillAddCandidate(hash [32]byte) (bool, error) {
	return _Voting.Contract.WillAddCandidate(&_Voting.CallOpts, hash)
}

// WillAddCandidate is a free data retrieval call binding the contract method 0x1f1b0299.
//
// Solidity: function willAddCandidate(hash bytes32) constant returns(out bool)
func (_Voting *VotingCallerSession) WillAddCandidate(hash [32]byte) (bool, error) {
	return _Voting.Contract.WillAddCandidate(&_Voting.CallOpts, hash)
}

// WillAddToCouncil is a free data retrieval call binding the contract method 0xc666b580.
//
// Solidity: function willAddToCouncil(addr address) constant returns(out bool)
func (_Voting *VotingCaller) WillAddToCouncil(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "willAddToCouncil", addr)
	return *ret0, err
}

// WillAddToCouncil is a free data retrieval call binding the contract method 0xc666b580.
//
// Solidity: function willAddToCouncil(addr address) constant returns(out bool)
func (_Voting *VotingSession) WillAddToCouncil(addr common.Address) (bool, error) {
	return _Voting.Contract.WillAddToCouncil(&_Voting.CallOpts, addr)
}

// WillAddToCouncil is a free data retrieval call binding the contract method 0xc666b580.
//
// Solidity: function willAddToCouncil(addr address) constant returns(out bool)
func (_Voting *VotingCallerSession) WillAddToCouncil(addr common.Address) (bool, error) {
	return _Voting.Contract.WillAddToCouncil(&_Voting.CallOpts, addr)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x7ff45d8b.
//
// Solidity: function addCandidate(hash bytes32, kind uint256, vote_by uint256) returns()
func (_Voting *VotingTransactor) AddCandidate(opts *bind.TransactOpts, hash [32]byte, kind *big.Int, vote_by *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "addCandidate", hash, kind, vote_by)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x7ff45d8b.
//
// Solidity: function addCandidate(hash bytes32, kind uint256, vote_by uint256) returns()
func (_Voting *VotingSession) AddCandidate(hash [32]byte, kind *big.Int, vote_by *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, hash, kind, vote_by)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x7ff45d8b.
//
// Solidity: function addCandidate(hash bytes32, kind uint256, vote_by uint256) returns()
func (_Voting *VotingTransactorSession) AddCandidate(hash [32]byte, kind *big.Int, vote_by *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, hash, kind, vote_by)
}

// AddToCouncil is a paid mutator transaction binding the contract method 0x471bb309.
//
// Solidity: function addToCouncil(member address) returns()
func (_Voting *VotingTransactor) AddToCouncil(opts *bind.TransactOpts, member common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "addToCouncil", member)
}

// AddToCouncil is a paid mutator transaction binding the contract method 0x471bb309.
//
// Solidity: function addToCouncil(member address) returns()
func (_Voting *VotingSession) AddToCouncil(member common.Address) (*types.Transaction, error) {
	return _Voting.Contract.AddToCouncil(&_Voting.TransactOpts, member)
}

// AddToCouncil is a paid mutator transaction binding the contract method 0x471bb309.
//
// Solidity: function addToCouncil(member address) returns()
func (_Voting *VotingTransactorSession) AddToCouncil(member common.Address) (*types.Transaction, error) {
	return _Voting.Contract.AddToCouncil(&_Voting.TransactOpts, member)
}

// RemoveCandidate is a paid mutator transaction binding the contract method 0x89bb617c.
//
// Solidity: function removeCandidate(hash bytes32) returns()
func (_Voting *VotingTransactor) RemoveCandidate(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "removeCandidate", hash)
}

// RemoveCandidate is a paid mutator transaction binding the contract method 0x89bb617c.
//
// Solidity: function removeCandidate(hash bytes32) returns()
func (_Voting *VotingSession) RemoveCandidate(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.RemoveCandidate(&_Voting.TransactOpts, hash)
}

// RemoveCandidate is a paid mutator transaction binding the contract method 0x89bb617c.
//
// Solidity: function removeCandidate(hash bytes32) returns()
func (_Voting *VotingTransactorSession) RemoveCandidate(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.RemoveCandidate(&_Voting.TransactOpts, hash)
}

// RemoveFromCouncil is a paid mutator transaction binding the contract method 0x335d8080.
//
// Solidity: function removeFromCouncil(member address) returns()
func (_Voting *VotingTransactor) RemoveFromCouncil(opts *bind.TransactOpts, member common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "removeFromCouncil", member)
}

// RemoveFromCouncil is a paid mutator transaction binding the contract method 0x335d8080.
//
// Solidity: function removeFromCouncil(member address) returns()
func (_Voting *VotingSession) RemoveFromCouncil(member common.Address) (*types.Transaction, error) {
	return _Voting.Contract.RemoveFromCouncil(&_Voting.TransactOpts, member)
}

// RemoveFromCouncil is a paid mutator transaction binding the contract method 0x335d8080.
//
// Solidity: function removeFromCouncil(member address) returns()
func (_Voting *VotingTransactorSession) RemoveFromCouncil(member common.Address) (*types.Transaction, error) {
	return _Voting.Contract.RemoveFromCouncil(&_Voting.TransactOpts, member)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0xb01e369f.
//
// Solidity: function setPrivileged(market address, parameterizer address) returns()
func (_Voting *VotingTransactor) SetPrivileged(opts *bind.TransactOpts, market common.Address, parameterizer common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "setPrivileged", market, parameterizer)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0xb01e369f.
//
// Solidity: function setPrivileged(market address, parameterizer address) returns()
func (_Voting *VotingSession) SetPrivileged(market common.Address, parameterizer common.Address) (*types.Transaction, error) {
	return _Voting.Contract.SetPrivileged(&_Voting.TransactOpts, market, parameterizer)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0xb01e369f.
//
// Solidity: function setPrivileged(market address, parameterizer address) returns()
func (_Voting *VotingTransactorSession) SetPrivileged(market common.Address, parameterizer common.Address) (*types.Transaction, error) {
	return _Voting.Contract.SetPrivileged(&_Voting.TransactOpts, market, parameterizer)
}

// Vote is a paid mutator transaction binding the contract method 0xa69beaba.
//
// Solidity: function vote(hash bytes32) returns()
func (_Voting *VotingTransactor) Vote(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "vote", hash)
}

// Vote is a paid mutator transaction binding the contract method 0xa69beaba.
//
// Solidity: function vote(hash bytes32) returns()
func (_Voting *VotingSession) Vote(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, hash)
}

// Vote is a paid mutator transaction binding the contract method 0xa69beaba.
//
// Solidity: function vote(hash bytes32) returns()
func (_Voting *VotingTransactorSession) Vote(hash [32]byte) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, hash)
}

// VotingCandidateAddedIterator is returned from FilterCandidateAdded and is used to iterate over the raw logs and unpacked data for CandidateAdded events raised by the Voting contract.
type VotingCandidateAddedIterator struct {
	Event *VotingCandidateAdded // Event containing the contract specifics and raw log

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
func (it *VotingCandidateAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingCandidateAdded)
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
		it.Event = new(VotingCandidateAdded)
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
func (it *VotingCandidateAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingCandidateAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingCandidateAdded represents a CandidateAdded event raised by the Voting contract.
type VotingCandidateAdded struct {
	Hash   [32]byte
	Kind   *big.Int
	VoteBy *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCandidateAdded is a free log retrieval operation binding the contract event 0x4566e4146679483436367293d719cef8eb5d21b6b774ec4cec37b82f224db161.
//
// Solidity: e CandidateAdded(hash indexed bytes32, kind indexed uint256, voteBy indexed uint256)
func (_Voting *VotingFilterer) FilterCandidateAdded(opts *bind.FilterOpts, hash [][32]byte, kind []*big.Int, voteBy []*big.Int) (*VotingCandidateAddedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var voteByRule []interface{}
	for _, voteByItem := range voteBy {
		voteByRule = append(voteByRule, voteByItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CandidateAdded", hashRule, kindRule, voteByRule)
	if err != nil {
		return nil, err
	}
	return &VotingCandidateAddedIterator{contract: _Voting.contract, event: "CandidateAdded", logs: logs, sub: sub}, nil
}

// WatchCandidateAdded is a free log subscription operation binding the contract event 0x4566e4146679483436367293d719cef8eb5d21b6b774ec4cec37b82f224db161.
//
// Solidity: e CandidateAdded(hash indexed bytes32, kind indexed uint256, voteBy indexed uint256)
func (_Voting *VotingFilterer) WatchCandidateAdded(opts *bind.WatchOpts, sink chan<- *VotingCandidateAdded, hash [][32]byte, kind []*big.Int, voteBy []*big.Int) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var voteByRule []interface{}
	for _, voteByItem := range voteBy {
		voteByRule = append(voteByRule, voteByItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CandidateAdded", hashRule, kindRule, voteByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingCandidateAdded)
				if err := _Voting.contract.UnpackLog(event, "CandidateAdded", log); err != nil {
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

// VotingCandidateRemovedIterator is returned from FilterCandidateRemoved and is used to iterate over the raw logs and unpacked data for CandidateRemoved events raised by the Voting contract.
type VotingCandidateRemovedIterator struct {
	Event *VotingCandidateRemoved // Event containing the contract specifics and raw log

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
func (it *VotingCandidateRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingCandidateRemoved)
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
		it.Event = new(VotingCandidateRemoved)
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
func (it *VotingCandidateRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingCandidateRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingCandidateRemoved represents a CandidateRemoved event raised by the Voting contract.
type VotingCandidateRemoved struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCandidateRemoved is a free log retrieval operation binding the contract event 0xb314642ce4aa3156c1090458cfccabff50f32bc85d110d7799c369e1a660bcf4.
//
// Solidity: e CandidateRemoved(hash indexed bytes32)
func (_Voting *VotingFilterer) FilterCandidateRemoved(opts *bind.FilterOpts, hash [][32]byte) (*VotingCandidateRemovedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CandidateRemoved", hashRule)
	if err != nil {
		return nil, err
	}
	return &VotingCandidateRemovedIterator{contract: _Voting.contract, event: "CandidateRemoved", logs: logs, sub: sub}, nil
}

// WatchCandidateRemoved is a free log subscription operation binding the contract event 0xb314642ce4aa3156c1090458cfccabff50f32bc85d110d7799c369e1a660bcf4.
//
// Solidity: e CandidateRemoved(hash indexed bytes32)
func (_Voting *VotingFilterer) WatchCandidateRemoved(opts *bind.WatchOpts, sink chan<- *VotingCandidateRemoved, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CandidateRemoved", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingCandidateRemoved)
				if err := _Voting.contract.UnpackLog(event, "CandidateRemoved", log); err != nil {
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

// VotingCouncilMemberAddedIterator is returned from FilterCouncilMemberAdded and is used to iterate over the raw logs and unpacked data for CouncilMemberAdded events raised by the Voting contract.
type VotingCouncilMemberAddedIterator struct {
	Event *VotingCouncilMemberAdded // Event containing the contract specifics and raw log

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
func (it *VotingCouncilMemberAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingCouncilMemberAdded)
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
		it.Event = new(VotingCouncilMemberAdded)
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
func (it *VotingCouncilMemberAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingCouncilMemberAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingCouncilMemberAdded represents a CouncilMemberAdded event raised by the Voting contract.
type VotingCouncilMemberAdded struct {
	Member common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCouncilMemberAdded is a free log retrieval operation binding the contract event 0xdcb46b4634af6065fbbeab63938d70a7cc803a67bc6a5bb9f28fb8d9a5c4ce14.
//
// Solidity: e CouncilMemberAdded(member indexed address)
func (_Voting *VotingFilterer) FilterCouncilMemberAdded(opts *bind.FilterOpts, member []common.Address) (*VotingCouncilMemberAddedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CouncilMemberAdded", memberRule)
	if err != nil {
		return nil, err
	}
	return &VotingCouncilMemberAddedIterator{contract: _Voting.contract, event: "CouncilMemberAdded", logs: logs, sub: sub}, nil
}

// WatchCouncilMemberAdded is a free log subscription operation binding the contract event 0xdcb46b4634af6065fbbeab63938d70a7cc803a67bc6a5bb9f28fb8d9a5c4ce14.
//
// Solidity: e CouncilMemberAdded(member indexed address)
func (_Voting *VotingFilterer) WatchCouncilMemberAdded(opts *bind.WatchOpts, sink chan<- *VotingCouncilMemberAdded, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CouncilMemberAdded", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingCouncilMemberAdded)
				if err := _Voting.contract.UnpackLog(event, "CouncilMemberAdded", log); err != nil {
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

// VotingCouncilMemberRemovedIterator is returned from FilterCouncilMemberRemoved and is used to iterate over the raw logs and unpacked data for CouncilMemberRemoved events raised by the Voting contract.
type VotingCouncilMemberRemovedIterator struct {
	Event *VotingCouncilMemberRemoved // Event containing the contract specifics and raw log

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
func (it *VotingCouncilMemberRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingCouncilMemberRemoved)
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
		it.Event = new(VotingCouncilMemberRemoved)
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
func (it *VotingCouncilMemberRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingCouncilMemberRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingCouncilMemberRemoved represents a CouncilMemberRemoved event raised by the Voting contract.
type VotingCouncilMemberRemoved struct {
	Member common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCouncilMemberRemoved is a free log retrieval operation binding the contract event 0xce6ed85c24e62fbd23486204879240944c169dfcde2952b0c15f12bf12029669.
//
// Solidity: e CouncilMemberRemoved(member indexed address)
func (_Voting *VotingFilterer) FilterCouncilMemberRemoved(opts *bind.FilterOpts, member []common.Address) (*VotingCouncilMemberRemovedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CouncilMemberRemoved", memberRule)
	if err != nil {
		return nil, err
	}
	return &VotingCouncilMemberRemovedIterator{contract: _Voting.contract, event: "CouncilMemberRemoved", logs: logs, sub: sub}, nil
}

// WatchCouncilMemberRemoved is a free log subscription operation binding the contract event 0xce6ed85c24e62fbd23486204879240944c169dfcde2952b0c15f12bf12029669.
//
// Solidity: e CouncilMemberRemoved(member indexed address)
func (_Voting *VotingFilterer) WatchCouncilMemberRemoved(opts *bind.WatchOpts, sink chan<- *VotingCouncilMemberRemoved, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CouncilMemberRemoved", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingCouncilMemberRemoved)
				if err := _Voting.contract.UnpackLog(event, "CouncilMemberRemoved", log); err != nil {
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

// VotingVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Voting contract.
type VotingVotedIterator struct {
	Event *VotingVoted // Event containing the contract specifics and raw log

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
func (it *VotingVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingVoted)
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
		it.Event = new(VotingVoted)
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
func (it *VotingVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingVoted represents a Voted event raised by the Voting contract.
type VotingVoted struct {
	Hash  [32]byte
	Voter common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x0b2f654b7e608ce51a82ce8157e79c350ed670605e8985266ad89fc85060e749.
//
// Solidity: e Voted(hash indexed bytes32, voter indexed address)
func (_Voting *VotingFilterer) FilterVoted(opts *bind.FilterOpts, hash [][32]byte, voter []common.Address) (*VotingVotedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "Voted", hashRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &VotingVotedIterator{contract: _Voting.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x0b2f654b7e608ce51a82ce8157e79c350ed670605e8985266ad89fc85060e749.
//
// Solidity: e Voted(hash indexed bytes32, voter indexed address)
func (_Voting *VotingFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *VotingVoted, hash [][32]byte, voter []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "Voted", hashRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingVoted)
				if err := _Voting.contract.UnpackLog(event, "Voted", log); err != nil {
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
