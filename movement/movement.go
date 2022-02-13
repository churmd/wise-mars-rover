package movement

import "github.com/churmd/wise-mars-rover/rover"

const (
	Left   Command = 'L'
	Right  Command = 'R'
	Foward Command = 'F'
)

type Command rune

type Journey struct {
	Rover    rover.Rover
	Commands []Command
}

func (j Journey) Run() rover.Rover {
	updatedRover := j.Rover
	for _, cmd := range j.Commands {
		switch cmd {
		case Left:
			updatedRover = updatedRover.RotateLeft()
		case Right:
			updatedRover = updatedRover.RotateRight()
		case Foward:
			updatedRover = updatedRover.Foward()
		}
	}

	return updatedRover
}
