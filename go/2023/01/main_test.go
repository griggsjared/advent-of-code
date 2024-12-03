package main_test

import (
	"testing"

	main "github.com/griggsjared/advent-of-code/go/2023/01"
)

func TestPart1(t *testing.T) {

  input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	sum, err := main.Part1(input)
	if err != nil {
		t.Error(err)
	}

	if sum != 142 {
		t.Errorf("Expected 142, got %d", sum)
	}
}

func TestPart2(t *testing.T) {

  input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
eightwo`

	sum, err := main.Part2(input)
	if err != nil {
		t.Error(err)
	}

	if sum != 363 {
		t.Errorf("Expected 363, got %d", sum)
	}
}


