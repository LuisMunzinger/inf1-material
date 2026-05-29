package index

import (
	"inf1-material/Vorlesungsprogramme/26.05.29/core/user"
	"inf1-material/Vorlesungsprogramme/26.05.29/view/element"
)

// New erstellt einen neuen leeren Index mit der gegebenen Schlüsselfunktion.
func New(key func(e *user.User) string) *Index {
	return &Index{
		root: element.Empty(),
		key:  key,
	}
}

// ById erstellt einen Index, der Benutzer anhand ihrer ID sortiert.
func ById() *Index {
	return New(func(u *user.User) string {
		return u.Id
	})
}

// ByNickname erstellt einen Index, der Benutzer anhand ihres Nickname sortiert.
func ByNickname() *Index {
	return New(func(u *user.User) string {
		return u.Nickname
	})
}

// ByName erstellt einen Index, der Benutzer anhand ihres Namens sortiert.
func ByName() *Index {
	return New(func(u *user.User) string {
		return u.Name
	})
}

// BySurname erstellt einen Index, der Benutzer anhand ihres Nachnamens sortiert.
func BySurname() *Index {
	return New(func(u *user.User) string {
		return u.Surname
	})
}
