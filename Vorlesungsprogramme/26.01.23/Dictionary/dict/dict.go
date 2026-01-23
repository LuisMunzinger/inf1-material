package dict

import "Dictionary/entry/"

type Dict struct {
	entries []entry.Entry
}

// New erzeugt ein neues leeres Dict-Objekt.
func New() Dict {
	return Dict{}
}

// Add fügt einen Eintrag zum Wörterbuch hinzu.
func (d *Dict) Add(e entry.Entry) {
	d.entries = append(d.entries, e)
}

// Size gibt die Anzahl der Einträge im Wörterbuch zurück.
func (d Dict) Size() int {
	// TODO
	return 0
}

// GetDe gibt den Eintrag mit dem deutschen Wort de zurück.
// Wenn kein Eintrag gefunden wird, wird ein leerer Eintrag zurückgegeben.
func (d Dict) GetDe(de string) entry.Entry {
	// TODO
	return entry.Empty()
}

// Lookup sucht nach dem ersten Eintrag mit dem deutschen Wort de.
// Wenn ein Eintrag gefunden wird, wird der entsprechende englische string geliefert.
// Wenn kein Eintrag gefunden wird, wird ein leerer string zurückgegeben.
func (d Dict) Lookup(de string) string {
	// TODO
	return ""
}

// GetAllDe gibt alle Einträge zurück, die das deutsche Wort de enthalten.
func (d Dict) GetAllDe(de string) []entry.Entry {
	// TODO
	return []entry.Entry{}
}

// LookupAll sucht nach allen Einträgen mit dem deutschen Wort de.
// Gibt eine Liste mit den entsprechenden englischen Wörtern zurück.
func (d Dict) LookupAll(de string) []string {
	// TODO
	return []string{}
}
