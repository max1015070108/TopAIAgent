// cmd/transfer.go
package cmd

import (
	"github.com/urfave/cli/v2"
)

var TransferCommand = &cli.Command{

	Name:  "transfer",
	Usage: "Transfer ETH to address",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "to",
			Usage:    "Recipient address",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "amount",
			Usage:    "Amount to transfer in ETH",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		// 处理转账命令
		return nil
	},
}
