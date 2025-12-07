package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	paths := [2]string{"test_input.txt", "input.txt"}
	var input_strings [2]string
	for i, path := range paths {
		content, _ := os.ReadFile(path)
		input := strings.TrimSpace(string(content))
		input_strings[i] = input
	}
	for _, input := range input_strings {
		fmt.Println(solve_part_one(input))
	}

	for _, input := range input_strings {
		fmt.Println(solve_part_two(input))
	}
}

type Coord struct {
	x, y int
}

func solve_part_one(input string) string {
	current := Coord{0, 0}
	visited := make(map[Coord]bool)
	visited[current] = true // Add starting location
	for _, r := range input {
		switch r {
		case '^':
			current.y++
		case 'v':
			current.y--
		case '>':
			current.x++
		case '<':
			current.x--
		}
		visited[current] = true
	}

	return strconv.Itoa(len(visited))
}

func solve_part_two(input string) string {
	santa_location := Coord{0, 0}
	robot_location := Coord{0, 0}
	visited := make(map[Coord]bool)
	visited[santa_location] = true // Add starting location
	for i, r := range input {
		var current_santa *Coord
		if i%2 == 1 {
			current_santa = &santa_location
		} else {
			current_santa = &robot_location
		}

		switch r {
		case '^':
			current_santa.y++
		case 'v':
			current_santa.y--
		case '>':
			current_santa.x++
		case '<':
			current_santa.x--
		}
		visited[*current_santa] = true
	}

	return strconv.Itoa(len(visited))
}
