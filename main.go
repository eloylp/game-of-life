package main

func main() {
	b := Board{
		{false, false, false, false, true},
		{false, true, false, false, true},
		{false, true, true, false, false},
		{false, true, false, false, true},
	}
	g := NewGame(b, 10, 1)
	g.Run()
}
