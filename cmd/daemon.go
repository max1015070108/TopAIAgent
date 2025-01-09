package cmd

import (
	"github.com/max1015070108/TopAIAgent/con_manager"
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
		_, err := con_manager.NewConManager(c.String("rpc"))
		if err != nil {
			return err
		}

		// conMan.AIModels.FilterUploadModeled(&bind.FilterOpts{}, modelId []*big.Int, uploader []common.Address)
		// nodereg, err := NodesRegistry.NewNodesRegistry(conAddress, client)
		// if err != nil {
		// 	return err
		// }

		// fmt.Printf("Transaction: %v\n", transaction.Hash().Hex())

		return nil
	},
}
