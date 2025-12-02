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

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func main() {
	path := os.Args[1]
	file, err := os.Open(path)
	check(err)

	scanner := bufio.NewScanner(file)

	clicks := 0
	dial_value := 50

	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])

		magnitude, err := strconv.Atoi(line[1:])
		check(err)

		clicks += magnitude / 100
		magnitude = magnitude % 100

		if direction == "R" {
			dial_value = dial_value + magnitude
			if dial_value >= 100 {
				clicks += 1
				dial_value -= 100
			}
		}
		if direction == "L" {
			if dial_value-magnitude <= 0 && dial_value != 0 {
				clicks += 1
			}
			dial_value -= magnitude
			if dial_value < 0 {
				dial_value += 100
			}
		}
	}

	fmt.Println(clicks)
}
