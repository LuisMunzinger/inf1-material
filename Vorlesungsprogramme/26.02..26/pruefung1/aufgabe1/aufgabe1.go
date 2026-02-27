package aufgabe1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ShortestAbc.
MAX. PUNKTE: 10
*/

// ShortestAbc erwartet eine Liste von Strings und liefert
// das kürzeste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
//
// Hinweis: Die Funktion muss nur mit kurzen Strings der Länge < 100 funktionieren.
func ShortestAbc(list []string) string {
	ergebnis := ""
	for i := 0; i < len(list); i++ {
		if len(list[i]) > 2 && list[i][:3] == "abc" {
			if len(list[i]) < len(ergebnis) || len(ergebnis) == 0 {
				ergebnis = list[i]
			}
		}
	}
	return ergebnis
}
