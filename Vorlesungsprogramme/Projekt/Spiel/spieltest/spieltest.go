package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
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
	playerW         = 40
	playerH         = 60
	speed           = 20
	interactionDist = 80
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
	StoneSword:      {Stone: 20},
	KupferSword:     {Kupferbarren: 20},
	IronSword:       {Eisenbarren: 20},
	StahlSword:      {Stahlbarren: 20},
	GoldSword:       {Goldbarren: 20},
	PlatinumSword:   {Platinbarren: 20},
	DiamondSword:    {GeschliffenerDiamant: 20},
	MythrilSword:    {Mythrilbarren: 20},
	AdamantiumSword: {Adamantiumbarren: 100},
}

var armorUpgradeCost = map[Armor]map[Resource]int{
	WoodArmor:       {Stone: 20},
	StoneArmor:      {Stone: 20},
	KupferArmor:     {Kupferbarren: 20},
	IronArmor:       {Eisenbarren: 20},
	StahlArmor:      {Stahlbarren: 20},
	GoldArmor:       {Goldbarren: 20},
	PlatinumArmor:   {Platinbarren: 20},
	DiamondArmor:    {GeschliffenerDiamant: 20},
	MythrilArmor:    {Mythrilbarren: 20},
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

var mines = []*Mine{
	{"Kohlemine", fyne.NewPos(100, 200), Kohle, 0, 1, WoodPickaxe, false, 0, "bild/Kohleerz.png", nil, nil, nil},
	{"Steinmine", fyne.NewPos(300, 200), Stone, 100, 1, WoodPickaxe, false, 0, "bild/Steinerz.png", nil, nil, nil},
	{"Kupfermine", fyne.NewPos(500, 200), Kupfererz, 200, 1, StonePickaxe, false, 0, "bild/Kupfererz.png", nil, nil, nil},
	{"Eisenmine", fyne.NewPos(700, 200), Eisenerz, 250, 1, KupferPickaxe, false, 0, "bild/Eisenerz.png", nil, nil, nil},
	{"Silbernmine", fyne.NewPos(900, 200), Silvererz, 500, 1, IronPickaxe, false, 0, "bild/Silbererz.png", nil, nil, nil},
	{"Goldmine", fyne.NewPos(1100, 200), Golderz, 500, 1, IronPickaxe, false, 0, "bild/Golderz.png", nil, nil, nil},
	{"Platinmine", fyne.NewPos(100, 400), Platinerz, 1000, 1, GoldPickaxe, false, 0, "bild/Platinerz.png", nil, nil, nil},
	{"Rubinnmine", fyne.NewPos(300, 400), Rubinerz, 2000, 1, PlatinumPickaxe, false, 0, "bild/Rubinerz.png", nil, nil, nil},
	{"Diamantmine", fyne.NewPos(500, 400), Diamanterz, 3000, 1, PlatinumPickaxe, false, 0, "bild/Diamanterz.png", nil, nil, nil},
	{"Mythrilmine", fyne.NewPos(700, 400), Mythrilerz, 5000, 1, DiamondPickaxe, false, 0, "bild/Mythrilerz.png", nil, nil, nil},
	{"Blutkristallmine", fyne.NewPos(900, 400), Blutkristallerz, 10000, 1, DiamondPickaxe, false, 0, "bild/Blutkristallerz.png", nil, nil, nil},
	{"Adamantiummine", fyne.NewPos(1100, 400), Adamantiumerz, 100000, 1, MythrilPickax, false, 0, "bild/Adamantiumerz.png", nil, nil, nil},
}
var inv = &Inventory{
	Resources:      make(map[Resource]int),
	Pickaxe:        WoodPickaxe,
	Sword:          WoodSword,
	Armor:          WoodArmor,
	Wiedergeburten: 0,
	Money:          1000,
	Kristalle:      1000,
}

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
	ImageFile    string
	Icon         *canvas.Image
	Border       *canvas.Rectangle
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
	Resources      map[Resource]int
	Pickaxe        Pickaxe
	Sword          Sword
	Armor          Armor
	Money          int
	Kristalle      int
	Wiedergeburten int
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
	showStartWindow(a)
	a.Run()
}
func showStartWindow(a fyne.App) {
	w := a.NewWindow("Mining Game – Start")
	w.Resize(fyne.NewSize(windowWidth, windowHeight))
	w.CenterOnScreen()

	// Hintergrund
	background := canvas.NewImageFromFile("bild/StartBild.png")
	background.FillMode = canvas.ImageFillStretch
	background.Resize(fyne.NewSize(windowWidth, windowHeight))
	background.Move(fyne.NewPos(0, 0))

	startBtnPos := fyne.NewPos(650, 390)
	startBtn := widget.NewButton("Start", func() { w.Close(); showGameWindow(a) })
	startBtn.Move(startBtnPos)
	startBtn.Resize(fyne.NewSize(200, 50))

	objects := []fyne.CanvasObject{background, startBtn}
	w.SetContent(container.NewWithoutLayout(objects...))
	w.Show()
}

func showGameWindow(a fyne.App) {
	w := a.NewWindow("Mining Game")
	w.Resize(fyne.NewSize(windowWidth, windowHeight))
	w.CenterOnScreen()

	player := canvas.NewImageFromFile("bild/player_down.png")
	player.FillMode = canvas.ImageFillContain
	player.Resize(fyne.NewSize(playerW, playerH))
	x, y := float32(50), float32(500)
	player.Move(fyne.NewPos(x, y))

	// Spieler-Lebensanzeige
	playerHealth := 100
	healthBarWidth := float32(playerHealth) / 10
	healthBarHeight := float32(5)

	healthBar := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
	healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))

	background := canvas.NewImageFromFile("bild/WeltenBild.png")
	background.FillMode = canvas.ImageFillStretch
	background.Resize(fyne.NewSize(windowWidth, windowHeight))

	objects := []fyne.CanvasObject{background, healthBar, player}

	// Dungeon erstellen
	dungeonPos1 := fyne.NewPos(950, 500)
	dungeonPos2 := fyne.NewPos(1100, 550)
	dungeon := canvas.NewImageFromFile("bild/Dungeon.png")
	dungeon.Move(dungeonPos1)
	dungeon.Resize(fyne.NewSize(400, 200))
	dungeon.FillMode = canvas.ImageFillContain

	shopPos1 := fyne.NewPos(130, 350)
	shopPos2 := fyne.NewPos(180, 400)
	shop := canvas.NewImageFromFile("bild/Shop.png")
	shop.Move(shopPos1)
	shop.Resize(fyne.NewSize(100, 100))
	shop.FillMode = canvas.ImageFillContain

	smithPos1 := fyne.NewPos(300, 350)
	smithPos2 := fyne.NewPos(350, 400)
	smith := canvas.NewImageFromFile("bild/Smith.png")
	smith.Move(smithPos1)
	smith.Resize(fyne.NewSize(100, 100))
	smith.FillMode = canvas.ImageFillContain

	schmelzofenPos1 := fyne.NewPos(430, 350)
	schmelzofenPos2 := fyne.NewPos(400, 400)
	schmelzofen := canvas.NewImageFromFile("bild/Schmelzofen.png")
	schmelzofen.Move(schmelzofenPos1)
	schmelzofen.Resize(fyne.NewSize(100, 100))
	schmelzofen.FillMode = canvas.ImageFillContain

	kristallShopPos1 := fyne.NewPos(560, 350)
	kristallShopPos2 := fyne.NewPos(600, 400)
	kristall := canvas.NewImageFromFile("bild/Kristallshop.png")
	kristall.Move(kristallShopPos1)
	kristall.Resize(fyne.NewSize(100, 100))
	kristall.FillMode = canvas.ImageFillContain

	wiedergeburtenLadnePos1 := fyne.NewPos(800, 150)
	wiedergeburtenLadnePos2 := fyne.NewPos(850, 200)
	wiedergeburtenladen := canvas.NewImageFromFile("bild/Wiedergeburtenladen.png")
	wiedergeburtenladen.Move(wiedergeburtenLadnePos1)
	wiedergeburtenladen.Resize(fyne.NewSize(100, 100))
	wiedergeburtenladen.FillMode = canvas.ImageFillContain

	miningweltPos1 := fyne.NewPos(1300, 180)
	miningweltPos2 := fyne.NewPos(1350, 230)
	miningwelt := canvas.NewImageFromFile("bild/Miningwelt.png")
	miningwelt.Move(miningweltPos1)
	miningwelt.Resize(fyne.NewSize(100, 100))
	miningwelt.FillMode = canvas.ImageFillContain

	objects = append(objects, smith, kristall, schmelzofen, wiedergeburtenladen, miningwelt, shop, dungeon)
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

	playerSprites := map[string]string{
		"up":    "bild/LaufenRückwärts.png",
		"down":  "bild/LaufenForwärts.png",
		"left":  "bild/LaufenLinks.png",
		"right": "bild/LaufenRechts.png",
	}

	currentDir := "down"

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyUp:
			y -= speed
			currentDir = "up"

		case fyne.KeyDown:
			y += speed
			currentDir = "down"

		case fyne.KeyLeft:
			x -= speed
			currentDir = "left"

		case fyne.KeyRight:
			x += speed
			currentDir = "right"

		case fyne.KeyReturn:
			p := fyne.NewPos(x, y)
			if dist2(p, shopPos2) < interactionDist*interactionDist {
				openShop(a, inv)
				return
			}
			if dist2(p, smithPos2) < interactionDist*interactionDist {
				openSmith(a, inv)
				return
			}
			if dist2(p, kristallShopPos2) < interactionDist*interactionDist {
				openKristallShop(a, inv, mines)
				return
			}
			if dist2(p, schmelzofenPos2) < interactionDist*interactionDist {
				openSchmelzofen(a, inv)
				return
			}
			if dist2(p, wiedergeburtenLadnePos2) < interactionDist*interactionDist {
				openWiedergeburtenLaden(a, inv, mines)
				return
			}
			if dist2(p, dungeonPos2) < interactionDist*interactionDist {
				openDungeon(a, w, inv)
				return
			}
			if dist2(p, miningweltPos2) < interactionDist*interactionDist {
				w.Close()
				openMiningwelt1(a, inv, mines)
				return
			}

		case fyne.KeySpace:
			openInventory(a, inv)
		}

		x = clamp(x, 0, windowWidth-playerW)
		y = clamp(y, 0, windowHeight-playerH)
		player.Move(fyne.NewPos(x, y))
		player.File = playerSprites[currentDir]
		player.Refresh()

		// Health Bar über dem Spieler aktualisieren
		healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))
		healthBar.Refresh()
	})

	w.SetContent(container.NewWithoutLayout(objects...))
	w.Show()
}

/* =================Miningwelt ========================*/
func openMiningwelt1(a fyne.App, inv *Inventory, mines []*Mine) {
	w := a.NewWindow("Miningwelt 1")
	w.Resize(fyne.NewSize(1470, 765))
	w.CenterOnScreen()

	var objects []fyne.CanvasObject
	xPositions := []float32{350, 650, 950}
	y := float32(350)

	hochPos := fyne.NewPos(700, 20)
	hoch := canvas.NewImageFromFile("bild/LeiterHoch.png")
	hoch.Move(hochPos)
	hoch.Resize(fyne.NewSize(100, 100))
	hoch.FillMode = canvas.ImageFillContain

	runterPos := fyne.NewPos(700, 600)
	runter := canvas.NewImageFromFile("bild/LeiterRunter.png")
	runter.Move(runterPos)
	runter.Resize(fyne.NewSize(100, 100))
	runter.FillMode = canvas.ImageFillContain

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Steinboden.png")
	background.FillMode = canvas.ImageFillStretch
	background.Resize(fyne.NewSize(windowWidth, windowHeight))
	background.Move(fyne.NewPos(0, 0))

	objects = append(objects, background, hoch, runter)

	// Welt 1: erste 3 Minen (Index 0-2)
	for i := 0; i < 3; i++ {
		m := mines[i]
		m.Pos = fyne.NewPos(xPositions[i], y)

		icon := canvas.NewImageFromFile(m.ImageFile)
		icon.Resize(fyne.NewSize(80, 80))
		icon.Move(m.Pos)
		icon.FillMode = canvas.ImageFillContain

		label := canvas.NewText(m.Name, color.Black)
		label.TextSize = 12
		label.Alignment = fyne.TextAlignCenter
		label.Move(fyne.NewPos(xPositions[i], y+90))
		m.Label = label

		objects = append(objects, icon, label)
	}

	player := canvas.NewImageFromFile("bild/player_down.png")
	player.FillMode = canvas.ImageFillContain
	player.Resize(fyne.NewSize(50, 100))
	x, y := float32(50), float32(50)
	player.Move(fyne.NewPos(x, y))

	playerHealth := 100
	healthBarWidth := float32(playerHealth) / 10
	healthBarHeight := float32(5)
	healthBar := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
	healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))

	playerSprites := map[string]string{
		"up":    "bild/LaufenRückwärts.png",
		"down":  "bild/LaufenForwärts.png",
		"left":  "bild/LaufenLinks.png",
		"right": "bild/LaufenRechts.png",
	}

	currentDir := "down"

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyUp:
			y -= speed
			currentDir = "up"

		case fyne.KeyDown:
			y += speed
			currentDir = "down"

		case fyne.KeyLeft:
			x -= speed
			currentDir = "left"

		case fyne.KeyRight:
			x += speed
			currentDir = "right"

		case fyne.KeyReturn:
			p := fyne.NewPos(x, y)
			for j := 0; j < 3; j++ {
				if dist2(p, mines[j].Pos) < interactionDist*interactionDist {
					openMineShop(a, mines[j], inv)
					return
				}
			}
			if dist2(p, hochPos) < interactionDist*interactionDist {
				w.Close()
				showGameWindow(a)
				return
			}
			if dist2(p, runterPos) < interactionDist*interactionDist {
				w.Close()
				openMiningwelt2(a, inv, mines)
				return
			}
		case fyne.KeySpace:
			openInventory(a, inv)
		}
		x = clamp(x, 0, windowWidth-playerW)
		y = clamp(y, 0, windowHeight-playerH)
		player.Move(fyne.NewPos(x, y))
		player.File = playerSprites[currentDir]
		player.Refresh()

		healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))
		healthBar.Refresh()
	})

	objects = append(objects, player, healthBar)
	w.SetContent(container.NewWithoutLayout(objects...))
	w.Show()
}

func openMiningwelt2(a fyne.App, inv *Inventory, mines []*Mine) {
	w := a.NewWindow("Miningwelt 2")
	w.Resize(fyne.NewSize(1470, 765))
	w.CenterOnScreen()

	var objects []fyne.CanvasObject
	xPositions := []float32{350, 650, 950}
	y := float32(350)

	hochPos := fyne.NewPos(700, 20)
	hoch := canvas.NewImageFromFile("bild/LeiterHoch.png")
	hoch.Move(hochPos)
	hoch.Resize(fyne.NewSize(100, 100))
	hoch.FillMode = canvas.ImageFillContain

	runterPos := fyne.NewPos(700, 600)
	runter := canvas.NewImageFromFile("bild/LeiterRunter.png")
	runter.Move(runterPos)
	runter.Resize(fyne.NewSize(100, 100))
	runter.FillMode = canvas.ImageFillContain

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Steinboden.png")
	background.FillMode = canvas.ImageFillStretch
	background.Resize(fyne.NewSize(windowWidth, windowHeight))
	background.Move(fyne.NewPos(0, 0))

	objects = append(objects, background, hoch, runter)

	// Welt 2: Minen 3-5
	for i := 3; i < 6; i++ {
		m := mines[i]
		m.Pos = fyne.NewPos(xPositions[i-3], y)

		icon := canvas.NewImageFromFile(m.ImageFile)
		icon.Resize(fyne.NewSize(80, 80))
		icon.Move(m.Pos)
		icon.FillMode = canvas.ImageFillContain

		label := canvas.NewText(m.Name, color.Black)
		label.TextSize = 12
		label.Move(fyne.NewPos(xPositions[i-3], y+90))
		m.Label = label

		objects = append(objects, icon, label)
	}

	player := canvas.NewImageFromFile("bild/player_down.png")
	player.FillMode = canvas.ImageFillContain
	player.Resize(fyne.NewSize(50, 100))
	x, y := float32(50), float32(50)
	player.Move(fyne.NewPos(x, y))

	playerHealth := 100
	healthBarWidth := float32(playerHealth) / 10
	healthBarHeight := float32(5)
	healthBar := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
	healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))

	playerSprites := map[string]string{
		"up":    "bild/LaufenRückwärts.png",
		"down":  "bild/LaufenForwärts.png",
		"left":  "bild/LaufenLinks.png",
		"right": "bild/LaufenRechts.png",
	}

	currentDir := "down"

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyUp:
			y -= speed
			currentDir = "up"

		case fyne.KeyDown:
			y += speed
			currentDir = "down"

		case fyne.KeyLeft:
			x -= speed
			currentDir = "left"

		case fyne.KeyRight:
			x += speed
			currentDir = "right"
		case fyne.KeyReturn:
			p := fyne.NewPos(x, y)
			for j := 3; j < 6; j++ {
				if dist2(p, mines[j].Pos) < interactionDist*interactionDist {
					openMineShop(a, mines[j], inv)
					return
				}
			}
			if dist2(p, hochPos) < interactionDist*interactionDist {
				w.Close()
				openMiningwelt1(a, inv, mines)
				return
			}
			if dist2(p, runterPos) < interactionDist*interactionDist {
				w.Close()
				openMiningwelt3(a, inv, mines)
				return
			}
		case fyne.KeySpace:
			openInventory(a, inv)
		}
		x = clamp(x, 0, windowWidth-playerW)
		y = clamp(y, 0, windowHeight-playerH)
		player.Move(fyne.NewPos(x, y))
		player.File = playerSprites[currentDir]
		player.Refresh()

		healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))
		healthBar.Refresh()
	})
	objects = append(objects, player, healthBar)
	w.SetContent(container.NewWithoutLayout(objects...))
	w.Show()
}

func openMiningwelt3(a fyne.App, inv *Inventory, mines []*Mine) {
	w := a.NewWindow("Miningwelt 3")
	w.Resize(fyne.NewSize(1470, 765))
	w.CenterOnScreen()

	var objects []fyne.CanvasObject
	xPositions := []float32{350, 650, 950}
	y := float32(350)

	hochPos := fyne.NewPos(700, 20)
	hoch := canvas.NewImageFromFile("bild/LeiterHoch.png")
	hoch.Move(hochPos)
	hoch.Resize(fyne.NewSize(100, 100))
	hoch.FillMode = canvas.ImageFillContain

	runterPos := fyne.NewPos(700, 600)
	runter := canvas.NewImageFromFile("bild/LeiterRunter.png")
	runter.Move(runterPos)
	runter.Resize(fyne.NewSize(100, 100))
	runter.FillMode = canvas.ImageFillContain

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Steinboden.png")
	background.FillMode = canvas.ImageFillStretch
	background.Resize(fyne.NewSize(windowWidth, windowHeight))
	background.Move(fyne.NewPos(0, 0))

	objects = append(objects, background, hoch, runter)

	// Welt 3: Minen 6-8
	for i := 6; i < 9; i++ {
		m := mines[i]
		m.Pos = fyne.NewPos(xPositions[i-6], y)

		icon := canvas.NewImageFromFile(m.ImageFile)
		icon.Resize(fyne.NewSize(80, 80))
		icon.Move(m.Pos)
		icon.FillMode = canvas.ImageFillContain

		label := canvas.NewText(m.Name, color.Black)
		label.TextSize = 12
		label.Move(fyne.NewPos(xPositions[i-6], y+90))
		m.Label = label

		objects = append(objects, icon, label)
	}

	player := canvas.NewImageFromFile("bild/player_down.png")
	player.FillMode = canvas.ImageFillContain
	player.Resize(fyne.NewSize(50, 100))
	x, y := float32(50), float32(50)
	player.Move(fyne.NewPos(x, y))

	playerHealth := 100
	healthBarWidth := float32(playerHealth) / 10
	healthBarHeight := float32(5)
	healthBar := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
	healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))

	playerSprites := map[string]string{
		"up":    "bild/LaufenRückwärts.png",
		"down":  "bild/LaufenForwärts.png",
		"left":  "bild/LaufenLinks.png",
		"right": "bild/LaufenRechts.png",
	}

	currentDir := "down"

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyUp:
			y -= speed
			currentDir = "up"

		case fyne.KeyDown:
			y += speed
			currentDir = "down"

		case fyne.KeyLeft:
			x -= speed
			currentDir = "left"

		case fyne.KeyRight:
			x += speed
			currentDir = "right"
		case fyne.KeyReturn:
			p := fyne.NewPos(x, y)
			for j := 6; j < 9; j++ {
				if dist2(p, mines[j].Pos) < interactionDist*interactionDist {
					openMineShop(a, mines[j], inv)
					return
				}
			}
			if dist2(p, hochPos) < interactionDist*interactionDist {
				w.Close()
				openMiningwelt2(a, inv, mines)
				return
			}
			if dist2(p, runterPos) < interactionDist*interactionDist {
				w.Close()
				openMiningwelt4(a, inv, mines)
				return
			}
		case fyne.KeySpace:
			openInventory(a, inv)
		}
		x = clamp(x, 0, windowWidth-playerW)
		y = clamp(y, 0, windowHeight-playerH)
		player.Move(fyne.NewPos(x, y))
		player.File = playerSprites[currentDir]
		player.Refresh()

		healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))
		healthBar.Refresh()
	})

	objects = append(objects, player, healthBar)
	w.SetContent(container.NewWithoutLayout(objects...))
	w.Show()

}

func openMiningwelt4(a fyne.App, inv *Inventory, mines []*Mine) {
	w := a.NewWindow("Miningwelt 4")
	w.Resize(fyne.NewSize(1470, 765))
	w.CenterOnScreen()

	var objects []fyne.CanvasObject
	xPositions := []float32{350, 650, 950}
	y := float32(350)

	hochPos := fyne.NewPos(700, 20)
	hoch := canvas.NewImageFromFile("bild/LeiterHoch.png")
	hoch.Move(hochPos)
	hoch.Resize(fyne.NewSize(100, 100))
	hoch.FillMode = canvas.ImageFillContain

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Steinboden.png")
	background.FillMode = canvas.ImageFillStretch
	background.Resize(fyne.NewSize(windowWidth, windowHeight))
	background.Move(fyne.NewPos(0, 0))

	objects = append(objects, background, hoch)
	// Welt 4: Minen 9-11
	for i := 9; i < 12 && i < len(mines); i++ {
		m := mines[i]
		m.Pos = fyne.NewPos(xPositions[i-9], y)

		icon := canvas.NewImageFromFile(m.ImageFile)
		icon.Resize(fyne.NewSize(80, 80))
		icon.Move(m.Pos)
		icon.FillMode = canvas.ImageFillContain

		label := canvas.NewText(m.Name, color.Black)
		label.TextSize = 12
		label.Move(fyne.NewPos(xPositions[i-9], y+90))
		m.Label = label

		objects = append(objects, icon, label)
	}

	player := canvas.NewImageFromFile("bild/player_down.png")
	player.FillMode = canvas.ImageFillContain
	player.Resize(fyne.NewSize(50, 100))
	x, y := float32(50), float32(50)
	player.Move(fyne.NewPos(x, y))

	playerHealth := 100
	healthBarWidth := float32(playerHealth) / 10
	healthBarHeight := float32(5)
	healthBar := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
	healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))

	playerSprites := map[string]string{
		"up":    "bild/LaufenRückwärts.png",
		"down":  "bild/LaufenForwärts.png",
		"left":  "bild/LaufenLinks.png",
		"right": "bild/LaufenRechts.png",
	}

	currentDir := "down"

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyUp:
			y -= speed
			currentDir = "up"

		case fyne.KeyDown:
			y += speed
			currentDir = "down"

		case fyne.KeyLeft:
			x -= speed
			currentDir = "left"

		case fyne.KeyRight:
			x += speed
			currentDir = "right"
		case fyne.KeyReturn:
			p := fyne.NewPos(x, y)
			for j := 9; j < 12 && j < len(mines); j++ {
				if dist2(p, mines[j].Pos) < interactionDist*interactionDist {
					openMineShop(a, mines[j], inv)
					return
				}
			}
			if dist2(p, hochPos) < interactionDist*interactionDist {
				w.Close()
				openMiningwelt3(a, inv, mines)
				return
			}
		case fyne.KeySpace:
			openInventory(a, inv)
		}
		x = clamp(x, 0, windowWidth-playerW)
		y = clamp(y, 0, windowHeight-playerH)
		player.Move(fyne.NewPos(x, y))
		player.File = playerSprites[currentDir]
		player.Refresh()

		healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))
		healthBar.Refresh()
	})
	objects = append(objects, player, healthBar)
	w.SetContent(container.NewWithoutLayout(objects...))
	w.Show()
}

/* ==================Wiedergeburten ==================*/
func openWiedergeburtenLaden(a fyne.App, inv *Inventory, mines []*Mine) {
	w := a.NewWindow("Wiedergeburten")
	enableBackspaceClose(w)
	w.CenterOnScreen()

	// ⚡ Kristallanzeige in Weiß
	kristallLabel := canvas.NewText(fmt.Sprintf("Kristalle: %d", inv.Kristalle), color.Black)
	kristallLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Skill-Level für Wiedergeburten
	currentLevel := inv.Wiedergeburten
	maxLevel := 5
	costs := map[int]int{
		1: 100,
		2: 200,
		3: 400,
		4: 600,
		5: 1000,
	}

	// ⚡ Skill-Label in Weiß
	skillLabel := canvas.NewText(fmt.Sprintf("Wiedergeburten-Skill Level %d", currentLevel), color.Black)

	// Button zum Leveln
	btn := widget.NewButton("", nil)
	updateButton := func() {
		if currentLevel >= maxLevel {
			btn.SetText("Maximales Level erreicht")
			btn.Disable()
		} else {
			btn.SetText(fmt.Sprintf("Wiedergeboren %d (Kosten: %d Kristalle)", currentLevel+1, costs[currentLevel+1]))
			btn.Enable()
		}
		skillLabel.Text = fmt.Sprintf("Wiedergeburt %d", currentLevel)
		canvas.Refresh(skillLabel)
	}
	updateButton()

	btn.OnTapped = func() {
		if currentLevel >= maxLevel {
			dialog.ShowInformation("Maximale Wiedergeburt", "Du hast bereits die höchste Wiedergeburt erreicht.", w)
			return
		}

		inv.Lock()
		defer inv.Unlock()

		cost := costs[currentLevel+1]
		if inv.Kristalle >= cost {

			currentLevel++
			inv.Wiedergeburten = currentLevel

			inv.Kristalle = 0
			inv.Money = 0
			inv.Armor = "Holzrüstung"
			inv.Pickaxe = "Holzspitzhacke"
			inv.Sword = "Holzschwert"
			inv.Resources = map[Resource]int{
				Kohle:                     0,
				Stone:                     0,
				Kupfererz:                 0,
				Kupferbarren:              0,
				Eisenerz:                  0,
				Eisenbarren:               0,
				Stahlbarren:               0,
				Silvererz:                 0,
				Silberbarren:              0,
				Golderz:                   0,
				Goldbarren:                0,
				Platinerz:                 0,
				Platinbarren:              0,
				Rubinerz:                  0,
				GeschliffenerRubin:        0,
				Diamanterz:                0,
				GeschliffenerDiamant:      0,
				Mythrilerz:                0,
				Mythrilbarren:             0,
				Blutkristallerz:           0,
				GeschliffenerBlutkristall: 0,
				Adamantiumerz:             0,
				Adamantiumbarren:          0,
			}
			for i := range mines {
				mines[i].Owned = false
				mines[i].Rate = mines[i].Rate * (inv.Wiedergeburten + 1)
			}

			kristallLabel.Text = fmt.Sprintf("Kristalle: %d", inv.Kristalle)
			canvas.Refresh(kristallLabel)
			updateButton()
			dialog.ShowInformation("Erfolg", fmt.Sprintf("Wiedergeburt %d freigeschaltet!", currentLevel), w)
		} else {
			dialog.ShowInformation("Nicht genug Kristalle", "Du hast nicht genügend Kristalle, um wiedergeboren zu werden.", w)
		}
	}

	// Layout
	mainBox := container.NewVBox(
		kristallLabel,
		widget.NewSeparator(),
		skillLabel,
		btn,
	)

	// Box zentrieren
	centeredBox := container.NewCenter(mainBox)

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Wiedergeburtladeninnen.png")
	background.FillMode = canvas.ImageFillStretch
	background.Resize(fyne.NewSize(windowWidth, windowHeight))
	background.Move(fyne.NewPos(0, 0))

	w.SetContent(container.NewMax(
		background,
		centeredBox,
	))

	w.Resize(fyne.NewSize(400, 200))
	w.Show()
}

/* ===================== MINE SHOP ===================== */
func openMineShop(a fyne.App, m *Mine, inv *Inventory) {
	w := a.NewWindow(m.Name)
	enableBackspaceClose(w)
	w.Resize(fyne.NewSize(400, 350))
	w.CenterOnScreen()

	// ⚡ Container für alle Infos
	infoBox := container.NewVBox()

	// Label als canvas.Text für schwarze Schrift
	nameLabel := canvas.NewText(m.Name, color.White)
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}
	rateLabel := canvas.NewText("", color.White)
	levelLabel := canvas.NewText("", color.White)
	reqLabel := canvas.NewText("", color.White)
	costLabel := canvas.NewText("", color.White)

	// Alle Labels in VBox
	infoBox.Add(nameLabel)
	infoBox.Add(rateLabel)
	infoBox.Add(levelLabel)
	infoBox.Add(reqLabel)
	infoBox.Add(costLabel)

	// Button
	actionBtn := widget.NewButton("", nil)
	centeredBox := container.NewCenter(actionBtn)

	// Update-Funktion
	updateLabels := func() {
		rateLabel.Text = fmt.Sprintf("Rate: %d/sec", m.Rate)
		levelLabel.Text = fmt.Sprintf("Upgrade-Level: %d/%d", m.UpgradeLevel, MaxMineUpgrade)
		reqLabel.Text = fmt.Sprintf("Benötigte Spitzhacke: %s", m.Required)

		if m.Owned {
			money, res, ok := getMineUpgradeCost(m)
			if ok {
				costText := fmt.Sprintf("Nächste Upgrade-Kosten:\nGeld: %d", money)
				for r, amt := range res {
					costText += fmt.Sprintf("\n%s: %d", r, amt)
				}
				costLabel.Text = costText
			} else {
				costLabel.Text = "(Max Level erreicht)"
			}
		} else {
			costLabel.Text = fmt.Sprintf("Kaufpreis: %d Geld", m.Cost)
		}

		// Refresh für canvas.Text
		canvas.Refresh(nameLabel)
		canvas.Refresh(rateLabel)
		canvas.Refresh(levelLabel)
		canvas.Refresh(reqLabel)
		canvas.Refresh(costLabel)
	}

	// Button-Logik
	actionBtn.OnTapped = func() {
		inv.Lock()
		defer inv.Unlock()

		pickaxeLevels := map[Pickaxe]int{
			WoodPickaxe: 1, StonePickaxe: 2, KupferPickaxe: 3, IronPickaxe: 4,
			StahlPickaxe: 5, GoldPickaxe: 6, PlatinumPickaxe: 7, DiamondPickaxe: 8,
			MythrilPickax: 9, AdamantiumPickax: 10,
		}

		if pickaxeLevels[inv.Pickaxe] < pickaxeLevels[m.Required] {
			costLabel.Text += "\nDu benötigst eine bessere Spitzhacke!"
			canvas.Refresh(costLabel)
			return
		}

		if !m.Owned {
			if inv.Money >= m.Cost {
				inv.Money -= m.Cost
				m.Owned = true
				updateLabels()
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
		updateLabels()

		if m.UpgradeLevel >= MaxMineUpgrade {
			actionBtn.Disable()
		}
	}

	if m.Owned {
		actionBtn.SetText("Upgrade Mine")
	} else {
		actionBtn.SetText(fmt.Sprintf("Kaufen für (%d Geld)", m.Cost))
	}

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Miningshopinnen.png")
	background.FillMode = canvas.ImageFillStretch

	w.SetContent(container.NewMax(
		background,
		container.NewVBox(
			infoBox,
			centeredBox,
		),
	))

	updateLabels()
	w.Show()
}

/*====================== Cristal Shop ==================*/
var currentLevels = make(map[string]int)

func openKristallShop(a fyne.App, inv *Inventory, mines []*Mine) {
	// Feste Reihenfolge für die Skills
	skillOrder := []string{
		"Kohlemine Multiplier",
		"Steinminen Multiplier",
		"Kupfermine Multiplier",
		"Eisenminen Multiplier",
		"Silberminen Multiplier",
		"Goldminen Multiplier",
		"Platinmine Multiplier",
		"Rubinmine Multiplier",
		"Diamantmine Multiplier",
		"Mythrilmine Multiplier",
		"Blutkristallmine Multiplier",
		"Adamantiummine Multiplier",
	}

	// Preis für jede Fähigkeit (in Kristallen)
	skills := map[string]map[int]int{
		"Kohlemine Multiplier":        {1: 10, 2: 25, 3: 50},
		"Steinminen Multiplier":       {1: 10, 2: 25, 3: 50},
		"Kupfermine Multiplier":       {1: 10, 2: 25, 3: 50},
		"Eisenminen Multiplier":       {1: 15, 2: 35, 3: 70},
		"Silberminen Multiplier":      {1: 10, 2: 25, 3: 50},
		"Goldminen Multiplier":        {1: 20, 2: 45, 3: 90},
		"Platinmine Multiplier":       {1: 10, 2: 25, 3: 50},
		"Rubinmine Multiplier":        {1: 10, 2: 25, 3: 50},
		"Diamantmine Multiplier":      {1: 15, 2: 35, 3: 70},
		"Mythrilmine Multiplier":      {1: 10, 2: 25, 3: 50},
		"Blutkristallmine Multiplier": {1: 10, 2: 25, 3: 50},
		"Adamantiummine Multiplier":   {1: 10, 2: 25, 3: 50},
	}

	// Level-Up-Funktionen für Skills
	levelUpFunctions := map[string]map[int]func(mines []*Mine){
		"Kohlemine Multiplier": {
			1: func(mines []*Mine) { mines[0].Rate = int(float64(mines[0].Rate) * 1.2) },
			2: func(mines []*Mine) { mines[0].Rate = int(float64(mines[0].Rate) * 1.5) },
			3: func(mines []*Mine) { mines[0].Rate = int(float64(mines[0].Rate) * 2.0) },
		},
		// ... gleiche Struktur für alle anderen Skills ...
	}

	// Fenster erstellen
	w := a.NewWindow("Skill Tree Shop")
	enableBackspaceClose(w)
	w.CenterOnScreen()

	// ⚡ Kristallanzeige in Weiß & fett
	kristallLabel := canvas.NewText(fmt.Sprintf("Kristalle: %d", inv.Kristalle), color.White)
	kristallLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Hauptbox für das Fenster
	mainBox := container.NewVBox(
		kristallLabel,
		widget.NewSeparator(),
		func() *canvas.Text {
			t := canvas.NewText("Fähigkeiten freischalten", color.White)
			t.TextStyle = fyne.TextStyle{Bold: true}
			return t
		}(),
	)

	// VBox für die Skills
	skillsBox := container.NewVBox()

	// Buttons für alle Fähigkeiten in fester Reihenfolge erstellen
	for _, skill := range skillOrder {
		levels := skills[skill]
		currentLevel := currentLevels[skill]
		skillBox := createSkillButtonWhite(skill, levels, currentLevel, inv, kristallLabel, levelUpFunctions, w)
		skillsBox.Add(skillBox)
	}

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Kristallshopinnen.png")
	background.FillMode = canvas.ImageFillStretch

	// Scrollbarer Inhalt
	scroll := container.NewVScroll(skillsBox)
	scroll.SetMinSize(fyne.NewSize(380, 500)) // Höhe und Breite des Scrollbereichs
	mainBox.Add(scroll)

	// Alles übereinander legen
	content := container.NewMax(
		background, // unten
		scroll,     // oben
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 600))
	w.Show()
}

// Hilfsfunktion für Skill-Buttons mit weißen, fetten Labels
func createSkillButtonWhite(skill string, levels map[int]int, currentLevel int, inv *Inventory, kristallLabel *canvas.Text, levelUpFunctions map[string]map[int]func(mines []*Mine), w fyne.Window) *fyne.Container {
	skillBox := container.NewVBox()
	// ⚡ Skill-Label in Weiß & fett
	skillLabel := canvas.NewText(skill, color.White)
	skillLabel.TextStyle = fyne.TextStyle{Bold: true}
	skillBox.Add(skillLabel)

	// Button erstellen
	btn := widget.NewButton(fmt.Sprintf("%s Level %d (Kosten: %d Kristalle)", skill, currentLevel+1, levels[currentLevel+1]), nil)

	// Button-Event
	btn.OnTapped = func() {
		handleSkillUnlockWhite(skill, levels, inv, kristallLabel, levelUpFunctions, btn, skillLabel, w)
	}

	skillBox.Add(btn)
	return skillBox
}

// Skill-Freischaltung für weiße Labels & fett
func handleSkillUnlockWhite(skill string, levels map[int]int, inv *Inventory, kristallLabel *canvas.Text, levelUpFunctions map[string]map[int]func(mines []*Mine), btn *widget.Button, skillLabel *canvas.Text, w fyne.Window) {
	inv.Lock()
	defer inv.Unlock()

	currentLevel := currentLevels[skill]

	if currentLevel < len(levels) && inv.Kristalle >= levels[currentLevel+1] {
		inv.Kristalle -= levels[currentLevel+1]
		kristallLabel.Text = fmt.Sprintf("Kristalle: %d", inv.Kristalle)
		canvas.Refresh(kristallLabel)

		currentLevel++
		currentLevels[skill] = currentLevel
		btn.SetText(fmt.Sprintf("%s Level %d (Kosten: %d Kristalle)", skill, currentLevel+1, levels[currentLevel+1]))
		skillLabel.Text = fmt.Sprintf("%s (Level %d freigeschaltet)", skill, currentLevel)
		canvas.Refresh(skillLabel)

		if fnMap, ok := levelUpFunctions[skill]; ok {
			if fn, ok := fnMap[currentLevel]; ok && fn != nil {
				fn(mines)
			}
		}
	} else if currentLevel < len(levels) {
		dialog.ShowInformation("Nicht genug Kristalle", "Du hast nicht genügend Kristalle, um dieses Level freizuschalten.", w)
	} else {
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
		Kohle, Stone, Kupfererz, Eisenerz, Silvererz, Golderz,
		Platinerz, Rubinerz, Diamanterz, Mythrilerz,
		Blutkristallerz, Adamantiumerz,
	}

	w := a.NewWindow("Shop")
	enableBackspaceClose(w)
	w.CenterOnScreen()

	// Money Label
	moneyText := canvas.NewText("", color.White)
	moneyText.TextStyle = fyne.TextStyle{Bold: true}

	labels := make(map[Resource]*canvas.Text)

	// Verkaufsmenge
	selectedAmount := 1
	amountSelector := widget.NewRadioGroup([]string{"1", "10", "100"}, func(value string) {
		switch value {
		case "10":
			selectedAmount = 10
		case "100":
			selectedAmount = 100
		default:
			selectedAmount = 1
		}
	})
	amountSelector.SetSelected("1")

	// Weißer Overlay-Text
	selectedText := canvas.NewText(amountSelector.Selected, color.White)
	selectedText.TextStyle = fyne.TextStyle{Bold: true}

	// Verkaufsblock pro Resource
	createSellBlock := func(r Resource) *fyne.Container {
		box := container.NewVBox()

		label := canvas.NewText("", color.White)
		label.TextStyle = fyne.TextStyle{}
		labels[r] = label

		price := prices[r]
		btn := widget.NewButton("Verkaufen", func() {
			inv.Lock()
			defer inv.Unlock()

			sell := selectedAmount
			if inv.Resources[r] < sell {
				sell = inv.Resources[r]
			}
			if sell > 0 {
				inv.Resources[r] -= sell
				inv.Money += sell * price
			}

			label.Text = fmt.Sprintf("%s: %d", r, inv.Resources[r])
			canvas.Refresh(label)

			moneyText.Text = fmt.Sprintf("Geld: %d", inv.Money)
			canvas.Refresh(moneyText)
		})

		box.Add(label) // nur das Label, kein Hintergrund
		box.Add(btn)
		box.Add(widget.NewSeparator())
		return box
	}

	// Überschriften (ohne Hintergrund)
	leftBox := container.NewVBox(canvas.NewText("Geschliffene Edelsteine", color.White))
	midBox := container.NewVBox(canvas.NewText("Barren", color.White))
	rightBox := container.NewVBox(canvas.NewText("Erze & Rohstoffe", color.White))

	for _, r := range edelsteine {
		leftBox.Add(createSellBlock(r))
	}
	for _, r := range barren {
		midBox.Add(createSellBlock(r))
	}
	for _, r := range erze {
		rightBox.Add(createSellBlock(r))
	}

	// Menge-Anzeige jetzt als canvas.Text in Weiß
	amountLabel := canvas.NewText("Menge:", color.White)

	topBar := container.NewHBox(
		moneyText,
		widget.NewSeparator(),
		amountLabel,
		amountSelector,
	)

	content := container.NewBorder(
		topBar,
		nil, nil, nil,
		container.NewGridWithColumns(
			3,
			container.NewVScroll(rightBox),
			container.NewVScroll(midBox),
			container.NewVScroll(leftBox),
		),
	)

	// Aktualisierungsticker
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			inv.Lock()
			moneyText.Text = fmt.Sprintf("Geld: %d", inv.Money)
			for r, label := range labels {
				label.Text = fmt.Sprintf("%s: %d", r, inv.Resources[r])
			}
			inv.Unlock()
			canvas.Refresh(moneyText)
			for _, label := range labels {
				canvas.Refresh(label)
			}
		}
	}()

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Shopinnen.png")
	background.FillMode = canvas.ImageFillStretch

	w.SetContent(container.NewMax(
		background,
		content,
	))

	w.Resize(fyne.NewSize(900, 450))
	w.Show()
}

/* ===================Schmelz ofen =================== */
func openSchmelzofen(a fyne.App, inv *Inventory) {
	w := a.NewWindow("Schmelzofen")
	enableBackspaceClose(w)
	w.CenterOnScreen()

	// Kohle-Kosten pro Einheit
	KohleKosten := map[Resource]int{
		Kupfererz: 1, Eisenerz: 1, Eisenbarren: 2, Silvererz: 2,
		Golderz: 3, Platinerz: 4, Mythrilerz: 6, Adamantiumerz: 10,
	}

	// Erz -> Barren
	OreSchmelzen := map[Resource]Resource{
		Kupfererz: Kupferbarren, Eisenerz: Eisenbarren, Eisenbarren: Stahlbarren,
		Silvererz: Silberbarren, Golderz: Goldbarren, Platinerz: Platinbarren,
		Mythrilerz: Mythrilbarren, Adamantiumerz: Adamantiumbarren,
	}

	// Kristall -> geschliffen
	Kristalleschleifen := map[Resource]Resource{
		Rubinerz: GeschliffenerRubin, Diamanterz: GeschliffenerDiamant, Blutkristallerz: GeschliffenerBlutkristall,
	}

	resourceOrderMetall := []Resource{Kupfererz, Eisenerz, Eisenbarren, Silvererz, Golderz, Platinerz, Mythrilerz, Adamantiumerz}
	resourceOrderKristall := []Resource{Rubinerz, Diamanterz, Blutkristallerz}

	labels := make(map[Resource]*canvas.Text)
	sellAmounts := []int{1, 10, 100}

	// 🔥 Brennstoff
	coalText := canvas.NewText("", color.White)
	coalText.TextStyle = fyne.TextStyle{Bold: true}

	// 🔥 Metall-Spalte
	titleFuel := canvas.NewText("Brennstoff", color.White)
	titleFuel.TextStyle = fyne.TextStyle{Bold: true}

	titleMetall := canvas.NewText("Metalle schmelzen", color.White)
	titleMetall.TextStyle = fyne.TextStyle{Bold: true}

	metallBox := container.NewVBox(
		titleFuel,
		coalText,
		titleMetall,
	)

	// 💎 Kristall-Spalte
	titleKristall := canvas.NewText("Kristalle schleifen", color.White)
	titleKristall.TextStyle = fyne.TextStyle{Bold: true}

	kristallBox := container.NewVBox(titleKristall)

	// Block-Erzeugung
	createResourceBlock := func(input Resource, output Resource, brauchtKohle bool) *fyne.Container {
		resBox := container.NewVBox()

		// Titel
		titleText := canvas.NewText(fmt.Sprintf("%s ➜ %s", input, output), color.White)
		titleText.TextStyle = fyne.TextStyle{Bold: true}
		resBox.Add(titleText)

		// Label
		label := canvas.NewText("", color.White)
		labels[input] = label
		resBox.Add(label)

		// Kohle-Hinweis
		if brauchtKohle {
			coalLabel := canvas.NewText(fmt.Sprintf("Benötigt Kohle pro Einheit: %d", KohleKosten[input]), color.White)
			resBox.Add(coalLabel)
		}

		// Buttons
		for _, amt := range sellAmounts {
			amount := amt
			btn := widget.NewButton(fmt.Sprintf("%d verarbeiten", amount), func() {
				inv.Lock()
				defer inv.Unlock()

				if inv.Resources[input] <= 0 {
					dialog.ShowInformation("Zu wenig Material", fmt.Sprintf("Du hast kein %s mehr.", input), w)
					return
				}

				use := amount
				if inv.Resources[input] < use {
					use = inv.Resources[input]
				}

				if brauchtKohle {
					maxByCoal := inv.Resources[Kohle] / KohleKosten[input]
					if maxByCoal <= 0 {
						dialog.ShowInformation("Zu wenig Kohle", fmt.Sprintf("Du brauchst %d Kohle pro Einheit, hast aber nur %d.", KohleKosten[input], inv.Resources[Kohle]), w)
						return
					}
					if maxByCoal < use {
						use = maxByCoal
					}
				}

				if use <= 0 {
					return
				}

				inv.Resources[input] -= use
				inv.Resources[output] += use
				if brauchtKohle {
					inv.Resources[Kohle] -= use * KohleKosten[input]
				}

				label.Text = fmt.Sprintf("%s: %d", input, inv.Resources[input])
				canvas.Refresh(label)
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
	content := container.NewGridWithColumns(2, container.NewVScroll(metallBox), container.NewVScroll(kristallBox))

	// UI Refresh
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			inv.Lock()
			coalText.Text = fmt.Sprintf("Kohle: %d", inv.Resources[Kohle])
			canvas.Refresh(coalText)
			for r, label := range labels {
				label.Text = fmt.Sprintf("%s: %d", r, inv.Resources[r])
				canvas.Refresh(label)
			}
			inv.Unlock()
		}
	}()

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Schmelzofeninnen.png")
	background.FillMode = canvas.ImageFillStretch

	// Content + Hintergrund kombinieren
	w.SetContent(container.NewMax(background, content))
	w.Resize(fyne.NewSize(700, 450))
	w.Show()
}

/*===================== SCHMIED ===================== */
func openSmith(a fyne.App, inv *Inventory) {
	w := a.NewWindow("Schmiede")
	w.Resize(fyne.NewSize(400, 350))
	w.CenterOnScreen()

	content := container.NewVBox()

	// Statusanzeige: canvas.Text für weiße Schrift
	moneyLabel := canvas.NewText("", color.White)
	moneyLabel.TextStyle = fyne.TextStyle{Bold: true}
	pickaxeLabel := canvas.NewText("", color.White)
	swordLabel := canvas.NewText("", color.White)
	armorLabel := canvas.NewText("", color.White)

	updateStatus := func() {
		inv.Lock()
		m := inv.Money
		p := inv.Pickaxe
		s := inv.Sword
		a := inv.Armor
		inv.Unlock()

		moneyLabel.Text = fmt.Sprintf("Geld: %d", m)
		pickaxeLabel.Text = fmt.Sprintf("Pickaxe: %s", p)
		swordLabel.Text = fmt.Sprintf("Sword: %s", s)
		armorLabel.Text = fmt.Sprintf("Armor: %s", a)

		// Refresh für Canvas-Text
		canvas.Refresh(moneyLabel)
		canvas.Refresh(pickaxeLabel)
		canvas.Refresh(swordLabel)
		canvas.Refresh(armorLabel)
	}

	statusContainer := container.NewVBox(
		moneyLabel,
		pickaxeLabel,
		swordLabel,
		armorLabel,
	)
	content.Add(statusContainer)

	separator := canvas.NewRectangle(color.Gray{Y: 150})
	separator.SetMinSize(fyne.NewSize(400, 2))
	content.Add(separator)

	// Upgrade-Button Helper (unverändert)
	addUpgradeButton := func(name string, currentItem interface{}) {
		btn := widget.NewButton("", nil)

		updateButtonText := func() {
			inv.Lock()
			var fromLevel, toLevel string
			var costText string
			canUpgrade := false

			switch ci := currentItem.(type) {
			case *Pickaxe:
				fromLevel = string(*ci)
				next, ok := pickaxeUpgrade[*ci]
				if !ok {
					inv.Unlock()
					btn.SetText(fmt.Sprintf("%s kann nicht weiter verbessert werden.", name))
					btn.Disable()
					return
				}
				toLevel = string(next)
				cost := pickaxeResourceCost[next]
				costStr := ""
				hasAll := true
				for r, v := range cost {
					costStr += fmt.Sprintf("%s: %d ", r, v)
					if inv.Resources[r] < v {
						hasAll = false
					}
				}
				costText = costStr
				canUpgrade = hasAll
			case *Sword:
				fromLevel = string(*ci)
				next, ok := swordUpgrade[*ci]
				if !ok {
					inv.Unlock()
					btn.SetText(fmt.Sprintf("%s kann nicht weiter verbessert werden.", name))
					btn.Disable()
					return
				}
				toLevel = string(next)
				cost := swordUpgradeCost[next]
				costStr := ""
				hasAll := true
				for r, v := range cost {
					costStr += fmt.Sprintf("%s: %d ", r, v)
					if inv.Resources[r] < v {
						hasAll = false
					}
				}
				costText = costStr
				canUpgrade = hasAll
			case *Armor:
				fromLevel = string(*ci)
				next, ok := armorUpgrade[*ci]
				if !ok {
					inv.Unlock()
					btn.SetText(fmt.Sprintf("%s kann nicht weiter verbessert werden.", name))
					btn.Disable()
					return
				}
				toLevel = string(next)
				cost := armorUpgradeCost[next]
				costStr := ""
				hasAll := true
				for r, v := range cost {
					costStr += fmt.Sprintf("%s: %d ", r, v)
					if inv.Resources[r] < v {
						hasAll = false
					}
				}
				costText = costStr
				canUpgrade = hasAll
			}
			inv.Unlock()

			btn.SetText(fmt.Sprintf("%s Upgrade: Von %s => %s | Kosten: %s", name, fromLevel, toLevel, costText))
			if canUpgrade {
				btn.Enable()
			} else {
				btn.Disable()
			}
		}

		btn.OnTapped = func() {
			var success bool
			var nextLevel string

			inv.Lock()
			switch ci := currentItem.(type) {
			case *Pickaxe:
				next := pickaxeUpgrade[*ci]
				cost := pickaxeResourceCost[next]
				hasAll := true
				for r, v := range cost {
					if inv.Resources[r] < v {
						hasAll = false
					}
				}
				if hasAll {
					for r, v := range cost {
						inv.Resources[r] -= v
					}
					*ci = next
					success = true
					nextLevel = string(next)
				}
			case *Sword:
				next := swordUpgrade[*ci]
				cost := swordUpgradeCost[next]
				hasAll := true
				for r, v := range cost {
					if inv.Resources[r] < v {
						hasAll = false
					}
				}
				if hasAll {
					for r, v := range cost {
						inv.Resources[r] -= v
					}
					*ci = next
					success = true
					nextLevel = string(next)
				}
			case *Armor:
				next := armorUpgrade[*ci]
				cost := armorUpgradeCost[next]
				hasAll := true
				for r, v := range cost {
					if inv.Resources[r] < v {
						hasAll = false
					}
				}
				if hasAll {
					for r, v := range cost {
						inv.Resources[r] -= v
					}
					*ci = next
					success = true
					nextLevel = string(next)
				}
			}
			inv.Unlock()

			if success {
				updateStatus() // Status-Labels aktualisieren
				dialog.ShowInformation("Erfolg", fmt.Sprintf("%s wurde auf %s verbessert!", name, nextLevel), w)
			} else {
				dialog.ShowError(fmt.Errorf("Nicht genug Ressourcen"), w)
			}

			updateButtonText() // Button aktualisieren
		}

		updateButtonText()
		content.Add(btn)
	}

	addUpgradeButton("Pickaxe", &inv.Pickaxe)
	addUpgradeButton("Sword", &inv.Sword)
	addUpgradeButton("Armor", &inv.Armor)

	updateStatus() // initiale Statusanzeige

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Smithinnen.png")
	background.FillMode = canvas.ImageFillStretch

	w.SetContent(container.NewMax(
		background,
		container.NewScroll(content),
	))

	w.Show()
}

/* ===================== INVENTAR ===================== */
func openInventory(a fyne.App, inv *Inventory) {
	w := a.NewWindow("Inventar")
	enableBackspaceClose(w)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(800, 600))

	inv.Lock()
	// Labels für Geld und Ausrüstung als canvas.Text (schwarz + fett)
	moneyLabel := canvas.NewText(fmt.Sprintf("Geld: %d", inv.Money), color.White)
	moneyLabel.TextStyle = fyne.TextStyle{Bold: true}

	cristalleLabel := canvas.NewText(fmt.Sprintf("Cristalle: %d", inv.Kristalle), color.White)
	cristalleLabel.TextStyle = fyne.TextStyle{Bold: true}

	wiedergeburtenLabel := canvas.NewText(fmt.Sprintf("Wiedergeburten: %d", inv.Wiedergeburten), color.White)
	wiedergeburtenLabel.TextStyle = fyne.TextStyle{Bold: true}

	pickaxeLabel := canvas.NewText(fmt.Sprintf("Spitzhacke: %s", inv.Pickaxe), color.White)
	pickaxeLabel.TextStyle = fyne.TextStyle{Bold: true}

	swordLabel := canvas.NewText(fmt.Sprintf("Schwert: %s (Attack: %d)", inv.Sword, swordAttack[inv.Sword]), color.White)
	swordLabel.TextStyle = fyne.TextStyle{Bold: true}

	armorLabel := canvas.NewText(fmt.Sprintf("Rüstung: %s (Verteidigung: %d)", inv.Armor, armorDefense[inv.Armor]), color.White)
	armorLabel.TextStyle = fyne.TextStyle{Bold: true}
	inv.Unlock()

	// Spalten
	erzBox := container.NewVBox()
	barrenBox := container.NewVBox()
	kristallBox := container.NewVBox()
	resourceLabels := make(map[Resource]*canvas.Text)

	// Überschriften
	erzBox.Add(canvas.NewText("Erze", color.White))
	barrenBox.Add(canvas.NewText("Barren", color.White))
	kristallBox.Add(canvas.NewText("Geschliffene Edelsteine", color.White))

	// Listen
	erzListe := []Resource{Kohle, Stone, Kupfererz, Eisenerz, Silvererz, Golderz, Platinerz, Mythrilerz, Adamantiumerz, Rubinerz, Diamanterz, Blutkristallerz}
	barrenListe := []Resource{Stahlbarren, Kupferbarren, Eisenbarren, Silberbarren, Goldbarren, Platinbarren, Mythrilbarren, Adamantiumbarren}
	kristallListe := []Resource{GeschliffenerRubin, GeschliffenerDiamant, GeschliffenerBlutkristall}

	inv.Lock()
	// Erz-Spalte
	for _, r := range erzListe {
		label := canvas.NewText(fmt.Sprintf("%s: %d", r, inv.Resources[r]), color.White)
		label.TextStyle = fyne.TextStyle{Bold: true}
		resourceLabels[r] = label
		erzBox.Add(label)
	}

	// Barren-Spalte
	for i := 0; i < 1; i++ {
		barrenBox.Add(canvas.NewText(" ", color.White))
	}

	for _, b := range barrenListe {
		label := canvas.NewText(fmt.Sprintf("%s: %d", b, inv.Resources[b]), color.White)
		label.TextStyle = fyne.TextStyle{Bold: true}
		resourceLabels[b] = label
		barrenBox.Add(label)
	}

	// Kristall-Spalte
	for _, r := range kristallListe {
		label := canvas.NewText(fmt.Sprintf("%s: %d", r, inv.Resources[r]), color.White)
		label.TextStyle = fyne.TextStyle{Bold: true}
		resourceLabels[r] = label
		kristallBox.Add(label)
	}
	inv.Unlock()

	ButtonFürKarte := widget.NewButton("Zur Karte", func() {
		öffneKarte(a, inv, w)
	})

	ButtonContainer := container.NewGridWrap(
		fyne.NewSize(180, 40),
		ButtonFürKarte,
	)

	// Layout: Oben Geld + Ausrüstung
	topBox := container.NewVBox(
		moneyLabel,
		cristalleLabel,
		wiedergeburtenLabel,
		pickaxeLabel,
		swordLabel,
		armorLabel,
		widget.NewSeparator(),
	)

	topRow := container.NewHBox(
		topBox,
		layout.NewSpacer(), // schiebt die Karte nach rechts
		ButtonContainer,
	)

	resourceColumns := container.NewHBox(
		erzBox,
		widget.NewSeparator(),
		barrenBox,
		widget.NewSeparator(),
		kristallBox,
	)

	mainBox := container.NewVBox(
		topRow,
		resourceColumns,
	)

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Inventar.png")
	background.FillMode = canvas.ImageFillStretch

	w.SetContent(container.NewMax(
		background,
		mainBox,
	))

	w.Show()

	// Live-Update
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			inv.Lock()
			moneyLabel.Text = fmt.Sprintf("Geld: %d", inv.Money)
			cristalleLabel.Text = fmt.Sprintf("Cristalle: %d", inv.Kristalle)
			wiedergeburtenLabel.Text = fmt.Sprintf("Wiedergeburten: %d", inv.Wiedergeburten)
			pickaxeLabel.Text = fmt.Sprintf("Spitzhacke: %s", inv.Pickaxe)
			swordLabel.Text = fmt.Sprintf("Schwert: %s (Attack: %d)", inv.Sword, swordAttack[inv.Sword])
			armorLabel.Text = fmt.Sprintf("Rüstung: %s (Verteidigung: %d)", inv.Armor, armorDefense[inv.Armor])

			for _, r := range erzListe {
				resourceLabels[r].Text = fmt.Sprintf("%s: %d", r, inv.Resources[r])
			}
			for _, b := range barrenListe {
				resourceLabels[b].Text = fmt.Sprintf("%s: %d", b, inv.Resources[b])
			}
			for _, r := range kristallListe {
				resourceLabels[r].Text = fmt.Sprintf("%s: %d", r, inv.Resources[r])
			}
			inv.Unlock()

			// Refresh canvas.Text
			canvas.Refresh(moneyLabel)
			canvas.Refresh(cristalleLabel)
			canvas.Refresh(wiedergeburtenLabel)
			canvas.Refresh(pickaxeLabel)
			canvas.Refresh(swordLabel)
			canvas.Refresh(armorLabel)
			for _, lbl := range resourceLabels {
				canvas.Refresh(lbl)
			}
		}
	}()
}

func öffneKarte(a fyne.App, inv *Inventory, invWindow fyne.Window) {
	// Inventar schließen
	invWindow.Close()

	// Kartenfenster
	karteWindow := a.NewWindow("Karte")
	karteWindow.Resize(fyne.NewSize(800, 600))
	karteWindow.CenterOnScreen()

	// Wenn Karte geschlossen wird → Inventar wieder öffnen
	karteWindow.SetOnClosed(func() {
		openInventory(a, inv)
	})

	// Karte als Hintergrund
	karte := canvas.NewImageFromFile("bild/Karte.png")
	karte.FillMode = canvas.ImageFillStretch

	karteWindow.SetContent(container.NewMax(karte))
	karteWindow.Show()
}

/* ===================== DUNGEON ===================== */
func openDungeon(a fyne.App, mainWindow fyne.Window, inv *Inventory) {
	// Hauptfenster ausblenden, nicht schließen
	mainWindow.Hide()

	w := a.NewWindow("Dungeon")
	w.Resize(fyne.NewSize(400, 300))
	w.CenterOnScreen()

	// Funktion zum Schließen des Dungeon-Fensters
	closeDungeon := func() {
		w.Close()
		mainWindow.Show() // Hauptfenster wieder anzeigen
	}

	// Lebenspunkte des Spielers
	playerHealth := 200
	playerHealth = playerHealth / 10 * armorDefense[inv.Armor]
	// Statt widget.NewLabel
	healthLabel := canvas.NewText(fmt.Sprintf("%d", playerHealth), color.White)
	healthLabel.TextStyle = fyne.TextStyle{Bold: true}
	healthLabel.Move(fyne.NewPos(150, 10))

	// Container für Dungeon-Level-Buttons
	levelButtonContainer := container.NewVBox()
	currentLevel := 1

	// Initiale Ebene erstellen
	createLevel(a, inv, w, mainWindow, &currentLevel, levelButtonContainer, healthLabel, playerHealth, closeDungeon)

	bg := canvas.NewImageFromFile("bild/Dungeoneingang.png")
	bg.FillMode = canvas.ImageFillStretch

	// Statt widget.NewLabelWithStyle
	title := canvas.NewText("Dungeon", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	content := container.NewVBox(
		title,
		levelButtonContainer,
		widget.NewButton("Dungeon verlassen", closeDungeon),
	)

	w.SetContent(
		container.NewMax(
			bg,
			container.NewPadded(content),
		),
	)

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == fyne.KeyBackspace || k.Name == fyne.KeyEscape {
			closeDungeon()
		}
	})
	w.Show()
}
func createLevel(a fyne.App, inv *Inventory, w fyne.Window, mainWindow fyne.Window, currentLevel *int,
	levelButtonContainer *fyne.Container, healthText *canvas.Text, playerHealth int, closeDungeon func()) {

	level := *currentLevel
	levelName := fmt.Sprintf("Ebene %d", level)

	adventureBtn := widget.NewButton(levelName, func() {
		w.Hide()
		w2 := a.NewWindow(levelName)
		w2.Resize(fyne.NewSize(1470, 765))
		w2.CenterOnScreen()

		// Spieler erstellen
		player := canvas.NewImageFromFile("bild/player_down.png")
		player.FillMode = canvas.ImageFillContain
		player.Resize(fyne.NewSize(50, 100))
		x, y := float32(50), float32(500)
		player.Move(fyne.NewPos(x, y))

		// Health Bar erstellen
		healthBarWidth := float32(playerHealth)
		healthBarHeight := float32(5)
		healthBar := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
		healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
		healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))

		background := canvas.NewImageFromFile("bild/Dungeonboden.png")
		background.FillMode = canvas.ImageFillStretch
		background.Resize(fyne.NewSize(windowWidth, windowHeight))
		background.Move(fyne.NewPos(0, 0))

		// Setze Position des Health Texts
		healthText.Move(fyne.NewPos(725, 10))

		objects := []fyne.CanvasObject{background, player, healthText}

		// Gegner-Typ
		type Enemy struct {
			Rect      *canvas.Image
			Health    int
			HealthBar *canvas.Rectangle
			Name      string
			NameLabel *canvas.Text
			X, Y      float32
			Speed     float32
		}

		// Gegner erstellen
		numEnemies := level + 1
		enemies := []*Enemy{}
		for i := 0; i < numEnemies; i++ {
			e := &Enemy{
				Rect:   canvas.NewImageFromFile("bild/Gegner.png"),
				Name:   "Magmamonster",
				X:      float32(rand.Intn(1420) + 50),
				Y:      float32(rand.Intn(715) + 50),
				Speed:  1 + float32(level)*0.5,
				Health: 50 + level*10,
			}

			e.Rect.Resize(fyne.NewSize(100, 100))
			e.Rect.Move(fyne.NewPos(e.X, e.Y))
			e.NameLabel = canvas.NewText(e.Name, color.White)
			e.NameLabel.TextStyle = fyne.TextStyle{Bold: true}
			e.NameLabel.Move(fyne.NewPos(e.X-5, e.Y-healthBarHeight-18))
			e.NameLabel.Resize(e.NameLabel.MinSize())

			e.HealthBar = canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
			e.HealthBar.Resize(fyne.NewSize(float32(e.Health)/10, healthBarHeight))
			e.HealthBar.Move(fyne.NewPos(e.X-5, e.Y-e.HealthBar.Size().Height-2))

			objects = append(objects, e.Rect, e.HealthBar, e.NameLabel) // NameLabel zuletzt
			enemies = append(enemies, e)
		}

		w2.SetContent(container.NewWithoutLayout(objects...))

		// Spielerbewegung
		playerSprites := map[string]string{
			"up":    "bild/LaufenRückwärts.png",
			"down":  "bild/LaufenForwärts.png",
			"left":  "bild/LaufenLinks.png",
			"right": "bild/LaufenRechts.png",
		}
		currentDir := "down"

		w2.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
			switch k.Name {
			case fyne.KeyUp:
				y -= speed
				currentDir = "up"
			case fyne.KeyDown:
				y += speed
				currentDir = "down"
			case fyne.KeyLeft:
				x -= speed
				currentDir = "left"
			case fyne.KeyRight:
				x += speed
				currentDir = "right"
			case fyne.KeyEscape, fyne.KeyBackspace:
				w2.Close()
				w.Show()
				return
			}

			x = clamp(x, 0, 1470-playerW)
			y = clamp(y, 0, 765-playerH)

			// Main Thread: UI-Update mit canvas.Refresh
			player.Move(fyne.NewPos(x, y))
			player.File = playerSprites[currentDir]
			player.Refresh()

			healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))
			healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
			healthBar.Refresh()

			healthText.Text = fmt.Sprintf("%d", int(healthBarWidth))
			canvas.Refresh(healthText)
		})

		// Spieler-Health-Ticker
		ticker := time.NewTicker(1 * time.Second)
		go func() {
			for range ticker.C {
				if healthBarWidth > 0 {
					playerHealth -= 1
					healthBarWidth = float32(playerHealth)

					healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
					healthBar.Refresh()

					healthText.Text = fmt.Sprintf("%d", int(healthBarWidth))
					canvas.Refresh(healthText)
				} else {
					ticker.Stop()
					healthText.Text = "Du bist gestorben!"
					canvas.Refresh(healthText)

					time.AfterFunc(2*time.Second, func() {
						w2.Close()
						closeDungeon()
					})
					return
				}
			}
		}()

		// Gegner-Logik
		go func() {
			for {
				allEnemiesDead := true
				for _, enemy := range enemies {
					e := enemy
					distance := math.Sqrt(float64((x-e.X)*(x-e.X) + (y-e.Y)*(y-e.Y)))

					if distance < 50 && e.Health > 0 {
						e.Health -= 10 * swordAttack[inv.Sword]
						e.HealthBar.Resize(fyne.NewSize(float32(e.Health)/10, healthBarHeight))
						e.HealthBar.Refresh()
					}

					if e.Health > 0 {
						allEnemiesDead = false
					} else if e.NameLabel.Text != "☠" {
						e.NameLabel.Text = "☠"
						e.NameLabel.TextSize = 24
						canvas.Refresh(e.NameLabel)
						e.Rect.Hide()
						e.HealthBar.Hide()
					}
				}

				if allEnemiesDead {
					ticker.Stop()
					inv.Kristalle += 1 * (inv.Wiedergeburten + 1) * level

					playerHealth = 100 / 10 * armorDefense[inv.Armor]
					healthText.Text = "Du hast Gewonnen!"
					canvas.Refresh(healthText)

					time.AfterFunc(2*time.Second, func() {
						w2.Close()
						w.Show()
						*currentLevel++
						createLevel(a, inv, w, mainWindow, currentLevel, levelButtonContainer, healthText, playerHealth, closeDungeon)
					})
					return
				}
				time.Sleep(1 * time.Second)
			}
		}()

		w2.Show()
	})

	levelButtonContainer.Add(adventureBtn)
	levelButtonContainer.Refresh()
}
