package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sum(array []string) int {
	result := 0
	for _, num := range array {
		intNum, _ := strconv.Atoi(num)
		result += intNum
	}
	return result
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("error", err)
	}

	str := string(data)
	splitStr := strings.Split(str, "\n\n")

	var calSum []int

	for _, s := range splitStr {
		stringArr := strings.Split(s, "\n")

		cals := sum(stringArr)
		calSum = append(calSum, cals)
	}

	sort.Ints(calSum)
	topThree := calSum[len(calSum)-3:]

	fmt.Println("s1 or s2?")
	var s string
	fmt.Scanln(&s)

	if s == "s1" {
		fmt.Println("Result:", topThree[2])
	} else if s == "s2" {
		sum := topThree[0] + topThree[1] + topThree[2]
		fmt.Println("Result:", sum)
	}

	sort.Ints(calSum)
}
