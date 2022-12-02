package day2

import "testing"

var testData = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestPart1(t *testing.T) {
	res := part1(testData)
	if res != 15 {
		t.Fatalf("got: %d, want: %d", res, 15)
	}
}

func TestPart2(t *testing.T) {
	res := part2(testData)
	if res != 12 {
		t.Fatalf("got: %d, want: %d", res, 12)
	}
}
