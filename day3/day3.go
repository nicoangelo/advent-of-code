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
	read_file(*input_path)
}

func read_file(filepath string) int {

	contents, line_count := get_file_contents(filepath)
	ones := count_ones(contents)
	most_common_bits := get_most_common_bits(ones, line_count)
	decimal, _ := strconv.ParseInt(most_common_bits, 2, 32)
	inverted_decimal, _ := strconv.ParseInt(flip_bits(most_common_bits), 2, 32)
	result := decimal * inverted_decimal

	fmt.Printf("Total lines: %d, ones: %d\n", line_count, ones)
	fmt.Println("gamma: ", decimal, ", epsilon:", inverted_decimal)
	fmt.Println("power consumption:", result)

	return int(result)
}

func get_file_contents(filepath string) ([][]int, int) {

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var contents [][]int
	line_count := 0

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
