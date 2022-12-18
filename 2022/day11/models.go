package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/nicoangelo/advent-of-code-2022/shared"
)

type Monkey struct {
	Items                 []int
	InspectionCounter     int
	Operation             *Operation
	DivisionTest          int
	TrueThrowMonkeyIndex  int
	FalseThrowMonkeyIndex int
}

func (m *Monkey) CatchItem(item int) {
	m.Items = append(m.Items, item)
}

type Operation struct {
	Operator rune
	Right    string
}

type KeepAwayGame struct {
	Monkeys           map[int]*Monkey
	CurrentRound      int
	WorryLevelDivider int
	BigNumberModulo   int
}

const (
	PREFIX_MONKEY     = "Monkey"
	PREFIX_ITEMS      = "  Starting items: "
	PREFIX_OPERATION  = "  Operation: new = old "
	PREFIX_TEST       = "  Test: divisible by "
	PREFIX_TEST_TRUE  = "    If true: throw to monkey "
	PREFIX_TEST_FALSE = "    If false: throw to monkey "
)

func (g *KeepAwayGame) MonkeysFromLines(lines []string) {
	var currentMonkey *Monkey
	g.Monkeys = map[int]*Monkey{}

	for _, l := range lines {
		switch true {
		case strings.HasPrefix(l, PREFIX_MONKEY):
			currentMonkey = &Monkey{}
			var index int
			fmt.Sscanf(l, PREFIX_MONKEY+" %d:", &index)
			g.Monkeys[index] = currentMonkey
		case strings.HasPrefix(l, PREFIX_ITEMS):
			items := strings.TrimPrefix(l, PREFIX_ITEMS)
			currentMonkey.Items = shared.SliceConvert(strings.Split(items, ", "), strconv.Atoi)
		case strings.HasPrefix(l, PREFIX_OPERATION):
			op_tokens := [2]string{}
			_, err := fmt.Sscanf(strings.TrimPrefix(l, PREFIX_OPERATION), "%s %s", &op_tokens[0], &op_tokens[1])
			if err == nil {
				currentMonkey.Operation = &Operation{rune(op_tokens[0][0]), op_tokens[1]}
			}
		case strings.HasPrefix(l, PREFIX_TEST):
			currentMonkey.DivisionTest, _ = strconv.Atoi(strings.TrimPrefix(l, PREFIX_TEST))
		case strings.HasPrefix(l, PREFIX_TEST_TRUE):
			currentMonkey.TrueThrowMonkeyIndex, _ = strconv.Atoi(strings.TrimPrefix(l, PREFIX_TEST_TRUE))
		case strings.HasPrefix(l, PREFIX_TEST_FALSE):
			currentMonkey.FalseThrowMonkeyIndex, _ = strconv.Atoi(strings.TrimPrefix(l, PREFIX_TEST_FALSE))
		case l == "":
			continue
		}
	}
}

func (g *KeepAwayGame) PlayOneRound() {
	g.CurrentRound++
	for i := 0; i < len(g.Monkeys); i++ {
		m := g.Monkeys[i]
		logger.Println("Turn Monkey #", i)
		if len(m.Items) == 0 {
			logger.Println("Monkey has no items. Moving on.")
			continue
		}

		for _, worryLevel := range m.Items {
			logger.Println("Playing with item", worryLevel)
			newLevel := m.Operation.CalculateNewValue(worryLevel) / int(g.WorryLevelDivider)
			if g.BigNumberModulo > 0 {
				newLevel %= int(g.BigNumberModulo)
			}
			var throwTarget int
			if newLevel%int(m.DivisionTest) == 0 {
				throwTarget = m.TrueThrowMonkeyIndex
			} else {
				throwTarget = m.FalseThrowMonkeyIndex
			}
			logger.Println("Throwing item", newLevel, "to monkey", throwTarget)
			g.Monkeys[throwTarget].CatchItem(newLevel)
		}
		m.InspectionCounter += len(m.Items)
		m.Items = []int{}
	}
}

func (g *KeepAwayGame) GetMostActiveMonkeys(topN int) []int {
	if topN == -1 {
		topN = len(g.Monkeys)
	}
	monkeyMoves := make([]int, len(g.Monkeys))
	for i, m := range g.Monkeys {
		monkeyMoves[i] = m.InspectionCounter
	}
	sort.Sort(sort.Reverse(sort.IntSlice(monkeyMoves)))
	return monkeyMoves[0:topN]
}

func (g *KeepAwayGame) PrintMonkeyStats() {
	s := g.GetMostActiveMonkeys(-1)
	fmt.Printf("== After round %d ==\n", g.CurrentRound)
	for i, v := range s {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, v)
	}
}

func (g *KeepAwayGame) PrintMonkeyStashes() {
	fmt.Printf("== After round %d ==\n", g.CurrentRound)
	for i, v := range g.Monkeys {
		fmt.Printf("Monkey %d: %v\n", i, v.Items)
	}
}

func (o *Operation) CalculateNewValue(old int) int {
	rightValue := int(0)
	if o.Right == "old" {
		rightValue = old
	} else {

		rightValue, _ = strconv.Atoi(o.Right)
	}
	switch o.Operator {
	case '+':
		return old + rightValue
	case '-':
		return old - rightValue
	case '*':
		return old * rightValue
	}
	return 0
}
