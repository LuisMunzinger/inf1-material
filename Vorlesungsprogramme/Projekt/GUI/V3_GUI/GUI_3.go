package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func createWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Rechner")

	// ------------------ Eingaben ------------------
	label1 := widget.NewLabel("Wählen sie ein System aus dem sie umwandeln wollen:")
	entry1 := widget.NewEntry()
	entry1.SetPlaceHolder("System eingeben...")

	label2 := widget.NewLabel("Wählen sie ein System in das sie umrechnen wollen:")
	entry2 := widget.NewEntry()
	entry2.SetPlaceHolder("System eingeben...")

	label3 := widget.NewLabel("Wählen sie eine Zahl zum umwandeln:")
	entry3 := widget.NewEntry()
	entry3.SetPlaceHolder("Zahl eingeben...")

	label4 := widget.NewLabel("Ergebnis:")
	label5 := widget.NewLabel("")

	// ------------------ Button ------------------
	button1 := widget.NewButton("Fertig", func() {
		vonSysteme, _ := strconv.Atoi(entry1.Text)
		zuSysteme, _ := strconv.Atoi(entry2.Text)

		zZahl := entry3.Text
		zahl, _ := strconv.ParseInt(zZahl, vonSysteme, 64)
		ergebnis := strconv.FormatInt(zahl, zuSysteme)

		label5.SetText(ergebnis)
	})

	zwischen := container.NewHBox(label4, label5)
	head := container.NewVBox(
		label1, entry1,
		label2, entry2,
		label3, entry3,
		zwischen,
		button1,
	)

	// ------------------ Bottom ------------------
	bottom := container.NewHBox(
		layout.NewSpacer(),
		widget.NewButton("Schließen", func() {
			w.Close()
		}),
	)

	// ------------------ Layout ------------------
	w.SetContent(
		container.NewBorder(
			head,
			bottom,
			nil,
			nil,
			nil,
		),
	)

	w.Resize(fyne.NewSize(500, 400))
	return w
}

func main() {
	a := app.New()

	w := createWindow(a)
	w.ShowAndRun()
}
