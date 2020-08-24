package game

const (
	nPawn     string = "\u6b69"
	nLance    string = "\u9999"
	nKnight   string = "\u6842"
	nSGeneral string = "\u9280"
	nGGeneral string = "\u91d1"
	nBishop   string = "\u89d2"
	nRook     string = "\u98db"
	nKing     string = "\u7389"
	nPPawn    string = "\u3068"
	nPLance   string = "\u674f"
	nPKnight  string = "\u572d"
	nPSilver  string = "\u5168"
	nPBishop  string = "\u99ac"
	nPRook    string = "\u9f8d"

	pawn     string = "p"
	lance    string = "l"
	knight   string = "n"
	sGeneral string = "s"
	gGeneral string = "g"
	bishop   string = "b"
	rook     string = "r"
	king     string = "k"
	pPawn    string = "P"
	pKnight  string = "N"
	pSilver  string = "S"
	pBishop  string = "B"
	pRook    string = "R"
)

// Base is base struct for a piece in shogi
type Base struct {
	isBlack bool
	kanji   string
	name    string
}

func newPiece(isBlack bool, kanji string, name string) *Base {
	b := Base{
		isBlack: isBlack,
		kanji:   kanji,
		name:    name,
	}

	return &b
}

func (p *Base) getKanji() string {
	return p.kanji
}

func (p *Base) getName() string {
	return p.name
}

func (p *Base) getColor() bool {
	return p.isBlack
}

func (this *Base) equals(other Piece) bool {
	if this.getName() != other.getName() {
		return false
	}

	if this.getKanji() != other.getKanji() {
		return false
	}

	if this.getColor() != other.getColor() {
		return false
	}

	return true
}

// King is a king piece
type King struct {
	*Base
}

func newKing(isBlack bool) *King {
	k := King{
		newPiece(isBlack, nKing, king),
	}

	return &k
}

func (k *King) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// Rook is a rook piece
type Rook struct {
	*Base
}

func newRook(isBlack bool) *Rook {
	r := Rook{
		newPiece(isBlack, nRook, rook),
	}

	return &r
}

func (r *Rook) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// Bishop is a bishop piece
type Bishop struct {
	*Base
}

func newBishop(isBlack bool) *Bishop {
	b := Bishop{
		newPiece(isBlack, nBishop, bishop),
	}

	return &b
}

func (b *Bishop) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// GoldGeneral is a gold general piece
type GoldGeneral struct {
	*Base
}

func newGoldGeneral(isBlack bool) *GoldGeneral {
	gg := GoldGeneral{
		newPiece(isBlack, nGGeneral, gGeneral),
	}

	return &gg
}

func (gg *GoldGeneral) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// SilverGeneral is a silver general piece
type SilverGeneral struct {
	*Base
}

func newSilverGeneral(isBlack bool) *SilverGeneral {
	ss := SilverGeneral{
		newPiece(isBlack, nSGeneral, sGeneral),
	}

	return &ss
}

func (sg *SilverGeneral) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// Knight is a knight piece
type Knight struct {
	*Base
}

func newKnight(isBlack bool) *Knight {
	n := Knight{
		newPiece(isBlack, nKnight, knight),
	}

	return &n
}

func (n *Knight) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// Lance is lance piece
type Lance struct {
	*Base
}

func newLance(isBlack bool) *Lance {
	l := Lance{
		newPiece(isBlack, nLance, lance),
	}

	return &l
}

func (l *Lance) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// Pawn is a pawn piece
type Pawn struct {
	*Base
}

func newPawn(isBlack bool) *Pawn {
	p := Pawn{
		newPiece(isBlack, nPawn, pawn),
	}

	return &p
}

func (p *Pawn) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// PPawn is a promoted pawn piece
type PPawn struct {
	*Base
}

func newPPawn(isBlack bool) *PPawn {
	pp := PPawn{
		newPiece(isBlack, nPPawn, pPawn),
	}

	return &pp
}

func (pp *PPawn) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// PKnight is a promoted knight piece
type PKnight struct {
	*Base
}

func newPKnight(isBlack bool) *PKnight {
	pk := PKnight{
		newPiece(isBlack, nPKnight, pKnight),
	}

	return &pk
}

func (pk *PKnight) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// PSilver is a promoted sliver general piece
type PSilver struct {
	*Base
}

func newPSilver(isBlack bool) *PSilver {
	ps := PSilver{
		newPiece(isBlack, nPSilver, pSilver),
	}

	return &ps
}

func (ps *PSilver) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// PBishop is a promoted bishop piece
type PBishop struct {
	*Base
}

func newPBishop(isBlack bool) *PBishop {
	pb := PBishop{
		newPiece(isBlack, nPBishop, pBishop),
	}

	return &pb
}

func (pb *PBishop) canMove(board *Board, start Spot, end Spot) bool {
	return true
}

// PRook is a promoted rook piece
type PRook struct {
	*Base
}

func newPRook(isBlack bool) *PRook {
	pr := PRook{
		newPiece(isBlack, nPRook, pRook),
	}

	return &pr
}

func (pr *PRook) canMove(board *Board, start Spot, end Spot) bool {
	return true
}
