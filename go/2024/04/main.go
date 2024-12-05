package main

import (
	"strings"

	"github.com/griggsjared/advent-of-code/go/internal"
)

const (
	day  = 4
	year = 2024
)

func main() {
	internal.ProcessAndDisplayResults(day, year, Part1, Part2)
}

func Part1(input string) (int, error) {

	grid := strings.Split(input, "\n")

	cCount := len(grid[0])
	rCount := len(grid)

	count := 0
	for i := 0; i < rCount; i++ {
		for j := 0; j < cCount; j++ {
			count += checkMas(i, j, grid)
		}
	}

	return count, nil
}

func Part2(input string) (int, error) {

	grid := strings.Split(input, "\n")

	cCount := len(grid[0])
	rCount := len(grid)

	count := 0
	for i := 0; i < rCount; i++ {
		for j := 0; j < cCount; j++ {
			if checkXMas(i, j, grid) {
				count++
			}
		}
	}

	return count, nil
}

// checkMas checks all 8 directions for X, M, then A, then S
func checkMas(x, y int, grid []string) int {

	directions := [][]int{
		{0, 1},   // right
		{1, 0},   // down
		{1, 1},   // down right
		{0, -1},  // left
		{-1, 0},  // up
		{-1, -1}, // up left
		{1, -1},  // down left
		{-1, 1},  // up right
	}

	count := 0

	xmas := []string{
		"X",
		"M",
		"A",
		"S",
	}

	for _, dir := range directions {
		found := 0

		//start off searching at the origin
		x1 := x
		y1 := y
		for _, char := range xmas {
			//cannot check outside of our bounds
			if x1 < 0 || x1 >= len(grid) || y1 < 0 || y1 >= len(grid[0]) {
				break
			}

			//if not the character we are looking for, break
			if string(grid[x1][y1]) != char {
				break
			}

			//it must be the character we are looking for, incerement found
			//and change x and y to the next direction
			found++
			x1 = dir[0] + x1
			y1 = dir[1] + y1
		}

		if found == len(xmas) {
			count++
		}
	}

	return count
}

func checkXMas(x, y int, grid []string) bool {

	directions := [][]int{
    {-1, 1},  // up right
    {-1, -1}, // up left
	}

	oChar := string(grid[x][y])

	if oChar != "A" {
		return false
	}

	for _, dir := range directions {

		//the direction we are checking
		x1 := x + dir[0]
		y1 := y + dir[1]

		//the opposite of the direction we are checking
		x2 := x + (dir[0] * -1)
		y2 := y + (dir[1] * -1)

		//is the direction within bounds?
		if x1 < 0 || x1 >= len(grid) || y1 < 0 || y1 >= len(grid[0]) {
      return false
		}

		//is the opposite of the direction within bounds?
		if x2 < 0 || x2 >= len(grid) || y2 < 0 || y2 >= len(grid[0]) {
      return false
		}

		dChar := string(grid[x1][y1])

		if dChar != "S" && dChar != "M" {
			return false
		}

		if dChar == "S" && string(grid[x2][y2]) != "M" {
			return false
		}

		if dChar == "M" && string(grid[x2][y2]) != "S" {
			return false
		}
  }

	return true
}
