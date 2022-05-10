package main

import (
	"fmt"

	"github.com/thisnat/go-rover/model"
)

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func main() {
	// data, err := os.ReadFile("control.txt")
	// check(err)
	// fmt.Print(string(data))

	mar := model.InitPlanet(24)
	rover := model.InitRover(*mar)

	rover.TurnRight()
	rover.Move()
	rover.TurnRight()
	rover.Move()
	rover.TurnRight()
	rover.TurnRight()
	rover.Move()
	rover.TurnLeft()

	fmt.Println(rover)
}
