package rover_test

import (
	"testing"

	"github.com/churmd/wise-mars-rover/domain/rover"
	"github.com/stretchr/testify/assert"
)

func TestNewRover(t *testing.T) {
	expectedRover := rover.Rover{
		X:           1,
		Y:           2,
		Orientation: rover.North,
	}

	actualRover := rover.New(1, 2, rover.North)

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
