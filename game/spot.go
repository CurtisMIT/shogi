package game

// Spot refers to a spot on the board
type Spot struct {
	x int
	y int
	p canMover
}

type canMover interface {
	canMove(Board, Spot, Spot) bool
}

func newSpot(x int, y int, p canMover) *Spot {
	s := &Spot{x, y, p}
	return s
}
