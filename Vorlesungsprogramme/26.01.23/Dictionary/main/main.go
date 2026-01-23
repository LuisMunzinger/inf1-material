package main

import (
	"fmt"
	"inf1-material/Vorlesungsprogramme/26.01.23/Dictionary/dict"
	"inf1-material/Vorlesungsprogramme/26.01.23/Dictionary/entry"
)

func main() {
	e1 := entry.New("Haus", "House")
	e2 := entry.New("Fahrrad", "Bicykel")

	d := dict.New()
	d.Add(e1)
	d.Add(e2)

	fmt.Print("Bitte ein deutsches wort eingeben: ")
	de := " "
	fmt.Scanln(&de)

	en := d.Lookup(de)
	if en != "" {
		fmt.Printf("Das Englische wort ist: %s\n", en)
	} else {
		fmt.Printf("Nicht Gefunden")
	}

}
