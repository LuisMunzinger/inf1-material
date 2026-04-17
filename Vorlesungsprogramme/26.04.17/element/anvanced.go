package element

// Append fügt ein neues Element mit dem gegebenen Wert am Ende der Liste ein.
func (e *Element) Append(value int) {
	// TODO
}

// Length gibt die Anzahl der Elemente in der Liste zurück.
func (e *Element) Length() int {
	// TODO
	return 0
}

// Contains gibt an, ob ein Element mit dem gegebenen Wert in der Liste enthalten ist.
func (e *Element) Contains(value int) bool {
	// TODO
	return false
}

// Count gibt die Anzahl der Elemente in der Liste zurück, die den gegebenen Wert enthalten.
func (e *Element) Count(value int) int {
	// TODO
	return 0
}

// Sum berechnet die Summe der Werte aller Elemente in der Liste.
func (e *Element) Sum() int {
	// TODO
	return 0
}

// Min gibt den kleinsten Wert aller Elemente in der Liste zurück.
// Falls die Liste leer ist, wird eine panic ausgelöst.
func (e *Element) Min() int {
	// TODO
	return 0
}

// Last gibt das letzte Element der Liste zurück.
// Falls die Liste leer ist, wird eine panic ausgelöst.
func (e *Element) Last() *Element {
	// TODO
	return nil
}

// At gibt das Element an der gegebenen Position zurück.
// Falls die Position nicht existiert, wird eine panic ausgelöst.
func (e *Element) At(position int) *Element {
	// TODO
	return nil
}

// String gibt eine textuelle Repräsentation der Liste zurück.
// Die Elemente werden durch " -> " getrennt.
// Falls die Liste leer ist, wird ein leerer String zurückgegeben.
func (e *Element) String() string {
	// TODO
	return ""
}

// Swap vertauscht die beiden Elemente an den Stellen i und j.
// Falls eine der Positionen nicht existiert, wird eine panic ausgelöst.
// Die Funktion liefert den neuen Kopf der Liste zurück.

// Es sollen nicht die Werte der Elemente vertauscht werden, sondern die Elemente selbst.
func (e *Element) Swap(i, j int) *Element {
	// TODO
	return e

}
