package rank

type Rank int

const (
	Tow Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	King
	Queen
	Ace
)

func (r Rank) string() string {
	switch r {
	case Tow:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "j"
	case King:
		return "K"
	case Queen:
		return "Q"
	case Ace:
		return "A"
	default:
		return "Unkown Rank"
	}
}

func (r Rank) String() string {
	strings := []string{
		"2", "3", "4", "5", "6", "7", "8", "9", "10",
		"J", "K", "Q", "A",
	}
	if r == strings 
	

}
