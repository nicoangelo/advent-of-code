package day9

import (
	"log"
	"os/user"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	u, _ := user.Current()

	lines := reader.ReadInputFile("./day9/input_" + strings.ToLower(string(u.Username[0])))
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {
	data := lines[0]
	checksum := 0
	backPos := len(data) - 1
	frontPos := 0
	backRemaining, _ := strconv.Atoi(string(data[len(data)-1]))
	fileId := 0

	for i := 0; i <= backPos; i++ {
		currData, _ := strconv.Atoi(string(data[i]))
		if i == backPos {
			currData = backRemaining
		}
		if i%2 == 0 {
			fileId = i / 2
			for j := 0; j < currData; j++ {
				checksum += frontPos * fileId
				frontPos++
			}
		} else if i%2 == 1 {
			for k := 0; k < currData; k++ {
				if backRemaining == 0 {
					backPos -= 2
					backRemaining, _ = strconv.Atoi(string(data[backPos]))
				}
				fileId = backPos / 2
				checksum += frontPos * fileId
				backRemaining--
				frontPos++
			}
		}
	}
	return checksum
}

func part2(lines []string) int {
	return 0
}
