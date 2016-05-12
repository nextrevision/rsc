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

var token string
var debug bool
var verbose bool
var dryRun bool
var bucket string

func init() {
	viper.SetEnvPrefix("rsc")
	viper.AutomaticEnv()
	RootCmd.PersistentFlags().StringVar(&token, "token", viper.GetString("token"), "runscope authentication token")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug output")
	RootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable verbose output")
	//RootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	//RootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	//RootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	//viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	//viper.BindPFlag("projectbase", RootCmd.PersistentFlags().Lookup("projectbase"))
	//viper.BindPFlag("useViper", RootCmd.PersistentFlags().Lookup("viper"))
	//viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	//viper.SetDefault("license", "apache")
}
