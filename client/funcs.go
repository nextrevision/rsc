package client

import (
	"strings"
	"time"
)

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

// intToRFC3339 converts epoch time to a RFC3339 date
func intToRFC3339(n int) string {
	u := time.Unix(int64(n), 0)
	return u.Format(time.RFC3339)
}

// floatToRFC3339 converts epoch time to a RFC3339 date
func floatToRFC3339(n float64) string {
	u := time.Unix(int64(n), 0)
	return u.Format(time.RFC3339)
}
