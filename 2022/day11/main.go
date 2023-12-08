package day11

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/nicoangelo/aoc-pkg/reader"
)

var logger = log.New(os.Stdout, "", 0)

func PrintSolutions() {
	logger.SetOutput(ioutil.Discard)
	lines := reader.ReadInputFile("./day11/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}
