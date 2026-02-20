package Chain

import "fmt"

func ExampleChain() {
	fmt.Println(Chain("abc", 3))
	fmt.Println(Chain("a", 2))
	fmt.Println(Chain("", 2))

	// Output:
	// abcabcabc
	// aa
	//
}
