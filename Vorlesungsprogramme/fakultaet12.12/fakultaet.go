package main

func Factorial(n int) int {
	result := 1

	for i := n; i > 1; i-- {
		result = result * i
	}
	return result
}
