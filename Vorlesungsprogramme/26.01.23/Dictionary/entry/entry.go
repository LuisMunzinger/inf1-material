package entry

type Entry struct {
	de string
	en string
}

// New erzeugt ein neues Entry-Objekt.
func New(de, en string) Entry {
	return Entry{de: de, en: en}
}

// Empty erzeugt ein leeres Entry-Objekt.
func Empty() Entry {
	return Entry{}
}

// De gibt den deutschen Eintrag zurück.
func (e Entry) De() string {
	// TODO
	return ""
}

// En gibt den englischen Eintrag zurück.
func (e Entry) En() string {
	// TODO
	return ""
}

// IsValid gibt true zurück, wenn der Eintrag gültig ist.
// Ein Eintrag ist gültig, wenn sowohl das deutsche als auch das englische Wort nicht leer sind.
func (e Entry) IsValid() bool {
	// TODO
	return false
}
