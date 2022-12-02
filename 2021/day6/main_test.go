package main

import "testing"

func TestGetFileContents(t *testing.T) {
	fish := getFileContents("./example")
	if len(fish) != 5 {
		t.Fatalf("Amount of fish wrong, got: %d, want: %d", len(fish), 5)
	}
}

func TestGetPart1After80Days(t *testing.T) {
	fish := getFileContents("./example")
	res := part1(fish, 80)
	if res != 5934 {
		t.Fatalf("Amount of fish wrong, got: %d, want: %d", res, 5934)
	}
}

func TestGetPart1After256Days(t *testing.T) {
	fish := getFileContents("./example")
	res := part1(fish, 256)
	if res != 26984457539 {
		t.Fatalf("Amount of fish wrong, got: %d, want: %d", res, 26984457539)
	}
}
