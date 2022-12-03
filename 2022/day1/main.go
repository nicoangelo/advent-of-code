package day1

import (
	"log"

	"github.com/nicoangelo/advent-of-code-2022/shared"
)

func PrintSolutions() {
	lines := shared.ReadInputFile("./day1/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}