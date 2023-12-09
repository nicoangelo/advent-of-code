package day3

import (
	"log"

	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./day3/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {
	s := &Schematic{}
	s.fromLines(lines)
	return s.PartsSum()
}

func part2(lines []string) int {
	return 0
}
