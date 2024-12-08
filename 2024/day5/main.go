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

func part1(lines []string) (res int) {
	orderingRules, pageUpdates := readInstructions(lines)
	for _, u := range pageUpdates {
		res += getMiddleItemIfCorrectlyOrdered(orderingRules, u)
	}
	return res
}

func readInstructions(lines []string) (preceedingPages map[int][]int, pageUpdates [][]int) {
	preceedingPages = map[int][]int{}
	pageUpdates = [][]int{}
	state := 1
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
	return preceedingPages, pageUpdates
}

// getMiddleItemIfCorrectlyOrdered finds the updates that are correctly ordered. Correct order is
// defined as those pages that adhere to the "ordering rules" in orderingRules. Each entry in
// that map holds all pages that must appear before it.
// It returns 0 if the pages are not correctly ordered, or the value of the middle item otherwise.
func getMiddleItemIfCorrectlyOrdered(orderingRules map[int][]int, pages []int) int {
	// accumulate the pages that must have already appeared before
	alreadyBefore := []int{}
	for _, p := range pages {
		// if the current page is already in the list of pages that must have appeared before...
		if slices.Contains(alreadyBefore, p) {
			return 0
		}
		alreadyBefore = append(alreadyBefore, orderingRules[p]...)
	}
	// we assume that there is an odd number of elements so there actually is a middle item
	return pages[len(pages)/2]
}

func part2(lines []string) (res int) {
	orderingRules, pageUpdates := readInstructions(lines)
	for _, u := range pageUpdates {
		res += getMiddleItemAfterReordering(orderingRules, u)
	}
	return res
}

func getMiddleItemAfterReordering(orderingRules map[int][]int, pages []int) int {
	reordered := false
	for i, p := range pages {
		mustBefore := orderingRules[p]

		for j, la := range pages[i+1:] {
			pos := i + j
			if slices.Contains(mustBefore, la) {
				reordered = true
				swapped := []int{la, p}
				if pos == 0 {
					pages = append(swapped, pages[pos+2:]...)
				} else if pos == len(pages)-2 {
					pages = append(pages[:pos], swapped...)
				} else {
					before := pages[:pos]
					after := pages[pos+2:]
					pages = append(before, swapped...)
					pages = append(pages, after...)
				}
			}
		}
	}
	if reordered {
		// we assume that there is an odd number of elements so there actually is a middle item
		return pages[len(pages)/2]
	}
	return 0
}
