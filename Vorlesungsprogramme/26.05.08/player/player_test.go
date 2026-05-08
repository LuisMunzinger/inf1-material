package player

import (
	"inf1-material/Vorlesungsprogramme/26.05.08/game"
	"testing"
)

func TestPlayer_HasPlayed(t *testing.T) {
	p := New("Alice")
	g := game.New("The Legend of Zelda", "Action-Adventure")

	if p.HasPlayed(g) {
		t.Errorf("Expected player to not have played the game")
	}

	p.PlayGame(g, 10)

	if !p.HasPlayed(g) {
		t.Errorf("Expected player to have played the game")
	}
}

func TestPlayer_HasPlayedMore(t *testing.T) {
	p := New("Alice")
	g := game.New("The Legend of Zelda", "Action-Adventure")

	if p.HasPlayedMore(g, 5) {
		t.Errorf("Expected player to not have played the game more than 5 hours")
	}

	p.PlayGame(g, 10)

	if !p.HasPlayedMore(g, 5) {
		t.Errorf("Expected player to have played the game for at least 5 hours")
	}

	if p.HasPlayedMore(g, 15) {
		t.Errorf("Expected player to not have played the game more than 15 hours")
	}
}
