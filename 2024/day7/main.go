package dayX

import (
	"log"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/sliceutils"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./dayX/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {

	numbers, results := readEquations(lines)

	for i := 0; i < len(results); i++ {
		nOps := len(numbers[i]) - 1

		permutations := [][]int{}

		for n := 0; n <= nOps; n++ {
			permutations = append(permutations, selectNX(n, nOps)...)
		}

	}

	return 0
}

func part2(lines []string) int {
	return 0
}

func readEquations(lines []string) (numbers [][]int, result []int) {
	numbers = [][]int{}
	result = []int{}

	for _, l := range lines {

		parts := strings.Split(l, " ")
		parts[0] = strings.TrimSuffix(parts[0], ":")

		u := sliceutils.SliceConvert(parts, strconv.Atoi)

		result = append(result, u[0])
		numbers = append(numbers, u[1:])

	}

	return numbers, result
}

func selectNX(n int, x int) [][]int {

	selected := [][]int{}

	if n == x {
		s := []int{}
		for i := 0; i < x; i++ {
			s = append(s, 1)
		}
		selected = append(selected, s)
	} else if n == 0 {
		s := []int{}
		for i := 0; i < x; i++ {
			s = append(s, 0)
		}
		selected = append(selected, s)
	} else {

		// opt1, add 0
		opt1 := selectNX(n, x-1)
		for _, v := range opt1 {
			selected = append(selected, append(v, 0))
		}

		// opt2, add 1
		opt2 := selectNX(n-1, x-1)
		for _, v := range opt2 {
			selected = append(selected, append(v, 1))
		}
	}

	return (selected)
}
