package client

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
