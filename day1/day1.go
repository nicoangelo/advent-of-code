package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./data")
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)

	previous_num := 0
	increased_count := 0
	first_run := true

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		if first_run {
			first_run = false
		} else if num > previous_num {
			increased_count++
		}
		previous_num = num
	}

	fmt.Println("Answer:", increased_count)

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
