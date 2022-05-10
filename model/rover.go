package model

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

func InitRover(planet Planet) *Rover {
	var rover = new(Rover)
	rover.x = 0
	rover.y = 0
	rover.direction = N
	rover.planet = planet

	return rover
}

func (rover *Rover) Move() {
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

func (rover *Rover) TurnRight() {
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

func (rover *Rover) TurnLeft() {
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
