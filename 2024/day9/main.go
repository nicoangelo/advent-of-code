package day9

import (
	"log"
	"os/user"
	"sort"
	"strconv"
	"strings"

	"github.com/nicoangelo/aoc-pkg/reader"
	"github.com/nicoangelo/aoc-pkg/slicemath"
)

func PrintSolutions() {
	u, _ := user.Current()

	lines := reader.ReadInputFile("./day9/input_" + strings.ToLower(string(u.Username[0])))
	part1 := part1(lines[0])
	log.Println("Part 1: ", part1)

	part2 := part2(lines[0])
	log.Println("Part 2: ", part2)
}

func part1(fs string) int {
	checksum := 0
	back_i := len(fs) - 1
	frontPos := 0
	backRemaining, _ := strconv.Atoi(string(fs[back_i]))
	fileId := 0

out:
	for i := 0; i <= back_i; i++ {
		data, _ := strconv.Atoi(string(fs[i]))
		if i == back_i {
			data = backRemaining
		}
		if i%2 == 0 {
			fileId = i / 2
			for range data {
				checksum += frontPos * fileId
				frontPos++
			}
		} else if i%2 == 1 {
			for range data {
				if backRemaining == 0 {
					back_i -= 2
					backRemaining, _ = strconv.Atoi(string(fs[back_i]))
				}
				// edge case: the last hole is bigger than the amount of blocks to move from the back
				if back_i < i {
					break out
				}
				fileId = back_i / 2
				checksum += frontPos * fileId
				backRemaining--
				frontPos++
			}
		}
	}
	return checksum
}

func part2(fs string) int {
	checksum := 0
	frontPos := 0
	fileId := 0

	holes := map[int][]int{}           // key=size, value=[]frontPos
	fileIdChecksums := map[int][]int{} // key=filedId, value=checksums for all its blocks

	for i := 0; i <= len(fs)-1; i++ {
		data, _ := strconv.Atoi(string(fs[i]))
		if i%2 == 0 {
			fileId = i / 2
			for range data {
				// pre calculate the checksum if the file is NOT moved later‚
				fileIdChecksums[fileId] = append(fileIdChecksums[fileId], frontPos*fileId)
				frontPos++
			}
		} else if i%2 == 1 {
			// collect holes
			holes[data] = append(holes[data], frontPos)
			frontPos += data
		}
	}
	for i := len(fs) - 1; i >= 0; i -= 2 {
		fileId = i / 2
		data, _ := strconv.Atoi(string(fs[i]))
		h := popFittingHole(holes, data)
		if h == 0 {
			continue
		}
		delete(fileIdChecksums, fileId)
		for j := h; j < data+h; j++ {
			checksum += j * fileId
		}
	}
	for _, c := range fileIdChecksums {
		checksum += slicemath.Sum(c)
	}
	return checksum
}

func popFittingHole(m map[int][]int, size int) int {
	// holes are between size and 9
	for i := size; i <= 9; i++ {
		if v, ok := m[i]; ok {
			m[i] = v[1:]
			if i != size {
				// if we found a hole that is bigger than size
				// a new hole is available with a smaller size
				newHoleSize := i - size
				m[newHoleSize] = append(m[newHoleSize], v[0]+size)
				sort.Slice(m[newHoleSize], func(i, j int) bool {
					return m[newHoleSize][i] < m[newHoleSize][j]
				})
			}
			if len(m[i]) == 0 {
				delete(m, i)
			}
			return v[0]
		}
	}
	return 0
}
