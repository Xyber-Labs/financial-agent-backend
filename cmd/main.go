package main

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"financial-agent-backend/config"
	"financial-agent-backend/core/transactor"
	"financial-agent-backend/core/utils"

	sgx_quote "github.com/Xyber-Labs/go-tee/sgx-quote"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "financial-agent-backend",
		Short: "Financial Agent Backend",
		Long:  `Financial Agent Backend`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.LoadConfig()
			if err != nil {
				log.Error().Err(err).Msg("Error loading config")
				return
			}
			log.Info().Interface("config", cfg).Msg("Config loaded")

			if err := cfg.Validate(); err != nil {
				log.Error().Err(err).Msg("Error validating config")
				return
			}

			ethClient, err := utils.ConnectToNode(cfg.Network.HttpRpcUrl)
			if err != nil {
				log.Error().Err(err).Msg("Error connecting to node")
				return
			}
			log.Info().Str("rpc_url", cfg.Network.HttpRpcUrl).Msg("Connected to node")

			chainId, err := ethClient.ChainID(context.Background())
			if err != nil {
				log.Error().Err(err).Msg("Error getting chain id")
				return
			}

			// Prepare transactor opts.
			transactOpts, err := utils.GetTransactOptsFromPrivateKeyString(cfg.Network.SenderPrivateKey, chainId)
			if err != nil {
				log.Error().Err(err).Msg("Error creating transactor")
				return
			}

			teeService := sgx_quote.NewTeeService(false)
			transactor, err := transactor.NewTransactor(
				ethClient,
				transactOpts,
				common.HexToAddress(cfg.Network.TrustManagementRouterAddress),
				common.HexToAddress(cfg.Network.TEEWalletAddress),
				teeService,
			)
			if err != nil {
				log.Error().Err(err).Msg("Error creating transactor")
				return
			}
			if err := transactor.InitializeOnChainSession(); err != nil {
				log.Error().Err(err).Msg("Error initializing transactor")
				return
			}
		},
	}
)

func main() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Info().Str("config_file", viper.ConfigFileUsed()).Msg("Using config file")
	}
}
