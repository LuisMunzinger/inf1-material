package game

import "testing"

func TestGame_HasGenre(t *testing.T) {
	g := New("The Legend of Zelda", "Action-Adventure")

	if !g.HasGenre("Action-Adventure") {
		t.Errorf("Expected game to have genre 'Action-Adventure'")
	}

	if g.HasGenre("RPG") {
		t.Errorf("Expected game to not have genre 'RPG'")
	}
}
