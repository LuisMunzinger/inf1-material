package main

import "fmt"

//BinDigits erwartete eine Zahl n und liefert eine Liste von Ziffern
func BinDigits(n int) []int {
	result := []int{}

	for n != 0 {
		last_digit := n % 2
		result = append(result, last_digit)
		//result = append([]int{last_digit}, result...) //Direkt umgekert anhängen
		n /= 2
	}

	return result
}

func ExampleBinDigits() {
	fmt.Println(BinDigits(42))

	// Output:
	//[ 1 0 1 0 1 0 ]
}
