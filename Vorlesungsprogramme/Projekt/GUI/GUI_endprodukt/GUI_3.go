package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// ---------------- Start-Fenster ----------------
func createStartWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Start")

	label := widget.NewLabel("Was wollen Sie rechnen?")

	button := widget.NewButton("Zum Systemumrechner", func() {
		sysWin := SystemRechnerWindow(a)
		sysWin.Show()
		w.Close()
	})

	button2 := widget.NewButton("Zum Taschenrechner", func() {
		taschenWin := TaschenrechnerWindow(a)
		taschenWin.Show()
		w.Close()
	})

	w.SetContent(container.NewVBox(
		layout.NewSpacer(),
		label,
		button,
		button2,
		layout.NewSpacer(),
	))

	w.Resize(fyne.NewSize(400, 200))
	return w
}

// ---------------- Systemrechner-Fenster ----------------
func SystemRechnerWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Systemumrechner")

	label1 := widget.NewLabel("Wählen Sie ein System aus dem Sie umwandeln wollen:")
	entry1 := widget.NewEntry()

	label2 := widget.NewLabel("Wählen Sie ein System in das umgerechnet werden soll:")
	entry2 := widget.NewEntry()

	label3 := widget.NewLabel("Zahl zum Umwandeln:")
	entry3 := widget.NewEntry()

	label4 := widget.NewLabel("Ergebnis:")
	label5 := widget.NewLabel("")

	button := widget.NewButton("Rechnen", func() {
		vonSystem, _ := strconv.Atoi(entry1.Text)
		zuSystem, _ := strconv.Atoi(entry2.Text)

		zahl, _ := strconv.ParseInt(entry3.Text, vonSystem, 64)
		ergebnis := strconv.FormatInt(zahl, zuSystem)

		label5.SetText(ergebnis)
	})

	content := container.NewVBox(
		label1, entry1,
		label2, entry2,
		label3, entry3,
		container.NewHBox(label4, label5),
		button,
	)

	bottom := container.NewHBox(
		widget.NewButton("Zurück", func() {
			start := createStartWindow(a)
			start.Show()
			w.Close()
		}),
		layout.NewSpacer(),
		widget.NewButton("Schließen", func() {
			w.Close()
		}),
	)

	w.SetContent(container.NewBorder(content, bottom, nil, nil))
	w.Resize(fyne.NewSize(500, 400))
	return w
}

// ---------------- Taschenrechner-Fenster ----------------
func TaschenrechnerWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Taschenrechner")

	label1 := widget.NewLabel("Was wollen Sie rechnen?")
	entry1 := widget.NewEntry()

	label2 := widget.NewLabel("Geben Sie die erste Zahl ein:")
	entry2 := widget.NewEntry()

	label3 := widget.NewLabel("Geben Sie die zweite Zahl ein:")
	entry3 := widget.NewEntry()

	label4 := widget.NewLabel("Ergebnis:")
	label5 := widget.NewLabel("")

	button := widget.NewButton("Rechnen", func() {
		zahl1, _ := strconv.Atoi(entry2.Text)
		zahl2, _ := strconv.Atoi(entry3.Text)

		ergebnis := zahl1 + zahl2
		label5.SetText(strconv.Itoa(ergebnis))
	})

	content := container.NewVBox(
		label1, entry1,
		label2, entry2,
		label3, entry3,
		container.NewHBox(label4, label5),
		button,
	)

	bottom := container.NewHBox(
		widget.NewButton("Zurück", func() {
			start := createStartWindow(a)
			start.Show()
			w.Close()
		}),
		layout.NewSpacer(),
		widget.NewButton("Schließen", func() {
			w.Close()
		}),
	)

	w.SetContent(container.NewBorder(content, bottom, nil, nil))
	w.Resize(fyne.NewSize(500, 400))
	return w
}

func main() {
	a := app.New()
	startWindow := createStartWindow(a)
	startWindow.ShowAndRun()
}
