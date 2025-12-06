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
	total := 0
	input_rows := strings.Split(input, "\n")

	operand_rows := make([][]int, len(input_rows)-1)
	var operator_row []string

	// Read input
	for i, input_row := range input_rows[:len(input_rows)-1] {
		for operand_str := range strings.FieldsSeq(input_row) {
			num, _ := strconv.Atoi(operand_str)
			operand_rows[i] = append(operand_rows[i], num)
		}
	}
	operator_row = strings.Fields(input_rows[len(input_rows)-1])

	// Do calculations
	FIRST_ROW_INDEX := 0
	num_problems := len(operand_rows[FIRST_ROW_INDEX])
	for problem := range num_problems {
		operator := operator_row[problem]

		subtotal := operand_rows[FIRST_ROW_INDEX][problem]
		for _, operand_row := range operand_rows[1:] {
			if operator == "+" {
				subtotal += operand_row[problem]
			} else {
				subtotal *= operand_row[problem]
			}
		}
		total += subtotal
	}
	return strconv.Itoa(total)
}

func solve_part_two(input string) string {
	total := 0
	input_rows := strings.Split(input, "\n")

	operator_row := input_rows[len(input_rows)-1] + "  "

	num_rows := len(input_rows)
	num_cols := len(input_rows[0])
	num_problems := len(strings.Fields(input_rows[0]))

	operands_by_problem := make([][]int, num_problems)
	operators_by_problem := make([]string, num_problems)

	problem_idx := 0
	operand_chars := make([]string, len(input_rows)-1)
	for col := num_cols - 1; col >= 0; col-- { // Step through input columns from right to left
		// Get the digits from the column
		for row := 0; row < num_rows-1; row++ {
			operand_chars[row] = string(input_rows[row][col])
		}

		// Convert completed operand to int and store
		combined_str := strings.Join(operand_chars, "")
		num, _ := strconv.Atoi(strings.TrimSpace(combined_str))
		operands_by_problem[problem_idx] = append(operands_by_problem[problem_idx], num)

		// If operator row contains an operator, move to the next problem
		operator_char := string(operator_row[col])
		problem_finished := operator_char != " "
		if problem_finished {
			operators_by_problem[problem_idx] = operator_char
			problem_idx++
			col-- // skip empty column
		}
	}

	// Do calculations
	FIRST_ROW_INDEX := 0
	for problem := range problem_idx {
		operands := operands_by_problem[problem]
		subtotal := operands[FIRST_ROW_INDEX]
		for i := range operands[1:] {
			operand := operands[i+1]
			if operators_by_problem[problem] == "+" {
				subtotal += operand
			} else {
				subtotal *= operand
			}
		}
		total += subtotal

	}

	return strconv.Itoa(total)
}
