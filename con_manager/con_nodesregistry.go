package con_manager

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (c *ConManager) GetLastEpoch() (*big.Int, error) {

	round_id, err := c.NodesGovernance.CurrentRoundId(&bind.CallOpts{})

	if err != nil {
		return nil, err
	}
	return round_id, nil
}

func (c *ConManager) GetAllNodeInfo() (*big.Int, error) {

	round_id, err := c.NodesGovernance.CurrentRoundId(&bind.CallOpts{})

	if err != nil {
		return nil, err
	}
	return round_id, nil
}
