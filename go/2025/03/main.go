package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/griggsjared/advent-of-code/go/internal"
)

const (
	day  = 3
	year = 2025
)

func main() {
	internal.ProcessAndDisplayResults(day, year, Part1, Part2)
}

func Part1(input string) (int, error) {
	return sumLines(input, 2), nil
}

func Part2(input string) (int, error) {
	return sumLines(input, 12), nil
}

func parseDigits(input string) []int {
	var digits []int
	numbers := strings.Split(input, "")
	for _, numStr := range numbers {
		digit, err := strconv.Atoi(numStr)
		if err != nil {
			continue
		}
		digits = append(digits, digit)
	}
	return digits
}

func maxNumber(input string, length int) (int, error) {
	digits := parseDigits(input)

	if len(digits) < length {
		return 0, fmt.Errorf("input length %d is less than required length %d", len(digits), length)
	}

	if length < 2 {
		return 0, fmt.Errorf("length must be at least 2")
	}

	var foundDigits []int
	currentIndex := 0
	for i := range length {
		endIndex := len(digits) - (length - i - 1)
		digitSlice := digits[currentIndex:endIndex]
		foundDigit := slices.Max(digitSlice)
		foundDigits = append(foundDigits, foundDigit)

		for i2 := currentIndex; i2 < endIndex; i2++ {
			if digits[i2] == foundDigit {
				currentIndex = i2 + 1
				break
			}
		}
	}

	combined := ""
	for _, digit := range foundDigits {
		combined += strconv.Itoa(digit)
	}
	combinedInt, err := strconv.Atoi(combined)
	if err != nil {
		return 0, err
	}

	return combinedInt, nil
}

func sumLines(input string, length int) int {
	lines := strings.Split(input, "\n")

	totalSum := 0
	for _, line := range lines {
		number, err := maxNumber(line, length)
		if err != nil {
			continue
		}

		totalSum += number
	}

	return totalSum
}
