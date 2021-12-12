package main

import (
	"fmt"
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
	part2 := resolveSignalsToDigits(entries)

	for _, v := range entries {
		fmt.Println(v)
	}

	if part2 != 26 {
		t.Fatalf("Amount of unique outputs wrong, got: %d, want: %d", part2, 26)
	}
}
