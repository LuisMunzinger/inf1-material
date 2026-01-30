package tennis

import "fmt"

type Game struct {
	p1 int
	p2 int
}

type Set struct {
	games []Game
}

func NewSet() Set {
	return Set{[]Game{}}
}

func (s *Set) AddGame(g Game) {
	s.games = append(s.games, g)
}

func (s *Set) NewGame() {
	g := Game{0, 0}
	s.games = append(s.games, g)
}

func (s *Set) Print() {
	current_game := len(s.games) - 1
	g := s.games[current_game]
	fmt.Printf("%d : %d\n", g.p1, g.p2)
}

func (s *Set) Score(p int) {
	current_game := len((s.games)) - 1
	if p == 1 {
		if s.games[current_game].p1 < 30 {
			s.games[current_game].p1 += 15
		} else {
			s.games[current_game].p1 += 10
		}
	} else {
		if s.games[current_game].p2 < 30 {
			s.games[current_game].p2 += 15
		} else {
			s.games[current_game].p2 += 10
		}
	}
}

func (s Set) PrintResults() {
	for i, g := range s.games {
		fmt.Printf("Ergebnis von Spiel %d: %d: %d\n", i+1, g.p1, g.p2)
	}
}
