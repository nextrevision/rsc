package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nextrevision/rsc/config"

	"github.com/nextrevision/go-runscope"
)

// ListTests prints all tests in a given bucket
func (rc *RunscopeClient) ListTests(b string, f string) error {
	rc.Log.Debugf("Listing tests in bucket '%s'", b)
	bucket, err := rc.GetBucketByName(b)
	if err != nil {
		return err
	}

	tests, err := rc.Runscope.ListTests(bucket.Key)
	if err != nil {
		return err
	}

	if f == "json" {
		data, err := json.MarshalIndent(*tests, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	} else {
		for _, t := range *tests {
			fmt.Println(t.Name)
		}
	}

	return nil
}

// ShowTest prints details for a given test
func (rc *RunscopeClient) ShowTest(b string, t string, f string) error {
	rc.Log.Debugf("Showing test '%s' in bucket '%s'", t, b)

	bucket, err := rc.GetBucketByName(b)
	if err != nil {
		return err
	}

	test, err := rc.GetTestByName(bucket.Key, t)
	if err != nil {
		return err
	}

	schedules, err := rc.Runscope.ListSchedules(bucket.Key, test.ID)
	if err != nil {
		return err
	}

	environments, err := rc.Runscope.ListTestEnvironments(bucket.Key, test.ID)
	if err != nil {
		return err
	}

	steps, err := rc.Runscope.ListSteps(bucket.Key, test.ID)
	if err != nil {
		return err
	}

	if f == "json" {
		jsonTest := test
		jsonTest.Schedules = *schedules
		jsonTest.Environments = *environments
		jsonTest.Steps = *steps
		data, err := json.MarshalIndent(jsonTest, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	} else {

		fmt.Printf("%s:\n", test.Name)
		fmt.Printf("  Name: %s\n", test.Name)
		fmt.Printf("  Bucket: %s (%s)\n", bucket.Name, bucket.Key)
		fmt.Printf("  ID: %s\n", test.ID)
		fmt.Printf("  Description: %s\n", test.Description)
		fmt.Printf("  TriggerURL: %s\n", test.TriggerURL)

		createdAt := time.Unix(int64(test.CreatedAt), 0)
		fmt.Printf("  Created: %s (%s)\n", createdAt.Format(time.RFC3339), test.CreatedBy.Name)

		finishedAt := time.Unix(int64(test.LastRun.FinishedAt), 0)
		fmt.Printf("  LastRun: %s (%d errors)\n", finishedAt.Format(time.RFC3339), test.LastRun.ErrorCount)

		if len(*schedules) != 0 {
			fmt.Printf("  Schedules (%d):\n", len(*schedules))
			for _, s := range *schedules {
				fmt.Printf("    %s", s.Interval)
				if s.Note != "" {
					fmt.Printf(" (%s)", s.Note)
				}
				fmt.Printf("\n")
			}
		}

		if len(*environments) != 0 {
			fmt.Printf("  Environments (%d):\n", len(*environments))
			for _, e := range *environments {
				fmt.Printf("    %s:\n", e.Name)
				if len(e.Regions) != 0 {
					fmt.Printf("    Regions (%d):\n", len(e.Regions))
				}
			}
		}

		if len(*steps) != 0 {
			rc.printSteps(*steps, "")
		}

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
func (rc *RunscopeClient) CreateOrUpdateTest(tc *config.TestConfig, d bool) error {
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
			if d {
				rc.Log.Infof("Would have updated test: %s", tc.Name)
				return nil
			}

			return rc.updateTest(tc, test.ID)
		}
	}

	if d {
		rc.Log.Infof("Would have created test: %s", tc.Name)
		return nil
	}
	return rc.createTest(tc)
}

// DeleteTest deletes a test from a bucket
func (rc *RunscopeClient) DeleteTest(b string, t string) error {
	bucket, err := rc.GetBucketByName(b)
	if err != nil {
		return err
	}

	test, err := rc.GetTestByName(bucket.Key, t)
	if err != nil {
		return err
	}

	rc.Log.Debugf("Deleting test from bucket '%s': %s (%s)", bucket.Name, test.Name, test.ID)
	return rc.Runscope.DeleteTest(bucket.Key, test.ID)
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

func (rc *RunscopeClient) printSteps(steps []runscope.Step, indent string) {
	baseIndent := indent + "  "
	fmt.Printf("%sSteps (%d):\n", baseIndent, len(steps))
	for i, s := range steps {
		i = i + 1
		if s.StepType == "request" {
			fmt.Printf("%s  %d. Request: %s %s\n", baseIndent, i, s.Method, s.URL)
		}

		if s.StepType == "condition" {
			fmt.Printf("%s  %d. Condition: %s %s %s\n", baseIndent, i, s.LeftValue, s.Comparison, s.RightValue)
		}

		if len(s.Variables) != 0 {
			fmt.Printf("%s    Variables: (%d):\n", baseIndent, len(s.Variables))
			for _, v := range s.Variables {
				varStr := fmt.Sprintf("%s = %s", v.Name, v.Source)
				if v.Property != "" {
					varStr = varStr + "." + v.Property
				}
				fmt.Printf("%s      %s\n", baseIndent, varStr)
			}
		}

		if len(s.Assertions) != 0 {
			fmt.Printf("%s    Assertions (%d):\n", baseIndent, len(s.Assertions))
			for _, a := range s.Assertions {
				aStr := a.Source
				if a.Property != "" {
					aStr = aStr + "." + a.Property
				}
				fmt.Printf("%s      %s %s %v\n", baseIndent, aStr, a.Comparison, a.Value.(interface{}))
			}
		}

		if len(s.Steps) != 0 {
			rc.printSteps(s.Steps, baseIndent)
		}
		fmt.Println("")
	}
}
