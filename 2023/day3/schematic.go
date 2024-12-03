package day3

import (
	"slices"

	"github.com/nicoangelo/aoc-pkg/slicemath"
)

type Schematic struct {
	Parts     []Part
	PartsMask slicemath.Matrix2D[*PartArea]
}

func (s *Schematic) fromLines(lines []string, partsMaskChars []rune) {
	s.Parts = make([]Part, 0)
	s.readPartsMask(lines, partsMaskChars)
	s.readParts(lines)
}

func (s *Schematic) readParts(lines []string) {
	numBuffer := 0
	var activePartArea *PartArea
	for y, row := range lines {
		numBuffer = 0
		for x, col := range row {
			if col >= '0' && col <= '9' {
				numBuffer *= 10
				numBuffer += int(col - '0')
				if activePartArea == nil {
					activePartArea = s.PartsMask.At(slicemath.Coord2D{X: x, Y: y})
				}
				continue
			}
			if numBuffer != 0 && activePartArea != nil {
				activePartArea.PartCount++
				s.addPartNumber(numBuffer, activePartArea)
			}
			numBuffer = 0
			activePartArea = nil
		}
		if numBuffer != 0 && activePartArea != nil {
			activePartArea.PartCount++
			s.addPartNumber(numBuffer, activePartArea)
		}
		numBuffer = 0
		activePartArea = nil
	}
}

func (s *Schematic) readPartsMask(lines []string, includeChars []rune) {
	s.PartsMask.Init(slicemath.Coord2D{Y: len(lines), X: len(lines[0])})

	for y, l := range lines {
		for x, col := range l {
			if !slices.Contains(includeChars, col) {
				continue
			}
			s.PartsMask.SetAndExpand(slicemath.Coord2D{X: x, Y: y}, &PartArea{}, 1)
		}
	}
}

func (s *Schematic) addPartNumber(number int, partArea *PartArea) {
	s.Parts = append(s.Parts, Part{Number: number, PartArea: partArea})
}

type Part struct {
	Number   int
	PartArea *PartArea
}

type PartArea struct {
	PartCount int
}
