package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./data")

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	increments := part_one(file)
	fmt.Println("Answer part one:", increments)

	file.Seek(0, io.SeekStart) // rewind file to the beginning
	sliding_increments := part_two(file)
	fmt.Println("Answer part two:", sliding_increments)
}

func part_one(file *os.File) int {

	scanner := bufio.NewScanner(file)

	previous_num := -1
	increments := 0

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		if previous_num != -1 && num > previous_num {
			increments++
		}
		previous_num = num
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return increments
}

func part_two(file *os.File) int {

	scanner := bufio.NewScanner(file)

	previous_sum := -1
	increments := 0

	rolling_sum := 0
	buffered_sums := 0
	sliding_buffer := []int{0, 0, 0}

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)

		rolling_sum += num - sliding_buffer[0]
		sliding_buffer = append(sliding_buffer[1:], num)

		if previous_sum != -1 && buffered_sums >= 3 && rolling_sum > previous_sum {
			increments++
		}

		previous_sum = rolling_sum
		buffered_sums++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return increments
}
