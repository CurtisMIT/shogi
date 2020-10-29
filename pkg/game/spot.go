package game

import "strconv"

// Spot refers to a spot on the board
type Spot struct {
	x int
	y int
	p Piece
}

// Piece refers to a piece on the board
type Piece interface {
	canMove(*Board, Spot, Spot) bool
	getKanji() string
	getName() string
	getColor() bool
	isPromotable() bool
	equals(Piece) bool
}

func NewSpot(x int, y int, p Piece) *Spot {
	s := &Spot{x, y, p}
	return s
}

func (s *Spot) getX() int {
	return s.x
}

func (s *Spot) getY() int {
	return s.y
}

func (s *Spot) getCanMover() Piece {
	return s.p
}

func (s *Spot) setX(x int) {
	s.x = x
}

func (s *Spot) setY(y int) {
	s.y = y
}

func (s *Spot) setCanMover(p Piece) {
	s.p = p
}

func (s *Spot) toString() string {
	var str string
	str += "x: "
	str += strconv.Itoa(s.x)
	str += " y: "
	str += strconv.Itoa(s.y)
	str += " #: "
	if s.p != nil {
		str += s.p.getKanji()
	}
	return str
}

func (this *Spot) equals(other *Spot) bool {
	if this.x != other.x {
		return false
	}

	if this.y != other.y {
		return false
	}

	if !this.p.equals(other.p) {
		return false
	}

	return true
}
