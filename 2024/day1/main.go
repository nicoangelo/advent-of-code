package day1

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/intmath"
	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/sliceutils"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./day1/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {

	len := len(lines)

	var a = make([]int, len)
	var b = make([]int, len)

	for i, line := range lines {
		l := sliceutils.SliceConvert(strings.Split(line, "   "), strconv.Atoi)
		a[i] = l[0]
		b[i] = l[1]
	}

	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.IntSlice(b))

	d := 0

	for i := 0; i < len; i++ {
		d += intmath.Abs(a[i] - b[i])
	}

	return d
}

func part2(lines []string) int {

	a := map[int]int{}
	b := map[int]int{}

	for _, line := range lines {
		l := sliceutils.SliceConvert(strings.Split(line, "   "), strconv.Atoi)

		if v, ok := a[l[0]]; ok {
			a[l[0]] = v + 1
		} else {
			a[l[0]] = 1
		}

		if w, ok := b[l[1]]; ok {
			b[l[1]] = w + 1
		} else {
			b[l[1]] = 1
		}
	}

	s := 0

	for k, v := range a {
		s = s + k*v*b[k]
	}

	return s
}
