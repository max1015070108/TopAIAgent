package con_manager

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/max1015070108/TopAIAgent/con_manager/AIWorkload"
)

func (c *ConManager) GetAIModelContract(
	worker string,
	workload, modelId, sessionId, epochId *big.Int,
	aiws []AIWorkload.Signature,
) (string, error) {

	account := common.HexToAddress(worker)

	aacount, err := c.Wallet.FindAccount(account)
	if err != nil {
		return "", err
	}

	err = c.Wallet.UnlockWallet(aacount, "123456")
	if err != nil {
		return "", err
	}

	privateKey, err := c.Wallet.ExportPrivateKey(aacount, "123456")
	if err != nil {
		return "", err
	}

	// fmt.Printf("privateKey: %v\n", privateKey)
	//get addr private from keystore
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey[2:])
	if err != nil {
		return "", err
	}

	auth, err := CreateLatestAuth(c.Client, privateKeyECDSA, c.Conf.ContractAddress.AIModels)
	if err != nil {
		return "", err
	}
	transaction, err := c.AIWorkload.ReportWorkload(
		auth,
		common.HexToAddress(worker),
		workload,
		modelId,
		sessionId,
		epochId,
		aiws,
	)
	if err != nil {
		return "", err
	}

	fmt.Println("tx....:", transaction.Hash().Hex())

	// return aim_contract, nil
	return transaction.Hash().Hex(), nil
}

// retrieve private key

func (c *ConManager) GetPrivateKeyByAddr(addr common.Address) (*ecdsa.PrivateKey, error) {

	// comAddr := common.HexToAddress(addr)

	aacount, err := c.Wallet.FindAccount(addr)
	if err != nil {
		return nil, err
	}

	err = c.Wallet.UnlockWallet(aacount, "123456")
	if err != nil {
		return nil, err
	}

	privateKey, err := c.Wallet.ExportPrivateKey(aacount, "123456")
	if err != nil {
		return nil, err
	}

	// fmt.Printf("privateKey: %v\n", privateKey)
	//get addr private from keystore
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey[2:])
	if err != nil {
		return nil, err
	}

	return privateKeyECDSA, nil
}

func (c *ConManager) ReportWorkload(reporters []string, workload, modelId, sessionId, epochID *big.Int) error {

	privList := []*ecdsa.PrivateKey{}
	addrlist := []string{}

	for _, key := range reporters {

		privateKeyECDSA, err := c.GetPrivateKeyByAddr(common.HexToAddress(key))
		if err != nil {
			return err
		}
		privList = append(privList, privateKeyECDSA)
		addrlist = append(addrlist, key)
	}

	auth, err := CreateLatestAuth(c.Client, privList[0], c.Conf.ContractAddress.AIWorkerload)
	if err != nil {
		return err
	}

	//mock data
	// workload := big.NewInt(10)
	// modelId := big.NewInt(1)
	// sessionId := big.NewInt(1)
	// epochID := big.NewInt(2)

	signatures, err := c.SignText(
		addrlist[0],
		workload,
		modelId,
		sessionId,
		epochID,
		privList,
	)

	tx, err := c.AIWorkload.ReportWorkload(
		auth,
		common.HexToAddress(addrlist[0]),
		workload,
		modelId,
		sessionId,
		epochID,
		signatures,
	)

	if err != nil {
		return err
	}
	fmt.Printf("tx.Hash: %+v\n", tx.Hash().Hex())

	time.Sleep(2 * time.Second)
	recipt, err := c.Client.TransactionReceipt(context.Background(), tx.Hash())
	// recipt, ispending, err := conMan.Client.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		return err
	}

	fmt.Printf("recipt:%+v", recipt)
	return nil
}

// GetAllWorkloadEvents 获取所有AIWorkload相关事件
func (c *ConManager) GetAllWorkloadEvents(fromBlock, toBlock *big.Int) ([]interface{}, error) {
	var events []interface{}

	end := toBlock.Uint64()
	// 创建过滤选项
	filterOpts := &bind.FilterOpts{
		Start:   fromBlock.Uint64(),
		End:     &end,
		Context: context.Background(),
	}

	// 获取WorkloadReported事件
	workloadIterator, err := c.AIWorkload.FilterWorkloadReported(filterOpts, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("filter WorkloadReported events error: %v", err)
	}
	defer workloadIterator.Close()

	for workloadIterator.Next() {
		events = append(events, workloadIterator.Event)
	}

	return events, nil
}

// WatchWorkloadEvents 监听AIWorkload合约事件
func (c *ConManager) WatchWorkloadEvents(sink chan<- interface{}) (event.Subscription, error) {
	// 创建监听选项
	watchOpts := &bind.WatchOpts{
		Context: context.Background(),
	}

	// 创建错误通道
	errChan := make(chan error)

	// 监听WorkloadReported事件
	workloadSink := make(chan *AIWorkload.AIWorkloadWorkloadReported)
	workloadSub, err := c.AIWorkload.WatchWorkloadReported(
		watchOpts,
		workloadSink,
		nil, // reporter
		nil, // modelId
	)
	if err != nil {
		return nil, err
	}

	// 启动事件处理
	go func() {
		defer workloadSub.Unsubscribe()

		for {
			select {
			case event := <-workloadSink:
				sink <- event
			case err := <-workloadSub.Err():
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

// 获取指定地址的工作量事件
func (c *ConManager) GetWorkerWorkloadEvents(
	worker common.Address,
	fromBlock, toBlock *big.Int,
) ([]*AIWorkload.AIWorkloadWorkloadReported, error) {
	var events []*AIWorkload.AIWorkloadWorkloadReported

	end := toBlock.Uint64()
	filterOpts := &bind.FilterOpts{
		Start:   fromBlock.Uint64(),
		End:     &end,
		Context: context.Background(),
	}

	// 过滤指定worker地址的事件
	iterator, err := c.AIWorkload.FilterWorkloadReported(
		filterOpts,
		nil,
		[]common.Address{worker}, // reporter
	)
	if err != nil {
		return nil, fmt.Errorf("filter worker events error: %v", err)
	}
	defer iterator.Close()

	for iterator.Next() {
		events = append(events, iterator.Event)
	}

	return events, nil
}

// 使用示例
func (c *ConManager) ExampleUsage() {
	// 1. 获取历史事件
	fromBlock := big.NewInt(0)
	toBlock := big.NewInt(1000000)

	events, err := c.GetAllWorkloadEvents(fromBlock, toBlock)
	if err != nil {
		fmt.Printf("Get events error: %v\n", err)
		return
	}

	for _, event := range events {
		if e, ok := event.(*AIWorkload.AIWorkloadWorkloadReported); ok {
			fmt.Printf("WorkloadReported - Reporter: %s, Workload: %s, ModelId: %s\n",
				e.Reporter.Hex(),
				e.Workload.String(),
				e.ModelId.String(),
			)
		}
	}

	// 2. 实时监听事件
	eventChan := make(chan interface{})
	sub, err := c.WatchWorkloadEvents(eventChan)
	if err != nil {
		fmt.Printf("Watch error: %v\n", err)
		return
	}
	defer sub.Unsubscribe()

	go func() {
		for {
			select {
			case event := <-eventChan:
				if e, ok := event.(*AIWorkload.AIWorkloadWorkloadReported); ok {
					fmt.Printf("New WorkloadReported - Reporter: %s, Workload: %s\n",
						e.Reporter.Hex(),
						e.Workload.String(),
					)
				}
			case err := <-sub.Err():
				fmt.Printf("Watch error: %v\n", err)
				return
			}
		}
	}()

	// 3. 获取特定worker的事件
	workerAddr := common.HexToAddress("0x...")
	workerEvents, err := c.GetWorkerWorkloadEvents(workerAddr, fromBlock, toBlock)
	if err != nil {
		fmt.Printf("Get worker events error: %v\n", err)
		return
	}

	for _, event := range workerEvents {
		fmt.Printf("Worker Workload - Workload: %s, ModelId: %s\n",
			event.Workload.String(),
			event.ModelId.String(),
		)
	}
}
