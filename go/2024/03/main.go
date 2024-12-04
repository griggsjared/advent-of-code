package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/griggsjared/advent-of-code/go/internal"
)

const (
	day  = 3
	year = 2024
)

func main() {
	internal.ProcessAndDisplayResults(day, year, Part1, Part2)
}

func Part1(input string) (int, error) {
	rows := strings.Split(input, "\n")
	return sumMulRows(rows), nil
}

func Part2(input string) (int, error) {
	//input with be lines of "corrupted" text strings,
	//this time there will be do() and don't() that will act like on off switches to prcess the proceeding mul(#,#) substrings,
	//the start of the string always starts with do(), all of mul() up to the don't() are active and will be processed, any mul() after the don't() will not be processed up til the next do()
	inDo := true //start is active
  row := ""
	for i := range input {
		if strings.HasPrefix(input[i:], "don't()") {
			inDo = false
		} else if strings.HasPrefix(input[i:], "do()") {
			inDo = true
		}
		if inDo {
			row += string(input[i])
		}
	}
  rows := make([]string, 0)
  rows = append(rows, row)

	return sumMulRows(rows), nil
}

func sumMulRows(rows []string) int {
	sum := 0
	for _, row := range rows {
		muls := extractMuls(row)
		for _, mul := range muls {
			mulVal, err := processMul(mul)
			if err != nil {
				panic(err)
			}
			sum += mulVal
		}
	}
	return sum
}

func extractMuls(s string) []string {
	//we can use regex to find the mul(#,#) substrings
	//we will return a slice of strings that contain the mul(#,#) substrings
	r, _ := regexp.Compile("mul\\(\\d+,\\d+\\)")

	muls := make([]string, 0)
	muls = r.FindAllString(s, -1)
	//find all the mul(#,#) substrings
	return muls
}

func processMul(mul string) (int, error) {

	//make sure the mul string is in the correct format
	//example: mul(2,4)
	r, _ := regexp.Compile("mul\\(\\d+,\\d+\\)")
	if !r.MatchString(mul) {
		return 0, fmt.Errorf("Invalid mul string: %s", mul)
	}

	//extract the digits out of the mul string
	r2, _ := regexp.Compile("\\d+")
	digits := r2.FindAllString(mul, -1)

	//multiply the digits together
	num1, _ := strconv.Atoi(digits[0])
	num2, _ := strconv.Atoi(digits[1])

	return num1 * num2, nil
}
