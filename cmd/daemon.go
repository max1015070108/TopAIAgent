package cmd

import (
	"sync"

	"github.com/max1015070108/TopAIAgent/con_manager"
	"github.com/max1015070108/TopAIAgent/database"
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
			Required: false,
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
			Required: false,
		},
	},
	Action: func(c *cli.Context) error {
		conMan, err := con_manager.NewConManager(c.String("rpc"))
		if err != nil {
			return err
		}

		conMan.DataEvent, err = database.NewEventStore() // DBName would be used for sqlite3
		//init database
		if err = conMan.DataEvent.InitTables(); err != nil {
			return err
		}

		var wg sync.WaitGroup
		wg.Add(1)
		conMan.WatchEvents(c.Context, &wg)

		wg.Wait()
		return nil
	},
}
