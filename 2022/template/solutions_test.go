package dayX

import "testing"

var testData = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestPart1(t *testing.T) {
	res := part1(testData)
	want := 15
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}

func TestPart2(t *testing.T) {
	res := part2(testData)
	want := 12
	if res != want {
		t.Fatalf("got: %d, want: %d", res, want)
	}
}
