package main

import "testing"

func TestGetFileContents(t *testing.T) {
	crabs, _ := getFileContents("./example")
	if len(crabs) != 10 {
		t.Fatalf("Amount of crab positions wrong, got: %d, want: %d", len(crabs), 10)
	}
}

func TestGetLeastExpensivePosition(t *testing.T) {
	crabs, max := getFileContents("./example")
	res := CalculatePositionFuels(crabs, max)
	if res[0].Position != 2 {
		t.Fatalf("Least expensive position wrong, got: %d, want: %d", res, 2)
	}
}
