package tictactoe

import "fmt"

func ExampleGame_run() {

	// Spielstart
	g := NewGame()
	fmt.Println(g.board)
	fmt.Println(g.PlayerWins("X"))
	fmt.Println(g.PlayerWins("O"))
	fmt.Println(g.GameOver())
	fmt.Println(g.IsDraw())
	fmt.Println()

	// Ein erlaubter Zug von Spieler "X".
	g.Set(1, 1, "X")
	fmt.Println(g.board)
	fmt.Println(g.PlayerWins("X"))
	fmt.Println(g.PlayerWins("O"))
	fmt.Println(g.GameOver())
	fmt.Println(g.IsDraw())
	fmt.Println()

	// Spieler "O" macht erst einen unerlaubten und dann einen erlaubten Zug.
	g.Set(1, 1, "O")
	g.Set(0, 0, "O")
	fmt.Println(g.board)
	fmt.Println(g.PlayerWins("X"))
	fmt.Println(g.PlayerWins("O"))
	fmt.Println(g.GameOver())
	fmt.Println(g.IsDraw())
	fmt.Println()

	// Beide Spieler machen weitere Züge, bis Spieler "X" gewinnt.
	g.Set(0, 2, "X")
	g.Set(2, 0, "O")
	g.Set(1, 0, "X")
	g.Set(2, 2, "O")
	g.Set(0, 1, "X")
	g.Set(2, 2, "O")
	g.Set(2, 1, "X")
	fmt.Println(g.board)
	fmt.Println(g.PlayerWins("X"))
	fmt.Println(g.PlayerWins("O"))
	fmt.Println(g.GameOver())
	fmt.Println(g.IsDraw())

	// Output:
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// false
	// false
	// false
	// false
	//
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// |   | X |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// false
	// false
	// false
	// false
	//
	// +---+---+---+
	// | O |   |   |
	// +---+---+---+
	// |   | X |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// false
	// false
	// false
	// false
	//
	// +---+---+---+
	// | O | X | X |
	// +---+---+---+
	// | X | X |   |
	// +---+---+---+
	// | O | X | O |
	// +---+---+---+
	// true
	// false
	// true
	// false
}
