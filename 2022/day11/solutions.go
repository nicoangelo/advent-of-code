package day11

import (
	"github.com/nicoangelo/aoc-pkg/slicemath"
)

func part1(lines []string) int {
	g := &KeepAwayGame{WorryLevelDivider: 3}
	g.MonkeysFromLines(lines)
	for i := 0; i < 20; i++ {
		g.PlayOneRound()
		// g.PrintMonkeyStashes()
	}
	/* Something is still off, because the wrong
	 * monkeys have the correct counts.
	 */
	// g.PrintMonkeyStats()
	return slicemath.Multiply(g.GetMostActiveMonkeys(2))
}

func part2(lines []string) int {
	g := &KeepAwayGame{WorryLevelDivider: 1}
	g.MonkeysFromLines(lines)
	g.BigNumberModulo = 1
	for _, m := range g.Monkeys {
		g.BigNumberModulo *= m.DivisionTest
	}
	for i := 0; i < 10000; i++ {
		g.PlayOneRound()
		// if (g.CurrentRound)%1000 == 0 || g.CurrentRound == 1 {
		// 	g.PrintMonkeyStats()
		// }
	}
	return slicemath.Multiply(g.GetMostActiveMonkeys(2))
}
