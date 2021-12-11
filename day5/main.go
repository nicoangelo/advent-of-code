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

type VentVector struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
}

func (vector *VentVector) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", vector.StartX, vector.StartY, vector.EndX, vector.EndY)
}

func main() {
	input_path := flag.String("input", "./input", "The input data")
	vectors, maxX, maxY := getFileContents(*input_path)
	system := markPerpendicularOnSystem(vectors, maxX, maxY)
	fmt.Println("Part 1:", countCrossingPoints(system))
	system = markAllOnSystem(vectors, maxX, maxY)
	fmt.Println("Part 2:", countCrossingPoints(system))
}

func (vector *VentVector) IsHorizontal() bool {
	return vector.StartY == vector.EndY
}

func (vector *VentVector) IsVertical() bool {
	return vector.StartX == vector.EndX
}

func (vector *VentVector) HasHorizontalComponent() bool {
	return vector.StartX != vector.EndX
}

func (vector *VentVector) HasVerticalComponent() bool {
	return vector.StartY != vector.EndY
}

func (vector *VentVector) IsPerpendicular() bool {
	return vector.IsHorizontal() || vector.IsVertical()
}

func (vector *VentVector) Length() (int, signX int, signY int) {
	lenX, signX := getDiffAndSign(vector.EndX, vector.StartX)
	lenY, signY := getDiffAndSign(vector.EndY, vector.StartY)

	if lenX > 0 && lenY > 0 {
		return lenX + 1, signX, signY
	} else if lenX > 0 {
		return lenX + 1, signX, 0
	}
	return lenY + 1, 0, signY
}

func getFileContents(filepath string) (ventVectors []*VentVector, maxX int, maxY int) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		vectorTokens := strings.Split(line, " -> ")

		if len(vectorTokens) != 2 {
			fmt.Printf("Line '%s' must contain exactly two coordinates. Skipping", line)
			continue
		}
		vector := new(VentVector)
		vector.StartX, vector.StartY = parseCoordinate(vectorTokens[0])
		vector.EndX, vector.EndY = parseCoordinate(vectorTokens[1])

		if vector.EndX > maxX {
			maxX = vector.EndX
		}
		if vector.EndY > maxY {
			maxY = vector.EndY
		}
		if vector.StartX > maxX {
			maxX = vector.StartX
		}
		if vector.StartY > maxY {
			maxY = vector.StartY
		}
		ventVectors = append(ventVectors, vector)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return ventVectors, maxX + 1, maxY + 1
}

func parseCoordinate(token string) (x int, y int) {
	coordTokens := strings.Split(token, ",")

	if len(coordTokens) != 2 {
		fmt.Printf("Token '%s' must contain exactly two values. Skipping", token)
		return 0, 0
	}
	x, _ = strconv.Atoi(strings.TrimSpace(coordTokens[0]))
	y, _ = strconv.Atoi(strings.TrimSpace(coordTokens[1]))
	return x, y
}

func markVectorsOnSystem(vectors []*VentVector, maxX int, maxY int, selector func(*VentVector) bool) [][]int {
	system := make2D(maxX, maxY)

	for _, vector := range vectors {
		if selector(vector) {
			fmt.Println(vector)
			len, hor, vert := vector.Length()
			for i := 0; i < len; i++ {
				system[vector.StartY+(i*vert)][vector.StartX+(i*hor)]++
			}
		}
	}
	return system
}

func markPerpendicularOnSystem(vectors []*VentVector, maxX int, maxY int) [][]int {
	return markVectorsOnSystem(
		vectors,
		maxX,
		maxY,
		func(v *VentVector) bool { return v.IsPerpendicular() })
}

func markAllOnSystem(vectors []*VentVector, maxX int, maxY int) [][]int {
	return markVectorsOnSystem(
		vectors,
		maxX,
		maxY,
		func(v *VentVector) bool { return true })
}

func make2D(x int, y int) [][]int {
	system := make([][]int, y)
	for i := 0; i < len(system); i++ {
		system[i] = make([]int, x)
	}
	return system
}

func countCrossingPoints(system [][]int) int {
	sum := 0
	for _, y := range system {
		for _, x := range y {
			if x > 1 {
				sum++
			}
		}
	}
	return sum
}

func printSystem(system [][]int) {
	for _, y := range system {
		fmt.Print("[")
		for _, x := range y {
			if x == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", x)
			}
		}
		fmt.Println("]")
	}
}

func getDiffAndSign(a int, b int) (abs int, sign int) {
	diff := a - b
	if diff < 0 {
		return diff * -1, -1
	}
	return diff, 1
}
