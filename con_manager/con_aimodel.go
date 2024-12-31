package con_manager

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *ConManager) GetNodeDeployment(model_addrss string) ([]*big.Int, error) {

	nodes_deploys, err := c.AIModels.GetNodeDeployment(&bind.CallOpts{}, common.HexToAddress(model_addrss))
	if err != nil {
		return nil, err
	}
	return nodes_deploys, nil
}

func (c *ConManager) GetModelDeploymentMap(modelId *big.Int) ([]common.Address, error) {

	nodes_deploys, err := c.AIModels.GetModelDistribution(&bind.CallOpts{}, modelId)
	if err != nil {
		return nil, err
	}
	return nodes_deploys, nil
}
