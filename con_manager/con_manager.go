package con_manager

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/topaiagent/con_manager/AIModels"
	"github.com/topaiagent/con_manager/AIWorkload"
	"github.com/topaiagent/con_manager/NodesRegistry"
	"github.com/topaiagent/config"
)

type ConManager struct {
	Client        *ethclient.Client
	AIModels      *AIModels.AIModels
	AIWorkload    *AIWorkload.AIWorkload
	NodesRegistry *NodesRegistry.NodesRegistry
}

func NewConManager(url string) (*ConManager, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	//load config file
	config := config.NewConfigLoader("config.json", ".env")
	confData, err := config.Load()
	if err != nil {
		return nil, err
	}

	fmt.Printf(".... %+v\n", confData.ContractAddress)

	return &ConManager{Client: client}, nil
}

// func CreateLatestAuth(client *ethclient.Client, privateKey *ecdsa.PrivateKey, contract string) (*bind.TransactOpts, error) {
// 	// 实现转账逻辑

// 	ctx := context.Background()

// 	// 1. 获取链ID
// 	chainID, err := client.ChainID(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("get chain id error: %v", err)
// 	}

// 	// 2. 创建基础auth对象
// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
// 	if err != nil {
// 		return nil, fmt.Errorf("create auth error: %v", err)
// 	}

// 	// 3. 获取发送者地址
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		return nil, fmt.Errorf("error casting public key to ECDSA")
// 	}
// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

// 	// 4. 获取最新nonce
// 	nonce, err := client.PendingNonceAt(ctx, fromAddress)
// 	if err != nil {
// 		return nil, fmt.Errorf("get nonce error: %v", err)
// 	}
// 	auth.Nonce = big.NewInt(int64(nonce))

// 	// 5. 获取最新gasPrice
// 	gasPrice, err := client.SuggestGasPrice(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("get gas price error: %v", err)
// 	}
// 	auth.GasPrice = gasPrice

// 	// 6. 获取最新gasTipCap (EIP-1559)
// 	gasTipCap, err := client.SuggestGasTipCap(ctx)
// 	if err != nil {
// 		// 如果网络不支持EIP-1559，使用普通gas price
// 		auth.GasFeeCap = gasPrice
// 		auth.GasTipCap = gasPrice
// 	} else {
// 		// 设置EIP-1559参数
// 		gasFeeCap := new(big.Int).Add(
// 			gasTipCap,
// 			new(big.Int).Mul(gasPrice, big.NewInt(2)),
// 		)
// 		auth.GasFeeCap = gasFeeCap
// 		auth.GasTipCap = gasTipCap
// 	}

// 	// 7. 估算gasLimit (这里展示如何为合约调用估算)
// 	// 注意：这是示例代码，实际使用时需要根据具体交易内容估算
// 	//
// 	//
// 	addr := common.HexToAddress(contract)
// 	gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{
// 		From:  fromAddress,
// 		To:    &addr, // 合约地址
// 		Gas:   0,
// 		Value: big.NewInt(0),
// 		Data:  nil, // 这里需要填入实际的合约调用数据
// 	})
// 	if err != nil {
// 		// 如果估算失败，使用默认值
// 		auth.GasLimit = uint64(300000)
// 	} else {
// 		// 为了安全，给估算值加上一个缓冲
// 		auth.GasLimit = uint64(float64(gasLimit) * 1.2)
// 	}

// 	return auth, nil
// }

func CreateLatestAuth(client *ethclient.Client, privateKey *ecdsa.PrivateKey, contract string) (*bind.TransactOpts, error) {
	ctx := context.Background()

	// 1. 获取链ID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain id error: %v", err)
	}

	// 2. 创建基础auth对象
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("create auth error: %v", err)
	}

	// 3. 获取发送者地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 4. 获取最新nonce
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, fmt.Errorf("get nonce error: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// 5. 尝试获取gasTipCap以检查是否支持EIP-1559
	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		// 如果不支持EIP-1559，使用传统的gasPrice
		gasPrice, err := client.SuggestGasPrice(ctx)
		if err != nil {
			return nil, fmt.Errorf("get gas price error: %v", err)
		}
		auth.GasPrice = gasPrice
		// 清除EIP-1559相关字段
		auth.GasFeeCap = nil
		auth.GasTipCap = nil
	} else {
		// 支持EIP-1559，设置相关参数
		baseFee, err := getBaseFee(client)
		if err != nil {
			return nil, fmt.Errorf("get base fee error: %v", err)
		}

		// 计算maxFeePerGas: baseFee * 2 + gasTipCap
		maxFeePerGas := new(big.Int).Add(
			new(big.Int).Mul(baseFee, big.NewInt(2)),
			gasTipCap,
		)

		// 设置EIP-1559参数
		auth.GasFeeCap = maxFeePerGas
		auth.GasTipCap = gasTipCap
		// 清除传统gasPrice
		auth.GasPrice = nil
	}

	// 6. 估算gasLimit
	addr := common.HexToAddress(contract)
	gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:  fromAddress,
		To:    &addr,
		Gas:   0,
		Value: big.NewInt(0),
		Data:  nil,
	})
	if err != nil {
		auth.GasLimit = uint64(300000)
	} else {
		auth.GasLimit = uint64(float64(gasLimit) * 1.2)
	}

	return auth, nil
}

// 获取当前区块的baseFee
func getBaseFee(client *ethclient.Client) (*big.Int, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	if header.BaseFee == nil {
		return big.NewInt(0), nil
	}
	return header.BaseFee, nil
}

// ==========================AIModels===============================
func (c *ConManager) GetNodeDeployment(opts *bind.CallOpts, contractAddr common.Address) (string, error) {
	// 实现合约调用逻辑
	return "", nil
}