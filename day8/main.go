package main

import (
	"aoc2022/common"
	"fmt"
	"os"
	"strings"
)

func makeTreeSlice(rows []string) [][]int {
	trees := [][]int{}

	for _, r := range rows {
		splitRow := strings.Split(r, "")
		ints := []int{}
		for _, s := range splitRow {
			ints = append(ints, common.ToInt(s))
		}
		trees = append(trees, ints)
	}
	return trees
}

func isTaller(treeSlice []int, currTree int) bool {
	for _, tree := range treeSlice {
		if tree >= currTree {
			return false
		}
	}
	return true
}

func treesVisible(treeSlice []int, currTree int) int {
	trees := 0
	for _, tree := range treeSlice {
		trees += 1
		if tree >= currTree {
			break
		}
	}
	return trees
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("error", err)
	}

	str := string(data)
	rows := strings.Split(str, "\n")
	trees := makeTreeSlice(rows)
	transposeTrees := common.Transpose(trees)

	fmt.Println("s1 or s2?")
	var s string
	fmt.Scanln(&s)

	if s == "s1" {
		fmt.Println("Visible trees:", s1(trees, transposeTrees))
	} else if s == "s2" {
		fmt.Println(s2(trees, transposeTrees))
	}
}

func s1(trees [][]int, transposeTrees [][]int) int {
	rowLen := len(trees[0])
	rowHei := len(trees)

	visibleTrees := 0

	for rowi := 0; rowi < rowLen; rowi++ {
		for coli := 0; coli < rowHei; coli++ {
			if coli == 0 || rowi == 0 || coli == rowHei-1 || rowi == rowLen-1 {
				visibleTrees += 1
				continue
			}
			currTree := trees[rowi][coli]
			treesLeft := trees[rowi][0:coli]
			treesRight := trees[rowi][coli+1 : rowLen]
			treesUp := transposeTrees[coli][0:rowi]
			treesDown := transposeTrees[coli][rowi+1 : rowHei]

			if isTaller(treesLeft, currTree) || isTaller(treesRight, currTree) {
				visibleTrees += 1
				continue
			}

			if isTaller(treesUp, currTree) || isTaller(treesDown, currTree) {
				visibleTrees += 1
				continue
			}
		}
	}

	return visibleTrees
}

func s2(trees [][]int, transposeTrees [][]int) int {
	rowLen := len(trees[0])
	rowHei := len(trees)

	highestScore := 0

	for rowi := 0; rowi < rowLen; rowi++ {
		for coli := 0; coli < rowHei; coli++ {
			if coli == 0 || rowi == 0 || coli == rowHei-1 || rowi == rowLen-1 {
				continue
			}

			// Reversing left and up to have closest trees first
			currTree := trees[rowi][coli]
			treesLeft := common.Reverse(trees[rowi][0:coli])
			treesRight := trees[rowi][coli+1 : rowLen]
			treesUp := common.Reverse(transposeTrees[coli][0:rowi])
			treesDown := common.Reverse(transposeTrees[coli][rowi+1 : rowHei])

			score := treesVisible(treesLeft, currTree) * treesVisible(treesRight, currTree) * treesVisible(treesUp, currTree) * treesVisible(treesDown, currTree)

			if score > highestScore {
				highestScore = score
			}
		}
	}
	return highestScore
}
