package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/deckarep/golang-set"
)

// TestConfig represents a JSON test file
type TestConfig struct {
	Name      string      `json:"name"`
	Template  string      `json:"template"`
	Bucket    string      `json:"bucket"`
	BucketKey string      `json:"bucket_key"`
	Vars      interface{} `json:"vars"`
	Data      interface{} `json:"data"`
	Bytes     []byte      `json:"-"`
	Depends   []string    `json:"depends"`
}

// GetTestData returns either a compiled template or inline data
func (t *TestConfig) GetTestData(templates *[]Template) ([]byte, error) {
	if t.Data != nil {
		return json.Marshal(t.Data)
	} else if t.Template != "" {
		template, err := GetTemplateByName(t.Template, templates)
		if err != nil {
			return []byte{}, err
		}

		jsonOut := new(bytes.Buffer)
		if err = template.template.Execute(jsonOut, t); err != nil {
			return []byte{}, err
		}

		return jsonOut.Bytes(), nil
	}

	return []byte{}, fmt.Errorf("Could not find any data for %s", t.Name)
}

// Largely taken from github.com/dnaeon/go-dependency-graph-algorithm
func OrderTests(tests []TestConfig) ([]TestConfig, error) {
	testNames := make(map[string]TestConfig)
	testDependencies := make(map[string]mapset.Set)

	// Populate the maps
	for _, test := range tests {
		testNames[test.Name] = test

		dependencySet := mapset.NewSet()
		for _, dep := range test.Depends {
			dependencySet.Add(dep)
		}
		testDependencies[test.Name] = dependencySet
	}

	// Iteratively find and remove nodes from the graph which have no dependencies.
	// If at some point there are still nodes in the graph and we cannot find
	// nodes without dependencies, that means we have a circular dependency
	var resolved []TestConfig
	for len(testDependencies) != 0 {
		// Get all nodes from the graph which have no dependencies
		readySet := mapset.NewSet()
		for name, deps := range testDependencies {
			if deps.Cardinality() == 0 {
				readySet.Add(name)
			}
		}

		// If there aren't any ready nodes, then we have a cicular dependency
		if readySet.Cardinality() == 0 {
			var g []TestConfig
			for name := range testDependencies {
				g = append(g, testNames[name])
			}

			return g, errors.New("Circular dependency found")
		}

		// Remove the ready nodes and add them to the resolved graph
		for name := range readySet.Iter() {
			delete(testDependencies, name.(string))
			resolved = append(resolved, testNames[name.(string)])
		}

		// Also make sure to remove the ready nodes from the
		// remaining node dependencies as well
		for name, deps := range testDependencies {
			diff := deps.Difference(readySet)
			testDependencies[name] = diff
		}
	}

	return resolved, nil
}
