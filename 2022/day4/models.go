package day4

import "strconv"

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

func newCleaningPairFromTokens(tokens []string) *CleaningPairs {
	pair1Start, _ := strconv.Atoi(tokens[0])
	pair1End, _ := strconv.Atoi(tokens[1])
	pair2Start, _ := strconv.Atoi(tokens[2])
	pair2End, _ := strconv.Atoi(tokens[3])

	return &CleaningPairs{
		&Pair{
			pair1Start,
			pair1End,
		},
		&Pair{
			pair2Start,
			pair2End,
		},
	}
}
