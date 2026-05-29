package index

import (
	"inf1-material/Vorlesungsprogramme/26.05.29/core/user"
	"inf1-material/Vorlesungsprogramme/26.05.29/view/element"
)

// Insert fügt einen neuen Benutzer in den Index ein.
func (idx *Index) Insert(u *user.User) {
	// TODO
}

// Find sucht einen Benutzer im Index anhand eines Schlüssels.
func (idx *Index) Find(key string) *element.Element {
	// TODO

	return nil
}

// List liefert alle Benutzer im Index in sortierter Reihenfolge.
func (idx *Index) List() []*user.User {
	// TODO
	return nil
}

// Keys liefert eine Liste aller Schlüssel im Index in sortierter Reihenfolge.
func (idx *Index) Keys() []string {
	keys := []string{}

	// TODO

	return keys
}

// String gibt eine menschenlesbare Listen-Darstellung aller Schlüssel im Index zurück.
func (idx *Index) String() string {
	// TODO
	return ""
}
