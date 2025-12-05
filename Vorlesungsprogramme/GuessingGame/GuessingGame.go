package GuessingGame

import "fmt"

func GuessingGame() {
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if NumberGood(guess) {
			fmt.Println("Richtig geraten! :-)")
			return
		} 
	}
	fmt.Println("Zu viele falsche Zahlen! :-(")
}