package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/nextrevision/go-runscope"
)

type bucketConfig struct {
	Name     string          `json:"name"`
	Template string          `json:"template"`
	TeamID   string          `json:"team_id"`
	Config   runscope.Bucket `json:"config"`
}

func processBuckets(configs *[]config, templates *[]tmpl) error {
	log.Info("Processing buckets...")
	for _, config := range *configs {
		if config.Buckets != nil {
			for _, bucket := range config.Buckets {
				log.Debugf("Processing bucket: %s", bucket.Name)

				if bucket.TeamID == "" {
					defaultTeam, err := getDefaultTeam()
					if err != nil {
						log.Error("Could not get default team")
						return err
					}

					bucket.TeamID = defaultTeam.ID
				}

				if err := bucket.createIfNotExists(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (b *bucketConfig) createIfNotExists() error {
	buckets, err := client.ListBuckets()
	if err != nil {
		log.Error("Could not retrieve bucket list")
		return err
	}

	for _, bucket := range *buckets {
		if bucket.Name == b.Name {
			// TODO: support updating buckets
			log.Debugf("Bucket already exists: %s", b.Name)
			return nil
		}
	}

	return b.create()
}

func (b *bucketConfig) create() error {
	log.Debugf("Creating bucket: %s", b.Name)
	_, err := client.NewBucket(&runscope.NewBucketRequest{
		Name:     b.Name,
		TeamUUID: b.TeamID,
	})
	if err == nil {
		log.Infof("Created bucket: %s", b.Name)
	} else {
		log.Errorf("Error creating bucket: %s", b.Name)
	}
	return err
}

func getBucketByName(name string) (*runscope.Bucket, error) {
	buckets, err := client.ListBuckets()
	if err != nil {
		log.Error("Error getting bucket by name")
		return &runscope.Bucket{}, err
	}

	for _, bucket := range *buckets {
		if name == bucket.Name {
			log.Debugf("Found bucket by name: %s", bucket.Name)
			return &bucket, nil
		}
	}

	return &runscope.Bucket{}, fmt.Errorf("Could not find bucket: %s", name)
}
