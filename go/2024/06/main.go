package main

import (
	"fmt"
	"strings"

	"github.com/griggsjared/advent-of-code/go/internal"
)

const (
	day  = 6    // Change this to the day you are working on
	year = 2024 // Change this to the year you are working on
)

func main() {
	internal.ProcessAndDisplayResults(day, year, Part1, Part2)
}

func Part1(input string) (int, error) {

	grid := strings.Split(input, "\n")

	cCount := len(grid[0])
	rCount := len(grid)

	//directionMap is a map keyed by the direction character and valued by that the character can move.
	directionMap := map[string][]int{
		"^": {0, -1},
		">": {1, 0},
		"v": {0, 1},
		"<": {-1, 0},
	}

	//nextDirection is map to change the direction of the character when it hits an obstacle
	nextDirection := map[string]string{
		"^": ">",
		">": "v",
		"v": "<",
		"<": "^",
	}

	currentPos := []int{0, 0}
	currentChar := "^"

	//look through the grid and find the starting location of the node arrow, it might be and of the four directions characters
	for i := 0; i < rCount; i++ {
		for j := 0; j < cCount; j++ {
			if grid[i][j] == '^' || grid[i][j] == 'v' || grid[i][j] == '<' || grid[i][j] == '>' {
				currentPos = []int{j, i}
				currentChar = string(grid[i][j])
				break
			}
		}
	}

	if currentChar == "" {
		return 0, fmt.Errorf("no starting location found")
	}

	//from the starting location, move the character through the grid til we reach the out of bounds or an obstacle,
	//if the character hits an obstacle, change the direction of the character according to the directionMap and continue moving
	//as we move map the locations we have been to so we can reference them later
	visited := make(map[string]bool)

	for {

		visitedKey := fmt.Sprintf("%d-%d", currentPos[0], currentPos[1])

		visited[visitedKey] = true

		nextPos := []int{currentPos[0] + directionMap[currentChar][0], currentPos[1] + directionMap[currentChar][1]}

		//if the next position is out of bounds, break the loop
		if nextPos[0] < 0 || nextPos[0] >= cCount || nextPos[1] < 0 || nextPos[1] >= rCount {
			break
		}

		//if the next position is an obstacle, change the direction of the character and continue
		if grid[nextPos[1]][nextPos[0]] == '#' {
			currentChar = nextDirection[currentChar]
			continue
		}

		//if the next position is a space, move the character to the next position
		currentPos = nextPos
	}

	visitedCount := len(visited)

	return visitedCount, nil
}

func Part2(input string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}
