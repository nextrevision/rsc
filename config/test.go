package config

import (
	"bytes"
	"encoding/json"
	"fmt"
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
