package aufgabe4

import "fmt"

func ExampleElementSums() {
	l1 := []int{1, 1, 1, 1, 1}
	l2 := []int{5, 4, 3, 2, 1}
	l3 := []int{3, 2, 1}

	fmt.Println(ElementSums(l1, l2))
	fmt.Println(ElementSums(l1, l3))
	fmt.Println(ElementSums(l2, l3))

	// Output:
	// [6 5 4 3 2]
	// [4 3 2 2 2]
	// [8 6 4 3 2]
}
