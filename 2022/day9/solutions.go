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
	r := &Rope{}
	for _, d := range *dirs {
		r.MoveHead(&d)
	}
	return len(*r.TailHistory)
}

func part2(lines []string) int {
	return 0
}
