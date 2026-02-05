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

	player := canvas.NewRectangle(color.RGBA{0, 200, 100, 255})
	player.Resize(fyne.NewSize(playerW, playerH))
	x, y := float32(50), float32(50)
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
	dungeonPos := fyne.NewPos(700, 600)
	dungeon := canvas.NewImageFromFile("bild/Dungeon.png")
	dungeon.Move(dungeonPos)
	dungeon.Resize(fyne.NewSize(64, 64))
	dungeon.FillMode = canvas.ImageFillContain

	shopPos := fyne.NewPos(100, 20)
	shop := canvas.NewImageFromFile("bild/Shop.png")
	shop.Move(shopPos)
	shop.Resize(fyne.NewSize(64, 64))
	shop.FillMode = canvas.ImageFillContain

	smithPos := fyne.NewPos(400, 20)
	smith := canvas.NewImageFromFile("bild/Smith.png")
	smith.Move(smithPos)
	smith.Resize(fyne.NewSize(64, 64))
	smith.FillMode = canvas.ImageFillContain

	kristallShopPos := fyne.NewPos(700, 20)
	kristall := canvas.NewImageFromFile("bild/Kristallshop.png")
	kristall.Move(kristallShopPos)
	kristall.Resize(fyne.NewSize(64, 64))
	kristall.FillMode = canvas.ImageFillContain

	schmelzofenPos := fyne.NewPos(1000, 20)
	schmelzofen := canvas.NewImageFromFile("bild/Schmelzofen.png")
	schmelzofen.Move(schmelzofenPos)
	schmelzofen.Resize(fyne.NewSize(64, 64))
	schmelzofen.FillMode = canvas.ImageFillContain

	wiedergeburtenLadnePos := fyne.NewPos(1200, 20)
	wiedergeburtenladen := canvas.NewImageFromFile("bild/Wiedergeburtenladen.png")
	wiedergeburtenladen.Move(wiedergeburtenLadnePos)
	wiedergeburtenladen.Resize(fyne.NewSize(64, 64))
	wiedergeburtenladen.FillMode = canvas.ImageFillContain

	miningweltPos := fyne.NewPos(1200, 400)
	miningwelt := canvas.NewImageFromFile("bild/Miningwelt.png")
	miningwelt.Move(miningweltPos)
	miningwelt.Resize(fyne.NewSize(64, 64))
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
				openKristallShop(a, inv, mines)
				return
			}
			if dist2(p, schmelzofenPos) < interactionDist*interactionDist {
				openSchmelzofen(a, inv)
				return
			}
			if dist2(p, wiedergeburtenLadnePos) < interactionDist*interactionDist {
				openWiedergeburtenLaden(a, inv, mines)
				return
			}
			if dist2(p, dungeonPos) < interactionDist*interactionDist {
				openDungeon(a, w, inv)
				return
			}
			if dist2(p, miningweltPos) < interactionDist*interactionDist {
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
	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()

	var objects []fyne.CanvasObject
	xPositions := []float32{100, 250, 400}
	y := float32(150)

	hochPos := fyne.NewPos(225, 20)
	hoch := canvas.NewImageFromFile("bild/LeiterHoch.png")
	hoch.Move(hochPos)
	hoch.Resize(fyne.NewSize(64, 64))
	hoch.FillMode = canvas.ImageFillContain

	runterPos := fyne.NewPos(225, 300)
	runter := canvas.NewImageFromFile("bild/LeiterRunter.png")
	runter.Move(runterPos)
	runter.Resize(fyne.NewSize(64, 64))
	runter.FillMode = canvas.ImageFillContain

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Steinboden1.png")
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

	player := canvas.NewRectangle(color.RGBA{0, 200, 100, 255})
	player.Resize(fyne.NewSize(playerW, playerH))
	x, y := float32(50), float32(50)
	player.Move(fyne.NewPos(x, y))

	playerHealth := 100
	healthBarWidth := float32(playerHealth) / 10
	healthBarHeight := float32(5)
	healthBar := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
	healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))

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

		healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))
		healthBar.Refresh()
	})

	objects = append(objects, player, healthBar)
	w.SetContent(container.NewWithoutLayout(objects...))
	w.Show()
}

func openMiningwelt2(a fyne.App, inv *Inventory, mines []*Mine) {
	w := a.NewWindow("Miningwelt 2")
	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()

	var objects []fyne.CanvasObject
	xPositions := []float32{100, 250, 400}
	y := float32(150)

	hochPos := fyne.NewPos(225, 20)
	hoch := canvas.NewImageFromFile("bild/LeiterHoch.png")
	hoch.Move(hochPos)
	hoch.Resize(fyne.NewSize(64, 64))
	hoch.FillMode = canvas.ImageFillContain

	runterPos := fyne.NewPos(225, 300)
	runter := canvas.NewImageFromFile("bild/LeiterRunter.png")
	runter.Move(runterPos)
	runter.Resize(fyne.NewSize(64, 64))
	runter.FillMode = canvas.ImageFillContain

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Steinboden1.png")
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

	player := canvas.NewRectangle(color.RGBA{0, 200, 100, 255})
	player.Resize(fyne.NewSize(playerW, playerH))
	x, y := float32(50), float32(50)
	player.Move(fyne.NewPos(x, y))

	playerHealth := 100
	healthBarWidth := float32(playerHealth) / 10
	healthBarHeight := float32(5)
	healthBar := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
	healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))

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

		healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))
		healthBar.Refresh()
	})
	objects = append(objects, player, healthBar)
	w.SetContent(container.NewWithoutLayout(objects...))
	w.Show()
}

func openMiningwelt3(a fyne.App, inv *Inventory, mines []*Mine) {
	w := a.NewWindow("Miningwelt 3")
	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()

	var objects []fyne.CanvasObject
	xPositions := []float32{100, 250, 400}
	y := float32(150)

	hochPos := fyne.NewPos(225, 20)
	hoch := canvas.NewImageFromFile("bild/LeiterHoch.png")
	hoch.Move(hochPos)
	hoch.Resize(fyne.NewSize(64, 64))
	hoch.FillMode = canvas.ImageFillContain

	runterPos := fyne.NewPos(225, 300)
	runter := canvas.NewImageFromFile("bild/LeiterRunter.png")
	runter.Move(runterPos)
	runter.Resize(fyne.NewSize(64, 64))
	runter.FillMode = canvas.ImageFillContain

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Steinboden1.png")
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

	player := canvas.NewRectangle(color.RGBA{0, 200, 100, 255})
	player.Resize(fyne.NewSize(playerW, playerH))
	x, y := float32(50), float32(50)
	player.Move(fyne.NewPos(x, y))

	playerHealth := 100
	healthBarWidth := float32(playerHealth) / 10
	healthBarHeight := float32(5)
	healthBar := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
	healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))

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

		healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))
		healthBar.Refresh()
	})

	objects = append(objects, player, healthBar)
	w.SetContent(container.NewWithoutLayout(objects...))
	w.Show()
}

func openMiningwelt4(a fyne.App, inv *Inventory, mines []*Mine) {
	w := a.NewWindow("Miningwelt 4")
	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()

	var objects []fyne.CanvasObject
	xPositions := []float32{100, 250, 400}
	y := float32(150)

	hochPos := fyne.NewPos(225, 20)
	hoch := canvas.NewImageFromFile("bild/LeiterHoch.png")
	hoch.Move(hochPos)
	hoch.Resize(fyne.NewSize(64, 64))
	hoch.FillMode = canvas.ImageFillContain

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Steinboden1.png")
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

	player := canvas.NewRectangle(color.RGBA{0, 200, 100, 255})
	player.Resize(fyne.NewSize(playerW, playerH))
	x, y := float32(50), float32(50)
	player.Move(fyne.NewPos(x, y))

	playerHealth := 100
	healthBarWidth := float32(playerHealth) / 10
	healthBarHeight := float32(5)
	healthBar := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
	healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))

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

	// Kristallanzeige
	kristallLabel := widget.NewLabel(fmt.Sprintf("Kristalle: %d", inv.Kristalle))
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

	skillLabel := widget.NewLabel(fmt.Sprintf("Wiedergeburten-Skill Level %d", currentLevel))

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
		skillLabel.SetText(fmt.Sprintf("Wiedergeburt %d", currentLevel))
	}
	updateButton()

	btn.OnTapped = func() {
		if currentLevel >= maxLevel {
			dialog.ShowInformation("Maximale Wiedergeburt", "Du hast bereits die höheste Wiedergeburt ereicht.", w)
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
			mines[0].Owned = false
			mines[1].Owned = false
			mines[2].Owned = false
			mines[3].Owned = false
			mines[4].Owned = false
			mines[5].Owned = false
			mines[6].Owned = false
			mines[7].Owned = false
			mines[8].Owned = false
			mines[9].Owned = false
			mines[10].Owned = false
			mines[11].Owned = false

			mines[0].Rate = mines[0].Rate * (inv.Wiedergeburten + 1)
			mines[1].Rate = mines[1].Rate * (inv.Wiedergeburten + 1)
			mines[2].Rate = mines[2].Rate * (inv.Wiedergeburten + 1)
			mines[3].Rate = mines[3].Rate * (inv.Wiedergeburten + 1)
			mines[4].Rate = mines[4].Rate * (inv.Wiedergeburten + 1)
			mines[5].Rate = mines[5].Rate * (inv.Wiedergeburten + 1)
			mines[6].Rate = mines[6].Rate * (inv.Wiedergeburten + 1)
			mines[7].Rate = mines[7].Rate * (inv.Wiedergeburten + 1)
			mines[8].Rate = mines[8].Rate * (inv.Wiedergeburten + 1)
			mines[9].Rate = mines[9].Rate * (inv.Wiedergeburten + 1)
			mines[10].Rate = mines[10].Rate * (inv.Wiedergeburten + 1)
			mines[11].Rate = mines[11].Rate * (inv.Wiedergeburten + 1)

			kristallLabel.SetText(fmt.Sprintf("Kristalle: %d", inv.Kristalle))
			updateButton()
			dialog.ShowInformation("Erfolg", fmt.Sprintf("Wiedergeburt %d freigeschaltet!", currentLevel), w)
		} else {
			dialog.ShowInformation("Nicht genug Kristalle", "Du hast nicht genügend Kristalle, um wiedergebore zu werden.", w)
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

	centeredBox := container.NewCenter(actionBtn)
	// Hintergrund
	background := canvas.NewImageFromFile("bild/Miningshopinnen.png")
	background.FillMode = canvas.ImageFillStretch
	background.Resize(fyne.NewSize(windowWidth, windowHeight))
	background.Move(fyne.NewPos(0, 0))

	w.SetContent(container.NewMax(
		background,
		label,
		centeredBox,
	))
	updateLabel()
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
		"Steinmine Multiplier": {
			1: func(mines []*Mine) { mines[1].Rate = int(float64(mines[1].Rate) * 1.2) },
			2: func(mines []*Mine) { mines[1].Rate = int(float64(mines[1].Rate) * 1.5) },
			3: func(mines []*Mine) { mines[1].Rate = int(float64(mines[1].Rate) * 2.0) },
		},
		"Kupfermine Multiplier": {
			1: func(mines []*Mine) { mines[2].Rate = int(float64(mines[2].Rate) * 1.2) },
			2: func(mines []*Mine) { mines[2].Rate = int(float64(mines[2].Rate) * 1.5) },
			3: func(mines []*Mine) { mines[2].Rate = int(float64(mines[2].Rate) * 2.0) },
		},
		"Eisenmine Multiplier": {
			1: func(mines []*Mine) { mines[3].Rate = int(float64(mines[3].Rate) * 1.2) },
			2: func(mines []*Mine) { mines[3].Rate = int(float64(mines[3].Rate) * 1.5) },
			3: func(mines []*Mine) { mines[3].Rate = int(float64(mines[3].Rate) * 2.0) },
		},
		"Silbermine Multiplier": {
			1: func(mines []*Mine) { mines[4].Rate = int(float64(mines[4].Rate) * 1.2) },
			2: func(mines []*Mine) { mines[4].Rate = int(float64(mines[4].Rate) * 1.5) },
			3: func(mines []*Mine) { mines[4].Rate = int(float64(mines[4].Rate) * 2.0) },
		},
		"Goldmine Multiplier": {
			1: func(mines []*Mine) { mines[5].Rate = int(float64(mines[5].Rate) * 1.2) },
			2: func(mines []*Mine) { mines[5].Rate = int(float64(mines[5].Rate) * 1.5) },
			3: func(mines []*Mine) { mines[5].Rate = int(float64(mines[5].Rate) * 2.0) },
		},
		"Platinmine Multiplier": {
			1: func(mines []*Mine) { mines[6].Rate = int(float64(mines[6].Rate) * 1.2) },
			2: func(mines []*Mine) { mines[6].Rate = int(float64(mines[6].Rate) * 1.5) },
			3: func(mines []*Mine) { mines[6].Rate = int(float64(mines[6].Rate) * 2.0) },
		},
		"Rubinmine Multiplier": {
			1: func(mines []*Mine) { mines[7].Rate = int(float64(mines[7].Rate) * 1.2) },
			2: func(mines []*Mine) { mines[7].Rate = int(float64(mines[7].Rate) * 1.5) },
			3: func(mines []*Mine) { mines[7].Rate = int(float64(mines[7].Rate) * 2.0) },
		},
		"Diamandmine Multiplier": {
			1: func(mines []*Mine) { mines[8].Rate = int(float64(mines[8].Rate) * 1.2) },
			2: func(mines []*Mine) { mines[8].Rate = int(float64(mines[8].Rate) * 1.5) },
			3: func(mines []*Mine) { mines[8].Rate = int(float64(mines[8].Rate) * 2.0) },
		},
		"Mythrilmine Multiplier": {
			1: func(mines []*Mine) { mines[9].Rate = int(float64(mines[9].Rate) * 1.2) },
			2: func(mines []*Mine) { mines[9].Rate = int(float64(mines[9].Rate) * 1.5) },
			3: func(mines []*Mine) { mines[9].Rate = int(float64(mines[9].Rate) * 2.0) },
		},
		"Blutkristallmine Multiplier": {
			1: func(mines []*Mine) { mines[10].Rate = int(float64(mines[10].Rate) * 1.2) },
			2: func(mines []*Mine) { mines[10].Rate = int(float64(mines[10].Rate) * 1.5) },
			3: func(mines []*Mine) { mines[10].Rate = int(float64(mines[10].Rate) * 2.0) },
		},
		"Adamantiummine Multiplier": {
			1: func(mines []*Mine) { mines[11].Rate = int(float64(mines[11].Rate) * 1.2) },
			2: func(mines []*Mine) { mines[11].Rate = int(float64(mines[11].Rate) * 1.5) },
			3: func(mines []*Mine) { mines[11].Rate = int(float64(mines[11].Rate) * 2.0) },
		},
	}

	// Fenster erstellen
	w := a.NewWindow("Skill Tree Shop")
	enableBackspaceClose(w)
	w.CenterOnScreen()

	// Kristallanzeige
	kristallLabel := createKristallLabel(inv)

	// Hauptbox für das Fenster
	mainBox := container.NewVBox(
		kristallLabel,
		widget.NewSeparator(),
		widget.NewLabel("Fähigkeiten freischalten"),
	)

	// VBox für die Skills
	skillsBox := container.NewVBox()

	// Buttons für alle Fähigkeiten in fester Reihenfolge erstellen

	// Hier übergibst du 'm' (das aktuelle Mine-Objekt) an createSkillButton
	for _, skill := range skillOrder {
		levels := skills[skill]
		currentLevel := currentLevels[skill]
		skillBox := createSkillButton(skill, levels, currentLevel, inv, kristallLabel, levelUpFunctions, w)
		skillsBox.Add(skillBox)
	}

	// Hintergrund
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

// Hilfsfunktion für die Kristallanzeige
func createKristallLabel(inv *Inventory) *widget.Label {
	label := widget.NewLabel(fmt.Sprintf("Kristalle: %d", inv.Kristalle))
	label.TextStyle = fyne.TextStyle{Bold: true}
	return label
}

// Hilfsfunktion für die Erstellung von Buttons
// Funktion zum Erstellen eines Skill-Buttons
func createSkillButton(skill string, levels map[int]int, currentLevel int, inv *Inventory, kristallLabel *widget.Label, levelUpFunctions map[string]map[int]func(mines []*Mine), w fyne.Window) *fyne.Container {
	skillBox := container.NewVBox()
	skillLabel := widget.NewLabel(fmt.Sprintf("%s", skill))
	skillBox.Add(skillLabel)

	// Button erstellen
	btn := widget.NewButton(fmt.Sprintf("%s Level %d (Kosten: %d Kristalle)", skill, currentLevel+1, levels[currentLevel+1]), nil)

	// Button-Event
	btn.OnTapped = func() {
		// Hier wird das *Mine-Objekt 'm' korrekt übergeben
		handleSkillUnlock(skill, levels, inv, kristallLabel, levelUpFunctions, btn, skillLabel, w)
	}

	// Button hinzufügen
	skillBox.Add(btn)

	// Anzeige des Levels aktualisieren
	updateSkillDisplay(skill, currentLevel, levels, btn, skillLabel)

	return skillBox
}

// Funktion zur Aktualisierung der Anzeige für den Skill
func updateSkillDisplay(skill string, currentLevel int, levels map[int]int, btn *widget.Button, skillLabel *widget.Label) {
	btn.SetText(fmt.Sprintf("%s Level %d (Kosten: %d Kristalle)", skill, currentLevel+1, levels[currentLevel+1]))
	skillLabel.SetText(fmt.Sprintf("%s (Level %d freigeschaltet)", skill, currentLevel+1))
}

// Funktion für das Freischalten des Skills
func handleSkillUnlock(skill string, levels map[int]int, inv *Inventory, kristallLabel *widget.Label, levelUpFunctions map[string]map[int]func(mines []*Mine), btn *widget.Button, skillLabel *widget.Label, w fyne.Window) {
	inv.Lock()
	defer inv.Unlock()

	currentLevel := currentLevels[skill]

	// Überprüfen, ob genügend Kristalle vorhanden sind und der Level noch nicht maximal ist
	if currentLevel < len(levels) && inv.Kristalle >= levels[currentLevel+1] {
		// Kristalle abziehen und das Label aktualisieren
		inv.Kristalle -= levels[currentLevel+1]
		kristallLabel.SetText(fmt.Sprintf("Kristalle: %d", inv.Kristalle))

		// Level erhöhen und die Anzeige aktualisieren
		currentLevel++
		currentLevels[skill] = currentLevel
		updateSkillDisplay(skill, currentLevel, levels, btn, skillLabel)

		// Sicherstellen, dass die Funktion existiert und aufrufen
		if fnMap, ok := levelUpFunctions[skill]; ok {
			if fn, ok := fnMap[currentLevel]; ok && fn != nil {
				// Hier übergeben wir das *Mine-Objekt an die Level-Up-Funktion
				fn(mines) // Überprüfe, ob 'm' hier korrekt übergeben wird
			}
		}
	} else if currentLevel < len(levels) {
		// Wenn nicht genug Kristalle vorhanden sind
		dialog.ShowInformation("Nicht genug Kristalle", "Du hast nicht genügend Kristalle, um dieses Level freizuschalten.", w)
	} else {
		// Wenn das maximale Level erreicht ist
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

	moneyLabel := widget.NewLabel("")
	moneyLabel.TextStyle = fyne.TextStyle{Bold: true}

	labels := make(map[Resource]*widget.Label)

	// 🔘 Verkaufsmenge (global)
	selectedAmount := 1
	amountSelector := widget.NewRadioGroup(
		[]string{"1", "10", "100"},
		func(value string) {
			switch value {
			case "10":
				selectedAmount = 10
			case "100":
				selectedAmount = 100
			default:
				selectedAmount = 1
			}
		},
	)
	amountSelector.SetSelected("1")

	// Verkaufsblock pro Resource
	createSellBlock := func(r Resource) *fyne.Container {
		box := container.NewVBox()

		label := widget.NewLabel("")
		labels[r] = label
		box.Add(label)

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

			label.SetText(fmt.Sprintf("%s: %d", r, inv.Resources[r]))
			moneyLabel.SetText(fmt.Sprintf("Geld: %d", inv.Money))
		})

		box.Add(btn)
		box.Add(widget.NewSeparator())
		return box
	}

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

	topBar := container.NewHBox(
		moneyLabel,
		widget.NewSeparator(),
		widget.NewLabel("Menge:"),
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
	// Hintergrund
	background := canvas.NewImageFromFile("bild/Shopinnen.png")
	background.FillMode = canvas.ImageFillStretch

	// Content + Hintergrund kombinieren
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

	// Hintergrund
	background := canvas.NewImageFromFile("bild/Schmelzofeninnen.png")
	background.FillMode = canvas.ImageFillStretch

	// Content + Hintergrund kombinieren
	w.SetContent(container.NewMax(
		background,
		content,
	))
	w.Resize(fyne.NewSize(700, 450))
	w.Show()
}

/* ===================== SCHMIED ===================== */
func openSmith(a fyne.App, inv *Inventory) {
	w := a.NewWindow("Schmiede")
	w.Resize(fyne.NewSize(400, 350))
	w.CenterOnScreen()

	content := container.NewVBox()

	// Statusanzeige
	moneyLabel := widget.NewLabel("")
	pickaxeLabel := widget.NewLabel("")
	swordLabel := widget.NewLabel("")
	armorLabel := widget.NewLabel("")

	updateStatus := func() {
		inv.Lock()
		m := inv.Money
		p := inv.Pickaxe
		s := inv.Sword
		a := inv.Armor
		inv.Unlock()

		moneyLabel.SetText(fmt.Sprintf("Geld: %d", m))
		pickaxeLabel.SetText(fmt.Sprintf("Pickaxe: %s", p))
		swordLabel.SetText(fmt.Sprintf("Sword: %s", s))
		armorLabel.SetText(fmt.Sprintf("Armor: %s", a))
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

	// Upgrade-Button Helper
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
	// Labels für Geld und Ausrüstung
	moneyLabel := widget.NewLabel(fmt.Sprintf("Geld: %d", inv.Money))
	cristalleLabel := widget.NewLabel(fmt.Sprintf("Cristalle: %d", inv.Kristalle))
	wiedergeburtenLabel := widget.NewLabel(fmt.Sprintf("Wiedergeburten: %d", inv.Wiedergeburten))
	pickaxeLabel := widget.NewLabel(fmt.Sprintf("Spitzhacke: %s", inv.Pickaxe))
	swordLabel := widget.NewLabel(fmt.Sprintf("Schwert: %s (Attack: %d)", inv.Sword, swordAttack[inv.Sword]))
	armorLabel := widget.NewLabel(fmt.Sprintf("Rüstung: %s (Verteidigung: %d)", inv.Armor, armorDefense[inv.Armor]))
	inv.Unlock()

	// Spalten
	erzBox := container.NewVBox()
	barrenBox := container.NewVBox()
	kristallBox := container.NewVBox()
	resourceLabels := make(map[Resource]*widget.Label)

	// Überschriften
	erzBox.Add(widget.NewLabelWithStyle("⛏ Erze", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
	barrenBox.Add(widget.NewLabelWithStyle("🔥 Barren", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
	kristallBox.Add(widget.NewLabelWithStyle("💎 Geschliffene Edelsteine", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))

	// Listen
	erzListe := []Resource{Kohle, Stone, Kupfererz, Eisenerz, Silvererz, Golderz, Platinerz, Mythrilerz, Adamantiumerz, Rubinerz, Diamanterz, Blutkristallerz}
	barrenListe := []Resource{Stahlbarren, Kupferbarren, Eisenbarren, Silberbarren, Goldbarren, Platinbarren, Mythrilbarren, Adamantiumbarren}
	kristallListe := []Resource{GeschliffenerRubin, GeschliffenerDiamant, GeschliffenerBlutkristall}

	inv.Lock()
	// Erz-Spalte
	for _, r := range erzListe {
		label := widget.NewLabel(fmt.Sprintf("%s: %d", r, inv.Resources[r]))
		resourceLabels[r] = label
		erzBox.Add(label)
	}

	// Barren-Spalte
	for i := 0; i < 1; i++ {
		barrenBox.Add(widget.NewLabel(" "))
	}

	for _, b := range barrenListe {
		label := widget.NewLabel(fmt.Sprintf("%s: %d", b, inv.Resources[b]))
		resourceLabels[b] = label
		barrenBox.Add(label)
	}

	// Kristall-Spalte
	for _, r := range kristallListe {
		label := widget.NewLabel(fmt.Sprintf("%s: %d", r, inv.Resources[r]))
		resourceLabels[r] = label
		kristallBox.Add(label)
	}
	inv.Unlock()

	karte := canvas.NewImageFromFile("bild/Karte.png")
	karte.FillMode = canvas.ImageFillContain
	karte.SetMinSize(fyne.NewSize(350, 250)) // Größe anpassen wie du willst

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
		karte,
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
			moneyLabel.SetText(fmt.Sprintf("Geld: %d", inv.Money))
			cristalleLabel.SetText(fmt.Sprintf("Cristalle: %d", inv.Kristalle))
			wiedergeburtenLabel.SetText(fmt.Sprintf("Wiedergeburten: %d", inv.Wiedergeburten)) // Hier wird der Text für das Label aktualisiert
			pickaxeLabel.SetText(fmt.Sprintf("Spitzhacke: %s", inv.Pickaxe))
			swordLabel.SetText(fmt.Sprintf("Schwert: %s (Attack: %d)", inv.Sword, swordAttack[inv.Sword]))
			armorLabel.SetText(fmt.Sprintf("Rüstung: %s (Verteidigung: %d)", inv.Armor, armorDefense[inv.Armor]))

			for _, r := range erzListe {
				resourceLabels[r].SetText(fmt.Sprintf("%s: %d", r, inv.Resources[r]))
			}
			for _, b := range barrenListe {
				resourceLabels[b].SetText(fmt.Sprintf("%s: %d", b, inv.Resources[b]))
			}
			for _, r := range kristallListe {
				resourceLabels[r].SetText(fmt.Sprintf("%s: %d", r, inv.Resources[r]))
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
	w.CenterOnScreen()

	// Funktion zum Schließen des Dungeon-Fensters
	closeDungeon := func() {
		w.Close()
		mainWindow.Show() // Hauptfenster wieder anzeigen
	}

	// Lebenspunkte des Spielers
	playerHealth := 100
	playerHealth = playerHealth / 10 * armorDefense[inv.Armor]
	healthLabel := widget.NewLabel(fmt.Sprintf("%d", playerHealth))

	// Container für Dungeon-Level-Buttons
	levelButtonContainer := container.NewVBox()
	currentLevel := 1

	// Initiale Ebene erstellen
	createLevel(a, inv, w, mainWindow, &currentLevel, levelButtonContainer, healthLabel, playerHealth, closeDungeon)

	w.SetContent(container.NewVBox(
		levelButtonContainer,
		widget.NewButton("Dungeon verlassen", closeDungeon),
	))

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == fyne.KeyBackspace || k.Name == fyne.KeyEscape {
			closeDungeon()
		}
	})

	w.Show()
}

// Ausgelagerte Level-Erstellungsfunktion
func createLevel(a fyne.App, inv *Inventory, w fyne.Window, mainWindow fyne.Window, currentLevel *int,
	levelButtonContainer *fyne.Container, healthLabel *widget.Label, playerHealth int, closeDungeon func()) {

	level := *currentLevel
	levelName := fmt.Sprintf("Ebene %d", level)

	adventureBtn := widget.NewButton(levelName, func() {
		w.Hide()
		w2 := a.NewWindow(levelName)
		w2.Resize(fyne.NewSize(400, 300))
		w.CenterOnScreen()

		// Spieler erstellen
		player := canvas.NewRectangle(color.RGBA{0, 200, 100, 255})
		player.Resize(fyne.NewSize(playerW, playerH))
		x, y := float32(50), float32(50)
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

		objects := []fyne.CanvasObject{background, player, healthLabel}

		// Gegner erstellen
		type Enemy struct {
			Rect      *canvas.Rectangle
			Health    int
			HealthBar *canvas.Rectangle
			Name      string
			NameLabel *widget.Label
			X, Y      float32
			Speed     float32
		}

		numEnemies := level + 1
		enemies := []*Enemy{}
		for i := 0; i < numEnemies; i++ {
			e := &Enemy{
				Rect:   canvas.NewRectangle(color.RGBA{255, 0, 0, 255}),
				Name:   "Magmamonster",
				X:      float32(50 + i*60),
				Y:      float32(50 + i*40),
				Speed:  1 + float32(level)*0.5,
				Health: 50 + level*10,
			}

			e.Rect.Resize(fyne.NewSize(20, 20))
			e.Rect.Move(fyne.NewPos(e.X, e.Y))
			e.NameLabel = widget.NewLabel(e.Name)
			e.NameLabel.Move(fyne.NewPos(
				e.X-5,
				e.Y-healthBarHeight-18, // über HealthBar
			))
			e.HealthBar = canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
			e.HealthBar.Resize(fyne.NewSize(float32(e.Health)/10, healthBarHeight))
			e.HealthBar.Move(fyne.NewPos(e.X-5, e.Y-e.HealthBar.Size().Height-2))
			objects = append(objects, e.Rect, e.NameLabel, e.HealthBar)
			enemies = append(enemies, e)
		}

		// Health Label oben positionieren
		healthLabel.Move(fyne.NewPos(150, 10))
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
			case fyne.KeyEscape, fyne.KeyBackspace:
				w2.Close()
				w.Show()
				return
			}
			x = clamp(x, 0, 400-playerW)
			y = clamp(y, 0, 300-playerH)
			player.Move(fyne.NewPos(x, y))
			player.Refresh()

			healthBar.Move(fyne.NewPos(x-5, y-healthBarHeight-2))
			healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
			healthBar.Refresh()
			healthLabel.SetText(fmt.Sprintf("%d", int(healthBarWidth)))
		})

		// Spieler-Health Ticker
		ticker := time.NewTicker(1 * time.Second)
		go func() {
			for range ticker.C {
				if healthBarWidth > 0 {
					playerHealth -= 1
					healthBarWidth = float32(playerHealth)
					healthBar.Resize(fyne.NewSize(healthBarWidth, healthBarHeight))
					healthBar.Refresh()
					healthLabel.SetText(fmt.Sprintf("%d", int(healthBarWidth)))
				}
				if healthBarWidth <= 0 {
					healthLabel.SetText("Du bist gestorben!")
					ticker.Stop()
					time.Sleep(2 * time.Second)
					w2.Close()
					closeDungeon()
					return
				}
				time.Sleep(2 * time.Second)
			}
		}()

		// Gegner-Logik
		go func() {
			for {
				allEnemiesDead := true
				for _, e := range enemies {
					distance := math.Sqrt(float64((x-e.X)*(x-e.X) + (y-e.Y)*(y-e.Y)))
					if distance < 50 && e.Health > 0 {
						e.Health -= 10 * swordAttack[inv.Sword]
						e.HealthBar.Resize(fyne.NewSize(float32(e.Health)/10, healthBarHeight))
						e.HealthBar.Refresh()
					}
					if e.Health > 0 {
						allEnemiesDead = false
					} else {
						e.Rect.FillColor = color.Gray{Y: 200}
						e.Rect.Refresh()
						e.NameLabel.SetText("☠")
					}
				}

				if allEnemiesDead {
					// Spieler hat gewonnen
					ticker.Stop()
					inv.Kristalle += 1 * (inv.Wiedergeburten + 1) * level

					// Spielerleben zurücksetzen
					playerHealth = 100 / 10 * armorDefense[inv.Armor]
					healthLabel.SetText(fmt.Sprintf("%d", playerHealth))
					healthLabel.SetText("Du hast Gewonnen!")

					time.Sleep(2 * time.Second)
					w2.Close()
					w.Show() // zurück zum Dungeon-Hauptfenster

					// Nächste Ebene als Button hinzufügen, aber nicht direkt starten
					*currentLevel++
					createLevel(a, inv, w, mainWindow, currentLevel, levelButtonContainer, healthLabel, playerHealth, closeDungeon)
					return
				}

				time.Sleep(1 * time.Second)
			}
		}()

		w2.Show()
	})

	// Button zum Dungeon-Fenster hinzufügen
	levelButtonContainer.Add(adventureBtn)
	levelButtonContainer.Refresh()
}
