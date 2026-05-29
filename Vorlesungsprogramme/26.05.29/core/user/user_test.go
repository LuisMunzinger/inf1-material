package user

import "fmt"

func ExampleNew() {
	u := New("1", "michkenntjeder", "Max", "Mustermann")

	fmt.Println(u.Id)
	fmt.Println(u.Nickname)
	fmt.Println(u.Name)
	fmt.Println(u.Surname)

	// Output:
	// 1
	// michkenntjeder
	// Max
	// Mustermann
}

func ExampleUser_AddContact() {
	u := New("1", "michkenntjeder", "Max", "Mustermann")

	u.AddContact("2")
	u.AddContact("3")

	fmt.Println(u.Contacts)

	// Output:
	// [2 3]
}
