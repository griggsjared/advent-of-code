package main_test

import (
	"testing"

	main "github.com/griggsjared/advent-of-code/go/2024/04"
)

const input string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPart1(t *testing.T) {

	count, err := main.Part1(input)
	if err != nil {
		t.Error(err)
	}

	if count != 18 {
		t.Errorf("Expected 18, got %d", count)
	}
}

func TestPart2(t *testing.T) {

  count, err := main.Part2(input)
	if err != nil {
		t.Error(err)
	}

	if count != 9 {
		t.Errorf("Expected 9, got %d", count)
	}
}
