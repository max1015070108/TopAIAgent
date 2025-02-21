package cmd

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/max1015070108/TopAIAgent/con_manager"
	"github.com/urfave/cli/v2"
)

var AIModelCommands = &cli.Command{
	Name:  "aimodel",
	Usage: "aimodel command to operate the contract",
	Subcommands: []*cli.Command{
		GetModelDeploymentMapCmd,
		RecordUploadModelCmd,
		GetAllModelCmd,
	},
}

var GetModelDeploymentMapCmd = &cli.Command{

	Name:  "deploymap",
	Usage: "deploymap command to operate the contract",
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

		addrlist, err := conMan.GetModelDeploymentMap()

		if err != nil {
			return nil
		}

		fmt.Printf(":\n%+v\n", addrlist)
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
		auth, err := con_manager.CreateLatestAuth(conMan.Client, privateKeyECDSA, conMan.Conf.ContractAddress.AIModels)
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
			return err
		}
		fmt.Printf("%+v\n", tx.Hash().Hex())

		// conMan.AIModels.AIModelsFilterer.FilterUploadModeled(opts *bind.FilterOpts, modelId []*big.Int, uploader []common.Address)

		receipt, err := bind.WaitMined(context.Background(), conMan.Client, tx)
		if err != nil {
			return fmt.Errorf("failed to wait for mining: %v", err)
		}

		fmt.Printf("receipt: %+v\n", receipt)
		// 检查交易是否成功
		if receipt.Status == 0 {
			return fmt.Errorf("transaction failed")
		}

		return nil
	},
}

var ReportDeploymentCmd = &cli.Command{

	Name:  "reportdeploy",
	Usage: "reportdeploy command to operate the contract",
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
		auth, err := con_manager.CreateLatestAuth(conMan.Client, privateKeyECDSA, conMan.Conf.ContractAddress.AIModels)
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
			return err
		}
		fmt.Printf("%+v\n", tx.Hash().Hex())

		// conMan.AIModels.AIModelsFilterer.FilterUploadModeled(opts *bind.FilterOpts, modelId []*big.Int, uploader []common.Address)

		receipt, err := bind.WaitMined(context.Background(), conMan.Client, tx)
		if err != nil {
			return fmt.Errorf("failed to wait for mining: %v", err)
		}

		fmt.Printf("receipt: %+v\n", receipt)
		// 检查交易是否成功
		if receipt.Status == 0 {
			return fmt.Errorf("transaction failed")
		}

		return nil
	},
}

var GetAllModelCmd = &cli.Command{

	Name:  "getall",
	Usage: "getall models command to operate the contract",
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

		nextid, err := conMan.AIModels.NextModelId(nil)
		if err != nil {
			return err
		}

		var allModels []interface{}
		for i := int64(1); i <= nextid.Int64(); i++ {

			model, err := conMan.AIModels.UploadModels(nil, big.NewInt(i))
			if err != nil {
				return err
			}
			allModels = append(allModels, model)
		}

		fmt.Printf(":\n%+v\n", allModels)
		return nil
	},
}
