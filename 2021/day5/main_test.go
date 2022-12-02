package main

import "testing"

func TestGetFileContents(t *testing.T) {
	vectors, _, _ := getFileContents("./example")
	if len(vectors) != 10 {
		t.Fatalf("Amount of vectors wrong, got: %d, want: %d", len(vectors), 10)
	}
}

func TestPart1(t *testing.T) {
	vectors, maxX, maxY := getFileContents("./example")
	system := markPerpendicularOnSystem(vectors, maxX, maxY)
	res := countCrossingPoints(system)
	if res != 5 {
		t.Fatalf("Number of crossing points wrong, got: %d, want: %d", res, 5)
	}
}

func TestPart2(t *testing.T) {
	println()
	vectors, maxX, maxY := getFileContents("./example")
	system := markAllOnSystem(vectors, maxX, maxY)
	res := countCrossingPoints(system)
	if res != 12 {
		t.Fatalf("Number of crossing points wrong, got: %d, want: %d", res, 12)
	}
}
