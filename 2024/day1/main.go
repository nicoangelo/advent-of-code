package day1

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/math"
	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/slices"
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
		l := slices.SliceConvert(strings.Split(line, "   "), strconv.Atoi)
		a[i] = l[0]
		b[i] = l[1]
	}

	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.IntSlice(b))

	d := 0

	for i := 0; i < len; i++ {
		d += math.AbsInt(a[i] - b[i])
	}

	return d
}

func part2(lines []string) int {
	return 0
}
