package main_test

import (
	"testing"

	main "github.com/griggsjared/advent-of-code/go/2025/01" 
)

const input string = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestPart1(t *testing.T) {

	count, err := main.Part1(input)
	if err != nil {
		t.Error(err)
		return
	}

	if count != 3 {
		t.Errorf("Expected 3, got %d", count)
	}
}

func TestPart2(t *testing.T) {

	count, err := main.Part2(input)
	if err != nil {
		t.Error(err)
    return
	}

	if count != 6 {
		t.Errorf("Expected 6, got %d", count)
	}
}
