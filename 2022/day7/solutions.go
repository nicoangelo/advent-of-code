package day7

import (
	"sort"
)

func parseHistoryLines(lines []string) (res *History) {
	res = &History{&[]Command{}}
	var currentContext *Command
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		if l[0] == '$' {
			if currentContext != nil {
				*res.Commands = append(*res.Commands, *currentContext)
			}
			currentContext = &Command{l[2:], &[]string{}}
		} else {
			*currentContext.Output = append(*currentContext.Output, l)
		}
	}
	if currentContext != nil {
		*res.Commands = append(*res.Commands, *currentContext)
	}
	return res
}

func part1(lines []string) (res int) {
	h := parseHistoryLines(lines)
	ft := h.ToFiletree()
	dirs := ft.FindByPredicate(func(d Directory) bool { return d.GetTotalSize() <= 100000 })

	for _, d := range dirs {
		res += d.GetTotalSize()
	}
	return res
}

func part2(lines []string) (res int) {
	h := parseHistoryLines(lines)
	ft := h.ToFiletree()
	used := ft.GetTotalSize()

	disk_size := 70000000
	needed := 30000000
	to_delete := needed - (disk_size - used)

	dirs := ft.FindByPredicate(func(d Directory) bool { return d.GetTotalSize() >= to_delete })
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].GetTotalSize() < dirs[j].GetTotalSize()
	})
	return dirs[0].GetTotalSize()

}
