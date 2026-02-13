package board

import (
	"fmt"
	"strings"
)

// String liefert eine menschenlesbare Darstellung des Spielfelds.
func (b *Board) String() string {
	var sb strings.Builder

	height := b.Height()
	width := b.Width()
	cellWidth := b.MaxCellWidth()

	// Mindestbreite 1, damit leere Boards schön aussehen
	if cellWidth < 1 {
		cellWidth = 1
	}

	// Funktion zum Zeichnen der horizontalen Linie
	drawLine := func() {
		sb.WriteString("+")
		for i := 0; i < width; i++ {
			sb.WriteString(strings.Repeat("-", cellWidth+2))
			sb.WriteString("+")
		}
		sb.WriteString("\n")
	}

	for r := 0; r < height; r++ {
		sb.WriteString("|")
		for c := 0; c < width; c++ {
			cell := b.rows[r][c]
			sb.WriteString(" ")
			sb.WriteString(fmt.Sprintf("%-*s", cellWidth, cell))
			sb.WriteString(" |")
		}
		sb.WriteString("\n")
		drawLine()
	}
	return sb.String()
}
