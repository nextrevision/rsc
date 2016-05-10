package client

import (
	"fmt"

	"github.com/nextrevision/rsc/config"

	"github.com/nextrevision/go-runscope"
)

// ListBuckets prints all buckets in the account
func (rc *RunscopeClient) ListBuckets() error {

	buckets, err := rc.Runscope.ListBuckets()
	if err != nil {
		return err
	}

	for _, b := range *buckets {
		fmt.Println(b.Name)
	}

	return nil
}

// GetBucketByName returns a bucket object for a given bucket name
func (rc *RunscopeClient) GetBucketByName(name string) (runscope.Bucket, error) {
	buckets, err := rc.Runscope.ListBuckets()
	if err != nil {
		rc.Log.Error("Error getting bucket by name")
		return runscope.Bucket{}, err
	}

	for _, bucket := range *buckets {
		if name == bucket.Name {
			rc.Log.Debugf("Found bucket by name: %s", bucket.Name)
			return bucket, nil
		}
	}

	return runscope.Bucket{}, fmt.Errorf("Could not find bucket: %s", name)
}

// CreateOrUpdateBucket searches for a bucket and creates
// a new one if not found, otherwise the bucket is updated
func (rc *RunscopeClient) CreateOrUpdateBucket(bc *config.BucketConfig) (runscope.Bucket, error) {
	buckets, err := rc.Runscope.ListBuckets()
	if err != nil {
		rc.Log.Error("Could not retrieve bucket list")
		return runscope.Bucket{}, err
	}

	for _, bucket := range *buckets {
		if bucket.Name == bc.Name {
			// TODO: support updating buckets
			rc.Log.Debugf("Bucket already exists: %s", bc.Name)
			return bucket, nil
		}
	}

	return rc.createBucket(bc)
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
	return *bucket, err
}
