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
	if e.User == nil || e.Left == nil || e.Right == nil {
		return true
	}
	return false
}

// SetUser setzt den Benutzer in diesem Element.
func (e *Element) SetUser(u *user.User) {
	e.User = u
	if e.Left == nil {
		e.Left = Empty()
	}
	if e.Right == nil {
		e.Right = Empty()
	}
}

// List gibt eine Liste aller Benutzer in diesem Element und seinen Kindern
// in sortierter Reihenfolge zurück.
func (e *Element) List() []*user.User {
	users := []*user.User{}

	if /* e.Left == nil || e.Right == nil */ e.IsEmpty() {
		return users
	}

	users = append(users, e.Left.List()...)
	users = append(users, e.User)
	users = append(users, e.Right.List()...)

	return users
}
