package internal

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func LoadDayInput(day int, year int) (string, error) {
	fpath := fmt.Sprintf("./%d/%02d/input.txt", year, day)

	i, err := os.ReadFile(fpath)
	if err != nil {
		return "", err
	}

	if len(i) == 0 {
		return "", fmt.Errorf("input file is empty")
	}

	str := string(i)

	//trim the input to remove empty lines
	str = strings.TrimRight(str, "\n")

	return str, nil
}

type FmtColor int

const (
	Black   FmtColor = 30
	Red     FmtColor = 31
	Green   FmtColor = 32
	Yellow  FmtColor = 33
	Blue    FmtColor = 34
	Magenta FmtColor = 35
	Cyan    FmtColor = 36
	White   FmtColor = 37
)

func (c FmtColor) Validate() bool {
	switch c {
	case Black, Red, Green, Yellow, Blue, Magenta, Cyan, White:
		return true
	default:
		return false
	}
}

func FmtWithColor(color FmtColor, format string, args ...any) string {
	if !color.Validate() {
		panic("Invalid color")
	}

	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", int(color), fmt.Sprintf(format, args...))
}

// partFunc is a function that takes a string and returns an int and an error,
// this should be the standard format for functions that solve part 1 and 2 of advent of code problems
type partFunc func(i string) (int, error)

func ProcessAndDisplayResults(day int, year int, p1 partFunc, p2 partFunc) {
	fmt.Println(FmtWithColor(Blue, "Advent of Code %d - Day %d", year, day))

	input, err := LoadDayInput(day, year)
	if err != nil {
		fmt.Println(FmtWithColor(Red, err.Error()))
		return
	}

	sTime := time.Now()
	fmt.Println(FmtWithColor(Green, "Part 1:"))
	result1, err := p1(input)
	if err != nil {
		fmt.Println(FmtWithColor(Red, err.Error()))
		return
  }
	eTime := time.Now()
	fmt.Println(FmtWithColor(Cyan, fmt.Sprintf("%d 🌟", result1)))
	fmt.Println(FmtWithColor(Yellow, "(Time: %v)", eTime.Sub(sTime)))

	sTime = time.Now()
	fmt.Println(FmtWithColor(Green, "Part 2:"))
	result2, err := p2(input)
	if err != nil {
		fmt.Println(FmtWithColor(Red, err.Error()))
		return
	}
	eTime = time.Now()
	fmt.Println(FmtWithColor(Cyan, fmt.Sprintf("%d 🌟", result2)))
	fmt.Println(FmtWithColor(Yellow, "(Time: %v)", eTime.Sub(sTime)))
}
