package day2

import (
	"log/slog"
	"strconv"
	"strings"
)

type GameRecord struct {
	GameId      int
	CubeReveals []CubeReveal
}

func GameRecordFromLine(line string) *GameRecord {

	tokens := strings.Split(line[5:], ":")
	num, err := strconv.Atoi(tokens[0])
	if err != nil {
		slog.Error("Cannot convert game id")
	}

	revealTokens := strings.Split(tokens[1], ";")
	gr := GameRecord{num, make([]CubeReveal, len(revealTokens))}

	for i, v := range revealTokens {
		gr.CubeReveals[i] = *cubeRevealFromString(v)
	}
	return &gr
}

// HasMaxCubes returns true if the given GameRecord list a maximum
// of the given colors in any CubeReveal.
func (gr *GameRecord) HasMaxCubes(target *CubeReveal) bool {
	for _, reveal := range gr.CubeReveals {
		if !reveal.HasMaxCubes(target) {
			return false
		}
	}
	return true
}

// GetMinimumNecessaryCubeColors returns the number of cubes of each color
// that must be present in the bag in order for the game to be possible
func (gr *GameRecord) GetMinimumNecessaryCubeColors() (cr *CubeReveal) {
	cr = &CubeReveal{}
	for _, reveal := range gr.CubeReveals {
		if reveal.Blue > cr.Blue {
			cr.Blue = reveal.Blue
		}
		if reveal.Green > cr.Green {
			cr.Green = reveal.Green
		}
		if reveal.Red > cr.Red {
			cr.Red = reveal.Red
		}
	}
	return cr
}

type CubeReveal struct {
	Red   int
	Green int
	Blue  int
}

func cubeRevealFromString(reveal string) *CubeReveal {
	cr := &CubeReveal{}
	for _, v := range strings.Split(reveal, ",") {
		v = strings.TrimSpace(v)
		tok := strings.Split(v, " ")
		if len(tok) != 2 {
			slog.Error("Cube reveal string is not two elements long", "reveal", reveal)
			return nil
		}
		amount, err := strconv.Atoi(tok[0])
		if err != nil {
			slog.Error("Cannot convert revealed number of cubes", "number", tok[0])
			return nil
		}
		switch tok[1] {
		case "red":
			cr.Red = amount
		case "green":
			cr.Green = amount
		case "blue":
			cr.Blue = amount
		}
	}
	return cr
}

func (cr *CubeReveal) HasMaxCubes(target *CubeReveal) bool {
	return cr.Blue <= target.Blue && cr.Green <= target.Green && cr.Red <= target.Red
}

func (cr *CubeReveal) GetPower() int {
	return cr.Blue * cr.Green * cr.Red
}
