package day6

func findUniqueRuneBlockStart(datastream string, blockSize int) int {
	for i := 0; i < len(datastream)-(blockSize-1); i++ {
		uniqueRunes := make(map[rune]int, 0)
		curr := datastream[i : i+blockSize]

		for _, v := range curr {
			if _, ok := uniqueRunes[v]; ok {
				continue
			}
			uniqueRunes[v] = 1
		}
		if len(uniqueRunes) == blockSize {
			return i + blockSize
		}
	}
	return -1
}

func part1(datastream string) int {
	return findUniqueRuneBlockStart(datastream, 4)
}

func part2(datastream string) int {
	return findUniqueRuneBlockStart(datastream, 14)
}
