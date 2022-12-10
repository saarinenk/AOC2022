package main

import (
	"aoc2022/common"
	"fmt"
	"math"
	"os"
	"strings"
)

type coordinate struct {
	lat int
	lng int
}

func splitCommand(comm string) (string, int) {
	split := strings.Split(comm, " ")
	return split[0], common.ToInt(split[1])
}

func tooFar(coord1 coordinate, coord2 coordinate) bool {
	distLng := math.Abs(float64(coord1.lng) - float64(coord2.lng))
	distLat := math.Abs(float64(coord1.lat) - float64(coord2.lat))
	if coord1 == coord2 {
		return false
	} else if distLng <= 1 && distLat <= 1 {
		return false
	}
	return true
}

func moves(dir string) (int, int) {
	if dir == "U" {
		return 0, 1
	} else if dir == "R" {
		return 1, 0
	} else if dir == "D" {
		return 0, -1
	} else {
		return -1, 0
	}
}

func moveH(coor coordinate, dir string) coordinate {
	lat, lng := moves(dir)
	return coordinate{coor.lat + lat, coor.lng + lng}
}

func moveKnot(coord1 coordinate, coord2 coordinate) coordinate {
	coord := coordinate{coord2.lat, coord2.lng}
	if coord1.lat > coord2.lat {
		coord = coordinate{coord2.lat + 1, coord.lng}
	} else if coord1.lat < coord2.lat {
		coord = coordinate{coord2.lat - 1, coord.lng}
	}
	if coord1.lng > coord2.lng {
		coord = coordinate{coord.lat, coord2.lng + 1}
	} else if coord1.lng < coord.lng {
		coord = coordinate{coord.lat, coord2.lng - 1}
	}
	return coord
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
		fmt.Println(makePath(commands, 2))
	} else if s == "s2" {
		fmt.Println(makePath(commands, 10))
	}
}

func makePath(commands []string, knots int) int {
	visited := map[coordinate]bool{}

	rope := make([]coordinate, knots)
	for i := 0; i < knots; i++ {
		rope[0] = coordinate{0, 0}
	}
	visited[rope[knots-1]] = true

	for _, comm := range commands {
		dir, steps := splitCommand(comm)
		for i := 0; i < steps; i++ {
			rope[0] = moveH(rope[0], dir)

			for i := 1; i < knots; i++ {
				if tooFar(rope[i-1], rope[i]) {
					rope[i] = moveKnot(rope[i-1], rope[i])
				}
			}
			visited[rope[knots-1]] = true
		}
	}
	return len(visited)
}
