package main_test

import (
	// "fmt"
	"testing"

	main "github.com/griggsjared/advent-of-code/go/2024/01"
)

func TestPart1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	distance, err := main.Part1(input)
	if err != nil {
		t.Error(err)
	}

	if distance != 11 {
		t.Errorf("Expected 11, got %d", distance)
	}
}

func TestPart2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	similarity, err := main.Part2(input)
	if err != nil {
		t.Error(err)
	}

	if similarity != 31 {
		t.Errorf("Expected 31, got %d", similarity)
	}
}

