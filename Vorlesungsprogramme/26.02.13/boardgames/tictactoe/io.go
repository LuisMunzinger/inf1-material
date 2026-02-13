package tictactoe

import (
	"fmt"
)

// Show zeigt den aktuellen Spielstand an.
func (g *Game) Show() {
	fmt.Println(g.board)
}

// Move fragt den Spieler nach einem Zug und führt diesen Zug aus, falls er erlaubt ist.
// Falls der Zug nicht erlaubt ist, wird der Spieler erneut nach einem Zug gefragt.
func (g *Game) Move(player string) {
	// TODO
}
