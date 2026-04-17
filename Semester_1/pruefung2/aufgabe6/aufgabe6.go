package aufgabe6

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// DuplicateSinglets erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste, bei der alle Elemente,
// die in list nur einmal vorkommen, verdoppelt sind,
// also zwei Mal hintereinander stehen.
// Elemente, die schon in list mehrfach vorkommen, sollen wie sie sind
// ins Ergebnis übertragen werden.
func DuplicateSinglets(list []int) []int {
	result := []int{}

	for i := 0; i < len(list); i++ {
		count := 0
		// Häufigkeit von list[i] bestimmen
		for j := 0; j < len(list); j++ {
			if list[j] == list[i] {
				count++
			}
		}
		if count == 1 {
			// einmaliges Element → doppelt anhängen
			result = append(result, list[i], list[i])
		} else {
			// mehrfaches Element → normal übernehmen
			result = append(result, list[i])
		}
	}
	return result
}
