package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type History struct {
	Commands *[]Command
}

type Command struct {
	Command string
	Output  *[]string
}

type Directory struct {
	Name        string
	Files       *map[string]File
	Directories *map[string]Directory
	Parent      *Directory
}

type File struct {
	Name string
	Size int
}

func newDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		name,
		&map[string]File{},
		&map[string]Directory{},
		parent,
	}
}

func (h *History) ToFiletree() (root *Directory) {
	root = newDirectory("", nil)
	var currentDir *Directory
	for _, c := range *h.Commands {
		cmdTokens := strings.Split(c.Command, " ")
		if cmdTokens[0] == "cd" && cmdTokens[1] == "/" {
			currentDir = root
		} else if cmdTokens[0] == "cd" && cmdTokens[1] != ".." {
			d := (*currentDir.Directories)[cmdTokens[1]]
			currentDir = &d
		} else if cmdTokens[0] == "cd" && cmdTokens[1] == ".." {
			currentDir = currentDir.Parent
		} else if cmdTokens[0] == "ls" {
			for _, f := range *c.Output {
				fileInfo := strings.Split(f, " ")
				if fileInfo[0] == "dir" {
					newDir := newDirectory(fileInfo[1], currentDir)
					(*currentDir.Directories)[fileInfo[1]] = *newDir
				} else {
					size, _ := strconv.Atoi(fileInfo[0])
					newFile := File{fileInfo[1], size}
					(*currentDir.Files)[fileInfo[1]] = newFile
				}
			}
		}
	}
	return root
}

func (dir Directory) FindByPredicate(p func(Directory) bool) (res []Directory) {
	res = make([]Directory, 0)

	if p(dir) {
		res = append(res, dir)
	}
	for _, d := range *dir.Directories {
		res = append(res, d.FindByPredicate(p)...)
	}
	return res
}

func (dir Directory) GetTotalSize() (sum int) {
	for _, f := range *dir.Files {
		sum += f.Size
	}
	for _, d := range *dir.Directories {
		sum += d.GetTotalSize()
	}
	return sum
}

func (d Directory) String() string {
	return fmt.Sprintf("%s: %d", d.Name, d.GetTotalSize())
}
