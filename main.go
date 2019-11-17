package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"os/exec"
	"time"
)

type game struct {
	board       Board
	generations int
	intervalSec int
}

func (g game) Run() {
	for i := 0; i < g.generations; i++ {
		g.clearTerminal()
		g.outputMatrix(i)
		time.Sleep(time.Second * time.Duration(g.intervalSec))
		g.board = g.board.NextGen()
	}
}

func (g game) outputMatrix(generationNum int) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(true)
	table.SetColWidth(10)
	table.SetRowLine(true)
	table.SetCaption(true, fmt.Sprintf("Generation %v", generationNum))
	table.AppendBulk(g.board.TextBoard())
	table.Render()
}

func (g game) clearTerminal() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func NewGame(initialBoard Board, generations, intervalSecs int) *game {
	return &game{board: initialBoard, generations: generations, intervalSec: intervalSecs}
}

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

func (b Board) TextBoard() [][]string {
	newBoard := make([][]string, 4)
	for y := 0; y < len(b); y++ {
		for x := 0; x < len(b[x]); x++ {
			c := b[y][x]
			if c {
				newBoard[y] = append(newBoard[y], "x")
			} else {
				newBoard[y] = append(newBoard[y], " ")
			}
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
	b := Board{
		{false, false, false, false},
		{false, true, false, false},
		{false, true, true, false},
		{false, true, false, false},
	}
	g := NewGame(b, 5, 1)
	g.Run()
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
