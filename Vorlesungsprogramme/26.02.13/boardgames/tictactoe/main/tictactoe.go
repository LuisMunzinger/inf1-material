package main

import (
	"fmt"
	"inf1-material/Vorlesungsprogramme/26.02.13/boardgames/tictactoe"
)

func main() {
	g := tictactoe.NewGame()
	p := "X"

	for !g.GameOver() {
		g.Show()
		g.Move(p)

		if p == "X" {
			p = "O"
		} else {
			p = "X"
		}
	}

	g.Show()
	if g.PlayerWins("X") {
		fmt.Println("Spieler X gewinnt!")
	}
	if g.PlayerWins("O") {
		fmt.Println("Spieler O gewinnt!")
	}
	if g.IsDraw() {
		fmt.Println("Unentschieden!")
	}
}
