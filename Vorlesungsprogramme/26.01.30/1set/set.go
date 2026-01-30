package tennis

import game "inf1-material/Vorlesungsprogramme/26.01.30/1game"

type Set struct {
	games []game.Game
}

// New erstellt einen neuen Tennis-Satz.
// Der Satz beginnt mit einem neuen Spiel.
func New() Set {
	return Set{games: []game.Game{game.New()}}
}

// CurrentGameNo gibt die Nummer des aktuellen Spiels im Satz zurück.
func (s Set) CurrentGameNo() int {
	return len(s.games)
}

// CurrentGame gibt das aktuelle Spiel im Satz zurück.
func (s Set) CurrentGame() *game.Game {
	return &s.games[len(s.games)-1]
}

// ScoreP1 erhöht den Punktestand von Spieler 1 im aktuellen Spiel.
func (s *Set) ScoreP1() {
	// TODO
}

// ScoreP2 erhöht den Punktestand von Spieler 2 im aktuellen Spiel.
func (s *Set) ScoreP2() {
	// TODO
}

// CurrentGamesP1 gibt die Anzahl der gewonnenen Spiele von Spieler 1 im Satz zurück.
func (s Set) CurrentGamesP1() int {
	return 0
}

// CurrentGamesP2 gibt die Anzahl der gewonnenen Spiele von Spieler 2 im Satz zurück.
func (s Set) CurrentGamesP2() int {
	// TODO
	return 0
}

// String gibt den aktuellen Satzstand als String zurück.
// Dabei wird zunächst der Stand der fertigen Spiele angezeigt,
// gefolgt vom Punktestand des aktuellen Spiels.
func (s Set) String() string {
	// TODO
	return ""
}
