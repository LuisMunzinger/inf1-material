package aufgabe3

import "fmt"

func ExampleNode_HasNeighbour() {
	n1 := NewNode("A")
	n1.AddNeighbours("B", "E", "F")
	fmt.Println(n1.HasNeighbour("B"))
	fmt.Println(n1.HasNeighbour("C"))

	// Output:
	// true
	// false
}
