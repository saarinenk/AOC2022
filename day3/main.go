package main

import (
	"fmt"
	"os"
	"strings"
)

func Intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("error", err)
	}

	str := string(data)
	rucksacks := strings.Split(str, "\n")

	fmt.Println("s1 or s2?")
	var s string
	fmt.Scanln(&s)

	if s == "s1" {
		fmt.Println(s1(rucksacks))
	} else if s == "s2" {
		fmt.Println(s2(rucksacks))
	}
}

func s1(rucksacks []string) int {
	prios := 0

	for _, rucksack := range rucksacks {
		strlen := len(rucksack)
		compartments := [2]string{rucksack[0:(strlen / 2)], rucksack[(strlen / 2):strlen]}
		commonLetters := Intersection(strings.Split(compartments[0], ""), strings.Split(compartments[1], ""))

		prio := strings.Index(alphabet, commonLetters[0]) + 1
		prios += prio
	}

	return prios
}

func s2(rucksacks []string) int {
	prios := 0

	for i := 0; i <= len(rucksacks)-1; i = i + 3 {
		commonLetters := Intersection(strings.Split(rucksacks[i], ""), strings.Split(rucksacks[i+1], ""))
		commonLettersOfThree := Intersection(commonLetters, strings.Split(rucksacks[i+2], ""))

		prio := strings.Index(alphabet, commonLettersOfThree[0]) + 1
		prios += prio
	}

	return prios
}
