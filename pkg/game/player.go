package game

// Player is a base class for a player
type Player struct {
	black    bool
	captured map[string]int
}

func (p *Player) getColor() bool {
	return p.black
}

func (p *Player) getCaptured() map[string]int {
	return p.captured
}

func (p *Player) toString() string {
	if p.black {
		return "black"
	}
	return "white"
}
