package board

import "fmt"

func ExampleNew() {
	b1 := New(3, 3)
	b2 := New(2, 4)

	fmt.Println(b1.rows)
	fmt.Println(b2.rows)

	// Output:
	// [[     ] [     ] [     ]]
	// [[       ] [       ]]
}

func ExampleNewWithNumbers() {
	b1 := NewWithNumbers(3, 3)
	b2 := NewWithNumbers(2, 4)

	fmt.Println(b1.rows)
	fmt.Println(b2.rows)

	// Output:
	// [[1 2 3] [4 5 6] [7 8 9]]
	// [[1 2 3 4] [5 6 7 8]]
}

func ExampleBoard_Get() {

	b := Board{
		rows: [][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		},
	}

	fmt.Println(b.Get(0, 0))
	fmt.Println(b.Get(0, 1))
	fmt.Println(b.Get(0, 2))
	fmt.Println(b.Get(1, 0))
	fmt.Println(b.Get(1, 1))
	fmt.Println(b.Get(1, 2))
	fmt.Println(b.Get(2, 0))
	fmt.Println(b.Get(2, 1))
	fmt.Println(b.Get(2, 2))
	fmt.Println(b.Get(3, 2) == "")
	fmt.Println(b.Get(2, 3) == "")

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// true
	// true
}

func ExampleBoard_Set() {

	b := New(3, 3)

	b.Set(0, 0, "1")
	b.Set(0, 1, "2")
	b.Set(0, 2, "3")
	b.Set(1, 0, "4")
	b.Set(1, 1, "5")
	b.Set(1, 2, "6")
	b.Set(2, 0, "7")
	b.Set(2, 1, "8")
	b.Set(2, 2, "9")
	b.Set(2, 3, "10")

	fmt.Println(b == NewWithNumbers(3, 3))
}
