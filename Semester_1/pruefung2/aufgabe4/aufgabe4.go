package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// ElementSums erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Summe der beiden Elemente enthält.
//
// Annahmen für die Berechnung:
// Falls eine Liste kürzer ist als die andere, soll für die Berechnung der
// hinteren Werte ihr letztes Element verwendet werden.
// Für leere Listen soll für die Berechnung ggf. 0 verwendet werden.
func ElementSums(l1, l2 []int) []int {
	result := []int{}
	zwischen := 0
	letztel1 := 0
	letztel2 := 0
	for i := 0; i < len(l1) && i < len(l2); i++ {
		zwischen = l1[i] + l2[i]
		result = append(result, zwischen)
		letztel1++
		letztel2++
	}
	for letztel1 < len(l1) {
		zwischen = l1[letztel1] + l2[letztel2-1]
		result = append(result, zwischen)
		letztel1++
	}
	for letztel2 < len(l2) {
		zwischen = l1[letztel1-1] + l2[letztel2]
		result = append(result, zwischen)
		letztel2++
	}

	return result
}
