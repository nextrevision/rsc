package config

import "testing"

func TestOrderTests(t *testing.T) {
	tests := []TestConfig{
		TestConfig{
			Name:    "b3",
			Depends: []string{"b2", "b4"},
		},
		TestConfig{
			Name:    "b1",
			Depends: []string{},
		},
		TestConfig{
			Name:    "b2",
			Depends: []string{"b1"},
		},
		TestConfig{
			Name:    "b4",
			Depends: []string{},
		},
	}

	orderedTests, err := OrderTests(tests)
	if err != nil {
		t.Fatal(err)
	}

	if orderedTests[0].Name != "b1" {
		t.Fatalf("First bucket is not b1")
	}
	if orderedTests[1].Name != "b4" {
		t.Fatalf("First bucket is not b4")
	}
	if orderedTests[2].Name != "b2" {
		t.Fatalf("First bucket is not b2")
	}
	if orderedTests[3].Name != "b3" {
		t.Fatalf("First bucket is not b3")
	}
}
