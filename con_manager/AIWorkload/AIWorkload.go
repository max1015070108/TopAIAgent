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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"WorkloadReported\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"}],\"name\":\"getLastEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"getNodeWorkload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"internalType\":\"structAIWorkload.Workload\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getTotalWorkReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getTotalWorkload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractNodesRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"reportWorkload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sessions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastEpochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalWorkReports\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalWorkload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161127a38038061127a83398101604081905261002f916100a1565b6001600160a01b03811661007c5760405162461bcd60e51b815260206004820152601060248201526f496e76616c696420726567697374727960801b604482015260640160405180910390fd5b600180546001600160a01b0319166001600160a01b03929092169190911790556100cf565b6000602082840312156100b2578081fd5b81516001600160a01b03811681146100c8578182fd5b9392505050565b61119c806100de6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80637b103999116100665780637b1039991461011357806383c4b7a31461013e578063980ab13a1461015e578063ad1f571714610187578063d9564c36146101a757610093565b806309de6d2f146100985780632c2e139b146100ad578063575eb436146100d35780635cc319f1146100f3575b600080fd5b6100ab6100a6366004610e0f565b610210565b005b6100c06100bb366004610dec565b6105a2565b6040519081526020015b60405180910390f35b6100c06100e1366004610dec565b60036020526000908152604090205481565b6100c0610101366004610dec565b60026020526000908152604090205481565b600154610126906001600160a01b031681565b6040516001600160a01b0390911681526020016100ca565b6100c061014c366004610f84565b60006020819052908152604090205481565b6100c061016c366004610dec565b6001600160a01b031660009081526002602052604090205490565b6100c0610195366004610f84565b60009081526020819052604090205490565b6101ba6101b5366004610f9c565b6105c1565b6040516100ca9190815181526020808301519082015260408083015190820152606080830151908201526080808301516001600160a01b039081169183019190915260a092830151169181019190915260c00190565b6001600160a01b0387166102635760405162461bcd60e51b8152602060048201526015602482015274496e76616c6964206f776e6572206164647265737360581b60448201526064015b60405180910390fd5b600086116102be5760405162461bcd60e51b815260206004820152602260248201527f576f726b6c6f6164206d7573742062652067726561746572207468616e207a65604482015261726f60f01b606482015260840161025a565b600381101561031d5760405162461bcd60e51b815260206004820152602560248201527f4c656e677468206f66207369676e617475726573206d757374206d6f7265207460448201526468616e203360d81b606482015260840161025a565b604080516001600160a01b0389166020820152908101879052606081018690526080810185905260a0810184905260009060c001604051602081830303815290604052905061036f883383868661067b565b6103af5760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b604482015260640161025a565b6000858152602081905260409020805485116104025760405162461bcd60e51b815260206004820152601260248201527122b837b1b41037baba1037b31037b93232b960711b604482015260640161025a565b6040518060c00160405280868152602001898152602001428152602001888152602001336001600160a01b031681526020018a6001600160a01b03168152508160010160008781526020019081526020016000206000820151816000015560208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060a08201518160050160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555090505084816000018190555087600260008b6001600160a01b03166001600160a01b031681526020019081526020016000206000828254610519919061106a565b909155505033600090815260036020526040812080548a929061053d90849061106a565b9091555050604080516001600160a01b038b1681526020810187905290810189905260608101889052339087907f59164c09eed89694722e80b610592178b04fbe4a72b982bf4503c8ab426f0ef69060800160405180910390a3505050505050505050565b6001600160a01b0381166000908152600360205260409020545b919050565b61060c6040518060c001604052806000815260200160008152602001600081526020016000815260200160006001600160a01b0316815260200160006001600160a01b031681525090565b506000828152602081815260408083208484526001908101835292819020815160c0810183528154815293810154928401929092526002820154908301526003810154606083015260048101546001600160a01b0390811660808401526005909101541660a082015292915050565b6000808267ffffffffffffffff8111156106a557634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156106ce578160200160208202803683370190505b50905060008080805b868110156109675760008089898481811061070257634e487b7160e01b600052603260045260246000fd5b905060600201600001358a8a8581811061072c57634e487b7160e01b600052603260045260246000fd5b905060600201602001358b8b8681811061075657634e487b7160e01b600052603260045260246000fd5b905060600201604001602081019061076e9190610fbd565b60405160200161079e93929190928352602083019190915260f81b6001600160f81b031916604082015260410190565b604051602081830303815290604052905060006107c36107bd8d6109b7565b836109f2565b506001546040516330af0bbf60e21b81526001600160a01b03808416600483015292935091169063c2bc2efc9060240160006040518083038186803b15801561080b57600080fd5b505afa15801561081f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108479190810190610eb9565b6060015161085757505050610955565b60005b878110156108be57816001600160a01b031689828151811061088c57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b031614156108ac57600193506108be565b806108b6816110dd565b91505061085a565b5082156108cd57505050610955565b8d6001600160a01b0316816001600160a01b031614156108ec57600194505b8c6001600160a01b0316816001600160a01b0316141561090b57600195505b8088888151811061092c57634e487b7160e01b600052603260045260246000fd5b6001600160a01b039092166020928302919091019091015261094f60018861106a565b96505050505b8061095f816110dd565b9150506106d7565b50600261097587600161106a565b61097f9190611082565b83108061098a575080155b80610993575081155b156109a55760009450505050506109ae565b60019450505050505b95945050505050565b60006109c38251610a62565b826040516020016109d5929190610fde565b604051602081830303815290604052805190602001209050919050565b600080825160411415610a295760208301516040840151606085015160001a610a1d87828585610b85565b94509450505050610a5b565b825160401415610a535760208301516040840151610a48868383610c72565b935093505050610a5b565b506000905060025b9250929050565b606081610a8757506040805180820190915260018152600360fc1b60208201526105bc565b8160005b8115610ab15780610a9b816110dd565b9150610aaa9050600a83611082565b9150610a8b565b60008167ffffffffffffffff811115610ada57634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015610b04576020820181803683370190505b5090505b8415610b7d57610b19600183611096565b9150610b26600a866110f8565b610b3190603061106a565b60f81b818381518110610b5457634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a905350610b76600a86611082565b9450610b08565b949350505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610bbc5750600090506003610c69565b8460ff16601b14158015610bd457508460ff16601c14155b15610be55750600090506004610c69565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610c39573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610c6257600060019250925050610c69565b9150600090505b94509492505050565b6000806001600160ff1b03831681610c8f60ff86901c601b61106a565b9050610c9d87828885610b85565b935093505050935093915050565b80516105bc8161114e565b600082601f830112610cc6578081fd5b8151602067ffffffffffffffff80831115610ce357610ce3611138565b610cf1828460051b01611039565b83815282810190868401865b86811015610d6c5781518901606080601f19838e03011215610d1d57898afd5b610d2681611039565b8883015188811115610d36578b8cfd5b610d448e8b83870101610d8a565b8252506040838101518a83015291909201519082015284529285019290850190600101610cfd565b509098975050505050505050565b805180151581146105bc57600080fd5b600082601f830112610d9a578081fd5b815167ffffffffffffffff811115610db457610db4611138565b610dc7601f8201601f1916602001611039565b818152846020838601011115610ddb578283fd5b610b7d8260208301602087016110ad565b600060208284031215610dfd578081fd5b8135610e088161114e565b9392505050565b600080600080600080600060c0888a031215610e29578283fd5b8735610e348161114e565b96506020880135955060408801359450606088013593506080880135925060a088013567ffffffffffffffff80821115610e6c578384fd5b818a0191508a601f830112610e7f578384fd5b813581811115610e8d578485fd5b8b6020606083028501011115610ea1578485fd5b60208301945080935050505092959891949750929550565b600060208284031215610eca578081fd5b815167ffffffffffffffff80821115610ee1578283fd5b9083019060c08286031215610ef4578283fd5b610efe60c0611039565b610f0783610cab565b8152602083015182811115610f1a578485fd5b610f2687828601610d8a565b60208301525060408301516040820152610f4260608401610d7a565b6060820152608083015182811115610f58578485fd5b610f6487828601610cb6565b608083015250610f7660a08401610cab565b60a082015295945050505050565b600060208284031215610f95578081fd5b5035919050565b60008060408385031215610fae578182fd5b50508035926020909101359150565b600060208284031215610fce578081fd5b813560ff81168114610e08578182fd5b60007f19457468657265756d205369676e6564204d6573736167653a0a0000000000008252835161101681601a8501602088016110ad565b83519083019061102d81601a8401602088016110ad565b01601a01949350505050565b604051601f8201601f1916810167ffffffffffffffff8111828210171561106257611062611138565b604052919050565b6000821982111561107d5761107d61110c565b500190565b60008261109157611091611122565b500490565b6000828210156110a8576110a861110c565b500390565b60005b838110156110c85781810151838201526020016110b0565b838111156110d7576000848401525b50505050565b60006000198214156110f1576110f161110c565b5060010190565b60008261110757611107611122565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461116357600080fd5b5056fea264697066735822122066d6f0b91cbcfc4b513b9eabac65d0a4b82901951f2d3777a3cd88b7f7e1e74364736f6c63430008030033",
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

// GetTotalWorkReport is a free data retrieval call binding the contract method 0x2c2e139b.
//
// Solidity: function getTotalWorkReport(address node) view returns(uint256)
func (_AIWorkload *AIWorkloadCaller) GetTotalWorkReport(opts *bind.CallOpts, node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "getTotalWorkReport", node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalWorkReport is a free data retrieval call binding the contract method 0x2c2e139b.
//
// Solidity: function getTotalWorkReport(address node) view returns(uint256)
func (_AIWorkload *AIWorkloadSession) GetTotalWorkReport(node common.Address) (*big.Int, error) {
	return _AIWorkload.Contract.GetTotalWorkReport(&_AIWorkload.CallOpts, node)
}

// GetTotalWorkReport is a free data retrieval call binding the contract method 0x2c2e139b.
//
// Solidity: function getTotalWorkReport(address node) view returns(uint256)
func (_AIWorkload *AIWorkloadCallerSession) GetTotalWorkReport(node common.Address) (*big.Int, error) {
	return _AIWorkload.Contract.GetTotalWorkReport(&_AIWorkload.CallOpts, node)
}

// GetTotalWorkload is a free data retrieval call binding the contract method 0x980ab13a.
//
// Solidity: function getTotalWorkload(address node) view returns(uint256)
func (_AIWorkload *AIWorkloadCaller) GetTotalWorkload(opts *bind.CallOpts, node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "getTotalWorkload", node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalWorkload is a free data retrieval call binding the contract method 0x980ab13a.
//
// Solidity: function getTotalWorkload(address node) view returns(uint256)
func (_AIWorkload *AIWorkloadSession) GetTotalWorkload(node common.Address) (*big.Int, error) {
	return _AIWorkload.Contract.GetTotalWorkload(&_AIWorkload.CallOpts, node)
}

// GetTotalWorkload is a free data retrieval call binding the contract method 0x980ab13a.
//
// Solidity: function getTotalWorkload(address node) view returns(uint256)
func (_AIWorkload *AIWorkloadCallerSession) GetTotalWorkload(node common.Address) (*big.Int, error) {
	return _AIWorkload.Contract.GetTotalWorkload(&_AIWorkload.CallOpts, node)
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

// TotalWorkReports is a free data retrieval call binding the contract method 0x575eb436.
//
// Solidity: function totalWorkReports(address ) view returns(uint256)
func (_AIWorkload *AIWorkloadCaller) TotalWorkReports(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "totalWorkReports", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWorkReports is a free data retrieval call binding the contract method 0x575eb436.
//
// Solidity: function totalWorkReports(address ) view returns(uint256)
func (_AIWorkload *AIWorkloadSession) TotalWorkReports(arg0 common.Address) (*big.Int, error) {
	return _AIWorkload.Contract.TotalWorkReports(&_AIWorkload.CallOpts, arg0)
}

// TotalWorkReports is a free data retrieval call binding the contract method 0x575eb436.
//
// Solidity: function totalWorkReports(address ) view returns(uint256)
func (_AIWorkload *AIWorkloadCallerSession) TotalWorkReports(arg0 common.Address) (*big.Int, error) {
	return _AIWorkload.Contract.TotalWorkReports(&_AIWorkload.CallOpts, arg0)
}

// TotalWorkload is a free data retrieval call binding the contract method 0x5cc319f1.
//
// Solidity: function totalWorkload(address ) view returns(uint256)
func (_AIWorkload *AIWorkloadCaller) TotalWorkload(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "totalWorkload", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWorkload is a free data retrieval call binding the contract method 0x5cc319f1.
//
// Solidity: function totalWorkload(address ) view returns(uint256)
func (_AIWorkload *AIWorkloadSession) TotalWorkload(arg0 common.Address) (*big.Int, error) {
	return _AIWorkload.Contract.TotalWorkload(&_AIWorkload.CallOpts, arg0)
}

// TotalWorkload is a free data retrieval call binding the contract method 0x5cc319f1.
//
// Solidity: function totalWorkload(address ) view returns(uint256)
func (_AIWorkload *AIWorkloadCallerSession) TotalWorkload(arg0 common.Address) (*big.Int, error) {
	return _AIWorkload.Contract.TotalWorkload(&_AIWorkload.CallOpts, arg0)
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
