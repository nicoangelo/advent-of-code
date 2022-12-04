package day4

import (
	"fmt"
)

type CleaningPairs struct {
	Pair1 *Pair
	Pair2 *Pair
}

func (cp *CleaningPairs) IsAnyFullyContained() bool {
	fullyContains := func(p1 *Pair, p2 *Pair) bool {
		return p1.Start <= p2.Start && p1.End >= p2.End
	}
	return fullyContains(cp.Pair1, cp.Pair2) || fullyContains(cp.Pair2, cp.Pair1)
}

func (cp *CleaningPairs) IsAnyOverlapping() bool {
	return cp.Pair1.Start <= cp.Pair2.End && cp.Pair2.Start <= cp.Pair1.End
}

type Pair struct {
	Start int
	End   int
}

func newCleaningPairFromTokens(line string) *CleaningPairs {
	p1 := &Pair{}
	p2 := &Pair{}
	fmt.Sscanf(line, "%d-%d,%d-%d", &p1.Start, &p1.End, &p2.Start, &p2.End)
	return &CleaningPairs{p1, p2}
}
