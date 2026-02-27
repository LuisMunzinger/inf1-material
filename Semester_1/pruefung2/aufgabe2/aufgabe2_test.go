package aufgabe2

import "fmt"

func ExampleFilterDigits() {

	fmt.Println(FilterDigits("a123b"))
	fmt.Println(FilterDigits("abcde"))
	fmt.Println(FilterDigits("0123456789"))
	fmt.Println(FilterDigits("1a9b"))
	fmt.Println(FilterDigits(""))

	// Output:
	// ab
	// abcde
	//
	// ab
	//
}
