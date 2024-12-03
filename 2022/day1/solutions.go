package day1

import (
	"sort"
	"strconv"

	"github.com/nicoangelo/aoc-pkg/slicemath"
)

func part1(lines []string) (maxCaloriesSum int) {
	maxCaloriesSum = 0
	currentCaloriesSum := 0

	for _, v := range lines {
		if v == "" {
			if currentCaloriesSum > maxCaloriesSum {
				maxCaloriesSum = currentCaloriesSum
			}
			currentCaloriesSum = 0
		}
		currentCalorieValue, _ := strconv.Atoi(v)
		currentCaloriesSum += currentCalorieValue
	}

	return maxCaloriesSum
}

func part2(lines []string) (maxCaloriesSum int) {
	maxCaloriesSums := make([]int, 0)
	currentCaloriesSum := 0

	for _, v := range lines {
		if v == "" {
			maxCaloriesSums = append(maxCaloriesSums, currentCaloriesSum)
			currentCaloriesSum = 0
		}
		currentCalorieValue, _ := strconv.Atoi(v)
		currentCaloriesSum += currentCalorieValue
	}
	maxCaloriesSums = append(maxCaloriesSums, currentCaloriesSum)

	sort.Sort(sort.Reverse(sort.IntSlice(maxCaloriesSums)))

	return slicemath.Sum(maxCaloriesSums[0:3])
}
