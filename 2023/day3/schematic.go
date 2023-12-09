package day3

import (
	"github.com/nicoangelo/aoc-pkg/math"
)

type Schematic struct {
	Parts     []Part
	PartsMask math.Matrix2D[bool]
}

func (s *Schematic) PartsSum() (sum int) {
	sum = 0
	for _, p := range s.Parts {
		sum += p.Number
	}
	return sum
}

func (s *Schematic) fromLines(lines []string) {
	s.Parts = make([]Part, 0)
	s.readPartsMask(lines)
	numBuffer := 0
	numBufferActive := false
	for y, row := range lines {
		numBuffer = 0
		for x, col := range row {
			if col >= '0' && col <= '9' {
				numBuffer *= 10
				numBuffer += int(col - '0')
				if !numBufferActive {
					numBufferActive = s.PartsMask.At(math.Coord2D{X: x, Y: y})
				}
				continue
			}
			if numBuffer != 0 && numBufferActive {
				s.addPartNumber(numBuffer)
			}
			numBuffer = 0
			numBufferActive = false
		}
		if numBuffer != 0 && numBufferActive {
			s.addPartNumber(numBuffer)
		}
		numBuffer = 0
		numBufferActive = false
	}
}

func (s *Schematic) readPartsMask(lines []string) {
	s.PartsMask.Init(math.Coord2D{Y: len(lines), X: len(lines[0])})

	for y, l := range lines {
		for x, col := range l {
			if col == '.' || (col >= '0' && col <= '9') {
				continue
			}
			s.PartsMask.SetAndExpand(math.Coord2D{X: x, Y: y}, true, 1)
		}
	}
}

func (s *Schematic) addPartNumber(number int) {
	s.Parts = append(s.Parts, Part{Number: number})
}

type Part struct {
	Number int
}
