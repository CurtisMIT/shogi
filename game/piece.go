package game

// Piece is base struct for a piece in shogi
type Piece struct {
	black bool
}

func newPiece(isBlack bool) *Piece {
	p := Piece{
		black: isBlack,
	}

	return &p
}

// King is a king piece
type King struct {
	black bool
}

func (k *King) canMove(board Board, start Spot, end Spot) bool {
	return true
}

// Rook is a rook piece
type Rook struct {
	black bool
}

func (r *Rook) canMove(board Board, start Spot, end Spot) bool {
	return true
}

// Bishop is a bishop piece
type Bishop struct {
	black bool
}

func (b *Bishop) canMove(board Board, start Spot, end Spot) bool {
	return true
}

// GoldGeneral is a gold general piece
type GoldGeneral struct {
	black bool
}

func (gg *GoldGeneral) canMove(board Board, start Spot, end Spot) bool {
	return true
}

// SilverGeneral is a silver general piece
type SilverGeneral struct {
	black bool
}

func (sg *SilverGeneral) canMove(board Board, start Spot, end Spot) bool {
	return true
}

// Knight is a knight piece
type Knight struct {
	black bool
}

func (n *Knight) canMove(board Board, start Spot, end Spot) bool {
	return true
}

// Lance is lance piece
type Lance struct {
	black bool
}

func (l *Lance) canMove(board Board, start Spot, end Spot) bool {
	return true
}

// Pawn is a pawn piece
type Pawn struct {
	black bool
}

func (p *Pawn) canMove(board Board, start Spot, end Spot) bool {
	return true
}
