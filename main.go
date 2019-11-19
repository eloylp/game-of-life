package main

func main() {

	cfg := Config()
	b := RandomBoard(cfg.boardX, cfg.boardY)
	g := NewGame(b, cfg.generations, cfg.intervalSecs)
	g.Run()
}
