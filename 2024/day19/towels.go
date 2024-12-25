package day19

type Towels struct {
	availablePatterns map[string]int
	maxPatternLength  int
	desiredPatterns   []string
}

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

func (t *Towels) TryFitPatterns() int {
	res := 0
	for _, p := range t.desiredPatterns {
		if t.tryFit(p) {
			res++
		}
	}
	return res
}

func (t *Towels) tryFit(remaining string) bool {
	for p, pLen := range t.availablePatterns {
		if len(remaining) < pLen || remaining[:pLen] != p {
			continue
		}
		if len(remaining) == pLen && remaining == p {
			return true
		}
		if t.tryFit(remaining[pLen:]) {
			return true
		}
	}
	return false
}
