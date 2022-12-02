package day1

import (
	"testing"
)

var testData = []string{
	"1000",
	"2000",
	"3000",
	"",
	"4000",
	"",
	"5000",
	"6000",
	"",
	"7000",
	"8000",
	"9000",
	"",
	"10000",
}

func TestPart1(t *testing.T) {
	res := part1(testData)
	if res != 24000 {
		t.Fatalf("got: %d, want: %d", res, 24000)
	}
}

func TestPart2(t *testing.T) {
	res := part2(testData)
	if res != 45000 {
		t.Fatalf("got: %d, want: %d", res, 45000)
	}
}
