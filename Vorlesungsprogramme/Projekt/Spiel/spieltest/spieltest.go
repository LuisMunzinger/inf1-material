package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

/* ===================== TYPEN & KONSTANTEN ===================== */
type Resource string
type Pickaxe string

const (
	Stone    Resource = "Stein"
	Iron     Resource = "Eisen"
	Gold     Resource = "Gold"
	Platinum Resource = "Platin"
	Diamond  Resource = "Diamant"

	WoodPickaxe     Pickaxe = "Holzspitzhacke"
	StonePickaxe    Pickaxe = "Steinspitzhacke"
	IronPickaxe     Pickaxe = "Eisenspitzhacke"
	GoldPickaxe     Pickaxe = "Goldspitzhacke"
	PlatinumPickaxe Pickaxe = "Platinspitzhacke"
	DiamondPickaxe  Pickaxe = "Diamantspitzhacke"
)

const (
	windowWidth     = 1200
	windowHeight    = 700
	playerW         = 20
	playerH         = 40
	speed           = 10
	iconSize        = 40
	interactionDist = 80
	borderWidth     = 4
	MaxMineUpgrade  = 9
)

/* ===================== UPGRADE-DATEN ===================== */
var pickaxeUpgrade = map[Pickaxe]Pickaxe{
	WoodPickaxe:     StonePickaxe,
	StonePickaxe:    IronPickaxe,
	IronPickaxe:     GoldPickaxe,
	GoldPickaxe:     PlatinumPickaxe,
	PlatinumPickaxe: DiamondPickaxe,
}

var pickaxeResourceCost = map[Pickaxe]map[Resource]int{
	WoodPickaxe:     {Stone: 50},
	IronPickaxe:     {Iron: 50},
	GoldPickaxe:     {Gold: 50},
	PlatinumPickaxe: {Platinum: 50},
}

var mineUpgradeResources = []Resource{Stone, Iron, Gold, Platinum, Diamond}
var mineUpgradeRateIncrease = 1

/* ===================== STRUKTUREN ===================== */
type Mine struct {
	Name         string
	Pos          fyne.Position
	Resource     Resource
	Cost         int
	Rate         int
	Required     Pickaxe
	Owned        bool
	UpgradeLevel int
	Icon         *canvas.Rectangle
	Label        *canvas.Text
}

type Inventory struct {
	sync.Mutex
	Resources map[Resource]int
	Pickaxe   Pickaxe
	Money     int
}

/* ===================== HELPER ===================== */
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

func getMineUpgradeCost(m *Mine) (money int, resources map[Resource]int, ok bool) {
	if m.UpgradeLevel >= MaxMineUpgrade {
		return 0, nil, false
	}
	resources = make(map[Resource]int)
	money = m.Cost * (m.UpgradeLevel + 1)
	for i := 0; i <= m.UpgradeLevel && i < len(mineUpgradeResources); i++ {
		resources[mineUpgradeResources[i]] = (m.UpgradeLevel + 1) * 10
	}
	return money, resources, true
}

/* ===================== MAIN ===================== */
func main() {
	a := app.New()
	w := a.NewWindow("Mining Game")
	w.Resize(fyne.NewSize(windowWidth, windowHeight))

	inv := &Inventory{
		Resources: make(map[Resource]int),
		Pickaxe:   WoodPickaxe,
		Money:     200,
	}

	player := canvas.NewRectangle(color.RGBA{0, 200, 100, 255})
	player.Resize(fyne.NewSize(playerW, playerH))
	x, y := float32(50), float32(50)
	player.Move(fyne.NewPos(x, y))

	mines := []*Mine{
		{"Steinmine", fyne.NewPos(100, 200), Stone, 100, 1, WoodPickaxe, false, 0, nil, nil},
		{"Eisenmine", fyne.NewPos(300, 200), Iron, 250, 1, StonePickaxe, false, 0, nil, nil},
		{"Goldmine", fyne.NewPos(500, 200), Gold, 500, 1, IronPickaxe, false, 0, nil, nil},
		{"Platinmine", fyne.NewPos(700, 200), Platinum, 1000, 1, GoldPickaxe, false, 0, nil, nil},
		{"Diamantmine", fyne.NewPos(900, 200), Diamond, 2000, 1, PlatinumPickaxe, false, 0, nil, nil},
	}

	objects := []fyne.CanvasObject{player}

	for _, m := range mines {
		icon := canvas.NewRectangle(color.Gray{Y: 180})
		icon.Resize(fyne.NewSize(iconSize, iconSize))
		icon.Move(m.Pos)
		icon.StrokeColor = color.Black
		icon.StrokeWidth = borderWidth
		m.Icon = icon

		label := canvas.NewText(m.Name, color.Black)
		label.TextSize = 12
		label.Move(fyne.NewPos(m.Pos.X, m.Pos.Y+45))
		m.Label = label

		objects = append(objects, icon, label)
	}

	shopPos := fyne.NewPos(200, 20)
	shop := widget.NewButtonWithIcon("Shop", theme.InfoIcon(), func() { openShop(a, inv) })
	shop.Move(shopPos)
	shop.Resize(fyne.NewSize(80, 40))

	smithPos := fyne.NewPos(400, 20)
	smith := widget.NewButtonWithIcon("Schmied", theme.InfoIcon(), func() { openSmith(a, inv) })
	smith.Move(smithPos)
	smith.Resize(fyne.NewSize(110, 40))

	objects = append(objects, shop, smith)
	w.SetContent(container.NewWithoutLayout(objects...))

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			inv.Lock()
			for _, m := range mines {
				if m.Owned {
					inv.Resources[m.Resource] += m.Rate
				}
			}
			inv.Unlock()
		}
	}()

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
			for _, m := range mines {
				if dist(p, m.Pos) < interactionDist {
					openMineShop(a, m, inv)
					return
				}
			}
			if dist(p, shopPos) < interactionDist {
				openShop(a, inv)
			} else if dist(p, smithPos) < interactionDist {
				openSmith(a, inv)
			}
		case fyne.KeySpace:
			openInventory(a, inv)
		}

		x = clamp(x, 0, windowWidth-playerW)
		y = clamp(y, 0, windowHeight-playerH)
		player.Move(fyne.NewPos(x, y))

		for _, m := range mines {
			d := dist(fyne.NewPos(x, y), m.Pos)
			if m.Owned {
				m.Icon.FillColor = color.RGBA{0, 255, 0, 255}
			} else if d < interactionDist {
				m.Icon.FillColor = color.RGBA{255, 255, 0, 255}
			} else {
				m.Icon.FillColor = color.Gray{Y: 180}
			}
			m.Icon.Refresh()
		}
	})

	w.ShowAndRun()
}

/* ===================== MINE SHOP ===================== */
func openMineShop(a fyne.App, m *Mine, inv *Inventory) {
	w := a.NewWindow(m.Name)
	w.Resize(fyne.NewSize(400, 350))

	label := widget.NewLabel("")
	var actionBtn *widget.Button

	updateLabel := func() {
		text := fmt.Sprintf("%s\nRate: %d/sec\nUpgrade-Level: %d/%d", m.Name, m.Rate, m.UpgradeLevel, MaxMineUpgrade)
		text += fmt.Sprintf("\nBenötigte Spitzhacke: %s", m.Required)
		if m.Owned {
			money, res, ok := getMineUpgradeCost(m)
			if ok {
				text += fmt.Sprintf("\nNächste Upgrade-Kosten:\nGeld: %d", money)
				for r, amt := range res {
					text += fmt.Sprintf("\n%s: %d", r, amt)
				}
			} else {
				text += "\n(Max Level erreicht)"
			}
		} else {
			text += fmt.Sprintf("\nKaufpreis: %d Geld", m.Cost)
		}
		label.SetText(text)
	}

	actionBtn = widget.NewButton("", func() {
		inv.Lock()
		defer inv.Unlock()

		pickaxeLevels := map[Pickaxe]int{
			WoodPickaxe:     1,
			StonePickaxe:    2,
			IronPickaxe:     3,
			GoldPickaxe:     4,
			PlatinumPickaxe: 5,
			DiamondPickaxe:  6,
		}
		if pickaxeLevels[inv.Pickaxe] < pickaxeLevels[m.Required] {
			label.SetText(label.Text + "\nDu benötigst eine bessere Spitzhacke!")
			return
		}

		if !m.Owned {
			if inv.Money >= m.Cost {
				inv.Money -= m.Cost
				m.Owned = true
				updateLabel()
				actionBtn.SetText("Upgrade Mine")
			}
			return
		}

		moneyCost, resCosts, ok := getMineUpgradeCost(m)
		if !ok {
			return
		}
		for r, amt := range resCosts {
			if inv.Resources[r] < amt {
				return
			}
		}
		if inv.Money < moneyCost {
			return
		}

		inv.Money -= moneyCost
		for r, amt := range resCosts {
			inv.Resources[r] -= amt
		}
		m.Rate += mineUpgradeRateIncrease
		m.UpgradeLevel++
		updateLabel()
		if m.UpgradeLevel >= MaxMineUpgrade {
			actionBtn.Disable()
		}
	})

	if m.Owned {
		actionBtn.SetText("Upgrade Mine")
	} else {
		actionBtn.SetText(fmt.Sprintf("Kaufen (%d Geld)", m.Cost))
	}

	w.SetContent(container.NewVBox(label, actionBtn))
	updateLabel()
	w.Show()
}

/* ===================== SHOP ===================== */
func openShop(a fyne.App, inv *Inventory) {
	prices := map[Resource]int{
		Stone:    1,
		Iron:     5,
		Gold:     10,
		Platinum: 20,
		Diamond:  50,
	}
	resourceOrder := []Resource{Stone, Iron, Gold, Platinum, Diamond}

	w := a.NewWindow("Shop")
	box := container.NewVBox(widget.NewLabel("Ressourcen verkaufen"))

	labels := make(map[Resource]*widget.Label)
	sellAmounts := []int{1, 10, 100}

	for _, r := range resourceOrder {
		p := prices[r]
		resBox := container.NewVBox()
		label := widget.NewLabel(fmt.Sprintf("%s: %d", r, inv.Resources[r]))
		labels[r] = label
		resBox.Add(label)
		for _, amt := range sellAmounts {
			amount := amt
			btn := widget.NewButton(fmt.Sprintf("Verkaufen %d", amount), func() {
				inv.Lock()
				defer inv.Unlock()
				sell := amount
				if inv.Resources[r] < sell {
					sell = inv.Resources[r]
				}
				if sell > 0 {
					inv.Resources[r] -= sell
					inv.Money += sell * p
				}
				label.SetText(fmt.Sprintf("%s: %d", r, inv.Resources[r]))
			})
			resBox.Add(btn)
		}
		box.Add(resBox)
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			inv.Lock()
			for r, label := range labels {
				label.SetText(fmt.Sprintf("%s: %d", r, inv.Resources[r]))
			}
			inv.Unlock()
		}
	}()

	w.SetContent(container.NewVScroll(box))
	w.Resize(fyne.NewSize(300, 400))
	w.Show()
}

/* ===================== SCHMIED ===================== */
func openSmith(a fyne.App, inv *Inventory) {
	w := a.NewWindow("Schmied")
	w.Resize(fyne.NewSize(320, 220))
	info := widget.NewLabel("")
	btn := widget.NewButton("Upgrade kaufen", func() {
		inv.Lock()
		defer inv.Unlock()
		next, ok := pickaxeUpgrade[inv.Pickaxe]
		if !ok {
			return
		}
		costs := pickaxeResourceCost[inv.Pickaxe]
		for res, amount := range costs {
			if inv.Resources[res] < amount {
				return
			}
		}
		for res, amount := range costs {
			inv.Resources[res] -= amount
		}
		inv.Pickaxe = next
		w.Close()
	})

	updateUI := func() {
		inv.Lock()
		defer inv.Unlock()
		next, ok := pickaxeUpgrade[inv.Pickaxe]
		if !ok {
			info.SetText("Spitzhacke: " + string(inv.Pickaxe) + "\n(Maximales Level)")
			btn.Disable()
			return
		}
		costText := ""
		for res, amount := range pickaxeResourceCost[inv.Pickaxe] {
			costText += fmt.Sprintf("%s: %d ", res, amount)
		}
		info.SetText(fmt.Sprintf(
			"Aktuell: %s\nNächste: %s\nBenötigte Ressourcen: %s",
			inv.Pickaxe, next, costText,
		))
		btn.Enable()
	}

	updateUI()
	w.SetContent(container.NewVBox(info, btn))
	w.Show()
}

/* ===================== INVENTAR ===================== */
func openInventory(a fyne.App, inv *Inventory) {
	w := a.NewWindow("Inventar")
	str := binding.NewString()
	label := widget.NewLabelWithData(str)

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			inv.Lock()
			value := fmt.Sprintf(
				"Geld: %d\nSpitzhacke: %s\n\nStein: %d\nEisen: %d\nGold: %d\nPlatin: %d\nDiamant: %d",
				inv.Money, inv.Pickaxe,
				inv.Resources[Stone],
				inv.Resources[Iron],
				inv.Resources[Gold],
				inv.Resources[Platinum],
				inv.Resources[Diamond],
			)
			inv.Unlock()
			_ = str.Set(value)
		}
	}()

	w.SetContent(label)
	w.Resize(fyne.NewSize(260, 300))
	w.Show()
}
