package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type PositionFuel struct {
	Position int
	SumFuel  int
}

type PositionFuelMap []PositionFuel

func (p PositionFuelMap) Len() int           { return len(p) }
func (p PositionFuelMap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PositionFuelMap) Less(i, j int) bool { return p[i].SumFuel < p[j].SumFuel }

func main() {
	input_path := flag.String("input", "./input", "The input data")
	crab_positions, max := getFileContents(*input_path)
	res1 := CalculatePositionFuels(crab_positions, max)
	part1 := res1[0]
	fmt.Println("Part 1 - position:", part1.Position, "/ Fuel:", part1.SumFuel)
}

func getFileContents(filepath string) (values []int, max int) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln(err)
	}
	tokens := strings.Split(string(content), ",")

	for _, v := range tokens {
		val, _ := strconv.Atoi(strings.TrimSpace(v))
		values = append(values, val)
		if val > max {
			max = val
		}
	}

	return values, max
}

func CalculatePositionFuels(crab_positions []int, max int) PositionFuelMap {
	position_fuels := make(PositionFuelMap, max+1)
	for i := range position_fuels {
		position_fuels[i] = PositionFuel{i, 0}
	}

	for _, crab_pos := range crab_positions {
		for i := 0; i < len(position_fuels); i++ {
			diff := int(math.Abs(float64(crab_pos - position_fuels[i].Position)))
			position_fuels[i].SumFuel += diff
		}
	}

	sort.Sort(position_fuels)
	return position_fuels
}
