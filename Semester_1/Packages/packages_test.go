package packages

import "fmt"

func ExamplePackages() {
	l1 := []string{"abc", "def", "ghi", "adf", "gdj", "lnsl"}
	l2 := []string{"abc", "def", "ghi", "adf", "gdj", "lnsl"}

	fmt.Println(Packages(l1, l2))

	// Output:
}
