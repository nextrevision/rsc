package cmd

import (
	"errors"

	"github.com/nextrevision/rsc/client"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show details of a resource",
}

var showBucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "show details of a bucket",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewRunscopeClient(token, debug, verbose)
		if err != nil {
			c.Log.Fatal(err)
		}

		for _, b := range args {
			err := c.ShowBucket(b, format)
			if err != nil {
				c.Log.Fatal(err)
			}
		}
	},
}

var showTestCmd = &cobra.Command{
	Use:   "test",
	Short: "show details of a test",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if bucket == "" {
			return errors.New("must specify a bucket")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewRunscopeClient(token, debug, verbose)
		if err != nil {
			c.Log.Fatal(err)
		}

		for _, t := range args {
			err := c.ShowTest(bucket, t, format)
			if err != nil {
				c.Log.Fatal(err)
			}
		}
	},
}

func init() {
	showTestCmd.Flags().StringVarP(&bucket, "bucket", "b", "", "bucket that contains the test")
	showCmd.PersistentFlags().StringVarP(&format, "format", "f", "", "output format (only json supported)")

	showCmd.AddCommand(showBucketCmd)
	showCmd.AddCommand(showTestCmd)

	RootCmd.AddCommand(showCmd)
}
