package cmd

import (
	"fmt"

	"github.com/max1015070108/TopAIAgent/wallet"
	"github.com/urfave/cli/v2"
)

var WalletCommand = &cli.Command{
	Name:  "wallet",
	Usage: "Manage wallet",
	Subcommands: []*cli.Command{
		NewCommand,
		ListCommand,
		BalanceCommand,
		ImportCommand,
	},
}

var NewCommand = &cli.Command{
	Name:  "new",
	Usage: "Create a new wallet",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "keystore",
			Usage:    "topaiagent keystore",
			Value:    "~/.topaiagent",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "password",
			Usage:    "password for the wallet",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {

		wallet := wallet.NewWalletManager(c.String("keystore"))

		account, err := wallet.CreateWallet(c.String("password"))
		if err != nil {
			return err
		}

		fmt.Printf("New account: %s\n", account.Address.Hex())
		return nil
	},
}

var ListCommand = &cli.Command{
	Name:  "list",
	Usage: "list all the address",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "keystore",
			Usage:    "topaiagent keystore",
			Value:    "~/.topaiagent",
			Required: false,
		},
	},
	Action: func(c *cli.Context) error {
		wallet := wallet.NewWalletManager(c.String("keystore"))
		addrlist := wallet.ListWalletAddresses()

		// fmt.Printf("Address list: %v\n", addrlist)

		for _, addr := range addrlist {
			fmt.Printf("Address: %s\n", addr)
		}
		return nil
	},
}

var BalanceCommand = &cli.Command{
	Name:  "balance",
	Usage: "balance of the address",
	Action: func(c *cli.Context) error {
		return nil
	},
}

var ImportCommand = &cli.Command{
	Name:  "importPrivate",
	Usage: "balance of the address",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "keystore",
			Aliases: []string{
				"k",
			},
			Usage:    "topaiagent keystore",
			Value:    "~/.topaiagent",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "password",
			Usage:    "password for the wallet",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "privateKey",
			Usage:    "private key",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		wallet := wallet.NewWalletManager(c.String("keystore"))
		aaccount, err := wallet.ImportWalletFromPrivateKey(c.String("privateKey"), c.String("password"))
		if err != nil {
			return err
		}

		fmt.Printf("Address: %s\n", aaccount.Address.Hex())
		return nil
	},
}
