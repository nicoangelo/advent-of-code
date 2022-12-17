package day9

import (
	"strconv"
	"strings"

	"github.com/nicoangelo/advent-of-code-2022/shared"
)

type Instruction struct {
	Direction            rune
	DirectionVector      [2]int // 0=x, 1=y
	DirectionUnityVector [2]int // 0=x, 1=y
	Distance             int
}

func (i *Instruction) FromLine(line string) {
	tokens := strings.Split(line, " ")
	i.Direction = rune(tokens[0][0])
	i.Distance, _ = strconv.Atoi(tokens[1])
	i.DirectionUnityVector = runeToVector(i.Direction)
	i.DirectionVector = shared.VectorMultiplyScalar(i.DirectionUnityVector, i.Distance)
}

func runeToVector(r rune) [2]int {
	if r == 'U' {
		return [2]int{0, 1}
	} else if r == 'R' {
		return [2]int{1, 0}
	} else if r == 'D' {
		return [2]int{0, -1}
	} else if r == 'L' {
		return [2]int{-1, 0}
	}
	return [2]int{0, 0}
}

type Rope struct {
	HeadPosition         [2]int
	KnotPositions        [][2]int
	KnotPositionsHistory []*map[[2]int]int
}

func NewRope(knots int) *Rope {
	knotHistory := shared.MakeSliceInit(knots, &map[[2]int]int{{0, 0}: 1})
	return &Rope{
		KnotPositions:        make([][2]int, knots),
		KnotPositionsHistory: knotHistory,
	}
}

func (r *Rope) IsKnotAdjacent(knot int) bool {
	diff := shared.VectorDiff(r.HeadPosition, r.KnotPositions[knot])
	return shared.AbsInt(diff[0]) <= 1 && shared.AbsInt(diff[1]) <= 1
}

func (r *Rope) IsKnotInSameRowOrCol(knot int) bool {
	return r.HeadPosition[0] == r.KnotPositions[knot][0] ||
		r.HeadPosition[1] == r.KnotPositions[knot][1]
}

func (r *Rope) MoveHead(in *Instruction) {
	for i := 0; i < in.Distance; i++ {
		r.HeadPosition = shared.VectorAdd(r.HeadPosition, in.DirectionUnityVector)
		for knot := 0; knot < len(r.KnotPositions); knot++ {
			isAdjacent := r.IsKnotAdjacent(knot)
			if !isAdjacent && !r.IsKnotInSameRowOrCol(knot) {
				diff_unity := shared.VectorUnity(shared.VectorDiff(r.HeadPosition, r.KnotPositions[knot]))
				r.MoveKnot(knot, diff_unity)
			} else if !isAdjacent {
				r.MoveKnot(knot, in.DirectionUnityVector)
			}
		}
	}
}

func (r *Rope) MoveKnot(knot int, dir [2]int) {
	r.KnotPositions[knot] = shared.VectorAdd(r.KnotPositions[knot], dir)

	if _, ok := (*r.KnotPositionsHistory[knot])[r.KnotPositions[knot]]; ok {
		(*r.KnotPositionsHistory[knot])[r.KnotPositions[knot]] += 1
	} else {
		(*r.KnotPositionsHistory[knot])[r.KnotPositions[knot]] = 1
	}
}
