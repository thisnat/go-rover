package model

type Planet struct {
	edge int
}

func NewPlanet(max int) *Planet {
	planet := Planet{max}

	return &planet
}

func (planet *Planet) GetEdge() int {
	return planet.edge
}
