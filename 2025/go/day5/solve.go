package main

import (
	"fmt"
	"os"
	"sort"
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

type ID_range struct {
	start int
	end   int
}

func get_input(input string) ([]ID_range, []int) {
	parts := strings.SplitN(strings.TrimSpace(input), "\n\n", 2)

	var id_ranges []ID_range
	var ingredients []int

	id_range_strings := strings.Split(strings.TrimSpace(parts[0]), "\n")
	id_ranges = make([]ID_range, 0, len(id_range_strings))
	for _, str := range id_range_strings {
		s := strings.TrimSpace(str)
		values := strings.Split(s, "-")
		start, _ := strconv.Atoi(strings.TrimSpace(values[0]))
		end, _ := strconv.Atoi(strings.TrimSpace(values[1]))
		id_ranges = append(id_ranges, ID_range{start: start, end: end})
	}

	ingredients_strings := strings.Split(strings.TrimSpace(parts[1]), "\n")
	ingredients = make([]int, 0, len(ingredients_strings))
	for _, str := range ingredients_strings {
		s := strings.TrimSpace(str)
		v, _ := strconv.Atoi(s)
		ingredients = append(ingredients, v)
	}

	return id_ranges, ingredients
}

func is_fresh(ingredient int, id_ranges []ID_range) bool {
	for _, id_range := range id_ranges {
		if id_range.start <= ingredient && ingredient <= id_range.end {
			return true
		}
	}
	return false
}

func solve_part_one(input string) string {
	id_ranges, ingredients := get_input(input)

	fresh_count := 0
	for _, ingredient := range ingredients {
		if is_fresh(ingredient, id_ranges) {
			fresh_count += 1
		}
	}

	return strconv.Itoa(fresh_count)
}

type ByStart []ID_range

func (s ByStart) Len() int           { return len(s) }
func (s ByStart) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByStart) Less(i, j int) bool { return s[i].start < s[j].start }

func solve_part_two(input string) string {
	id_ranges, _ := get_input(input)

	// Combine ranges
	sort.Sort(ByStart(id_ranges))
	var combined_ranges []ID_range

	start := id_ranges[0].start
	end := id_ranges[0].end
	for i := 1; i < len(id_ranges); i++ {
		current := id_ranges[i]
		if current.start > end {
			// Disjoint Range, append and start new
			combined_ranges = append(combined_ranges, ID_range{start: start, end: end})
			start = current.start
			end = current.end
		} else if current.end > end {
			// Extend current range
			end = current.end
		}
	}
	combined_ranges = append(combined_ranges, ID_range{start: start, end: end})

	// Sum ranges
	fresh_count := 0
	for _, id_range := range combined_ranges {
		fresh_count += (id_range.end - id_range.start) + 1
	}

	return strconv.Itoa(fresh_count)
}
