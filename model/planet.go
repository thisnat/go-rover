package model

type Planet struct {
	edge int
}

func InitPlanet(max int) *Planet {
	var planet = new(Planet)
	planet.edge = max

	return planet
}

func (planet *Planet) GetEdge() int {
	return planet.edge
}
