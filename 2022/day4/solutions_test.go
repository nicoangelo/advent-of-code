package day4

import "testing"

var testData = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestPart1(t *testing.T) {
	res := part1(testData)
	want := 2
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}

// func TestPart2(t *testing.T) {
// 	res := part2(testData)
// 	want := 12
// 	if res != want {
// 		t.Fatalf("got: %d, want: %d", res, want)
// 	}
// }
