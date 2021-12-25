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
	lowPoints := findLowPoints(heightMap)
	part1 := calculateRiskLevel(lowPoints)
	println(part1)
}

func getFileContents(filepath string) (heightMap *CoordinateSystem) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	y := 0
	rows := make([][]*CoordinateValue, 0)
	for scanner.Scan() {
		line := scanner.Text()

		digits := make([]*CoordinateValue, 0)
		for x, v := range line {
			digits = append(digits, &CoordinateValue{x, y, int(v - '0')})
		}
		rows = append(rows, digits)
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return &CoordinateSystem{rows}
}

func findLowPoints(heightMap *CoordinateSystem) []*CoordinateValue {
	values := make([]*CoordinateValue, 0)
	for y := 0; y < heightMap.GetMaxY(); y++ {
		for x := 0; x < heightMap.GetMaxX(); x++ {
			v := heightMap.Values[y][x]
			neighbors := heightMap.GetNeighborsOf(v)
			if v.IsLowerThan(neighbors) {
				values = append(values, v)
			}
		}
	}
	return values
}

func calculateRiskLevel(lowPoints []*CoordinateValue) int {
	riskLevel := 0
	for _, v := range lowPoints {
		riskLevel += v.Value + 1
	}
	return riskLevel
}
