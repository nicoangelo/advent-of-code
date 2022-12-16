package day8

import (
	"github.com/nicoangelo/advent-of-code-2022/shared"
)

const ASCII_CODE_ZERO = 48

type TreeGrid struct {
	Visibility [][]int
}

func (g *TreeGrid) FillFromLines(lines []string) {
	g.Visibility = make([][]int, len(lines))
	for row, l := range lines {
		newRow := make([]int, len(l))
		for col, r := range l {
			newRow[col] = int(r) - ASCII_CODE_ZERO
		}
		g.Visibility[row] = newRow
	}
}

func (g *TreeGrid) CountVisibleTrees() int {
	visibleMap := map[[2]int]int{}

	// init from edges
	right := len(g.Visibility[0]) - 1
	for i := range g.Visibility {
		coordLeft := [2]int{i, 0}
		coordRight := [2]int{i, right}
		visibleMap[coordLeft] = 1
		visibleMap[coordRight] = 1
	}
	bottom := len(g.Visibility) - 1
	for i := range g.Visibility {
		coordLeft := [2]int{0, i}
		coordRight := [2]int{0, bottom}
		visibleMap[coordLeft] = 1
		visibleMap[coordRight] = 1
	}

	colMax := shared.MakeSliceInit(len(g.Visibility[0]), -1)

	for i := 0; i < len(g.Visibility); i++ {
		rowMax := g.Visibility[i][0]
		for j := 0; j < len(g.Visibility[i]); j++ {
			current := g.Visibility[i][j]
			coord := [2]int{i, j}
			// visible from left
			if current > rowMax {
				visibleMap[coord] = 1
				rowMax = current
				// continue
			}
			// visible from top
			if current > colMax[j] {
				visibleMap[coord] = 1
				colMax[j] = current
				// continue
			}
		}
	}

	colMax = shared.MakeSliceInit(len(g.Visibility[0]), -1)

	for i := len(g.Visibility) - 1; i >= 0; i-- {
		rowMax := g.Visibility[i][len(g.Visibility[i])-1]
		for j := len(g.Visibility[i]) - 1; j >= 0; j-- {
			current := g.Visibility[i][j]
			coord := [2]int{i, j}
			// visible from right
			if current > rowMax {
				visibleMap[coord] = 1
				rowMax = current
				// continue
			}
			// visible from bottom
			if current > colMax[j] {
				visibleMap[coord] = 1
				colMax[j] = current
				// continue
			}
		}
	}
	return len(visibleMap)
}

func (tg *TreeGrid) GetVisibilityScores() (res *[]int) {
	res = &[]int{}
	for row := 0; row < len(tg.Visibility)-1; row++ {
		for col := 0; col < len(tg.Visibility[0])-1; col++ {
			currentHeight := tg.Visibility[row][col]
			visibility := [4]int{}
			if row == 0 || col == 0 {
				continue
			}
			// look up
			for i := 1; i <= row; i++ {
				visibility[0] = i
				if tg.Visibility[row-i][col] >= currentHeight {
					break
				}
			}
			// look down
			for i := 1; i <= len(tg.Visibility)-1-row; i++ {
				visibility[1] = i
				if tg.Visibility[row+i][col] >= currentHeight {
					break
				}
			}
			// look left
			for i := 1; i <= col; i++ {
				visibility[2] = i
				if tg.Visibility[row][col-i] >= currentHeight {
					break
				}
			}
			// look right
			for i := 1; i <= len(tg.Visibility[row])-1-col; i++ {
				visibility[3] = i
				if tg.Visibility[row][col+i] >= currentHeight {
					break
				}
			}
			*res = append(*res, shared.Multiply(visibility[:]))
		}
	}
	return res
}
