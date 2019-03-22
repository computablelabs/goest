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
const VotingABI = "[{\"name\":\"CandidateAdded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"kind\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"voteBy\",\"indexed\":false,\"unit\":\"sec\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CandidateRemoved\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CouncilMemberAdded\",\"inputs\":[{\"type\":\"address\",\"name\":\"member\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CouncilMemberRemoved\",\"inputs\":[{\"type\":\"address\",\"name\":\"member\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Voted\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"voter\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"getPrivileged\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1191},{\"name\":\"setPrivileged\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"parameterizer\"},{\"type\":\"address\",\"name\":\"datatrust\"},{\"type\":\"address\",\"name\":\"listing\"},{\"type\":\"address\",\"name\":\"investing\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":140819},{\"name\":\"has_privilege\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"sender\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1284},{\"name\":\"inCouncil\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"member\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":745},{\"name\":\"addToCouncil\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"member\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":39419},{\"name\":\"removeFromCouncil\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"member\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":24446},{\"name\":\"candidateIs\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"kind\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":883},{\"name\":\"isCandidate\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":913},{\"name\":\"getCandidate\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"sec\"},{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2623},{\"name\":\"getCandidateOwner\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":964},{\"name\":\"addCandidate\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"kind\"},{\"type\":\"address\",\"name\":\"owner\"},{\"type\":\"uint256\",\"name\":\"vote_by\",\"unit\":\"sec\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":111685},{\"name\":\"removeCandidate\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":20545731},{\"name\":\"didPass\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"quorum\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":5716},{\"name\":\"didVote\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"address\",\"name\":\"member\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":724940},{\"name\":\"pollClosed\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1609},{\"name\":\"vote\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"option\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":861595}]"

// VotingBin is the compiled bytecode used for deploying new contracts.
const VotingBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a052341561009e57600080fd5b33600655610fbe56600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263d4c1753960005114156100dd5734156100ac57600080fd5b60806101405261016060025481526003548160200152600454816040015260055481606001525061014051610160f3005b6339db04e6600051141561017257608060046101403734156100fe57600080fd5b600435602051811061010f57600080fd5b50602435602051811061012157600080fd5b50604435602051811061013357600080fd5b50606435602051811061014557600080fd5b50600654331461015457600080fd5b6101405160025561016051600355610180516004556101a051600555005b636a395d2360005114156101d2576020600461014037341561019357600080fd5b60043560205181106101a457600080fd5b50600354610140511460025461014051141760045461014051141760055461014051141760005260206000f3005b6362f9a55d600051141561022157602060046101403734156101f357600080fd5b600435602051811061020457600080fd5b5060016101405160e05260c052604060c0205460005260206000f3005b63471bb30960005114156102e3576020600461014037341561024257600080fd5b600435602051811061025357600080fd5b5060206101e06024636a395d2361016052336101805261017c6000305af161027a57600080fd5b6101e05161028757600080fd5b60016101405160e05260c052604060c02054156102a357600080fd5b600160016101405160e05260c052604060c02055610140517fdcb46b4634af6065fbbeab63938d70a7cc803a67bc6a5bb9f28fb8d9a5c4ce1460006000a2005b63335d808060005114156103a4576020600461014037341561030457600080fd5b600435602051811061031557600080fd5b5060206101e06024636a395d2361016052336101805261017c6000305af161033c57600080fd5b6101e05161034957600080fd5b60016101405160e05260c052604060c0205461036457600080fd5b600060016101405160e05260c052604060c02055610140517fce6ed85c24e62fbd23486204879240944c169dfcde2952b0c15f12bf1202966960006000a2005b63af61f76060005114156103ee57604060046101403734156103c557600080fd5b6101605160006101405160e05260c052604060c02060c052602060c020541460005260206000f3005b63b89694c6600051141561043a576020600461014037341561040f57600080fd5b6000600160006101405160e05260c052604060c02060c052602060c0200154141560005260206000f3005b63dfb6419f6000511415610513576020600461014037341561045b57600080fd5b60a06101605261018060006101405160e05260c052604060c02060c052602060c020548152600160006101405160e05260c052604060c02060c052602060c02001548160200152600260006101405160e05260c052604060c02060c052602060c02001548160400152600460006101405160e05260c052604060c02060c052602060c02001548160600152600560006101405160e05260c052604060c02060c052602060c020015481608001525061016051610180f3005b63eb3014fd600051141561055b576020600461014037341561053457600080fd5b600160006101405160e05260c052604060c02060c052602060c020015460005260206000f3005b63e8b969a460005114156106a1576080600461014037341561057c57600080fd5b604435602051811061058d57600080fd5b5060206102406024636a395d236101c052336101e0526101dc6000305af16105b457600080fd5b610240516105c157600080fd5b600160006101405160e05260c052604060c02060c052602060c0200154156105e857600080fd5b426101a051420110156105fa57600080fd5b6101a0514201610260526101605160006101405160e05260c052604060c02060c052602060c0205561018051600160006101405160e05260c052604060c02060c052602060c020015561026051600260006101405160e05260c052604060c02060c052602060c020015561026051610280526101805161016051610140517f51baeaa00f5969dc416d8e5299022ac3886d317d13685d1c6906ced9bac71b026020610280a4005b6389bb617c600051141561090657602060046101403734156106c257600080fd5b60206101e06024636a395d2361016052336101805261017c6000305af16106e857600080fd5b6101e0516106f557600080fd5b6000600160006101405160e05260c052604060c02060c052602060c0200154141561071f57600080fd5b600060006101405160e05260c052604060c02060c052602060c020556000600160006101405160e05260c052604060c02060c052602060c02001556000600260006101405160e05260c052604060c02060c052602060c0200155600460006101405160e05260c052604060c02060c052602060c0200154600560006101405160e05260c052604060c02060c052602060c0200154600460006101405160e05260c052604060c02060c052602060c02001540110156107dc57600080fd5b600560006101405160e05260c052604060c02060c052602060c0200154600460006101405160e05260c052604060c02060c052602060c02001540160405181111561082657600080fd5b6102005261022060006103e8818352015b6102005161022051141561084e5761089a56610889565b6000610220516103e8811061086257600080fd5b600360006101405160e05260c052604060c02060c052602060c0200160c052602060c02001555b5b8151600101808352811415610837575b50506000600460006101405160e05260c052604060c02060c052602060c02001556000600560006101405160e05260c052604060c02060c052602060c0200155610140517fb314642ce4aa3156c1090458cfccabff50f32bc85d110d7799c369e1a660bcf460006000a2005b638f354b796000511415610aef576040600461014037341561092757600080fd5b6000600160006101405160e05260c052604060c02060c052602060c0200154141561095157600080fd5b42600260006101405160e05260c052604060c02060c052602060c02001541061097957600080fd5b600460006101405160e05260c052604060c02060c052602060c0200154600560006101405160e05260c052604060c02060c052602060c0200154600460006101405160e05260c052604060c02060c052602060c02001540110156109dc57600080fd5b600560006101405160e05260c052604060c02060c052602060c0200154600460006101405160e05260c052604060c02060c052602060c02001540161018052610180511515610a3757610160511560005260206000f3610aed565b6101605161018051610a4857600080fd5b61018051600460006101405160e05260c052604060c02060c052602060c02001541515610a76576000610ae1565b6064600460006101405160e05260c052604060c02060c052602060c02001546064600460006101405160e05260c052604060c02060c052602060c0200154020414610ac057600080fd5b6064600460006101405160e05260c052604060c02060c052602060c0200154025b04101560005260206000f35b005b63187a9c076000511415610c6c5760406004610140373415610b1057600080fd5b6024356020518110610b2157600080fd5b50600061018052600460006101405160e05260c052604060c02060c052602060c0200154600560006101405160e05260c052604060c02060c052602060c0200154600460006101405160e05260c052604060c02060c052602060c0200154011015610b8b57600080fd5b600560006101405160e05260c052604060c02060c052602060c0200154600460006101405160e05260c052604060c02060c052602060c020015401604051811115610bd557600080fd5b6101a0526101c060006103e8818352015b6101a0516101c0511415610bfd57610c5c56610c4b565b610160516101c0516103e88110610c1357600080fd5b600360006101405160e05260c052604060c02060c052602060c0200160c052602060c02001541415610c4a57600161018052610c5c565b5b5b8151600101808352811415610be6575b50506101805160005260206000f3005b63327322c86000511415610ce05760206004610140373415610c8d57600080fd5b6000600160006101405160e05260c052604060c02060c052602060c02001541415610cb757600080fd5b42600260006101405160e05260c052604060c02060c052602060c02001541060005260206000f3005b639ef1204c6000511415610f115760406004610140373415610d0157600080fd5b60013360e05260c052604060c02054610d1957600080fd5b6000600160006101405160e05260c052604060c02060c052602060c02001541415610d4357600080fd5b42600260006101405160e05260c052604060c02060c052602060c020015411610d6b57600080fd5b6020610220604463187a9c0761018052610140516101a052336101c05261019c6000305af1610d9957600080fd5b6102205115610da757600080fd5b600460006101405160e05260c052604060c02060c052602060c0200154600560006101405160e05260c052604060c02060c052602060c0200154600460006101405160e05260c052604060c02060c052602060c0200154011015610e0a57600080fd5b600560006101405160e05260c052604060c02060c052602060c0200154600460006101405160e05260c052604060c02060c052602060c020015401604051811115610e5457600080fd5b6102405233610240516103e88110610e6b57600080fd5b600360006101405160e05260c052604060c02060c052602060c0200160c052602060c02001556001610160511415610ed857600460006101405160e05260c052604060c02060c052602060c02001805460018254011015610ecb57600080fd5b6001815401815550610f0f565b600560006101405160e05260c052604060c02060c052602060c02001805460018254011015610f0657600080fd5b60018154018155505b005b60006000fd5b6100a7610fbe036100a76000396100a7610fbe036000f3`

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
// Solidity: function getCandidate(hash bytes32) constant returns(out uint256, out address, out uint256, out uint256, out uint256)
func (_Voting *VotingCaller) GetCandidate(opts *bind.CallOpts, hash [32]byte) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Voting.contract.Call(opts, out, "getCandidate", hash)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetCandidate is a free data retrieval call binding the contract method 0xdfb6419f.
//
// Solidity: function getCandidate(hash bytes32) constant returns(out uint256, out address, out uint256, out uint256, out uint256)
func (_Voting *VotingSession) GetCandidate(hash [32]byte) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _Voting.Contract.GetCandidate(&_Voting.CallOpts, hash)
}

// GetCandidate is a free data retrieval call binding the contract method 0xdfb6419f.
//
// Solidity: function getCandidate(hash bytes32) constant returns(out uint256, out address, out uint256, out uint256, out uint256)
func (_Voting *VotingCallerSession) GetCandidate(hash [32]byte) (*big.Int, common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _Voting.Contract.GetCandidate(&_Voting.CallOpts, hash)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xeb3014fd.
//
// Solidity: function getCandidateOwner(hash bytes32) constant returns(out address)
func (_Voting *VotingCaller) GetCandidateOwner(opts *bind.CallOpts, hash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "getCandidateOwner", hash)
	return *ret0, err
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xeb3014fd.
//
// Solidity: function getCandidateOwner(hash bytes32) constant returns(out address)
func (_Voting *VotingSession) GetCandidateOwner(hash [32]byte) (common.Address, error) {
	return _Voting.Contract.GetCandidateOwner(&_Voting.CallOpts, hash)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xeb3014fd.
//
// Solidity: function getCandidateOwner(hash bytes32) constant returns(out address)
func (_Voting *VotingCallerSession) GetCandidateOwner(hash [32]byte) (common.Address, error) {
	return _Voting.Contract.GetCandidateOwner(&_Voting.CallOpts, hash)
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address, out address, out address, out address)
func (_Voting *VotingCaller) GetPrivileged(opts *bind.CallOpts) (common.Address, common.Address, common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(common.Address)
		ret3 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Voting.contract.Call(opts, out, "getPrivileged")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address, out address, out address, out address)
func (_Voting *VotingSession) GetPrivileged() (common.Address, common.Address, common.Address, common.Address, error) {
	return _Voting.Contract.GetPrivileged(&_Voting.CallOpts)
}

// GetPrivileged is a free data retrieval call binding the contract method 0xd4c17539.
//
// Solidity: function getPrivileged() constant returns(out address, out address, out address, out address)
func (_Voting *VotingCallerSession) GetPrivileged() (common.Address, common.Address, common.Address, common.Address, error) {
	return _Voting.Contract.GetPrivileged(&_Voting.CallOpts)
}

// HasPrivilege is a free data retrieval call binding the contract method 0x6a395d23.
//
// Solidity: function has_privilege(sender address) constant returns(out bool)
func (_Voting *VotingCaller) HasPrivilege(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "has_privilege", sender)
	return *ret0, err
}

// HasPrivilege is a free data retrieval call binding the contract method 0x6a395d23.
//
// Solidity: function has_privilege(sender address) constant returns(out bool)
func (_Voting *VotingSession) HasPrivilege(sender common.Address) (bool, error) {
	return _Voting.Contract.HasPrivilege(&_Voting.CallOpts, sender)
}

// HasPrivilege is a free data retrieval call binding the contract method 0x6a395d23.
//
// Solidity: function has_privilege(sender address) constant returns(out bool)
func (_Voting *VotingCallerSession) HasPrivilege(sender common.Address) (bool, error) {
	return _Voting.Contract.HasPrivilege(&_Voting.CallOpts, sender)
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

// AddCandidate is a paid mutator transaction binding the contract method 0xe8b969a4.
//
// Solidity: function addCandidate(hash bytes32, kind uint256, owner address, vote_by uint256) returns()
func (_Voting *VotingTransactor) AddCandidate(opts *bind.TransactOpts, hash [32]byte, kind *big.Int, owner common.Address, vote_by *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "addCandidate", hash, kind, owner, vote_by)
}

// AddCandidate is a paid mutator transaction binding the contract method 0xe8b969a4.
//
// Solidity: function addCandidate(hash bytes32, kind uint256, owner address, vote_by uint256) returns()
func (_Voting *VotingSession) AddCandidate(hash [32]byte, kind *big.Int, owner common.Address, vote_by *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, hash, kind, owner, vote_by)
}

// AddCandidate is a paid mutator transaction binding the contract method 0xe8b969a4.
//
// Solidity: function addCandidate(hash bytes32, kind uint256, owner address, vote_by uint256) returns()
func (_Voting *VotingTransactorSession) AddCandidate(hash [32]byte, kind *big.Int, owner common.Address, vote_by *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.AddCandidate(&_Voting.TransactOpts, hash, kind, owner, vote_by)
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

// SetPrivileged is a paid mutator transaction binding the contract method 0x39db04e6.
//
// Solidity: function setPrivileged(parameterizer address, datatrust address, listing address, investing address) returns()
func (_Voting *VotingTransactor) SetPrivileged(opts *bind.TransactOpts, parameterizer common.Address, datatrust common.Address, listing common.Address, investing common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "setPrivileged", parameterizer, datatrust, listing, investing)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0x39db04e6.
//
// Solidity: function setPrivileged(parameterizer address, datatrust address, listing address, investing address) returns()
func (_Voting *VotingSession) SetPrivileged(parameterizer common.Address, datatrust common.Address, listing common.Address, investing common.Address) (*types.Transaction, error) {
	return _Voting.Contract.SetPrivileged(&_Voting.TransactOpts, parameterizer, datatrust, listing, investing)
}

// SetPrivileged is a paid mutator transaction binding the contract method 0x39db04e6.
//
// Solidity: function setPrivileged(parameterizer address, datatrust address, listing address, investing address) returns()
func (_Voting *VotingTransactorSession) SetPrivileged(parameterizer common.Address, datatrust common.Address, listing common.Address, investing common.Address) (*types.Transaction, error) {
	return _Voting.Contract.SetPrivileged(&_Voting.TransactOpts, parameterizer, datatrust, listing, investing)
}

// Vote is a paid mutator transaction binding the contract method 0x9ef1204c.
//
// Solidity: function vote(hash bytes32, option uint256) returns()
func (_Voting *VotingTransactor) Vote(opts *bind.TransactOpts, hash [32]byte, option *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "vote", hash, option)
}

// Vote is a paid mutator transaction binding the contract method 0x9ef1204c.
//
// Solidity: function vote(hash bytes32, option uint256) returns()
func (_Voting *VotingSession) Vote(hash [32]byte, option *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, hash, option)
}

// Vote is a paid mutator transaction binding the contract method 0x9ef1204c.
//
// Solidity: function vote(hash bytes32, option uint256) returns()
func (_Voting *VotingTransactorSession) Vote(hash [32]byte, option *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, hash, option)
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
	Owner  common.Address
	VoteBy *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCandidateAdded is a free log retrieval operation binding the contract event 0x51baeaa00f5969dc416d8e5299022ac3886d317d13685d1c6906ced9bac71b02.
//
// Solidity: e CandidateAdded(hash indexed bytes32, kind indexed uint256, owner indexed address, voteBy uint256)
func (_Voting *VotingFilterer) FilterCandidateAdded(opts *bind.FilterOpts, hash [][32]byte, kind []*big.Int, owner []common.Address) (*VotingCandidateAddedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "CandidateAdded", hashRule, kindRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VotingCandidateAddedIterator{contract: _Voting.contract, event: "CandidateAdded", logs: logs, sub: sub}, nil
}

// WatchCandidateAdded is a free log subscription operation binding the contract event 0x51baeaa00f5969dc416d8e5299022ac3886d317d13685d1c6906ced9bac71b02.
//
// Solidity: e CandidateAdded(hash indexed bytes32, kind indexed uint256, owner indexed address, voteBy uint256)
func (_Voting *VotingFilterer) WatchCandidateAdded(opts *bind.WatchOpts, sink chan<- *VotingCandidateAdded, hash [][32]byte, kind []*big.Int, owner []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "CandidateAdded", hashRule, kindRule, ownerRule)
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
