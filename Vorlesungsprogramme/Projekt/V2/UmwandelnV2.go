package main

import "fmt"

func umwandeln(Zahl, R int) {
	result := []int{}
	for Zahl != 0 {
		if Zahl%R == 0 {
			result = append([]int{0}, result...)
			Zahl = Zahl / R
		}
		if Zahl%R != 0 {
			result = append([]int{1}, result...)
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
