package wallet

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// DefaultKeyDir 默认密钥目录
	DefaultKeyDir = "./keystore"
	Passphrase    = "123456"
)

// WalletManager 钱包管理结构体
type WalletManager struct {
	ks     *keystore.KeyStore
	keyDir string
}

// NewWalletManager 创建钱包管理器
func NewWalletManager(keyDir string) *WalletManager {
	ks := keystore.NewKeyStore(
		keyDir,
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)
	return &WalletManager{
		ks:     ks,
		keyDir: keyDir,
	}
}

// CreateWallet 创建新钱包
func (wm *WalletManager) CreateWallet(password string) (accounts.Account, error) {
	account, err := wm.ks.NewAccount(password)
	if err != nil {
		return accounts.Account{}, fmt.Errorf("create wallet error: %v", err)
	}
	return account, nil
}

// ImportWalletFromPrivateKey 从私钥导入钱包
func (wm *WalletManager) ImportWalletFromPrivateKey(privateKeyHex string, password string) (accounts.Account, error) {
	// 去除可能的 "0x" 前缀
	if len(privateKeyHex) > 2 && privateKeyHex[0:2] == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return accounts.Account{}, fmt.Errorf("invalid private key: %v", err)
	}

	account, err := wm.ks.ImportECDSA(privateKey, password)
	if err != nil {
		return accounts.Account{}, fmt.Errorf("import wallet error: %v", err)
	}

	return account, nil
}

// ExportPrivateKey 导出私钥
func (wm *WalletManager) ExportPrivateKey(account accounts.Account, password string) (string, error) {
	keyJson, err := os.ReadFile(account.URL.Path)
	if err != nil {
		return "", fmt.Errorf("read keystore file error: %v", err)
	}

	key, err := keystore.DecryptKey(keyJson, password)
	if err != nil {
		return "", fmt.Errorf("decrypt key error: %v", err)
	}

	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)
	return hexutil.Encode(privateKeyBytes), nil
}

// UnlockWallet 解锁钱包
func (wm *WalletManager) UnlockWallet(account accounts.Account, password string) error {
	return wm.ks.Unlock(account, password)
}

// GetBalance 获取账户余额
func (wm *WalletManager) GetBalance(client *ethclient.Client, account accounts.Account) (*big.Int, error) {
	balance, err := client.BalanceAt(context.Background(), account.Address, nil)
	if err != nil {
		return nil, fmt.Errorf("get balance error: %v", err)
	}
	return balance, nil
}

// GetAllAccounts 获取所有账户
func (wm *WalletManager) GetAllAccounts() []accounts.Account {
	return wm.ks.Accounts()
}

// ListWalletAddresses 列出keystore中所有钱包地址
func (wm *WalletManager) ListWalletAddresses() []string {
	accounts := wm.ks.Accounts()
	addresses := make([]string, len(accounts))

	for i, account := range accounts {
		addresses[i] = account.Address.Hex()
	}

	return addresses
}

// 方法3：通过地址查找获取 Account
func (wm *WalletManager) FindAccount(address common.Address) (accounts.Account, error) {
	// 获取所有账户
	allAccounts := wm.ks.Accounts()

	// 查找匹配地址的账户
	for _, account := range allAccounts {

		fmt.Printf("%v<=>%v\n", address.Hex(), account.Address.Hex())
		if strings.EqualFold(account.Address.Hex(), address.Hex()) {
			return account, nil
		}
	}

	return accounts.Account{}, fmt.Errorf("account not found for address: %s", address.Hex())
}
