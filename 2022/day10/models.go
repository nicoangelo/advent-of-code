package day10

import "strconv"

var OpMap = map[string]func(cpu *CPU, i *Instruction){
	"noop": func(cpu *CPU, i *Instruction) {
		cpu.Tick()
	},
	"addx": func(cpu *CPU, inst *Instruction) {
		cpu.Tick()
		v, _ := strconv.Atoi(inst.Args[0])
		cpu.RegisterX += v
		cpu.Tick()
	},
}

type CPU struct {
	Cycle              int
	RegisterX          int
	observer           *CpuObserver
	instructions       []*Instruction
	InstructionPointer *Instruction
	instructionIndex   int
}

func (cpu *CPU) Init(instructions []*Instruction, observer *CpuObserver) {
	cpu.Cycle = 1
	cpu.RegisterX = 1
	cpu.instructionIndex = -1
	cpu.observer = observer
	cpu.instructions = instructions
}

func (cpu *CPU) Tick() {
	cpu.Cycle += 1
	(*cpu.observer).OnTick(cpu.Cycle, cpu.RegisterX)
}

func (cpu *CPU) ExecuteNextInstruction() bool {
	cpu.instructionIndex++
	if cpu.instructionIndex < len(cpu.instructions) {
		cpu.InstructionPointer = cpu.instructions[cpu.instructionIndex]
		OpMap[cpu.InstructionPointer.Op](cpu, cpu.InstructionPointer)
		return true
	}
	return false
}

func (cpu *CPU) ExecuteAllInstructions() {
	for {
		if !cpu.ExecuteNextInstruction() {
			break
		}
	}
}

type Instruction struct {
	Op   string
	Args []string
}

type CpuObserver interface {
	OnTick(cycle int, registerX int)
}

type RegisterObserver struct {
	NextObserveCycle  int
	SignalStrengthSum int
}

func (ro *RegisterObserver) OnTick(cycle int, registerX int) {
	if cycle == ro.NextObserveCycle {
		ro.SignalStrengthSum += cycle * registerX
		ro.NextObserveCycle += 40
	}
}

type ScreenWriter struct {
	LineBuffer string
	Lines      []string
}

func (sw *ScreenWriter) OnTick(cycle int, registerX int) {
	if sw.Lines == nil {
		sw.Lines = make([]string, 0)
	}
	horizPos := (cycle % 40) - 1
	if horizPos >= registerX-1 && horizPos <= registerX+1 {
		sw.LineBuffer += "#"
	} else {
		sw.LineBuffer += " "
	}
	if horizPos == 0 {
		sw.Lines = append(sw.Lines, sw.LineBuffer)
		sw.LineBuffer = ""
	}
}
