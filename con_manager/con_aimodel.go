package con_manager

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/max1015070108/TopAIAgent/con_manager/AIModels"
)

func (c *ConManager) GetNodeDeployment(model_addrss string) ([]*big.Int, error) {

	nodes_deploys, err := c.AIModels.GetNodeDeployment(&bind.CallOpts{}, common.HexToAddress(model_addrss))
	if err != nil {
		return nil, err
	}
	return nodes_deploys, nil
}

func (c *ConManager) GetModelDeploymentMap() (map[int64][]common.Address, error) {
	// Create map to store results using int64 instead of *big.Int
	distri := make(map[int64][]common.Address)

	// Get the next model ID (total count)
	modelIdsBI, err := c.AIModels.NextModelId(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get NextModelId: %v", err)
	}

	// Convert *big.Int to int64
	modelIds := modelIdsBI.Int64()
	fmt.Printf("Total models: %v\n", modelIds)

	// Iterate through each model ID from 1 to modelIds-1
	for i := int64(1); i < modelIds; i++ {
		// Create a *big.Int for calling the contract method
		currentId := big.NewInt(i)

		// Get distribution for this model ID
		addrlist, err := c.AIModels.GetModelDistribution(nil, currentId)
		if err != nil {
			return nil, fmt.Errorf("failed to get distribution for model %v: %v", i, err)
		}

		// Store in our map using int64 as key
		distri[i] = addrlist

		// fmt.Printf("Model ID: %v, Addresses: %v\n", i, addrlist)
	}

	return distri, nil
}

// GetAllEvents 获取指定区块范围内的所有事件
func (c *ConManager) GetAllEvents(aiModels *AIModels.AIModels, fromBlock *big.Int, toBlock *big.Int) ([]interface{}, error) {
	var events []interface{}

	// 创建过滤选项
	end := toBlock.Uint64()
	filterOpts := &bind.FilterOpts{
		Start:   fromBlock.Uint64(),
		End:     &end,
		Context: context.Background(),
	}

	// 1. 获取 ModelDeployed 事件
	modelDeployedIterator, err := aiModels.FilterModelDeployed(filterOpts, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("filter ModelDeployed events error: %v", err)
	}
	defer modelDeployedIterator.Close()

	for modelDeployedIterator.Next() {
		events = append(events, modelDeployedIterator.Event)
	}

	// 2. 获取 ModelRemoved 事件
	modelRemovedIterator, err := aiModels.FilterModelRemoved(filterOpts, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("filter ModelRemoved events error: %v", err)
	}
	defer modelRemovedIterator.Close()

	for modelRemovedIterator.Next() {
		events = append(events, modelRemovedIterator.Event)
	}

	// 3. 获取 UploadModeled 事件
	uploadModeledIterator, err := aiModels.FilterUploadModeled(filterOpts, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("filter UploadModeled events error: %v", err)
	}
	defer uploadModeledIterator.Close()

	for uploadModeledIterator.Next() {
		events = append(events, uploadModeledIterator.Event)
	}

	return events, nil
}
