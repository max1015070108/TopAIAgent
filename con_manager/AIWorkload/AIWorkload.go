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
	Bin: "0x608034620001b757601f62001f7f38819003918201601f191683019291906001600160401b03841183851017620001bc578160809284926040968752833981010312620001b7576200005181620001d2565b6200005f60208301620001d2565b916200007b606062000073868401620001d2565b9201620001d2565b610e10600c556001600160a01b03928316939091908415620001735783169081156200012f578316908115620000eb5760018060a01b03199485600154161760015584600254161760025542600d5583600e541617600e551690600f541617600f5551611d979081620001e88239f35b855162461bcd60e51b815260206004820152601360248201527f496e76616c6964207374616b6520746f6b656e000000000000000000000000006044820152606490fd5b855162461bcd60e51b815260206004820152601660248201527f496e76616c6964206d6f64656c207265676973747279000000000000000000006044820152606490fd5b855162461bcd60e51b815260206004820152601560248201527f496e76616c6964206e6f646520726567697374727900000000000000000000006044820152606490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b0382168203620001b75756fe6080604052600436101561001257600080fd5b60003560e01c80630f17c347146100f257806320737af1146100ed5780632ba36f78146100e85780632c2e139b146100e35780634ca3ce67146100de57806351160630146100d957806351ed6a30146100d45780637ec4142c146100cf57806383c4b7a3146100ca578063ad1f5717146100ca578063d14c9fa7146100c5578063d9564c36146100c0578063d9b5c4a5146100bb5763e926805e146100b657600080fd5b61083e565b610815565b6106de565b6103e4565b6102f8565b6102da565b6102b1565b610288565b610205565b6101bb565b610181565b610125565b610107565b600091031261010257565b600080fd5b34610102576000366003190112610102576020600d54604051908152f35b346101025760203660031901126101025761013e611ad0565b506004356000526006602052604080600020600182519161015e836109ce565b805483520154602082015261017f8251809260208091805184520151910152565bf35b34610102576000366003190112610102576002546040516001600160a01b039091168152602090f35b6001600160a01b0381160361010257565b34610102576020366003190112610102576004356101d8816101aa565b6101e0611ad0565b5060018060a01b03166000526009602052604080600020600182519161015e836109ce565b346101025760e036600319011261010257600435610222816101aa565b60243561022e816101aa565b60c4359167ffffffffffffffff91828411610102573660238501121561010257836004013592831161010257366024606085028601011161010257602461028694019160a43591608435916064359160443591610edf565b005b3461010257600036600319011261010257600f546040516001600160a01b039091168152602090f35b3461010257600036600319011261010257600e546040516001600160a01b039091168152602090f35b34610102576000366003190112610102576020600c54604051908152f35b346101025760203660031901126101025760043560005260006020526020604060002054604051908152f35b90815180825260208080930193019160005b828110610344575050505090565b835180516001600160a01b031686528201518583015260409094019392810192600101610336565b93929061038190606086526060860190610324565b936020948181038683015285808551928381520194019060005b8181106103bb575050506103b89394506040818403910152610324565b90565b909194876040826103d96001948a5160208091805184520151910152565b01960192910161039b565b34610102576000806003193601126106db57610410610408600d54600c5490610d89565b421015611ae9565b60045461041c81611b2e565b90829083905b80821061065757505061043481611b2e565b91835b828110610644575050506007549161044e83611b2e565b92819082905b8082106105cc57505061046681611b2e565b8094835b8381106105a35750505050600a5461048181611b2e565b829183905b8082106104f057505061049882611b2e565b925b8281106104c1575050506104bd906104b142600d55565b6040519384938461036c565b0390f35b806104cf6104eb9284610e5b565b516104da8287610e5b565b526104e58186610e5b565b5061137f565b61049a565b909261052a61051061050461050487611be8565b6001600160a01b031690565b6001600160a01b0316600090815260096020526040902090565b600181549101549081811115610597579161057761058b92610591946105726105586105046105048c611be8565b610562868b610e5b565b516001600160a01b039091169052565b6119e7565b60206105838388610e5b565b51015261137f565b9361137f565b90610486565b5050926105919061137f565b806105b16105c79284610e5b565b516105bc8286610e5b565b526104e58185610e5b565b61046a565b90916105ea6105da84611bb3565b6000526006602052604060002090565b600181549101549081811115610638579161062061062c926106329461060f88611bb3565b610619858d610e5b565b51526119e7565b6020610583838b610e5b565b9261137f565b90610454565b5050916106329061137f565b806104cf6106529284610e5b565b610437565b909161068561066b61050461050486611b7e565b6001600160a01b0316600090815260036020526040902090565b6001815491015490818111156106cf57916106bd61062c926106c9946105726106b36105046105048b611b7e565b610562868c610e5b565b60206105838389610e5b565b90610422565b5050916106c99061137f565b80fd5b346101025760408060031936011261010257808060c06104bd9351610702816109ef565b600091818380935282602082015282858201528260608201528260808201528260a08201520152600435815280602052600182822001602435825260205220906107bb6107ab6006835194610756866109ef565b80548652600181015460208701526002810154858701526003810154606087015260048101546001600160a01b0316608087015260058101546001600160a01b031660a087015201546001600160a01b031690565b6001600160a01b031660c0840152565b519182918291909160c060e0820193805183526020810151602084015260408101516040840152606081015160608401528160018060a01b03918260808201511660808601528260a08201511660a0860152015116910152565b34610102576000366003190112610102576001546040516001600160a01b039091168152602090f35b346101025760203660031901126101025760043561085b816101aa565b610863611ad0565b5060018060a01b03166000526003602052604080600020600182519161015e836109ce565b1561088f57565b60405162461bcd60e51b8152602060048201526015602482015274496e76616c6964206f776e6572206164647265737360581b6044820152606490fd5b156108d357565b60405162461bcd60e51b815260206004820152602260248201527f576f726b6c6f6164206d7573742062652067726561746572207468616e207a65604482015261726f60f01b6064820152608490fd5b1561092a57565b60405162461bcd60e51b815260206004820152602560248201527f4c656e677468206f66207369676e617475726573206d757374206d6f7265207460448201526468616e203360d81b6064820152608490fd5b1561098457565b60405162461bcd60e51b815260206004820152600c60248201526b24b73b30b634b2103ab9b2b960a11b6044820152606490fd5b634e487b7160e01b600052604160045260246000fd5b6040810190811067ffffffffffffffff8211176109ea57604052565b6109b8565b60e0810190811067ffffffffffffffff8211176109ea57604052565b67ffffffffffffffff81116109ea57604052565b6060810190811067ffffffffffffffff8211176109ea57604052565b90601f8019910116810190811067ffffffffffffffff8211176109ea57604052565b60405190610a6a826109ef565b565b67ffffffffffffffff81116109ea57601f01601f191660200190565b60405190610a95826109ce565b602082527f3d3d3d3d207265706f7274576f726b6c6f61642065706f6368496420697320326020830152565b60005b838110610ad45750506000910152565b8181015183820152602001610ac4565b81601f82011215610102578051610afa81610a6c565b92610b086040519485610a3b565b81845260208284010111610102576103b89160208085019101610ac1565b5190610a6a826101aa565b909160e0828403126101025781519260208301519167ffffffffffffffff928381116101025782610b63918601610ae4565b9360408101518481116101025783610b7c918301610ae4565b936060820151610b8b816101aa565b93608083015191821161010257610ba3918301610ae4565b9160c060a083015192015190565b6040513d6000823e3d90fd5b60405190610bca826109ce565b600b82526a3a36b826b7b232b624b21d60a91b6020830152565b15610beb57565b60405162461bcd60e51b815260206004820152600f60248201526e135bd9195b081b9bdd08195e1a5cdd608a1b6044820152606490fd5b60405190610c2f826109ce565b602082527f3d3d3d3d207265706f7274576f726b6c6f61642065706f6368496420697320336020830152565b15610c6257565b60405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b6044820152606490fd5b15610ca257565b60405162461bcd60e51b815260206004820152601260248201527122b837b1b41037baba1037b31037b93232b960711b6044820152606490fd5b8151815560208201516001820155604082015160028201556060820151600382015560808201516004820180546001600160a01b03199081166001600160a01b039384161790915560a084015160058401805491841691831691909117905560c090930151600690920180549093169116179055565b634e487b7160e01b600052601160045260246000fd5b9060018201809211610d7657565b610d52565b6030019081603011610d7657565b91908201809211610d7657565b60405190610da3826109ce565b601a82527f3d3d3d3d2041495061792065706f636849642069733a202573200000000000006020830152565b67ffffffffffffffff81116109ea5760051b60200190565b60405190610df4826109ce565b6001825260203681840137565b90610e0b82610dcf565b610e186040519182610a3b565b8281528092610e29601f1991610dcf565b0190602036910137565b634e487b7160e01b600052603260045260246000fd5b805115610e565760200190565b610e33565b8051821015610e565760209160051b010190565b9396959492919060c0850190855260018060a01b039182602091168187015260c0604087015283518092528060e087019401926000905b838210610ec65750505050509060a0929195606083015260808201520152565b8451811686529482019493820193600190910190610ea6565b94919395909296610f6d60018060a01b0391610efe8389161515610888565b610f098715156108cc565b610f166003821015610923565b610f23838716151561097d565b610f33610f2e610a88565b611c42565b6002546000908b90610f4d906001600160a01b0316610504565b6040518080968194634dd86b0960e11b8352600483019190602083019252565b03915afa9081156111f8578594876110098b610ffb8f8f818f610fb38e946110109c6110159e6000916111fd575b50610fad81610fa8610bbd565b611c8b565b14610be4565b610fbe610f2e610c22565b604051978896602088019260a09491979695929760c0850198600180881b0380921686521660208501526040840152606083015260808201520152565b03601f198101835282610a3b565b338b6115ed565b610c5b565b6110c061102c886000526000602052604060002090565b9161103983548511610c9b565b611094611044610a5d565b85815260208101899052426040820152606081018c90523360808201526001600160a01b038a1660a08201526001600160a01b03881660c082015260008681526001860160205260409020610cdc565b6001600160a01b03881660009081526003602052604090206110b7888254610d89565b90558716611226565b506110d5886000526006602052604060002090565b6110e0868254610d89565b90556110eb8861129d565b50336000908152600960205260409020611106868254610d89565b90556111113361130e565b5055604080516001600160a01b03861681526020810183905290810184905260608101879052339086907f59164c09eed89694722e80b610592178b04fbe4a72b982bf4503c8ab426f0ef690608090a361116d81610fa8610d96565b611191611178610de7565b9461118286610e49565b6001600160a01b039091169052565b600f546111a6906001600160a01b0316610504565b92833b15610102576111d4600096928793604051998a988997889663155219b560e21b885260048801610e6f565b03925af180156111f8576111e55750565b806111f2610a6a92610a0b565b806100f7565b610bb1565b61121a91503d806000833e6112128183610a3b565b810190610b31565b50505050505038610f9b565b8060005260056020526040600020541560001461129757600454600160401b8110156109ea576001810180600455811015610e565781907f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0155600454906000526005602052604060002055600190565b50600090565b8060005260086020526040600020541560001461129757600754600160401b8110156109ea576001810180600755811015610e565781907fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880155600754906000526008602052604060002055600190565b80600052600b6020526040600020541560001461129757600a54600160401b8110156109ea576001810180600a55811015610e565781907fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80155600a5490600052600b602052604060002055600190565b6000198114610d765760010190565b9190811015610e56576060020190565b3560ff811681036101025790565b604051906113b9826109ce565b60098252682fb0b2323932b9b99d60b91b6020830152565b5190811515820361010257565b81601f820112156101025780516113f481610dcf565b9260409161140483519586610a3b565b808552602093848087019260051b8201019383851161010257858201925b858410611433575050505050505090565b835167ffffffffffffffff9081811161010257840191606080601f19858a0301126101025784519061146482610a1f565b8a85015193841161010257846114818a8d80989781980101610ae4565b83528681015185840152015185820152815201930192611422565b9060208282031261010257815167ffffffffffffffff92838211610102570160e081830312610102576114cd610a5d565b926114d782610b26565b8452602082015181811161010257836114f1918401610ae4565b60208501526040820151604085015261150c606083016113d1565b606085015260808201519081116101025760c09261152b9183016113de565b608084015261153c60a08201610b26565b60a0840152015160c082015290565b60405190611558826109ce565b60078252663bb7b935b2b91d60c91b6020830152565b6040519061157b826109ce565b60098252683932b837b93a32b91d60b91b6020830152565b604051906115a0826109ce565b601082526f031b7b73a30b4b739abb7b935b2b91d160851b6020830152565b604051906115cc826109ce565b6012825271031b7b73a30b4b739a932b837b93a32b91d160751b6020830152565b6115f685610e01565b946000908195829583916001996116166105048c5460018060a01b031690565b85851061169e57505050505061162e61163491610d68565b60011c90565b118015611696575b801561168e575b61164e575050505090565b6116899450611681929161166c6116749261166761154b565b611d13565b61166761156e565b61167c611593565b611cce565b61167c6115bf565b600090565b508315611643565b50821561163c565b90919293958560006117196116b48a848961138e565b3561170b6020610ffb8d6116e48c6116de856116d1858d8561138e565b01359360409b8c9361138e565b0161139e565b908851958694850191604193918352602083015260ff60f81b9060f81b1660408201520190565b6117148961196d565b611855565b5091611727836116676113ac565b516330af0bbf60e21b81526001600160a01b03831660048201528181602481885afa80156111f85760606117689161176c938591611833575b500151151590565b1590565b61182757805b8381106117e4575b506117d9578c6001600160a01b03828116908b811682146117d0575b8c16146117c8575b50816117b46117b9926111826117bf9588610e5b565b610d68565b9661137f565b93929190611616565b9b508c61179e565b9c50819c611796565b50956117bf9061137f565b6117fe6117f18288610e5b565b516001600160a01b031690565b6001600160a01b0384811691161461181e576118199061137f565b611772565b50508c3861177a565b5050956117bf9061137f565b61184f91503d8087833e6118478183610a3b565b81019061149c565b38611760565b81516041810361188257509061187e916020820151906060604084015193015160001a906118c5565b9091565b6040036118bb5760406020830151920151918260ff1c91601b8301809311610d765761187e936001600160ff1b03169260ff16906118c5565b5050600090600290565b9291907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083116119615760ff16601b81141580611956575b61194a579160809493916020936040519384528484015260408301526060820152600093849182805260015afa156111f85781516001600160a01b03811615611944579190565b50600190565b50505050600090600490565b50601c8114156118fd565b50505050600090600390565b6119778151611a05565b906119e1603a604051809360208201957f19457468657265756d205369676e6564204d6573736167653a0a00000000000087526119bd8151809260208787019101610ac1565b82016119d28251809360208785019101610ac1565b0103601a810184520182610a3b565b51902090565b91908203918211610d7657565b908151811015610e56570160200190565b8015611ab257806000908282935b611a9e5750611a2183610a6c565b92611a2f6040519485610a3b565b80845281601f19611a3f83610a6c565b013660208701375b611a515750505090565b6000198101908111610d76578091600a91611a95611a8c611a7c611a76868606610d7b565b60ff1690565b60f81b6001600160f81b03191690565b861a91876119f4565b53049081611a47565b92611aaa600a9161137f565b930480611a13565b50604051611abf816109ce565b60018152600360fc1b602082015290565b60405190611add826109ce565b60006020838281520152565b15611af057565b60405162461bcd60e51b815260206004820152601660248201527514d95d1d1b195b595b9d081b9bdd08191d59481e595d60521b6044820152606490fd5b90611b3882610dcf565b611b456040519182610a3b565b8281528092611b56601f1991610dcf565b019060005b828110611b6757505050565b602090611b72611ad0565b82828501015201611b5b565b600454811015610e565760046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b015490565b600754811015610e565760076000527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688015490565b600a54811015610e5657600a6000527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8015490565b90602091611c3681518092818552858086019101610ac1565b601f01601f1916010190565b610ffb611c70610a6a9260405192839163104c13eb60e21b6020840152602060248401526044830190611c1d565b600080916020815191016a636f6e736f6c652e6c6f675afa50565b611c70611cba91610a6a93604051938492632d839cb360e21b6020850152604060248501526064840190611c1d565b90604483015203601f198101835282610a3b565b611c70611cfd91610a6a9360405193849263c3b5563560e01b6020850152604060248501526064840190611c1d565b901515604483015203601f198101835282610a3b565b611c70611d4291610a6a9360405193849263319af33360e01b6020850152604060248501526064840190611c1d565b6001600160a01b0391909116604483015203601f198101835282610a3b56fea264697066735822122066db0c5c9057d72b70f1b6a05e95b9ac126659d55b0094a34b619072ee36f64464736f6c63430008140033",
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
