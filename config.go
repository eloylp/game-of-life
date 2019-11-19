package main

import (
	"flag"
)

type Cfg struct {
	boardX, boardY, generations, intervalSecs int
}

func Config() Cfg {

	c := Cfg{}

	flag.IntVar(&c.boardX, "x", 50, "Size of row for the board")
	flag.IntVar(&c.boardY, "y", 5, "Size of col for the board")
	flag.IntVar(&c.generations, "g", 40, "Num of generations to process")
	flag.IntVar(&c.intervalSecs, "i", 1, "Interval secs for board generation")
	flag.Parse()

	return c
}
