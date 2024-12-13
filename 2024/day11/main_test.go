package day11

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []string{
	"125 17",
}

func TestMain(m *testing.M) {
	loggerOpts := slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &loggerOpts))
	slog.SetDefault(logger)
	code := m.Run()
	os.Exit(code)
}

func TestPart1Blink6(t *testing.T) {
	res := part1(testData, 6)
	want := 22
	assert.Equal(t, want, res)
}

func TestPart1Blink25(t *testing.T) {
	res := part1(testData, 25)
	want := 55312
	assert.Equal(t, want, res)
}
