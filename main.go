package main

import (
	"log"
	"os"

	"github.com/topaiagent/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "eth-cli",
		Usage:    "Ethereum command line interface",
		Commands: cmd.CommandSum,
		// Commands: []*cli.Command{
		// 	cmd.TransferCommand,
		// 	cmd.RegisterCommand,
		// 	cmd.ContractCommand(),
		// 	cmd.WalletCommand,
		// },
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
