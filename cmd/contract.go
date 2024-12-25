// cmd/contract.go
package cmd

import (
	"github.com/urfave/cli/v2"
	// "github.com/ethereum/go-ethereum"
	//   "github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func ContractCommand() *cli.Command {
	return &cli.Command{
		Name:  "contract",
		Usage: "Call smart contract",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "address",
				Usage:    "Contract address",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "method",
				Usage:    "Contract method to call",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			// 处理合约调用命令
			// client_obj := client.NewEthClient("http://localhost:8545")

			// contractAddr := c.String("address")
			// method := c.String("method")

			// client_obj.CallContract(contractAddr, method)

			return nil
		},
	}
}
