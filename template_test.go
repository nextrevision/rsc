package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadTemplates(t *testing.T) {
	templates, err := loadTemplates("examples/")

	assert.Nil(t, err)
	assert.Equal(t, len(templates), 1)
	assert.Equal(t, templates[0].Name, "exported_test.json.tmpl")
	assert.Equal(t, templates[0].Path, "examples/exported_test.json.tmpl")
}

func TestGetTemplateByName(t *testing.T) {
	templates := []tmpl{
		tmpl{
			Name: "foo",
			Path: "foo.json.tmpl",
		},
		tmpl{
			Name: "bar",
			Path: "bar.json.tmpl",
		},
	}

	tmpl, err := getTemplateByName("bar", &templates)

	assert.Nil(t, err)
	assert.Equal(t, tmpl.Name, "bar")
	assert.Equal(t, tmpl.Path, "bar.json.tmpl")
}
