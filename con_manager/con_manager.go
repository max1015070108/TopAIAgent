package con_manager

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/max1015070108/TopAIAgent/con_manager/AIModels"
	"github.com/max1015070108/TopAIAgent/con_manager/AIWorkload"
	"github.com/max1015070108/TopAIAgent/con_manager/NodesGovernance"
	"github.com/max1015070108/TopAIAgent/con_manager/NodesRegistry"
	"github.com/max1015070108/TopAIAgent/config"
	"github.com/max1015070108/TopAIAgent/database"
	"github.com/max1015070108/TopAIAgent/wallet"
)

type ConManager struct {
	Client          *ethclient.Client
	AIModels        *AIModels.AIModels
	AIWorkload      *AIWorkload.AIWorkload
	NodesRegistry   *NodesRegistry.NodesRegistry
	NodesGovernance *NodesGovernance.NodesGovernance
	//add wallet
	Wallet    wallet.WalletManager
	Conf      *config.Config
	DataEvent *database.EventStore
}

func NewConManager(url string) (*ConManager, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	//load config file
	config := config.NewConfigLoader("./config.json", ".env")
	confData, err := config.Load()
	if err != nil {
		return nil, err
	}

	// DataEvent, err := database.NewEventStore() // DBName would be used for sqlite3
	// if err != nil {
	// 	return nil, err
	// }

	// err = DataEvent.InitTables()
	// if err != nil {
	// 	return nil, err
	// }

	if confData.Ethereum.KeystoreDir == "" {
		confData.Ethereum.KeystoreDir = wallet.DefaultKeyDir
	}
	walletMan := wallet.NewWalletManager(confData.Ethereum.KeystoreDir)

	// AIModels contract
	aim_contract, err := AIModels.NewAIModels(common.HexToAddress(confData.ContractAddress.AIModels), client)
	if err != nil {
		return nil, err
	}

	// AIWorkload contract
	aiw_contract, err := AIWorkload.NewAIWorkload(common.HexToAddress(confData.ContractAddress.AIWorkerload), client)

	if err != nil {
		return nil, err
	}

	nodes_registry, err := NodesRegistry.NewNodesRegistry(common.HexToAddress(confData.ContractAddress.NodeRegister), client)

	if err != nil {
		return nil, err
	}

	Nodes_governance, err := NodesGovernance.NewNodesGovernance(common.HexToAddress(confData.ContractAddress.NodesGovernance), client)
	if err != nil {
		return nil, err
	}

	return &ConManager{
		Client:          client,
		Wallet:          *walletMan,
		AIModels:        aim_contract,
		AIWorkload:      aiw_contract,
		NodesRegistry:   nodes_registry,
		NodesGovernance: Nodes_governance,
		Conf:            confData,
		// DataEvent:       DataEvent,
	}, nil
}

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
	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	return nil, fmt.Errorf("error casting public key to ECDSA")
	// }
	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// // 4. 获取最新nonce
	// nonce, err := client.PendingNonceAt(ctx, fromAddress)
	// if err != nil {
	// 	return nil, fmt.Errorf("get nonce error: %v", err)
	// }
	// auth.Nonce = big.NewInt(int64(nonce))

	// // 5. 尝试获取gasTipCap以检查是否支持EIP-1559
	// gasTipCap, err := client.SuggestGasTipCap(ctx)
	// if err != nil {
	// 	// 如果不支持EIP-1559，使用传统的gasPrice
	// 	gasPrice, err := client.SuggestGasPrice(ctx)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("get gas price error: %v", err)
	// 	}
	// 	auth.GasPrice = big.NewInt(1).Mul(gasPrice, big.NewInt(10))
	// 	// 清除EIP-1559相关字段
	// 	auth.GasFeeCap = nil
	// 	auth.GasTipCap = nil
	// } else {
	// 	// 支持EIP-1559，设置相关参数
	// 	baseFee, err := getBaseFee(client)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("get base fee error: %v", err)
	// 	}

	// 	// 计算maxFeePerGas: baseFee * 2 + gasTipCap
	// 	maxFeePerGas := new(big.Int).Add(
	// 		new(big.Int).Mul(baseFee, big.NewInt(10)),
	// 		gasTipCap,
	// 	)

	// 	// 设置EIP-1559参数
	// 	auth.GasFeeCap = maxFeePerGas
	// 	auth.GasTipCap = gasTipCap
	// 	// 清除传统gasPrice
	// 	auth.GasPrice = nil
	// }

	// // 6. 估算gasLimit
	// addr := common.HexToAddress(contract)
	// gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{
	// 	From:  fromAddress,
	// 	To:    &addr,
	// 	Gas:   0,
	// 	Value: big.NewInt(0),
	// 	Data:  nil,
	// })
	// if err != nil {
	// 	auth.GasLimit = uint64(1000000)
	// } else {
	// 	auth.GasLimit = uint64(float64(gasLimit) * 1.2)
	// }

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

func (c *ConManager) GetLastBlock() (*types.Header, error) {

	headInfo, err := c.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return headInfo, err
	// return struct {
	// 	TimeStamp uint64
	// 	Hash      string
	// }{
	// 	TimeStamp: headInfo.Time,
	// 	Hash:      headInfo.Hash().Hex(),
	// }, nil
}

func (c *ConManager) GetBlockByHeighNumber(id int64) (*types.Header, error) {

	headInfo, err := c.Client.HeaderByNumber(context.Background(), big.NewInt(id))
	if err != nil {
		return nil, err
	}
	return headInfo, nil
}

func (c *ConManager) GetLatestBlock(id int64) (struct {
	TimeStamp uint64
	Hash      string
}, error) {

	headInfo, err := c.Client.HeaderByNumber(context.Background(), big.NewInt(id))
	if err != nil {
		return struct {
			TimeStamp uint64
			Hash      string
		}{}, err
	}
	return struct {
		TimeStamp uint64
		Hash      string
	}{
		TimeStamp: headInfo.Time,
		Hash:      headInfo.Hash().Hex(),
	}, nil
}
