package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path := os.Args[1]
	file, err := os.Open(path)
	check(err)

	scanner := bufio.NewScanner(file)

	zeros := 0
	dial_value := 50

	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])

		magnitude, err := strconv.Atoi(line[1:])
		check(err)

		if direction == "R" {
			dial_value = (dial_value + magnitude) % 100
		} else {
			dial_value = (dial_value + (100 - magnitude)) % 100
		}

		if dial_value == 0 {
			zeros += 1
		}
	}

	fmt.Println(zeros)
}
