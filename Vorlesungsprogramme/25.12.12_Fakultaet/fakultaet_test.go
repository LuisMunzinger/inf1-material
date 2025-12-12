package main

import "fmt"

func ExampleFactorial() {
	fmt.Println(Factorial(0))
	fmt.Println(Factorial(1))
	fmt.Println(Factorial(2))
	fmt.Println(Factorial(3))
	fmt.Println(Factorial(4))
	fmt.Println(Factorial(5))

	// Output:
	// 1
	// 1
	// 2
	// 6
	// 24
	// 120
}
func ExampleFactorial_v2() {
	fmt.Println(Factorial_v2(0))
	fmt.Println(Factorial_v2(1))
	fmt.Println(Factorial_v2(2))
	fmt.Println(Factorial_v2(3))
	fmt.Println(Factorial_v2(4))
	fmt.Println(Factorial_v2(5))

	// Output:
	// 1
	// 1
	// 2
	// 6
	// 24
	// 120
}
