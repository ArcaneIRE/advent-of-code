package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	paths := [2]string{"test_input.txt", "input.txt"}
	// paths := [1]string{"test_input.txt"}
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

const EMPTY = '.'
const SOURCE = 'S'
const BEAM = '|'
const SPLITTER = '^'

func solve_part_one(input string) string {
	input_lines := strings.Split(input, "\n")
	grid := make([][]rune, len(input_lines))
	for i, row := range input_lines {
		trimmed_row := strings.TrimSpace(row)
		grid[i] = []rune(trimmed_row)
	}

	split_count := 0

	for i, row := range grid {
		if i == 0 {
			continue // Skip first row
		}
		for j, symbol := range row {
			symbol_above := grid[i-1][j]
			is_below_beam := symbol_above == SOURCE || symbol_above == BEAM
			if !is_below_beam {
				continue // do nothing
			}
			switch symbol {
			case EMPTY:
				row[j] = BEAM
			case SPLITTER:
				split_count++
				// Set adjacents to BEAMs, if in bounds
				if (j) > 0 { // bounds check
					row[j-1] = BEAM
				}
				if j < len(row)-1 { // bounds check
					row[j+1] = BEAM
				}
			}
		}
	}

	return strconv.Itoa(split_count)
}

type Coordinate struct {
	Row, Col int
}

// Starting at the given coordinates, trace a tachyon path recursively
//
// Returns the number of timelines the tachyon ends up on from this point
func trace_tachyon(grid []string, row int, col int, splitter_cache map[Coordinate]int) int {
	new_row := row

	// Descend until we hit a splitter or the end
	for new_row < len(grid)-1 {
		new_row++
		key := Coordinate{new_row, col}

		if grid[new_row][col] == SPLITTER {
			// Hit a splitter
			// Check cache to see if we've visited before
			val, splitter_cached := splitter_cache[key]
			if splitter_cached {
				return val
			}

			// New splitter, count timelines
			timelines_reachable := 0
			if col > 0 { // can branch left
				// Count paths for left branch
				timelines_reachable += trace_tachyon(grid, new_row, col-1, splitter_cache)
			}
			if col < len(grid[0]) { // can branch right
				// Count paths for right branch
				timelines_reachable += trace_tachyon(grid, new_row, col+1, splitter_cache)
			}
			splitter_cache[key] = timelines_reachable // Cache result
			return timelines_reachable
		}
	}
	// Base case: we've hit the bottom
	return 1
}

func solve_part_two(input string) string {
	input_lines := strings.Split(input, "\n")
	grid := make([]string, len(input_lines))
	for i, row := range input_lines {
		trimmed_row := strings.TrimSpace(row)
		grid[i] = trimmed_row
	}

	var source_col int
	for j, char := range grid[0] {
		if char == 'S' {
			source_col = j
		}
	}
	return strconv.Itoa(trace_tachyon(grid, 0, source_col, map[Coordinate]int{}))
}
