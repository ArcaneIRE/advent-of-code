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

func calculate_paper_needed(w int, h int, l int) int {
	paper_needed := 0

	l_w_side := l * w
	w_h_side := w * h
	h_l_side := h * l

	paper_needed += 2 * l_w_side
	paper_needed += 2 * w_h_side
	paper_needed += 2 * h_l_side

	paper_needed += min(l_w_side, w_h_side, h_l_side)

	return paper_needed
}

func solve_part_one(input string) string {
	wrapping_paper_needed := 0
	presents := strings.Split(input, "\n")
	for _, present := range presents {
		sides := make([]int, 3)
		for i, side_str := range strings.Split(present, "x") {
			sides[i], _ = strconv.Atoi(side_str)
		}
		wrapping_paper_needed += calculate_paper_needed(sides[0], sides[1], sides[2])

	}
	return strconv.Itoa(wrapping_paper_needed)
}

func calculate_ribbon_needed(w int, h int, l int) int {
	paper_needed := 0

	l_w_side := 2 * (l + w)
	w_h_side := 2 * (w + h)
	h_l_side := 2 * (h + l)
	paper_needed += min(l_w_side, w_h_side, h_l_side)

	paper_needed += l * w * h

	return paper_needed
}

func solve_part_two(input string) string {
	ribbon_needed := 0
	presents := strings.SplitSeq(input, "\n")
	for present := range presents {
		sides := make([]int, 3)
		for i, side_str := range strings.Split(present, "x") {
			sides[i], _ = strconv.Atoi(side_str)
		}
		ribbon_needed += calculate_ribbon_needed(sides[0], sides[1], sides[2])
	}
	return strconv.Itoa(ribbon_needed)
}
