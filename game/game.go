package game

// Game is the struct for a game of shogi
type Game struct {
	Players     [2]Player
	Board       *Board
	Turn        int
	MovesPlayed []Move
}

// NewGame is the constructor for a game
func NewGame(p0 Player, p1 Player) *Game {
	g := Game{
		Players: [2]Player{p0, p1},
		Board:   newBoard(),
		Turn:    0,
	}

	return &g
}
