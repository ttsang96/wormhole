package processor

import (
	"context"
	"encoding/hex"

	"github.com/mr-tron/base58"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"

	"vebridge/node/pkg/common"
	"vebridge/node/pkg/supervisor"
	"vebridge/node/pkg/vaa"
)

var (
	// SECURITY: source_chain/target_chain are untrusted uint8 values. An attacker could cause a maximum of 255**2 label
	// pairs to be created, which is acceptable.

	lockupsObservedTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "wormhole_lockups_observed_total",
			Help: "Total number of lockups received on-chain",
		},
		[]string{"source_chain", "target_chain"})

	lockupsSignedTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "wormhole_lockups_signed_total",
			Help: "Total number of lockups that were successfully signed",
		},
		[]string{"source_chain", "target_chain"})
)

func init() {
	prometheus.MustRegister(lockupsObservedTotal)
	prometheus.MustRegister(lockupsSignedTotal)
}

// handleLockup processes a lockup received from a chain and instantiates our deterministic copy of the VAA. A lockup
// event may be received multiple times until it has been successfully completed.
func (p *Processor) handleLockup(ctx context.Context, k *common.ChainLock) {
	supervisor.Logger(ctx).Info("lockup confirmed",
		zap.Stringer("source_chain", k.SourceChain),
		zap.Stringer("target_chain", k.TargetChain),
		zap.Stringer("source_addr", k.SourceAddress),
		zap.Stringer("target_addr", k.TargetAddress),
		zap.Stringer("token_chain", k.TokenChain),
		zap.Stringer("token_addr", k.TokenAddress),
		zap.Stringer("amount", k.Amount),
		zap.Stringer("txhash", k.TxHash),
		zap.Time("timestamp", k.Timestamp),
	)

	lockupsObservedTotal.With(prometheus.Labels{
		"source_chain": k.SourceChain.String(),
		"target_chain": k.TargetChain.String()}).Add(1)

	// All nodes will create the exact same VAA and sign its digest.
	// Consensus is established on this digest.

	v := &vaa.VAA{
		Version:          vaa.SupportedVAAVersion,
		GuardianSetIndex: p.gs.Index,
		Signatures:       nil,
		Timestamp:        k.Timestamp,
		Payload: &vaa.BodyTransfer{
			Nonce:         k.Nonce,
			SourceChain:   k.SourceChain,
			TargetChain:   k.TargetChain,
			SourceAddress: k.SourceAddress,
			TargetAddress: k.TargetAddress,
			Asset: &vaa.AssetMeta{
				Chain:    k.TokenChain,
				Address:  k.TokenAddress,
				Decimals: k.TokenDecimals,
			},
			Amount: k.Amount,
		},
	}

	// Generate digest of the unsigned VAA.
	digest, err := v.SigningMsg()
	if err != nil {
		panic(err)
	}

	// Sign the digest using our node's guardian key.
	s, err := crypto.Sign(digest.Bytes(), p.gk)
	if err != nil {
		panic(err)
	}

	p.logger.Info("observed and signed confirmed lockup",
		zap.Stringer("source_chain", k.SourceChain),
		zap.Stringer("target_chain", k.TargetChain),
		zap.Stringer("txhash", k.TxHash),
		zap.String("txhash_b58", base58.Encode(k.TxHash.Bytes())),
		zap.String("digest", hex.EncodeToString(digest.Bytes())),
		zap.String("signature", hex.EncodeToString(s)))

	lockupsSignedTotal.With(prometheus.Labels{
		"source_chain": k.SourceChain.String(),
		"target_chain": k.TargetChain.String()}).Add(1)

	p.broadcastSignature(v, s, k.TxHash.Bytes())
}
