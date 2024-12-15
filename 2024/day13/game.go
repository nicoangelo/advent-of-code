package day13

import (
	"fmt"
	"math"

	"github.com/nicoangelo/aoc-pkg/slicemath"
)

type Game struct {
	ButtonA   slicemath.Coord2D
	ButtonB   slicemath.Coord2D
	Prize     slicemath.Coord2D
	solutionA int
	solutionB int
}

func (g *Game) FromLines(lines []string) {
	if len(lines) != 3 {
		panic("Must be exactly 3 input lines")
	}
	for i, l := range lines {
		if i == 0 {
			fmt.Sscanf(l, "Button A: X+%d, Y+%d", &g.ButtonA.X, &g.ButtonA.Y)
		} else if i == 1 {
			fmt.Sscanf(l, "Button B: X+%d, Y+%d", &g.ButtonB.X, &g.ButtonB.Y)
		} else {
			fmt.Sscanf(l, "Prize: X=%d, Y=%d", &g.Prize.X, &g.Prize.Y)
		}
	}
	g.Solve()
}

func (g *Game) Solve() {
	// B=py*ax-ay*px/by*ax-bx*ay
	B := float64(g.Prize.Y*g.ButtonA.X-g.Prize.X*g.ButtonA.Y) / float64(g.ButtonB.Y*g.ButtonA.X-g.ButtonB.X*g.ButtonA.Y)
	// A=(px-bx*B)/ax
	A := (float64(g.Prize.X) - float64(g.ButtonB.X)*B) / float64(g.ButtonA.X)

	if math.Mod(A, 1) == 0 && math.Mod(B, 1) == 0 {
		g.solutionA = int(A)
		g.solutionB = int(B)
	} else {
		g.solutionB = 0
		g.solutionA = 0
	}
}

func (g *Game) GetSolution() (A int, B int) {
	return g.solutionA, g.solutionB
}

func (g *Game) RequiredTokens() int {
	if g.solutionA <= 100 && g.solutionB <= 100 {
		return g.solutionA*3 + g.solutionB*1
	}
	return 0
}
