package day8

import (
	"log"

	"github.com/nicoangelo/advent-of-code-2023/day8/nav"
	"github.com/nicoangelo/aoc-pkg/reader"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./day8/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {
	m := &nav.Map{}
	m.FromLines(lines)
	ctx := m.NewNavigationContext("AAA", func(n *nav.Node) bool { return n.Name == "ZZZ" })

	return ctx.Navigate()
}

func part2(lines []string) int {
	m := &nav.Map{}
	m.FromLines(lines)
	ctxs := []*nav.NavigationContext{}
	for _, v := range m.GetNodeKeys() {
		if v[len(v)-1:] != "A" {
			continue
		}
		c := m.NewNavigationContext(v, nil)
		c.StartNavigation()
		ctxs = append(ctxs, c)
	}

	for {
		for i := range ctxs {
			ctxs[i].NavigateStep()
		}
		finished := 0
		for _, c := range ctxs {
			nn := c.GetCurrentNode().Name
			if nn[len(nn)-1:] == "Z" {
				finished++
			}
		}
		if finished == len(ctxs) {
			break
		}
	}

	return ctxs[0].GetCurrentStepCount()
}
