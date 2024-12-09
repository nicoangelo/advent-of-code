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
	antennas := getAntennas(lines)

	width := len(lines[0])
	height := len(lines)

	return (getAntinodes2(antennas, width, height))
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

func getAntinodes2(nodesMap map[byte][][]int, width int, height int) int {

	antinodes := map[node]bool{}
	count := 0

	for _, nodes := range nodesMap {
		for i := 0; i < len(nodes); i++ {
			for j := i + 1; j < len(nodes); j++ {

				dist := node{nodes[i][0] - nodes[j][0], nodes[i][1] - nodes[j][1]}

				minDist := getMinDist(dist)

				an := node{nodes[i][0], nodes[i][1]}
				tempAn := an

				for (0 <= tempAn.x) && (tempAn.x < height) &&
					(0 <= tempAn.y) && (tempAn.y < width) {

					_, okay := antinodes[tempAn]

					if !okay {
						antinodes[tempAn] = true
						count += 1
					}

					tempAn = substractDist(tempAn, minDist)
				}

				// we already looked at an
				tempAn = addDist(an, minDist)

				for (0 <= tempAn.x) && (tempAn.x < height) &&
					(0 <= tempAn.y) && (tempAn.y < width) {

					_, okay := antinodes[tempAn]

					if !okay {
						antinodes[tempAn] = true
						count += 1
					}

					tempAn = addDist(tempAn, minDist)
				}

			}
		}
	}

	return count

}

type node struct {
	x, y int
}

func addDist(oldNode node, dist node) node {
	newNode := node{oldNode.x + dist.x, oldNode.y + dist.y}
	return (newNode)
}

func substractDist(oldNode node, dist node) node {
	newNode := node{oldNode.x - dist.x, oldNode.y - dist.y}
	return (newNode)
}

func getMinDist(dist node) node {
	div := GCD(dist.x, dist.y)
	return (node{dist.x / div, dist.y / div})
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
