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
	windowWidth     = 1470
	windowHeight    = 765
	playerW         = 20
	playerH         = 40
	speed           = 20
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
	WoodSword:     1,
	StoneSword:    3,
	IronSword:     5,
	GoldSword:     8,
	PlatinumSword: 10,
	DiamondSword:  13,
}

var armorDefense = map[Armor]int{
	WoodArmor:     1,
	StoneArmor:    3,
	IronArmor:     5,
	GoldArmor:     8,
	PlatinumArmor: 10,
	DiamondArmor:  13,
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

type Dungeon struct {
	Name  string
	Pos   fyne.Position
	Icon  *canvas.Rectangle
	Label *canvas.Text
}

type Inventory struct {
	sync.Mutex
	Resources map[Resource]int
	Pickaxe   Pickaxe
	Sword     Sword
	Armor     Armor
	Money     int
	Cristalle int
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

func enableBackspaceClose(w fyne.Window) {
	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == fyne.KeyBackspace || k.Name == fyne.KeyEscape {
			w.Close()
		}
	})
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
		Cristalle: 0,
	}

	player := canvas.NewRectangle(color.RGBA{0, 200, 100, 255})
	player.Resize(fyne.NewSize(playerW, playerH))
	x, y := float32(50), float32(50)
	player.Move(fyne.NewPos(x, y))

	// Spieler-Lebensanzeige
	playerHealth := 100
	healthBarWidth := float32(playerHealth) / 10
	healthBarHeight := float32(5)

	healthBar := canvas.NewRectangle(color.RGBA{255, 0, 0, 255}) // rot
	healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
	healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2)) // über dem Spieler

	objects := []fyne.CanvasObject{healthBar, player}

	// Minen erstellen
	mines := []*Mine{
		{"Steinmine", fyne.NewPos(100, 200), Stone, 100, 1, WoodPickaxe, false, 0, nil, nil},
		{"Eisenmine", fyne.NewPos(300, 200), Iron, 250, 1, StonePickaxe, false, 0, nil, nil},
		{"Goldmine", fyne.NewPos(500, 200), Gold, 500, 1, IronPickaxe, false, 0, nil, nil},
		{"Platinmine", fyne.NewPos(700, 200), Platinum, 1000, 1, GoldPickaxe, false, 0, nil, nil},
		{"Diamantmine", fyne.NewPos(900, 200), Diamond, 2000, 1, PlatinumPickaxe, false, 0, nil, nil},
	}

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

	// Dungeon erstellen
	dungeonPos := fyne.NewPos(1100, 500)
	dungeon := &Dungeon{
		Name: "Dungeon",
		Pos:  dungeonPos,
	}
	dungeon.Icon = canvas.NewRectangle(color.RGBA{200, 0, 0, 255})
	dungeon.Icon.Resize(fyne.NewSize(iconSize, iconSize))
	dungeon.Icon.Move(dungeon.Pos)
	dungeon.Icon.StrokeColor = color.Black
	dungeon.Icon.StrokeWidth = borderWidth

	dungeon.Label = canvas.NewText(dungeon.Name, color.Black)
	dungeon.Label.TextSize = 12
	dungeon.Label.Move(fyne.NewPos(dungeon.Pos.X, dungeon.Pos.Y+45))

	objects = append(objects, dungeon.Icon, dungeon.Label)

	// Shop & Schmied
	shopPos := fyne.NewPos(200, 20)
	shop := widget.NewButtonWithIcon("Shop", theme.InfoIcon(), func() { openShop(a, inv) })
	shop.Move(shopPos)
	shop.Resize(fyne.NewSize(80, 40))

	smithPos := fyne.NewPos(500, 20)
	smith := widget.NewButtonWithIcon("Schmied", theme.InfoIcon(), func() { openSmith(a, inv) })
	smith.Move(smithPos)
	smith.Resize(fyne.NewSize(110, 40))

	cristalShopPos := fyne.NewPos(800, 20)
	cristal := widget.NewButtonWithIcon("Cristal Shop", theme.InfoIcon(), func() { openCristallShop(a, inv) })
	cristal.Move(cristalShopPos)
	cristal.Resize(fyne.NewSize(110, 40))

	objects = append(objects, shop, smith, cristal)
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
				return
			}
			if dist2(p, smithPos) < interactionDist*interactionDist {
				openSmith(a, inv)
				return
			}
			if dist2(p, cristalShopPos) < interactionDist*interactionDist {
				openCristallShop(a, inv)
				return
			}
			if dist2(p, dungeon.Pos) < interactionDist*interactionDist {
				openDungeon(a, w, inv) // hier übergeben wir das Hauptfenster
				return
			}

		case fyne.KeySpace:
			openInventory(a, inv)
		}

		x = clamp(x, 0, windowWidth-playerW)
		y = clamp(y, 0, windowHeight-playerH)
		player.Move(fyne.NewPos(x, y))

		// Health Bar über dem Spieler aktualisieren
		healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))
		healthBar.Refresh()

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

		// Dungeon hervorheben, wenn nah
		if dist2(fyne.NewPos(x, y), dungeon.Pos) < interactionDist*interactionDist {
			dungeon.Icon.FillColor = color.RGBA{255, 255, 0, 255}
		} else {
			dungeon.Icon.FillColor = color.RGBA{200, 0, 0, 255}
		}
		dungeon.Icon.Refresh()
	})

	w.ShowAndRun()
}

/* ===================== MINE SHOP ===================== */
func openMineShop(a fyne.App, m *Mine, inv *Inventory) {
	w := a.NewWindow(m.Name)
	enableBackspaceClose(w)
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

/*====================== Cristal Shop ==================*/
func openCristallShop(a fyne.App, inv *Inventory) {
	w := a.NewWindow("Cristal Shop")
	enableBackspaceClose(w)

	prices := map[Resource]int{
		Stone:    1,
		Iron:     5,
		Gold:     10,
		Platinum: 20,
		Diamond:  50,
	}
	resourceOrder := []Resource{Stone, Iron, Gold, Platinum, Diamond}

	moneyLabel := widget.NewLabel(fmt.Sprintf("Geld: %d", inv.Money))
	moneyLabel.TextStyle = fyne.TextStyle{Bold: true}
	box := container.NewVBox(
		moneyLabel,
		widget.NewSeparator(),
		widget.NewLabel("Ressourcen verkaufen"),
	)
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
				moneyLabel.SetText(fmt.Sprintf("Geld: %d", inv.Money))

			})
			resBox.Add(btn)
		}
		box.Add(resBox)
	}

	w.SetContent(container.NewVScroll(box))
	w.Resize(fyne.NewSize(300, 400))
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
	enableBackspaceClose(w)
	moneyLabel := widget.NewLabel(fmt.Sprintf("Geld: %d", inv.Money))
	moneyLabel.TextStyle = fyne.TextStyle{Bold: true}
	box := container.NewVBox(
		moneyLabel,
		widget.NewSeparator(),
		widget.NewLabel("Ressourcen verkaufen"),
	)
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
				moneyLabel.SetText(fmt.Sprintf("Geld: %d", inv.Money))

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
			moneyLabel.SetText(fmt.Sprintf("Geld: %d", inv.Money))
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
	enableBackspaceClose(w)
	w.Resize(fyne.NewSize(600, 500))

	info := widget.NewLabel("")
	msgLabel := widget.NewLabel("") // Nur für Upgrade-Meldungen

	pickaxeCostLabel := widget.NewLabel("")
	swordCostLabel := widget.NewLabel("")
	armorCostLabel := widget.NewLabel("")

	// Update nur der Ressourcen / Kosten Labels
	updateCosts := func() {
		inv.Lock()
		defer inv.Unlock()

		info.SetText(fmt.Sprintf(
			"Geld: %d\nSpitzhacke: %s\nSchwert: %s (Attack: %d)\nRüstung: %s (Verteidigung: %d)",
			inv.Money,
			inv.Pickaxe,
			inv.Sword, swordAttack[inv.Sword],
			inv.Armor, armorDefense[inv.Armor],
		))

		updateLabel := func(current interface{}, label *widget.Label, itemType string) {
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

		updateLabel(inv.Pickaxe, pickaxeCostLabel, "Spitzhacke")
		updateLabel(inv.Sword, swordCostLabel, "Schwert")
		updateLabel(inv.Armor, armorCostLabel, "Rüstung")
	}

	tryUpgrade := func(current interface{}) string {
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
		msgLabel.SetText(tryUpgrade(inv.Pickaxe))
		updateCosts() // nur die Ressourcen Labels neu zeichnen
	})

	swordBtn := widget.NewButton("Upgrade Schwert", func() {
		msgLabel.SetText(tryUpgrade(inv.Sword))
		updateCosts()
	})

	armorBtn := widget.NewButton("Upgrade Rüstung", func() {
		msgLabel.SetText(tryUpgrade(inv.Armor))
		updateCosts()
	})

	box := container.NewVBox(info, msgLabel, pickaxeCostLabel, pickaxeBtn, swordCostLabel, swordBtn, armorCostLabel, armorBtn)
	w.SetContent(container.NewVScroll(box))
	updateCosts()
	w.Show()

	// Live-Update nur für die Ressourcen, msgLabel bleibt unberührt
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			updateCosts()
		}
	}()
}

/* ===================== INVENTAR ===================== */
func openInventory(a fyne.App, inv *Inventory) {
	w := a.NewWindow("Inventar")
	enableBackspaceClose(w)
	w.Resize(fyne.NewSize(300, 400))
	box := container.NewVBox()

	inv.Lock()
	// Labels für Geld und Ausrüstung
	moneyLabel := widget.NewLabel(fmt.Sprintf("Geld: %d", inv.Money))
	cristalleLabel := widget.NewLabel(fmt.Sprintf("Cristalle: %d", inv.Cristalle))
	pickaxeLabel := widget.NewLabel(fmt.Sprintf("Spitzhacke: %s", inv.Pickaxe))
	swordLabel := widget.NewLabel(fmt.Sprintf("Schwert: %s (Attack: %d)", inv.Sword, swordAttack[inv.Sword]))
	armorLabel := widget.NewLabel(fmt.Sprintf("Rüstung: %s (Verteidigung: %d)", inv.Armor, armorDefense[inv.Armor]))

	box.Add(moneyLabel)
	box.Add(cristalleLabel)
	box.Add(pickaxeLabel)
	box.Add(swordLabel)
	box.Add(armorLabel)

	// Ressourcen-Labels
	box.Add(widget.NewLabel("--- Ressourcen ---"))
	resourceLabels := make(map[Resource]*widget.Label)
	for _, r := range []Resource{Stone, Iron, Gold, Platinum, Diamond} {
		label := widget.NewLabel(fmt.Sprintf("%s: %d", r, inv.Resources[r]))
		resourceLabels[r] = label
		box.Add(label)
	}
	inv.Unlock()

	w.SetContent(container.NewVScroll(box))
	w.Show()

	// Live-Update
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			inv.Lock()
			moneyLabel.SetText(fmt.Sprintf("Geld: %d", inv.Money))
			cristalleLabel.SetText(fmt.Sprintf("Cristalle: %d", inv.Cristalle))
			pickaxeLabel.SetText(fmt.Sprintf("Spitzhacke: %s", inv.Pickaxe))
			swordLabel.SetText(fmt.Sprintf("Schwert: %s (Attack: %d)", inv.Sword, swordAttack[inv.Sword]))
			armorLabel.SetText(fmt.Sprintf("Rüstung: %s (Verteidigung: %d)", inv.Armor, armorDefense[inv.Armor]))
			for r, label := range resourceLabels {
				label.SetText(fmt.Sprintf("%s: %d", r, inv.Resources[r]))
			}
			inv.Unlock()
		}
	}()
}

/* ===================== DUNGEON ===================== */
func openDungeon(a fyne.App, mainWindow fyne.Window, inv *Inventory) {
	// Hauptfenster ausblenden, nicht schließen
	mainWindow.Hide()

	w := a.NewWindow("Dungeon")
	w.Resize(fyne.NewSize(400, 300))

	// Funktion zum Schließen des Dungeon-Fensters
	closeDungeon := func() {
		w.Close()
		mainWindow.Show() // Hauptfenster wieder anzeigen
	}

	// Lebenspunkte des Spielers
	var playerHealth int = 100
	playerHealth = playerHealth / 10 * armorDefense[inv.Armor]
	healthLabel := widget.NewLabel(fmt.Sprintf("%d", playerHealth))

	adventureBtn := widget.NewButton("Ebene 1", func() {
		w.Hide()
		w2 := a.NewWindow("Ebene 1")
		w2.Resize(fyne.NewSize(400, 300))

		// Spieler erstellen
		player := canvas.NewRectangle(color.RGBA{0, 200, 100, 255})
		player.Resize(fyne.NewSize(playerW, playerH))
		x, y := float32(50), float32(50)
		player.Move(fyne.NewPos(x, y))

		// Health Bar erstellen
		healthBarWidth := float32(playerHealth) // Berechnung des Health Bar mit Skalierung
		healthBarHeight := float32(5)

		healthBar := canvas.NewRectangle(color.RGBA{255, 0, 0, 255}) // rot
		healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
		healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2)) // über dem Spieler

		objects := []fyne.CanvasObject{player, healthLabel}

		// Gegner (rote Blöcke) erstellen
		type Enemy struct {
			Rect      *canvas.Rectangle
			Health    int
			HealthBar *canvas.Rectangle
			X, Y      float32
			Speed     float32
		}

		// Gegner erstellen und Lebensbalken hinzufügen
		enemies := []*Enemy{
			{Rect: canvas.NewRectangle(color.RGBA{255, 0, 0, 255}), X: 300, Y: 50, Speed: 2, Health: 50},
			{Rect: canvas.NewRectangle(color.RGBA{255, 0, 0, 255}), X: 200, Y: 200, Speed: 1.5, Health: 50},
		}

		for _, e := range enemies {
			e.Rect.Resize(fyne.NewSize(20, 20))
			e.Rect.Move(fyne.NewPos(e.X, e.Y))

			// Health Bar für den Gegner erstellen
			e.HealthBar = canvas.NewRectangle(color.RGBA{255, 0, 0, 255})           // rot
			e.HealthBar.Resize(fyne.NewSize(float32(e.Health)/10, healthBarHeight)) // Skalierung des HealthBars
			e.HealthBar.Move(fyne.NewPos(e.X-5, e.Y-e.HealthBar.Size().Height-2))   // Position über dem Gegner

			// Gegner + HealthBar zu Objekten hinzufügen
			objects = append(objects, e.Rect, e.HealthBar)
		}

		// Health Label anpassen, damit es sichtbar ist
		healthLabel.Move(fyne.NewPos(150, 10)) // Position des Labels oben im Fenster

		w2.SetContent(container.NewWithoutLayout(objects...))

		// Spielerbewegung
		w2.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
			speed := float32(10)
			switch k.Name {
			case fyne.KeyUp:
				y -= speed
			case fyne.KeyDown:
				y += speed
			case fyne.KeyLeft:
				x -= speed
			case fyne.KeyRight:
				x += speed
			case fyne.KeyEscape, fyne.KeyBackspace: // Dungeon zurück
				w2.Close()
				w.Show() // Hauptfenster wieder anzeigen
				return
			}

			// Health Bar über dem Spieler aktualisieren
			healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))
			healthBar.Refresh()

			// Position begrenzen
			x = clamp(x, 0, 400-playerW)
			y = clamp(y, 0, 300-playerH)
			player.Move(fyne.NewPos(x, y))
			player.Refresh()

			// Health Bar anpassen
			healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight)) // Breite anpassen
			healthBar.Refresh()

			// Lebenspunkte anzeigen
			healthLabel.SetText(fmt.Sprintf("%d", int(healthBarWidth))) // Update health label
		})

		// Ticker, der alle 2 Sekunden den Health Bar reduziert
		ticker := time.NewTicker(1 * time.Second)
		go func() {
			for range ticker.C {
				// Health Bar breiter machen (überprüfen, ob sie nicht unter 0 geht)
				if healthBarWidth > 0 {
					playerHealth -= 1
					healthBarWidth = float32(playerHealth)
					// Health Bar anpassen
					healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
					healthBar.Refresh()

					// Lebenspunkte anzeigen
					healthLabel.SetText(fmt.Sprintf("%d", int(healthBarWidth))) // Update health label
				}
				// Wenn die Lebenspunkte 0 erreichen, stoppen wir den Ticker
				if healthBarWidth == 0 {
					healthLabel.SetText("Du bist gestorben!")
					ticker.Stop() // Stoppe den Ticker, wenn der Spieler gestorben ist
					// Hauptfenster anzeigen und Dungeon-Fenster schließen
					time.Sleep(2 * time.Second)
					w2.Close() // Dungeon-Fenster schließen
					closeDungeon()
					return
				}

				// Hier alle 2 Sekunden wiederholen
				time.Sleep(2 * time.Second)
			}
		}()

		// Gegner bewegen und Schaden nehmen, wenn der Spieler in der Nähe ist
		go func() {
			for {
				allEnemiesDead := true // Flag, um zu prüfen, ob alle Gegner tot sind
				for _, e := range enemies {
					// Berechne die Distanz zwischen dem Spieler und dem Gegner
					distance := math.Sqrt(float64((x-e.X)*(x-e.X) + (y-e.Y)*(y-e.Y)))

					// Wenn der Spieler nahe genug am Gegner ist (z.B. Abstand < 50)
					if distance < 50 {
						// Gegner verliert Lebenspunkte jede Sekunde
						if e.Health > 0 {
							e.Health -= 10 * swordAttack[inv.Sword]
							// Health Bar des Gegners aktualisieren
							e.HealthBar.Resize(fyne.NewSize(float32(e.Health)/10, healthBarHeight)) // Health Bar anpassen
							e.HealthBar.Refresh()                                                   // Health Bar aktualisieren
						}
					}

					// Wenn der Gegner keine Lebenspunkte mehr hat
					if e.Health <= 0 {
						e.Rect.FillColor = color.Gray{Y: 200} // Gegner "sterben"
					} else {
						// Wenn der Gegner noch lebt, setzen wir das Flag auf false
						allEnemiesDead = false
					}
				}

				// Wenn alle Gegner tot sind, schließen wir das Dungeon-Fenster
				if allEnemiesDead {
					healthLabel.SetText("Du hast Gewonnen!")
					ticker.Stop() // Stoppe den Ticker
					inv.Cristalle += 1
					time.Sleep(2 * time.Second) // Kurze Pause, damit der Spieler den Tod der Gegner sehen kann
					w2.Close()                  // Dungeon-Fenster schließen
					closeDungeon()
					return
				}

				// Alle 1 Sekunde Schaden zugefügt werden
				time.Sleep(1 * time.Second)
			}
		}()

		w2.Show()
	})

	w.SetContent(container.NewVBox(
		adventureBtn,
		widget.NewButton("Dungeon verlassen", closeDungeon),
	))

	// Backspace oder Escape schließt auch das Dungeon
	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == fyne.KeyBackspace || k.Name == fyne.KeyEscape {
			closeDungeon()
		}
	})

	w.Show()
}
