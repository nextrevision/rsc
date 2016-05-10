package config

import "github.com/nextrevision/go-runscope"

// BucketConfig represents a JSON bucket file
type BucketConfig struct {
	Name     string          `json:"name"`
	Template string          `json:"template"`
	TeamID   string          `json:"team_id"`
	Config   runscope.Bucket `json:"config"`
}
