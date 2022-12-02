package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input_path := flag.String("input", "./input", "The input data")
	contents, line_count := get_file_contents(*input_path)
	part1(contents, line_count)
	part2(contents, line_count)
}

func part1(contents [][]int, line_count int) int64 {
	ones := count_ones(contents)
	most_common_bits := get_most_common_bits(ones, line_count)
	decimal, _ := strconv.ParseInt(most_common_bits, 2, 32)
	inverted_decimal, _ := strconv.ParseInt(flip_bits(most_common_bits), 2, 32)
	result := decimal * inverted_decimal

	fmt.Printf("Total lines: %d, ones: %d\n", line_count, ones)
	fmt.Println("gamma: ", decimal, ", epsilon:", inverted_decimal)
	fmt.Println("power consumption:", result)

	return result
}

func part2(contents [][]int, line_count int) int64 {
	most_common_bits := find_column_value_by_criterion(contents, true)
	decimal, _ := strconv.ParseInt(most_common_bits, 2, 32)
	least_common_bits := find_column_value_by_criterion(contents, false)
	inverted_decimal, _ := strconv.ParseInt(least_common_bits, 2, 32)
	result := decimal * inverted_decimal

	fmt.Println("oxygen generator rating: ", decimal, ", co2 scrubber rating:", inverted_decimal)
	fmt.Println("life support rating:", result)

	return result
}

func find_column_value_by_criterion(contents [][]int, most_common bool) string {
	columns_count := len(contents[0])
	result := ""
	var selected_indices []int = nil

	for i := 0; i < columns_count; i++ {
		iteration_length := len(contents)
		if selected_indices != nil {
			iteration_length = len(selected_indices)
		}
		if iteration_length == 1 {
			result += strconv.Itoa(contents[selected_indices[0]][i])
			continue
		}

		one_indices := []int{}
		zero_indices := []int{}
		for j := 0; j < iteration_length; j++ {
			current_index := j
			if selected_indices != nil {
				current_index = selected_indices[j]
			}
			if contents[current_index][i] == 0 {
				zero_indices = append(zero_indices, current_index)
			} else {
				one_indices = append(one_indices, current_index)
			}
		}

		if (most_common && len(one_indices) >= len(zero_indices)) ||
			(!most_common && len(one_indices) < len(zero_indices)) {
			selected_indices = one_indices
			result += "1"
		} else {
			selected_indices = zero_indices
			result += "0"
		}
	}
	return result
}

func get_file_contents(filepath string) (contents [][]int, line_count int) {

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.SplitAfter(line, "")
		var row []int
		for _, v := range tokens {
			bit, _ := strconv.Atoi(v)
			row = append(row, bit)
		}
		contents = append(contents, row)
		line_count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return contents, line_count
}

func count_ones(bits [][]int) []int {
	var ones []int

	for _, row := range bits {

		for i, bit := range row {
			if len(ones) < i+1 {
				ones = append(ones, bit)
			} else {
				ones[i] += bit
			}
		}
	}
	return ones
}

func get_most_common_bits(ones []int, line_count int) string {
	binary_string := ""
	for _, v := range ones {
		if v > line_count/2 {
			binary_string += "1"
		} else {
			binary_string += "0"
		}
	}
	return binary_string
}

func flip_bits(bits string) string {
	flipped := ""
	for _, v := range bits {
		if v == '0' {
			flipped += "1"
		} else {
			flipped += "0"
		}
	}
	return flipped
}
