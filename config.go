package main

import (
	"errors"
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

func CheckConfig(cfg Cfg) error {
	if cfg.boardX < 1 {
		return errors.New("Cannot set X axis under 1")
	}
	if cfg.boardY < 1 {
		return errors.New("Cannot set Y axis under 1")
	}
	if cfg.generations < 1 {
		return errors.New("Cannot set generations  under 1")
	}
	if cfg.intervalSecs < 1 {
		return errors.New("Cannot set intervals  under 1")
	}
	return nil
}
