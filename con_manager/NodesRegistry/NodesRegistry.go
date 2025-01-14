// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NodesRegistry

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

// NodeComputeUsed is an auto generated low-level Go binding around an user-defined struct.
type NodeComputeUsed struct {
	Identifier common.Address
	GpuType    string
	Used       *big.Int
}

// NodesRegistryComputeAvailable is an auto generated low-level Go binding around an user-defined struct.
type NodesRegistryComputeAvailable struct {
	GpuType  string
	TotalNum *big.Int
	Used     *big.Int
}

// NodesRegistryNode is an auto generated low-level Go binding around an user-defined struct.
type NodesRegistryNode struct {
	Identifier       common.Address
	AliasIdentifier  string
	RegistrationTime *big.Int
	Active           bool
	Gpus             []NodesRegistryComputeAvailable
	Wallet           common.Address
	Stake            *big.Int
}

// NodesRegistryMetaData contains all meta data concerning the NodesRegistry contract.
var NodesRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"Authorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"}],\"name\":\"NodeActived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"}],\"name\":\"NodeDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"}],\"name\":\"NodeRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"gpuTypes\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gpuNums\",\"type\":\"uint256[]\"}],\"name\":\"allocGPU\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"internalType\":\"structNodeComputeUsed[]\",\"name\":\"gpuNodes\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizedPerson\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"at\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"registrationTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"internalType\":\"structNodesRegistry.ComputeAvailable[]\",\"name\":\"gpus\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"internalType\":\"structNodesRegistry.Node\",\"name\":\"node\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregisterNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"}],\"name\":\"deregisterNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"internalType\":\"structNodeComputeUsed[]\",\"name\":\"gpuNodes\",\"type\":\"tuple[]\"}],\"name\":\"freeGPU\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"registrationTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"internalType\":\"structNodesRegistry.ComputeAvailable[]\",\"name\":\"gpus\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"internalType\":\"structNodesRegistry.Node\",\"name\":\"node\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"gpuSummary\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"gpuTypeOfNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"gpuTypes\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gpuNums\",\"type\":\"uint256[]\"}],\"name\":\"registerNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NodesRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use NodesRegistryMetaData.ABI instead.
var NodesRegistryABI = NodesRegistryMetaData.ABI

// NodesRegistry is an auto generated Go binding around an Ethereum contract.
type NodesRegistry struct {
	NodesRegistryCaller     // Read-only binding to the contract
	NodesRegistryTransactor // Write-only binding to the contract
	NodesRegistryFilterer   // Log filterer for contract events
}

// NodesRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodesRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodesRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodesRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodesRegistrySession struct {
	Contract     *NodesRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodesRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodesRegistryCallerSession struct {
	Contract *NodesRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// NodesRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodesRegistryTransactorSession struct {
	Contract     *NodesRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// NodesRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodesRegistryRaw struct {
	Contract *NodesRegistry // Generic contract binding to access the raw methods on
}

// NodesRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodesRegistryCallerRaw struct {
	Contract *NodesRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// NodesRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodesRegistryTransactorRaw struct {
	Contract *NodesRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodesRegistry creates a new instance of NodesRegistry, bound to a specific deployed contract.
func NewNodesRegistry(address common.Address, backend bind.ContractBackend) (*NodesRegistry, error) {
	contract, err := bindNodesRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodesRegistry{NodesRegistryCaller: NodesRegistryCaller{contract: contract}, NodesRegistryTransactor: NodesRegistryTransactor{contract: contract}, NodesRegistryFilterer: NodesRegistryFilterer{contract: contract}}, nil
}

// NewNodesRegistryCaller creates a new read-only instance of NodesRegistry, bound to a specific deployed contract.
func NewNodesRegistryCaller(address common.Address, caller bind.ContractCaller) (*NodesRegistryCaller, error) {
	contract, err := bindNodesRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodesRegistryCaller{contract: contract}, nil
}

// NewNodesRegistryTransactor creates a new write-only instance of NodesRegistry, bound to a specific deployed contract.
func NewNodesRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*NodesRegistryTransactor, error) {
	contract, err := bindNodesRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodesRegistryTransactor{contract: contract}, nil
}

// NewNodesRegistryFilterer creates a new log filterer instance of NodesRegistry, bound to a specific deployed contract.
func NewNodesRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*NodesRegistryFilterer, error) {
	contract, err := bindNodesRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodesRegistryFilterer{contract: contract}, nil
}

// bindNodesRegistry binds a generic wrapper to an already deployed contract.
func bindNodesRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodesRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodesRegistry *NodesRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodesRegistry.Contract.NodesRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodesRegistry *NodesRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodesRegistry.Contract.NodesRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodesRegistry *NodesRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodesRegistry.Contract.NodesRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodesRegistry *NodesRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodesRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodesRegistry *NodesRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodesRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodesRegistry *NodesRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodesRegistry.Contract.contract.Transact(opts, method, params...)
}

// Allocator is a free data retrieval call binding the contract method 0xaa5dcecc.
//
// Solidity: function allocator() view returns(address)
func (_NodesRegistry *NodesRegistryCaller) Allocator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "allocator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Allocator is a free data retrieval call binding the contract method 0xaa5dcecc.
//
// Solidity: function allocator() view returns(address)
func (_NodesRegistry *NodesRegistrySession) Allocator() (common.Address, error) {
	return _NodesRegistry.Contract.Allocator(&_NodesRegistry.CallOpts)
}

// Allocator is a free data retrieval call binding the contract method 0xaa5dcecc.
//
// Solidity: function allocator() view returns(address)
func (_NodesRegistry *NodesRegistryCallerSession) Allocator() (common.Address, error) {
	return _NodesRegistry.Contract.Allocator(&_NodesRegistry.CallOpts)
}

// At is a free data retrieval call binding the contract method 0xe0886f90.
//
// Solidity: function at(uint256 index) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256) node)
func (_NodesRegistry *NodesRegistryCaller) At(opts *bind.CallOpts, index *big.Int) (NodesRegistryNode, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "at", index)

	if err != nil {
		return *new(NodesRegistryNode), err
	}

	out0 := *abi.ConvertType(out[0], new(NodesRegistryNode)).(*NodesRegistryNode)

	return out0, err

}

// At is a free data retrieval call binding the contract method 0xe0886f90.
//
// Solidity: function at(uint256 index) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256) node)
func (_NodesRegistry *NodesRegistrySession) At(index *big.Int) (NodesRegistryNode, error) {
	return _NodesRegistry.Contract.At(&_NodesRegistry.CallOpts, index)
}

// At is a free data retrieval call binding the contract method 0xe0886f90.
//
// Solidity: function at(uint256 index) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256) node)
func (_NodesRegistry *NodesRegistryCallerSession) At(index *big.Int) (NodesRegistryNode, error) {
	return _NodesRegistry.Contract.At(&_NodesRegistry.CallOpts, index)
}

// Check is a free data retrieval call binding the contract method 0xc23697a8.
//
// Solidity: function check(address identifier) view returns(bool)
func (_NodesRegistry *NodesRegistryCaller) Check(opts *bind.CallOpts, identifier common.Address) (bool, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "check", identifier)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Check is a free data retrieval call binding the contract method 0xc23697a8.
//
// Solidity: function check(address identifier) view returns(bool)
func (_NodesRegistry *NodesRegistrySession) Check(identifier common.Address) (bool, error) {
	return _NodesRegistry.Contract.Check(&_NodesRegistry.CallOpts, identifier)
}

// Check is a free data retrieval call binding the contract method 0xc23697a8.
//
// Solidity: function check(address identifier) view returns(bool)
func (_NodesRegistry *NodesRegistryCallerSession) Check(identifier common.Address) (bool, error) {
	return _NodesRegistry.Contract.Check(&_NodesRegistry.CallOpts, identifier)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address identifier) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256) node)
func (_NodesRegistry *NodesRegistryCaller) Get(opts *bind.CallOpts, identifier common.Address) (NodesRegistryNode, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "get", identifier)

	if err != nil {
		return *new(NodesRegistryNode), err
	}

	out0 := *abi.ConvertType(out[0], new(NodesRegistryNode)).(*NodesRegistryNode)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address identifier) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256) node)
func (_NodesRegistry *NodesRegistrySession) Get(identifier common.Address) (NodesRegistryNode, error) {
	return _NodesRegistry.Contract.Get(&_NodesRegistry.CallOpts, identifier)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address identifier) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256) node)
func (_NodesRegistry *NodesRegistryCallerSession) Get(identifier common.Address) (NodesRegistryNode, error) {
	return _NodesRegistry.Contract.Get(&_NodesRegistry.CallOpts, identifier)
}

// GpuSummary is a free data retrieval call binding the contract method 0xf67c5bdc.
//
// Solidity: function gpuSummary(string ) view returns(string gpuType, uint256 totalNum, uint256 used)
func (_NodesRegistry *NodesRegistryCaller) GpuSummary(opts *bind.CallOpts, arg0 string) (struct {
	GpuType  string
	TotalNum *big.Int
	Used     *big.Int
}, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "gpuSummary", arg0)

	outstruct := new(struct {
		GpuType  string
		TotalNum *big.Int
		Used     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GpuType = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TotalNum = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Used = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GpuSummary is a free data retrieval call binding the contract method 0xf67c5bdc.
//
// Solidity: function gpuSummary(string ) view returns(string gpuType, uint256 totalNum, uint256 used)
func (_NodesRegistry *NodesRegistrySession) GpuSummary(arg0 string) (struct {
	GpuType  string
	TotalNum *big.Int
	Used     *big.Int
}, error) {
	return _NodesRegistry.Contract.GpuSummary(&_NodesRegistry.CallOpts, arg0)
}

// GpuSummary is a free data retrieval call binding the contract method 0xf67c5bdc.
//
// Solidity: function gpuSummary(string ) view returns(string gpuType, uint256 totalNum, uint256 used)
func (_NodesRegistry *NodesRegistryCallerSession) GpuSummary(arg0 string) (struct {
	GpuType  string
	TotalNum *big.Int
	Used     *big.Int
}, error) {
	return _NodesRegistry.Contract.GpuSummary(&_NodesRegistry.CallOpts, arg0)
}

// GpuTypeOfNodes is a free data retrieval call binding the contract method 0xed38ed0d.
//
// Solidity: function gpuTypeOfNodes(address , string ) view returns(uint256)
func (_NodesRegistry *NodesRegistryCaller) GpuTypeOfNodes(opts *bind.CallOpts, arg0 common.Address, arg1 string) (*big.Int, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "gpuTypeOfNodes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GpuTypeOfNodes is a free data retrieval call binding the contract method 0xed38ed0d.
//
// Solidity: function gpuTypeOfNodes(address , string ) view returns(uint256)
func (_NodesRegistry *NodesRegistrySession) GpuTypeOfNodes(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _NodesRegistry.Contract.GpuTypeOfNodes(&_NodesRegistry.CallOpts, arg0, arg1)
}

// GpuTypeOfNodes is a free data retrieval call binding the contract method 0xed38ed0d.
//
// Solidity: function gpuTypeOfNodes(address , string ) view returns(uint256)
func (_NodesRegistry *NodesRegistryCallerSession) GpuTypeOfNodes(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _NodesRegistry.Contract.GpuTypeOfNodes(&_NodesRegistry.CallOpts, arg0, arg1)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_NodesRegistry *NodesRegistryCaller) Length(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "length")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_NodesRegistry *NodesRegistrySession) Length() (*big.Int, error) {
	return _NodesRegistry.Contract.Length(&_NodesRegistry.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_NodesRegistry *NodesRegistryCallerSession) Length() (*big.Int, error) {
	return _NodesRegistry.Contract.Length(&_NodesRegistry.CallOpts)
}

// AllocGPU is a paid mutator transaction binding the contract method 0x6252e1c2.
//
// Solidity: function allocGPU(uint256 startIndex, string[] gpuTypes, uint256[] gpuNums) returns((address,string,uint256)[] gpuNodes, uint256 len)
func (_NodesRegistry *NodesRegistryTransactor) AllocGPU(opts *bind.TransactOpts, startIndex *big.Int, gpuTypes []string, gpuNums []*big.Int) (*types.Transaction, error) {
	return _NodesRegistry.contract.Transact(opts, "allocGPU", startIndex, gpuTypes, gpuNums)
}

// AllocGPU is a paid mutator transaction binding the contract method 0x6252e1c2.
//
// Solidity: function allocGPU(uint256 startIndex, string[] gpuTypes, uint256[] gpuNums) returns((address,string,uint256)[] gpuNodes, uint256 len)
func (_NodesRegistry *NodesRegistrySession) AllocGPU(startIndex *big.Int, gpuTypes []string, gpuNums []*big.Int) (*types.Transaction, error) {
	return _NodesRegistry.Contract.AllocGPU(&_NodesRegistry.TransactOpts, startIndex, gpuTypes, gpuNums)
}

// AllocGPU is a paid mutator transaction binding the contract method 0x6252e1c2.
//
// Solidity: function allocGPU(uint256 startIndex, string[] gpuTypes, uint256[] gpuNums) returns((address,string,uint256)[] gpuNodes, uint256 len)
func (_NodesRegistry *NodesRegistryTransactorSession) AllocGPU(startIndex *big.Int, gpuTypes []string, gpuNums []*big.Int) (*types.Transaction, error) {
	return _NodesRegistry.Contract.AllocGPU(&_NodesRegistry.TransactOpts, startIndex, gpuTypes, gpuNums)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address authorizedPerson) returns(bool)
func (_NodesRegistry *NodesRegistryTransactor) Approve(opts *bind.TransactOpts, authorizedPerson common.Address) (*types.Transaction, error) {
	return _NodesRegistry.contract.Transact(opts, "approve", authorizedPerson)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address authorizedPerson) returns(bool)
func (_NodesRegistry *NodesRegistrySession) Approve(authorizedPerson common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.Approve(&_NodesRegistry.TransactOpts, authorizedPerson)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address authorizedPerson) returns(bool)
func (_NodesRegistry *NodesRegistryTransactorSession) Approve(authorizedPerson common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.Approve(&_NodesRegistry.TransactOpts, authorizedPerson)
}

// DeregisterNode is a paid mutator transaction binding the contract method 0x18b1c081.
//
// Solidity: function deregisterNode() returns()
func (_NodesRegistry *NodesRegistryTransactor) DeregisterNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodesRegistry.contract.Transact(opts, "deregisterNode")
}

// DeregisterNode is a paid mutator transaction binding the contract method 0x18b1c081.
//
// Solidity: function deregisterNode() returns()
func (_NodesRegistry *NodesRegistrySession) DeregisterNode() (*types.Transaction, error) {
	return _NodesRegistry.Contract.DeregisterNode(&_NodesRegistry.TransactOpts)
}

// DeregisterNode is a paid mutator transaction binding the contract method 0x18b1c081.
//
// Solidity: function deregisterNode() returns()
func (_NodesRegistry *NodesRegistryTransactorSession) DeregisterNode() (*types.Transaction, error) {
	return _NodesRegistry.Contract.DeregisterNode(&_NodesRegistry.TransactOpts)
}

// DeregisterNode0 is a paid mutator transaction binding the contract method 0xe7658aaf.
//
// Solidity: function deregisterNode(address authorizer) returns()
func (_NodesRegistry *NodesRegistryTransactor) DeregisterNode0(opts *bind.TransactOpts, authorizer common.Address) (*types.Transaction, error) {
	return _NodesRegistry.contract.Transact(opts, "deregisterNode0", authorizer)
}

// DeregisterNode0 is a paid mutator transaction binding the contract method 0xe7658aaf.
//
// Solidity: function deregisterNode(address authorizer) returns()
func (_NodesRegistry *NodesRegistrySession) DeregisterNode0(authorizer common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.DeregisterNode0(&_NodesRegistry.TransactOpts, authorizer)
}

// DeregisterNode0 is a paid mutator transaction binding the contract method 0xe7658aaf.
//
// Solidity: function deregisterNode(address authorizer) returns()
func (_NodesRegistry *NodesRegistryTransactorSession) DeregisterNode0(authorizer common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.DeregisterNode0(&_NodesRegistry.TransactOpts, authorizer)
}

// FreeGPU is a paid mutator transaction binding the contract method 0x036fe9c2.
//
// Solidity: function freeGPU((address,string,uint256)[] gpuNodes) returns()
func (_NodesRegistry *NodesRegistryTransactor) FreeGPU(opts *bind.TransactOpts, gpuNodes []NodeComputeUsed) (*types.Transaction, error) {
	return _NodesRegistry.contract.Transact(opts, "freeGPU", gpuNodes)
}

// FreeGPU is a paid mutator transaction binding the contract method 0x036fe9c2.
//
// Solidity: function freeGPU((address,string,uint256)[] gpuNodes) returns()
func (_NodesRegistry *NodesRegistrySession) FreeGPU(gpuNodes []NodeComputeUsed) (*types.Transaction, error) {
	return _NodesRegistry.Contract.FreeGPU(&_NodesRegistry.TransactOpts, gpuNodes)
}

// FreeGPU is a paid mutator transaction binding the contract method 0x036fe9c2.
//
// Solidity: function freeGPU((address,string,uint256)[] gpuNodes) returns()
func (_NodesRegistry *NodesRegistryTransactorSession) FreeGPU(gpuNodes []NodeComputeUsed) (*types.Transaction, error) {
	return _NodesRegistry.Contract.FreeGPU(&_NodesRegistry.TransactOpts, gpuNodes)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xefca74d2.
//
// Solidity: function registerNode(address wallet, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums) payable returns()
func (_NodesRegistry *NodesRegistryTransactor) RegisterNode(opts *bind.TransactOpts, wallet common.Address, aliasIdentifier string, gpuTypes []string, gpuNums []*big.Int) (*types.Transaction, error) {
	return _NodesRegistry.contract.Transact(opts, "registerNode", wallet, aliasIdentifier, gpuTypes, gpuNums)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xefca74d2.
//
// Solidity: function registerNode(address wallet, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums) payable returns()
func (_NodesRegistry *NodesRegistrySession) RegisterNode(wallet common.Address, aliasIdentifier string, gpuTypes []string, gpuNums []*big.Int) (*types.Transaction, error) {
	return _NodesRegistry.Contract.RegisterNode(&_NodesRegistry.TransactOpts, wallet, aliasIdentifier, gpuTypes, gpuNums)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xefca74d2.
//
// Solidity: function registerNode(address wallet, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums) payable returns()
func (_NodesRegistry *NodesRegistryTransactorSession) RegisterNode(wallet common.Address, aliasIdentifier string, gpuTypes []string, gpuNums []*big.Int) (*types.Transaction, error) {
	return _NodesRegistry.Contract.RegisterNode(&_NodesRegistry.TransactOpts, wallet, aliasIdentifier, gpuTypes, gpuNums)
}

// NodesRegistryAuthorizedIterator is returned from FilterAuthorized and is used to iterate over the raw logs and unpacked data for Authorized events raised by the NodesRegistry contract.
type NodesRegistryAuthorizedIterator struct {
	Event *NodesRegistryAuthorized // Event containing the contract specifics and raw log

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
func (it *NodesRegistryAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRegistryAuthorized)
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
		it.Event = new(NodesRegistryAuthorized)
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
func (it *NodesRegistryAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRegistryAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRegistryAuthorized represents a Authorized event raised by the NodesRegistry contract.
type NodesRegistryAuthorized struct {
	Owner   common.Address
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuthorized is a free log retrieval operation binding the contract event 0xf5a7f4fb8a92356e8c8c4ae7ac3589908381450500a7e2fd08c95600021ee889.
//
// Solidity: event Authorized(address indexed owner, address indexed spender)
func (_NodesRegistry *NodesRegistryFilterer) FilterAuthorized(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NodesRegistryAuthorizedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NodesRegistry.contract.FilterLogs(opts, "Authorized", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NodesRegistryAuthorizedIterator{contract: _NodesRegistry.contract, event: "Authorized", logs: logs, sub: sub}, nil
}

// WatchAuthorized is a free log subscription operation binding the contract event 0xf5a7f4fb8a92356e8c8c4ae7ac3589908381450500a7e2fd08c95600021ee889.
//
// Solidity: event Authorized(address indexed owner, address indexed spender)
func (_NodesRegistry *NodesRegistryFilterer) WatchAuthorized(opts *bind.WatchOpts, sink chan<- *NodesRegistryAuthorized, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NodesRegistry.contract.WatchLogs(opts, "Authorized", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRegistryAuthorized)
				if err := _NodesRegistry.contract.UnpackLog(event, "Authorized", log); err != nil {
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

// ParseAuthorized is a log parse operation binding the contract event 0xf5a7f4fb8a92356e8c8c4ae7ac3589908381450500a7e2fd08c95600021ee889.
//
// Solidity: event Authorized(address indexed owner, address indexed spender)
func (_NodesRegistry *NodesRegistryFilterer) ParseAuthorized(log types.Log) (*NodesRegistryAuthorized, error) {
	event := new(NodesRegistryAuthorized)
	if err := _NodesRegistry.contract.UnpackLog(event, "Authorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesRegistryNodeActivedIterator is returned from FilterNodeActived and is used to iterate over the raw logs and unpacked data for NodeActived events raised by the NodesRegistry contract.
type NodesRegistryNodeActivedIterator struct {
	Event *NodesRegistryNodeActived // Event containing the contract specifics and raw log

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
func (it *NodesRegistryNodeActivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRegistryNodeActived)
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
		it.Event = new(NodesRegistryNodeActived)
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
func (it *NodesRegistryNodeActivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRegistryNodeActivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRegistryNodeActived represents a NodeActived event raised by the NodesRegistry contract.
type NodesRegistryNodeActived struct {
	Wallet          common.Address
	Identifier      common.Address
	Time            *big.Int
	AliasIdentifier string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeActived is a free log retrieval operation binding the contract event 0x25cdb56344c72b35eff48499be35553cea1cf3ee21c613692d20ea5a5c539e37.
//
// Solidity: event NodeActived(address indexed wallet, address identifier, uint256 time, string aliasIdentifier)
func (_NodesRegistry *NodesRegistryFilterer) FilterNodeActived(opts *bind.FilterOpts, wallet []common.Address) (*NodesRegistryNodeActivedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _NodesRegistry.contract.FilterLogs(opts, "NodeActived", walletRule)
	if err != nil {
		return nil, err
	}
	return &NodesRegistryNodeActivedIterator{contract: _NodesRegistry.contract, event: "NodeActived", logs: logs, sub: sub}, nil
}

// WatchNodeActived is a free log subscription operation binding the contract event 0x25cdb56344c72b35eff48499be35553cea1cf3ee21c613692d20ea5a5c539e37.
//
// Solidity: event NodeActived(address indexed wallet, address identifier, uint256 time, string aliasIdentifier)
func (_NodesRegistry *NodesRegistryFilterer) WatchNodeActived(opts *bind.WatchOpts, sink chan<- *NodesRegistryNodeActived, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _NodesRegistry.contract.WatchLogs(opts, "NodeActived", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRegistryNodeActived)
				if err := _NodesRegistry.contract.UnpackLog(event, "NodeActived", log); err != nil {
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

// ParseNodeActived is a log parse operation binding the contract event 0x25cdb56344c72b35eff48499be35553cea1cf3ee21c613692d20ea5a5c539e37.
//
// Solidity: event NodeActived(address indexed wallet, address identifier, uint256 time, string aliasIdentifier)
func (_NodesRegistry *NodesRegistryFilterer) ParseNodeActived(log types.Log) (*NodesRegistryNodeActived, error) {
	event := new(NodesRegistryNodeActived)
	if err := _NodesRegistry.contract.UnpackLog(event, "NodeActived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesRegistryNodeDeregisteredIterator is returned from FilterNodeDeregistered and is used to iterate over the raw logs and unpacked data for NodeDeregistered events raised by the NodesRegistry contract.
type NodesRegistryNodeDeregisteredIterator struct {
	Event *NodesRegistryNodeDeregistered // Event containing the contract specifics and raw log

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
func (it *NodesRegistryNodeDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRegistryNodeDeregistered)
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
		it.Event = new(NodesRegistryNodeDeregistered)
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
func (it *NodesRegistryNodeDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRegistryNodeDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRegistryNodeDeregistered represents a NodeDeregistered event raised by the NodesRegistry contract.
type NodesRegistryNodeDeregistered struct {
	Identifier      common.Address
	Time            *big.Int
	AliasIdentifier string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeDeregistered is a free log retrieval operation binding the contract event 0x60d01d146c7aa1a7d4e3fdd5543872f7d5b2a241980a66b3552ae1a86ae18453.
//
// Solidity: event NodeDeregistered(address indexed identifier, uint256 time, string aliasIdentifier)
func (_NodesRegistry *NodesRegistryFilterer) FilterNodeDeregistered(opts *bind.FilterOpts, identifier []common.Address) (*NodesRegistryNodeDeregisteredIterator, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _NodesRegistry.contract.FilterLogs(opts, "NodeDeregistered", identifierRule)
	if err != nil {
		return nil, err
	}
	return &NodesRegistryNodeDeregisteredIterator{contract: _NodesRegistry.contract, event: "NodeDeregistered", logs: logs, sub: sub}, nil
}

// WatchNodeDeregistered is a free log subscription operation binding the contract event 0x60d01d146c7aa1a7d4e3fdd5543872f7d5b2a241980a66b3552ae1a86ae18453.
//
// Solidity: event NodeDeregistered(address indexed identifier, uint256 time, string aliasIdentifier)
func (_NodesRegistry *NodesRegistryFilterer) WatchNodeDeregistered(opts *bind.WatchOpts, sink chan<- *NodesRegistryNodeDeregistered, identifier []common.Address) (event.Subscription, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _NodesRegistry.contract.WatchLogs(opts, "NodeDeregistered", identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRegistryNodeDeregistered)
				if err := _NodesRegistry.contract.UnpackLog(event, "NodeDeregistered", log); err != nil {
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

// ParseNodeDeregistered is a log parse operation binding the contract event 0x60d01d146c7aa1a7d4e3fdd5543872f7d5b2a241980a66b3552ae1a86ae18453.
//
// Solidity: event NodeDeregistered(address indexed identifier, uint256 time, string aliasIdentifier)
func (_NodesRegistry *NodesRegistryFilterer) ParseNodeDeregistered(log types.Log) (*NodesRegistryNodeDeregistered, error) {
	event := new(NodesRegistryNodeDeregistered)
	if err := _NodesRegistry.contract.UnpackLog(event, "NodeDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesRegistryNodeRegisteredIterator is returned from FilterNodeRegistered and is used to iterate over the raw logs and unpacked data for NodeRegistered events raised by the NodesRegistry contract.
type NodesRegistryNodeRegisteredIterator struct {
	Event *NodesRegistryNodeRegistered // Event containing the contract specifics and raw log

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
func (it *NodesRegistryNodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRegistryNodeRegistered)
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
		it.Event = new(NodesRegistryNodeRegistered)
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
func (it *NodesRegistryNodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRegistryNodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRegistryNodeRegistered represents a NodeRegistered event raised by the NodesRegistry contract.
type NodesRegistryNodeRegistered struct {
	Wallet          common.Address
	Identifier      common.Address
	Time            *big.Int
	AliasIdentifier string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeRegistered is a free log retrieval operation binding the contract event 0x0b464c895406310478ed7e414061d7621fd5f7444d7618e96d7fbb917a51f302.
//
// Solidity: event NodeRegistered(address indexed wallet, address identifier, uint256 time, string aliasIdentifier)
func (_NodesRegistry *NodesRegistryFilterer) FilterNodeRegistered(opts *bind.FilterOpts, wallet []common.Address) (*NodesRegistryNodeRegisteredIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _NodesRegistry.contract.FilterLogs(opts, "NodeRegistered", walletRule)
	if err != nil {
		return nil, err
	}
	return &NodesRegistryNodeRegisteredIterator{contract: _NodesRegistry.contract, event: "NodeRegistered", logs: logs, sub: sub}, nil
}

// WatchNodeRegistered is a free log subscription operation binding the contract event 0x0b464c895406310478ed7e414061d7621fd5f7444d7618e96d7fbb917a51f302.
//
// Solidity: event NodeRegistered(address indexed wallet, address identifier, uint256 time, string aliasIdentifier)
func (_NodesRegistry *NodesRegistryFilterer) WatchNodeRegistered(opts *bind.WatchOpts, sink chan<- *NodesRegistryNodeRegistered, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _NodesRegistry.contract.WatchLogs(opts, "NodeRegistered", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRegistryNodeRegistered)
				if err := _NodesRegistry.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
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

// ParseNodeRegistered is a log parse operation binding the contract event 0x0b464c895406310478ed7e414061d7621fd5f7444d7618e96d7fbb917a51f302.
//
// Solidity: event NodeRegistered(address indexed wallet, address identifier, uint256 time, string aliasIdentifier)
func (_NodesRegistry *NodesRegistryFilterer) ParseNodeRegistered(log types.Log) (*NodesRegistryNodeRegistered, error) {
	event := new(NodesRegistryNodeRegistered)
	if err := _NodesRegistry.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
