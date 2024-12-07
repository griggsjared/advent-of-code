package main

import (
	"strconv"
	"strings"

	"github.com/griggsjared/advent-of-code/go/internal"
)

const (
	day  = 7
	year = 2024
)

func main() {
	internal.ProcessAndDisplayResults(day, year, Part1, Part2)
}

func Part1(input string) (int, error) {

	rows := strings.Split(input, "\n")

	sum := 0
	for _, row := range rows {
		parts := strings.Split(row, ": ")
		answer, _ := strconv.Atoi(parts[0])

		var nums []int
		for _, s := range strings.Split(parts[1], " ") {
			s = strings.TrimSpace(s)
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}

		if areNumbersValidForAnswer(nums, answer, false) {
			sum += answer
		}
	}

	return sum, nil
}

func Part2(input string) (int, error) {
	rows := strings.Split(input, "\n")

	sum := 0
	for _, row := range rows {
		parts := strings.Split(row, ": ")
		answer, _ := strconv.Atoi(parts[0])

		var nums []int
		for _, s := range strings.Split(parts[1], " ") {
			s = strings.TrimSpace(s)
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}

		if areNumbersValidForAnswer(nums, answer, true) {
			sum += answer
		}
	}

	return sum, nil
}

// areNumbersValidForAnswer will return true if the numbers can be added or multiplied to equal the answer.
// If the total is equal to the answer, it will return true.
func areNumbersValidForAnswer(nums []int, answer int, concat bool) bool {

	if len(nums) < 2 {
		return nums[0] == answer
	}

	var callback func(pos int, current int) bool
	callback = func(pos int, current int) bool {

    //base case once we have gone through all the numbers 
    //if that number is equal to the answer, return true
		if pos == len(nums) {
			return current == answer
		}

		//addition
		if callback(pos+1, current+nums[pos]) {
			return true
		}

		//multiplication
		if callback(pos+1, current*nums[pos]) {
			return true
		}

		if concat {
			//concatenation
      currentStr := strconv.Itoa(current)
      nextStr := strconv.Itoa(nums[pos])
      //conact the two strings together and change back to an int
      num, _ := strconv.Atoi(currentStr + nextStr)
      if callback(pos+1, num) {
        return true
      }
		}

		return false
	}

	return callback(1, nums[0])
}
