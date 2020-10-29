package game

import "fmt"

/*
   8    7    6    5    4    3    2    1    0
+--------------------------------------------+
| wL | wN | wS | wG | wK | wG | wS | wN | wL |  a
+--------------------------------------------+
|    | wR |    |    |    |    |    | wB |    |  b
+--------------------------------------------+
| wP | wP | wP | wP | wP | wP | wP | wP | wP |  c
+--------------------------------------------+
|    |    |    |    |    |    |    |    |    |  d
+--------------------------------------------+
|    |    |    |    |    |    |    |    |    |  e
+--------------------------------------------+
|    |    |    |    |    |    |    |    |    |  f
+--------------------------------------------+
| bP | bP | bP | bP | bP | bP | bP | bP | bP |  g
+--------------------------------------------+
|    | bB |    |    |    |    |    | bR |    |  h
+--------------------------------------------+
| bL | bN | bS | bG | bK | bG | bS | bN | bL |  i
+--------------------------------------------+
*/

// Board is the board of shogi
type Board struct {
	grid [9][9]*Spot
}

// GetSpot returns a spot on the board
func (b *Board) getSpot(x int, y int) *Spot {
	return b.grid[x][y]
}

// SetSpot sets a spot on the board
func (b *Board) setSpot(x int, y int, s *Spot) {
	b.grid[x][y] = s
}

// NewBoard returns a New instance of a fresh shogi board
func NewBoard() *Board {
	var grid [9][9]*Spot
	grid[0][0] = NewSpot(0, 0, NewLance(true))
	grid[1][0] = NewSpot(1, 0, NewKnight(true))
	grid[2][0] = NewSpot(2, 0, NewSilverGeneral(true))
	grid[3][0] = NewSpot(3, 0, NewGoldGeneral(true))
	grid[4][0] = NewSpot(4, 0, NewKing(true))
	grid[5][0] = NewSpot(5, 0, NewGoldGeneral(true))
	grid[6][0] = NewSpot(6, 0, NewSilverGeneral(true))
	grid[7][0] = NewSpot(7, 0, NewKnight(true))
	grid[8][0] = NewSpot(8, 0, NewLance(true))

	grid[0][1] = NewSpot(0, 1, nil)
	grid[1][1] = NewSpot(1, 1, NewRook(true))
	for i := 2; i < 7; i++ {
		grid[i][1] = NewSpot(i, 1, nil)
	}
	grid[7][1] = NewSpot(7, 1, NewBishop(true))
	grid[8][1] = NewSpot(8, 1, nil)

	for i := 0; i < 9; i++ {
		grid[i][2] = NewSpot(i, 2, NewPawn(true))
	}

	for i := 0; i < 9; i++ {
		for j := 3; j < 6; j++ {
			grid[i][j] = NewSpot(i, j, nil)
		}
	}

	for i := 0; i < 9; i++ {
		grid[i][6] = NewSpot(i, 6, NewPawn(false))
	}

	grid[0][7] = NewSpot(0, 7, nil)
	grid[1][7] = NewSpot(1, 7, NewBishop(false))
	for i := 2; i < 7; i++ {
		grid[i][7] = NewSpot(i, 7, nil)
	}
	grid[7][7] = NewSpot(7, 7, NewRook(false))
	grid[8][7] = NewSpot(8, 7, nil)

	grid[0][8] = NewSpot(0, 8, NewLance(false))
	grid[1][8] = NewSpot(1, 8, NewKnight(false))
	grid[2][8] = NewSpot(2, 8, NewSilverGeneral(false))
	grid[3][8] = NewSpot(3, 8, NewGoldGeneral(false))
	grid[4][8] = NewSpot(4, 8, NewKing(false))
	grid[5][8] = NewSpot(5, 8, NewGoldGeneral(false))
	grid[6][8] = NewSpot(6, 8, NewSilverGeneral(false))
	grid[7][8] = NewSpot(7, 8, NewKnight(false))
	grid[8][8] = NewSpot(8, 8, NewLance(false))

	b := Board{
		grid: grid,
	}

	return &b
}

func (b *Board) resetBoard() {
	grid := b.grid

	grid[0][0] = NewSpot(0, 0, NewLance(false))
	grid[1][0] = NewSpot(1, 0, NewKnight(false))
	grid[2][0] = NewSpot(2, 0, NewSilverGeneral(false))
	grid[3][0] = NewSpot(3, 0, NewGoldGeneral(false))
	grid[4][0] = NewSpot(4, 0, NewKing(false))
	grid[5][0] = NewSpot(5, 0, NewGoldGeneral(false))
	grid[6][0] = NewSpot(6, 0, NewSilverGeneral(false))
	grid[7][0] = NewSpot(7, 0, NewKnight(false))
	grid[8][0] = NewSpot(8, 0, NewLance(false))

	grid[0][1] = NewSpot(0, 1, nil)
	grid[1][1] = NewSpot(1, 1, NewRook(false))
	for i := 2; i < 7; i++ {
		grid[i][1] = NewSpot(i, 1, nil)
	}
	grid[7][1] = NewSpot(7, 1, NewBishop(false))
	grid[8][1] = NewSpot(8, 1, nil)

	for i := 0; i < 9; i++ {
		grid[i][2] = NewSpot(i, 2, NewPawn(false))
	}

	for i := 0; i < 9; i++ {
		for j := 3; j < 6; j++ {
			grid[i][j] = NewSpot(i, j, nil)
		}
	}

	for i := 0; i < 9; i++ {
		grid[i][6] = NewSpot(i, 6, NewPawn(true))
	}

	grid[0][7] = NewSpot(0, 7, nil)
	grid[1][7] = NewSpot(1, 7, NewBishop(true))
	for i := 2; i < 7; i++ {
		grid[i][7] = NewSpot(i, 7, nil)
	}
	grid[7][7] = NewSpot(7, 7, NewRook(true))
	grid[8][7] = NewSpot(8, 7, nil)

	grid[0][8] = NewSpot(0, 8, NewLance(true))
	grid[1][8] = NewSpot(1, 8, NewKnight(true))
	grid[2][8] = NewSpot(2, 8, NewSilverGeneral(true))
	grid[3][8] = NewSpot(3, 8, NewGoldGeneral(true))
	grid[4][8] = NewSpot(4, 8, NewKing(true))
	grid[5][8] = NewSpot(5, 8, NewGoldGeneral(true))
	grid[6][8] = NewSpot(6, 8, NewSilverGeneral(true))
	grid[7][8] = NewSpot(7, 8, NewKnight(true))
	grid[8][8] = NewSpot(8, 8, NewLance(true))

}

// PrintBoard prints the current game board
func (b *Board) PrintBoard() {
	rowNames := [9]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	for i := 8; i >= 0; i-- {
		fmt.Print("| ")
		fmt.Print(i)
		fmt.Print("|")
	}
	fmt.Println()
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			p := b.grid[x][y].p
			if p != nil {
				fmt.Printf("|%s|", p.getKanji())
			} else {
				fmt.Print("|  |")
			}
		}
		fmt.Print(rowNames[y])
		fmt.Println()
	}
}
