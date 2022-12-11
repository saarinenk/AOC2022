package main

import (
	"aoc2022/common"
	"fmt"
	"os"
	"strings"
)

func printResult(slices [][]string) {
	for i := 0; i < len(slices); i++ {
		fmt.Println(slices[i])
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("error", err)
	}

	str := string(data)
	commands := strings.Split(str, "\n")

	fmt.Println("s1 or s2?")
	var s string
	fmt.Scanln(&s)

	if s == "s1" {
		fmt.Println(s1(commands))
	} else if s == "s2" {
		printResult(s2(commands))
	}
}

func s1(commands []string) int {
	cycles := make([]int, 240)
	cycles[0] = 1
	cycleIndex := 1

	for _, command := range commands {
		if command == "noop" {
			cycleIndex++
		} else {
			_, num := common.SplitCommand(command)
			cycleIndex += 2
			cycles[cycleIndex] = num
		}
	}

	value := 0
	values := []int{}
	for i := 0; i <= 220; i++ {
		value += cycles[i]
		if (i%40 - 20) == 0 {
			values = append(values, value*i)
		}
	}

	return common.SliceSum(values)
}

func s2(commands []string) [][]string {
	cycles := make([]int, 240)
	cycles[0] = 1
	cycleIndex := 0

	for _, command := range commands {
		if command == "noop" {
			cycleIndex++
		} else {
			_, num := common.SplitCommand(command)
			cycleIndex += 2
			cycles[cycleIndex] = num
		}
	}

	sum := 0
	marks := make([][]string, 6)
	for i := range marks {
		marks[i] = make([]string, 40)
	}
	for i := 0; i < 6; i++ {
		for i2 := 0; i2 < 40; i2++ {
			sum += cycles[i*40+i2]

			if i2 >= sum-1 && i2 <= sum+1 {
				marks[i][i2] = "#"
			} else {
				marks[i][i2] = "."
			}
		}
	}

	return marks
}
