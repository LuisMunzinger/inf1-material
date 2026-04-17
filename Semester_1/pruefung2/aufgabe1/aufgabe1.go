package aufgabe1

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// PrefixBelow10 erwartet eine Liste "list" von Zahlen und liefert
// die längste Teil-Liste, mit der "list" beginnt und die nur Zahlen < 10 enthält.
func PrefixBelow10(list []int) []int {
	ergebnis := []int{}
	for i := 0; i < len(list) && list[i] < 10; i++ {
		ergebnis = append(ergebnis, list[i])
	}

	return ergebnis
}
