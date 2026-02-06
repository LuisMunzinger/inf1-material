package suit

type Suit int

const (
	Clubs Suit = iota
	Spades
	Hearts
	Diamonds
)

func (s Suit) string() string {
	switch s {
	case Clubs:
		return "Peak"
	case Spades:
		return "Schipe"
	case Hearts:
		return "Herz"
	case Diamonds:
		return "Karo"
	default:
		return "Unkown Suit"
	}
}
