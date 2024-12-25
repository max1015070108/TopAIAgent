// client/client.go
package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	Client *ethclient.Client
}

func NewEthClient(url string) (*EthClient, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &EthClient{Client: client}, nil
}

func (c *EthClient) Transfer(to string, amount *big.Int, privateKey string) error {
	// 实现转账逻辑
	return nil
}

func (c *EthClient) CallContract(contractAddr string, abi abi.ABI, method string, args ...interface{}) error {

	// 实现合约调用逻辑
	return nil
}
