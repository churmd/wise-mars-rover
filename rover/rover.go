package rover

import "fmt"

const (
	North orientation = 'N'
	East  orientation = 'E'
	South orientation = 'S'
	West  orientation = 'W'
)

type orientation rune

func (o orientation) String() string {
	return string(o)
}

type Grid struct {
	Width  int
	Height int
}

type Rover struct {
	X           int
	Y           int
	Orientation orientation
	Grid        Grid
	Lost        bool
}

func New(x, y int, orientation orientation, mars Grid) Rover {
	return Rover{X: x, Y: y, Orientation: orientation, Grid: mars}
}

func (r Rover) String() string {
	s := fmt.Sprintf("(%d, %d, %s)", r.X, r.Y, r.Orientation)
	if r.Lost {
		s += " LOST"
	}

	return s
}

func (r Rover) Foward() Rover {
	if r.Lost {
		return r
	}

	updated := r
	switch r.Orientation {
	case North:
		updated.Y++
	case East:
		updated.X++
	case South:
		updated.Y--
	case West:
		updated.X--
	}

	if updated.X < 0 || updated.Y < 0 || updated.X >= updated.Grid.Width || updated.Y >= updated.Grid.Height {
		return r
	}

	return updated
}
