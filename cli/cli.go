package cli

import (
	"fmt"
	"io"
	"log"

	"github.com/churmd/wise-mars-rover/movement"
	"github.com/churmd/wise-mars-rover/rover"
)

type Interactor struct {
	Input io.Reader
	Output io.Writer
}

func New(in io.Reader, out io.Writer) Interactor {
	return Interactor{Input: in, Output: out}
}

func (i Interactor) Run() {
	grid := i.getGrid()
	journeys := i.getRoverJourneys(grid)

	finalRovers := make([]rover.Rover, len(journeys))
	for i, j := range journeys {
		r := j.Run()
		finalRovers[i] = r
	}

	fmt.Fprintln(i.Output, "Final Rover States")
	for _, r := range finalRovers {
		fmt.Fprintln(i.Output, r)
	}
}

func (i Interactor) getRoverJourneys(grid rover.Grid) []movement.Journey {
	var journeys []movement.Journey
	for {
		fmt.Fprintln(i.Output, "Enter a rover and list of commands, e.g. 2 3 E LFRFF. Or an empty line to continue")
		var x int
		var y int
		var oreintation rover.Orientation
		var cmds string
		num, err := fmt.Fscanf(i.Input, "%d %d %c %s", &x, &y, &oreintation, &cmds)
		if num == 0 && err.Error() == "unexpected newline" {
			break
		}

		if err != nil {
			log.Fatalf("could not read rover and commands, only received %d values. %s", num, err)
		}

		if !oreintationValid(oreintation) {
			log.Printf("oreientation %s is not valid, must be one of N E S W\n", oreintation)
			continue
		}

		if !commandsValid(cmds) {
			log.Printf("commands %s are not valid, must contain only R L F\n", oreintation)
			continue
		}

		r := rover.New(x, y, oreintation, grid)
		j := movement.Journey{Rover: r, Commands: []movement.Command(cmds)}
		journeys = append(journeys, j)
	}
	return journeys
}

func (i Interactor) getGrid() rover.Grid {
	fmt.Fprintln(i.Output, "Enter the size of the grid, e.g. 4 8")

	var gridWidth int
	var gridHeight int
	num, err := fmt.Fscanf(i.Input, "%d %d", &gridWidth, &gridHeight)
	if err != nil {
		log.Fatalf("could not read grid size, only received %d values. %s", num, err)
	}

	grid := rover.Grid{Width: gridWidth, Height: gridHeight}
	return grid
}

func oreintationValid(o rover.Orientation) bool {
	return o == rover.North || o == rover.East|| o == rover.South || o ==rover.West
}

func commandsValid(s string) bool {
	for _, c := range s {
		cmd := movement.Command(c)
		if !(cmd == movement.Left || cmd == movement.Right || cmd == movement.Foward) {
			return false
		}
	}

	return true
}