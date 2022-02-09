package rover_test

import (
	"testing"

	"github.com/churmd/wise-mars-rover/rover"
	"github.com/stretchr/testify/assert"
)

func TestNewRover(t *testing.T) {
	expectedRover := rover.Rover{
		X:           1,
		Y:           2,
		Orientation: rover.North,
		Grid:        rover.Grid{Width: 4, Height: 8},
	}

	actualRover := rover.New(1, 2, rover.North, rover.Grid{Width: 4, Height: 8})

	assert.Equal(t, expectedRover, actualRover)
}

func TestString(t *testing.T) {
	cases := map[rover.Rover]string{
		{X: 1, Y: 2, Orientation: rover.North}:               "(1, 2, N)",
		{X: 10, Y: 20, Orientation: rover.South, Lost: true}: "(10, 20, S) LOST",
	}

	for r, s := range cases {
		actualString := r.String()
		assert.Equal(t, s, actualString)
	}
}

func TestForwardNorth(t *testing.T) {
	mars := rover.Grid{Width: 5, Height: 5}
	r := rover.New(2, 2, rover.North, mars)

	updatedR := r.Foward()

	assert.Equal(t, 3, updatedR.Y)
}

func TestForwardEast(t *testing.T) {
	mars := rover.Grid{Width: 5, Height: 5}
	r := rover.New(2, 2, rover.East, mars)

	updatedR := r.Foward()

	assert.Equal(t, 3, updatedR.X)
}

func TestForwardSoth(t *testing.T) {
	mars := rover.Grid{Width: 5, Height: 5}
	r := rover.New(2, 2, rover.South, mars)

	updatedR := r.Foward()

	assert.Equal(t, 1, updatedR.Y)
}

func TestForwardWest(t *testing.T) {
	mars := rover.Grid{Width: 5, Height: 5}
	r := rover.New(2, 2, rover.West, mars)

	updatedR := r.Foward()

	assert.Equal(t, 1, updatedR.X)
}

func TestFowardDoesNothingWhenLost(t *testing.T) {
	mars := rover.Grid{Width: 5, Height: 5}
	r := rover.New(2, 2, rover.North, mars)
	r.Lost = true

	updatedR := r.Foward()

	assert.Equal(t, 2, updatedR.Y)
}

func TestForwardSetsLostAndDoesNotMoveWhenGoingOffGrid(t *testing.T) {
	mars := rover.Grid{Width: 5, Height: 5}
	r := rover.New(4, 4, rover.North, mars)

	updatedR := r.Foward()

	assert.Equal(t, r, updatedR)
}
