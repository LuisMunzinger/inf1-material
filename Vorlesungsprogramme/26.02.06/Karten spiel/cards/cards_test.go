package card

import (
	"cards/rank"
	"cards/suit"
	"fmt"
)

func ExampleCard_Front() {
	c1 := Card{Rank: rank.Ace, Suit: suit.Spades}
	fmt.Println(c1.Front())
	c2 := Card{Rank: rank.Ten, Suit: suit.Hearts}
	fmt.Println(c2.Front())

	// Output:
	// ┌─────┐
	// │A    │
	// │  ♠  │
	// │    A│
	// └─────┘
	// ┌─────┐
	// │10   │
	// │  ♥  │
	// │   10│
	// └─────┘
}

func ExampleCard_Back() {
	c := Card{Rank: rank.King, Suit: suit.Diamonds}
	fmt.Println(c.Back())

	// Output:
	// ┌─────┐
	// │\   /│
	// │ | | │
	// │/   \│
	// └─────┘
}

func ExampleCard_MatchesRank() {
	c1 := Card{Rank: rank.Queen, Suit: suit.Clubs}
	c2 := Card{Rank: rank.Queen, Suit: suit.Hearts}
	c3 := Card{Rank: rank.Jack, Suit: suit.Clubs}

	fmt.Println(c1.MatchesRank(c2))
	fmt.Println(c1.MatchesRank(c3))
	fmt.Println(c2.MatchesRank(c3))
	fmt.Println(c1.MatchesRank(c1))
	fmt.Println(c2.MatchesRank(c2))
	fmt.Println(c3.MatchesRank(c3))

	// Output:
	// true
	// false
	// false
	// true
	// true
	// true
}

func ExampleCard_MatchesSuit() {
	c1 := Card{Rank: rank.Queen, Suit: suit.Clubs}
	c2 := Card{Rank: rank.Jack, Suit: suit.Clubs}
	c3 := Card{Rank: rank.Ace, Suit: suit.Hearts}

	fmt.Println(c1.MatchesSuit(c2))
	fmt.Println(c1.MatchesSuit(c3))
	fmt.Println(c2.MatchesSuit(c3))
	fmt.Println(c1.MatchesSuit(c1))
	fmt.Println(c2.MatchesSuit(c2))
	fmt.Println(c3.MatchesSuit(c3))

	// Output:
	// true
	// false
	// false
	// true
	// true
	// true
}
