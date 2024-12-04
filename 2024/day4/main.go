package day4

import (
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

	mas := [3]byte{'M', 'A', 'S'}
	sam := [3]byte{'S', 'A', 'M'}

	count := 0

	width := len(lines[0])

	for r := 0; r < (len(lines) - 2); r++ {

		for c := 0; c < (width - 2); c++ {
			fw := getFwDiag(lines, r, c)
			if fw == mas || fw == sam {
				bw := getBwDiag(lines, r, c+2)
				if bw == mas || bw == sam {
					count += 1
				}
			}
		}

	}

	return count
}

func getFwDiag(lines []string, r int, c int) [3]byte {

	var fw [3]byte

	for i := 0; i < 3; i++ {
		fw[i] = lines[r+i][c+i]
	}

	return (fw)

}

func getBwDiag(lines []string, r int, c int) [3]byte {

	var fw [3]byte

	for i := 0; i < 3; i++ {
		fw[i] = lines[r+i][c-i]
	}

	return (fw)

}
