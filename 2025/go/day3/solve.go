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

func calculate_joltage(bank string, battery_length int) int {
	bank_length := len(bank)
	joltage_str := ""

	max_value := -1
	start_of_range := -1

	for i := range battery_length {
		remaining_batteries := battery_length - i - 1

		// Find max voltage in available range
		max_value = -1
		max_value_index := 0
		end_of_range := bank_length - 1 - remaining_batteries

		for j := end_of_range; j > start_of_range; j-- {
			candidate, _ := strconv.Atoi(string(bank[j]))
			if candidate >= max_value {
				max_value = candidate
				max_value_index = j
			}
		}

		// Add digit to string
		joltage_str += strconv.Itoa(max_value)
		// Limit range for next iteration
		start_of_range = max_value_index
	}

	joltage, _ := strconv.Atoi(joltage_str)
	return joltage
}

func solve_part_one(input string) string {
	battery_banks := strings.SplitSeq(input, "\n")
	total_joltage := 0

	for bank := range battery_banks {
		joltage := calculate_joltage(bank, 2)
		total_joltage += joltage
	}

	return strconv.Itoa(total_joltage)
}

func solve_part_two(input string) string {
	battery_banks := strings.SplitSeq(input, "\n")
	total_joltage := 0

	for bank := range battery_banks {
		joltage := calculate_joltage(bank, 12)
		total_joltage += joltage
	}

	return strconv.Itoa(total_joltage)
}
