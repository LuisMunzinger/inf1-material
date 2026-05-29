package user

// User repräsentiert einen Benutzer in unserem sozialen Netzwerk.
type User struct {
	Id       string
	Nickname string
	Name     string
	Surname  string
	Contacts []string
}

// New erstellt einen neuen Benutzer mit dem angegebenen Nickname
// sowie Vor- und Nachname. Die Kontaktliste ist zunächst leer.
func New(id string, nickname string, name string, surname string) *User {
	return &User{
		Id:       id,
		Nickname: nickname,
		Name:     name,
		Surname:  surname,
		Contacts: []string{},
	}
}

// AddContact fügt einen Kontakt zur Kontaktliste des Benutzers hinzu.
func (u *User) AddContact(contactID string) {
	// TODO
}
