// Advent of Code 2023
// Day 1
// Will Harrington
// Go 1.21

package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	str "strings"
)

var print = fmt.Println

func main() {
	dat, _ := os.ReadFile("input.txt")
	lines := str.Split(string(dat), "\n")

	// Part 1
	partOne := 0
	for _, line := range lines {
		var digits []int
		for _, char := range line {
			digit, err := strconv.Atoi(string(char))
			if err == nil {
				digits = append(digits, digit)
			}
		}
		first := digits[0]
		last := digits[len(digits)-1]
		calibration := 10*first + last
		partOne += calibration
	}

	// Part 2
	var numbers = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	partTwo := 0
	for _, line := range lines {
		var first int
		var substring []rune
	LineLoop:
		for _, char := range line {
			digit, err := strconv.Atoi(string(char))
			if err == nil {
				first = digit
				break LineLoop
			} else {
				substring = append(substring, char)
				for number, value := range numbers {
					if str.Contains(string(substring), number) {
						first = value
						break LineLoop
					}
				}
			}
		}

		var last int
		substring = []rune{}
		lineSlice := []rune(line)
		max := len(lineSlice) - 1
	LineLoopReverse:
		for i := range line {
			char := lineSlice[max-i]
			digit, err := strconv.Atoi(string(char))
			if err == nil {
				last = digit
				break LineLoopReverse
			} else {
				substring = append(substring, char)
				for number, value := range numbers {
					reverseSubstring := slices.Clone(substring)
					slices.Reverse(reverseSubstring)
					if str.Contains(string(reverseSubstring), number) {
						last = value
						break LineLoopReverse
					}
				}
			}
		}

		calibration := 10*first + last
		partTwo += calibration
	}

	print("Part 1: ", partOne)
	print("Part 2: ", partTwo)
}
