package main_test

import (
	"testing"

	main "github.com/griggsjared/advent-of-code/go/2025/02"
)

const input string = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestPart1(t *testing.T) {

	count, err := main.Part1(input)
	if err != nil {
		t.Error(err)
		return
	}

	if count != 1227775554 {
		t.Errorf("Expected 1227775554, got %d", count)
	}

}

func TestPart2(t *testing.T) {

	count, err := main.Part2(input)
	if err != nil {
		t.Error(err)
		return
	}

	if count != 4174379265 {
		t.Errorf("Expected 4174379265, got %d", count)
	}
}
