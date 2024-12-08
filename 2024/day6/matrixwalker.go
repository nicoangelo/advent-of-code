package day6

import (
	"fmt"

	"github.com/nicoangelo/aoc-pkg/slicemath"
)

var lookingDirMap map[rune]slicemath.Coord2D = map[rune]slicemath.Coord2D{
	'^': {X: 0, Y: -1},
	'>': {X: 1, Y: 0},
	'v': {X: 0, Y: 1},
	'<': {X: -1, Y: 0},
}

type MatrixWalker struct {
	matrix     slicemath.Matrix2D[rune]
	currPos    slicemath.Coord2D
	visited    map[slicemath.Coord2D]bool
	lookingDir rune
}

func (mw *MatrixWalker) Init(m *slicemath.Matrix2D[rune]) {
	mw.matrix = *m
}

func (mw *MatrixWalker) SetStart() {
	mw.visited = map[slicemath.Coord2D]bool{}
	mw.lookingDir = '^'
	currPos, ok := mw.matrix.FindFirst(mw.lookingDir)
	if ok {
		mw.currPos = currPos
	} else {
		panic("Cannot find starting position")
	}
}

func (mw *MatrixWalker) Walk() bool {
	dir := lookingDirMap[mw.lookingDir]
	lookAtPos := mw.currPos.Add(dir)
	if mw.matrix.IsOutOfBounds(lookAtPos) {
		return false
	}
	if mw.matrix.At(lookAtPos) == '#' {
		mw.turnRight()
		return true
	}
	mw.currPos = lookAtPos
	mw.visited[mw.currPos] = true
	return true
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

func (mw *MatrixWalker) turnRight() {
	switch mw.lookingDir {
	case '^':
		mw.lookingDir = '>'
	case '>':
		mw.lookingDir = 'v'
	case 'v':
		mw.lookingDir = '<'
	case '<':
		mw.lookingDir = '^'
	}
}
