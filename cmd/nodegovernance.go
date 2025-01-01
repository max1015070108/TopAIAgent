package cmd

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/topaiagent/con_manager"
	"github.com/urfave/cli/v2"
)

var NodeGovernanceCmd = &cli.Command{

	Name:  "nodegovernance",
	Usage: "daemon Node to monitor",
	Subcommands: []*cli.Command{
		NodeGovernanceInitCmd,
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

		ROUND_DURATION_TIME := big.NewInt(3600)

		fmt.Println("current address...", identifiersAddr[0])
		privateKeyECDSA, err := conMan.GetPrivateKeyByAddr(identifiersAddr[0])
		if err != nil {
			return err
		}
		auth, err := con_manager.CreateLatestAuth(conMan.Client, privateKeyECDSA, c.String("address"))

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
			ROUND_DURATION_TIME,
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
