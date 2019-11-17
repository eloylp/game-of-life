package main

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
