package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	Network NetworkConfig
}

type NetworkConfig struct {
	// Name of the network (e.g. ethereum)
	Name string
	// Http RPC endpoint to use for interacting with the network
	HttpRpcUrl string
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

	return nil
}
