package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thisnat/go-rover/model"
)

func TestPlanetGetSize(t *testing.T) {
	t.Parallel()

	mar := model.InitPlanet(24)

	expectedSize := 24
	actualSize := mar.GetEdge()

	assert.Equal(t, expectedSize, actualSize)
}
