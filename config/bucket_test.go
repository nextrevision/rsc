package config

import "testing"

func TestOrderBuckets(t *testing.T) {
	buckets := []BucketConfig{
		BucketConfig{
			Name:    "b3",
			Depends: []string{"b2", "b4"},
		},
		BucketConfig{
			Name:    "b1",
			Depends: []string{},
		},
		BucketConfig{
			Name:    "b2",
			Depends: []string{"b1"},
		},
		BucketConfig{
			Name:    "b4",
			Depends: []string{},
		},
	}

	orderedBuckets, err := OrderBuckets(buckets)
	if err != nil {
		t.Fatal(err)
	}

	if orderedBuckets[0].Name != "b1" {
		t.Fatalf("First bucket is not b1")
	}
	if orderedBuckets[1].Name != "b4" {
		t.Fatalf("First bucket is not b4")
	}
	if orderedBuckets[2].Name != "b2" {
		t.Fatalf("First bucket is not b2")
	}
	if orderedBuckets[3].Name != "b3" {
		t.Fatalf("First bucket is not b3")
	}
}
