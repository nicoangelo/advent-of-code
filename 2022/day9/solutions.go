package day9

func linesToDirections(lines []string) *[]Instruction {
	directions := &[]Instruction{}
	for _, v := range lines {
		d := &Instruction{}
		d.FromLine(v)
		*directions = append(*directions, *d)
	}
	return directions
}

func part1(lines []string) int {
	dirs := linesToDirections(lines)
	r := NewRope(1)
	for _, d := range *dirs {
		r.MoveHead(&d)
	}
	return len(*r.KnotPositionsHistory[0])
}

func part2(lines []string) int {
	dirs := linesToDirections(lines)
	r := NewRope(9)
	for _, d := range *dirs {
		r.MoveHead(&d)
		// r.PrintKnotPositions()
	}
	// r.PrintKnotHistoryPositions(8)
	return len(*r.KnotPositionsHistory[8])
}
