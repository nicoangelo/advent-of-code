package day4

import (
	"strings"
)

func part1(lines []string) (sumAnyOverlapping int) {
	return loadInputAndCountByFunc(lines, func(cp *CleaningPairs) bool { return cp.IsAnyFullyContained() })
}

func part2(lines []string) int {
	return loadInputAndCountByFunc(lines, func(cp *CleaningPairs) bool { return cp.IsAnyOverlapping() })
}

func loadInputAndCountByFunc(lines []string, countFunc func(*CleaningPairs) bool) int {
	res := 0
	separators := ",-"
	for _, v := range lines {
		pairs := strings.FieldsFunc(v, func(r rune) bool { return strings.ContainsRune(separators, r) })
		cp := newCleaningPairFromTokens(pairs)
		if countFunc(cp) {
			res += 1
		}
	}
	return res
}
