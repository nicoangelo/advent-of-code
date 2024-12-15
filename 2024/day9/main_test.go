package day9

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []string{
	"2333133121414131402",
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
	want := 1928
	assert.Equal(t, want, res)
}

func TestPart2(t *testing.T) {
	res := part2(testData)
	want := 2858
	assert.Equal(t, want, res)
}
