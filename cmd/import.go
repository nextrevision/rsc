package cmd

import (
	"strings"

	"github.com/nextrevision/rsc/client"
	"github.com/nextrevision/rsc/config"

	"github.com/spf13/cobra"
)

var importPath string
var importInclude string

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import tests from a path of configs and templates",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewRunscopeClient(token, debug, verbose)
		if err != nil {
			c.Log.Fatal(err)
		}

		configs, err := config.LoadConfigs(importPath)
		if err != nil {
			c.Log.Fatal(err)
		}

		templates, err := config.LoadTemplates(importPath)
		if err != nil {
			c.Log.Fatal(err)
		}

		for _, config := range configs {
			if config.Buckets != nil && strings.Contains("buckets", importInclude) {
				for _, bucket := range config.Buckets {

					if bucket.TeamID == "" {
						defaultTeam, err := c.GetDefaultTeam()
						if err != nil {
							c.Log.Fatal(err)
						}

						bucket.TeamID = defaultTeam.ID
					}

					bucket, err := c.CreateOrUpdateBucket(&bucket)
					if err != nil {
						c.Log.Fatal(err)
					}
				}
			}

			if config.Tests != nil && strings.Contains("tests", importInclude) {
				for _, test := range config.Tests {
					if test.Bucket == "" {
						c.Log.Fatalf("Must specify a bucket for test: %s", test.Name)
					}

					data, err := test.GetTestData(&templates)
					if err != nil {
						c.Log.Fatal(err)
					}

					test.Bytes = data

					if err = c.CreateOrUpdateTest(&test); err != nil {
						c.Log.Fatal(err)
					}
				}
			}
		}
	},
}

func init() {
	importCmd.Flags().StringVarP(&importPath, "path", "p", ".", "base path to search for configs and templates")
	importCmd.Flags().StringVarP(&importInclude, "include", "i", "", "comma-separated list of types to import (bucket, tests)")

	RootCmd.AddCommand(importCmd)
}
