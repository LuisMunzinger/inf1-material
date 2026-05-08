package db

import (
	"inf1-material/Vorlesungsprogramme/26.05.08/player"
)

// GetPlayer sucht einen Spieler in der Datenbank anhand eines Namens.
// Liefert einen Zeiger auf den Spieler zurück, wenn er gefunden wird,
// oder nil, wenn er nicht gefunden wird.
func (db *GameDb) GetPlayer(name string) *player.Player {
	// TODO
	return nil
}

// GetPlayersByGame sucht Spieler in der Datenbank, die ein bestimmtes Spiel gespielt haben.
// Erwartet den Titel des Spiels und die Mindestanzahl gespielter Stunden.
func (db *GameDb) GetPlayersByGame(title string, min_played int) []*player.Player {
	players := []*player.Player{}
	// TODO
	return players
}
