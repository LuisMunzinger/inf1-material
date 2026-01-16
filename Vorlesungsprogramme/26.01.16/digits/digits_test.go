package digits

import (
	"fmt"
	"slices"
)

// BinDigits erwartete eine Zahl n und liefert eine Liste von Ziffern
func Digits(n, base int) []int {
	result := []int{}

	for n != 0 {
		last_digit := n % base
		result = append(result, last_digit)
		n /= base
	}
	slices.Reverse(result)

	return result
}

// Sum erwartet eine Liste von Zahlen und berechnet deren Summe.
func Sum(number []int) int {
	result := 0
	for _, n := range number {
		result += n
	}
	return result
}

// Parity erwartet eine Zahl n und liefert:
//
// 1: Falls die Anazhl der Einsen in der Binärdarstellung von n ungerade ist.
// 2: Falls die Anazhl der Einsen in der Binärdarstellung von n ungerade ist.
func ParityBit(n int) int {
	return Sum(Digits(n, 2)) % 2
}

func DigitSum(n, base int) int {
	return Sum(Digits(n, base))
}

func ExampleDigits() {
	fmt.Println(Digits(42, 2))
	fmt.Println(Digits(42, 16))
	fmt.Println(Digits(42, 10))
	fmt.Println(Digits(42, 8))

	fmt.Println(ParityBit(42))

	// Output:
	//[ 1 0 1 0 1 0 ]
	//[ 2 10 ]
	//[ 4 2 ]
	//[ 5 2 ]
	// 1
}
