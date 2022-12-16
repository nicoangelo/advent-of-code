package day8

import "testing"

var testData = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

func TestPart1(t *testing.T) {
	res := part1(testData)
	want := 21
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}

func TestPart2(t *testing.T) {
	res := part2(testData)
	want := 8
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}
