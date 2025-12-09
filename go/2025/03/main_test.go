package main_test

import (
	"testing"

	main "github.com/griggsjared/advent-of-code/go/2025/03"
)

const input string = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestPart1(t *testing.T) {

	expect := 357
	value, err := main.Part1(input)
	if err != nil {
		t.Error(err)
		return
	}

	if value != expect {
		t.Errorf("Expected %d, got %d", expect, value)
	}
}

func TestPart2(t *testing.T) {

	expect := 3121910778619
	value, err := main.Part2(input)
	if err != nil {
		t.Error(err)
		return
	}

	if value != expect {
		t.Errorf("Expected %d, got %d", expect, value)
	}
}
