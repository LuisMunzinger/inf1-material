package tictactoe

import "inf1-material/Vorlesungsprogramme/26.02.13/boardgames/board"

type Game struct {
	board *board.Board
}

// NewGame erstellt ein neues Spiel mit einem leeren 3x3-Spielfeld.
func NewGame() *Game {
	return &Game{
		board: board.New(3, 3),
	}
}

// Get liefert den Inhalt der Zelle an der angegebenen Position zurück.
// Liefert einen leeren String, falls die Position außerhalb des Spielfelds liegt.
func (g *Game) Get(row, col int) string {
	// TODO
	return ""
}

// Set setzt den Inhalt der Zelle an der angegebenen Position auf den angegebenen Wert.
// Ignoriert die Anweisung, falls der Zug nicht erlaubt ist.
func (g *Game) Set(row, col int, value string) {
	// TODO
}
