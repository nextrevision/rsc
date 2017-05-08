package config

import "testing"

func TestOrderTests(t *testing.T) {
	tests := []TestConfig{
		TestConfig{
			Name:    "b3",
			Depends: []string{"b2"},
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
			Depends: []string{"b2", "b3"},
		},
	}

	orderedTests, err := OrderTests(tests)
	if err != nil {
		t.Fatalf("Error ordering tests: %+v\n", err)
	}

	if orderedTests[0].Name != "b1" {
		t.Fatalf("First test is not b1: %+v", orderedTests)
	}
	if orderedTests[1].Name != "b2" {
		t.Fatalf("Second test is not b2: %+v", orderedTests)
	}
	if orderedTests[2].Name != "b3" {
		t.Fatalf("Third test is not b3: %+v", orderedTests)
	}
	if orderedTests[3].Name != "b4" {
		t.Fatalf("Last test is not b4: %+v", orderedTests)
	}
}
