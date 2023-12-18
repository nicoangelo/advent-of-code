package day7

import (
	"log"
	"sort"

	"github.com/nicoangelo/advent-of-code-2023/day7/camelcards"
	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./day7/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {
	hands := make([]camelcards.Hand, len(lines))
	for i, l := range lines {
		hands[i].FromLine(l)
	}
	sort.Slice(hands, func(i, j int) bool { return hands[i].Compare(&hands[j]) < 0 })
	sum := 0
	for i, h := range hands {
		// fmt.Println(h)
		sum += (i + 1) * h.GetBid()
	}
	return sum
}

func part2(lines []string) int {
	return 0
}
