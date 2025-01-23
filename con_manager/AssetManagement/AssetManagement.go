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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_topWallet\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_whiteContracts\",\"type\":\"address[]\"}],\"name\":\"AssetManagement_initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockPeriod\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteContracts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506109ea806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630c51b88f1461006757806323fe0ea91461007c57806382dda22d1461008f5780638cf8ab5a146100fc578063f2888dbb1461012f578063fc0c546a14610142575b600080fd5b61007a610075366004610919565b610173565b005b61007a61008a366004610898565b6103b5565b6100e261009d366004610866565b6001600160a01b039182166000908152600260209081526040808320939094168252918252829020825180840190935280548084526001909101549290910182905291565b604080519283526020830191909152015b60405180910390f35b61011f61010a366004610845565b60016020526000908152604090205460ff1681565b60405190151581526020016100f3565b61007a61013d366004610845565b6105d9565b60005461015b906201000090046001600160a01b031681565b6040516001600160a01b0390911681526020016100f3565b6001600160a01b03831660009081526001602052604090205460ff166101db5760405162461bcd60e51b815260206004820152601860248201527710dbdb9d1c9858dd081b9bdd081dda1a5d195b1a5cdd195960421b60448201526064015b60405180910390fd5b811561030457600082116102315760405162461bcd60e51b815260206004820152601d60248201527f416d6f756e74206d7573742062652067726561746572207468616e203000000060448201526064016101d2565b6000546040516323b872dd60e01b815233600482015230602482015260448101849052620100009091046001600160a01b0316906323b872dd90606401602060405180830381600087803b15801561028857600080fd5b505af115801561029c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c0919061094b565b6103045760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b60448201526064016101d2565b6000610310824261096b565b6001600160a01b038516600090815260026020908152604080832033845290915281208054929350859290919061034890849061096b565b90915550506001600160a01b0384166000818152600260209081526040808320338085529083529281902060010185905580518781529182018590527f6c86f3fd5118b3aa8bb4f389a617046de0a3d3d477de1a1673d227f802f616dc910160405180910390a350505050565b600054610100900460ff166103d05760005460ff16156103d8565b6103d861080a565b61043b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016101d2565b600054610100900460ff1615801561045d576000805461ffff19166101011790555b6001600160a01b0384166104b35760405162461bcd60e51b815260206004820152601960248201527f496e76616c696420746f7020746f6b656e20616464726573730000000000000060448201526064016101d2565b60005b828110156105c05760008484838181106104e057634e487b7160e01b600052603260045260246000fd5b90506020020160208101906104f59190610845565b6001600160a01b0316141561054c5760405162461bcd60e51b815260206004820152601e60248201527f496e76616c696420776869746520636f6e74726163742061646472657373000060448201526064016101d2565b600180600086868581811061057157634e487b7160e01b600052603260045260246000fd5b90506020020160208101906105869190610845565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055806105b881610983565b9150506104b6565b5080156105d3576000805461ff00191690555b50505050565b6001600160a01b03811660009081526001602052604090205460ff1661063c5760405162461bcd60e51b815260206004820152601860248201527710dbdb9d1c9858dd081b9bdd081dda1a5d195b1a5cdd195960421b60448201526064016101d2565b3360009081526002602090815260408083206001600160a01b0385168452909152902080546106a05760405162461bcd60e51b815260206004820152601060248201526f139bc81cdd185ad95908185b5bdd5b9d60821b60448201526064016101d2565b80600101544210156106ec5760405162461bcd60e51b815260206004820152601560248201527414dd185ad9481a5cc81cdd1a5b1b081b1bd8dad959605a1b60448201526064016101d2565b80546000808355600183018190555460405163a9059cbb60e01b815233600482015260248101839052620100009091046001600160a01b03169063a9059cbb90604401602060405180830381600087803b15801561074957600080fd5b505af115801561075d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610781919061094b565b6107c55760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b60448201526064016101d2565b6040518181526001600160a01b0384169033907fd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e39060200160405180910390a3505050565b60006108153061081b565b15905090565b6001600160a01b0381163b15155b919050565b80356001600160a01b038116811461082957600080fd5b600060208284031215610856578081fd5b61085f8261082e565b9392505050565b60008060408385031215610878578081fd5b6108818361082e565b915061088f6020840161082e565b90509250929050565b6000806000604084860312156108ac578081fd5b6108b58461082e565b9250602084013567ffffffffffffffff808211156108d1578283fd5b818601915086601f8301126108e4578283fd5b8135818111156108f2578384fd5b8760208260051b8501011115610906578384fd5b6020830194508093505050509250925092565b60008060006060848603121561092d578283fd5b6109368461082e565b95602085013595506040909401359392505050565b60006020828403121561095c578081fd5b8151801515811461085f578182fd5b6000821982111561097e5761097e61099e565b500190565b60006000198214156109975761099761099e565b5060010190565b634e487b7160e01b600052601160045260246000fdfea26469706673582212205dd9f2296546b8f76e6f3fc64638153c51bf7a50f598f1e368c46d79155f307b64736f6c63430008030033",
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
