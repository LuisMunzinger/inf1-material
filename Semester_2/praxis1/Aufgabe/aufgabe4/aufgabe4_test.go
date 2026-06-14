package aufgabe4

import "fmt"

func ExampleNode_ReachableNodes_short_nocycle() {
	a1 := NewNode("a")
	a2 := NewNode("a")
	b1 := NewNode("b")
	b2 := NewNode("b")

	a1.AddNeighbourNode(b1)
	b1.AddNeighbourNode(a2)
	b1.AddNeighbourNode(b2)

	PrintNodePtrList(a1.ReachableNodes())
	PrintNodePtrList(a2.ReachableNodes())
	PrintNodePtrList(b1.ReachableNodes())
	PrintNodePtrList(b2.ReachableNodes())

	// Output:
	// [a b b a]
	// [a]
	// [b b a]
	// [b]
}

func ExampleNode_ReachableNodes_long_nocycle() {
	a1 := NewNode("a")
	a2 := NewNode("a")
	b1 := NewNode("b")
	b2 := NewNode("b")
	b3 := NewNode("b")
	c1 := NewNode("c")
	c2 := NewNode("c")

	a1.AddNeighbourNode(b1)
	b1.AddNeighbourNode(c1)
	b1.AddNeighbourNode(c2)
	c1.AddNeighbourNode(a2)
	c2.AddNeighbourNode(b2)
	c2.AddNeighbourNode(b3)

	PrintNodePtrList(a1.ReachableNodes())
	PrintNodePtrList(a2.ReachableNodes())
	PrintNodePtrList(b1.ReachableNodes())

	// Output:
	// [a b c b b c a]
	// [a]
	// [b c b b c a]
}

func ExampleNode_ReachableNodes_short_cycle() {
	a1 := NewNode("a")
	a2 := NewNode("a")
	b1 := NewNode("b")

	a1.AddNeighbourNode(a2)
	a2.AddNeighbourNode(b1)
	b1.AddNeighbourNode(a1)

	PrintNodePtrList(a1.ReachableNodes())
	PrintNodePtrList(a2.ReachableNodes())
	PrintNodePtrList(b1.ReachableNodes())

	// Output:
	// [a a b]
	// [a b a]
	// [b a a]
}

/* Hilfsfunktion für die Tests */
func PrintNodePtrList(list []*Node) {
	labels := []string{}
	for _, n := range list {
		labels = append(labels, n.Label)
	}
	fmt.Println(labels)
}
