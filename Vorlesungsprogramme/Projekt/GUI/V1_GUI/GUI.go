package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	w.SetContent(widget.NewLabel("Luis"))

	btn := widget.NewButton("Klick mich", func() { println("Button wurde geklickt") })

	w.SetContent(btn)
	w.ShowAndRun()
}
