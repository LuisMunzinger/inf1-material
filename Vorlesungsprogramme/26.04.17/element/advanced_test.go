package element

import "fmt"

func ExampleElement_Append() {
	e := Empty()

	e.Append(42)
	e.Append(23)

	fmt.Println(e.Value())
	fmt.Println(e.Next().Value())

	fmt.Println(e.IsEmpty())
	fmt.Println(e.Next().IsEmpty())
	fmt.Println(e.Next().Next().IsEmpty())

	// Output:
	// 42
	// 23
	// false
	// false
	// true
}

func ExampleElement_Length() {
	e := Empty()

	fmt.Println(e.Length())

	e.Append(42)
	fmt.Println(e.Length())

	e.Append(23)
	fmt.Println(e.Length())

	// Output:
	// 0
	// 1
	// 2
}

func ExampleElement_Contains() {
	e := Empty()

	fmt.Println(e.Contains(42))

	e.Append(42)
	fmt.Println(e.Contains(42))
	fmt.Println(e.Contains(23))

	e.Append(23)
	fmt.Println(e.Contains(42))
	fmt.Println(e.Contains(23))

	// Output:
	// false
	// true
	// false
	// true
	// true
}

func ExampleElement_Count() {
	e := Empty()

	e.Append(42)
	e.Append(23)
	e.Append(42)
	e.Append(15)

	fmt.Println(e.Count(42))
	fmt.Println(e.Count(23))
	fmt.Println(e.Count(15))
	fmt.Println(e.Count(99))

	// Output:
	// 2
	// 1
	// 1
	// 0
}

func ExampleElement_Sum() {
	e := Empty()

	e.Append(42)
	e.Append(23)
	e.Append(15)

	fmt.Println(e.Sum())

	// Output:
	// 80
}

func ExampleElement_Min() {
	e := Empty()

	e.Append(42)
	e.Append(15)
	e.Append(23)

	fmt.Println(e.Min())

	// Output:
	// 15
}

func ExampleElement_Last() {
	e := Empty()

	e.Append(42)
	e.Append(15)
	e.Append(23)

	fmt.Println(e.Last().Value())

	// Output:
	// 23
}

func ExampleElement_At_existing() {
	e := Empty()

	e.Append(42)
	e.Append(15)
	e.Append(23)

	fmt.Println(e.At(0).Value())
	fmt.Println(e.At(1).Value())
	fmt.Println(e.At(2).Value())

	// Output:
	// 42
	// 15
	// 23
}

func ExampleElement_At_nonexisting() {
	e := Empty()

	e.Append(42)
	e.Append(15)
	e.Append(23)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println(e.At(3).Value())

	// Output:
	// index out of bounds
}

func ExampleElement_String() {
	e := Empty()

	fmt.Println(e)

	e.Append(42)
	e.Append(15)
	e.Append(23)

	fmt.Println(e)

	// Output:
	//
	// 42 -> 15 -> 23
}

func ExampleElement_Swap_regular() {
	e := Empty()

	e.Append(10)
	e.Append(20)
	e.Append(30)
	e.Append(40)
	e.Append(50)
	e.Append(60)
	e.Append(70)
	fmt.Println(e)

	e = e.Swap(1, 5)
	fmt.Println(e)

	// Output:
	// 10 -> 20 -> 30 -> 40 -> 50 -> 60 -> 70
	// 10 -> 60 -> 30 -> 40 -> 50 -> 20 -> 70
}

func ExampleElement_Swap_neighbours() {
	e := Empty()

	e.Append(10)
	e.Append(20)
	e.Append(30)
	e.Append(40)
	fmt.Println(e)

	e = e.Swap(1, 2)
	fmt.Println(e)

	// Output:
	// 10 -> 20 -> 30 -> 40
	// 10 -> 30 -> 20 -> 40
}

func ExampleElement_Swap_zero() {
	e := Empty()

	e.Append(10)
	e.Append(20)
	e.Append(30)
	fmt.Println(e)

	e = e.Swap(0, 2)
	fmt.Println(e)

	// Output:
	// 10 -> 20 -> 30
	// 30 -> 20 -> 10
}
