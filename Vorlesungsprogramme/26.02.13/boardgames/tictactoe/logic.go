package tictactoe

// PlayerWins prüft, ob der angegebene Spieler das Spiel gewonnen hat.
// Der Spieler wird dabei durch das Zeichen repräsentiert,
// das er auf dem Spielfeld verwendet (z.B. "X" oder "O").
func (g *Game) PlayerWins(player string) bool {
	// TODO
	return false
}

// GameOver prüft, ob das Spiel vorbei ist, d.h. ob ein Spieler gewonnen hat oder das Spielfeld voll ist.
func (g *Game) GameOver() bool {
	// TODO
	return false
}

// IsDraw prüft, ob das Spiel unentschieden ist, d.h. ob das Spielfeld voll ist und kein Spieler gewonnen hat.
func (g *Game) IsDraw() bool {
	// TODO
	return false
}

// MoveAllowed prüft, ob ein Zug an der angegebenen Position erlaubt ist.
// Ein Zug ist erlaubt, wenn die Position innerhalb des Spielfelds liegt und die Zelle an dieser Position leer ist.
func (g *Game) MoveAllowed(row, col int) bool {
	// TODO
	return false
}
