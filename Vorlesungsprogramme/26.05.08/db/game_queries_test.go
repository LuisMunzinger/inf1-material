package db

import (
	"testing"
)

func TestGameDb_GetPlayedGames(t *testing.T) {
	db := testdata()

	games := db.GetPlayedGames("Alice", 50)

	if len(games) != 2 {
		t.Errorf("Expected 2 played games, got %d", len(games))
	}

	expectedTitles := map[string]bool{
		"The Legend of Zelda": true,
		"Minecraft":           true,
	}

	for _, g := range games {
		if !expectedTitles[g.Title] {
			t.Errorf("Unexpected game '%s' in played games", g.Title)
		}
	}
}

func TestGameDb_GetGamesByGenre(t *testing.T) {
	db := testdata()

	games := db.GetGamesByGenre("Survival")

	if len(games) != 2 {
		t.Errorf("Expected 2 games in genre 'Survival', got %d", len(games))
	}

	expectedTitles := map[string]bool{
		"Minecraft": true,
		"Valheim":   true,
	}

	for _, g := range games {
		if !expectedTitles[g.Title] {
			t.Errorf("Unexpected game '%s' in genre 'Survival'", g.Title)
		}
	}
}
