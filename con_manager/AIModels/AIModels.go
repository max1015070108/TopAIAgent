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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"ModelDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"ModelRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelInfo\",\"type\":\"string\"}],\"name\":\"UploadModeled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"getModelDistribution\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getNodeDeployment\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"modelDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"modelIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"modelStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextModelId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeDeployment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelInfo\",\"type\":\"string\"}],\"name\":\"recordModelUpload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractNodesRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"removeDeployment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"reportDeployment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uploadModels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"extendInfo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052600160045534801561001557600080fd5b50604051611573380380611573833981016040819052610034916100b3565b6001600160a01b03811661008e5760405162461bcd60e51b815260206004820152601860248201527f496e76616c696420726567697374727920616464726573730000000000000000604482015260640160405180910390fd5b600080546001600160a01b0319166001600160a01b03929092169190911790556100e1565b6000602082840312156100c4578081fd5b81516001600160a01b03811681146100da578182fd5b9392505050565b611483806100f06000396000f3fe6080604052600436106100c25760003560e01c806369f4296a1161007f578063907d004211610059578063907d00421461023b5780639bb0d6121461025b578063dcca1c921461028d578063e472ae8b146102ad576100c2565b806369f4296a146101ab5780637b103999146101e35780637d5ef48e1461021b576100c2565b806309364a7e146100c75780632ef64537146100fa5780632f6ab275146101275780633a4b66f11461015457806357793e681461015e57806364290cbf1461018b575b600080fd5b3480156100d357600080fd5b506100e76100e2366004611019565b6102c3565b6040519081526020015b60405180910390f35b34801561010657600080fd5b5061011a610115366004610faf565b61058e565b6040516100f1919061128d565b34801561013357600080fd5b506100e7610142366004610faf565b60036020526000908152604090205481565b61015c6105fb565b005b34801561016a57600080fd5b5061017e610179366004611159565b610628565b6040516100f19190611240565b34801561019757600080fd5b5061015c6101a6366004611159565b610693565b3480156101b757600080fd5b506100e76101c63660046110af565b805160208183018101805160018252928201919093012091525481565b3480156101ef57600080fd5b50600054610203906001600160a01b031681565b6040516001600160a01b0390911681526020016100f1565b34801561022757600080fd5b506100e7610236366004610fd0565b6107e5565b34801561024757600080fd5b5061015c610256366004611159565b610816565b34801561026757600080fd5b5061027b610276366004611159565b61085a565b6040516100f19695949392919061130e565b34801561029957600080fd5b506102036102a8366004611171565b610a36565b3480156102b957600080fd5b506100e760045481565b60008061033988888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8c018190048102820181019092528a815292508a9150899081908401838280828437600092019190915250610a6e92505050565b905060018160405161034b91906111e8565b90815260200160405180910390205460001461039c5760405162461bcd60e51b815260206004820152600b60248201526a135bd9195b08195e1a5cdd60aa1b60448201526064015b60405180910390fd5b6040518060c00160405280600454815260200189898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f8a01819004810282018101909252888152918101919089908990819084018382808284376000920191909152505050908252503360208083019190915260408051601f880183900483028101830182528781529201919087908790819084018382808284376000920182905250938552505042602093840152506004548152600282526040902082518155828201518051919261049592600185019290910190610eb8565b50604082015180516104b1916002840191602090910190610eb8565b5060608201516003820180546001600160a01b0319166001600160a01b03909216919091179055608082015180516104f3916004840191602090910190610eb8565b5060a0820151816005015590505060045460018260405161051491906111e8565b9081526040519081900360200181209190915560045433917fc10402236a0d3324d97ba17b997a8c5cd276aa435f3ebc1b9b14866eaba568ac90610563908c908c908c908c908c908c906112c5565b60405180910390a36004805492508290600061057e83611406565b9190505550509695505050505050565b6001600160a01b0381166000908152600660209081526040918290208054835181840281018401909452808452606093928301828280156105ee57602002820191906000526020600020905b8154815260200190600101908083116105da575b505050505090505b919050565b3415610626573360009081526003602052604081208054349290610620908490611370565b90915550505b565b6000818152600560209081526040918290208054835181840281018401909452808452606093928301828280156105ee57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161066a5750505050509050919050565b600054604051631846d2f560e31b81523360048201526001600160a01b039091169063c23697a89060240160206040518083038186803b1580156106d657600080fd5b505afa1580156106ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070e9190610ff9565b61074f5760405162461bcd60e51b81526020600482015260126024820152714e6f6465206973206e6f742061637469766560701b6044820152606401610393565b600081815260026020526040902080546107a05760405162461bcd60e51b8152602060048201526012602482015271135bd9195b081a5cc81b9bdd08195e1a5cdd60721b6044820152606401610393565b6107aa8233610a9a565b6107b48233610b85565b604051829033907fb81e8e03e6a54d2b07bc11d3cddad332ded9a7018c0e554bd78262a1d4cd20d190600090a35050565b6006602052816000526040600020818154811061080157600080fd5b90600052602060002001600091509150505481565b6108208133610c5c565b61082a8133610db1565b604051819033907f271097aa25aca38eb083e68e5cafcbd8247d421aea1887c9630402cda1706b6390600090a350565b6002602052600090815260409020805460018201805491929161087c906113cb565b80601f01602080910402602001604051908101604052809291908181526020018280546108a8906113cb565b80156108f55780601f106108ca576101008083540402835291602001916108f5565b820191906000526020600020905b8154815290600101906020018083116108d857829003601f168201915b50505050509080600201805461090a906113cb565b80601f0160208091040260200160405190810160405280929190818152602001828054610936906113cb565b80156109835780601f1061095857610100808354040283529160200191610983565b820191906000526020600020905b81548152906001019060200180831161096657829003601f168201915b505050600384015460048501805494956001600160a01b039092169491935091506109ad906113cb565b80601f01602080910402602001604051908101604052809291908181526020018280546109d9906113cb565b8015610a265780601f106109fb57610100808354040283529160200191610a26565b820191906000526020600020905b815481529060010190602001808311610a0957829003601f168201915b5050505050908060050154905086565b60056020528160005260406000208181548110610a5257600080fd5b6000918252602090912001546001600160a01b03169150829050565b60608282604051602001610a83929190611204565b604051602081830303815290604052905092915050565b6000828152600560205260408120905b8154811015610b5257826001600160a01b0316828281548110610add57634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b03161415610b405760405162461bcd60e51b815260206004820181905260248201527f4d6f64656c20646973747269627574696f6e20616c72656164792065786973746044820152606401610393565b80610b4a81611406565b915050610aaa565b50505060009081526005602090815260408220805460018101825590835291200180546001600160a01b03191633179055565b6001600160a01b0381166000908152600660205260408120905b8154811015610c345783828281548110610bc957634e487b7160e01b600052603260045260246000fd5b90600052602060002001541415610c225760405162461bcd60e51b815260206004820152601d60248201527f4e6f6465206465706c6f796d656e7420616c72656164792065786973740000006044820152606401610393565b80610c2c81611406565b915050610b9f565b5050336000908152600660209081526040822080546001810182559083529120019190915550565b6000828152600560205260408120905b8154811015610dab57826001600160a01b0316828281548110610c9f57634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b03161415610d995781548290610cca90600190611388565b81548110610ce857634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b0316828281548110610d2657634e487b7160e01b600052603260045260246000fd5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081805480610d7257634e487b7160e01b600052603160045260246000fd5b600082815260209020810160001990810180546001600160a01b0319169055019055610dab565b80610da381611406565b915050610c6c565b50505050565b6001600160a01b0381166000908152600660205260408120905b8154811015610dab5783828281548110610df557634e487b7160e01b600052603260045260246000fd5b90600052602060002001541415610ea65781548290610e1690600190611388565b81548110610e3457634e487b7160e01b600052603260045260246000fd5b9060005260206000200154828281548110610e5f57634e487b7160e01b600052603260045260246000fd5b906000526020600020018190555081805480610e8b57634e487b7160e01b600052603160045260246000fd5b60019003818190600052602060002001600090559055610dab565b80610eb081611406565b915050610dcb565b828054610ec4906113cb565b90600052602060002090601f016020900481019282610ee65760008555610f2c565b82601f10610eff57805160ff1916838001178555610f2c565b82800160010185558215610f2c579182015b82811115610f2c578251825591602001919060010190610f11565b50610f38929150610f3c565b5090565b5b80821115610f385760008155600101610f3d565b80356001600160a01b03811681146105f657600080fd5b60008083601f840112610f79578081fd5b50813567ffffffffffffffff811115610f90578182fd5b602083019150836020828501011115610fa857600080fd5b9250929050565b600060208284031215610fc0578081fd5b610fc982610f51565b9392505050565b60008060408385031215610fe2578081fd5b610feb83610f51565b946020939093013593505050565b60006020828403121561100a578081fd5b81518015158114610fc9578182fd5b60008060008060008060608789031215611031578182fd5b863567ffffffffffffffff80821115611048578384fd5b6110548a838b01610f68565b9098509650602089013591508082111561106c578384fd5b6110788a838b01610f68565b90965094506040890135915080821115611090578384fd5b5061109d89828a01610f68565b979a9699509497509295939492505050565b6000602082840312156110c0578081fd5b813567ffffffffffffffff808211156110d7578283fd5b818401915084601f8301126110ea578283fd5b8135818111156110fc576110fc611437565b604051601f8201601f19908116603f0116810190838211818310171561112457611124611437565b8160405282815287602084870101111561113c578586fd5b826020860160208301379182016020019490945295945050505050565b60006020828403121561116a578081fd5b5035919050565b60008060408385031215611183578182fd5b50508035926020909101359150565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b600081518084526111d481602086016020860161139f565b601f01601f19169290920160200192915050565b600082516111fa81846020870161139f565b9190910192915050565b6000835161121681846020880161139f565b602f60f81b908301908152835161123481600184016020880161139f565b01600101949350505050565b6020808252825182820181905260009190848201906040850190845b818110156112815783516001600160a01b03168352928401929184019160010161125c565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b81811015611281578351835292840192918401916001016112a9565b6000606082526112d960608301888a611192565b82810360208401526112ec818789611192565b90508281036040840152611301818587611192565b9998505050505050505050565b600087825260c0602083015261132760c08301886111bc565b828103604084015261133981886111bc565b6001600160a01b03871660608501528381036080850152905061135c81866111bc565b9150508260a0830152979650505050505050565b6000821982111561138357611383611421565b500190565b60008282101561139a5761139a611421565b500390565b60005b838110156113ba5781810151838201526020016113a2565b83811115610dab5750506000910152565b600181811c908216806113df57607f821691505b6020821081141561140057634e487b7160e01b600052602260045260246000fd5b50919050565b600060001982141561141a5761141a611421565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212205247c392543bef4cc9aea1eacf03f9f1507bdd12aef787191cde835e782bbc6264736f6c63430008030033",
}

// AIModelsABI is the input ABI used to generate the binding from.
// Deprecated: Use AIModelsMetaData.ABI instead.
var AIModelsABI = AIModelsMetaData.ABI

// AIModelsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AIModelsMetaData.Bin instead.
var AIModelsBin = AIModelsMetaData.Bin

// DeployAIModels deploys a new Ethereum contract, binding an instance of AIModels to it.
func DeployAIModels(auth *bind.TransactOpts, backend bind.ContractBackend, _registry common.Address) (common.Address, *types.Transaction, *AIModels, error) {
	parsed, err := AIModelsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AIModelsBin), backend, _registry)
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

// ModelStakes is a free data retrieval call binding the contract method 0x2f6ab275.
//
// Solidity: function modelStakes(address ) view returns(uint256)
func (_AIModels *AIModelsCaller) ModelStakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "modelStakes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModelStakes is a free data retrieval call binding the contract method 0x2f6ab275.
//
// Solidity: function modelStakes(address ) view returns(uint256)
func (_AIModels *AIModelsSession) ModelStakes(arg0 common.Address) (*big.Int, error) {
	return _AIModels.Contract.ModelStakes(&_AIModels.CallOpts, arg0)
}

// ModelStakes is a free data retrieval call binding the contract method 0x2f6ab275.
//
// Solidity: function modelStakes(address ) view returns(uint256)
func (_AIModels *AIModelsCallerSession) ModelStakes(arg0 common.Address) (*big.Int, error) {
	return _AIModels.Contract.ModelStakes(&_AIModels.CallOpts, arg0)
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

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_AIModels *AIModelsTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIModels.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_AIModels *AIModelsSession) Stake() (*types.Transaction, error) {
	return _AIModels.Contract.Stake(&_AIModels.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_AIModels *AIModelsTransactorSession) Stake() (*types.Transaction, error) {
	return _AIModels.Contract.Stake(&_AIModels.TransactOpts)
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
