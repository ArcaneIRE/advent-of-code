package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]
	input, _ := os.Open(path)

	part_one_answer := solve_part_one(input)
	fmt.Println(part_one_answer)

	_, _ = input.Seek(0, io.SeekStart)

	part_two_answer := solve_part_two(input)
	fmt.Println(part_two_answer)
}

func solve_part_one(f *os.File) string {
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	id_ranges := strings.SplitSeq(scanner.Text(), ",")
	invalid_id_sum := 0

	for id_range := range id_ranges {
		range_boundaries := strings.Split(id_range, "-")

		start, _ := strconv.Atoi(range_boundaries[0])
		end, _ := strconv.Atoi(range_boundaries[1])

		for id := start; id <= end; id++ {
			s := strconv.Itoa(id)
			s1 := s[:len(s)/2]
			s2 := s[len(s)/2:]
			if s1 == s2 {
				invalid_id_sum += id
			}
		}
	}

	return strconv.Itoa(invalid_id_sum)
}

func solve_part_two(f *os.File) string {
	scanner := bufio.NewScanner(f)

	invalid_id_sum := 0
	scanner.Scan()

	is_repeated_substring := func(full_str string, sub_str string) bool {
		sub_str_length := len(sub_str)
		full_str_length := len(full_str)
		if full_str_length%sub_str_length != 0 {
			return false
		}
		for start_pos := len(sub_str); start_pos <= full_str_length-sub_str_length; start_pos += sub_str_length {
			chunk := full_str[start_pos : start_pos+sub_str_length]
			if sub_str != chunk {
				return false
			}
		}
		return true
	}

	is_invalid := func(s string) bool {
		for sub_str_len := 1; sub_str_len <= len(s)/2; sub_str_len++ {
			if is_repeated_substring(s, s[:sub_str_len]) {
				return true
			}
		}
		return false
	}

	for product_range := range strings.SplitSeq(scanner.Text(), ",") {
		range_boundaries := strings.Split(product_range, "-")

		start, _ := strconv.Atoi(range_boundaries[0])
		end, _ := strconv.Atoi(range_boundaries[1])

		for id := start; id <= end; id++ {
			s := strconv.Itoa(id)
			if is_invalid(s) {
				invalid_id_sum += id
			}
		}
	}

	return strconv.Itoa(invalid_id_sum)
}
