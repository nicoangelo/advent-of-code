package almanac_test

import (
	"testing"

	"github.com/nicoangelo/advent-of-code-2023/day5/almanac"
	"github.com/stretchr/testify/assert"
)

func TestTranslateNumber1(t *testing.T) {
	e := &almanac.ConversionEntry{SourceStart: 1, SourceEnd: 10, DestinationStart: 20}

	assert.Equal(t, 24, e.TranslateNumber(5))
}

func TestTranslateNumberBigNumber(t *testing.T) {
	e := &almanac.ConversionEntry{SourceStart: 2321931404, SourceEnd: 2446354472, DestinationStart: 2067746708}

	assert.Equal(t, 2067746709, e.TranslateNumber(2321931405))
}
