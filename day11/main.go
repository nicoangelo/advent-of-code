package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	input_path := flag.String("input", "./input", "The input data")
	octopuses := getFileContents(*input_path)
	totalFlashes := 0
	for i := 1; i <= 100; i++ {
		totalFlashes += octopuses.Step()
	}
	fmt.Println(totalFlashes)

	octopuses = getFileContents(*input_path)
	i := 0
	for i = 1; true; i++ {
		flashes := octopuses.Step()
		if flashes == 100 {
			break
		}
	}
	fmt.Println(i)
}

func getFileContents(filepath string) (octopuses *CoordinateSystem) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	y := 0
	rows := make([][]*Octopus, 0)
	for scanner.Scan() {
		line := scanner.Text()

		energyLevels := make([]*Octopus, 0)
		for x, v := range line {
			energyLevels = append(energyLevels, &Octopus{x, y, int(v - '0'), false})
		}
		rows = append(rows, energyLevels)
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return &CoordinateSystem{rows}
}
