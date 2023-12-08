package day6

import (
	"log"

	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	line := reader.ReadInputFile("./day6/input")[0]
	part1 := part1(line)
	log.Println("Part 1: ", part1)

	part2 := part2(line)
	log.Println("Part 2: ", part2)
}
