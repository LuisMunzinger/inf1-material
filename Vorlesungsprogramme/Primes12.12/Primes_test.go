package primes

import "fmt"

func ExampleIsPrime() {
	fmt.Println(IsPrime(1))
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(3))
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(6))
	fmt.Println(IsPrime(7))
	fmt.Println(IsPrime(8))
	fmt.Println(IsPrime(9))

	// Output:
	// false
	// true
	// true
	// false
	// true
	// false
	// true
}

func ExamplePrimes() {
	p := Primes(10)

	fmt.Println(p)
	fmt.Println(p[0])
	fmt.Println(p[3])
	fmt.Println(p[1:3]) //Von stelle 1 bis ausschlieslich stelle 3 aus der liste

	// Output:
	// [2 3 5 7]
}
