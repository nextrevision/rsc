package client

import "strings"

// TriggerURL returns the trigger URL for a given test
func (rc *RunscopeClient) TriggerURL(bucketName string, testName string) (string, error) {
	bucket, err := rc.GetBucketByName(bucketName)
	if err != nil {
		return "", err
	}

	test, err := rc.GetTestByName(bucket.Key, testName)
	if err != nil {
		return "", err
	}

	return test.TriggerURL, nil
}

// BatchURL returns the batch URL for a given test
func (rc *RunscopeClient) BatchURL(bucketName string, testName string) (string, error) {
	bucket, err := rc.GetBucketByName(bucketName)
	if err != nil {
		return "", err
	}

	test, err := rc.GetTestByName(bucket.Key, testName)
	if err != nil {
		return "", err
	}

	return strings.Replace(test.TriggerURL, "trigger", "batch", 1), nil
}
