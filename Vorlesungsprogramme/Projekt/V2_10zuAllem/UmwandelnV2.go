package main

import (
	"fmt"
	"strconv"
)

//import "math"

func umwandeln(Zahl, Systeme int64) {
	fmt.Println("Ihre zahl ist:  ", strconv.FormatInt(Zahl, int(Systeme)))
}

func main() {
	var Zahl int64
	var Systeme int64
	fmt.Print("Wählen sie eine Zahl zum umwandeln:  ")
	fmt.Scan(&Zahl)

	fmt.Print("Wählen sie ein rechensystem:  ")
	fmt.Scan(&Systeme)

	umwandeln(Zahl, Systeme)
}
