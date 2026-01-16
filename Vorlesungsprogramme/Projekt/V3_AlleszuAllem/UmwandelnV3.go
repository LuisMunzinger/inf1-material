package main

import (
	"fmt"
	"strconv"
)

func umwandeln(zZahl, vonSysteme, zuSysteme int64) {
	s := "1010"
	Zahl, _ := strconv.ParseInt(s, int(vonSysteme), 64)
	println("")
	fmt.Println("Ihre zahl ist:  ", strconv.FormatInt(Zahl, int(zuSysteme)))
	println("")
}

func main() {
	var zZahl int64
	var zuSysteme int64
	var vonSysteme int64
	fmt.Print("Wählen sie eine Systeme aus dem sie umwandeln wollen:  ")
	fmt.Scan(&vonSysteme)

	fmt.Print("Wählen sie eine Zahl zum umwandeln:  ")
	fmt.Scan(&zZahl)

	fmt.Print("Wählen sie ein rechensystem:  ")
	fmt.Scan(&zuSysteme)

	umwandeln(zZahl, vonSysteme, zuSysteme)
}
