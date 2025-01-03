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

func (c *ConManager) GetModelDeploymentMap(modelId *big.Int) (map[*big.Int][]common.Address, error) {

	//get model list index
	distri := make(map[*big.Int][]common.Address)
	modelIds, err := c.AIModels.NextModelId(nil)
	if err != nil {
		return nil, err
	}

	for i := big.NewInt(0); i.Cmp(modelIds) < 0; i.Add(i, big.NewInt(1)) {
		addrlist, err := c.AIModels.GetModelDistribution(nil, i)
		if err != nil {
			return nil, err
		}
		distri[i] = addrlist
	}
	return distri, nil
}
