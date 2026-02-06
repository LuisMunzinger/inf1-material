package deck

import (
	"cards/card"
	"cards/rank"
	"cards/suit"
	"fmt"
)

func ExampleEmpty() {
	d := Empty()

	fmt.Println(len(d.cards))

	// Output:
	// 0
}

func ExampleDeck_Add() {
	d := Empty()

	d.Add(card.Card{Rank: rank.Queen, Suit: suit.Spades})
	fmt.Println(len(d.cards))

	// Output:
	// 1
}
