package game

import "fmt"

func Run() {
	for i := 0; i < 3; i++ {
		userGuess := ReadNumber()
		if NumberGood(userGuess) {
			fmt.Println("Richtige Antwort")
			return
		}
		fmt.Println("Falsche Antwort")
	}
	fmt.Println("Zu viele falsche Antworten, Game Over!")
}

func ReadNumber() int {
	var n int
	fmt.Print("Rate eine Zahl: ")
	fmt.Scan(&n)
	return n
}
