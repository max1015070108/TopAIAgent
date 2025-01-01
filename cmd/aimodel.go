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

var AIModelCommands = &cli.Command{
	Name:  "aimodel",
	Usage: "aimodel command to operate the contract",
	Subcommands: []*cli.Command{
		GetModelDeploymentMapCmd,
		RecordUploadModelCmd,
		UploadModelCmd,
	},
}

var GetModelDeploymentMapCmd = &cli.Command{

	Name:  "aimodel",
	Usage: "aimodel command to operate the contract",
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
	},
	Action: func(c *cli.Context) error {

		conMan, err := con_manager.NewConManager(c.String("rpc"))
		if err != nil {
			return err
		}

		addrlist, err := conMan.GetModelDeploymentMap(
			big.NewInt(1000),
		)

		if err != nil {
			return nil
		}

		fmt.Println(":\n%+v", addrlist)
		return nil
	},
}

var RecordUploadModelCmd = &cli.Command{

	Name:  "recordupload",
	Usage: "recordupload command to operate the contract",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"c"}, // 命令简写
			Usage:    "config path",
			Value:    "~/.config/config.json",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "modelname",
			Usage:    "model name for upload",
			Value:    "TestModel",
			Required: false,
		},

		&cli.StringFlag{
			Name:     "modelversion",
			Usage:    "model version for upload",
			Value:    "v1.0",
			Required: false,
		},

		&cli.StringFlag{
			Name:     "modelinfo",
			Usage:    "model info for upload",
			Value:    "Test model description",
			Required: false,
		},

		&cli.StringFlag{
			Name:     "uploader",
			Usage:    "who will upload this info",
			Value:    "0xC4AB424f86c9C9bAfDc02B2D3fE0d97950c7dd17",
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

		privateKeyECDSA, err := conMan.GetPrivateKeyByAddr(common.HexToAddress(c.String("uploader")))
		if err != nil {
			return err
		}
		auth, err := con_manager.CreateLatestAuth(conMan.Client, privateKeyECDSA, c.String("address"))

		if err != nil {
			return err
		}

		tx, err := conMan.AIModels.RecordModelUpload(
			auth,
			c.String("modelname"),
			c.String("modelversion"),
			c.String("modelinfo"),
		)

		if err != nil {
			return nil
		}

		fmt.Printf("%+v\n", tx.Hash().Hex())
		return nil
	},
}

var UploadModelCmd = &cli.Command{

	Name:  "upload",
	Usage: "upload command to operate the contract",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"c"}, // 命令简写
			Usage:    "config path",
			Value:    "~/.config/config.json",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "modelname",
			Usage:    "model name for upload",
			Value:    "TestModel",
			Required: false,
		},

		&cli.StringFlag{
			Name:     "modelversion",
			Usage:    "model version for upload",
			Value:    "v1.0",
			Required: false,
		},

		&cli.StringFlag{
			Name:     "modelinfo",
			Usage:    "model info for upload",
			Value:    "Test model description",
			Required: false,
		},

		&cli.StringFlag{
			Name:     "uploader",
			Usage:    "who will upload this info",
			Value:    "0xC4AB424f86c9C9bAfDc02B2D3fE0d97950c7dd17",
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

		privateKeyECDSA, err := conMan.GetPrivateKeyByAddr(common.HexToAddress(c.String("uploader")))
		if err != nil {
			return err
		}
		auth, err := con_manager.CreateLatestAuth(conMan.Client, privateKeyECDSA, c.String("address"))

		if err != nil {
			return err
		}

		//get next model id
		nid, err := conMan.AIModels.NextModelId(nil)

		fmt.Println("next model id:", nid)

		tx, err := conMan.AIModels.RecordModelUpload(
			auth,
			c.String("modelname"),
			c.String("modelversion"),
			c.String("modelinfo"),
		)

		if err != nil {
			return nil
		}

		fmt.Printf("%+v\n", tx.Hash().Hex())

		time.Sleep(5 * time.Second)
		tx_out, _, err := conMan.Client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			return err
		}

		fmt.Print("tx:%+v", tx_out)
		return nil
	},
}
