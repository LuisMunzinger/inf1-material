package list

import (
	"fmt"
)

func ExampleEmpty() {
	l := Empty()

	fmt.Println(l.IsEmpty())
	fmt.Println(l)

	// Output:
	// true
	// []
}

func ExampleList_Append() {
	l := Empty()

	l.Append(42)
	l.Append(23)

	fmt.Println(l.IsEmpty())
	fmt.Println(l)

	// Output:
	// false
	// [42 -> 23]
}

func ExampleList_Length() {
	l := Empty()

	fmt.Println(l.Length())

	l.Append(42)
	fmt.Println(l.Length())

	l.Append(23)
	fmt.Println(l.Length())

	// Output:
	// 0
	// 1
	// 2
}

func ExampleList_Contains() {
	l := Empty()

	fmt.Println(l.Contains(42))

	l.Append(42)
	fmt.Println(l.Contains(42))
	fmt.Println(l.Contains(23))

	l.Append(23)
	fmt.Println(l.Contains(42))
	fmt.Println(l.Contains(23))

	// Output:
	// false
	// true
	// false
	// true
	// true
}

func ExampleList_Swap() {
	l := Empty()

	l.Append(42)
	l.Append(23)
	l.Append(17)
	l.Append(5)
	l.Append(11)
	l.Append(3)

	fmt.Println(l)

	l.Swap(0, 5)
	fmt.Println(l)

	l.Swap(1, 4)
	fmt.Println(l)

	l.Swap(2, 3)
	fmt.Println(l)

	// Output:
	// [42 -> 23 -> 17 -> 5 -> 11 -> 3]
	// [3 -> 23 -> 17 -> 5 -> 11 -> 42]
	// [3 -> 11 -> 17 -> 5 -> 23 -> 42]
	// [3 -> 11 -> 5 -> 17 -> 23 -> 42]
}
