package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GUI V2")
	//---------------------------------------------------------------------------------------------------
	label1 := widget.NewLabel("Wählen sie eine Systeme aus dem sie umwandeln wollen:  ")
	entry1 := widget.NewEntry()
	entry1.SetPlaceHolder("Systeme eingeben...")

	label2 := widget.NewLabel("Wählen sie ein System in das sie umrechnen wollen:  ")
	entry2 := widget.NewEntry()
	entry2.SetPlaceHolder("Zahl eingeben...")

	label3 := widget.NewLabel("Wählen sie eine Zahl zum umwandeln:  ")
	entry3 := widget.NewEntry()
	entry3.SetPlaceHolder("Zahl eingeben...")

	label4 := widget.NewLabel("Ergebnis:")
	label5 := widget.NewLabel("")

	button1 := widget.NewButton("Fertig",
		func() {
			Zwischen1 := entry1.Text
			Zwischen2 := entry2.Text
			zZahl := entry3.Text

			vonSysteme, _ := strconv.Atoi(Zwischen1)
			zuSysteme, _ := strconv.Atoi(Zwischen2)

			Zahl, _ := strconv.ParseInt(zZahl, int(vonSysteme), 64)
			a := (strconv.FormatInt(Zahl, int(zuSysteme)))
			label5.SetText(a)
		},
	)
	zwischen := container.NewHBox(label4, label5)
	head := container.NewVBox(label1, entry1, label2, entry2, label3, entry3, zwischen, button1)
	//---------------------------------------------------------------------------------------------------
	bottom := container.NewHBox(
		layout.NewSpacer(), // Nach rechts verschoben
		widget.NewButton("Schließen", func() { w.Close() }),
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

	w.Resize(fyne.NewSize(500, 400)) // Fenster Größe X Y
	w.ShowAndRun()
}
