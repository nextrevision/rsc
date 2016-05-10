package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadTemplates(t *testing.T) {
	templates, err := LoadTemplates("../examples/")

	assert.Nil(t, err)
	assert.Equal(t, len(templates), 1)
	assert.Equal(t, templates[0].Name, "exported_test.json.tmpl")
	assert.Equal(t, templates[0].Path, "../examples/exported_test.json.tmpl")
}

func TestGetTemplateByName(t *testing.T) {
	templates := []Template{
		Template{
			Name: "foo",
			Path: "foo.json.tmpl",
		},
		Template{
			Name: "bar",
			Path: "bar.json.tmpl",
		},
	}

	tmpl, err := GetTemplateByName("bar", &templates)

	assert.Nil(t, err)
	assert.Equal(t, tmpl.Name, "bar")
	assert.Equal(t, tmpl.Path, "bar.json.tmpl")
}
