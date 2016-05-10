package cmd

import (
	"errors"

	"github.com/nextrevision/rsc/client"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a resource from Runscope",
}

var deleteBucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "deletes the specified buckets",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewRunscopeClient(token, debug, verbose)
		if err != nil {
			c.Log.Fatal(err)
		}

		for _, b := range args {
			err := c.DeleteBucket(b)
			if err != nil {
				c.Log.Fatal(err)
			}
		}
	},
}

var deleteTestCmd = &cobra.Command{
	Use:   "test",
	Short: "deletes the specified tests from a bucket",
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
			err := c.DeleteTest(bucket, t)
			if err != nil {
				c.Log.Fatal(err)
			}
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteBucketCmd)

	deleteTestCmd.Flags().StringVarP(&bucket, "bucket", "b", "", "bucket to list tests in")
	deleteCmd.AddCommand(deleteTestCmd)

	RootCmd.AddCommand(deleteCmd)
}
