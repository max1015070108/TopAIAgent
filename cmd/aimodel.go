package cmd

import (
	"fmt"
	"math/big"

	"github.com/topaiagent/con_manager"
	"github.com/urfave/cli/v2"
)

var AIModelCommand = &cli.Command{

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
		&cli.StringFlag{
			Name:     "keystore",
			Usage:    "topaiagent keystore",
			Value:    "~/.topaiagent",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "password",
			Usage:    "password for the wallet",
			Value:    "101507",
			Required: false,
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

		// conAddress := common.HexToAddress(c.String("contractaddress"))

		// client, err := ethclient.Dial(c.String("rpc"))
		// if err != nil {
		// 	return err
		// }

		// walletMan := wallet.NewWalletManager(c.String("keystore"))

		// //unlock wallet
		// account := common.HexToAddress(c.String("address"))

		// aacount, err := walletMan.FindAccount(account)
		// if err != nil {
		// 	return err
		// }

		// err = walletMan.UnlockWallet(aacount, c.String("password"))
		// if err != nil {
		// 	return err
		// }

		// privateKey, err := walletMan.ExportPrivateKey(aacount, c.String("password"))
		// if err != nil {
		// 	return err
		// }

		// fmt.Printf("privateKey: %v\n", privateKey)
		// //get addr private from keystore
		// privateKeyECDSA, err := crypto.HexToECDSA(privateKey[2:])
		// if err != nil {
		// 	return err
		// }
		// auth, err := con_manager.CreateLatestAuth(client, privateKeyECDSA, c.String("address"))

		// nodereg, err := NodesRegistry.NewNodesRegistry(conAddress, client)
		// if err != nil {
		// 	return err
		// }

		// transaction, err := nodereg.RegisterNode(auth, account, "", []string{""}, []*big.Int{big.NewInt(1)})
		// if err != nil {
		// 	return err
		// }

		// fmt.Printf("Transaction: %v\n", transaction.Hash().Hex())

		return nil
	},
}
