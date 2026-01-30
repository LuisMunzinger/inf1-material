package game

import "fmt"

func ExampleGame_scores1() {
	g := New()
	g.ScoreP1()
	fmt.Println(g)
	g.ScoreP1()
	fmt.Println(g)
	g.ScoreP2()
	fmt.Println(g)
	g.ScoreP1()
	fmt.Println(g)
	g.ScoreP1()
	fmt.Println(g)

	fmt.Println(g.P1Wins())
	fmt.Println(g.P2Wins())

	// Output:
	// 1:0
	// 2:0
	// 2:1
	// 3:1
	// Player 1 wins
	// true
	// false
}

func ExampleGame_scores2() {
	g := New()
	g.ScoreP1()
	g.ScoreP1()
	g.ScoreP2()
	g.ScoreP2()
	g.ScoreP1()
	g.ScoreP2()

	fmt.Println(g)

	fmt.Println(g.P1Wins())
	fmt.Println(g.P2Wins())

	// Output:
	// 3:3
	// false
	// false
}
