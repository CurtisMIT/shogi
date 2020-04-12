package game

import (
	"strconv"
)

// Spot refers to a spot on the board
type Spot struct {
	x int
	y int
	p canMover
}

type canMover interface {
	canMove(*Board, Spot, Spot) bool
	getName() string
}

func newSpot(x int, y int, p canMover) *Spot {
	s := &Spot{x, y, p}
	return s
}

func (s *Spot) toString() string {
	var str string
	str += "x: "
	str += strconv.Itoa(s.x)
	str += " y: "
	str += strconv.Itoa(s.y)
	str += " #: "
	str += s.p.getName()
	return str
}
