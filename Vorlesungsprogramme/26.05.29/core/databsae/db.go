package database

import (
	"inf1-material/Vorlesungsprogramme/26.05.29/core/user"
)

// Database repräsentiert unsere Datenbank, die alle Benutzer enthält.
type Database struct {
	Users []*user.User
}

// New erstellt eine neue leere Datenbank.
func New() *Database {
	return &Database{
		Users: []*user.User{},
	}
}

// AddUser fügt einen neuen Benutzer zur Datenbank hinzu.
func (db *Database) AddUser(u *user.User) {
	db.Users = append(db.Users, u)
}
