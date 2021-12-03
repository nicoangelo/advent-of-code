package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input")

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	depth, horizontal_pos := calculate_position_without_aim(file)
	fmt.Printf("Answer part one: %d (depth: %d, horizontal: %d)\n", depth*horizontal_pos, depth, horizontal_pos)

	file.Seek(0, io.SeekStart) // rewind file to the beginning
	depth, horizontal_pos = calculate_position_with_aim(file)
	fmt.Printf("Answer part two: %d (depth: %d, horizontal: %d)\n", depth*horizontal_pos, depth, horizontal_pos)
}

func calculate_position_without_aim(file *os.File) (int, int) {

	scanner := bufio.NewScanner(file)

	horizontal_pos := 0
	depth := 0

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.SplitN(line, " ", 2)
		move_length, _ := strconv.Atoi(tokens[1])

		switch tokens[0] {
		case "forward":
			horizontal_pos += move_length
		case "up":
			depth -= move_length
		case "down":
			depth += move_length
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return depth, horizontal_pos
}

func calculate_position_with_aim(file *os.File) (int, int) {

	scanner := bufio.NewScanner(file)

	horizontal_pos := 0
	depth := 0
	aim := 0

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.SplitN(line, " ", 2)
		move_length, _ := strconv.Atoi(tokens[1])

		switch tokens[0] {
		case "forward":
			horizontal_pos += move_length
			depth += aim * move_length
		case "up":
			aim -= move_length
		case "down":
			aim += move_length
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return depth, horizontal_pos
}
