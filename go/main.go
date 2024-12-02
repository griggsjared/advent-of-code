package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	var year int
	var day int

	fmt.Println("Enter year:")
	fmt.Scanln(&year)

	fmt.Println("Enter day:")
	fmt.Scanln(&day)

	pathGo := fmt.Sprintf("./%d/%02d/main.go", year, day)
	pathInput := fmt.Sprintf("./%d/%02d/input.txt", year, day)

	if _, err := os.Stat(pathGo); os.IsNotExist(err) {
		fmt.Println("That day has not been implemented yet.")
		return
	}

	if _, err := os.Stat(pathInput); os.IsNotExist(err) {
		fmt.Println("That day has been implemented, but the input file is missing.")
		return
	}

	cmnd := exec.Command("go", "run", pathGo)
	cmnd.Stdout = os.Stdout
	cmnd.Stderr = os.Stderr
	cmnd.Run()
}
