package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	Networks NetworksConfig
}

type NetworksConfig []NetworkConfig

type NetworkConfig struct {
	Name string
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
	if len(c.Networks) == 0 {
		return errors.New("no networks configured")
	}

	// Check the networks
	for _, network := range c.Networks {
		if network.Name == "" {
			return errors.New("network name is empty")
		}

		if network.HttpRpcUrl == "" {
			return errors.New("network http rpc url is empty")
		}
	}

	return nil
}
