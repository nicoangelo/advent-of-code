package day19

type Towels struct {
	availablePatterns map[string]int
	maxPatternLength  int
	desiredPatterns   []string
}

var knownCombinations = map[string]int{}

func (t *Towels) AddAvailablePattern(p string) {
	if t.availablePatterns == nil {
		t.availablePatterns = map[string]int{}
	}
	if len(p) > t.maxPatternLength {
		t.maxPatternLength = len(p)
	}
	t.availablePatterns[p] = len(p)
}

func (t *Towels) AddDesiredPattern(p string) {
	t.desiredPatterns = append(t.desiredPatterns, p)
}

func (t *Towels) CountFittablePatterns() int {
	res := 0
	for _, p := range t.desiredPatterns {
		fits := t.countFits(p)
		if fits > 0 {
			res++
		}
	}
	return res
}

func (t *Towels) CountAllFittablePatternCombinations() int {
	res := 0
	for _, p := range t.desiredPatterns {
		res += t.countFits(p)
	}
	return res
}

func (t *Towels) countFits(remaining string) int {
	known := knownCombinations[remaining]
	if known > 0 {
		return known
	}
	res := 0
	for p, pLen := range t.availablePatterns {
		if len(remaining) < pLen || remaining[:pLen] != p {
			continue
		}
		if len(remaining) == pLen && remaining == p {
			knownCombinations[remaining]++
			res++
		} else {
			cnts := t.countFits(remaining[pLen:])
			knownCombinations[remaining[pLen:]] = cnts
			res += cnts
		}
	}
	return res
}
