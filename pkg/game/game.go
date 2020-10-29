package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	BoardDimension = 9
	DropLength     = 3
	DefaultLength  = 4
	PromoteLength  = 5
)

// Game is the struct for a game of shogi
type Game struct {
	Players       [2]Player
	Board         *Board
	Turn          int
	MovesPlayed   []Move
	Status        string
	CurrentPlayer Player
}

// NewGame is the constructor for a game
func NewGame(p0 Player, p1 Player) *Game {
	g := Game{
		Players: [2]Player{p0, p1},
		Board:   NewBoard(),
		Turn:    0,
	}

	return &g
}

func isValidLength(input string) (int, error) {
	length := len(input)
	if length == 3 || length == 4 || length == 5 {
		return length, nil
	}
	return 0, errors.New("Invalid length")
}

func (g *Game) isValidFrom(input []string) (*Spot, error) {
	fromXStr := input[0]
	fromYStr := input[1]

	fromX, err := strconv.Atoi(fromXStr)
	if err != nil {
		return nil, errors.New("Failure to convert FROM X coordinate.")
	}

	fromY, err := strconv.Atoi(fromYStr)
	if err != nil {
		return nil, errors.New("Failure to convert FROM Y coordinate.")
	}

	if fromX >= BoardDimension {
		return nil, errors.New("Invalid FROM X coordinate: outside of board")
	}

	if fromY >= BoardDimension {
		return nil, errors.New("Invalid FROM Y coordinate: outside of board")
	}

	spot := g.Board.getSpot(fromX, fromY)

	if spot.p == nil {
		return nil, errors.New("Invalid FROM Y coordinate: no piece exists in that spot")
	}

	if g.CurrentPlayer.getColor() != spot.p.getColor() {
		return nil, errors.New("Invalid piece: Attempted to move opponent's piece.")
	}

	return spot, nil
}

func (g *Game) isValidTo(toCoordinates []string) (*Spot, error) {
	toXStr := toCoordinates[0]
	toYStr := toCoordinates[1]

	toX, err := strconv.Atoi(toXStr)
	if err != nil {
		return nil, errors.New("Failure to convert TO X coordinate.")
	}

	toY, err := strconv.Atoi(toYStr)
	if err != nil {
		return nil, errors.New("Failure to convert TO Y coordinate.")
	}

	if toX >= BoardDimension {
		return nil, errors.New("Invalid TO X coordinate: outside of board.")
	}

	if toY >= BoardDimension {
		return nil, errors.New("Invalid TO Y coordinate: outside of board.")
	}

	spot := g.Board.getSpot(toX, toY)

	if spot.p != nil && spot.p.getColor() == g.CurrentPlayer.getColor() {
		return nil, errors.New("Invalid TO coordinate: arriving at your own piece.")
	}

	return spot, nil
}

func (g *Game) inPromotionZone(to *Spot) bool {
	if g.CurrentPlayer.getColor() {
		return to.y <= 2
	}

	return to.y >= 6
}

func (g *Game) isPromotable(to *Spot) error {
	if !to.p.isPromotable() {
		return errors.New("Promotion error: Piece is not promotable.")
	}

	if !g.inPromotionZone(to) {
		return errors.New("Promotion error: TO coordinate not in promotion zone.")
	}

	return nil
}

func (g *Game) isCaptured(input []string) (Piece, error) {
	pieceStr := input[0]

	captured := g.CurrentPlayer.getCaptured()

	if count, ok := captured[pieceStr]; !ok || count == 0 {
		return nil, errors.New("Invalid drop: you haven't captured a " + pieceStr)
	}

	captured[pieceStr]--
	initializer, _ := pieceInitMap[pieceStr]
	piece := initializer(g.CurrentPlayer.getColor())

	return piece, nil
}

func (g *Game) isValid(input string) (Move, error) {
	length, err := isValidLength(input)
	if err != nil {
		return nil, err
	}

	s := strings.Split(input, "")
	var from *Spot
	var to *Spot
	var piece Piece
	var isPromotion bool
	if length != DropLength {
		fromCoordinates := s[:2]
		from, err = g.isValidFrom(fromCoordinates)
		if err != nil {
			return nil, err
		}

		piece = from.p
		toCoordinates := s[2:4]
		to, err = g.isValidTo(toCoordinates)

		if err != nil {
			return nil, err
		}

		if length == PromoteLength {
			if err = g.isPromotable(to); err != nil {
				return nil, err
			}
			isPromotion = true
		}
	} else {
		piece, err = g.isCaptured(s)

		if err != nil {
			return nil, err
		}
	}
	move := newDefaultMove(g.CurrentPlayer, from, to, piece, isPromotion)

	fmt.Println(move.toString())
	return move, nil
}

func (g *Game) execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	if len(input) == 0 {
		return errors.New("No input")
	}

	// TODO: create execMove function
	move, err := g.isValid(input)
	if err != nil {
		return err
	}
	fmt.Println(move.toString())

	return nil
}

// Init initializes a shogi game
func Init() {

	p0 := Player{false, make(map[string]int)}
	p1 := Player{true, make(map[string]int)}
	g := NewGame(p0, p1)
	fmt.Println("in pkg game")
	reader := bufio.NewReader(os.Stdin)
	for g.Status != "GAME OVER" {

		g.Board.PrintBoard()
		fmt.Println("Enter move: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = g.execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}
}
