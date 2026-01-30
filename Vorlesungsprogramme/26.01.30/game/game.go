package game

import (
	"fmt"
)

type Game struct {
	p1 int
	p2 int
}

// New erstellt ein neues Tennis-Spiel mit dem Punktestand 0:0.
func New() Game {
	return Game{0, 0}
}

// String gibt den aktuellen Punktestand des Spiels als String in der Form "P1:P2" zurück.
// Falls ein Spieler im Vorteil ist, wird "Advantage Player 1" zurückgegeben.
// Falls ein Spieler gewonnen hat, wird "Player X wins" zurückgegeben.
func (g Game) String() string {
	if g.P1Wins() {
		return "Player 1 wins"
	}
	if g.P2Wins() {
		return "Player 2 wins"
	}
	if g.AdvantageP1() {
		return "Advantage Player 1"
	}
	if g.AdvantageP2() {
		return "Advantage Player 2"
	}

	return fmt.Sprintf("%d:%d", g.p1, g.p2)
}

// ScoreP1 erhöht den Punktestand von Spieler 1 um einen Punkt.
// Wenn Spieler 1 gewinnt, wird sein Punktestand auf "Win" gesetzt.
func (g *Game) ScoreP1() {
	g.p1++
}

// ScoreP2 erhöht den Punktestand von Spieler 2 um einen Punkt.
// Wenn Spieler 2 gewinnt, wird sein Punktestand auf "Win" gesetzt.
func (g *Game) ScoreP2() {
	g.p2++
}

// AdvantageP1 überprüft, ob Spieler 1 im Vorteil ist.
func (g Game) AdvantageP1() bool {
	return g.p1 >= 4 && g.p1-g.p2 == 1
}

// AdvantageP2 überprüft, ob Spieler 2 im Vorteil ist.
func (g Game) AdvantageP2() bool {
	return g.p2 >= 4 && g.p2-g.p1 == 1
}

// P1Wins überprüft, ob Spieler 1 das Spiel gewonnen hat.
func (g Game) P1Wins() bool {
	return g.p1 > 3 && g.p1-g.p2 >= 2
}

// P2Wins überprüft, ob Spieler 2 das Spiel gewonnen hat.
func (g Game) P2Wins() bool {
	return g.p2 > 3 && g.p2-g.p1 >= 2
}
