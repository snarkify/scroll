//go:build snarkify_chunk
// +build snarkify_chunk

package snarkify

import (
	"context"
	"fmt"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/snarkify/snarkify-sdk-go/snarkify"
	"github.com/urfave/cli/v2"
	"scroll-tech/common/utils"
	"scroll-tech/prover"
	"scroll-tech/prover/config"
	"scroll-tech/prover/core"
)

type ProofInput struct {
	BlockHashes []common.Hash `json:"block_hashes"`
}

type ProofOutput struct {
	Proof []byte `json:"proof"`
}

func Run(ctx *cli.Context) {

	snarkify.Run[ProofInput, ProofOutput](func(input ProofInput) (ProofOutput, error) {
		// Load config file.
		cfgFile := ctx.String(utils.ConfigFileFlag.Name)
		cfg, err := config.NewConfig(cfgFile)
		if err != nil {
			log.Crit("failed to load config file", "config file", cfgFile, "error", err)
		}
		// Create prover
		r, err := prover.NewChunkProver(context.Background(), cfg)
		traces, err := r.GetSortedTracesByHashes(input.BlockHashes)
		if err != nil {
			return ProofOutput{}, fmt.Errorf("get traces from eth node failed, block hashes: %v, err: %v", input.BlockHashes, err)
		}
		cfg.Core.ProofType = message.ProofTypeChunk
		c, err := core.NewProverCore(cfg.Core)
		chunkProof, err := c.ProveChunk("default", traces)
		if err != nil {
			return ProofOutput{}, fmt.Errorf("prove chunk failed, err: %v", err)
		}
		return ProofOutput{Proof: chunkProof.Proof}, nil
	})
}
