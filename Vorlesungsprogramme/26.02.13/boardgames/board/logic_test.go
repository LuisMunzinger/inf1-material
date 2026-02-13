package board

import "fmt"

func ExampleBoard_RowContains() {

	// ["1" "2" "3"]
	// ["4" "5" "6"]
	// ["7" "8" "9"]
	b1 := NewWithNumbers(3, 3)
	fmt.Println(b1.RowContains(0, "2"))
	fmt.Println(b1.RowContains(0, "5"))
	fmt.Println(b1.RowContains(1, "5"))
	fmt.Println(b1.RowContains(1, "8"))
	fmt.Println(b1.RowContains(2, "8"))
	fmt.Println(b1.RowContains(2, "2"))

	// Output:
	// true
	// false
	// true
	// false
	// true
	// false
}

func ExampleBoard_RowContainsChain() {

	b := Board{
		rows: [][]string{
			{"*", "*", "*"},
			{"*", " ", "*"},
			{" ", " ", " "},
		},
	}

	fmt.Println(b.RowContainsChain(0, "*", 3))
	fmt.Println(b.RowContainsChain(0, "*", 2))
	fmt.Println(b.RowContainsChain(0, "*", 1))
	fmt.Println(b.RowContainsChain(0, "*", 0))
	fmt.Println(b.RowContainsChain(1, "*", 2))
	fmt.Println(b.RowContainsChain(1, "*", 1))

	// Output:
	// true
	// true
	// true
	// true
	// false
	// true
}

func ExampleBoard_RowContainsOnly() {

	b := Board{
		rows: [][]string{
			{"*", "*", "*"},
			{"*", " ", "*"},
			{" ", " ", " "},
		},
	}

	fmt.Println(b.RowContainsOnly(0, "*"))
	fmt.Println(b.RowContainsOnly(1, "*"))
	fmt.Println(b.RowContainsOnly(2, "*"))
	fmt.Println(b.RowContainsOnly(2, " "))

	// Output:
	// true
	// false
	// false
	// true

}

func ExampleBoard_ColContainsChain() {
	b := Board{
		rows: [][]string{
			{"*", "*", "*"},
			{"*", " ", "*"},
			{"*", " ", " "},
		},
	}

	fmt.Println(b.ColContainsChain(0, "*", 3))
	fmt.Println(b.ColContainsChain(0, "*", 2))
	fmt.Println(b.ColContainsChain(0, "*", 1))
	fmt.Println(b.ColContainsChain(0, "*", 0))
	fmt.Println(b.ColContainsChain(1, "*", 2))
	fmt.Println(b.ColContainsChain(1, "*", 1))

	// Output:
	// true
	// true
	// true
	// true
	// false
	// true
}

func ExampleBoard_ColContains() {

	// ["1" "2" "3"]
	// ["4" "5" "6"]
	// ["7" "8" "9"]
	b1 := NewWithNumbers(3, 3)
	fmt.Println(b1.ColContains(0, "4"))
	fmt.Println(b1.ColContains(0, "5"))
	fmt.Println(b1.ColContains(1, "5"))
	fmt.Println(b1.ColContains(1, "6"))
	fmt.Println(b1.ColContains(2, "6"))
	fmt.Println(b1.ColContains(2, "4"))

	// Output:
	// true
	// false
	// true
	// false
	// true
	// false
}

func ExampleBoard_ColContainsOnly() {

	b := Board{
		rows: [][]string{
			{"*", "*", " "},
			{"*", " ", " "},
			{"*", " ", " "},
		},
	}

	fmt.Println(b.ColContainsOnly(0, "*"))
	fmt.Println(b.ColContainsOnly(1, "*"))
	fmt.Println(b.ColContainsOnly(2, "*"))
	fmt.Println(b.ColContainsOnly(2, " "))

	// Output:
	// true
	// false
	// false
	// true
}

func ExampleBoard_DiagDownRightContains() {

	// ["1" "2" "3"]
	// ["4" "5" "6"]
	// ["7" "8" "9"]
	b1 := NewWithNumbers(3, 3)
	fmt.Println(b1.DiagDownRightContains(0, "1"))
	fmt.Println(b1.DiagDownRightContains(0, "5"))
	fmt.Println(b1.DiagDownRightContains(0, "9"))
	fmt.Println(b1.DiagDownRightContains(0, "2"))
	fmt.Println(b1.DiagDownRightContains(1, "2"))
	fmt.Println(b1.DiagDownRightContains(1, "6"))
	fmt.Println(b1.DiagDownRightContains(2, "3"))
	fmt.Println(b1.DiagDownRightContains(2, "7"))

	// Output:
	// true
	// true
	// true
	// false
	// true
	// true
	// true
	// false
}

func ExampleBoard_DiagDownRightContainsChain() {

	b := Board{
		rows: [][]string{
			{"*", "*", "*"},
			{" ", "*", " "},
			{" ", " ", "*"},
		},
	}

	fmt.Println(b.DiagDownRightContainsChain(0, "*", 3))
	fmt.Println(b.DiagDownRightContainsChain(0, "*", 2))
	fmt.Println(b.DiagDownRightContainsChain(0, "*", 1))
	fmt.Println(b.DiagDownRightContainsChain(0, "*", 0))
	fmt.Println(b.DiagDownRightContainsChain(1, "*", 2))
	fmt.Println(b.DiagDownRightContainsChain(1, "*", 1))

	// Output:
	// true
	// true
	// true
	// true
	// false
	// true
}

func ExampleBoard_DiagDownRightContainsOnly() {

	b := Board{
		rows: [][]string{
			{"*", " ", " "},
			{" ", "*", " "},
			{" ", "*", "*"},
		},
	}

	fmt.Println(b.DiagDownRightContainsOnly(0, "*"))
	fmt.Println(b.DiagDownRightContainsOnly(1, "*"))
	fmt.Println(b.DiagDownRightContainsOnly(1, " "))
	fmt.Println(b.DiagDownRightContainsOnly(-1, "*"))
	fmt.Println(b.DiagDownRightContainsOnly(-1, " "))

	// Output:
	// true
	// false
	// true
	// false
	// false
}

func ExampleBoard_DiagDownLeftContains() {

	// ["1" "2" "3"]
	// ["4" "5" "6"]
	// ["7" "8" "9"]
	b1 := NewWithNumbers(3, 3)
	fmt.Println(b1.DiagDownLeftContains(2, "3"))
	fmt.Println(b1.DiagDownLeftContains(2, "5"))
	fmt.Println(b1.DiagDownLeftContains(2, "7"))
	fmt.Println(b1.DiagDownLeftContains(2, "2"))
	fmt.Println(b1.DiagDownLeftContains(1, "2"))
	fmt.Println(b1.DiagDownLeftContains(1, "4"))
	fmt.Println(b1.DiagDownLeftContains(0, "1"))
	fmt.Println(b1.DiagDownLeftContains(0, "5"))

	// Output:
	// true
	// true
	// true
	// false
	// true
	// true
	// true
	// false
}

func ExampleBoard_DiagDownLeftContainsChain() {

	b := Board{
		rows: [][]string{
			{" ", "*", "*"},
			{" ", "*", " "},
			{"*", " ", " "},
		},
	}

	fmt.Println(b.DiagDownLeftContainsChain(2, "*", 3))
	fmt.Println(b.DiagDownLeftContainsChain(2, "*", 2))
	fmt.Println(b.DiagDownLeftContainsChain(2, "*", 1))
	fmt.Println(b.DiagDownLeftContainsChain(2, "*", 0))
	fmt.Println(b.DiagDownLeftContainsChain(1, "*", 2))
	fmt.Println(b.DiagDownLeftContainsChain(1, "*", 1))

	// Output:
	// true
	// true
	// true
	// true
	// false
	// true
}

func ExampleBoard_DiagDownLeftContainsOnly() {

	b := Board{
		rows: [][]string{
			{" ", "*", "*"},
			{" ", "*", " "},
			{"*", " ", " "},
		},
	}

	fmt.Println(b.DiagDownLeftContainsOnly(2, "*"))
	fmt.Println(b.DiagDownLeftContainsOnly(1, "*"))
	fmt.Println(b.DiagDownLeftContainsOnly(1, " "))
	fmt.Println(b.DiagDownLeftContainsOnly(3, "*"))
	fmt.Println(b.DiagDownLeftContainsOnly(3, " "))

	// Output:
	// true
	// false
	// false
	// false
	// true
}

func ExampleBoard_BoardContains() {

	// ["1" "2" "3"]
	// ["4" "5" "6"]
	// ["7" "8" "9"]
	b1 := NewWithNumbers(3, 3)
	fmt.Println(b1.BoardContains("5"))
	fmt.Println(b1.BoardContains("0"))

	// Output:
	// true
	// false
}

func ExampleBoard_BoardContainsChain() {

	b := Board{
		rows: [][]string{
			{"*", "*", "*"},
			{" ", "*", " "},
			{" ", " ", "*"},
		},
	}

	fmt.Println(b.BoardContainsChain("*", 3))
	fmt.Println(b.BoardContainsChain("*", 4))
	fmt.Println(b.BoardContainsChain("*", 2))
	fmt.Println(b.BoardContainsChain("*", 1))
	fmt.Println(b.BoardContainsChain("*", 0))
	fmt.Println(b.BoardContainsChain(" ", 3))
	fmt.Println(b.BoardContainsChain(" ", 2))

	// Output:
	// true
	// false
	// true
	// true
	// true
	// false
	// true
}

func ExampleBoard_BoardContainsOnly() {

	b1 := Board{
		rows: [][]string{
			{"*", "*", "*"},
			{" ", "*", " "},
			{" ", " ", "*"},
		},
	}

	b2 := Board{
		rows: [][]string{
			{"*", "*", "*"},
			{"*", "*", "*"},
			{"*", "*", "*"},
		},
	}

	fmt.Println(b1.BoardContainsOnly("*"))
	fmt.Println(b1.BoardContainsOnly(" "))
	fmt.Println(b2.BoardContainsOnly("*"))
	fmt.Println(b2.BoardContainsOnly(" "))

	// Output:
	// false
	// false
	// true
	// false
}
