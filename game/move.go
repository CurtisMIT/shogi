package game

// Move is struct for a move in shogi
type Move struct {
	player      Player
	start       Spot
	end         Spot
	pieceMoved  Piece
	pieceKilled Piece
}

func newMove(player Player, start Spot, end Spot) Move {
	m := Move{
		player: player,
		start:  start,
		end:    end,
	}

	return m
}
