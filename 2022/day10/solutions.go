package day10

import (
	"fmt"
	"strings"
)

func InstructionFromLines(lines []string) (instructions []*Instruction) {
	instructions = make([]*Instruction, len(lines))
	for i, l := range lines {
		tokens := strings.Split(l, " ")
		instructions[i] = &Instruction{
			Op:   tokens[0],
			Args: tokens[1:],
		}
	}
	return instructions
}

func part1(lines []string) int {
	var obs CpuObserver
	obs = &RegisterObserver{NextObserveCycle: 20, SignalStrengthSum: 0}
	cpu := &CPU{}
	cpu.Init(InstructionFromLines(lines), &obs)
	cpu.ExecuteAllInstructions()

	return obs.(*RegisterObserver).SignalStrengthSum
}

func part2(lines []string) int {
	var obs CpuObserver
	obs = &ScreenWriter{}
	cpu := &CPU{}
	cpu.Init(InstructionFromLines(lines), &obs)
	cpu.ExecuteAllInstructions()

	for _, l := range obs.(*ScreenWriter).Lines {
		fmt.Println(l)
	}

	return 0
}
