package con_manager

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/topaiagent/con_manager/AIWorkload"
)

func (c *ConManager) GetAIModelContract(
	worker string,
	workload, modelId, sessionId, epochId *big.Int,
	aiws []AIWorkload.Signature,
) (string, error) {

	account := common.HexToAddress(worker)

	aacount, err := c.Wallet.FindAccount(account)
	if err != nil {
		return "", err
	}

	err = c.Wallet.UnlockWallet(aacount, "123456")
	if err != nil {
		return "", err
	}

	privateKey, err := c.Wallet.ExportPrivateKey(aacount, "123456")
	if err != nil {
		return "", err
	}

	fmt.Printf("privateKey: %v\n", privateKey)
	//get addr private from keystore
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey[2:])
	if err != nil {
		return "", err
	}

	auth, err := CreateLatestAuth(c.Client, privateKeyECDSA, c.Conf.ContractAddress.AIModels)
	if err != nil {
		return "", err
	}
	transaction, err := c.AIWorkload.ReportWorkload(
		auth,
		common.HexToAddress(worker),
		workload,
		modelId,
		sessionId,
		epochId,
		aiws,
	)
	if err != nil {
		return "", err
	}

	fmt.Println("tx....:", transaction.Hash().Hex())

	// return aim_contract, nil
	return transaction.Hash().Hex(), nil
}

// retrieve private key

func (c *ConManager) GetPrivateKeyByAddr(addr common.Address) (*ecdsa.PrivateKey, error) {

	// comAddr := common.HexToAddress(addr)

	aacount, err := c.Wallet.FindAccount(addr)
	if err != nil {
		return nil, err
	}

	err = c.Wallet.UnlockWallet(aacount, "123456")
	if err != nil {
		return nil, err
	}

	privateKey, err := c.Wallet.ExportPrivateKey(aacount, "123456")
	if err != nil {
		return nil, err
	}

	fmt.Printf("privateKey: %v\n", privateKey)
	//get addr private from keystore
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey[2:])
	if err != nil {
		return nil, err
	}

	return privateKeyECDSA, nil
}
