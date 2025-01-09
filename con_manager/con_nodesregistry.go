package con_manager

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/max1015070108/TopAIAgent/con_manager/NodesRegistry"
)

// func (c *ConManager) GetAllNodeInfo(ctx context.Context) ([]NodesRegistry.NodesRegistryNode, error) {

// idLen, err := c.NodesRegistry.Length(nil)
// if err != nil {
// 	return nil, err
// }

// nodes := make([]NodesRegistry.NodesRegistryNode, idLen.Int64())
// for i := 0; i < int(idLen.Int64()); i++ {
// 	node, err := c.NodesRegistry.At(nil, big.NewInt(int64(i)))
// 	if err != nil {
// 		return nil, err
// 	}
// 	nodes[i] = node
// }
// return nodes, nil
// }

// 4. 定义结果结构
type nodeResult struct {
	index int64
	node  NodesRegistry.NodesRegistryNode
	err   error
}

func (c *ConManager) GetAllNodeInfo(ctx context.Context) ([]NodesRegistry.NodesRegistryNode, error) {
	// 1. 使用带超时的 context
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// 2. 获取长度
	idLen, err := c.NodesRegistry.Length(&bind.CallOpts{Context: ctxWithTimeout})
	if err != nil {
		return nil, fmt.Errorf("failed to get length: %w", err)
	}

	totalNodes := idLen.Int64()
	// nodes := make([]NodesRegistry.NodesRegistryNode, 0, totalNodes)

	// 3. 使用并发控制
	sem := make(chan struct{}, 5) // 限制并发数
	var wg sync.WaitGroup
	nodeChan := make(chan nodeResult, totalNodes)

	// 5. 并发获取节点
	for i := int64(0); i < totalNodes; i++ {
		wg.Add(1)
		go func(index int64) {
			defer wg.Done()

			// 限制并发
			sem <- struct{}{}
			defer func() { <-sem }()

			// 创建新的 context 避免连接复用问题
			reqCtx, reqCancel := context.WithTimeout(ctx, 5*time.Second)
			defer reqCancel()

			// 使用独立的 CallOpts
			opts := &bind.CallOpts{Context: reqCtx}
			node, err := c.NodesRegistry.At(opts, big.NewInt(index))

			nodeChan <- nodeResult{
				index: index,
				node:  node,
				err:   err,
			}
		}(i)
	}

	// 6. 等待所有请求完成
	go func() {
		wg.Wait()
		close(nodeChan)
	}()

	// 7. 收集结果
	results := make([]NodesRegistry.NodesRegistryNode, totalNodes)
	for result := range nodeChan {
		if result.err != nil {
			return nil, fmt.Errorf("failed to get node at %d: %w", result.index, result.err)
		}
		results[result.index] = result.node
	}

	return results, nil
}
