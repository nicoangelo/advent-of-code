package day8

import (
	"log"

	"github.com/nicoangelo/advent-of-code-2023/day8/nav"
	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./day8/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {
	m := &nav.Map{}
	m.FromLines(lines)

	return m.Navigate("AAA", "ZZZ")
}

func part2(lines []string) int {
	return 0
}
