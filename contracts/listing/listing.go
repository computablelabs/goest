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
const ListingABI = "[{\"name\":\"ApplicationFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"applicant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Applied\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"applicant\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Challenged\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"challenger\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ChallengeFailed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"challenger\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"reward\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ChallengeSucceeded\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"challenger\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"reward\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Listed\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"reward\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingConverted\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingDeposit\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"deposited\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingRemoved\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ListingWithdraw\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\",\"indexed\":true},{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"withdrawn\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"market_token_addr\"},{\"type\":\"address\",\"name\":\"voting_addr\"},{\"type\":\"address\",\"name\":\"p11r_addr\"},{\"type\":\"address\",\"name\":\"data_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"isListed\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":697},{\"name\":\"isListingOwner\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":698},{\"name\":\"depositToListing\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":41294},{\"name\":\"withdrawFromListing\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"},{\"type\":\"uint256\",\"name\":\"amount\",\"unit\":\"wei\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":41103},{\"name\":\"list\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"listing\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":8550},{\"name\":\"getHash\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"string\",\"name\":\"str\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":663},{\"name\":\"getListing\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"},{\"type\":\"uint256\",\"name\":\"out\",\"unit\":\"wei\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1747},{\"name\":\"convertListing\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":118689},{\"name\":\"resolveApplication\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":132027},{\"name\":\"challenge\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":12626},{\"name\":\"resolveChallenge\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":232868},{\"name\":\"exit\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"hash\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":157875}]"

// ListingBin is the compiled bytecode used for deploying new contracts.
const ListingBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a0526080611d936101403934156100a757600080fd5b6020611d9360c03960c05160205181106100c057600080fd5b5060206020611d930160c03960c05160205181106100dd57600080fd5b5060206040611d930160c03960c05160205181106100fa57600080fd5b5060206060611d930160c03960c051602051811061011757600080fd5b506101405160025561016051600355610180516004556101a051600555611d7b56600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263ecefbdc660005114156100dc57602060046101403734156100b457600080fd5b600060006101405160e05260c052604060c02060c052602060c02054141560005260206000f3005b6302a47cd7600051141561012f57602060046101403734156100fd57600080fd5b600435602051811061010e57600080fd5b50600160016101405160e05260c052604060c02054101560005260206000f3005b63341ff089600051141561023f576040600461014037341561015057600080fd5b600060006101405160e05260c052604060c02060c052602060c02054141561017757600080fd5b6002543b61018457600080fd5b60025430141561019357600080fd5b602061024060646323b872dd61018052336101a052306101c052610160516101e05261019c60006002545af16101c857600080fd5b60005061024050600160006101405160e05260c052604060c02060c052602060c0200180546101605182540110156101ff57600080fd5b61016051815401815550610160516102605233610140517f6f27446a1e23c7e7dd81a6a77f81438be78145de331a5806f2c44e68738bdcef6020610260a3005b63a7e786756000511415610345576040600461014037341561026057600080fd5b3360006101405160e05260c052604060c02060c052602060c020541461028557600080fd5b600160006101405160e05260c052604060c02060c052602060c0200161016051815410156102b257600080fd5b610160518154038155506002543b6102c957600080fd5b6002543014156102d857600080fd5b6020610220604463a9059cbb61018052336101a052610160516101c05261019c60006002545af161030857600080fd5b60005061022050610160516102405233610140517f87c48f853a68e9374cf8389dab199f46adaffa48468fe9a4e2dc2edf75286db76020610240a3005b63fe42bf1a60005114156104d6576020600461014037341561036657600080fd5b606060043560040161016037604060043560040135111561038657600080fd5b6101608051602082012090506101e0526003543b6103a357600080fd5b6003543014156103b257600080fd5b6020610280602463b89694c6610200526101e0516102205261021c6003545afa6103db57600080fd5b60005061028051156103ec57600080fd5b60006101e05160e05260c052604060c02060c052602060c020541561041057600080fd5b6003543b61041d57600080fd5b60035430141561042c57600080fd5b60006000608463e8b969a4610320526101e0516103405260016103605233610380526004543b61045b57600080fd5b60045430141561046a57600080fd5b60206103006004632d0d2bc66102a0526102bc6004545afa61048b57600080fd5b600050610300516103a05261033c60006003545af16104a957600080fd5b336101e0517f26f7438941d95dc8495c99ef33fac91b2b04f98c73290027562cd3e612a3c85f60006000a3005b635b6beeb9600051141561052d57602060046101403734156104f757600080fd5b606060043560040161016037604060043560040135111561051757600080fd5b61016080516020820120905060005260206000f3005b63175c0d1660005114156105c2576020600461014037341561054e57600080fd5b60606101605261018060006101405160e05260c052604060c02060c052602060c020548152600160006101405160e05260c052604060c02060c052602060c02001548160200152600260006101405160e05260c052604060c02060c052602060c020015481604001525061016051610180f3005b63d748484d60005114156107e557602060046101403734156105e357600080fd5b3360006101405160e05260c052604060c02060c052602060c020541461060857600080fd5b600160006101405160e05260c052604060c02060c052602060c0200154600260006101405160e05260c052604060c02060c052602060c0200154600160006101405160e05260c052604060c02060c052602060c020015401101561066b57600080fd5b600260006101405160e05260c052604060c02060c052602060c0200154600160006101405160e05260c052604060c02060c052602060c020015401610160526000610160511115610760576000600160006101405160e05260c052604060c02060c052602060c02001556000600260006101405160e05260c052604060c02060c052602060c02001556002543b61070157600080fd5b60025430141561071057600080fd5b6020610220604463a9059cbb6101805260006101405160e05260c052604060c02060c052602060c020546101a052610160516101c05261019c60006002545af161075957600080fd5b6000506102205b503060006101405160e05260c052604060c02060c052602060c02055600060013360e05260c052604060c0205411156107b95760013360e05260c052604060c0206001815410156107b057600080fd5b60018154038155505b610140517ffd6a9ca1840ac9f8b7c592783e5415fdc282281630286c93b07b6d40f3c0ba5760006000a2005b600015610b8b575b6101605261014052600060006101405160e05260c052604060c02060c052602060c02054141561081c57600080fd5b6000600160006101405160e05260c052604060c02060c052602060c020015411156108e6576002543b61084e57600080fd5b60025430141561085d57600080fd5b6020610220604463a9059cbb6101805260006101405160e05260c052604060c02060c052602060c020546101a052600160006101405160e05260c052604060c02060c052602060c02001546101c05261019c60006002545af16108bf57600080fd5b600050610220506000600160006101405160e05260c052604060c02060c052602060c02001555b6000600260006101405160e05260c052604060c02060c052602060c0200154111561098a576002543b61091857600080fd5b60025430141561092757600080fd5b6000600060246342966c6861024052600260006101405160e05260c052604060c02060c052602060c02001546102605261025c60006002545af161096a57600080fd5b6000600260006101405160e05260c052604060c02060c052602060c02001555b6000600160006101405160e05260c052604060c02060c052602060c0205460e05260c052604060c0205411156109f957600160006101405160e05260c052604060c02060c052602060c0205460e05260c052604060c0206001815410156109f057600080fd5b60018154038155505b6001600160006101405160e05260c052604060c02060c052602060c0205460e05260c052604060c02054106003543b610a3157600080fd5b600354301415610a4057600080fd5b60206103e060246362f9a55d6103605260006101405160e05260c052604060c02060c052602060c020546103805261037c6003545afa610a7f57600080fd5b6000506103e0511615610ae9576003543b610a9957600080fd5b600354301415610aa857600080fd5b60006000602463335d80806104005260006101405160e05260c052604060c02060c052602060c020546104205261041c60006003545af1610ae857600080fd5b5b60006101405160e05260c052604060c02060c052602060c020600081556000600182015560006002820155506005543b610b2257600080fd5b600554301415610b3157600080fd5b60006000602463bd82badc61048052610140516104a05261049c60006005545af1610b5b57600080fd5b610140517f50425cae216bd151d26c8e8bb9779cc899c99d72f78081f2ceec9a99f001ff7960006000a261016051565b6318e28e1c60005114156110b25760206004610140373415610bac57600080fd5b6003543b610bb957600080fd5b600354301415610bc857600080fd5b60206101e060246362f9a55d61016052336101805261017c6003545afa610bee57600080fd5b6000506101e051610bfe57600080fd5b6003543b610c0b57600080fd5b600354301415610c1a57600080fd5b60206102a0604463af61f76061020052610140516102205260016102405261021c6003545afa610c4957600080fd5b6000506102a051610c5957600080fd5b6003543b610c6657600080fd5b600354301415610c7557600080fd5b6020610340602463327322c86102c052610140516102e0526102dc6003545afa610c9e57600080fd5b60005061034051610cae57600080fd5b6005543b610cbb57600080fd5b600554301415610cca57600080fd5b60206104006024637a639f6e61038052610140516103a05261039c6005545afa610cf357600080fd5b60005061040051610360526003543b610d0b57600080fd5b600354301415610d1a57600080fd5b60206104c0602463eb3014fd61044052610140516104605261045c6003545afa610d4357600080fd5b6000506104c0516104205260006103605114151561103b576003543b610d6857600080fd5b600354301415610d7757600080fd5b60206106806044638f354b796105e05261014051610600526004543b610d9c57600080fd5b600454301415610dab57600080fd5b60206105c0600463c26c12eb6105605261057c6004545afa610dcc57600080fd5b6000506105c051610620526105fc6003545afa610de857600080fd5b6000506106805115610fc1576104205160006101405160e05260c052604060c02060c052602060c020556004543b610e1f57600080fd5b600454301415610e2e57600080fd5b6020610720600463dc2b28536106c0526106dc6004545afa610e4f57600080fd5b600050610720516106a0526002543b610e6757600080fd5b600254301415610e7657600080fd5b60006000602463a0712d68610740526106a0516107605261075c60006002545af1610ea057600080fd5b6106a051600260006101405160e05260c052604060c02060c052602060c02001556003543b610ece57600080fd5b600354301415610edd57600080fd5b602061084060246362f9a55d6107c052610420516107e0526107dc6003545afa610f0657600080fd5b600050610840511515610f5a576003543b610f2057600080fd5b600354301415610f2f57600080fd5b60006000602463471bb30961086052610420516108805261087c60006003545af1610f5957600080fd5b5b6106a0516108e05261042051610140517f7f1a2c3e2883554425dc0f9f24dcfcdf54213b186c550080373dcc65813aa8d060206108e0a360016104205160e05260c052604060c020805460018254011015610fb457600080fd5b6001815401815550611036565b6005543b610fce57600080fd5b600554301415610fdd57600080fd5b60006000602463bd82badc6104e05261014051610500526104fc60006005545af161100757600080fd5b61042051610140517f163b66974440f1696d484dfc777282273c42ff2fa247d72dbdbb3510d92cbda660006000a35b61106a565b61042051610140517f163b66974440f1696d484dfc777282273c42ff2fa247d72dbdbb3510d92cbda660006000a35b6003543b61107757600080fd5b60035430141561108657600080fd5b6000600060246389bb617c61090052610140516109205261091c60006003545af16110b057600080fd5b005b63cffd46dc60005114156112b657602060046101403734156110d357600080fd5b600060006101405160e05260c052604060c02060c052602060c0205414156110fa57600080fd5b6004543b61110757600080fd5b60045430141561111657600080fd5b60206101e06004635aebe7686101805261019c6004545afa61113757600080fd5b6000506101e051610160526003543b61114f57600080fd5b60035430141561115e57600080fd5b6020610280602463b89694c661020052610140516102205261021c6003545afa61118757600080fd5b600050610280511561119857600080fd5b6002543b6111a557600080fd5b6002543014156111b457600080fd5b602061036060646323b872dd6102a052336102c052306102e05261016051610300526102bc60006002545af16111e957600080fd5b600050610360506003543b6111fd57600080fd5b60035430141561120c57600080fd5b60006000608463e8b969a461040052610140516104205260026104405233610460526004543b61123b57600080fd5b60045430141561124a57600080fd5b60206103e06004632d0d2bc66103805261039c6004545afa61126b57600080fd5b6000506103e0516104805261041c60006003545af161128957600080fd5b33610140517fe9479421670c3425a1497ce47a53af8bd96ce5bd0741e96221ba0acace3f7d4760006000a3005b63d32c943a6000511415611b4d57602060046101403734156112d757600080fd5b6003543b6112e457600080fd5b6003543014156112f357600080fd5b60206101e060246362f9a55d61016052336101805261017c6003545afa61131957600080fd5b6000506101e05161132957600080fd5b6003543b61133657600080fd5b60035430141561134557600080fd5b60206102a0604463af61f76061020052610140516102205260026102405261021c6003545afa61137457600080fd5b6000506102a05161138457600080fd5b6003543b61139157600080fd5b6003543014156113a057600080fd5b6020610340602463327322c86102c052610140516102e0526102dc6003545afa6113c957600080fd5b600050610340516113d957600080fd5b6003543b6113e657600080fd5b6003543014156113f557600080fd5b6020610400602463eb3014fd61038052610140516103a05261039c6003545afa61141e57600080fd5b60005061040051610360526004543b61143657600080fd5b60045430141561144557600080fd5b60206104a06004635aebe7686104405261045c6004545afa61146657600080fd5b6000506104a05161042052600160006101405160e05260c052604060c02060c052602060c02001546104c052600260006101405160e05260c052604060c02060c052602060c02001546104e052610420516104c0516104e0516104c0510110156114cf57600080fd5b6104e0516104c0510110156116cf5760006104c0516104e0516104c0510110156114f857600080fd5b6104e0516104c0510111156115b8576000600160006101405160e05260c052604060c02060c052602060c02001556000600260006101405160e05260c052604060c02060c052602060c02001556002543b61155257600080fd5b60025430141561156157600080fd5b60206108e0604463a9059cbb6108405261036051610860526104c0516104e0516104c05101101561159157600080fd5b6104e0516104c051016108805261085c60006002545af16115b157600080fd5b6000506108e05b506101405161016051610180516101a0516101c0516101e05161020051610220516102405161026051610280516102a0516102c0516102e05161030051610320516103405161036051610380516103a0516103c0516103e05161040051610420516104405161046051610480516104a0516104c0516104e05163550ee13361090052610140516109205261092051600658016107ed565b6104e0526104c0526104a05261048052610460526104405261042052610400526103e0526103c0526103a05261038052610360526103405261032052610300526102e0526102c0526102a05261028052610260526102405261022052610200526101e0526101c0526101a052610180526101605261014052600050611b4b565b6003543b6116dc57600080fd5b6003543014156116eb57600080fd5b60206106406044638f354b796105a052610140516105c0526004543b61171057600080fd5b60045430141561171f57600080fd5b6020610580600463c26c12eb6105205261053c6004545afa61174057600080fd5b600050610580516105e0526105bc6003545afa61175c57600080fd5b6000506106405115611a9257610420516104c0511015156117b357600160006101405160e05260c052604060c02060c052602060c0200161042051815410156117a457600080fd5b6104205181540381555061183a565b6000600160006101405160e05260c052604060c02060c052602060c0200155600260006101405160e05260c052604060c02060c052602060c020016104c05161042051101561180157600080fd5b6104c05161042051038154101561181757600080fd5b6104c05161042051101561182a57600080fd5b6104c05161042051038154038155505b6002543b61184757600080fd5b60025430141561185657600080fd5b6020610700604463a9059cbb61066052610360516106805261042051151561187f57600061189f565b60026104205160026104205102041461189757600080fd5b600261042051025b6106a05261067c60006002545af16118b657600080fd5b600050610700506101405161016051610180516101a0516101c0516101e05161020051610220516102405161026051610280516102a0516102c0516102e05161030051610320516103405161036051610380516103a0516103c0516103e05161040051610420516104405161046051610480516104a0516104c0516104e05161050051610520516105405161056051610580516105a0516105c0516105e05161060051610620516106405161066051610680516106a0516106c0516106e0516107005163550ee13361072052610140516107405261074051600658016107ed565b610700526106e0526106c0526106a05261068052610660526106405261062052610600526105e0526105c0526105a05261058052610560526105405261052052610500526104e0526104c0526104a05261048052610460526104405261042052610400526103e0526103c0526103a05261038052610360526103405261032052610300526102e0526102c0526102a05261028052610260526102405261022052610200526101e0526101c0526101a052610180526101605261014052600050610420516107a05261036051610140517fb2432a6d7681dc3f7416efc0e321723b96bd7a96d2d02d228dc2a67c51cf032260206107a0a3611b04565b600260006101405160e05260c052604060c02060c052602060c020018054610420518254011015611ac257600080fd5b61042051815401815550610420516105005261036051610140517fc7aab91fe7b50a881bd71df2cf07e9a6e5351721087cdb09bc70603f618e70576020610500a35b6003543b611b1157600080fd5b600354301415611b2057600080fd5b6000600060246389bb617c6107c052610140516107e0526107dc60006003545af1611b4a57600080fd5b5b005b630ca362636000511415611c3c5760206004610140373415611b6e57600080fd5b3360006101405160e05260c052604060c02060c052602060c0205414611b9357600080fd5b6003543b611ba057600080fd5b600354301415611baf57600080fd5b60206101e0602463b89694c661016052610140516101805261017c6003545afa611bd857600080fd5b6000506101e05115611be957600080fd5b6101405161016051610180516101a0516101c0516101e05163550ee13361020052610140516102205261022051600658016107ed565b6101e0526101c0526101a052610180526101605261014052600050005b60006000fd5b610139611d7b03610139600039610139611d7b036000f3`

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
