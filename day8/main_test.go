package main

import (
	"testing"
)

func TestGetFileContents(t *testing.T) {
	entries := getFileContents("./example")
	if len(entries) != 10 {
		t.Fatalf("Amount of entries wrong, got: %d, want: %d", len(entries), 10)
	}
}

func TestCountUniqueOutputs(t *testing.T) {
	entries := getFileContents("./example")
	part1 := countUniqueOutputs(entries)
	if part1 != 26 {
		t.Fatalf("Amount of unique outputs wrong, got: %d, want: %d", part1, 26)
	}
}

func TestResolveSignals(t *testing.T) {
	entries := getFileContents("./example")
	part2 := resolveOutputToNumbers(entries)

	if len(part2) != 10 {
		t.Fatalf("Amount of unique outputs wrong, got: %d, want: %d", len(part2), 10)
	}
	numberWants := []int{
		8394,
		9781,
		1197,
		9361,
		4873,
		8418,
		4548,
		1625,
		8717,
		4315,
	}

	for i, want := range numberWants {
		if part2[i] != want {
			t.Fatalf("Number 0 wrong, got %d, want: %d", part2[i], want)
		}
	}

}
