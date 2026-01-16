package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Button Beispiel")

	content := container.NewGridWrap(
		fyne.NewSize(150, 50),
		widget.NewButton("Ja", func() { println("Ja") }),
		widget.NewButton("Nein", func() { println("Nein") }),
		widget.NewButton("Abbruch", func() { println("Abbruch") }),
	)
	w.SetContent(content)
	//w.SetContent(container.NewVBox(widget.NewLabel("Hello Fyne!")))
	w.ShowAndRun()
}
