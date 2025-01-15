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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"WorkloadReported\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"}],\"name\":\"getLastEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"getNodeWorkload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"}],\"internalType\":\"structAIWorkload.Workload\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"getTotalModelWorkload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalWorkload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledWorkload\",\"type\":\"uint256\"}],\"internalType\":\"structAIWorkload.WorkLoad\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"getTotalWorkReport\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalWorkload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledWorkload\",\"type\":\"uint256\"}],\"internalType\":\"structAIWorkload.WorkLoad\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getTotalWorkerWorkload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalWorkload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledWorkload\",\"type\":\"uint256\"}],\"internalType\":\"structAIWorkload.WorkLoad\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSettlementTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractNodesRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"reportWorkload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sessions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastEpochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settleRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"}],\"internalType\":\"structNodeSettleWorkload[]\",\"name\":\"settledWorkers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"}],\"internalType\":\"structModelSettleWorkload[]\",\"name\":\"settledModels\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"}],\"internalType\":\"structNodeSettleWorkload[]\",\"name\":\"settledReporters\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052610e10600b5534801561001657600080fd5b50604051611c8a380380611c8a833981016040819052610035916100ab565b6001600160a01b0381166100825760405162461bcd60e51b815260206004820152601060248201526f496e76616c696420726567697374727960801b604482015260640160405180910390fd5b600180546001600160a01b0319166001600160a01b039290921691909117905542600c556100d9565b6000602082840312156100bc578081fd5b81516001600160a01b03811681146100d2578182fd5b9392505050565b611ba2806100e86000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80637ec4142c116100715780637ec4142c1461017f57806383c4b7a314610188578063ad1f5717146101a8578063d14c9fa7146101c8578063d9564c36146101df578063e926805e14610248576100a9565b806309de6d2f146100ae5780630f17c347146100c357806320737af1146100df5780632c2e139b146100ff5780637b10399914610154575b600080fd5b6100c16100bc366004611729565b61029d565b005b6100cc600c5481565b6040519081526020015b60405180910390f35b6100f26100ed3660046118a9565b610686565b6040516100d69190611a29565b6100f261010d36600461170d565b6040805180820190915260008082526020820152506001600160a01b0316600090815260086020908152604091829020825180840190935280548352600101549082015290565b600154610167906001600160a01b031681565b6040516001600160a01b0390911681526020016100d6565b6100cc600b5481565b6100cc6101963660046118a9565b60006020819052908152604090205481565b6100cc6101b63660046118a9565b60009081526020819052604090205490565b6101d06106c0565b6040516100d6939291906119ac565b6101f26101ed3660046118c1565b610e20565b6040516100d69190815181526020808301519082015260408083015190820152606080830151908201526080808301516001600160a01b039081169183019190915260a092830151169181019190915260c00190565b6100f261025636600461170d565b6040805180820190915260008082526020820152506001600160a01b0316600090815260026020908152604091829020825180840190935280548352600101549082015290565b6001600160a01b0387166102f05760405162461bcd60e51b8152602060048201526015602482015274496e76616c6964206f776e6572206164647265737360581b60448201526064015b60405180910390fd5b6000861161034b5760405162461bcd60e51b815260206004820152602260248201527f576f726b6c6f6164206d7573742062652067726561746572207468616e207a65604482015261726f60f01b60648201526084016102e7565b60038110156103aa5760405162461bcd60e51b815260206004820152602560248201527f4c656e677468206f66207369676e617475726573206d757374206d6f7265207460448201526468616e203360d81b60648201526084016102e7565b604080516001600160a01b0389166020820152908101879052606081018690526080810185905260a0810184905260009060c00160405160208183030381529060405290506103fc8833838686610edb565b61043c5760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b60448201526064016102e7565b60008581526020819052604090208054851161048f5760405162461bcd60e51b815260206004820152601260248201527122b837b1b41037baba1037b31037b93232b960711b60448201526064016102e7565b6040518060c00160405280868152602001898152602001428152602001888152602001336001600160a01b031681526020018a6001600160a01b03168152508160010160008781526020019081526020016000206000820151816000015560208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060a08201518160050160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050506000600260008b6001600160a01b03166001600160a01b031681526020019081526020016000209050888160000160008282546105a59190611a70565b909155506105b6905060038b611216565b506000888152600560205260408120805490918b918391906105d9908490611a70565b909155506105ea905060068a611232565b50336000908152600860205260408120805490918c9183919061060e908490611a70565b9091555061061f9050600933611216565b50878455604080516001600160a01b038e168152602081018a90529081018c9052606081018b905233908a907f59164c09eed89694722e80b610592178b04fbe4a72b982bf4503c8ab426f0ef69060800160405180910390a3505050505050505050505050565b604080518082018252600080825260209182018190528381526005825282902082518084019093528054835260010154908201525b919050565b6060806060600b54600c546106d59190611a70565b42101561071d5760405162461bcd60e51b815260206004820152601660248201527514d95d1d1b195b595b9d081b9bdd08191d59481e595d60521b60448201526064016102e7565b6000610729600361123e565b6001600160401b0381111561074e57634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561079357816020015b604080518082019091526000808252602082015281526020019060019003908161076c5790505b5090506000805b6107a4600361123e565b8110156108955760006002816107bb600385611248565b6001600160a01b03166001600160a01b03168152602001908152602001600020905080600101548160000154116107f25750610883565b6107fd600383611248565b84848151811061081d57634e487b7160e01b600052603260045260246000fd5b60209081029190910101516001600160a01b039091169052600181015481546108469190611a9c565b84848151811061086657634e487b7160e01b600052603260045260246000fd5b60209081029190910181015101528261087e81611ae3565b935050505b8061088d81611ae3565b91505061079a565b50806001600160401b038111156108bc57634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561090157816020015b60408051808201909152600080825260208201528152602001906001900390816108da5790505b50945060005b818110156109755782818151811061092f57634e487b7160e01b600052603260045260246000fd5b602002602001015186828151811061095757634e487b7160e01b600052603260045260246000fd5b6020026020010181905250808061096d90611ae3565b915050610907565b506000610982600661123e565b6001600160401b038111156109a757634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156109ec57816020015b60408051808201909152600080825260208201528152602001906001900390816109c55790505b5090506000915060005b610a00600661123e565b811015610ad3576000600581610a17600685611248565b815260200190815260200160002090508060010154816000015411610a3c5750610ac1565b610a47600683611248565b838581518110610a6757634e487b7160e01b600052603260045260246000fd5b60209081029190910101515260018101548154610a849190611a9c565b838581518110610aa457634e487b7160e01b600052603260045260246000fd5b602090810291909101810151015283610abc81611ae3565b945050505b80610acb81611ae3565b9150506109f6565b50816001600160401b03811115610afa57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610b3f57816020015b6040805180820190915260008082526020820152815260200190600190039081610b185790505b50945060005b82811015610bb357818181518110610b6d57634e487b7160e01b600052603260045260246000fd5b6020026020010151868281518110610b9557634e487b7160e01b600052603260045260246000fd5b60200260200101819052508080610bab90611ae3565b915050610b45565b506000610bc0600961123e565b6001600160401b03811115610be557634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610c2a57816020015b6040805180820190915260008082526020820152815260200190600190039081610c035790505b5090506000925060005b610c3e600961123e565b811015610d2f576000600881610c55600985611248565b6001600160a01b03166001600160a01b0316815260200190815260200160002090508060010154816000015411610c8c5750610d1d565b610c97600983611248565b838681518110610cb757634e487b7160e01b600052603260045260246000fd5b60209081029190910101516001600160a01b03909116905260018101548154610ce09190611a9c565b838681518110610d0057634e487b7160e01b600052603260045260246000fd5b602090810291909101810151015284610d1881611ae3565b955050505b80610d2781611ae3565b915050610c34565b50826001600160401b03811115610d5657634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610d9b57816020015b6040805180820190915260008082526020820152815260200190600190039081610d745790505b50945060005b83811015610e0f57818181518110610dc957634e487b7160e01b600052603260045260246000fd5b6020026020010151868281518110610df157634e487b7160e01b600052603260045260246000fd5b60200260200101819052508080610e0790611ae3565b915050610da1565b5042600c8190555050505050909192565b610e6b6040518060c001604052806000815260200160008152602001600081526020016000815260200160006001600160a01b0316815260200160006001600160a01b031681525090565b506000828152602081815260408083208484526001908101835292819020815160c0810183528154815293810154928401929092526002820154908301526003810154606083015260048101546001600160a01b0390811660808401526005909101541660a08201525b92915050565b600080826001600160401b03811115610f0457634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610f2d578160200160208202803683370190505b50905060008080805b868110156111c657600080898984818110610f6157634e487b7160e01b600052603260045260246000fd5b905060600201600001358a8a85818110610f8b57634e487b7160e01b600052603260045260246000fd5b905060600201602001358b8b86818110610fb557634e487b7160e01b600052603260045260246000fd5b9050606002016040016020810190610fcd91906118e2565b604051602001610ffd93929190928352602083019190915260f81b6001600160f81b031916604082015260410190565b6040516020818303038152906040529050600061102261101c8d611254565b8361128f565b506001546040516330af0bbf60e21b81526001600160a01b03808416600483015292935091169063c2bc2efc9060240160006040518083038186803b15801561106a57600080fd5b505afa15801561107e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110a691908101906117d2565b606001516110b6575050506111b4565b60005b8781101561111d57816001600160a01b03168982815181106110eb57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b0316141561110b576001935061111d565b8061111581611ae3565b9150506110b9565b50821561112c575050506111b4565b8d6001600160a01b0316816001600160a01b0316141561114b57600194505b8c6001600160a01b0316816001600160a01b0316141561116a57600195505b8088888151811061118b57634e487b7160e01b600052603260045260246000fd5b6001600160a01b03909216602092830291909101909101526111ae600188611a70565b96505050505b806111be81611ae3565b915050610f36565b5060026111d4876001611a70565b6111de9190611a88565b8310806111e9575080155b806111f2575081155b1561120457600094505050505061120d565b60019450505050505b95945050505050565b600061122b836001600160a01b0384166112ff565b9392505050565b600061122b83836112ff565b6000610ed5825490565b600061122b838361134e565b60006112608251611386565b82604051602001611272929190611951565b604051602081830303815290604052805190602001209050919050565b6000808251604114156112c65760208301516040840151606085015160001a6112ba878285856114a8565b945094505050506112f8565b8251604014156112f057602083015160408401516112e5868383611595565b9350935050506112f8565b506000905060025b9250929050565b600081815260018301602052604081205461134657508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610ed5565b506000610ed5565b600082600001828154811061137357634e487b7160e01b600052603260045260246000fd5b9060005260206000200154905092915050565b6060816113ab57506040805180820190915260018152600360fc1b60208201526106bb565b8160005b81156113d557806113bf81611ae3565b91506113ce9050600a83611a88565b91506113af565b6000816001600160401b038111156113fd57634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015611427576020820181803683370190505b5090505b84156114a05761143c600183611a9c565b9150611449600a86611afe565b611454906030611a70565b60f81b81838151811061147757634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a905350611499600a86611a88565b945061142b565b949350505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156114df575060009050600361158c565b8460ff16601b141580156114f757508460ff16601c14155b15611508575060009050600461158c565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561155c573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166115855760006001925092505061158c565b9150600090505b94509492505050565b6000806001600160ff1b038316816115b260ff86901c601b611a70565b90506115c0878288856114a8565b935093505050935093915050565b80516106bb81611b54565b600082601f8301126115e9578081fd5b815160206001600160401b038083111561160557611605611b3e565b611613828460051b01611a40565b83815282810190868401865b8681101561168e5781518901606080601f19838e0301121561163f57898afd5b61164881611a40565b8883015188811115611658578b8cfd5b6116668e8b838701016116ac565b8252506040838101518a8301529190920151908201528452928501929085019060010161161f565b509098975050505050505050565b805180151581146106bb57600080fd5b600082601f8301126116bc578081fd5b81516001600160401b038111156116d5576116d5611b3e565b6116e8601f8201601f1916602001611a40565b8181528460208386010111156116fc578283fd5b6114a0826020830160208701611ab3565b60006020828403121561171e578081fd5b813561122b81611b54565b600080600080600080600060c0888a031215611743578283fd5b873561174e81611b54565b96506020880135955060408801359450606088013593506080880135925060a08801356001600160401b0380821115611785578384fd5b818a0191508a601f830112611798578384fd5b8135818111156117a6578485fd5b8b60206060830285010111156117ba578485fd5b60208301945080935050505092959891949750929550565b6000602082840312156117e3578081fd5b81516001600160401b03808211156117f9578283fd5b9083019060e0828603121561180c578283fd5b61181660e0611a40565b61181f836115ce565b8152602083015182811115611832578485fd5b61183e878286016116ac565b6020830152506040830151604082015261185a6060840161169c565b6060820152608083015182811115611870578485fd5b61187c878286016115d9565b60808301525061188e60a084016115ce565b60a082015260c083015160c082015280935050505092915050565b6000602082840312156118ba578081fd5b5035919050565b600080604083850312156118d3578182fd5b50508035926020909101359150565b6000602082840312156118f3578081fd5b813560ff8116811461122b578182fd5b6000815180845260208085019450808401835b8381101561194657815180516001600160a01b031688528301518388015260409096019590820190600101611916565b509495945050505050565b60007f19457468657265756d205369676e6564204d6573736167653a0a0000000000008252835161198981601a850160208801611ab3565b8351908301906119a081601a840160208801611ab3565b01601a01949350505050565b6000606082526119bf6060830186611903565b828103602084810191909152855180835286820192820190845b81811015611a09576119f683865180518252602090810151910152565b93830193604092909201916001016119d9565b50508481036040860152611a1d8187611903565b98975050505050505050565b815181526020808301519082015260408101610ed5565b604051601f8201601f191681016001600160401b0381118282101715611a6857611a68611b3e565b604052919050565b60008219821115611a8357611a83611b12565b500190565b600082611a9757611a97611b28565b500490565b600082821015611aae57611aae611b12565b500390565b60005b83811015611ace578181015183820152602001611ab6565b83811115611add576000848401525b50505050565b6000600019821415611af757611af7611b12565b5060010190565b600082611b0d57611b0d611b28565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114611b6957600080fd5b5056fea2646970667358221220e59bbe6f33c081ec936bc74f39c9a70431de00775a4b152da3351581f836693564736f6c63430008030033",
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
