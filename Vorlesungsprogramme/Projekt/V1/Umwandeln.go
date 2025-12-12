package main

import "fmt"

func umwandeln(Zahl, R int) {

}

func main() {
	var Zahl int
	result := []int{}
	fmt.Print("Wählen sie eine Zahl zum umwandeln: ")
	fmt.Scan(&Zahl)

	for Zahl != 0 {
		if Zahl%2 == 0 {
			result = append([]int{0}, result...)
			Zahl = Zahl / 2
		}
		if Zahl%2 != 0 {
			result = append([]int{1}, result...)
			Zahl = (Zahl - 1) / 2
		}
	}
	fmt.Println(result)

}
