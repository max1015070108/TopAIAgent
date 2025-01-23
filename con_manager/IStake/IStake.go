// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IStake

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

// IStakeMetaData contains all meta data concerning the IStake contract.
var IStakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockPeriod\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IStakeABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakeMetaData.ABI instead.
var IStakeABI = IStakeMetaData.ABI

// IStake is an auto generated Go binding around an Ethereum contract.
type IStake struct {
	IStakeCaller     // Read-only binding to the contract
	IStakeTransactor // Write-only binding to the contract
	IStakeFilterer   // Log filterer for contract events
}

// IStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakeSession struct {
	Contract     *IStake           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakeCallerSession struct {
	Contract *IStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakeTransactorSession struct {
	Contract     *IStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakeRaw struct {
	Contract *IStake // Generic contract binding to access the raw methods on
}

// IStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakeCallerRaw struct {
	Contract *IStakeCaller // Generic read-only contract binding to access the raw methods on
}

// IStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakeTransactorRaw struct {
	Contract *IStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStake creates a new instance of IStake, bound to a specific deployed contract.
func NewIStake(address common.Address, backend bind.ContractBackend) (*IStake, error) {
	contract, err := bindIStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStake{IStakeCaller: IStakeCaller{contract: contract}, IStakeTransactor: IStakeTransactor{contract: contract}, IStakeFilterer: IStakeFilterer{contract: contract}}, nil
}

// NewIStakeCaller creates a new read-only instance of IStake, bound to a specific deployed contract.
func NewIStakeCaller(address common.Address, caller bind.ContractCaller) (*IStakeCaller, error) {
	contract, err := bindIStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakeCaller{contract: contract}, nil
}

// NewIStakeTransactor creates a new write-only instance of IStake, bound to a specific deployed contract.
func NewIStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakeTransactor, error) {
	contract, err := bindIStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakeTransactor{contract: contract}, nil
}

// NewIStakeFilterer creates a new log filterer instance of IStake, bound to a specific deployed contract.
func NewIStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakeFilterer, error) {
	contract, err := bindIStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakeFilterer{contract: contract}, nil
}

// bindIStake binds a generic wrapper to an already deployed contract.
func bindIStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStake *IStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStake.Contract.IStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStake *IStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStake.Contract.IStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStake *IStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStake.Contract.IStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStake *IStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStake *IStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStake *IStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStake.Contract.contract.Transact(opts, method, params...)
}

// GetStake is a free data retrieval call binding the contract method 0x82dda22d.
//
// Solidity: function getStake(address contractAddress, address user) view returns(uint256 amount, uint256 unlockTime)
func (_IStake *IStakeCaller) GetStake(opts *bind.CallOpts, contractAddress common.Address, user common.Address) (struct {
	Amount     *big.Int
	UnlockTime *big.Int
}, error) {
	var out []interface{}
	err := _IStake.contract.Call(opts, &out, "getStake", contractAddress, user)

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
func (_IStake *IStakeSession) GetStake(contractAddress common.Address, user common.Address) (struct {
	Amount     *big.Int
	UnlockTime *big.Int
}, error) {
	return _IStake.Contract.GetStake(&_IStake.CallOpts, contractAddress, user)
}

// GetStake is a free data retrieval call binding the contract method 0x82dda22d.
//
// Solidity: function getStake(address contractAddress, address user) view returns(uint256 amount, uint256 unlockTime)
func (_IStake *IStakeCallerSession) GetStake(contractAddress common.Address, user common.Address) (struct {
	Amount     *big.Int
	UnlockTime *big.Int
}, error) {
	return _IStake.Contract.GetStake(&_IStake.CallOpts, contractAddress, user)
}

// Stake is a paid mutator transaction binding the contract method 0x0c51b88f.
//
// Solidity: function stake(address contractAddress, uint256 amount, uint256 lockPeriod) returns()
func (_IStake *IStakeTransactor) Stake(opts *bind.TransactOpts, contractAddress common.Address, amount *big.Int, lockPeriod *big.Int) (*types.Transaction, error) {
	return _IStake.contract.Transact(opts, "stake", contractAddress, amount, lockPeriod)
}

// Stake is a paid mutator transaction binding the contract method 0x0c51b88f.
//
// Solidity: function stake(address contractAddress, uint256 amount, uint256 lockPeriod) returns()
func (_IStake *IStakeSession) Stake(contractAddress common.Address, amount *big.Int, lockPeriod *big.Int) (*types.Transaction, error) {
	return _IStake.Contract.Stake(&_IStake.TransactOpts, contractAddress, amount, lockPeriod)
}

// Stake is a paid mutator transaction binding the contract method 0x0c51b88f.
//
// Solidity: function stake(address contractAddress, uint256 amount, uint256 lockPeriod) returns()
func (_IStake *IStakeTransactorSession) Stake(contractAddress common.Address, amount *big.Int, lockPeriod *big.Int) (*types.Transaction, error) {
	return _IStake.Contract.Stake(&_IStake.TransactOpts, contractAddress, amount, lockPeriod)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address contractAddress) returns()
func (_IStake *IStakeTransactor) Unstake(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _IStake.contract.Transact(opts, "unstake", contractAddress)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address contractAddress) returns()
func (_IStake *IStakeSession) Unstake(contractAddress common.Address) (*types.Transaction, error) {
	return _IStake.Contract.Unstake(&_IStake.TransactOpts, contractAddress)
}

// Unstake is a paid mutator transaction binding the contract method 0xf2888dbb.
//
// Solidity: function unstake(address contractAddress) returns()
func (_IStake *IStakeTransactorSession) Unstake(contractAddress common.Address) (*types.Transaction, error) {
	return _IStake.Contract.Unstake(&_IStake.TransactOpts, contractAddress)
}
