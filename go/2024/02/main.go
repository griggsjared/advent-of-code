package main

import (
	"strconv"
	"strings"

	"github.com/griggsjared/advent-of-code/go/internal"
)

const (
	day  = 2
	year = 2024
)

func main() {
	internal.ProcessAndDisplayResults(day, year, Part1, Part2)
}

func Part1(i string) (int, error) {

	rows := strings.Split(i, "\n")

	safe := 0
	for _, row := range rows {
		rowValues := strings.Split(row, " ")
		if rowIsSafe(rowValues) {
			safe++
		}
	}

	return safe, nil
}

func Part2(i string) (int, error) {

	rows := strings.Split(i, "\n")

	safe := 0

	for _, row := range rows {
		rowValues := strings.Split(row, " ")

		rowSafe := rowIsSafe(rowValues)
		if !rowSafe {
			//we will remove a single value from the row and check if it is safe with that value removed,
			//if it is, we will increment the safe count and move to the next row
			for i := range rowValues {

				//copy the row so we can remove a value
				newRow := make([]string, len(rowValues))
				copy(newRow, rowValues)

				//remove the value at index i
				newRow = append(newRow[:i], newRow[i+1:]...)

				//check if the new row is safe
				if rowIsSafe(newRow) {
					safe++
					break
				}
			}
		}

		if rowSafe {
			safe++
		}
	}

	return safe, nil
}

func rowIsSafe(row []string) bool {
	isSafe := true
	incOrDec := ""
	for i, value := range row {

		//skip the first value since we have nothing to compare it to
		if i == 0 {
			continue
		}

		// covnerting the input values to numbers
		num, err := strconv.Atoi(value)
		if err != nil {
			num = 0
		}

		pnum, err := strconv.Atoi(row[i-1])
		if err != nil {
			pnum = 0
		}

		//if the numbers are the same they are not safe
		if num == pnum {
			isSafe = false
			break
		}

		//check to see if prev to current is increasing or decreasing, and matches the current
		valDir := "inc"
		if num < pnum {
			valDir = "dec"
		}

		//setting the direction of the first two values
		if incOrDec == "" {
			incOrDec = valDir
		}

		//if the direction changes, it is not safe
		if incOrDec != valDir {
			isSafe = false
			break
		}

		//difference of 3 or more is not safe
		diff := num - pnum
		if diff > 3 || diff < -3 {
			isSafe = false
			break
		}
	}

	return isSafe
}
