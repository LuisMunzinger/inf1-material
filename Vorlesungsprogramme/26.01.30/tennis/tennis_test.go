package tennis

import "fmt"

func ExamplePrint() {
	s1 := NewSet()
	s1.NewGame()

	s1.Score(1)
	s1.Print()

	s1.Score(2)
	s1.Print()

	s1.Score(1)
	s1.Print()

	s1.Score(1)
	s1.Print()

	s1.NewGame()
	s1.Print()

	// Output:
}
func Example() {
	n1 := "Alcaraz"
	n2 := "Zverev"

	s1 := NewSet()

	g1 := Game{30, 40}
	s1.AddGame(g1)

	g2 := Game{40, 30}
	s1.AddGame(g2)
	g3 := Game{0, 40}
	s1.AddGame(g3)

	g4 := Game{40, 15}
	s1.AddGame(g4)

	g5 := Game{40, 15}
	s1.AddGame(g5)

	g6 := Game{40, 15}
	s1.AddGame(g6)

	g7 := Game{40, 50}
	s1.AddGame(g7)

	g8 := Game{40, 30}
	s1.AddGame(g8)

	g9 := Game{30, 40}
	s1.AddGame(g9)

	g10 := Game{40, 30}
	s1.AddGame(g10)

	fmt.Println(s1)
	fmt.Println(s1.games)
	fmt.Println(s1.games[0])
	fmt.Println(s1.games[0].p1)

	fmt.Printf("Spiel: %s gegen %s\n", n1, n2)
	s1.PrintResults()

	// Output:

}
