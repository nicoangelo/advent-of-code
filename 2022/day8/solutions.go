package day8

func part1(lines []string) int {
	tg := &TreeGrid{}
	tg.FillFromLines(lines)

	return tg.CountVisibleTrees()
}

func part2(lines []string) int {
	return 0
}
