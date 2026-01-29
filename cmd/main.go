package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"financial-agent-backend/config"
	"financial-agent-backend/core/abi/bindings/AavePool"
	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"
	"financial-agent-backend/core/abi/bindings/WETH"
	"financial-agent-backend/core/network"
	"financial-agent-backend/core/onchain"
	"financial-agent-backend/core/server"
	"financial-agent-backend/core/transactor"
	"financial-agent-backend/core/utils"
	_ "financial-agent-backend/docs"

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
			ctx := cmd.Context()
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

			chainId, err := ethClient.ChainID(ctx)
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
			transactOpts.GasLimit = 1000000

			trustManagementRouter, err := TrustManagementRouter.NewTrustManagementRouter(
				common.HexToAddress(cfg.Network.TrustManagementRouterAddress),
				ethClient,
			)
			if err != nil {
				log.Error().Err(err).Msg("Error creating trust management router")
				return
			}

			teeService := sgx_quote.NewTeeService(false)
			transactor, err := transactor.NewTransactor(
				ethClient,
				chainId,
				transactOpts,
				common.HexToAddress(cfg.Network.TrustManagementRouterAddress),
				trustManagementRouter,
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

			aavePool, err := AavePool.NewAavePool(
				common.HexToAddress(cfg.Network.AavePoolAddress),
				ethClient,
			)
			if err != nil {
				log.Error().Err(err).Msg("Error creating aave pool")
				return
			}

			nativeErc20Address := common.HexToAddress(cfg.Network.NativeErc20Address)
			nativeErc20, err := WETH.NewWETH(nativeErc20Address, ethClient)
			if err != nil {
				log.Error().Err(err).Msg("Error creating native erc20")
				return
			}

			trustManagementProvider := onchain.NewTrustManagementProvider(
				ethClient,
				transactor,
				trustManagementRouter,
				common.HexToAddress(cfg.Network.AavePoolAddress),
				aavePool,
				&bind.CallOpts{},
				&nativeErc20Address,
				nativeErc20,
			)

			startBlock := -1
			if cfg.Network.StartBlock != nil {
				startBlock = int(*cfg.Network.StartBlock)
			}
			blockLimit := 100
			eventHandler := network.NewEvmEventHandler(
				trustManagementProvider,
				trustManagementRouter,
				ethClient,
				uint64(blockLimit),
				int64(startBlock),
			)
			go eventHandler.Start(ctx)

			agentServer := server.NewHttpAgentServer(
				&config.HttpServerConfig{
					Port: cfg.Http.Port,
				},
				transactor,
				trustManagementProvider,
			)
			err = agentServer.Start(ctx)
			if err != nil {
				log.Error().Err(err).Msg("Error starting agent server")
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
