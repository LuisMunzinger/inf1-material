package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	ergebnis := []string{}
	vorhanden := false
	if len(list) == 0 {
		return ergebnis
	}
	for a := 0; a < len(list); a++ {
		if list[a] == first {
			vorhanden = true
		}
	}

	if vorhanden == true {
		for i := 0; i < len(list); i++ {
			for i < len(list) && list[i] != first && list[i] != last {
				ergebnis = append(ergebnis, list[i])
				i++
			}
			i++
			for i < len(list) && list[i] != last {
				i++
			}
		}
	}
	return ergebnis
}
