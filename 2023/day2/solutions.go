package day2

func part1(lines []string) int {
	targetMax := &CubeReveal{Red: 12, Green: 13, Blue: 14}
	sum := 0
	for _, l := range lines {
		gr := GameRecordFromLine(l)
		if gr.HasMaxCubes(targetMax) {
			sum += gr.GameId
		}
	}

	return sum
}

func part2(lines []string) int {
	totalPower := 0
	for _, l := range lines {
		gr := GameRecordFromLine(l)
		totalPower += gr.GetMinimumNecessaryCubeColors().GetPower()
	}

	return totalPower
}
