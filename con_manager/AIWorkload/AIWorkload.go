// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AIWorkload

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

// AIWorkloadWorkLoad is an auto generated low-level Go binding around an user-defined struct.
type AIWorkloadWorkLoad struct {
	TotalWorkload   *big.Int
	SettledWorkload *big.Int
}

// AIWorkloadWorkload is an auto generated low-level Go binding around an user-defined struct.
type AIWorkloadWorkload struct {
	EpochId   *big.Int
	Workload  *big.Int
	Timestamp *big.Int
	ModelId   *big.Int
	Reporter  common.Address
	Worker    common.Address
}

// ModelSettleWorkload is an auto generated low-level Go binding around an user-defined struct.
type ModelSettleWorkload struct {
	ModelId  *big.Int
	Workload *big.Int
}

// NodeSettleWorkload is an auto generated low-level Go binding around an user-defined struct.
type NodeSettleWorkload struct {
	Node     common.Address
	Workload *big.Int
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	R [32]byte
	S [32]byte
	V uint8
}

// AIWorkloadMetaData contains all meta data concerning the AIWorkload contract.
var AIWorkloadMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_modelRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"WorkloadReported\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"}],\"name\":\"getLastEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"getNodeWorkload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"internalType\":\"structAIWorkload.Workload\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"getTotalModelWorkload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalWorkload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledWorkload\",\"type\":\"uint256\"}],\"internalType\":\"structAIWorkload.WorkLoad\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"getTotalWorkReport\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalWorkload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledWorkload\",\"type\":\"uint256\"}],\"internalType\":\"structAIWorkload.WorkLoad\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getTotalWorkerWorkload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalWorkload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledWorkload\",\"type\":\"uint256\"}],\"internalType\":\"structAIWorkload.WorkLoad\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSettlementTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modelRegistry\",\"outputs\":[{\"internalType\":\"contractAIModels\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeRegistry\",\"outputs\":[{\"internalType\":\"contractNodesRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"reportWorkload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sessions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastEpochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settleRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"}],\"internalType\":\"structNodeSettleWorkload[]\",\"name\":\"settledWorkers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"}],\"internalType\":\"structModelSettleWorkload[]\",\"name\":\"settledModels\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"}],\"internalType\":\"structNodeSettleWorkload[]\",\"name\":\"settledReporters\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIStake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052610e10600c553480156200001757600080fd5b5060405162001f6038038062001f608339810160408190526200003a91620001aa565b6001600160a01b038316620000965760405162461bcd60e51b815260206004820152601560248201527f496e76616c6964206e6f6465207265676973747279000000000000000000000060448201526064015b60405180910390fd5b6001600160a01b038216620000ee5760405162461bcd60e51b815260206004820152601660248201527f496e76616c6964206d6f64656c2072656769737472790000000000000000000060448201526064016200008d565b6001600160a01b038116620001465760405162461bcd60e51b815260206004820152601360248201527f496e76616c6964207374616b6520746f6b656e0000000000000000000000000060448201526064016200008d565b600180546001600160a01b039485166001600160a01b031991821617909155600280549385169382169390931790925542600d55600e8054919093169116179055620001f3565b80516001600160a01b0381168114620001a557600080fd5b919050565b600080600060608486031215620001bf578283fd5b620001ca846200018d565b9250620001da602085016200018d565b9150620001ea604085016200018d565b90509250925092565b611d5d80620002036000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80637ec4142c1161008c578063d14c9fa711610066578063d14c9fa714610201578063d9564c3614610218578063d9b5c4a514610281578063e926805e14610294576100cf565b80637ec4142c146101b857806383c4b7a3146101c1578063ad1f5717146101e1576100cf565b806309de6d2f146100d45780630f17c347146100e957806320737af1146101055780632ba36f78146101255780632c2e139b1461015057806351ed6a30146101a5575b600080fd5b6100e76100e236600461183b565b6102e9565b005b6100f2600d5481565b6040519081526020015b60405180910390f35b6101186101133660046119bb565b610798565b6040516100fc9190611be4565b600254610138906001600160a01b031681565b6040516001600160a01b0390911681526020016100fc565b61011861015e36600461181f565b6040805180820190915260008082526020820152506001600160a01b0316600090815260096020908152604091829020825180840190935280548352600101549082015290565b600e54610138906001600160a01b031681565b6100f2600c5481565b6100f26101cf3660046119bb565b60006020819052908152604090205481565b6100f26101ef3660046119bb565b60009081526020819052604090205490565b6102096107d2565b6040516100fc93929190611b67565b61022b610226366004611a7c565b610f32565b6040516100fc9190815181526020808301519082015260408083015190820152606080830151908201526080808301516001600160a01b039081169183019190915260a092830151169181019190915260c00190565b600154610138906001600160a01b031681565b6101186102a236600461181f565b6040805180820190915260008082526020820152506001600160a01b0316600090815260036020908152604091829020825180840190935280548352600101549082015290565b6001600160a01b03871661033c5760405162461bcd60e51b8152602060048201526015602482015274496e76616c6964206f776e6572206164647265737360581b60448201526064015b60405180910390fd5b600086116103975760405162461bcd60e51b815260206004820152602260248201527f576f726b6c6f6164206d7573742062652067726561746572207468616e207a65604482015261726f60f01b6064820152608401610333565b60038110156103f65760405162461bcd60e51b815260206004820152602560248201527f4c656e677468206f66207369676e617475726573206d757374206d6f7265207460448201526468616e203360d81b6064820152608401610333565b600254604051634dd86b0960e11b8152600481018790526000916001600160a01b031690639bb0d6129060240160006040518083038186803b15801561043b57600080fd5b505afa15801561044f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261047791908101906119d3565b505050505090508581146104bf5760405162461bcd60e51b815260206004820152600f60248201526e135bd9195b081b9bdd08195e1a5cdd608a1b6044820152606401610333565b604080516001600160a01b038a166020820152908101889052606081018790526080810186905260a0810185905261050e908990339060c0016040516020818303038152906040528686610fed565b61054e5760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b6044820152606401610333565b6000858152602081905260409020805485116105a15760405162461bcd60e51b815260206004820152601260248201527122b837b1b41037baba1037b31037b93232b960711b6044820152606401610333565b6040518060c00160405280868152602001898152602001428152602001888152602001336001600160a01b031681526020018a6001600160a01b03168152508160010160008781526020019081526020016000206000820151816000015560208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060a08201518160050160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050506000600360008b6001600160a01b03166001600160a01b031681526020019081526020016000209050888160000160008282546106b79190611c2b565b909155506106c8905060048b611328565b506000888152600660205260408120805490918b918391906106eb908490611c2b565b909155506106fc905060078a611344565b50336000908152600960205260408120805490918c91839190610720908490611c2b565b909155506107319050600a33611328565b50878455604080516001600160a01b038e168152602081018a90529081018c9052606081018b905233908a907f59164c09eed89694722e80b610592178b04fbe4a72b982bf4503c8ab426f0ef69060800160405180910390a3505050505050505050505050565b604080518082018252600080825260209182018190528381526006825282902082518084019093528054835260010154908201525b919050565b6060806060600c54600d546107e79190611c2b565b42101561082f5760405162461bcd60e51b815260206004820152601660248201527514d95d1d1b195b595b9d081b9bdd08191d59481e595d60521b6044820152606401610333565b600061083b6004611350565b6001600160401b0381111561086057634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156108a557816020015b604080518082019091526000808252602082015281526020019060019003908161087e5790505b5090506000805b6108b66004611350565b8110156109a75760006003816108cd60048561135a565b6001600160a01b03166001600160a01b03168152602001908152602001600020905080600101548160000154116109045750610995565b61090f60048361135a565b84848151811061092f57634e487b7160e01b600052603260045260246000fd5b60209081029190910101516001600160a01b039091169052600181015481546109589190611c57565b84848151811061097857634e487b7160e01b600052603260045260246000fd5b60209081029190910181015101528261099081611c9e565b935050505b8061099f81611c9e565b9150506108ac565b50806001600160401b038111156109ce57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610a1357816020015b60408051808201909152600080825260208201528152602001906001900390816109ec5790505b50945060005b81811015610a8757828181518110610a4157634e487b7160e01b600052603260045260246000fd5b6020026020010151868281518110610a6957634e487b7160e01b600052603260045260246000fd5b60200260200101819052508080610a7f90611c9e565b915050610a19565b506000610a946007611350565b6001600160401b03811115610ab957634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610afe57816020015b6040805180820190915260008082526020820152815260200190600190039081610ad75790505b5090506000915060005b610b126007611350565b811015610be5576000600681610b2960078561135a565b815260200190815260200160002090508060010154816000015411610b4e5750610bd3565b610b5960078361135a565b838581518110610b7957634e487b7160e01b600052603260045260246000fd5b60209081029190910101515260018101548154610b969190611c57565b838581518110610bb657634e487b7160e01b600052603260045260246000fd5b602090810291909101810151015283610bce81611c9e565b945050505b80610bdd81611c9e565b915050610b08565b50816001600160401b03811115610c0c57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610c5157816020015b6040805180820190915260008082526020820152815260200190600190039081610c2a5790505b50945060005b82811015610cc557818181518110610c7f57634e487b7160e01b600052603260045260246000fd5b6020026020010151868281518110610ca757634e487b7160e01b600052603260045260246000fd5b60200260200101819052508080610cbd90611c9e565b915050610c57565b506000610cd2600a611350565b6001600160401b03811115610cf757634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610d3c57816020015b6040805180820190915260008082526020820152815260200190600190039081610d155790505b5090506000925060005b610d50600a611350565b811015610e41576000600981610d67600a8561135a565b6001600160a01b03166001600160a01b0316815260200190815260200160002090508060010154816000015411610d9e5750610e2f565b610da9600a8361135a565b838681518110610dc957634e487b7160e01b600052603260045260246000fd5b60209081029190910101516001600160a01b03909116905260018101548154610df29190611c57565b838681518110610e1257634e487b7160e01b600052603260045260246000fd5b602090810291909101810151015284610e2a81611c9e565b955050505b80610e3981611c9e565b915050610d46565b50826001600160401b03811115610e6857634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610ead57816020015b6040805180820190915260008082526020820152815260200190600190039081610e865790505b50945060005b83811015610f2157818181518110610edb57634e487b7160e01b600052603260045260246000fd5b6020026020010151868281518110610f0357634e487b7160e01b600052603260045260246000fd5b60200260200101819052508080610f1990611c9e565b915050610eb3565b5042600d8190555050505050909192565b610f7d6040518060c001604052806000815260200160008152602001600081526020016000815260200160006001600160a01b0316815260200160006001600160a01b031681525090565b506000828152602081815260408083208484526001908101835292819020815160c0810183528154815293810154928401929092526002820154908301526003810154606083015260048101546001600160a01b0390811660808401526005909101541660a08201525b92915050565b600080826001600160401b0381111561101657634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561103f578160200160208202803683370190505b50905060008080805b868110156112d85760008089898481811061107357634e487b7160e01b600052603260045260246000fd5b905060600201600001358a8a8581811061109d57634e487b7160e01b600052603260045260246000fd5b905060600201602001358b8b868181106110c757634e487b7160e01b600052603260045260246000fd5b90506060020160400160208101906110df9190611a9d565b60405160200161110f93929190928352602083019190915260f81b6001600160f81b031916604082015260410190565b6040516020818303038152906040529050600061113461112e8d611366565b836113a1565b506001546040516330af0bbf60e21b81526001600160a01b03808416600483015292935091169063c2bc2efc9060240160006040518083038186803b15801561117c57600080fd5b505afa158015611190573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111b891908101906118e4565b606001516111c8575050506112c6565b60005b8781101561122f57816001600160a01b03168982815181106111fd57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b0316141561121d576001935061122f565b8061122781611c9e565b9150506111cb565b50821561123e575050506112c6565b8d6001600160a01b0316816001600160a01b0316141561125d57600194505b8c6001600160a01b0316816001600160a01b0316141561127c57600195505b8088888151811061129d57634e487b7160e01b600052603260045260246000fd5b6001600160a01b03909216602092830291909101909101526112c0600188611c2b565b96505050505b806112d081611c9e565b915050611048565b5060026112e6876001611c2b565b6112f09190611c43565b8310806112fb575080155b80611304575081155b1561131657600094505050505061131f565b60019450505050505b95945050505050565b600061133d836001600160a01b038416611411565b9392505050565b600061133d8383611411565b6000610fe7825490565b600061133d8383611460565b60006113728251611498565b82604051602001611384929190611b0c565b604051602081830303815290604052805190602001209050919050565b6000808251604114156113d85760208301516040840151606085015160001a6113cc878285856115ba565b9450945050505061140a565b82516040141561140257602083015160408401516113f78683836116a7565b93509350505061140a565b506000905060025b9250929050565b600081815260018301602052604081205461145857508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610fe7565b506000610fe7565b600082600001828154811061148557634e487b7160e01b600052603260045260246000fd5b9060005260206000200154905092915050565b6060816114bd57506040805180820190915260018152600360fc1b60208201526107cd565b8160005b81156114e757806114d181611c9e565b91506114e09050600a83611c43565b91506114c1565b6000816001600160401b0381111561150f57634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015611539576020820181803683370190505b5090505b84156115b25761154e600183611c57565b915061155b600a86611cb9565b611566906030611c2b565b60f81b81838151811061158957634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a9053506115ab600a86611c43565b945061153d565b949350505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156115f1575060009050600361169e565b8460ff16601b1415801561160957508460ff16601c14155b1561161a575060009050600461169e565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561166e573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166116975760006001925092505061169e565b9150600090505b94509492505050565b6000806001600160ff1b038316816116c460ff86901c601b611c2b565b90506116d2878288856115ba565b935093505050935093915050565b80516107cd81611d0f565b600082601f8301126116fb578081fd5b815160206001600160401b038083111561171757611717611cf9565b611725828460051b01611bfb565b83815282810190868401865b868110156117a05781518901606080601f19838e0301121561175157898afd5b61175a81611bfb565b888301518881111561176a578b8cfd5b6117788e8b838701016117be565b8252506040838101518a83015291909201519082015284529285019290850190600101611731565b509098975050505050505050565b805180151581146107cd57600080fd5b600082601f8301126117ce578081fd5b81516001600160401b038111156117e7576117e7611cf9565b6117fa601f8201601f1916602001611bfb565b81815284602083860101111561180e578283fd5b6115b2826020830160208701611c6e565b600060208284031215611830578081fd5b813561133d81611d0f565b600080600080600080600060c0888a031215611855578283fd5b873561186081611d0f565b96506020880135955060408801359450606088013593506080880135925060a08801356001600160401b0380821115611897578384fd5b818a0191508a601f8301126118aa578384fd5b8135818111156118b8578485fd5b8b60206060830285010111156118cc578485fd5b60208301945080935050505092959891949750929550565b6000602082840312156118f5578081fd5b81516001600160401b038082111561190b578283fd5b9083019060e0828603121561191e578283fd5b61192860e0611bfb565b611931836116e0565b8152602083015182811115611944578485fd5b611950878286016117be565b6020830152506040830151604082015261196c606084016117ae565b6060820152608083015182811115611982578485fd5b61198e878286016116eb565b6080830152506119a060a084016116e0565b60a082015260c083015160c082015280935050505092915050565b6000602082840312156119cc578081fd5b5035919050565b60008060008060008060c087890312156119eb578182fd5b8651955060208701516001600160401b0380821115611a08578384fd5b611a148a838b016117be565b96506040890151915080821115611a29578384fd5b611a358a838b016117be565b955060608901519150611a4782611d0f565b608089015191945080821115611a5b578384fd5b50611a6889828a016117be565b92505060a087015190509295509295509295565b60008060408385031215611a8e578182fd5b50508035926020909101359150565b600060208284031215611aae578081fd5b813560ff8116811461133d578182fd5b6000815180845260208085019450808401835b83811015611b0157815180516001600160a01b031688528301518388015260409096019590820190600101611ad1565b509495945050505050565b60007f19457468657265756d205369676e6564204d6573736167653a0a00000000000082528351611b4481601a850160208801611c6e565b835190830190611b5b81601a840160208801611c6e565b01601a01949350505050565b600060608252611b7a6060830186611abe565b828103602084810191909152855180835286820192820190845b81811015611bc457611bb183865180518252602090810151910152565b9383019360409290920191600101611b94565b50508481036040860152611bd88187611abe565b98975050505050505050565b815181526020808301519082015260408101610fe7565b604051601f8201601f191681016001600160401b0381118282101715611c2357611c23611cf9565b604052919050565b60008219821115611c3e57611c3e611ccd565b500190565b600082611c5257611c52611ce3565b500490565b600082821015611c6957611c69611ccd565b500390565b60005b83811015611c89578181015183820152602001611c71565b83811115611c98576000848401525b50505050565b6000600019821415611cb257611cb2611ccd565b5060010190565b600082611cc857611cc8611ce3565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114611d2457600080fd5b5056fea2646970667358221220f44d5f96fe74a5c36b5638cd67eaf0b331579181e453765f5b53ad11f6070b9764736f6c63430008030033",
}

// AIWorkloadABI is the input ABI used to generate the binding from.
// Deprecated: Use AIWorkloadMetaData.ABI instead.
var AIWorkloadABI = AIWorkloadMetaData.ABI

// AIWorkloadBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AIWorkloadMetaData.Bin instead.
var AIWorkloadBin = AIWorkloadMetaData.Bin

// DeployAIWorkload deploys a new Ethereum contract, binding an instance of AIWorkload to it.
func DeployAIWorkload(auth *bind.TransactOpts, backend bind.ContractBackend, _nodeRegistry common.Address, _modelRegistry common.Address, _stakeToken common.Address) (common.Address, *types.Transaction, *AIWorkload, error) {
	parsed, err := AIWorkloadMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AIWorkloadBin), backend, _nodeRegistry, _modelRegistry, _stakeToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AIWorkload{AIWorkloadCaller: AIWorkloadCaller{contract: contract}, AIWorkloadTransactor: AIWorkloadTransactor{contract: contract}, AIWorkloadFilterer: AIWorkloadFilterer{contract: contract}}, nil
}

// AIWorkload is an auto generated Go binding around an Ethereum contract.
type AIWorkload struct {
	AIWorkloadCaller     // Read-only binding to the contract
	AIWorkloadTransactor // Write-only binding to the contract
	AIWorkloadFilterer   // Log filterer for contract events
}

// AIWorkloadCaller is an auto generated read-only Go binding around an Ethereum contract.
type AIWorkloadCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIWorkloadTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AIWorkloadTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIWorkloadFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AIWorkloadFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIWorkloadSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AIWorkloadSession struct {
	Contract     *AIWorkload       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AIWorkloadCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AIWorkloadCallerSession struct {
	Contract *AIWorkloadCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AIWorkloadTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AIWorkloadTransactorSession struct {
	Contract     *AIWorkloadTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AIWorkloadRaw is an auto generated low-level Go binding around an Ethereum contract.
type AIWorkloadRaw struct {
	Contract *AIWorkload // Generic contract binding to access the raw methods on
}

// AIWorkloadCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AIWorkloadCallerRaw struct {
	Contract *AIWorkloadCaller // Generic read-only contract binding to access the raw methods on
}

// AIWorkloadTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AIWorkloadTransactorRaw struct {
	Contract *AIWorkloadTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAIWorkload creates a new instance of AIWorkload, bound to a specific deployed contract.
func NewAIWorkload(address common.Address, backend bind.ContractBackend) (*AIWorkload, error) {
	contract, err := bindAIWorkload(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AIWorkload{AIWorkloadCaller: AIWorkloadCaller{contract: contract}, AIWorkloadTransactor: AIWorkloadTransactor{contract: contract}, AIWorkloadFilterer: AIWorkloadFilterer{contract: contract}}, nil
}

// NewAIWorkloadCaller creates a new read-only instance of AIWorkload, bound to a specific deployed contract.
func NewAIWorkloadCaller(address common.Address, caller bind.ContractCaller) (*AIWorkloadCaller, error) {
	contract, err := bindAIWorkload(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AIWorkloadCaller{contract: contract}, nil
}

// NewAIWorkloadTransactor creates a new write-only instance of AIWorkload, bound to a specific deployed contract.
func NewAIWorkloadTransactor(address common.Address, transactor bind.ContractTransactor) (*AIWorkloadTransactor, error) {
	contract, err := bindAIWorkload(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AIWorkloadTransactor{contract: contract}, nil
}

// NewAIWorkloadFilterer creates a new log filterer instance of AIWorkload, bound to a specific deployed contract.
func NewAIWorkloadFilterer(address common.Address, filterer bind.ContractFilterer) (*AIWorkloadFilterer, error) {
	contract, err := bindAIWorkload(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AIWorkloadFilterer{contract: contract}, nil
}

// bindAIWorkload binds a generic wrapper to an already deployed contract.
func bindAIWorkload(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AIWorkloadMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AIWorkload *AIWorkloadRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AIWorkload.Contract.AIWorkloadCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AIWorkload *AIWorkloadRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIWorkload.Contract.AIWorkloadTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AIWorkload *AIWorkloadRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AIWorkload.Contract.AIWorkloadTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AIWorkload *AIWorkloadCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AIWorkload.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AIWorkload *AIWorkloadTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIWorkload.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AIWorkload *AIWorkloadTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AIWorkload.Contract.contract.Transact(opts, method, params...)
}

// GetLastEpoch is a free data retrieval call binding the contract method 0xad1f5717.
//
// Solidity: function getLastEpoch(uint256 sessionId) view returns(uint256)
func (_AIWorkload *AIWorkloadCaller) GetLastEpoch(opts *bind.CallOpts, sessionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "getLastEpoch", sessionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastEpoch is a free data retrieval call binding the contract method 0xad1f5717.
//
// Solidity: function getLastEpoch(uint256 sessionId) view returns(uint256)
func (_AIWorkload *AIWorkloadSession) GetLastEpoch(sessionId *big.Int) (*big.Int, error) {
	return _AIWorkload.Contract.GetLastEpoch(&_AIWorkload.CallOpts, sessionId)
}

// GetLastEpoch is a free data retrieval call binding the contract method 0xad1f5717.
//
// Solidity: function getLastEpoch(uint256 sessionId) view returns(uint256)
func (_AIWorkload *AIWorkloadCallerSession) GetLastEpoch(sessionId *big.Int) (*big.Int, error) {
	return _AIWorkload.Contract.GetLastEpoch(&_AIWorkload.CallOpts, sessionId)
}

// GetNodeWorkload is a free data retrieval call binding the contract method 0xd9564c36.
//
// Solidity: function getNodeWorkload(uint256 sessionId, uint256 epochId) view returns((uint256,uint256,uint256,uint256,address,address))
func (_AIWorkload *AIWorkloadCaller) GetNodeWorkload(opts *bind.CallOpts, sessionId *big.Int, epochId *big.Int) (AIWorkloadWorkload, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "getNodeWorkload", sessionId, epochId)

	if err != nil {
		return *new(AIWorkloadWorkload), err
	}

	out0 := *abi.ConvertType(out[0], new(AIWorkloadWorkload)).(*AIWorkloadWorkload)

	return out0, err

}

// GetNodeWorkload is a free data retrieval call binding the contract method 0xd9564c36.
//
// Solidity: function getNodeWorkload(uint256 sessionId, uint256 epochId) view returns((uint256,uint256,uint256,uint256,address,address))
func (_AIWorkload *AIWorkloadSession) GetNodeWorkload(sessionId *big.Int, epochId *big.Int) (AIWorkloadWorkload, error) {
	return _AIWorkload.Contract.GetNodeWorkload(&_AIWorkload.CallOpts, sessionId, epochId)
}

// GetNodeWorkload is a free data retrieval call binding the contract method 0xd9564c36.
//
// Solidity: function getNodeWorkload(uint256 sessionId, uint256 epochId) view returns((uint256,uint256,uint256,uint256,address,address))
func (_AIWorkload *AIWorkloadCallerSession) GetNodeWorkload(sessionId *big.Int, epochId *big.Int) (AIWorkloadWorkload, error) {
	return _AIWorkload.Contract.GetNodeWorkload(&_AIWorkload.CallOpts, sessionId, epochId)
}

// GetTotalModelWorkload is a free data retrieval call binding the contract method 0x20737af1.
//
// Solidity: function getTotalModelWorkload(uint256 modelId) view returns((uint256,uint256))
func (_AIWorkload *AIWorkloadCaller) GetTotalModelWorkload(opts *bind.CallOpts, modelId *big.Int) (AIWorkloadWorkLoad, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "getTotalModelWorkload", modelId)

	if err != nil {
		return *new(AIWorkloadWorkLoad), err
	}

	out0 := *abi.ConvertType(out[0], new(AIWorkloadWorkLoad)).(*AIWorkloadWorkLoad)

	return out0, err

}

// GetTotalModelWorkload is a free data retrieval call binding the contract method 0x20737af1.
//
// Solidity: function getTotalModelWorkload(uint256 modelId) view returns((uint256,uint256))
func (_AIWorkload *AIWorkloadSession) GetTotalModelWorkload(modelId *big.Int) (AIWorkloadWorkLoad, error) {
	return _AIWorkload.Contract.GetTotalModelWorkload(&_AIWorkload.CallOpts, modelId)
}

// GetTotalModelWorkload is a free data retrieval call binding the contract method 0x20737af1.
//
// Solidity: function getTotalModelWorkload(uint256 modelId) view returns((uint256,uint256))
func (_AIWorkload *AIWorkloadCallerSession) GetTotalModelWorkload(modelId *big.Int) (AIWorkloadWorkLoad, error) {
	return _AIWorkload.Contract.GetTotalModelWorkload(&_AIWorkload.CallOpts, modelId)
}

// GetTotalWorkReport is a free data retrieval call binding the contract method 0x2c2e139b.
//
// Solidity: function getTotalWorkReport(address reporter) view returns((uint256,uint256))
func (_AIWorkload *AIWorkloadCaller) GetTotalWorkReport(opts *bind.CallOpts, reporter common.Address) (AIWorkloadWorkLoad, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "getTotalWorkReport", reporter)

	if err != nil {
		return *new(AIWorkloadWorkLoad), err
	}

	out0 := *abi.ConvertType(out[0], new(AIWorkloadWorkLoad)).(*AIWorkloadWorkLoad)

	return out0, err

}

// GetTotalWorkReport is a free data retrieval call binding the contract method 0x2c2e139b.
//
// Solidity: function getTotalWorkReport(address reporter) view returns((uint256,uint256))
func (_AIWorkload *AIWorkloadSession) GetTotalWorkReport(reporter common.Address) (AIWorkloadWorkLoad, error) {
	return _AIWorkload.Contract.GetTotalWorkReport(&_AIWorkload.CallOpts, reporter)
}

// GetTotalWorkReport is a free data retrieval call binding the contract method 0x2c2e139b.
//
// Solidity: function getTotalWorkReport(address reporter) view returns((uint256,uint256))
func (_AIWorkload *AIWorkloadCallerSession) GetTotalWorkReport(reporter common.Address) (AIWorkloadWorkLoad, error) {
	return _AIWorkload.Contract.GetTotalWorkReport(&_AIWorkload.CallOpts, reporter)
}

// GetTotalWorkerWorkload is a free data retrieval call binding the contract method 0xe926805e.
//
// Solidity: function getTotalWorkerWorkload(address node) view returns((uint256,uint256))
func (_AIWorkload *AIWorkloadCaller) GetTotalWorkerWorkload(opts *bind.CallOpts, node common.Address) (AIWorkloadWorkLoad, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "getTotalWorkerWorkload", node)

	if err != nil {
		return *new(AIWorkloadWorkLoad), err
	}

	out0 := *abi.ConvertType(out[0], new(AIWorkloadWorkLoad)).(*AIWorkloadWorkLoad)

	return out0, err

}

// GetTotalWorkerWorkload is a free data retrieval call binding the contract method 0xe926805e.
//
// Solidity: function getTotalWorkerWorkload(address node) view returns((uint256,uint256))
func (_AIWorkload *AIWorkloadSession) GetTotalWorkerWorkload(node common.Address) (AIWorkloadWorkLoad, error) {
	return _AIWorkload.Contract.GetTotalWorkerWorkload(&_AIWorkload.CallOpts, node)
}

// GetTotalWorkerWorkload is a free data retrieval call binding the contract method 0xe926805e.
//
// Solidity: function getTotalWorkerWorkload(address node) view returns((uint256,uint256))
func (_AIWorkload *AIWorkloadCallerSession) GetTotalWorkerWorkload(node common.Address) (AIWorkloadWorkLoad, error) {
	return _AIWorkload.Contract.GetTotalWorkerWorkload(&_AIWorkload.CallOpts, node)
}

// LastSettlementTime is a free data retrieval call binding the contract method 0x0f17c347.
//
// Solidity: function lastSettlementTime() view returns(uint256)
func (_AIWorkload *AIWorkloadCaller) LastSettlementTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "lastSettlementTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSettlementTime is a free data retrieval call binding the contract method 0x0f17c347.
//
// Solidity: function lastSettlementTime() view returns(uint256)
func (_AIWorkload *AIWorkloadSession) LastSettlementTime() (*big.Int, error) {
	return _AIWorkload.Contract.LastSettlementTime(&_AIWorkload.CallOpts)
}

// LastSettlementTime is a free data retrieval call binding the contract method 0x0f17c347.
//
// Solidity: function lastSettlementTime() view returns(uint256)
func (_AIWorkload *AIWorkloadCallerSession) LastSettlementTime() (*big.Int, error) {
	return _AIWorkload.Contract.LastSettlementTime(&_AIWorkload.CallOpts)
}

// ModelRegistry is a free data retrieval call binding the contract method 0x2ba36f78.
//
// Solidity: function modelRegistry() view returns(address)
func (_AIWorkload *AIWorkloadCaller) ModelRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "modelRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ModelRegistry is a free data retrieval call binding the contract method 0x2ba36f78.
//
// Solidity: function modelRegistry() view returns(address)
func (_AIWorkload *AIWorkloadSession) ModelRegistry() (common.Address, error) {
	return _AIWorkload.Contract.ModelRegistry(&_AIWorkload.CallOpts)
}

// ModelRegistry is a free data retrieval call binding the contract method 0x2ba36f78.
//
// Solidity: function modelRegistry() view returns(address)
func (_AIWorkload *AIWorkloadCallerSession) ModelRegistry() (common.Address, error) {
	return _AIWorkload.Contract.ModelRegistry(&_AIWorkload.CallOpts)
}

// NodeRegistry is a free data retrieval call binding the contract method 0xd9b5c4a5.
//
// Solidity: function nodeRegistry() view returns(address)
func (_AIWorkload *AIWorkloadCaller) NodeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "nodeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeRegistry is a free data retrieval call binding the contract method 0xd9b5c4a5.
//
// Solidity: function nodeRegistry() view returns(address)
func (_AIWorkload *AIWorkloadSession) NodeRegistry() (common.Address, error) {
	return _AIWorkload.Contract.NodeRegistry(&_AIWorkload.CallOpts)
}

// NodeRegistry is a free data retrieval call binding the contract method 0xd9b5c4a5.
//
// Solidity: function nodeRegistry() view returns(address)
func (_AIWorkload *AIWorkloadCallerSession) NodeRegistry() (common.Address, error) {
	return _AIWorkload.Contract.NodeRegistry(&_AIWorkload.CallOpts)
}

// Sessions is a free data retrieval call binding the contract method 0x83c4b7a3.
//
// Solidity: function sessions(uint256 ) view returns(uint256 lastEpochId)
func (_AIWorkload *AIWorkloadCaller) Sessions(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "sessions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sessions is a free data retrieval call binding the contract method 0x83c4b7a3.
//
// Solidity: function sessions(uint256 ) view returns(uint256 lastEpochId)
func (_AIWorkload *AIWorkloadSession) Sessions(arg0 *big.Int) (*big.Int, error) {
	return _AIWorkload.Contract.Sessions(&_AIWorkload.CallOpts, arg0)
}

// Sessions is a free data retrieval call binding the contract method 0x83c4b7a3.
//
// Solidity: function sessions(uint256 ) view returns(uint256 lastEpochId)
func (_AIWorkload *AIWorkloadCallerSession) Sessions(arg0 *big.Int) (*big.Int, error) {
	return _AIWorkload.Contract.Sessions(&_AIWorkload.CallOpts, arg0)
}

// SettlementInterval is a free data retrieval call binding the contract method 0x7ec4142c.
//
// Solidity: function settlementInterval() view returns(uint256)
func (_AIWorkload *AIWorkloadCaller) SettlementInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "settlementInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SettlementInterval is a free data retrieval call binding the contract method 0x7ec4142c.
//
// Solidity: function settlementInterval() view returns(uint256)
func (_AIWorkload *AIWorkloadSession) SettlementInterval() (*big.Int, error) {
	return _AIWorkload.Contract.SettlementInterval(&_AIWorkload.CallOpts)
}

// SettlementInterval is a free data retrieval call binding the contract method 0x7ec4142c.
//
// Solidity: function settlementInterval() view returns(uint256)
func (_AIWorkload *AIWorkloadCallerSession) SettlementInterval() (*big.Int, error) {
	return _AIWorkload.Contract.SettlementInterval(&_AIWorkload.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_AIWorkload *AIWorkloadCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_AIWorkload *AIWorkloadSession) StakeToken() (common.Address, error) {
	return _AIWorkload.Contract.StakeToken(&_AIWorkload.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_AIWorkload *AIWorkloadCallerSession) StakeToken() (common.Address, error) {
	return _AIWorkload.Contract.StakeToken(&_AIWorkload.CallOpts)
}

// ReportWorkload is a paid mutator transaction binding the contract method 0x09de6d2f.
//
// Solidity: function reportWorkload(address worker, uint256 workload, uint256 modelId, uint256 sessionId, uint256 epochId, (bytes32,bytes32,uint8)[] signatures) returns()
func (_AIWorkload *AIWorkloadTransactor) ReportWorkload(opts *bind.TransactOpts, worker common.Address, workload *big.Int, modelId *big.Int, sessionId *big.Int, epochId *big.Int, signatures []Signature) (*types.Transaction, error) {
	return _AIWorkload.contract.Transact(opts, "reportWorkload", worker, workload, modelId, sessionId, epochId, signatures)
}

// ReportWorkload is a paid mutator transaction binding the contract method 0x09de6d2f.
//
// Solidity: function reportWorkload(address worker, uint256 workload, uint256 modelId, uint256 sessionId, uint256 epochId, (bytes32,bytes32,uint8)[] signatures) returns()
func (_AIWorkload *AIWorkloadSession) ReportWorkload(worker common.Address, workload *big.Int, modelId *big.Int, sessionId *big.Int, epochId *big.Int, signatures []Signature) (*types.Transaction, error) {
	return _AIWorkload.Contract.ReportWorkload(&_AIWorkload.TransactOpts, worker, workload, modelId, sessionId, epochId, signatures)
}

// ReportWorkload is a paid mutator transaction binding the contract method 0x09de6d2f.
//
// Solidity: function reportWorkload(address worker, uint256 workload, uint256 modelId, uint256 sessionId, uint256 epochId, (bytes32,bytes32,uint8)[] signatures) returns()
func (_AIWorkload *AIWorkloadTransactorSession) ReportWorkload(worker common.Address, workload *big.Int, modelId *big.Int, sessionId *big.Int, epochId *big.Int, signatures []Signature) (*types.Transaction, error) {
	return _AIWorkload.Contract.ReportWorkload(&_AIWorkload.TransactOpts, worker, workload, modelId, sessionId, epochId, signatures)
}

// SettleRewards is a paid mutator transaction binding the contract method 0xd14c9fa7.
//
// Solidity: function settleRewards() returns((address,uint256)[] settledWorkers, (uint256,uint256)[] settledModels, (address,uint256)[] settledReporters)
func (_AIWorkload *AIWorkloadTransactor) SettleRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIWorkload.contract.Transact(opts, "settleRewards")
}

// SettleRewards is a paid mutator transaction binding the contract method 0xd14c9fa7.
//
// Solidity: function settleRewards() returns((address,uint256)[] settledWorkers, (uint256,uint256)[] settledModels, (address,uint256)[] settledReporters)
func (_AIWorkload *AIWorkloadSession) SettleRewards() (*types.Transaction, error) {
	return _AIWorkload.Contract.SettleRewards(&_AIWorkload.TransactOpts)
}

// SettleRewards is a paid mutator transaction binding the contract method 0xd14c9fa7.
//
// Solidity: function settleRewards() returns((address,uint256)[] settledWorkers, (uint256,uint256)[] settledModels, (address,uint256)[] settledReporters)
func (_AIWorkload *AIWorkloadTransactorSession) SettleRewards() (*types.Transaction, error) {
	return _AIWorkload.Contract.SettleRewards(&_AIWorkload.TransactOpts)
}

// AIWorkloadWorkloadReportedIterator is returned from FilterWorkloadReported and is used to iterate over the raw logs and unpacked data for WorkloadReported events raised by the AIWorkload contract.
type AIWorkloadWorkloadReportedIterator struct {
	Event *AIWorkloadWorkloadReported // Event containing the contract specifics and raw log

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
func (it *AIWorkloadWorkloadReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIWorkloadWorkloadReported)
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
		it.Event = new(AIWorkloadWorkloadReported)
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
func (it *AIWorkloadWorkloadReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIWorkloadWorkloadReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIWorkloadWorkloadReported represents a WorkloadReported event raised by the AIWorkload contract.
type AIWorkloadWorkloadReported struct {
	SessionId *big.Int
	Reporter  common.Address
	Worker    common.Address
	EpochId   *big.Int
	Workload  *big.Int
	ModelId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWorkloadReported is a free log retrieval operation binding the contract event 0x59164c09eed89694722e80b610592178b04fbe4a72b982bf4503c8ab426f0ef6.
//
// Solidity: event WorkloadReported(uint256 indexed sessionId, address indexed reporter, address worker, uint256 epochId, uint256 workload, uint256 modelId)
func (_AIWorkload *AIWorkloadFilterer) FilterWorkloadReported(opts *bind.FilterOpts, sessionId []*big.Int, reporter []common.Address) (*AIWorkloadWorkloadReportedIterator, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _AIWorkload.contract.FilterLogs(opts, "WorkloadReported", sessionIdRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return &AIWorkloadWorkloadReportedIterator{contract: _AIWorkload.contract, event: "WorkloadReported", logs: logs, sub: sub}, nil
}

// WatchWorkloadReported is a free log subscription operation binding the contract event 0x59164c09eed89694722e80b610592178b04fbe4a72b982bf4503c8ab426f0ef6.
//
// Solidity: event WorkloadReported(uint256 indexed sessionId, address indexed reporter, address worker, uint256 epochId, uint256 workload, uint256 modelId)
func (_AIWorkload *AIWorkloadFilterer) WatchWorkloadReported(opts *bind.WatchOpts, sink chan<- *AIWorkloadWorkloadReported, sessionId []*big.Int, reporter []common.Address) (event.Subscription, error) {

	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _AIWorkload.contract.WatchLogs(opts, "WorkloadReported", sessionIdRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIWorkloadWorkloadReported)
				if err := _AIWorkload.contract.UnpackLog(event, "WorkloadReported", log); err != nil {
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

// ParseWorkloadReported is a log parse operation binding the contract event 0x59164c09eed89694722e80b610592178b04fbe4a72b982bf4503c8ab426f0ef6.
//
// Solidity: event WorkloadReported(uint256 indexed sessionId, address indexed reporter, address worker, uint256 epochId, uint256 workload, uint256 modelId)
func (_AIWorkload *AIWorkloadFilterer) ParseWorkloadReported(log types.Log) (*AIWorkloadWorkloadReported, error) {
	event := new(AIWorkloadWorkloadReported)
	if err := _AIWorkload.contract.UnpackLog(event, "WorkloadReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
