package forms

import "fmt"

func ExampleAreaRectangle() {
	fmt.Println(AreaRectangle(3, 4))
	fmt.Println(AreaRectangle(2.5, 10))
	fmt.Println(AreaRectangle(1.5, 25))

	// Output:
	// 12
	// 25
	// 37.5
}
