package cmd

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/topaiagent/con_manager"
	"github.com/topaiagent/con_manager/NodesRegistry"
	"github.com/topaiagent/wallet"
	"github.com/urfave/cli/v2"
)

var RegisterCommand = &cli.Command{

	Name:  "register",
	Usage: "register Node to address",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "contractaddress",
			Aliases:  []string{"c"}, // 命令简写
			Usage:    "contract address",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "rpc",
			Usage:    "blockchain rpc url",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "address",
			Aliases:  []string{"a"}, // 命令简写
			Usage:    "which address to sign the transaction",
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

		conAddress := common.HexToAddress(c.String("contractaddress"))

		client, err := ethclient.Dial(c.String("rpc"))
		if err != nil {
			return err
		}

		walletMan := wallet.NewWalletManager(c.String("keystore"))

		//unlock wallet
		account := common.HexToAddress(c.String("address"))

		aacount, err := walletMan.FindAccount(account)
		if err != nil {
			return err
		}

		err = walletMan.UnlockWallet(aacount, c.String("password"))
		if err != nil {
			return err
		}

		privateKey, err := walletMan.ExportPrivateKey(aacount, c.String("password"))
		if err != nil {
			return err
		}

		fmt.Printf("privateKey: %v\n", privateKey)
		//get addr private from keystore
		privateKeyECDSA, err := crypto.HexToECDSA(privateKey[2:])
		if err != nil {
			return err
		}
		auth, err := con_manager.CreateLatestAuth(client, privateKeyECDSA, c.String("address"))

		nodereg, err := NodesRegistry.NewNodesRegistry(conAddress, client)
		if err != nil {
			return err
		}

		transaction, err := nodereg.RegisterNode(auth, account, "", []string{""}, []*big.Int{big.NewInt(1)})
		if err != nil {
			return err
		}

		fmt.Printf("Transaction: %v\n", transaction.Hash().Hex())

		return nil
	},
}
