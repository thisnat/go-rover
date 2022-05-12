package model

import (
	"fmt"
	"strconv"
)

type Rover struct {
	x         int
	y         int
	direction string
	planet    Planet
}

const (
	N = "N"
	S = "S"
	W = "W"
	E = "E"
)

func (rover *Rover) ControlByInput(command []string) {
	rover.x = 0
	rover.y = 0
	rover.direction = N

	for index, line := range command {
		if index == 0 {
			maxEdge, err := strconv.Atoi(line)
			if err != nil {
				panic("First index should be a numer")
			}

			mar := NewPlanet(maxEdge)
			rover.planet = *mar

			fmt.Printf("Map size: %d\n", rover.planet.GetEdge())
		} else {
			switch line {
			case "F":
				rover.move()
			case "L":
				rover.turnLeft()
			case "R":
				rover.turnRight()
			default:
				fmt.Println("unknow case")
			}
		}
		fmt.Printf("%s:%d,%d\n", rover.direction, rover.x, rover.y)
	}
}

func (rover *Rover) move() {
	if rover.direction == N || rover.direction == S {
		if rover.direction == N {
			if rover.y+1 <= rover.planet.GetEdge() {
				rover.y += 1
			}
		} else if rover.direction == S {
			if rover.y-1 >= 0 {
				rover.y -= 1
			}
		}
	} else if rover.direction == W || rover.direction == E {
		if rover.direction == E {
			if rover.x+1 <= rover.planet.GetEdge() {
				rover.x += 1
			}
		} else if rover.direction == W {
			if rover.x-1 >= 0 {
				rover.x -= 1
			}
		}
	}
}

func (rover *Rover) turnRight() {
	switch rover.direction {
	case N:
		rover.direction = E
	case S:
		rover.direction = W
	case W:
		rover.direction = N
	case E:
		rover.direction = S
	}
}

func (rover *Rover) turnLeft() {
	switch rover.direction {
	case N:
		rover.direction = W
	case S:
		rover.direction = E
	case W:
		rover.direction = S
	case E:
		rover.direction = N
	}
}

func (rover *Rover) GetX() int {
	return rover.x
}

func (rover *Rover) GetY() int {
	return rover.y
}

func (rover *Rover) GetDirecttion() string {
	return rover.direction
}
