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
	IsPublic         bool
}

// NodesRegistryMetaData contains all meta data concerning the NodesRegistry contract.
var NodesRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"Authorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"gpuTypes\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"gpuNums\",\"type\":\"uint256[]\"}],\"name\":\"NodeActived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identifierOfProvider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identifierOfServer\",\"type\":\"address\"}],\"name\":\"NodeAttached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"}],\"name\":\"NodeDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identifierOfProvider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identifierOfServer\",\"type\":\"address\"}],\"name\":\"NodeDetached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"gpuTypes\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"gpuNums\",\"type\":\"uint256[]\"}],\"name\":\"NodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"ProxyNodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"gpuTypes\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gpuNums\",\"type\":\"uint256[]\"}],\"name\":\"allocGPU\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"internalType\":\"structNodeComputeUsed[]\",\"name\":\"gpuNodes\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"at\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"registrationTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"internalType\":\"structNodesRegistry.ComputeAvailable[]\",\"name\":\"gpus\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"internalType\":\"structNodesRegistry.Node\",\"name\":\"node\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"server\",\"type\":\"address\"}],\"name\":\"attach\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregisterNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"server\",\"type\":\"address\"}],\"name\":\"dettach\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"internalType\":\"structNodeComputeUsed[]\",\"name\":\"gpuNodes\",\"type\":\"tuple[]\"}],\"name\":\"freeGPU\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"registrationTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"internalType\":\"structNodesRegistry.ComputeAvailable[]\",\"name\":\"gpus\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"internalType\":\"structNodesRegistry.Node\",\"name\":\"node\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getAttach\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"gpuSummary\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"gpuTypeOfNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proxyNodes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"gpuTypes\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gpuNums\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"name\":\"registerNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"registerProxyNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIStake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_NodesRegistry *NodesRegistryCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_NodesRegistry *NodesRegistrySession) ADMINROLE() ([32]byte, error) {
	return _NodesRegistry.Contract.ADMINROLE(&_NodesRegistry.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_NodesRegistry *NodesRegistryCallerSession) ADMINROLE() ([32]byte, error) {
	return _NodesRegistry.Contract.ADMINROLE(&_NodesRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NodesRegistry *NodesRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NodesRegistry *NodesRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NodesRegistry.Contract.DEFAULTADMINROLE(&_NodesRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NodesRegistry *NodesRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NodesRegistry.Contract.DEFAULTADMINROLE(&_NodesRegistry.CallOpts)
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
// Solidity: function at(uint256 index) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256,bool) node)
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
// Solidity: function at(uint256 index) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256,bool) node)
func (_NodesRegistry *NodesRegistrySession) At(index *big.Int) (NodesRegistryNode, error) {
	return _NodesRegistry.Contract.At(&_NodesRegistry.CallOpts, index)
}

// At is a free data retrieval call binding the contract method 0xe0886f90.
//
// Solidity: function at(uint256 index) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256,bool) node)
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
// Solidity: function get(address identifier) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256,bool) node)
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
// Solidity: function get(address identifier) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256,bool) node)
func (_NodesRegistry *NodesRegistrySession) Get(identifier common.Address) (NodesRegistryNode, error) {
	return _NodesRegistry.Contract.Get(&_NodesRegistry.CallOpts, identifier)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address identifier) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256,bool) node)
func (_NodesRegistry *NodesRegistryCallerSession) Get(identifier common.Address) (NodesRegistryNode, error) {
	return _NodesRegistry.Contract.Get(&_NodesRegistry.CallOpts, identifier)
}

// GetAttach is a free data retrieval call binding the contract method 0xc7edca7a.
//
// Solidity: function getAttach(address provider) view returns(address[])
func (_NodesRegistry *NodesRegistryCaller) GetAttach(opts *bind.CallOpts, provider common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "getAttach", provider)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAttach is a free data retrieval call binding the contract method 0xc7edca7a.
//
// Solidity: function getAttach(address provider) view returns(address[])
func (_NodesRegistry *NodesRegistrySession) GetAttach(provider common.Address) ([]common.Address, error) {
	return _NodesRegistry.Contract.GetAttach(&_NodesRegistry.CallOpts, provider)
}

// GetAttach is a free data retrieval call binding the contract method 0xc7edca7a.
//
// Solidity: function getAttach(address provider) view returns(address[])
func (_NodesRegistry *NodesRegistryCallerSession) GetAttach(provider common.Address) ([]common.Address, error) {
	return _NodesRegistry.Contract.GetAttach(&_NodesRegistry.CallOpts, provider)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NodesRegistry *NodesRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NodesRegistry *NodesRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NodesRegistry.Contract.GetRoleAdmin(&_NodesRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NodesRegistry *NodesRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NodesRegistry.Contract.GetRoleAdmin(&_NodesRegistry.CallOpts, role)
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

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NodesRegistry *NodesRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NodesRegistry *NodesRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NodesRegistry.Contract.HasRole(&_NodesRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NodesRegistry *NodesRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NodesRegistry.Contract.HasRole(&_NodesRegistry.CallOpts, role, account)
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

// ProxyNodes is a free data retrieval call binding the contract method 0xb8afa39c.
//
// Solidity: function proxyNodes(address ) view returns(bool)
func (_NodesRegistry *NodesRegistryCaller) ProxyNodes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "proxyNodes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProxyNodes is a free data retrieval call binding the contract method 0xb8afa39c.
//
// Solidity: function proxyNodes(address ) view returns(bool)
func (_NodesRegistry *NodesRegistrySession) ProxyNodes(arg0 common.Address) (bool, error) {
	return _NodesRegistry.Contract.ProxyNodes(&_NodesRegistry.CallOpts, arg0)
}

// ProxyNodes is a free data retrieval call binding the contract method 0xb8afa39c.
//
// Solidity: function proxyNodes(address ) view returns(bool)
func (_NodesRegistry *NodesRegistryCallerSession) ProxyNodes(arg0 common.Address) (bool, error) {
	return _NodesRegistry.Contract.ProxyNodes(&_NodesRegistry.CallOpts, arg0)
}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_NodesRegistry *NodesRegistryCaller) RenounceRole(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) error {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "renounceRole", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_NodesRegistry *NodesRegistrySession) RenounceRole(arg0 [32]byte, arg1 common.Address) error {
	return _NodesRegistry.Contract.RenounceRole(&_NodesRegistry.CallOpts, arg0, arg1)
}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_NodesRegistry *NodesRegistryCallerSession) RenounceRole(arg0 [32]byte, arg1 common.Address) error {
	return _NodesRegistry.Contract.RenounceRole(&_NodesRegistry.CallOpts, arg0, arg1)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_NodesRegistry *NodesRegistryCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_NodesRegistry *NodesRegistrySession) StakeToken() (common.Address, error) {
	return _NodesRegistry.Contract.StakeToken(&_NodesRegistry.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_NodesRegistry *NodesRegistryCallerSession) StakeToken() (common.Address, error) {
	return _NodesRegistry.Contract.StakeToken(&_NodesRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NodesRegistry *NodesRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NodesRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NodesRegistry *NodesRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NodesRegistry.Contract.SupportsInterface(&_NodesRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NodesRegistry *NodesRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NodesRegistry.Contract.SupportsInterface(&_NodesRegistry.CallOpts, interfaceId)
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

// Attach is a paid mutator transaction binding the contract method 0x7a0ca1e2.
//
// Solidity: function attach(address server) returns()
func (_NodesRegistry *NodesRegistryTransactor) Attach(opts *bind.TransactOpts, server common.Address) (*types.Transaction, error) {
	return _NodesRegistry.contract.Transact(opts, "attach", server)
}

// Attach is a paid mutator transaction binding the contract method 0x7a0ca1e2.
//
// Solidity: function attach(address server) returns()
func (_NodesRegistry *NodesRegistrySession) Attach(server common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.Attach(&_NodesRegistry.TransactOpts, server)
}

// Attach is a paid mutator transaction binding the contract method 0x7a0ca1e2.
//
// Solidity: function attach(address server) returns()
func (_NodesRegistry *NodesRegistryTransactorSession) Attach(server common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.Attach(&_NodesRegistry.TransactOpts, server)
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

// Dettach is a paid mutator transaction binding the contract method 0x0e22e7f8.
//
// Solidity: function dettach(address server) returns()
func (_NodesRegistry *NodesRegistryTransactor) Dettach(opts *bind.TransactOpts, server common.Address) (*types.Transaction, error) {
	return _NodesRegistry.contract.Transact(opts, "dettach", server)
}

// Dettach is a paid mutator transaction binding the contract method 0x0e22e7f8.
//
// Solidity: function dettach(address server) returns()
func (_NodesRegistry *NodesRegistrySession) Dettach(server common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.Dettach(&_NodesRegistry.TransactOpts, server)
}

// Dettach is a paid mutator transaction binding the contract method 0x0e22e7f8.
//
// Solidity: function dettach(address server) returns()
func (_NodesRegistry *NodesRegistryTransactorSession) Dettach(server common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.Dettach(&_NodesRegistry.TransactOpts, server)
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

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NodesRegistry *NodesRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodesRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NodesRegistry *NodesRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.GrantRole(&_NodesRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NodesRegistry *NodesRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.GrantRole(&_NodesRegistry.TransactOpts, role, account)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xff7b178d.
//
// Solidity: function registerNode(address wallet, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums, bool isPublic) payable returns()
func (_NodesRegistry *NodesRegistryTransactor) RegisterNode(opts *bind.TransactOpts, wallet common.Address, aliasIdentifier string, gpuTypes []string, gpuNums []*big.Int, isPublic bool) (*types.Transaction, error) {
	return _NodesRegistry.contract.Transact(opts, "registerNode", wallet, aliasIdentifier, gpuTypes, gpuNums, isPublic)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xff7b178d.
//
// Solidity: function registerNode(address wallet, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums, bool isPublic) payable returns()
func (_NodesRegistry *NodesRegistrySession) RegisterNode(wallet common.Address, aliasIdentifier string, gpuTypes []string, gpuNums []*big.Int, isPublic bool) (*types.Transaction, error) {
	return _NodesRegistry.Contract.RegisterNode(&_NodesRegistry.TransactOpts, wallet, aliasIdentifier, gpuTypes, gpuNums, isPublic)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xff7b178d.
//
// Solidity: function registerNode(address wallet, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums, bool isPublic) payable returns()
func (_NodesRegistry *NodesRegistryTransactorSession) RegisterNode(wallet common.Address, aliasIdentifier string, gpuTypes []string, gpuNums []*big.Int, isPublic bool) (*types.Transaction, error) {
	return _NodesRegistry.Contract.RegisterNode(&_NodesRegistry.TransactOpts, wallet, aliasIdentifier, gpuTypes, gpuNums, isPublic)
}

// RegisterProxyNode is a paid mutator transaction binding the contract method 0xd8fd0eb1.
//
// Solidity: function registerProxyNode(address proxy) returns()
func (_NodesRegistry *NodesRegistryTransactor) RegisterProxyNode(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _NodesRegistry.contract.Transact(opts, "registerProxyNode", proxy)
}

// RegisterProxyNode is a paid mutator transaction binding the contract method 0xd8fd0eb1.
//
// Solidity: function registerProxyNode(address proxy) returns()
func (_NodesRegistry *NodesRegistrySession) RegisterProxyNode(proxy common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.RegisterProxyNode(&_NodesRegistry.TransactOpts, proxy)
}

// RegisterProxyNode is a paid mutator transaction binding the contract method 0xd8fd0eb1.
//
// Solidity: function registerProxyNode(address proxy) returns()
func (_NodesRegistry *NodesRegistryTransactorSession) RegisterProxyNode(proxy common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.RegisterProxyNode(&_NodesRegistry.TransactOpts, proxy)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NodesRegistry *NodesRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodesRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NodesRegistry *NodesRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.RevokeRole(&_NodesRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NodesRegistry *NodesRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodesRegistry.Contract.RevokeRole(&_NodesRegistry.TransactOpts, role, account)
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

// NodesRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NodesRegistry contract.
type NodesRegistryInitializedIterator struct {
	Event *NodesRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *NodesRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRegistryInitialized)
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
		it.Event = new(NodesRegistryInitialized)
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
func (it *NodesRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRegistryInitialized represents a Initialized event raised by the NodesRegistry contract.
type NodesRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodesRegistry *NodesRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*NodesRegistryInitializedIterator, error) {

	logs, sub, err := _NodesRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodesRegistryInitializedIterator{contract: _NodesRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodesRegistry *NodesRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodesRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _NodesRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRegistryInitialized)
				if err := _NodesRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NodesRegistry *NodesRegistryFilterer) ParseInitialized(log types.Log) (*NodesRegistryInitialized, error) {
	event := new(NodesRegistryInitialized)
	if err := _NodesRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
	GpuTypes        []string
	GpuNums         []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeActived is a free log retrieval operation binding the contract event 0xb0a73f160683fa6bd1601bd25cd9e2addc090f0a74959bf51c6a2a4560af6f5f.
//
// Solidity: event NodeActived(address indexed wallet, address identifier, uint256 time, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums)
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

// WatchNodeActived is a free log subscription operation binding the contract event 0xb0a73f160683fa6bd1601bd25cd9e2addc090f0a74959bf51c6a2a4560af6f5f.
//
// Solidity: event NodeActived(address indexed wallet, address identifier, uint256 time, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums)
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

// ParseNodeActived is a log parse operation binding the contract event 0xb0a73f160683fa6bd1601bd25cd9e2addc090f0a74959bf51c6a2a4560af6f5f.
//
// Solidity: event NodeActived(address indexed wallet, address identifier, uint256 time, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums)
func (_NodesRegistry *NodesRegistryFilterer) ParseNodeActived(log types.Log) (*NodesRegistryNodeActived, error) {
	event := new(NodesRegistryNodeActived)
	if err := _NodesRegistry.contract.UnpackLog(event, "NodeActived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesRegistryNodeAttachedIterator is returned from FilterNodeAttached and is used to iterate over the raw logs and unpacked data for NodeAttached events raised by the NodesRegistry contract.
type NodesRegistryNodeAttachedIterator struct {
	Event *NodesRegistryNodeAttached // Event containing the contract specifics and raw log

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
func (it *NodesRegistryNodeAttachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRegistryNodeAttached)
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
		it.Event = new(NodesRegistryNodeAttached)
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
func (it *NodesRegistryNodeAttachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRegistryNodeAttachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRegistryNodeAttached represents a NodeAttached event raised by the NodesRegistry contract.
type NodesRegistryNodeAttached struct {
	IdentifierOfProvider common.Address
	IdentifierOfServer   common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNodeAttached is a free log retrieval operation binding the contract event 0x042e2dbbe7288392f225b5cc7f610946ae1a1e8673b3f8ce4f6fc3bc6dee5aa6.
//
// Solidity: event NodeAttached(address indexed identifierOfProvider, address indexed identifierOfServer)
func (_NodesRegistry *NodesRegistryFilterer) FilterNodeAttached(opts *bind.FilterOpts, identifierOfProvider []common.Address, identifierOfServer []common.Address) (*NodesRegistryNodeAttachedIterator, error) {

	var identifierOfProviderRule []interface{}
	for _, identifierOfProviderItem := range identifierOfProvider {
		identifierOfProviderRule = append(identifierOfProviderRule, identifierOfProviderItem)
	}
	var identifierOfServerRule []interface{}
	for _, identifierOfServerItem := range identifierOfServer {
		identifierOfServerRule = append(identifierOfServerRule, identifierOfServerItem)
	}

	logs, sub, err := _NodesRegistry.contract.FilterLogs(opts, "NodeAttached", identifierOfProviderRule, identifierOfServerRule)
	if err != nil {
		return nil, err
	}
	return &NodesRegistryNodeAttachedIterator{contract: _NodesRegistry.contract, event: "NodeAttached", logs: logs, sub: sub}, nil
}

// WatchNodeAttached is a free log subscription operation binding the contract event 0x042e2dbbe7288392f225b5cc7f610946ae1a1e8673b3f8ce4f6fc3bc6dee5aa6.
//
// Solidity: event NodeAttached(address indexed identifierOfProvider, address indexed identifierOfServer)
func (_NodesRegistry *NodesRegistryFilterer) WatchNodeAttached(opts *bind.WatchOpts, sink chan<- *NodesRegistryNodeAttached, identifierOfProvider []common.Address, identifierOfServer []common.Address) (event.Subscription, error) {

	var identifierOfProviderRule []interface{}
	for _, identifierOfProviderItem := range identifierOfProvider {
		identifierOfProviderRule = append(identifierOfProviderRule, identifierOfProviderItem)
	}
	var identifierOfServerRule []interface{}
	for _, identifierOfServerItem := range identifierOfServer {
		identifierOfServerRule = append(identifierOfServerRule, identifierOfServerItem)
	}

	logs, sub, err := _NodesRegistry.contract.WatchLogs(opts, "NodeAttached", identifierOfProviderRule, identifierOfServerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRegistryNodeAttached)
				if err := _NodesRegistry.contract.UnpackLog(event, "NodeAttached", log); err != nil {
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

// ParseNodeAttached is a log parse operation binding the contract event 0x042e2dbbe7288392f225b5cc7f610946ae1a1e8673b3f8ce4f6fc3bc6dee5aa6.
//
// Solidity: event NodeAttached(address indexed identifierOfProvider, address indexed identifierOfServer)
func (_NodesRegistry *NodesRegistryFilterer) ParseNodeAttached(log types.Log) (*NodesRegistryNodeAttached, error) {
	event := new(NodesRegistryNodeAttached)
	if err := _NodesRegistry.contract.UnpackLog(event, "NodeAttached", log); err != nil {
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

// NodesRegistryNodeDetachedIterator is returned from FilterNodeDetached and is used to iterate over the raw logs and unpacked data for NodeDetached events raised by the NodesRegistry contract.
type NodesRegistryNodeDetachedIterator struct {
	Event *NodesRegistryNodeDetached // Event containing the contract specifics and raw log

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
func (it *NodesRegistryNodeDetachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRegistryNodeDetached)
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
		it.Event = new(NodesRegistryNodeDetached)
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
func (it *NodesRegistryNodeDetachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRegistryNodeDetachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRegistryNodeDetached represents a NodeDetached event raised by the NodesRegistry contract.
type NodesRegistryNodeDetached struct {
	IdentifierOfProvider common.Address
	IdentifierOfServer   common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNodeDetached is a free log retrieval operation binding the contract event 0x864dd06d15940858b627540246db79e66082ff76b9712b0bb5d483d168b1cd18.
//
// Solidity: event NodeDetached(address indexed identifierOfProvider, address indexed identifierOfServer)
func (_NodesRegistry *NodesRegistryFilterer) FilterNodeDetached(opts *bind.FilterOpts, identifierOfProvider []common.Address, identifierOfServer []common.Address) (*NodesRegistryNodeDetachedIterator, error) {

	var identifierOfProviderRule []interface{}
	for _, identifierOfProviderItem := range identifierOfProvider {
		identifierOfProviderRule = append(identifierOfProviderRule, identifierOfProviderItem)
	}
	var identifierOfServerRule []interface{}
	for _, identifierOfServerItem := range identifierOfServer {
		identifierOfServerRule = append(identifierOfServerRule, identifierOfServerItem)
	}

	logs, sub, err := _NodesRegistry.contract.FilterLogs(opts, "NodeDetached", identifierOfProviderRule, identifierOfServerRule)
	if err != nil {
		return nil, err
	}
	return &NodesRegistryNodeDetachedIterator{contract: _NodesRegistry.contract, event: "NodeDetached", logs: logs, sub: sub}, nil
}

// WatchNodeDetached is a free log subscription operation binding the contract event 0x864dd06d15940858b627540246db79e66082ff76b9712b0bb5d483d168b1cd18.
//
// Solidity: event NodeDetached(address indexed identifierOfProvider, address indexed identifierOfServer)
func (_NodesRegistry *NodesRegistryFilterer) WatchNodeDetached(opts *bind.WatchOpts, sink chan<- *NodesRegistryNodeDetached, identifierOfProvider []common.Address, identifierOfServer []common.Address) (event.Subscription, error) {

	var identifierOfProviderRule []interface{}
	for _, identifierOfProviderItem := range identifierOfProvider {
		identifierOfProviderRule = append(identifierOfProviderRule, identifierOfProviderItem)
	}
	var identifierOfServerRule []interface{}
	for _, identifierOfServerItem := range identifierOfServer {
		identifierOfServerRule = append(identifierOfServerRule, identifierOfServerItem)
	}

	logs, sub, err := _NodesRegistry.contract.WatchLogs(opts, "NodeDetached", identifierOfProviderRule, identifierOfServerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRegistryNodeDetached)
				if err := _NodesRegistry.contract.UnpackLog(event, "NodeDetached", log); err != nil {
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

// ParseNodeDetached is a log parse operation binding the contract event 0x864dd06d15940858b627540246db79e66082ff76b9712b0bb5d483d168b1cd18.
//
// Solidity: event NodeDetached(address indexed identifierOfProvider, address indexed identifierOfServer)
func (_NodesRegistry *NodesRegistryFilterer) ParseNodeDetached(log types.Log) (*NodesRegistryNodeDetached, error) {
	event := new(NodesRegistryNodeDetached)
	if err := _NodesRegistry.contract.UnpackLog(event, "NodeDetached", log); err != nil {
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
	GpuTypes        []string
	GpuNums         []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeRegistered is a free log retrieval operation binding the contract event 0x99c250edfb33141684f230eecda8aa955bf0d62d17bcaaab0ab43f318f3637b0.
//
// Solidity: event NodeRegistered(address indexed wallet, address identifier, uint256 time, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums)
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

// WatchNodeRegistered is a free log subscription operation binding the contract event 0x99c250edfb33141684f230eecda8aa955bf0d62d17bcaaab0ab43f318f3637b0.
//
// Solidity: event NodeRegistered(address indexed wallet, address identifier, uint256 time, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums)
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

// ParseNodeRegistered is a log parse operation binding the contract event 0x99c250edfb33141684f230eecda8aa955bf0d62d17bcaaab0ab43f318f3637b0.
//
// Solidity: event NodeRegistered(address indexed wallet, address identifier, uint256 time, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums)
func (_NodesRegistry *NodesRegistryFilterer) ParseNodeRegistered(log types.Log) (*NodesRegistryNodeRegistered, error) {
	event := new(NodesRegistryNodeRegistered)
	if err := _NodesRegistry.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesRegistryProxyNodeRegisteredIterator is returned from FilterProxyNodeRegistered and is used to iterate over the raw logs and unpacked data for ProxyNodeRegistered events raised by the NodesRegistry contract.
type NodesRegistryProxyNodeRegisteredIterator struct {
	Event *NodesRegistryProxyNodeRegistered // Event containing the contract specifics and raw log

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
func (it *NodesRegistryProxyNodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRegistryProxyNodeRegistered)
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
		it.Event = new(NodesRegistryProxyNodeRegistered)
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
func (it *NodesRegistryProxyNodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRegistryProxyNodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRegistryProxyNodeRegistered represents a ProxyNodeRegistered event raised by the NodesRegistry contract.
type NodesRegistryProxyNodeRegistered struct {
	Proxy common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProxyNodeRegistered is a free log retrieval operation binding the contract event 0x34093c616a3ab1dd56c3a4780155eb800bf5d1c3d024468ea10940d25fa9538d.
//
// Solidity: event ProxyNodeRegistered(address indexed proxy)
func (_NodesRegistry *NodesRegistryFilterer) FilterProxyNodeRegistered(opts *bind.FilterOpts, proxy []common.Address) (*NodesRegistryProxyNodeRegisteredIterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _NodesRegistry.contract.FilterLogs(opts, "ProxyNodeRegistered", proxyRule)
	if err != nil {
		return nil, err
	}
	return &NodesRegistryProxyNodeRegisteredIterator{contract: _NodesRegistry.contract, event: "ProxyNodeRegistered", logs: logs, sub: sub}, nil
}

// WatchProxyNodeRegistered is a free log subscription operation binding the contract event 0x34093c616a3ab1dd56c3a4780155eb800bf5d1c3d024468ea10940d25fa9538d.
//
// Solidity: event ProxyNodeRegistered(address indexed proxy)
func (_NodesRegistry *NodesRegistryFilterer) WatchProxyNodeRegistered(opts *bind.WatchOpts, sink chan<- *NodesRegistryProxyNodeRegistered, proxy []common.Address) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _NodesRegistry.contract.WatchLogs(opts, "ProxyNodeRegistered", proxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRegistryProxyNodeRegistered)
				if err := _NodesRegistry.contract.UnpackLog(event, "ProxyNodeRegistered", log); err != nil {
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

// ParseProxyNodeRegistered is a log parse operation binding the contract event 0x34093c616a3ab1dd56c3a4780155eb800bf5d1c3d024468ea10940d25fa9538d.
//
// Solidity: event ProxyNodeRegistered(address indexed proxy)
func (_NodesRegistry *NodesRegistryFilterer) ParseProxyNodeRegistered(log types.Log) (*NodesRegistryProxyNodeRegistered, error) {
	event := new(NodesRegistryProxyNodeRegistered)
	if err := _NodesRegistry.contract.UnpackLog(event, "ProxyNodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the NodesRegistry contract.
type NodesRegistryRoleAdminChangedIterator struct {
	Event *NodesRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NodesRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRegistryRoleAdminChanged)
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
		it.Event = new(NodesRegistryRoleAdminChanged)
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
func (it *NodesRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the NodesRegistry contract.
type NodesRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NodesRegistry *NodesRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NodesRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _NodesRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NodesRegistryRoleAdminChangedIterator{contract: _NodesRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NodesRegistry *NodesRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NodesRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _NodesRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRegistryRoleAdminChanged)
				if err := _NodesRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_NodesRegistry *NodesRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*NodesRegistryRoleAdminChanged, error) {
	event := new(NodesRegistryRoleAdminChanged)
	if err := _NodesRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the NodesRegistry contract.
type NodesRegistryRoleGrantedIterator struct {
	Event *NodesRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *NodesRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRegistryRoleGranted)
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
		it.Event = new(NodesRegistryRoleGranted)
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
func (it *NodesRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRegistryRoleGranted represents a RoleGranted event raised by the NodesRegistry contract.
type NodesRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NodesRegistry *NodesRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NodesRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _NodesRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NodesRegistryRoleGrantedIterator{contract: _NodesRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NodesRegistry *NodesRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NodesRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NodesRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRegistryRoleGranted)
				if err := _NodesRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_NodesRegistry *NodesRegistryFilterer) ParseRoleGranted(log types.Log) (*NodesRegistryRoleGranted, error) {
	event := new(NodesRegistryRoleGranted)
	if err := _NodesRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the NodesRegistry contract.
type NodesRegistryRoleRevokedIterator struct {
	Event *NodesRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NodesRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesRegistryRoleRevoked)
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
		it.Event = new(NodesRegistryRoleRevoked)
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
func (it *NodesRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesRegistryRoleRevoked represents a RoleRevoked event raised by the NodesRegistry contract.
type NodesRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NodesRegistry *NodesRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NodesRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _NodesRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NodesRegistryRoleRevokedIterator{contract: _NodesRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NodesRegistry *NodesRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NodesRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NodesRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesRegistryRoleRevoked)
				if err := _NodesRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_NodesRegistry *NodesRegistryFilterer) ParseRoleRevoked(log types.Log) (*NodesRegistryRoleRevoked, error) {
	event := new(NodesRegistryRoleRevoked)
	if err := _NodesRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
