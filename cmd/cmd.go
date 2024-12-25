package cmd

import "github.com/urfave/cli/v2"

var CommandSum = []*cli.Command{
	WalletCommand,
	RegisterCommand,
}
