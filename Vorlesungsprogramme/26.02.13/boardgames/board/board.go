package board

import "fmt"

// Ein Spielfeld für Brettspiele.
type Board struct {
	rows [][]string
}

// New erstellt ein neues leeres Spielfeld mit der angegebenen Anzahl von Zeilen und Spalten.
// Die Zellen des Spielfelds werden mit Leerzeichen initialisiert.
func New(rows, cols int) *Board {
	b := &Board{
		rows: make([][]string, rows),
	}
	for i := range b.rows {
		b.rows[i] = make([]string, cols)
		for j := range b.rows[i] {
			b.rows[i][j] = " "
		}
	}
	return b
}

// NewWithNumbers erstellt ein neues Spielfeld mit der angegebenen Anzahl von Zeilen und Spalten.
// Füllt die Zellen des Spielfelds mit aufsteigenden Zahlen als Strings, beginnend bei "1".
func NewWithNumbers(rows, cols int) *Board {
	b := &Board{
		rows: make([][]string, rows),
	}
	num := 1
	for i := range b.rows {
		b.rows[i] = make([]string, cols)
		for j := range b.rows[i] {
			b.rows[i][j] = fmt.Sprintf("%d", num)
			num++
		}
	}
	return b
}

// Get liefert den Inhalt der Zelle an der angegebenen Position zurück.
// Liefert einen leeren String, falls die Position außerhalb des Spielfelds liegt.
func (b *Board) Get(row, col int) string {
	// TODO
	return ""
}

// Set setzt den Inhalt der Zelle an der angegebenen Position auf den angegebenen Wert.
// Ignoriert die Anweisung, falls die Position außerhalb des Spielfelds liegt.
func (b *Board) Set(row, col int, value string) {
	// TODO
}
