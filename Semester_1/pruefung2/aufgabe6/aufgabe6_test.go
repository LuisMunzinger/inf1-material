package aufgabe6

import "fmt"

func ExampleDuplicateSinglets() {
	l1 := []int{1, 2, 3}
	l2 := []int{1, 2, 3, 1}
	l3 := []int{1, 2, 3, 1, 2, 3}

	fmt.Println(DuplicateSinglets(l1))
	fmt.Println(DuplicateSinglets(l2))
	fmt.Println(DuplicateSinglets(l3))

	// Output:
	// [1 1 2 2 3 3]
	// [1 2 2 3 3 1]
	// [1 2 3 1 2 3]
}
