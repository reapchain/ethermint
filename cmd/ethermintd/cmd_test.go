package main_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/reapchain/cosmos-sdk/client/flags"
	svrcmd "github.com/reapchain/cosmos-sdk/server/cmd"
	"github.com/reapchain/cosmos-sdk/x/genutil/client/cli"

	"github.com/tharsis/ethermint/app"
	ethermintd "github.com/tharsis/ethermint/cmd/ethermintd"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := ethermintd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",          // Test the init cmd
		"etherminttest", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, "ethermint_9000-1"),
	})

	err := svrcmd.Execute(rootCmd, app.DefaultNodeHome)
	require.NoError(t, err)
}
