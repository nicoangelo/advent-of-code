package day19

import (
	"log"
	"os/user"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	u, _ := user.Current()

	towels := reader.ReadInputIntoStruct(
		"./day19/input_"+strings.ToLower(string(u.Username[0])),
		func(s string, t *Towels) {
			for _, v := range strings.Split(s, ",") {
				t.AddAvailablePattern(strings.TrimSpace(v))
			}
		},
		func(s string, t *Towels) {
			t.AddDesiredPattern(s)
		},
	)
	part1 := part1(towels)
	log.Println("Part 1: ", part1)

	part2 := part2(towels)
	log.Println("Part 2: ", part2)
}

func part1(towels *Towels) int {
	return towels.TryFitPatterns()
}

func part2(towels *Towels) int {
	return 0
}
