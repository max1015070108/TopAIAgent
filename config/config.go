package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config 配置结构体
type Config struct {
	Database struct {
		Host     string `json:"host" env:"DB_HOST"`
		Port     int    `json:"port" env:"DB_PORT"`
		User     string `json:"user" env:"DB_USER"`
		Password string `json:"password" env:"DB_PASSWORD"`
		DBName   string `json:"dbname" env:"DB_NAME"`
	} `json:"database"`

	Ethereum struct {
		RPCURL          string `json:"rpc_url" env:"ETH_RPC_URL"`
		PrivateKey      string `json:"private_key" env:"ETH_PRIVATE_KEY"`
		ChainID         int64  `json:"chain_id" env:"ETH_CHAIN_ID"`
		ContractAddress string `json:"contract_address" env:"ETH_CONTRACT_ADDRESS"`
	} `json:"ethereum"`

	Server struct {
		Host    string `json:"host" env:"SERVER_HOST"`
		Port    int    `json:"port" env:"SERVER_PORT"`
		EnvMode string `json:"env_mode" env:"ENV_MODE"`
	} `json:"server"`

	ContractAddress struct {
		AIModels     string `json:"ai_models" env:"CONTRACT_AI_MODELS"`
		AIWorkerload string `json:"ai_workerload" env:"CONTRACT_AI_WORKERLOAD"`
		NodeRegister string `json:"node_register" env:"CONTRACT_NODE_REGISTER"`
	} `json:"contractaddress"`
}

// ConfigLoader 配置加载器
type ConfigLoader struct {
	jsonPath string
	envPath  string
}

// NewConfigLoader 创建新的配置加载器
func NewConfigLoader(jsonPath, envPath string) *ConfigLoader {
	return &ConfigLoader{
		jsonPath: jsonPath,
		envPath:  envPath,
	}
}

// Load 加载并合并配置
func (cl *ConfigLoader) Load() (*Config, error) {
	// 1. 首先加载JSON配置作为基础配置
	config, err := cl.loadJSONConfig()
	if err != nil {
		return nil, fmt.Errorf("loading json config: %w", err)
	}

	// 2. 加载.env文件（如果存在）
	if cl.envPath != "" {
		if err := godotenv.Load(cl.envPath); err != nil {
			// 如果.env文件不存在，仅记录警告但不返回错误
			fmt.Printf("Warning: .env file not found at %s\n", cl.envPath)
		}
	}

	// 3. 用环境变量覆盖配置
	if err := cl.overrideWithEnv(config); err != nil {
		return nil, fmt.Errorf("overriding with env vars: %w", err)
	}

	// 4. 验证配置
	if err := cl.validateConfig(config); err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}

	return config, nil
}

// loadJSONConfig 加载JSON配置文件
func (cl *ConfigLoader) loadJSONConfig() (*Config, error) {
	var config Config

	data, err := os.ReadFile(cl.jsonPath)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("parsing config file: %w", err)
	}

	return &config, nil
}

// overrideWithEnv 使用环境变量覆盖配置
func (cl *ConfigLoader) overrideWithEnv(config *Config) error {
	// Database overrides
	if host := os.Getenv("DB_HOST"); host != "" {
		config.Database.Host = host
	}
	if port := os.Getenv("DB_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.Database.Port = p
		}
	}
	if user := os.Getenv("DB_USER"); user != "" {
		config.Database.User = user
	}
	if pass := os.Getenv("DB_PASSWORD"); pass != "" {
		config.Database.Password = pass
	}
	if name := os.Getenv("DB_NAME"); name != "" {
		config.Database.DBName = name
	}

	// Ethereum overrides
	if url := os.Getenv("ETH_RPC_URL"); url != "" {
		config.Ethereum.RPCURL = url
	}
	if key := os.Getenv("ETH_PRIVATE_KEY"); key != "" {
		config.Ethereum.PrivateKey = key
	}
	if chainID := os.Getenv("ETH_CHAIN_ID"); chainID != "" {
		if id, err := strconv.ParseInt(chainID, 10, 64); err == nil {
			config.Ethereum.ChainID = id
		}
	}
	if addr := os.Getenv("ETH_CONTRACT_ADDRESS"); addr != "" {
		config.Ethereum.ContractAddress = addr
	}

	// Server overrides
	if host := os.Getenv("SERVER_HOST"); host != "" {
		config.Server.Host = host
	}
	if port := os.Getenv("SERVER_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.Server.Port = p
		}
	}
	if mode := os.Getenv("ENV_MODE"); mode != "" {
		config.Server.EnvMode = mode
	}

	// ContractAddress overrides
	if addr := os.Getenv("CONTRACT_AI_MODELS"); addr != "" {
		config.ContractAddress.AIModels = addr
	}
	if addr := os.Getenv("CONTRACT_AI_WORKERLOAD"); addr != "" {
		config.ContractAddress.AIWorkerload = addr
	}
	if addr := os.Getenv("CONTRACT_NODE_REGISTER"); addr != "" {
		config.ContractAddress.NodeRegister = addr
	}

	return nil
}

// validateConfig 验证配置
func (cl *ConfigLoader) validateConfig(config *Config) error {
	// 验证数据库配置
	if config.Database.Port < 0 || config.Database.Port > 65535 {
		return fmt.Errorf("invalid database port: %d", config.Database.Port)
	}

	// 验证服务器配置
	if config.Server.Port < 0 || config.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", config.Server.Port)
	}

	// 验证以太坊配置
	if config.Ethereum.ChainID <= 0 {
		return fmt.Errorf("invalid chain ID: %d", config.Ethereum.ChainID)
	}

	if config.ContractAddress.AIModels == "" {
		return fmt.Errorf("ai_models contract address is empty")
	}
	if config.ContractAddress.AIWorkerload == "" {
		return fmt.Errorf("ai_workerload contract address is empty")
	}
	if config.ContractAddress.NodeRegister == "" {
		return fmt.Errorf("node_register contract address is empty")
	}

	return nil
}

// 使用示例
func Example() {
	// config.json 示例内容
	jsonConfig := `{
        "database": {
            "host": "localhost",
            "port": 5432,
            "user": "default_user",
            "password": "default_password",
            "dbname": "mydb"
        },
        "ethereum": {
            "rpc_url": "http://localhost:8545",
            "chain_id": 1,
            "contract_address": "0x..."
        },
        "server": {
            "host": "localhost",
            "port": 8080,
            "env_mode": "development"
        },
        "contractaddress": {
                    "ai_models": "0x...",
                    "ai_workerload": "0x...",
                    "node_register": "0x..."
                }
    }`

	// 创建临时配置文件
	tmpfile, err := os.CreateTemp("", "config*.json")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(jsonConfig)); err != nil {
		panic(err)
	}

	// 创建配置加载器
	loader := NewConfigLoader(tmpfile.Name(), ".env")

	// 加载配置
	config, err := loader.Load()
	if err != nil {
		panic(err)
	}

	// 使用配置
	fmt.Printf("Database Host: %s\n", config.Database.Host)
	fmt.Printf("Ethereum RPC URL: %s\n", config.Ethereum.RPCURL)
	fmt.Printf("Server Port: %d\n", config.Server.Port)
}
