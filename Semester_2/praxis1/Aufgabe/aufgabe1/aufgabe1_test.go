package aufgabe1

import "fmt"

func ExampleNode_LengthGreater5() {
	n := NewNode()
	fmt.Println(n.LengthGreater5())

	n.PushBack(1)
	fmt.Println(n.LengthGreater5())

	n.PushBack(2)
	n.PushBack(3)
	n.PushBack(4)
	n.PushBack(5)
	fmt.Println(n.LengthGreater5())

	n.PushBack(6)
	fmt.Println(n.LengthGreater5())

	n.PushBack(7)
	fmt.Println(n.LengthGreater5())

	// Output:
	// false
	// false
	// false
	// true
	// true
}
