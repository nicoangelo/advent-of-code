package day9

import "testing"

var testData = []string{
	"R 4",
	"U 4",
	"L 3",
	"D 1",
	"R 4",
	"D 1",
	"L 5",
	"R 2",
}

var testDataLarge = []string{
	"R 5",
	"U 8",
	"L 8",
	"D 3",
	"R 17",
	"D 10",
	"L 25",
	"U 20",
}

func TestPart1(t *testing.T) {
	res := part1(testData)
	want := 13
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}

func TestPart2(t *testing.T) {
	res := part2(testData)
	want := 1 // never moved, stayed at 0,0
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}

func TestPart2Large(t *testing.T) {
	res := part2(testDataLarge)
	want := 36
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}
