package cmd

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/max1015070108/TopAIAgent/con_manager"
	"github.com/urfave/cli/v2"
)

var AIWorkLoadCommands = &cli.Command{
	Name:  "aiworkload",
	Usage: "aiworkload command to operate the contract",
	Subcommands: []*cli.Command{
		ReportWorkLoadCmd,
	},
}

var ReportWorkLoadCmd = &cli.Command{

	Name:  "reportworkload",
	Usage: "report work load command to operate the contract",
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
		&cli.StringSliceFlag{
			Name:     "reporter",
			Aliases:  []string{"rp"},
			Usage:    "reporter who sign workload",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "proxy",
			Aliases:  []string{"p"},
			Usage:    "proxy who report workload",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "worker",
			Aliases:  []string{"w"},
			Usage:    "reporter who report workload",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "user",
			Aliases:  []string{"u"},
			Usage:    "reporter who report workload",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {

		conMan, err := con_manager.NewConManager(c.String("rpc"))
		if err != nil {
			return err
		}

		// privList := []*ecdsa.PrivateKey{}
		// addrlist := []string{}

		// for _, key := range c.StringSlice("reporter") {

		// 	privateKeyECDSA, err := conMan.GetPrivateKeyByAddr(common.HexToAddress(key))
		// 	if err != nil {
		// 		return err
		// 	}
		// 	privList = append(privList, privateKeyECDSA)
		// 	addrlist = append(addrlist, key)
		// }

		// auth, err := con_manager.CreateLatestAuth(conMan.Client, privList[0], conMan.Conf.ContractAddress.AIWorkerload)
		// if err != nil {
		// 	return err
		// }

		// //mock data
		workload := big.NewInt(10)
		modelId := big.NewInt(1)
		sessionId := big.NewInt(1)
		epochID := big.NewInt(50)

		tx, err := conMan.ReportWorkload(
			c.StringSlice("reporter"),
			c.String("proxy"),
			c.String("worker"),
			c.String("user"),
			workload,
			modelId,
			sessionId,
			epochID,
		)

		if err != nil {
			return err
		}

		receipt, err := bind.WaitMined(context.Background(), conMan.Client, tx)
		if err != nil {
			return fmt.Errorf("failed to wait for mining: %v", err)
		}

		fmt.Printf("receipt:%+v\n", receipt)
		return nil
	},
}
