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
func LoadTemplates(path string) ([]Template, error) {
	var templates = []Template{}

	//log.Debug("Finding templates...")
	err := filepath.Walk(path, func(filePath string, f os.FileInfo, err error) error {
		funcs := template.FuncMap{"triggerURL": triggerURL}

		if filepath.Ext(filePath) == ".tmpl" {
			//log.Debugf("Loading template: %s", filePath)

			name := strings.Replace(filePath, fmt.Sprintf("%s", path), "", 1)
			name = strings.TrimPrefix(name, "/")

			t, err := template.New("").Delims("<%", "%>").Funcs(funcs).ParseFiles(filePath)
			if err != nil {
				//log.Errorf("Error parsing template %s: %s", filePath, err.Error())
				return err
			}

			for _, temp := range t.Templates() {
				if temp.Name() != "" {
					//log.Debugf("Adding template: %s", name)
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
