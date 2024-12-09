package day7

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []string{
	"190: 10 19",
	"3267: 81 40 27",
	"83: 17 5",
	"156: 15 6",
	"7290: 6 8 6 15",
	"161011: 16 10 13",
	"192: 17 8 14",
	"21037: 9 7 18 13",
	"292: 11 6 16 20",
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
	want := 3749
	assert.Equal(t, want, res)
}

func TestPart2(t *testing.T) {
	res := part2(testData)
	want := 0
	assert.Equal(t, want, res)
}
