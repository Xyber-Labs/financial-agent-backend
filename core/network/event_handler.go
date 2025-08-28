package network

import (
	"context"
	"financial-agent-backend/core/onchain"
	"time"

	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rs/zerolog/log"
)

type EvmEventHandler struct {
	Client BlockNumberReader
	// Limit of blocks to scan in one iteration
	BlockLimit uint64
	// Start block to scan from
	StartBlock            uint64
	Provider              *onchain.TrustManagementProvider
	TrustManagementRouter *TrustManagementRouter.TrustManagementRouter
}

type BlockNumberReader interface {
	BlockNumber(ctx context.Context) (uint64, error)
}

func NewEvmEventHandler(
	provider *onchain.TrustManagementProvider,
	trustManagementRouter *TrustManagementRouter.TrustManagementRouter,
	client BlockNumberReader,
	blockLimit uint64,
	startBlock uint64,
) *EvmEventHandler {
	return &EvmEventHandler{
		Client:                client,
		BlockLimit:            blockLimit,
		StartBlock:            startBlock,
		Provider:              provider,
		TrustManagementRouter: trustManagementRouter,
	}
}

func (p *EvmEventHandler) Start(ctx context.Context) error {
	ticker := time.NewTicker(10 * time.Second)
	for r := range ListenBlocks(ctx, p.Client, p.StartBlock, p.BlockLimit, ticker) {
		log.Info().Uint64("start", r.Start).Uint64("end", r.End).Msg("Processing block range")
		if err := p.HandleDepositedEvents(r); err != nil {
			log.Error().Err(err).Msg("handleDepositedEvents: error handling deposited events")
		}
	}

	return nil
}

func (p *EvmEventHandler) HandleDepositedEvents(r BlockRange) error {
	depositsIter, err := p.TrustManagementRouter.FilterDeposited(&bind.FilterOpts{
		Start: r.Start,
		End:   &r.End,
	}, nil, nil)
	if err != nil {
		return err
	}
	for depositsIter.Next() {
		depositedEvent := depositsIter.Event
		log.Info().Interface("event", depositedEvent).Msg("Deposited event catched")
		tx, err := p.Provider.Deposit(
			depositedEvent.User,
			depositedEvent.Token,
			depositedEvent.Amount,
		)
		if err != nil {
			log.Error().Err(err).Msg("handleDepositedEvents: error executing deposit")
			continue
		}
		log.Info().Interface("tx", tx).Msg("Deposited event processed")
	}
	return nil
}

type BlockRange struct {
	Start uint64
	End   uint64
}

func ListenBlocks(ctx context.Context, client BlockNumberReader, start uint64, limit uint64, ticker *time.Ticker) <-chan BlockRange {
	startBlock := start
	ch := make(chan BlockRange)
	go func() {
		firstScan := true
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case <-ticker.C:
				currentBlock, err := client.BlockNumber(context.Background())
				if err != nil {
					log.Error().Err(err).Msg("ListenBlocks: error getting block number")
					continue
				}

				// If started block is higher for some reason, reset it
				if currentBlock < startBlock {
					startBlock = currentBlock
				}

				blockRange := currentBlock - startBlock
				if blockRange == 0 {
					continue
				}

				if blockRange > limit {
					blockRange = limit
				}
				endBlock := startBlock + blockRange

				if !firstScan {
					startBlock++ // skip 1 block so we don't process it again
				}
				firstScan = false

				ch <- BlockRange{
					Start: startBlock,
					End:   endBlock,
				}

				startBlock += blockRange
			}
		}
	}()
	return ch
}
