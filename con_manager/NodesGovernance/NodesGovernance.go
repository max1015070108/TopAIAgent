// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NodesGovernance

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

// NodeInfo is an auto generated low-level Go binding around an user-defined struct.
type NodeInfo struct {
	Identifier      common.Address
	AliasIdentifier string
	Wallet          common.Address
	GpuTypes        []string
	GpuNums         []*big.Int
}

// NodeState is an auto generated low-level Go binding around an user-defined struct.
type NodeState struct {
	FailedCnt     uint64
	SuccessfulCnt uint64
	ExpectCnt     *big.Int
	Wallet        common.Address
	Identifier    common.Address
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

// NodesGovernanceMetaData contains all meta data concerning the NodesGovernance contract.
var NodesGovernanceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"Authorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"gpuTypes\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"gpuNums\",\"type\":\"uint256[]\"}],\"name\":\"NodeActived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identifierOfProvider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identifierOfServer\",\"type\":\"address\"}],\"name\":\"NodeAttached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"}],\"name\":\"NodeDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identifierOfProvider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identifierOfServer\",\"type\":\"address\"}],\"name\":\"NodeDetached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"gpuTypes\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"gpuNums\",\"type\":\"uint256[]\"}],\"name\":\"NodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"ProxyNodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"failedCnt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"successfulCnt\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"expectCnt\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structNodeState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalQuota\",\"type\":\"uint256\"}],\"name\":\"SettlementResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"ValidationResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expectedCompletionTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"ValidationStarted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_CANDIDATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_PER_CANDIDATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"gpuTypes\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gpuNums\",\"type\":\"uint256[]\"}],\"name\":\"allocGPU\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"internalType\":\"structNodeComputeUsed[]\",\"name\":\"gpuNodes\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"at\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"registrationTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"internalType\":\"structNodesRegistry.ComputeAvailable[]\",\"name\":\"gpus\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"internalType\":\"structNodesRegistry.Node\",\"name\":\"node\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"server\",\"type\":\"address\"}],\"name\":\"attach\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidatePerRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numOfNodes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedCompletionTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentDetectCircleId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentDetectCircleStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRoundStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregisterNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"detectDurationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"detectPeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"server\",\"type\":\"address\"}],\"name\":\"dettach\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"internalType\":\"structNodeComputeUsed[]\",\"name\":\"gpuNodes\",\"type\":\"tuple[]\"}],\"name\":\"freeGPU\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"registrationTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"internalType\":\"structNodesRegistry.ComputeAvailable[]\",\"name\":\"gpus\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"internalType\":\"structNodesRegistry.Node\",\"name\":\"node\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getAttach\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"detectPeriodId\",\"type\":\"uint256\"}],\"name\":\"getOnePeriodSettlement\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"failedCnt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"successfulCnt\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"expectCnt\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"}],\"internalType\":\"structNodeState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"totalQuotas\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getRoundCandidates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"candidates\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"getValidatorsOfCandidate\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"gpuSummary\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"gpuTypeOfNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"gpuTypes\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gpuNums\",\"type\":\"uint256[]\"}],\"internalType\":\"structNodeInfo[]\",\"name\":\"_nodesInfos\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_allocator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_roundDurationTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"}],\"name\":\"nodesGovernance_initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proxyNodes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"aliasIdentifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"gpuTypes\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gpuNums\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"name\":\"registerNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"registerProxyNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundDurationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"detectPeriodId\",\"type\":\"uint256\"}],\"name\":\"settlementOnePeriod\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"failedCnt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"successfulCnt\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"expectCnt\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"identifier\",\"type\":\"address\"}],\"internalType\":\"structNodeState[]\",\"name\":\"states\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"totalQuotas\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIStake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startNewValidationRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"detectId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorsPerCandidate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votedPerCandidate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"yesVotes\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"noVotes\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234610021576110e0600c5560b4600d55614d7190816100278239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a7146102d2578063036fe9c2146102cd5780630732bd7e146102c85780630e22e7f8146102c357806318767f31146102be57806318b1c081146102b9578063191805d8146102b45780631bb12a43146102af5780631f7b6d32146102aa578063248a9ca3146102a557806326cc3399146102465780632ad4f0ad146102a05780632f2ff15d1461029b57806336568abe14610296578063380dd9011461029157806349d2b4201461028c5780634d4fc0b81461028757806351ed6a30146102825780635ff5c03c1461027d5780636252e1c2146102785780636300c9511461027357806363c941991461026e578063698083af1461026957806375b238fc146102645780637a0ca1e21461025f57806391d148541461025a57806393e9d413146102555780639cbe5efd14610250578063a217fddf1461024b578063a7f3231414610246578063aa5dcecc14610241578063b8afa39c1461023c578063c23697a814610237578063c2bc2efc14610232578063c7edca7a1461022d578063d547741f14610228578063d8fd0eb114610223578063de1372701461021e578063e0886f9014610219578063e2cdd42a14610214578063ed38ed0d1461020f578063efa219e01461020a578063f67c5bdc146102055763ff7b178d1461020057600080fd5b612437565b6123a3565b612197565b612119565b611f77565b611f49565b611f2b565b611e5d565b611d94565b611d1a565b611c3b565b611afc565b611aba565b611a91565b61096e565b611a75565b611a57565b6119e4565b61198e565b6118d3565b611898565b61184a565b6116c0565b611605565b611256565b611139565b611110565b610bfd565b610b4b565b610b2d565b610ade565b610a11565b61098a565b61093f565b610921565b610903565b6108e5565b61078c565b610713565b6105ea565b61053f565b61035d565b346103285760203660031901126103285760043563ffffffff60e01b811680910361032857602090637965db0b60e01b8114908115610317575b506040519015158152f35b6301ffc9a760e01b1490503861030c565b600080fd5b9181601f84011215610328578235916001600160401b038311610328576020808501948460051b01011161032857565b346103285760208060031936011261032857600480356001600160401b0381116103285761038e903690830161032d565b600a549093906103b9906103b2906001600160a01b03165b6001600160a01b031690565b3314613e01565b60005b8481106103c557005b80856103dd6103d861041e948388613fad565b6136c5565b6103f16001600160a01b0382161515613fcf565b6001600160a01b0381166000908152600360205260409020600381015460ff16610423575b505050612bd8565b6103bc565b6104946104656104486104e09460018060a01b03166000526005602052604060002090565b61045f61045688888d613fad565b8a8101906136cf565b90613ec1565b5491610472831515614010565b8961048d8a610485896040998a93613fad565b013594612b70565b9101613f91565b50906104a66002809301918254612b7f565b90556104c86104c26104b9868c8b613fad565b888101906136cf565b90613e8f565b926104d4858b8a613fad565b01359201918254612b7f565b9055853880610416565b6001600160a01b0381160361032857565b6020908160408183019282815285518094520193019160005b828110610522575050505090565b83516001600160a01b031685529381019392810192600101610514565b346103285760403660031901126103285760243561055c816104ea565b600435600052610588602091600f835260406000209060018060a01b0316600052602052604060002090565b90604051908181845491828152019360005281600020916000905b8282106105ca576105c6856105ba81890382611fff565b604051918291826104fb565b0390f35b83546001600160a01b0316865294850194600193840193909101906105a3565b3461032857602036600319011261032857600435610607816104ea565b600090338252600360205261064260018060a01b0361062d81604086205416151561395f565b33845260026020526040842092168092613bd6565b1561066f57337f864dd06d15940858b627540246db79e66082ff76b9712b0bb5d483d168b1cd188380a380f35b60405162461bcd60e51b815260206004820152601260248201527118da1a5b19081a5cc81b9bdd08195e1a5cdd60721b6044820152606490fd5b634e487b7160e01b600052603260045260246000fd5b6008548110156106f65760086000527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30190600090565b6106a9565b80548210156106f65760005260206000200190600090565b3461032857606036600319011261032857602435610730816104ea565b61075e60443591600435600052600f60205260406000209060018060a01b0316600052602052604060002090565b805482101561032857602091610773916106fb565b905460405160039290921b1c6001600160a01b03168152f35b34610328576000806003193601126108e2576107a9331515613fcf565b33600090815260036020526040902080546107d7906107d0906001600160a01b03166103a6565b151561404f565b3360009081526002602052604090206107f1905415614092565b6001806107ff818401612388565b926004859101925b610884575b50505061082b61081b826120f3565b80546001600160a01b0319169055565b336000908152600360205260409020610843906141d1565b61084c33613b37565b507f60d01d146c7aa1a7d4e3fdd5543872f7d5b2a241980a66b3552ae1a86ae184536040518061087e3394428361425a565b0390a280f35b82548110156108dd576108d7818661089d859487613f91565b50848101546108b8866108af846140de565b01918254612b7f565b90556002906108cf828201926108af8454936140de565b905555612bd8565b90610807565b61080c565b80fd5b34610328576000366003190112610328576020601654604051908152f35b34610328576000366003190112610328576020601754604051908152f35b34610328576000366003190112610328576020600854604051908152f35b346103285760203660031901126103285760043560005260016020526020600160406000200154604051908152f35b3461032857600036600319011261032857602060405160058152f35b346103285760403660031901126103285760806109d36024356109ac816104ea565b600435600052601060205260406000209060018060a01b0316600052602052604060002090565b60018060a01b038154169060ff6002600183015492015416906040519283526001600160801b0381166020840152831c604083015215156060820152f35b3461032857604036600319011261032857600435602435610a31816104ea565b6000918083526001602052610a4c60016040852001546125b1565b8083526001602090815260408085206001600160a01b0385166000908152925290205460ff1615610a7b578280f35b8083526001602090815260408085206001600160a01b038516600090815292529020805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8480a438808280f35b3461032857604036600319011261032857610afa6024356104ea565b60405162461bcd60e51b815260206004820152600b60248201526a1b9bdd081cdd5c1c1bdc9d60aa1b6044820152606490fd5b34610328576000366003190112610328576020601954604051908152f35b34610328576000366003190112610328576020600c54604051908152f35b9291906040808501908086528251809252606091828701926020809501926000915b86848410610b9e57505050505050930152565b855180516001600160401b039081168952828201511682890152838101516001600160801b031684890152828101516001600160a01b03908116848a0152608091820151169088015260a0909601959094019360019290920191610b8b565b346103285760208060031936011261032857600435906000610c2260165484106134b9565b610c36836000526012602052604060002090565b92835492610c45841515613505565b610c59826000526014602052604060002090565b93610c65855415613551565b600195610c7e610c79888301938454612b7f565b612777565b6001600160801b039390841686895b610f58575b5050509086929154965b610dd0575b5050509080929350610cbd826000526015602052604060002090565b558254610cc9816135cc565b9160005b828110610d19575050507fd51417935ddbb98970f20a5f6f9c5070ce90768d0e3bfaba49e7e2f8621debac60405180610d07858583610b69565b0390a16105c660405192839283610b69565b80610d3f610d2a610d8f93896106fb565b905460039190911b1c6001600160a01b031690565b610d846103a66002610d7684610d5f896000526013602052604060002090565b9060018060a01b0316600052602052604060002090565b01546001600160a01b031690565b610d94575b50612bd8565b610ccd565b610daf610db491610d5f866000526013602052604060002090565b613654565b610dbe8287613640565b52610dc98186613640565b5038610d89565b80548711610f5357610dfa610def886000526011602052604060002090565b958487015490612785565b946000845b610e1c575b5050610e138394959697612bd8565b96959493610c9c565b8154811015610f4e57610efc81610d5f87610e528d610e3f610d2a84988a6106fb565b9384916000526010602052604060002090565b8054909290610e69906001600160a01b03166103a6565b1561041657610e8690610d5f8c6000526013602052604060002090565b91610e9c610e98600283015460ff1690565b1590565b610f34570154878160801c91161115610f025780610ed0610ecb610ef793546001600160401b039060401c1690565b61359d565b67ffffffffffffffff60401b82549160401b169067ffffffffffffffff60401b1916179055565b612bd8565b90610dff565b80610f1a610ecb610ef793546001600160401b031690565b6001600160401b03166001600160401b0319825416179055565b505080610f1a610ecb610ef793546001600160401b031690565b610e04565b610ca1565b9089969798999291888a8c600854841015611102578392918691610f7e6110d996613d5b565b9087606095610f92610e9888860151151590565b80156110f1575b6110e55761102d6110619561101d61103d94610ef79a610fcc61104b98610fc68b5160018060a01b031690565b90612df9565b60a089015161100c906001600160a01b03168a519093906001600160a01b031695610ff5612020565b8181529d8e01526001600160801b031660408d0152565b6001600160a01b03909116908a0152565b6001600160a01b03166080880152565b6000526013602052604060002090565b91516001600160a01b031690565b60018060a01b0316600052602052604060002090565b815160208301516040808501516001600160401b0390931691901b67ffffffffffffffff60401b16176001600160801b0319608092831b1617825560608301516001830180546001600160a01b03199081166001600160a01b0393841617909155919093015160029092018054929093169116179055565b90919299989796610c8d565b50505050505050612bd8565b50604084015160028e015410610f99565b505050509897969598610c92565b3461032857600036600319011261032857600b546040516001600160a01b039091168152602090f35b34610328576020366003190112610328576004356000526011602052604060002060026001820154910154906105c66040519283928360209093929193604081019481520152565b60005b8381106111945750506000910152565b8181015183820152602001611184565b906020916111bd81518092818552858086019101611181565b601f01601f1916010190565b9291906040808501818652825180915260609081870192828260051b890101936020809601936000925b878585106112075750505050505050930152565b806001929394959697988d605f19908203018752895190848060a01b0382511681528580611240858501518a878601528a8501906111a4565b93015191015298019401940192949391906111f3565b346103285760603660031901126103285760048035906001600160401b036024358181116103285761128b903690840161032d565b9091604435908111610328576112a4903690850161032d565b600a54600093906112c1906103b2906001600160a01b03166103a6565b6112ca81613e42565b91849385915b83831061157b575050506008936112f06112eb838754612764565b613f2a565b885b86546112fe8183612785565b8b1080611572575b156115615761131861131d918c612bfd565b613d99565b6001600160a01b03811660009081526003602052604090209093909761134a610e9860038b015460ff1690565b61154e576001600160a01b0385166000908152600560205260408120989c9b989a9993989384019390925b87841080611545575b15611521578b9c9d8b9c611395868c9e9d9e613640565b5115611514576113aa9061045f878c8f613e74565b548015611514576113bd6113c391612b70565b87613f91565b50906001820190600282549301926113dd84548092612b7f565b9d8e156115045761143a8f918f8f8f8f60208f95611421878f966114299361141c6114339861140c8b8a613640565b516001600160a01b039091169052565b613e74565b959093613640565b510192369161207b565b9052613640565b51111561149f57509161149393916114758e61148d958f61146f8c611469856114638386613640565b51612b7f565b92613640565b52612b7f565b9d6040611482858d613640565b510152549055612777565b93612bd8565b929c9b9a99989c611375565b611493949d506114d79250926114c461148d946114bd8f8b90613640565b5190612785565b90556114d0878d613640565b5190612b7f565b9a6114e2868c613640565b5160406114ef838b613640565b51015260006114fe878d613640565b52612777565b50509b5050509261149390612bd8565b5099509261149390612bd8565b999a509a909b979394506112fe925061153a9150612bd8565b9996929190506112f2565b508d151561137e565b9a61153a9198506112fe92939450612bd8565b50506105c6604051928392836111c9565b50861515611306565b9091946115f86115fe916115936104c289888d613e74565b6115d4600260018301549201916115c26115af84548093612b7f565b6115ba8d8a8c613eda565b351115613eea565b6115cd8b888a613eda565b3590612785565b90556115e1888587613eda565b356115ec8989613640565b526115cd888587613eda565b95612bd8565b91906112d0565b3461032857602036600319011261032857600435806000526014602052604060002090815490611634826135cc565b9160005b8181106116655783611654846000526015602052604060002090565b54906105c660405192839283610b69565b80611676610d2a6116a093886106fb565b6116966103a66002610d7684610d5f8a6000526013602052604060002090565b6116a55750612bd8565b611638565b610daf610db491610d5f876000526013602052604060002090565b34610328576000366003190112610328576116e96116e060195442612b7f565b600d5410612b8c565b60186116fd6116f88254612bd8565b601855565b61170642601955565b61171260175442612b7f565b600c54101561180e5761172e611729601654612bd8565b601655565b61173742601755565b611794815461174461202f565b90808252602082015242604082015242606082015261176f6016546000526012602052604060002090565b9060606003918051845560208101516001850155604081015160028501550151910155565b6117a0600d5442612785565b60026117b783546000526011602052604060002090565b01556117c460165461307d565b60016117db83546000526011602052604060002090565b01556117f760165482546117f1600d5442612785565b91612c82565b601654905460408051928352602083019190915290f35b805460016118286016546000526012602052604060002090565b01554260036118436016546000526012602052604060002090565b0155611794565b34610328576020366003190112610328576004356000526012602052608060406000208054906001810154906003600282015491015491604051938452602084015260408301526060820152f35b346103285760003660031901126103285760206040517fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217758152f35b34610328576020366003190112610328576004356118f0816104ea565b600090338252600360205261192b60018060a01b0361191681604086205416151561395f565b33845260026020526040842092168092613a40565b1561195857337f042e2dbbe7288392f225b5cc7f610946ae1a1e8673b3f8ce4f6fc3bc6dee5aa68380a380f35b60405162461bcd60e51b815260206004820152600e60248201526d18da1a5b19081a5cc8195e1a5cdd60921b6044820152606490fd5b3461032857604036600319011261032857602060ff6119d86024356119b2816104ea565b6004356000526001845260406000209060018060a01b0316600052602052604060002090565b54166040519015158152f35b34610328576020806003193601126103285760043560005260118152604060002090604051908181845491828152019360005281600020916000905b828210611a37576105c6856105ba81890382611fff565b83546001600160a01b031686529485019460019384019390910190611a20565b34610328576000366003190112610328576020601854604051908152f35b3461032857600036600319011261032857602060405160008152f35b3461032857600036600319011261032857600a546040516001600160a01b039091168152602090f35b3461032857602036600319011261032857600435611ad7816104ea565b60018060a01b03166000526007602052602060ff604060002054166040519015158152f35b34610328576020366003190112610328576020611b23600435611b1e816104ea565b613dd7565b6040519015158152f35b908082519081815260208091019281808460051b8301019501936000915b848310611b5b5750505050505090565b9091929394958480600192601f19858203018652895190611b84825160608084528301906111a4565b9183810151848301526040809101519101529801930193019194939290611b4b565b611c38906020815260018060a01b038351166020820152602083015160e0611c12611bdf610100938460408701526101208601906111a4565b60408701516060860152611bfc6060880151608087019015159052565b6080870151858203601f190160a0870152611b2d565b60a08601516001600160a01b031660c08501529460c08101518483015201511515910152565b90565b3461032857602036600319011261032857600435611c58816104ea565b611c60613c7f565b506001600160a01b031660009081526003602052604090206105c690611d0e611d056007611c8c61204e565b84546001600160a01b0316815293611ca660018201612388565b602086015260028101546040860152611ccf611cc6600383015460ff1690565b15156060870152565b611cdb60048201613cd0565b608086015260058101546001600160a01b031660a0860152600681015460c0860152015460ff1690565b151560e0830152565b60405191829182611ba6565b346103285760208060031936011261032857600435611d38816104ea565b6001600160a01b03166000908152600282526040808220905181548082529183528383208185019491939092915b828210611d7d576105c6856105ba81890382611fff565b835486529485019460019384019390910190611d66565b3461032857604036600319011261032857600435602435611db4816104ea565b6000918083526001602052611dcf60016040852001546125b1565b8083526001602090815260408085206001600160a01b0385166000908152925290205460ff16611dfd578280f35b8083526001602090815260408085206001600160a01b038516600090815292529020805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b8480a438808280f35b3461032857602036600319011261032857600435611e7a816104ea565b611e826124cf565b6001600160a01b03811660008181526007602052604081205490929060ff16611ef5576001600160a01b03166000908152600760205260409020611ece905b805460ff19166001179055565b7f34093c616a3ab1dd56c3a4780155eb800bf5d1c3d024468ea10940d25fa9538d8280a280f35b60405162461bcd60e51b815260206004820152600e60248201526d141c9bde1e481a5cc8195e1a5cdd60921b6044820152606490fd5b34610328576000366003190112610328576020600d54604051908152f35b34610328576020366003190112610328576105c6611d0e600435613d5b565b60843590811515820361032857565b3461032857606036600319011261032857602435611f94816104ea565b604435801515810361032857611fac91600435613208565b005b634e487b7160e01b600052604160045260246000fd5b60a081019081106001600160401b03821117611fdf57604052565b611fae565b606081019081106001600160401b03821117611fdf57604052565b90601f801991011681019081106001600160401b03821117611fdf57604052565b6040519061202d82611fc4565b565b60405190608082018281106001600160401b03821117611fdf57604052565b6040519061010082018281106001600160401b03821117611fdf57604052565b6040519061202d82611fe4565b9291926001600160401b038211611fdf57604051916120a4601f8201601f191660200184611fff565b829481845281830111610328578281602093846000960137010152565b9080601f8301121561032857816020611c389335910161207b565b906120ef60209282815194859201611181565b0190565b602061210c918160405193828580945193849201611181565b8101600481520301902090565b3461032857604036600319011261032857600435612136816104ea565b602435906001600160401b0382116103285760206121859161215d829436906004016120c1565b9060018060a01b03166000526005825260406000208260405194838680955193849201611181565b82019081520301902054604051908152f35b34610328576080366003190112610328576004356001600160401b038111610328576121ca61222c91369060040161032d565b9190602435906121d9826104ea565b606435916121e6836104ea565b6000549461220b60ff8760081c1615809781986122aa575b811561228a575b5061295e565b8561221e600160ff196000541617600055565b612271575b604435926129c1565b61223257005b61224261ff001960005416600055565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1005b61228561010061ff00196000541617600055565b612223565b303b1591508161229c575b5038612205565b6001915060ff161438612295565b600160ff82161091506121fe565b90600182811c921680156122e8575b60208310146122d257565b634e487b7160e01b600052602260045260246000fd5b91607f16916122c7565b9060009291805491612303836122b8565b9182825260019384811690816000146123655750600114612325575b50505050565b90919394506000526020928360002092846000945b83861061235157505050500101903880808061231f565b80548587018301529401938590820161233a565b9294505050602093945060ff191683830152151560051b0101903880808061231f565b9061202d61239c92604051938480926122f2565b0383611fff565b34610328576020366003190112610328576004356001600160401b038111610328576123ee60206123db6124289336906004016120c1565b8160405193828580945193849201611181565b810160068152030190206040519061240a8261239c81846122f2565b600260018201549101546040519384936060855260608501906111a4565b91602084015260408301520390f35b60a03660031901126103285760043561244f816104ea565b6024356001600160401b0391828211610328573660238301121561032857816004013590838211610328573660248385010111610328576044358481116103285761249e90369060040161032d565b90606435958611610328576124ba611fac96369060040161032d565b94909360246124c7611f68565b9701906137b8565b3360009081527f50efbde2d46c37e9785f1791697f77e94bb7b701e19f1930a668820722d37694602052604090205460ff161561250857565b6125ad604861259561251933612843565b6125876125246128cd565b6040519485937f416363657373436f6e74726f6c3a206163636f756e74200000000000000000006020860152612564815180926020603789019101611181565b84017001034b99036b4b9b9b4b733903937b6329607d1b603782015201906120dc565b03601f198101835282611fff565b60405162461bcd60e51b81529182916004830161266c565b0390fd5b60008181526001602081815260408084203385529091529091205490919060ff16156125db575050565b6125e433612843565b906125ed612792565b9260306125f9856127bd565b536078612605856127ca565b536041905b808211612628576125ad6048612595866125878961252489156127f8565b9091600f81169060108210156106f657612666916f181899199a1a9b1b9c1cb0b131b232b360811b901a61265c85886127da565b5360041c926127eb565b9061260a565b906020611c389281815201906111a4565b6001600160a01b03811660009081527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49602052604081205460ff16156126c1575050565b8080526001602090815260408083206001600160a01b038516600090815292529020805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4565b634e487b7160e01b600052601160045260246000fd5b600181901b91906001600160ff1b0381160361274e57565b612720565b908160180291601883040361274e57565b8181029291811591840414171561274e57565b906001820180921161274e57565b9190820180921161274e57565b60405190608082018281106001600160401b03821117611fdf57604052604282526060366020840137565b8051156106f65760200190565b8051600110156106f65760210190565b9081518110156106f6570160200190565b801561274e576000190190565b156127ff57565b606460405162461bcd60e51b815260206004820152602060248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152fd5b6040519061285082611fe4565b602a825260403660208401376030612867836127bd565b536078612873836127ca565b536029905b6001821161288b57611c389150156127f8565b600f81169060108210156106f6576128c7916f181899199a1a9b1b9c1cb0b131b232b360811b901a6128bd84866127da565b5360041c916127eb565b90612878565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756128f6612792565b906030612902836127bd565b53607861290e836127ca565b536041905b6001821161292657611c389150156127f8565b600f81169060108210156106f657612958916f181899199a1a9b1b9c1cb0b131b232b360811b901a6128bd84866127da565b90612913565b1561296557565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b94939060ff60005460081c1615612b175760005b818110612a7c575093945061202d9350612a54916001600160a01b039190612a2690612a048482161515613736565b60018060a01b03166bffffffffffffffffffffffff60a01b600a541617600a55565b16612a32811515613776565b60018060a01b03166bffffffffffffffffffffffff60a01b600b541617600b55565b612a5d3361267d565b612a6642601955565b612a77612a7282612753565b600c55565b600d55565b80612afc612a996040612a93612b1295878d6136a3565b016136c5565b612aa76103d884878d6136a3565b90858b612ac2612ab88784846136a3565b60208101906136cf565b90612af4612aea89612ae2612ad88289896136a3565b6060810190613701565b9790966136a3565b6080810190613701565b9690956146b4565b610ef7612b0d6103d883868c6136a3565b614c0a565b6129d5565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b60001981019190821161274e57565b9190820391821161274e57565b15612b9357565b60405162461bcd60e51b815260206004820152601c60248201527f50726576696f757320726f756e64206973206e6f7420656e64696e67000000006044820152606490fd5b600019811461274e5760010190565b634e487b7160e01b600052600060045260246000fd5b8115612c07570690565b634e487b7160e01b600052601260045260246000fd5b60a090939192936080810193815260209485820152600180831b038093166040820152608060608201528554809452019360005282600020926000915b838310612c6957505050505090565b8454811686529481019460019485019490920191612c5a565b9091612c9a6000926000526012602052604060002090565b60020192825b60058410612caf575050505050565b612cb843612b70565b4093604090612d0a612d05835197612cf889612cea886020830194428691606093918352602083015260408201520190565b03601f1981018b528a611fff565b8851902060085490612bfd565b613d5b565b90828201518854108015612de3575b612dd657612d44612d3486600052600f602052604060002090565b83516001600160a01b031661104b565b54612dd657917f71afff60b83105500984ce43d4633544224775a10de240da021704c056b58bdb91612d91612d7c612dd19695612bd8565b825190999088906001600160a01b0316612ed8565b8051612dc990612dbc906001600160a01b03169261104b61103d8a600052600f602052604060002090565b9251928392898985612c1d565b0390a1612bd8565b612ca0565b95505050612dd190612bd8565b50612df4610e986060840151151590565b612d19565b8054600160401b811015611fdf57612e16916001820181556106fb565b819291549060031b9160018060a01b03809116831b921b1916179055565b815181546001600160a01b039091166001600160a01b03199091161781556020820151604083015160801b6001600160801b0319166001600160801b0390911617600182015561202d9160029060600151151591019060ff801983541691151516179055565b6020929190612eb0849282815194859201611181565b019081520190565b60041115612ec257565b634e487b7160e01b600052602160045260246000fd5b929192600090612efb81612ef6856000526011602052604060002090565b612df9565b612f0361202f565b6001600160a01b03821681526020838183015260408481840152612f456060938685820152612f4086610d5f8a6000526010602052604060002090565b612e34565b845b600586108061306a575b1561305f57612f7d612d058a8451612f7181612587878a83019586612e9a565b51902060085490612bfd565b612f8c610e9886830151151590565b6130555787612fa987610d5f83600052600e602052604060002090565b82516001600160a01b038981169116908114159081613021575b50612fda575b505050612fd590612bd8565b612f47565b82610fc661103d8a610d5f612fd5989d96613008611ec1610ef79861104b6130189c5160018060a01b031690565b600052600f602052604060002090565b95908738612fc9565b6001600160a01b0316600090815260208390526040902061304591505b5460ff1690565b61304e81612eb8565b1538612fc3565b50612fd590612bd8565b505050505050509050565b50613076600854612736565b8110612f51565b6000906000526012602052604080600020600091600854915b8284106130a4575050505090565b9091929360606130b386613d5b565b01511580156130e5575b6130db576130cd6130d391612bd8565b94612bd8565b929190613096565b936130d390612bd8565b50826130f086613d5b565b01516002830154106130bd565b1561310457565b60405162461bcd60e51b815260206004820152601860248201527f56616c69646174696f6e2074696d6520657863656564656400000000000000006044820152606490fd5b1561315057565b60405162461bcd60e51b815260206004820152601160248201527024b73b30b634b2103b30b634b230ba37b960791b6044820152606490fd5b1561319057565b60405162461bcd60e51b815260206004820152601c60248201527f56616c69646174696f6e20616c726561647920636f6d706c65746564000000006044820152606490fd5b6001600160801b0380911690811461274e5760010190565b9190916001600160801b038080941691160191821161274e57565b9161322c6002613222856000526011602052604060002090565b01544211156130fd565b61324482610d5f85600052600e602052604060002090565b336000908152602082905260409020613272906001906132639061303e565b61326c81612eb8565b14613149565b61328a83610d5f866000526010602052604060002090565b9060028201926132a66132a1610e98865460ff1690565b613189565b1561346b576132ff61330c916132ec600185016132d26132cd82546001600160801b031690565b6131d5565b6001600160801b03166001600160801b0319825416179055565b3360009081526020919091526040902090565b805460ff19166003179055565b600161332684610d5f87600052600f602052604060002090565b549101908154916001600160801b03928381169060801c838561334983856131ed565b16149182613461575b505061342657549060011c80928216116000146133c7575050805460ff191660011790557f92de7c81b7cf6c7977d7cd091ccd01996264a02b7dbbce5d2a2524a8daabe769916133c2916133a590614c0a565b604080519182523360208301526001908201529081906060820190565b0390a1565b9091925060801c116133d7575050565b805460ff191660011790557f92de7c81b7cf6c7977d7cd091ccd01996264a02b7dbbce5d2a2524a8daabe769906133c2905b604080519182523360208301526000908201529081906060820190565b5050815460ff1916600117909155507f92de7c81b7cf6c7977d7cd091ccd01996264a02b7dbbce5d2a2524a8daabe769916133c29150613409565b1490503880613352565b6134a76134b4916132ec600185016134876132cd825460801c90565b81546001600160801b031660809190911b6001600160801b031916179055565b805460ff19166002179055565b61330c565b156134c057565b60405162461bcd60e51b815260206004820152601e60248201527f536574746c656d656e7420666f7220646574656374656420706572696f6400006044820152606490fd5b1561350c57565b60405162461bcd60e51b815260206004820152601d60248201527f44657465637420706572696f64206964206973206e6f742065786973740000006044820152606490fd5b1561355857565b60405162461bcd60e51b815260206004820152601760248201527f506572696f6420686173206265656e20736574746c65640000000000000000006044820152606490fd5b6001600160401b0380911690811461274e5760010190565b6001600160401b038111611fdf5760051b60200190565b906135d6826135b5565b6040906135e582519182611fff565b83815280936135f6601f19916135b5565b019160005b8381106136085750505050565b602090825161361681611fc4565b600081528260008183015260008583015260006060830152600060808301528286010152016135fb565b80518210156106f65760209160051b010190565b9060405161366181611fc4565b6080819380546001600160401b0380821685528160401c166020850152821c6040840152600260018060a01b0391826001820154166060860152015416910152565b91908110156106f65760051b81013590609e1981360301821215610328570190565b35611c38816104ea565b903590601e198136030182121561032857018035906001600160401b0382116103285760200191813603831361032857565b903590601e198136030182121561032857018035906001600160401b03821161032857602001918160051b3603831361032857565b1561373d57565b60405162461bcd60e51b815260206004820152601160248201527024b73b30b634b21030b63637b1b0ba37b960791b6044820152606490fd5b1561377d57565b60405162461bcd60e51b815260206004820152601360248201527224b73b30b634b21039ba30b5b2903a37b5b2b760691b6044820152606490fd5b906137ca97969594939291339061493b565b60186137d96116f88254612bd8565b6137e560175442612b7f565b600c541015613923576137fc611729601654612bd8565b61380542601755565b613812815461174461202f565b61381e600d5442612785565b600261383583546000526011602052604060002090565b015561384260165461307d565b600161385983546000526011602052604060002090565b01556138a4613890825461389e61386f43612b70565b60408051426020820152914090820152600160608201529283906080820190565b03601f198101845283611fff565b33612ed8565b546138bd33610d5f83600052600f602052604060002090565b546138cc575061202d33614c0a565b7f71afff60b83105500984ce43d4633544224775a10de240da021704c056b58bdb906138fa600d5442612785565b6133c261391533610d5f85600052600f602052604060002090565b604051938493339185612c1d565b8054600161393d6016546000526012602052604060002090565b01554260036139586016546000526012602052604060002090565b0155613812565b1561396657565b60405162461bcd60e51b815260206004820152601860248201527f4964656e746966696572206d75737420626520657869737400000000000000006044820152606490fd5b916139c59183549060031b91821b91600019901b19161790565b9055565b80600052600960205260406000205415600014613a3a57600854600160401b811015611fdf5760018101806008558110156106f65781907ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155600854906000526009602052604060002055600190565b50600090565b6001810190826000528160205260406000205415600014613aa9578054600160401b811015611fdf57613a94613a7d8260018794018555846106fb565b819391549060031b91821b91600019901b19161790565b90555491600052602052604060002055600190565b505050600090565b6008548015613af757600019810190808210156106f6577ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee2600091600883520155600855565b634e487b7160e01b600052603160045260246000fd5b8054908115613af75760001991820191613b2783836106fb565b909182549160031b1b1916905555565b6000818152600960205260409020548015613bcf57600019918183019180831161274e5760085493840193841161274e578383613b929460009603613b98575b505050613b82613ab1565b6000526009602052604060002090565b55600190565b613b82613bc091613bb8613bae613bc6956106bf565b90549060031b1c90565b9283916106bf565b906139ab565b55388080613b77565b5050600090565b6001810191806000528260205260406000205492831515600014613c7657600019928484019085821161274e57805494850194851161274e576000958583613c2a94613b929803613c39575b505050613b0d565b90600052602052604060002090565b613c5d613bc091613c4d613c6d94876106fb565b90549060031b1c928391876106fb565b8590600052602052604060002090565b55388080613c22565b50505050600090565b6040519061010082018281106001600160401b03821117611fdf576040528160e060009182815260606020820152826040820152826060820152606060808201528260a08201528260c08201520152565b908154613cdc816135b5565b92604093613cec85519182611fff565b828152809460208092019260005281600020906000935b858510613d1257505050505050565b6003846001928451613d2381611fe4565b8551613d3a81613d33818b6122f2565b0382611fff565b81528487015483820152600287015486820152815201930194019391613d03565b613d6d90613d67613c7f565b506106bf565b60018060a01b0391549060031b1c1660005260036020526040600020611c38611d056007611c8c61204e565b6008548110156106f65760086000527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee301546001600160a01b031690565b60018060a01b0316600052600360205260ff60036040600020015416613dfc57600090565b600190565b15613e0857565b60405162461bcd60e51b815260206004820152601260248201527127b7363c903337b91030b63637b1b0ba37b960711b6044820152606490fd5b90613e4c826135b5565b613e596040519182611fff565b8281528092613e6a601f19916135b5565b0190602036910137565b908210156106f657613e8b9160051b8101906136cf565b9091565b6020908260405193849283378101600681520301902090565b6020908260405193849283378101600481520301902090565b6020919283604051948593843782019081520301902090565b91908110156106f65760051b0190565b15613ef157565b60405162461bcd60e51b81526020600482015260116024820152700cee0ea40d2e640dcdee840cadcdeeaced607b1b6044820152606490fd5b90613f34826135b5565b604090613f4382519182611fff565b8381528093613f54601f19916135b5565b0191600091825b848110613f69575050505050565b6020908351613f7781611fe4565b858152826060818301528686830152828501015201613f5b565b80548210156106f6576000526003602060002091020190600090565b91908110156106f65760051b81013590605e1981360301821215610328570190565b15613fd657565b60405162461bcd60e51b815260206004820152601260248201527124b73b30b634b21034b232b73a34b334b2b960711b6044820152606490fd5b1561401757565b60405162461bcd60e51b815260206004820152601060248201526f496e76616c696420677075207479706560801b6044820152606490fd5b1561405657565b60405162461bcd60e51b81526020600482015260146024820152731259195b9d1a599a595c881b9bdd08195e1a5cdd60621b6044820152606490fd5b1561409957565b60405162461bcd60e51b815260206004820152601760248201527f4964656e74696669657220686173206368696c6472656e0000000000000000006044820152606490fd5b6040519081600082546140f0816122b8565b936001918083169081156141555750600114614118575b505060209250600681520301902090565b90915060005260209081600020906000915b858310614141575050505060209181013880614107565b80548784015286945091830191810161412a565b92505050602093915060ff191682528015150281013880614107565b81811061417c575050565b60008155600101614171565b61419281546122b8565b908161419c575050565b81601f600093116001146141ae575055565b9080839182526141cd601f60208420940160051c840160018501614171565b5555565b9060008083556001926141e5848201614188565b60029382858301556003948386840155600483019182549285815583614220575b505050508160079293945060058201558260068201550155565b838802938885040361274e57855260208520928301925b83811015614206578061424a8992614188565b8683820155868482015501614237565b604090611c389392815281602082015201906111a4565b1561427857565b60405162461bcd60e51b815260206004820152601060248201526f496e76616c696420475055206461746160801b6044820152606490fd5b156142b757565b60405162461bcd60e51b815260206004820152601c60248201527f496e76616c69642077616c6c6574206f72206964656e746966696572000000006044820152606490fd5b1561430357565b60405162461bcd60e51b815260206004820152601060248201526f1259195b9d1a599a595c88195e1a5cdd60821b6044820152606490fd5b1561434257565b60405162461bcd60e51b8152602060048201526016602482015275105b1a585cc81a59195b9d1a599a595c88195e1a5cdd60521b6044820152606490fd5b9190601f811161438f57505050565b61202d926000526020600020906020601f840160051c830193106143bb575b601f0160051c0190614171565b90915081906143ae565b9092916001600160401b038111611fdf576143ea816143e484546122b8565b84614380565b6000601f82116001146144245781929394600092614419575b50508160011b916000199060031b1c1916179055565b013590503880614403565b601f1982169461443984600052602060002090565b91805b87811061447457508360019596971061445a575b505050811b019055565b0135600019600384901b60f8161c19169055388080614450565b9092602060018192868601358155019401910161443c565b91908254600160401b811015611fdf576144ad906001948582018155613f91565b91909161459e5780519384516001600160401b038111611fdf576144db816144d586546122b8565b86614380565b60209081601f821160011461452c5790806040959493926002979899600092614521575b5050600019600383901b1c191690831b1786555b820151908501550151910155565b0151905038806144ff565b601f1982169061454187600052602060002090565b9160005b8181106145895750918391600298999a879695604099989510614570575b505050811b018655614513565b015160001960f88460031b161c19169055388080614563565b8a830151845592860192918501918501614545565b612be7565b908060209392818452848401376000828201840152601f01601f1916010190565b81835290916001600160fb1b0383116103285760209260051b809284830137010190565b9897959290936146179297959460018060a01b03168a52602097888b015260a060408b015260a08a01916145a3565b94878603606089015281865280860195818360051b82010196846000925b8584106146555750505050505050846080611c38959685039101526145c4565b90919293949598601f198282030184528935601e19843603018112156103285783018681019190356001600160401b038111610328578036038313610328576146a3889283926001956145a3565b9b0194019401929594939190614635565b9693909491959284831480614932575b6146cd90614271565b600160a01b6001900387818a169983898c151561474761475a9361472161475494846103a695614927575b8061491e575b614707906142b0565b6001600160a01b0316600090815260036020526040902090565b80549097906147429061473c906001600160a01b03166103a6565b156142fc565b613ea8565b546001600160a01b031690565b1561433b565b6001600160a01b03891660009081526005602052604090206147b99083546001600160a01b0319166001600160a01b038c16178455914260028501556005840180546001600160a01b0319166001600160a01b03909216919091179055565b6001906147c9858c8486016143c5565b34600684015560078301805460ff19166001179055600092600401835b888110614869575050505050918695939161482861486496947f99c250edfb33141684f230eecda8aa955bf0d62d17bcaaab0ab43f318f3637b09a99166139c9565b5061485687614837838b613ea8565b80546001600160a01b0319166001600160a01b03909216919091179055565b6040519788974290896145e8565b0390a2565b80886148f3866148ea8f8f968f6148e382848f9b8a8f8f906149199f86896148dd946148d66104c29e61045f958a6148c46148b2886148aa818f8b90613e74565b969094613eda565b35916148bc61206e565b94369161207b565b8352602083015260408201528261448c565b5495613e74565b55613eda565b3596613e74565b01918254612785565b9055610ef78a8a6149146104c28561490c818686613e74565b959094613e74565b6143c5565b6147e6565b508215156146fe565b5089811615156146f8565b508415156146c4565b969492909593919685841480614aeb575b61495590614271565b60018060a01b03908888614a478484169c8d151580614ae0575b80614ad7575b61497e906142b0565b6001600160a01b0383166000908152600360205260409020614a1d9080549096906149b59061473c906001600160a01b03166103a6565b6149c86147546103a66147478c8a613ea8565b6001600160a01b0394909416600081815260056020526040902087546001600160a01b0319169091178755934260028801556005870180546001600160a01b0319166001600160a01b03909216919091179055565b614a2c876001958688016143c5565b346006860155600785019060ff801983541691151516179055565b600092600401835b888110614a91575050505050918695939161482861486496947f99c250edfb33141684f230eecda8aa955bf0d62d17bcaaab0ab43f318f3637b09a99166139c9565b80886148f3866148ea8f8f968f6148e382848f9b8a8f8f90614ad29f86896148dd946148d66104c29e61045f958a6148c46148b2886148aa818f8b90613e74565b614a4f565b50861515614975565b50858316151561496f565b5085151561494c565b90614afe826135b5565b614b0b6040519182611fff565b8281528092614b1c601f19916135b5565b019060005b828110614b2d57505050565b806060602080938501015201614b21565b90815180825260208080930193019160005b828110614b5e575050505090565b835185529381019392810192600101614b50565b91959492614b9c9160018060a01b031683526020968784015260a0604084015260a08301906122f2565b94818603606083015283518087528187019180808360051b8a01019601926000905b838210614bdd575050505050611c389394506080818403910152614b3e565b90919293968380614bfb6001938d601f199082030186528b516111a4565b99019201920190939291614bbe565b6001600160a01b0390811660009081526003602052604090208054919291614c35908416151561404f565b60038101614c47610e98825460ff1690565b614c52575b50509050565b805460ff1916600117905560048101928354614c6d81614af4565b90614c7781613e42565b9060005b818110614cea575050600584015494955092937fb0a73f160683fa6bd1601bd25cd9e2addc090f0a74959bf51c6a2a4560af6f5f93614cdf916001600160a01b031686549093906001600160a01b0316600160405196879616980190429086614b72565b0390a2803880614c4c565b80614d01614cfb614d36938b613f91565b50612388565b614d0b8287613640565b52614d168186613640565b506001614d23828b613f91565b500154614d308286613640565b52612bd8565b614c7b56fea26469706673582212204bbc24d4dad341879380cae41426d94547f7bc2d2c72c046ed2074bd63d124c564736f6c63430008140033",
}

// NodesGovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use NodesGovernanceMetaData.ABI instead.
var NodesGovernanceABI = NodesGovernanceMetaData.ABI

// NodesGovernanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NodesGovernanceMetaData.Bin instead.
var NodesGovernanceBin = NodesGovernanceMetaData.Bin

// DeployNodesGovernance deploys a new Ethereum contract, binding an instance of NodesGovernance to it.
func DeployNodesGovernance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodesGovernance, error) {
	parsed, err := NodesGovernanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NodesGovernanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodesGovernance{NodesGovernanceCaller: NodesGovernanceCaller{contract: contract}, NodesGovernanceTransactor: NodesGovernanceTransactor{contract: contract}, NodesGovernanceFilterer: NodesGovernanceFilterer{contract: contract}}, nil
}

// NodesGovernance is an auto generated Go binding around an Ethereum contract.
type NodesGovernance struct {
	NodesGovernanceCaller     // Read-only binding to the contract
	NodesGovernanceTransactor // Write-only binding to the contract
	NodesGovernanceFilterer   // Log filterer for contract events
}

// NodesGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodesGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodesGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodesGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodesGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodesGovernanceSession struct {
	Contract     *NodesGovernance  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodesGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodesGovernanceCallerSession struct {
	Contract *NodesGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NodesGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodesGovernanceTransactorSession struct {
	Contract     *NodesGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NodesGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodesGovernanceRaw struct {
	Contract *NodesGovernance // Generic contract binding to access the raw methods on
}

// NodesGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodesGovernanceCallerRaw struct {
	Contract *NodesGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// NodesGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodesGovernanceTransactorRaw struct {
	Contract *NodesGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodesGovernance creates a new instance of NodesGovernance, bound to a specific deployed contract.
func NewNodesGovernance(address common.Address, backend bind.ContractBackend) (*NodesGovernance, error) {
	contract, err := bindNodesGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodesGovernance{NodesGovernanceCaller: NodesGovernanceCaller{contract: contract}, NodesGovernanceTransactor: NodesGovernanceTransactor{contract: contract}, NodesGovernanceFilterer: NodesGovernanceFilterer{contract: contract}}, nil
}

// NewNodesGovernanceCaller creates a new read-only instance of NodesGovernance, bound to a specific deployed contract.
func NewNodesGovernanceCaller(address common.Address, caller bind.ContractCaller) (*NodesGovernanceCaller, error) {
	contract, err := bindNodesGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceCaller{contract: contract}, nil
}

// NewNodesGovernanceTransactor creates a new write-only instance of NodesGovernance, bound to a specific deployed contract.
func NewNodesGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*NodesGovernanceTransactor, error) {
	contract, err := bindNodesGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceTransactor{contract: contract}, nil
}

// NewNodesGovernanceFilterer creates a new log filterer instance of NodesGovernance, bound to a specific deployed contract.
func NewNodesGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*NodesGovernanceFilterer, error) {
	contract, err := bindNodesGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceFilterer{contract: contract}, nil
}

// bindNodesGovernance binds a generic wrapper to an already deployed contract.
func bindNodesGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodesGovernanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodesGovernance *NodesGovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodesGovernance.Contract.NodesGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodesGovernance *NodesGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodesGovernance.Contract.NodesGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodesGovernance *NodesGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodesGovernance.Contract.NodesGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodesGovernance *NodesGovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodesGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodesGovernance *NodesGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodesGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodesGovernance *NodesGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodesGovernance.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_NodesGovernance *NodesGovernanceCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_NodesGovernance *NodesGovernanceSession) ADMINROLE() ([32]byte, error) {
	return _NodesGovernance.Contract.ADMINROLE(&_NodesGovernance.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_NodesGovernance *NodesGovernanceCallerSession) ADMINROLE() ([32]byte, error) {
	return _NodesGovernance.Contract.ADMINROLE(&_NodesGovernance.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NodesGovernance *NodesGovernanceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NodesGovernance *NodesGovernanceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NodesGovernance.Contract.DEFAULTADMINROLE(&_NodesGovernance.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NodesGovernance *NodesGovernanceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NodesGovernance.Contract.DEFAULTADMINROLE(&_NodesGovernance.CallOpts)
}

// MINCANDIDATE is a free data retrieval call binding the contract method 0x26cc3399.
//
// Solidity: function MIN_CANDIDATE() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCaller) MINCANDIDATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "MIN_CANDIDATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINCANDIDATE is a free data retrieval call binding the contract method 0x26cc3399.
//
// Solidity: function MIN_CANDIDATE() view returns(uint256)
func (_NodesGovernance *NodesGovernanceSession) MINCANDIDATE() (*big.Int, error) {
	return _NodesGovernance.Contract.MINCANDIDATE(&_NodesGovernance.CallOpts)
}

// MINCANDIDATE is a free data retrieval call binding the contract method 0x26cc3399.
//
// Solidity: function MIN_CANDIDATE() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCallerSession) MINCANDIDATE() (*big.Int, error) {
	return _NodesGovernance.Contract.MINCANDIDATE(&_NodesGovernance.CallOpts)
}

// VALIDATORPERCANDIDATE is a free data retrieval call binding the contract method 0xa7f32314.
//
// Solidity: function VALIDATOR_PER_CANDIDATE() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCaller) VALIDATORPERCANDIDATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "VALIDATOR_PER_CANDIDATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VALIDATORPERCANDIDATE is a free data retrieval call binding the contract method 0xa7f32314.
//
// Solidity: function VALIDATOR_PER_CANDIDATE() view returns(uint256)
func (_NodesGovernance *NodesGovernanceSession) VALIDATORPERCANDIDATE() (*big.Int, error) {
	return _NodesGovernance.Contract.VALIDATORPERCANDIDATE(&_NodesGovernance.CallOpts)
}

// VALIDATORPERCANDIDATE is a free data retrieval call binding the contract method 0xa7f32314.
//
// Solidity: function VALIDATOR_PER_CANDIDATE() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCallerSession) VALIDATORPERCANDIDATE() (*big.Int, error) {
	return _NodesGovernance.Contract.VALIDATORPERCANDIDATE(&_NodesGovernance.CallOpts)
}

// Allocator is a free data retrieval call binding the contract method 0xaa5dcecc.
//
// Solidity: function allocator() view returns(address)
func (_NodesGovernance *NodesGovernanceCaller) Allocator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "allocator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Allocator is a free data retrieval call binding the contract method 0xaa5dcecc.
//
// Solidity: function allocator() view returns(address)
func (_NodesGovernance *NodesGovernanceSession) Allocator() (common.Address, error) {
	return _NodesGovernance.Contract.Allocator(&_NodesGovernance.CallOpts)
}

// Allocator is a free data retrieval call binding the contract method 0xaa5dcecc.
//
// Solidity: function allocator() view returns(address)
func (_NodesGovernance *NodesGovernanceCallerSession) Allocator() (common.Address, error) {
	return _NodesGovernance.Contract.Allocator(&_NodesGovernance.CallOpts)
}

// At is a free data retrieval call binding the contract method 0xe0886f90.
//
// Solidity: function at(uint256 index) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256,bool) node)
func (_NodesGovernance *NodesGovernanceCaller) At(opts *bind.CallOpts, index *big.Int) (NodesRegistryNode, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "at", index)

	if err != nil {
		return *new(NodesRegistryNode), err
	}

	out0 := *abi.ConvertType(out[0], new(NodesRegistryNode)).(*NodesRegistryNode)

	return out0, err

}

// At is a free data retrieval call binding the contract method 0xe0886f90.
//
// Solidity: function at(uint256 index) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256,bool) node)
func (_NodesGovernance *NodesGovernanceSession) At(index *big.Int) (NodesRegistryNode, error) {
	return _NodesGovernance.Contract.At(&_NodesGovernance.CallOpts, index)
}

// At is a free data retrieval call binding the contract method 0xe0886f90.
//
// Solidity: function at(uint256 index) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256,bool) node)
func (_NodesGovernance *NodesGovernanceCallerSession) At(index *big.Int) (NodesRegistryNode, error) {
	return _NodesGovernance.Contract.At(&_NodesGovernance.CallOpts, index)
}

// CandidatePerRound is a free data retrieval call binding the contract method 0x5ff5c03c.
//
// Solidity: function candidatePerRound(uint256 ) view returns(uint256 numOfNodes, uint256 expectedCompletionTime)
func (_NodesGovernance *NodesGovernanceCaller) CandidatePerRound(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NumOfNodes             *big.Int
	ExpectedCompletionTime *big.Int
}, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "candidatePerRound", arg0)

	outstruct := new(struct {
		NumOfNodes             *big.Int
		ExpectedCompletionTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NumOfNodes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ExpectedCompletionTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CandidatePerRound is a free data retrieval call binding the contract method 0x5ff5c03c.
//
// Solidity: function candidatePerRound(uint256 ) view returns(uint256 numOfNodes, uint256 expectedCompletionTime)
func (_NodesGovernance *NodesGovernanceSession) CandidatePerRound(arg0 *big.Int) (struct {
	NumOfNodes             *big.Int
	ExpectedCompletionTime *big.Int
}, error) {
	return _NodesGovernance.Contract.CandidatePerRound(&_NodesGovernance.CallOpts, arg0)
}

// CandidatePerRound is a free data retrieval call binding the contract method 0x5ff5c03c.
//
// Solidity: function candidatePerRound(uint256 ) view returns(uint256 numOfNodes, uint256 expectedCompletionTime)
func (_NodesGovernance *NodesGovernanceCallerSession) CandidatePerRound(arg0 *big.Int) (struct {
	NumOfNodes             *big.Int
	ExpectedCompletionTime *big.Int
}, error) {
	return _NodesGovernance.Contract.CandidatePerRound(&_NodesGovernance.CallOpts, arg0)
}

// Check is a free data retrieval call binding the contract method 0xc23697a8.
//
// Solidity: function check(address identifier) view returns(bool)
func (_NodesGovernance *NodesGovernanceCaller) Check(opts *bind.CallOpts, identifier common.Address) (bool, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "check", identifier)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Check is a free data retrieval call binding the contract method 0xc23697a8.
//
// Solidity: function check(address identifier) view returns(bool)
func (_NodesGovernance *NodesGovernanceSession) Check(identifier common.Address) (bool, error) {
	return _NodesGovernance.Contract.Check(&_NodesGovernance.CallOpts, identifier)
}

// Check is a free data retrieval call binding the contract method 0xc23697a8.
//
// Solidity: function check(address identifier) view returns(bool)
func (_NodesGovernance *NodesGovernanceCallerSession) Check(identifier common.Address) (bool, error) {
	return _NodesGovernance.Contract.Check(&_NodesGovernance.CallOpts, identifier)
}

// CurrentDetectCircleId is a free data retrieval call binding the contract method 0x191805d8.
//
// Solidity: function currentDetectCircleId() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCaller) CurrentDetectCircleId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "currentDetectCircleId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentDetectCircleId is a free data retrieval call binding the contract method 0x191805d8.
//
// Solidity: function currentDetectCircleId() view returns(uint256)
func (_NodesGovernance *NodesGovernanceSession) CurrentDetectCircleId() (*big.Int, error) {
	return _NodesGovernance.Contract.CurrentDetectCircleId(&_NodesGovernance.CallOpts)
}

// CurrentDetectCircleId is a free data retrieval call binding the contract method 0x191805d8.
//
// Solidity: function currentDetectCircleId() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCallerSession) CurrentDetectCircleId() (*big.Int, error) {
	return _NodesGovernance.Contract.CurrentDetectCircleId(&_NodesGovernance.CallOpts)
}

// CurrentDetectCircleStartTime is a free data retrieval call binding the contract method 0x1bb12a43.
//
// Solidity: function currentDetectCircleStartTime() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCaller) CurrentDetectCircleStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "currentDetectCircleStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentDetectCircleStartTime is a free data retrieval call binding the contract method 0x1bb12a43.
//
// Solidity: function currentDetectCircleStartTime() view returns(uint256)
func (_NodesGovernance *NodesGovernanceSession) CurrentDetectCircleStartTime() (*big.Int, error) {
	return _NodesGovernance.Contract.CurrentDetectCircleStartTime(&_NodesGovernance.CallOpts)
}

// CurrentDetectCircleStartTime is a free data retrieval call binding the contract method 0x1bb12a43.
//
// Solidity: function currentDetectCircleStartTime() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCallerSession) CurrentDetectCircleStartTime() (*big.Int, error) {
	return _NodesGovernance.Contract.CurrentDetectCircleStartTime(&_NodesGovernance.CallOpts)
}

// CurrentRoundId is a free data retrieval call binding the contract method 0x9cbe5efd.
//
// Solidity: function currentRoundId() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCaller) CurrentRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "currentRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRoundId is a free data retrieval call binding the contract method 0x9cbe5efd.
//
// Solidity: function currentRoundId() view returns(uint256)
func (_NodesGovernance *NodesGovernanceSession) CurrentRoundId() (*big.Int, error) {
	return _NodesGovernance.Contract.CurrentRoundId(&_NodesGovernance.CallOpts)
}

// CurrentRoundId is a free data retrieval call binding the contract method 0x9cbe5efd.
//
// Solidity: function currentRoundId() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCallerSession) CurrentRoundId() (*big.Int, error) {
	return _NodesGovernance.Contract.CurrentRoundId(&_NodesGovernance.CallOpts)
}

// CurrentRoundStartTime is a free data retrieval call binding the contract method 0x380dd901.
//
// Solidity: function currentRoundStartTime() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCaller) CurrentRoundStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "currentRoundStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRoundStartTime is a free data retrieval call binding the contract method 0x380dd901.
//
// Solidity: function currentRoundStartTime() view returns(uint256)
func (_NodesGovernance *NodesGovernanceSession) CurrentRoundStartTime() (*big.Int, error) {
	return _NodesGovernance.Contract.CurrentRoundStartTime(&_NodesGovernance.CallOpts)
}

// CurrentRoundStartTime is a free data retrieval call binding the contract method 0x380dd901.
//
// Solidity: function currentRoundStartTime() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCallerSession) CurrentRoundStartTime() (*big.Int, error) {
	return _NodesGovernance.Contract.CurrentRoundStartTime(&_NodesGovernance.CallOpts)
}

// DetectDurationTime is a free data retrieval call binding the contract method 0x49d2b420.
//
// Solidity: function detectDurationTime() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCaller) DetectDurationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "detectDurationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DetectDurationTime is a free data retrieval call binding the contract method 0x49d2b420.
//
// Solidity: function detectDurationTime() view returns(uint256)
func (_NodesGovernance *NodesGovernanceSession) DetectDurationTime() (*big.Int, error) {
	return _NodesGovernance.Contract.DetectDurationTime(&_NodesGovernance.CallOpts)
}

// DetectDurationTime is a free data retrieval call binding the contract method 0x49d2b420.
//
// Solidity: function detectDurationTime() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCallerSession) DetectDurationTime() (*big.Int, error) {
	return _NodesGovernance.Contract.DetectDurationTime(&_NodesGovernance.CallOpts)
}

// DetectPeriods is a free data retrieval call binding the contract method 0x698083af.
//
// Solidity: function detectPeriods(uint256 ) view returns(uint256 startId, uint256 endId, uint256 startTime, uint256 endTime)
func (_NodesGovernance *NodesGovernanceCaller) DetectPeriods(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartId   *big.Int
	EndId     *big.Int
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "detectPeriods", arg0)

	outstruct := new(struct {
		StartId   *big.Int
		EndId     *big.Int
		StartTime *big.Int
		EndTime   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DetectPeriods is a free data retrieval call binding the contract method 0x698083af.
//
// Solidity: function detectPeriods(uint256 ) view returns(uint256 startId, uint256 endId, uint256 startTime, uint256 endTime)
func (_NodesGovernance *NodesGovernanceSession) DetectPeriods(arg0 *big.Int) (struct {
	StartId   *big.Int
	EndId     *big.Int
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _NodesGovernance.Contract.DetectPeriods(&_NodesGovernance.CallOpts, arg0)
}

// DetectPeriods is a free data retrieval call binding the contract method 0x698083af.
//
// Solidity: function detectPeriods(uint256 ) view returns(uint256 startId, uint256 endId, uint256 startTime, uint256 endTime)
func (_NodesGovernance *NodesGovernanceCallerSession) DetectPeriods(arg0 *big.Int) (struct {
	StartId   *big.Int
	EndId     *big.Int
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _NodesGovernance.Contract.DetectPeriods(&_NodesGovernance.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address identifier) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256,bool) node)
func (_NodesGovernance *NodesGovernanceCaller) Get(opts *bind.CallOpts, identifier common.Address) (NodesRegistryNode, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "get", identifier)

	if err != nil {
		return *new(NodesRegistryNode), err
	}

	out0 := *abi.ConvertType(out[0], new(NodesRegistryNode)).(*NodesRegistryNode)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address identifier) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256,bool) node)
func (_NodesGovernance *NodesGovernanceSession) Get(identifier common.Address) (NodesRegistryNode, error) {
	return _NodesGovernance.Contract.Get(&_NodesGovernance.CallOpts, identifier)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address identifier) view returns((address,string,uint256,bool,(string,uint256,uint256)[],address,uint256,bool) node)
func (_NodesGovernance *NodesGovernanceCallerSession) Get(identifier common.Address) (NodesRegistryNode, error) {
	return _NodesGovernance.Contract.Get(&_NodesGovernance.CallOpts, identifier)
}

// GetAttach is a free data retrieval call binding the contract method 0xc7edca7a.
//
// Solidity: function getAttach(address provider) view returns(address[])
func (_NodesGovernance *NodesGovernanceCaller) GetAttach(opts *bind.CallOpts, provider common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "getAttach", provider)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAttach is a free data retrieval call binding the contract method 0xc7edca7a.
//
// Solidity: function getAttach(address provider) view returns(address[])
func (_NodesGovernance *NodesGovernanceSession) GetAttach(provider common.Address) ([]common.Address, error) {
	return _NodesGovernance.Contract.GetAttach(&_NodesGovernance.CallOpts, provider)
}

// GetAttach is a free data retrieval call binding the contract method 0xc7edca7a.
//
// Solidity: function getAttach(address provider) view returns(address[])
func (_NodesGovernance *NodesGovernanceCallerSession) GetAttach(provider common.Address) ([]common.Address, error) {
	return _NodesGovernance.Contract.GetAttach(&_NodesGovernance.CallOpts, provider)
}

// GetOnePeriodSettlement is a free data retrieval call binding the contract method 0x6300c951.
//
// Solidity: function getOnePeriodSettlement(uint256 detectPeriodId) view returns((uint64,uint64,uint128,address,address)[] states, uint256 totalQuotas)
func (_NodesGovernance *NodesGovernanceCaller) GetOnePeriodSettlement(opts *bind.CallOpts, detectPeriodId *big.Int) (struct {
	States      []NodeState
	TotalQuotas *big.Int
}, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "getOnePeriodSettlement", detectPeriodId)

	outstruct := new(struct {
		States      []NodeState
		TotalQuotas *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.States = *abi.ConvertType(out[0], new([]NodeState)).(*[]NodeState)
	outstruct.TotalQuotas = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOnePeriodSettlement is a free data retrieval call binding the contract method 0x6300c951.
//
// Solidity: function getOnePeriodSettlement(uint256 detectPeriodId) view returns((uint64,uint64,uint128,address,address)[] states, uint256 totalQuotas)
func (_NodesGovernance *NodesGovernanceSession) GetOnePeriodSettlement(detectPeriodId *big.Int) (struct {
	States      []NodeState
	TotalQuotas *big.Int
}, error) {
	return _NodesGovernance.Contract.GetOnePeriodSettlement(&_NodesGovernance.CallOpts, detectPeriodId)
}

// GetOnePeriodSettlement is a free data retrieval call binding the contract method 0x6300c951.
//
// Solidity: function getOnePeriodSettlement(uint256 detectPeriodId) view returns((uint64,uint64,uint128,address,address)[] states, uint256 totalQuotas)
func (_NodesGovernance *NodesGovernanceCallerSession) GetOnePeriodSettlement(detectPeriodId *big.Int) (struct {
	States      []NodeState
	TotalQuotas *big.Int
}, error) {
	return _NodesGovernance.Contract.GetOnePeriodSettlement(&_NodesGovernance.CallOpts, detectPeriodId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NodesGovernance *NodesGovernanceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NodesGovernance *NodesGovernanceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NodesGovernance.Contract.GetRoleAdmin(&_NodesGovernance.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NodesGovernance *NodesGovernanceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NodesGovernance.Contract.GetRoleAdmin(&_NodesGovernance.CallOpts, role)
}

// GetRoundCandidates is a free data retrieval call binding the contract method 0x93e9d413.
//
// Solidity: function getRoundCandidates(uint256 roundId) view returns(address[] candidates)
func (_NodesGovernance *NodesGovernanceCaller) GetRoundCandidates(opts *bind.CallOpts, roundId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "getRoundCandidates", roundId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoundCandidates is a free data retrieval call binding the contract method 0x93e9d413.
//
// Solidity: function getRoundCandidates(uint256 roundId) view returns(address[] candidates)
func (_NodesGovernance *NodesGovernanceSession) GetRoundCandidates(roundId *big.Int) ([]common.Address, error) {
	return _NodesGovernance.Contract.GetRoundCandidates(&_NodesGovernance.CallOpts, roundId)
}

// GetRoundCandidates is a free data retrieval call binding the contract method 0x93e9d413.
//
// Solidity: function getRoundCandidates(uint256 roundId) view returns(address[] candidates)
func (_NodesGovernance *NodesGovernanceCallerSession) GetRoundCandidates(roundId *big.Int) ([]common.Address, error) {
	return _NodesGovernance.Contract.GetRoundCandidates(&_NodesGovernance.CallOpts, roundId)
}

// GetValidatorsOfCandidate is a free data retrieval call binding the contract method 0x0732bd7e.
//
// Solidity: function getValidatorsOfCandidate(uint256 roundId, address candidate) view returns(address[] validators)
func (_NodesGovernance *NodesGovernanceCaller) GetValidatorsOfCandidate(opts *bind.CallOpts, roundId *big.Int, candidate common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "getValidatorsOfCandidate", roundId, candidate)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorsOfCandidate is a free data retrieval call binding the contract method 0x0732bd7e.
//
// Solidity: function getValidatorsOfCandidate(uint256 roundId, address candidate) view returns(address[] validators)
func (_NodesGovernance *NodesGovernanceSession) GetValidatorsOfCandidate(roundId *big.Int, candidate common.Address) ([]common.Address, error) {
	return _NodesGovernance.Contract.GetValidatorsOfCandidate(&_NodesGovernance.CallOpts, roundId, candidate)
}

// GetValidatorsOfCandidate is a free data retrieval call binding the contract method 0x0732bd7e.
//
// Solidity: function getValidatorsOfCandidate(uint256 roundId, address candidate) view returns(address[] validators)
func (_NodesGovernance *NodesGovernanceCallerSession) GetValidatorsOfCandidate(roundId *big.Int, candidate common.Address) ([]common.Address, error) {
	return _NodesGovernance.Contract.GetValidatorsOfCandidate(&_NodesGovernance.CallOpts, roundId, candidate)
}

// GpuSummary is a free data retrieval call binding the contract method 0xf67c5bdc.
//
// Solidity: function gpuSummary(string ) view returns(string gpuType, uint256 totalNum, uint256 used)
func (_NodesGovernance *NodesGovernanceCaller) GpuSummary(opts *bind.CallOpts, arg0 string) (struct {
	GpuType  string
	TotalNum *big.Int
	Used     *big.Int
}, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "gpuSummary", arg0)

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
func (_NodesGovernance *NodesGovernanceSession) GpuSummary(arg0 string) (struct {
	GpuType  string
	TotalNum *big.Int
	Used     *big.Int
}, error) {
	return _NodesGovernance.Contract.GpuSummary(&_NodesGovernance.CallOpts, arg0)
}

// GpuSummary is a free data retrieval call binding the contract method 0xf67c5bdc.
//
// Solidity: function gpuSummary(string ) view returns(string gpuType, uint256 totalNum, uint256 used)
func (_NodesGovernance *NodesGovernanceCallerSession) GpuSummary(arg0 string) (struct {
	GpuType  string
	TotalNum *big.Int
	Used     *big.Int
}, error) {
	return _NodesGovernance.Contract.GpuSummary(&_NodesGovernance.CallOpts, arg0)
}

// GpuTypeOfNodes is a free data retrieval call binding the contract method 0xed38ed0d.
//
// Solidity: function gpuTypeOfNodes(address , string ) view returns(uint256)
func (_NodesGovernance *NodesGovernanceCaller) GpuTypeOfNodes(opts *bind.CallOpts, arg0 common.Address, arg1 string) (*big.Int, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "gpuTypeOfNodes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GpuTypeOfNodes is a free data retrieval call binding the contract method 0xed38ed0d.
//
// Solidity: function gpuTypeOfNodes(address , string ) view returns(uint256)
func (_NodesGovernance *NodesGovernanceSession) GpuTypeOfNodes(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _NodesGovernance.Contract.GpuTypeOfNodes(&_NodesGovernance.CallOpts, arg0, arg1)
}

// GpuTypeOfNodes is a free data retrieval call binding the contract method 0xed38ed0d.
//
// Solidity: function gpuTypeOfNodes(address , string ) view returns(uint256)
func (_NodesGovernance *NodesGovernanceCallerSession) GpuTypeOfNodes(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _NodesGovernance.Contract.GpuTypeOfNodes(&_NodesGovernance.CallOpts, arg0, arg1)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NodesGovernance *NodesGovernanceCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NodesGovernance *NodesGovernanceSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NodesGovernance.Contract.HasRole(&_NodesGovernance.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NodesGovernance *NodesGovernanceCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NodesGovernance.Contract.HasRole(&_NodesGovernance.CallOpts, role, account)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCaller) Length(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "length")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_NodesGovernance *NodesGovernanceSession) Length() (*big.Int, error) {
	return _NodesGovernance.Contract.Length(&_NodesGovernance.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCallerSession) Length() (*big.Int, error) {
	return _NodesGovernance.Contract.Length(&_NodesGovernance.CallOpts)
}

// ProxyNodes is a free data retrieval call binding the contract method 0xb8afa39c.
//
// Solidity: function proxyNodes(address ) view returns(bool)
func (_NodesGovernance *NodesGovernanceCaller) ProxyNodes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "proxyNodes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProxyNodes is a free data retrieval call binding the contract method 0xb8afa39c.
//
// Solidity: function proxyNodes(address ) view returns(bool)
func (_NodesGovernance *NodesGovernanceSession) ProxyNodes(arg0 common.Address) (bool, error) {
	return _NodesGovernance.Contract.ProxyNodes(&_NodesGovernance.CallOpts, arg0)
}

// ProxyNodes is a free data retrieval call binding the contract method 0xb8afa39c.
//
// Solidity: function proxyNodes(address ) view returns(bool)
func (_NodesGovernance *NodesGovernanceCallerSession) ProxyNodes(arg0 common.Address) (bool, error) {
	return _NodesGovernance.Contract.ProxyNodes(&_NodesGovernance.CallOpts, arg0)
}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_NodesGovernance *NodesGovernanceCaller) RenounceRole(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) error {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "renounceRole", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_NodesGovernance *NodesGovernanceSession) RenounceRole(arg0 [32]byte, arg1 common.Address) error {
	return _NodesGovernance.Contract.RenounceRole(&_NodesGovernance.CallOpts, arg0, arg1)
}

// RenounceRole is a free data retrieval call binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 , address ) pure returns()
func (_NodesGovernance *NodesGovernanceCallerSession) RenounceRole(arg0 [32]byte, arg1 common.Address) error {
	return _NodesGovernance.Contract.RenounceRole(&_NodesGovernance.CallOpts, arg0, arg1)
}

// RoundDurationTime is a free data retrieval call binding the contract method 0xde137270.
//
// Solidity: function roundDurationTime() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCaller) RoundDurationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "roundDurationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundDurationTime is a free data retrieval call binding the contract method 0xde137270.
//
// Solidity: function roundDurationTime() view returns(uint256)
func (_NodesGovernance *NodesGovernanceSession) RoundDurationTime() (*big.Int, error) {
	return _NodesGovernance.Contract.RoundDurationTime(&_NodesGovernance.CallOpts)
}

// RoundDurationTime is a free data retrieval call binding the contract method 0xde137270.
//
// Solidity: function roundDurationTime() view returns(uint256)
func (_NodesGovernance *NodesGovernanceCallerSession) RoundDurationTime() (*big.Int, error) {
	return _NodesGovernance.Contract.RoundDurationTime(&_NodesGovernance.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_NodesGovernance *NodesGovernanceCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_NodesGovernance *NodesGovernanceSession) StakeToken() (common.Address, error) {
	return _NodesGovernance.Contract.StakeToken(&_NodesGovernance.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_NodesGovernance *NodesGovernanceCallerSession) StakeToken() (common.Address, error) {
	return _NodesGovernance.Contract.StakeToken(&_NodesGovernance.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NodesGovernance *NodesGovernanceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NodesGovernance *NodesGovernanceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NodesGovernance.Contract.SupportsInterface(&_NodesGovernance.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NodesGovernance *NodesGovernanceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NodesGovernance.Contract.SupportsInterface(&_NodesGovernance.CallOpts, interfaceId)
}

// ValidatorsPerCandidate is a free data retrieval call binding the contract method 0x18767f31.
//
// Solidity: function validatorsPerCandidate(uint256 , address , uint256 ) view returns(address)
func (_NodesGovernance *NodesGovernanceCaller) ValidatorsPerCandidate(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "validatorsPerCandidate", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorsPerCandidate is a free data retrieval call binding the contract method 0x18767f31.
//
// Solidity: function validatorsPerCandidate(uint256 , address , uint256 ) view returns(address)
func (_NodesGovernance *NodesGovernanceSession) ValidatorsPerCandidate(arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _NodesGovernance.Contract.ValidatorsPerCandidate(&_NodesGovernance.CallOpts, arg0, arg1, arg2)
}

// ValidatorsPerCandidate is a free data retrieval call binding the contract method 0x18767f31.
//
// Solidity: function validatorsPerCandidate(uint256 , address , uint256 ) view returns(address)
func (_NodesGovernance *NodesGovernanceCallerSession) ValidatorsPerCandidate(arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _NodesGovernance.Contract.ValidatorsPerCandidate(&_NodesGovernance.CallOpts, arg0, arg1, arg2)
}

// VotedPerCandidate is a free data retrieval call binding the contract method 0x2ad4f0ad.
//
// Solidity: function votedPerCandidate(uint256 , address ) view returns(address candidate, uint128 yesVotes, uint128 noVotes, bool completed)
func (_NodesGovernance *NodesGovernanceCaller) VotedPerCandidate(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Candidate common.Address
	YesVotes  *big.Int
	NoVotes   *big.Int
	Completed bool
}, error) {
	var out []interface{}
	err := _NodesGovernance.contract.Call(opts, &out, "votedPerCandidate", arg0, arg1)

	outstruct := new(struct {
		Candidate common.Address
		YesVotes  *big.Int
		NoVotes   *big.Int
		Completed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Candidate = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.YesVotes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NoVotes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Completed = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// VotedPerCandidate is a free data retrieval call binding the contract method 0x2ad4f0ad.
//
// Solidity: function votedPerCandidate(uint256 , address ) view returns(address candidate, uint128 yesVotes, uint128 noVotes, bool completed)
func (_NodesGovernance *NodesGovernanceSession) VotedPerCandidate(arg0 *big.Int, arg1 common.Address) (struct {
	Candidate common.Address
	YesVotes  *big.Int
	NoVotes   *big.Int
	Completed bool
}, error) {
	return _NodesGovernance.Contract.VotedPerCandidate(&_NodesGovernance.CallOpts, arg0, arg1)
}

// VotedPerCandidate is a free data retrieval call binding the contract method 0x2ad4f0ad.
//
// Solidity: function votedPerCandidate(uint256 , address ) view returns(address candidate, uint128 yesVotes, uint128 noVotes, bool completed)
func (_NodesGovernance *NodesGovernanceCallerSession) VotedPerCandidate(arg0 *big.Int, arg1 common.Address) (struct {
	Candidate common.Address
	YesVotes  *big.Int
	NoVotes   *big.Int
	Completed bool
}, error) {
	return _NodesGovernance.Contract.VotedPerCandidate(&_NodesGovernance.CallOpts, arg0, arg1)
}

// AllocGPU is a paid mutator transaction binding the contract method 0x6252e1c2.
//
// Solidity: function allocGPU(uint256 startIndex, string[] gpuTypes, uint256[] gpuNums) returns((address,string,uint256)[] gpuNodes, uint256 len)
func (_NodesGovernance *NodesGovernanceTransactor) AllocGPU(opts *bind.TransactOpts, startIndex *big.Int, gpuTypes []string, gpuNums []*big.Int) (*types.Transaction, error) {
	return _NodesGovernance.contract.Transact(opts, "allocGPU", startIndex, gpuTypes, gpuNums)
}

// AllocGPU is a paid mutator transaction binding the contract method 0x6252e1c2.
//
// Solidity: function allocGPU(uint256 startIndex, string[] gpuTypes, uint256[] gpuNums) returns((address,string,uint256)[] gpuNodes, uint256 len)
func (_NodesGovernance *NodesGovernanceSession) AllocGPU(startIndex *big.Int, gpuTypes []string, gpuNums []*big.Int) (*types.Transaction, error) {
	return _NodesGovernance.Contract.AllocGPU(&_NodesGovernance.TransactOpts, startIndex, gpuTypes, gpuNums)
}

// AllocGPU is a paid mutator transaction binding the contract method 0x6252e1c2.
//
// Solidity: function allocGPU(uint256 startIndex, string[] gpuTypes, uint256[] gpuNums) returns((address,string,uint256)[] gpuNodes, uint256 len)
func (_NodesGovernance *NodesGovernanceTransactorSession) AllocGPU(startIndex *big.Int, gpuTypes []string, gpuNums []*big.Int) (*types.Transaction, error) {
	return _NodesGovernance.Contract.AllocGPU(&_NodesGovernance.TransactOpts, startIndex, gpuTypes, gpuNums)
}

// Attach is a paid mutator transaction binding the contract method 0x7a0ca1e2.
//
// Solidity: function attach(address server) returns()
func (_NodesGovernance *NodesGovernanceTransactor) Attach(opts *bind.TransactOpts, server common.Address) (*types.Transaction, error) {
	return _NodesGovernance.contract.Transact(opts, "attach", server)
}

// Attach is a paid mutator transaction binding the contract method 0x7a0ca1e2.
//
// Solidity: function attach(address server) returns()
func (_NodesGovernance *NodesGovernanceSession) Attach(server common.Address) (*types.Transaction, error) {
	return _NodesGovernance.Contract.Attach(&_NodesGovernance.TransactOpts, server)
}

// Attach is a paid mutator transaction binding the contract method 0x7a0ca1e2.
//
// Solidity: function attach(address server) returns()
func (_NodesGovernance *NodesGovernanceTransactorSession) Attach(server common.Address) (*types.Transaction, error) {
	return _NodesGovernance.Contract.Attach(&_NodesGovernance.TransactOpts, server)
}

// DeregisterNode is a paid mutator transaction binding the contract method 0x18b1c081.
//
// Solidity: function deregisterNode() returns()
func (_NodesGovernance *NodesGovernanceTransactor) DeregisterNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodesGovernance.contract.Transact(opts, "deregisterNode")
}

// DeregisterNode is a paid mutator transaction binding the contract method 0x18b1c081.
//
// Solidity: function deregisterNode() returns()
func (_NodesGovernance *NodesGovernanceSession) DeregisterNode() (*types.Transaction, error) {
	return _NodesGovernance.Contract.DeregisterNode(&_NodesGovernance.TransactOpts)
}

// DeregisterNode is a paid mutator transaction binding the contract method 0x18b1c081.
//
// Solidity: function deregisterNode() returns()
func (_NodesGovernance *NodesGovernanceTransactorSession) DeregisterNode() (*types.Transaction, error) {
	return _NodesGovernance.Contract.DeregisterNode(&_NodesGovernance.TransactOpts)
}

// Dettach is a paid mutator transaction binding the contract method 0x0e22e7f8.
//
// Solidity: function dettach(address server) returns()
func (_NodesGovernance *NodesGovernanceTransactor) Dettach(opts *bind.TransactOpts, server common.Address) (*types.Transaction, error) {
	return _NodesGovernance.contract.Transact(opts, "dettach", server)
}

// Dettach is a paid mutator transaction binding the contract method 0x0e22e7f8.
//
// Solidity: function dettach(address server) returns()
func (_NodesGovernance *NodesGovernanceSession) Dettach(server common.Address) (*types.Transaction, error) {
	return _NodesGovernance.Contract.Dettach(&_NodesGovernance.TransactOpts, server)
}

// Dettach is a paid mutator transaction binding the contract method 0x0e22e7f8.
//
// Solidity: function dettach(address server) returns()
func (_NodesGovernance *NodesGovernanceTransactorSession) Dettach(server common.Address) (*types.Transaction, error) {
	return _NodesGovernance.Contract.Dettach(&_NodesGovernance.TransactOpts, server)
}

// FreeGPU is a paid mutator transaction binding the contract method 0x036fe9c2.
//
// Solidity: function freeGPU((address,string,uint256)[] gpuNodes) returns()
func (_NodesGovernance *NodesGovernanceTransactor) FreeGPU(opts *bind.TransactOpts, gpuNodes []NodeComputeUsed) (*types.Transaction, error) {
	return _NodesGovernance.contract.Transact(opts, "freeGPU", gpuNodes)
}

// FreeGPU is a paid mutator transaction binding the contract method 0x036fe9c2.
//
// Solidity: function freeGPU((address,string,uint256)[] gpuNodes) returns()
func (_NodesGovernance *NodesGovernanceSession) FreeGPU(gpuNodes []NodeComputeUsed) (*types.Transaction, error) {
	return _NodesGovernance.Contract.FreeGPU(&_NodesGovernance.TransactOpts, gpuNodes)
}

// FreeGPU is a paid mutator transaction binding the contract method 0x036fe9c2.
//
// Solidity: function freeGPU((address,string,uint256)[] gpuNodes) returns()
func (_NodesGovernance *NodesGovernanceTransactorSession) FreeGPU(gpuNodes []NodeComputeUsed) (*types.Transaction, error) {
	return _NodesGovernance.Contract.FreeGPU(&_NodesGovernance.TransactOpts, gpuNodes)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NodesGovernance *NodesGovernanceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodesGovernance.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NodesGovernance *NodesGovernanceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodesGovernance.Contract.GrantRole(&_NodesGovernance.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NodesGovernance *NodesGovernanceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodesGovernance.Contract.GrantRole(&_NodesGovernance.TransactOpts, role, account)
}

// NodesGovernanceInitialize is a paid mutator transaction binding the contract method 0xefa219e0.
//
// Solidity: function nodesGovernance_initialize((address,string,address,string[],uint256[])[] _nodesInfos, address _allocator, uint256 _roundDurationTime, address _stakeToken) returns()
func (_NodesGovernance *NodesGovernanceTransactor) NodesGovernanceInitialize(opts *bind.TransactOpts, _nodesInfos []NodeInfo, _allocator common.Address, _roundDurationTime *big.Int, _stakeToken common.Address) (*types.Transaction, error) {
	return _NodesGovernance.contract.Transact(opts, "nodesGovernance_initialize", _nodesInfos, _allocator, _roundDurationTime, _stakeToken)
}

// NodesGovernanceInitialize is a paid mutator transaction binding the contract method 0xefa219e0.
//
// Solidity: function nodesGovernance_initialize((address,string,address,string[],uint256[])[] _nodesInfos, address _allocator, uint256 _roundDurationTime, address _stakeToken) returns()
func (_NodesGovernance *NodesGovernanceSession) NodesGovernanceInitialize(_nodesInfos []NodeInfo, _allocator common.Address, _roundDurationTime *big.Int, _stakeToken common.Address) (*types.Transaction, error) {
	return _NodesGovernance.Contract.NodesGovernanceInitialize(&_NodesGovernance.TransactOpts, _nodesInfos, _allocator, _roundDurationTime, _stakeToken)
}

// NodesGovernanceInitialize is a paid mutator transaction binding the contract method 0xefa219e0.
//
// Solidity: function nodesGovernance_initialize((address,string,address,string[],uint256[])[] _nodesInfos, address _allocator, uint256 _roundDurationTime, address _stakeToken) returns()
func (_NodesGovernance *NodesGovernanceTransactorSession) NodesGovernanceInitialize(_nodesInfos []NodeInfo, _allocator common.Address, _roundDurationTime *big.Int, _stakeToken common.Address) (*types.Transaction, error) {
	return _NodesGovernance.Contract.NodesGovernanceInitialize(&_NodesGovernance.TransactOpts, _nodesInfos, _allocator, _roundDurationTime, _stakeToken)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xff7b178d.
//
// Solidity: function registerNode(address wallet, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums, bool isPublic) payable returns()
func (_NodesGovernance *NodesGovernanceTransactor) RegisterNode(opts *bind.TransactOpts, wallet common.Address, aliasIdentifier string, gpuTypes []string, gpuNums []*big.Int, isPublic bool) (*types.Transaction, error) {
	return _NodesGovernance.contract.Transact(opts, "registerNode", wallet, aliasIdentifier, gpuTypes, gpuNums, isPublic)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xff7b178d.
//
// Solidity: function registerNode(address wallet, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums, bool isPublic) payable returns()
func (_NodesGovernance *NodesGovernanceSession) RegisterNode(wallet common.Address, aliasIdentifier string, gpuTypes []string, gpuNums []*big.Int, isPublic bool) (*types.Transaction, error) {
	return _NodesGovernance.Contract.RegisterNode(&_NodesGovernance.TransactOpts, wallet, aliasIdentifier, gpuTypes, gpuNums, isPublic)
}

// RegisterNode is a paid mutator transaction binding the contract method 0xff7b178d.
//
// Solidity: function registerNode(address wallet, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums, bool isPublic) payable returns()
func (_NodesGovernance *NodesGovernanceTransactorSession) RegisterNode(wallet common.Address, aliasIdentifier string, gpuTypes []string, gpuNums []*big.Int, isPublic bool) (*types.Transaction, error) {
	return _NodesGovernance.Contract.RegisterNode(&_NodesGovernance.TransactOpts, wallet, aliasIdentifier, gpuTypes, gpuNums, isPublic)
}

// RegisterProxyNode is a paid mutator transaction binding the contract method 0xd8fd0eb1.
//
// Solidity: function registerProxyNode(address proxy) returns()
func (_NodesGovernance *NodesGovernanceTransactor) RegisterProxyNode(opts *bind.TransactOpts, proxy common.Address) (*types.Transaction, error) {
	return _NodesGovernance.contract.Transact(opts, "registerProxyNode", proxy)
}

// RegisterProxyNode is a paid mutator transaction binding the contract method 0xd8fd0eb1.
//
// Solidity: function registerProxyNode(address proxy) returns()
func (_NodesGovernance *NodesGovernanceSession) RegisterProxyNode(proxy common.Address) (*types.Transaction, error) {
	return _NodesGovernance.Contract.RegisterProxyNode(&_NodesGovernance.TransactOpts, proxy)
}

// RegisterProxyNode is a paid mutator transaction binding the contract method 0xd8fd0eb1.
//
// Solidity: function registerProxyNode(address proxy) returns()
func (_NodesGovernance *NodesGovernanceTransactorSession) RegisterProxyNode(proxy common.Address) (*types.Transaction, error) {
	return _NodesGovernance.Contract.RegisterProxyNode(&_NodesGovernance.TransactOpts, proxy)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NodesGovernance *NodesGovernanceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodesGovernance.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NodesGovernance *NodesGovernanceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodesGovernance.Contract.RevokeRole(&_NodesGovernance.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NodesGovernance *NodesGovernanceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodesGovernance.Contract.RevokeRole(&_NodesGovernance.TransactOpts, role, account)
}

// SettlementOnePeriod is a paid mutator transaction binding the contract method 0x4d4fc0b8.
//
// Solidity: function settlementOnePeriod(uint256 detectPeriodId) returns((uint64,uint64,uint128,address,address)[] states, uint256 totalQuotas)
func (_NodesGovernance *NodesGovernanceTransactor) SettlementOnePeriod(opts *bind.TransactOpts, detectPeriodId *big.Int) (*types.Transaction, error) {
	return _NodesGovernance.contract.Transact(opts, "settlementOnePeriod", detectPeriodId)
}

// SettlementOnePeriod is a paid mutator transaction binding the contract method 0x4d4fc0b8.
//
// Solidity: function settlementOnePeriod(uint256 detectPeriodId) returns((uint64,uint64,uint128,address,address)[] states, uint256 totalQuotas)
func (_NodesGovernance *NodesGovernanceSession) SettlementOnePeriod(detectPeriodId *big.Int) (*types.Transaction, error) {
	return _NodesGovernance.Contract.SettlementOnePeriod(&_NodesGovernance.TransactOpts, detectPeriodId)
}

// SettlementOnePeriod is a paid mutator transaction binding the contract method 0x4d4fc0b8.
//
// Solidity: function settlementOnePeriod(uint256 detectPeriodId) returns((uint64,uint64,uint128,address,address)[] states, uint256 totalQuotas)
func (_NodesGovernance *NodesGovernanceTransactorSession) SettlementOnePeriod(detectPeriodId *big.Int) (*types.Transaction, error) {
	return _NodesGovernance.Contract.SettlementOnePeriod(&_NodesGovernance.TransactOpts, detectPeriodId)
}

// StartNewValidationRound is a paid mutator transaction binding the contract method 0x63c94199.
//
// Solidity: function startNewValidationRound() returns(uint256 detectId, uint256 roundId)
func (_NodesGovernance *NodesGovernanceTransactor) StartNewValidationRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodesGovernance.contract.Transact(opts, "startNewValidationRound")
}

// StartNewValidationRound is a paid mutator transaction binding the contract method 0x63c94199.
//
// Solidity: function startNewValidationRound() returns(uint256 detectId, uint256 roundId)
func (_NodesGovernance *NodesGovernanceSession) StartNewValidationRound() (*types.Transaction, error) {
	return _NodesGovernance.Contract.StartNewValidationRound(&_NodesGovernance.TransactOpts)
}

// StartNewValidationRound is a paid mutator transaction binding the contract method 0x63c94199.
//
// Solidity: function startNewValidationRound() returns(uint256 detectId, uint256 roundId)
func (_NodesGovernance *NodesGovernanceTransactorSession) StartNewValidationRound() (*types.Transaction, error) {
	return _NodesGovernance.Contract.StartNewValidationRound(&_NodesGovernance.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 roundId, address candidate, bool result) returns()
func (_NodesGovernance *NodesGovernanceTransactor) Vote(opts *bind.TransactOpts, roundId *big.Int, candidate common.Address, result bool) (*types.Transaction, error) {
	return _NodesGovernance.contract.Transact(opts, "vote", roundId, candidate, result)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 roundId, address candidate, bool result) returns()
func (_NodesGovernance *NodesGovernanceSession) Vote(roundId *big.Int, candidate common.Address, result bool) (*types.Transaction, error) {
	return _NodesGovernance.Contract.Vote(&_NodesGovernance.TransactOpts, roundId, candidate, result)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 roundId, address candidate, bool result) returns()
func (_NodesGovernance *NodesGovernanceTransactorSession) Vote(roundId *big.Int, candidate common.Address, result bool) (*types.Transaction, error) {
	return _NodesGovernance.Contract.Vote(&_NodesGovernance.TransactOpts, roundId, candidate, result)
}

// NodesGovernanceAuthorizedIterator is returned from FilterAuthorized and is used to iterate over the raw logs and unpacked data for Authorized events raised by the NodesGovernance contract.
type NodesGovernanceAuthorizedIterator struct {
	Event *NodesGovernanceAuthorized // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceAuthorized)
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
		it.Event = new(NodesGovernanceAuthorized)
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
func (it *NodesGovernanceAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceAuthorized represents a Authorized event raised by the NodesGovernance contract.
type NodesGovernanceAuthorized struct {
	Owner   common.Address
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuthorized is a free log retrieval operation binding the contract event 0xf5a7f4fb8a92356e8c8c4ae7ac3589908381450500a7e2fd08c95600021ee889.
//
// Solidity: event Authorized(address indexed owner, address indexed spender)
func (_NodesGovernance *NodesGovernanceFilterer) FilterAuthorized(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NodesGovernanceAuthorizedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "Authorized", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceAuthorizedIterator{contract: _NodesGovernance.contract, event: "Authorized", logs: logs, sub: sub}, nil
}

// WatchAuthorized is a free log subscription operation binding the contract event 0xf5a7f4fb8a92356e8c8c4ae7ac3589908381450500a7e2fd08c95600021ee889.
//
// Solidity: event Authorized(address indexed owner, address indexed spender)
func (_NodesGovernance *NodesGovernanceFilterer) WatchAuthorized(opts *bind.WatchOpts, sink chan<- *NodesGovernanceAuthorized, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "Authorized", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceAuthorized)
				if err := _NodesGovernance.contract.UnpackLog(event, "Authorized", log); err != nil {
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
func (_NodesGovernance *NodesGovernanceFilterer) ParseAuthorized(log types.Log) (*NodesGovernanceAuthorized, error) {
	event := new(NodesGovernanceAuthorized)
	if err := _NodesGovernance.contract.UnpackLog(event, "Authorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesGovernanceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NodesGovernance contract.
type NodesGovernanceInitializedIterator struct {
	Event *NodesGovernanceInitialized // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceInitialized)
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
		it.Event = new(NodesGovernanceInitialized)
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
func (it *NodesGovernanceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceInitialized represents a Initialized event raised by the NodesGovernance contract.
type NodesGovernanceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodesGovernance *NodesGovernanceFilterer) FilterInitialized(opts *bind.FilterOpts) (*NodesGovernanceInitializedIterator, error) {

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceInitializedIterator{contract: _NodesGovernance.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodesGovernance *NodesGovernanceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodesGovernanceInitialized) (event.Subscription, error) {

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceInitialized)
				if err := _NodesGovernance.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NodesGovernance *NodesGovernanceFilterer) ParseInitialized(log types.Log) (*NodesGovernanceInitialized, error) {
	event := new(NodesGovernanceInitialized)
	if err := _NodesGovernance.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesGovernanceNodeActivedIterator is returned from FilterNodeActived and is used to iterate over the raw logs and unpacked data for NodeActived events raised by the NodesGovernance contract.
type NodesGovernanceNodeActivedIterator struct {
	Event *NodesGovernanceNodeActived // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceNodeActivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceNodeActived)
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
		it.Event = new(NodesGovernanceNodeActived)
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
func (it *NodesGovernanceNodeActivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceNodeActivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceNodeActived represents a NodeActived event raised by the NodesGovernance contract.
type NodesGovernanceNodeActived struct {
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
func (_NodesGovernance *NodesGovernanceFilterer) FilterNodeActived(opts *bind.FilterOpts, wallet []common.Address) (*NodesGovernanceNodeActivedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "NodeActived", walletRule)
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceNodeActivedIterator{contract: _NodesGovernance.contract, event: "NodeActived", logs: logs, sub: sub}, nil
}

// WatchNodeActived is a free log subscription operation binding the contract event 0xb0a73f160683fa6bd1601bd25cd9e2addc090f0a74959bf51c6a2a4560af6f5f.
//
// Solidity: event NodeActived(address indexed wallet, address identifier, uint256 time, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums)
func (_NodesGovernance *NodesGovernanceFilterer) WatchNodeActived(opts *bind.WatchOpts, sink chan<- *NodesGovernanceNodeActived, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "NodeActived", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceNodeActived)
				if err := _NodesGovernance.contract.UnpackLog(event, "NodeActived", log); err != nil {
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
func (_NodesGovernance *NodesGovernanceFilterer) ParseNodeActived(log types.Log) (*NodesGovernanceNodeActived, error) {
	event := new(NodesGovernanceNodeActived)
	if err := _NodesGovernance.contract.UnpackLog(event, "NodeActived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesGovernanceNodeAttachedIterator is returned from FilterNodeAttached and is used to iterate over the raw logs and unpacked data for NodeAttached events raised by the NodesGovernance contract.
type NodesGovernanceNodeAttachedIterator struct {
	Event *NodesGovernanceNodeAttached // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceNodeAttachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceNodeAttached)
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
		it.Event = new(NodesGovernanceNodeAttached)
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
func (it *NodesGovernanceNodeAttachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceNodeAttachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceNodeAttached represents a NodeAttached event raised by the NodesGovernance contract.
type NodesGovernanceNodeAttached struct {
	IdentifierOfProvider common.Address
	IdentifierOfServer   common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNodeAttached is a free log retrieval operation binding the contract event 0x042e2dbbe7288392f225b5cc7f610946ae1a1e8673b3f8ce4f6fc3bc6dee5aa6.
//
// Solidity: event NodeAttached(address indexed identifierOfProvider, address indexed identifierOfServer)
func (_NodesGovernance *NodesGovernanceFilterer) FilterNodeAttached(opts *bind.FilterOpts, identifierOfProvider []common.Address, identifierOfServer []common.Address) (*NodesGovernanceNodeAttachedIterator, error) {

	var identifierOfProviderRule []interface{}
	for _, identifierOfProviderItem := range identifierOfProvider {
		identifierOfProviderRule = append(identifierOfProviderRule, identifierOfProviderItem)
	}
	var identifierOfServerRule []interface{}
	for _, identifierOfServerItem := range identifierOfServer {
		identifierOfServerRule = append(identifierOfServerRule, identifierOfServerItem)
	}

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "NodeAttached", identifierOfProviderRule, identifierOfServerRule)
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceNodeAttachedIterator{contract: _NodesGovernance.contract, event: "NodeAttached", logs: logs, sub: sub}, nil
}

// WatchNodeAttached is a free log subscription operation binding the contract event 0x042e2dbbe7288392f225b5cc7f610946ae1a1e8673b3f8ce4f6fc3bc6dee5aa6.
//
// Solidity: event NodeAttached(address indexed identifierOfProvider, address indexed identifierOfServer)
func (_NodesGovernance *NodesGovernanceFilterer) WatchNodeAttached(opts *bind.WatchOpts, sink chan<- *NodesGovernanceNodeAttached, identifierOfProvider []common.Address, identifierOfServer []common.Address) (event.Subscription, error) {

	var identifierOfProviderRule []interface{}
	for _, identifierOfProviderItem := range identifierOfProvider {
		identifierOfProviderRule = append(identifierOfProviderRule, identifierOfProviderItem)
	}
	var identifierOfServerRule []interface{}
	for _, identifierOfServerItem := range identifierOfServer {
		identifierOfServerRule = append(identifierOfServerRule, identifierOfServerItem)
	}

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "NodeAttached", identifierOfProviderRule, identifierOfServerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceNodeAttached)
				if err := _NodesGovernance.contract.UnpackLog(event, "NodeAttached", log); err != nil {
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
func (_NodesGovernance *NodesGovernanceFilterer) ParseNodeAttached(log types.Log) (*NodesGovernanceNodeAttached, error) {
	event := new(NodesGovernanceNodeAttached)
	if err := _NodesGovernance.contract.UnpackLog(event, "NodeAttached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesGovernanceNodeDeregisteredIterator is returned from FilterNodeDeregistered and is used to iterate over the raw logs and unpacked data for NodeDeregistered events raised by the NodesGovernance contract.
type NodesGovernanceNodeDeregisteredIterator struct {
	Event *NodesGovernanceNodeDeregistered // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceNodeDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceNodeDeregistered)
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
		it.Event = new(NodesGovernanceNodeDeregistered)
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
func (it *NodesGovernanceNodeDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceNodeDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceNodeDeregistered represents a NodeDeregistered event raised by the NodesGovernance contract.
type NodesGovernanceNodeDeregistered struct {
	Identifier      common.Address
	Time            *big.Int
	AliasIdentifier string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeDeregistered is a free log retrieval operation binding the contract event 0x60d01d146c7aa1a7d4e3fdd5543872f7d5b2a241980a66b3552ae1a86ae18453.
//
// Solidity: event NodeDeregistered(address indexed identifier, uint256 time, string aliasIdentifier)
func (_NodesGovernance *NodesGovernanceFilterer) FilterNodeDeregistered(opts *bind.FilterOpts, identifier []common.Address) (*NodesGovernanceNodeDeregisteredIterator, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "NodeDeregistered", identifierRule)
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceNodeDeregisteredIterator{contract: _NodesGovernance.contract, event: "NodeDeregistered", logs: logs, sub: sub}, nil
}

// WatchNodeDeregistered is a free log subscription operation binding the contract event 0x60d01d146c7aa1a7d4e3fdd5543872f7d5b2a241980a66b3552ae1a86ae18453.
//
// Solidity: event NodeDeregistered(address indexed identifier, uint256 time, string aliasIdentifier)
func (_NodesGovernance *NodesGovernanceFilterer) WatchNodeDeregistered(opts *bind.WatchOpts, sink chan<- *NodesGovernanceNodeDeregistered, identifier []common.Address) (event.Subscription, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "NodeDeregistered", identifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceNodeDeregistered)
				if err := _NodesGovernance.contract.UnpackLog(event, "NodeDeregistered", log); err != nil {
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
func (_NodesGovernance *NodesGovernanceFilterer) ParseNodeDeregistered(log types.Log) (*NodesGovernanceNodeDeregistered, error) {
	event := new(NodesGovernanceNodeDeregistered)
	if err := _NodesGovernance.contract.UnpackLog(event, "NodeDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesGovernanceNodeDetachedIterator is returned from FilterNodeDetached and is used to iterate over the raw logs and unpacked data for NodeDetached events raised by the NodesGovernance contract.
type NodesGovernanceNodeDetachedIterator struct {
	Event *NodesGovernanceNodeDetached // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceNodeDetachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceNodeDetached)
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
		it.Event = new(NodesGovernanceNodeDetached)
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
func (it *NodesGovernanceNodeDetachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceNodeDetachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceNodeDetached represents a NodeDetached event raised by the NodesGovernance contract.
type NodesGovernanceNodeDetached struct {
	IdentifierOfProvider common.Address
	IdentifierOfServer   common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNodeDetached is a free log retrieval operation binding the contract event 0x864dd06d15940858b627540246db79e66082ff76b9712b0bb5d483d168b1cd18.
//
// Solidity: event NodeDetached(address indexed identifierOfProvider, address indexed identifierOfServer)
func (_NodesGovernance *NodesGovernanceFilterer) FilterNodeDetached(opts *bind.FilterOpts, identifierOfProvider []common.Address, identifierOfServer []common.Address) (*NodesGovernanceNodeDetachedIterator, error) {

	var identifierOfProviderRule []interface{}
	for _, identifierOfProviderItem := range identifierOfProvider {
		identifierOfProviderRule = append(identifierOfProviderRule, identifierOfProviderItem)
	}
	var identifierOfServerRule []interface{}
	for _, identifierOfServerItem := range identifierOfServer {
		identifierOfServerRule = append(identifierOfServerRule, identifierOfServerItem)
	}

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "NodeDetached", identifierOfProviderRule, identifierOfServerRule)
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceNodeDetachedIterator{contract: _NodesGovernance.contract, event: "NodeDetached", logs: logs, sub: sub}, nil
}

// WatchNodeDetached is a free log subscription operation binding the contract event 0x864dd06d15940858b627540246db79e66082ff76b9712b0bb5d483d168b1cd18.
//
// Solidity: event NodeDetached(address indexed identifierOfProvider, address indexed identifierOfServer)
func (_NodesGovernance *NodesGovernanceFilterer) WatchNodeDetached(opts *bind.WatchOpts, sink chan<- *NodesGovernanceNodeDetached, identifierOfProvider []common.Address, identifierOfServer []common.Address) (event.Subscription, error) {

	var identifierOfProviderRule []interface{}
	for _, identifierOfProviderItem := range identifierOfProvider {
		identifierOfProviderRule = append(identifierOfProviderRule, identifierOfProviderItem)
	}
	var identifierOfServerRule []interface{}
	for _, identifierOfServerItem := range identifierOfServer {
		identifierOfServerRule = append(identifierOfServerRule, identifierOfServerItem)
	}

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "NodeDetached", identifierOfProviderRule, identifierOfServerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceNodeDetached)
				if err := _NodesGovernance.contract.UnpackLog(event, "NodeDetached", log); err != nil {
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
func (_NodesGovernance *NodesGovernanceFilterer) ParseNodeDetached(log types.Log) (*NodesGovernanceNodeDetached, error) {
	event := new(NodesGovernanceNodeDetached)
	if err := _NodesGovernance.contract.UnpackLog(event, "NodeDetached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesGovernanceNodeRegisteredIterator is returned from FilterNodeRegistered and is used to iterate over the raw logs and unpacked data for NodeRegistered events raised by the NodesGovernance contract.
type NodesGovernanceNodeRegisteredIterator struct {
	Event *NodesGovernanceNodeRegistered // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceNodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceNodeRegistered)
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
		it.Event = new(NodesGovernanceNodeRegistered)
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
func (it *NodesGovernanceNodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceNodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceNodeRegistered represents a NodeRegistered event raised by the NodesGovernance contract.
type NodesGovernanceNodeRegistered struct {
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
func (_NodesGovernance *NodesGovernanceFilterer) FilterNodeRegistered(opts *bind.FilterOpts, wallet []common.Address) (*NodesGovernanceNodeRegisteredIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "NodeRegistered", walletRule)
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceNodeRegisteredIterator{contract: _NodesGovernance.contract, event: "NodeRegistered", logs: logs, sub: sub}, nil
}

// WatchNodeRegistered is a free log subscription operation binding the contract event 0x99c250edfb33141684f230eecda8aa955bf0d62d17bcaaab0ab43f318f3637b0.
//
// Solidity: event NodeRegistered(address indexed wallet, address identifier, uint256 time, string aliasIdentifier, string[] gpuTypes, uint256[] gpuNums)
func (_NodesGovernance *NodesGovernanceFilterer) WatchNodeRegistered(opts *bind.WatchOpts, sink chan<- *NodesGovernanceNodeRegistered, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "NodeRegistered", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceNodeRegistered)
				if err := _NodesGovernance.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
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
func (_NodesGovernance *NodesGovernanceFilterer) ParseNodeRegistered(log types.Log) (*NodesGovernanceNodeRegistered, error) {
	event := new(NodesGovernanceNodeRegistered)
	if err := _NodesGovernance.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesGovernanceProxyNodeRegisteredIterator is returned from FilterProxyNodeRegistered and is used to iterate over the raw logs and unpacked data for ProxyNodeRegistered events raised by the NodesGovernance contract.
type NodesGovernanceProxyNodeRegisteredIterator struct {
	Event *NodesGovernanceProxyNodeRegistered // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceProxyNodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceProxyNodeRegistered)
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
		it.Event = new(NodesGovernanceProxyNodeRegistered)
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
func (it *NodesGovernanceProxyNodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceProxyNodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceProxyNodeRegistered represents a ProxyNodeRegistered event raised by the NodesGovernance contract.
type NodesGovernanceProxyNodeRegistered struct {
	Proxy common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProxyNodeRegistered is a free log retrieval operation binding the contract event 0x34093c616a3ab1dd56c3a4780155eb800bf5d1c3d024468ea10940d25fa9538d.
//
// Solidity: event ProxyNodeRegistered(address indexed proxy)
func (_NodesGovernance *NodesGovernanceFilterer) FilterProxyNodeRegistered(opts *bind.FilterOpts, proxy []common.Address) (*NodesGovernanceProxyNodeRegisteredIterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "ProxyNodeRegistered", proxyRule)
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceProxyNodeRegisteredIterator{contract: _NodesGovernance.contract, event: "ProxyNodeRegistered", logs: logs, sub: sub}, nil
}

// WatchProxyNodeRegistered is a free log subscription operation binding the contract event 0x34093c616a3ab1dd56c3a4780155eb800bf5d1c3d024468ea10940d25fa9538d.
//
// Solidity: event ProxyNodeRegistered(address indexed proxy)
func (_NodesGovernance *NodesGovernanceFilterer) WatchProxyNodeRegistered(opts *bind.WatchOpts, sink chan<- *NodesGovernanceProxyNodeRegistered, proxy []common.Address) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "ProxyNodeRegistered", proxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceProxyNodeRegistered)
				if err := _NodesGovernance.contract.UnpackLog(event, "ProxyNodeRegistered", log); err != nil {
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
func (_NodesGovernance *NodesGovernanceFilterer) ParseProxyNodeRegistered(log types.Log) (*NodesGovernanceProxyNodeRegistered, error) {
	event := new(NodesGovernanceProxyNodeRegistered)
	if err := _NodesGovernance.contract.UnpackLog(event, "ProxyNodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesGovernanceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the NodesGovernance contract.
type NodesGovernanceRoleAdminChangedIterator struct {
	Event *NodesGovernanceRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceRoleAdminChanged)
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
		it.Event = new(NodesGovernanceRoleAdminChanged)
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
func (it *NodesGovernanceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceRoleAdminChanged represents a RoleAdminChanged event raised by the NodesGovernance contract.
type NodesGovernanceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NodesGovernance *NodesGovernanceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NodesGovernanceRoleAdminChangedIterator, error) {

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

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceRoleAdminChangedIterator{contract: _NodesGovernance.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NodesGovernance *NodesGovernanceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NodesGovernanceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceRoleAdminChanged)
				if err := _NodesGovernance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_NodesGovernance *NodesGovernanceFilterer) ParseRoleAdminChanged(log types.Log) (*NodesGovernanceRoleAdminChanged, error) {
	event := new(NodesGovernanceRoleAdminChanged)
	if err := _NodesGovernance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesGovernanceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the NodesGovernance contract.
type NodesGovernanceRoleGrantedIterator struct {
	Event *NodesGovernanceRoleGranted // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceRoleGranted)
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
		it.Event = new(NodesGovernanceRoleGranted)
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
func (it *NodesGovernanceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceRoleGranted represents a RoleGranted event raised by the NodesGovernance contract.
type NodesGovernanceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NodesGovernance *NodesGovernanceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NodesGovernanceRoleGrantedIterator, error) {

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

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceRoleGrantedIterator{contract: _NodesGovernance.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NodesGovernance *NodesGovernanceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NodesGovernanceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceRoleGranted)
				if err := _NodesGovernance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_NodesGovernance *NodesGovernanceFilterer) ParseRoleGranted(log types.Log) (*NodesGovernanceRoleGranted, error) {
	event := new(NodesGovernanceRoleGranted)
	if err := _NodesGovernance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesGovernanceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the NodesGovernance contract.
type NodesGovernanceRoleRevokedIterator struct {
	Event *NodesGovernanceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceRoleRevoked)
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
		it.Event = new(NodesGovernanceRoleRevoked)
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
func (it *NodesGovernanceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceRoleRevoked represents a RoleRevoked event raised by the NodesGovernance contract.
type NodesGovernanceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NodesGovernance *NodesGovernanceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NodesGovernanceRoleRevokedIterator, error) {

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

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceRoleRevokedIterator{contract: _NodesGovernance.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NodesGovernance *NodesGovernanceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NodesGovernanceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceRoleRevoked)
				if err := _NodesGovernance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_NodesGovernance *NodesGovernanceFilterer) ParseRoleRevoked(log types.Log) (*NodesGovernanceRoleRevoked, error) {
	event := new(NodesGovernanceRoleRevoked)
	if err := _NodesGovernance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesGovernanceSettlementResultIterator is returned from FilterSettlementResult and is used to iterate over the raw logs and unpacked data for SettlementResult events raised by the NodesGovernance contract.
type NodesGovernanceSettlementResultIterator struct {
	Event *NodesGovernanceSettlementResult // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceSettlementResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceSettlementResult)
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
		it.Event = new(NodesGovernanceSettlementResult)
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
func (it *NodesGovernanceSettlementResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceSettlementResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceSettlementResult represents a SettlementResult event raised by the NodesGovernance contract.
type NodesGovernanceSettlementResult struct {
	States     []NodeState
	TotalQuota *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSettlementResult is a free log retrieval operation binding the contract event 0xd51417935ddbb98970f20a5f6f9c5070ce90768d0e3bfaba49e7e2f8621debac.
//
// Solidity: event SettlementResult((uint64,uint64,uint128,address,address)[] states, uint256 totalQuota)
func (_NodesGovernance *NodesGovernanceFilterer) FilterSettlementResult(opts *bind.FilterOpts) (*NodesGovernanceSettlementResultIterator, error) {

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "SettlementResult")
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceSettlementResultIterator{contract: _NodesGovernance.contract, event: "SettlementResult", logs: logs, sub: sub}, nil
}

// WatchSettlementResult is a free log subscription operation binding the contract event 0xd51417935ddbb98970f20a5f6f9c5070ce90768d0e3bfaba49e7e2f8621debac.
//
// Solidity: event SettlementResult((uint64,uint64,uint128,address,address)[] states, uint256 totalQuota)
func (_NodesGovernance *NodesGovernanceFilterer) WatchSettlementResult(opts *bind.WatchOpts, sink chan<- *NodesGovernanceSettlementResult) (event.Subscription, error) {

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "SettlementResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceSettlementResult)
				if err := _NodesGovernance.contract.UnpackLog(event, "SettlementResult", log); err != nil {
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

// ParseSettlementResult is a log parse operation binding the contract event 0xd51417935ddbb98970f20a5f6f9c5070ce90768d0e3bfaba49e7e2f8621debac.
//
// Solidity: event SettlementResult((uint64,uint64,uint128,address,address)[] states, uint256 totalQuota)
func (_NodesGovernance *NodesGovernanceFilterer) ParseSettlementResult(log types.Log) (*NodesGovernanceSettlementResult, error) {
	event := new(NodesGovernanceSettlementResult)
	if err := _NodesGovernance.contract.UnpackLog(event, "SettlementResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesGovernanceValidationResultIterator is returned from FilterValidationResult and is used to iterate over the raw logs and unpacked data for ValidationResult events raised by the NodesGovernance contract.
type NodesGovernanceValidationResultIterator struct {
	Event *NodesGovernanceValidationResult // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceValidationResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceValidationResult)
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
		it.Event = new(NodesGovernanceValidationResult)
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
func (it *NodesGovernanceValidationResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceValidationResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceValidationResult represents a ValidationResult event raised by the NodesGovernance contract.
type NodesGovernanceValidationResult struct {
	RoundId   *big.Int
	Validator common.Address
	Result    bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidationResult is a free log retrieval operation binding the contract event 0x92de7c81b7cf6c7977d7cd091ccd01996264a02b7dbbce5d2a2524a8daabe769.
//
// Solidity: event ValidationResult(uint256 roundId, address validator, bool result)
func (_NodesGovernance *NodesGovernanceFilterer) FilterValidationResult(opts *bind.FilterOpts) (*NodesGovernanceValidationResultIterator, error) {

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "ValidationResult")
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceValidationResultIterator{contract: _NodesGovernance.contract, event: "ValidationResult", logs: logs, sub: sub}, nil
}

// WatchValidationResult is a free log subscription operation binding the contract event 0x92de7c81b7cf6c7977d7cd091ccd01996264a02b7dbbce5d2a2524a8daabe769.
//
// Solidity: event ValidationResult(uint256 roundId, address validator, bool result)
func (_NodesGovernance *NodesGovernanceFilterer) WatchValidationResult(opts *bind.WatchOpts, sink chan<- *NodesGovernanceValidationResult) (event.Subscription, error) {

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "ValidationResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceValidationResult)
				if err := _NodesGovernance.contract.UnpackLog(event, "ValidationResult", log); err != nil {
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

// ParseValidationResult is a log parse operation binding the contract event 0x92de7c81b7cf6c7977d7cd091ccd01996264a02b7dbbce5d2a2524a8daabe769.
//
// Solidity: event ValidationResult(uint256 roundId, address validator, bool result)
func (_NodesGovernance *NodesGovernanceFilterer) ParseValidationResult(log types.Log) (*NodesGovernanceValidationResult, error) {
	event := new(NodesGovernanceValidationResult)
	if err := _NodesGovernance.contract.UnpackLog(event, "ValidationResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodesGovernanceValidationStartedIterator is returned from FilterValidationStarted and is used to iterate over the raw logs and unpacked data for ValidationStarted events raised by the NodesGovernance contract.
type NodesGovernanceValidationStartedIterator struct {
	Event *NodesGovernanceValidationStarted // Event containing the contract specifics and raw log

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
func (it *NodesGovernanceValidationStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodesGovernanceValidationStarted)
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
		it.Event = new(NodesGovernanceValidationStarted)
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
func (it *NodesGovernanceValidationStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodesGovernanceValidationStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodesGovernanceValidationStarted represents a ValidationStarted event raised by the NodesGovernance contract.
type NodesGovernanceValidationStarted struct {
	RoundId                *big.Int
	ExpectedCompletionTime *big.Int
	Candidate              common.Address
	Validators             []common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterValidationStarted is a free log retrieval operation binding the contract event 0x71afff60b83105500984ce43d4633544224775a10de240da021704c056b58bdb.
//
// Solidity: event ValidationStarted(uint256 roundId, uint256 expectedCompletionTime, address candidate, address[] validators)
func (_NodesGovernance *NodesGovernanceFilterer) FilterValidationStarted(opts *bind.FilterOpts) (*NodesGovernanceValidationStartedIterator, error) {

	logs, sub, err := _NodesGovernance.contract.FilterLogs(opts, "ValidationStarted")
	if err != nil {
		return nil, err
	}
	return &NodesGovernanceValidationStartedIterator{contract: _NodesGovernance.contract, event: "ValidationStarted", logs: logs, sub: sub}, nil
}

// WatchValidationStarted is a free log subscription operation binding the contract event 0x71afff60b83105500984ce43d4633544224775a10de240da021704c056b58bdb.
//
// Solidity: event ValidationStarted(uint256 roundId, uint256 expectedCompletionTime, address candidate, address[] validators)
func (_NodesGovernance *NodesGovernanceFilterer) WatchValidationStarted(opts *bind.WatchOpts, sink chan<- *NodesGovernanceValidationStarted) (event.Subscription, error) {

	logs, sub, err := _NodesGovernance.contract.WatchLogs(opts, "ValidationStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodesGovernanceValidationStarted)
				if err := _NodesGovernance.contract.UnpackLog(event, "ValidationStarted", log); err != nil {
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

// ParseValidationStarted is a log parse operation binding the contract event 0x71afff60b83105500984ce43d4633544224775a10de240da021704c056b58bdb.
//
// Solidity: event ValidationStarted(uint256 roundId, uint256 expectedCompletionTime, address candidate, address[] validators)
func (_NodesGovernance *NodesGovernanceFilterer) ParseValidationStarted(log types.Log) (*NodesGovernanceValidationStarted, error) {
	event := new(NodesGovernanceValidationStarted)
	if err := _NodesGovernance.contract.UnpackLog(event, "ValidationStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
