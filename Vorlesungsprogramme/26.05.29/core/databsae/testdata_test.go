package database

import (
	"inf1-material/Vorlesungsprogramme/26.05.29/core/user"
)

func testdata() *Database {
	db := New()

	u1 := user.New("1", "michkenntjeder", "Max", "Mustermann")
	u2 := user.New("2", "zombie43", "Arno", "Nym")
	u3 := user.New("3", "superstar123", "Eva", "Einbildung")

	db.AddUser(u1)
	db.AddUser(u2)
	db.AddUser(u3)

	u1.AddContact("3")
	u2.AddContact("1")
	u3.AddContact("1")
	u3.AddContact("2")

	return db
}
