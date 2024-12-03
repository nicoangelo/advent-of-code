package day2

import (
	"log"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/slices"
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

		l := slices.SliceConvert(strings.Split(line, " "), strconv.Atoi)

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

		l := slices.SliceConvert(strings.Split(line, " "), strconv.Atoi)

		if is_safe(l, 0) {
			safe += 1
		}

	}

	return safe
}

func delete_at_index(slice []int, index int) []int {
	slice_new := make([]int, len(slice))
	copy(slice_new, slice)
	return append(slice_new[:index], slice_new[index+1:]...)
}

func is_safe(l []int, t int) bool {

	if t > 1 {
		return false
	}

	// init
	var pos int

	// take last entry cause second oen could be eliminated
	if l[len(l)-1]-l[0] > 0 {
		pos = 1
	} else {
		pos = -1
	}

	i := 1

	for i < len(l) {
		if (pos*(l[i]-l[i-1]) < 1) || (pos*(l[i]-l[i-1]) > 3) {

			//             for j, _ := range l {
			//                if is_safe(delete_at_index(l,j),t+1) {
			//                     return true
			//                }
			//             }
			//
			//             return false

			safe1 := is_safe(delete_at_index(l, i), t+1)
			safe2 := is_safe(delete_at_index(l, i-1), t+1)

			// 			firstSave := is_safe(delete_at_index(l, 0), t+1)
			// 			lastSave := is_safe(delete_at_index(l, len(l)-1), t+1)

			if safe1 || safe2 {
				return true
			} else {
				return false
			}

		}
		i += 1
	}

	return true

}
