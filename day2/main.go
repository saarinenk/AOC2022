package main

// A for Rock,
// B for Paper,
// C for Scissors

// X for Rock, 1
// Y for Paper, 2
// Z for Scissors, 3

// 0 if you lost,
// 3 if the round was a draw,
// 6 if you won

import (
	"fmt"
	"os"
	"strings"
)

var wins = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var equals = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var losses = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

func win(game []string) bool {
	return game[1] == wins[game[0]]
}

func tie(game []string) bool {
	return game[1] == equals[game[0]]
}

func main() {
	pointmap := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("error", err)
	}

	str := string(data)
	games := strings.Split(str, "\n")

	fmt.Println("s1 or s2?")
	var s string
	fmt.Scanln(&s)

	if s == "s1" {
		fmt.Println(s1(games, pointmap))
	} else if s == "s2" {
		fmt.Println(s2(games, pointmap))
	}
}

func s1(games []string, pointmap map[string]int) int {
	points := 0

	for _, g := range games {
		game := strings.Split(g, " ")

		if win(game) {
			points += pointmap[game[1]] + 6
		} else if tie(game) {
			points += pointmap[game[1]] + 3
		} else {
			points += pointmap[game[1]]
		}
	}

	return points
}

// X means you need to lose,
// Y means you need to end the round in a draw,
// Z means you need to win

func s2(games []string, pointmap map[string]int) int {
	points := 0

	for _, g := range games {
		game := strings.Split(g, " ")
		move := game[1]

		if move == "X" {
			points += pointmap[losses[game[0]]]
		} else if move == "Y" {
			points += pointmap[equals[game[0]]] + 3
		} else {
			points += pointmap[wins[game[0]]] + 6
		}
	}

	return points
}
