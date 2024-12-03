package day2

import (
	"log"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/sliceutils"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./day2/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {

	safe := 0

	for _, line := range lines {

		l := sliceutils.SliceConvert(strings.Split(line, " "), strconv.Atoi)

		pos := (l[1]-l[0] > 0)

		s := true
		i := 1

		for s && (i < len(l)) {

			if pos {
				if (l[i]-l[i-1] < 1) || (l[i]-l[i-1] > 3) {
					s = false
				}
			} else {
				if (l[i]-l[i-1] > -1) || (l[i]-l[i-1] < -3) {
					s = false
				}
			}
			i += 1
		}

		if s {
			safe += 1
		}

	}

	return safe
}

func part2(lines []string) int {
	safe := 0

	for _, line := range lines {

		l := sliceutils.SliceConvert(strings.Split(line, " "), strconv.Atoi)

		if is_safe(l, 0) {
			safe += 1
		}

	}

	return safe
}

func is_safe(l []int, t int) bool {

	if t > 1 {
		return false
	}

	// init
	var pos int

	// take last entry cause second one could be eliminated
	if l[len(l)-1]-l[0] > 0 {
		pos = 1
	} else {
		pos = -1
	}

	for i := 1; i < len(l); i++ {

		if (pos*(l[i]-l[i-1]) < 1) || (pos*(l[i]-l[i-1]) > 3) {

			safe1 := is_safe(sliceutils.RemoveAtIndex(l, i), t+1)
			safe2 := is_safe(sliceutils.RemoveAtIndex(l, i-1), t+1)

			return safe1 || safe2
		}
	}

	return true
}
