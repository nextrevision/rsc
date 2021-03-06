package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/nextrevision/rsc/config"
	"github.com/nextrevision/rsc/helper"

	"github.com/nextrevision/go-runscope"
)

// ListBuckets prints all buckets in the account
func (rc *RunscopeClient) ListBuckets(format string) error {

	buckets, err := rc.Runscope.ListBuckets()
	if err != nil {
		return err
	}

	if format == "json" {
		data, err := json.MarshalIndent(buckets, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	} else {
		header := []string{"Name", "Team", "Default"}
		rows := [][]string{}
		for _, b := range buckets {
			rows = append(rows, []string{b.Name, b.Team.Name, strconv.FormatBool(b.Default)})
		}
		helper.WriteTable(header, rows)
	}

	return nil
}

// ShowBucket prints details for a given bucket
func (rc *RunscopeClient) ShowBucket(name string, format string) error {

	bucket, err := rc.GetBucketByName(name)
	if err != nil {
		return err
	}

	// only supports json format for now
	data, err := json.MarshalIndent(bucket, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))

	return nil
}

// GetBucketByName returns a bucket object for a given bucket name
func (rc *RunscopeClient) GetBucketByName(name string) (runscope.Bucket, error) {
	buckets, err := rc.Runscope.ListBuckets()
	if err != nil {
		rc.Log.Error("Error getting bucket by name")
		return runscope.Bucket{}, err
	}

	for _, bucket := range buckets {
		if name == bucket.Name {
			rc.Log.Debugf("Found bucket by name: %s", bucket.Name)
			return bucket, nil
		}
	}

	return runscope.Bucket{}, fmt.Errorf("Could not find bucket: %s", name)
}

// CreateOrUpdateBucket searches for a bucket and creates
// a new one if not found, otherwise the bucket is updated
func (rc *RunscopeClient) CreateOrUpdateBucket(bc *config.BucketConfig, d bool) (runscope.Bucket, error) {
	buckets, err := rc.Runscope.ListBuckets()
	if err != nil {
		rc.Log.Error("Could not retrieve bucket list")
		return runscope.Bucket{}, err
	}

	for _, bucket := range buckets {
		if bucket.Name == bc.Name {
			// TODO: support updating buckets
			rc.Log.Debugf("Bucket already exists: %s", bc.Name)
			return bucket, nil
		}
	}

	if d {
		rc.Log.Infof("Would have created bucket: %s", bc.Name)
		return runscope.Bucket{}, nil
	}

	return rc.createBucket(bc)
}

// DeleteBucket deletes a bucket and all tests within
func (rc *RunscopeClient) DeleteBucket(b string) error {
	bucket, err := rc.GetBucketByName(b)
	if err != nil {
		rc.Log.Debugf("Bucket does not exist")
		return nil
	}

	rc.Log.Debugf("Deleting bucket: %s (%s)", bucket.Name, bucket.Key)
	return rc.Runscope.DeleteBucket(bucket.Key)
}

func (rc *RunscopeClient) createBucket(bc *config.BucketConfig) (runscope.Bucket, error) {
	rc.Log.Debugf("Creating bucket: %s", bc.Name)
	bucket, err := rc.Runscope.NewBucket(&runscope.NewBucketRequest{
		Name:     bc.Name,
		TeamUUID: bc.TeamID,
	})
	if err == nil {
		rc.Log.Infof("Created bucket: %s", bc.Name)
	} else {
		rc.Log.Errorf("Error creating bucket: %s", bc.Name)
	}
	return bucket, err
}
