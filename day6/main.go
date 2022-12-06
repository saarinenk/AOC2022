package main

import (
	"aoc2022/common"
	"fmt"
	"os"
)

func hasDuplicates(str string) bool {
	chars := make(map[string]bool, 0)
	for i := 0; i < len(str); i++ {
		if chars[string(str[i])] == true {
			return true
		} else {
			chars[string(str[i])] = true
		}
	}
	return false
}

func markerPosition(str string, numOfChars int) int {
	pos := 0
	for i := 0; i <= len(str)-numOfChars; i++ {
		if !hasDuplicates(str[i : i+numOfChars]) {
			pos = i
			break
		}
	}
	return pos + numOfChars
}

// s1 and s2
func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("error", err)
	}

	str := string(data)

	fmt.Println("How many chars?")
	var num string
	fmt.Scanln(&num)

	markerPos := markerPosition(str, common.ToInt(num))
	fmt.Println(markerPos)
}
