package cmd

import (
	"errors"

	"github.com/nextrevision/rsc/client"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show a listing of resources",
}

var listBucketsCmd = &cobra.Command{
	Use:   "buckets",
	Short: "show a listing of buckets",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewRunscopeClient(token, debug, verbose)
		if err != nil {
			c.Log.Fatal(err)
		}

		err = c.ListBuckets(format)
		if err != nil {
			c.Log.Fatal(err)
		}
	},
}

var listTestsCmd = &cobra.Command{
	Use:   "tests",
	Short: "show a listing of tests in a bucket",
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

		err = c.ListTests(bucket, format)
		if err != nil {
			c.Log.Fatal(err)
		}
	},
}

func init() {
	listCmd.AddCommand(listBucketsCmd)

	listTestsCmd.Flags().StringVarP(&bucket, "bucket", "b", "", "bucket to list tests in")
	listCmd.PersistentFlags().StringVarP(&format, "format", "f", "", "output format (cli, json)")
	listCmd.AddCommand(listTestsCmd)

	RootCmd.AddCommand(listCmd)
}
