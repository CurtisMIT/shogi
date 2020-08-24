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
	CurrentPlayer bool
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
	fmt.Println("In from validation")
	fromX, err := strconv.Atoi(fromXStr)
	if err != nil {
		return nil, errors.New("Failure to convert FROM X coordinate.")
	}

	fromY, err := strconv.Atoi(fromYStr)
	if err != nil {
		return nil, errors.New("Failure to conver FROM Y coordinate.")
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

	if g.CurrentPlayer != spot.p.getColor() {
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

	if spot.p != nil {
		// TODO: call isValidCapture here
		fmt.Println("isValidCapture will be called")
	}

	return spot, nil
}

func (g *Game) isValid(input string) error {
	length, err := isValidLength(input)
	if err != nil {
		return err
	}
	fmt.Print(length)

	s := strings.Split(input, "")
	var from *Spot
	var to *Spot
	if length == DefaultLength || length == PromoteLength {
		fromCoordinates := s[:2]
		from, err = g.isValidFrom(fromCoordinates)

		toCoordinates := s[2:4]
		to, err = g.isValidTo(toCoordinates)
	}

	if err != nil {
		return err
	}

	fmt.Println(from)
	fmt.Println(to)
	return nil
}

func (g *Game) execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	if len(input) == 0 {
		return errors.New("No input")
	}

	// TODO: create execMove function
	err := g.isValid(input)
	if err != nil {
		return err
	}

	return nil
}

// Init initializes a shogi game
func Init() {

	p0 := Player{false}
	p1 := Player{true}
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
