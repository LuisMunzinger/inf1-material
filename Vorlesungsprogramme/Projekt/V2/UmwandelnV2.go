package main

import "fmt"

func umwandeln(Zahl, R int) {
	result := []int{}
	for Zahl != 0 {
		if Zahl%R == 0 {
			Z := Zahl / R
			result = append([]int{Z}, result...)
			Zahl = Zahl / R
		}
		if Zahl%R != 0 {
			Z := Zahl / R
			result = append([]int{Z}, result...)
			Zahl = (Zahl - 1) / R
		}
	}
	fmt.Println(result)
}

func main() {
	var Zahl int
	var R int
	fmt.Print("Wählen sie eine Zahl zum umwandeln: ")
	fmt.Scan(&Zahl)

	fmt.Println("Wählen sie ein rechensystem")
	fmt.Scan(&R)

	umwandeln(Zahl, R)
}
