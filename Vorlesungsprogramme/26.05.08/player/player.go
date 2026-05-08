package player

import "inf1-material/Vorlesungsprogramme/26.05.08/game"

// Repräsentiert einen Spieler in einer Spiele-Datenbank.
type Player struct {
	Name        string
	playedGames map[string]int
}

// New erstellt einen neuen Spieler mit dem gegebenen Namen und Geburtsjahr.
func New(name string) *Player {
	return &Player{
		Name:        name,
		playedGames: map[string]int{},
	}
}

// PlayGame fügt ein Spiel zu den gespielten Spielen des Spielers hinzu und erhöht die Anzahl der gespielten Stunden.
func (p *Player) PlayGame(g *game.Game, hours int) {
	p.playedGames[g.Title] = hours
}

// HasPlayed prüft, ob der Spieler ein bestimmtes Spiel gespielt hat.
func (p *Player) HasPlayed(g *game.Game) bool {
	player := p.playedGames
	titel := p.playedGames["Titel"]

	for _, i := range player {
		if player[] == titel {
			return true
		}
	}
	return false
}

// HasPlayedMore prüft, ob der Spieler ein bestimmtes Spiel mindestens `hours` Stunden gespielt hat.
func (p *Player) HasPlayedMore(g *game.Game, hours int) bool {
	// TODO
	return false
}

// PlayedGames liefert eine Liste mit den Spielen, die der
// Spieler mehr als `hours` Stunden gespielt hat.
func (p *Player) PlayedGames(hours int) []string {
	games := []string{}
	// TODO
	return games
}
