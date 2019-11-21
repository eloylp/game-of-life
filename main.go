package main

import (
	"log"
)

func main() {

	cfg := Config()
	if err := CheckConfig(cfg); err != nil {
		log.Fatal(err)
	}
	b := RandomBoard(cfg.boardX, cfg.boardY)
	g := NewGame(b, cfg.generations, cfg.intervalSecs)
	g.Run()
}
