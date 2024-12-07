package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type year uint

func (y year) Valid() bool {
  currentYear := time.Now().Year()
  if int(y) < 2015 || int(y) > currentYear {
    return false
  }
  return true
}

type day uint

func (d day) Valid() bool {
  if int(d) < 1 || int(d) > 24 {
    return false
  }
  return true
}

func main() {

	if len(os.Args) > 1 && os.Args[1] == "--all" {
		startYear := 2015
		currentYear := time.Now().Year()
		for y := startYear; y <= currentYear; y++ {
			for d := 1; d <= 24; d++ {
				runDay(year(y), day(d))
			}
		}
		return
	}

	var y year
	var d day

	fmt.Println("Enter year:")
	fmt.Scanln(&y)

  if !y.Valid() {
    fmt.Println("Invalid year, must be between 2015 and current year.")
    return
  }

	fmt.Println("Enter day:")
	fmt.Scanln(&d)

  if !d.Valid() {
    fmt.Println("Invalid day, must be between 1 and 24.")
    return
  }

	err := runDay(y, d)
	if err != nil {
		fmt.Println(err)
	}
}

func runDay(y year, d day) error {
	pathGo := fmt.Sprintf("./%d/%02d/main.go", y, d)
	pathInput := fmt.Sprintf("./%d/%02d/input.txt", y, d)

	if _, err := os.Stat(pathGo); os.IsNotExist(err) {
		return fmt.Errorf("That day has not been implemented yet.")
	}

	if _, err := os.Stat(pathInput); os.IsNotExist(err) {
		return fmt.Errorf("That day has been implemented, but the input file is missing.")
	}

	cmnd := exec.Command("go", "run", pathGo)
	cmnd.Stdout = os.Stdout
	cmnd.Stderr = os.Stderr
	cmnd.Run()

	return nil
}
