package day13

import (
	"log"
	"os/user"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	u, _ := user.Current()

	lines := reader.ReadInputFile("./day13/input_" + strings.ToLower(string(u.Username[0])))
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {
	tokensAmount := 0
	for i := 0; i < len(lines); i += 4 {
		g := Game{}
		g.FromLines(lines[i:i+3], 0)
		g.Solve()
		tokensAmount += g.RequiredTokensWithLimit(100)
	}
	return tokensAmount
}

func part2(lines []string) int {
	tokensAmount := 0
	for i := 0; i < len(lines); i += 4 {
		g := Game{}
		g.FromLines(lines[i:i+3], 10000000000000)
		g.Solve()
		tokensAmount += g.RequiredTokensWithLimit(0)
	}
	return tokensAmount
}
