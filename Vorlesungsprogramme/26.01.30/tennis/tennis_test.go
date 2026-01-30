package tennis

import "fmt"

func Example() {
	n1 := "Alcaraz"
	n2 := "Zverev"

	g1 := Game{30, 40}
	g2 := Game{40, 30}
	g3 := Game{0, 40}
	g4 := Game{40, 15}
	g5 := Game{40, 15}
	g6 := Game{40, 15}
	g7 := Game{40, 50}
	g8 := Game{40, 30}
	g9 := Game{30, 40}
	g10 := Game{40, 30}

	s1 := Set{[]Game{g1, g2, g3, g4, g5, g6, g7, g8, g9, g10}}

	fmt.Println(s1)
	fmt.Println(s1.games[0])
	fmt.Println(s1.games[0].p1)

	fmt.Printf("Spiel: %s gegen %s\n", n1, n2)
	fmt.Printf("Ergebns von Spiel 1: %d : %d\n", g1.p1, g1.p2)
	fmt.Printf("Ergebns von Spiel 2: %d : %d\n", g2.p1, g2.p2)
	fmt.Printf("Ergebns von Spiel 3: %d : %d\n", g3.p1, g3.p2)
	fmt.Printf("Ergebns von Spiel 4: %d : %d\n", g4.p1, g4.p2)
	fmt.Printf("Ergebns von Spiel 5: %d : %d\n", g5.p1, g5.p2)
	fmt.Printf("Ergebns von Spiel 6: %d : %d\n", g6.p1, g6.p2)
	fmt.Printf("Ergebns von Spiel 7: %d : %d\n", g7.p1, g7.p2)
	fmt.Printf("Ergebns von Spiel 8: %d : %d\n", g8.p1, g8.p2)
	fmt.Printf("Ergebns von Spiel 9: %d : %d\n", g9.p1, g9.p2)
	fmt.Printf("Ergebns von Spiel 10: %d : %d\n", g10.p1, g10.p2)

	// Output:

}
