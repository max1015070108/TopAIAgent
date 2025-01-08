package cmd

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/max1015070108/topaiagent/con_manager"
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

		time.Sleep(5 * time.Second)
		recipt, err := conMan.Client.TransactionReceipt(context.Background(), tx.Hash())
		// recipt, ispending, err := conMan.Client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			return err
		}

		fmt.Printf("recipt:%+v, ispending:%+v", recipt, "ispending")
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
