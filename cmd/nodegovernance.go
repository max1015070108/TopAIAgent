package cmd

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/max1015070108/TopAIAgent/con_manager"
	"github.com/max1015070108/TopAIAgent/con_manager/NodesRegistry"
	"github.com/max1015070108/TopAIAgent/con_manager/NodesRegistryImpl"
	"github.com/urfave/cli/v2"
)

const ROUND_DURATION_TIME = 3600

var NodeGovernanceCmd = &cli.Command{

	Name:  "nodegovernance",
	Usage: "daemon Node to monitor",
	Subcommands: []*cli.Command{
		NodeGovernanceInitCmd,
		RegisterCommand,
		GetNodeInfoByAddrCmd,
		MonitorFilterNodeEventsCmd,
		WatchNodeEventsCmd,
	},
}

var NodeGovernanceInitCmd = &cli.Command{

	Name:  "nodegovernanceinit",
	Usage: "node governance initialize",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"c"}, // 命令简写
			Usage:    "config path",
			Value:    "~/.config/config.json",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "rpc",
			Usage:    "blockchain rpc url",
			Required: true,
		},

		//address string array
		&cli.StringSliceFlag{
			Name:    "identifiers",
			Aliases: []string{"if"},
			Usage:   "address slice for that",
			Value:   cli.NewStringSlice("0xc4ab424f86c9c9bafdc02b2d3fe0d97950c7dd17", "0x27763c1dafd6704223b286643b9ec596213d41fd"),
		},
	},
	Action: func(c *cli.Context) error {
		conMan, err := con_manager.NewConManager(c.String("rpc"))
		if err != nil {
			return err
		}

		identifiers := c.StringSlice("identifiers")
		identifiersAddr := []common.Address{}
		for _, value := range identifiers {

			fmt.Printf("if is....%+v\n", value)

			identifiersAddr = append(identifiersAddr, common.HexToAddress(value))
		}
		// fmt.Println(value)
		// wallets := c.StringSlice("identifiers")
		alias_identifiers := []string{"11111111111111111", "21111111111111111"}
		gpuTypes := [][]string{{"A100", "V100"}, {"A100", "V100"}}
		gpuNums := [][]*big.Int{
			{
				big.NewInt(2), big.NewInt(3)}, {big.NewInt(2), big.NewInt(3)},
		}

		fmt.Println("current address...", identifiersAddr[0])
		privateKeyECDSA, err := conMan.GetPrivateKeyByAddr(identifiersAddr[0])
		if err != nil {
			return err
		}
		auth, err := con_manager.CreateLatestAuth(conMan.Client, privateKeyECDSA, conMan.Conf.ContractAddress.NodeRegister)

		if err != nil {
			return err
		}

		// auth := conMan.
		// conMan.A
		tx, err := conMan.NodesGovernance.NodesGovernanceInitialize(

			auth,
			identifiersAddr,
			alias_identifiers,
			identifiersAddr,
			gpuTypes,
			gpuNums,
			identifiersAddr[0],
			big.NewInt(ROUND_DURATION_TIME),
		)

		if err != nil {
			return nil
		}

		fmt.Printf("%+v\n", tx.Hash().Hex())

		time.Sleep(5 * time.Second)
		recipt, err := conMan.Client.TransactionReceipt(context.Background(), tx.Hash())

		if err != nil {
			return err
		}

		fmt.Printf("recipt:%+v", recipt)
		return nil

	},
}

var RegisterCommand = &cli.Command{

	Name:  "register",
	Usage: "register node",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"c"}, // 命令简写
			Usage:    "config path",
			Value:    "~/.config/config.json",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "rpc",
			Usage:    "blockchain rpc url",
			Required: true,
		},

		//address string array
		&cli.StringSliceFlag{
			Name:    "identifiers",
			Aliases: []string{"if"},
			Usage:   "address slice for that",
			Value:   cli.NewStringSlice("0xc4ab424f86c9c9bafdc02b2d3fe0d97950c7dd17", "0x27763c1dafd6704223b286643b9ec596213d41fd"),
		},
	},
	Action: func(c *cli.Context) error {
		conMan, err := con_manager.NewConManager(c.String("rpc"))
		if err != nil {
			return err
		}

		identifiers := c.StringSlice("identifiers")
		identifiersAddr := []common.Address{}
		for _, value := range identifiers {

			fmt.Printf("if is....%+v\n", value)

			identifiersAddr = append(identifiersAddr, common.HexToAddress(value))
		}
		// fmt.Println(value)
		// wallets := c.StringSlice("identifiers")
		alias_identifiers := []string{"11111111111111111", "21111111111111111"}
		gpuTypes := [][]string{{"A100", "V100"}, {"A100", "V100"}}
		gpuNums := [][]*big.Int{
			{big.NewInt(2), big.NewInt(3)},
			{big.NewInt(2), big.NewInt(3)},
		}

		fmt.Println("current address...", identifiersAddr[0])
		fmt.Println("current contract...", conMan.Conf.ContractAddress.NodeRegister)
		privateKeyECDSA, err := conMan.GetPrivateKeyByAddr(identifiersAddr[0])
		if err != nil {
			return err
		}

		auth, err := con_manager.CreateLatestAuth(conMan.Client, privateKeyECDSA, conMan.Conf.ContractAddress.NodeRegister)

		if err != nil {
			return err
		}

		tx, err := conMan.NodesRegistry.RegisterNode(
			auth, identifiersAddr[0], alias_identifiers[0], gpuTypes[0], gpuNums[0],
		)
		if err != nil {
			return err
		}

		fmt.Printf("%+v\n", tx.Hash().Hex())

		receipt, err := bind.WaitMined(context.Background(), conMan.Client, tx)
		if err != nil {
			return fmt.Errorf("failed to wait for mining: %v", err)
		}

		fmt.Printf("receipt: %+v\n", receipt)
		// 检查交易是否成功
		if receipt.Status == 0 {
			return fmt.Errorf("transaction failed")
		}
		// time.Sleep(10 * time.Second)
		// recipt, err := conMan.Client.TransactionReceipt(context.Background(), tx.Hash())
		// // recipt, ispending, err := conMan.Client.TransactionByHash(context.Background(), tx.Hash())
		// if err != nil {
		// 	return err
		// }

		// fmt.Printf("recipt:%+v, ispending:%+v", recipt, "ispending")
		return nil

	},
}

var GetNodeInfoByAddrCmd = &cli.Command{

	Name:  "getnodeinfo",
	Usage: "get node info",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"c"}, // 命令简写
			Usage:    "config path",
			Value:    "~/.config/config.json",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "rpc",
			Usage:    "blockchain rpc url",
			Required: true,
		},

		//address string array
		&cli.StringFlag{
			Name:  "address",
			Usage: "address for get info",
		},
	},
	Action: func(c *cli.Context) error {
		conMan, err := con_manager.NewConManager(c.String("rpc"))
		if err != nil {
			return err
		}

		nodes, err := conMan.GetAllNodeInfo(c.Context)
		if err != nil {
			return err
		}
		fmt.Printf("nodes:%+v\n", nodes)
		return nil
	},
}

var MonitorFilterNodeEventsCmd = &cli.Command{

	Name:  "filterevents",
	Usage: "filterevents command to operate the contract",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"c"}, // 命令简写
			Usage:    "config path",
			Value:    "~/.config/config.json",
			Required: false,
		},

		&cli.Int64Flag{
			Name:     "fromblock",
			Aliases:  []string{"f"}, // 命令简写
			Usage:    "fromblock for filter event",
			Value:    0,
			Required: false,
		},

		&cli.Int64Flag{
			Name:     "endblock",
			Aliases:  []string{"e"}, // 命令简写
			Usage:    "endblock for filter event",
			Value:    0,
			Required: false,
		},

		&cli.StringFlag{
			Name:     "rpc",
			Usage:    "blockchain rpc url",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {

		conMan, err := con_manager.NewConManager(c.String("rpc"))
		if err != nil {
			return err
		}

		latestBlock, err := conMan.Client.BlockNumber(c.Context)
		if err != nil {
			return err
		}

		fmt.Println("latestBlock:", latestBlock)

		endBlock := big.NewInt(c.Int64("endblock"))
		if endBlock.Int64() == 0 {
			endBlock = big.NewInt(int64(latestBlock))
		}

		events, err := conMan.GetNodeEvents(common.HexToAddress(conMan.Conf.ContractAddress.NodeRegister), big.NewInt(c.Int64("fromblock")), endBlock)
		if err != nil {
			return err
		}

		// events, err := conMan.GetAllNodesRegistryEvents(big.NewInt(c.Int64("fromblock")), endBlock)
		// if err != nil {
		// 	return err
		// }
		fmt.Printf("how much event:%+v\n", events[0])

		for _, event := range events {
			switch e := event.(type) {
			case *NodesRegistryImpl.NodesRegistryImplAuthorized:
				fmt.Printf("Authorized - Owner: %s, Spender: %s\n",
					e.Owner.Hex(),
					e.Spender.Hex(),
				)
			case *NodesRegistryImpl.NodesRegistryImplNodeActived:
				fmt.Printf("NodeActived - Miner: %s, Identifier: %s, Time: %s, AliasIdentifier: %s\n",
					e.Wallet.Hex(),
					e.Identifier.Hex(),
					e.Time.String(),
					e.AliasIdentifier,
				)
			case *NodesRegistryImpl.NodesRegistryImplNodeDeregistered:
				fmt.Printf("NodeDeregistered - Identifier: %s, Time: %s, AliasIdentifier: %s\n",
					e.Identifier.Hex(),
					e.Time.String(),
					e.AliasIdentifier,
				)
			case *NodesRegistry.NodesRegistryNodeRegistered:
				fmt.Printf("NodeRegistered - Miner: %s, Identifier: %s, Time: %s, AliasIdentifier: %s\n",
					e.Wallet.Hex(),
					e.Identifier.Hex(),
					e.Time.String(),
					e.AliasIdentifier,
				)
			}
		}
		return nil
	},
}

var WatchNodeEventsCmd = &cli.Command{

	Name:  "watchevents",
	Usage: "watchevents command to operate the contract",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"c"}, // 命令简写
			Usage:    "config path",
			Value:    "~/.config/config.json",
			Required: false,
		},
		&cli.Int64Flag{
			Name:     "endblock",
			Aliases:  []string{"e"}, // 命令简写
			Usage:    "endblock for filter event",
			Value:    0,
			Required: false,
		},

		&cli.StringFlag{
			Name:     "rpc",
			Usage:    "blockchain rpc url",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {

		conMan, err := con_manager.NewConManager(c.String("rpc"))
		if err != nil {
			return err
		}

		eventChan := make(chan interface{})
		sub, err := conMan.WatchNodesRegistryEvents(eventChan)
		if err != nil {
			fmt.Printf("Watch error: %v\n", err)
			return err
		}
		defer sub.Unsubscribe()

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for {
				select {
				case event := <-eventChan:
					switch e := event.(type) {
					case *NodesRegistry.NodesRegistryNodeRegistered:
						fmt.Printf("New NodeRegistered - Miner: %s\n", e.Wallet.Hex())
					case *NodesRegistry.NodesRegistryNodeActived:
						fmt.Printf("New NodeActived - Miner: %s\n", e.Wallet.Hex())
						// ... 处理其他事件类型
					}
				case err := <-sub.Err():
					fmt.Printf("Watch error: %v\n", err)
					wg.Done()
					return
				}
			}

		}(&wg)

		wg.Wait()
		return nil
	},
}
