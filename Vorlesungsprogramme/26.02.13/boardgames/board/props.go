package board

// MaxCellWidth bestimmt die maximale Breite des Inhalts einer Zelle.
func (b *Board) MaxCellWidth() int {
	maxWidth := 0
	for _, row := range b.rows {
		for _, cell := range row {
			if len(cell) > maxWidth {
				maxWidth = len(cell)
			}
		}
	}
	return maxWidth
}

// Width bestimmt die Breite des Spielfelds, d.h. die Anzahl der Spalten.
func (b *Board) Width() int {
	return len(b.rows[0])
}

// Height bestimmt die Höhe des Spielfelds, d.h. die Anzahl der Zeilen.
func (b *Board) Height() int {
	return len(b.rows)
}

// Row liefert die Zeile an der angegebenen Position zurück.
func (b *Board) Row(index int) []string {
	return b.rows[index]
}

// Col liefert die Spalte an der angegebenen Position zurück.
func (b *Board) Col(index int) []string {
	col := make([]string, len(b.rows))
	for i, row := range b.rows {
		col[i] = row[index]
	}
	return col
}

// DiagDownRight liefert eine Diagonale von oben links nach unten rechts zurück.
func (b *Board) DiagDownRight(startCol int) []string {
	diag := []string{}

	height := len(b.rows)
	width := len(b.rows[0])

	var row, col int

	if startCol >= 0 {
		if startCol >= width {
			return diag
		}
		row = 0
		col = startCol
	} else {
		row = -startCol
		col = 0
		if row >= height {
			return diag
		}
	}

	for row < height && col < width {
		diag = append(diag, b.rows[row][col])
		row++
		col++
	}
	return diag
}

// DiagDownLeft liefert eine Diagonale von oben rechts nach unten links zurück.
func (b *Board) DiagDownLeft(startCol int) []string {
	diag := []string{}

	height := len(b.rows)
	width := len(b.rows[0])

	var row, col int

	if startCol >= 0 && startCol < width {
		row = 0
		col = startCol
	} else if startCol >= width {
		row = startCol - (width - 1)
		col = width - 1
		if row >= height {
			return diag
		}
	} else { // startCol < 0
		return diag
	}

	for row < height && col >= 0 {
		diag = append(diag, b.rows[row][col])
		row++
		col--
	}

	return diag
}
