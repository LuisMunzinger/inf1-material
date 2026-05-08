package db

import (
	"inf1-material/Vorlesungsprogramme/26.05.08/game"
)

// GetPlayedGames sucht alle Spiele, die ein gegebener Spieler gespielt hat.
// Erwartet dabei den Namen des Spielers und die Mindestanzahl gespielter Stunden.
func (db *GameDb) GetPlayedGames(name string, min_played int) []*game.Game {
	games := []*game.Game{}

	for _, spieler := range db.Players {
		if name == spieler.Name {
			spiele := spieler.PlayedGames(min_played)

			_ = spiele
		}
	}

	return games
}

// GetGamesByGenre sucht Spiele in der Datenbank anhand ihres Genres.
func (db *GameDb) GetGamesByGenre(genre string) []*game.Game {
	games := []*game.Game{}
	// TODO
	return games
}
