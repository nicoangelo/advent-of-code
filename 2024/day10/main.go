package day10

import (
	"log"
	"os/user"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	u, _ := user.Current()

	lines := reader.ReadInputFile("./day10/input_" + strings.ToLower(string(u.Username[0])))
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {

	for r, l := range lines {
		for c := 0; c < len(l); c++ {
			if(l[c]=="0")
			
		}
	}
	return 0
}

func part2(lines []string) int {
	return 0
}

func linesToMatrix(lines []string) *slicemath.Matrix2D[int] {
	m := &slicemath.Matrix2D[int]{}
	m.Init(slicemath.Coord2D{X: len(lines), Y: len(lines[0])})
	for y, l := range lines {
		for x, r := range l {
			m.Set(slicemath.Coord2D{X: x, Y: y}, strconv.Atoi(r))
		}
	}
	return m
}