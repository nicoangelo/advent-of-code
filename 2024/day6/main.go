package day6

import (
	"log"

	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/slicemath"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./day6/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {
	mw := &MatrixWalker{}
	mw.Init(linesToMatrix(lines))
	mw.SetStart()
	for {
		if more, _ := mw.Walk(); !more {
			break
		}
	}
	// mw.Print()
	return mw.VisitedPlaces()
}

func part2(lines []string) int {
	mw := &MatrixWalker{}
	mw.Init(linesToMatrix(lines))
	mw.SetStart()
	newObstacles := 0
	for {
		more, turned := mw.Walk()
		if !more {
			break
		}
		if turned {
			continue
		}
		// check if a turn here would hit an obstacle (i.e. known previous turn)
		// assumption: potential obstacles have been visited before
		tryTurn := mw.GetRightTurnVector()
		if mw.WouldTurnInDirection(tryTurn) {
			newObstacles += 1
		}
	}
	// mw.Print()
	return newObstacles
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
