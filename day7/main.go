package main

import (
	"aoc2022/common"
	"fmt"
	"os"
	"strings"
)

var sizemap = map[string]int{}

func getSize(lines []string, index int) int {
	return 0
}

func buildMap(lines []string) {
	path := ""
	for i, l := range lines {
		splitLine := strings.Fields(l)
		switch splitLine[0] {
		case "dir":
			continue
		case "$":
			switch splitLine[1] {
			case "ls":
				continue
			case "cd":
				if splitLine[2] == ".." {
					lastSlashIndex := strings.LastIndex(path, "/")
					path = path[:lastSlashIndex]
				} else {
					// adding i in case of duplicate folder names
					dir := "root"

					if splitLine[2] == "/" {
						path += dir
					} else {
						dir = splitLine[2] + fmt.Sprint(i)
						path += "/" + dir
					}
					sizemap[dir] = 0
				}
			}
		default:
			for _, dir := range strings.Split(path, "/") {
				sizemap[dir] += common.ToInt(splitLine[0])
			}
		}

	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("error", err)
	}

	str := string(data)
	lines := strings.Split(str, "\n")
	buildMap(lines)

	fmt.Println("s1 or s2?")
	var s string
	fmt.Scanln(&s)

	if s == "s1" {
		fmt.Println(s1())
	} else if s == "s2" {
		fmt.Println(s2())
	}
}

func s1() int {
	sizes := 0
	for _, value := range sizemap {
		if value < 100000 {
			sizes += value
		}
	}
	return sizes
}

func s2() int {
	maxSpace := 70000000
	neededSpace := 30000000
	spaceLeft := maxSpace - sizemap["root"]
	dirSize := maxSpace

	for _, value := range sizemap {
		if spaceLeft+value > neededSpace {
			if value < dirSize {
				dirSize = value
			}
		}
	}

	return dirSize
}
