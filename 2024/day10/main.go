package day10

import (
	"log"
	"os/user"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/slicemath"
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

	matrix := linesToMatrix(lines)

	counts := 0

	for y := 0; y <= matrix.MaxY(); y++ {
		for x := 0; x <= matrix.MaxX(); x++ {

			curPos := slicemath.Coord2D{X: x, Y: y}

			if matrix.At(curPos) == 0 {

				nines := findPaths(matrix, curPos, map[slicemath.Coord2D]bool{})

				counts += len(nines)
			}
		}
	}

	return counts
}

func part2(lines []string) int {
	return 0
}

func linesToMatrix(lines []string) *slicemath.Matrix2D[int] {
	m := &slicemath.Matrix2D[int]{}
	m.Init(slicemath.Coord2D{X: len(lines), Y: len(lines[0])})
	for y, l := range lines {
		for x, r := range l {
			v, _ := strconv.Atoi(string(r))
			m.Set(slicemath.Coord2D{X: x, Y: y}, v)
		}
	}
	return m
}

func findPaths(matrix *slicemath.Matrix2D[int], pos slicemath.Coord2D, nines map[slicemath.Coord2D]bool) map[slicemath.Coord2D]bool {

	curVal := matrix.At(pos)

	if curVal == 9 {
		nines[pos] = true
		return (nines)
	}

	for _, d := range directions {
		newPos := pos.Add(d)
		if !matrix.IsOutOfBounds(newPos) && matrix.At(newPos) == curVal+1 {

			nines = findPaths(matrix, newPos, nines)

		}
	}

	return (nines)

}

var directions []slicemath.Coord2D = []slicemath.Coord2D{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
}
