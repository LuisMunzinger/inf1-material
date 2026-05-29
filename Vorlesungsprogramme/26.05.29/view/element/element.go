package element

import "inf1-material/Vorlesungsprogramme/26.05.29/core/user"

// Element ist ein Element in einem binären Suchbaum, der Benutzer enthält.
type Element struct {
	User  *user.User
	Left  *Element
	Right *Element
}

// Empty erzeugt ein neues leeres Element.
func Empty() *Element {
	return &Element{
		User:  nil,
		Left:  nil,
		Right: nil,
	}
}

// IsEmpty prüft, ob dieses Element leer ist.
// Ein Element ist leer, wenn irgendeines der Felder nil ist.
func (e *Element) IsEmpty() bool {
	// TODO
	return false
}

// SetUser setzt den Benutzer in diesem Element.
func (e *Element) SetUser(u *user.User) {
	// TODO
}

// List gibt eine Liste aller Benutzer in diesem Element und seinen Kindern
// in sortierter Reihenfolge zurück.
func (e *Element) List() []*user.User {
	users := []*user.User{}

	// TODO

	return users
}
