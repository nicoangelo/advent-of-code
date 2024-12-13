package day11

import (
	"log"
	"math"
	"os/user"
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

	next := map[int]int{}
	for i := 0; i < len(stones); i++ {
		next[stones[i]]++
	}
	totalFreqs := 0
	for b := range blinks {
		numFreqs := make(map[int]int, len(next))
		for k, v := range next {
			numFreqs[k] = v
		}
		next = map[int]int{}
		totalFreqs = 0
		for n, f := range numFreqs {
			if n == 0 {
				next[1] += f
				totalFreqs += f
			} else if isEven, len := digitLenIsEven(n); isEven {
				n1, n2 := splitDigits(n, len)
				next[n1] += f
				totalFreqs += f
				next[n2] += f
				totalFreqs += f
			} else {
				next[n*2024] += f
				totalFreqs += f
			}
		}
		log.Println("Blinked", b, "times:", totalFreqs, "stones")
	}
	return totalFreqs
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

func splitDigits(n int, len int) (int, int) {
	splitter := int(math.Pow10(len / 2))
	return n / splitter, n % splitter
}
