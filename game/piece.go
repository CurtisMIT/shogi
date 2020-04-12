package game

const (
	pawn     string = "\u6b69"
	lance    string = "\u9999"
	knight   string = "\u6842"
	sGeneral string = "\u9280"
	gGeneral string = "\u91d1"
	bishop   string = "\u89d2"
	rook     string = "\u98db"
	king     string = "\u7389"
	pPawn    string = "\u3068"
	pLance   string = "\u674f"
	pKnight  string = "\u572d"
	pSilver  string = "\u5168"
	pBishop  string = "\u99ac"
	pRook    string = "\u9f8d"
)

// Piece is base struct for a piece in shogi
type Piece struct {
	black bool
	name  string
}

func newPiece(isBlack bool, name string) *Piece {
	p := Piece{
		black: isBlack,
		name:  name,
	}

	return &p
}

// Name returns name of the piece
func (p *Piece) getName() string {
	return p.name
}

// King is a king piece
type King struct {
	*Piece
}

func newKing(isBlack bool) *King {
	k := King{
		newPiece(isBlack, king),
	}

	return &k
}

func (k *King) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// Rook is a rook piece
type Rook struct {
	*Piece
}

func newRook(isBlack bool) *Rook {
	r := Rook{
		newPiece(isBlack, rook),
	}

	return &r
}

func (r *Rook) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// Bishop is a bishop piece
type Bishop struct {
	*Piece
}

func newBishop(isBlack bool) *Bishop {
	b := Bishop{
		newPiece(isBlack, bishop),
	}

	return &b
}

func (b *Bishop) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// GoldGeneral is a gold general piece
type GoldGeneral struct {
	*Piece
}

func newGoldGeneral(isBlack bool) *GoldGeneral {
	gg := GoldGeneral{
		newPiece(isBlack, gGeneral),
	}

	return &gg
}

func (gg *GoldGeneral) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// SilverGeneral is a silver general piece
type SilverGeneral struct {
	*Piece
}

func newSilverGeneral(isBlack bool) *SilverGeneral {
	ss := SilverGeneral{
		newPiece(isBlack, sGeneral),
	}

	return &ss
}

func (sg *SilverGeneral) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// Knight is a knight piece
type Knight struct {
	*Piece
}

func newKnight(isBlack bool) *Knight {
	n := Knight{
		newPiece(isBlack, knight),
	}

	return &n
}

func (n *Knight) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// Lance is lance piece
type Lance struct {
	*Piece
}

func newLance(isBlack bool) *Lance {
	l := Lance{
		newPiece(isBlack, lance),
	}

	return &l
}

func (l *Lance) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// Pawn is a pawn piece
type Pawn struct {
	*Piece
}

func newPawn(isBlack bool) *Pawn {
	p := Pawn{
		newPiece(isBlack, pawn),
	}

	return &p
}

func (p *Pawn) canMove(board *Board, start Spot, end Spot) bool {
	return true
}
