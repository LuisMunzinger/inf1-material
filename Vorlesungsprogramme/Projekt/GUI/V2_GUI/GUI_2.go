package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GUI V2")

	label := widget.NewLabel("Wie heist du?")
	label2 := widget.NewLabel("")

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Name eingeben...")

	content := container.NewGridWrap(
		fyne.NewSize(150, 50),
		widget.NewButton("Fertig", func() { label2.SetText("Hallo " + entry.Text) }),
		widget.NewButton("Close", func() { w.Close() }),
	)

	w.SetContent(container.NewVBox(label, label2, entry, content))
	w.SetContent(
		(container.NewWithoutLayout()),
	)
	w.ShowAndRun()
}
