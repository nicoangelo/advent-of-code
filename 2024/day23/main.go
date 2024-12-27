package day23

import (
	"fmt"
	"log"
	"os/user"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/sliceutils"
)

func PrintSolutions() {
	u, _ := user.Current()

	lines := reader.ReadInputFile("./day23/input_" + strings.ToLower(string(u.Username[0])))
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {
	hops := map[string][]string{}
	for _, l := range lines {
		tok := strings.Split(l, "-")
		sliceutils.SortAlphabeticallyIgnoreCase(tok)
		from := tok[0]
		to := tok[1]
		hops[from] = append(hops[from], to)
	}

	res := 0
	for k := range hops {
		res += findThreeHopsWithT(hops, []string{k}, strings.Contains(k, "t"))
	}

	return res
}

func findThreeHopsWithT(hops map[string][]string, stack []string, hasT bool) int {
	if len(stack) == 3 && hasT {
		fmt.Println(stack)
		return 1
	} else if len(stack) == 3 {
		return 0
	}
	nextHops := hops[stack[len(stack)-1]]

	res := 0
	for _, h := range nextHops {
		res += findThreeHopsWithT(hops, append(stack, h), hasT || strings.Contains(h, "t"))
	}
	return res
}

func part2(lines []string) int {
	return 0
}
