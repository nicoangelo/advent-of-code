package day17

import (
	"github.com/nicoangelo/aoc-pkg/intmath"
)

type Computer struct {
	Registers    map[string]int
	Instructions []Instruction
	Output       []int
}

// opcodes
const (
	ADV = 0
	BXL = 1
	BST = 2
	JNZ = 3
	BXC = 4
	OUT = 5
	BDV = 6
	CDV = 7
)

// combo operands
const (
	Literal0  = 0
	Literal1  = 1
	Literal2  = 2
	Literal3  = 3
	RegisterA = 4
	RegisterB = 5
	RegisterC = 6
	Reserved  = 7
)

type Instruction struct {
	Opcode  int
	Operand int
}

func (c *Computer) getComboOperandValue(operand int) int {
	if operand <= Literal3 {
		return operand
	}
	if operand == RegisterA {
		return c.Registers["A"]
	}
	if operand == RegisterB {
		return c.Registers["B"]
	}
	if operand == RegisterC {
		return c.Registers["C"]
	}
	return 0
}

func (c *Computer) ExecuteAllInstructions() {
	for i := 0; i < len(c.Instructions); {
		jmp := c.Execute(c.Instructions[i])
		if jmp >= 0 {
			i = jmp // if this would jump beyond boundaries, the for condition "halts"
		} else {
			i += 1
		}
	}
}
func (c *Computer) Execute(i Instruction) (jump int) {
	cmb := c.getComboOperandValue(i.Operand)
	switch i.Opcode {
	case ADV:
		nom := c.Registers["A"]
		denom := intmath.Pow(2, cmb)
		// log.Println("adv", nom, denom)
		c.Registers["A"] = nom / denom
	case BXL:
		// log.Println("bxl", c.Registers["B"], i.Operand)
		c.Registers["B"] = c.Registers["B"] ^ i.Operand
	case BST:
		// log.Println("bst", cmb, "%", 8)
		c.Registers["B"] = cmb % 8
	case JNZ:
		if c.Registers["A"] != 0 {
			// log.Println("jnz", i.Operand)
			return i.Operand
		}
		// log.Println("jnz", "noop")
	case BXC:
		// log.Println("bxc", c.Registers["B"], c.Registers["C"])
		c.Registers["B"] = c.Registers["B"] ^ c.Registers["C"]
	case OUT:
		// log.Println("out", cmb, "%", 8)
		c.Output = append(c.Output, cmb%8)
	case BDV:
		nom := c.Registers["A"]
		denom := intmath.Pow(2, cmb)
		// log.Println("bdv", nom, denom)
		c.Registers["B"] = nom / denom
	case CDV:
		nom := c.Registers["A"]
		denom := intmath.Pow(2, cmb)
		// log.Println("cdv", nom, denom)
		c.Registers["C"] = nom / denom
	}
	return -1
}
