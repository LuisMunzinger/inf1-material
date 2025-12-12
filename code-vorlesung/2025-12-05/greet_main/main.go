package main

import (
	"fmt"
	"inf1-material/code-vorlesung/2025-12-05/greet"
)

func main() {
	fmt.Print("Wie heißt du? ")
	var n string
	fmt.Scanln(&n)

	s1 := greet.Greet(n)

	fmt.Println(s1)
}
