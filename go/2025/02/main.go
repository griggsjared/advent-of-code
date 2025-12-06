package main

import (
	"strconv"
	"strings"

	"github.com/griggsjared/advent-of-code/go/internal"
)

const (
	day  = 2
	year = 2025
)

func main() {
	internal.ProcessAndDisplayResults(day, year, Part1, Part2)
}

func Part1(input string) (int, error) {
	invalidSum := 0
	for _, item := range parseInput(input) {
		for i := item.start; i <= item.end; i++ {
			valueStr := strconv.Itoa(i)
			numLen := len(valueStr)

			if numLen%2 != 0 {
				continue
			}

			partsCount := numLen / 2
			part1Str := valueStr[:partsCount]
			part2Str := valueStr[partsCount:]

			if part1Str == part2Str {
				invalidSum += i
			}
		}
	}

	return invalidSum, nil
}

func Part2(input string) (int, error) {
	invalidSum := 0
	for _, item := range parseInput(input) {
		for i := item.start; i <= item.end; i++ {

			valueStr := strconv.Itoa(i)
			numLen := len(valueStr)
			maxWindow := numLen / 2

			for j := 1; j <= maxWindow; j++ {
				evenSplit := numLen % j == 0
				if !evenSplit {
					continue
				}

				partStr := valueStr[:j]
				fullStr := strings.Repeat(partStr, numLen/j)

				if fullStr == valueStr {
					invalidSum += i
					break
				}
			}
		}
	}

	return invalidSum, nil
}

type part struct {
	start int
	end   int
}

func parseInput(input string) []part {

	items := strings.Split(input, ",")

	var parts []part
	for _, item := range items {

		p := strings.Split(item, "-")

		start, err := strconv.Atoi(p[0])
		if err != nil {
			continue
		}

		end, err := strconv.Atoi(p[1])
		if err != nil {
			continue
		}

		parts = append(parts, part{
			start: start,
			end:   end,
		})
	}

	return parts
}
