package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigs(t *testing.T) {
	configs, err := LoadConfigs("../examples/")

	assert.Nil(t, err)
	assert.Equal(t, len(configs), 1)
	assert.Equal(t, configs[0].Path, "../examples/config.json")

	buckets := configs[0].Buckets
	assert.Equal(t, len(buckets), 1)
	assert.Equal(t, buckets[0].Name, "MyBucket")

	tests := configs[0].Tests
	assert.Equal(t, len(tests), 2)

	firstTest := tests[0]
	assert.Equal(t, firstTest.Name, "MyFirstTest")
	assert.Equal(t, firstTest.Bucket, "MyBucket")
	assert.Empty(t, firstTest.Vars)
	assert.Empty(t, firstTest.Template)
	assert.NotEmpty(t, firstTest.Data)

	secondTest := tests[1]
	assert.Equal(t, secondTest.Name, "MySecondTest")
	assert.Equal(t, secondTest.Bucket, "MyBucket")
	assert.Equal(t, secondTest.Template, "exported_test.json.tmpl")
	assert.Empty(t, secondTest.Data)
	assert.NotEmpty(t, secondTest.Vars)
}
