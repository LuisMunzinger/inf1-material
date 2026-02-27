package aufgabe5

/*
/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
die auf dem darunter definierten Datentyp arbeitet.
MAX. PUNKTE: 10
*/

// ReplaceEn sucht in der Liste nach dem Eintrag für
// de und ersetzt dessen En-Wert mit dem gegebenen en.
// Gibt es mehrere solche Einträge, soll nur der erste ersetzt werden.
func ReplaceEn(dict []Entry, de, en string) {
	// TODO
}

// Entry repräsentiert einen Eintrag in einem Wörterbuch
type Entry struct {
	De string
	En string
}
