package scheresteinpapier

import "fmt"

func ExampleValue_String() {
	rock := Rock
	paper := Paper
	scissors := Scissors
	spock := Spock
	lizard := Lizard

	fmt.Println(rock)
	fmt.Println(paper)
	fmt.Println(scissors)
	fmt.Println(spock)
	fmt.Println(lizard)

	// Output:
	// Stein
	// Papier
	// Schere
	// Spock
	// Eidechse
}

func ExampleValue_Beats() {
	fmt.Println(Rock.Beats(Rock))
	fmt.Println(Rock.Beats(Paper))
	fmt.Println(Rock.Beats(Scissors))
	fmt.Println(Rock.Beats(Spock))
	fmt.Println(Rock.Beats(Lizard))
	fmt.Println()

	fmt.Println(Paper.Beats(Rock))
	fmt.Println(Paper.Beats(Paper))
	fmt.Println(Paper.Beats(Scissors))
	fmt.Println(Paper.Beats(Spock))
	fmt.Println(Paper.Beats(Lizard))
	fmt.Println()

	fmt.Println(Scissors.Beats(Rock))
	fmt.Println(Scissors.Beats(Paper))
	fmt.Println(Scissors.Beats(Scissors))
	fmt.Println(Scissors.Beats(Spock))
	fmt.Println(Scissors.Beats(Lizard))
	fmt.Println()

	fmt.Println(Spock.Beats(Rock))
	fmt.Println(Spock.Beats(Paper))
	fmt.Println(Spock.Beats(Scissors))
	fmt.Println(Spock.Beats(Spock))
	fmt.Println(Spock.Beats(Lizard))
	fmt.Println()

	fmt.Println(Lizard.Beats(Rock))
	fmt.Println(Lizard.Beats(Paper))
	fmt.Println(Lizard.Beats(Scissors))
	fmt.Println(Lizard.Beats(Spock))
	fmt.Println(Lizard.Beats(Lizard))

	// Output:
	// false
	// false
	// true
	// false
	// true
	//
	// true
	// false
	// false
	// true
	// false
	//
	// false
	// true
	// false
	// false
	// true
	//
	// true
	// false
	// true
	// false
	// false
	//
	// false
	// true
	// false
	// true
	// false
}
