package element

// Ein Element einer einfach verketteten Liste.
type Element struct {
	value int
	next  *Element
}

// Empty liefert ein neues leeres Element.
func Empty() *Element {
	return &Element{value: 0, next: nil}
}

// New liefert ein neues Element mit dem gegebenen Wert.
// Das Nachfolger-Element wird als leeres Element initialisiert.
func New(value int) *Element {
	return &Element{value: value, next: Empty()}
}
