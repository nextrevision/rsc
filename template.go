package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type tmpl struct {
	Name     string
	Path     string
	template *template.Template
}

func loadTemplates(path string) ([]tmpl, error) {
	var templates = []tmpl{}

	log.Debug("Finding templates...")
	err := filepath.Walk(path, func(filePath string, f os.FileInfo, err error) error {
		funcs := template.FuncMap{"triggerURL": triggerURL}

		if filepath.Ext(filePath) == ".tmpl" {
			log.Debugf("Loading template: %s", filePath)

			name := strings.Replace(filePath, fmt.Sprintf("%s", path), "", 1)
			name = strings.TrimPrefix(name, "/")

			t, err := template.New("").Delims("<%", "%>").Funcs(funcs).ParseFiles(filePath)
			if err != nil {
				log.Errorf("Error parsing template %s: %s", filePath, err.Error())
				return err
			}

			for _, temp := range t.Templates() {
				if temp.Name() != "" {
					log.Debugf("Adding template: %s", name)
					templates = append(templates, tmpl{
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

func getTemplateByName(name string, templates *[]tmpl) (*tmpl, error) {
	for _, t := range *templates {
		if t.Name == name {
			return &t, nil
		}
	}

	return &tmpl{}, fmt.Errorf("Could not find template with name: %s", name)
}
