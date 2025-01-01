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

var NodeInitialCommand = &cli.Command{

	Name:  "nodesinitial",
	Usage: "nodesinitial command to initial the codes",
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

		//address string array
		&cli.StringSliceFlag{
			Name:    "identifiers",
			Aliases: []string{"if"},
			Usage:   "address slice for that",
			Value:   cli.NewStringSlice("0xc4ab424f86c9c9bafdc02b2d3fe0d97950c7dd17", "0x27763c1dafd6704223b286643b9ec596213d41fd"),
		},
	},
	Action: func(c *cli.Context) error {

		conMan, err := con_manager.NewConManager(c.String("rpc"))
		if err != nil {
			return err
		}

		identifiers := c.StringSlice("identifiers")
		wallets := c.StringSlice("identifiers")
		alias_identifiers := []string{"11111111111111111", "21111111111111111"}
		gpuTypes := [][]string{{"A100", "V100"}, {"A100", "V100"}}
		gpuNums := [][]int{{2, 3}, {2, 3}}

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
