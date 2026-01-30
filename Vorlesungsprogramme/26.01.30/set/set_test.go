package tennis

import (
	"fmt"
	"inf1-material/Vorlesungsprogramme/26.01.30/game"
)

func ExampleSet_CurrentGameNo() {
	s := New()
	fmt.Println(s.CurrentGameNo())

	s.games = append(s.games, game.New())
	fmt.Println(s.CurrentGameNo())

	// Output:
	// 1
	// 2
}

func ExampleSet_CurrentGame() {
	s := New()
	fmt.Println(s.CurrentGame())

	s.CurrentGame().ScoreP1()
	fmt.Println(s.CurrentGame())

	s.games = append(s.games, game.New())
	fmt.Println(s.CurrentGame())

	// Output:
	// 0:0
	// 1:0
	// 0:0
}

func ExampleSet_String() {
	s := New()
	fmt.Println(s.String())

	s.ScoreP1()
	s.ScoreP1()
	s.ScoreP2()
	fmt.Println(s.String())

	s.ScoreP1()
	s.ScoreP1()
	fmt.Println(s.String())

	// Output:
	// 0-0
	// 0:0
	//
	// 0-0
	// 2:1
	//
	// 1-0
	// Player 1 wins
}
