package main_test

import (
	"testing"

	main "github.com/griggsjared/advent-of-code/go/2024/03"
)


func TestPart1(t *testing.T) {

  input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	count, err := main.Part1(input)
	if err != nil {
		t.Error(err)
	}

	if count != 161 {
		t.Errorf("Expected 161, got %d", count)
	}
}

func TestPart2(t *testing.T) {

  input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	count, err := main.Part2(input)
	if err != nil {
		t.Error(err)
	}

	if count != 48 {
		t.Errorf("Expected 48, got %d", count)
	}
}

