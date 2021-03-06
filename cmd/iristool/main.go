package main

import (
	"os"

	irisInit "github.com/irisnet/irishub/server/init"
	"github.com/spf13/cobra"
	debugcmd "github.com/irisnet/irishub/tools/debug"
	"github.com/irisnet/irishub/tools/prometheus"
	"github.com/irisnet/irishub/app"
	"github.com/tendermint/tendermint/libs/cli"
)

func init() {

	irisInit.InitBech32Prefix()
	cdc := app.MakeCodec()

	rootCmd.AddCommand(debugcmd.RootCmd)
	rootCmd.AddCommand(prometheus.MonitorCommand(cdc))
}

var rootCmd = &cobra.Command{
	Use:          "iristool",
	Short:        "Iris tool",
	SilenceUsage: true,
}

func main() {
	executor := cli.PrepareMainCmd(rootCmd, "IRIS", app.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
