package main

import (
	"aoc2022/common"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Stack []string

func (s *Stack) push(str string) {
	*s = append(*s, str)
}

func (s *Stack) pop() string {
	if len(*s) == 0 {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func instructionToSlice(instructions string) []int {
	reg, err := regexp.Compile(`\b\d{1,2}\b`)
	if err != nil {
		panic(err)
	}

	numsFromText := reg.FindAllStringSubmatch(instructions, 3)
	nums := []int{}
	for _, num := range numsFromText {
		nums = append(nums, common.ToInt(num[0]))
	}
	return nums
}

func formatStack(stack string) []string {
	slice := []string{}
	for i := 1; i < len(stack); i += 4 {
		slice = append(slice, string(stack[i]))
	}
	return slice
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("error", err)
	}

	str := string(data)
	input := strings.Split(str, "\n\n")
	instructions := strings.Split(input[1], "\n")

	rows := strings.Split(input[0], "\n")
	rows = rows[:len(rows)-1]
	numOfStacks := len(formatStack(rows[0]))

	stacks := make([]Stack, numOfStacks)

	// add crates to stacks
	for i := 0; i < len(rows); i++ {
		stackSlice := formatStack(rows[len(rows)-1-i])
		for i2, c := range stackSlice {
			if c != " " {
				stacks[i2].push(c)
			}
		}
	}

	fmt.Println("s1 or s2?")
	var s string
	fmt.Scanln(&s)

	if s == "s1" {
		fmt.Println(s1(instructions, stacks))
	} else if s == "s2" {
		fmt.Println(s2(instructions, stacks))
	}
}

func s1(instructions []string, stacks []Stack) string {
	for _, ins := range instructions {
		is := instructionToSlice(ins)

		for i := 0; i < is[0]; i++ {
			crate := stacks[is[1]-1].pop()
			stacks[is[2]-1].push(crate)
		}
	}

	resultStr := ""

	for _, stack := range stacks {
		topCrane := stack.pop()
		resultStr += topCrane
	}

	return resultStr
}

func s2(instructions []string, stacks []Stack) string {
	for _, ins := range instructions {
		is := instructionToSlice(ins)

		storage := Stack{}

		for i := 0; i < is[0]; i++ {
			crate := stacks[is[1]-1].pop()
			storage.push(crate)
		}

		for len(storage) > 0 {
			stacks[is[2]-1].push(storage.pop())
		}
	}

	resultStr := ""

	for _, stack := range stacks {
		topCrane := stack.pop()
		resultStr += topCrane
	}

	return resultStr
}
