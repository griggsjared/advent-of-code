package main_test

import (
	"testing"

	main "github.com/griggsjared/advent-of-code/go/2024/06" 
)

const input string = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPart1(t *testing.T) {

	count, err := main.Part1(input)
	if err != nil {
		t.Error(err)
		return
	}

	if count != 41 {
		t.Errorf("Expected 41, got %d", count)
	}
}

// func TestPart2(t *testing.T) {
//
// 	count, err := main.Part2(input)
// 	if err != nil {
// 		t.Error(err)
//     return
// 	}
//
// 	if count != 0 {
// 		t.Errorf("Expected 0, got %d", count)
// 	}
// }
