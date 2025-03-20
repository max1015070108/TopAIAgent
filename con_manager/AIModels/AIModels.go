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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"ModelDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"ModelRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelInfo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"UploadModeled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"modelInfo\",\"type\":\"bytes\"}],\"name\":\"decodeModelInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelExtendInfo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelExtendInfo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"encodeModelInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"getModelDistribution\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getNodeDeployment\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"modelDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"modelIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"modelOwns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextModelId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeDeployment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelExtendInfo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"recordModelUpload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractNodesRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"removeDeployment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"reportDeployment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIStake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uploadModels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"modelName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"modelVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"extendInfo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608034620001a85762001c8990601f38839003908101601f19168201906001600160401b03821183831017620001ad5780839160409586948552833981010312620001a8576200004f81620001c3565b6200005e6020809301620001c3565b60016005556001600160a01b0391821691908215620001585716908115620001145760018060a01b0319908160015416176001556006541617600655600090818052818152828220338352815260ff838320541615620000c8575b8251611ab09081620001d98239f35b818052818152828220338084529152828220805460ff191660011790559081907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a43880620000b9565b835162461bcd60e51b815260048101849052601360248201527f496e76616c6964207374616b6520746f6b656e000000000000000000000000006044820152606490fd5b845162461bcd60e51b8152600481018590526024808201527f496e76616c6964207175616e74697479206f66207265676973747279206164646044820152637265737360e01b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b0382168203620001a85756fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a71461135e575080630d8bc09614611270578063248a9ca3146112415780632ef645371461118d5780632f2ff15d146110da57806336568abe1461108d57806351ed6a301461106457806357793e6814610fac57806364290cbf14610cab57806369f4296a14610c715780637b10399914610c485780637d5ef48e14610bf0578063907d004214610a0457806391d14854146109b75780639bb0d61214610903578063a217fddf146108e7578063c1ce7185146102ce578063c78717d514610288578063c8dad36014610235578063d547741f14610186578063dcca1c92146101345763e472ae8b1461011157600080fd5b3461012f57600036600319011261012f576020600554604051908152f35b600080fd5b3461012f57604036600319011261012f5760243560043560005260076020526040600020805482101561012f5760209161016d916114d9565b905460405160039290921b1c6001600160a01b03168152f35b3461012f57604036600319011261012f576004356101a2611484565b8160005260006020526101bc60016040600020015461167f565b81600052600060205260406000209060018060a01b0316908160005260205260ff604060002054166101ea57005b816000526000602052604060002081600052602052604060002060ff19815416905533917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b600080a4005b3461012f5761028461026261027061024c36611615565b93604098959893919351988997602089016119e6565b03601f1981018352826113b1565b604051918291602083526020830190611449565b0390f35b3461012f5760206102983661149a565b816102ac6040519283815193849201611426565b600490820190815281900382019020546040516001600160a01b039091168152f35b3461012f576102dc36611615565b6102ed9694929591963683876113ef565b9661034660216102fe3688886113ef565b996040519a816103188d9351809260208087019101611426565b8201602f60f81b60208201526103378251809360208785019101611426565b0103600181018b5201896113b1565b6040516020818161035d8c83815193849201611426565b81016002815203019020546108b45760055497604051988960e081011067ffffffffffffffff60e08c0111176107085760e08a01604052808a526103a236868a6113ef565b9060208b019182526103b53689896113ef565b60408c01523360608c01526103cb36858c6113ef565b60808c01524260a08c01528460c08c015260005260036020526040600020908a5182555180519067ffffffffffffffff82116107085781906104106001850154611507565b601f8111610861575b50602090601f83116001146107f4576000926107e9575b50508160011b916000199060031b1c19161760018201555b60408a015180519067ffffffffffffffff821161070857819061046e6002850154611507565b601f8111610796575b50602090601f83116001146107295760009261071e575b50508160011b916000199060031b1c19161760028201555b6003810160018060a01b0360608c0151166bffffffffffffffffffffffff60a01b82541617905560808a0151998a5167ffffffffffffffff8111610708576104f16004840154611507565b9b601f8d116106bd575b60209c508c90601f8311600114610620577fa17734f4a6749e7cc4947af9f12fb87299c2c2ab175158d74d24c8287fdee4369998979593836105ce98969460069460c094600092610615575b50508160011b916000199060031b1c19161760048501555b60a0810151600585015501519101556105898b600554928160405193828580945193849201611426565b81016002815203019020556040518489823760048186019081528190038b019020546001600160a01b0316156105e5575b600554988997604051968796339b886119e6565b0390a36105da81611a29565b600555604051908152f35b604051848982378a818681016004815203019020336bffffffffffffffffffffffff60a01b8254161790556105ba565b015190503880610547565b90600485016000528d6000209160005b601f19851681106106a657509360018460069460c0947fa17734f4a6749e7cc4947af9f12fb87299c2c2ab175158d74d24c8287fdee4369e9d9c9a986105ce9c9a98601f1981161061068d575b505050811b01600485015561055f565b015160001960f88460031b161c1916905538808061067d565b91928f600181928685015181550194019201610630565b600484016000526020600020601f830160051c810160208410610701575b601f8f5b0160051c820181106106f25750506104fb565b60008155600101601f8f6106df565b50806106db565b634e487b7160e01b600052604160045260246000fd5b015190508c8061048e565b600285016000908152602081209350601f198516905b81811061077e5750908460019594939210610765575b505050811b0160028201556104a6565b015160001960f88460031b161c191690558c8080610755565b9293602060018192878601518155019501930161073f565b909150600284016000526020600020601f840160051c8101602085106107e2575b90849392915b601f830160051c820181106107d3575050610477565b600081558594506001016107bd565b50806107b7565b015190508c80610430565b600185016000908152602081209350601f198516905b8181106108495750908460019594939210610830575b505050811b016001820155610448565b015160001960f88460031b161c191690558c8080610820565b9293602060018192878601518155019501930161080a565b909150600184016000526020600020601f840160051c8101602085106108ad575b90849392915b601f830160051c8201811061089e575050610419565b60008155859450600101610888565b5080610882565b60405162461bcd60e51b815260206004820152600b60248201526a135bd9195b08195e1a5cdd60aa1b6044820152606490fd5b3461012f57600036600319011261012f57602060405160008152f35b3461012f57602036600319011261012f5760043560005260036020526040600020805461093260018301611541565b9161093f60028201611541565b60038201546109a8906001600160a01b031661095d60048501611541565b6109956006600587015496015494610987604051998a998a5260e060208b015260e08a0190611449565b9088820360408a0152611449565b9160608701528582036080870152611449565b9160a084015260c08301520390f35b3461012f57604036600319011261012f576109d0611484565b600435600052600060205260406000209060018060a01b0316600052602052602060ff604060002054166040519015158152f35b3461012f57602036600319011261012f5760043560008181526007602052604081205b80549182811015610be857610a3c81836114d9565b90546003946001600160a01b0392909190861b1c82163314610a6b575050610a65919250611a29565b90610a27565b90916000199180830191908211610bc957610a9b84610a8d610aba94886114d9565b905490891b1c1691866114d9565b90919082549060031b9160018060a01b03809116831b921b1916179055565b82548015610bb3570192610ace84846114d9565b81939154921b1b19169055555b33600052600860205260406000208160005b825480821015610bdf578290610b0383866114d9565b929054600393841b1c14610b21575050610b1c90611a29565b610aed565b9093925060001991828201918211610bc957610b50610b43610b6993866114d9565b905490871b1c91856114d9565b90919082549060031b91821b91600019901b1916179055565b81548015610bb357810192610b7e84846114d9565b81939154921b1b19169055555b337f271097aa25aca38eb083e68e5cafcbd8247d421aea1887c9630402cda1706b63600080a3005b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b50505050610b8b565b505050610adb565b3461012f57604036600319011261012f57610c0961146e565b6001600160a01b031660009081526008602052604090208054602435919082101561012f57602091610c3a916114d9565b90546040519160031b1c8152f35b3461012f57600036600319011261012f576001546040516001600160a01b039091168152602090f35b3461012f576020610c9881610c853661149a565b8160405193828580945193849201611426565b8101600281520301902054604051908152f35b3461012f5760208060031936011261012f57600154604051631846d2f560e31b81523360048281019190915235926001600160a01b039260249290829082908590829088165afa908115610fa057600091610f6a575b5015610f33578360005260039283825260406000205415610efb5760008581526007835260408120805492915b838110610e93575050505083600052600781526040600020928354680100000000000000009485821015610e7e5781610d6f916001610d8d940181556114d9565b81546001600160a01b0360039290921b91821b19163390911b179055565b33600052600882526040600020906000918054925b838110610e1657505050506008903360005252604060002090815492831015610e025750610b508284926001610dda950181556114d9565b337fb81e8e03e6a54d2b07bc11d3cddad332ded9a7018c0e554bd78262a1d4cd20d1600080a3005b634e487b7160e01b60009081526041600452fd5b87610e2182846114d9565b905490851b1c14610e3a57610e3590611a29565b610da2565b60405162461bcd60e51b815260048101869052601d818801527f4e6f6465206465706c6f796d656e7420616c72656164792065786973740000006044820152606490fd5b84634e487b7160e01b60005260416004526000fd5b82610e9e82846114d9565b905490891b1c163314610eb957610eb490611a29565b610d2e565b606486866040519162461bcd60e51b83528160048401528201527f4d6f64656c20646973747269627574696f6e20616c72656164792065786973746044820152fd5b5060126064926040519262461bcd60e51b8452600484015282015271135bd9195b081a5cc81b9bdd08195e1a5cdd60721b6044820152fd5b60126064926040519262461bcd60e51b84526004840152820152714e6f6465206973206e6f742061637469766560701b6044820152fd5b90508181813d8311610f99575b610f8181836113b1565b8101031261012f5751801515810361012f5785610d01565b503d610f77565b6040513d6000823e3d90fd5b3461012f5760208060031936011261012f5760043560005260078152604060002060405190818382549182815201908192600052846000209060005b86828210611047578686610ffe828803836113b1565b604051928392818401908285525180915260408401929160005b82811061102757505050500390f35b83516001600160a01b031685528695509381019392810192600101611018565b83546001600160a01b031685529093019260019283019201610fe8565b3461012f57600036600319011261012f576006546040516001600160a01b039091168152602090f35b3461012f57604036600319011261012f576110a6611484565b5060405162461bcd60e51b815260206004820152600b60248201526a1b9bdd081cdd5c1c1bdc9d60aa1b6044820152606490fd5b3461012f57604036600319011261012f576004356110f6611484565b81600052600060205261111060016040600020015461167f565b81600052600060205260406000209060018060a01b0316908160005260205260ff604060002054161561113f57005b8160005260006020526040600020816000526020526040600020600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4005b3461012f5760208060031936011261012f576001600160a01b036111af61146e565b1660005260088152604060002060405190818382549182815201908192600052846000209060005b8682821061122d5786866111ed828803836113b1565b604051928392818401908285525180915260408401929160005b82811061121657505050500390f35b835185528695509381019392810192600101611207565b8354855290930192600192830192016111d7565b3461012f57602036600319011261012f5760043560005260006020526020600160406000200154604051908152f35b3461012f5760208060031936011261012f5767ffffffffffffffff60043581811161012f573660238201121561012f576112b49036906024816004013591016113ef565b908151820160808385830192031261012f578383015182811161012f5781856112df92860101611a38565b92604081015183811161012f5782866112fa92840101611a38565b92606082015190811161012f57856080611321611354956113399961134695870101611a38565b93015194604051978897608089526080890190611449565b9187830390880152611449565b908482036040860152611449565b9060608301520390f35b3461012f57602036600319011261012f576004359063ffffffff60e01b821680920361012f57602091637965db0b60e01b81149081156113a0575b5015158152f35b6301ffc9a760e01b14905083611399565b90601f8019910116810190811067ffffffffffffffff82111761070857604052565b67ffffffffffffffff811161070857601f01601f191660200190565b9291926113fb826113d3565b9161140960405193846113b1565b82948184528183011161012f578281602093846000960137010152565b60005b8381106114395750506000910152565b8181015183820152602001611429565b9060209161146281518092818552858086019101611426565b601f01601f1916010190565b600435906001600160a01b038216820361012f57565b602435906001600160a01b038216820361012f57565b602060031982011261012f576004359067ffffffffffffffff821161012f578060238301121561012f578160246114d6936004013591016113ef565b90565b80548210156114f15760005260206000200190600090565b634e487b7160e01b600052603260045260246000fd5b90600182811c92168015611537575b602083101461152157565b634e487b7160e01b600052602260045260246000fd5b91607f1691611516565b906040519182600082549261155584611507565b9081845260019485811690816000146115c45750600114611581575b505061157f925003836113b1565b565b9093915060005260209081600020936000915b8183106115ac57505061157f93508201013880611571565b85548884018501529485019487945091830191611594565b91505061157f94506020925060ff191682840152151560051b8201013880611571565b9181601f8401121561012f5782359167ffffffffffffffff831161012f576020838186019501011161012f57565b90608060031983011261012f5767ffffffffffffffff60043581811161012f5783611642916004016115e7565b9390939260243583811161012f578261165d916004016115e7565b9390939260443591821161012f57611677916004016115e7565b909160643590565b600090808252602090828252604092838120338252835260ff8482205416156116a85750505050565b83519167ffffffffffffffff903360608501838111868210176119a0578752602a8552858501918736843785511561198c5760308353855191600192831015611978576078602188015360295b83811161190e57506118cc57908751936080850190858210908211176118b8578852604284528684019460603687378451156118a4576030865384518210156118a45790607860218601536041915b818311611836575050506117f4576117f09386936117d4936117c560489461179c9a519a8b957f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008c8801525180926037880190611426565b8401917001034b99036b4b9b9b4b733903937b6329607d1b603784015251809386840190611426565b010360288101875201856113b1565b5192839262461bcd60e51b845260048401526024830190611449565b0390fd5b60648587519062461bcd60e51b825280600483015260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152fd5b909192600f81166010811015611890576f181899199a1a9b1b9c1cb0b131b232b360811b901a61186685886119b4565b5360041c92801561187c57600019019190611744565b634e487b7160e01b82526011600452602482fd5b634e487b7160e01b83526032600452602483fd5b634e487b7160e01b81526032600452602490fd5b634e487b7160e01b86526041600452602486fd5b60648789519062461bcd60e51b825280600483015260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152fd5b90600f81166010811015611964576f181899199a1a9b1b9c1cb0b131b232b360811b901a61193c838a6119b4565b5360041c90801561195057600019016116f5565b634e487b7160e01b87526011600452602487fd5b634e487b7160e01b88526032600452602488fd5b634e487b7160e01b86526032600452602486fd5b634e487b7160e01b85526032600452602485fd5b634e487b7160e01b85526041600452602485fd5b9081518110156114f1570160200190565b908060209392818452848401376000828201840152601f01601f1916010190565b9796959192611a2494611a08611a16936060989660808d5260808d01916119c5565b918a830360208c01526119c5565b9187830360408901526119c5565b930152565b6000198114610bc95760010190565b81601f8201121561012f578051611a4e816113d3565b92611a5c60405194856113b1565b8184526020828401011161012f576114d6916020808501910161142656fea2646970667358221220d808cceca151c0da8ce5de1f6eb6a7a67e33127a66f6bdff0c8303fc2de3672864736f6c63430008140033",
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AIModels *AIModelsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AIModels *AIModelsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AIModels.Contract.DEFAULTADMINROLE(&_AIModels.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AIModels *AIModelsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AIModels.Contract.DEFAULTADMINROLE(&_AIModels.CallOpts)
}

// DecodeModelInfo is a free data retrieval call binding the contract method 0x0d8bc096.
//
// Solidity: function decodeModelInfo(bytes modelInfo) pure returns(string modelName, string modelVersion, string modelExtendInfo, uint256 price)
func (_AIModels *AIModelsCaller) DecodeModelInfo(opts *bind.CallOpts, modelInfo []byte) (struct {
	ModelName       string
	ModelVersion    string
	ModelExtendInfo string
	Price           *big.Int
}, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "decodeModelInfo", modelInfo)

	outstruct := new(struct {
		ModelName       string
		ModelVersion    string
		ModelExtendInfo string
		Price           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ModelName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ModelVersion = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ModelExtendInfo = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DecodeModelInfo is a free data retrieval call binding the contract method 0x0d8bc096.
//
// Solidity: function decodeModelInfo(bytes modelInfo) pure returns(string modelName, string modelVersion, string modelExtendInfo, uint256 price)
func (_AIModels *AIModelsSession) DecodeModelInfo(modelInfo []byte) (struct {
	ModelName       string
	ModelVersion    string
	ModelExtendInfo string
	Price           *big.Int
}, error) {
	return _AIModels.Contract.DecodeModelInfo(&_AIModels.CallOpts, modelInfo)
}

// DecodeModelInfo is a free data retrieval call binding the contract method 0x0d8bc096.
//
// Solidity: function decodeModelInfo(bytes modelInfo) pure returns(string modelName, string modelVersion, string modelExtendInfo, uint256 price)
func (_AIModels *AIModelsCallerSession) DecodeModelInfo(modelInfo []byte) (struct {
	ModelName       string
	ModelVersion    string
	ModelExtendInfo string
	Price           *big.Int
}, error) {
	return _AIModels.Contract.DecodeModelInfo(&_AIModels.CallOpts, modelInfo)
}

// EncodeModelInfo is a free data retrieval call binding the contract method 0xc8dad360.
//
// Solidity: function encodeModelInfo(string modelName, string modelVersion, string modelExtendInfo, uint256 price) pure returns(bytes)
func (_AIModels *AIModelsCaller) EncodeModelInfo(opts *bind.CallOpts, modelName string, modelVersion string, modelExtendInfo string, price *big.Int) ([]byte, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "encodeModelInfo", modelName, modelVersion, modelExtendInfo, price)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeModelInfo is a free data retrieval call binding the contract method 0xc8dad360.
//
// Solidity: function encodeModelInfo(string modelName, string modelVersion, string modelExtendInfo, uint256 price) pure returns(bytes)
func (_AIModels *AIModelsSession) EncodeModelInfo(modelName string, modelVersion string, modelExtendInfo string, price *big.Int) ([]byte, error) {
	return _AIModels.Contract.EncodeModelInfo(&_AIModels.CallOpts, modelName, modelVersion, modelExtendInfo, price)
}

// EncodeModelInfo is a free data retrieval call binding the contract method 0xc8dad360.
//
// Solidity: function encodeModelInfo(string modelName, string modelVersion, string modelExtendInfo, uint256 price) pure returns(bytes)
func (_AIModels *AIModelsCallerSession) EncodeModelInfo(modelName string, modelVersion string, modelExtendInfo string, price *big.Int) ([]byte, error) {
	return _AIModels.Contract.EncodeModelInfo(&_AIModels.CallOpts, modelName, modelVersion, modelExtendInfo, price)
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

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AIModels *AIModelsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AIModels *AIModelsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AIModels.Contract.GetRoleAdmin(&_AIModels.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AIModels *AIModelsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AIModels.Contract.GetRoleAdmin(&_AIModels.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AIModels *AIModelsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AIModels *AIModelsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AIModels.Contract.HasRole(&_AIModels.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AIModels *AIModelsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AIModels.Contract.HasRole(&_AIModels.CallOpts, role, account)
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

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_AIModels *AIModelsCaller) RenounceRole(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) error {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "renounceRole", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_AIModels *AIModelsSession) RenounceRole(arg0 [32]byte, arg1 common.Address) error {
	return _AIModels.Contract.RenounceRole(&_AIModels.CallOpts, arg0, arg1)
}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_AIModels *AIModelsCallerSession) RenounceRole(arg0 [32]byte, arg1 common.Address) error {
	return _AIModels.Contract.RenounceRole(&_AIModels.CallOpts, arg0, arg1)
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

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AIModels *AIModelsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AIModels.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AIModels *AIModelsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AIModels.Contract.SupportsInterface(&_AIModels.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AIModels *AIModelsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AIModels.Contract.SupportsInterface(&_AIModels.CallOpts, interfaceId)
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

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AIModels *AIModelsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AIModels.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AIModels *AIModelsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AIModels.Contract.GrantRole(&_AIModels.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AIModels *AIModelsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AIModels.Contract.GrantRole(&_AIModels.TransactOpts, role, account)
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

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AIModels *AIModelsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AIModels.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AIModels *AIModelsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AIModels.Contract.RevokeRole(&_AIModels.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AIModels *AIModelsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AIModels.Contract.RevokeRole(&_AIModels.TransactOpts, role, account)
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

// AIModelsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AIModels contract.
type AIModelsRoleAdminChangedIterator struct {
	Event *AIModelsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AIModelsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIModelsRoleAdminChanged)
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
		it.Event = new(AIModelsRoleAdminChanged)
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
func (it *AIModelsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIModelsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIModelsRoleAdminChanged represents a RoleAdminChanged event raised by the AIModels contract.
type AIModelsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AIModels *AIModelsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AIModelsRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AIModels.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AIModelsRoleAdminChangedIterator{contract: _AIModels.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AIModels *AIModelsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AIModelsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AIModels.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIModelsRoleAdminChanged)
				if err := _AIModels.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AIModels *AIModelsFilterer) ParseRoleAdminChanged(log types.Log) (*AIModelsRoleAdminChanged, error) {
	event := new(AIModelsRoleAdminChanged)
	if err := _AIModels.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIModelsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AIModels contract.
type AIModelsRoleGrantedIterator struct {
	Event *AIModelsRoleGranted // Event containing the contract specifics and raw log

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
func (it *AIModelsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIModelsRoleGranted)
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
		it.Event = new(AIModelsRoleGranted)
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
func (it *AIModelsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIModelsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIModelsRoleGranted represents a RoleGranted event raised by the AIModels contract.
type AIModelsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AIModels *AIModelsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AIModelsRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AIModels.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AIModelsRoleGrantedIterator{contract: _AIModels.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AIModels *AIModelsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AIModelsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AIModels.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIModelsRoleGranted)
				if err := _AIModels.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AIModels *AIModelsFilterer) ParseRoleGranted(log types.Log) (*AIModelsRoleGranted, error) {
	event := new(AIModelsRoleGranted)
	if err := _AIModels.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIModelsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AIModels contract.
type AIModelsRoleRevokedIterator struct {
	Event *AIModelsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AIModelsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIModelsRoleRevoked)
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
		it.Event = new(AIModelsRoleRevoked)
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
func (it *AIModelsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIModelsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIModelsRoleRevoked represents a RoleRevoked event raised by the AIModels contract.
type AIModelsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AIModels *AIModelsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AIModelsRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AIModels.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AIModelsRoleRevokedIterator{contract: _AIModels.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AIModels *AIModelsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AIModelsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AIModels.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIModelsRoleRevoked)
				if err := _AIModels.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AIModels *AIModelsFilterer) ParseRoleRevoked(log types.Log) (*AIModelsRoleRevoked, error) {
	event := new(AIModelsRoleRevoked)
	if err := _AIModels.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
