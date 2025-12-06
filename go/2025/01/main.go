package main

import (
	"strconv"
	"strings"

	"github.com/griggsjared/advent-of-code/go/internal"
)

const (
	day  = 1    // Change this to the day you are working on
	year = 2025 // Change this to the year you are working on
)

func main() {
	internal.ProcessAndDisplayResults(day, year, Part1, Part2)
}

func Part1(input string) (int, error) {
	stopsOnZero, _ := zeroStats(input, 50)
	return stopsOnZero, nil
}

func Part2(input string) (int, error) {
	_, numberZeros := zeroStats(input, 50)
	return numberZeros, nil
}

type step struct {
	direction byte
	distance  int
}

func parseSteps(input string) []step {
	items := strings.Split(input, "\n")

	var steps []step
	for _, item := range items {
		// left or right depeding on the first character
		direction := item[0]
		distance, err := strconv.Atoi(item[1:])
		if err != nil {
			continue
		}

		steps = append(steps, step{
			direction: direction,
			distance:  distance,
		})
	}

	return steps
}

func zeroStats(input string, start int) (int, int) {
	position := start
	numberZeros := 0
	stopsOnZero := 0
	for _, step := range parseSteps(input) {

		var distance = step.distance

		for distance >= 100 {
			distance = distance - 100
			numberZeros++
		}

		var next int
		if step.direction == 'L' {
			next = position - distance
			if next < 0 {
				next = 100 + next
				numberZeros++
			}
		}

		if step.direction == 'R' {
			next = position + distance
			if next > 99 {
				next = next - 100
				numberZeros++
			}
		}

		if next == 0 {
			stopsOnZero++
		}

		position = next

	}

	return stopsOnZero, numberZeros
}
