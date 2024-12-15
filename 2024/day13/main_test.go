package day13

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []string{
	"Button A: X+94, Y+34",
	"Button B: X+22, Y+67",
	"Prize: X=8400, Y=5400",
	"",
	"Button A: X+26, Y+66",
	"Button B: X+67, Y+21",
	"Prize: X=12748, Y=12176",
	"",
	"Button A: X+17, Y+86",
	"Button B: X+84, Y+37",
	"Prize: X=7870, Y=6450",
	"",
	"Button A: X+69, Y+23",
	"Button B: X+27, Y+71",
	"Prize: X=18641, Y=10279",
}

func TestMain(m *testing.M) {
	loggerOpts := slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &loggerOpts))
	slog.SetDefault(logger)
	code := m.Run()
	os.Exit(code)
}

func TestPart1(t *testing.T) {
	res := part1(testData)
	want := 480
	assert.Equal(t, want, res)
}
