package rank

// Rank definiert die Werte der Karten.
type Rank int

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

// String gibt den Wert als String zurück und verwendet dabei die üblichen Symbole für die Werte.
func (r Rank) String() string {
	strings := []string{
		"2", "3", "4", "5", "6", "7", "8", "9", "10",
		"J", "Q", "K", "A",
	}
	if r >= Two && r <= Ace {
		return strings[r]
	}
	return "Unknown Rank"
}
