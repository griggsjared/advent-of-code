package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/griggsjared/advent-of-code/go/internal"
)

const (
	day  = 1
	year = 2024
)

func main() {
	internal.ProcessAndDisplayResults(day, year, Part1, Part2)
}

func parseLists(i string) ([]int, []int) {
	//input has 2 columns of lists in random order, we need to create 2 lists from the input
	//Example input:
	// 3   4
	// 4   3
	// 2   5
	// 1   3
	// 3   9
	// 3   3

	//split the input into a slice of strings on the newline character
	mainList := strings.Split(i, "\n")

	var list1 []int
	var list2 []int

	for _, v := range mainList {
		//there are 3 spaces between the 2 numbers, so we can split on the space character
		split := strings.Split(v, "   ")

		//convert the strings to integer
		num1, _ := strconv.Atoi(split[0])
		num2, _ := strconv.Atoi(split[1])

		//append the integers to the lists so we can sort them and then compare them
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	//the loops must be the same size
	if len(list1) != len(list2) {
		panic("Lists are not the same size")
	}

	return list1, list2
}

func Part1(i string) (int, error) {

	list1, list2 := parseLists(i)

	//sort each list lowest to highest
	slices.Sort(list1)
	slices.Sort(list2)

	distance := 0
	for i := 0; i < len(list1); i++ {
		if list1[i] == list2[i] {
			continue
		}

		if list1[i] > list2[i] {
			distance += list1[i] - list2[i]
		} else {
			distance += list2[i] - list1[i]
		}
	}

	return distance, nil
}

func Part2(i string) (int, error) {

	list1, list2 := parseLists(i)

	similarity := 0
	for _, num := range list1 {

		//counting the number of times the subject number (num) appears in the second list
		count := 0
		for _, num2 := range list2 {
			if num == num2 {
				count++
			}
		}

		if count == 0 {
			continue
		}

		//similarity for this number is the count of the number is the second list time the subject number
		//add them to the total similarity
		similarity += count * num
	}

	return similarity, nil
}
