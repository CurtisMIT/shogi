package game

// Move is struct for a move in shogi
type Move struct {
	player        Player
	from          *Spot
	to            *Spot
	pieceMoved    Piece
	pieceCaptured Piece
	promote       string
	strForm       string
}

func newMove(player Player, from *Spot, to *Spot) Move {
	m := Move{
		player: player,
		from:   from,
		to:     to,
	}

	return m
}

func (m *Move) getPlayer() Player {
	return m.player
}

func (m *Move) getFrom() *Spot {
	return m.from
}

func (m *Move) getTo() *Spot {
	return m.to
}
