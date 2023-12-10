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

func part1(lines []string) (sum int) {
	includeChars := []rune{'#', '$', '%', '&', '*', '-', '/', '=', '@', '+'}
	s := &Schematic{}
	s.fromLines(lines, includeChars)

	for _, p := range s.Parts {
		sum += p.Number
	}
	return sum
}

func part2(lines []string) (sum int) {
	includeChars := []rune{'*'}
	s := &Schematic{}
	s.fromLines(lines, includeChars)

	partAreas := map[*PartArea]int{}
	for _, p := range s.Parts {
		_, ok := partAreas[p.PartArea]

		if p.PartArea.PartCount == 2 {
			if !ok {
				partAreas[p.PartArea] = p.Number
			} else {
				partAreas[p.PartArea] *= p.Number
			}
		}
	}

	for _, v := range partAreas {
		sum += v
	}
	return sum
}
