package GuessingGame

import "fmt"
import "math/rand/v2"

func GuessingGame() {
	correct := rand.IntN(100) 
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if NumberGood(guess, correct) {
			fmt.Println("Richtig geraten! :-)")
			return
		} 
	}
	fmt.Printf("Zu viele falsche Zahlen!",correct)
}