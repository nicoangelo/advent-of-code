package day5

import (
	"fmt"
	"strings"
)

type CargoOperations struct {
	CrateStacks *[]CrateStack
	Operations  *[]CraneOperation
}

type CrateStack struct {
	Crates []rune
}

func (co *CargoOperations) GetTopCrates() (res []rune) {
	res = make([]rune, 0)
	for _, v := range *co.CrateStacks {
		res = append(res, v.Crates[len(v.Crates)-1])
	}
	return res
}

func (co *CargoOperations) FillFromLines(lines *[]string) {
	section := Crates
	stacks := make([]CrateStack, 0)
	ops := make([]CraneOperation, 0)
	for i, v := range *lines {
		if strings.TrimSpace(v) == "" || v[1] == '1' {
			section = Operations
			continue
		}
		switch section {
		case Crates:
			stack := 0
			for j := 1; j < len(v); j += 4 {
				crateRune := rune(v[j])
				if i == 0 {
					stacks = append(stacks, *new(CrateStack))
				}
				if crateRune != ' ' {
					stacks[stack].Crates = append([]rune{crateRune}, stacks[stack].Crates...)
				}
				stack++
			}
		case Operations:
			op := &CraneOperation{}
			tokens := make([]int, 2)
			fmt.Sscanf(v, "move %d from %d to %d", &op.Amount, &tokens[0], &tokens[1])
			op.Source = &stacks[tokens[0]-1]
			op.Destination = &stacks[tokens[1]-1]
			ops = append(ops, *op)
		}
	}
	co.CrateStacks = &stacks
	co.Operations = &ops
}

type CraneOperation struct {
	Amount      int
	Source      *CrateStack
	Destination *CrateStack
}

type ReadSection int

const (
	Crates ReadSection = iota + 1
	Operations
)
