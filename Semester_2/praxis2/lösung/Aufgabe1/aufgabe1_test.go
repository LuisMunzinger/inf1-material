package aufgabe1

import "fmt"

func ExampleNode_Sum() {
	n := NewNode()
	fmt.Println(n.Sum())

	n.PushBack(1)
	fmt.Println(n.Sum())

	n.PushBack(2)
	n.PushBack(3)
	n.PushBack(4)
	n.PushBack(5)
	fmt.Println(n.Sum())

	n.PushBack(6)
	fmt.Println(n.Sum())

	n.PushBack(7)
	fmt.Println(n.Sum())

	// Output:
	// 0
	// 1
	// 15
	// 21
	// 28
}
