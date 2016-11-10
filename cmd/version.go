package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "0.2.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version number of rsc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("rsc %s\n", Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
