package Datentypen

import "fmt"

type Length int

func (l Length) Centimeters() int {
	return int(l)
}

func (l Length) Meters() int {
	return l.Centimeters() / 100
}

func (l *Length) Scale(factor int) {
	*l = Length(l.Centimeters() * factor)
}

func Example() {
	var a Length = 1000000
	var b int = 2

	fmt.Println(a)
	a.Scale(b)

	println(a.Centimeters())
	println(a.Meters())

	//Output:
	// 1000000

}
