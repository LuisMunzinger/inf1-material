package aufgabe4

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion MaxElements.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein!
*/

// MaxElements erwartet zwei int-Listen und liefert eine Liste, die an jeder Position
// jeweils das größere der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt, gilt dieses Element als das größere.
func MaxElements(l1, l2 []int) []int {
	ergebnis := []int{}
	if len(l1) == 0 && len(l2) == 0 {
		return ergebnis
	}
	if len(l1) == 0 {
		ergebnis = append(ergebnis, l2[0])
		MaxElements(l1, l2[:1])
	}
	if len(l2) == 0 {
		ergebnis = append(ergebnis, l1[0])
		MaxElements(l1[:1], l2)
	}
	if l1[0] > l2[0] {
		ergebnis = append(ergebnis, l1[0])
		MaxElements(l1[1:], l2[1:])
	}
	if l1[0] < l2[0] {
		ergebnis = append(ergebnis, l2[0])
		MaxElements(l1[1:], l2[1:])
	} else {
		ergebnis = append(ergebnis, l1[0])
		MaxElements(l1[1:], l2[1:])
	}
	return append(ergebnis, MaxElements(l1[1:], l2[1:])...)
}
