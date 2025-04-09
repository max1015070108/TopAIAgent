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
	User      common.Address
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_modelRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_settlement\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"WorkloadReported\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"}],\"name\":\"getLastEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"getNodeWorkload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"internalType\":\"structAIWorkload.Workload\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"getTotalModelWorkload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalWorkload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledWorkload\",\"type\":\"uint256\"}],\"internalType\":\"structAIWorkload.WorkLoad\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"getTotalWorkReport\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalWorkload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledWorkload\",\"type\":\"uint256\"}],\"internalType\":\"structAIWorkload.WorkLoad\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getTotalWorkerWorkload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalWorkload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settledWorkload\",\"type\":\"uint256\"}],\"internalType\":\"structAIWorkload.WorkLoad\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSettlementTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modelRegistry\",\"outputs\":[{\"internalType\":\"contractAIModels\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeRegistry\",\"outputs\":[{\"internalType\":\"contractNodesRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"reportWorkload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sessions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastEpochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settleRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"}],\"internalType\":\"structNodeSettleWorkload[]\",\"name\":\"settledWorkers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"}],\"internalType\":\"structModelSettleWorkload[]\",\"name\":\"settledModels\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workload\",\"type\":\"uint256\"}],\"internalType\":\"structNodeSettleWorkload[]\",\"name\":\"settledReporters\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlement\",\"outputs\":[{\"internalType\":\"contractSettlement\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIStake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608034620001bc57601f62001e2238819003918201601f191683019291906001600160401b03841183851017620001c1578160809284926040968752833981010312620001bc576200005181620001d7565b6200005f60208301620001d7565b916200007b606062000073868401620001d7565b9201620001d7565b6001600055610e10600d556001600160a01b039283169390919084156200017857831690811562000134578316908115620000f05760018060a01b03199485600254161760025584600354161760035542600e5583600f541617600f551690601054161760105551611c359081620001ed8239f35b855162461bcd60e51b815260206004820152601360248201527f496e76616c6964207374616b6520746f6b656e000000000000000000000000006044820152606490fd5b855162461bcd60e51b815260206004820152601660248201527f496e76616c6964206d6f64656c207265676973747279000000000000000000006044820152606490fd5b855162461bcd60e51b815260206004820152601560248201527f496e76616c6964206e6f646520726567697374727900000000000000000000006044820152606490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b0382168203620001bc5756fe6080604052600436101561001257600080fd5b60003560e01c80630f17c347146100f257806320737af1146100ed5780632ba36f78146100e85780632c2e139b146100e35780634ca3ce67146100de57806351160630146100d957806351ed6a30146100d45780637ec4142c146100cf57806383c4b7a3146100ca578063ad1f5717146100ca578063d14c9fa7146100c5578063d9564c36146100c0578063d9b5c4a5146100bb5763e926805e146100b657600080fd5b61078a565b610761565b610629565b6103d4565b6102e8565b6102ca565b6102a1565b610278565b6101f5565b6101ab565b610171565b610115565b34610110576000366003190112610110576020600e54604051908152f35b600080fd5b346101105760203660031901126101105761012e6119f6565b506004356000526007602052604080600020600182519161014e83610c1d565b805483520154602082015261016f8251809260208091805184520151910152565bf35b34610110576000366003190112610110576003546040516001600160a01b039091168152602090f35b6001600160a01b0381160361011057565b34610110576020366003190112610110576004356101c88161019a565b6101d06119f6565b5060018060a01b0316600052600a602052604080600020600182519161014e83610c1d565b346101105760e0366003190112610110576004356102128161019a565b60243561021e8161019a565b60c4359167ffffffffffffffff91828411610110573660238501121561011057836004013592831161011057366024606085028601011161011057602461027694019160a435916084359160643591604435916107d4565b005b34610110576000366003190112610110576010546040516001600160a01b039091168152602090f35b3461011057600036600319011261011057600f546040516001600160a01b039091168152602090f35b34610110576000366003190112610110576020600d54604051908152f35b346101105760203660031901126101105760043560005260016020526020604060002054604051908152f35b90815180825260208080930193019160005b828110610334575050505090565b835180516001600160a01b031686528201518583015260409094019392810192600101610326565b93929061037190606086526060860190610314565b936020948181038683015285808551928381520194019060005b8181106103ab575050506103a89394506040818403910152610314565b90565b909194876040826103c96001948a5160208091805184520151910152565b01960192910161038b565b3461011057600080600319360112610626576104006103f8600e54600d5490610f4a565b421015611a0f565b60055461040c81611a54565b90825b8181106105965750506008549161042583611a54565b92815b818110610518575050600b5461043d81611a54565b91805b828110610467575050506104639061045742600e55565b6040519384938461035c565b0390f35b806104a361048961047d61047d6104f695611b0e565b6001600160a01b031690565b6001600160a01b03166000908152600a6020526040902090565b80546001909101548082116104fb5750508260206104c18388610fe3565b5101525b6104f16104d761047d61047d84611b0e565b6104e18388610fe3565b516001600160a01b039091169052565b611150565b610440565b6105049161190d565b60206105108388610fe3565b5101526104c5565b8061053861052861057493611ad9565b6000526007602052604060002090565b80546001909101548082116105795750508360206105568389610fe3565b5101525b61056381611ad9565b61056d8288610fe3565b5152611150565b610428565b6105829161190d565b602061058e8389610fe3565b51015261055a565b806105c66105ac61047d61047d61060495611aa4565b6001600160a01b0316600090815260046020526040902090565b80546001909101548082116106095750508460206105e48387610fe3565b5101525b6104f16105fa61047d61047d84611aa4565b6104e18387610fe3565b61040f565b6106129161190d565b602061061e8387610fe3565b5101526105e8565b80fd5b346101105760408060031936011261011057808060c0610463935161064d81610c3e565b600091818380935282602082015282858201528260608201528260808201528260a0820152015260043581526001602052600182822001602435825260205220906107076106f760068351946106a286610c3e565b80548652600181015460208701526002810154858701526003810154606087015260048101546001600160a01b0316608087015260058101546001600160a01b031660a087015201546001600160a01b031690565b6001600160a01b031660c0840152565b519182918291909160c060e0820193805183526020810151602084015260408101516040840152606081015160608401528160018060a01b03918260808201511660808601528260a08201511660a0860152015116910152565b34610110576000366003190112610110576002546040516001600160a01b039091168152602090f35b34610110576020366003190112610110576004356107a78161019a565b6107af6119f6565b5060018060a01b03166000526004602052604080600020600182519161014e83610c1d565b9592919096600260005414610a925760026000556001600160a01b038781169890610825906108048b1515610ad7565b61080f851515610b1b565b61081c6003851015610b72565b82161515610bcc565b60035461083a906001600160a01b031661047d565b604051634dd86b0960e11b8152600481018690529590600090879060249082905afa998a15610a8d5761091261090d8a956109c29488610a5a9f8f90898f7f59164c09eed89694722e80b610592178b04fbe4a72b982bf4503c8ab426f0ef69f8f91828f926108f894610906976108ba92600091610a66575b5014610dde565b60405197889588602088019260a09491979695929760c0850198600180881b0380921686521660208501526040840152606083015260808201520152565b03601f198101845283610c76565b33906114a1565b610e1c565b6109988a61098361092d8b6000526001602052604060002090565b9561093a87548911610e5c565b610973610945610c98565b898152602081018b9052426040820152606081018c9052336080820152936001600160a01b031660a0850152565b6001600160a01b031660c0830152565b60008681526001860160205260409020610e9d565b6001600160a01b038a1660009081526004602052604090206109bb868254610f4a565b9055610ff7565b506109d7846000526007602052604060002090565b6109e2848254610f4a565b90556109ed8461106e565b50336000908152600a60205260409020610a08848254610f4a565b9055610a13336110df565b5055604080516001600160a01b0388168152602081019690965285015260608401523392608090a3610a4b610a46610f6f565b610fd1565b6001600160a01b039091169052565b610a646001600055565b565b610a81913d8091833e610a798183610c76565b810190610d52565b505050505050386108b3565b610dd2565b60405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606490fd5b15610ade57565b60405162461bcd60e51b8152602060048201526015602482015274496e76616c6964206f776e6572206164647265737360581b6044820152606490fd5b15610b2257565b60405162461bcd60e51b815260206004820152602260248201527f576f726b6c6f6164206d7573742062652067726561746572207468616e207a65604482015261726f60f01b6064820152608490fd5b15610b7957565b60405162461bcd60e51b815260206004820152602560248201527f4c656e677468206f66207369676e617475726573206d757374206d6f7265207460448201526468616e203360d81b6064820152608490fd5b15610bd357565b60405162461bcd60e51b815260206004820152600c60248201526b24b73b30b634b2103ab9b2b960a11b6044820152606490fd5b634e487b7160e01b600052604160045260246000fd5b6040810190811067ffffffffffffffff821117610c3957604052565b610c07565b60e0810190811067ffffffffffffffff821117610c3957604052565b6060810190811067ffffffffffffffff821117610c3957604052565b90601f8019910116810190811067ffffffffffffffff821117610c3957604052565b60405190610a6482610c3e565b60405190610100820182811067ffffffffffffffff821117610c3957604052565b67ffffffffffffffff8111610c3957601f01601f191660200190565b60005b838110610cf55750506000910152565b8181015183820152602001610ce5565b81601f82011215610110578051610d1b81610cc6565b92610d296040519485610c76565b81845260208284010111610110576103a89160208085019101610ce2565b5190610a648261019a565b909160e0828403126101105781519260208301519167ffffffffffffffff928381116101105782610d84918601610d05565b9360408101518481116101105783610d9d918301610d05565b936060820151610dac8161019a565b93608083015191821161011057610dc4918301610d05565b9160c060a083015192015190565b6040513d6000823e3d90fd5b15610de557565b60405162461bcd60e51b815260206004820152600f60248201526e135bd9195b081b9bdd08195e1a5cdd608a1b6044820152606490fd5b15610e2357565b60405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b6044820152606490fd5b15610e6357565b60405162461bcd60e51b815260206004820152601260248201527122b837b1b41037baba1037b31037b93232b960711b6044820152606490fd5b8151815560208201516001820155604082015160028201556060820151600382015560808201516004820180546001600160a01b03199081166001600160a01b039384161790915560a084015160058401805491841691831691909117905560c090930151600690920180549093169116179055565b634e487b7160e01b600052601160045260246000fd5b9060018201809211610f3757565b610f13565b6030019081603011610f3757565b91908201809211610f3757565b67ffffffffffffffff8111610c395760051b60200190565b60405190610f7c82610c1d565b6001825260203681840137565b90610f9382610f57565b610fa06040519182610c76565b8281528092610fb1601f1991610f57565b0190602036910137565b634e487b7160e01b600052603260045260246000fd5b805115610fde5760200190565b610fbb565b8051821015610fde5760209160051b010190565b8060005260066020526040600020541560001461106857600554600160401b811015610c39576001810180600555811015610fde5781907f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00155600554906000526006602052604060002055600190565b50600090565b8060005260096020526040600020541560001461106857600854600160401b811015610c39576001810180600855811015610fde5781907ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155600854906000526009602052604060002055600190565b80600052600c6020526040600020541560001461106857600b54600160401b811015610c39576001810180600b55811015610fde5781907f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db90155600b5490600052600c602052604060002055600190565b6000198114610f375760010190565b9190811015610fde576060020190565b3560ff811681036101105790565b6040519061118a82610c1d565b60098252682fb0b2323932b9b99d60b91b6020830152565b604051906111af82610c1d565b60078252663bb7b935b2b91d60c91b6020830152565b604051906111d282610c1d565b60098252683932b837b93a32b91d60b91b6020830152565b604051906111f782610c1d565b600f82526e3737b232a932b3b4b9ba393c97171760891b6020830152565b5190811515820361011057565b81601f8201121561011057805161123881610f57565b9260409161124883519586610c76565b808552602093848087019260051b8201019383851161011057858201925b858410611277575050505050505090565b835167ffffffffffffffff9081811161011057840191606080601f19858a030112610110578451906112a882610c5a565b8a85015193841161011057846112c58a8d80989781980101610d05565b83528681015185840152015185820152815201930192611266565b60208183031261011057805167ffffffffffffffff918282116101105701916101008382031261011057611312610ca5565b9261131c81610d47565b845260208101518381116101105782611336918301610d05565b60208501526040810151604085015261135160608201611215565b606085015260808101519283116101105761137360e092611399948301611222565b608085015261138460a08201610d47565b60a085015260c081015160c085015201611215565b60e082015290565b604051906113ae82610c1d565b60128252711b9bd919549959da5cdd1c9e4b8b8b995b9960721b6020830152565b604051906113dc82610c1d565b600c82526b323ab83634b1b0ba3297171760a11b6020830152565b6040519061140482610c1d565b60128252712fb0b2323932b9b9901e9e903bb7b935b2b960711b6020830152565b6040519061143282610c1d565b60148252732fb0b2323932b9b9901e9e903932b837b93a32b960611b6020830152565b90816020910312610110576103a890611215565b1561147057565b60405162461bcd60e51b81526020600482015260096024820152684e6f742070726f787960b81b6044820152606490fd5b9190936114ad81610f89565b6002546000969193879290916114cb906001600160a01b031661047d565b925b8489106114e257505050505050505050600190565b8488600061156a8c61155c6114f882878761115f565b3561154e61152460209761151e89611511888d8561115f565b01359660409b8c9361115f565b0161116f565b8851888101938452602084019590955260f81b6001600160f81b0319166040830152839160410190565b03601f198101835282610c76565b61156588611893565b61177b565b509261157d8461157861117d565b611bb1565b6115898b6115786111a2565b611595866115786111c5565b6115a56115a06111ea565b611b68565b80516330af0bbf60e21b81526001600160a01b038516600480830191909152929081816024818d5afa8015610a8d5760606115ed916115f1938591611759575b500151151590565b1590565b61174a576116006115a06113a1565b805b8b878210611705575b50506116186115a06113cf565b6116f7576116276115a06113f7565b6116326115a0611425565b6001600160a01b038481169087161461166f575b5050508161165e61166392610a4b611669958b610fe3565b610f29565b98611150565b976114cd565b51632e2be8e760e21b81526001600160a01b0384169181019182529193919084908290819060200103818a5afa928315610a8d576116c16116639461165e93611669976000926116ca575b5050611469565b92829450611646565b6116e99250803d106116f0575b6116e18183610c76565b810190611455565b38806116ba565b503d6116d7565b505050509761166990611150565b6117128261171f92610fe3565b516001600160a01b031690565b6001600160a01b0387811691161461173f5761173a90611150565b611602565b50506001388b61160b565b50505050509761166990611150565b61177591503d8087833e61176d8183610c76565b8101906112e0565b386115e5565b8151604181036117a85750906117a4916020820151906060604084015193015160001a906117eb565b9091565b6040036117e15760406020830151920151918260ff1c91601b8301809311610f37576117a4936001600160ff1b03169260ff16906117eb565b5050600090600290565b9291907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083116118875760ff16601b8114158061187c575b611870579160809493916020936040519384528484015260408301526060820152600093849182805260015afa15610a8d5781516001600160a01b0381161561186a579190565b50600190565b50505050600090600490565b50601c811415611823565b50505050600090600390565b61189d815161192b565b90611907603a604051809360208201957f19457468657265756d205369676e6564204d6573736167653a0a00000000000087526118e38151809260208787019101610ce2565b82016118f88251809360208785019101610ce2565b0103601a810184520182610c76565b51902090565b91908203918211610f3757565b908151811015610fde570160200190565b80156119d857806000908282935b6119c4575061194783610cc6565b926119556040519485610c76565b80845281601f1961196583610cc6565b013660208701375b6119775750505090565b6000198101908111610f37578091600a916119bb6119b26119a261199c868606610f3c565b60ff1690565b60f81b6001600160f81b03191690565b861a918761191a565b5304908161196d565b926119d0600a91611150565b930480611939565b506040516119e581610c1d565b60018152600360fc1b602082015290565b60405190611a0382610c1d565b60006020838281520152565b15611a1657565b60405162461bcd60e51b815260206004820152601660248201527514d95d1d1b195b595b9d081b9bdd08191d59481e595d60521b6044820152606490fd5b90611a5e82610f57565b611a6b6040519182610c76565b8281528092611a7c601f1991610f57565b019060005b828110611a8d57505050565b602090611a986119f6565b82828501015201611a81565b600554811015610fde5760056000527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0015490565b600854811015610fde5760086000527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3015490565b600b54811015610fde57600b6000527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db9015490565b90602091611b5c81518092818552858086019101610ce2565b601f01601f1916010190565b61154e611b96610a649260405192839163104c13eb60e21b6020840152602060248401526044830190611b43565b600080916020815191016a636f6e736f6c652e6c6f675afa50565b611b96611be091610a649360405193849263319af33360e01b6020850152604060248501526064840190611b43565b6001600160a01b0391909116604483015203601f198101835282610c7656fea26469706673582212201c75c5186b357fdec397a3c81efd9eb3071e518762ab55dd8c837b3e07b59cc064736f6c63430008140033",
}

// AIWorkloadABI is the input ABI used to generate the binding from.
// Deprecated: Use AIWorkloadMetaData.ABI instead.
var AIWorkloadABI = AIWorkloadMetaData.ABI

// AIWorkloadBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AIWorkloadMetaData.Bin instead.
var AIWorkloadBin = AIWorkloadMetaData.Bin

// DeployAIWorkload deploys a new Ethereum contract, binding an instance of AIWorkload to it.
func DeployAIWorkload(auth *bind.TransactOpts, backend bind.ContractBackend, _nodeRegistry common.Address, _modelRegistry common.Address, _stakeToken common.Address, _settlement common.Address) (common.Address, *types.Transaction, *AIWorkload, error) {
	parsed, err := AIWorkloadMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AIWorkloadBin), backend, _nodeRegistry, _modelRegistry, _stakeToken, _settlement)
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
// Solidity: function getNodeWorkload(uint256 sessionId, uint256 epochId) view returns((uint256,uint256,uint256,uint256,address,address,address))
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
// Solidity: function getNodeWorkload(uint256 sessionId, uint256 epochId) view returns((uint256,uint256,uint256,uint256,address,address,address))
func (_AIWorkload *AIWorkloadSession) GetNodeWorkload(sessionId *big.Int, epochId *big.Int) (AIWorkloadWorkload, error) {
	return _AIWorkload.Contract.GetNodeWorkload(&_AIWorkload.CallOpts, sessionId, epochId)
}

// GetNodeWorkload is a free data retrieval call binding the contract method 0xd9564c36.
//
// Solidity: function getNodeWorkload(uint256 sessionId, uint256 epochId) view returns((uint256,uint256,uint256,uint256,address,address,address))
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

// Settlement is a free data retrieval call binding the contract method 0x51160630.
//
// Solidity: function settlement() view returns(address)
func (_AIWorkload *AIWorkloadCaller) Settlement(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIWorkload.contract.Call(opts, &out, "settlement")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Settlement is a free data retrieval call binding the contract method 0x51160630.
//
// Solidity: function settlement() view returns(address)
func (_AIWorkload *AIWorkloadSession) Settlement() (common.Address, error) {
	return _AIWorkload.Contract.Settlement(&_AIWorkload.CallOpts)
}

// Settlement is a free data retrieval call binding the contract method 0x51160630.
//
// Solidity: function settlement() view returns(address)
func (_AIWorkload *AIWorkloadCallerSession) Settlement() (common.Address, error) {
	return _AIWorkload.Contract.Settlement(&_AIWorkload.CallOpts)
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

// ReportWorkload is a paid mutator transaction binding the contract method 0x4ca3ce67.
//
// Solidity: function reportWorkload(address worker, address user, uint256 workload, uint256 modelId, uint256 sessionId, uint256 epochId, (bytes32,bytes32,uint8)[] signatures) returns()
func (_AIWorkload *AIWorkloadTransactor) ReportWorkload(opts *bind.TransactOpts, worker common.Address, user common.Address, workload *big.Int, modelId *big.Int, sessionId *big.Int, epochId *big.Int, signatures []Signature) (*types.Transaction, error) {
	return _AIWorkload.contract.Transact(opts, "reportWorkload", worker, user, workload, modelId, sessionId, epochId, signatures)
}

// ReportWorkload is a paid mutator transaction binding the contract method 0x4ca3ce67.
//
// Solidity: function reportWorkload(address worker, address user, uint256 workload, uint256 modelId, uint256 sessionId, uint256 epochId, (bytes32,bytes32,uint8)[] signatures) returns()
func (_AIWorkload *AIWorkloadSession) ReportWorkload(worker common.Address, user common.Address, workload *big.Int, modelId *big.Int, sessionId *big.Int, epochId *big.Int, signatures []Signature) (*types.Transaction, error) {
	return _AIWorkload.Contract.ReportWorkload(&_AIWorkload.TransactOpts, worker, user, workload, modelId, sessionId, epochId, signatures)
}

// ReportWorkload is a paid mutator transaction binding the contract method 0x4ca3ce67.
//
// Solidity: function reportWorkload(address worker, address user, uint256 workload, uint256 modelId, uint256 sessionId, uint256 epochId, (bytes32,bytes32,uint8)[] signatures) returns()
func (_AIWorkload *AIWorkloadTransactorSession) ReportWorkload(worker common.Address, user common.Address, workload *big.Int, modelId *big.Int, sessionId *big.Int, epochId *big.Int, signatures []Signature) (*types.Transaction, error) {
	return _AIWorkload.Contract.ReportWorkload(&_AIWorkload.TransactOpts, worker, user, workload, modelId, sessionId, epochId, signatures)
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
