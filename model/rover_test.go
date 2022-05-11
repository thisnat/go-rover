package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thisnat/go-rover/model"
)

func TestInitPostionShouldBeNorthZeroZero(t *testing.T) {
	t.Parallel()

	mar := model.InitPlanet(24)
	rover := model.InitRover(*mar)

	expectedX := 0
	expectedY := 0
	expectedDirection := "N"

	actualX := rover.GetX()
	actualY := rover.GetY()
	actualDirection := rover.GetDirecttion()

	assert.Equal(t, expectedX, actualX)
	assert.Equal(t, expectedY, actualY)
	assert.Equal(t, expectedDirection, actualDirection)
}
