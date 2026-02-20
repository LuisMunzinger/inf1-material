package main

import (
	"fmt"
	"inf1-material/Vorlesungsprogramme/26.02.20/hanoi_v2"
)

func main() {
	hanoi_v2.Hanoi(21, "A", "B", "C")

	fmt.Printf("Anzahl der Moves: %d", hanoi_v2.Counter)
}
