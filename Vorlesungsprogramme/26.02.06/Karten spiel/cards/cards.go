package cards

import (
	"cards/rank"
	"cards/suit"
	"fmt"
	"strings"
)

type Card struct {
	Rank rank.Rank
	Suit suit.Suit
}

func (c Card) String() string {
	rank := c.Rank.String()
	suit := c.Suit.String()

	lines := []string{
		"-------",
		fmt.Sprintf("|%-2s....|", rank),
		fmt.Sprintf("|..%s..|", suit),
		fmt.Sprintf("|....%2s|", rank),
		"-------",
	}
	return strings.Join(lines, "\n")
}
