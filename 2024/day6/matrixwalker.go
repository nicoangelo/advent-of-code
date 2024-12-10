package day6

import (
	"fmt"

	"github.com/nicoangelo/aoc-pkg/slicemath"
)

var directions []slicemath.Coord2D = []slicemath.Coord2D{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
}

type MatrixWalker struct {
	matrix          slicemath.Matrix2D[rune]
	currPos         slicemath.Coord2D
	visited         map[slicemath.Coord2D]bool
	currentDirIndex int
	turns           map[slicemath.Coord2D]bool
}

func (mw *MatrixWalker) Init(m *slicemath.Matrix2D[rune]) {
	mw.matrix = *m
}

func (mw *MatrixWalker) SetStart() {
	mw.visited = map[slicemath.Coord2D]bool{}
	mw.turns = map[slicemath.Coord2D]bool{}
	mw.currentDirIndex = 0
	currPos, ok := mw.matrix.FindFirst('^')
	if ok {
		mw.currPos = currPos
	} else {
		panic("Cannot find starting position")
	}
}

func (mw *MatrixWalker) Walk() (more bool, turned bool) {
	lookAtPos := mw.currPos.Add(directions[mw.currentDirIndex])
	if mw.matrix.IsOutOfBounds(lookAtPos) {
		return false, false
	}
	if mw.matrix.At(lookAtPos) == '#' {
		mw.turns[mw.currPos] = true
		mw.turnRight()
		return true, true
	}
	mw.currPos = lookAtPos
	mw.visited[mw.currPos] = true
	return true, false
}

func (mw *MatrixWalker) VisitedPlaces() int {
	return len(mw.visited)
}

func (mw *MatrixWalker) Print() {
	for y := 0; y < mw.matrix.MaxY(); y++ {
		for x := 0; x < mw.matrix.MaxX(); x++ {
			pos := slicemath.Coord2D{X: x, Y: y}
			v := mw.matrix.At(pos)
			visited := mw.visited[pos]
			if v == 0 {
				v = ' '
			}
			if visited {
				v = 'X'
			}
			fmt.Print(string(v))
		}
		fmt.Println()
	}
}

func (mw *MatrixWalker) WouldTurnInDirection(dir slicemath.Coord2D) bool {
	newPos := mw.currPos
	for {
		newPos = newPos.Add(dir)
		if mw.matrix.IsOutOfBounds(newPos) {
			return false
		}
		if _, ok := mw.turns[newPos]; ok {
			return true
		}
	}
}

func (mw *MatrixWalker) GetRightTurnVector() slicemath.Coord2D {
	if mw.currentDirIndex == 3 {
		return directions[0]
	}
	return directions[mw.currentDirIndex+1]
}

func (mw *MatrixWalker) turnRight() {
	if mw.currentDirIndex == 3 {
		mw.currentDirIndex = 0
	} else {
		mw.currentDirIndex++
	}
}
