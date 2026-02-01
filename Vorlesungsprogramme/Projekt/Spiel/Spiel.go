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
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

/* ===================== TYPEN & KONSTANTEN ===================== */
type Resource string
type Pickaxe string
type Sword string
type Armor string

const (
	Kohle                     Resource = "Kohle"
	Stone                     Resource = "Stein"
	Kupfererz                 Resource = "Kupferore"
	Kupferbarren              Resource = "Kupferbarren"
	Eisenerz                  Resource = "Eisenerz"
	Eisenbarren               Resource = "Eisenbarren"
	Stahlbarren               Resource = "Stahlbarren"
	Silvererz                 Resource = "Silbererz"
	Silberbarren              Resource = "Silberbarren"
	Golderz                   Resource = "Golderz"
	Goldbarren                Resource = "Goldbarren"
	Platinerz                 Resource = "Platinerz"
	Platinbarren              Resource = "Platinbarren"
	Rubinerz                  Resource = "Rubinerz"
	GeschliffenerRubin        Resource = "Geschliffener Rubin"
	Diamanterz                Resource = "Diamanterz"
	GeschliffenerDiamant      Resource = "GeschliffenerDiamant"
	Mythrilerz                Resource = "Mythrilerz"
	Mythrilbarren             Resource = "Mythrilbarren "
	Blutkristallerz           Resource = "Blutkristallerz"
	GeschliffenerBlutkristall Resource = "GeschliffenerBlutkristall"
	Adamantiumerz             Resource = "Adamantiumerz"
	Adamantiumbarren          Resource = "Adamantiumbarren "

	WoodPickaxe      Pickaxe = "Holzspitzhacke"
	StonePickaxe     Pickaxe = "Steinspitzhacke"
	KupferPickaxe    Pickaxe = "Kupferspitzhacke"
	IronPickaxe      Pickaxe = "Eisenspitzhacke"
	StahlPickaxe     Pickaxe = "Stahlspitzhacke"
	GoldPickaxe      Pickaxe = "Goldspitzhacke"
	PlatinumPickaxe  Pickaxe = "Platinspitzhacke"
	DiamondPickaxe   Pickaxe = "Diamantspitzhacke"
	MythrilPickax    Pickaxe = "Mythrilspitzhacke"
	AdamantiumPickax Pickaxe = "Adamantiumspitzhacke"

	WoodSword       Sword = "Holzschwert"
	StoneSword      Sword = "Steinschwert"
	KupferSword     Sword = "Kupferschwert"
	IronSword       Sword = "Eisenschwert"
	StahlSword      Sword = "Stahlschwert"
	GoldSword       Sword = "Goldschwert"
	PlatinumSword   Sword = "Platinschwert"
	DiamondSword    Sword = "Diamantschwert"
	MythrilSword    Sword = "Mythrilschwert"
	AdamantiumSword Sword = "Adamantiumschwert"

	WoodArmor       Armor = "Holzrüstung"
	StoneArmor      Armor = "Steinrüstung"
	KupferArmor     Armor = "Kupferrüstung"
	IronArmor       Armor = "Eisenrüstung"
	StahlArmor      Armor = "Stahlrüstung"
	GoldArmor       Armor = "Goldrüstung"
	PlatinumArmor   Armor = "Platinrüstung"
	DiamondArmor    Armor = "Diamantrüstung"
	MythrilArmor    Armor = "Mythrilrüstung"
	AdamantiumArmor Armor = "Adamantiumrüstung"
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

var pickaxeResourceCost = map[Pickaxe]map[Resource]int{
	WoodPickaxe:      {Stone: 50},
	StonePickaxe:     {Stone: 50},
	KupferPickaxe:    {Kupferbarren: 50},
	IronPickaxe:      {Eisenbarren: 50},
	StahlPickaxe:     {Stahlbarren: 50},
	GoldPickaxe:      {Goldbarren: 50},
	PlatinumPickaxe:  {Platinbarren: 50},
	DiamondPickaxe:   {GeschliffenerDiamant: 50},
	MythrilPickax:    {Mythrilbarren: 50},
	AdamantiumPickax: {Adamantiumbarren: 50},
}

var swordUpgradeCost = map[Sword]map[Resource]int{
	WoodSword:       {Stone: 20},
	StoneSword:      {Kupferbarren: 20},
	KupferSword:     {Eisenbarren: 20},
	IronSword:       {Stahlbarren: 20},
	StahlSword:      {Goldbarren: 20},
	GoldSword:       {Platinbarren: 20},
	PlatinumSword:   {GeschliffenerDiamant: 20},
	DiamondSword:    {Mythrilbarren: 20},
	MythrilSword:    {Adamantiumbarren: 20},
	AdamantiumSword: {Adamantiumbarren: 100},
}

var armorUpgradeCost = map[Armor]map[Resource]int{
	WoodArmor:       {Stone: 20},
	StoneArmor:      {Kupferbarren: 20},
	KupferArmor:     {Eisenbarren: 20},
	IronArmor:       {Stahlbarren: 20},
	StahlArmor:      {Goldbarren: 20},
	GoldArmor:       {Platinbarren: 20},
	PlatinumArmor:   {GeschliffenerDiamant: 20},
	DiamondArmor:    {Mythrilbarren: 20},
	MythrilArmor:    {Adamantiumbarren: 20},
	AdamantiumArmor: {Adamantiumbarren: 100},
}

var swordAttack = map[Sword]int{
	WoodSword:       1,
	StoneSword:      3,
	KupferSword:     5,
	IronSword:       7,
	StahlSword:      10,
	GoldSword:       12,
	PlatinumSword:   15,
	DiamondSword:    18,
	MythrilSword:    21,
	AdamantiumSword: 25,
}

var armorDefense = map[Armor]int{
	WoodArmor:       1,
	StoneArmor:      3,
	KupferArmor:     5,
	IronArmor:       7,
	StahlArmor:      10,
	GoldArmor:       12,
	PlatinumArmor:   15,
	DiamondArmor:    18,
	MythrilArmor:    21,
	AdamantiumArmor: 25,
}

var pickaxeUpgrade = map[Pickaxe]Pickaxe{
	WoodPickaxe:     StonePickaxe,
	StonePickaxe:    KupferPickaxe,
	KupferPickaxe:   IronPickaxe,
	IronPickaxe:     StahlPickaxe,
	StahlPickaxe:    GoldPickaxe,
	GoldPickaxe:     PlatinumPickaxe,
	PlatinumPickaxe: DiamondPickaxe,
	DiamondPickaxe:  MythrilPickax,
	MythrilPickax:   AdamantiumPickax,
}

var swordUpgrade = map[Sword]Sword{
	WoodSword:     StoneSword,
	StoneSword:    KupferSword,
	KupferSword:   IronSword,
	IronSword:     StahlSword,
	StahlSword:    GoldSword,
	GoldSword:     PlatinumSword,
	PlatinumSword: DiamondSword,
	DiamondSword:  MythrilSword,
	MythrilSword:  AdamantiumSword,
}

var armorUpgrade = map[Armor]Armor{
	WoodArmor:     StoneArmor,
	StoneArmor:    KupferArmor,
	KupferArmor:   IronArmor,
	IronArmor:     StahlArmor,
	StahlArmor:    GoldArmor,
	GoldArmor:     PlatinumArmor,
	PlatinumArmor: DiamondArmor,
	DiamondArmor:  MythrilArmor,
	MythrilArmor:  AdamantiumArmor,
}

var mineUpgradeResources = []Resource{Kohle, Stone, Kupferbarren, Eisenbarren, Stahlbarren, Silberbarren, Goldbarren, Platinbarren, GeschliffenerRubin, GeschliffenerDiamant, Mythrilbarren, GeschliffenerBlutkristall, Adamantiumbarren}

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
	Kristalle int
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
		Money:     0,
		Kristalle: 0,
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
		{"Kohlemine", fyne.NewPos(100, 200), Kohle, 0, 1, WoodPickaxe, false, 0, nil, nil},
		{"Steinmine", fyne.NewPos(300, 200), Stone, 100, 1, WoodPickaxe, false, 0, nil, nil},
		{"Kupfermine", fyne.NewPos(500, 200), Kupfererz, 200, 1, StonePickaxe, false, 0, nil, nil},
		{"Eisenmine", fyne.NewPos(700, 200), Eisenerz, 250, 1, KupferPickaxe, false, 0, nil, nil},
		{"Silbernmine", fyne.NewPos(900, 200), Silvererz, 500, 1, IronPickaxe, false, 0, nil, nil},
		{"Goldmine", fyne.NewPos(1100, 200), Golderz, 500, 1, IronPickaxe, false, 0, nil, nil},
		{"Platinmine", fyne.NewPos(100, 400), Platinerz, 1000, 1, GoldPickaxe, false, 0, nil, nil},
		{"Rubinnmine", fyne.NewPos(300, 400), Rubinerz, 2000, 1, PlatinumPickaxe, false, 0, nil, nil},
		{"Diamantmine", fyne.NewPos(500, 400), Diamanterz, 3000, 1, PlatinumPickaxe, false, 0, nil, nil},
		{"Mythrilmine", fyne.NewPos(700, 400), Mythrilerz, 5000, 1, DiamondPickaxe, false, 0, nil, nil},
		{"Blutkristallmine", fyne.NewPos(900, 400), Blutkristallerz, 10000, 1, DiamondPickaxe, false, 0, nil, nil},
		{"Adamantiummine", fyne.NewPos(1100, 400), Adamantiumerz, 100000, 1, MythrilPickax, false, 0, nil, nil},
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
	dungeonPos := fyne.NewPos(700, 600)
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
	shopPos := fyne.NewPos(100, 20)
	shop := widget.NewButtonWithIcon("Shop", theme.InfoIcon(), func() { openShop(a, inv) })
	shop.Move(shopPos)
	shop.Resize(fyne.NewSize(80, 40))

	smithPos := fyne.NewPos(400, 20)
	smith := widget.NewButtonWithIcon("Schmiede", theme.InfoIcon(), func() { openSmith(a, inv) })
	smith.Move(smithPos)
	smith.Resize(fyne.NewSize(110, 40))

	kristallShopPos := fyne.NewPos(700, 20)
	kristall := widget.NewButtonWithIcon("Kristall Shop", theme.InfoIcon(), func() { openKristallShop(a, inv) })
	kristall.Move(kristallShopPos)
	kristall.Resize(fyne.NewSize(110, 40))

	schmelzofenPos := fyne.NewPos(1000, 20)
	schmelzofen := widget.NewButtonWithIcon("Schmelzofen", theme.InfoIcon(), func() { openSchmelzofen(a, inv) })
	schmelzofen.Move(schmelzofenPos)
	schmelzofen.Resize(fyne.NewSize(110, 40))

	objects = append(objects, shop, smith, kristall, schmelzofen)
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
			if dist2(p, kristallShopPos) < interactionDist*interactionDist {
				openKristallShop(a, inv)
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
			WoodPickaxe:      1,
			StonePickaxe:     2,
			KupferPickaxe:    3,
			IronPickaxe:      4,
			StahlPickaxe:     5,
			GoldPickaxe:      6,
			PlatinumPickaxe:  7,
			DiamondPickaxe:   8,
			MythrilPickax:    9,
			AdamantiumPickax: 10,
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
// Globale Map für die gespeicherten Levels
var currentLevels = make(map[string]int)

func openKristallShop(a fyne.App, inv *Inventory) {
	// Preis für jede Fähigkeit (in Kristallen)
	skills := map[string]map[int]int{
		"Steinminen Multiplier":  {1: 10, 2: 25, 3: 50},
		"Eisenminen Multiplier":  {1: 15, 2: 35, 3: 70},
		"Goldminen Multiplier":   {1: 20, 2: 45, 3: 90},
		"Platinmine Multiplier":  {1: 10, 2: 25, 3: 50},
		"Diamantmine Multiplier": {1: 15, 2: 35, 3: 70},
	}

	// Level-Up-Funktionen für Skills
	levelUpFunctions := map[string]map[int]func(){
		"Steinminen Multiplier": {
			1: func() { fmt.Println("Steinminen Level 1 aktiviert - Mehr Ressourcen sammeln!") },
			2: func() { fmt.Println("Steinminen Level 2 aktiviert - Kristallgewinn erhöht!") },
			3: func() { fmt.Println("Steinminen Level 3 aktiviert - Maximale Ressourcengewinnung!") },
		},
		"Eisenminen Multiplier": {
			1: func() { fmt.Println("Eisenminen Level 1 aktiviert - Eisenproduktion startet.") },
			2: func() { fmt.Println("Eisenminen Level 2 aktiviert - Eisenproduktion verbessert.") },
			3: func() { fmt.Println("Eisenminen Level 3 aktiviert - Maximale Eisenproduktion.") },
		},
		"Goldminen Multiplier": {
			1: func() { fmt.Println("Goldminen Level 1 aktiviert - Goldproduktion startet.") },
			2: func() { fmt.Println("Goldminen Level 2 aktiviert - Goldproduktion optimiert.") },
			3: func() { fmt.Println("Goldminen Level 3 aktiviert - Maximale Goldproduktion.") },
		},
		"Platinmine Multiplier": {
			1: func() { fmt.Println("Platinminen Level 1 aktiviert - Platinproduktion startet.") },
			2: func() { fmt.Println("Platinminen Level 2 aktiviert - Platinproduktion verbessert.") },
			3: func() { fmt.Println("Platinminen Level 3 aktiviert - Maximale Platinproduktion.") },
		},
		"Diamantmine Multiplier": {
			1: func() { fmt.Println("Diamantminen Level 1 aktiviert - Diamantproduktion startet.") },
			2: func() { fmt.Println("Diamantminen Level 2 aktiviert - Diamantproduktion optimiert.") },
			3: func() { fmt.Println("Diamantminen Level 3 aktiviert - Maximale Diamantproduktion.") },
		},
	}

	// Fenster erstellen
	w := a.NewWindow("Skill Tree Shop")
	enableBackspaceClose(w)

	// Kristallanzeige
	kristallLabel := createKristallLabel(inv)

	// Shop-Layout
	box := container.NewVBox(
		kristallLabel,
		widget.NewSeparator(),
		widget.NewLabel("Fähigkeiten freischalten"),
	)

	// Buttons für alle Fähigkeiten
	for skill, levels := range skills {
		// Lade das aktuelle Level aus der globalen Map
		currentLevel := currentLevels[skill]

		// Button und Anzeige für jedes Skill erstellen
		createSkillButton(skill, levels, currentLevel, inv, kristallLabel, levelUpFunctions, w, box)
	}

	// Fenster anzeigen
	w.SetContent(box)
	w.Resize(fyne.NewSize(400, 600))
	w.Show()
}

// Hilfsfunktion für die Kristallanzeige
func createKristallLabel(inv *Inventory) *widget.Label {
	label := widget.NewLabel(fmt.Sprintf("Kristalle: %d", inv.Kristalle))
	label.TextStyle = fyne.TextStyle{Bold: true}
	return label
}

// Hilfsfunktion für die Erstellung von Buttons
func createSkillButton(skill string, levels map[int]int, currentLevel int, inv *Inventory, kristallLabel *widget.Label, levelUpFunctions map[string]map[int]func(), w fyne.Window, box *fyne.Container) {
	skillBox := container.NewVBox()
	skillLabel := widget.NewLabel(fmt.Sprintf("%s", skill))
	skillBox.Add(skillLabel)

	// Button erstellen
	btn := widget.NewButton(fmt.Sprintf("%s Level %d (Kosten: %d Kristalle)", skill, currentLevel+1, levels[currentLevel+1]), nil)

	// Button-Event
	btn.OnTapped = func() {
		handleSkillUnlock(skill, levels, &currentLevel, inv, kristallLabel, levelUpFunctions, btn, skillLabel, w)
	}

	// Button hinzufügen
	skillBox.Add(btn)

	// SkillBox zur Hauptbox hinzufügen
	box.Add(skillBox)

	// Anzeige des Levels aktualisieren
	updateSkillDisplay(skill, currentLevel, levels, btn, skillLabel)
}

// Funktion zur Aktualisierung der Anzeige für den Skill
func updateSkillDisplay(skill string, currentLevel int, levels map[int]int, btn *widget.Button, skillLabel *widget.Label) {
	// Button und Label mit den aktuellen Level-Informationen aktualisieren
	btn.SetText(fmt.Sprintf("%s Level %d (Kosten: %d Kristalle)", skill, currentLevel+1, levels[currentLevel+1]))
	skillLabel.SetText(fmt.Sprintf("%s (Level %d freigeschaltet)", skill, currentLevel+1))
}

// Funktion für das Freischalten des Skills
func handleSkillUnlock(skill string, levels map[int]int, currentLevel *int, inv *Inventory, kristallLabel *widget.Label, levelUpFunctions map[string]map[int]func(), btn *widget.Button, skillLabel *widget.Label, w fyne.Window) {
	inv.Lock()
	defer inv.Unlock()

	// Prüfen, ob genügend Kristalle und Level verfügbar sind
	if *currentLevel < len(levels) && inv.Kristalle >= levels[*currentLevel+1] {
		// Kristalle abziehen und Text aktualisieren
		inv.Kristalle -= levels[*currentLevel+1]
		kristallLabel.SetText(fmt.Sprintf("Kristalle: %d", inv.Kristalle))

		// Level erhöhen und Button-Text anpassen
		*currentLevel++
		updateSkillDisplay(skill, *currentLevel, levels, btn, skillLabel)

		// Level-Up-Funktion ausführen
		levelUpFunctions[skill][*currentLevel]()

		// Aktuelles Level in der globalen Map speichern
		currentLevels[skill] = *currentLevel
	} else if *currentLevel < len(levels) {
		// Fehlermeldung, wenn nicht genügend Kristalle
		dialog.ShowInformation("Nicht genug Kristalle", "Du hast nicht genügend Kristalle, um dieses Level freizuschalten.", w)
	} else {
		// Maximales Level erreicht
		dialog.ShowInformation("Maximales Level erreicht", "Du hast bereits das höchste Level für dieses Skill freigeschaltet.", w)
	}
}

/* ===================== SHOP ===================== */
func openShop(a fyne.App, inv *Inventory) {

	prices := map[Resource]int{
		Kohle:                     1,
		Stone:                     1,
		Kupfererz:                 2,
		Kupferbarren:              2,
		Eisenerz:                  5,
		Eisenbarren:               5,
		Stahlbarren:               5,
		Silvererz:                 10,
		Silberbarren:              10,
		Golderz:                   10,
		Goldbarren:                10,
		Platinerz:                 20,
		Platinbarren:              20,
		Rubinerz:                  30,
		GeschliffenerRubin:        30,
		Diamanterz:                50,
		GeschliffenerDiamant:      50,
		Mythrilerz:                100,
		Mythrilbarren:             100,
		Blutkristallerz:           200,
		GeschliffenerBlutkristall: 200,
		Adamantiumerz:             1000,
		Adamantiumbarren:          1000,
	}

	// 🟩 Kategorien
	edelsteine := []Resource{
		GeschliffenerRubin,
		GeschliffenerDiamant,
		GeschliffenerBlutkristall,
	}

	barren := []Resource{
		Kupferbarren,
		Eisenbarren,
		Stahlbarren,
		Silberbarren,
		Goldbarren,
		Platinbarren,
		Mythrilbarren,
		Adamantiumbarren,
	}

	erze := []Resource{
		Kohle,
		Stone,
		Kupfererz,
		Eisenerz,
		Silvererz,
		Golderz,
		Platinerz,
		Rubinerz,
		Diamanterz,
		Mythrilerz,
		Blutkristallerz,
		Adamantiumerz,
	}

	w := a.NewWindow("Shop")
	enableBackspaceClose(w)

	moneyLabel := widget.NewLabel("")
	moneyLabel.TextStyle = fyne.TextStyle{Bold: true}

	labels := make(map[Resource]*widget.Label)
	sellAmounts := []int{1, 10, 100}

	// Hilfsfunktion für Verkaufsblock
	createSellBlock := func(r Resource) *fyne.Container {
		resBox := container.NewVBox()
		label := widget.NewLabel("")
		labels[r] = label
		resBox.Add(label)

		price := prices[r]

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
					inv.Money += sell * price
				}

				label.SetText(fmt.Sprintf("%s: %d", r, inv.Resources[r]))
				moneyLabel.SetText(fmt.Sprintf("Geld: %d", inv.Money))
			})
			resBox.Add(btn)
		}

		resBox.Add(widget.NewSeparator())
		return resBox
	}

	// 🟦 Spalten
	leftBox := container.NewVBox(widget.NewLabel("💎 Geschliffene Edelsteine"))
	midBox := container.NewVBox(widget.NewLabel("🔩 Barren"))
	rightBox := container.NewVBox(widget.NewLabel("⛏️ Erze & Rohstoffe"))

	for _, r := range edelsteine {
		leftBox.Add(createSellBlock(r))
	}
	for _, r := range barren {
		midBox.Add(createSellBlock(r))
	}
	for _, r := range erze {
		rightBox.Add(createSellBlock(r))
	}

	// Layout
	content := container.NewBorder(
		container.NewHBox(moneyLabel),
		nil, nil, nil,
		container.NewGridWithColumns(
			3,
			container.NewVScroll(rightBox),
			container.NewVScroll(midBox),
			container.NewVScroll(leftBox),
		),
	)

	// UI Refresh
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

	w.SetContent(content)
	w.Resize(fyne.NewSize(900, 450))
	w.Show()
}

/* ===================Schmelz ofen =================== */
func openSchmelzofen(a fyne.App, inv *Inventory) {
	w := a.NewWindow("Schmelzofen")
	enableBackspaceClose(w)

	// Kohle-Kosten pro Einheit
	KohleKosten := map[Resource]int{
		Kupfererz:     1,
		Eisenerz:      1,
		Eisenbarren:   2,
		Silvererz:     2,
		Golderz:       3,
		Platinerz:     4,
		Mythrilerz:    6,
		Adamantiumerz: 10,
	}

	// Erz -> Barren
	OreSchmelzen := map[Resource]Resource{
		Kupfererz:     Kupferbarren,
		Eisenerz:      Eisenbarren,
		Eisenbarren:   Stahlbarren,
		Silvererz:     Silberbarren,
		Golderz:       Goldbarren,
		Platinerz:     Platinbarren,
		Mythrilerz:    Mythrilbarren,
		Adamantiumerz: Adamantiumbarren,
	}

	// Kristall -> geschliffen
	Kristalleschleifen := map[Resource]Resource{
		Rubinerz:        GeschliffenerRubin,
		Diamanterz:      GeschliffenerDiamant,
		Blutkristallerz: GeschliffenerBlutkristall,
	}

	resourceOrderMetall := []Resource{
		Kupfererz, Eisenerz, Eisenbarren,
		Silvererz, Golderz, Platinerz,
		Mythrilerz, Adamantiumerz,
	}

	resourceOrderKristall := []Resource{
		Rubinerz, Diamanterz, Blutkristallerz,
	}

	labels := make(map[Resource]*widget.Label)
	sellAmounts := []int{1, 10, 100}
	coalLabel := widget.NewLabel("")
	coalLabel.TextStyle = fyne.TextStyle{Bold: true}

	// 🔥 Metall-Spalte
	metallBox := container.NewVBox(
		widget.NewLabel("🔥 Brennstoff"),
		coalLabel,
		widget.NewLabel("🔥 Metalle schmelzen"),
	)

	// 💎 Kristall-Spalte
	kristallBox := container.NewVBox(
		widget.NewSeparator(),
		widget.NewLabel("💎 Kristalle schleifen"),
	)

	// Block-Erzeugung
	createResourceBlock := func(input Resource, output Resource, brauchtKohle bool) *fyne.Container {
		resBox := container.NewVBox()

		title := widget.NewLabel(fmt.Sprintf("%s ➜ %s", input, output))
		title.TextStyle = fyne.TextStyle{Bold: true}
		resBox.Add(title)

		label := widget.NewLabel("")
		labels[input] = label
		resBox.Add(label)

		var localCoalLabel *widget.Label
		coalCost := 0
		if brauchtKohle {
			coalCost = KohleKosten[input]
			localCoalLabel = widget.NewLabel(fmt.Sprintf("Benötigt Kohle pro Einheit: %d", coalCost))
			resBox.Add(localCoalLabel)
		}

		for _, amt := range sellAmounts {
			amount := amt
			btn := widget.NewButton(fmt.Sprintf("%d verarbeiten", amount), func() {
				inv.Lock()
				defer inv.Unlock()

				if inv.Resources[input] <= 0 {
					dialog.ShowInformation(
						"Zu wenig Material",
						fmt.Sprintf("Du hast kein %s mehr.", input),
						w,
					)
					return
				}

				use := amount
				if inv.Resources[input] < use {
					use = inv.Resources[input]
				}

				if brauchtKohle {
					maxByCoal := inv.Resources[Kohle] / coalCost
					if maxByCoal <= 0 {
						dialog.ShowInformation(
							"Zu wenig Kohle",
							fmt.Sprintf(
								"Du brauchst %d Kohle pro Einheit, hast aber nur %d.",
								coalCost,
								inv.Resources[Kohle],
							),
							w,
						)
						return
					}
					if maxByCoal < use {
						use = maxByCoal
					}
				}

				if use <= 0 {
					return
				}

				// Ressourcen umwandeln
				inv.Resources[input] -= use
				inv.Resources[output] += use

				if brauchtKohle {
					inv.Resources[Kohle] -= use * coalCost
				}

				label.SetText(fmt.Sprintf("%s: %d", input, inv.Resources[input]))
			})
			resBox.Add(btn)
		}

		resBox.Add(widget.NewSeparator())
		return resBox
	}

	// Metalle füllen
	for _, r := range resourceOrderMetall {
		if out, ok := OreSchmelzen[r]; ok {
			metallBox.Add(createResourceBlock(r, out, true))
		}
	}

	// Kristalle füllen
	for _, r := range resourceOrderKristall {
		if out, ok := Kristalleschleifen[r]; ok {
			kristallBox.Add(createResourceBlock(r, out, false))
		}
	}

	// 🔲 Layout: nebeneinander
	content := container.NewGridWithColumns(
		2,
		container.NewVScroll(metallBox),
		container.NewVScroll(kristallBox),
	)

	// UI Refresh
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			inv.Lock()

			// Kohle immer aktualisieren
			coalLabel.SetText(fmt.Sprintf("Kohle: %d", inv.Resources[Kohle]))

			// Ressourcen aktualisieren
			for r, label := range labels {
				label.SetText(fmt.Sprintf("%s: %d", r, inv.Resources[r]))
			}

			inv.Unlock()
		}
	}()

	w.SetContent(content)
	w.Resize(fyne.NewSize(700, 450))
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

	tryUpgrade := func(current interface{}) {
		inv.Lock()
		defer inv.Unlock()

		switch c := current.(type) {
		case Pickaxe:
			next, ok := pickaxeUpgrade[c]
			if !ok {
				dialog.ShowInformation("Spitzhacke ist max. Level!", "Deine Spitzhacke ist schon auf dem maximalen Level", w)
				return
			}
			for r, amt := range pickaxeResourceCost[c] {
				if inv.Resources[r] < amt {
					dialog.ShowInformation("Nicht genügend Ressourcen für die Spitzhacke!", "Dir fehlen Ressourcen für das Spitzhacke-Upgraden", w)
					return
				}
			}
			for r, amt := range pickaxeResourceCost[c] {
				inv.Resources[r] -= amt
			}
			inv.Pickaxe = next
			dialog.ShowInformation("", "Spitzhacke erfolgreich geupgradet", w)

		case Sword:
			next, ok := swordUpgrade[c]
			if !ok {
				dialog.ShowInformation("Schwert ist max. Level!", "Deine Schwert ist schon auf dem maximalen Level", w)
				return
			}
			for r, amt := range swordUpgradeCost[c] {
				if inv.Resources[r] < amt {
					dialog.ShowInformation("Nicht genügend Ressourcen für das Schwert!", "Dir fehlen Ressourcen für das Schwert upzugraden", w)
					return
				}
			}
			for r, amt := range swordUpgradeCost[c] {
				inv.Resources[r] -= amt
			}
			inv.Sword = next
			dialog.ShowInformation("", "Schwert erfolgreich geupgradet", w)

		case Armor:
			next, ok := armorUpgrade[c]
			if !ok {
				dialog.ShowInformation("Rüstung ist max. Level!", "Deine Rüstung ist schon auf dem maximalen Level", w)
				return
			}
			for r, amt := range armorUpgradeCost[c] {
				if inv.Resources[r] < amt {
					dialog.ShowInformation("Nicht genügend Ressourcen für Rüstung!", "Dir fehlen Ressourcen für das Rüstungs-Upgrade", w)
					return
				}
			}
			for r, amt := range armorUpgradeCost[c] {
				inv.Resources[r] -= amt
			}
			inv.Armor = next
			dialog.ShowInformation("", "Rüstung erfolgreich geupgradet", w)

		default:
			dialog.ShowInformation("Fehler", "Unbekannter Gegenstand zum Upgraden", w)
		}
		updateCosts() // Hier rufst du die Funktion zur Aktualisierung der Ressourcen auf
	}
	pickaxeBtn := widget.NewButton("Upgrade Spitzhacke", func() {
		tryUpgrade(inv.Pickaxe) // Nur die Funktion aufrufen, ohne msgLabel.SetText
	})

	swordBtn := widget.NewButton("Upgrade Schwert", func() {
		tryUpgrade(inv.Sword)
	})

	armorBtn := widget.NewButton("Upgrade Rüstung", func() {
		tryUpgrade(inv.Armor)
	})
	box := container.NewVBox(
		info, msgLabel,
		pickaxeCostLabel, pickaxeBtn,
		swordCostLabel, swordBtn,
		armorCostLabel, armorBtn,
	)
	w.SetContent(container.NewVScroll(box))
	updateCosts() // Initiale Ressourcen-Updates, falls erforderlich
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
	cristalleLabel := widget.NewLabel(fmt.Sprintf("Cristalle: %d", inv.Kristalle))
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
	for _, r := range []Resource{Kohle, Stone, Kupfererz, Eisenerz, Silvererz, Golderz, Platinerz, Rubinerz, Diamanterz, Mythrilerz, Blutkristallerz, Adamantiumerz, Kupferbarren, Eisenbarren, Stahlbarren, Silberbarren, Goldbarren, Platinbarren, Mythrilbarren, Adamantiumbarren, GeschliffenerRubin, GeschliffenerDiamant, GeschliffenerBlutkristall} {
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
			cristalleLabel.SetText(fmt.Sprintf("Cristalle: %d", inv.Kristalle))
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
					inv.Kristalle += 1
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
