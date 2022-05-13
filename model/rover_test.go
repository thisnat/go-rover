package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thisnat/go-rover/model"
)

func TestInitPostionShouldBeNorthZeroZero(t *testing.T) {
	t.Parallel()

	var rover model.Rover
	var command []string = []string{"24"}

	rover.ControlByInput(command)

	expectedX := 0
	expectedY := 0
	expectedDirection := "N"

	assert.Equal(t, expectedX, rover.GetX())
	assert.Equal(t, expectedY, rover.GetY())
	assert.Equal(t, expectedDirection, rover.GetDirecttion())
}

func TestGivenTestCaseLastPostionShouldBeWestOneZero(t *testing.T) {
	t.Parallel()

	var rover model.Rover
	var command []string = []string{"24", "R", "F", "L", "F", "L", "L", "F", "R"}

	rover.ControlByInput(command)

	expectedX := 1
	expectedY := 0
	expectedDirection := "W"

	assert.Equal(t, expectedX, rover.GetX())
	assert.Equal(t, expectedY, rover.GetY())
	assert.Equal(t, expectedDirection, rover.GetDirecttion())
}

func TestMoveFowardFromInitPostionShouldBeNorthZeroOne(t *testing.T) {
	t.Parallel()

	var rover model.Rover
	var command []string = []string{"24", "F"}

	rover.ControlByInput(command)

	expectedX := 0
	expectedY := 1
	expectedDirection := "N"

	assert.Equal(t, expectedX, rover.GetX())
	assert.Equal(t, expectedY, rover.GetY())
	assert.Equal(t, expectedDirection, rover.GetDirecttion())
}

func TestTurnLeftFromInitPostionShouldBeWestZeroZero(t *testing.T) {
	t.Parallel()

	var rover model.Rover
	var command []string = []string{"24", "L"}

	rover.ControlByInput(command)

	expectedX := 0
	expectedY := 0
	expectedDirection := "W"

	assert.Equal(t, expectedX, rover.GetX())
	assert.Equal(t, expectedY, rover.GetY())
	assert.Equal(t, expectedDirection, rover.GetDirecttion())
}

func TestTurnRightFromInitPostionShouldBeEastZeroZero(t *testing.T) {
	t.Parallel()

	var rover model.Rover
	var command []string = []string{"24", "R"}

	rover.ControlByInput(command)

	expectedX := 0
	expectedY := 0
	expectedDirection := "E"

	assert.Equal(t, expectedX, rover.GetX())
	assert.Equal(t, expectedY, rover.GetY())
	assert.Equal(t, expectedDirection, rover.GetDirecttion())
}

func TestReachingTheEdgeMustMaintainThePosition(t *testing.T) {
	t.Parallel()

	var rover model.Rover
	var command []string = []string{"2", "F", "F", "F", "F"}

	rover.ControlByInput(command)

	expectedX := 0
	expectedY := 2
	expectedDirection := "N"

	assert.Equal(t, expectedX, rover.GetX())
	assert.Equal(t, expectedY, rover.GetY())
	assert.Equal(t, expectedDirection, rover.GetDirecttion())
}

func TestTurnRightAndMoveShouldIncreaseX(t *testing.T) {
	t.Parallel()

	var rover model.Rover
	var command []string = []string{"24", "R", "F", "F"}

	rover.ControlByInput(command)

	expectedX := 2
	expectedY := 0
	expectedDirection := "E"

	assert.Equal(t, expectedX, rover.GetX())
	assert.Equal(t, expectedY, rover.GetY())
	assert.Equal(t, expectedDirection, rover.GetDirecttion())
}

func TestTurnLeftAndMoveFromInitPostionShouldBeWestZeroZero(t *testing.T) {
	t.Parallel()

	var rover model.Rover
	var command []string = []string{"24", "L", "F", "F"}

	rover.ControlByInput(command)

	expectedX := 0
	expectedY := 0
	expectedDirection := "W"

	assert.Equal(t, expectedX, rover.GetX())
	assert.Equal(t, expectedY, rover.GetY())
	assert.Equal(t, expectedDirection, rover.GetDirecttion())
}

func TestTurnRightAndMoveThenTurnLeftTwiceAndMoveShouldBeWestZeroZero(t *testing.T) {
	t.Parallel()

	var rover model.Rover
	var command []string = []string{"24", "R", "F", "L", "L", "F"}

	rover.ControlByInput(command)

	expectedX := 0
	expectedY := 0
	expectedDirection := "W"

	assert.Equal(t, expectedX, rover.GetX())
	assert.Equal(t, expectedY, rover.GetY())
	assert.Equal(t, expectedDirection, rover.GetDirecttion())
}

func TestTurnRighFourTimeShouldBeNorthZeroZero(t *testing.T) {
	t.Parallel()

	var rover model.Rover
	var command []string = []string{"24", "R", "R", "R", "R"}

	rover.ControlByInput(command)

	expectedX := 0
	expectedY := 0
	expectedDirection := "N"

	assert.Equal(t, expectedX, rover.GetX())
	assert.Equal(t, expectedY, rover.GetY())
	assert.Equal(t, expectedDirection, rover.GetDirecttion())
}

func TestTurnLeftFourTimeShouldBeNorthZeroZero(t *testing.T) {
	t.Parallel()

	var rover model.Rover
	var command []string = []string{"24", "L", "L", "L", "L"}

	rover.ControlByInput(command)

	expectedX := 0
	expectedY := 0
	expectedDirection := "N"

	assert.Equal(t, expectedX, rover.GetX())
	assert.Equal(t, expectedY, rover.GetY())
	assert.Equal(t, expectedDirection, rover.GetDirecttion())
}

func TestRoverShouldNotResponseUnknowCommand(t *testing.T) {
	t.Parallel()

	var rover model.Rover
	var command []string = []string{"24", "Q"}

	rover.ControlByInput(command)

	expectedX := 0
	expectedY := 0
	expectedDirection := "N"

	assert.Equal(t, expectedX, rover.GetX())
	assert.Equal(t, expectedY, rover.GetY())
	assert.Equal(t, expectedDirection, rover.GetDirecttion())
}

func TestPanicWhenFirstIndexIsNotANumber(t *testing.T) {
	t.Parallel()

	var rover model.Rover
	var command []string = []string{"test", "F"}

	assert.Panics(t, func() { rover.ControlByInput(command) }, "First index must be a numer")
}
