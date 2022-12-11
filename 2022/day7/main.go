package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	name   string
	size   int
	isDir  bool
	parent *file
	files  []*file
}

func (f *file) findDir(name string) *file {
	for _, file := range f.files {
		if file.name == name && file.isDir {
			return file
		}
	}
	return nil
}

func (f *file) addFile(name string, size int) {
	f.files = append(f.files, &file{name: name, size: size, parent: f})
	f.increaseSize(size)
}

func (f *file) addDir(name string) {
	f.files = append(f.files, &file{name: name, isDir: true, parent: f})
}

func (f *file) increaseSize(size int) {
	f.size += size
	if f.parent != nil {
		f.parent.increaseSize(size)
	}
}

func main() {
	data, _ := ReadInput()

	tree := walk(data)

	maxSize := 100_000
	fmt.Println("Part 1: ", findCandidatesFileSize(tree, maxSize))

	totalDisk := 70_000_000
	minimumFree := 30_000_000
	usedSpace := tree.size
	freeSpace := totalDisk - usedSpace
	neededSpace := minimumFree - freeSpace
	candidates := findDeletableFile(tree, neededSpace)

	sort.Slice(candidates, func(i int, j int) bool {
		return candidates[i].size < candidates[j].size
	})

	fmt.Println("Part 2: ", candidates[0].size)
}

func walk(input []string) *file {
	root := &file{name: "root", files: []*file{}}
	current := root

	for i := 0; i < len(input); i++ {
		split := strings.Split(input[i], " ")

		switch {
		case split[0] == "$":
			switch {
			case split[1] == "cd":
				switch {
				case split[2] == "..":
					current = current.parent
				case split[2] == "/":
					current = root
				default:
					current = current.findDir(split[2])
				}

			case split[1] == "ls":
				for {
					if i == len(input)-1 || input[i+1][0] == '$' {
						break
					}

					i++
					stdin := input[i]
					stdout := strings.Split(stdin, " ")

					if stdout[0] == "dir" {
						current.addDir(stdout[1])
					} else {
						size, _ := strconv.Atoi(stdout[0])
						current.addFile(stdout[1], size)
					}
				}
			}
		}
	}

	return root
}

func findCandidatesFileSize(dir *file, maxSize int) int {
	total := 0

	if dir.size <= maxSize {
		total += dir.size
	}

	for _, file := range dir.files {
		if file.isDir {
			total += findCandidatesFileSize(file, maxSize)
		}
	}

	return total
}

func findDeletableFile(dir *file, neededSpace int) []*file {
	candidates := []*file{}

	if dir.size >= neededSpace {
		candidates = append(candidates, dir)
	}

	for _, file := range dir.files {
		candidates = append(candidates, findDeletableFile(file, neededSpace)...)
	}

	return candidates
}

// 6483228 too high
// 8005232 Too high
