package day5

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/sliceutils"
)

func PrintSolutions() {
	lines := reader.ReadInputFile("./day5/input")
	part1 := part1(lines)
	log.Println("Part 1: ", part1)

	part2 := part2(lines)
	log.Println("Part 2: ", part2)
}

func part1(lines []string) int {
	state := 1
	preceedingPages := map[int][]int{}
	pageUpdates := [][]int{}
	for _, l := range lines {
		if l == "" {
			state = 2
			continue
		}
		if state == 1 {
			var page, before int
			fmt.Sscanf(l, "%d|%d", &page, &before)
			preceedingPages[before] = append(preceedingPages[before], page)
		} else if state == 2 {
			u := sliceutils.SliceConvert(strings.Split(l, ","), strconv.Atoi)
			pageUpdates = append(pageUpdates, u)
		}
	}
	return getCorrectlyOrderedMiddleItemSum(preceedingPages, pageUpdates)
}

// getCorrectlyOrderedMiddleItemSum finds the updates that are correctly ordered. Correct order is
// defined as those pages that adhere to the "ordering rules" in preceedingPages. Each entry in
// that map holds all pages that must appear before it.
func getCorrectlyOrderedMiddleItemSum(preceedingPages map[int][]int, updates [][]int) (res int) {
out:
	for _, u := range updates {
		// accumulate the pages that must have already appeared before
		alreadyBefore := []int{}
		for _, p := range u {
			// if the current page is already in the list of pages that must have appeared before...
			if slices.Contains(alreadyBefore, p) {
				continue out
			}
			alreadyBefore = append(alreadyBefore, preceedingPages[p]...)
		}
		// we assume that there is an odd number of elements so there actually is a middle item
		res += u[len(u)/2]
	}
	return res
}

func part2(lines []string) int {
	return 0
}
