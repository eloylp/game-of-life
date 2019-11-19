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

func NewGame(initialBoard Board, generations, intervalSecs int) *game {
	return &game{board: initialBoard, generations: generations, intervalSec: intervalSecs}
}

func (g game) Run() {
	for i := 0; i <= g.generations; i++ {
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
