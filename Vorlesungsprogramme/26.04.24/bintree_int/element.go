package bintree_int

// Ein Element eines binären Baums, das einen Integer-Wert speichert.
type Element struct {
	value int
	left  *Element
	right *Element
}

// Empty erstellt ein neues leeres Element.
func Empty() *Element {
	return &Element{}
}

// New erstellt ein neues Element mit einem gegebenen Wert.
func New(value int) *Element {
	return &Element{value: value, left: Empty(), right: Empty()}
}

// Value gibt den Wert des Elements zurück.
func (e *Element) Value() int {
	return e.value
}

// Left gibt das linke Kind des Elements zurück.
func (e *Element) Left() *Element {
	return e.left
}

// Right gibt das rechte Kind des Elements zurück.
func (e *Element) Right() *Element {
	return e.right
}
