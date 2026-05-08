package db

import "inf1-material/Vorlesungsprogramme/26.05.08/game"

// RecommendGames erwartet einen Spielernamen und generiert Spiele-Empfehlungen.
//
// Die Funktion sucht nach Spielen, die der Spieler oft gespielt hat.
// Für diese Spiele werden Spiele mit gleichem Genre gesucht,
// die von anderen Spielern häufig gespielt wurden.
func (db *GameDb) RecommendGames(playerName string) []*game.Game {
	recommendedGames := []*game.Game{}
	// TODO
	return recommendedGames
}
