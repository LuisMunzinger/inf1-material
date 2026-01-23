package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GUI V2")
	fragen := []string{
		"Was ist die Hauptstadt von Deutschland?",
		"Wie viele Planeten hat unser Sonnensystem?",
		"Was ist 5 + 7?",
	}
	//---------------------------------------------------------------------------------------------------
	label1 := widget.NewLabel("" + fragen[0])
	label2 := widget.NewLabel("")

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Name eingeben...")

	button := widget.NewButton("Fertig", func() { label2.SetText("Hallo " + entry.Text) })

	head := container.NewVBox(
		label1, label2, entry, button)
	//---------------------------------------------------------------------------------------------------
	bottom := container.NewHBox(
		layout.NewSpacer(),
		widget.NewButton("Close", func() { w.Close() }),
	)
	//---------------------------------------------------------------------------------------------------
	w.SetContent(
		container.NewBorder(
			head,   // oben fix
			bottom, // unten fix
			nil,    // links fixiert
			nil,    // rechts fixiert
			nil,    // mitte flexibel
		),
	)

	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}
