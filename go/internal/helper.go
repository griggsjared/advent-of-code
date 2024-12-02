package internal

import (
	"fmt"
	"os"
	"strings"
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

func DisplayDayHeader(day int, year int) {
  fmt.Printf("Advent of Code %d - Day %d\n", year, day)
}
