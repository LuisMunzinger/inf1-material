package database

import (
	"fmt"
)

func ExampleDatabase_GetDirectContactIds() {
	db := testdata()

	fmt.Println(db.GetDirectContactIds("1"))
	fmt.Println(db.GetDirectContactIds("2"))
	fmt.Println(db.GetDirectContactIds("3"))

	// Output:
	// [3]
	// [1]
	// [1 2]
}

func ExampleDatabase_GetContactIds_depth_1() {
	db := testdata()

	fmt.Println(db.GetContactIds("1", 1))
	fmt.Println(db.GetContactIds("2", 1))
	fmt.Println(db.GetContactIds("3", 1))

	// Output:
	// [3]
	// [1]
	// [1 2]
}

func ExampleDatabase_GetContactIds_depth_2() {
	db := testdata()

	fmt.Println(db.GetContactIds("1", 2))
	fmt.Println(db.GetContactIds("2", 2))
	fmt.Println(db.GetContactIds("3", 2))

	// Output:
	// [3 2]
	// [1 3]
	// [1 2]
}

func ExampleDatabase_GetContactIds_depth_0() {
	db := testdata()

	fmt.Println(db.GetContactIds("1", 0))
	fmt.Println(db.GetContactIds("2", 0))
	fmt.Println(db.GetContactIds("3", 0))

	// Output:
	// []
	// []
	// []
}
