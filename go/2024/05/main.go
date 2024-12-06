package main

import (
	"strconv"
	"strings"

	"github.com/griggsjared/advent-of-code/go/internal"
)

const (
	day  = 5
	year = 2024
)

func main() {
	internal.ProcessAndDisplayResults(day, year, Part1, Part2)
}

func Part1(input string) (int, error) {

	sets, rules := getRulesAndSetsFromInput(input)
	validSets, _ := getValidAndInvalidSets(sets, rules)

	sum := getSumForSetSets(validSets)
	return sum, nil
}

func Part2(input string) (int, error) {

	sets, rules := getRulesAndSetsFromInput(input)
	_, invalidSets := getValidAndInvalidSets(sets, rules)

	//invalid sets are not in the correct order
	//we need to loop through these sets and correct them according to the rules

	var correctedSets [][]string
	for _, s := range invalidSets {

		//while the set is not in the correct order, try to shuffle the the set using the rules pretty much the same way we check if the set is in the correct order
		for !setIsCorrectOrder(s, rules) {

			for _, rule := range rules {
				num1Index := sliceIndex(s, rule[0])
				num2Index := sliceIndex(s, rule[1])

				//if either of the nubmers in the rule are not in the set, then the rule does not apply
				if num1Index == -1 || num2Index == -1 {
					continue //continue to the next rule, this one did not apply
				}

				//the first index must be less than the second index
				if num1Index > num2Index {
					//swap the indexes
					s[num1Index], s[num2Index] = s[num2Index], s[num1Index]
				}
			}
		}

		correctedSets = append(correctedSets, s)
	}

	sum := getSumForSetSets(invalidSets)
	return sum, nil
}

func sliceIndex(haystack []string, needle string) int {
	for i, item := range haystack {
		if item == needle {
			return i
		}
	}
	return -1
}

func getValidAndInvalidSets(sets [][]string, rules [][]string) (validSets [][]string, invalidSets [][]string) {

	var vs [][]string
	var ivs [][]string

	for _, s := range sets {
		correctOrder := setIsCorrectOrder(s, rules)
		if correctOrder {
			vs = append(vs, s)
		} else {
			ivs = append(ivs, s)
		}
	}

	return vs, ivs
}

func setIsCorrectOrder(s []string, rules [][]string) bool {

	for _, rule := range rules {

		num1Index := sliceIndex(s, rule[0])
		num2Index := sliceIndex(s, rule[1])

		//if either of the nubmers in the rule are not in the set, then the rule does not apply
		if num1Index == -1 || num2Index == -1 {
			continue //continue to the next rule, this one did not apply
		}

		//the first index must be less than the second index
		if num1Index > num2Index {
			return false
		}
	}

	return true
}

func getRulesAndSetsFromInput(i string) ([][]string, [][]string) {

	//the input has 2 groups of data seperated by a blank line (\n)
	rows := strings.Split(i, "\n")

	var data1 []string
	var data2 []string
	for _, row := range rows {
		//if the row as a | it goes into data1
		//blank line gets rejected
		//all other rows go into data2
		if strings.Contains(row, "|") {
			data1 = append(data1, row)
		} else if row != "" {
			data2 = append(data2, row)
		}
	}

	//rules will be 2 numbers (as strings) in a slice of a length of 2
	var rules [][]string
	for _, row := range data1 {
		rule := strings.Split(row, "|")
		rules = append(rules, rule)
	}

	//sets will be a slice of numbers (as strings) in a slice of variable length
	var sets [][]string
	for _, row := range data2 {
		set := strings.Split(row, ",")
		sets = append(sets, set)
	}

	return sets, rules
}

func getSumForSetSets(sets [][]string) int {
	var values []int
	for _, s := range sets {
		//the middle index in the set is the value we want
		mIndex := len(s) / 2

		vInt, err := strconv.Atoi(s[mIndex])
		if err != nil {
			values = append(values, 0)
		}

		values = append(values, vInt)
	}

	//sum the values of the valid sets
	sum := 0
	for _, v := range values {
		sum += v
	}

	return sum
}
