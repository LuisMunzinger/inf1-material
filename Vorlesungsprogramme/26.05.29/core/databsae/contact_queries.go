package database

import (
	"inf1-material/Vorlesungsprogramme/26.05.29/core/user"
)

// GetDirectContacts liefert eine Liste mit den direkten Kontakten eines Benutzers zurück.
// Parameter:
//   - `id`: Die ID des Benutzers, dessen direkte Kontakte abgerufen werden sollen.
func (db *Database) GetDirectContacts(id string) []user.User {
	result := []user.User{}

	//TODO

	return result
}

// GetDirectContactIds liefert eine Liste mit den IDs aller direkten Kontakte eines Benutzers zurück.
// Parameter:
//   - `id`: Die ID des Benutzers, dessen direkte Kontakt-IDs abgerufen werden sollen.
func (db *Database) GetDirectContactIds(id string) []string {
	contact_ids := []string{}

	// TODO

	return contact_ids
}

// GetContacts liefert eine Liste mit allen direkten und indirekten Kontakten eines Benutzers zurück.
// Parameter:
//   - `id`: Die ID des Benutzers, dessen Kontakte abgerufen werden sollen.
//   - `depth`: Die Anzahl der Ebenen von Kontakten, die berücksichtigt werden sollen.
//     Eine Tiefe von 1 bedeutet nur direkte Kontakte, 2 bedeutet direkte Kontakte und deren Kontakte usw.
func (db *Database) GetContacts(id string, depth int) []user.User {
	result := []user.User{}

	// TODO

	return result
}

// GetContactIds liefert eine Liste mit den IDs aller direkten und indirekten Kontakte eines Benutzers zurück.
// Parameter:
//   - `id`: Die ID des Benutzers, dessen Kontakt-IDs abgerufen werden sollen.
//   - `depth`: Die Anzahl der Ebenen von Kontakten, die berücksichtigt werden sollen.
//     Eine Tiefe von 1 bedeutet nur direkte Kontakte, 2 bedeutet direkte Kontakte und deren Kontakte usw.
func (db *Database) GetContactIds(id string, depth int) []string {
	contact_ids := []string{}

	// TODO

	return contact_ids
}
