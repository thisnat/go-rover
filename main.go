package main

import (
	"bufio"
	"os"

	"github.com/thisnat/go-rover/model"
)

func main() {
	var command []string
	var rover model.Rover

	file, err := os.Open("control.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		command = append(command, scanner.Text())
	}

	rover.ControlByInput(command)
}
