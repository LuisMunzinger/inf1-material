package element

import "fmt"

func ExampleElement_IsEmpty() {
	e1 := Empty()
	e2 := New(42)

	fmt.Println(e1.IsEmpty())
	fmt.Println(e2.IsEmpty())
	fmt.Print(e2.Next().IsEmpty())

	// Output:
	// true
	// false
	// true
}

func ExampleElement_Value_nonempty() {
	e := New(42)

	fmt.Println(e.Value())

	// Output:
	// 42
}

func ExampleElement_Value_empty() {
	e := Empty()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println(e.Value())

	// Output:
	// value for empty element requested
}

func ExampleElement_Next_nonempty() {
	e1 := New(42)
	e1.next = New(23)

	fmt.Println(e1.Next().value)
	fmt.Println(e1.Next().IsEmpty())

	// Output:
	// 23
	// false
}

func ExampleElement_SetValue() {
	e := Empty()

	e.SetValue(42)
	fmt.Println(e.Value())
	fmt.Println(e.IsEmpty())

	e.SetValue(23)
	fmt.Println(e.Value())
	fmt.Println(e.IsEmpty())

	// Output:
	// 42
	// false
	// 23
	// false
}
