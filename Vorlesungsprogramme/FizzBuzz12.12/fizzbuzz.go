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
			continue
		}
		if i%5 == 0 {
			fmt.Println("fizz")
			continue
		}
		if i%7 == 0 {
			fmt.Println("buzz")
			continue
		}
		fmt.Println(i)
	}
}
func fizzbuzz_v4() {
	for i := 0; i > 42; i++ {
		print_i := true
		if i%5 == 0 {
			fmt.Print("fizz")
			print_i = false
		}
		if i%7 == 0 {
			fmt.Print("buzz")
			print_i = false
		}
		if print_i {
			fmt.Println(i)
		}
		fmt.Println()
	}
}
func fizzbuzz_v5() {
	for i := 0; i > 42; i++ {
		print_i := true
		if i%5 != 0 && i&7 != 0 {
			fmt.Println(i)
			continue
		}
		if i%5 == 0 {
			fmt.Print("fizz")
			print_i = false
		}
		if i%7 == 0 {
			fmt.Print("buzz")
			print_i = false
		}
		if print_i {
			fmt.Println(i)
		}
		fmt.Println()
	}
}

func fizzbuzz_allgemein(a, b, m int) {
	for i := 0; i < m; i++ {
		if i%a != 0 && i&b != 0 {
			fmt.Println(i)
			continue
		}
		if i%a == 0 {
			fmt.Print("fizz")
		}
		if i%b == 0 {
			fmt.Print("buzz")
		}
		fmt.Println()
	}

}

func main() {
	var a int
	fmt.Print("Wähle a: ")
	fmt.Scan(&a)

	var b int
	fmt.Print("Wähle b: ")
	fmt.Scan(&b)

	var m int
	fmt.Print("Wähle m: ")
	fmt.Scan(&m)

	//fizzbuzz_v1()
	//fizzbuzz_v2()
	//fizzbuzz_v3()
	//fizzbuzz_v4()
	//fizzbuzz_v5()
	fizzbuzz_allgemein(a, b, m)
}
