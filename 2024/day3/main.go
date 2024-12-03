package day3

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./day3/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) (res int) {
	rx, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)

	instructions := strings.Join(lines, "")
	m := rx.FindAllStringSubmatch(instructions, -1)
	for _, v := range m {
		v1, _ := strconv.Atoi(v[1])
		v2, _ := strconv.Atoi(v[2])
		res += v1 * v2
	}
	return res
}

func part2(lines []string) (res int) {
	rx, _ := regexp.Compile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	enabled := true
	instructions := strings.Join(lines, "")
	m := rx.FindAllStringSubmatch(instructions, -1)
	for _, v := range m {
		if v[0] == "do()" {
			enabled = true
		} else if v[0] == "don't()" {
			enabled = false
		} else if enabled {
			v1, _ := strconv.Atoi(v[1])
			v2, _ := strconv.Atoi(v[2])
			res += v1 * v2
		}
	}
	return res
}
