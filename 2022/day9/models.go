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
	HeadPosition [2]int
	TailPosition [2]int
	TailHistory  *map[[2]int]int
}

func (r *Rope) IsTailAdjacent() bool {
	diff := shared.VectorDiff(r.HeadPosition, r.TailPosition)
	return shared.AbsInt(diff[0]) <= 1 && shared.AbsInt(diff[1]) <= 1
}

func (r *Rope) IsTailInSameRowOrCol() bool {
	return r.HeadPosition[0] == r.TailPosition[0] ||
		r.HeadPosition[1] == r.TailPosition[1]
}

func (r *Rope) MoveHead(in *Instruction) {
	for i := 0; i < in.Distance; i++ {
		r.HeadPosition = shared.VectorAdd(r.HeadPosition, in.DirectionUnityVector)
		isAdjacent := r.IsTailAdjacent()
		if !isAdjacent && !r.IsTailInSameRowOrCol() {
			diff_unity := shared.VectorUnity(shared.VectorDiff(r.HeadPosition, r.TailPosition))
			r.MoveTail(diff_unity)
		} else if !isAdjacent {
			r.MoveTail(in.DirectionUnityVector)
		}
	}
}

func (r *Rope) MoveTail(dir [2]int) {
	if r.TailHistory == nil {
		r.TailHistory = &map[[2]int]int{{0, 0}: 1}
	}
	r.TailPosition = shared.VectorAdd(r.TailPosition, dir)

	if _, ok := (*r.TailHistory)[r.TailPosition]; ok {
		(*r.TailHistory)[r.TailPosition] += 1
	} else {
		(*r.TailHistory)[r.TailPosition] = 1
	}
}
