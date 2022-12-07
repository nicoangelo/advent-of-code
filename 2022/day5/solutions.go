package day5

func part1(lines []string) string {
	co := &CargoOperations{}
	co.FillFromLines(&lines)
	for _, op := range *co.Operations {
		for i := 0; i < op.Amount; i++ {
			lastRune := op.Source.Crates[len(op.Source.Crates)-1]
			op.Source.Crates = op.Source.Crates[0 : len(op.Source.Crates)-1]
			op.Destination.Crates = append(op.Destination.Crates, lastRune)
		}
	}
	return string(co.GetTopCrates())
}

func part2(lines []string) string {
	co := &CargoOperations{}
	co.FillFromLines(&lines)
	for _, op := range *co.Operations {
		sliceLength := len(op.Source.Crates) - op.Amount
		lastRunes := op.Source.Crates[sliceLength:]
		op.Source.Crates = op.Source.Crates[0:sliceLength]
		op.Destination.Crates = append(op.Destination.Crates, lastRunes...)
	}
	return string(co.GetTopCrates())
}
