package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {
	input_path := flag.String("input", "./input", "The input data")
	heightMap := getFileContents(*input_path)
	part1 := calculateRiskLevel(heightMap)
	println(part1)
}

func getFileContents(filepath string) (heightMap [][]int) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		digits := make([]int, 0)
		for _, v := range line {
			digits = append(digits, int(v-'0'))
		}
		heightMap = append(heightMap, digits)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return heightMap
}

func calculateRiskLevel(heightMap [][]int) int {
	riskLevel := 0
	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[y]); x++ {
			v := heightMap[y][x]
			if x == 0 && y == 0 {
				riskLevel += IncrementedValueIfLowerThan(v, heightMap[y][x+1], heightMap[y+1][x])
			} else if x == 0 && y == len(heightMap)-1 {
				riskLevel += IncrementedValueIfLowerThan(v, heightMap[y][x+1], heightMap[y-1][x])
			} else if x == len(heightMap[y])-1 && y == 0 {
				riskLevel += IncrementedValueIfLowerThan(v, heightMap[y][x-1], heightMap[y+1][x])
			} else if x == len(heightMap[y])-1 && y == len(heightMap)-1 {
				riskLevel += IncrementedValueIfLowerThan(v, heightMap[y][x-1], heightMap[y-1][x])
			} else if y == 0 {
				riskLevel += IncrementedValueIfLowerThan(v, heightMap[y][x-1], heightMap[y][x+1], heightMap[y+1][x])
			} else if y == len(heightMap)-1 {
				riskLevel += IncrementedValueIfLowerThan(v, heightMap[y][x-1], heightMap[y][x+1], heightMap[y-1][x])
			} else if x == 0 {
				riskLevel += IncrementedValueIfLowerThan(v, heightMap[y][x+1], heightMap[y+1][x], heightMap[y-1][x])
			} else if x == len(heightMap[y])-1 {
				riskLevel += IncrementedValueIfLowerThan(v, heightMap[y][x-1], heightMap[y+1][x], heightMap[y-1][x])
			} else {
				riskLevel += IncrementedValueIfLowerThan(v, heightMap[y][x-1], heightMap[y-1][x], heightMap[y][x+1], heightMap[y+1][x])
			}
		}
	}
	return riskLevel
}

func IncrementedValueIfLowerThan(value int, others ...int) int {
	for _, v := range others {
		if v <= value {
			return 0
		}
	}
	return value + 1
}
