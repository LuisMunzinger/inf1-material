package deck

import "cards/card"

// Repräsentiert einen Kartenstapel.
type Deck struct {
	cards []card.Card
}

// Empty erzeugt neues leeres Deck
func Empty() *Deck {
	return &Deck{}
}

// New32 erzeugt ein neues Skatblatt mit 32 Karten.
func New32() *Deck {
	// TODO
	return &Deck{}
}

// New52 erzeugt ein neues französisches Blatt mit 52 Karten.
func New52() *Deck {
	// TODO
	return &Deck{}
}

// Add fügt eine Karte zum Deck hinzu.
func (d *Deck) Add(c card.Card) {
	d.cards = append(d.cards, c)
}

// Shuffle mischt das Deck.
func (d *Deck) Shuffle() {
	// TODO
}

// Draw zieht die oberste Karte.
// Entfernt die Karte aus dem Deck und liefert sie zurück.
func (d *Deck) Draw() card.Card {
	// TODO
	return card.Card{}
}

// Top liefert die oberste Karte.
// Entfernt sie aber nicht.
func (d Deck) Top() card.Card {
	// TODO
	return card.Card{}
}
