package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Mini-Spiel mit Charakter")

	// Charakter (Rechteck)
	playerSize := float32(40)
	player := canvas.NewRectangle(color.RGBA{0, 200, 100, 255})
	player.Resize(fyne.NewSize(playerSize, playerSize))

	// Startposition
	x, y := float32(100), float32(100)
	player.Move(fyne.NewPos(x, y))

	game := container.NewWithoutLayout(player)
	w.SetContent(game)

	speed := float32(10)

	// Tastatursteuerung
	w.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {
		case fyne.KeyUp:
			y -= speed
		case fyne.KeyDown:
			y += speed
		case fyne.KeyLeft:
			x -= speed
		case fyne.KeyRight:
			x += speed
		}
		player.Move(fyne.NewPos(x, y))
		canvas.Refresh(player)
	})

	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}
