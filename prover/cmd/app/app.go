package app

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"scroll-tech/common/utils"
	"scroll-tech/common/version"
	"scroll-tech/prover/snarkify"
)

var app *cli.App

func init() {
	app = cli.NewApp()
	app.Action = action
	app.Name = "prover"
	app.Usage = "The Scroll L2 Prover"
	app.Version = version.Version
	app.Flags = append(app.Flags, utils.CommonFlags...)
	app.Before = func(ctx *cli.Context) error {
		return utils.LogSetup(ctx)
	}

	// Register `prover-test` app for integration-test.
	utils.RegisterSimulation(app, utils.ChunkProverApp)
	utils.RegisterSimulation(app, utils.BatchProverApp)
}

func action(ctx *cli.Context) error {

	snarkify.Run(ctx)

	return nil
}

// Run the prover cmd func.
func Run() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
