package main

func triggerURL(bucketName string, testName string) (string, error) {
	bucket, err := getBucketByName(bucketName)
	if err != nil {
		return "", err
	}

	test, err := getTestByName(bucket.Key, testName)
	if err != nil {
		return "", err
	}

	return test.TriggerURL, nil
}
