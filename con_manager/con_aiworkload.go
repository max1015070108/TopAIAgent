package con_manager

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/max1015070108/topaiagent/con_manager/AIWorkload"
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

func (c *ConManager) ReportWorkload(reporters []string, workload, modelId, sessionId, epochID *big.Int) error {

	privList := []*ecdsa.PrivateKey{}
	addrlist := []string{}

	for _, key := range reporters {

		privateKeyECDSA, err := c.GetPrivateKeyByAddr(common.HexToAddress(key))
		if err != nil {
			return err
		}
		privList = append(privList, privateKeyECDSA)
		addrlist = append(addrlist, key)
	}

	auth, err := CreateLatestAuth(c.Client, privList[0], c.Conf.ContractAddress.AIWorkerload)
	if err != nil {
		return err
	}

	//mock data
	// workload := big.NewInt(10)
	// modelId := big.NewInt(1)
	// sessionId := big.NewInt(1)
	// epochID := big.NewInt(2)

	signatures, err := c.SignText(
		addrlist[0],
		workload,
		modelId,
		sessionId,
		epochID,
		privList,
	)

	tx, err := c.AIWorkload.ReportWorkload(
		auth,
		common.HexToAddress(addrlist[0]),
		workload,
		modelId,
		sessionId,
		epochID,
		signatures,
	)

	if err != nil {
		return err
	}
	fmt.Printf("tx.Hash: %+v\n", tx.Hash().Hex())

	time.Sleep(2 * time.Second)
	recipt, err := c.Client.TransactionReceipt(context.Background(), tx.Hash())
	// recipt, ispending, err := conMan.Client.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		return err
	}

	fmt.Printf("recipt:%+v", recipt)
	return nil
}
