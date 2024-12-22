package day9

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []string{
	"2333133121414131402",
	"673253833464635054191677274350925861527651788483",
	"23222120202525282820202020272722212121",
	"22222228282828222222282829212324252627282920",
}

func TestMain(m *testing.M) {
	loggerOpts := slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &loggerOpts))
	slog.SetDefault(logger)
	code := m.Run()
	os.Exit(code)
}

func TestPart1(t *testing.T) {
	res := part1(testData[0])
	want := 1928
	assert.Equal(t, want, res)
}

func TestPart2Fs1(t *testing.T) {
	res := part2(testData[0])
	want := 2858
	assert.Equal(t, want, res)
}

func TestPart2Fs2(t *testing.T) {
	res := part2(testData[1])
	want := 149706
	assert.Equal(t, want, res)
}

func TestPart2Fs3(t *testing.T) {
	res := part2(testData[2])
	want := 7705
	assert.Equal(t, want, res)
}

func TestPart2Fs4(t *testing.T) {
	res := part2(testData[3])
	want := 9447
	assert.Equal(t, want, res)
}
