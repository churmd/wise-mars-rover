package movement_test

import (
	"testing"

	"github.com/churmd/wise-mars-rover/movement"
	"github.com/churmd/wise-mars-rover/rover"
	"github.com/stretchr/testify/assert"
)

func TestJourneyNew(t *testing.T) {
	for i, s := range inputStrings {
		j := movement.Parse(s)

		assert.Equal(t, journeys[i], j)
	}
}

func TestJourneyRun(t *testing.T) {
	for i, j := range journeys {
		r := j.Run()

		assert.Equal(t, expectedRovers[i], r)
	}
}

var inputStrings = []string {
	"(2, 3, E) LFRFF",
"(0, 2, N) FFLFRFF",
"(2, 3, N) FLLFR",
"(1, 0, S) FFRLF",
}

// Exmaples define the grid 4x8 and allow 4 as final X coordinate, must be inclusive
var g = rover.Grid{Width: 5, Height: 9}

// (2, 3, E) LFRFF
// (0, 2, N) FFLFRFF
// (2, 3, N) FLLFR
// (1, 0, S) FFRLF
var journeys = []movement.Journey{
	{Rover: rover.New(0, 0, rover.North, g), Commands: []movement.Command{}},
	{Rover: rover.New(2, 3, rover.East, g), Commands: []movement.Command{movement.Left, movement.Foward, movement.Right, movement.Foward, movement.Foward}},
	{Rover: rover.New(0, 2, rover.North, g), Commands: []movement.Command{movement.Foward, movement.Foward, movement.Left, movement.Foward, movement.Right, movement.Foward, movement.Foward}},
	{Rover: rover.New(2, 3, rover.North, g), Commands: []movement.Command{movement.Foward, movement.Left, movement.Left, movement.Foward, movement.Right}},
	{Rover: rover.New(1, 0, rover.South, g), Commands: []movement.Command{movement.Foward, movement.Foward, movement.Right, movement.Left, movement.Foward}},
}

// > (4, 4, E)
// > (0, 4, W) LOST
// > (2, 3, W)
// > (1, 0, S) LOST
var expectedRovers = []rover.Rover{
	rover.New(0, 0, rover.North, g),
	rover.New(4, 4, rover.East, g),
	{
		X:           0,
		Y:           4,
		Orientation: rover.West,
		Grid:        g,
		Lost:        true,
	},
	rover.New(2, 3, rover.West, g),
	{
		X:           1,
		Y:           0,
		Orientation: rover.South,
		Grid:        g,
		Lost:        true,
	},
}
