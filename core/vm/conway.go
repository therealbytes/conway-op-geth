// Based on: https://github.com/pinpox/go-game-of-life/blob/master/game-of-life.go

package vm

const (
	Alive uint8 = 1
	Dead  uint8 = 0
)

type GameBoard struct {
	xSize, ySize int
	cells        []uint8
}

func NewGameBoard(x, y int) *GameBoard {
	cells := make([]uint8, x*y)
	return &GameBoard{xSize: x, ySize: y, cells: cells}
}

func (gb *GameBoard) Iterate() {

	gbOld := NewGameBoard(gb.xSize, gb.ySize)
	copy(gbOld.cells, gb.cells)

	for y := 0; y < gb.ySize; y++ {
		for x := 0; x < gb.xSize; x++ {

			neighbors := gbOld.Neighbors(x, y)

			// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
			if gbOld.Get(x, y) == Dead {
				if neighbors == 3 {
					gb.Set(x, y, Alive)
				}
				continue
			}

			// Any live cell with fewer than two live neighbors dies, as if by underpopulation.
			if neighbors < 2 {
				gb.Set(x, y, Dead)
				continue
			}

			// Any live cell with more than three live neighbors dies, as if by overpopulation.
			if neighbors > 3 {
				gb.Set(x, y, Dead)
				continue
			}

			// Any live cell with two or three live neighbors lives on to the next generation.
			// No need to set, already alive
		}
	}
}

func (gb *GameBoard) Equal(gb2 *GameBoard) bool {
	if gb.xSize != gb2.xSize || gb.ySize != gb2.ySize {
		return false
	}

	if len(gb.cells) != len(gb2.cells) {
		return false
	}

	for k := range gb.cells {
		if gb.cells[k] != gb2.cells[k] {
			return false
		}
	}

	return true
}

func (gb *GameBoard) Set(x, y int, val uint8) {
	if !gb.InBounds(x, y) {
		return
	}
	gb.cells[y*(gb.xSize)+x] = val
}

func (gb *GameBoard) Get(x, y int) uint8 {
	if !gb.InBounds(x, y) {
		return Dead
	}
	return gb.cells[y*(gb.xSize)+x]
}

// Neighbors returns the number of alive Neighbors of a given cell
func (gb *GameBoard) Neighbors(x, y int) int {
	count := 0
	arr := []int{-1, 0, 1}

	for _, v1 := range arr {
		for _, v2 := range arr {
			if v1 == 0 && v2 == 0 {
				continue
			}
			if gb.InBounds(x+v1, y+v2) {
				if gb.Get(x+v1, y+v2) == Alive {
					count++
				}
			}
		}
	}
	return count
}

// InBounds is a helper function to check if a coordinate is inside the grid.
func (gb *GameBoard) InBounds(x int, y int) bool {
	return (x >= 0 &&
		x < gb.xSize &&
		y >= 0 &&
		y < gb.ySize)
}
