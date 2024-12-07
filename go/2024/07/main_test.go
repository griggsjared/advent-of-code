package main_test

import (
	"testing"

	main "github.com/griggsjared/advent-of-code/go/2024/07"
)

const input string = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestPart1(t *testing.T) {

	count, err := main.Part1(input)
	if err != nil {
		t.Error(err)
		return
	}

	if count != 3749 {
		t.Errorf("Expected 3749, got %d", count)
	}
}

func TestPart2(t *testing.T) {

	count, err := main.Part2(input)
	if err != nil {
		t.Error(err)
		return
	}

	if count != 11387 {
		t.Errorf("Expected 11387, got %d", count)
	}
}
