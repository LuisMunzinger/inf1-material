package bintree_int

// IsEmpty liefert true zurück, wenn das Element leer ist.
//
// Ein Element gilt als leer, wenn es keine Kinder hat.
// D.h. wenn die Kind-Pointer Left und Right beide nil sind.
// Der Wert des Elements spielt dabei keine Rolle.
func (e *Element) IsEmpty() bool {
	// TODO
	return false
}

// IsLeaf liefert true zurück, wenn das Element ein Blatt ist.
//
// Ein Element gilt als Blatt, wenn beide Kinder leer sind.
func (e *Element) IsLeaf() bool {
	// TODO
	return false
}

// Count zählt die Anzahl der Elemente im Baum, beginnend bei diesem Element.
// Dabei werden nur nicht-leere Elemente gezählt.
func (e *Element) Count() int {
	// TODO
	return 0
}

// Height berechnet die Höhe des Baums.
// Ein leeres Element hat die Höhe 0, ein Blatt hat die Höhe 1, usw.
func (e *Element) Height() int {
	// TODO
	return 0
}
