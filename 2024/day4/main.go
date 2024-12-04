package day4

import (
	"fmt"
	"log"
	"regexp"

	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./day4/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {

	count := 0

	height := len(lines)
	width := len(lines[0])

	fo, _ := regexp.Compile(`XMAS`)
	re, _ := regexp.Compile(`SAMX`)

	columns := make([]string, width)
	diags1 := make([]string, width+height-1)
	diags2 := make([]string, width+height-1)

	// var diag1 []string
	// var diag2 []string

	for r, row := range lines {

		count += len(fo.FindAllStringSubmatch(row, -1))
		count += len(re.FindAllStringSubmatch(row, -1))

		for c := 0; c < width; c++ {
			columns[c] = columns[c] + string(row[c])
			diags1[r+c] = diags1[r+c] + string(row[c])
			diags2[r+(width-c-1)] = diags2[r+(width-c-1)] + string(row[c])

		}

	}

	for _, col := range columns {
		count += len(fo.FindAllStringSubmatch(col, -1))
		count += len(re.FindAllStringSubmatch(col, -1))
	}

	for _, diag1 := range diags1 {
		count += len(fo.FindAllStringSubmatch(diag1, -1))
		count += len(re.FindAllStringSubmatch(diag1, -1))
	}

	for _, diag2 := range diags2 {
		count += len(fo.FindAllStringSubmatch(diag2, -1))
		count += len(re.FindAllStringSubmatch(diag2, -1))
	}

	return count
}

func part2(lines []string) int {

	height := len(lines)
	width := len(lines[0])

	fo, _ := regexp.Compile(`MAS`)
	re, _ := regexp.Compile(`SAM`)

	diags1 := make([]string, width+height-1)
	diags2 := make([]string, width+height-1)

	// var diag1 []string
	// var diag2 []string

	for r, row := range lines {

		for c := 0; c < width; c++ {
			diags1[r+c] = diags1[r+c] + string(row[c])
			diags2[r+(width-c-1)] = diags2[r+(width-c-1)] + string(row[c])

		}

	}

	bws := make([][]int, 0)

	for d, diag1 := range diags1 {

		diag1f := fo.FindAllStringIndex(diag1, -1)
		diag1r := re.FindAllStringIndex(diag1, -1)

		for _, v := range append(diag1f, diag1r...) {

			i := v[0]
			if d < width {
				line := []int{i, d - i}
				bws = append(bws, line)
			} else {
				line := []int{i + d - width, d - i}
				bws = append(bws, line)
			}

		}
	}

	fws := make([][]int, 0)

	for d, diag2 := range diags2 {

		diag2f := fo.FindAllStringIndex(diag2, -1)
		diag2r := re.FindAllStringIndex(diag2, -1)

		for _, v := range append(diag2f, diag2r...) {

			i := v[0]
			if d < width {
				line := []int{i, width - (d - i) - 1}
				fws = append(fws, line)
			} else {
				line := []int{i + d - width, width - (d - i) - 1}
				fws = append(fws, line)
			}

		}
	}

	fmt.Print(bws)
	fmt.Print(fws)

	return 0
}
