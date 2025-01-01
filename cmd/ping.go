package cmd

import (
	"fmt"

	"github.com/topaiagent/con_manager"
	"github.com/urfave/cli/v2"
)

var PingTestCmd = &cli.Command{

	Name:  "getBlock",
	Usage: "block info",
	Flags: []cli.Flag{

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
			Value:    "123456",
			Required: false,
		},
	},
	Action: func(c *cli.Context) error {

		conMan, err := con_manager.NewConManager(c.String("rpc"))
		if err != nil {
			return err
		}

		// conMan.Client.SubscribeNewHead(, ch chan<- *types.Header)

		roundID, err := conMan.GetLastEpoch()

		if err != nil {
			return nil
		}

		fmt.Printf("current ID:%+v\n", roundID)

		return nil
	},
}

var HeadBlockCmd = &cli.Command{

	Name:  "gethead",
	Usage: "block info",
	Flags: []cli.Flag{

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

		head, err := conMan.GetLastEpoch()
		if err != nil {
			return err
		}

		fmt.Printf("Headdata:%+v", head)
		return nil
	},
}
