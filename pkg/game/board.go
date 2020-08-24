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

func (b *Board) getSpot(x int, y int) *Spot {
	return b.grid[x][y]
}

func (b *Board) setSpot(x int, y int, s *Spot) {
	b.grid[x][y] = s
}

func newBoard() *Board {
	var grid [9][9]*Spot
	grid[0][0] = newSpot(0, 0, newLance(false))
	grid[1][0] = newSpot(1, 0, newKnight(false))
	grid[2][0] = newSpot(2, 0, newSilverGeneral(false))
	grid[3][0] = newSpot(3, 0, newGoldGeneral(false))
	grid[4][0] = newSpot(4, 0, newKing(false))
	grid[5][0] = newSpot(5, 0, newGoldGeneral(false))
	grid[6][0] = newSpot(6, 0, newSilverGeneral(false))
	grid[7][0] = newSpot(7, 0, newKnight(false))
	grid[8][0] = newSpot(8, 0, newLance(false))

	grid[0][1] = newSpot(0, 1, nil)
	grid[1][1] = newSpot(1, 1, newRook(false))
	for i := 2; i < 7; i++ {
		grid[i][1] = newSpot(i, 1, nil)
	}
	grid[7][1] = newSpot(7, 1, newBishop(false))
	grid[8][1] = newSpot(8, 1, nil)

	for i := 0; i < 9; i++ {
		grid[i][2] = newSpot(i, 2, newPawn(false))
	}

	for i := 0; i < 9; i++ {
		for j := 3; j < 6; j++ {
			grid[i][j] = newSpot(i, j, nil)
		}
	}

	for i := 0; i < 9; i++ {
		grid[i][6] = newSpot(i, 6, newPawn(true))
	}

	grid[0][7] = newSpot(0, 7, nil)
	grid[1][7] = newSpot(1, 7, newBishop(true))
	for i := 2; i < 7; i++ {
		grid[i][7] = newSpot(i, 7, nil)
	}
	grid[7][7] = newSpot(7, 7, newRook(true))
	grid[8][7] = newSpot(8, 7, nil)

	grid[0][8] = newSpot(0, 8, newLance(true))
	grid[1][8] = newSpot(1, 8, newKnight(true))
	grid[2][8] = newSpot(2, 8, newSilverGeneral(true))
	grid[3][8] = newSpot(3, 8, newGoldGeneral(true))
	grid[4][8] = newSpot(4, 8, newKing(true))
	grid[5][8] = newSpot(5, 8, newGoldGeneral(true))
	grid[6][8] = newSpot(6, 8, newSilverGeneral(true))
	grid[7][8] = newSpot(7, 8, newKnight(true))
	grid[8][8] = newSpot(8, 8, newLance(true))

	b := Board{
		grid: grid,
	}

	return &b
}

func (b *Board) resetBoard() {
	grid := b.grid

	grid[0][0] = newSpot(0, 0, newLance(false))
	grid[1][0] = newSpot(1, 0, newKnight(false))
	grid[2][0] = newSpot(2, 0, newSilverGeneral(false))
	grid[3][0] = newSpot(3, 0, newGoldGeneral(false))
	grid[4][0] = newSpot(4, 0, newKing(false))
	grid[5][0] = newSpot(5, 0, newGoldGeneral(false))
	grid[6][0] = newSpot(6, 0, newSilverGeneral(false))
	grid[7][0] = newSpot(7, 0, newKnight(false))
	grid[8][0] = newSpot(8, 0, newLance(false))

	grid[0][1] = newSpot(0, 1, nil)
	grid[1][1] = newSpot(1, 1, newRook(false))
	for i := 2; i < 7; i++ {
		grid[i][1] = newSpot(i, 1, nil)
	}
	grid[7][1] = newSpot(7, 1, newBishop(false))
	grid[8][1] = newSpot(8, 1, nil)

	for i := 0; i < 9; i++ {
		grid[i][2] = newSpot(i, 2, newPawn(false))
	}

	for i := 0; i < 9; i++ {
		for j := 3; j < 6; j++ {
			grid[i][j] = newSpot(i, j, nil)
		}
	}

	for i := 0; i < 9; i++ {
		grid[i][6] = newSpot(i, 6, newPawn(true))
	}

	grid[0][7] = newSpot(0, 7, nil)
	grid[1][7] = newSpot(1, 7, newBishop(true))
	for i := 2; i < 7; i++ {
		grid[i][7] = newSpot(i, 7, nil)
	}
	grid[7][7] = newSpot(7, 7, newRook(true))
	grid[8][7] = newSpot(8, 7, nil)

	grid[0][8] = newSpot(0, 8, newLance(true))
	grid[1][8] = newSpot(1, 8, newKnight(true))
	grid[2][8] = newSpot(2, 8, newSilverGeneral(true))
	grid[3][8] = newSpot(3, 8, newGoldGeneral(true))
	grid[4][8] = newSpot(4, 8, newKing(true))
	grid[5][8] = newSpot(5, 8, newGoldGeneral(true))
	grid[6][8] = newSpot(6, 8, newSilverGeneral(true))
	grid[7][8] = newSpot(7, 8, newKnight(true))
	grid[8][8] = newSpot(8, 8, newLance(true))

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

func (b *Board) execMove(m *Move) {

}
