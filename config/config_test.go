package config

import (
	"os"
	"testing"
)

func TestLoadJSONConfig(t *testing.T) {
	jsonConfig := `{
  "database": {
   "host": "localhost",
   "port": 5432,
   "user": "test_user",
   "password": "test_password",
   "dbname": "testdb"
  },
  "ethereum": {
   "rpc_url": "http://localhost:8545",
   "chain_id": 1,
   "contract_address": "0x123"
  },
  "server": {
   "host": "localhost",
   "port": 8080,
   "env_mode": "test"
  },
  "contractaddress": {
   "ai_models": "0x456",
   "ai_workerload": "0x789",
   "node_register": "0xabc"
  }
 }`

	tmpfile, err := os.CreateTemp("", "config*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(jsonConfig)); err != nil {
		t.Fatal(err)
	}

	loader := NewConfigLoader(tmpfile.Name(), "")
	config, err := loader.loadJSONConfig()
	if err != nil {
		t.Fatal(err)
	}

	// Test database config
	if config.Database.Host != "localhost" {
		t.Errorf("Expected database host localhost, got %s", config.Database.Host)
	}
	if config.Database.Port != 5432 {
		t.Errorf("Expected database port 5432, got %d", config.Database.Port)
	}

	// Test ethereum config
	if config.Ethereum.RPCURL != "http://localhost:8545" {
		t.Errorf("Expected ethereum RPC URL http://localhost:8545, got %s", config.Ethereum.RPCURL)
	}
	if config.Ethereum.ChainID != 1 {
		t.Errorf("Expected chain ID 1, got %d", config.Ethereum.ChainID)
	}

	// Test server config
	if config.Server.Port != 8080 {
		t.Errorf("Expected server port 8080, got %d", config.Server.Port)
	}

	// Test contract addresses
	if config.ContractAddress.AIModels != "0x456" {
		t.Errorf("Expected AI models address 0x456, got %s", config.ContractAddress.AIModels)
	}
}

func TestOverrideWithEnv(t *testing.T) {
	config := &Config{}
	loader := NewConfigLoader("", "")

	os.Setenv("DB_HOST", "testhost")
	os.Setenv("DB_PORT", "5555")
	os.Setenv("ETH_RPC_URL", "http://testeth:8545")
	os.Setenv("SERVER_PORT", "9090")
	os.Setenv("CONTRACT_AI_MODELS", "0xtest")

	if err := loader.overrideWithEnv(config); err != nil {
		t.Fatal(err)
	}

	if config.Database.Host != "testhost" {
		t.Errorf("Expected database host testhost, got %s", config.Database.Host)
	}
	if config.Database.Port != 5555 {
		t.Errorf("Expected database port 5555, got %d", config.Database.Port)
	}
	if config.Ethereum.RPCURL != "http://testeth:8545" {
		t.Errorf("Expected ethereum RPC URL http://testeth:8545, got %s", config.Ethereum.RPCURL)
	}
	if config.Server.Port != 9090 {
		t.Errorf("Expected server port 9090, got %d", config.Server.Port)
	}
	if config.ContractAddress.AIModels != "0xtest" {
		t.Errorf("Expected AI models address 0xtest, got %s", config.ContractAddress.AIModels)
	}
}

func TestValidateConfig(t *testing.T) {
	loader := NewConfigLoader("", "")

	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "Valid config",
			config: &Config{
				Database: struct {
					Host     string `json:"host" env:"DB_HOST"`
					Port     int    `json:"port" env:"DB_PORT"`
					User     string `json:"user" env:"DB_USER"`
					Password string `json:"password" env:"DB_PASSWORD"`
					DBName   string `json:"dbname" env:"DB_NAME"`
				}{
					Port: 5432,
				},
				Server: struct {
					Host    string `json:"host" env:"SERVER_HOST"`
					Port    int    `json:"port" env:"SERVER_PORT"`
					EnvMode string `json:"env_mode" env:"ENV_MODE"`
				}{
					Port: 8080,
				},
				Ethereum: struct {
					RPCURL          string `json:"rpc_url" env:"ETH_RPC_URL"`
					PrivateKey      string `json:"private_key" env:"ETH_PRIVATE_KEY"`
					ChainID         int64  `json:"chain_id" env:"ETH_CHAIN_ID"`
					ContractAddress string `json:"contract_address" env:"ETH_CONTRACT_ADDRESS"`
				}{
					ChainID: 1,
				},
				ContractAddress: struct {
					AIModels     string `json:"ai_models" env:"CONTRACT_AI_MODELS"`
					AIWorkerload string `json:"ai_workerload" env:"CONTRACT_AI_WORKERLOAD"`
					NodeRegister string `json:"node_register" env:"CONTRACT_NODE_REGISTER"`
				}{
					AIModels:     "0x123",
					AIWorkerload: "0x456",
					NodeRegister: "0x789",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid database port",
			config: &Config{
				Database: struct {
					Host     string `json:"host" env:"DB_HOST"`
					Port     int    `json:"port" env:"DB_PORT"`
					User     string `json:"user" env:"DB_USER"`
					Password string `json:"password" env:"DB_PASSWORD"`
					DBName   string `json:"dbname" env:"DB_NAME"`
				}{
					Port: -1,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := loader.validateConfig(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
