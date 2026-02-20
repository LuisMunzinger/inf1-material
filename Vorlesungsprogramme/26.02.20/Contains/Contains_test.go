package contains

import "fmt"

func ExampleContains() {
	fmt.Println(Contains("abaaba", "a"))
	fmt.Println(Contains("abaaba", "ab"))
	fmt.Println(Contains("abaaba", "aa"))
	fmt.Println(Contains("abaaba", ""))
	fmt.Println(Contains("abaaba", "baab"))
	fmt.Println(Contains("abaaba", "baaab"))

	// Output:
	// true
	// true
	// true
	// true
	// true
	// false
}
