package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nicoangelo/advent-of-code-2022/shared"
)

type Instruction struct {
	DirectionVector [2]int // 0=x, 1=y
	Distance        int
}

func (i *Instruction) FromLine(line string) {
	tokens := strings.Split(line, " ")
	i.Distance, _ = strconv.Atoi(tokens[1])
	i.DirectionVector = runeToVector(rune(tokens[0][0]))
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
	BoundingBox          [4]int
}

func NewRope(knots int) *Rope {
	knotHistory := make([]*map[[2]int]int, knots)
	for i := 0; i < knots; i++ {
		knotHistory[i] = &map[[2]int]int{{0, 0}: 1}
	}
	return &Rope{
		KnotPositions:        make([][2]int, knots),
		KnotPositionsHistory: knotHistory,
		BoundingBox:          [4]int{},
	}
}

func (r *Rope) IsKnotAdjacentToPosition(knot int, tailPos [2]int) bool {
	diff := shared.VectorDiff(tailPos, r.KnotPositions[knot])
	return shared.AbsInt(diff[0]) <= 1 && shared.AbsInt(diff[1]) <= 1
}

func (r *Rope) IsKnotInSameRowOrColAsPosition(knot int, tailPos [2]int) bool {
	return tailPos[0] == r.KnotPositions[knot][0] ||
		tailPos[1] == r.KnotPositions[knot][1]
}

func (r *Rope) MoveHead(in *Instruction) {
	for i := 0; i < in.Distance; i++ {
		r.HeadPosition = shared.VectorAdd(r.HeadPosition, in.DirectionVector)
		currentHead := r.HeadPosition
		for knot := 0; knot < len(r.KnotPositions); knot++ {
			if r.IsKnotAdjacentToPosition(knot, currentHead) {
				break
			} else {
				diffToCurrentHead := shared.VectorDiff(currentHead, r.KnotPositions[knot])
				diff_unity := shared.VectorUnity(diffToCurrentHead)
				r.MoveKnot(knot, diff_unity)
				currentHead = r.KnotPositions[knot]
			}
			// r.PrintKnotPositions()
		}
	}
}

func (r *Rope) MoveKnot(knot int, dir [2]int) {
	r.KnotPositions[knot] = shared.VectorAdd(r.KnotPositions[knot], dir)
	pos := r.KnotPositions[knot]
	r.ExtendBoundingBox(pos)

	if _, ok := (*r.KnotPositionsHistory[knot])[pos]; ok {
		(*r.KnotPositionsHistory[knot])[pos] += 1
	} else {
		(*r.KnotPositionsHistory[knot])[pos] = 1
	}
}

func (r *Rope) ExtendBoundingBox(pos [2]int) {
	if pos[0] >= 0 && pos[0] > r.BoundingBox[0] {
		r.BoundingBox[0] = pos[0]
	}
	if pos[1] >= 0 && pos[1] > r.BoundingBox[1] {
		r.BoundingBox[1] = pos[1]
	}
	if pos[0] < 0 && pos[0] < r.BoundingBox[2] {
		r.BoundingBox[2] = pos[0]
	}
	if pos[1] < 0 && pos[1] < r.BoundingBox[3] {
		r.BoundingBox[3] = pos[1]
	}
}

func (r *Rope) PrintKnotPositions() {
	r.PrintBox(func(pos [2]int) (string, bool) { return r.KnotAtPosition(pos) })
}

func (r *Rope) PrintKnotHistoryPositions(knot int) {
	r.PrintBox(func(pos [2]int) (string, bool) { return r.KnotHistoryAtPosition(knot, pos) })
}

func (r *Rope) PrintBox(f func(pos [2]int) (string, bool)) {
	prevY := r.BoundingBox[1]
	fmt.Printf("%4d ", prevY)

	for y := r.BoundingBox[1]; y >= r.BoundingBox[3]; y-- {
		if y != prevY {
			fmt.Println()
			fmt.Printf("%4d ", y)
		}
		for x := r.BoundingBox[2]; x <= r.BoundingBox[0]; x++ {
			if x == 0 && y == 0 {
				fmt.Print("s")
			} else if value, ok := f([2]int{x, y}); ok {
				fmt.Print(value)
			} else {
				fmt.Print(".")
			}
		}
		prevY = y
	}
	fmt.Println()
	fmt.Println()
}

func (r *Rope) KnotAtPosition(pos [2]int) (v string, ok bool) {
	for i, k := range r.KnotPositions {
		if k == pos {
			return fmt.Sprint(i + 1), true
		}
	}
	return "", false
}

func (r *Rope) KnotHistoryAtPosition(knotIndex int, pos [2]int) (v string, ok bool) {
	_, ok = (*r.KnotPositionsHistory[knotIndex])[pos]
	return "#", ok
}
