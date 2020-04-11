package game

// Board is the board of shogi
type Board struct {
	grid [9][9]*Spot
}

func (b *Board) resetBoard() {
	grid := b.grid

	grid[0][0] = newSpot(0, 0, &Lance{false})
	grid[1][0] = newSpot(1, 0, &Knight{false})
	grid[2][0] = newSpot(2, 0, &SilverGeneral{false})
	grid[3][0] = newSpot(3, 0, &GoldGeneral{false})
	grid[4][0] = newSpot(4, 0, &King{false})
	grid[5][0] = newSpot(5, 0, &GoldGeneral{false})
	grid[6][0] = newSpot(6, 0, &SilverGeneral{false})
	grid[7][0] = newSpot(7, 0, &Knight{false})
	grid[8][0] = newSpot(8, 0, &Lance{false})

	grid[0][1] = newSpot(0, 1, nil)
	grid[1][1] = newSpot(1, 1, &Rook{false})
	for i := 2; i < 7; i++ {
		grid[i][1] = newSpot(i, 1, nil)
	}
	grid[7][1] = newSpot(7, 1, &Bishop{false})
	grid[8][1] = newSpot(8, 1, nil)

	for i := 0; i < 9; i++ {
		grid[i][2] = newSpot(i, 2, &Pawn{false})
	}

	for i := 0; i < 9; i++ {
		for j := 3; j < 6; j++ {
			grid[i][j] = newSpot(i, j, nil)
		}
	}

	for i := 0; i < 9; i++ {
		grid[i][6] = newSpot(i, 6, &Pawn{true})
	}

	grid[0][7] = newSpot(0, 7, nil)
	grid[1][7] = newSpot(1, 7, &Bishop{true})
	for i := 2; i < 7; i++ {
		grid[i][7] = newSpot(i, 7, nil)
	}
	grid[7][7] = newSpot(7, 7, &Rook{true})
	grid[8][7] = newSpot(8, 7, nil)

	grid[0][8] = newSpot(0, 8, &Lance{true})
	grid[1][8] = newSpot(1, 8, &Knight{true})
	grid[2][8] = newSpot(2, 8, &SilverGeneral{true})
	grid[3][8] = newSpot(3, 8, &GoldGeneral{true})
	grid[4][8] = newSpot(4, 8, &King{true})
	grid[5][8] = newSpot(5, 8, &GoldGeneral{true})
	grid[6][8] = newSpot(6, 8, &SilverGeneral{true})
	grid[7][8] = newSpot(7, 8, &Knight{true})
	grid[8][8] = newSpot(8, 8, &Lance{true})

}
