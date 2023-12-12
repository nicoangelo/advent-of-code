package day4

import (
	"log"

	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./day4/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {
	sum := 0
	for _, l := range lines {
		c := Card{}
		c.FromLine(l)
		sum += c.GetWinningScore()
	}
	return sum
}

func part2(lines []string) int {
	sum := 0
	agioMap := map[int]int{}
	for _, l := range lines {
		c := Card{}
		c.FromLine(l)
		matchCount := c.GetWinningMatchCount()

		cardCount := 1 + agioMap[c.Number]
		sum += cardCount
		for i := 1; i <= matchCount; i++ {
			agioMap[c.Number+i] += cardCount
		}
	}
	return sum
}
