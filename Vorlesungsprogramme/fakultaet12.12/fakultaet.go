package main

import "fmt"

func facultaet(n int) int {
	result := 1
	for i := n; i > 1; i-- {
		result = result * i
	}
	return result
}
func main() {
	fmt.Println(facultaet(0))
	fmt.Println(facultaet(1))
	fmt.Println(facultaet(2))
	fmt.Println(facultaet(3))
	fmt.Println(facultaet(4))
	fmt.Println(facultaet(5))
	fmt.Println(facultaet(6))
}
