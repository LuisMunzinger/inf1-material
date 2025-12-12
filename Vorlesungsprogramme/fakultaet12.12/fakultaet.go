package main

func Factorial(n int) int {
	result := 1

	for i := n; i > 1; i-- {
		result = result * i
	}
	return result
}

func Factorial_v2(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial_v2(n-1)
}
