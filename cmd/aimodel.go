package cmd

import (
	"fmt"
	"math/big"

	"github.com/topaiagent/con_manager"
	"github.com/urfave/cli/v2"
)

var AIModelCommands = &cli.Command{
	Name:  "aimodel",
	Usage: "aimodel command to operate the contract",
	Subcommands: []*cli.Command{
		AIModelCommand,
	},
	// Flags: []cli.Flag{
	// 	&cli.StringFlag{
	// 		Name:     "config",
	// 		Aliases:  []string{"c"}, // 命令简写
	// 		Usage:    "config path",
	// 		Value:    "~/.config/config.json",
	// 		Required: false,
	// 	},
	// 	&cli.StringFlag{
	// 		Name:     "rpc",
	// 		Usage:    "blockchain rpc url",
	// 		Required: true,
	// 	},
	// },
	// Action: func(c *cli.Context) error {

	// 	conMan, err := con_manager.NewConManager(c.String("rpc"))
	// 	if err != nil {
	// 		return err
	// 	}

	// 	addrlist, err := conMan.GetModelDeploymentMap(
	// 		big.NewInt(1000),
	// 	)

	// 	if err != nil {
	// 		return nil
	// 	}

	// 	fmt.Println(":\n%+v", addrlist)
	// 	return nil
	// },
}

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
