package aufgabe1

import (
	"fmt"
)

func ExamplePrefixBelow10() {
	l1 := []int{1, 7, 3, 2, 5, 10, 3, 11, 4, 6}
	l2 := []int{10, 2, 4, 6}
	l3 := []int{2, 4, 6}
	l4 := []int{10, 11, 12, 13, 14, 15}
	l5 := []int{}

	fmt.Println(PrefixBelow10(l1))
	fmt.Println(PrefixBelow10(l2))
	fmt.Println(PrefixBelow10(l3))
	fmt.Println(PrefixBelow10(l4))
	fmt.Println(PrefixBelow10(l5))

	// Output:
	// [1 7 3 2 5]
	// []
	// [2 4 6]
	// []
	// []
}
