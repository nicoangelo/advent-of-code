package day17

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = Computer{
	Registers: map[string]int{
		"A": 729,
		"B": 0,
		"C": 0,
	},
	Instructions: []Instruction{
		{Opcode: 0, Operand: 1},
		{Opcode: 5, Operand: 4},
		{Opcode: 3, Operand: 0},
	},
}

func TestMain(m *testing.M) {
	loggerOpts := slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &loggerOpts))
	slog.SetDefault(logger)
	code := m.Run()
	os.Exit(code)
}

func TestPart1(t *testing.T) {
	res := part1(&testData)
	want := "4,6,3,5,6,3,5,2,1,0"
	assert.Equal(t, want, res)
}

func TestPart2(t *testing.T) {
	res := part2(&testData)
	want := 0
	assert.Equal(t, want, res)
}
