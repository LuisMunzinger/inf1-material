package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GUI V1")

	content := container.NewGridWrap(
		fyne.NewSize(150, 50),
		widget.NewButton("Ja", func() { println("Ja") }),
		widget.NewButton("Nein", func() { println("Nein") }),
		widget.NewButton("Abbruch", func() { w.Close() }),
	)
	w.SetContent(content)
	w.ShowAndRun()
}
