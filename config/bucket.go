package config

import (
	"errors"

	"github.com/deckarep/golang-set"
	"github.com/nextrevision/go-runscope"
)

// BucketConfig represents a JSON bucket file
type BucketConfig struct {
	Name     string          `json:"name"`
	Template string          `json:"template"`
	TeamID   string          `json:"team_id"`
	Config   runscope.Bucket `json:"config"`
	Depends  []string        `json:"depends"`
}

// Largely taken from github.com/dnaeon/go-dependency-graph-algorithm
func OrderBuckets(buckets []BucketConfig) ([]BucketConfig, error) {
	bucketNames := make(map[string]BucketConfig)
	bucketDependencies := make(map[string]mapset.Set)

	// Populate the maps
	for _, bucket := range buckets {
		bucketNames[bucket.Name] = bucket

		dependencySet := mapset.NewSet()
		for _, dep := range bucket.Depends {
			dependencySet.Add(dep)
		}
		bucketDependencies[bucket.Name] = dependencySet
	}

	// Iteratively find and remove nodes from the graph which have no dependencies.
	// If at some point there are still nodes in the graph and we cannot find
	// nodes without dependencies, that means we have a circular dependency
	var resolved []BucketConfig
	for len(bucketDependencies) != 0 {
		// Get all nodes from the graph which have no dependencies
		readySet := mapset.NewSet()
		for name, deps := range bucketDependencies {
			if deps.Cardinality() == 0 {
				readySet.Add(name)
			}
		}

		// If there aren't any ready nodes, then we have a cicular dependency
		if readySet.Cardinality() == 0 {
			var g []BucketConfig
			for name := range bucketDependencies {
				g = append(g, bucketNames[name])
			}

			return g, errors.New("Circular dependency found")
		}

		// Remove the ready nodes and add them to the resolved graph
		for name := range readySet.Iter() {
			delete(bucketDependencies, name.(string))
			resolved = append(resolved, bucketNames[name.(string)])
		}

		// Also make sure to remove the ready nodes from the
		// remaining node dependencies as well
		for name, deps := range bucketDependencies {
			diff := deps.Difference(readySet)
			bucketDependencies[name] = diff
		}
	}

	return resolved, nil
}
