package day23

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []string{
	"kh-tc",
	"qp-kh",
	"de-cg",
	"ka-co",
	"yn-aq",
	"qp-ub",
	"cg-tb",
	"vc-aq",
	"tb-ka",
	"wh-tc",
	"yn-cg",
	"kh-ub",
	"ta-co",
	"de-co",
	"tc-td",
	"tb-wq",
	"wh-td",
	"ta-ka",
	"td-qp",
	"aq-cg",
	"wq-ub",
	"ub-vc",
	"de-ta",
	"wq-aq",
	"wq-vc",
	"wh-yn",
	"ka-de",
	"kh-ta",
	"co-tc",
	"wh-qp",
	"tb-vc",
	"td-yn",
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
	want := 7
	assert.Equal(t, want, res)
}

func TestPart2(t *testing.T) {
	res := part2(testData)
	want := 0
	assert.Equal(t, want, res)
}
