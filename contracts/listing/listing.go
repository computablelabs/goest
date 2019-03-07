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
const ListingABI = "[{\"name\":\"ApplicationFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Applied\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"applicant\",\"indexed\":true},{\"type\":\"bytes32\",\"name\":\"dataHash\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Challenged\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"challenger\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ChallengeFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"challenger\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"reward\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ChallengeSucceeded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"challenger\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"reward\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Listed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"reward\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingConverted\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingDeposit\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"deposited\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingRemoved\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingWithdraw\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"withdrawn\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"market_token_addr\"},{\"type\":\"address\",\"name\":\"voting_addr\"},{\"type\":\"address\",\"name\":\"p11r_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"isListed\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":703},{\"name\":\"isListing\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1033},{\"name\":\"wasListing\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":754},{\"name\":\"isListingOwner\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":758},{\"name\":\"depositToListing\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":41660},{\"name\":\"withdrawFromListing\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":41959},{\"name\":\"list\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"listing\"},{\"type\":\"bytes32\",\"name\":\"data_hash\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":224392},{\"name\":\"getListingCount\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":693},{\"name\":\"getListingKey\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"index\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":940},{\"name\":\"getListing\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"bytes32\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2251},{\"name\":\"convertListing\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":82759},{\"name\":\"resolveApplication\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":305172},{\"name\":\"challenge\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"chall\"},{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":12950},{\"name\":\"resolveChallenge\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":370547},{\"name\":\"exit\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":295560}]"

// ListingBin is the compiled bytecode used for deploying new contracts.
const ListingBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05260606120536101403934156100a757600080fd5b602061205360c03960c05160205181106100c057600080fd5b50602060206120530160c03960c05160205181106100dd57600080fd5b50602060406120530160c03960c05160205181106100fa57600080fd5b503360045561014051600555610160516006556101805160075561203b56600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263ecefbdc660005114156100df57602060046101403734156100b457600080fd5b6000600160006101405160e05260c052604060c02060c052602060c0200154141560005260206000f3005b636aa0966b6000511415610144576020600461014037341561010057600080fd5b6101405160006101405160e05260c052604060c02060c052602060c02054620f4240811061012d57600080fd5b600260c052602060c02001541460005260206000f3005b6382c3eb3e60005114156101ab576020600461014037341561016557600080fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60006101405160e05260c052604060c02060c052602060c020541460005260206000f3005b6302a47cd760005114156101fe57602060046101403734156101cc57600080fd5b60043560205181106101dd57600080fd5b50600160036101405160e05260c052604060c02054101560005260206000f3005b63341ff089600051141561032a576040600461014037341561021f57600080fd5b6101405160006101405160e05260c052604060c02060c052602060c02054620f4240811061024c57600080fd5b600260c052602060c02001541461026257600080fd5b6005543b61026f57600080fd5b60055430141561027e57600080fd5b602061024060646323b872dd61018052336101a052306101c052610160516101e05261019c60006005545af16102b357600080fd5b60005061024050600360006101405160e05260c052604060c02060c052602060c0200180546101605182540110156102ea57600080fd5b61016051815401815550610160516102605233610140517f6f27446a1e23c7e7dd81a6a77f81438be78145de331a5806f2c44e68738bdcef6020610260a3005b63a7e786756000511415610476576040600461014037341561034b57600080fd5b6101405160006101405160e05260c052604060c02060c052602060c02054620f4240811061037857600080fd5b600260c052602060c02001541461038e57600080fd5b33600160006101405160e05260c052604060c02060c052602060c0200154146103b657600080fd5b600360006101405160e05260c052604060c02060c052602060c0200161016051815410156103e357600080fd5b610160518154038155506005543b6103fa57600080fd5b60055430141561040957600080fd5b6020610220604463a9059cbb61018052336101a052610160516101c05261019c60006005545af161043957600080fd5b60005061022050610160516102405233610140517f87c48f853a68e9374cf8389dab199f46adaffa48468fe9a4e2dc2edf75286db76020610240a3005b635148cd0f600051141561073d576060600461014037341561049757600080fd5b60606004356004016101a03760406004356004013511156104b757600080fd5b620f4240600154106104c857600080fd5b6101a0805160208201209050610220526006543b6104e557600080fd5b6006543014156104f457600080fd5b60206102c06024631f1b029961024052610220516102605261025c6006545afa61051d57600080fd5b6000506102c05161052d57600080fd5b6006543b61053a57600080fd5b60065430141561054957600080fd5b60006000608463e8b969a461036052610220516103805260016103a052336103c0526007543b61057857600080fd5b60075430141561058757600080fd5b60206103406004632d0d2bc66102e0526102fc6007545afa6105a857600080fd5b600050610340516103e05261037c60006006545af16105c657600080fd5b61016051600260006102205160e05260c052604060c02060c052602060c020015560015460006102205160e05260c052604060c02060c052602060c0205561022051600154620f4240811061061a57600080fd5b600260c052602060c0200155600160605160018254018060405190131561064057600080fd5b809190121561064e57600080fd5b81555060033360e05260c052604060c02080546001825401101561067157600080fd5b600181540181555060006101805111156106ff576005543b61069257600080fd5b6005543014156106a157600080fd5b602061050060646323b872dd6104405233610460523061048052610180516104a05261045c60006005545af16106d657600080fd5b6000506105005061018051600360006102205160e05260c052604060c02060c052602060c02001555b6101605161052052610180516105405233610220517f2be7f407f8117561d2a0100982c56a4a1ec8ffc24bbeb7b087b2ca8a224cc56d6040610520a3005b6387ed92d7600051141561076357341561075657600080fd5b60015460005260206000f3005b63c48dc1af60005114156107d2576020600461014037341561078457600080fd5b6060516004358060405190131561079a57600080fd5b80919012156107a857600080fd5b5061014051620f424081106107bc57600080fd5b600260c052602060c020015460005260206000f3005b63175c0d16600051141561088c57602060046101403734156107f357600080fd5b608061016052610180600160006101405160e05260c052604060c02060c052602060c02001548152600260006101405160e05260c052604060c02060c052602060c02001548160200152600360006101405160e05260c052604060c02060c052602060c02001548160400152600460006101405160e05260c052604060c02060c052602060c020015481606001525061016051610180f3005b63d748484d6000511415610a7b57602060046101403734156108ad57600080fd5b33600160006101405160e05260c052604060c02060c052602060c0200154146108d557600080fd5b600360006101405160e05260c052604060c02060c052602060c0200154600460006101405160e05260c052604060c02060c052602060c0200154600360006101405160e05260c052604060c02060c052602060c020015401101561093857600080fd5b600460006101405160e05260c052604060c02060c052602060c0200154600360006101405160e05260c052604060c02060c052602060c020015401610160526000610160511115610a30576000600360006101405160e05260c052604060c02060c052602060c02001556000600460006101405160e05260c052604060c02060c052602060c02001556005543b6109ce57600080fd5b6005543014156109dd57600080fd5b6020610220604463a9059cbb61018052600160006101405160e05260c052604060c02060c052602060c02001546101a052610160516101c05261019c60006005545af1610a2957600080fd5b6000506102205b5030600160006101405160e05260c052604060c02060c052602060c0200155610140517ffd6a9ca1840ac9f8b7c592783e5415fdc282281630286c93b07b6d40f3c0ba5760006000a2005b600015610e94575b61016052610140526101405160006101405160e05260c052604060c02060c052602060c02054620f42408110610ab857600080fd5b600260c052602060c020015414610ace57600080fd5b6000600360006101405160e05260c052604060c02060c052602060c02001541115610b9b576005543b610b0057600080fd5b600554301415610b0f57600080fd5b6020610220604463a9059cbb61018052600160006101405160e05260c052604060c02060c052602060c02001546101a052600360006101405160e05260c052604060c02060c052602060c02001546101c05261019c60006005545af1610b7457600080fd5b600050610220506000600360006101405160e05260c052604060c02060c052602060c02001555b6000600460006101405160e05260c052604060c02060c052602060c02001541115610c3f576005543b610bcd57600080fd5b600554301415610bdc57600080fd5b6000600060246342966c6861024052600460006101405160e05260c052604060c02060c052602060c02001546102605261025c60006005545af1610c1f57600080fd5b6000600460006101405160e05260c052604060c02060c052602060c02001555b6001606051600182540380604051901315610c5957600080fd5b8091901215610c6757600080fd5b81555060006001541315610cf75760006101405160e05260c052604060c02060c052602060c020546102c052600154620f42408110610ca557600080fd5b600260c052602060c02001546102e0526102e0516102c051620f42408110610ccc57600080fd5b600260c052602060c02001556102c05160006102e05160e05260c052604060c02060c052602060c020555b6000600154620f42408110610d0b57600080fd5b600260c052602060c02001556003600160006101405160e05260c052604060c02060c052602060c020015460e05260c052604060c020600181541015610d5057600080fd5b600181540381555060016003600160006101405160e05260c052604060c02060c052602060c020015460e05260c052604060c020541015610deb576006543b610d9857600080fd5b600654301415610da757600080fd5b60006000602463335d808061030052600160006101405160e05260c052604060c02060c052602060c02001546103205261031c60006006545af1610dea57600080fd5b5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60006101405160e05260c052604060c02060c052602060c020556000600160006101405160e05260c052604060c02060c052602060c02001556000600260006101405160e05260c052604060c02060c052602060c0200155610140517f50425cae216bd151d26c8e8bb9779cc899c99d72f78081f2ceec9a99f001ff7960006000a261016051565b6318e28e1c600051141561135d5760206004610140373415610eb557600080fd5b6006543b610ec257600080fd5b600654301415610ed157600080fd5b60206101e060246362f9a55d61016052336101805261017c6006545afa610ef757600080fd5b6000506101e051610f0757600080fd5b6006543b610f1457600080fd5b600654301415610f2357600080fd5b60206102a0604463af61f76061020052610140516102205260016102405261021c6006545afa610f5257600080fd5b6000506102a051610f6257600080fd5b600160006101405160e05260c052604060c02060c052602060c020015415610f8957600080fd5b6006543b610f9657600080fd5b600654301415610fa557600080fd5b6020610340602463327322c86102c052610140516102e0526102dc6006545afa610fce57600080fd5b60005061034051610fde57600080fd5b6006543b610feb57600080fd5b600654301415610ffa57600080fd5b60206105006044638f354b796104605261014051610480526007543b61101f57600080fd5b60075430141561102e57600080fd5b6020610440600463c26c12eb6103e0526103fc6007545afa61104f57600080fd5b600050610440516104a05261047c6006545afa61106b57600080fd5b600050610500511561126b576006543b61108457600080fd5b60065430141561109357600080fd5b60206105c0602463eb3014fd61054052610140516105605261055c6006545afa6110bc57600080fd5b6000506105c0516105205261052051600160006101405160e05260c052604060c02060c052602060c02001556007543b6110f557600080fd5b60075430141561110457600080fd5b6020610660600463dc2b28536106005261061c6007545afa61112557600080fd5b600050610660516105e0526005543b61113d57600080fd5b60055430141561114c57600080fd5b60006000602463a0712d68610680526105e0516106a05261069c60006005545af161117657600080fd5b6105e051600460006101405160e05260c052604060c02060c052602060c02001556006543b6111a457600080fd5b6006543014156111b357600080fd5b6020610780602463c666b58061070052610520516107205261071c6006545afa6111dc57600080fd5b600050610780511561122f576006543b6111f557600080fd5b60065430141561120457600080fd5b60006000602463471bb3096107a052610520516107c0526107bc60006006545af161122e57600080fd5b5b6105e0516108205261052051610140517f7f1a2c3e2883554425dc0f9f24dcfcdf54213b186c550080373dcc65813aa8d06020610820a3611315565b6101405161016051610180516101a0516101c0516101e05161020051610220516102405161026051610280516102a0516102c0516102e05161030051610320516103405163550ee1336103605261014051610380526103805160065801610a83565b6103405261032052610300526102e0526102c0526102a05261028052610260526102405261022052610200526101e0526101c0526101a0526101805261016052610140526000505b6006543b61132257600080fd5b60065430141561133157600080fd5b6000600060246389bb617c61084052610140516108605261085c60006006545af161135b57600080fd5b005b63cd32dcc56000511415611593576040600461014037341561137e57600080fd5b606060043560040161018037604060043560040135111561139e57600080fd5b6000600160006101605160e05260c052604060c02060c052602060c020015414156113c857600080fd5b6007543b6113d557600080fd5b6007543014156113e457600080fd5b60206102806004635aebe7686102205261023c6007545afa61140557600080fd5b60005061028051610200526101808051602082012090506102a0526006543b61142d57600080fd5b60065430141561143c57600080fd5b60206103406024631f1b02996102c0526102a0516102e0526102dc6006545afa61146557600080fd5b6000506103405161147557600080fd5b6005543b61148257600080fd5b60055430141561149157600080fd5b602061042060646323b872dd610360523361038052306103a052610200516103c05261037c60006005545af16114c657600080fd5b600050610420506006543b6114da57600080fd5b6006543014156114e957600080fd5b60006000608463e8b969a46104c052610160516104e05260026105005233610520526007543b61151857600080fd5b60075430141561152757600080fd5b60206104a06004632d0d2bc66104405261045c6007545afa61154857600080fd5b6000506104a051610540526104dc60006006545af161156657600080fd5b33610160517fe9479421670c3425a1497ce47a53af8bd96ce5bd0741e96221ba0acace3f7d4760006000a3005b63d32c943a6000511415611e2a57602060046101403734156115b457600080fd5b6006543b6115c157600080fd5b6006543014156115d057600080fd5b60206101e060246362f9a55d61016052336101805261017c6006545afa6115f657600080fd5b6000506101e05161160657600080fd5b6006543b61161357600080fd5b60065430141561162257600080fd5b60206102a0604463af61f76061020052610140516102205260026102405261021c6006545afa61165157600080fd5b6000506102a05161166157600080fd5b6006543b61166e57600080fd5b60065430141561167d57600080fd5b6020610340602463327322c86102c052610140516102e0526102dc6006545afa6116a657600080fd5b600050610340516116b657600080fd5b6006543b6116c357600080fd5b6006543014156116d257600080fd5b6020610400602463eb3014fd61038052610140516103a05261039c6006545afa6116fb57600080fd5b60005061040051610360526007543b61171357600080fd5b60075430141561172257600080fd5b60206104a06004635aebe7686104405261045c6007545afa61174357600080fd5b6000506104a05161042052600360006101405160e05260c052604060c02060c052602060c02001546104c052600460006101405160e05260c052604060c02060c052602060c02001546104e052610420516104c0516104e0516104c0510110156117ac57600080fd5b6104e0516104c0510110156119ac5760006104c0516104e0516104c0510110156117d557600080fd5b6104e0516104c051011115611895576000600360006101405160e05260c052604060c02060c052602060c02001556000600460006101405160e05260c052604060c02060c052602060c02001556005543b61182f57600080fd5b60055430141561183e57600080fd5b60206108e0604463a9059cbb6108405261036051610860526104c0516104e0516104c05101101561186e57600080fd5b6104e0516104c051016108805261085c60006005545af161188e57600080fd5b6000506108e05b506101405161016051610180516101a0516101c0516101e05161020051610220516102405161026051610280516102a0516102c0516102e05161030051610320516103405161036051610380516103a0516103c0516103e05161040051610420516104405161046051610480516104a0516104c0516104e05163550ee1336109005261014051610920526109205160065801610a83565b6104e0526104c0526104a05261048052610460526104405261042052610400526103e0526103c0526103a05261038052610360526103405261032052610300526102e0526102c0526102a05261028052610260526102405261022052610200526101e0526101c0526101a052610180526101605261014052600050611e28565b6006543b6119b957600080fd5b6006543014156119c857600080fd5b60206106406044638f354b796105a052610140516105c0526007543b6119ed57600080fd5b6007543014156119fc57600080fd5b6020610580600463c26c12eb6105205261053c6007545afa611a1d57600080fd5b600050610580516105e0526105bc6006545afa611a3957600080fd5b6000506106405115611d6f57610420516104c051101515611a9057600360006101405160e05260c052604060c02060c052602060c020016104205181541015611a8157600080fd5b61042051815403815550611b17565b6000600360006101405160e05260c052604060c02060c052602060c0200155600460006101405160e05260c052604060c02060c052602060c020016104c051610420511015611ade57600080fd5b6104c051610420510381541015611af457600080fd5b6104c051610420511015611b0757600080fd5b6104c05161042051038154038155505b6005543b611b2457600080fd5b600554301415611b3357600080fd5b6020610700604463a9059cbb610660526103605161068052610420511515611b5c576000611b7c565b600261042051600261042051020414611b7457600080fd5b600261042051025b6106a05261067c60006005545af1611b9357600080fd5b600050610700506101405161016051610180516101a0516101c0516101e05161020051610220516102405161026051610280516102a0516102c0516102e05161030051610320516103405161036051610380516103a0516103c0516103e05161040051610420516104405161046051610480516104a0516104c0516104e05161050051610520516105405161056051610580516105a0516105c0516105e05161060051610620516106405161066051610680516106a0516106c0516106e0516107005163550ee1336107205261014051610740526107405160065801610a83565b610700526106e0526106c0526106a05261068052610660526106405261062052610600526105e0526105c0526105a05261058052610560526105405261052052610500526104e0526104c0526104a05261048052610460526104405261042052610400526103e0526103c0526103a05261038052610360526103405261032052610300526102e0526102c0526102a05261028052610260526102405261022052610200526101e0526101c0526101a052610180526101605261014052600050610420516107a05261036051610140517fb2432a6d7681dc3f7416efc0e321723b96bd7a96d2d02d228dc2a67c51cf032260206107a0a3611de1565b600460006101405160e05260c052604060c02060c052602060c020018054610420518254011015611d9f57600080fd5b61042051815401815550610420516105005261036051610140517fc7aab91fe7b50a881bd71df2cf07e9a6e5351721087cdb09bc70603f618e70576020610500a35b6006543b611dee57600080fd5b600654301415611dfd57600080fd5b6000600060246389bb617c6107c052610140516107e0526107dc60006006545af1611e2757600080fd5b5b005b630ca362636000511415611f1c5760206004610140373415611e4b57600080fd5b33600160006101405160e05260c052604060c02060c052602060c020015414611e7357600080fd5b6006543b611e8057600080fd5b600654301415611e8f57600080fd5b60206101e0602463b89694c661016052610140516101805261017c6006545afa611eb857600080fd5b6000506101e05115611ec957600080fd5b6101405161016051610180516101a0516101c0516101e05163550ee1336102005261014051610220526102205160065801610a83565b6101e0526101c0526101a052610180526101605261014052600050005b60006000fd5b61011961203b0361011960003961011961203b036000f3`

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

// GetListing is a free data retrieval call binding the contract method 0x175c0d16.
//
// Solidity: function getListing(hash bytes32) constant returns(out address, out bytes32, out uint256, out uint256)
func (_Listing *ListingCaller) GetListing(opts *bind.CallOpts, hash [32]byte) (common.Address, [32]byte, *big.Int, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new([32]byte)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Listing.contract.Call(opts, out, "getListing", hash)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetListing is a free data retrieval call binding the contract method 0x175c0d16.
//
// Solidity: function getListing(hash bytes32) constant returns(out address, out bytes32, out uint256, out uint256)
func (_Listing *ListingSession) GetListing(hash [32]byte) (common.Address, [32]byte, *big.Int, *big.Int, error) {
	return _Listing.Contract.GetListing(&_Listing.CallOpts, hash)
}

// GetListing is a free data retrieval call binding the contract method 0x175c0d16.
//
// Solidity: function getListing(hash bytes32) constant returns(out address, out bytes32, out uint256, out uint256)
func (_Listing *ListingCallerSession) GetListing(hash [32]byte) (common.Address, [32]byte, *big.Int, *big.Int, error) {
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
