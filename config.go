package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
)

type config struct {
	Path    string
	Buckets []bucketConfig `json:"buckets"`
	Tests   []testConfig   `json:"tests"`
}

func loadConfigs(path string) ([]config, error) {
	var configs = []config{}

	log.Debug("Finding configs...")
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			log.Debugf("Loading config: %s", path)
			var config = config{Path: path}

			contents, err := os.Open(path)
			if err != nil {
				log.Errorf("Error opening config: %s", path)
				return err
			}

			defer contents.Close()

			jsonParser := json.NewDecoder(contents)
			if err = jsonParser.Decode(&config); err != nil {
				log.Errorf("Error decoding config: %s", path)
				return err
			}

			configs = append(configs, config)
		}

		return err
	})

	return configs, err
}
