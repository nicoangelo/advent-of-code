package day4

import (
	"strings"
)

func part1(lines []string) (sumAnyOverlapping int) {
	separators := ",-"
	for _, v := range lines {
		pairs := strings.FieldsFunc(v, func(r rune) bool { return strings.ContainsRune(separators, r) })
		cp := newCleaningPairFromTokens(pairs)
		if cp.IsAnyFullyContained() {
			sumAnyOverlapping += 1
		}
	}
	return sumAnyOverlapping
}

func part2(lines []string) int {
	return 0
}
