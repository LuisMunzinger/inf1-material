package db

import (
	"testing"
)

func TestGameDb_RecommendGames(t *testing.T) {
	db := testdata()

	games := db.RecommendGames("Alice")

	if len(games) != 1 {
		t.Errorf("Expected 1 recommended game, got %d", len(games))
	}

	if games[0].Title != "Valheim" {
		t.Errorf("Expected recommended game to be 'Valheim', got '%s'", games[0].Title)
	}
}
