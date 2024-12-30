package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
)

var DaemonCommand = &cli.Command{

	Name:  "daemon",
	Usage: "daemon Node to monitor",
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
	},
	Action: func(c *cli.Context) error {

		conAddress := common.HexToAddress(c.String("contractaddress"))

		client, err := ethclient.Dial(c.String("rpc"))
		if err != nil {
			return err
		}

		// nodereg, err := NodesRegistry.NewNodesRegistry(conAddress, client)
		// if err != nil {
		// 	return err
		// }

		// fmt.Printf("Transaction: %v\n", transaction.Hash().Hex())

		return nil
	},
}
