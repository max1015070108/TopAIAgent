package con_manager

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/max1015070108/TopAIAgent/con_manager/NodesRegistry"
	"github.com/max1015070108/TopAIAgent/con_manager/NodesRegistryImpl"
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

// GetAllNodesRegistryEvents 获取所有节点注册相关事件
func (c *ConManager) GetAllNodesRegistryEvents(fromBlock, toBlock *big.Int) ([]interface{}, error) {
	var events []interface{}

	end := toBlock.Uint64()
	// 创建过滤选项
	filterOpts := &bind.FilterOpts{
		Start:   fromBlock.Uint64(),
		End:     &end,
		Context: context.Background(),
	}

	// 1. 获取 Authorized 事件
	authorizedIterator, err := c.NodesRegistry.FilterAuthorized(filterOpts, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("filter Authorized events error: %v", err)
	}
	defer authorizedIterator.Close()

	for authorizedIterator.Next() {
		events = append(events, authorizedIterator.Event)
	}

	// 2. 获取 NodeActived 事件
	nodeActivedIterator, err := c.NodesRegistry.FilterNodeActived(filterOpts, nil)
	if err != nil {
		return nil, fmt.Errorf("filter NodeActived events error: %v", err)
	}
	defer nodeActivedIterator.Close()

	for nodeActivedIterator.Next() {
		events = append(events, nodeActivedIterator.Event)
	}

	// 3. 获取 NodeDeregistered 事件
	nodeDeregisteredIterator, err := c.NodesRegistry.FilterNodeDeregistered(filterOpts, nil)
	if err != nil {
		return nil, fmt.Errorf("filter NodeDeregistered events error: %v", err)
	}
	defer nodeDeregisteredIterator.Close()

	for nodeDeregisteredIterator.Next() {
		events = append(events, nodeDeregisteredIterator.Event)
	}

	// 4. 获取 NodeRegistered 事件
	nodeRegisteredIterator, err := c.NodesRegistry.FilterNodeRegistered(filterOpts, nil)
	if err != nil {
		return nil, fmt.Errorf("filter NodeRegistered events error: %v", err)
	}
	defer nodeRegisteredIterator.Close()

	for nodeRegisteredIterator.Next() {
		events = append(events, nodeRegisteredIterator.Event)
	}

	return events, nil
}

// WatchNodesRegistryEvents 监听所有节点注册相关事件
func (c *ConManager) WatchNodesRegistryEvents(sink chan<- interface{}) (event.Subscription, error) {
	// 创建监听选项
	watchOpts := &bind.WatchOpts{
		Context: context.Background(),
	}

	// 创建错误通道
	errChan := make(chan error)

	// 1. 监听 Authorized 事件
	authorizedSink := make(chan *NodesRegistry.NodesRegistryAuthorized)
	authorizedSub, err := c.NodesRegistry.WatchAuthorized(
		watchOpts,
		authorizedSink,
		nil, // owner
		nil, // spender
	)
	if err != nil {
		return nil, err
	}

	// 2. 监听 NodeActived 事件
	nodeActivedSink := make(chan *NodesRegistry.NodesRegistryNodeActived)
	nodeActivedSub, err := c.NodesRegistry.WatchNodeActived(
		watchOpts,
		nodeActivedSink,
		nil, // miner
	)
	if err != nil {
		return nil, err
	}

	// 3. 监听 NodeDeregistered 事件
	nodeDeregisteredSink := make(chan *NodesRegistry.NodesRegistryNodeDeregistered)
	nodeDeregisteredSub, err := c.NodesRegistry.WatchNodeDeregistered(
		watchOpts,
		nodeDeregisteredSink,
		nil, // identifier
	)
	if err != nil {
		return nil, err
	}

	// 4. 监听 NodeRegistered 事件
	nodeRegisteredSink := make(chan *NodesRegistry.NodesRegistryNodeRegistered)
	nodeRegisteredSub, err := c.NodesRegistry.WatchNodeRegistered(
		watchOpts,
		nodeRegisteredSink,
		nil, // miner
	)
	if err != nil {
		return nil, err
	}

	// 启动事件处理
	go func() {
		defer authorizedSub.Unsubscribe()
		defer nodeActivedSub.Unsubscribe()
		defer nodeDeregisteredSub.Unsubscribe()
		defer nodeRegisteredSub.Unsubscribe()

		for {
			select {
			case event := <-authorizedSink:
				sink <- event
			case event := <-nodeActivedSink:
				sink <- event
			case event := <-nodeDeregisteredSink:
				sink <- event
			case event := <-nodeRegisteredSink:
				sink <- event
			case err := <-authorizedSub.Err():
				errChan <- err
				return
			case err := <-nodeActivedSub.Err():
				errChan <- err
				return
			case err := <-nodeDeregisteredSub.Err():
				errChan <- err
				return
			case err := <-nodeRegisteredSub.Err():
				errChan <- err
				return
			}
		}
	}()

	return event.NewSubscription(func(quit <-chan struct{}) error {
		select {
		case <-quit:
			return nil
		case err := <-errChan:
			return err
		}
	}), nil
}

// GetNodeEvents 获取特定节点的所有相关事件
func (c *ConManager) GetNodeEvents(
	nodeAddr common.Address,
	fromBlock, toBlock *big.Int,
) ([]interface{}, error) {
	var events []interface{}

	end := toBlock.Uint64()
	filterOpts := &bind.FilterOpts{
		Start:   fromBlock.Uint64(),
		End:     &end,
		Context: context.Background(),
	}

	// 获取节点相关的各类事件
	// 1. NodeActived 事件
	activedEvents, err := c.NodesRegistry.FilterNodeActived(
		filterOpts,
		// []common.Address{nodeAddr},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("filter NodeActived events error: %v", err)
	}
	defer activedEvents.Close()

	for activedEvents.Next() {
		events = append(events, activedEvents.Event)
	}

	// 2. NodeDeregistered 事件
	deregEvents, err := c.NodesRegistry.FilterNodeDeregistered(
		filterOpts,
		// []common.Address{nodeAddr},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("filter NodeDeregistered events error: %v", err)
	}
	defer deregEvents.Close()

	for deregEvents.Next() {
		events = append(events, deregEvents.Event)
	}

	// 3. NodeRegistered 事件
	regEvents, err := c.NodesRegistry.FilterNodeRegistered(
		filterOpts,
		// []common.Address{nodeAddr},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("filter NodeRegistered events error: %v", err)
	}
	defer regEvents.Close()

	for regEvents.Next() {
		events = append(events, regEvents.Event)
	}

	return events, nil
}

// 使用示例
func (c *ConManager) ExampleNodeRegistryEvents() {
	// 1. 获取历史事件
	fromBlock := big.NewInt(0)
	toBlock := big.NewInt(1000000)

	events, err := c.GetAllNodesRegistryEvents(fromBlock, toBlock)
	if err != nil {
		fmt.Printf("Get events error: %v\n", err)
		return
	}

	for _, event := range events {
		switch e := event.(type) {
		case *NodesRegistryImpl.NodesRegistryImplAuthorized:
			fmt.Printf("Authorized - Owner: %s, Spender: %s\n",
				e.Owner.Hex(),
				e.Spender.Hex(),
			)
		case *NodesRegistryImpl.NodesRegistryImplNodeActived:
			fmt.Printf("NodeActived - Miner: %s, Identifier: %s\n",
				e.Miner.Hex(),
				e.Identifier.Hex(),
			)
		case *NodesRegistryImpl.NodesRegistryImplNodeDeregistered:
			fmt.Printf("NodeDeregistered - Identifier: %s, Time: %s\n",
				e.Identifier.Hex(),
				e.Time.String(),
			)
		case *NodesRegistryImpl.NodesRegistryImplNodeRegistered:
			fmt.Printf("NodeRegistered - Miner: %s, Identifier: %s\n",
				e.Miner.Hex(),
				e.Identifier.Hex(),
			)
		}
	}

	// 2. 监听实时事件
	eventChan := make(chan interface{})
	sub, err := c.WatchNodesRegistryEvents(eventChan)
	if err != nil {
		fmt.Printf("Watch error: %v\n", err)
		return
	}
	defer sub.Unsubscribe()

	go func() {
		for {
			select {
			case event := <-eventChan:
				switch e := event.(type) {
				case *NodesRegistryImpl.NodesRegistryImplNodeRegistered:
					fmt.Printf("New NodeRegistered - Miner: %s\n", e.Miner.Hex())
				case *NodesRegistryImpl.NodesRegistryImplNodeActived:
					fmt.Printf("New NodeActived - Miner: %s\n", e.Miner.Hex())
					// ... 处理其他事件类型
				}
			case err := <-sub.Err():
				fmt.Printf("Watch error: %v\n", err)
				return
			}
		}
	}()

	// 3. 获取特定节点的事件
	nodeAddr := common.HexToAddress("0x...")
	nodeEvents, err := c.GetNodeEvents(nodeAddr, fromBlock, toBlock)
	if err != nil {
		fmt.Printf("Get node events error: %v\n", err)
		return
	}

	for _, event := range nodeEvents {
		// 处理节点相关事件...
		fmt.Printf("Node Event: %+v\n", event)
	}
}
