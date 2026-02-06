package suit

// Suit definiert die Farben der Karten.
type Suit int

const (
	Clubs Suit = iota
	Spades
	Hearts
	Diamonds
)

// String gibt die Farbe als String zurück
// und verwendet dabei die üblichen Symbole für die Farben.
func (s Suit) String() string {
	switch s {
	case Clubs:
		return "♣"
	case Spades:
		return "♠"
	case Hearts:
		return "♥"
	case Diamonds:
		return "♦"
	default:
		return "Unknown Suit"
	}
}
