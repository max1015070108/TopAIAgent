package cmd

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/topaiagent/con_manager"
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
			Usage:    "reporter who report workload",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {

		conMan, err := con_manager.NewConManager(c.String("rpc"))
		if err != nil {
			return err
		}

		privList := []*ecdsa.PrivateKey{}
		addrlist := []string{}

		for _, key := range c.StringSlice("reporter") {

			privateKeyECDSA, err := conMan.GetPrivateKeyByAddr(common.HexToAddress(key))
			if err != nil {
				return err
			}
			privList = append(privList, privateKeyECDSA)
			addrlist = append(addrlist, key)
		}

		auth, err := con_manager.CreateLatestAuth(conMan.Client, privList[0], conMan.Conf.ContractAddress.AIWorkerload)
		if err != nil {
			return err
		}

		//mock data
		workload := big.NewInt(10)
		modelId := big.NewInt(1)
		sessionId := big.NewInt(1)
		epochID := big.NewInt(3)

		signatures, err := conMan.SignText(
			addrlist[0],
			workload,
			modelId,
			sessionId,
			epochID,
			privList,
		)

		tx, err := conMan.AIWorkload.ReportWorkload(
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
