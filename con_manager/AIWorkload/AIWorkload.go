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

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	R [32]byte
	S [32]byte
	V uint8
}

// AIWorkloadMetaData contains all meta data concerning the AIWorkload contract.
var AIWorkloadMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"WorkloadReported\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"}],\"name\":\"getLastEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"getNodeWorkload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"internalType\":\"structAIWorkload.Workload\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"getTotalModelWorkload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalWorkload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledWorkload\",\"type\":\"uint256\"}],\"internalType\":\"structAIWorkload.WorkLoad\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"getTotalWorkReport\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalWorkload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledWorkload\",\"type\":\"uint256\"}],\"internalType\":\"structAIWorkload.WorkLoad\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getTotalWorkerWorkload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalWorkload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledWorkload\",\"type\":\"uint256\"}],\"internalType\":\"structAIWorkload.WorkLoad\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSettlementTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractNodesRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"reportWorkload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sessions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastEpochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settleRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052610e10600b5534801561001657600080fd5b50604051611578380380611578833981016040819052610035916100ab565b6001600160a01b0381166100825760405162461bcd60e51b815260206004820152601060248201526f496e76616c696420726567697374727960801b604482015260640160405180910390fd5b600180546001600160a01b0319166001600160a01b039290921691909117905542600c556100d9565b6000602082840312156100bc578081fd5b81516001600160a01b03811681146100d2578182fd5b9392505050565b611490806100e86000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80637ec4142c116100715780637ec4142c1461018d57806383c4b7a314610196578063ad1f5717146101b6578063d14c9fa7146101d6578063d9564c36146101de578063e926805e14610247576100a9565b806309de6d2f146100ae5780630f17c347146100c357806320737af1146100df5780632c2e139b1461010d5780637b10399914610162575b600080fd5b6100c16100bc3660046110f6565b61029c565b005b6100cc600c5481565b6040519081526020015b60405180910390f35b6100f26100ed366004611278565b610685565b604080518251815260209283015192810192909252016100d6565b6100f261011b3660046110da565b6040805180820190915260008082526020820152506001600160a01b0316600090815260086020908152604091829020825180840190935280548352600101549082015290565b600154610175906001600160a01b031681565b6040516001600160a01b0390911681526020016100d6565b6100cc600b5481565b6100cc6101a4366004611278565b60006020819052908152604090205481565b6100cc6101c4366004611278565b60009081526020819052604090205490565b6100c16106bf565b6101f16101ec366004611290565b6107e9565b6040516100d69190815181526020808301519082015260408083015190820152606080830151908201526080808301516001600160a01b039081169183019190915260a092830151169181019190915260c00190565b6100f26102553660046110da565b6040805180820190915260008082526020820152506001600160a01b0316600090815260026020908152604091829020825180840190935280548352600101549082015290565b6001600160a01b0387166102ef5760405162461bcd60e51b8152602060048201526015602482015274496e76616c6964206f776e6572206164647265737360581b60448201526064015b60405180910390fd5b6000861161034a5760405162461bcd60e51b815260206004820152602260248201527f576f726b6c6f6164206d7573742062652067726561746572207468616e207a65604482015261726f60f01b60648201526084016102e6565b60038110156103a95760405162461bcd60e51b815260206004820152602560248201527f4c656e677468206f66207369676e617475726573206d757374206d6f7265207460448201526468616e203360d81b60648201526084016102e6565b604080516001600160a01b0389166020820152908101879052606081018690526080810185905260a0810184905260009060c00160405160208183030381529060405290506103fb88338386866108a4565b61043b5760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b60448201526064016102e6565b60008581526020819052604090208054851161048e5760405162461bcd60e51b815260206004820152601260248201527122b837b1b41037baba1037b31037b93232b960711b60448201526064016102e6565b6040518060c00160405280868152602001898152602001428152602001888152602001336001600160a01b031681526020018a6001600160a01b03168152508160010160008781526020019081526020016000206000820151816000015560208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060a08201518160050160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050506000600260008b6001600160a01b03166001600160a01b031681526020019081526020016000209050888160000160008282546105a4919061135e565b909155506105b5905060038b610be0565b506000888152600560205260408120805490918b918391906105d890849061135e565b909155506105e9905060068a610bfc565b50336000908152600860205260408120805490918c9183919061060d90849061135e565b9091555061061e9050600933610be0565b50878455604080516001600160a01b038e168152602081018a90529081018c9052606081018b905233908a907f59164c09eed89694722e80b610592178b04fbe4a72b982bf4503c8ab426f0ef69060800160405180910390a3505050505050505050505050565b604080518082018252600080825260209182018190528381526005825282902082518084019093528054835260010154908201525b919050565b600b54600c546106cf919061135e565b4210156107175760405162461bcd60e51b815260206004820152601660248201527514d95d1d1b195b595b9d081b9bdd08191d59481e595d60521b60448201526064016102e6565b60005b6107246003610c08565b81101561075d57600060028161073b600385610c12565b6001600160a01b03168152602001525080610755816113d1565b91505061071a565b5060005b61076b6006610c08565b81101561079b576000600581610782600685610c12565b8152602001525080610793816113d1565b915050610761565b5060005b6107a96009610c08565b8110156107e25760006008816107c0600985610c12565b6001600160a01b031681526020015250806107da816113d1565b91505061079f565b5042600c55565b6108346040518060c001604052806000815260200160008152602001600081526020016000815260200160006001600160a01b0316815260200160006001600160a01b031681525090565b506000828152602081815260408083208484526001908101835292819020815160c0810183528154815293810154928401929092526002820154908301526003810154606083015260048101546001600160a01b0390811660808401526005909101541660a08201525b92915050565b6000808267ffffffffffffffff8111156108ce57634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156108f7578160200160208202803683370190505b50905060008080805b86811015610b905760008089898481811061092b57634e487b7160e01b600052603260045260246000fd5b905060600201600001358a8a8581811061095557634e487b7160e01b600052603260045260246000fd5b905060600201602001358b8b8681811061097f57634e487b7160e01b600052603260045260246000fd5b905060600201604001602081019061099791906112b1565b6040516020016109c793929190928352602083019190915260f81b6001600160f81b031916604082015260410190565b604051602081830303815290604052905060006109ec6109e68d610c1e565b83610c59565b506001546040516330af0bbf60e21b81526001600160a01b03808416600483015292935091169063c2bc2efc9060240160006040518083038186803b158015610a3457600080fd5b505afa158015610a48573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a7091908101906111a0565b60600151610a8057505050610b7e565b60005b87811015610ae757816001600160a01b0316898281518110610ab557634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03161415610ad55760019350610ae7565b80610adf816113d1565b915050610a83565b508215610af657505050610b7e565b8d6001600160a01b0316816001600160a01b03161415610b1557600194505b8c6001600160a01b0316816001600160a01b03161415610b3457600195505b80888881518110610b5557634e487b7160e01b600052603260045260246000fd5b6001600160a01b0390921660209283029190910190910152610b7860018861135e565b96505050505b80610b88816113d1565b915050610900565b506002610b9e87600161135e565b610ba89190611376565b831080610bb3575080155b80610bbc575081155b15610bce576000945050505050610bd7565b60019450505050505b95945050505050565b6000610bf5836001600160a01b038416610cc9565b9392505050565b6000610bf58383610cc9565b600061089e825490565b6000610bf58383610d18565b6000610c2a8251610d50565b82604051602001610c3c9291906112d2565b604051602081830303815290604052805190602001209050919050565b600080825160411415610c905760208301516040840151606085015160001a610c8487828585610e73565b94509450505050610cc2565b825160401415610cba5760208301516040840151610caf868383610f60565b935093505050610cc2565b506000905060025b9250929050565b6000818152600183016020526040812054610d105750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561089e565b50600061089e565b6000826000018281548110610d3d57634e487b7160e01b600052603260045260246000fd5b9060005260206000200154905092915050565b606081610d7557506040805180820190915260018152600360fc1b60208201526106ba565b8160005b8115610d9f5780610d89816113d1565b9150610d989050600a83611376565b9150610d79565b60008167ffffffffffffffff811115610dc857634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015610df2576020820181803683370190505b5090505b8415610e6b57610e0760018361138a565b9150610e14600a866113ec565b610e1f90603061135e565b60f81b818381518110610e4257634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a905350610e64600a86611376565b9450610df6565b949350505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610eaa5750600090506003610f57565b8460ff16601b14158015610ec257508460ff16601c14155b15610ed35750600090506004610f57565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610f27573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610f5057600060019250925050610f57565b9150600090505b94509492505050565b6000806001600160ff1b03831681610f7d60ff86901c601b61135e565b9050610f8b87828885610e73565b935093505050935093915050565b80516106ba81611442565b600082601f830112610fb4578081fd5b8151602067ffffffffffffffff80831115610fd157610fd161142c565b610fdf828460051b0161132d565b83815282810190868401865b8681101561105a5781518901606080601f19838e0301121561100b57898afd5b6110148161132d565b8883015188811115611024578b8cfd5b6110328e8b83870101611078565b8252506040838101518a83015291909201519082015284529285019290850190600101610feb565b509098975050505050505050565b805180151581146106ba57600080fd5b600082601f830112611088578081fd5b815167ffffffffffffffff8111156110a2576110a261142c565b6110b5601f8201601f191660200161132d565b8181528460208386010111156110c9578283fd5b610e6b8260208301602087016113a1565b6000602082840312156110eb578081fd5b8135610bf581611442565b600080600080600080600060c0888a031215611110578283fd5b873561111b81611442565b96506020880135955060408801359450606088013593506080880135925060a088013567ffffffffffffffff80821115611153578384fd5b818a0191508a601f830112611166578384fd5b813581811115611174578485fd5b8b6020606083028501011115611188578485fd5b60208301945080935050505092959891949750929550565b6000602082840312156111b1578081fd5b815167ffffffffffffffff808211156111c8578283fd5b9083019060e082860312156111db578283fd5b6111e560e061132d565b6111ee83610f99565b8152602083015182811115611201578485fd5b61120d87828601611078565b6020830152506040830151604082015261122960608401611068565b606082015260808301518281111561123f578485fd5b61124b87828601610fa4565b60808301525061125d60a08401610f99565b60a082015260c083015160c082015280935050505092915050565b600060208284031215611289578081fd5b5035919050565b600080604083850312156112a2578182fd5b50508035926020909101359150565b6000602082840312156112c2578081fd5b813560ff81168114610bf5578182fd5b60007f19457468657265756d205369676e6564204d6573736167653a0a0000000000008252835161130a81601a8501602088016113a1565b83519083019061132181601a8401602088016113a1565b01601a01949350505050565b604051601f8201601f1916810167ffffffffffffffff811182821017156113565761135661142c565b604052919050565b6000821982111561137157611371611400565b500190565b60008261138557611385611416565b500490565b60008282101561139c5761139c611400565b500390565b60005b838110156113bc5781810151838201526020016113a4565b838111156113cb576000848401525b50505050565b60006000198214156113e5576113e5611400565b5060010190565b6000826113fb576113fb611416565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461145757600080fd5b5056fea264697066735822122089dde7daf09317f97ee2a2351195b504c62552a0b6bc485a60f3ca97c5f5790664736f6c63430008030033",
}

// AIWorkloadABI is the input ABI used to generate the binding from.
// Deprecated: Use AIWorkloadMetaData.ABI instead.
var AIWorkloadABI = AIWorkloadMetaData.ABI

// AIWorkloadBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AIWorkloadMetaData.Bin instead.
var AIWorkloadBin = AIWorkloadMetaData.Bin

// DeployAIWorkload deploys a new Ethereum contract, binding an instance of AIWorkload to it.
func DeployAIWorkload(auth *bind.TransactOpts, backend bind.ContractBackend, _registry common.Address) (common.Address, *types.Transaction, *AIWorkload, error) {
	parsed, err := AIWorkloadMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AIWorkloadBin), backend, _registry)
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

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AIWorkload *AIWorkloadCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AIWorkload *AIWorkloadSession) Registry() (common.Address, error) {
	return _AIWorkload.Contract.Registry(&_AIWorkload.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AIWorkload *AIWorkloadCallerSession) Registry() (common.Address, error) {
	return _AIWorkload.Contract.Registry(&_AIWorkload.CallOpts)
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
// Solidity: function settleRewards() returns()
func (_AIWorkload *AIWorkloadTransactor) SettleRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIWorkload.contract.Transact(opts, "settleRewards")
}

// SettleRewards is a paid mutator transaction binding the contract method 0xd14c9fa7.
//
// Solidity: function settleRewards() returns()
func (_AIWorkload *AIWorkloadSession) SettleRewards() (*types.Transaction, error) {
	return _AIWorkload.Contract.SettleRewards(&_AIWorkload.TransactOpts)
}

// SettleRewards is a paid mutator transaction binding the contract method 0xd14c9fa7.
//
// Solidity: function settleRewards() returns()
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
