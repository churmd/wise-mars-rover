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
	return j.Rover
}
