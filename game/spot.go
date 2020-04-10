package game

type spot struct {
	x int
	y int
}

func Spot(x int, y int) *spot {
	s := spot{x, y}
	return &s
}
