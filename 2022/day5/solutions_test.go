package day5

import "testing"

var testData = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestPart1(t *testing.T) {
	res := part1(testData)
	want := "CMZ"
	if res != want {
		t.Fatalf("got: %s, want: %s", res, want)
	}
}

func TestPart2(t *testing.T) {
	res := part2(testData)
	want := "MCD"
	if res != want {
		t.Fatalf("got: %s, want: %s", res, want)
	}
}
