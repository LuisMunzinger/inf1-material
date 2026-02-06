package scheresteinpapier

type Value int

const (
	Rock Value = iota
	Paper
	Scissors
	Spock
	Lizard
)

// String liefert den Wert als String.
func (v Value) String() string {
	switch v {
	case Rock:
		return "Stein"
	case Paper:
		return "Papier"
	case Scissors:
		return "Schere"
	case Spock:
		return "Spock"
	case Lizard:
		return "Eidechse"
	default:
		return "Unkown Suit"
	}
}

// Beats gibt an, ob der Wert v den Wert w schlägt.
func (v Value) Beats(w Value) bool {
	switch v {
	case Rock:
		return w == Scissors || w == Lizard
	case Paper:
		return w == Rock || w == Spock
	case Scissors:
		return w == Paper || w == Lizard
	case Spock:
		return w == Rock || w == Scissors
	case Lizard:
		return w == Paper || w == Spock
	}
	return false
}
