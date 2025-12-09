package main_test

import (
	"testing"

	main "github.com/griggsjared/advent-of-code/go/year/day"
)

const input string = ``

func TestPart1(t *testing.T) {

	expected := 0;
	count, err := main.Part1(input)
	if err != nil {
		t.Error(err)
		return
	}

	if count != expected {
		t.Errorf("Expected %d, got %d", expected, count)
	}
}

func TestPart2(t *testing.T) {

	expected := 0;
	count, err := main.Part2(input)
	if err != nil {
		t.Error(err)
		return
	}

	if count != expected {
		t.Errorf("Expected %d, got %d", expected, count)
	}
}
