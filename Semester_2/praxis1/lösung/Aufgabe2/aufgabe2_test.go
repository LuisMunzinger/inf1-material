package aufgabe2

import "fmt"

// Test für das Entfernen von gültigen Positionen, aber nicht am Anfang der Liste.
func ExampleNode_RemoveAt_valid_notfirst() {
	n := NewNode() // []
	PrintValuesAndLengths(n)

	n.PushBack(42)
	n.PushBack(25)
	n.PushBack(17)
	n.PushBack(53)
	n.PushBack(38)
	PrintValuesAndLengths(n) // [42 25 17 53 38]

	n.RemoveAt(2) // [42 25 53 38]
	PrintValuesAndLengths(n)

	n.RemoveAt(1) // [42 53 38]
	PrintValuesAndLengths(n)

	// Output:
	// Values: []
	// Lengths: []
	//
	// Values: [42 25 17 53 38]
	// Lengths: [5 4 3 2 1]
	//
	// Values: [42 25 53 38]
	// Lengths: [4 3 2 1]
	//
	// Values: [42 53 38]
	// Lengths: [3 2 1]
}

// Test für das Entfernen von ungültigen Positionen.
func ExampleNode_RemoveAt_invalid() {
	n := NewNode() // []
	n.PushBack(42)
	n.PushBack(38) // [42 38]

	n.RemoveAt(-1) // [42 38]
	PrintValuesAndLengths(n)

	n.RemoveAt(3) // [42 38]
	PrintValuesAndLengths(n)

	// Output:
	// Values: [42 38]
	// Lengths: [2 1]
	//
	// Values: [42 38]
	// Lengths: [2 1]
}

// Test für das Entfernen der ersten Position.
func ExampleNode_RemoveAt_valid_first() {
	n := NewNode() // []
	n.PushBack(42)
	n.PushBack(38) // [42 38]

	n.RemoveAt(0) // [38]
	PrintValuesAndLengths(n)

	// Output:
	// Values: [38]
	// Lengths: [1]
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
