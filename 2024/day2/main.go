package day2

import (
	"log"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/slices"
)


func PrintSolutions() {
	lines := reader.ReadInputFile("./day2/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {

    safe:=0

	for _, line := range lines {

		l := slices.SliceConvert(strings.Split(line, " "), strconv.Atoi)

        pos:=(l[1]-l[0] > 0)

		s:=true
		i:=1

		for s && (i < len(l)){

		    if pos {
		        if (l[i]-l[i-1] < 1 ) || (l[i]-l[i-1] > 3 ){
		            s=false
		        }
		    } else {
		        if (l[i]-l[i-1] > -1 ) || (l[i]-l[i-1] < -3 ){
		            s=false
		        }
		    }
            i+=1
        }

        if s {
            safe += 1
        }

	}

	return safe
}

func part2(lines []string) int {
	return 0
}
