package entry

import "fmt"

func ExampleNew() {
	e1 := New("Apfel", "apple")

	fmt.Println(e1.de)
	fmt.Println(e1.en)

	// Output:
	// Apfel
	// apple
}

func ExampleEmpty() {
	e1 := Empty()

	fmt.Println(e1.de)
	fmt.Println(e1.en)

	// Output:
	//
	//
}

func ExampleEntry_De() {
	e1 := New("Apfel", "apple")

	fmt.Println(e1.De())

	// Output:
	// Apfel
}

func ExampleEntry_En() {
	e1 := New("Apfel", "apple")

	fmt.Println(e1.En())

	// Output:
	// apple
}

func ExampleEntry_IsValid() {
	e1 := New("Apfel", "apple")
	e2 := New("", "apple")
	e3 := New("Apfel", "")
	e4 := New("", "")
	e5 := Empty()

	fmt.Println(e1.IsValid())
	fmt.Println(e2.IsValid())
	fmt.Println(e3.IsValid())
	fmt.Println(e4.IsValid())
	fmt.Println(e5.IsValid())

	// Output:
	// true
	// false
	// false
	// false
	// false
}
