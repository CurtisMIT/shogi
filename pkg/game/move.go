package game

import "fmt"

// Move is struct for a move in shogi
type Move struct {
	player        Player
	from          *Spot
	to            *Spot
	pieceMoved    Piece
	pieceCaptured Piece
	isPromotion   bool
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

func newDefaultMove(player Player, from *Spot, to *Spot, piece Piece, isPromotion bool) Move {
	var pieceCaptured Piece
	if to.p != nil {
		pieceCaptured = to.p
	}
	m := Move{
		player:        player,
		from:          from,
		to:            to,
		pieceMoved:    piece,
		pieceCaptured: pieceCaptured,
		isPromotion:   isPromotion,
	}

	return m
}

func newPromoteMove(player Player, from *Spot, to *Spot, piece Piece) Move {
	m := newDefaultMove(player, from, to, piece, true)
	return m
}

func newDropMove(player Player, to *Spot, piece Piece) Move {
	m := Move{
		player:        player,
		from:          nil,
		to:            to,
		pieceMoved:    piece,
		pieceCaptured: nil,
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

func (m *Move) toString() string {
	s := fmt.Sprintf("%s %s %s %t", m.player.toString(), m.from.toString(), m.to.toString(), m.isPromotion)
	return s
}
