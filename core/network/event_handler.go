package network

import (
	"context"
	"financial-agent-backend/core/onchain"
	"time"

	"financial-agent-backend/core/abi/bindings/TrustManagementRouter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// EvmEventHandler periodically scans EVM blocks and handles TrustManagement
// events. It listens for TrustManagementRouter.Deposited events in block
// ranges and, for each event, calls the on-chain provider to execute the
// corresponding deposit.
//
// Fields:
//   - Client: BlockNumberReader used to fetch the latest chain head. It is
//     queried in Start() (when StartBlock < 0) and by ListenBlocks() on each
//     tick to determine the current block number.
//   - BlockLimit: Maximum number of blocks to process per iteration. Passed
//     into ListenBlocks() as the 'limit' to bound range size.
//   - StartBlock: Starting block number to begin scanning from. If negative,
//     Start() resolves it to the latest head via Client.BlockNumber().
//   - Provider: onchain.TrustManagementProvider used to execute the deposit
//     action (Provider.Deposit) when a Deposited event is observed.
//   - TrustManagementRouter: Contract binding used to filter and iterate
//     Deposited events over a block range (FilterDeposited).
type EvmEventHandler struct {
	Client BlockNumberReader
	// Limit of blocks to scan in one iteration
	BlockLimit uint64
	// Start block to scan from
	StartBlock            int64
	Provider              *onchain.TrustManagementProvider
	TrustManagementRouter *TrustManagementRouter.TrustManagementRouter
	logger                zerolog.Logger
}

// BlockNumberReader abstracts obtaining the latest block number from an EVM
// node. Implementations must return the current chain head used to compute
// scan ranges.
type BlockNumberReader interface {
	BlockNumber(ctx context.Context) (uint64, error)
}

// NewEvmEventHandler constructs an EvmEventHandler.
//
// Arguments:
//   - provider: on-chain executor for TrustManagement operations. Used by
//     HandleDepositedEvents() to call Provider.Deposit for each emitted event.
//   - trustManagementRouter: generated contract binding for the
//     TrustManagementRouter contract. Used to filter Deposited events via
//     FilterDeposited over a block range.
//   - client: BlockNumberReader supplying current chain head. Used by Start()
//     when StartBlock < 0 to resolve the initial start block, and by
//     ListenBlocks() on each tick to compute new block ranges.
//   - blockLimit: maximum number of blocks to scan per iteration. Passed as
//     'limit' into ListenBlocks() to avoid scanning excessively large ranges.
//   - startBlock: block number to start scanning from. If < 0, Start() will set
//     the effective start block to the latest block reported by
//     client.BlockNumber().
//
// Returns a handler ready to Start() scanning and handling Deposited events.
func NewEvmEventHandler(
	provider *onchain.TrustManagementProvider,
	trustManagementRouter *TrustManagementRouter.TrustManagementRouter,
	client BlockNumberReader,
	blockLimit uint64,
	startBlock int64,
) *EvmEventHandler {
	return &EvmEventHandler{
		Client:                client,
		BlockLimit:            blockLimit,
		StartBlock:            startBlock,
		Provider:              provider,
		TrustManagementRouter: trustManagementRouter,
		logger:                log.With().Str("component", "EvmEventHandler").Logger(),
	}
}

func (p *EvmEventHandler) Start(ctx context.Context) error {
	p.logger.Info().Msg("Executing EvmEventHandler.Start()")
	ticker := time.NewTicker(10 * time.Second)
	var startBlock uint64
	if p.StartBlock < 0 {
		latestBlockNumber, err := p.Client.BlockNumber(ctx)
		if err != nil {
			return err
		}
		startBlock = latestBlockNumber
	} else {
		startBlock = uint64(p.StartBlock)
	}

	p.logger.Info().Uint64("start_block", startBlock).Msg("Start block resolved")
	for r := range ListenBlocks(ctx, p.Client, startBlock, p.BlockLimit, ticker) {
		p.logger.Info().Uint64("start", r.Start).Uint64("end", r.End).Msg("Processing block range")
		if err := p.HandleDepositedEvents(r); err != nil {
			p.logger.Error().Err(err).Msg("handleDepositedEvents: error handling deposited events")
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

		p.logger.Info().Interface("event", depositedEvent).Msg("Deposited event catched")

		if depositedEvent.Token == ethcommon.HexToAddress(onchain.TrustManagementNativeTokenLabel) {
			tx, err := p.Provider.DepositNative(
				depositedEvent.User,
				depositedEvent.Amount,
			)
			if err != nil {
				p.logger.Error().Err(err).Msg("error handling native deposit event")
				continue
			}
			p.logger.Info().Interface("tx", tx.Hash()).Msg("Native deposit event processed")
		} else {
			tx, err := p.Provider.Deposit(
				depositedEvent.User,
				depositedEvent.Token,
				depositedEvent.Amount,
			)
			if err != nil {
				p.logger.Error().Err(err).Msg("error handling deposit event")
				continue
			}
			p.logger.Info().Interface("tx", tx.Hash()).Msg("Deposit event processed")
		}
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
