package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Custom Icon Button")

	// Bild als Icon-Ressource laden
	iconResource, err := fyne.LoadResourceFromPath("C:/Users/silke/Desktop/inf1-material/Vorlesungsprogramme/Projekt/Spiel/bild/Screenshot 2025-02-05 153606.png")
	if err != nil {
		log.Fatal(err)
	}

	// Button mit eigenem Icon
	button := widget.NewButtonWithIcon("Klick mich", iconResource, func() {
		fmt.Println("Button mit Custom Icon geklickt!")
	})

	myWindow.SetContent(button)
	myWindow.Resize(fyne.NewSize(500, 300))
	myWindow.ShowAndRun()
}
