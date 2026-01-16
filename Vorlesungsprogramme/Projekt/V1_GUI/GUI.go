package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func a() {

}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Hello GUI in Go")

	label := widget.NewLabel("Hallo Welt!")
	button := widget.NewButton("Klick mich", func() {
		label.SetText("Du hast geklickt!")
	})

	myWindow.SetContent(container.NewVBox(label, button))
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
