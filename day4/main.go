package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func pairToInt(pair string) []int {
	areas := strings.Split(pair, "-")
	slice := []int{toInt(areas[0]), toInt(areas[1])}
	return slice
}

func pairsToSlice(pairs []string) [][]int {
	var areas [][]int
	for _, pair := range pairs {
		areas = append(areas, pairToInt(pair))
	}
	return areas
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("error", err)
	}

	str := string(data)
	listOfAssignments := strings.Split(str, "\n")

	fmt.Println("s1 or s2?")
	var s string
	fmt.Scanln(&s)

	if s == "s1" {
		fmt.Println(s1(listOfAssignments))
	} else if s == "s2" {
		fmt.Println(s2(listOfAssignments))
	}
}

func s1(list []string) int {
	count := 0
	for _, assignments := range list {
		pairs := strings.Split(assignments, ",")

		areas := pairsToSlice(pairs)
		lastInFirst := areas[0][0] <= areas[1][0] && areas[0][1] >= areas[1][1]
		firstInLast := areas[1][0] <= areas[0][0] && areas[1][1] >= areas[0][1]
		if lastInFirst || firstInLast {
			count += 1
		}
	}
	return count
}

func s2(list []string) int {
	count := 0
	for _, assignments := range list {
		pairs := strings.Split(assignments, ",")

		areas := pairsToSlice(pairs)
		outsideOther := areas[0][0] > areas[1][1] || areas[0][1] < areas[1][0]
		if !outsideOther {
			count += 1
		}
	}
	return count
}
