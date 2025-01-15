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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"ModelDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"ModelRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"ModelUploadStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"ModelUploadUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"modelUser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"ModelUsedStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"modelUser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"ModelUsedUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelInfo\",\"type\":\"string\"}],\"name\":\"UploadModeled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"getModelDistribution\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getNodeDeployment\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"modelDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"modelEvaluations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"parameter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"modelIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"modelUploadStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"modelUsedStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextModelId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeDeployment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelInfo\",\"type\":\"string\"}],\"name\":\"recordModelUpload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractNodesRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"removeDeployment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"reportDeployment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeModelUpload\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeModelUsed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uploadModels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"extendInfo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052600160065534801561001557600080fd5b50604051611734380380611734833981016040819052610034916100bf565b6001600160a01b03811661009a5760405162461bcd60e51b8152602060048201526024808201527f496e76616c6964207175616e74697479206f66207265676973747279206164646044820152637265737360e01b606482015260840160405180910390fd5b600080546001600160a01b0319166001600160a01b03929092169190911790556100ed565b6000602082840312156100d0578081fd5b81516001600160a01b03811681146100e6578182fd5b9392505050565b611638806100fc6000396000f3fe6080604052600436106100f35760003560e01c80638981d2251161008a578063ab96d7bc11610059578063ab96d7bc146102eb578063baf70a46146102f3578063dcca1c9214610320578063e472ae8b14610340576100f3565b80638981d2251461023f578063907d00421461026c5780639bb0d6121461028c578063a3217acc146102be576100f3565b806364290cbf116100c657806364290cbf1461018f57806369f4296a146101af5780637b103999146101e75780637d5ef48e1461021f576100f3565b806309364a7e146100f85780632ef645371461012b57806357793e68146101585780636268e03e14610185575b600080fd5b34801561010457600080fd5b506101186101133660046111ce565b610356565b6040519081526020015b60405180910390f35b34801561013757600080fd5b5061014b610146366004611164565b610621565b6040516101229190611442565b34801561016457600080fd5b5061017861017336600461130e565b61068e565b60405161012291906113f5565b61018d6106f9565b005b34801561019b57600080fd5b5061018d6101aa36600461130e565b6107a8565b3480156101bb57600080fd5b506101186101ca366004611264565b805160208183018101805160018252928201919093012091525481565b3480156101f357600080fd5b50600054610207906001600160a01b031681565b6040516001600160a01b039091168152602001610122565b34801561022b57600080fd5b5061011861023a366004611185565b6108fa565b34801561024b57600080fd5b5061011861025a36600461130e565b60036020526000908152604090205481565b34801561027857600080fd5b5061018d61028736600461130e565b61092b565b34801561029857600080fd5b506102ac6102a736600461130e565b61096f565b604051610122969594939291906114c3565b3480156102ca57600080fd5b506101186102d9366004611164565b60046020526000908152604090205481565b61018d610b4b565b3480156102ff57600080fd5b5061011861030e366004611164565b60056020526000908152604090205481565b34801561032c57600080fd5b5061020761033b366004611326565b610beb565b34801561034c57600080fd5b5061011860065481565b6000806103cc88888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8c018190048102820181019092528a815292508a9150899081908401838280828437600092019190915250610c2392505050565b90506001816040516103de919061139d565b90815260200160405180910390205460001461042f5760405162461bcd60e51b815260206004820152600b60248201526a135bd9195b08195e1a5cdd60aa1b60448201526064015b60405180910390fd5b6040518060c00160405280600654815260200189898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f8a01819004810282018101909252888152918101919089908990819084018382808284376000920191909152505050908252503360208083019190915260408051601f88018390048302810183018252878152920191908790879081908401838280828437600092018290525093855250504260209384015250600654815260028252604090208251815582820151805191926105289260018501929091019061106d565b506040820151805161054491600284019160209091019061106d565b5060608201516003820180546001600160a01b0319166001600160a01b039092169190911790556080820151805161058691600484019160209091019061106d565b5060a082015181600501559050506006546001826040516105a7919061139d565b9081526040519081900360200181209190915560065433917fc10402236a0d3324d97ba17b997a8c5cd276aa435f3ebc1b9b14866eaba568ac906105f6908c908c908c908c908c908c9061147a565b60405180910390a360068054925082906000610611836115bb565b9190505550509695505050505050565b6001600160a01b03811660009081526008602090815260409182902080548351818402810184019094528084526060939283018282801561068157602002820191906000526020600020905b81548152602001906001019080831161066d575b505050505090505b919050565b60008181526007602090815260409182902080548351818402810184019094528084526060939283018282801561068157602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116106d05750505050509050919050565b346107555760405162461bcd60e51b815260206004820152602660248201527f496e76616c6964207175616e74697479206f66206d6f64656c2075706c6f6164604482015265207374616b6560d01b6064820152608401610426565b3360009081526004602052604081208054349290610774908490611525565b9091555050604051349033907f9cd60d356d4a3dcc1fe32caaefdb28d5c57595181bcae18dd63934949464382290600090a3565b600054604051631846d2f560e31b81523360048201526001600160a01b039091169063c23697a89060240160206040518083038186803b1580156107eb57600080fd5b505afa1580156107ff573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061082391906111ae565b6108645760405162461bcd60e51b81526020600482015260126024820152714e6f6465206973206e6f742061637469766560701b6044820152606401610426565b600081815260026020526040902080546108b55760405162461bcd60e51b8152602060048201526012602482015271135bd9195b081a5cc81b9bdd08195e1a5cdd60721b6044820152606401610426565b6108bf8233610c4f565b6108c98233610d3a565b604051829033907fb81e8e03e6a54d2b07bc11d3cddad332ded9a7018c0e554bd78262a1d4cd20d190600090a35050565b6008602052816000526040600020818154811061091657600080fd5b90600052602060002001600091509150505481565b6109358133610e11565b61093f8133610f66565b604051819033907f271097aa25aca38eb083e68e5cafcbd8247d421aea1887c9630402cda1706b6390600090a350565b6002602052600090815260409020805460018201805491929161099190611580565b80601f01602080910402602001604051908101604052809291908181526020018280546109bd90611580565b8015610a0a5780601f106109df57610100808354040283529160200191610a0a565b820191906000526020600020905b8154815290600101906020018083116109ed57829003601f168201915b505050505090806002018054610a1f90611580565b80601f0160208091040260200160405190810160405280929190818152602001828054610a4b90611580565b8015610a985780601f10610a6d57610100808354040283529160200191610a98565b820191906000526020600020905b815481529060010190602001808311610a7b57829003601f168201915b505050600384015460048501805494956001600160a01b03909216949193509150610ac290611580565b80601f0160208091040260200160405190810160405280929190818152602001828054610aee90611580565b8015610b3b5780601f10610b1057610100808354040283529160200191610b3b565b820191906000526020600020905b815481529060010190602001808311610b1e57829003601f168201915b5050505050908060050154905086565b34610b985760405162461bcd60e51b815260206004820152601860248201527f496e76616c6964206d6f64656c2075736564207374616b6500000000000000006044820152606401610426565b3360009081526005602052604081208054349290610bb7908490611525565b9091555050604051349033907fd915a351426f3619ff44a69ba34964f8a74f0b302d08a457b3eac7ae0e90cd0590600090a3565b60076020528160005260406000208181548110610c0757600080fd5b6000918252602090912001546001600160a01b03169150829050565b60608282604051602001610c389291906113b9565b604051602081830303815290604052905092915050565b6000828152600760205260408120905b8154811015610d0757826001600160a01b0316828281548110610c9257634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b03161415610cf55760405162461bcd60e51b815260206004820181905260248201527f4d6f64656c20646973747269627574696f6e20616c72656164792065786973746044820152606401610426565b80610cff816115bb565b915050610c5f565b50505060009081526007602090815260408220805460018101825590835291200180546001600160a01b03191633179055565b6001600160a01b0381166000908152600860205260408120905b8154811015610de95783828281548110610d7e57634e487b7160e01b600052603260045260246000fd5b90600052602060002001541415610dd75760405162461bcd60e51b815260206004820152601d60248201527f4e6f6465206465706c6f796d656e7420616c72656164792065786973740000006044820152606401610426565b80610de1816115bb565b915050610d54565b5050336000908152600860209081526040822080546001810182559083529120019190915550565b6000828152600760205260408120905b8154811015610f6057826001600160a01b0316828281548110610e5457634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b03161415610f4e5781548290610e7f9060019061153d565b81548110610e9d57634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b0316828281548110610edb57634e487b7160e01b600052603260045260246000fd5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081805480610f2757634e487b7160e01b600052603160045260246000fd5b600082815260209020810160001990810180546001600160a01b0319169055019055610f60565b80610f58816115bb565b915050610e21565b50505050565b6001600160a01b0381166000908152600860205260408120905b8154811015610f605783828281548110610faa57634e487b7160e01b600052603260045260246000fd5b9060005260206000200154141561105b5781548290610fcb9060019061153d565b81548110610fe957634e487b7160e01b600052603260045260246000fd5b906000526020600020015482828154811061101457634e487b7160e01b600052603260045260246000fd5b90600052602060002001819055508180548061104057634e487b7160e01b600052603160045260246000fd5b60019003818190600052602060002001600090559055610f60565b80611065816115bb565b915050610f80565b82805461107990611580565b90600052602060002090601f01602090048101928261109b57600085556110e1565b82601f106110b457805160ff19168380011785556110e1565b828001600101855582156110e1579182015b828111156110e15782518255916020019190600101906110c6565b506110ed9291506110f1565b5090565b5b808211156110ed57600081556001016110f2565b80356001600160a01b038116811461068957600080fd5b60008083601f84011261112e578081fd5b50813567ffffffffffffffff811115611145578182fd5b60208301915083602082850101111561115d57600080fd5b9250929050565b600060208284031215611175578081fd5b61117e82611106565b9392505050565b60008060408385031215611197578081fd5b6111a083611106565b946020939093013593505050565b6000602082840312156111bf578081fd5b8151801515811461117e578182fd5b600080600080600080606087890312156111e6578182fd5b863567ffffffffffffffff808211156111fd578384fd5b6112098a838b0161111d565b90985096506020890135915080821115611221578384fd5b61122d8a838b0161111d565b90965094506040890135915080821115611245578384fd5b5061125289828a0161111d565b979a9699509497509295939492505050565b600060208284031215611275578081fd5b813567ffffffffffffffff8082111561128c578283fd5b818401915084601f83011261129f578283fd5b8135818111156112b1576112b16115ec565b604051601f8201601f19908116603f011681019083821181831017156112d9576112d96115ec565b816040528281528760208487010111156112f1578586fd5b826020860160208301379182016020019490945295945050505050565b60006020828403121561131f578081fd5b5035919050565b60008060408385031215611338578182fd5b50508035926020909101359150565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b60008151808452611389816020860160208601611554565b601f01601f19169290920160200192915050565b600082516113af818460208701611554565b9190910192915050565b600083516113cb818460208801611554565b602f60f81b90830190815283516113e9816001840160208801611554565b01600101949350505050565b6020808252825182820181905260009190848201906040850190845b818110156114365783516001600160a01b031683529284019291840191600101611411565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156114365783518352928401929184019160010161145e565b60006060825261148e60608301888a611347565b82810360208401526114a1818789611347565b905082810360408401526114b6818587611347565b9998505050505050505050565b600087825260c060208301526114dc60c0830188611371565b82810360408401526114ee8188611371565b6001600160a01b0387166060850152838103608085015290506115118186611371565b9150508260a0830152979650505050505050565b60008219821115611538576115386115d6565b500190565b60008282101561154f5761154f6115d6565b500390565b60005b8381101561156f578181015183820152602001611557565b83811115610f605750506000910152565b600181811c9082168061159457607f821691505b602082108114156115b557634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156115cf576115cf6115d6565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea264697066735822122087f6578f80ecf5aeb613be9bd5ed4fe9565370e54551d0dd2d95cfb152aecaab64736f6c63430008030033",
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

// ModelUploadStakes is a free data retrieval call binding the contract method 0xa3217acc.
//
// Solidity: function modelUploadStakes(address ) view returns(uint256)
func (_AIModels *AIModelsCaller) ModelUploadStakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "modelUploadStakes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModelUploadStakes is a free data retrieval call binding the contract method 0xa3217acc.
//
// Solidity: function modelUploadStakes(address ) view returns(uint256)
func (_AIModels *AIModelsSession) ModelUploadStakes(arg0 common.Address) (*big.Int, error) {
	return _AIModels.Contract.ModelUploadStakes(&_AIModels.CallOpts, arg0)
}

// ModelUploadStakes is a free data retrieval call binding the contract method 0xa3217acc.
//
// Solidity: function modelUploadStakes(address ) view returns(uint256)
func (_AIModels *AIModelsCallerSession) ModelUploadStakes(arg0 common.Address) (*big.Int, error) {
	return _AIModels.Contract.ModelUploadStakes(&_AIModels.CallOpts, arg0)
}

// ModelUsedStakes is a free data retrieval call binding the contract method 0xbaf70a46.
//
// Solidity: function modelUsedStakes(address ) view returns(uint256)
func (_AIModels *AIModelsCaller) ModelUsedStakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "modelUsedStakes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModelUsedStakes is a free data retrieval call binding the contract method 0xbaf70a46.
//
// Solidity: function modelUsedStakes(address ) view returns(uint256)
func (_AIModels *AIModelsSession) ModelUsedStakes(arg0 common.Address) (*big.Int, error) {
	return _AIModels.Contract.ModelUsedStakes(&_AIModels.CallOpts, arg0)
}

// ModelUsedStakes is a free data retrieval call binding the contract method 0xbaf70a46.
//
// Solidity: function modelUsedStakes(address ) view returns(uint256)
func (_AIModels *AIModelsCallerSession) ModelUsedStakes(arg0 common.Address) (*big.Int, error) {
	return _AIModels.Contract.ModelUsedStakes(&_AIModels.CallOpts, arg0)
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

// StakeModelUpload is a paid mutator transaction binding the contract method 0x6268e03e.
//
// Solidity: function stakeModelUpload() payable returns()
func (_AIModels *AIModelsTransactor) StakeModelUpload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIModels.contract.Transact(opts, "stakeModelUpload")
}

// StakeModelUpload is a paid mutator transaction binding the contract method 0x6268e03e.
//
// Solidity: function stakeModelUpload() payable returns()
func (_AIModels *AIModelsSession) StakeModelUpload() (*types.Transaction, error) {
	return _AIModels.Contract.StakeModelUpload(&_AIModels.TransactOpts)
}

// StakeModelUpload is a paid mutator transaction binding the contract method 0x6268e03e.
//
// Solidity: function stakeModelUpload() payable returns()
func (_AIModels *AIModelsTransactorSession) StakeModelUpload() (*types.Transaction, error) {
	return _AIModels.Contract.StakeModelUpload(&_AIModels.TransactOpts)
}

// StakeModelUsed is a paid mutator transaction binding the contract method 0xab96d7bc.
//
// Solidity: function stakeModelUsed() payable returns()
func (_AIModels *AIModelsTransactor) StakeModelUsed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIModels.contract.Transact(opts, "stakeModelUsed")
}

// StakeModelUsed is a paid mutator transaction binding the contract method 0xab96d7bc.
//
// Solidity: function stakeModelUsed() payable returns()
func (_AIModels *AIModelsSession) StakeModelUsed() (*types.Transaction, error) {
	return _AIModels.Contract.StakeModelUsed(&_AIModels.TransactOpts)
}

// StakeModelUsed is a paid mutator transaction binding the contract method 0xab96d7bc.
//
// Solidity: function stakeModelUsed() payable returns()
func (_AIModels *AIModelsTransactorSession) StakeModelUsed() (*types.Transaction, error) {
	return _AIModels.Contract.StakeModelUsed(&_AIModels.TransactOpts)
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

// AIModelsModelUploadStakedIterator is returned from FilterModelUploadStaked and is used to iterate over the raw logs and unpacked data for ModelUploadStaked events raised by the AIModels contract.
type AIModelsModelUploadStakedIterator struct {
	Event *AIModelsModelUploadStaked // Event containing the contract specifics and raw log

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
func (it *AIModelsModelUploadStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIModelsModelUploadStaked)
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
		it.Event = new(AIModelsModelUploadStaked)
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
func (it *AIModelsModelUploadStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIModelsModelUploadStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIModelsModelUploadStaked represents a ModelUploadStaked event raised by the AIModels contract.
type AIModelsModelUploadStaked struct {
	Uploader common.Address
	Stake    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterModelUploadStaked is a free log retrieval operation binding the contract event 0x9cd60d356d4a3dcc1fe32caaefdb28d5c57595181bcae18dd639349494643822.
//
// Solidity: event ModelUploadStaked(address indexed uploader, uint256 indexed stake)
func (_AIModels *AIModelsFilterer) FilterModelUploadStaked(opts *bind.FilterOpts, uploader []common.Address, stake []*big.Int) (*AIModelsModelUploadStakedIterator, error) {

	var uploaderRule []interface{}
	for _, uploaderItem := range uploader {
		uploaderRule = append(uploaderRule, uploaderItem)
	}
	var stakeRule []interface{}
	for _, stakeItem := range stake {
		stakeRule = append(stakeRule, stakeItem)
	}

	logs, sub, err := _AIModels.contract.FilterLogs(opts, "ModelUploadStaked", uploaderRule, stakeRule)
	if err != nil {
		return nil, err
	}
	return &AIModelsModelUploadStakedIterator{contract: _AIModels.contract, event: "ModelUploadStaked", logs: logs, sub: sub}, nil
}

// WatchModelUploadStaked is a free log subscription operation binding the contract event 0x9cd60d356d4a3dcc1fe32caaefdb28d5c57595181bcae18dd639349494643822.
//
// Solidity: event ModelUploadStaked(address indexed uploader, uint256 indexed stake)
func (_AIModels *AIModelsFilterer) WatchModelUploadStaked(opts *bind.WatchOpts, sink chan<- *AIModelsModelUploadStaked, uploader []common.Address, stake []*big.Int) (event.Subscription, error) {

	var uploaderRule []interface{}
	for _, uploaderItem := range uploader {
		uploaderRule = append(uploaderRule, uploaderItem)
	}
	var stakeRule []interface{}
	for _, stakeItem := range stake {
		stakeRule = append(stakeRule, stakeItem)
	}

	logs, sub, err := _AIModels.contract.WatchLogs(opts, "ModelUploadStaked", uploaderRule, stakeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIModelsModelUploadStaked)
				if err := _AIModels.contract.UnpackLog(event, "ModelUploadStaked", log); err != nil {
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

// ParseModelUploadStaked is a log parse operation binding the contract event 0x9cd60d356d4a3dcc1fe32caaefdb28d5c57595181bcae18dd639349494643822.
//
// Solidity: event ModelUploadStaked(address indexed uploader, uint256 indexed stake)
func (_AIModels *AIModelsFilterer) ParseModelUploadStaked(log types.Log) (*AIModelsModelUploadStaked, error) {
	event := new(AIModelsModelUploadStaked)
	if err := _AIModels.contract.UnpackLog(event, "ModelUploadStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIModelsModelUploadUnstakedIterator is returned from FilterModelUploadUnstaked and is used to iterate over the raw logs and unpacked data for ModelUploadUnstaked events raised by the AIModels contract.
type AIModelsModelUploadUnstakedIterator struct {
	Event *AIModelsModelUploadUnstaked // Event containing the contract specifics and raw log

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
func (it *AIModelsModelUploadUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIModelsModelUploadUnstaked)
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
		it.Event = new(AIModelsModelUploadUnstaked)
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
func (it *AIModelsModelUploadUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIModelsModelUploadUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIModelsModelUploadUnstaked represents a ModelUploadUnstaked event raised by the AIModels contract.
type AIModelsModelUploadUnstaked struct {
	Uploader common.Address
	Stake    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterModelUploadUnstaked is a free log retrieval operation binding the contract event 0xf5d3791607f24f018982a65d6f2d56a93a557cc13e0415e92511322c6bbc376c.
//
// Solidity: event ModelUploadUnstaked(address indexed uploader, uint256 indexed stake)
func (_AIModels *AIModelsFilterer) FilterModelUploadUnstaked(opts *bind.FilterOpts, uploader []common.Address, stake []*big.Int) (*AIModelsModelUploadUnstakedIterator, error) {

	var uploaderRule []interface{}
	for _, uploaderItem := range uploader {
		uploaderRule = append(uploaderRule, uploaderItem)
	}
	var stakeRule []interface{}
	for _, stakeItem := range stake {
		stakeRule = append(stakeRule, stakeItem)
	}

	logs, sub, err := _AIModels.contract.FilterLogs(opts, "ModelUploadUnstaked", uploaderRule, stakeRule)
	if err != nil {
		return nil, err
	}
	return &AIModelsModelUploadUnstakedIterator{contract: _AIModels.contract, event: "ModelUploadUnstaked", logs: logs, sub: sub}, nil
}

// WatchModelUploadUnstaked is a free log subscription operation binding the contract event 0xf5d3791607f24f018982a65d6f2d56a93a557cc13e0415e92511322c6bbc376c.
//
// Solidity: event ModelUploadUnstaked(address indexed uploader, uint256 indexed stake)
func (_AIModels *AIModelsFilterer) WatchModelUploadUnstaked(opts *bind.WatchOpts, sink chan<- *AIModelsModelUploadUnstaked, uploader []common.Address, stake []*big.Int) (event.Subscription, error) {

	var uploaderRule []interface{}
	for _, uploaderItem := range uploader {
		uploaderRule = append(uploaderRule, uploaderItem)
	}
	var stakeRule []interface{}
	for _, stakeItem := range stake {
		stakeRule = append(stakeRule, stakeItem)
	}

	logs, sub, err := _AIModels.contract.WatchLogs(opts, "ModelUploadUnstaked", uploaderRule, stakeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIModelsModelUploadUnstaked)
				if err := _AIModels.contract.UnpackLog(event, "ModelUploadUnstaked", log); err != nil {
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

// ParseModelUploadUnstaked is a log parse operation binding the contract event 0xf5d3791607f24f018982a65d6f2d56a93a557cc13e0415e92511322c6bbc376c.
//
// Solidity: event ModelUploadUnstaked(address indexed uploader, uint256 indexed stake)
func (_AIModels *AIModelsFilterer) ParseModelUploadUnstaked(log types.Log) (*AIModelsModelUploadUnstaked, error) {
	event := new(AIModelsModelUploadUnstaked)
	if err := _AIModels.contract.UnpackLog(event, "ModelUploadUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIModelsModelUsedStakedIterator is returned from FilterModelUsedStaked and is used to iterate over the raw logs and unpacked data for ModelUsedStaked events raised by the AIModels contract.
type AIModelsModelUsedStakedIterator struct {
	Event *AIModelsModelUsedStaked // Event containing the contract specifics and raw log

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
func (it *AIModelsModelUsedStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIModelsModelUsedStaked)
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
		it.Event = new(AIModelsModelUsedStaked)
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
func (it *AIModelsModelUsedStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIModelsModelUsedStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIModelsModelUsedStaked represents a ModelUsedStaked event raised by the AIModels contract.
type AIModelsModelUsedStaked struct {
	ModelUser common.Address
	Stake     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModelUsedStaked is a free log retrieval operation binding the contract event 0xd915a351426f3619ff44a69ba34964f8a74f0b302d08a457b3eac7ae0e90cd05.
//
// Solidity: event ModelUsedStaked(address indexed modelUser, uint256 indexed stake)
func (_AIModels *AIModelsFilterer) FilterModelUsedStaked(opts *bind.FilterOpts, modelUser []common.Address, stake []*big.Int) (*AIModelsModelUsedStakedIterator, error) {

	var modelUserRule []interface{}
	for _, modelUserItem := range modelUser {
		modelUserRule = append(modelUserRule, modelUserItem)
	}
	var stakeRule []interface{}
	for _, stakeItem := range stake {
		stakeRule = append(stakeRule, stakeItem)
	}

	logs, sub, err := _AIModels.contract.FilterLogs(opts, "ModelUsedStaked", modelUserRule, stakeRule)
	if err != nil {
		return nil, err
	}
	return &AIModelsModelUsedStakedIterator{contract: _AIModels.contract, event: "ModelUsedStaked", logs: logs, sub: sub}, nil
}

// WatchModelUsedStaked is a free log subscription operation binding the contract event 0xd915a351426f3619ff44a69ba34964f8a74f0b302d08a457b3eac7ae0e90cd05.
//
// Solidity: event ModelUsedStaked(address indexed modelUser, uint256 indexed stake)
func (_AIModels *AIModelsFilterer) WatchModelUsedStaked(opts *bind.WatchOpts, sink chan<- *AIModelsModelUsedStaked, modelUser []common.Address, stake []*big.Int) (event.Subscription, error) {

	var modelUserRule []interface{}
	for _, modelUserItem := range modelUser {
		modelUserRule = append(modelUserRule, modelUserItem)
	}
	var stakeRule []interface{}
	for _, stakeItem := range stake {
		stakeRule = append(stakeRule, stakeItem)
	}

	logs, sub, err := _AIModels.contract.WatchLogs(opts, "ModelUsedStaked", modelUserRule, stakeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIModelsModelUsedStaked)
				if err := _AIModels.contract.UnpackLog(event, "ModelUsedStaked", log); err != nil {
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

// ParseModelUsedStaked is a log parse operation binding the contract event 0xd915a351426f3619ff44a69ba34964f8a74f0b302d08a457b3eac7ae0e90cd05.
//
// Solidity: event ModelUsedStaked(address indexed modelUser, uint256 indexed stake)
func (_AIModels *AIModelsFilterer) ParseModelUsedStaked(log types.Log) (*AIModelsModelUsedStaked, error) {
	event := new(AIModelsModelUsedStaked)
	if err := _AIModels.contract.UnpackLog(event, "ModelUsedStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIModelsModelUsedUnstakedIterator is returned from FilterModelUsedUnstaked and is used to iterate over the raw logs and unpacked data for ModelUsedUnstaked events raised by the AIModels contract.
type AIModelsModelUsedUnstakedIterator struct {
	Event *AIModelsModelUsedUnstaked // Event containing the contract specifics and raw log

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
func (it *AIModelsModelUsedUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIModelsModelUsedUnstaked)
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
		it.Event = new(AIModelsModelUsedUnstaked)
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
func (it *AIModelsModelUsedUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIModelsModelUsedUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIModelsModelUsedUnstaked represents a ModelUsedUnstaked event raised by the AIModels contract.
type AIModelsModelUsedUnstaked struct {
	ModelUser common.Address
	Stake     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModelUsedUnstaked is a free log retrieval operation binding the contract event 0x286fcb14163c58337da50510614f4b2259e3d9b86507935bd78af0eb1fe4a5d1.
//
// Solidity: event ModelUsedUnstaked(address indexed modelUser, uint256 indexed stake)
func (_AIModels *AIModelsFilterer) FilterModelUsedUnstaked(opts *bind.FilterOpts, modelUser []common.Address, stake []*big.Int) (*AIModelsModelUsedUnstakedIterator, error) {

	var modelUserRule []interface{}
	for _, modelUserItem := range modelUser {
		modelUserRule = append(modelUserRule, modelUserItem)
	}
	var stakeRule []interface{}
	for _, stakeItem := range stake {
		stakeRule = append(stakeRule, stakeItem)
	}

	logs, sub, err := _AIModels.contract.FilterLogs(opts, "ModelUsedUnstaked", modelUserRule, stakeRule)
	if err != nil {
		return nil, err
	}
	return &AIModelsModelUsedUnstakedIterator{contract: _AIModels.contract, event: "ModelUsedUnstaked", logs: logs, sub: sub}, nil
}

// WatchModelUsedUnstaked is a free log subscription operation binding the contract event 0x286fcb14163c58337da50510614f4b2259e3d9b86507935bd78af0eb1fe4a5d1.
//
// Solidity: event ModelUsedUnstaked(address indexed modelUser, uint256 indexed stake)
func (_AIModels *AIModelsFilterer) WatchModelUsedUnstaked(opts *bind.WatchOpts, sink chan<- *AIModelsModelUsedUnstaked, modelUser []common.Address, stake []*big.Int) (event.Subscription, error) {

	var modelUserRule []interface{}
	for _, modelUserItem := range modelUser {
		modelUserRule = append(modelUserRule, modelUserItem)
	}
	var stakeRule []interface{}
	for _, stakeItem := range stake {
		stakeRule = append(stakeRule, stakeItem)
	}

	logs, sub, err := _AIModels.contract.WatchLogs(opts, "ModelUsedUnstaked", modelUserRule, stakeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIModelsModelUsedUnstaked)
				if err := _AIModels.contract.UnpackLog(event, "ModelUsedUnstaked", log); err != nil {
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

// ParseModelUsedUnstaked is a log parse operation binding the contract event 0x286fcb14163c58337da50510614f4b2259e3d9b86507935bd78af0eb1fe4a5d1.
//
// Solidity: event ModelUsedUnstaked(address indexed modelUser, uint256 indexed stake)
func (_AIModels *AIModelsFilterer) ParseModelUsedUnstaked(log types.Log) (*AIModelsModelUsedUnstaked, error) {
	event := new(AIModelsModelUsedUnstaked)
	if err := _AIModels.contract.UnpackLog(event, "ModelUsedUnstaked", log); err != nil {
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
