package day12

import (
	"log"
	"os/user"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/slicemath"
)

func PrintSolutions() {
	u, _ := user.Current()

	lines := reader.ReadInputFile("./day12/input_" + strings.ToLower(string(u.Username[0])))
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {

	res := 0

	matrix := linesToMatrix(lines)
	allCollections := map[rune]*FieldCollection{}

	for y := 0; y <= matrix.MaxY(); y++ {
		for x := 0; x <= matrix.MaxX(); x++ {
			tile := slicemath.Coord2D{X: x, Y: y}
			_, ok := allCollections[matrix.At(tile)]
			if !ok {
				allCollections[matrix.At(tile)] = &FieldCollection{fields: map[int]*Field{}, curFieldKey: 0}
			}
			allCollections[matrix.At(tile)].AddTile(matrix, tile)
		}
	}

	for _, v := range allCollections {
		for _, w := range v.fields {
			res += w.area * w.circumference
		}
	}

	return res
}

func part2(lines []string) int {
	return 0
}

func linesToMatrix(lines []string) *slicemath.Matrix2D[rune] {
	m := &slicemath.Matrix2D[rune]{}
	m.Init(slicemath.Coord2D{X: len(lines), Y: len(lines[0])})
	for y, l := range lines {
		for x, r := range l {
			m.Set(slicemath.Coord2D{X: x, Y: y}, r)
		}
	}
	return m
}
