package day8

import (
	"sort"
)

func part1(lines []string) int {
	tg := &TreeGrid{}
	tg.FillFromLines(lines)

	return tg.CountVisibleTrees()
}

func part2(lines []string) int {
	tg := &TreeGrid{}
	tg.FillFromLines(lines)

	scores := tg.GetVisibilityScores()

	sort.Sort(sort.Reverse(sort.IntSlice(*scores)))
	return (*scores)[0]
}
