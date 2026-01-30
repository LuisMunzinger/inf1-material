package main

import (
	"fmt"
	"image/color"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

/* ===================== TYPEN & KONSTANTEN ===================== */
type Resource string
type Pickaxe string
type Sword string
type Armor string

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

	WoodSword     Sword = "Holzschwert"
	StoneSword    Sword = "Steinschwert"
	IronSword     Sword = "Eisenschwert"
	GoldSword     Sword = "Goldschwert"
	PlatinumSword Sword = "Platinschwert"
	DiamondSword  Sword = "Diamantschwert"

	WoodArmor     Armor = "Holzrüstung"
	StoneArmor    Armor = "Steinrüstung"
	IronArmor     Armor = "Eisenrüstung"
	GoldArmor     Armor = "Goldrüstung"
	PlatinumArmor Armor = "Platinrüstung"
	DiamondArmor  Armor = "Diamantrüstung"
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
	mineUpgradeRate = 1
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
	StonePickaxe:    {Stone: 50},
	IronPickaxe:     {Iron: 50},
	GoldPickaxe:     {Gold: 50},
	PlatinumPickaxe: {Platinum: 50},
}

var swordAttack = map[Sword]int{
	WoodSword:     5,
	StoneSword:    10,
	IronSword:     20,
	GoldSword:     30,
	PlatinumSword: 40,
	DiamondSword:  50,
}

var armorDefense = map[Armor]int{
	WoodArmor:     2,
	StoneArmor:    5,
	IronArmor:     10,
	GoldArmor:     15,
	PlatinumArmor: 20,
	DiamondArmor:  30,
}

var swordUpgrade = map[Sword]Sword{
	WoodSword:     StoneSword,
	StoneSword:    IronSword,
	IronSword:     GoldSword,
	GoldSword:     PlatinumSword,
	PlatinumSword: DiamondSword,
}

var armorUpgrade = map[Armor]Armor{
	WoodArmor:     StoneArmor,
	StoneArmor:    IronArmor,
	IronArmor:     GoldArmor,
	GoldArmor:     PlatinumArmor,
	PlatinumArmor: DiamondArmor,
}

var swordUpgradeCost = map[Sword]map[Resource]int{
	WoodSword:     {Stone: 20},
	StoneSword:    {Iron: 20},
	IronSword:     {Gold: 20},
	GoldSword:     {Platinum: 20},
	PlatinumSword: {Diamond: 20},
}

var armorUpgradeCost = map[Armor]map[Resource]int{
	WoodArmor:     {Stone: 20},
	StoneArmor:    {Iron: 20},
	IronArmor:     {Gold: 20},
	GoldArmor:     {Platinum: 20},
	PlatinumArmor: {Diamond: 20},
}

var mineUpgradeResources = []Resource{Stone, Iron, Gold, Platinum, Diamond}

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
	Sword     Sword
	Armor     Armor
	Money     int
}

/* ===================== HELPER ===================== */
func dist2(a, b fyne.Position) float32 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return dx*dx + dy*dy
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

/* ===================== SPIEL ===================== */
func main() {
	a := app.New()
	w := a.NewWindow("Mining Game")
	w.Resize(fyne.NewSize(windowWidth, windowHeight))

	inv := &Inventory{
		Resources: make(map[Resource]int),
		Pickaxe:   WoodPickaxe,
		Sword:     WoodSword,
		Armor:     WoodArmor,
		Money:     200,
	}

	player := canvas.NewRectangle(color.RGBA{0, 200, 100, 255})
	player.Resize(fyne.NewSize(playerW, playerH))
	x, y := float32(50), float32(50)
	player.Move(fyne.NewPos(x, y))

	// Minen erstellen
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

	// Zentraler Tick
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
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
				if dist2(p, m.Pos) < interactionDist*interactionDist {
					openMineShop(a, m, inv)
					return
				}
			}
			if dist2(p, shopPos) < interactionDist*interactionDist {
				openShop(a, inv)
			} else if dist2(p, smithPos) < interactionDist*interactionDist {
				openSmith(a, inv)
			}
		case fyne.KeySpace:
			openInventory(a, inv)
		}

		x = clamp(x, 0, windowWidth-playerW)
		y = clamp(y, 0, windowHeight-playerH)
		player.Move(fyne.NewPos(x, y))

		for _, m := range mines {
			d := dist2(fyne.NewPos(x, y), m.Pos)
			if m.Owned {
				m.Icon.FillColor = color.RGBA{0, 255, 0, 255}
			} else if d < interactionDist*interactionDist {
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
	actionBtn := widget.NewButton("", nil)

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

	actionBtn.OnTapped = func() {
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
		m.Rate += mineUpgradeRate
		m.UpgradeLevel++
		updateLabel()
		if m.UpgradeLevel >= MaxMineUpgrade {
			actionBtn.Disable()
		}
	}

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
	w.Resize(fyne.NewSize(400, 500))

	info := widget.NewLabel("")
	msgLabel := widget.NewLabel("") // Für Fehler- oder Statusmeldungen

	pickaxeCostLabel := widget.NewLabel("")
	swordCostLabel := widget.NewLabel("")
	armorCostLabel := widget.NewLabel("")

	updateUI := func() {
		inv.Lock()
		defer inv.Unlock()

		info.SetText(fmt.Sprintf(
			"Geld: %d\nSpitzhacke: %s\nSchwert: %s (Attack: %d)\nRüstung: %s (Verteidigung: %d)",
			inv.Money,
			inv.Pickaxe,
			inv.Sword, swordAttack[inv.Sword],
			inv.Armor, armorDefense[inv.Armor],
		))

		updateLabel := func(current interface{}, upgradeMap interface{}, costMap interface{}, label *widget.Label, itemType string) {
			switch c := current.(type) {
			case Pickaxe:
				if next, ok := pickaxeUpgrade[c]; ok {
					text := fmt.Sprintf("Für Upgrade auf %s benötigst du:\n", next)
					for r, amt := range pickaxeResourceCost[c] {
						text += fmt.Sprintf("%s: %d (Du hast: %d)\n", r, amt, inv.Resources[r])
					}
					label.SetText(text)
				} else {
					label.SetText(itemType + " ist max. Level")
				}
			case Sword:
				if next, ok := swordUpgrade[c]; ok {
					text := fmt.Sprintf("Für Upgrade auf %s benötigst du:\n", next)
					for r, amt := range swordUpgradeCost[c] {
						text += fmt.Sprintf("%s: %d (Du hast: %d)\n", r, amt, inv.Resources[r])
					}
					label.SetText(text)
				} else {
					label.SetText(itemType + " ist max. Level")
				}
			case Armor:
				if next, ok := armorUpgrade[c]; ok {
					text := fmt.Sprintf("Für Upgrade auf %s benötigst du:\n", next)
					for r, amt := range armorUpgradeCost[c] {
						text += fmt.Sprintf("%s: %d (Du hast: %d)\n", r, amt, inv.Resources[r])
					}
					label.SetText(text)
				} else {
					label.SetText(itemType + " ist max. Level")
				}
			}
		}

		updateLabel(inv.Pickaxe, pickaxeUpgrade, pickaxeResourceCost, pickaxeCostLabel, "Spitzhacke")
		updateLabel(inv.Sword, swordUpgrade, swordUpgradeCost, swordCostLabel, "Schwert")
		updateLabel(inv.Armor, armorUpgrade, armorUpgradeCost, armorCostLabel, "Rüstung")
	}

	tryUpgrade := func(current interface{}, upgradeMap interface{}, costMap interface{}) string {
		inv.Lock()
		defer inv.Unlock()

		switch c := current.(type) {
		case Pickaxe:
			next, ok := pickaxeUpgrade[c]
			if !ok {
				return "Spitzhacke ist max. Level!"
			}
			for r, amt := range pickaxeResourceCost[c] {
				if inv.Resources[r] < amt {
					return "Nicht genügend Ressourcen für Spitzhacke!"
				}
			}
			for r, amt := range pickaxeResourceCost[c] {
				inv.Resources[r] -= amt
			}
			inv.Pickaxe = next
			return "Spitzhacke erfolgreich upgegradet!"
		case Sword:
			next, ok := swordUpgrade[c]
			if !ok {
				return "Schwert ist max. Level!"
			}
			for r, amt := range swordUpgradeCost[c] {
				if inv.Resources[r] < amt {
					return "Nicht genügend Ressourcen für Schwert!"
				}
			}
			for r, amt := range swordUpgradeCost[c] {
				inv.Resources[r] -= amt
			}
			inv.Sword = next
			return "Schwert erfolgreich upgegradet!"
		case Armor:
			next, ok := armorUpgrade[c]
			if !ok {
				return "Rüstung ist max. Level!"
			}
			for r, amt := range armorUpgradeCost[c] {
				if inv.Resources[r] < amt {
					return "Nicht genügend Ressourcen für Rüstung!"
				}
			}
			for r, amt := range armorUpgradeCost[c] {
				inv.Resources[r] -= amt
			}
			inv.Armor = next
			return "Rüstung erfolgreich upgegradet!"
		}
		return "Upgrade fehlgeschlagen"
	}

	pickaxeBtn := widget.NewButton("Upgrade Spitzhacke", func() {
		msgLabel.SetText(tryUpgrade(inv.Pickaxe, pickaxeUpgrade, pickaxeResourceCost))
		updateUI()
	})

	swordBtn := widget.NewButton("Upgrade Schwert", func() {
		msgLabel.SetText(tryUpgrade(inv.Sword, swordUpgrade, swordUpgradeCost))
		updateUI()
	})

	armorBtn := widget.NewButton("Upgrade Rüstung", func() {
		msgLabel.SetText(tryUpgrade(inv.Armor, armorUpgrade, armorUpgradeCost))
		updateUI()
	})

	box := container.NewVBox(info, msgLabel, pickaxeCostLabel, pickaxeBtn, swordCostLabel, swordBtn, armorCostLabel, armorBtn)
	w.SetContent(container.NewVScroll(box))
	updateUI()
	w.Show()
}

/* ===================== INVENTAR ===================== */
func openInventory(a fyne.App, inv *Inventory) {
	w := a.NewWindow("Inventar")
	w.Resize(fyne.NewSize(300, 400))
	box := container.NewVBox()

	inv.Lock()
	defer inv.Unlock()

	// Zeige Geld
	box.Add(widget.NewLabel(fmt.Sprintf("Geld: %d", inv.Money)))

	// Zeige Ausrüstung
	box.Add(widget.NewLabel(fmt.Sprintf("Spitzhacke: %s", inv.Pickaxe)))
	box.Add(widget.NewLabel(fmt.Sprintf("Schwert: %s (Attack: %d)", inv.Sword, swordAttack[inv.Sword])))
	box.Add(widget.NewLabel(fmt.Sprintf("Rüstung: %s (Verteidigung: %d)", inv.Armor, armorDefense[inv.Armor])))

	// Zeige Ressourcen
	box.Add(widget.NewLabel("--- Ressourcen ---"))
	for _, r := range []Resource{Stone, Iron, Gold, Platinum, Diamond} {
		box.Add(widget.NewLabel(fmt.Sprintf("%s: %d", r, inv.Resources[r])))
	}

	w.SetContent(container.NewVScroll(box))
	w.Show()
}
