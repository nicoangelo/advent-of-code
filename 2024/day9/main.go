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
	fs := lines[0]
	checksum := 0
	back_i := len(fs) - 1
	frontPos := 0
	backRemaining, _ := strconv.Atoi(string(fs[back_i]))
	fileId := 0

out:
	for i := 0; i <= back_i; i++ {
		data, _ := strconv.Atoi(string(fs[i]))
		if i == back_i {
			data = backRemaining
		}
		if i%2 == 0 {
			fileId = i / 2
			for range data {
				checksum += frontPos * fileId
				frontPos++
			}
		} else if i%2 == 1 {
			for range data {
				if backRemaining == 0 {
					back_i -= 2
					backRemaining, _ = strconv.Atoi(string(fs[back_i]))
				}
				// edge case: the last hole is bigger than the amount of blocks to move from the back
				if back_i < i {
					break out
				}
				fileId = back_i / 2
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
