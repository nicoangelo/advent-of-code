package day11

import (
	"log"
	"math"
	"os/user"
	"slices"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/sliceutils"
)

func PrintSolutions() {
	u, _ := user.Current()

	lines := reader.ReadInputFile("./day11/input_" + strings.ToLower(string(u.Username[0])))
	res1 := part1(lines, 25)
	log.Println("Part 1: ", res1)

	res2 := part1(lines, 75)
	log.Println("Part 2: ", res2)
}

func part1(lines []string, blinks int) int {
	line := lines[0]
	stones := sliceutils.SliceConvert(strings.Split(line, " "), strconv.Atoi)
	for range blinks {
		for i := 0; i < len(stones); i++ {
			if stones[i] == 0 {
				stones[i] = 1
			} else if isEven, len := digitLenIsEven(stones[i]); isEven {
				newDigits := splitDigits(stones[i], len)
				stones = slices.Replace(stones, i, i+1, newDigits...)
				i++
			} else {
				stones[i] = stones[i] * 2024
			}
		}
	}
	return len(stones)
}

func digitLenIsEven(n int) (even bool, len int) {
	if n == 0 {
		return false, 1
	}
	for n != 0 {
		n /= 10
		len++
	}
	return len%2 == 0, len
}

func splitDigits(n int, len int) []int {
	splitter := int(math.Pow10(len / 2))
	return []int{n / splitter, n % splitter}
}
