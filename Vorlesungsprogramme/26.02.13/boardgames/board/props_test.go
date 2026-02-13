package board

import "fmt"

func ExampleBoard_maxCellWidth() {
	b1 := New(3, 3)
	b2 := NewWithNumbers(3, 4)

	fmt.Println(b1.MaxCellWidth()) // 1 (Leerzeichen)
	fmt.Println(b2.MaxCellWidth()) // 2 (Zahl "10" hat 2 Zeichen)

	// Output:
	// 1
	// 2
}

func ExampleBoard_Width() {
	b1 := New(3, 3)
	b2 := NewWithNumbers(3, 4)

	fmt.Println(b1.Width())
	fmt.Println(b2.Width())

	// Output:
	// 3
	// 4
}

func ExampleBoard_Height() {
	b1 := New(3, 3)
	b2 := NewWithNumbers(3, 4)

	fmt.Println(b1.Height())
	fmt.Println(b2.Height())

	// Output:
	// 3
	// 3
}

func ExampleBoard_Row() {
	b1 := NewWithNumbers(3, 3)
	fmt.Println(b1.Row(0)) // ["1" "2" "3"]
	fmt.Println(b1.Row(1)) // ["4" "5" "6"]
	fmt.Println(b1.Row(2)) // ["7" "8" "9"]

	b2 := NewWithNumbers(2, 4)
	fmt.Println(b2.Row(0)) // ["1" "2" "3" "4"]
	fmt.Println(b2.Row(1)) // ["5" "6" "7" "8"]

	// Output:
	// [1 2 3]
	// [4 5 6]
	// [7 8 9]
	// [1 2 3 4]
	// [5 6 7 8]
}

func ExampleBoard_Col() {
	b1 := NewWithNumbers(3, 3)
	fmt.Println(b1.Col(0)) // ["1" "4" "7"]
	fmt.Println(b1.Col(1)) // ["2" "5" "8"]
	fmt.Println(b1.Col(2)) // ["3" "6" "9"]

	b2 := NewWithNumbers(2, 4)
	fmt.Println(b2.Col(0)) // ["1" "5"]
	fmt.Println(b2.Col(1)) // ["2" "6"]
	fmt.Println(b2.Col(2)) // ["3" "7"]
	fmt.Println(b2.Col(3)) // ["4" "8"]

	// Output:
	// [1 4 7]
	// [2 5 8]
	// [3 6 9]
	// [1 5]
	// [2 6]
	// [3 7]
	// [4 8]
}

func ExampleBoard_DiagDownRight() {

	// ["1" "2" "3"]
	// ["4" "5" "6"]
	// ["7" "8" "9"]
	b1 := NewWithNumbers(3, 3)
	fmt.Println(b1.DiagDownRight(0))  // ["1" "5" "9"]
	fmt.Println(b1.DiagDownRight(1))  // ["2" "6"]
	fmt.Println(b1.DiagDownRight(2))  // ["3"]
	fmt.Println(b1.DiagDownRight(3))  // []
	fmt.Println(b1.DiagDownRight(-1)) // ["4" "8"]
	fmt.Println(b1.DiagDownRight(-2)) // ["7"]
	fmt.Println(b1.DiagDownRight(-3)) // []

	fmt.Println()

	// ["1" "2" "3" "4"]
	// ["5" "6" "7" "8"]
	b2 := NewWithNumbers(2, 4)
	fmt.Println(b2.DiagDownRight(0))  // ["1" "6"]
	fmt.Println(b2.DiagDownRight(1))  // ["2" "7"]
	fmt.Println(b2.DiagDownRight(2))  // ["3" "8"]
	fmt.Println(b2.DiagDownRight(3))  // ["4"]
	fmt.Println(b2.DiagDownRight(4))  // []
	fmt.Println(b2.DiagDownRight(-1)) // ["5"]
	fmt.Println(b2.DiagDownRight(-2)) // []

	// Output:
	// [1 5 9]
	// [2 6]
	// [3]
	// []
	// [4 8]
	// [7]
	// []
	//
	// [1 6]
	// [2 7]
	// [3 8]
	// [4]
	// []
	// [5]
	// []
}

func ExampleBoard_DiagDownLeft() {

	// ["1" "2" "3"]
	// ["4" "5" "6"]
	// ["7" "8" "9"]
	b1 := NewWithNumbers(3, 3)
	fmt.Println(b1.DiagDownLeft(2))  // ["3" "5" "7"]
	fmt.Println(b1.DiagDownLeft(1))  // ["2" "4"]
	fmt.Println(b1.DiagDownLeft(0))  // ["1"]
	fmt.Println(b1.DiagDownLeft(-1)) // []
	fmt.Println(b1.DiagDownLeft(3))  // ["6" "8"]
	fmt.Println(b1.DiagDownLeft(4))  // ["9"]
	fmt.Println(b1.DiagDownLeft(5))  // []

	fmt.Println()

	// ["1" "2" "3" "4"]
	// ["5" "6" "7" "8"]
	b2 := NewWithNumbers(2, 4)
	fmt.Println(b2.DiagDownLeft(3))  // ["4" "7"]
	fmt.Println(b2.DiagDownLeft(2))  // ["3" "6"]
	fmt.Println(b2.DiagDownLeft(1))  // ["2" "5"]
	fmt.Println(b2.DiagDownLeft(0))  // ["1"]
	fmt.Println(b2.DiagDownLeft(-1)) // []
	fmt.Println(b2.DiagDownLeft(4))  // ["8"]
	fmt.Println(b2.DiagDownLeft(5))  // []

	// Output:
	// [3 5 7]
	// [2 4]
	// [1]
	// []
	// [6 8]
	// [9]
	// []
	//
	// [4 7]
	// [3 6]
	// [2 5]
	// [1]
	// []
	// [8]
	// []
}
