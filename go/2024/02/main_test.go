package main_test

import (
	// "fmt"
	"testing"

	main "github.com/griggsjared/advent-of-code/go/2024/02"
)

const input string = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 4 7 9`

func TestPart1(t *testing.T) {
	count, err := main.Part1(input)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestPart2(t *testing.T) {
	count, err := main.Part2(input)
	if err != nil {
		t.Error(err)
	}

	if count != 4 {
		t.Errorf("Expected 4, got %d", count)
	}
}

