package main

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

/* ===================== TYPEN ===================== */

type Mine struct {
	Name string
	Pos  fyne.Position
	Rate int
}

type Inventory struct {
	Stone, Iron, Gold, Platinum, Diamond int
	Pickaxe                              string
}

/* ===================== KONSTANTEN ===================== */

const (
	windowWidth     = 1200
	windowHeight    = 700
	playerW         = 20
	playerH         = 40
	speed           = 10
	interactionDist = 80
	iconSize        = 40
)

/* ===================== HILFSFUNKTIONEN ===================== */

func dist(a, b fyne.Position) float64 {
	return math.Hypot(float64(a.X-b.X), float64(a.Y-b.Y))
}

func clamp(v, min, max float32) float32 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

/* ===================== MAIN ===================== */

func main() {
	a := app.New()
	w := a.NewWindow("Mining Game")
	w.Resize(fyne.NewSize(windowWidth, windowHeight))

	/* ----- SPIELER ----- */
	player := canvas.NewRectangle(color.RGBA{0, 200, 100, 255})
	player.Resize(fyne.NewSize(playerW, playerH))
	x, y := float32(50), float32(50)
	player.Move(fyne.NewPos(x, y))

	/* ----- INVENTAR ----- */
	inv := &Inventory{Pickaxe: "Holzspitzhacke"}
	money := 0

	/* ----- MINEN ----- */
	mines := []Mine{
		{"Stein", fyne.NewPos(100, 200), 1},
		{"Eisen", fyne.NewPos(300, 200), 1},
		{"Gold", fyne.NewPos(500, 200), 1},
		{"Platin", fyne.NewPos(700, 200), 1},
		{"Diamant", fyne.NewPos(900, 200), 1},
	}

	objects := []fyne.CanvasObject{player}

	for _, m := range mines {
		icon := widget.NewButtonWithIcon("", theme.StorageIcon(), nil)
		icon.Resize(fyne.NewSize(iconSize, iconSize))
		icon.Move(m.Pos)

		label := canvas.NewText(m.Name, color.Black)
		label.Move(fyne.NewPos(m.Pos.X, m.Pos.Y+45))

		objects = append(objects, icon, label)
	}

	/* ----- SHOP & SCHMIED ----- */
	shopPos := fyne.NewPos(200, 20)
	shop := widget.NewButtonWithIcon("Shop", theme.InfoIcon(), nil)
	shop.Move(shopPos)
	shop.Resize(fyne.NewSize(80, 40))

	smithPos := fyne.NewPos(300, 20)
	smith := widget.NewButtonWithIcon("Schmied", theme.InfoIcon(), nil)
	smith.Move(smithPos)
	smith.Resize(fyne.NewSize(100, 40))

	objects = append(objects, shop, smith)

	w.SetContent(container.NewWithoutLayout(objects...))

	/* ----- RESSOURCEN TICK ----- */
	go func() {
		for range time.Tick(time.Second) {
			inv.Stone += mines[0].Rate
			inv.Iron += mines[1].Rate
			inv.Gold += mines[2].Rate
			inv.Platinum += mines[3].Rate
			inv.Diamond += mines[4].Rate
		}
	}()

	/* ----- STEUERUNG + BARRIERE ----- */
	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyUp:
			y -= speed
		case fyne.KeyDown:
			y += speed
		case fyne.KeyLeft:
			x -= speed
		case fyne.KeyRight:
			x += speed
		case fyne.KeyReturn:
			p := fyne.NewPos(x, y)
			if dist(p, shopPos) < interactionDist {
				openShop(a, inv, &money)
			}
			if dist(p, smithPos) < interactionDist {
				openSmith(a, inv)
			}
		case fyne.KeySpace:
			openInventory(a, inv, &money)
		}

		/* ----- FENSTER-BARRIERE ----- */
		x = clamp(x, 0, windowWidth-playerW)
		y = clamp(y, 0, windowHeight-playerH)

		player.Move(fyne.NewPos(x, y))
	})

	w.ShowAndRun()
}

/* ===================== SHOP ===================== */

func openShop(a fyne.App, inv *Inventory, money *int) {
	w := a.NewWindow("Shop – Verkaufen")

	prices := map[string]int{
		"Stein": 1, "Eisen": 5, "Gold": 10, "Platin": 20, "Diamant": 50,
	}

	box := container.NewVBox(widget.NewLabel("Ressourcen verkaufen"))

	add := func(name string, val *int) {
		btn := widget.NewButton("", func() {
			if *val > 0 {
				*money += *val * prices[name]
				*val = 0
				w.Close()
			}
		})
		btn.SetText(fmt.Sprintf("%s verkaufen (%d)", name, *val))
		box.Add(btn)
	}

	add("Stein", &inv.Stone)
	add("Eisen", &inv.Iron)
	add("Gold", &inv.Gold)
	add("Platin", &inv.Platinum)
	add("Diamant", &inv.Diamond)

	w.SetContent(box)
	w.Resize(fyne.NewSize(300, 300))
	w.Show()
}

/* ===================== SCHMIED ===================== */

func openSmith(a fyne.App, inv *Inventory) {
	w := a.NewWindow("Schmied")

	make := func(name string, cost int, res *int) *widget.Button {
		return widget.NewButton(name, func() {
			if *res >= cost {
				*res -= cost
				inv.Pickaxe = name
				w.Close()
			}
		})
	}

	w.SetContent(container.NewVBox(
		make("Holzspitzhacke", 0, &inv.Stone),
		make("Eisenspitzhacke (10)", 10, &inv.Iron),
		make("Goldspitzhacke (10)", 10, &inv.Gold),
		make("Platinspitzhacke (10)", 10, &inv.Platinum),
		make("Diamantspitzhacke (5)", 5, &inv.Diamond),
	))

	w.Resize(fyne.NewSize(300, 300))
	w.Show()
}

/* ===================== INVENTAR ===================== */

func openInventory(a fyne.App, inv *Inventory, money *int) {
	w := a.NewWindow("Inventar")

	label := widget.NewLabel("")
	update := func() {
		label.SetText(fmt.Sprintf(
			"Geld: %d\nSpitzhacke: %s\n\nStein: %d\nEisen: %d\nGold: %d\nPlatin: %d\nDiamant: %d",
			*money, inv.Pickaxe, inv.Stone, inv.Iron, inv.Gold, inv.Platinum, inv.Diamond,
		))
	}

	update()
	go func() {
		for range time.Tick(time.Second) {
			update()
		}
	}()

	w.SetContent(label)
	w.Resize(fyne.NewSize(250, 300))
	w.Show()
}
