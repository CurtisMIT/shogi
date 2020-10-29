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

var pieceInitMap map[string]func(bool) Piece

// Base is base struct for a piece in shogi
type Base struct {
	isBlack    bool
	kanji      string
	name       string
	promotable bool
}

func NewPiece(isBlack bool, kanji string, name string, promotable bool) *Base {
	b := Base{
		isBlack:    isBlack,
		kanji:      kanji,
		name:       name,
		promotable: promotable,
	}

	return &b
}

// GetKanji returns the kanji of the piece
func (b *Base) getKanji() string {
	return b.kanji
}

// GetName returns the name of the piece
func (b *Base) getName() string {
	return b.name
}

// GetColor returns the color of the piece
func (b *Base) getColor() bool {
	return b.isBlack
}

// IsPromotable returns whether the piece is promotable or not
func (b *Base) isPromotable() bool {
	return b.promotable
}

// Equals returns whether this piece is equal to that other piece
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

func NewKing(isBlack bool) *King {
	k := King{
		NewPiece(isBlack, nKing, king, false),
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

func NewRook(isBlack bool) *Rook {
	r := Rook{
		NewPiece(isBlack, nRook, rook, true),
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

func NewBishop(isBlack bool) *Bishop {
	b := Bishop{
		NewPiece(isBlack, nBishop, bishop, true),
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

func NewGoldGeneral(isBlack bool) *GoldGeneral {
	gg := GoldGeneral{
		NewPiece(isBlack, nGGeneral, gGeneral, false),
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

func NewSilverGeneral(isBlack bool) *SilverGeneral {
	ss := SilverGeneral{
		NewPiece(isBlack, nSGeneral, sGeneral, true),
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

func NewKnight(isBlack bool) *Knight {
	n := Knight{
		NewPiece(isBlack, nKnight, knight, true),
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

func NewLance(isBlack bool) *Lance {
	l := Lance{
		NewPiece(isBlack, nLance, lance, true),
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

func NewPawn(isBlack bool) *Pawn {
	p := Pawn{
		NewPiece(isBlack, nPawn, pawn, true),
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

func NewPPawn(isBlack bool) *PPawn {
	pp := PPawn{
		NewPiece(isBlack, nPPawn, pPawn, false),
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

func NewPKnight(isBlack bool) *PKnight {
	pk := PKnight{
		NewPiece(isBlack, nPKnight, pKnight, false),
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

func NewPSilver(isBlack bool) *PSilver {
	ps := PSilver{
		NewPiece(isBlack, nPSilver, pSilver, false),
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

func NewPBishop(isBlack bool) *PBishop {
	pb := PBishop{
		NewPiece(isBlack, nPBishop, pBishop, false),
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

func NewPRook(isBlack bool) *PRook {
	pr := PRook{
		NewPiece(isBlack, nPRook, pRook, false),
	}

	return &pr
}

func (pr *PRook) canMove(board *Board, start Spot, end Spot) bool {
	return true
}
