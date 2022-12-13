package main

import (
	"aoc2022/common"
	"fmt"
	"os"
	"strings"
)

func enqueue(queue []pos, element pos) []pos {
	queue = append(queue, element)
	return queue
}

func makeMap(rows []string) [][]string {
	trees := [][]string{}

	for _, r := range rows {
		splitRow := strings.Split(r, "")
		trees = append(trees, splitRow)
	}
	return trees
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("error", err)
	}

	str := string(data)
	rowSlice := strings.Split(str, "\n")
	strmap := makeMap(rowSlice)

	fmt.Println("s1 or s2?")
	var s string
	fmt.Scanln(&s)

	if s == "s1" {
		fmt.Println(s1(strmap))
	} else if s == "s2" {
		fmt.Println(s2(strmap))
	}
}

type pos struct {
	x int
	y int
}

func neighbors(strmap [][]string, n pos, h int, w int) []pos {
	xn := n.x
	yn := n.y
	neighbors := []pos{{xn, yn - 1}, {xn, yn + 1}, {xn - 1, yn}, {xn + 1, yn}}
	suitables := []pos{}
	for _, nbr := range neighbors {
		if nbr.x < 0 || nbr.y < 0 || nbr.x >= h || nbr.y >= w {
			continue
		}
		char := strmap[xn][yn]
		nchar := strmap[nbr.x][nbr.y]
		if strings.Index(common.Alphabet, nchar)-strings.Index(common.Alphabet, char) <= 1 {
			suitables = append(suitables, nbr)
		}
	}
	return suitables
}

func bfs(strmap [][]string, start, end pos) int {
	visited := make(map[pos]int)
	height := len(strmap)
	width := len(strmap[0])
	path := height * width
	q := []pos{start}
	visited[start] = 0

	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		if v == end {
			path = visited[end]
			break
		}
		neighbors := neighbors(strmap, v, height, width)
		for _, nbr := range neighbors {
			if _, ok := visited[nbr]; ok {
				continue
			}
			q = enqueue(q, nbr)
			visited[nbr] = visited[v] + 1
		}
	}

	return path
}

func s1(strmap [][]string) int {
	modifiedMap := strmap
	start := pos{0, 0}
	end := pos{0, 0}
	height := len(strmap)
	width := len(strmap[0])

	// replacing start and end with the right heights
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if strmap[i][j] == "S" {
				start = pos{i, j}
				modifiedMap[i][j] = "a"
			}
			if strmap[i][j] == "E" {
				end = pos{i, j}
				modifiedMap[i][j] = "z"
			}
		}
	}

	return bfs(modifiedMap, start, end)
}

func s2(strmap [][]string) int {
	modifiedMap := strmap
	starts := []pos{}
	end := pos{0, 0}
	height := len(strmap)
	width := len(strmap[0])

	// replacing start and end with the right heights
	// adding possible starts to a slice
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if strmap[i][j] == "a" {
				starts = append(starts, pos{i, j})
			}
			if strmap[i][j] == "S" {
				starts = append(starts, pos{i, j})
				modifiedMap[i][j] = "a"
			}
			if strmap[i][j] == "E" {
				end = pos{i, j}
				modifiedMap[i][j] = "z"
			}
		}
	}

	path := height * width
	for _, start := range starts {
		newPath := bfs(modifiedMap, start, end)

		if newPath < path {
			path = newPath
		}
	}
	return path
}
