package card

import (
	"cards/rank"
	"cards/suit"
	"fmt"
	"strings"
)

// Card repräsentiert eine Spielkarte mit einem Wert und einer Farbe.
type Card struct {
	Rank rank.Rank
	Suit suit.Suit
}

// Front gibt die Karte als String zurück.
// Dabei soll die Karte in AsciiArt mit
// einem Rahmen dargestellt werden, z.B.:
// ┌─────┐
// │A    │
// │  ♠  │
// │    A│
// └─────┘
func (c Card) Front() string {
	rank := c.Rank
	suit := c.Suit

	lines := []string{
		"┌─────┐",
		fmt.Sprintf("│%-2s   │", rank),
		fmt.Sprintf("│  %s  │", suit),
		fmt.Sprintf("│   %2s│", rank),
		"└─────┘",
	}
	return strings.Join(lines, "\n")
}

// Back gibt die Rückseite der Karte als String zurück.
func (c Card) Back() string {
	lines := []string{
		"┌─────┐",
		"│\\   /│",
		"│ | | │",
		"│/   \\│",
		"└─────┘",
	}
	return strings.Join(lines, "\n")
}

// MatchesRank prüft, ob die Karte den gleichen Wert wie eine andere Karte hat.
func (c Card) MatchesRank(other Card) bool {
	return c.Rank == other.Rank
}

// MatchesSuit prüft, ob die Karte die gleiche Farbe wie eine andere Karte hat.
func (c Card) MatchesSuit(other Card) bool {
	return c.Suit == other.Suit
}
