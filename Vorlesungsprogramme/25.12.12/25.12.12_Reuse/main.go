package main

import "fmt"

func GetInput() int {
	fmt.Println("Gib eine Zahl Zwischen 1 und 10 ein: ")
	var n int
	fmt.Scan(&n)

	if n > 0 && n <= 10 {
		return n
	}
	fmt.Println("Ungültige Eingabe, versuch es noch einmal.")
	return GetInput()
}

func main() {
	n := GetInput()
	fmt.Printf("%v ist eine gute Zahl.", n)

}
