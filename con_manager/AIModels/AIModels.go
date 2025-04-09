// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AIModels

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AIModelsMetaData contains all meta data concerning the AIModels contract.
var AIModelsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"ModelDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"ModelRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelInfo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"UploadModeled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"getModelDistribution\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getNodeDeployment\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"modelDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"modelIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"modelOwns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextModelId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeDeployment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelExtendInfo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"recordModelUpload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractNodesRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"removeDeployment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"reportDeployment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIStake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uploadModels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"extendInfo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60803461013157601f61141a38819003918201601f19168301916001600160401b0383118484101761013657808492604094855283398101031261013157610052602061004b8361014c565b920161014c565b60016004556001600160a01b03918216919082156100e0571690811561009b5760018060a01b03199081600054161760005560055416176005556040516112b990816101618239f35b60405162461bcd60e51b815260206004820152601360248201527f496e76616c6964207374616b6520746f6b656e000000000000000000000000006044820152606490fd5b60405162461bcd60e51b8152602060048201526024808201527f496e76616c6964207175616e74697479206f66207265676973747279206164646044820152637265737360e01b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101315756fe608080604052600436101561001357600080fd5b60003560e01c9081632ef6453714610f615750806351ed6a3014610f3857806357793e6814610e8057806364290cbf14610b8357806369f4296a14610b495780637b10399914610b205780637d5ef48e14610ac8578063907d0042146108dc5780639bb0d61214610828578063c1ce718514610169578063c78717d514610123578063dcca1c92146100d15763e472ae8b146100ae57600080fd5b346100cc5760003660031901126100cc576020600454604051908152f35b600080fd5b346100cc5760403660031901126100cc576024356004356000526006602052604060002080548210156100cc5760209161010a916110f2565b905460405160039290921b1c6001600160a01b03168152f35b346100cc57602061013336611090565b8161014760405192838151938492016110cf565b600390820190815281900382019020546040516001600160a01b039091168152f35b346100cc5760803660031901126100cc5760043567ffffffffffffffff81116100cc5761019a903690600401611225565b9060243567ffffffffffffffff81116100cc576101bb903690600401611225565b92909160443567ffffffffffffffff81116100cc576101de903690600401611225565b90936101eb368585611049565b9561024460216101fc368486611049565b9860405199816102168c93518092602080870191016110cf565b8201602f60f81b602082015261023582518093602087850191016110cf565b0103600181018a520188611027565b6040516020818161025b8b838151938492016110cf565b81016001815203019020546107f55760045496604051978860e081011067ffffffffffffffff60e08b01111761063f5760e089016040528089526102a0368888611049565b60208a01526102b0368486611049565b60408a01523360608a01526102c636868a611049565b60808a01524260a08a015260643560c08a01526000526002602052604060002088518155602089015180519067ffffffffffffffff821161063f5781906103106001850154611120565b601f81116107a2575b50602090601f831160011461073057600092610725575b50508160011b916000199060031b1c19161760018201555b604089015180519067ffffffffffffffff821161063f57819061036e6002850154611120565b601f81116106d2575b50602090601f831160011461066057600092610655575b50508160011b916000199060031b1c19161760028201555b6003810160018060a01b0360608b0151166bffffffffffffffffffffffff60a01b825416179055608089015198895167ffffffffffffffff811161063f576103f16004840154611120565b9a601f8c116105f4575b60209b508b90601f83116001146105515793600660c06104f39995856104d89f9a967fa17734f4a6749e7cc4947af9f12fb87299c2c2ab175158d74d24c8287fdee4369d9a966104e59a600092610546575b50508160011b916000199060031b1c19161760048501555b60a08101516005850155015191015561048f8c6004549281604051938285809451938492016110cf565b8101600181520301902055604051898982376003818b019081528190038c019020546001600160a01b031615610516575b600454998a986040519a60808c5260808c0191611253565b918983038c8b0152611253565b918683036040880152611253565b9260643560608201528033940390a361050b81611274565b600455604051908152f35b604051898982378b818b81016003815203019020336bffffffffffffffffffffffff60a01b8254161790556104c0565b01519050388061044d565b90600485016000528c6000209160005b601f19851681106105dd575060c06104f3999560016104d89f9a967fa17734f4a6749e7cc4947af9f12fb87299c2c2ab175158d74d24c8287fdee4369d9a966104e59a9660069683601f198116106105c4575b505050811b016004850155610465565b015160001960f88460031b161c191690553880806105b4565b91928e600181928685015181550194019201610561565b600484016000526020600020601f830160051c810160208410610638575b601f8e5b0160051c820181106106295750506103fb565b60008155600101601f8e610616565b5080610612565b634e487b7160e01b600052604160045260246000fd5b015190508b8061038e565b9250600284016000526020600020906000935b601f19841685106106b7576001945083601f1981161061069e575b505050811b0160028201556103a6565b015160001960f88460031b161c191690558b808061068e565b81810151835560209485019460019093019290910190610673565b909150600284016000526020600020601f840160051c81016020851061071e575b90849392915b601f830160051c8201811061070f575050610377565b600081558594506001016106f9565b50806106f3565b015190508b80610330565b9250600184016000526020600020906000935b601f1984168510610787576001945083601f1981161061076e575b505050811b016001820155610348565b015160001960f88460031b161c191690558b808061075e565b81810151835560209485019460019093019290910190610743565b909150600184016000526020600020601f840160051c8101602085106107ee575b90849392915b601f830160051c820181106107df575050610319565b600081558594506001016107c9565b50806107c3565b60405162461bcd60e51b815260206004820152600b60248201526a135bd9195b08195e1a5cdd60aa1b6044820152606490fd5b346100cc5760203660031901126100cc576004356000526002602052604060002080546108576001830161115a565b916108646002820161115a565b60038201546108cd906001600160a01b03166108826004850161115a565b6108ba60066005870154960154946108ac604051998a998a5260e060208b015260e08a0190611200565b9088820360408a0152611200565b9160608701528582036080870152611200565b9160a084015260c08301520390f35b346100cc5760203660031901126100cc5760043560008181526006602052604081205b80549182811015610ac05761091481836110f2565b90546003946001600160a01b0392909190861b1c8216331461094357505061093d919250611274565b906108ff565b90916000199180830191908211610aa1576109738461096561099294886110f2565b905490891b1c1691866110f2565b90919082549060031b9160018060a01b03809116831b921b1916179055565b82548015610a8b5701926109a684846110f2565b81939154921b1b19169055555b33600052600760205260406000208160005b825480821015610ab75782906109db83866110f2565b929054600393841b1c146109f95750506109f490611274565b6109c5565b9093925060001991828201918211610aa157610a28610a1b610a4193866110f2565b905490871b1c91856110f2565b90919082549060031b91821b91600019901b1916179055565b81548015610a8b57810192610a5684846110f2565b81939154921b1b19169055555b337f271097aa25aca38eb083e68e5cafcbd8247d421aea1887c9630402cda1706b63600080a3005b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b50505050610a63565b5050506109b3565b346100cc5760403660031901126100cc57610ae1611011565b6001600160a01b03166000908152600760205260409020805460243591908210156100cc57602091610b12916110f2565b90546040519160031b1c8152f35b346100cc5760003660031901126100cc576000546040516001600160a01b039091168152602090f35b346100cc576020610b7081610b5d36611090565b81604051938285809451938492016110cf565b8101600181520301902054604051908152f35b346100cc576020806003193601126100cc57600435908060018060a01b03806000541660405192838092631846d2f560e31b825233600483015260249586915afa908115610e7457600091610e3e575b5015610e0557836000526002835260406000205415610dcc5760008481526006845260408120805492915b838110610d62575050505082600052600682526040600020918254680100000000000000009384821015610d4d5781610c3f916001610c5d940181556110f2565b81546001600160a01b0360039290921b91821b19163390911b179055565b336000526007815260406000206000908054915b828110610ce4575050506007903360005252604060002090815492831015610cd05750610a288284926001610ca8950181556110f2565b337fb81e8e03e6a54d2b07bc11d3cddad332ded9a7018c0e554bd78262a1d4cd20d1600080a3005b634e487b7160e01b60009081526041600452fd5b86610cef82846110f2565b90549060031b1c14610d0957610d0490611274565b610c71565b60405162461bcd60e51b815260048101859052601d818701527f4e6f6465206465706c6f796d656e7420616c72656164792065786973740000006044820152606490fd5b83634e487b7160e01b60005260416004526000fd5b82610d6d82846110f2565b919054339260031b1c1614610d8a57610d8590611274565b610bfe565b606485876040519162461bcd60e51b83528160048401528201527f4d6f64656c20646973747269627574696f6e20616c72656164792065786973746044820152fd5b60405162461bcd60e51b81526004810184905260128184015271135bd9195b081a5cc81b9bdd08195e1a5cdd60721b6044820152606490fd5b60405162461bcd60e51b815260048101849052601281840152714e6f6465206973206e6f742061637469766560701b6044820152606490fd5b90508381813d8311610e6d575b610e558183611027565b810103126100cc575180151581036100cc5785610bd3565b503d610e4b565b6040513d6000823e3d90fd5b346100cc576020806003193601126100cc5760043560005260068152604060002060405190818382549182815201908192600052846000209060005b86828210610f1b578686610ed282880383611027565b604051928392818401908285525180915260408401929160005b828110610efb57505050500390f35b83516001600160a01b031685528695509381019392810192600101610eec565b83546001600160a01b031685529093019260019283019201610ebc565b346100cc5760003660031901126100cc576005546040516001600160a01b039091168152602090f35b346100cc57602090816003193601126100cc576001600160a01b03610f84611011565b166000526007825260406000208054808352838301908192600052846000209060005b86828210610ffd578686610fbd82880383611027565b604051928392818401908285525180915260408401929160005b828110610fe657505050500390f35b835185528695509381019392810192600101610fd7565b835485529093019260019283019201610fa7565b600435906001600160a01b03821682036100cc57565b90601f8019910116810190811067ffffffffffffffff82111761063f57604052565b92919267ffffffffffffffff821161063f5760405191611073601f8201601f191660200184611027565b8294818452818301116100cc578281602093846000960137010152565b60206003198201126100cc576004359067ffffffffffffffff82116100cc57806023830112156100cc578160246110cc93600401359101611049565b90565b60005b8381106110e25750506000910152565b81810151838201526020016110d2565b805482101561110a5760005260206000200190600090565b634e487b7160e01b600052603260045260246000fd5b90600182811c92168015611150575b602083101461113a57565b634e487b7160e01b600052602260045260246000fd5b91607f169161112f565b906040519182600082549261116e84611120565b9081845260019485811690816000146111dd575060011461119a575b505061119892500383611027565b565b9093915060005260209081600020936000915b8183106111c55750506111989350820101388061118a565b855488840185015294850194879450918301916111ad565b91505061119894506020925060ff191682840152151560051b820101388061118a565b90602091611219815180928185528580860191016110cf565b601f01601f1916010190565b9181601f840112156100cc5782359167ffffffffffffffff83116100cc57602083818601950101116100cc57565b908060209392818452848401376000828201840152601f01601f1916010190565b6000198114610aa1576001019056fea2646970667358221220901eac38f56634372bef3342e6837af139daaff6e3f1dbaa72da9e6264822c0064736f6c63430008140033",
}

// AIModelsABI is the input ABI used to generate the binding from.
// Deprecated: Use AIModelsMetaData.ABI instead.
var AIModelsABI = AIModelsMetaData.ABI

// AIModelsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AIModelsMetaData.Bin instead.
var AIModelsBin = AIModelsMetaData.Bin

// DeployAIModels deploys a new Ethereum contract, binding an instance of AIModels to it.
func DeployAIModels(auth *bind.TransactOpts, backend bind.ContractBackend, _registry common.Address, _stakeToken common.Address) (common.Address, *types.Transaction, *AIModels, error) {
	parsed, err := AIModelsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AIModelsBin), backend, _registry, _stakeToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AIModels{AIModelsCaller: AIModelsCaller{contract: contract}, AIModelsTransactor: AIModelsTransactor{contract: contract}, AIModelsFilterer: AIModelsFilterer{contract: contract}}, nil
}

// AIModels is an auto generated Go binding around an Ethereum contract.
type AIModels struct {
	AIModelsCaller     // Read-only binding to the contract
	AIModelsTransactor // Write-only binding to the contract
	AIModelsFilterer   // Log filterer for contract events
}

// AIModelsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AIModelsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIModelsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AIModelsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIModelsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AIModelsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIModelsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AIModelsSession struct {
	Contract     *AIModels         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AIModelsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AIModelsCallerSession struct {
	Contract *AIModelsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AIModelsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AIModelsTransactorSession struct {
	Contract     *AIModelsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AIModelsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AIModelsRaw struct {
	Contract *AIModels // Generic contract binding to access the raw methods on
}

// AIModelsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AIModelsCallerRaw struct {
	Contract *AIModelsCaller // Generic read-only contract binding to access the raw methods on
}

// AIModelsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AIModelsTransactorRaw struct {
	Contract *AIModelsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAIModels creates a new instance of AIModels, bound to a specific deployed contract.
func NewAIModels(address common.Address, backend bind.ContractBackend) (*AIModels, error) {
	contract, err := bindAIModels(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AIModels{AIModelsCaller: AIModelsCaller{contract: contract}, AIModelsTransactor: AIModelsTransactor{contract: contract}, AIModelsFilterer: AIModelsFilterer{contract: contract}}, nil
}

// NewAIModelsCaller creates a new read-only instance of AIModels, bound to a specific deployed contract.
func NewAIModelsCaller(address common.Address, caller bind.ContractCaller) (*AIModelsCaller, error) {
	contract, err := bindAIModels(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AIModelsCaller{contract: contract}, nil
}

// NewAIModelsTransactor creates a new write-only instance of AIModels, bound to a specific deployed contract.
func NewAIModelsTransactor(address common.Address, transactor bind.ContractTransactor) (*AIModelsTransactor, error) {
	contract, err := bindAIModels(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AIModelsTransactor{contract: contract}, nil
}

// NewAIModelsFilterer creates a new log filterer instance of AIModels, bound to a specific deployed contract.
func NewAIModelsFilterer(address common.Address, filterer bind.ContractFilterer) (*AIModelsFilterer, error) {
	contract, err := bindAIModels(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AIModelsFilterer{contract: contract}, nil
}

// bindAIModels binds a generic wrapper to an already deployed contract.
func bindAIModels(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AIModelsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AIModels *AIModelsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AIModels.Contract.AIModelsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AIModels *AIModelsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIModels.Contract.AIModelsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AIModels *AIModelsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AIModels.Contract.AIModelsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AIModels *AIModelsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AIModels.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AIModels *AIModelsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIModels.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AIModels *AIModelsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AIModels.Contract.contract.Transact(opts, method, params...)
}

// GetModelDistribution is a free data retrieval call binding the contract method 0x57793e68.
//
// Solidity: function getModelDistribution(uint256 modelId) view returns(address[])
func (_AIModels *AIModelsCaller) GetModelDistribution(opts *bind.CallOpts, modelId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "getModelDistribution", modelId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModelDistribution is a free data retrieval call binding the contract method 0x57793e68.
//
// Solidity: function getModelDistribution(uint256 modelId) view returns(address[])
func (_AIModels *AIModelsSession) GetModelDistribution(modelId *big.Int) ([]common.Address, error) {
	return _AIModels.Contract.GetModelDistribution(&_AIModels.CallOpts, modelId)
}

// GetModelDistribution is a free data retrieval call binding the contract method 0x57793e68.
//
// Solidity: function getModelDistribution(uint256 modelId) view returns(address[])
func (_AIModels *AIModelsCallerSession) GetModelDistribution(modelId *big.Int) ([]common.Address, error) {
	return _AIModels.Contract.GetModelDistribution(&_AIModels.CallOpts, modelId)
}

// GetNodeDeployment is a free data retrieval call binding the contract method 0x2ef64537.
//
// Solidity: function getNodeDeployment(address node) view returns(uint256[])
func (_AIModels *AIModelsCaller) GetNodeDeployment(opts *bind.CallOpts, node common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "getNodeDeployment", node)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNodeDeployment is a free data retrieval call binding the contract method 0x2ef64537.
//
// Solidity: function getNodeDeployment(address node) view returns(uint256[])
func (_AIModels *AIModelsSession) GetNodeDeployment(node common.Address) ([]*big.Int, error) {
	return _AIModels.Contract.GetNodeDeployment(&_AIModels.CallOpts, node)
}

// GetNodeDeployment is a free data retrieval call binding the contract method 0x2ef64537.
//
// Solidity: function getNodeDeployment(address node) view returns(uint256[])
func (_AIModels *AIModelsCallerSession) GetNodeDeployment(node common.Address) ([]*big.Int, error) {
	return _AIModels.Contract.GetNodeDeployment(&_AIModels.CallOpts, node)
}

// ModelDistribution is a free data retrieval call binding the contract method 0xdcca1c92.
//
// Solidity: function modelDistribution(uint256 , uint256 ) view returns(address)
func (_AIModels *AIModelsCaller) ModelDistribution(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "modelDistribution", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ModelDistribution is a free data retrieval call binding the contract method 0xdcca1c92.
//
// Solidity: function modelDistribution(uint256 , uint256 ) view returns(address)
func (_AIModels *AIModelsSession) ModelDistribution(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _AIModels.Contract.ModelDistribution(&_AIModels.CallOpts, arg0, arg1)
}

// ModelDistribution is a free data retrieval call binding the contract method 0xdcca1c92.
//
// Solidity: function modelDistribution(uint256 , uint256 ) view returns(address)
func (_AIModels *AIModelsCallerSession) ModelDistribution(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _AIModels.Contract.ModelDistribution(&_AIModels.CallOpts, arg0, arg1)
}

// ModelIds is a free data retrieval call binding the contract method 0x69f4296a.
//
// Solidity: function modelIds(string ) view returns(uint256)
func (_AIModels *AIModelsCaller) ModelIds(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "modelIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModelIds is a free data retrieval call binding the contract method 0x69f4296a.
//
// Solidity: function modelIds(string ) view returns(uint256)
func (_AIModels *AIModelsSession) ModelIds(arg0 string) (*big.Int, error) {
	return _AIModels.Contract.ModelIds(&_AIModels.CallOpts, arg0)
}

// ModelIds is a free data retrieval call binding the contract method 0x69f4296a.
//
// Solidity: function modelIds(string ) view returns(uint256)
func (_AIModels *AIModelsCallerSession) ModelIds(arg0 string) (*big.Int, error) {
	return _AIModels.Contract.ModelIds(&_AIModels.CallOpts, arg0)
}

// ModelOwns is a free data retrieval call binding the contract method 0xc78717d5.
//
// Solidity: function modelOwns(string ) view returns(address)
func (_AIModels *AIModelsCaller) ModelOwns(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "modelOwns", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ModelOwns is a free data retrieval call binding the contract method 0xc78717d5.
//
// Solidity: function modelOwns(string ) view returns(address)
func (_AIModels *AIModelsSession) ModelOwns(arg0 string) (common.Address, error) {
	return _AIModels.Contract.ModelOwns(&_AIModels.CallOpts, arg0)
}

// ModelOwns is a free data retrieval call binding the contract method 0xc78717d5.
//
// Solidity: function modelOwns(string ) view returns(address)
func (_AIModels *AIModelsCallerSession) ModelOwns(arg0 string) (common.Address, error) {
	return _AIModels.Contract.ModelOwns(&_AIModels.CallOpts, arg0)
}

// NextModelId is a free data retrieval call binding the contract method 0xe472ae8b.
//
// Solidity: function nextModelId() view returns(uint256)
func (_AIModels *AIModelsCaller) NextModelId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "nextModelId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextModelId is a free data retrieval call binding the contract method 0xe472ae8b.
//
// Solidity: function nextModelId() view returns(uint256)
func (_AIModels *AIModelsSession) NextModelId() (*big.Int, error) {
	return _AIModels.Contract.NextModelId(&_AIModels.CallOpts)
}

// NextModelId is a free data retrieval call binding the contract method 0xe472ae8b.
//
// Solidity: function nextModelId() view returns(uint256)
func (_AIModels *AIModelsCallerSession) NextModelId() (*big.Int, error) {
	return _AIModels.Contract.NextModelId(&_AIModels.CallOpts)
}

// NodeDeployment is a free data retrieval call binding the contract method 0x7d5ef48e.
//
// Solidity: function nodeDeployment(address , uint256 ) view returns(uint256)
func (_AIModels *AIModelsCaller) NodeDeployment(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "nodeDeployment", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeDeployment is a free data retrieval call binding the contract method 0x7d5ef48e.
//
// Solidity: function nodeDeployment(address , uint256 ) view returns(uint256)
func (_AIModels *AIModelsSession) NodeDeployment(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AIModels.Contract.NodeDeployment(&_AIModels.CallOpts, arg0, arg1)
}

// NodeDeployment is a free data retrieval call binding the contract method 0x7d5ef48e.
//
// Solidity: function nodeDeployment(address , uint256 ) view returns(uint256)
func (_AIModels *AIModelsCallerSession) NodeDeployment(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AIModels.Contract.NodeDeployment(&_AIModels.CallOpts, arg0, arg1)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AIModels *AIModelsCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AIModels *AIModelsSession) Registry() (common.Address, error) {
	return _AIModels.Contract.Registry(&_AIModels.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AIModels *AIModelsCallerSession) Registry() (common.Address, error) {
	return _AIModels.Contract.Registry(&_AIModels.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_AIModels *AIModelsCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_AIModels *AIModelsSession) StakeToken() (common.Address, error) {
	return _AIModels.Contract.StakeToken(&_AIModels.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_AIModels *AIModelsCallerSession) StakeToken() (common.Address, error) {
	return _AIModels.Contract.StakeToken(&_AIModels.CallOpts)
}

// UploadModels is a free data retrieval call binding the contract method 0x9bb0d612.
//
// Solidity: function uploadModels(uint256 ) view returns(uint256 modelId, string modelName, string modelVersion, address uploader, string extendInfo, uint256 timestamp, uint256 price)
func (_AIModels *AIModelsCaller) UploadModels(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ModelId      *big.Int
	ModelName    string
	ModelVersion string
	Uploader     common.Address
	ExtendInfo   string
	Timestamp    *big.Int
	Price        *big.Int
}, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "uploadModels", arg0)

	outstruct := new(struct {
		ModelId      *big.Int
		ModelName    string
		ModelVersion string
		Uploader     common.Address
		ExtendInfo   string
		Timestamp    *big.Int
		Price        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ModelId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ModelName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ModelVersion = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Uploader = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ExtendInfo = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UploadModels is a free data retrieval call binding the contract method 0x9bb0d612.
//
// Solidity: function uploadModels(uint256 ) view returns(uint256 modelId, string modelName, string modelVersion, address uploader, string extendInfo, uint256 timestamp, uint256 price)
func (_AIModels *AIModelsSession) UploadModels(arg0 *big.Int) (struct {
	ModelId      *big.Int
	ModelName    string
	ModelVersion string
	Uploader     common.Address
	ExtendInfo   string
	Timestamp    *big.Int
	Price        *big.Int
}, error) {
	return _AIModels.Contract.UploadModels(&_AIModels.CallOpts, arg0)
}

// UploadModels is a free data retrieval call binding the contract method 0x9bb0d612.
//
// Solidity: function uploadModels(uint256 ) view returns(uint256 modelId, string modelName, string modelVersion, address uploader, string extendInfo, uint256 timestamp, uint256 price)
func (_AIModels *AIModelsCallerSession) UploadModels(arg0 *big.Int) (struct {
	ModelId      *big.Int
	ModelName    string
	ModelVersion string
	Uploader     common.Address
	ExtendInfo   string
	Timestamp    *big.Int
	Price        *big.Int
}, error) {
	return _AIModels.Contract.UploadModels(&_AIModels.CallOpts, arg0)
}

// RecordModelUpload is a paid mutator transaction binding the contract method 0xc1ce7185.
//
// Solidity: function recordModelUpload(string modelName, string modelVersion, string modelExtendInfo, uint256 price) returns(uint256 modelId)
func (_AIModels *AIModelsTransactor) RecordModelUpload(opts *bind.TransactOpts, modelName string, modelVersion string, modelExtendInfo string, price *big.Int) (*types.Transaction, error) {
	return _AIModels.contract.Transact(opts, "recordModelUpload", modelName, modelVersion, modelExtendInfo, price)
}

// RecordModelUpload is a paid mutator transaction binding the contract method 0xc1ce7185.
//
// Solidity: function recordModelUpload(string modelName, string modelVersion, string modelExtendInfo, uint256 price) returns(uint256 modelId)
func (_AIModels *AIModelsSession) RecordModelUpload(modelName string, modelVersion string, modelExtendInfo string, price *big.Int) (*types.Transaction, error) {
	return _AIModels.Contract.RecordModelUpload(&_AIModels.TransactOpts, modelName, modelVersion, modelExtendInfo, price)
}

// RecordModelUpload is a paid mutator transaction binding the contract method 0xc1ce7185.
//
// Solidity: function recordModelUpload(string modelName, string modelVersion, string modelExtendInfo, uint256 price) returns(uint256 modelId)
func (_AIModels *AIModelsTransactorSession) RecordModelUpload(modelName string, modelVersion string, modelExtendInfo string, price *big.Int) (*types.Transaction, error) {
	return _AIModels.Contract.RecordModelUpload(&_AIModels.TransactOpts, modelName, modelVersion, modelExtendInfo, price)
}

// RemoveDeployment is a paid mutator transaction binding the contract method 0x907d0042.
//
// Solidity: function removeDeployment(uint256 modelId) returns()
func (_AIModels *AIModelsTransactor) RemoveDeployment(opts *bind.TransactOpts, modelId *big.Int) (*types.Transaction, error) {
	return _AIModels.contract.Transact(opts, "removeDeployment", modelId)
}

// RemoveDeployment is a paid mutator transaction binding the contract method 0x907d0042.
//
// Solidity: function removeDeployment(uint256 modelId) returns()
func (_AIModels *AIModelsSession) RemoveDeployment(modelId *big.Int) (*types.Transaction, error) {
	return _AIModels.Contract.RemoveDeployment(&_AIModels.TransactOpts, modelId)
}

// RemoveDeployment is a paid mutator transaction binding the contract method 0x907d0042.
//
// Solidity: function removeDeployment(uint256 modelId) returns()
func (_AIModels *AIModelsTransactorSession) RemoveDeployment(modelId *big.Int) (*types.Transaction, error) {
	return _AIModels.Contract.RemoveDeployment(&_AIModels.TransactOpts, modelId)
}

// ReportDeployment is a paid mutator transaction binding the contract method 0x64290cbf.
//
// Solidity: function reportDeployment(uint256 modelId) returns()
func (_AIModels *AIModelsTransactor) ReportDeployment(opts *bind.TransactOpts, modelId *big.Int) (*types.Transaction, error) {
	return _AIModels.contract.Transact(opts, "reportDeployment", modelId)
}

// ReportDeployment is a paid mutator transaction binding the contract method 0x64290cbf.
//
// Solidity: function reportDeployment(uint256 modelId) returns()
func (_AIModels *AIModelsSession) ReportDeployment(modelId *big.Int) (*types.Transaction, error) {
	return _AIModels.Contract.ReportDeployment(&_AIModels.TransactOpts, modelId)
}

// ReportDeployment is a paid mutator transaction binding the contract method 0x64290cbf.
//
// Solidity: function reportDeployment(uint256 modelId) returns()
func (_AIModels *AIModelsTransactorSession) ReportDeployment(modelId *big.Int) (*types.Transaction, error) {
	return _AIModels.Contract.ReportDeployment(&_AIModels.TransactOpts, modelId)
}

// AIModelsModelDeployedIterator is returned from FilterModelDeployed and is used to iterate over the raw logs and unpacked data for ModelDeployed events raised by the AIModels contract.
type AIModelsModelDeployedIterator struct {
	Event *AIModelsModelDeployed // Event containing the contract specifics and raw log

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
func (it *AIModelsModelDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIModelsModelDeployed)
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
		it.Event = new(AIModelsModelDeployed)
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
func (it *AIModelsModelDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIModelsModelDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIModelsModelDeployed represents a ModelDeployed event raised by the AIModels contract.
type AIModelsModelDeployed struct {
	Node    common.Address
	ModelId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelDeployed is a free log retrieval operation binding the contract event 0xb81e8e03e6a54d2b07bc11d3cddad332ded9a7018c0e554bd78262a1d4cd20d1.
//
// Solidity: event ModelDeployed(address indexed node, uint256 indexed modelId)
func (_AIModels *AIModelsFilterer) FilterModelDeployed(opts *bind.FilterOpts, node []common.Address, modelId []*big.Int) (*AIModelsModelDeployedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _AIModels.contract.FilterLogs(opts, "ModelDeployed", nodeRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &AIModelsModelDeployedIterator{contract: _AIModels.contract, event: "ModelDeployed", logs: logs, sub: sub}, nil
}

// WatchModelDeployed is a free log subscription operation binding the contract event 0xb81e8e03e6a54d2b07bc11d3cddad332ded9a7018c0e554bd78262a1d4cd20d1.
//
// Solidity: event ModelDeployed(address indexed node, uint256 indexed modelId)
func (_AIModels *AIModelsFilterer) WatchModelDeployed(opts *bind.WatchOpts, sink chan<- *AIModelsModelDeployed, node []common.Address, modelId []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _AIModels.contract.WatchLogs(opts, "ModelDeployed", nodeRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIModelsModelDeployed)
				if err := _AIModels.contract.UnpackLog(event, "ModelDeployed", log); err != nil {
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

// ParseModelDeployed is a log parse operation binding the contract event 0xb81e8e03e6a54d2b07bc11d3cddad332ded9a7018c0e554bd78262a1d4cd20d1.
//
// Solidity: event ModelDeployed(address indexed node, uint256 indexed modelId)
func (_AIModels *AIModelsFilterer) ParseModelDeployed(log types.Log) (*AIModelsModelDeployed, error) {
	event := new(AIModelsModelDeployed)
	if err := _AIModels.contract.UnpackLog(event, "ModelDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIModelsModelRemovedIterator is returned from FilterModelRemoved and is used to iterate over the raw logs and unpacked data for ModelRemoved events raised by the AIModels contract.
type AIModelsModelRemovedIterator struct {
	Event *AIModelsModelRemoved // Event containing the contract specifics and raw log

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
func (it *AIModelsModelRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIModelsModelRemoved)
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
		it.Event = new(AIModelsModelRemoved)
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
func (it *AIModelsModelRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIModelsModelRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIModelsModelRemoved represents a ModelRemoved event raised by the AIModels contract.
type AIModelsModelRemoved struct {
	Node    common.Address
	ModelId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelRemoved is a free log retrieval operation binding the contract event 0x271097aa25aca38eb083e68e5cafcbd8247d421aea1887c9630402cda1706b63.
//
// Solidity: event ModelRemoved(address indexed node, uint256 indexed modelId)
func (_AIModels *AIModelsFilterer) FilterModelRemoved(opts *bind.FilterOpts, node []common.Address, modelId []*big.Int) (*AIModelsModelRemovedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _AIModels.contract.FilterLogs(opts, "ModelRemoved", nodeRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &AIModelsModelRemovedIterator{contract: _AIModels.contract, event: "ModelRemoved", logs: logs, sub: sub}, nil
}

// WatchModelRemoved is a free log subscription operation binding the contract event 0x271097aa25aca38eb083e68e5cafcbd8247d421aea1887c9630402cda1706b63.
//
// Solidity: event ModelRemoved(address indexed node, uint256 indexed modelId)
func (_AIModels *AIModelsFilterer) WatchModelRemoved(opts *bind.WatchOpts, sink chan<- *AIModelsModelRemoved, node []common.Address, modelId []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _AIModels.contract.WatchLogs(opts, "ModelRemoved", nodeRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIModelsModelRemoved)
				if err := _AIModels.contract.UnpackLog(event, "ModelRemoved", log); err != nil {
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

// ParseModelRemoved is a log parse operation binding the contract event 0x271097aa25aca38eb083e68e5cafcbd8247d421aea1887c9630402cda1706b63.
//
// Solidity: event ModelRemoved(address indexed node, uint256 indexed modelId)
func (_AIModels *AIModelsFilterer) ParseModelRemoved(log types.Log) (*AIModelsModelRemoved, error) {
	event := new(AIModelsModelRemoved)
	if err := _AIModels.contract.UnpackLog(event, "ModelRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIModelsUploadModeledIterator is returned from FilterUploadModeled and is used to iterate over the raw logs and unpacked data for UploadModeled events raised by the AIModels contract.
type AIModelsUploadModeledIterator struct {
	Event *AIModelsUploadModeled // Event containing the contract specifics and raw log

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
func (it *AIModelsUploadModeledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIModelsUploadModeled)
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
		it.Event = new(AIModelsUploadModeled)
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
func (it *AIModelsUploadModeledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIModelsUploadModeledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIModelsUploadModeled represents a UploadModeled event raised by the AIModels contract.
type AIModelsUploadModeled struct {
	ModelId      *big.Int
	Uploader     common.Address
	ModelName    string
	ModelVersion string
	ModelInfo    string
	Price        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUploadModeled is a free log retrieval operation binding the contract event 0xa17734f4a6749e7cc4947af9f12fb87299c2c2ab175158d74d24c8287fdee436.
//
// Solidity: event UploadModeled(uint256 indexed modelId, address indexed uploader, string modelName, string modelVersion, string modelInfo, uint256 price)
func (_AIModels *AIModelsFilterer) FilterUploadModeled(opts *bind.FilterOpts, modelId []*big.Int, uploader []common.Address) (*AIModelsUploadModeledIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var uploaderRule []interface{}
	for _, uploaderItem := range uploader {
		uploaderRule = append(uploaderRule, uploaderItem)
	}

	logs, sub, err := _AIModels.contract.FilterLogs(opts, "UploadModeled", modelIdRule, uploaderRule)
	if err != nil {
		return nil, err
	}
	return &AIModelsUploadModeledIterator{contract: _AIModels.contract, event: "UploadModeled", logs: logs, sub: sub}, nil
}

// WatchUploadModeled is a free log subscription operation binding the contract event 0xa17734f4a6749e7cc4947af9f12fb87299c2c2ab175158d74d24c8287fdee436.
//
// Solidity: event UploadModeled(uint256 indexed modelId, address indexed uploader, string modelName, string modelVersion, string modelInfo, uint256 price)
func (_AIModels *AIModelsFilterer) WatchUploadModeled(opts *bind.WatchOpts, sink chan<- *AIModelsUploadModeled, modelId []*big.Int, uploader []common.Address) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var uploaderRule []interface{}
	for _, uploaderItem := range uploader {
		uploaderRule = append(uploaderRule, uploaderItem)
	}

	logs, sub, err := _AIModels.contract.WatchLogs(opts, "UploadModeled", modelIdRule, uploaderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIModelsUploadModeled)
				if err := _AIModels.contract.UnpackLog(event, "UploadModeled", log); err != nil {
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

// ParseUploadModeled is a log parse operation binding the contract event 0xa17734f4a6749e7cc4947af9f12fb87299c2c2ab175158d74d24c8287fdee436.
//
// Solidity: event UploadModeled(uint256 indexed modelId, address indexed uploader, string modelName, string modelVersion, string modelInfo, uint256 price)
func (_AIModels *AIModelsFilterer) ParseUploadModeled(log types.Log) (*AIModelsUploadModeled, error) {
	event := new(AIModelsUploadModeled)
	if err := _AIModels.contract.UnpackLog(event, "UploadModeled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
