package day6

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
)

/* This problem is a simple quadratic equation
 * While we could probably enumerate all possible solutions
 * for part1, that's gonna take a while for part 2.

d...Distance traveled
c...Charge time
r...Total Race time

d = c * (r - c)

or...

c^2 - r*c + d = 0

then just apply the well-known solution of a quadratic equation
**/

func PrintSolutions() {
	lines := reader.ReadInputFile("./day6/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {
	return solveForRaceData(readRaceData(lines))
}

func part2(lines []string) int {
	return solveForRaceData(readRaceDataDekern(lines))
}

func readRaceData(lines []string) (res []RaceData) {
	times := strings.Fields(lines[0][5:])
	record := strings.Fields(lines[1][9:])
	res = make([]RaceData, len(times))
	for i, v := range times {
		trt, err := strconv.Atoi(v)
		rr, err := strconv.Atoi(record[i])
		if err != nil {
			return nil
		}
		res[i] = RaceData{TotalRaceTime: trt, RaceRecord: rr}
	}
	return res
}

func readRaceDataDekern(lines []string) (res []RaceData) {
	rd := RaceData{}
	for _, r := range lines[0][5:] {
		if r == ' ' {
			continue
		}
		rd.TotalRaceTime *= 10
		rd.TotalRaceTime += int(r - '0')
	}
	for _, r := range lines[1][9:] {
		if r == ' ' {
			continue
		}
		rd.RaceRecord *= 10
		rd.RaceRecord += int(r - '0')
	}
	res = make([]RaceData, 1)
	res[0] = rd
	return res
}

func solveForRaceData(d []RaceData) int {
	ans := 1
	for _, v := range d {
		totalRaceTime := float64(v.TotalRaceTime)
		recordTime := float64(v.RaceRecord)

		// solving for the record time, we would not win
		// but just be equal. Therefore, solve for the next integer step
		x1, x2 := getChargeTimeSolutions(totalRaceTime, recordTime+1)
		fmt.Println("x1:", x1, "x2:", x2)
		ans *= x2 - x1 + 1
	}

	return ans
}

func getChargeTimeSolutions(totalRaceTime float64, timeToBeat float64) (min int, max int) {
	sqrt := math.Sqrt(math.Pow(totalRaceTime, 2) - 4*timeToBeat)
	x1 := math.Ceil((totalRaceTime - sqrt) / 2)
	x2 := math.Floor((totalRaceTime + sqrt) / 2)
	return int(x1), int(x2)
}
