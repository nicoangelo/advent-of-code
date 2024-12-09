package day8

import (
	"log"
	"os/user"
	"strings"
	"unicode"

	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	u, _ := user.Current()

	lines := reader.ReadInputFile("./day8/input_" + strings.ToLower(string(u.Username[0])))
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {

	antennas := getAntennas(lines)

	width := len(lines[0])
	height := len(lines)

	return (getAntinodes(antennas, width, height))

}

func part2(lines []string) int {
	return 0
}

func getAntennas(lines []string) map[byte][][]int {

	antennas := map[byte][][]int{}

	for r, line := range lines {
		for c := 0; c < len(line); c++ {
			l := rune(line[c])

			if unicode.IsLetter(l) || unicode.IsDigit(l) {

				temp, okay := antennas[line[c]]

				if !okay {
					antennas[line[c]] = append([][]int{}, []int{r, c})
				} else {
					antennas[line[c]] = append(temp, []int{r, c})
				}

			}
		}
	}

	return (antennas)

}

func getAntinodes(nodesMap map[byte][][]int, width int, height int) int {

	antinodes := map[node]bool{}
	count := 0

	for _, nodes := range nodesMap {
		for i := 0; i < len(nodes); i++ {
			for j := i + 1; j < len(nodes); j++ {

				dr := nodes[i][0] - nodes[j][0]
				dc := nodes[i][1] - nodes[j][1]

				a1r := nodes[i][0] + dr
				a1c := nodes[i][1] + dc

				a1 := node{a1r, a1c}

				a2r := nodes[j][0] - dr
				a2c := nodes[j][1] - dc

				a2 := node{a2r, a2c}

				if (0 <= a1r) && (a1r < height) &&
					(0 <= a1c) && (a1c < width) {

					_, okay := antinodes[a1]

					if !okay {
						antinodes[a1] = true
						count += 1
					}

				}

				if (0 <= a2r) && (a2r < height) &&
					(0 <= a2c) && (a2c < width) {

					_, okay := antinodes[a2]

					if !okay {
						antinodes[a2] = true
						count += 1
					}
				}

			}
		}
	}

	return count

}

type node struct {
	x, y int
}
