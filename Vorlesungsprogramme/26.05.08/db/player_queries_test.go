package db

import (
	"testing"
)

func TestGameDb_GetPlayer(t *testing.T) {
	db := testdata()

	player := db.GetPlayer("Alice")

	if player == nil {
		t.Errorf("Expected to find player 'Alice', got nil")
	} else if player.Name != "Alice" {
		t.Errorf("Expected player name to be 'Alice', got '%s'", player.Name)
	}
}

func TestGameDb_GetPlayersByGame_zelda(t *testing.T) {
	db := testdata()

	players := db.GetPlayersByGame("The Legend of Zelda", 50)

	if len(players) != 2 {
		t.Errorf("Expected 2 players who played 'The Legend of Zelda' for at least 50 hours, got %d", len(players))
	}

	expectedNames := map[string]bool{
		"Alice":   true,
		"Charlie": true,
	}

	for _, player := range players {
		if !expectedNames[player.Name] {
			t.Errorf("Unexpected player '%s' in players who played 'The Legend of Zelda'", player.Name)
		}
	}
}

func TestGameDb_GetPlayersByGame_valheim(t *testing.T) {
	db := testdata()

	players := db.GetPlayersByGame("Valheim", 50)

	if len(players) != 2 {
		t.Errorf("Expected 2 players who played 'Valheim' for at least 50 hours, got %d", len(players))
	}

	expectedNames := map[string]bool{
		"Bob":     true,
		"Charlie": true,
	}

	for _, player := range players {
		if !expectedNames[player.Name] {
			t.Errorf("Unexpected player '%s' in players who played 'Valheim'", player.Name)
		}
	}
}

func TestGameDb_GetPlayersByGame_minecraft(t *testing.T) {
	db := testdata()

	players := db.GetPlayersByGame("Minecraft", 50)

	if len(players) != 1 {
		t.Errorf("Expected 1 player who played 'Minecraft' for at least 50 hours, got %d", len(players))
	}

	expectedNames := map[string]bool{
		"Alice": true,
	}

	for _, player := range players {
		if !expectedNames[player.Name] {
			t.Errorf("Unexpected player '%s' in players who played 'Minecraft'", player.Name)
		}
	}
}
