package con_manager

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/topaiagent/con_manager/AIModels"
)

func (c *ConManager) GetAIModelContract(contract string, client *ethclient.Client) (*AIModels.AIModels, error) {

	aim_contract, err := AIModels.NewAIModels(common.HexToAddress(contract), client)
	if err != nil {
		return nil, err
	}
	return aim_contract, nil
}
