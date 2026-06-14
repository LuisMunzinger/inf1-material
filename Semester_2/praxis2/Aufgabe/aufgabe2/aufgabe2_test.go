package aufgabe2

import "fmt"

func ExampleNode_RemoveAll() {
	n := NewNode() // []
	PrintValuesAndLengths(n)

	n.PushBack(42)
	n.PushBack(25)
	n.PushBack(17)
	n.PushBack(53)
	n.PushBack(42)
	n.PushBack(38)
	PrintValuesAndLengths(n) // [42 25 17 53 42 38]

	n = n.RemoveAll(25) // [42 17 53 42 38]
	PrintValuesAndLengths(n)

	n = n.RemoveAll(42) //  [17 53 38]
	PrintValuesAndLengths(n)

	n = n.RemoveAll(13) //  [17 53 38]
	PrintValuesAndLengths(n)

	// Output:
	// Values: []
	// Lengths: []
	//
	// Values: [42 25 17 53 42 38]
	// Lengths: [6 5 4 3 2 1]
	//
	// Values: [42 17 53 42 38]
	// Lengths: [5 4 3 2 1]
	//
	// Values: [17 53 38]
	// Lengths: [3 2 1]
	//
	// Values: [17 53 38]
	// Lengths: [3 2 1]
}

/* Hilfsfunktion für den Test */
func PrintValuesAndLengths(n *Node) {
	values := []int{}
	lengths := []int{}
	for !n.IsEmpty() {
		values = append(values, n.Value)
		lengths = append(lengths, n.Length)
		n = n.Next
	}
	fmt.Printf("Values: %v\nLengths: %v\n\n", values, lengths)
}
