package main

import (
	"os"

	"github.com/reapchain/cosmos-sdk/server"
	svrcmd "github.com/reapchain/cosmos-sdk/server/cmd"
	sdk "github.com/reapchain/cosmos-sdk/types"

	"github.com/reapchain/ethermint/app"
	cmdcfg "github.com/reapchain/ethermint/cmd/config"
)

func main() {
	setupConfig()
	cmdcfg.RegisterDenoms()

	rootCmd, _ := NewRootCmd()

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}

func setupConfig() {
	// set the address prefixes
	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	cmdcfg.SetBip44CoinType(config)
	config.Seal()
}
