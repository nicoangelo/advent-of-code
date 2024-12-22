package day17

import (
	"fmt"
	"log"
	"os/user"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/sliceutils"
)

func PrintSolutions() {
	u, _ := user.Current()

	readRegister := func(s string, c *Computer) {
		tokens := strings.Split(s, ":")
		reg := tokens[0][len(tokens[0])-1:]
		val, _ := strconv.Atoi(strings.TrimSpace(tokens[1]))
		if c.Registers == nil {
			c.Registers = map[string]int{}
		}
		c.Registers[strings.TrimSuffix(reg, ":")] = val
	}
	readInstructions := func(s string, c *Computer) {
		s = strings.TrimPrefix(s, "Program: ")
		ins := strings.Split(s, ",")
		c.Instructions = make([]Instruction, len(ins)/2)
		for i := 0; i < len(ins); i += 2 {
			oc, _ := strconv.Atoi(ins[i])
			op, _ := strconv.Atoi(ins[i+1])
			c.Instructions[i/2] = Instruction{Opcode: oc, Operand: op}
		}
	}
	c := reader.ReadInputIntoStruct(
		"./day17/input_"+strings.ToLower(string(u.Username[0])),
		readRegister,
		readInstructions)
	part1 := part1(c)
	log.Println("Part 1: ", part1)

	c = reader.ReadInputIntoStruct(
		"./day17/input_"+strings.ToLower(string(u.Username[0])),
		readRegister,
		readInstructions)
	part2 := part2(c)
	log.Println("Part 2: ", part2)
}

func part1(c *Computer) string {
	c.ExecuteAllInstructions()
	str := sliceutils.SliceConvert(c.Output, func(in int) (string, error) { return fmt.Sprintf("%d", in), nil })
	return strings.Join(str, ",")
}

func part2(c *Computer) int {
	return 0
}
