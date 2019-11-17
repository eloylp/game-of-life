package main

type Cell bool
type Board [4][4]Cell

func (b Board) NextGen() Board {
	newBoard := Board{}
	for y := 0; y < len(b); y++ {
		for x := 0; x < len(b[x]); x++ {
			newBoard[y][x] = b.NextGenCell(x, y)
		}
	}
	return newBoard
}

var neighbourOffsets = [8][2]int{
	{+1, +1},
	{+1, 0},
	{+1, -1},
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, +1},
	{0, +1},
}

func (b Board) NextGenCell(x, y int) Cell {
	neighbours := 0
	currentCell := b[y][x]
	for _, o := range neighbourOffsets {
		if Neighbour(b, x, y, o[0], o[1]) {
			neighbours++
		}
	}
	if !currentCell && neighbours == 3 {
		return true
	}
	if currentCell && (neighbours == 2 || neighbours == 3) {
		return true
	}
	return false
}

func main() {

}

func IsReachableNeighbour(board Board, x, y, ox, oy int) bool {
	if y+oy > len(board)-1 || y+oy < 0 {
		return false
	}
	if x+ox > len(board[y])-1 || x+ox < 0 {
		return false
	}
	return true
}

func Neighbour(board Board, x, y, ox, oy int) Cell {
	if !IsReachableNeighbour(board, x, y, ox, oy) {
		return Cell(false)
	}
	return board[y+oy][x+ox]
}
