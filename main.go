package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/thisnat/go-rover/model"
)

func main() {
	var text []string
	var mar *model.Planet
	var rover *model.Rover

	file, err := os.Open("control.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for index, line := range text {
		if index == 0 {
			maxEdge, _ := strconv.Atoi(line)

			mar = model.InitPlanet(maxEdge)
			rover = model.InitRover(*mar)

			fmt.Printf("Map size: %d\n", mar.GetEdge())
		} else {
			switch line {
			case "F":
				rover.Move()
			case "L":
				rover.TurnLeft()
			case "R":
				rover.TurnRight()
			default:
				fmt.Println("unknow case")
			}
		}
		fmt.Printf("%s:%d,%d\n", rover.GetDirecttion(), rover.GetX(), rover.GetY())
	}
}
