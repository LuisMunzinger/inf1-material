package database

import (
	"fmt"
)

func ExampleDatabase_UsersById() {
	db := testdata()
	users := db.UsersById()

	fmt.Println(users)

	// Output:
	// [1 2 3]
}

func ExampleDatabase_UsersByNickname() {
	db := testdata()
	users := db.UsersByNickname()

	fmt.Println(users)

	// Output:
	// [michkenntjeder superstar123 zombie43]
}

func ExampleDatabase_UsersByName() {
	db := testdata()
	users := db.UsersByName()

	fmt.Println(users)

	// Output:
	// [Arno Eva Max]
}

func ExampleDatabase_UsersBySurname() {
	db := testdata()
	users := db.UsersBySurname()

	fmt.Println(users)

	// Output:
	// [Einbildung Mustermann Nym]
}
