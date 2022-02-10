package movement_test

import (
	"testing"

	"github.com/churmd/wise-mars-rover/movement"
	"github.com/churmd/wise-mars-rover/rover"
	"github.com/stretchr/testify/assert"
)

func TestJourneyEmptyCommandsDoNothing(t *testing.T) {
	g := rover.Grid{Width: 4, Height: 8}
	r := rover.New(0, 0, rover.North, g)
	j := movement.Journey{
		Rover:    r,
		Commands: []movement.Command{},
	}

	updatedRover := j.Run()

	assert.Equal(t, r, updatedRover)
}
