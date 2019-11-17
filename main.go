package main

type Cell bool
type Board [4][4]Cell

func (b Board) NextGen() Board {
	newBoard := Board{}
	for y := 0; y < len(b); y++ {
		for x := 0; x < len(b[x]); x++ {
			newBoard[y][x] = b.NextCell(x, y)
		}
	}
	return newBoard
}

var offsets = [8][2]int{
	{+1, +1},
	{+1, 0},
	{+1, -1},
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, +1},
	{0, +1},
}

func (b Board) NextCell(x, y int) Cell {
	c := 0
	currentCell := b[y][x]
	for _, o := range offsets {
		if (CheckNeighbour(b, x, y, o[0], o[1])) {
			c++
		}
	}
	if !currentCell && c == 3 {
		return true
	}
	if currentCell && (c == 2 || c == 3) {
		return true
	}
	return false
}

func CheckNeighbour(board Board, x, y, ox, oy int) Cell {
	if y+oy > len(board) || y+oy < 0 {
		return false
	}
	if x+ox > len(board[y]) || x+ox < 0 {
		return false
	}
	return board[y+oy][x+ox]
}

func main() {
}
