package day7

import (
	"log"

	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./day7/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}
