package rover

import "fmt"

const (
	North Orientation = 'N'
	East  Orientation = 'E'
	South Orientation = 'S'
	West  Orientation = 'W'
)

type Orientation rune

func (o Orientation) String() string {
	return string(o)
}

type Rover struct {
	X           int
	Y           int
	Orientation Orientation
	Lost        bool
}

func New(x, y int, orientation Orientation) Rover {
	return Rover{X: x, Y: y, Orientation: orientation, Lost: false}
}

func (r Rover) String() string {
	s := fmt.Sprintf("(%d, %d, %s)", r.X, r.Y, r.Orientation)
	if r.Lost {
		s += " LOST"
	}

	return s
}
