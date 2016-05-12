package cmd

import (
	"html/template"
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

		funcs := template.FuncMap{
			"triggerURL": c.TriggerURL,
		}

		c.Log.Infof("Loading configs from %s...", importPath)
		configs, err := config.LoadConfigs(importPath)
		if err != nil {
			c.Log.Fatal(err)
		}
		c.Log.Infof("Found %d configs...", len(configs))

		c.Log.Infof("Loading templates from %s...", importPath)
		templates, err := config.LoadTemplates(importPath, funcs)
		if err != nil {
			c.Log.Fatal(err)
		}
		c.Log.Infof("Found %d templates...", len(templates))

		for _, config := range configs {
			if config.Buckets != nil && strings.Contains("buckets", importInclude) {
				c.Log.Infof("Importing %d buckets...", len(config.Buckets))
				for _, bucket := range config.Buckets {

					if bucket.TeamID == "" {
						defaultTeam, err := c.GetDefaultTeam()
						if err != nil {
							c.Log.Fatal(err)
						}

						bucket.TeamID = defaultTeam.ID
					}

					_, err := c.CreateOrUpdateBucket(&bucket)
					if err != nil {
						c.Log.Fatal(err)
					}
				}
			}

			if config.Tests != nil && strings.Contains("tests", importInclude) {
				c.Log.Infof("Importing %d tests...", len(config.Tests))
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
