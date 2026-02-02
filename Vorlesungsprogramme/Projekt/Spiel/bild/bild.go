package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne Hintergrundbild")

	// Bild aus dem gleichen Ordner laden
	background := canvas.NewImageFromFile("GoldMine.png")
	background.FillMode = canvas.ImageFillStretch // oder ImageFillContain

	// Container, damit das Bild den ganzen Hintergrund füllt
	content := container.NewMax(background)

	w.SetContent(content)
	w.Resize(fyne.NewSize(900, 600))
	w.ShowAndRun()
}
