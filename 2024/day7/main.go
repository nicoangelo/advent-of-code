package day7

import (
	"fmt"
	"log"
	"math"
	"os/user"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {

	u, _ := user.Current()

	lines := reader.ReadInputFile("./day7/input_" + strings.ToLower(string(u.Username[0])))

	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {

	total := uint64(0)

	numbers, results := readEquations(lines)

	for i := 0; i < len(results); i++ {
		nOps := len(numbers[i]) - 1

		permutations := [][]int{}

		for n := 0; n <= nOps; n++ {
			permutations = append(permutations, selectNX(n, nOps)...)
		}

		total += testOperators(results[i], numbers[i], permutations)

	}

	return int(total)
}

func part2(lines []string) int {

	total := uint64(0)

	// for i := 4; i < 5; i++ {
	// 	fmt.Println(permuteNX(i))
	// }

	numbers, results := readEquations(lines)

	permutations := map[int][][]int{}

	for i := 0; i < len(results); i++ {

		// log.Println(i, numbers[i], results[i], total)

		nOps := len(numbers[i]) - 1

		perm, okay := permutations[nOps]

		if !okay {
			perm = permuteNX(nOps)
			permutations[nOps] = perm
		}

		// log.Println("permutations calculated", nOps, len(perm))
		temp := total
		total += testOperators2(results[i], numbers[i], perm)

		if total < temp {
			fmt.Println("Overflow")
		}

	}

	return int(total)
}

func readEquations(lines []string) (numbers [][]uint64, result []uint64) {

	numbers = [][]uint64{}
	result = []uint64{}

	for _, l := range lines {

		parts := strings.Split(l, " ")
		parts[0] = strings.TrimSuffix(parts[0], ":")

		v, _ := strconv.ParseUint(parts[0], 10, 64)
		result = append(result, v)

		nums := []uint64{}

		for i := 1; i < len(parts); i++ {
			w, _ := strconv.ParseUint(parts[i], 10, 64)
			nums = append(nums, w)
		}

		numbers = append(numbers, nums)

	}

	return numbers, result
}

func testOperators(result uint64, numbers []uint64, permutations [][]int) uint64 {

	for _, v := range permutations {

		out := numbers[0]

		i := 0

		for i < len(v) {
			if v[i] == 1 {
				out = out * numbers[i+1]
			} else {
				out += numbers[i+1]
			}

			i += 1
		}

		if out == result {
			return (result)
		}

	}

	return (0)

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

///////////////////////// part2

func permuteNX(x int) [][]int {

	s := [][]int{}

	if x == 1 {
		s = append(s, []int{2})
		s = append(s, []int{0})
		s = append(s, []int{1})
	} else {
		part := permuteNX(x - 1)

		for _, v := range part {

			for i := 0; i < 3; i++ {

				// This is too stupid to be true
				cpy := make([]int, len(v))
				copy(cpy, v)
				temp := append(cpy, i)

				s = append(s, temp)
			}
			// fmt.Println("s: ", s)

		}
	}
	return (s)
}

func testOperators2(result uint64, numbers []uint64, permutations [][]int) uint64 {

	j := 0
	for _, v := range permutations {

		out := numbers[0]

		for i := 0; i < len(v); i++ {
			temp := out

			// the order should not change results since all permutations are testes
			// however the order/mapping of operators to 0/1/2 does change the results!!!

			if v[i] == 0 {
				// append
				new := float64(numbers[i+1])

				for new >= 1 {
					new = math.Floor(new / 10)
					out = out * 10
				}

				out = out + numbers[i+1]

			} else if v[i] == 1 {

				//multiply
				out = out * numbers[i+1]

			} else if v[i] == 2 {

				//add
				out = out + numbers[i+1]

			}

			if out < temp {
				fmt.Println("Overflow")
			}

			// skipping early if result gets to high
			if out > result {
				continue
			}

		}

		if out == result {
			return (result)
		}

		j += 1

	}

	if float64(j) != math.Pow(3, float64(len(numbers)-1)) {
		fmt.Println(result, numbers, math.Pow(3, float64(len(numbers)-1)), j)
	}

	return (0)

}
