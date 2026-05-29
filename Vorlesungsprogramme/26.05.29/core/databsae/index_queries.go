package database

import (
	"inf1-material/Vorlesungsprogramme/26.05.29/view/index"
)

// UsersById Gibt einen Index zurück, der Benutzer anhand ihrer IDs sortiert.
func (db *Database) UsersById() *index.Index {
	// Hinweis:
	// - Erstellen Sie einen neuen Index mit `index.ById()`.
	// - Durchlaufen Sie dann die Liste aller Benutzer in `db.Users`
	//   und fügen Sie sie mit der `Insert`-Methode in den Index ein.

	idx := index.ById()

	for _, u := range db.Users {
		idx.Insert(u)
	}

	return idx
}

// UsersByNickname Gibt einen Index zurück, der Benutzer anhand ihrer Nicknames sortiert.
func (db *Database) UsersByNickname() *index.Index {
	// Hinweis: Gehen Sie analog zu `UsersById` vor.

	idx := index.ByNickname()

	for _, u := range db.Users {
		idx.Insert(u)
	}

	return idx
}

// UsersByName Gibt einen Index zurück, der Benutzer anhand ihrer Namen sortiert.
func (db *Database) UsersByName() *index.Index {
	// TODO
	return nil
}

// UsersBySurname Gibt einen Index zurück, der Benutzer anhand ihrer Nachnamen sortiert.
func (db *Database) UsersBySurname() *index.Index {
	// TODO
	return nil
}
