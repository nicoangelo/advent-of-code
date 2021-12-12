package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input_path := flag.String("input", "./input", "The input data")
	initial_fish := getFileContents(*input_path)
	res80 := part1(initial_fish, 80)
	fmt.Println("Part 1:", res80)
	res256 := part1(initial_fish, 256)
	fmt.Println("Part 2:", res256)
}

func getFileContents(filepath string) []int {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln(err)
	}
	tokens := strings.Split(string(content), ",")

	values := []int{}
	for _, v := range tokens {
		val, _ := strconv.Atoi(strings.TrimSpace(v))
		values = append(values, val)
	}

	return values
}

func part1(initial_fish_timer []int, days int) int {
	current_generation := make([]int, 9)
	for _, initial := range initial_fish_timer {
		current_generation[initial]++
	}

	for i := 0; i < days; i++ {
		rollover := current_generation[0]
		next_generation := append(current_generation[1:], rollover)
		next_generation[6] += rollover
		current_generation = next_generation
	}

	return sum(current_generation)
}

func sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}
