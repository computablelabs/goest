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
const ListingABI = "[{\"name\":\"ApplicationFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Applied\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"applicant\",\"indexed\":true},{\"type\":\"bytes32\",\"name\":\"dataHash\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Challenged\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"challenger\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ChallengeFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"challenger\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"reward\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ChallengeSucceeded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"challenger\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"reward\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Listed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"reward\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingConverted\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingDeposit\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"deposited\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingRemoved\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingWithdraw\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"withdrawn\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"market_token_addr\"},{\"type\":\"address\",\"name\":\"voting_addr\"},{\"type\":\"address\",\"name\":\"p11r_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"isListed\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":700},{\"name\":\"isListing\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1033},{\"name\":\"wasListing\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":754},{\"name\":\"isListingOwner\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":758},{\"name\":\"depositToListing\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":41660},{\"name\":\"withdrawFromListing\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":41959},{\"name\":\"list\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"listing\"},{\"type\":\"bytes32\",\"name\":\"data_hash\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":259285},{\"name\":\"getListingCount\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":693},{\"name\":\"getListingKey\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"index\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":940},{\"name\":\"getListing\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"bytes32\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2659},{\"name\":\"convertListing\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":83246},{\"name\":\"resolveApplication\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":285441},{\"name\":\"challenge\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"chall\"},{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":48050},{\"name\":\"getChallenge\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1036},{\"name\":\"resolveChallenge\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":369788},{\"name\":\"exit\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":276359}]"

// ListingBin is the compiled bytecode used for deploying new contracts.
const ListingBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05260606120b46101403934156100a757600080fd5b60206120b460c03960c05160205181106100c057600080fd5b50602060206120b40160c03960c05160205181106100dd57600080fd5b50602060406120b40160c03960c05160205181106100fa57600080fd5b503360055561014051600655610160516007556101805160085561209c56600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263ecefbdc660005114156100de57602060046101403734156100b457600080fd5b6001600160006101405160e05260c052604060c02060c052602060c02001541460005260206000f3005b636aa0966b600051141561014357602060046101403734156100ff57600080fd5b6101405160006101405160e05260c052604060c02060c052602060c02054620f4240811061012c57600080fd5b600260c052602060c02001541460005260206000f3005b6382c3eb3e60005114156101aa576020600461014037341561016457600080fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60006101405160e05260c052604060c02060c052602060c020541460005260206000f3005b6302a47cd760005114156101fd57602060046101403734156101cb57600080fd5b60043560205181106101dc57600080fd5b50600160036101405160e05260c052604060c02054101560005260206000f3005b63341ff0896000511415610329576040600461014037341561021e57600080fd5b6101405160006101405160e05260c052604060c02060c052602060c02054620f4240811061024b57600080fd5b600260c052602060c02001541461026157600080fd5b6006543b61026e57600080fd5b60065430141561027d57600080fd5b602061024060646323b872dd61018052336101a052306101c052610160516101e05261019c60006006545af16102b257600080fd5b60005061024050600460006101405160e05260c052604060c02060c052602060c0200180546101605182540110156102e957600080fd5b61016051815401815550610160516102605233610140517f6f27446a1e23c7e7dd81a6a77f81438be78145de331a5806f2c44e68738bdcef6020610260a3005b63a7e786756000511415610475576040600461014037341561034a57600080fd5b6101405160006101405160e05260c052604060c02060c052602060c02054620f4240811061037757600080fd5b600260c052602060c02001541461038d57600080fd5b33600260006101405160e05260c052604060c02060c052602060c0200154146103b557600080fd5b600460006101405160e05260c052604060c02060c052602060c0200161016051815410156103e257600080fd5b610160518154038155506006543b6103f957600080fd5b60065430141561040857600080fd5b6020610220604463a9059cbb61018052336101a052610160516101c05261019c60006006545af161043857600080fd5b60005061022050610160516102405233610140517f87c48f853a68e9374cf8389dab199f46adaffa48468fe9a4e2dc2edf75286db76020610240a3005b635148cd0f6000511415610744576060600461014037341561049657600080fd5b60606004356004016101a03760406004356004013511156104b657600080fd5b6101a0805160208201209050610220526007543b6104d357600080fd5b6007543014156104e257600080fd5b60206102c06024631f1b029961024052610220516102605261025c6007545afa61050b57600080fd5b6000506102c05161051b57600080fd5b6007543b61052857600080fd5b60075430141561053757600080fd5b600060006064637ff45d8b61036052610220516103805260016103a0526008543b61056157600080fd5b60085430141561057057600080fd5b60206103406004632d0d2bc66102e0526102fc6008545afa61059157600080fd5b600050610340516103c05261037c60006007545af16105af57600080fd5b33600260006102205160e05260c052604060c02060c052602060c020015561016051600360006102205160e05260c052604060c02060c052602060c020015560015460006102205160e05260c052604060c02060c052602060c0205561022051600154620f4240811061062157600080fd5b600260c052602060c0200155600160605160018254018060405190131561064757600080fd5b809190121561065557600080fd5b81555060033360e05260c052604060c02080546001825401101561067857600080fd5b60018154018155506000610180511115610706576006543b61069957600080fd5b6006543014156106a857600080fd5b60206104e060646323b872dd6104205233610440523061046052610180516104805261043c60006006545af16106dd57600080fd5b6000506104e05061018051600460006102205160e05260c052604060c02060c052602060c02001555b6101605161050052610180516105205233610220517f2be7f407f8117561d2a0100982c56a4a1ec8ffc24bbeb7b087b2ca8a224cc56d6040610500a3005b6387ed92d7600051141561076a57341561075d57600080fd5b60015460005260206000f3005b63c48dc1af60005114156107d9576020600461014037341561078b57600080fd5b606051600435806040519013156107a157600080fd5b80919012156107af57600080fd5b5061014051620f424081106107c357600080fd5b600260c052602060c020015460005260206000f3005b63175c0d1660005114156108b557602060046101403734156107fa57600080fd5b60a061016052610180600160006101405160e05260c052604060c02060c052602060c02001548152600260006101405160e05260c052604060c02060c052602060c02001548160200152600360006101405160e05260c052604060c02060c052602060c02001548160400152600460006101405160e05260c052604060c02060c052602060c02001548160600152600560006101405160e05260c052604060c02060c052602060c020015481608001525061016051610180f3005b63d748484d6000511415610acd57602060046101403734156108d657600080fd5b6001600160006101405160e05260c052604060c02060c052602060c0200154146108ff57600080fd5b33600260006101405160e05260c052604060c02060c052602060c02001541461092757600080fd5b600460006101405160e05260c052604060c02060c052602060c0200154600560006101405160e05260c052604060c02060c052602060c0200154600460006101405160e05260c052604060c02060c052602060c020015401101561098a57600080fd5b600560006101405160e05260c052604060c02060c052602060c0200154600460006101405160e05260c052604060c02060c052602060c020015401610160526000610160511115610a82576000600460006101405160e05260c052604060c02060c052602060c02001556000600560006101405160e05260c052604060c02060c052602060c02001556006543b610a2057600080fd5b600654301415610a2f57600080fd5b6020610220604463a9059cbb61018052600260006101405160e05260c052604060c02060c052602060c02001546101a052610160516101c05261019c60006006545af1610a7b57600080fd5b6000506102205b5030600260006101405160e05260c052604060c02060c052602060c0200155610140517ffd6a9ca1840ac9f8b7c592783e5415fdc282281630286c93b07b6d40f3c0ba5760006000a2005b600015610eed575b61016052610140526101405160006101405160e05260c052604060c02060c052602060c02054620f42408110610b0a57600080fd5b600260c052602060c020015414610b2057600080fd5b6000600460006101405160e05260c052604060c02060c052602060c02001541115610bed576006543b610b5257600080fd5b600654301415610b6157600080fd5b6020610220604463a9059cbb61018052600260006101405160e05260c052604060c02060c052602060c02001546101a052600460006101405160e05260c052604060c02060c052602060c02001546101c05261019c60006006545af1610bc657600080fd5b600050610220506000600460006101405160e05260c052604060c02060c052602060c02001555b6000600560006101405160e05260c052604060c02060c052602060c02001541115610c91576006543b610c1f57600080fd5b600654301415610c2e57600080fd5b6000600060246342966c6861024052600560006101405160e05260c052604060c02060c052602060c02001546102605261025c60006006545af1610c7157600080fd5b6000600560006101405160e05260c052604060c02060c052602060c02001555b6001606051600182540380604051901315610cab57600080fd5b8091901215610cb957600080fd5b81555060006001541315610d495760006101405160e05260c052604060c02060c052602060c020546102c052600154620f42408110610cf757600080fd5b600260c052602060c02001546102e0526102e0516102c051620f42408110610d1e57600080fd5b600260c052602060c02001556102c05160006102e05160e05260c052604060c02060c052602060c020555b6000600154620f42408110610d5d57600080fd5b600260c052602060c02001556003600260006101405160e05260c052604060c02060c052602060c020015460e05260c052604060c020600181541015610da257600080fd5b600181540381555060016003600260006101405160e05260c052604060c02060c052602060c020015460e05260c052604060c020541015610e3d576007543b610dea57600080fd5b600754301415610df957600080fd5b60006000602463335d808061030052600260006101405160e05260c052604060c02060c052602060c02001546103205261031c60006007545af1610e3c57600080fd5b5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60006101405160e05260c052604060c02060c052602060c020556001600160006101405160e05260c052604060c02060c052602060c02001541415610ebd576000600160006101405160e05260c052604060c02060c052602060c02001555b610140517f50425cae216bd151d26c8e8bb9779cc899c99d72f78081f2ceec9a99f001ff7960006000a261016051565b6318e28e1c60005114156113b25760206004610140373415610f0e57600080fd5b6007543b610f1b57600080fd5b600754301415610f2a57600080fd5b60206101e060246362f9a55d61016052336101805261017c6007545afa610f5057600080fd5b6000506101e051610f6057600080fd5b6007543b610f6d57600080fd5b600754301415610f7c57600080fd5b60206102a0604463af61f76061020052610140516102205260016102405261021c6007545afa610fab57600080fd5b6000506102a051610fbb57600080fd5b6001600160006101405160e05260c052604060c02060c052602060c02001541415610fe557600080fd5b6007543b610ff257600080fd5b60075430141561100157600080fd5b6020610340602463327322c86102c052610140516102e0526102dc6007545afa61102a57600080fd5b6000506103405161103a57600080fd5b6007543b61104757600080fd5b60075430141561105657600080fd5b60206105006044638f354b796104605261014051610480526008543b61107b57600080fd5b60085430141561108a57600080fd5b6020610440600463c26c12eb6103e0526103fc6008545afa6110ab57600080fd5b600050610440516104a05261047c6007545afa6110c757600080fd5b60005061050051156112c0576001600160006101405160e05260c052604060c02060c052602060c02001556008543b6110ff57600080fd5b60085430141561110e57600080fd5b60206105a0600463dc2b28536105405261055c6008545afa61112f57600080fd5b6000506105a051610520526006543b61114757600080fd5b60065430141561115657600080fd5b60006000602463a0712d686105c052610520516105e0526105dc60006006545af161118057600080fd5b61052051600560006101405160e05260c052604060c02060c052602060c02001556007543b6111ae57600080fd5b6007543014156111bd57600080fd5b60206106c0602463c666b58061064052600260006101405160e05260c052604060c02060c052602060c02001546106605261065c6007545afa6111ff57600080fd5b6000506106c0511561126b576007543b61121857600080fd5b60075430141561122757600080fd5b60006000602463471bb3096106e052600260006101405160e05260c052604060c02060c052602060c0200154610700526106fc60006007545af161126a57600080fd5b5b6105205161076052600260006101405160e05260c052604060c02060c052602060c0200154610140517f7f1a2c3e2883554425dc0f9f24dcfcdf54213b186c550080373dcc65813aa8d06020610760a361136a565b6101405161016051610180516101a0516101c0516101e05161020051610220516102405161026051610280516102a0516102c0516102e05161030051610320516103405163550ee1336103605261014051610380526103805160065801610ad5565b6103405261032052610300526102e0526102c0526102a05261028052610260526102405261022052610200526101e0526101c0526101a0526101805261016052610140526000505b6007543b61137757600080fd5b60075430141561138657600080fd5b6000600060246389bb617c61078052610140516107a05261079c60006007545af16113b057600080fd5b005b63cd32dcc560005114156115f257604060046101403734156113d357600080fd5b60606004356004016101803760406004356004013511156113f357600080fd5b600160006101605160e05260c052604060c02060c052602060c020015461141957600080fd5b6008543b61142657600080fd5b60085430141561143557600080fd5b60206102806004635aebe7686102205261023c6008545afa61145657600080fd5b60005061028051610200526101808051602082012090506102a0526007543b61147e57600080fd5b60075430141561148d57600080fd5b60206103406024631f1b02996102c0526102a0516102e0526102dc6007545afa6114b657600080fd5b600050610340516114c657600080fd5b6006543b6114d357600080fd5b6006543014156114e257600080fd5b602061042060646323b872dd610360523361038052306103a052610200516103c05261037c60006006545af161151757600080fd5b600050610420503360046101605160e05260c052604060c020556007543b61153e57600080fd5b60075430141561154d57600080fd5b600060006064637ff45d8b6104c052610160516104e0526002610500526008543b61157757600080fd5b60085430141561158657600080fd5b60206104a06004632d0d2bc66104405261045c6008545afa6115a757600080fd5b6000506104a051610520526104dc60006007545af16115c557600080fd5b33610160517fe9479421670c3425a1497ce47a53af8bd96ce5bd0741e96221ba0acace3f7d4760006000a3005b63458d2bf1600051141561162f576020600461014037341561161357600080fd5b60046101405160e05260c052604060c0205460005260206000f3005b63d32c943a6000511415611e62576020600461014037341561165057600080fd5b6007543b61165d57600080fd5b60075430141561166c57600080fd5b60206101e060246362f9a55d61016052336101805261017c6007545afa61169257600080fd5b6000506101e0516116a257600080fd5b6007543b6116af57600080fd5b6007543014156116be57600080fd5b60206102a0604463af61f76061020052610140516102205260026102405261021c6007545afa6116ed57600080fd5b6000506102a0516116fd57600080fd5b6007543b61170a57600080fd5b60075430141561171957600080fd5b6020610340602463327322c86102c052610140516102e0526102dc6007545afa61174257600080fd5b6000506103405161175257600080fd5b6008543b61175f57600080fd5b60085430141561176e57600080fd5b60206103e06004635aebe7686103805261039c6008545afa61178f57600080fd5b6000506103e05161036052600460006101405160e05260c052604060c02060c052602060c020015461040052600560006101405160e05260c052604060c02060c052602060c020015461042052610360516104005161042051610400510110156117f857600080fd5b61042051610400510110156119d657600061040051610420516104005101101561182157600080fd5b61042051610400510111156118ef576000600460006101405160e05260c052604060c02060c052602060c02001556000600560006101405160e05260c052604060c02060c052602060c02001556006543b61187b57600080fd5b60065430141561188a57600080fd5b6020610820604463a9059cbb6107805260046101405160e05260c052604060c020546107a0526104005161042051610400510110156118c857600080fd5b6104205161040051016107c05261079c60006006545af16118e857600080fd5b6000506108205b506101405161016051610180516101a0516101c0516101e05161020051610220516102405161026051610280516102a0516102c0516102e05161030051610320516103405161036051610380516103a0516103c0516103e051610400516104205163550ee1336108405261014051610860526108605160065801610ad5565b61042052610400526103e0526103c0526103a05261038052610360526103405261032052610300526102e0526102c0526102a05261028052610260526102405261022052610200526101e0526101c0526101a052610180526101605261014052600050611e60565b6007543b6119e357600080fd5b6007543014156119f257600080fd5b60206105806044638f354b796104e05261014051610500526008543b611a1757600080fd5b600854301415611a2657600080fd5b60206104c0600463c26c12eb6104605261047c6008545afa611a4757600080fd5b6000506104c051610520526104fc6007545afa611a6357600080fd5b6000506105805115611d85576103605161040051101515611aba57600460006101405160e05260c052604060c02060c052602060c020016103605181541015611aab57600080fd5b61036051815403815550611b41565b6000600460006101405160e05260c052604060c02060c052602060c0200155600560006101405160e05260c052604060c02060c052602060c0200161040051610360511015611b0857600080fd5b61040051610360510381541015611b1e57600080fd5b61040051610360511015611b3157600080fd5b6104005161036051038154038155505b6006543b611b4e57600080fd5b600654301415611b5d57600080fd5b6020610640604463a9059cbb6105a05260046101405160e05260c052604060c020546105c052610360511515611b94576000611bb4565b600261036051600261036051020414611bac57600080fd5b600261036051025b6105e0526105bc60006006545af1611bcb57600080fd5b600050610640506101405161016051610180516101a0516101c0516101e05161020051610220516102405161026051610280516102a0516102c0516102e05161030051610320516103405161036051610380516103a0516103c0516103e05161040051610420516104405161046051610480516104a0516104c0516104e05161050051610520516105405161056051610580516105a0516105c0516105e05161060051610620516106405163550ee1336106605261014051610680526106805160065801610ad5565b6106405261062052610600526105e0526105c0526105a05261058052610560526105405261052052610500526104e0526104c0526104a05261048052610460526104405261042052610400526103e0526103c0526103a05261038052610360526103405261032052610300526102e0526102c0526102a05261028052610260526102405261022052610200526101e0526101c0526101a052610180526101605261014052600050610360516106e05260046101405160e05260c052604060c02054610140517fb2432a6d7681dc3f7416efc0e321723b96bd7a96d2d02d228dc2a67c51cf032260206106e0a3611e05565b600560006101405160e05260c052604060c02060c052602060c020018054610360518254011015611db557600080fd5b61036051815401815550610360516104405260046101405160e05260c052604060c02054610140517fc7aab91fe7b50a881bd71df2cf07e9a6e5351721087cdb09bc70603f618e70576020610440a35b600060046101405160e05260c052604060c020556007543b611e2657600080fd5b600754301415611e3557600080fd5b6000600060246389bb617c61070052610140516107205261071c60006007545af1611e5f57600080fd5b5b005b630ca362636000511415611f7d5760206004610140373415611e8357600080fd5b33600260006101405160e05260c052604060c02060c052602060c020015414611eab57600080fd5b6001600160006101405160e05260c052604060c02060c052602060c020015414611ed457600080fd5b6007543b611ee157600080fd5b600754301415611ef057600080fd5b60206101e0602463b89694c661016052610140516101805261017c6007545afa611f1957600080fd5b6000506101e05115611f2a57600080fd5b6101405161016051610180516101a0516101c0516101e05163550ee1336102005261014051610220526102205160065801610ad5565b6101e0526101c0526101a052610180526101605261014052600050005b60006000fd5b61011961209c0361011960003961011961209c036000f3`

// DeployListing deploys a new Ethereum contract, binding an instance of Listing to it.
func DeployListing(auth *bind.TransactOpts, backend bind.ContractBackend, market_token_addr common.Address, voting_addr common.Address, p11r_addr common.Address) (common.Address, *types.Transaction, *Listing, error) {
	parsed, err := abi.JSON(strings.NewReader(ListingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ListingBin), backend, market_token_addr, voting_addr, p11r_addr)
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

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(hash bytes32) constant returns(out address)
func (_Listing *ListingCaller) GetChallenge(opts *bind.CallOpts, hash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Listing.contract.Call(opts, out, "getChallenge", hash)
	return *ret0, err
}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(hash bytes32) constant returns(out address)
func (_Listing *ListingSession) GetChallenge(hash [32]byte) (common.Address, error) {
	return _Listing.Contract.GetChallenge(&_Listing.CallOpts, hash)
}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(hash bytes32) constant returns(out address)
func (_Listing *ListingCallerSession) GetChallenge(hash [32]byte) (common.Address, error) {
	return _Listing.Contract.GetChallenge(&_Listing.CallOpts, hash)
}

// GetListing is a free data retrieval call binding the contract method 0x175c0d16.
//
// Solidity: function getListing(hash bytes32) constant returns(out bool, out address, out bytes32, out uint256, out uint256)
func (_Listing *ListingCaller) GetListing(opts *bind.CallOpts, hash [32]byte) (bool, common.Address, [32]byte, *big.Int, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(common.Address)
		ret2 = new([32]byte)
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
	err := _Listing.contract.Call(opts, out, "getListing", hash)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetListing is a free data retrieval call binding the contract method 0x175c0d16.
//
// Solidity: function getListing(hash bytes32) constant returns(out bool, out address, out bytes32, out uint256, out uint256)
func (_Listing *ListingSession) GetListing(hash [32]byte) (bool, common.Address, [32]byte, *big.Int, *big.Int, error) {
	return _Listing.Contract.GetListing(&_Listing.CallOpts, hash)
}

// GetListing is a free data retrieval call binding the contract method 0x175c0d16.
//
// Solidity: function getListing(hash bytes32) constant returns(out bool, out address, out bytes32, out uint256, out uint256)
func (_Listing *ListingCallerSession) GetListing(hash [32]byte) (bool, common.Address, [32]byte, *big.Int, *big.Int, error) {
	return _Listing.Contract.GetListing(&_Listing.CallOpts, hash)
}

// GetListingCount is a free data retrieval call binding the contract method 0x87ed92d7.
//
// Solidity: function getListingCount() constant returns(out int128)
func (_Listing *ListingCaller) GetListingCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Listing.contract.Call(opts, out, "getListingCount")
	return *ret0, err
}

// GetListingCount is a free data retrieval call binding the contract method 0x87ed92d7.
//
// Solidity: function getListingCount() constant returns(out int128)
func (_Listing *ListingSession) GetListingCount() (*big.Int, error) {
	return _Listing.Contract.GetListingCount(&_Listing.CallOpts)
}

// GetListingCount is a free data retrieval call binding the contract method 0x87ed92d7.
//
// Solidity: function getListingCount() constant returns(out int128)
func (_Listing *ListingCallerSession) GetListingCount() (*big.Int, error) {
	return _Listing.Contract.GetListingCount(&_Listing.CallOpts)
}

// GetListingKey is a free data retrieval call binding the contract method 0xc48dc1af.
//
// Solidity: function getListingKey(index int128) constant returns(out bytes32)
func (_Listing *ListingCaller) GetListingKey(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Listing.contract.Call(opts, out, "getListingKey", index)
	return *ret0, err
}

// GetListingKey is a free data retrieval call binding the contract method 0xc48dc1af.
//
// Solidity: function getListingKey(index int128) constant returns(out bytes32)
func (_Listing *ListingSession) GetListingKey(index *big.Int) ([32]byte, error) {
	return _Listing.Contract.GetListingKey(&_Listing.CallOpts, index)
}

// GetListingKey is a free data retrieval call binding the contract method 0xc48dc1af.
//
// Solidity: function getListingKey(index int128) constant returns(out bytes32)
func (_Listing *ListingCallerSession) GetListingKey(index *big.Int) ([32]byte, error) {
	return _Listing.Contract.GetListingKey(&_Listing.CallOpts, index)
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

// IsListing is a free data retrieval call binding the contract method 0x6aa0966b.
//
// Solidity: function isListing(hash bytes32) constant returns(out bool)
func (_Listing *ListingCaller) IsListing(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Listing.contract.Call(opts, out, "isListing", hash)
	return *ret0, err
}

// IsListing is a free data retrieval call binding the contract method 0x6aa0966b.
//
// Solidity: function isListing(hash bytes32) constant returns(out bool)
func (_Listing *ListingSession) IsListing(hash [32]byte) (bool, error) {
	return _Listing.Contract.IsListing(&_Listing.CallOpts, hash)
}

// IsListing is a free data retrieval call binding the contract method 0x6aa0966b.
//
// Solidity: function isListing(hash bytes32) constant returns(out bool)
func (_Listing *ListingCallerSession) IsListing(hash [32]byte) (bool, error) {
	return _Listing.Contract.IsListing(&_Listing.CallOpts, hash)
}

// IsListingOwner is a free data retrieval call binding the contract method 0x02a47cd7.
//
// Solidity: function isListingOwner(addr address) constant returns(out bool)
func (_Listing *ListingCaller) IsListingOwner(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Listing.contract.Call(opts, out, "isListingOwner", addr)
	return *ret0, err
}

// IsListingOwner is a free data retrieval call binding the contract method 0x02a47cd7.
//
// Solidity: function isListingOwner(addr address) constant returns(out bool)
func (_Listing *ListingSession) IsListingOwner(addr common.Address) (bool, error) {
	return _Listing.Contract.IsListingOwner(&_Listing.CallOpts, addr)
}

// IsListingOwner is a free data retrieval call binding the contract method 0x02a47cd7.
//
// Solidity: function isListingOwner(addr address) constant returns(out bool)
func (_Listing *ListingCallerSession) IsListingOwner(addr common.Address) (bool, error) {
	return _Listing.Contract.IsListingOwner(&_Listing.CallOpts, addr)
}

// WasListing is a free data retrieval call binding the contract method 0x82c3eb3e.
//
// Solidity: function wasListing(hash bytes32) constant returns(out bool)
func (_Listing *ListingCaller) WasListing(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Listing.contract.Call(opts, out, "wasListing", hash)
	return *ret0, err
}

// WasListing is a free data retrieval call binding the contract method 0x82c3eb3e.
//
// Solidity: function wasListing(hash bytes32) constant returns(out bool)
func (_Listing *ListingSession) WasListing(hash [32]byte) (bool, error) {
	return _Listing.Contract.WasListing(&_Listing.CallOpts, hash)
}

// WasListing is a free data retrieval call binding the contract method 0x82c3eb3e.
//
// Solidity: function wasListing(hash bytes32) constant returns(out bool)
func (_Listing *ListingCallerSession) WasListing(hash [32]byte) (bool, error) {
	return _Listing.Contract.WasListing(&_Listing.CallOpts, hash)
}

// Challenge is a paid mutator transaction binding the contract method 0xcd32dcc5.
//
// Solidity: function challenge(chall string, hash bytes32) returns()
func (_Listing *ListingTransactor) Challenge(opts *bind.TransactOpts, chall string, hash [32]byte) (*types.Transaction, error) {
	return _Listing.contract.Transact(opts, "challenge", chall, hash)
}

// Challenge is a paid mutator transaction binding the contract method 0xcd32dcc5.
//
// Solidity: function challenge(chall string, hash bytes32) returns()
func (_Listing *ListingSession) Challenge(chall string, hash [32]byte) (*types.Transaction, error) {
	return _Listing.Contract.Challenge(&_Listing.TransactOpts, chall, hash)
}

// Challenge is a paid mutator transaction binding the contract method 0xcd32dcc5.
//
// Solidity: function challenge(chall string, hash bytes32) returns()
func (_Listing *ListingTransactorSession) Challenge(chall string, hash [32]byte) (*types.Transaction, error) {
	return _Listing.Contract.Challenge(&_Listing.TransactOpts, chall, hash)
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

// List is a paid mutator transaction binding the contract method 0x5148cd0f.
//
// Solidity: function list(listing string, data_hash bytes32, amount uint256) returns()
func (_Listing *ListingTransactor) List(opts *bind.TransactOpts, listing string, data_hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Listing.contract.Transact(opts, "list", listing, data_hash, amount)
}

// List is a paid mutator transaction binding the contract method 0x5148cd0f.
//
// Solidity: function list(listing string, data_hash bytes32, amount uint256) returns()
func (_Listing *ListingSession) List(listing string, data_hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Listing.Contract.List(&_Listing.TransactOpts, listing, data_hash, amount)
}

// List is a paid mutator transaction binding the contract method 0x5148cd0f.
//
// Solidity: function list(listing string, data_hash bytes32, amount uint256) returns()
func (_Listing *ListingTransactorSession) List(listing string, data_hash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Listing.Contract.List(&_Listing.TransactOpts, listing, data_hash, amount)
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
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterApplicationFailed is a free log retrieval operation binding the contract event 0x1e766ffad15a98a0dfb403123edc5f30e83e4981b1799b32a2cec29468fc84fd.
//
// Solidity: e ApplicationFailed(hash indexed bytes32)
func (_Listing *ListingFilterer) FilterApplicationFailed(opts *bind.FilterOpts, hash [][32]byte) (*ListingApplicationFailedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Listing.contract.FilterLogs(opts, "ApplicationFailed", hashRule)
	if err != nil {
		return nil, err
	}
	return &ListingApplicationFailedIterator{contract: _Listing.contract, event: "ApplicationFailed", logs: logs, sub: sub}, nil
}

// WatchApplicationFailed is a free log subscription operation binding the contract event 0x1e766ffad15a98a0dfb403123edc5f30e83e4981b1799b32a2cec29468fc84fd.
//
// Solidity: e ApplicationFailed(hash indexed bytes32)
func (_Listing *ListingFilterer) WatchApplicationFailed(opts *bind.WatchOpts, sink chan<- *ListingApplicationFailed, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Listing.contract.WatchLogs(opts, "ApplicationFailed", hashRule)
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
	DataHash  [32]byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterApplied is a free log retrieval operation binding the contract event 0x2be7f407f8117561d2a0100982c56a4a1ec8ffc24bbeb7b087b2ca8a224cc56d.
//
// Solidity: e Applied(hash indexed bytes32, applicant indexed address, dataHash bytes32, amount uint256)
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

// WatchApplied is a free log subscription operation binding the contract event 0x2be7f407f8117561d2a0100982c56a4a1ec8ffc24bbeb7b087b2ca8a224cc56d.
//
// Solidity: e Applied(hash indexed bytes32, applicant indexed address, dataHash bytes32, amount uint256)
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
	Reward     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallengeFailed is a free log retrieval operation binding the contract event 0xc7aab91fe7b50a881bd71df2cf07e9a6e5351721087cdb09bc70603f618e7057.
//
// Solidity: e ChallengeFailed(hash indexed bytes32, challenger indexed address, reward uint256)
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

// WatchChallengeFailed is a free log subscription operation binding the contract event 0xc7aab91fe7b50a881bd71df2cf07e9a6e5351721087cdb09bc70603f618e7057.
//
// Solidity: e ChallengeFailed(hash indexed bytes32, challenger indexed address, reward uint256)
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
	Reward     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallengeSucceeded is a free log retrieval operation binding the contract event 0xb2432a6d7681dc3f7416efc0e321723b96bd7a96d2d02d228dc2a67c51cf0322.
//
// Solidity: e ChallengeSucceeded(hash indexed bytes32, challenger indexed address, reward uint256)
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

// WatchChallengeSucceeded is a free log subscription operation binding the contract event 0xb2432a6d7681dc3f7416efc0e321723b96bd7a96d2d02d228dc2a67c51cf0322.
//
// Solidity: e ChallengeSucceeded(hash indexed bytes32, challenger indexed address, reward uint256)
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
