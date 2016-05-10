package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version number of rsc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rsc - Runscope Command Line Client -- HEAD")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
