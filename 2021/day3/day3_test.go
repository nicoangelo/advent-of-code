package main

import "testing"

const power_consumption = 198
const life_support_rating = 230

func TestPart1(t *testing.T) {
	contents, line_count := get_file_contents("./example")
	res := part1(contents, line_count)
	if res != power_consumption {
		t.Fatalf("Epsilon wrong, got: %d, want: %d", res, power_consumption)
	}
}

func TestPart2(t *testing.T) {
	contents, line_count := get_file_contents("./example")
	res := part2(contents, line_count)
	if res != life_support_rating {
		t.Fatalf("Life support rating wrong, got: %d, want: %d", res, life_support_rating)
	}
}
