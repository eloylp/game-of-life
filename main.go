package main

import (
	"fmt"
	"log"
)

var Version string

func main() {

	cfg := Config()
	if err := CheckConfig(cfg); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conway`s game of life")
	fmt.Println(Version)
	b := RandomBoard(cfg.boardX, cfg.boardY)
	g := NewGame(b, cfg.generations, cfg.intervalSecs)
	g.Run()
}
