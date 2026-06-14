package aufgabe4

import (
	"fmt"
	"slices"
)

func ExampleNode_NodesWithDistance() {
	a := NewNode("A")
	b := NewNode("B")
	c := NewNode("C")
	d := NewNode("D")

	// Distanz 1: [B C]
	a.AddNeighbourNode(b)
	a.AddNeighbourNode(c)

	// Distanz 2: [A C D]
	b.AddNeighbourNode(a)
	b.AddNeighbourNode(d)
	c.AddNeighbourNode(a)
	c.AddNeighbourNode(c)

	// Distanz 3: [A B C D]
	d.AddNeighbourNode(b)
	d.AddNeighbourNode(c)

	PrintNodeLabels(a.NodesWithDistance(0))
	PrintNodeLabels(a.NodesWithDistance(1))
	PrintNodeLabels(a.NodesWithDistance(2))
	PrintNodeLabels(a.NodesWithDistance(3))

	// Output:
	// [A]
	// [B C]
	// [A C D]
	// [A B C]
}

// Hilfsfunktion für die Tests
// Gibt eine sortierte Liste der Labels der Knoten in list aus.
func PrintNodeLabels(list []*Node) {
	labels := []string{}
	for _, n := range list {
		labels = append(labels, n.Label)
	}
	slices.Sort(labels)
	fmt.Println(labels)
}
