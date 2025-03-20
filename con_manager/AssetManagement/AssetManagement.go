// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AssetManagement

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

// AssetManagementMetaData contains all meta data concerning the AssetManagement contract.
var AssetManagementMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_topWallet\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_whiteContracts\",\"type\":\"address[]\"}],\"name\":\"AssetManagement_initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockPeriod\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteContracts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657610889908161001c8239f35b600080fdfe6040608081526004908136101561001557600080fd5b600091823560e01c9081630c51b88f146104b757816323fe0ea9146102f857816382dda22d146102605781638cf8ab5a14610222578163f2888dbb14610093575063fc0c546a1461006557600080fd5b3461008f578160031936011261008f579054905160109190911c6001600160a01b03168152602090f35b5080fd5b9190503461021e5760208060031936011261021a576001600160a01b0391826100ba610658565b1693848652600183526100d260ff8388205416610750565b3386526002835281862085875283528186209384549485156101e4576001810190815442106101a957859289809593816044945555845460101c16918551948593849263a9059cbb60e01b845233908401528960248401525af190811561019f577fd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3939291610168918891610172575b506107ec565b519283523392a380f35b6101929150843d8611610198575b61018a818361079c565b8101906107d4565b38610162565b503d610180565b82513d88823e3d90fd5b845162461bcd60e51b8152808501879052601560248201527414dd185ad9481a5cc81cdd1a5b1b081b1bd8dad959605a1b6044820152606490fd5b835162461bcd60e51b8152808401869052601060248201526f139bc81cdd185ad95908185b5bdd5b9d60821b6044820152606490fd5b8380fd5b8280fd5b50503461008f57602036600319011261008f5760209160ff9082906001600160a01b0361024d610658565b1681526001855220541690519015158152f35b8284346102f557816003193601126102f55761027a610658565b6001600160a01b03602435818116929083900361021a571682526002602052828220908252602052818120908251908382019082821067ffffffffffffffff8311176102e2575090602091845260018354938483520154918291015282519182526020820152f35b634e487b7160e01b815260418652602490fd5b80fd5b90503461021e578160031936011261021e57610312610658565b9060249081359067ffffffffffffffff938483116104b357366023840112156104b357828201359485116104b357838301928436918760051b0101116104b35761035a610710565b94856104a1575b6001600160a01b039182161561045f57875b8181106103bb57888888610385575080f35b60207f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989161ff001984541684555160018152a180f35b826103cf6103ca838589610673565b610699565b161561041d57826103e46103ca838589610673565b168952600180602052888a209060ff19825416179055600019811461040b57600101610373565b634e487b7160e01b8952601184528589fd5b875162461bcd60e51b8152602081860152601e818801527f496e76616c696420776869746520636f6e7472616374206164647265737300006044820152606490fd5b865162461bcd60e51b81526020818501526019818701527f496e76616c696420746f7020746f6b656e2061646472657373000000000000006044820152606490fd5b875461ff001916610100178855610361565b8680fd5b9190503461021e57606036600319011261021e576104d3610658565b906024359160018060a01b0380911693848652602091600183526104fc60ff8589205416610750565b8415801561057c575b505050907f6c86f3fd5118b3aa8bb4f389a617046de0a3d3d477de1a1673d227f802f616dc9161053760443542610830565b90858752600281528287203388528152828720610555868254610830565b9055858752600281528287203388528152816001848920015582519485528401523392a380f35b61061657869160648492845460101c1691865194859384926323b872dd60e01b845233908401523060248401528960448401525af190811561060c577f6c86f3fd5118b3aa8bb4f389a617046de0a3d3d477de1a1673d227f802f616dc9392916105ec9188916105f557506107ec565b90913880610505565b6101929150833d85116101985761018a818361079c565b83513d88823e3d90fd5b50915162461bcd60e51b815291820152601d60248201527f416d6f756e74206d7573742062652067726561746572207468616e20300000006044820152606490fd5b600435906001600160a01b038216820361066e57565b600080fd5b91908110156106835760051b0190565b634e487b7160e01b600052603260045260246000fd5b356001600160a01b038116810361066e5790565b156106b457565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b60005460ff8160081c16600014610732575061072d303b156106ad565b600090565b80610743600160ff819416106106ad565b60ff191617600055600190565b1561075757565b60405162461bcd60e51b815260206004820152601860248201527f436f6e7472616374206e6f742077686974656c697374656400000000000000006044820152606490fd5b90601f8019910116810190811067ffffffffffffffff8211176107be57604052565b634e487b7160e01b600052604160045260246000fd5b9081602091031261066e5751801515810361066e5790565b156107f357565b60405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b6044820152606490fd5b9190820180921161083d57565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220ffc4c75f22829371802235fa360c4feceb511c4c07a3053d19b04483e6aa414064736f6c63430008140033",
}

// AssetManagementABI is the input ABI used to generate the binding from.
// Deprecated: Use AssetManagementMetaData.ABI instead.
var AssetManagementABI = AssetManagementMetaData.ABI

// AssetManagementBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AssetManagementMetaData.Bin instead.
var AssetManagementBin = AssetManagementMetaData.Bin

// DeployAssetManagement deploys a new Ethereum contract, binding an instance of AssetManagement to it.
func DeployAssetManagement(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AssetManagement, error) {
	parsed, err := AssetManagementMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AssetManagementBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AssetManagement{AssetManagementCaller: AssetManagementCaller{contract: contract}, AssetManagementTransactor: AssetManagementTransactor{contract: contract}, AssetManagementFilterer: AssetManagementFilterer{contract: contract}}, nil
}

// AssetManagement is an auto generated Go binding around an Ethereum contract.
type AssetManagement struct {
	AssetManagementCaller     // Read-only binding to the contract
	AssetManagementTransactor // Write-only binding to the contract
	AssetManagementFilterer   // Log filterer for contract events
}

// AssetManagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetManagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetManagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetManagementSession struct {
	Contract     *AssetManagement  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetManagementCallerSession struct {
	Contract *AssetManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AssetManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetManagementTransactorSession struct {
	Contract     *AssetManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AssetManagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetManagementRaw struct {
	Contract *AssetManagement // Generic contract binding to access the raw methods on
}

// AssetManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetManagementCallerRaw struct {
	Contract *AssetManagementCaller // Generic read-only contract binding to access the raw methods on
}

// AssetManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetManagementTransactorRaw struct {
	Contract *AssetManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetManagement creates a new instance of AssetManagement, bound to a specific deployed contract.
func NewAssetManagement(address common.Address, backend bind.ContractBackend) (*AssetManagement, error) {
	contract, err := bindAssetManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetManagement{AssetManagementCaller: AssetManagementCaller{contract: contract}, AssetManagementTransactor: AssetManagementTransactor{contract: contract}, AssetManagementFilterer: AssetManagementFilterer{contract: contract}}, nil
}

// NewAssetManagementCaller creates a new read-only instance of AssetManagement, bound to a specific deployed contract.
func NewAssetManagementCaller(address common.Address, caller bind.ContractCaller) (*AssetManagementCaller, error) {
	contract, err := bindAssetManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetManagementCaller{contract: contract}, nil
}

// NewAssetManagementTransactor creates a new write-only instance of AssetManagement, bound to a specific deployed contract.
func NewAssetManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetManagementTransactor, error) {
	contract, err := bindAssetManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetManagementTransactor{contract: contract}, nil
}

// NewAssetManagementFilterer creates a new log filterer instance of AssetManagement, bound to a specific deployed contract.
func NewAssetManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetManagementFilterer, error) {
	contract, err := bindAssetManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetManagementFilterer{contract: contract}, nil
}

// bindAssetManagement binds a generic wrapper to an already deployed contract.
func bindAssetManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssetManagementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetManagement *AssetManagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetManagement.Contract.AssetManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetManagement *AssetManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetManagement.Contract.AssetManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetManagement *AssetManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetManagement.Contract.AssetManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetManagement *AssetManagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetManagement *AssetManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetManagement *AssetManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetManagement.Contract.contract.Transact(opts, method, params...)
}

// GetStake is a free data retrieval call binding the contract method 0x82dda22d.
//
// Solidity: function getStake(address contractAddress, address user) view returns(uint256 amount, uint256 unlockTime)
func (_AssetManagement *AssetManagementCaller) GetStake(opts *bind.CallOpts, contractAddress common.Address, user common.Address) (struct {
	Amount     *big.Int
	UnlockTime *big.Int
}, error) {
	var out []interface{}
	err := _AssetManagement.contract.Call(opts, &out, "getStake", contractAddress, user)

	outstruct := new(struct {
		Amount     *big.Int
		UnlockTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStake is a free data retrieval call binding the contract method 0x82dda22d.
//
// Solidity: function getStake(address contractAddress, address user) view returns(uint256 amount, uint256 unlockTime)
func (_AssetManagement *AssetManagementSession) GetStake(contractAddress common.Address, user common.Address) (struct {
	Amount     *big.Int
	UnlockTime *big.Int
}, error) {
	return _AssetManagement.Contract.GetStake(&_AssetManagement.CallOpts, contractAddress, user)
}

// GetStake is a free data retrieval call binding the contract method 0x82dda22d.
//
// Solidity: function getStake(address contractAddress, address user) view returns(uint256 amount, uint256 unlockTime)
func (_AssetManagement *AssetManagementCallerSession) GetStake(contractAddress common.Address, user common.Address) (struct {
	Amount     *big.Int
	UnlockTime *big.Int
}, error) {
	return _AssetManagement.Contract.GetStake(&_AssetManagement.CallOpts, contractAddress, user)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_AssetManagement *AssetManagementCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetManagement.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_AssetManagement *AssetManagementSession) Token() (common.Address, error) {
	return _AssetManagement.Contract.Token(&_AssetManagement.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_AssetManagement *AssetManagementCallerSession) Token() (common.Address, error) {
	return _AssetManagement.Contract.Token(&_AssetManagement.CallOpts)
}

// WhiteContracts is a free data retrieval call binding the contract method 0x8cf8ab5a.
//
// Solidity: function whiteContracts(address ) view returns(bool)
func (_AssetManagement *AssetManagementCaller) WhiteContracts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AssetManagement.contract.Call(opts, &out, "whiteContracts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteContracts is a free data retrieval call binding the contract method 0x8cf8ab5a.
//
// Solidity: function whiteContracts(address ) view returns(bool)
func (_AssetManagement *AssetManagementSession) WhiteContracts(arg0 common.Address) (bool, error) {
	return _AssetManagement.Contract.WhiteContracts(&_AssetManagement.CallOpts, arg0)
}

// WhiteContracts is a free data retrieval call binding the contract method 0x8cf8ab5a.
//
// Solidity: function whiteContracts(address ) view returns(bool)
func (_AssetManagement *AssetManagementCallerSession) WhiteContracts(arg0 common.Address) (bool, error) {
	return _AssetManagement.Contract.WhiteContracts(&_AssetManagement.CallOpts, arg0)
}

// AssetManagementInitialize is a paid mutator transaction binding the contract method 0x23fe0ea9.
//
// Solidity: function AssetManagement_initialize(address _topWallet, address[] _whiteContracts) returns()
func (_AssetManagement *AssetManagementTransactor) AssetManagementInitialize(opts *bind.TransactOpts, _topWallet common.Address, _whiteContracts []common.Address) (*types.Transaction, error) {
	return _AssetManagement.contract.Transact(opts, "AssetManagement_initialize", _topWallet, _whiteContracts)
}

// AssetManagementInitialize is a paid mutator transaction binding the contract method 0x23fe0ea9.
//
// Solidity: function AssetManagement_initialize(address _topWallet, address[] _whiteContracts) returns()
func (_AssetManagement *AssetManagementSession) AssetManagementInitialize(_topWallet common.Address, _whiteContracts []common.Address) (*types.Transaction, error) {
	return _AssetManagement.Contract.AssetManagementInitialize(&_AssetManagement.TransactOpts, _topWallet, _whiteContracts)
}

// AssetManagementInitialize is a paid mutator transaction binding the contract method 0x23fe0ea9.
//
// Solidity: function AssetManagement_initialize(address _topWallet, address[] _whiteContracts) returns()
func (_AssetManagement *AssetManagementTransactorSession) AssetManagementInitialize(_topWallet common.Address, _whiteContracts []common.Address) (*types.Transaction, error) {
	return _AssetManagement.Contract.AssetManagementInitialize(&_AssetManagement.TransactOpts, _topWallet, _whiteContracts)
}

// Stake is a paid mutator transaction binding the contract method 0x0c51b88f.
//
// Solidity: function stake(address contractAddress, uint256 amount, uint256 lockPeriod) returns()
func (_AssetManagement *AssetManagementTransactor) Stake(opts *bind.TransactOpts, contractAddress common.Address, amount *big.Int, lockPeriod *big.Int) (*types.Transaction, error) {
	return _AssetManagement.contract.Transact(opts, "stake", contractAddress, amount, lockPeriod)
}

// Stake is a paid mutator transaction binding the contract method 0x0c51b88f.
//
// Solidity: function stake(address contractAddress, uint256 amount, uint256 lockPeriod) returns()
func (_AssetManagement *AssetManagementSession) Stake(contractAddress common.Address, amount *big.Int, lockPeriod *big.Int) (*types.Transaction, error) {
	return _AssetManagement.Contract.Stake(&_AssetManagement.TransactOpts, contractAddress, amount, lockPeriod)
}

// Stake is a paid mutator transaction binding the contract method 0x0c51b88f.
//
// Solidity: function stake(address contractAddress, uint256 amount, uint256 lockPeriod) returns()
func (_AssetManagement *AssetManagementTransactorSession) Stake(contractAddress common.Address, amount *big.Int, lockPeriod *big.Int) (*types.Transaction, error) {
	return _AssetManagement.Contract.Stake(&_AssetManagement.TransactOpts, contractAddress, amount, lockPeriod)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address contractAddress) returns()
func (_AssetManagement *AssetManagementTransactor) Unstake(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _AssetManagement.contract.Transact(opts, "unstake", contractAddress)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address contractAddress) returns()
func (_AssetManagement *AssetManagementSession) Unstake(contractAddress common.Address) (*types.Transaction, error) {
	return _AssetManagement.Contract.Unstake(&_AssetManagement.TransactOpts, contractAddress)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address contractAddress) returns()
func (_AssetManagement *AssetManagementTransactorSession) Unstake(contractAddress common.Address) (*types.Transaction, error) {
	return _AssetManagement.Contract.Unstake(&_AssetManagement.TransactOpts, contractAddress)
}

// AssetManagementInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AssetManagement contract.
type AssetManagementInitializedIterator struct {
	Event *AssetManagementInitialized // Event containing the contract specifics and raw log

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
func (it *AssetManagementInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagementInitialized)
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
		it.Event = new(AssetManagementInitialized)
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
func (it *AssetManagementInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagementInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagementInitialized represents a Initialized event raised by the AssetManagement contract.
type AssetManagementInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AssetManagement *AssetManagementFilterer) FilterInitialized(opts *bind.FilterOpts) (*AssetManagementInitializedIterator, error) {

	logs, sub, err := _AssetManagement.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AssetManagementInitializedIterator{contract: _AssetManagement.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AssetManagement *AssetManagementFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AssetManagementInitialized) (event.Subscription, error) {

	logs, sub, err := _AssetManagement.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagementInitialized)
				if err := _AssetManagement.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AssetManagement *AssetManagementFilterer) ParseInitialized(log types.Log) (*AssetManagementInitialized, error) {
	event := new(AssetManagementInitialized)
	if err := _AssetManagement.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagementStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the AssetManagement contract.
type AssetManagementStakedIterator struct {
	Event *AssetManagementStaked // Event containing the contract specifics and raw log

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
func (it *AssetManagementStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagementStaked)
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
		it.Event = new(AssetManagementStaked)
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
func (it *AssetManagementStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagementStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagementStaked represents a Staked event raised by the AssetManagement contract.
type AssetManagementStaked struct {
	User            common.Address
	ContractAddress common.Address
	Amount          *big.Int
	UnlockTime      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x6c86f3fd5118b3aa8bb4f389a617046de0a3d3d477de1a1673d227f802f616dc.
//
// Solidity: event Staked(address indexed user, address indexed contractAddress, uint256 amount, uint256 unlockTime)
func (_AssetManagement *AssetManagementFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address, contractAddress []common.Address) (*AssetManagementStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _AssetManagement.contract.FilterLogs(opts, "Staked", userRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagementStakedIterator{contract: _AssetManagement.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x6c86f3fd5118b3aa8bb4f389a617046de0a3d3d477de1a1673d227f802f616dc.
//
// Solidity: event Staked(address indexed user, address indexed contractAddress, uint256 amount, uint256 unlockTime)
func (_AssetManagement *AssetManagementFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *AssetManagementStaked, user []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _AssetManagement.contract.WatchLogs(opts, "Staked", userRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagementStaked)
				if err := _AssetManagement.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x6c86f3fd5118b3aa8bb4f389a617046de0a3d3d477de1a1673d227f802f616dc.
//
// Solidity: event Staked(address indexed user, address indexed contractAddress, uint256 amount, uint256 unlockTime)
func (_AssetManagement *AssetManagementFilterer) ParseStaked(log types.Log) (*AssetManagementStaked, error) {
	event := new(AssetManagementStaked)
	if err := _AssetManagement.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagementUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the AssetManagement contract.
type AssetManagementUnstakedIterator struct {
	Event *AssetManagementUnstaked // Event containing the contract specifics and raw log

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
func (it *AssetManagementUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagementUnstaked)
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
		it.Event = new(AssetManagementUnstaked)
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
func (it *AssetManagementUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagementUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagementUnstaked represents a Unstaked event raised by the AssetManagement contract.
type AssetManagementUnstaked struct {
	User            common.Address
	ContractAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed user, address indexed contractAddress, uint256 amount)
func (_AssetManagement *AssetManagementFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address, contractAddress []common.Address) (*AssetManagementUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _AssetManagement.contract.FilterLogs(opts, "Unstaked", userRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagementUnstakedIterator{contract: _AssetManagement.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed user, address indexed contractAddress, uint256 amount)
func (_AssetManagement *AssetManagementFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *AssetManagementUnstaked, user []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _AssetManagement.contract.WatchLogs(opts, "Unstaked", userRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagementUnstaked)
				if err := _AssetManagement.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed user, address indexed contractAddress, uint256 amount)
func (_AssetManagement *AssetManagementFilterer) ParseUnstaked(log types.Log) (*AssetManagementUnstaked, error) {
	event := new(AssetManagementUnstaked)
	if err := _AssetManagement.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
