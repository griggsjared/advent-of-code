package main_test

import (
	"testing"

	main "github.com/griggsjared/advent-of-code/go/starter" // Change to the day you are working on (e.g. 2024/01 instead of starter)
)

const input string = ``

func TestPart1(t *testing.T) {

	count, err := main.Part1(input)
	if err != nil {
		t.Error(err)
		return
	}

	if count != 0 {
		t.Errorf("Expected 0, got %d", count)
	}
}

func TestPart2(t *testing.T) {

	count, err := main.Part2(input)
	if err != nil {
		t.Error(err)
    return
	}

	if count != 0 {
		t.Errorf("Expected 0, got %d", count)
	}
}
