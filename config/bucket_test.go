package config

import "testing"

func TestOrderBuckets(t *testing.T) {
	buckets := []BucketConfig{
		BucketConfig{
			Name:    "b3",
			Depends: []string{"b2"},
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
			Depends: []string{"b2", "b3"},
		},
	}

	orderedBuckets, err := OrderBuckets(buckets)
	if err != nil {
		t.Fatalf("Error ordering buckets: %+v", err)
	}

	if orderedBuckets[0].Name != "b1" {
		t.Fatalf("First bucket is not b1: %+v", orderedBuckets)
	}
	if orderedBuckets[1].Name != "b2" {
		t.Fatalf("Second bucket is not b2: %+v", orderedBuckets)
	}
	if orderedBuckets[2].Name != "b3" {
		t.Fatalf("Third bucket is not b3: %+v", orderedBuckets)
	}
	if orderedBuckets[3].Name != "b4" {
		t.Fatalf("Last bucket is not b4: %+v", orderedBuckets)
	}
}
