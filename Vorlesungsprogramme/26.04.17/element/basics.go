package element

// IsEmpty gibt an, ob das Element leer ist.
// Ein Element ist leer, wenn es kein Nachfolger-Element hat.
// D.h. wenn der Nachfolger-Pointer nil ist.
func (e *Element) IsEmpty() bool {
	if e.next == nil {
		return true
	} else {
		return false
	}
}

// Value liefert den Wert des Elements.
// Für ein leeres Element wird eine panic ausgelöst.
func (e *Element) Value() int {
	if e.value == Empty().value {
		panic("value for empty element requested")
	}
	return e.value
}

// Next liefert das Nachfolger-Element.
// Für ein leeres Element wird eine panic ausgelöst.
func (e *Element) Next() *Element {
	if e.IsEmpty() {
		panic("Leeres Element")
	}
	return e.next
}

// SetValue setzt den Wert des Elements.
// Falls das Element bisher leer war, wird es mit dem gegebenen Wert initialisiert.
func (e *Element) SetValue(value int) {
	if e.IsEmpty() {
		e.next = Empty()
	}
	e.value = value
}
