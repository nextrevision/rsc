package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/nextrevision/go-runscope"
)

type testConfig struct {
	Name      string      `json:"name"`
	Template  string      `json:"template"`
	Bucket    string      `json:"bucket"`
	BucketKey string      `json:"bucket_key"`
	Vars      interface{} `json:"vars"`
	Data      interface{} `json:"data"`
	Bytes     []byte      `json:"-"`
}

func processTests(configs *[]config, templates *[]tmpl) error {
	log.Info("Processing tests...")
	for _, config := range *configs {
		if config.Tests != nil {
			for _, test := range config.Tests {
				log.Debugf("Found test: %s", test.Name)
				if test.Bucket == "" {
					return fmt.Errorf("Must specify a bucket for test: %s", test.Name)
				}

				data, err := test.getTestData(templates)
				if err != nil {
					log.Errorf("Could not get data for test: %s", test.Name)
					return err
				}

				test.Bytes = data

				if err = test.createOrUpdateTest(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (t *testConfig) createOrUpdateTest() error {
	if t.BucketKey == "" {
		bucket, err := getBucketByName(t.Bucket)
		if err != nil {
			log.Error("Could determine bucket for test: %s", t.Name)
			return err
		}

		t.BucketKey = bucket.Key
	}

	tests, err := client.ListTests(t.BucketKey)
	if err != nil {
		log.Error("Could list tests for bucket: %s", t.BucketKey)
		return err
	}

	for _, test := range *tests {
		if test.Name == t.Name {
			return t.update(test.ID)
		}
	}

	return t.create()
}

func (t *testConfig) create() error {
	log.Debugf("Creating test: %s", t.Name)

	_, err := client.ImportTest(t.BucketKey, t.Bytes)
	if err == nil {
		log.Infof("Created test: %s", t.Name)
	} else {
		log.Errorf("Error creating test: %s", t.Name)
	}

	return err
}

func (t *testConfig) update(testID string) error {
	log.Debugf("Updating test: %s", t.Name)

	_, err := client.ReimportTest(t.BucketKey, testID, t.Bytes)
	if err == nil {
		log.Infof("Updated test: %s", t.Name)
	} else {
		log.Errorf("Error updating test: %s", t.Name)
	}

	return err
}

func (t *testConfig) getTestData(templates *[]tmpl) ([]byte, error) {
	if t.Data != nil {
		return json.Marshal(t.Data)
	} else if t.Template != "" {
		template, err := getTemplateByName(t.Template, templates)
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

func getTestByName(bucketKey string, testName string) (*runscope.Test, error) {
	var test = &runscope.Test{}

	tests, err := client.ListTests(bucketKey)
	if err != nil {
		return test, err
	}

	for _, t := range *tests {
		if t.Name == testName {
			return &t, nil
		}
	}

	return test, fmt.Errorf("No such test: %s", testName)
}
