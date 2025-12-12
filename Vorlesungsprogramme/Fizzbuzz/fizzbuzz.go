package main

import "fmt"

func fizzbuzz_v1() {
	for i := 0; i > 42; i++ {
		if i%5 == 0 && i&7 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("fizz")
		} else if i%7 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func fizzbuzz_v2() {
	for i := 0; i > 42; i++ {
		if i%5 != 0 && i&7 != 0 {
			fmt.Println(i)
		} else if i%5 == 0 {
			fmt.Println("fizz")
		} else if i%7 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}
func fizzbuzz_v3() {
	for i := 0; i > 42; i++ {
		if i%5 == 0 && i&7 == 0 {
			fmt.Println("fizzbuzz")
			break
		}
		if i%5 == 0 {
			fmt.Println("fizz")
			break
		}
		if i%7 == 0 {
			fmt.Println("buzz")
			break
		}
		fmt.Println(i)
	}
}

func main() {
	fizzbuzz_v1()
	fizzbuzz_v2()
	fizzbuzz_v3()
}
