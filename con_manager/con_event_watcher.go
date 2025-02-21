package con_manager

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/max1015070108/TopAIAgent/con_manager/AIModels"
	AIWorkerload "github.com/max1015070108/TopAIAgent/con_manager/AIWorkload"
	"github.com/max1015070108/TopAIAgent/con_manager/NodesGovernance"
	"github.com/max1015070108/TopAIAgent/con_manager/NodesRegistry"
	"github.com/max1015070108/TopAIAgent/database"
)

func (c *ConManager) WatchEvents(ctx context.Context, wg *sync.WaitGroup) ([]common.Address, error) {

	// blockInfoChan := database.InfoChan{}

	channel := c.DataEvent.EventDataStoreChannel(1000)
	// defer blockInfoChan.Close()

	// Start background goroutine to watch headers
	go func(wg *sync.WaitGroup) {

		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-channel.Err:
				log.Printf("Error processing block info: %v", err)
			default:
				lastBlockInfo, err := c.DataEvent.GetLatestBlockInfo()
				if err != nil && err.Error() != "sql: no rows in result set" {
					time.Sleep(1 * time.Second)
					continue
				}

				currentHeader, err := c.Client.HeaderByNumber(ctx, nil)
				if err != nil {
					continue
				}

				storeNumber := int64(0)
				if lastBlockInfo != nil {
					storeNumber = lastBlockInfo.Number
				}
				if storeNumber < currentHeader.Number.Int64() {
					logs, err := c.GetLogs(storeNumber, currentHeader.Number.Int64())
					if err != nil {
						continue
					}

					if err := c.HandleLogs(logs); err != nil {

						fmt.Printf("HandleLogs error: ", err)
						continue
					}

					bInfo := database.BlockInfo{
						Number:    currentHeader.Number.Int64(),
						BlockHash: currentHeader.Hash().Hex(),
						Timestamp: int64(currentHeader.Time),
					}

					channel.Channel <- bInfo
				}
			}

			time.Sleep(1 * time.Second)
		}
	}(wg)

	return nil, nil
}

//get log from blockinfo

type LogHandler struct {
	client          *ethclient.Client
	contractAddress common.Address
}

func (c *ConManager) GetLogs(fromBlock, toBlock int64) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{
			common.HexToAddress(c.Conf.ContractAddress.AIModels),
			common.HexToAddress(c.Conf.ContractAddress.AIWorkerload),
			common.HexToAddress(c.Conf.ContractAddress.NodeRegister),
			common.HexToAddress(c.Conf.ContractAddress.NodesGovernance),
		}, // 合约地址
		Topics: [][]common.Hash{},
	}

	return c.Client.FilterLogs(context.Background(), query)
}

// 处理日志的函数
func (c *ConManager) HandleLogs(logs []types.Log) error {
	// 获取合约事件的ABI
	aiWorkloadAbi, err := AIWorkerload.AIWorkloadMetaData.GetAbi()
	if err != nil {
		return err
	}

	aiModelsAbi, err := AIModels.AIModelsMetaData.GetAbi()
	if err != nil {
		return err
	}

	nodeRegistryAbi, err := NodesRegistry.NodesRegistryMetaData.GetAbi()
	if err != nil {
		return err
	}

	nodesGovernanceAbi, err := NodesGovernance.NodesGovernanceMetaData.GetAbi()
	if err != nil {
		return err
	}

	// 遍历所有日志
	for _, log := range logs {
		// 根据事件签名判断事件类型
		switch log.Topics[0].Hex() {
		case aiModelsAbi.Events["UploadModeled"].ID.Hex():
			event := new(AIModels.AIModelsUploadModeled)
			err := aiModelsAbi.UnpackIntoInterface(event, "UploadModeled", log.Data)
			if err != nil {
				fmt.Println("UnpackIntoInterface error: ", err)
				continue
			}
			fmt.Println("log.Topics[0].Hex():", log.Topics[0].Hex())

		case aiModelsAbi.Events["ModelDeployed"].ID.Hex():
			event := new(AIModels.AIModelsModelDeployed)
			err := aiModelsAbi.UnpackIntoInterface(event, "ModelDeployed", log.Data)
			if err != nil {
				fmt.Println("UnpackIntoInterface error: ", err)
				continue
			}
			fmt.Println("log.Topics[0].Hex():", log.Topics[0].Hex())

		case aiModelsAbi.Events["ModelRemoved"].ID.Hex():
			event := new(AIModels.AIModelsModelRemoved)
			err := aiModelsAbi.UnpackIntoInterface(event, "ModelRemoved", log.Data)
			if err != nil {
				fmt.Println("UnpackIntoInterface error: ", err)
				continue
			}
			fmt.Println("log.Topics[0].Hex():", log.Topics[0].Hex())

		// case aiModelsAbi.Events["ModelUploadStaked"].ID.Hex():
		// 	event := new(AIModels.AIModelsModelUploadStaked)
		// 	err := aiModelsAbi.UnpackIntoInterface(event, "ModelUploadStaked", log.Data)
		// 	if err != nil {
		// 		fmt.Println("UnpackIntoInterface error: ", err)
		// 		continue
		// 	}

		// case aiModelsAbi.Events["ModelUsedStaked"].ID.Hex():
		// 	event := new(AIModels.AIModelsModelUsedStaked)
		// 	err := aiModelsAbi.UnpackIntoInterface(event, "ModelUsedStaked", log.Data)
		// 	if err != nil {
		// 		fmt.Println("UnpackIntoInterface error: ", err)
		// 		continue
		// 	}

		// case aiModelsAbi.Events["ModelUploadUnstaked"].ID.Hex():
		// 	event := new(AIModels.AIModelsModelUploadUnstaked)
		// 	err := aiModelsAbi.UnpackIntoInterface(event, "ModelUploadUnstaked", log.Data)
		// 	if err != nil {
		// 		fmt.Println("UnpackIntoInterface error: ", err)
		// 		continue
		// 	}

		// case aiModelsAbi.Events["ModelUsedUnstaked"].ID.Hex():
		// 	event := new(AIModels.AIModelsModelUsedUnstaked)
		// 	err := aiModelsAbi.UnpackIntoInterface(event, "ModelUsedUnstaked", log.Data)
		// 	if err != nil {
		// 		fmt.Println("UnpackIntoInterface error: ", err)
		// 		continue
		// 	}

		case aiWorkloadAbi.Events["WorkloadReported"].ID.Hex():
			event := new(AIWorkerload.AIWorkloadWorkloadReported)
			err = aiWorkloadAbi.UnpackIntoInterface(event, "WorkloadReported", log.Data)
			if err != nil {
				fmt.Println("UnpackIntoInterface error: ", err)
				continue
			}
			fmt.Println("log.Topics[0].Hex():", log.Topics[0].Hex())

		case nodesGovernanceAbi.Events["ValidationStarted"].ID.Hex():
			event := new(NodesGovernance.NodesGovernanceValidationStarted)
			err = nodesGovernanceAbi.UnpackIntoInterface(event, "ValidationStarted", log.Data)
			if err != nil {
				fmt.Println("UnpackIntoInterface error: ", err)
				continue
			}
			fmt.Println("log.Topics[0].Hex():", log.Topics[0].Hex())

		case nodesGovernanceAbi.Events["ValidationResult"].ID.Hex():
			event := new(NodesGovernance.NodesGovernanceValidationResult)
			err = nodesGovernanceAbi.UnpackIntoInterface(event, "ValidationResult", log.Data)
			if err != nil {
				fmt.Println("UnpackIntoInterface error: ", err)
				continue
			}
			fmt.Println("log.Topics[0].Hex():", log.Topics[0].Hex())

		case nodesGovernanceAbi.Events["SettlementResult"].ID.Hex():
			event := new(NodesGovernance.NodesGovernanceSettlementResult)
			err = nodesGovernanceAbi.UnpackIntoInterface(event, "SettlementResult", log.Data)
			if err != nil {
				fmt.Println("UnpackIntoInterface error: ", err)
				continue
			}
			fmt.Println("log.Topics[0].Hex():", log.Topics[0].Hex())

		case nodeRegistryAbi.Events["NodeRegistered"].ID.Hex():
			event := new(NodesRegistry.NodesRegistryNodeRegistered)
			err = nodeRegistryAbi.UnpackIntoInterface(event, "NodeRegistered", log.Data)
			if err != nil {
				fmt.Println("UnpackIntoInterface error: ", err)
				continue
			}
			fmt.Println("log.Topics[0].Hex():", log.Topics[0].Hex())

		case nodeRegistryAbi.Events["NodeActived"].ID.Hex():
			event := new(NodesRegistry.NodesRegistryNodeActived)
			err = nodeRegistryAbi.UnpackIntoInterface(event, "NodeActived", log.Data)
			if err != nil {
				fmt.Println("UnpackIntoInterface error: ", err)
				continue
			}
			fmt.Println("log.Topics[0].Hex():", log.Topics[0].Hex())

		case nodeRegistryAbi.Events["NodeDeregistered"].ID.Hex():
			event := new(NodesRegistry.NodesRegistryNodeDeregistered)
			err = nodeRegistryAbi.UnpackIntoInterface(event, "NodeDeregistered", log.Data)
			if err != nil {
				fmt.Println("UnpackIntoInterface error: ", err)
				continue
			}
			fmt.Println("log.Topics[0].Hex():", log.Topics[0].Hex())

		case nodeRegistryAbi.Events["Authorized"].ID.Hex():
			event := new(NodesRegistry.NodesRegistryAuthorized)
			err = nodeRegistryAbi.UnpackIntoInterface(event, "Authorized", log.Data)
			if err != nil {
				fmt.Println("UnpackIntoInterface error: ", err)
				continue
			}
			fmt.Println("log.Topics[0].Hex():", log.Topics[0].Hex())

		case nodeRegistryAbi.Events["NodeAttached"].ID.Hex():
			event := new(NodesRegistry.NodesRegistryNodeAttached)
			err = nodeRegistryAbi.UnpackIntoInterface(event, "NodeAttached", log.Data)
			if err != nil {
				fmt.Println("UnpackIntoInterface error: ", err)
				continue
			}
			fmt.Println("log.Topics[0].Hex():", log.Topics[0].Hex())

		case nodeRegistryAbi.Events["NodeDetached"].ID.Hex():
			event := new(NodesRegistry.NodesRegistryNodeDetached)
			err = nodeRegistryAbi.UnpackIntoInterface(event, "NodeDetached", log.Data)
			if err != nil {
				fmt.Println("UnpackIntoInterface error: ", err)
				continue
			}
			fmt.Println("log.Topics[0].Hex():", log.Topics[0].Hex())
		}
	}

	return nil
}
