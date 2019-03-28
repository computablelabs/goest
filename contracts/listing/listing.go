// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package listing

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

// ListingABI is the input ABI used to generate the binding from.
const ListingABI = "[{\"name\":\"ApplicationFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"applicant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Applied\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"applicant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Challenged\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"challenger\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ChallengeFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"challenger\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ChallengeSucceeded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"challenger\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Listed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"reward\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingConverted\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingDeposit\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"deposited\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingRemoved\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingWithdraw\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"withdrawn\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"market_token_addr\"},{\"type\":\"address\",\"name\":\"voting_addr\"},{\"type\":\"address\",\"name\":\"p11r_addr\"},{\"type\":\"address\",\"name\":\"data_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"isListed\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":697},{\"name\":\"depositToListing\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":41264},{\"name\":\"withdrawFromListing\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":41073},{\"name\":\"list\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"listing\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":10247},{\"name\":\"getHash\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"string\",\"name\":\"str\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":633},{\"name\":\"getListing\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1717},{\"name\":\"convertListing\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":82621},{\"name\":\"resolveApplication\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":90163},{\"name\":\"challenge\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":10209},{\"name\":\"resolveChallenge\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":125992},{\"name\":\"exit\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":115221}]"

// ListingBin is the compiled bytecode used for deploying new contracts.
const ListingBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05260806115d86101403934156100a757600080fd5b60206115d860c03960c05160205181106100c057600080fd5b50602060206115d80160c03960c05160205181106100dd57600080fd5b50602060406115d80160c03960c05160205181106100fa57600080fd5b50602060606115d80160c03960c051602051811061011757600080fd5b506101405160015561016051600255610180516003556101a0516004556115c056600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263ecefbdc660005114156100dc57602060046101403734156100b457600080fd5b600060006101405160e05260c052604060c02060c052602060c02054141560005260206000f3005b63341ff08960005114156101ec57604060046101403734156100fd57600080fd5b600060006101405160e05260c052604060c02060c052602060c02054141561012457600080fd5b6001543b61013157600080fd5b60015430141561014057600080fd5b602061024060646323b872dd61018052336101a052306101c052610160516101e05261019c60006001545af161017557600080fd5b60005061024050600160006101405160e05260c052604060c02060c052602060c0200180546101605182540110156101ac57600080fd5b61016051815401815550610160516102605233610140517f6f27446a1e23c7e7dd81a6a77f81438be78145de331a5806f2c44e68738bdcef6020610260a3005b63a7e7867560005114156102f2576040600461014037341561020d57600080fd5b3360006101405160e05260c052604060c02060c052602060c020541461023257600080fd5b600160006101405160e05260c052604060c02060c052602060c02001610160518154101561025f57600080fd5b610160518154038155506001543b61027657600080fd5b60015430141561028557600080fd5b6020610220604463a9059cbb61018052336101a052610160516101c05261019c60006001545af16102b557600080fd5b60005061022050610160516102405233610140517f87c48f853a68e9374cf8389dab199f46adaffa48468fe9a4e2dc2edf75286db76020610240a3005b63fe42bf1a60005114156104cb576020600461014037341561031357600080fd5b606060043560040161016037604060043560040135111561033357600080fd5b6101608051602082012090506101e0526002543b61035057600080fd5b60025430141561035f57600080fd5b6020610280602463b89694c6610200526101e0516102205261021c6002545afa61038857600080fd5b600050610280511561039957600080fd5b60006101e05160e05260c052604060c02060c052602060c02054156103bd57600080fd5b6002543b6103ca57600080fd5b6002543014156103d957600080fd5b6000600060a463bb12c49e6103a0526101e0516103c05260016103e05233610400526003543b61040857600080fd5b60035430141561041757600080fd5b6020610300600463fc0e3d906102a0526102bc6003545afa61043857600080fd5b60005061030051610420526003543b61045057600080fd5b60035430141561045f57600080fd5b60206103806004632d0d2bc66103205261033c6003545afa61048057600080fd5b60005061038051610440526103bc60006002545af161049e57600080fd5b336101e0517f26f7438941d95dc8495c99ef33fac91b2b04f98c73290027562cd3e612a3c85f60006000a3005b635b6beeb9600051141561052257602060046101403734156104ec57600080fd5b606060043560040161016037604060043560040135111561050c57600080fd5b61016080516020820120905060005260206000f3005b63175c0d1660005114156105b7576020600461014037341561054357600080fd5b60606101605261018060006101405160e05260c052604060c02060c052602060c020548152600160006101405160e05260c052604060c02060c052602060c02001548160200152600260006101405160e05260c052604060c02060c052602060c020015481604001525061016051610180f3005b63d748484d600051141561079d57602060046101403734156105d857600080fd5b3360006101405160e05260c052604060c02060c052602060c02054146105fd57600080fd5b600160006101405160e05260c052604060c02060c052602060c0200154600260006101405160e05260c052604060c02060c052602060c0200154600160006101405160e05260c052604060c02060c052602060c020015401101561066057600080fd5b600260006101405160e05260c052604060c02060c052602060c0200154600160006101405160e05260c052604060c02060c052602060c020015401610160526000610160511115610755576000600160006101405160e05260c052604060c02060c052602060c02001556000600260006101405160e05260c052604060c02060c052602060c02001556001543b6106f657600080fd5b60015430141561070557600080fd5b6020610220604463a9059cbb6101805260006101405160e05260c052604060c02060c052602060c020546101a052610160516101c05261019c60006001545af161074e57600080fd5b6000506102205b503060006101405160e05260c052604060c02060c052602060c02055610140517ffd6a9ca1840ac9f8b7c592783e5415fdc282281630286c93b07b6d40f3c0ba5760006000a2005b6000156109e4575b6101605261014052600060006101405160e05260c052604060c02060c052602060c0205414156107d457600080fd5b6000600160006101405160e05260c052604060c02060c052602060c0200154111561089e576001543b61080657600080fd5b60015430141561081557600080fd5b6020610220604463a9059cbb6101805260006101405160e05260c052604060c02060c052602060c020546101a052600160006101405160e05260c052604060c02060c052602060c02001546101c05261019c60006001545af161087757600080fd5b600050610220506000600160006101405160e05260c052604060c02060c052602060c02001555b6000600260006101405160e05260c052604060c02060c052602060c02001541115610942576001543b6108d057600080fd5b6001543014156108df57600080fd5b6000600060246342966c6861024052600260006101405160e05260c052604060c02060c052602060c02001546102605261025c60006001545af161092257600080fd5b6000600260006101405160e05260c052604060c02060c052602060c02001555b60006101405160e05260c052604060c02060c052602060c020600081556000600182015560006002820155506004543b61097b57600080fd5b60045430141561098a57600080fd5b60006000602463bd82badc6102c052610140516102e0526102dc60006004545af16109b457600080fd5b610140517f50425cae216bd151d26c8e8bb9779cc899c99d72f78081f2ceec9a99f001ff7960006000a261016051565b6318e28e1c6000511415610df55760206004610140373415610a0557600080fd5b6002543b610a1257600080fd5b600254301415610a2157600080fd5b6020610200604463af61f76061016052610140516101805260016101a05261017c6002545afa610a5057600080fd5b60005061020051610a6057600080fd5b6002543b610a6d57600080fd5b600254301415610a7c57600080fd5b60206102a0602463327322c861022052610140516102405261023c6002545afa610aa557600080fd5b6000506102a051610ab557600080fd5b6004543b610ac257600080fd5b600454301415610ad157600080fd5b60206103606024637a639f6e6102e05261014051610300526102fc6004545afa610afa57600080fd5b600050610360516102c0526002543b610b1257600080fd5b600254301415610b2157600080fd5b6020610420602463eb3014fd6103a052610140516103c0526103bc6002545afa610b4a57600080fd5b600050610420516103805260006102c051141515610d7e576002543b610b6f57600080fd5b600254301415610b7e57600080fd5b60206105e06044638f354b796105405261014051610560526003543b610ba357600080fd5b600354301415610bb257600080fd5b6020610520600463c26c12eb6104c0526104dc6003545afa610bd357600080fd5b600050610520516105805261055c6002545afa610bef57600080fd5b6000506105e05115610d04576103805160006101405160e05260c052604060c02060c052602060c020556003543b610c2657600080fd5b600354301415610c3557600080fd5b6020610680600463dc2b28536106205261063c6003545afa610c5657600080fd5b60005061068051610600526001543b610c6e57600080fd5b600154301415610c7d57600080fd5b60006000602463a0712d686106a052610600516106c0526106bc60006001545af1610ca757600080fd5b61060051600260006101405160e05260c052604060c02060c052602060c0200155610600516107205261038051610140517f7f1a2c3e2883554425dc0f9f24dcfcdf54213b186c550080373dcc65813aa8d06020610720a3610d79565b6004543b610d1157600080fd5b600454301415610d2057600080fd5b60006000602463bd82badc61044052610140516104605261045c60006004545af1610d4a57600080fd5b61038051610140517f163b66974440f1696d484dfc777282273c42ff2fa247d72dbdbb3510d92cbda660006000a35b610dad565b61038051610140517f163b66974440f1696d484dfc777282273c42ff2fa247d72dbdbb3510d92cbda660006000a35b6002543b610dba57600080fd5b600254301415610dc957600080fd5b6000600060246389bb617c61074052610140516107605261075c60006002545af1610df357600080fd5b005b63cffd46dc6000511415610fa15760206004610140373415610e1657600080fd5b600060006101405160e05260c052604060c02060c052602060c020541415610e3d57600080fd5b6002543b610e4a57600080fd5b600254301415610e5957600080fd5b60206101e0602463b89694c661016052610140516101805261017c6002545afa610e8257600080fd5b6000506101e05115610e9357600080fd5b6002543b610ea057600080fd5b600254301415610eaf57600080fd5b6000600060a463bb12c49e61030052610140516103205260026103405233610360526003543b610ede57600080fd5b600354301415610eed57600080fd5b6020610260600463fc0e3d906102005261021c6003545afa610f0e57600080fd5b60005061026051610380526003543b610f2657600080fd5b600354301415610f3557600080fd5b60206102e06004632d0d2bc66102805261029c6003545afa610f5657600080fd5b6000506102e0516103a05261031c60006002545af1610f7457600080fd5b33610140517fe9479421670c3425a1497ce47a53af8bd96ce5bd0741e96221ba0acace3f7d4760006000a3005b63d32c943a60005114156113925760206004610140373415610fc257600080fd5b6002543b610fcf57600080fd5b600254301415610fde57600080fd5b6020610200604463af61f76061016052610140516101805260026101a05261017c6002545afa61100d57600080fd5b6000506102005161101d57600080fd5b6002543b61102a57600080fd5b60025430141561103957600080fd5b60206102a0602463327322c861022052610140516102405261023c6002545afa61106257600080fd5b6000506102a05161107257600080fd5b6002543b61107f57600080fd5b60025430141561108e57600080fd5b6020610360602463eb3014fd6102e05261014051610300526102fc6002545afa6110b757600080fd5b600050610360516102c0526002543b6110cf57600080fd5b6002543014156110de57600080fd5b60206105406044638f354b796104a052610140516104c0526003543b61110357600080fd5b60035430141561111257600080fd5b6020610480600463c26c12eb6104205261043c6003545afa61113357600080fd5b600050610480516104e0526104bc6002545afa61114f57600080fd5b60005061054051156112b7576101405161016051610180516101a0516101c0516101e05161020051610220516102405161026051610280516102a0516102c0516102e05161030051610320516103405161036051610380516103a0516103c0516103e05161040051610420516104405161046051610480516104a0516104c0516104e05161050051610520516105405163550ee13361056052610140516105805261058051600658016107a5565b6105405261052052610500526104e0526104c0526104a05261048052610460526104405261042052610400526103e0526103c0526103a05261038052610360526103405261032052610300526102e0526102c0526102a05261028052610260526102405261022052610200526101e0526101c0526101a0526101805261016052610140526000506102c051610140517fb7e10908a54924c8e096789495f5a4958fed82b49d856d22b208a23b398306bb60006000a361134a565b6002543b6112c457600080fd5b6002543014156112d357600080fd5b60006000604463b6b6920661038052610140516103a05260006101405160e05260c052604060c02060c052602060c020546103c05261039c60006002545af161131b57600080fd5b6102c051610140517fa0e6c0bd204f59362c8b2ebdf77c63242bf779d8ac30c347933db099abcc037060006000a35b6002543b61135757600080fd5b60025430141561136657600080fd5b6000600060246389bb617c6105e05261014051610600526105fc60006002545af161139057600080fd5b005b630ca36263600051141561148157602060046101403734156113b357600080fd5b3360006101405160e05260c052604060c02060c052602060c02054146113d857600080fd5b6002543b6113e557600080fd5b6002543014156113f457600080fd5b60206101e0602463b89694c661016052610140516101805261017c6002545afa61141d57600080fd5b6000506101e0511561142e57600080fd5b6101405161016051610180516101a0516101c0516101e05163550ee13361020052610140516102205261022051600658016107a5565b6101e0526101c0526101a052610180526101605261014052600050005b60006000fd5b6101396115c0036101396000396101396115c0036000f3`

// DeployListing deploys a new Ethereum contract, binding an instance of Listing to it.
func DeployListing(auth *bind.TransactOpts, backend bind.ContractBackend, market_token_addr common.Address, voting_addr common.Address, p11r_addr common.Address, data_addr common.Address) (common.Address, *types.Transaction, *Listing, error) {
	parsed, err := abi.JSON(strings.NewReader(ListingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ListingBin), backend, market_token_addr, voting_addr, p11r_addr, data_addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Listing{ListingCaller: ListingCaller{contract: contract}, ListingTransactor: ListingTransactor{contract: contract}, ListingFilterer: ListingFilterer{contract: contract}}, nil
}

// Listing is an auto generated Go binding around an Ethereum contract.
type Listing struct {
	ListingCaller     // Read-only binding to the contract
	ListingTransactor // Write-only binding to the contract
	ListingFilterer   // Log filterer for contract events
}

// ListingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ListingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ListingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ListingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ListingSession struct {
	Contract     *Listing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ListingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ListingCallerSession struct {
	Contract *ListingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ListingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ListingTransactorSession struct {
	Contract     *ListingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ListingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ListingRaw struct {
	Contract *Listing // Generic contract binding to access the raw methods on
}

// ListingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ListingCallerRaw struct {
	Contract *ListingCaller // Generic read-only contract binding to access the raw methods on
}

// ListingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ListingTransactorRaw struct {
	Contract *ListingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewListing creates a new instance of Listing, bound to a specific deployed contract.
func NewListing(address common.Address, backend bind.ContractBackend) (*Listing, error) {
	contract, err := bindListing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Listing{ListingCaller: ListingCaller{contract: contract}, ListingTransactor: ListingTransactor{contract: contract}, ListingFilterer: ListingFilterer{contract: contract}}, nil
}

// NewListingCaller creates a new read-only instance of Listing, bound to a specific deployed contract.
func NewListingCaller(address common.Address, caller bind.ContractCaller) (*ListingCaller, error) {
	contract, err := bindListing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ListingCaller{contract: contract}, nil
}

// NewListingTransactor creates a new write-only instance of Listing, bound to a specific deployed contract.
func NewListingTransactor(address common.Address, transactor bind.ContractTransactor) (*ListingTransactor, error) {
	contract, err := bindListing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ListingTransactor{contract: contract}, nil
}

// NewListingFilterer creates a new log filterer instance of Listing, bound to a specific deployed contract.
func NewListingFilterer(address common.Address, filterer bind.ContractFilterer) (*ListingFilterer, error) {
	contract, err := bindListing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ListingFilterer{contract: contract}, nil
}

// bindListing binds a generic wrapper to an already deployed contract.
func bindListing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ListingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Listing *ListingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Listing.Contract.ListingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Listing *ListingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Listing.Contract.ListingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Listing *ListingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Listing.Contract.ListingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Listing *ListingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Listing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Listing *ListingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Listing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Listing *ListingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Listing.Contract.contract.Transact(opts, method, params...)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(str string) constant returns(out bytes32)
func (_Listing *ListingCaller) GetHash(opts *bind.CallOpts, str string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Listing.contract.Call(opts, out, "getHash", str)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(str string) constant returns(out bytes32)
func (_Listing *ListingSession) GetHash(str string) ([32]byte, error) {
	return _Listing.Contract.GetHash(&_Listing.CallOpts, str)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(str string) constant returns(out bytes32)
func (_Listing *ListingCallerSession) GetHash(str string) ([32]byte, error) {
	return _Listing.Contract.GetHash(&_Listing.CallOpts, str)
}

// GetListing is a free data retrieval call binding the contract method 0x175c0d16.
//
// Solidity: function getListing(hash bytes32) constant returns(out address, out uint256, out uint256)
func (_Listing *ListingCaller) GetListing(opts *bind.CallOpts, hash [32]byte) (common.Address, *big.Int, *big.Int, error) {
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
	err := _Listing.contract.Call(opts, out, "getListing", hash)
	return *ret0, *ret1, *ret2, err
}

// GetListing is a free data retrieval call binding the contract method 0x175c0d16.
//
// Solidity: function getListing(hash bytes32) constant returns(out address, out uint256, out uint256)
func (_Listing *ListingSession) GetListing(hash [32]byte) (common.Address, *big.Int, *big.Int, error) {
	return _Listing.Contract.GetListing(&_Listing.CallOpts, hash)
}

// GetListing is a free data retrieval call binding the contract method 0x175c0d16.
//
// Solidity: function getListing(hash bytes32) constant returns(out address, out uint256, out uint256)
func (_Listing *ListingCallerSession) GetListing(hash [32]byte) (common.Address, *big.Int, *big.Int, error) {
	return _Listing.Contract.GetListing(&_Listing.CallOpts, hash)
}

// IsListed is a free data retrieval call binding the contract method 0xecefbdc6.
//
// Solidity: function isListed(hash bytes32) constant returns(out bool)
func (_Listing *ListingCaller) IsListed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Listing.contract.Call(opts, out, "isListed", hash)
	return *ret0, err
}

// IsListed is a free data retrieval call binding the contract method 0xecefbdc6.
//
// Solidity: function isListed(hash bytes32) constant returns(out bool)
func (_Listing *ListingSession) IsListed(hash [32]byte) (bool, error) {
	return _Listing.Contract.IsListed(&_Listing.CallOpts, hash)
}

// IsListed is a free data retrieval call binding the contract method 0xecefbdc6.
//
// Solidity: function isListed(hash bytes32) constant returns(out bool)
func (_Listing *ListingCallerSession) IsListed(hash [32]byte) (bool, error) {
	return _Listing.Contract.IsListed(&_Listing.CallOpts, hash)
}

// Challenge is a paid mutator transaction binding the contract method 0xcffd46dc.
//
// Solidity: function challenge(hash bytes32) returns()
func (_Listing *ListingTransactor) Challenge(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Listing.contract.Transact(opts, "challenge", hash)
}

// Challenge is a paid mutator transaction binding the contract method 0xcffd46dc.
//
// Solidity: function challenge(hash bytes32) returns()
func (_Listing *ListingSession) Challenge(hash [32]byte) (*types.Transaction, error) {
	return _Listing.Contract.Challenge(&_Listing.TransactOpts, hash)
}

// Challenge is a paid mutator transaction binding the contract method 0xcffd46dc.
//
// Solidity: function challenge(hash bytes32) returns()
func (_Listing *ListingTransactorSession) Challenge(hash [32]byte) (*types.Transaction, error) {
	return _Listing.Contract.Challenge(&_Listing.TransactOpts, hash)
}

// ConvertListing is a paid mutator transaction binding the contract method 0xd748484d.
//
// Solidity: function convertListing(hash bytes32) returns()
func (_Listing *ListingTransactor) ConvertListing(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Listing.contract.Transact(opts, "convertListing", hash)
}

// ConvertListing is a paid mutator transaction binding the contract method 0xd748484d.
//
// Solidity: function convertListing(hash bytes32) returns()
func (_Listing *ListingSession) ConvertListing(hash [32]byte) (*types.Transaction, error) {
	return _Listing.Contract.ConvertListing(&_Listing.TransactOpts, hash)
}

// ConvertListing is a paid mutator transaction binding the contract method 0xd748484d.
//
// Solidity: function convertListing(hash bytes32) returns()
func (_Listing *ListingTransactorSession) ConvertListing(hash [32]byte) (*types.Transaction, error) {
	return _Listing.Contract.ConvertListing(&_Listing.TransactOpts, hash)
}

// DepositToListing is a paid mutator transaction binding the contract method 0x341ff089.
//
// Solidity: function depositToListing(hash bytes32, amount uint256) returns()
func (_Listing *ListingTransactor) DepositToListing(opts *bind.TransactOpts, hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Listing.contract.Transact(opts, "depositToListing", hash, amount)
}

// DepositToListing is a paid mutator transaction binding the contract method 0x341ff089.
//
// Solidity: function depositToListing(hash bytes32, amount uint256) returns()
func (_Listing *ListingSession) DepositToListing(hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Listing.Contract.DepositToListing(&_Listing.TransactOpts, hash, amount)
}

// DepositToListing is a paid mutator transaction binding the contract method 0x341ff089.
//
// Solidity: function depositToListing(hash bytes32, amount uint256) returns()
func (_Listing *ListingTransactorSession) DepositToListing(hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Listing.Contract.DepositToListing(&_Listing.TransactOpts, hash, amount)
}

// Exit is a paid mutator transaction binding the contract method 0x0ca36263.
//
// Solidity: function exit(hash bytes32) returns()
func (_Listing *ListingTransactor) Exit(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Listing.contract.Transact(opts, "exit", hash)
}

// Exit is a paid mutator transaction binding the contract method 0x0ca36263.
//
// Solidity: function exit(hash bytes32) returns()
func (_Listing *ListingSession) Exit(hash [32]byte) (*types.Transaction, error) {
	return _Listing.Contract.Exit(&_Listing.TransactOpts, hash)
}

// Exit is a paid mutator transaction binding the contract method 0x0ca36263.
//
// Solidity: function exit(hash bytes32) returns()
func (_Listing *ListingTransactorSession) Exit(hash [32]byte) (*types.Transaction, error) {
	return _Listing.Contract.Exit(&_Listing.TransactOpts, hash)
}

// List is a paid mutator transaction binding the contract method 0xfe42bf1a.
//
// Solidity: function list(listing string) returns()
func (_Listing *ListingTransactor) List(opts *bind.TransactOpts, listing string) (*types.Transaction, error) {
	return _Listing.contract.Transact(opts, "list", listing)
}

// List is a paid mutator transaction binding the contract method 0xfe42bf1a.
//
// Solidity: function list(listing string) returns()
func (_Listing *ListingSession) List(listing string) (*types.Transaction, error) {
	return _Listing.Contract.List(&_Listing.TransactOpts, listing)
}

// List is a paid mutator transaction binding the contract method 0xfe42bf1a.
//
// Solidity: function list(listing string) returns()
func (_Listing *ListingTransactorSession) List(listing string) (*types.Transaction, error) {
	return _Listing.Contract.List(&_Listing.TransactOpts, listing)
}

// ResolveApplication is a paid mutator transaction binding the contract method 0x18e28e1c.
//
// Solidity: function resolveApplication(hash bytes32) returns()
func (_Listing *ListingTransactor) ResolveApplication(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Listing.contract.Transact(opts, "resolveApplication", hash)
}

// ResolveApplication is a paid mutator transaction binding the contract method 0x18e28e1c.
//
// Solidity: function resolveApplication(hash bytes32) returns()
func (_Listing *ListingSession) ResolveApplication(hash [32]byte) (*types.Transaction, error) {
	return _Listing.Contract.ResolveApplication(&_Listing.TransactOpts, hash)
}

// ResolveApplication is a paid mutator transaction binding the contract method 0x18e28e1c.
//
// Solidity: function resolveApplication(hash bytes32) returns()
func (_Listing *ListingTransactorSession) ResolveApplication(hash [32]byte) (*types.Transaction, error) {
	return _Listing.Contract.ResolveApplication(&_Listing.TransactOpts, hash)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0xd32c943a.
//
// Solidity: function resolveChallenge(hash bytes32) returns()
func (_Listing *ListingTransactor) ResolveChallenge(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Listing.contract.Transact(opts, "resolveChallenge", hash)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0xd32c943a.
//
// Solidity: function resolveChallenge(hash bytes32) returns()
func (_Listing *ListingSession) ResolveChallenge(hash [32]byte) (*types.Transaction, error) {
	return _Listing.Contract.ResolveChallenge(&_Listing.TransactOpts, hash)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0xd32c943a.
//
// Solidity: function resolveChallenge(hash bytes32) returns()
func (_Listing *ListingTransactorSession) ResolveChallenge(hash [32]byte) (*types.Transaction, error) {
	return _Listing.Contract.ResolveChallenge(&_Listing.TransactOpts, hash)
}

// WithdrawFromListing is a paid mutator transaction binding the contract method 0xa7e78675.
//
// Solidity: function withdrawFromListing(hash bytes32, amount uint256) returns()
func (_Listing *ListingTransactor) WithdrawFromListing(opts *bind.TransactOpts, hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Listing.contract.Transact(opts, "withdrawFromListing", hash, amount)
}

// WithdrawFromListing is a paid mutator transaction binding the contract method 0xa7e78675.
//
// Solidity: function withdrawFromListing(hash bytes32, amount uint256) returns()
func (_Listing *ListingSession) WithdrawFromListing(hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Listing.Contract.WithdrawFromListing(&_Listing.TransactOpts, hash, amount)
}

// WithdrawFromListing is a paid mutator transaction binding the contract method 0xa7e78675.
//
// Solidity: function withdrawFromListing(hash bytes32, amount uint256) returns()
func (_Listing *ListingTransactorSession) WithdrawFromListing(hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Listing.Contract.WithdrawFromListing(&_Listing.TransactOpts, hash, amount)
}

// ListingApplicationFailedIterator is returned from FilterApplicationFailed and is used to iterate over the raw logs and unpacked data for ApplicationFailed events raised by the Listing contract.
type ListingApplicationFailedIterator struct {
	Event *ListingApplicationFailed // Event containing the contract specifics and raw log

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
func (it *ListingApplicationFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingApplicationFailed)
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
		it.Event = new(ListingApplicationFailed)
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
func (it *ListingApplicationFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingApplicationFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingApplicationFailed represents a ApplicationFailed event raised by the Listing contract.
type ListingApplicationFailed struct {
	Hash      [32]byte
	Applicant common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterApplicationFailed is a free log retrieval operation binding the contract event 0x163b66974440f1696d484dfc777282273c42ff2fa247d72dbdbb3510d92cbda6.
//
// Solidity: e ApplicationFailed(hash indexed bytes32, applicant indexed address)
func (_Listing *ListingFilterer) FilterApplicationFailed(opts *bind.FilterOpts, hash [][32]byte, applicant []common.Address) (*ListingApplicationFailedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var applicantRule []interface{}
	for _, applicantItem := range applicant {
		applicantRule = append(applicantRule, applicantItem)
	}

	logs, sub, err := _Listing.contract.FilterLogs(opts, "ApplicationFailed", hashRule, applicantRule)
	if err != nil {
		return nil, err
	}
	return &ListingApplicationFailedIterator{contract: _Listing.contract, event: "ApplicationFailed", logs: logs, sub: sub}, nil
}

// WatchApplicationFailed is a free log subscription operation binding the contract event 0x163b66974440f1696d484dfc777282273c42ff2fa247d72dbdbb3510d92cbda6.
//
// Solidity: e ApplicationFailed(hash indexed bytes32, applicant indexed address)
func (_Listing *ListingFilterer) WatchApplicationFailed(opts *bind.WatchOpts, sink chan<- *ListingApplicationFailed, hash [][32]byte, applicant []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var applicantRule []interface{}
	for _, applicantItem := range applicant {
		applicantRule = append(applicantRule, applicantItem)
	}

	logs, sub, err := _Listing.contract.WatchLogs(opts, "ApplicationFailed", hashRule, applicantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingApplicationFailed)
				if err := _Listing.contract.UnpackLog(event, "ApplicationFailed", log); err != nil {
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

// ListingAppliedIterator is returned from FilterApplied and is used to iterate over the raw logs and unpacked data for Applied events raised by the Listing contract.
type ListingAppliedIterator struct {
	Event *ListingApplied // Event containing the contract specifics and raw log

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
func (it *ListingAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingApplied)
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
		it.Event = new(ListingApplied)
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
func (it *ListingAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingApplied represents a Applied event raised by the Listing contract.
type ListingApplied struct {
	Hash      [32]byte
	Applicant common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterApplied is a free log retrieval operation binding the contract event 0x26f7438941d95dc8495c99ef33fac91b2b04f98c73290027562cd3e612a3c85f.
//
// Solidity: e Applied(hash indexed bytes32, applicant indexed address)
func (_Listing *ListingFilterer) FilterApplied(opts *bind.FilterOpts, hash [][32]byte, applicant []common.Address) (*ListingAppliedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var applicantRule []interface{}
	for _, applicantItem := range applicant {
		applicantRule = append(applicantRule, applicantItem)
	}

	logs, sub, err := _Listing.contract.FilterLogs(opts, "Applied", hashRule, applicantRule)
	if err != nil {
		return nil, err
	}
	return &ListingAppliedIterator{contract: _Listing.contract, event: "Applied", logs: logs, sub: sub}, nil
}

// WatchApplied is a free log subscription operation binding the contract event 0x26f7438941d95dc8495c99ef33fac91b2b04f98c73290027562cd3e612a3c85f.
//
// Solidity: e Applied(hash indexed bytes32, applicant indexed address)
func (_Listing *ListingFilterer) WatchApplied(opts *bind.WatchOpts, sink chan<- *ListingApplied, hash [][32]byte, applicant []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var applicantRule []interface{}
	for _, applicantItem := range applicant {
		applicantRule = append(applicantRule, applicantItem)
	}

	logs, sub, err := _Listing.contract.WatchLogs(opts, "Applied", hashRule, applicantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingApplied)
				if err := _Listing.contract.UnpackLog(event, "Applied", log); err != nil {
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

// ListingChallengeFailedIterator is returned from FilterChallengeFailed and is used to iterate over the raw logs and unpacked data for ChallengeFailed events raised by the Listing contract.
type ListingChallengeFailedIterator struct {
	Event *ListingChallengeFailed // Event containing the contract specifics and raw log

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
func (it *ListingChallengeFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingChallengeFailed)
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
		it.Event = new(ListingChallengeFailed)
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
func (it *ListingChallengeFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingChallengeFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingChallengeFailed represents a ChallengeFailed event raised by the Listing contract.
type ListingChallengeFailed struct {
	Hash       [32]byte
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallengeFailed is a free log retrieval operation binding the contract event 0xa0e6c0bd204f59362c8b2ebdf77c63242bf779d8ac30c347933db099abcc0370.
//
// Solidity: e ChallengeFailed(hash indexed bytes32, challenger indexed address)
func (_Listing *ListingFilterer) FilterChallengeFailed(opts *bind.FilterOpts, hash [][32]byte, challenger []common.Address) (*ListingChallengeFailedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Listing.contract.FilterLogs(opts, "ChallengeFailed", hashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ListingChallengeFailedIterator{contract: _Listing.contract, event: "ChallengeFailed", logs: logs, sub: sub}, nil
}

// WatchChallengeFailed is a free log subscription operation binding the contract event 0xa0e6c0bd204f59362c8b2ebdf77c63242bf779d8ac30c347933db099abcc0370.
//
// Solidity: e ChallengeFailed(hash indexed bytes32, challenger indexed address)
func (_Listing *ListingFilterer) WatchChallengeFailed(opts *bind.WatchOpts, sink chan<- *ListingChallengeFailed, hash [][32]byte, challenger []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Listing.contract.WatchLogs(opts, "ChallengeFailed", hashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingChallengeFailed)
				if err := _Listing.contract.UnpackLog(event, "ChallengeFailed", log); err != nil {
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

// ListingChallengeSucceededIterator is returned from FilterChallengeSucceeded and is used to iterate over the raw logs and unpacked data for ChallengeSucceeded events raised by the Listing contract.
type ListingChallengeSucceededIterator struct {
	Event *ListingChallengeSucceeded // Event containing the contract specifics and raw log

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
func (it *ListingChallengeSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingChallengeSucceeded)
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
		it.Event = new(ListingChallengeSucceeded)
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
func (it *ListingChallengeSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingChallengeSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingChallengeSucceeded represents a ChallengeSucceeded event raised by the Listing contract.
type ListingChallengeSucceeded struct {
	Hash       [32]byte
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallengeSucceeded is a free log retrieval operation binding the contract event 0xb7e10908a54924c8e096789495f5a4958fed82b49d856d22b208a23b398306bb.
//
// Solidity: e ChallengeSucceeded(hash indexed bytes32, challenger indexed address)
func (_Listing *ListingFilterer) FilterChallengeSucceeded(opts *bind.FilterOpts, hash [][32]byte, challenger []common.Address) (*ListingChallengeSucceededIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Listing.contract.FilterLogs(opts, "ChallengeSucceeded", hashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ListingChallengeSucceededIterator{contract: _Listing.contract, event: "ChallengeSucceeded", logs: logs, sub: sub}, nil
}

// WatchChallengeSucceeded is a free log subscription operation binding the contract event 0xb7e10908a54924c8e096789495f5a4958fed82b49d856d22b208a23b398306bb.
//
// Solidity: e ChallengeSucceeded(hash indexed bytes32, challenger indexed address)
func (_Listing *ListingFilterer) WatchChallengeSucceeded(opts *bind.WatchOpts, sink chan<- *ListingChallengeSucceeded, hash [][32]byte, challenger []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Listing.contract.WatchLogs(opts, "ChallengeSucceeded", hashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingChallengeSucceeded)
				if err := _Listing.contract.UnpackLog(event, "ChallengeSucceeded", log); err != nil {
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

// ListingChallengedIterator is returned from FilterChallenged and is used to iterate over the raw logs and unpacked data for Challenged events raised by the Listing contract.
type ListingChallengedIterator struct {
	Event *ListingChallenged // Event containing the contract specifics and raw log

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
func (it *ListingChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingChallenged)
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
		it.Event = new(ListingChallenged)
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
func (it *ListingChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingChallenged represents a Challenged event raised by the Listing contract.
type ListingChallenged struct {
	Hash       [32]byte
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallenged is a free log retrieval operation binding the contract event 0xe9479421670c3425a1497ce47a53af8bd96ce5bd0741e96221ba0acace3f7d47.
//
// Solidity: e Challenged(hash indexed bytes32, challenger indexed address)
func (_Listing *ListingFilterer) FilterChallenged(opts *bind.FilterOpts, hash [][32]byte, challenger []common.Address) (*ListingChallengedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Listing.contract.FilterLogs(opts, "Challenged", hashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ListingChallengedIterator{contract: _Listing.contract, event: "Challenged", logs: logs, sub: sub}, nil
}

// WatchChallenged is a free log subscription operation binding the contract event 0xe9479421670c3425a1497ce47a53af8bd96ce5bd0741e96221ba0acace3f7d47.
//
// Solidity: e Challenged(hash indexed bytes32, challenger indexed address)
func (_Listing *ListingFilterer) WatchChallenged(opts *bind.WatchOpts, sink chan<- *ListingChallenged, hash [][32]byte, challenger []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Listing.contract.WatchLogs(opts, "Challenged", hashRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingChallenged)
				if err := _Listing.contract.UnpackLog(event, "Challenged", log); err != nil {
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

// ListingListedIterator is returned from FilterListed and is used to iterate over the raw logs and unpacked data for Listed events raised by the Listing contract.
type ListingListedIterator struct {
	Event *ListingListed // Event containing the contract specifics and raw log

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
func (it *ListingListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingListed)
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
		it.Event = new(ListingListed)
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
func (it *ListingListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingListed represents a Listed event raised by the Listing contract.
type ListingListed struct {
	Hash   [32]byte
	Owner  common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterListed is a free log retrieval operation binding the contract event 0x7f1a2c3e2883554425dc0f9f24dcfcdf54213b186c550080373dcc65813aa8d0.
//
// Solidity: e Listed(hash indexed bytes32, owner indexed address, reward uint256)
func (_Listing *ListingFilterer) FilterListed(opts *bind.FilterOpts, hash [][32]byte, owner []common.Address) (*ListingListedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Listing.contract.FilterLogs(opts, "Listed", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ListingListedIterator{contract: _Listing.contract, event: "Listed", logs: logs, sub: sub}, nil
}

// WatchListed is a free log subscription operation binding the contract event 0x7f1a2c3e2883554425dc0f9f24dcfcdf54213b186c550080373dcc65813aa8d0.
//
// Solidity: e Listed(hash indexed bytes32, owner indexed address, reward uint256)
func (_Listing *ListingFilterer) WatchListed(opts *bind.WatchOpts, sink chan<- *ListingListed, hash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Listing.contract.WatchLogs(opts, "Listed", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingListed)
				if err := _Listing.contract.UnpackLog(event, "Listed", log); err != nil {
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

// ListingListingConvertedIterator is returned from FilterListingConverted and is used to iterate over the raw logs and unpacked data for ListingConverted events raised by the Listing contract.
type ListingListingConvertedIterator struct {
	Event *ListingListingConverted // Event containing the contract specifics and raw log

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
func (it *ListingListingConvertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingListingConverted)
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
		it.Event = new(ListingListingConverted)
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
func (it *ListingListingConvertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingListingConvertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingListingConverted represents a ListingConverted event raised by the Listing contract.
type ListingListingConverted struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterListingConverted is a free log retrieval operation binding the contract event 0xfd6a9ca1840ac9f8b7c592783e5415fdc282281630286c93b07b6d40f3c0ba57.
//
// Solidity: e ListingConverted(hash indexed bytes32)
func (_Listing *ListingFilterer) FilterListingConverted(opts *bind.FilterOpts, hash [][32]byte) (*ListingListingConvertedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Listing.contract.FilterLogs(opts, "ListingConverted", hashRule)
	if err != nil {
		return nil, err
	}
	return &ListingListingConvertedIterator{contract: _Listing.contract, event: "ListingConverted", logs: logs, sub: sub}, nil
}

// WatchListingConverted is a free log subscription operation binding the contract event 0xfd6a9ca1840ac9f8b7c592783e5415fdc282281630286c93b07b6d40f3c0ba57.
//
// Solidity: e ListingConverted(hash indexed bytes32)
func (_Listing *ListingFilterer) WatchListingConverted(opts *bind.WatchOpts, sink chan<- *ListingListingConverted, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Listing.contract.WatchLogs(opts, "ListingConverted", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingListingConverted)
				if err := _Listing.contract.UnpackLog(event, "ListingConverted", log); err != nil {
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

// ListingListingDepositIterator is returned from FilterListingDeposit and is used to iterate over the raw logs and unpacked data for ListingDeposit events raised by the Listing contract.
type ListingListingDepositIterator struct {
	Event *ListingListingDeposit // Event containing the contract specifics and raw log

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
func (it *ListingListingDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingListingDeposit)
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
		it.Event = new(ListingListingDeposit)
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
func (it *ListingListingDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingListingDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingListingDeposit represents a ListingDeposit event raised by the Listing contract.
type ListingListingDeposit struct {
	Hash      [32]byte
	Owner     common.Address
	Deposited *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterListingDeposit is a free log retrieval operation binding the contract event 0x6f27446a1e23c7e7dd81a6a77f81438be78145de331a5806f2c44e68738bdcef.
//
// Solidity: e ListingDeposit(hash indexed bytes32, owner indexed address, deposited uint256)
func (_Listing *ListingFilterer) FilterListingDeposit(opts *bind.FilterOpts, hash [][32]byte, owner []common.Address) (*ListingListingDepositIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Listing.contract.FilterLogs(opts, "ListingDeposit", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ListingListingDepositIterator{contract: _Listing.contract, event: "ListingDeposit", logs: logs, sub: sub}, nil
}

// WatchListingDeposit is a free log subscription operation binding the contract event 0x6f27446a1e23c7e7dd81a6a77f81438be78145de331a5806f2c44e68738bdcef.
//
// Solidity: e ListingDeposit(hash indexed bytes32, owner indexed address, deposited uint256)
func (_Listing *ListingFilterer) WatchListingDeposit(opts *bind.WatchOpts, sink chan<- *ListingListingDeposit, hash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Listing.contract.WatchLogs(opts, "ListingDeposit", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingListingDeposit)
				if err := _Listing.contract.UnpackLog(event, "ListingDeposit", log); err != nil {
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

// ListingListingRemovedIterator is returned from FilterListingRemoved and is used to iterate over the raw logs and unpacked data for ListingRemoved events raised by the Listing contract.
type ListingListingRemovedIterator struct {
	Event *ListingListingRemoved // Event containing the contract specifics and raw log

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
func (it *ListingListingRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingListingRemoved)
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
		it.Event = new(ListingListingRemoved)
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
func (it *ListingListingRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingListingRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingListingRemoved represents a ListingRemoved event raised by the Listing contract.
type ListingListingRemoved struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterListingRemoved is a free log retrieval operation binding the contract event 0x50425cae216bd151d26c8e8bb9779cc899c99d72f78081f2ceec9a99f001ff79.
//
// Solidity: e ListingRemoved(hash indexed bytes32)
func (_Listing *ListingFilterer) FilterListingRemoved(opts *bind.FilterOpts, hash [][32]byte) (*ListingListingRemovedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Listing.contract.FilterLogs(opts, "ListingRemoved", hashRule)
	if err != nil {
		return nil, err
	}
	return &ListingListingRemovedIterator{contract: _Listing.contract, event: "ListingRemoved", logs: logs, sub: sub}, nil
}

// WatchListingRemoved is a free log subscription operation binding the contract event 0x50425cae216bd151d26c8e8bb9779cc899c99d72f78081f2ceec9a99f001ff79.
//
// Solidity: e ListingRemoved(hash indexed bytes32)
func (_Listing *ListingFilterer) WatchListingRemoved(opts *bind.WatchOpts, sink chan<- *ListingListingRemoved, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Listing.contract.WatchLogs(opts, "ListingRemoved", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingListingRemoved)
				if err := _Listing.contract.UnpackLog(event, "ListingRemoved", log); err != nil {
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

// ListingListingWithdrawIterator is returned from FilterListingWithdraw and is used to iterate over the raw logs and unpacked data for ListingWithdraw events raised by the Listing contract.
type ListingListingWithdrawIterator struct {
	Event *ListingListingWithdraw // Event containing the contract specifics and raw log

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
func (it *ListingListingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ListingListingWithdraw)
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
		it.Event = new(ListingListingWithdraw)
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
func (it *ListingListingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ListingListingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ListingListingWithdraw represents a ListingWithdraw event raised by the Listing contract.
type ListingListingWithdraw struct {
	Hash      [32]byte
	Owner     common.Address
	Withdrawn *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterListingWithdraw is a free log retrieval operation binding the contract event 0x87c48f853a68e9374cf8389dab199f46adaffa48468fe9a4e2dc2edf75286db7.
//
// Solidity: e ListingWithdraw(hash indexed bytes32, owner indexed address, withdrawn uint256)
func (_Listing *ListingFilterer) FilterListingWithdraw(opts *bind.FilterOpts, hash [][32]byte, owner []common.Address) (*ListingListingWithdrawIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Listing.contract.FilterLogs(opts, "ListingWithdraw", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ListingListingWithdrawIterator{contract: _Listing.contract, event: "ListingWithdraw", logs: logs, sub: sub}, nil
}

// WatchListingWithdraw is a free log subscription operation binding the contract event 0x87c48f853a68e9374cf8389dab199f46adaffa48468fe9a4e2dc2edf75286db7.
//
// Solidity: e ListingWithdraw(hash indexed bytes32, owner indexed address, withdrawn uint256)
func (_Listing *ListingFilterer) WatchListingWithdraw(opts *bind.WatchOpts, sink chan<- *ListingListingWithdraw, hash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Listing.contract.WatchLogs(opts, "ListingWithdraw", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ListingListingWithdraw)
				if err := _Listing.contract.UnpackLog(event, "ListingWithdraw", log); err != nil {
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
