package db

import (
	"inf1-material/Vorlesungsprogramme/26.05.08/game"
	"inf1-material/Vorlesungsprogramme/26.05.08/player"
)

func testdata() *GameDb {
	db := New(50, 2)

	// Testdaten erstellen
	g1 := game.New("The Legend of Zelda", "Action-Adventure")
	g2 := game.New("Minecraft", "Survival")
	g3 := game.New("Tomb Raider", "Action-Adventure")
	g4 := game.New("Valheim", "Survival")

	p1 := player.New("Alice")
	p2 := player.New("Bob")
	p3 := player.New("Charlie")
	p4 := player.New("Diana")

	// Spiele hinzufügen
	db.AddGame(g1)
	db.AddGame(g2)
	db.AddGame(g3)
	db.AddGame(g4)

	// Spieler hinzufügen
	db.AddPlayer(p1)
	db.AddPlayer(p2)
	db.AddPlayer(p3)
	db.AddPlayer(p4)

	// Spieler spielen Spiele.

	// - Alice spielt "Zelda" und "Minecraft" beide sehr oft
	//         aber nicht "Tomb Raider" und "Valheim".
	p1.PlayGame(g1, 100)
	p1.PlayGame(g2, 75)

	// - Bob spielt "Tomb Raider" und "Valheim" beide sehr oft.
	p2.PlayGame(g3, 55)
	p2.PlayGame(g4, 60)

	// - Charlie spielt "Zelda" und "Valheim" beide sehr oft.
	p3.PlayGame(g1, 80)
	p3.PlayGame(g4, 70)

	// - Diana spielt "Tomb Raider" und "Valheim", aber beide nur selten.
	p4.PlayGame(g3, 10)
	p4.PlayGame(g4, 15)

	// Erwartete Empfehlungen für Alice:
	//
	// - Alice bekommt "Valheim" empfohlen, da sie "Minecraft" oft gespielt hat,
	//   "Valheim" das gleiche Genre hat und es auch von mindestens 2
	//   Spielern (Bob und Charlie) für mindestens 50 Stunden gespielt wurde.
	// - Alice bekommt nicht "Tomb Raider" empfohlen, da dies zwar zu "Zelda"
	//   passen würde, aber nur von einem Spieler (Bob) genügend gespielt wurde.
	// - Alice bekommt weder "Zelda" noch "Minecraft" empfohlen,
	//   da sie diese Spiele bereits gespielt hat.

	return db
}
