package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/griggsjared/advent-of-code/go/internal"
)

const (
	day  = 1
	year = 2023
)

func main() {
	internal.ProcessAndDisplayResults(day, year, Part1, Part2)
}

func Part1(i string) (int, error) {

	rows := strings.Split(i, "\n")

	return sumForRows(rows), nil
}

func Part2(i string) (int, error) {

	wordNumMap := make(map[string]string)
	wordNumMap["one"] = "1"
	wordNumMap["two"] = "2"
	wordNumMap["three"] = "3"
	wordNumMap["four"] = "4"
	wordNumMap["five"] = "5"
	wordNumMap["six"] = "6"
	wordNumMap["seven"] = "7"
	wordNumMap["eight"] = "8"
	wordNumMap["nine"] = "9"
	wordNumMap["1"] = "1"
	wordNumMap["2"] = "2"
	wordNumMap["3"] = "3"
	wordNumMap["4"] = "4"
	wordNumMap["5"] = "5"
	wordNumMap["6"] = "6"
	wordNumMap["7"] = "7"
	wordNumMap["8"] = "8"
	wordNumMap["9"] = "9"

	rows := strings.Split(i, "\n")
	newRows := make([]string, 0)
	for _, row := range rows {
		newRow := ""
		for i := range row {
			//the string is the row up to to i places
			subStr := row[i:]

			//for each key in the map, see if the subStr starts with the key
			for k, num := range wordNumMap {
				if strings.HasPrefix(subStr, k) {
					newRow += num
					break
				}
			}
		}

		newRows = append(newRows, newRow)
	}

	return sumForRows(newRows), nil
}

func sumForRows(rows []string) int {
	//split each row into a slice of characters, then filter out any non-numeric characters
	nums := make([][]int, 0)
	for _, row := range rows {
		split := strings.Split(row, "")

		//filter out any non-numeric characters
		var rowNums []int
		for _, v := range split {
			num, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			rowNums = append(rowNums, num)
		}
		nums = append(nums, rowNums)
	}

	// the first and last index will be a number, se want to concatenate them and them add them all together
	sum := 0
	for _, row := range nums {
		first := row[0]
		last := row[len(row)-1]

		num, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
		if err != nil {
			num = 0
		}

		sum += num
	}

	return sum
}
