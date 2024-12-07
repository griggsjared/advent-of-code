package internal

import (
	"fmt"
	"os"
	"strings"
	"sync"
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

  str = strings.TrimLeft(str, "\n")
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

	i, err := LoadDayInput(day, year)
	if err != nil {
		fmt.Println(FmtWithColor(Red, err.Error()))
		return
	}

	var wg sync.WaitGroup

	wg.Add(2)

  go runPart(p1, i, 1, &wg)
  go runPart(p2, i, 2, &wg)

	wg.Wait()
}

func runPart(p partFunc, i string, pn int, wg *sync.WaitGroup) {

	var outputLns []string
  defer func() {
    fmt.Println(strings.Join(outputLns, "\n"))
  }()
  
  defer wg.Done()

	sTime := time.Now()
	outputLns = append(outputLns, fmt.Sprintf(FmtWithColor(Green, fmt.Sprintf("Part %d:", pn))))
  eTime := time.Now()
	res, err := p(i)
	if err != nil {
    outputLns = append(outputLns, FmtWithColor(Red, err.Error()))
		return
	}
  outputLns = append(outputLns, FmtWithColor(Cyan, fmt.Sprintf("%d 🌟", res)))
  outputLns = append(outputLns, FmtWithColor(Yellow, "(Time: %v)", eTime.Sub(sTime))) 
}
