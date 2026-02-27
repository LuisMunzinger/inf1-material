package aufgabe3

import "fmt"

func ExampleCountSquares() {
	l1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	l2 := []int{1, 4, 9, 16}
	l3 := []int{2, 3, 5}
	l4 := []int{}

	fmt.Println(CountSquares(l1))
	fmt.Println(CountSquares(l2))
	fmt.Println(CountSquares(l3))
	fmt.Println(CountSquares(l4))

	// Output:
	// 3
	// 4
	// 0
	// 0
}
