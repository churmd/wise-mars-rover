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

type Grid struct {
	Width  int
	Height int
}

type Rover struct {
	X           int
	Y           int
	Orientation Orientation
	Grid        Grid
	Lost        bool
}

func New(x, y int, orientation Orientation, mars Grid) Rover {
	return Rover{X: x, Y: y, Orientation: orientation, Grid: mars}
}

func (r Rover) String() string {
	s := fmt.Sprintf("(%d, %d, %s)", r.X, r.Y, r.Orientation)
	if r.Lost {
		s += " LOST"
	}

	return s
}

func (r Rover) RotateRight() Rover {
	if r.Lost {
		return r
	}

	switch r.Orientation {
	case North:
		r.Orientation = East
	case East:
		r.Orientation = South
	case South:
		r.Orientation = West
	case West:
		r.Orientation = North
	}
	return r
}

func (r Rover) RotateLeft() Rover {
	if r.Lost {
		return r
	}

	switch r.Orientation {
	case North:
		r.Orientation = West
	case East:
		r.Orientation = North
	case South:
		r.Orientation = East
	case West:
		r.Orientation = South
	}
	return r
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

	if updated.becomeLost() {
		r.Lost = true
		return r
	}

	return updated
}

func (r Rover) becomeLost() bool {
	return r.X < 0 || r.Y < 0 || r.X > r.Grid.Width || r.Y > r.Grid.Height
}
