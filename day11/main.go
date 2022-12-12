package main

import (
	"aoc2022/common"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

var reg, _ = regexp.Compile(`\b\d{1,2}\b`)

type monkey struct {
	items       []int
	operation   func(int) int
	divisible   int
	next        func(int) int
	inspections int
}

func parseOperation(str string) func(int) int {
	op := strings.Fields(strings.Split(str, " = ")[1])
	num := op[2]
	if op[1] == "*" {
		if num == "old" {
			return func(old int) int {
				return old * old
			}
		} else {
			return func(old int) int {
				return common.ToInt(num) * old
			}
		}
	} else {
		return func(old int) int {
			return common.ToInt(num) + old
		}
	}
}

func parseNext(str []string) func(int) int {
	divisible := common.ToInt(strings.Split(str[0], "by ")[1])
	trueMonkey := common.ToInt(strings.Split(str[1], "monkey ")[1])
	falseMonkey := common.ToInt(strings.Split(str[2], "monkey ")[1])

	return func(num int) int {
		if num%divisible == 0 {
			return trueMonkey
		}
		return falseMonkey
	}
}

func createMonkey(lines []string) monkey {
	m := monkey{}
	nums := reg.FindAllString(lines[1], -1)
	items := make([]int, len(nums))
	for i, num := range nums {
		items[i] = common.ToInt(num)
	}
	m.items = items
	m.operation = parseOperation(lines[2])
	m.divisible = common.ToInt(strings.Split(lines[3], "by ")[1])
	m.next = parseNext(lines[3:6])
	m.inspections = 0

	return m
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("error", err)
	}

	str := string(data)
	monkeysString := strings.Split(str, "\n\n")
	monkeys := make([]monkey, len(monkeysString))
	for i, m := range monkeysString {
		lines := strings.Split(m, "\n")

		monkeys[i] = createMonkey(lines)
	}

	fmt.Println("s1 or s2?")
	var s string
	fmt.Scanln(&s)

	if s == "s1" {
		fmt.Println(s1(monkeys))
	} else if s == "s2" {
		fmt.Println(s2(monkeys))
	}
}

func s1(monkeys []monkey) int {
	for i := 0; i < 20; i++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				worry := m.operation(item) / 3
				nextMonkey := m.next(worry)
				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, worry)
				monkeys[i].inspections++
			}
			monkeys[i].items = []int{}
		}
	}

	inspections := []int{}
	for _, m := range monkeys {
		inspections = append(inspections, m.inspections)
	}
	insLen := len(inspections)
	sort.Ints(inspections)
	return (inspections[insLen-1] * inspections[insLen-2])
}

func s2(monkeys []monkey) int {
	mod := 1
	for _, m := range monkeys {
		mod *= m.divisible
	}

	for i := 0; i < 10000; i++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				worry := m.operation(item) % mod
				nextMonkey := m.next(worry)
				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, worry)
				monkeys[i].inspections++
			}
			monkeys[i].items = []int{}
		}
	}

	inspections := []int{}
	for _, m := range monkeys {
		inspections = append(inspections, m.inspections)
	}
	insLen := len(inspections)
	sort.Ints(inspections)
	return (inspections[insLen-1] * inspections[insLen-2])
}
