package main_test

import (
	"testing"

	main "github.com/griggsjared/advent-of-code/go/2024/05"
)

const input string = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestPart1(t *testing.T) {

	count, err := main.Part1(input)
	if err != nil {
		t.Error(err)
		return
	}

	if count != 143 {
		t.Errorf("Expected 143, got %d", count)
	}
}

func TestPart2(t *testing.T) {

	count, err := main.Part2(input)
	if err != nil {
		t.Error(err)
		return
	}

	if count != 123 {
		t.Errorf("Expected 123, got %d", count)
	}
}
