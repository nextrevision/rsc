package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd is the entrypoint for the client command
var RootCmd = &cobra.Command{
	Use:   "rsc",
	Short: "rsc is a command-line client to Runscope",
	Long: `Runscope Client (rsc) provides a CLI for interacting
                with the Runscope service.`,
}

var (
	token   string
	debug   bool
	verbose bool
	dryRun  bool
	bucket  string
	format  = "json"
)

func init() {
	viper.SetEnvPrefix("rsc")
	viper.AutomaticEnv()
	RootCmd.PersistentFlags().StringVar(&token, "token", viper.GetString("token"), "runscope authentication token")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug output")
	RootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable verbose output")
}
