package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config represents a listing of all config files
type Config struct {
	Path    string
	Buckets []BucketConfig `json:"buckets"`
	Tests   []TestConfig   `json:"tests"`
}

// LoadConfigs walks a path and returns a list of configs
func LoadConfigs(path string) ([]Config, error) {
	var configs = []Config{}

	//log.Debug("Finding configs...")
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			//log.Debugf("Loading config: %s", path)
			var config = Config{Path: path}

			contents, err := os.Open(path)
			if err != nil {
				//log.Errorf("Error opening config: %s", path)
				return err
			}

			defer contents.Close()

			jsonParser := json.NewDecoder(contents)
			if err = jsonParser.Decode(&config); err != nil {
				//log.Errorf("Error decoding config: %s", path)
				return err
			}

			configs = append(configs, config)
		}

		return err
	})

	return configs, err
}
