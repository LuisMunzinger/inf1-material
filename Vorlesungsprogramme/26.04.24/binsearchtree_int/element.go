package bintree_int

import (
	"bintree_int"
)

// Diese Datei definiert die Element-Struktur für einen binären Suchbaum,
// der Integer-Werte speichert. Sie erweitert die grundlegende Element-Struktur
// aus dem bintree_int-Paket um die spezifischen Eigenschaften eines binären Suchbaums.
//
// Hier werden außerdem einige Methoden für den grundlegenden Zugriff auf die Kinder
// definiert. Wenn Sie den Baum verwenden, sollten Sie diese Methoden nutzen,
// da ein direkter Zugriff auf die eingebettete Element-Struktur nicht funktioniert.
// Dies liegt daran, dass die privaten Felder der eingebetteten Struktur nicht direkt
// zugänglich sind, weil wir hier in einem anderen Package arbeiten.

// Ein Element eines binären Suchbaums, das einen Integer-Wert speichert.
type Element struct {
	*bintree_int.Element
}

// Empty erstellt ein neues leeres Element.
func Empty() *Element {
	return &Element{Element: bintree_int.Empty()}
}

// New erstellt ein neues Element mit einem gegebenen Wert.
func New(value int) *Element {
	return &Element{Element: bintree_int.New(value)}
}

// Left gibt das linke Kind des Elements zurück.
func (e *Element) Left() *Element {
	return &Element{Element: e.Element.Left()}
}

// Right gibt das rechte Kind des Elements zurück.
func (e *Element) Right() *Element {
	return &Element{Element: e.Element.Right()}
}
