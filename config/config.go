package config

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/spf13/viper"
)

type Config struct {
	Network NetworkConfig
	Http    HttpServerConfig
}

type NetworkConfig struct {
	// Name of the network (e.g. ethereum)
	Name string
	// Http RPC endpoint to use for interacting with the network
	HttpRpcUrl string
	// Private key to be used for sending transactions
	SenderPrivateKey string
	// Trust Management Router address
	TrustManagementRouterAddress string
	// Aave Pool address
	AavePoolAddress string
	// Optional block number to start processing events from
	StartBlock *uint64
	// Address of ERC20 Wrapped native token on the network
	NativeErc20Address string
}

type HttpServerConfig struct {
	Port int
}

func LoadConfig() (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) Validate() error {
	if c.Network.Name == "" {
		return errors.New("network name is empty")
	}

	if c.Network.HttpRpcUrl == "" {
		return errors.New("network http rpc url is empty")
	}

	if c.Network.SenderPrivateKey == "" {
		return errors.New("network sender private key is empty")
	}

	if !regexp.MustCompile(`^0x[0-9a-fA-F]{64}$`).MatchString(c.Network.SenderPrivateKey) {
		return fmt.Errorf("invalid sender private key format: %s", c.Network.SenderPrivateKey)
	}

	if c.Network.TrustManagementRouterAddress == "" {
		return errors.New("network trust management router address is empty")
	}

	if c.Network.AavePoolAddress == "" {
		return errors.New("network aave pool address is empty")
	}

	return nil
}
