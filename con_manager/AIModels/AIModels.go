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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"ModelDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"ModelRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelInfo\",\"type\":\"string\"}],\"name\":\"UploadModeled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"getModelDistribution\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getNodeDeployment\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"modelDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"modelEvaluations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"parameter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"modelIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextModelId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeDeployment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelInfo\",\"type\":\"string\"}],\"name\":\"recordModelUpload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractNodesRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"removeDeployment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"reportDeployment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIStake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uploadModels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"extendInfo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260016004553480156200001657600080fd5b506040516200154938038062001549833981016040819052620000399162000148565b6001600160a01b038216620000a15760405162461bcd60e51b8152602060048201526024808201527f496e76616c6964207175616e74697479206f66207265676973747279206164646044820152637265737360e01b60648201526084015b60405180910390fd5b6001600160a01b038116620000f95760405162461bcd60e51b815260206004820152601360248201527f496e76616c6964207374616b6520746f6b656e00000000000000000000000000604482015260640162000098565b600080546001600160a01b039384166001600160a01b031991821617909155600580549290931691161790556200017f565b80516001600160a01b03811681146200014357600080fd5b919050565b600080604083850312156200015b578182fd5b62000166836200012b565b915062000176602084016200012b565b90509250929050565b6113ba806200018f6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80637b1039991161008c578063907d004211610066578063907d0042146101eb5780639bb0d612146101fe578063dcca1c9214610223578063e472ae8b14610236576100cf565b80637b103999146101a55780637d5ef48e146101b85780638981d225146101cb576100cf565b806309364a7e146100d45780632ef64537146100fa57806351ed6a301461011a57806357793e681461014557806364290cbf1461016557806369f4296a1461017a575b600080fd5b6100e76100e2366004610f68565b61023f565b6040519081526020015b60405180910390f35b61010d610108366004610efe565b61050a565b6040516100f191906111dc565b60055461012d906001600160a01b031681565b6040516001600160a01b0390911681526020016100f1565b6101586101533660046110a8565b610577565b6040516100f1919061118f565b6101786101733660046110a8565b6105e2565b005b6100e7610188366004610ffe565b805160208183018101805160018252928201919093012091525481565b60005461012d906001600160a01b031681565b6100e76101c6366004610f1f565b610734565b6100e76101d93660046110a8565b60036020526000908152604090205481565b6101786101f93660046110a8565b610765565b61021161020c3660046110a8565b6107a9565b6040516100f19695949392919061125d565b61012d6102313660046110c0565b610985565b6100e760045481565b6000806102b588888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8c018190048102820181019092528a815292508a91508990819084018382808284376000920191909152506109bd92505050565b90506001816040516102c79190611137565b9081526020016040518091039020546000146103185760405162461bcd60e51b815260206004820152600b60248201526a135bd9195b08195e1a5cdd60aa1b60448201526064015b60405180910390fd5b6040518060c00160405280600454815260200189898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f8a01819004810282018101909252888152918101919089908990819084018382808284376000920191909152505050908252503360208083019190915260408051601f880183900483028101830182528781529201919087908790819084018382808284376000920182905250938552505042602093840152506004548152600282526040902082518155828201518051919261041192600185019290910190610e07565b506040820151805161042d916002840191602090910190610e07565b5060608201516003820180546001600160a01b0319166001600160a01b039092169190911790556080820151805161046f916004840191602090910190610e07565b5060a082015181600501559050506004546001826040516104909190611137565b9081526040519081900360200181209190915560045433917fc10402236a0d3324d97ba17b997a8c5cd276aa435f3ebc1b9b14866eaba568ac906104df908c908c908c908c908c908c90611214565b60405180910390a3600480549250829060006104fa8361133d565b9190505550509695505050505050565b6001600160a01b03811660009081526007602090815260409182902080548351818402810184019094528084526060939283018282801561056a57602002820191906000526020600020905b815481526020019060010190808311610556575b505050505090505b919050565b60008181526006602090815260409182902080548351818402810184019094528084526060939283018282801561056a57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116105b95750505050509050919050565b600054604051631846d2f560e31b81523360048201526001600160a01b039091169063c23697a89060240160206040518083038186803b15801561062557600080fd5b505afa158015610639573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061065d9190610f48565b61069e5760405162461bcd60e51b81526020600482015260126024820152714e6f6465206973206e6f742061637469766560701b604482015260640161030f565b600081815260026020526040902080546106ef5760405162461bcd60e51b8152602060048201526012602482015271135bd9195b081a5cc81b9bdd08195e1a5cdd60721b604482015260640161030f565b6106f982336109e9565b6107038233610ad4565b604051829033907fb81e8e03e6a54d2b07bc11d3cddad332ded9a7018c0e554bd78262a1d4cd20d190600090a35050565b6007602052816000526040600020818154811061075057600080fd5b90600052602060002001600091509150505481565b61076f8133610bab565b6107798133610d00565b604051819033907f271097aa25aca38eb083e68e5cafcbd8247d421aea1887c9630402cda1706b6390600090a350565b600260205260009081526040902080546001820180549192916107cb90611302565b80601f01602080910402602001604051908101604052809291908181526020018280546107f790611302565b80156108445780601f1061081957610100808354040283529160200191610844565b820191906000526020600020905b81548152906001019060200180831161082757829003601f168201915b50505050509080600201805461085990611302565b80601f016020809104026020016040519081016040528092919081815260200182805461088590611302565b80156108d25780601f106108a7576101008083540402835291602001916108d2565b820191906000526020600020905b8154815290600101906020018083116108b557829003601f168201915b505050600384015460048501805494956001600160a01b039092169491935091506108fc90611302565b80601f016020809104026020016040519081016040528092919081815260200182805461092890611302565b80156109755780601f1061094a57610100808354040283529160200191610975565b820191906000526020600020905b81548152906001019060200180831161095857829003601f168201915b5050505050908060050154905086565b600660205281600052604060002081815481106109a157600080fd5b6000918252602090912001546001600160a01b03169150829050565b606082826040516020016109d2929190611153565b604051602081830303815290604052905092915050565b6000828152600660205260408120905b8154811015610aa157826001600160a01b0316828281548110610a2c57634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b03161415610a8f5760405162461bcd60e51b815260206004820181905260248201527f4d6f64656c20646973747269627574696f6e20616c7265616479206578697374604482015260640161030f565b80610a998161133d565b9150506109f9565b50505060009081526006602090815260408220805460018101825590835291200180546001600160a01b03191633179055565b6001600160a01b0381166000908152600760205260408120905b8154811015610b835783828281548110610b1857634e487b7160e01b600052603260045260246000fd5b90600052602060002001541415610b715760405162461bcd60e51b815260206004820152601d60248201527f4e6f6465206465706c6f796d656e7420616c7265616479206578697374000000604482015260640161030f565b80610b7b8161133d565b915050610aee565b5050336000908152600760209081526040822080546001810182559083529120019190915550565b6000828152600660205260408120905b8154811015610cfa57826001600160a01b0316828281548110610bee57634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b03161415610ce85781548290610c19906001906112bf565b81548110610c3757634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b0316828281548110610c7557634e487b7160e01b600052603260045260246000fd5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081805480610cc157634e487b7160e01b600052603160045260246000fd5b600082815260209020810160001990810180546001600160a01b0319169055019055610cfa565b80610cf28161133d565b915050610bbb565b50505050565b6001600160a01b0381166000908152600760205260408120905b8154811015610cfa5783828281548110610d4457634e487b7160e01b600052603260045260246000fd5b90600052602060002001541415610df55781548290610d65906001906112bf565b81548110610d8357634e487b7160e01b600052603260045260246000fd5b9060005260206000200154828281548110610dae57634e487b7160e01b600052603260045260246000fd5b906000526020600020018190555081805480610dda57634e487b7160e01b600052603160045260246000fd5b60019003818190600052602060002001600090559055610cfa565b80610dff8161133d565b915050610d1a565b828054610e1390611302565b90600052602060002090601f016020900481019282610e355760008555610e7b565b82601f10610e4e57805160ff1916838001178555610e7b565b82800160010185558215610e7b579182015b82811115610e7b578251825591602001919060010190610e60565b50610e87929150610e8b565b5090565b5b80821115610e875760008155600101610e8c565b80356001600160a01b038116811461057257600080fd5b60008083601f840112610ec8578081fd5b50813567ffffffffffffffff811115610edf578182fd5b602083019150836020828501011115610ef757600080fd5b9250929050565b600060208284031215610f0f578081fd5b610f1882610ea0565b9392505050565b60008060408385031215610f31578081fd5b610f3a83610ea0565b946020939093013593505050565b600060208284031215610f59578081fd5b81518015158114610f18578182fd5b60008060008060008060608789031215610f80578182fd5b863567ffffffffffffffff80821115610f97578384fd5b610fa38a838b01610eb7565b90985096506020890135915080821115610fbb578384fd5b610fc78a838b01610eb7565b90965094506040890135915080821115610fdf578384fd5b50610fec89828a01610eb7565b979a9699509497509295939492505050565b60006020828403121561100f578081fd5b813567ffffffffffffffff80821115611026578283fd5b818401915084601f830112611039578283fd5b81358181111561104b5761104b61136e565b604051601f8201601f19908116603f011681019083821181831017156110735761107361136e565b8160405282815287602084870101111561108b578586fd5b826020860160208301379182016020019490945295945050505050565b6000602082840312156110b9578081fd5b5035919050565b600080604083850312156110d2578182fd5b50508035926020909101359150565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b600081518084526111238160208601602086016112d6565b601f01601f19169290920160200192915050565b600082516111498184602087016112d6565b9190910192915050565b600083516111658184602088016112d6565b602f60f81b90830190815283516111838160018401602088016112d6565b01600101949350505050565b6020808252825182820181905260009190848201906040850190845b818110156111d05783516001600160a01b0316835292840192918401916001016111ab565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156111d0578351835292840192918401916001016111f8565b60006060825261122860608301888a6110e1565b828103602084015261123b8187896110e1565b905082810360408401526112508185876110e1565b9998505050505050505050565b600087825260c0602083015261127660c083018861110b565b8281036040840152611288818861110b565b6001600160a01b0387166060850152838103608085015290506112ab818661110b565b9150508260a0830152979650505050505050565b6000828210156112d1576112d1611358565b500390565b60005b838110156112f15781810151838201526020016112d9565b83811115610cfa5750506000910152565b600181811c9082168061131657607f821691505b6020821081141561133757634e487b7160e01b600052602260045260246000fd5b50919050565b600060001982141561135157611351611358565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220e389bb7637c41c3885d555cacb2885ce57aad7f3b7fc1879837fc6176266426f64736f6c63430008030033",
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

// ModelEvaluations is a free data retrieval call binding the contract method 0x8981d225.
//
// Solidity: function modelEvaluations(uint256 ) view returns(uint256 parameter)
func (_AIModels *AIModelsCaller) ModelEvaluations(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "modelEvaluations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModelEvaluations is a free data retrieval call binding the contract method 0x8981d225.
//
// Solidity: function modelEvaluations(uint256 ) view returns(uint256 parameter)
func (_AIModels *AIModelsSession) ModelEvaluations(arg0 *big.Int) (*big.Int, error) {
	return _AIModels.Contract.ModelEvaluations(&_AIModels.CallOpts, arg0)
}

// ModelEvaluations is a free data retrieval call binding the contract method 0x8981d225.
//
// Solidity: function modelEvaluations(uint256 ) view returns(uint256 parameter)
func (_AIModels *AIModelsCallerSession) ModelEvaluations(arg0 *big.Int) (*big.Int, error) {
	return _AIModels.Contract.ModelEvaluations(&_AIModels.CallOpts, arg0)
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
// Solidity: function uploadModels(uint256 ) view returns(uint256 modelId, string modelName, string modelVersion, address uploader, string extendInfo, uint256 timestamp)
func (_AIModels *AIModelsCaller) UploadModels(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ModelId      *big.Int
	ModelName    string
	ModelVersion string
	Uploader     common.Address
	ExtendInfo   string
	Timestamp    *big.Int
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

	return *outstruct, err

}

// UploadModels is a free data retrieval call binding the contract method 0x9bb0d612.
//
// Solidity: function uploadModels(uint256 ) view returns(uint256 modelId, string modelName, string modelVersion, address uploader, string extendInfo, uint256 timestamp)
func (_AIModels *AIModelsSession) UploadModels(arg0 *big.Int) (struct {
	ModelId      *big.Int
	ModelName    string
	ModelVersion string
	Uploader     common.Address
	ExtendInfo   string
	Timestamp    *big.Int
}, error) {
	return _AIModels.Contract.UploadModels(&_AIModels.CallOpts, arg0)
}

// UploadModels is a free data retrieval call binding the contract method 0x9bb0d612.
//
// Solidity: function uploadModels(uint256 ) view returns(uint256 modelId, string modelName, string modelVersion, address uploader, string extendInfo, uint256 timestamp)
func (_AIModels *AIModelsCallerSession) UploadModels(arg0 *big.Int) (struct {
	ModelId      *big.Int
	ModelName    string
	ModelVersion string
	Uploader     common.Address
	ExtendInfo   string
	Timestamp    *big.Int
}, error) {
	return _AIModels.Contract.UploadModels(&_AIModels.CallOpts, arg0)
}

// RecordModelUpload is a paid mutator transaction binding the contract method 0x09364a7e.
//
// Solidity: function recordModelUpload(string modelName, string modelVersion, string modelInfo) returns(uint256 modelId)
func (_AIModels *AIModelsTransactor) RecordModelUpload(opts *bind.TransactOpts, modelName string, modelVersion string, modelInfo string) (*types.Transaction, error) {
	return _AIModels.contract.Transact(opts, "recordModelUpload", modelName, modelVersion, modelInfo)
}

// RecordModelUpload is a paid mutator transaction binding the contract method 0x09364a7e.
//
// Solidity: function recordModelUpload(string modelName, string modelVersion, string modelInfo) returns(uint256 modelId)
func (_AIModels *AIModelsSession) RecordModelUpload(modelName string, modelVersion string, modelInfo string) (*types.Transaction, error) {
	return _AIModels.Contract.RecordModelUpload(&_AIModels.TransactOpts, modelName, modelVersion, modelInfo)
}

// RecordModelUpload is a paid mutator transaction binding the contract method 0x09364a7e.
//
// Solidity: function recordModelUpload(string modelName, string modelVersion, string modelInfo) returns(uint256 modelId)
func (_AIModels *AIModelsTransactorSession) RecordModelUpload(modelName string, modelVersion string, modelInfo string) (*types.Transaction, error) {
	return _AIModels.Contract.RecordModelUpload(&_AIModels.TransactOpts, modelName, modelVersion, modelInfo)
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
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUploadModeled is a free log retrieval operation binding the contract event 0xc10402236a0d3324d97ba17b997a8c5cd276aa435f3ebc1b9b14866eaba568ac.
//
// Solidity: event UploadModeled(uint256 indexed modelId, address indexed uploader, string modelName, string modelVersion, string modelInfo)
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

// WatchUploadModeled is a free log subscription operation binding the contract event 0xc10402236a0d3324d97ba17b997a8c5cd276aa435f3ebc1b9b14866eaba568ac.
//
// Solidity: event UploadModeled(uint256 indexed modelId, address indexed uploader, string modelName, string modelVersion, string modelInfo)
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

// ParseUploadModeled is a log parse operation binding the contract event 0xc10402236a0d3324d97ba17b997a8c5cd276aa435f3ebc1b9b14866eaba568ac.
//
// Solidity: event UploadModeled(uint256 indexed modelId, address indexed uploader, string modelName, string modelVersion, string modelInfo)
func (_AIModels *AIModelsFilterer) ParseUploadModeled(log types.Log) (*AIModelsUploadModeled, error) {
	event := new(AIModelsUploadModeled)
	if err := _AIModels.contract.UnpackLog(event, "UploadModeled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
