package config

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// Template represents a template file
type Template struct {
	Name     string
	Path     string
	template *template.Template
}

// LoadTemplates walks a path in search for all template files
func LoadTemplates(path string, funcs template.FuncMap) ([]Template, error) {
	var templates = []Template{}

	err := filepath.Walk(path, func(filePath string, f os.FileInfo, err error) error {
		if filepath.Ext(filePath) == ".tmpl" {

			name := strings.Replace(filePath, fmt.Sprintf("%s", path), "", 1)
			name = strings.TrimPrefix(name, "/")

			t, err := template.New("").Delims("<%", "%>").Funcs(funcs).ParseFiles(filePath)
			if err != nil {
				return err
			}

			for _, temp := range t.Templates() {
				if temp.Name() != "" {
					templates = append(templates, Template{
						Name:     name,
						Path:     filePath,
						template: temp,
					})
				}
			}
		}

		return nil
	})

	return templates, err
}

// GetTemplateByName searches a list of Templates for the specified
// name of the template
func GetTemplateByName(name string, templates *[]Template) (Template, error) {
	for _, t := range *templates {
		if t.Name == name {
			return t, nil
		}
	}

	return Template{}, fmt.Errorf("Could not find template with name: %s", name)
}
