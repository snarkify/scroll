//go:build snarkify_batch
// +build snarkify_batch

package snarkify

import (
	"fmt"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/snarkify/snarkify-sdk-go/snarkify"
	"github.com/urfave/cli/v2"
	"scroll-tech/common/types/message"
	"scroll-tech/common/utils"
	"scroll-tech/prover/config"
	"scroll-tech/prover/core"
)

type ProofInput struct {
	ChunkInfos  []*message.ChunkInfo  `json:"chunk_infos"`
	ChunkProofs []*message.ChunkProof `json:"chunk_proofs"`
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
		c, err := core.NewProverCore(cfg.Core)
		batchProof, err := c.ProveBatch("default", input.ChunkInfos, input.ChunkProofs)
		if err != nil {
			return ProofOutput{}, fmt.Errorf("prove chunk failed, err: %v", err)
		}
		return ProofOutput{Proof: batchProof.Proof}, nil
	})
}
