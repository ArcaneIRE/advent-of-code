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

func create_grid_from_input(input string) [][]bool {
	input_rows := strings.SplitSeq(input, "\n")
	var output_grid [][]bool

	row_count := 0
	for input_row := range input_rows {
		output_grid = append(output_grid, make([]bool, len(input_row)))
		for i, char := range input_row {
			if char == '@' {
				output_grid[row_count][i] = true
			} else {
				output_grid[row_count][i] = false
			}
		}
		row_count += 1
	}

	return output_grid
}

func is_in_bounds(i int, j int, grid [][]int) bool {
	if i < 0 || j < 0 || i > (len(grid)-1) || j > len(grid[0]) {
		return false
	}
	return true
}

type Coord struct {
	i int
	j int
}

func count_adjacent_cells(i int, j int, grid [][]bool) int {
	count := 0
	adjacents := []Coord{
		{-1, 0},  // N
		{1, 0},   // S
		{0, -1},  // W
		{0, 1},   // E
		{-1, -1}, // NW
		{-1, 1},  // NE
		{1, -1},  // SW
		{1, 1},   // SE
	}

	row_count := len(grid)
	col_count := len(grid[0])

	for _, dir := range adjacents {
		new_i := i + dir.i
		new_j := j + dir.j

		if new_i >= 0 && new_i < row_count && new_j >= 0 && new_j < col_count {
			if grid[new_i][new_j] == true {
				count += 1
			}
		}
	}

	return count
}

func solve_part_one(input string) string {
	input_grid := create_grid_from_input(input)

	moveable_count := 0

	for i, row := range input_grid {
		for j, cell := range row {
			cell_has_paper := cell == true
			if cell_has_paper {
				if count_adjacent_cells(i, j, input_grid) < 4 {
					moveable_count += 1
				}
			}
		}
	}

	return strconv.Itoa(moveable_count)
}

func solve_part_two(input string) string {
	input_grid := create_grid_from_input(input)

	moveable_count := 0
	moved_last_iteration := true
	for moved_last_iteration {
		moved_last_iteration = false
		for i, row := range input_grid {
			for j, cell := range row {
				cell_has_paper := cell == true
				if cell_has_paper {
					if count_adjacent_cells(i, j, input_grid) < 4 {
						input_grid[i][j] = false
						moveable_count += 1
						moved_last_iteration = true
					}
				}
			}
		}

	}

	return strconv.Itoa(moveable_count)
}
