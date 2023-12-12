package day4

import (
	"slices"
)

type Card struct {
	Number         int
	WinningNumbers []int
	MyNumbers      []int
}

func (c *Card) FromLine(line string) {
	line = line[4:] // skip initial "Card" string

	state := 0
	numBuffer := 0
	for _, r := range line {
		switch state {
		case 0:
			if r >= '0' && r <= '9' {
				numBuffer *= 10
				numBuffer += int(r - '0')
			}
			if r == ':' {
				c.Number = numBuffer
				numBuffer = 0
				state = 1
			}
		case 1, 2:
			// Edge case: if the actual number is 0 it won't work
			if r == ' ' && numBuffer > 0 {
				if state == 1 {
					c.WinningNumbers = append(c.WinningNumbers, numBuffer)
				} else {
					c.MyNumbers = append(c.MyNumbers, numBuffer)
				}
				numBuffer = 0
			}
			if r >= '0' && r <= '9' {
				numBuffer *= 10
				numBuffer += int(r - '0')
			}
			if r == '|' {
				state = 2
			}
		}
	}
	c.MyNumbers = append(c.MyNumbers, numBuffer)
}

// GetWinningScore checks if MyNumbers has an overlap
// with the WinningNumbers and calculates the score from it
// The first overlap counts as score 1, with every subsequent overlap
// doubling the score
func (c *Card) GetWinningScore() int {
	matches := 0
	for _, my := range c.MyNumbers {
		if slices.Contains(c.WinningNumbers, my) {
			matches += 1
		}
	}
	if matches == 0 {
		return 0
	}
	return 1 * (1 << (matches - 1))
}

// GetWinningMatchCount counts how many WinningNumbers
// are contained in the MyNumbers
func (c *Card) GetWinningMatchCount() int {
	matches := 0
	for _, my := range c.MyNumbers {
		if slices.Contains(c.WinningNumbers, my) {
			matches += 1
		}
	}
	return matches
}
