package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]

	content, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(content))

	part_one_answer := solve_part_one(input)
	fmt.Println(part_one_answer)

	part_two_answer := solve_part_two(input)
	fmt.Println(part_two_answer)
}

func solve_part_one(input string) string {
	floor := 0
	for _, char := range input {
		if char == '(' {
			floor++
		} else {
			floor--
		}
	}
	return strconv.Itoa(floor)
}

func solve_part_two(input string) string {
	floor := 0
	for i, char := range input {
		if char == '(' {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			return strconv.Itoa(i + 1)
		}
	}
	return "Error"
}
