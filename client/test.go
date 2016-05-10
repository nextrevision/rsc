package client

import (
	"fmt"

	"github.com/nextrevision/rsc/config"

	"github.com/nextrevision/go-runscope"
)

// ListTests prints all tests in a given bucket
func (rc *RunscopeClient) ListTests(b string) error {
	bucket, err := rc.GetBucketByName(b)
	if err != nil {
		return err
	}

	tests, err := rc.Runscope.ListTests(bucket.Key)
	if err != nil {
		return err
	}

	for _, t := range *tests {
		fmt.Println(t.Name)
	}

	return nil
}

// GetTestByName searches for a test in the supplied bucket
// and returns a Test object if found
func (rc *RunscopeClient) GetTestByName(bucketKey string, testName string) (runscope.Test, error) {
	var test = runscope.Test{}

	tests, err := rc.Runscope.ListTests(bucketKey)
	if err != nil {
		return test, err
	}

	for _, t := range *tests {
		if t.Name == testName {
			return t, nil
		}
	}

	return test, fmt.Errorf("No such test: %s", testName)
}

// CreateOrUpdateTest searches for a test in a bucket and
// creates a new test if it does not exists, otherwise the
// test is updated
func (rc *RunscopeClient) CreateOrUpdateTest(tc *config.TestConfig) error {
	if tc.BucketKey == "" {
		bucket, err := rc.GetBucketByName(tc.Bucket)
		if err != nil {
			rc.Log.Errorf("Could determine bucket for test: %s", tc.Name)
			return err
		}

		tc.BucketKey = bucket.Key
	}

	tests, err := rc.Runscope.ListTests(tc.BucketKey)
	if err != nil {
		rc.Log.Errorf("Could list tests for bucket: %s", tc.BucketKey)
		return err
	}

	for _, test := range *tests {
		if test.Name == tc.Name {
			return rc.updateTest(tc, test.ID)
		}
	}

	return rc.createTest(tc)
}

func (rc *RunscopeClient) createTest(tc *config.TestConfig) error {
	rc.Log.Debugf("Creating test: %s", tc.Name)

	_, err := rc.Runscope.ImportTest(tc.BucketKey, tc.Bytes)
	if err == nil {
		rc.Log.Infof("Created test: %s", tc.Name)
	} else {
		rc.Log.Errorf("Error creating test: %s", tc.Name)
	}

	return err
}

func (rc *RunscopeClient) updateTest(tc *config.TestConfig, testID string) error {
	rc.Log.Debugf("Updating test: %s", tc.Name)

	_, err := rc.Runscope.ReimportTest(tc.BucketKey, testID, tc.Bytes)
	if err == nil {
		rc.Log.Infof("Updated test: %s", tc.Name)
	} else {
		rc.Log.Errorf("Error updating test: %s", tc.Name)
	}

	return err
}
