package main

import "testing"

func TestGetFileContents(t *testing.T) {
	crabs, _ := getFileContents("./example")
	if len(crabs) != 10 {
		t.Fatalf("Amount of crab positions wrong, got: %d, want: %d", len(crabs), 10)
	}
}

func TestGetLeastExpensiveLinearFuelPosition(t *testing.T) {
	crabs, max := getFileContents("./example")
	res := CalculateLinearPositionFuels(crabs, max)
	if res[0].SumFuel != 37 {
		t.Fatalf("Least expensive consumption wrong, got: %d, want: %d", res[0].SumFuel, 37)
	}
}

func TestGetLeastExpensiveSummedFuelPosition(t *testing.T) {
	crabs, max := getFileContents("./example")
	res := CalculateSummedPositionFuels(crabs, max)
	if res[0].SumFuel != 168 {
		t.Fatalf("Least expensive position (linear calc) wrong, got: %d, want: %d", res[0].SumFuel, 168)
	}
}
