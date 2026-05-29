package index

import (
	"inf1-material/Vorlesungsprogramme/26.05.29/core/user"
	"inf1-material/Vorlesungsprogramme/26.05.29/view/element"
)

// Index repräsentiert einen binären Suchbaum, der Benutzer enthält.
// Der Baum dient als Suchindex, um Benutzer schnell anhand eines Schlüssels zu finden.
type Index struct {
	root *element.Element
	key  func(e *user.User) string
}
