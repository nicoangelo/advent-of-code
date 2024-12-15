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

func (g *Game) FromLines(lines []string, locationIncrease int) {
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
	g.Prize.X += locationIncrease
	g.Prize.Y += locationIncrease
}

/*
 * Solve solves the equation system with two unknowns (Button presses A and B).
 * The two equations are:
 *   a_x*A + b_x*B = p_x
 *   a_y*A + b_y*B = p_y
 *
 *   a_x/a_y: increase on X/Y when pressing Button A
 *   b_x/b_y: increase on X/Y when pressing Button B
 *   p_x/p_y: location of Prize on X/Y axis
 *   A/B: number of button presses
 *
 * Solving the first for A so that it can be used in the second equation which then is solved for B:
 *   A = (p_x - b_x*B) / a_x
 *   a_y*((p_x - b_x*B) / a_x) + b_y*B = p_y
 *   B = (p_y * a_x - a_y * p_x) / (b_y * a_x - b_x * a_y)
 * Boundary condition: A and B must be integers
**/
func (g *Game) Solve() {
	B := float64(g.Prize.Y*g.ButtonA.X-g.Prize.X*g.ButtonA.Y) / float64(g.ButtonB.Y*g.ButtonA.X-g.ButtonB.X*g.ButtonA.Y)
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

func (g *Game) RequiredTokensWithLimit(limit int) int {
	if limit <= 0 || (limit > 0 && g.solutionA <= limit && g.solutionB <= limit) {
		return g.solutionA*3 + g.solutionB*1
	}
	return 0
}
