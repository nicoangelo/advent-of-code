package day6

import "testing"

var testData1 = map[string]int{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
	"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
	"nppdvjthqldpwncqszvftbrmjlhg":      6,
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
}
var testData2 = map[string]int{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
	"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
	"nppdvjthqldpwncqszvftbrmjlhg":      23,
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
}

func TestPart1(t *testing.T) {
	for k, v := range testData1 {
		res := part1(k)
		want := v
		if res != want {
			t.Fatalf("%s, got: %d, want: %d", k, res, want)
		}
	}
}

func TestPart2(t *testing.T) {
	for k, v := range testData2 {
		res := part2(k)
		want := v
		if res != want {
			t.Fatalf("%s, got: %d, want: %d", k, res, want)
		}
	}
}
