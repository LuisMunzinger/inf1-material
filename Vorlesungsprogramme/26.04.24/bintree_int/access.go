package bintree_int

import "fmt"

// SetValue setzt den Wert des Elements auf den gegebenen Wert.
// Falls das Element vorher leer war, werden leere Kind-Elemente erstellt.
func (e *Element) SetValue(value int) {
	// TODO
}

// GetValueAt gibt den Wert des Elements zurück, das durch den gegebenen Pfad erreicht wird.
// Beispiele:
// - "LR": Gehe einmal nach links und dann einmal nach rechts.
// - "LLR": Gehe zwei Mal nach links und dann einmal nach rechts.
// - "": Gib den Wert der Wurzel zurück.
//
// Als zweiter Rückgabewert wird ein Fehlerwert geliefert.
// Wenn der Pfad zu einem leeren Element führt oder ein Element auf dem Pfad nicht existiert,
// liefert die Funktion 0 und einen Fehler, andernfalls den Wert des Elements und nil.
func (e *Element) GetValueAt(path string) (int, error) {
	// TODO
	return 0, fmt.Errorf("Not implemented")
}

// InsertValueAt fügt einen neuen Wert in den Baum ein.
// Die Stelle, an der eingefügt wird, wird durch den gegebenen Pfad bestimmt.
// Beispiele:
// - "LR": Gehe einmal nach links und dann einmal nach rechts.
// - "LLR": Gehe zwei Mal nach links und dann einmal nach rechts.
// - "": Füge direkt in der Wurzel ein.
//
// Wenn der Pfad zu einem leeren Element führt, wird dort ein neues Element
// mit dem gegebenen Wert erstellt.
// Wenn ein Element auf dem Pfad nicht existiert (außer dem letzten),
// liefert die Funktion einen Fehler, andernfalls true.
func (e *Element) InsertValueAt(path string, value int) error {
	// TODO
	return fmt.Errorf("Not implemented")
}
