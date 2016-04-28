package main

import (
	"flag"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/nextrevision/go-runscope"
)

var (
	path    string
	token   string
	debug   bool
	verbose bool
	client  *runscope.Client
)

func init() {
	flag.StringVar(&path, "path", ".", "path to search for files")
	flag.StringVar(&token, "token", "", "Runscope token")
	flag.BoolVar(&debug, "debug", false, "Enable debug logging")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	flag.Parse()
	if debug {
		log.SetLevel(log.DebugLevel)
	} else if verbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	osToken := os.Getenv("RUNSCOPE_TOKEN")

	if token == "" {
		if osToken == "" {
			log.Fatal("RUNSCOPE_TOKEN or -token must be set")
		}

		token = osToken
	}

	dir, err := isDirectory(path)
	if err != nil {
		log.Fatal(err)
	} else if !dir {
		log.Fatal("-path must be a directory")
	}

}

func main() {
	client = runscope.NewClient(&runscope.Options{Token: token})
	err := Run(path)
	if err != nil {
		log.Fatal(err)
	}
}

func Run(p string) error {
	configs, err := loadConfigs(p)
	if err != nil {
		return err
	}

	templates, err := loadTemplates(p)
	if err != nil {
		return err
	}

	if err = processBuckets(&configs, &templates); err != nil {
		return err
	}

	if err = processTests(&configs, &templates); err != nil {
		return err
	}

	return nil
}
