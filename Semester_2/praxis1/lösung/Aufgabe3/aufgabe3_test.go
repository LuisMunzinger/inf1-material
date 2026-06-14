package aufgabe3

import "fmt"

func ExampleNode_FirstNeighbourStartingWith() {
	n1 := NewNode("1")
	n1.AddNeighbours("aabc", "abc", "aa", "ab", "cde", "ba")
	fmt.Println(n1.FirstNeighbourStartingWith("c").Label)
	fmt.Println(n1.FirstNeighbourStartingWith("ab").Label)

	n2 := NewNode("2")
	n2.AddNeighbours("a", "b", "c")
	fmt.Println(n2.FirstNeighbourStartingWith("a").Label)
	fmt.Println(n2.FirstNeighbourStartingWith("d"))

	// Output:
	// cde
	// abc
	// a
	// <nil>
}

func ExampleNode_FirstNeighbourStartingWith_empty() {
	n3 := NewNode("3")
	fmt.Println(n3.FirstNeighbourStartingWith("a"))

	// Output:
	// <nil>
}
