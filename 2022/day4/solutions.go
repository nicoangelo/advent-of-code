package day4

func part1(lines []string) (sumAnyOverlapping int) {
	return loadInputAndCountByFunc(lines, func(cp *CleaningPairs) bool { return cp.IsAnyFullyContained() })
}

func part2(lines []string) int {
	return loadInputAndCountByFunc(lines, func(cp *CleaningPairs) bool { return cp.IsAnyOverlapping() })
}

func loadInputAndCountByFunc(lines []string, countFunc func(*CleaningPairs) bool) int {
	res := 0
	for _, v := range lines {
		cp := newCleaningPairFromTokens(v)
		if countFunc(cp) {
			res += 1
		}
	}
	return res
}
